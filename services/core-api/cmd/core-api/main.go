package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/phatnguyen03022001/ilets/services/core-api/internal/db"
	"github.com/phatnguyen03022001/ilets/services/core-api/internal/httpapi"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	pool, err := db.Open(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	cfg := httpapi.Config{
		Environment:  getenv("ILETS_ENV", "development"),
		WebOrigins:   splitCSV(getenv("WEB_ORIGINS", "http://127.0.0.1:3000,http://localhost:3000")),
		BuildVersion: getenv("BUILD_VERSION", "dev"),
	}
	handler := httpapi.New(pool, cfg, slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	server := &http.Server{Addr: getenv("CORE_ADDR", "127.0.0.1:8080"), Handler: handler, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 15 * time.Second, WriteTimeout: 15 * time.Second, IdleTimeout: 60 * time.Second}
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdownCtx)
	}()
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func getenv(k, fallback string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return fallback
}
func splitCSV(v string) []string {
	var out []string
	for _, s := range strings.Split(v, ",") {
		if s = strings.TrimSpace(s); s != "" {
			out = append(out, s)
		}
	}
	return out
}
