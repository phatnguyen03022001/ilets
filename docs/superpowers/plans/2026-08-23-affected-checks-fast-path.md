# Affected Checks Fast Path Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a fail-closed affected-check fast path without weakening the authoritative root `./verify` procedure.

**Architecture:** `tools/check-affected` is a small Python entrypoint that resolves an explicit baseline, reads NUL-delimited Git name-status output, detects submodule/parse anomalies, delegates pure path classification to `tools/affected/classifier.py`, and then runs either docs-only/no-change, the full Web lane, the full Go lane, both lanes sequentially, or full `./verify`. Classifier tests are standard-library `unittest` and are invoked by root `./verify`.

**Tech Stack:** Python 3 standard library, Git CLI, existing pnpm/Web scripts, existing Go toolchain, existing root `./verify`.

**Spec:** `docs/superpowers/specs/2026-08-23-affected-checks-design.md`

## Global Constraints

- Root `./verify` remains the only authoritative PASS definition.
- `run_verify` remains full verification.
- `tools/check-affected` emits only `CHECK_PASS` / `CHECK_FAIL` status labels, never `VERIFY_PASS`.
- Baseline precedence: `--base` → `AFFECTED_BASE` → full fallback; no `origin/main` default.
- Changed files use `git diff --name-status -z`; rename/copy classify old and new paths.
- Any baseline, merge-base, parser, shallow-history resolution, submodule, critical-path, or unknown-path uncertainty falls back to full `./verify`.
- v1 is sequential; no parallel execution.
- Web runs the complete Web lane; Go runs the complete non-integration Go lane.
- No broad `research/**` docs-only whitelist.
- Do not modify IELTS Runner, local tunnel, or stable OrbStack runner design.

---

### Task 1: Pure classifier and NUL parser

**Files:**
- Create: `tools/affected/classifier.py`
- Create: `tools/affected/test_classifier.py`

**Interfaces:**
- Produces: `Change`, `Classification`, `ParseError`, `parse_name_status_z(data: bytes)`, `classify_changes(changes: Sequence[Change])`.
- Classification modes: `no-changes`, `docs`, `web`, `go`, `web+go`, `full`.

- [ ] **Step 1: Write failing classifier tests**

Cover docs-only, Web-only, Go-only, DB critical, contract critical, unknown, mixed Web+Go, rename/copy both-path behavior, malformed NUL input, unsupported status, and classifier/self-change.

- [ ] **Step 2: Run test to verify RED**

Run: `python3 tools/affected/test_classifier.py`

Expected: failure because `tools/affected/classifier.py` does not exist or required interfaces are absent.

- [ ] **Step 3: Implement minimal parser/classifier**

Use strict UTF-8 decoding, strict NUL record arity, explicit status handling, path validation, narrow docs whitelist, explicit critical rules, and unknown→full behavior.

- [ ] **Step 4: Run test to verify GREEN**

Run: `python3 tools/affected/test_classifier.py`

Expected: all tests pass.

### Task 2: Fast-path entrypoint

**Files:**
- Create: `tools/check-affected`
- Test: `tools/affected/test_classifier.py`

**Interfaces:**
- Consumes: classifier interfaces from Task 1.
- CLI: `tools/check-affected [--base <ref>]`; environment fallback `AFFECTED_BASE`.

- [ ] **Step 1: Add failing tests for helper behavior that can remain pure**

Add coverage for classification precedence assumptions needed by the entrypoint, especially critical override and rename/copy path union.

- [ ] **Step 2: Run test to verify RED**

Run: `python3 tools/affected/test_classifier.py`

Expected: new assertion fails until supporting classifier behavior exists.

- [ ] **Step 3: Implement entrypoint**

Implement explicit baseline resolution, `git rev-parse`, `git merge-base`, `git diff --name-status -z --find-renames --find-copies`, submodule-mode detection (`160000`) in base/current state, sequential Web/Go lanes, and full fallback. Every unexpected exception routes to full fallback rather than a partial PASS.

- [ ] **Step 4: Exercise CLI against synthetic Git repositories**

Use temporary Git repositories to verify: missing baseline→full fallback dispatch, explicit no-change baseline→`CHECK_PASS`, docs-only→`CHECK_PASS`, rename to critical→full classification, and malformed parser unit behavior.

### Task 3: Root verification ownership

**Files:**
- Modify: `verify`

**Interfaces:**
- Root verify invokes `python3 "$ROOT/tools/affected/test_classifier.py"` before existing repository checks.

- [ ] **Step 1: Add classifier tests to root verify**

Do not alter or skip any existing root verification step.

- [ ] **Step 2: Run focused test**

Run: `python3 tools/affected/test_classifier.py`

Expected: pass.

- [ ] **Step 3: Review root verify diff**

Confirm the only semantic change is adding classifier-test coverage; full Web, codegen, Go, DB, integration, Playwright, build, and clean-tree steps remain intact.

### Task 4: Final verification and integration

**Files:**
- Review all files from Tasks 1–3.

- [ ] **Step 1: Run classifier suite fresh**

Run: `python3 tools/affected/test_classifier.py`

Expected: pass with zero failures.

- [ ] **Step 2: Fast-forward GitHub main only after branch code is green**

Keep intermediate RED work off `main`.

- [ ] **Step 3: Sync existing local runner to GitHub main**

Use existing IELTS Runner `sync_main`; do not modify tunnel or runner implementation.

- [ ] **Step 4: Run authoritative full verification**

Use existing IELTS Runner `run_verify`, which invokes full root `./verify`.

Expected: exit code 0 for the exact final GitHub `main` SHA.

- [ ] **Step 5: Inspect final verify log and GitHub SHA**

Record exact final SHA, files changed, classification/lane matrix, fallback rules, RED/GREEN evidence, full verify evidence, and remaining risks.
