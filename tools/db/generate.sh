#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT/services/core-api"
exec go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.31.1 generate
