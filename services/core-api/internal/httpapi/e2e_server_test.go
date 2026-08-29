package httpapi

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"testing"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/phatnguyen03022001/ilets/services/core-api/internal/db"
)

func TestPlaywrightServer(t *testing.T) {
	if os.Getenv("ILETS_E2E_SERVER") != "1" {
		t.Skip("Playwright-only Core server")
	}

	pool, err := db.Open(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	if _, err := pool.Exec(context.Background(), `TRUNCATE idempotency_operations, observations, attempts, practice_activities, target_profiles, external_principals, learners CASCADE`); err != nil {
		t.Fatal(err)
	}

	publicKey := os.Getenv("ILETS_E2E_PUBLIC_KEY_PEM")
	if publicKey == "" {
		t.Fatal("ILETS_E2E_PUBLIC_KEY_PEM is required")
	}
	cfg := Config{
		Environment:            "test",
		WebOrigins:             splitNonEmptyCSV(os.Getenv("WEB_ORIGINS")),
		BuildVersion:           "playwright",
		ClerkIssuer:            os.Getenv("CLERK_ISSUER"),
		ClerkAudience:          os.Getenv("CLERK_AUDIENCE"),
		ClerkAuthorizedParties: splitNonEmptyCSV(os.Getenv("CLERK_AUTHORIZED_PARTIES")),
	}
	handler := newWithClerkAuthorizationOptions(
		pool,
		cfg,
		slog.New(slog.NewJSONHandler(io.Discard, nil)),
		clerkhttp.JSONWebKey(publicKey),
		clerkhttp.Leeway(0),
	)

	server := &http.Server{Addr: os.Getenv("CORE_ADDR"), Handler: handler}
	if server.Addr == "" {
		t.Fatal("CORE_ADDR is required")
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		t.Fatal(err)
	}
}

func splitNonEmptyCSV(value string) []string {
	var out []string
	for _, item := range strings.Split(value, ",") {
		if item = strings.TrimSpace(item); item != "" {
			out = append(out, item)
		}
	}
	return out
}
