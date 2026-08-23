# Persistent OrbStack IELTS Runner Design

## Goal

Make the local verification environment easy to manage in OrbStack by leaving one stable `ielts` Compose project visible after a run, while preserving repository-native `./verify` as the only PASS definition.

## Current state

`tools/verify-local` builds the pinned `ielts:latest` image and launches a one-shot `ielts-runner` container with `docker run --rm`. The container invokes `./verify`, which owns a disposable PostgreSQL Compose lifecycle. A successful run removes the runner, PostgreSQL container, and verification network, so OrbStack shows the image but no stable IELTS container/project.

## Design

Add a dedicated local runner Compose file under `tools/local-verify/` and make `tools/verify-local` use it to create/start the runner service instead of `docker run --rm`.

The stable Compose project is named `ielts` and contains one service, `runner`. The service uses the existing `ielts:latest` image, bind-mounts the checkout at `/workspace`, bind-mounts the active Docker Unix socket, uses host networking, and executes `./verify` as its command. After `./verify` exits, the runner remains as a stopped container, so OrbStack keeps showing the `ielts` project.

The repository-native verifier's PostgreSQL lifecycle remains separate and disposable. `tools/verify-local` passes `ILETS_VERIFY_COMPOSE_PROJECT=ielts-verify` unless the caller explicitly overrides it. This separation is required because root `./verify` runs `docker compose ... down --volumes --remove-orphans`; putting the runner in the same Compose project could cause the verifier to remove its own runner container.

Therefore OrbStack's steady state after a successful run is one stopped `ielts` project containing `runner`. During verification, a transient `ielts-verify` PostgreSQL project may be visible; root `./verify` removes it before exit.

## Naming

- Runner image: `ielts:latest`
- Stable Compose project: `ielts`
- Stable service: `runner`
- Stable container: Compose-generated under project `ielts` (for example `ielts-runner-1`)
- Disposable verification Compose project: `ielts-verify`
- Existing `ILETS_*` environment variables remain unchanged.

## Lifecycle

1. `tools/verify-local` resolves and validates the pinned toolchain versions and Docker socket exactly as today.
2. It builds or reuses `ielts:latest` from `tools/local-verify/Dockerfile`.
3. It uses the dedicated runner Compose file with project name `ielts`.
4. Before a run, it recreates the runner service so the bind mount always points at the checkout that invoked `tools/verify-local` and stale stopped containers do not retain an obsolete checkout path.
5. It starts the runner and attaches to its output.
6. The runner executes the checkout's own `./verify` without wrapping or replacing its PASS criteria.
7. `tools/verify-local` returns the runner process exit code.
8. The stopped runner container remains visible in OrbStack. The verifier-owned `ielts-verify` PostgreSQL resources are cleaned up by `./verify`.

## Correctness and safety constraints

- GitHub `main` remains the source of truth; repository source is never copied into the image.
- Root `./verify` remains the single PASS definition and is not modified for this change.
- A local wrapper PASS is exactly exit code 0 from `./verify`.
- The runner continues to use the pinned Node, pnpm, Go, Python, Docker CLI/Compose, and Playwright toolchain already established by the image.
- The pnpm store remains outside the bind-mounted checkout so verification leaves the working tree clean.
- No arbitrary shell interface, Git write capability, or new plugin authority is introduced.
- The Docker socket is exposed only to the local runner container because root `./verify` requires Docker Compose; it is not exposed as a ChatGPT/plugin tool.
- `tools/verify-local` must safely replace an earlier stopped runner before rebinding a different checkout.

## Failure behavior

If image build fails, no verification PASS is claimed. If `./verify` exits nonzero, `tools/verify-local` exits nonzero and the stopped runner is retained for inspection. Root verifier cleanup remains responsible for its disposable PostgreSQL resources.

## Verification

The implementation is accepted only when an exact committed SHA is tested under OrbStack and all of the following are observed:

1. `./tools/verify-local` returns exit code 0.
2. The runner output proves repository-native `./verify` completed.
3. The checkout is clean at the final `./verify` gate.
4. `docker compose -p ielts ... ps -a` shows a stopped runner after completion.
5. No `ielts-verify` container/network/volume remains after successful completion.
6. `ielts:latest` remains available for reuse.
7. Re-running from a different disposable checkout recreates the runner with the new bind mount and still verifies that checkout rather than stale source.

## Non-goals

- Making PostgreSQL persistent between verification runs.
- Changing product runtime topology.
- Changing root verifier semantics or weakening any gate.
- Adding paid GitHub compute.
- Migrating the ChatGPT MCP/plugin to the persistent runner in the same change; that is a follow-up after the runner design is proven locally.
