#!/bin/sh
set -eu

ROOT="$(CDPATH= cd -- "$(dirname -- "$0")/../.." && pwd)"
PUBLIC="$ROOT/contracts/http/public.openapi.yaml"
EVALUATOR="$ROOT/contracts/http/evaluator.openapi.yaml"
PUBLIC_GO_DIR="$ROOT/services/core-api/internal/generated/openapi/public"
EVALUATOR_GO_DIR="$ROOT/services/core-api/internal/generated/openapi/evaluator"
PUBLIC_TS_DIR="$ROOT/apps/web/src/generated/public"

rm -rf "$PUBLIC_GO_DIR" "$EVALUATOR_GO_DIR" "$PUBLIC_TS_DIR"
mkdir -p "$PUBLIC_GO_DIR" "$EVALUATOR_GO_DIR"

cd "$ROOT/services/core-api"
go tool oapi-codegen -config ../../tools/contracts/oapi-public.yaml "$PUBLIC"
go tool oapi-codegen -config ../../tools/contracts/oapi-evaluator.yaml "$EVALUATOR"

cd "$ROOT/apps/web"
corepack pnpm exec openapi-ts   --input "$PUBLIC"   --output "$PUBLIC_TS_DIR"   --client @hey-api/client-fetch   --no-log-file   --silent
