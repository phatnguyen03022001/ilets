# Local verification runner

This runner packages the toolchain needed by the repository-native `./verify` command. It does not define another PASS condition: the runner service command is still `./verify`, unchanged.

The image contains only tools. Repository source is never copied into the image. `tools/verify-local` bind-mounts the checkout at `/workspace`, so the invoking checkout remains the source being verified.

Pinned runner toolchain (base images are also pinned by immutable multi-platform digest):

- Node from `.node-version` (`24.19.0` on the current main line)
- pnpm from `apps/web/package.json` (`11.21.0` on the current main line)
- Go from `.go-version`
- Python `3.14.7` (the Python 3.14 line used by repository CI)
- Docker CLI `29.7.2` with its Compose plugin
- Playwright version from `apps/web/package.json`, with Chromium and Linux browser dependencies installed in the image

## Docker resource names

The wrapper keeps a stable local verification project visible in OrbStack:

- runner image: `ielts:latest`
- stable Compose project: `ielts`
- stable service: `runner`
- stable stopped container after a run: normally `ielts-runner-1`
- disposable verifier Compose project: `ielts-verify` by default

The stable runner and verifier-owned PostgreSQL resources intentionally use separate Compose projects. Repository-native `./verify` owns the PostgreSQL lifecycle and calls `docker compose down --volumes --remove-orphans`; keeping PostgreSQL under `ielts-verify` prevents that cleanup from removing the parent runner.

After verification finishes, the `ielts` runner container remains stopped so OrbStack continues to show the project. PostgreSQL, its verification network, and any verifier-owned volumes are removed by `./verify`.

`ILETS_VERIFY_COMPOSE_PROJECT` remains supported when an explicit verifier project override is needed. Other existing verifier environment variable names are unchanged.

## Run from a disposable checkout

```bash
checkout="$(mktemp -d)/ilets"
git clone --depth 1 --branch main https://github.com/phatnguyen03022001/ilets.git "$checkout"
cd "$checkout"
./tools/verify-local
cd /
rm -rf "$checkout"
```

The first run builds the local `ielts:latest` image. Subsequent runs may reuse Docker build layers. Before each verification, `tools/verify-local` removes an earlier stopped `ielts` runner and recreates it with the checkout that invoked the wrapper, then attaches to the runner until `./verify` exits. The wrapper returns the runner's exit code.

A stopped runner created from a disposable checkout is retained for OrbStack visibility and inspection only. If that checkout has since been deleted, do not manually restart the stopped container from OrbStack because its bind mount points at the old checkout path. Invoke `./tools/verify-local` from a valid checkout instead; the wrapper recreates the runner with the current path.

If an `ielts` runner is already running, the wrapper refuses to replace it so concurrent verification runs cannot silently stop each other.

## OrbStack prerequisites

- OrbStack is installed and running.
- The active Docker context points at OrbStack and exposes a Unix Docker socket.
- Docker host networking is available. The runner uses host networking because the repository-native verifier intentionally reaches its Compose PostgreSQL port and E2E servers on `127.0.0.1`.
- The machine has internet access for the initial image build and for repository dependency/tool downloads performed by `./verify`.

No host installation of Node, pnpm, Go, Python, PostgreSQL, or Playwright is required.

The usual verifier overrides remain available through `ILETS_VERIFY_COMPOSE_PROJECT`, `ILETS_VERIFY_POSTGRES_PORT`, `ILETS_E2E_CORE_PORT`, and `ILETS_E2E_WEB_PORT`; the runner Compose service forwards them to `./verify`. The pnpm content-addressable store is kept at `/tmp/ielts-pnpm-store` inside the runner so verification does not dirty the bind-mounted checkout.

## Inspect the stable project

After a run:

```bash
docker image ls ielts
docker ps -a --filter label=com.docker.compose.project=ielts
docker ps -a --filter label=com.docker.compose.project=ielts-verify
docker network ls --filter label=com.docker.compose.project=ielts-verify
```

The expected steady state after a successful verification is `ielts:latest` plus one stopped runner in project `ielts`, with no remaining `ielts-verify` container or network.
