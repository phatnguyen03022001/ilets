#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$ROOT/services/core-api"
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.8.0 \
  -generate types,chi-server,spec \
  -package generated \
  -o internal/httpapi/generated/public_v1.gen.go \
  ../../contracts/http/public-v1.json
cd "$ROOT/apps/web"
pnpm contract:generate
