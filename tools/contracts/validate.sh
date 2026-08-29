#!/bin/sh
set -eu

ROOT="$(CDPATH= cd -- "$(dirname -- "$0")/../.." && pwd)"
PUBLIC="$ROOT/contracts/http/public.openapi.yaml"
EVALUATOR="$ROOT/contracts/http/evaluator.openapi.yaml"
LEGACY="$ROOT/contracts/http/public-v1.json"

if [ -e "$LEGACY" ]; then
  echo "legacy bootstrap contract must not remain an active authority: $LEGACY" >&2
  exit 1
fi

unformatted="$(gofmt -l "$ROOT/tools/contracts/validate.go")"
if [ -n "$unformatted" ]; then
  echo "contract validator needs gofmt: $unformatted" >&2
  exit 1
fi

cd "$ROOT/services/core-api"
go run ../../tools/contracts/validate.go --repository "$PUBLIC" "$EVALUATOR"

cd "$ROOT"
python3 tools/contracts/test_validate.py
