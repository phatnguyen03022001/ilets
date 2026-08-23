# Persistent OrbStack IELTS Runner Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Keep one stable stopped `ielts` Compose project visible in OrbStack after local verification while preserving repository-native `./verify` as the only PASS definition.

**Architecture:** `tools/verify-local` continues to build the pinned `ielts:latest` toolchain image, but launches verification through a dedicated Compose project named `ielts` instead of `docker run --rm`. The stable `runner` service bind-mounts the invoking checkout and Docker socket, runs `./verify`, retains its stopped container for OrbStack visibility, and delegates PostgreSQL to the separate disposable `ielts-verify` project owned by root `./verify`.

**Tech Stack:** Bash, Docker/Compose, OrbStack, repository-native `./verify`.

**Spec:** `docs/superpowers/specs/2026-08-23-orbstack-persistent-ielts-runner-design.md`

## Global Constraints

- GitHub `main` remains the sole source of truth.
- Root `./verify` must not change.
- A local verification PASS remains exactly exit code 0 from root `./verify`.
- Repository source remains bind-mounted and is never copied into the runner image.
- Runner image remains `ielts:latest` with the existing pinned toolchain.
- Stable runner Compose project is `ielts`; verifier-owned PostgreSQL Compose project defaults to `ielts-verify`.
- Existing `ILETS_*` environment variable contracts remain unchanged.
- pnpm store must remain outside the checkout.
- Do not expose arbitrary shell, Git write authority, or Docker socket through the ChatGPT plugin surface.

---

### Task 1: Define the persistent runner Compose service

**Files:**
- Create: `tools/local-verify/compose.runner.yml`
- Verify: temporary local/static contract check plus OrbStack validation in Task 3

**Interfaces:**
- Consumes: `ielts:latest`, `IELTS_RUNNER_ROOT`, `IELTS_RUNNER_DOCKER_SOCKET`, existing `ILETS_*` overrides.
- Produces: Compose service `runner` under project `ielts`; command `./verify`; stopped container retained after exit.

- [ ] **Step 1: Write the failing contract check**

Use a temporary Python assertion that the Compose file exists and specifies: `runner`, `image: ielts:latest`, `network_mode: host`, checkout bind mount to `/workspace`, Docker socket bind mount, `pnpm_config_store_dir=/tmp/ielts-pnpm-store`, `ILETS_VERIFY_COMPOSE_PROJECT`, and command `./verify`. Run it before creating the file and confirm failure because the file is absent.

- [ ] **Step 2: Create the minimal Compose file**

```yaml
services:
  runner:
    image: ielts:latest
    init: true
    network_mode: host
    shm_size: 1g
    working_dir: /workspace
    command: ["./verify"]
    volumes:
      - type: bind
        source: ${IELTS_RUNNER_ROOT}
        target: /workspace
      - type: bind
        source: ${IELTS_RUNNER_DOCKER_SOCKET}
        target: /var/run/docker.sock
    environment:
      DOCKER_HOST: unix:///var/run/docker.sock
      pnpm_config_store_dir: /tmp/ielts-pnpm-store
      ILETS_VERIFY_COMPOSE_PROJECT: ${ILETS_VERIFY_COMPOSE_PROJECT:-ielts-verify}
      ILETS_VERIFY_POSTGRES_PORT: ${ILETS_VERIFY_POSTGRES_PORT:-}
      ILETS_E2E_CORE_PORT: ${ILETS_E2E_CORE_PORT:-}
      ILETS_E2E_WEB_PORT: ${ILETS_E2E_WEB_PORT:-}
```

- [ ] **Step 3: Run the contract check again**

Expected: PASS for all required runner properties.

### Task 2: Route `tools/verify-local` through the stable Compose project

**Files:**
- Modify: `tools/verify-local`
- Modify: `tools/local-verify/README.md`

**Interfaces:**
- Consumes: active Docker Unix socket and invoking checkout path.
- Produces: image build plus `docker compose -p ielts -f tools/local-verify/compose.runner.yml up ... runner`, with runner exit code propagated.

- [ ] **Step 1: Write the failing wrapper contract check**

