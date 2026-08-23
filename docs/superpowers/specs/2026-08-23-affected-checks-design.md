# Affected Checks Fast Path Design

## Status

Approved implementation design for a non-authoritative fast feedback path. The repository-native root `./verify` remains the only authoritative PASS definition.

## Invariants

- `./verify` remains unchanged in meaning and always runs the full repository correctness procedure.
- `run_verify` remains full verification; this task does not alter IELTS Runner or the tunnel.
- `tools/check-affected` may return only `CHECK_PASS` or `CHECK_FAIL` status labels. It must never emit `VERIFY_PASS`.
- Any uncertainty in baseline resolution, diff parsing, classification, path ownership, submodule handling, working-tree completeness, or critical-path impact falls back to full `./verify`.
- v1 is sequential. Determinism is preferred over speculative parallelism.

## Baseline and changed-file discovery

Baseline precedence is exact:

1. `--base <ref>`
2. `AFFECTED_BASE`
3. no trusted baseline → full `./verify`

There is no `origin/main` default. A missing baseline must never produce a no-change `CHECK_PASS`.

The fast path resolves the supplied ref, computes `git merge-base <base> HEAD`, and obtains changed paths with `git diff --name-status -z --find-renames --find-copies <merge-base> HEAD --`.

Parsing is NUL-delimited. Rename/copy records classify both old and new paths. Malformed records, unsupported statuses, undecodable paths, merge-base failure, unresolved shallow history, or other Git anomalies trigger full fallback.

V1 classifies committed `HEAD` history only. Before classification it requires a clean Git working tree, including no untracked files. A dirty worktree would make the committed baseline diff incomplete and therefore triggers full fallback rather than permitting a false no-change/partial `CHECK_PASS`.

Changed paths are checked for Git mode `160000` in the base tree and current index/tree. Any submodule involvement triggers full fallback.

`tools/check-affected` disables Python bytecode writes before importing its classifier so invoking the fast path itself does not dirty the checkout with `__pycache__` artifacts.

## Classification

Classification is fail-closed and path based. A change set may accumulate non-critical lanes; any critical or unknown path overrides the result to full verification.

### Docs-only

The docs-only whitelist is intentionally narrow:

- `README.md`
- `docs/**/*.md`

No broad `research/**`, `evidence/**`, or `archive/**` exemption exists in v1. Unrecognized documentation/support paths therefore fall back to full verification.

Docs-only runs no runtime lane and may return `CHECK_PASS` after classification succeeds.

### Web lane

Ordinary `apps/web/**` source changes run the complete Web lane:

1. repository Node version check
2. `corepack enable`
3. `pnpm install --frozen-lockfile`
4. `pnpm format:check`
5. `pnpm lint`
6. `pnpm typecheck`
7. `pnpm test` (all Vitest tests)
8. `pnpm build`

No filename-based Vitest selection is used.

### Go lane

Ordinary non-DB `services/core-api/**` implementation changes run the complete Go lane:

1. `gofmt -l` cleanliness check
2. `go vet ./...`
3. `go test ./...`
4. `go build` for `./cmd/core-api` into a temporary directory

No package/test narrowing is used in v1. Files ending in `_integration_test.go` force full fallback because the ordinary Go lane does not start PostgreSQL or set `ILETS_INTEGRATION=1`; classifying such a test change as ordinary Go could skip the changed test body.

### Mixed Web + Go

If a change set contains only recognized non-critical Web and Go paths, run the full Web lane followed by the full Go lane, sequentially.

## Full-fallback paths

At minimum, these paths/classes force full `./verify`:

- root `verify`
- `tools/verify-local` and `tools/local-verify/**`
- `compose*.yml`
- `.github/workflows/**`
- `.node-version`, `.go-version`
- package-manager/module/lock/toolchain files including Web `package.json`, `pnpm-lock.yaml`, `pnpm-workspace.yaml`, Go `go.mod`, `go.sum`, and relevant Web tool configs
- `contracts/**` and `tools/contracts/**`
- generated contract bindings
- DB migrations/schema/query/sqlc/generation paths, seed/bootstrap data, and `tools/db/**`
- Go `*_integration_test.go` files
- `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, `design/**`
- canonical materialization/registry sources and tooling under `tools/canonical/**` and `tools/slice/**`
- founder-closure semantic inputs/verifiers
- `tools/check-affected` and `tools/affected/**`
- Web E2E / Playwright paths
- submodules
- dirty working tree or untracked files
- any unknown or unclassified path

## Full fallback execution

Fallback directly invokes root `./verify`. A successful fallback is reported as `CHECK_PASS mode=full-fallback`; a failure is `CHECK_FAIL mode=full-fallback`. These labels describe the fast-path command result and do not create a second authoritative repository PASS definition.

## Testing

`tools/affected/test_classifier.py` is called from root `./verify` so classifier correctness is part of the existing authoritative procedure.

Required classifier coverage includes:

- docs-only
- Web-only
- Go-only
- DB/contract critical → full
- Go integration test → full
- unknown → full
- mixed Web + Go
- rename/copy old + new path handling
- malformed NUL parser input
- classifier/self-change → full
- dirty worktree → full fallback
- submodule → full fallback
- unresolved shallow baseline → full fallback
- direct invocation does not create Python bytecode in the checkout

Implementation follows RED → GREEN. Final acceptance requires a fresh full root `./verify` through the existing local OrbStack runner.

## Concurrency

No new parallel lanes are introduced in v1. Existing `tools/verify-local` concurrent-run protection is left unchanged. The fast path does not modify runner Compose, ports, tunnel behavior, or the stable OrbStack project design.
