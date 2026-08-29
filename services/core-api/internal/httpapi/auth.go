package httpapi

import (
	"context"
	"fmt"
	"net/http"

	clerk "github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/clerk/clerk-sdk-go/v2/jwks"
)

type clerkAuthConfig struct {
	Issuer            string
	Audience          string
	AuthorizedParties []string
}

type externalPrincipalKey struct{}

type externalPrincipalIdentity struct {
	Provider string
	Issuer   string
	Subject  string
}

func newClerkAuthMiddleware(cfg clerkAuthConfig, extra ...clerkhttp.AuthorizationOption) func(http.Handler) http.Handler {
	parties := make(map[string]struct{}, len(cfg.AuthorizedParties))
	for _, party := range cfg.AuthorizedParties {
		if party != "" {
			parties[party] = struct{}{}
		}
	}
	opts := []clerkhttp.AuthorizationOption{
		clerkhttp.AuthorizedParty(func(azp string) bool {
			if azp == "" || len(parties) == 0 {
				return false
			}
			_, ok := parties[azp]
			return ok
		}),
	}
	opts = append(opts, extra...)
	verifyHeader := clerkhttp.WithHeaderAuthorization(opts...)

	return func(next http.Handler) http.Handler {
		exactClaims := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := clerk.SessionClaimsFromContext(r.Context())
			if !ok || claims == nil || claims.Subject == "" || claims.SessionID == "" || claims.Expiry == nil || claims.NotBefore == nil || claims.IssuedAt == nil || claims.Issuer != cfg.Issuer || len(claims.Audience) != 1 || claims.Audience[0] != cfg.Audience {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), externalPrincipalKey{}, externalPrincipalIdentity{
				Provider: "clerk",
				Issuer:   claims.Issuer,
				Subject:  claims.Subject,
			})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
		checked := verifyHeader(exactClaims)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Authorization") == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			checked.ServeHTTP(w, r)
		})
	}
}

func authenticatedExternalPrincipal(ctx context.Context) (externalPrincipalIdentity, bool) {
	value, ok := ctx.Value(externalPrincipalKey{}).(externalPrincipalIdentity)
	return value, ok && value.Provider != "" && value.Issuer != "" && value.Subject != ""
}

func productionClerkAuthorizationOptions(cfg Config) ([]clerkhttp.AuthorizationOption, error) {
	if cfg.ClerkIssuer == "" || cfg.ClerkAudience == "" || len(cfg.ClerkAuthorizedParties) == 0 {
		return nil, fmt.Errorf("CLERK_ISSUER, CLERK_AUDIENCE, and CLERK_AUTHORIZED_PARTIES are required")
	}
	if cfg.ClerkSecretKey == "" {
		return nil, fmt.Errorf("CLERK_SECRET_KEY is required")
	}
	clientConfig := &clerk.ClientConfig{
		BackendConfig: clerk.BackendConfig{
			Key: clerk.String(cfg.ClerkSecretKey),
		},
	}
	return []clerkhttp.AuthorizationOption{clerkhttp.JWKSClient(jwks.NewClient(clientConfig))}, nil
}
