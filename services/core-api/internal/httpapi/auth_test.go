package httpapi

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	clerk "github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/go-jose/go-jose/v3"
	josejwt "github.com/go-jose/go-jose/v3/jwt"
)

const (
	testClerkIssuer = "https://test-instance.clerk.accounts.dev"
	testAudience    = "ilets-core"
	testAzp         = "http://127.0.0.1:3000"
)

type fixedClock struct{ now time.Time }

func (c fixedClock) Now() time.Time { return c.now }

func TestClerkBearerValidation(t *testing.T) {
	now := time.Date(2026, 8, 29, 10, 0, 0, 0, time.UTC)
	goodKey := mustRSAKey(t)
	badKey := mustRSAKey(t)
	publicPEM := mustPublicPEM(t, &goodKey.PublicKey)

	cfg := clerkAuthConfig{
		Issuer:            testClerkIssuer,
		Audience:          testAudience,
		AuthorizedParties: []string{testAzp},
	}
	middleware := newClerkAuthMiddleware(cfg,
		clerkhttp.JSONWebKey(publicPEM),
		clerkhttp.Clock(fixedClock{now: now}),
		clerkhttp.Leeway(0),
	)

	valid := tokenClaims{Issuer: testClerkIssuer, Subject: "user_alpha", Audience: []string{testAudience}, AuthorizedParty: testAzp, ExpiresAt: now.Add(time.Minute), NotBefore: now.Add(-time.Minute), IssuedAt: now.Add(-time.Minute)}
	cases := []struct {
		name   string
		claims tokenClaims
		key    *rsa.PrivateKey
		want   int
	}{
		{"valid signed token", valid, goodKey, http.StatusNoContent},
		{"invalid signature", valid, badKey, http.StatusUnauthorized},
		{"wrong issuer", withClaims(valid, func(c *tokenClaims) { c.Issuer = "https://other.clerk.accounts.dev" }), goodKey, http.StatusUnauthorized},
		{"missing audience", withClaims(valid, func(c *tokenClaims) { c.Audience = nil }), goodKey, http.StatusUnauthorized},
		{"wrong audience", withClaims(valid, func(c *tokenClaims) { c.Audience = []string{"other-api"} }), goodKey, http.StatusUnauthorized},
		{"extra audience", withClaims(valid, func(c *tokenClaims) { c.Audience = []string{testAudience, "other-api"} }), goodKey, http.StatusUnauthorized},
		{"unauthorized azp", withClaims(valid, func(c *tokenClaims) { c.AuthorizedParty = "https://evil.example" }), goodKey, http.StatusUnauthorized},
		{"expired", withClaims(valid, func(c *tokenClaims) { c.ExpiresAt = now.Add(-time.Second) }), goodKey, http.StatusUnauthorized},
		{"nbf violation", withClaims(valid, func(c *tokenClaims) { c.NotBefore = now.Add(time.Minute) }), goodKey, http.StatusUnauthorized},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			token := mustSignToken(t, tc.key, tc.claims)
			req := httptest.NewRequest(http.MethodGet, "/v1/me", nil)
			req.Header.Set("Authorization", "Bearer "+token)
			rr := httptest.NewRecorder()
			middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				principal, ok := authenticatedExternalPrincipal(r.Context())
				if !ok || principal.Provider != "clerk" || principal.Issuer != testClerkIssuer || principal.Subject != "user_alpha" {
					t.Fatalf("verified principal = %#v, %v", principal, ok)
				}
				w.WriteHeader(http.StatusNoContent)
			})).ServeHTTP(rr, req)
			if rr.Code != tc.want {
				t.Fatalf("status=%d body=%s want=%d", rr.Code, rr.Body.String(), tc.want)
			}
		})
	}
}

type tokenClaims struct {
	Issuer          string
	Subject         string
	Audience        []string
	AuthorizedParty string
	ExpiresAt       time.Time
	NotBefore       time.Time
	IssuedAt        time.Time
}

func withClaims(base tokenClaims, change func(*tokenClaims)) tokenClaims { change(&base); return base }
func mustRSAKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	return k
}
func mustPublicPEM(t *testing.T, key *rsa.PublicKey) string {
	t.Helper()
	der, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: der}))
}
func mustSignToken(t *testing.T, key *rsa.PrivateKey, c tokenClaims) string {
	t.Helper()
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: key}, (&jose.SignerOptions{}).WithType("JWT").WithHeader("kid", "test-key"))
	if err != nil {
		t.Fatal(err)
	}
	claims := josejwt.Claims{Issuer: c.Issuer, Subject: c.Subject, Audience: josejwt.Audience(c.Audience), Expiry: josejwt.NewNumericDate(c.ExpiresAt), NotBefore: josejwt.NewNumericDate(c.NotBefore), IssuedAt: josejwt.NewNumericDate(c.IssuedAt)}
	raw, err := josejwt.Signed(signer).Claims(claims).Claims(map[string]any{"azp": c.AuthorizedParty, "sid": "sess_test"}).CompactSerialize()
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

var _ clerk.Clock = fixedClock{}
var _ = context.Background
