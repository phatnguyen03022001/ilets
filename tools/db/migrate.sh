#!/usr/bin/env bash
set -euo pipefail

: "${DATABASE_URL:?DATABASE_URL is required}"
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT/services/core-api"
exec go run github.com/jackc/tern/v2@v2.4.2 migrate --migrations "$ROOT/services/core-api/migrations" --conn-string "$DATABASE_URL"
