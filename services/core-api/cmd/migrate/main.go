package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/phatnguyen03022001/ilets/services/core-api/internal/db"
)

func main() {
	ctx := context.Background()
	pool, err := db.Open(ctx)
	if err != nil { log.Fatal(err) }
	defer pool.Close()
	if _, err := pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (name text PRIMARY KEY, applied_at timestamptz NOT NULL DEFAULT now())`); err != nil { log.Fatal(err) }
	dir := os.Getenv("MIGRATIONS_DIR")
	if dir == "" { dir = "migrations" }
	entries, err := os.ReadDir(dir)
	if err != nil { log.Fatal(err) }
	var names []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") { names = append(names, entry.Name()) }
	}
	sort.Strings(names)
	for _, name := range names {
		var exists bool
		if err := pool.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE name=$1)`, name).Scan(&exists); err != nil { log.Fatal(err) }
		if exists { continue }
		body, err := os.ReadFile(filepath.Join(dir, name)); if err != nil { log.Fatal(err) }
		tx, err := pool.Begin(ctx); if err != nil { log.Fatal(err) }
		if _, err = tx.Exec(ctx, string(body)); err == nil {
			_, err = tx.Exec(ctx, `INSERT INTO schema_migrations(name) VALUES($1)`, name)
		}
		if err != nil { _ = tx.Rollback(ctx); log.Fatalf("migration %s: %v", name, err) }
		if err := tx.Commit(ctx); err != nil { log.Fatal(err) }
		fmt.Println("applied", name)
	}
}