Use a temporary static check requiring the wrapper to export `IELTS_RUNNER_ROOT`, `IELTS_RUNNER_DOCKER_SOCKET`, default `ILETS_VERIFY_COMPOSE_PROJECT` to `ielts-verify`, define the runner Compose command with project `ielts`, reject a concurrently running `runner`, remove an old stopped runner, and use `up --no-deps --abort-on-container-exit --exit-code-from runner runner`. Confirm the current wrapper fails these assertions because it still uses `docker run`.

- [ ] **Step 2: Implement minimal wrapper changes**

Keep the existing version resolution, Docker socket discovery, and image build. Replace the one-shot `docker run --rm` block with exports for the Compose interpolation variables and a fixed Compose command:

```bash
IMAGE_TAG="ielts:latest"
RUNNER_PROJECT="ielts"
RUNNER_COMPOSE_FILE="$ROOT/tools/local-verify/compose.runner.yml"
VERIFY_COMPOSE_PROJECT="${ILETS_VERIFY_COMPOSE_PROJECT:-ielts-verify}"

export IELTS_RUNNER_ROOT="$ROOT"
export IELTS_RUNNER_DOCKER_SOCKET="$DOCKER_SOCKET"
export ILETS_VERIFY_COMPOSE_PROJECT="$VERIFY_COMPOSE_PROJECT"

RUNNER_COMPOSE=(docker compose --project-name "$RUNNER_PROJECT" -f "$RUNNER_COMPOSE_FILE")

if [[ -n "$("${RUNNER_COMPOSE[@]}" ps --status running -q runner)" ]]; then
  echo "IELTS verification runner is already running" >&2
  exit 1
fi

"${RUNNER_COMPOSE[@]}" rm --stop --force runner >/dev/null 2>&1 || true

exec "${RUNNER_COMPOSE[@]}" up \
  --no-deps \
  --abort-on-container-exit \
  --exit-code-from runner \
  runner
```

The image build immediately before this block remains unchanged except for using the explicit `ielts:latest` tag.

- [ ] **Step 3: Update README lifecycle and management guidance**

Document that successful verification leaves `ielts-runner-1` stopped and visible in OrbStack, while `ielts-verify` resources are transient. Explicitly state that a stopped runner created from a disposable checkout is for visibility/inspection only; users must invoke `./tools/verify-local` from a valid checkout to recreate it rather than manually restarting a runner whose original bind path may have been deleted.

- [ ] **Step 4: Run static checks**

Run the wrapper contract check and `bash -n tools/verify-local`. Parse `compose.runner.yml` as YAML and confirm the required service fields. Expected: all PASS.

### Task 3: Prove the exact committed SHA under OrbStack

**Files:**
- No source changes unless validation finds a defect.

**Interfaces:**
- Consumes: exact implementation commit SHA and local OrbStack Docker context.
- Produces: evidence for runner exit code, stable stopped Compose project, cleanup of transient verifier resources, and correct rebinding behavior.

- [ ] **Step 1: Run exact-SHA verification from a disposable checkout**

```bash
(
  set -e
  checkout="$(mktemp -d)"
  trap 'rm -rf "$checkout"' EXIT
  git clone https://github.com/phatnguyen03022001/ilets.git "$checkout/ilets"
  cd "$checkout/ilets"
  git checkout --detach <IMPLEMENTATION_SHA>
  echo "HEAD=$(git rev-parse HEAD)"
  ./tools/verify-local
)
echo $?
```

Expected: root `./verify` completes and final exit code is `0`.

- [ ] **Step 2: Verify OrbStack steady state**

```bash
docker image ls ielts
docker ps -a --filter label=com.docker.compose.project=ielts
docker ps -a --filter label=com.docker.compose.project=ielts-verify
docker network ls --filter label=com.docker.compose.project=ielts-verify
```

Expected: `ielts:latest` exists; one stopped `ielts` runner remains; no `ielts-verify` container or network remains.

- [ ] **Step 3: Verify rebinding from a second checkout**

Run the exact-SHA block again from a fresh disposable checkout. Expected: the earlier stopped runner is replaced, the second checkout verifies successfully, and the final stopped runner belongs to the same `ielts` project without stale source use.

- [ ] **Step 4: Record final status**

Do not claim completion until the exact implementation SHA has `./tools/verify-local -> ./verify -> exit 0` evidence and the OrbStack steady-state checks match the spec.
