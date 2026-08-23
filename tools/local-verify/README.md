# Local verification runner

This runner packages the toolchain needed by the repository-native `./verify` command. It does not define another PASS condition: the container entry command is still `./verify`, unchanged.

The image contains only tools. Repository source is never copied into the image. `tools/verify-local` bind-mounts the checkout at `/workspace`, so the checkout remains the source being verified and can be discarded after the run.

Pinned runner toolchain (base images are also pinned by immutable multi-platform digest):

- Node from `.node-version` (`24.19.0` on the current main line)
- pnpm from `apps/web/package.json` (`11.21.0` on the current main line)
- Go from `.go-version`
- Python `3.14.7` (the Python 3.14 line used by repository CI)
- Docker CLI `29.7.2` with its Compose plugin
- Playwright version from `apps/web/package.json`, with Chromium and Linux browser dependencies installed in the image

## Docker resource names

The wrapper keeps local verification resources under one `ielts` namespace for easier OrbStack management:

- runner image: `ielts`
- runner container: `ielts-runner`
- verifier Compose project: `ielts` by default, so Compose-managed resources use the `ielts-*` prefix

The runner container is deliberately not attached to the Compose project. Repository-native `./verify` owns that Compose lifecycle and calls `docker compose down --remove-orphans`; keeping the runner outside the project prevents the verifier from deleting its own parent container.

`ILETS_VERIFY_COMPOSE_PROJECT` remains supported when an explicit project override is needed. Other existing verifier environment variable names are unchanged.

## Run from a disposable checkout

```bash
checkout="$(mktemp -d)/ilets"
git clone --depth 1 --branch main https://github.com/phatnguyen03022001/ilets.git "$checkout"
cd "$checkout"
./tools/verify-local
cd /
rm -rf "$checkout"
```

The first run builds the local `ielts` runner image. Subsequent runs may reuse Docker build layers, but each verification still executes the checkout's own `./verify` and its disposable Compose database lifecycle.

## OrbStack prerequisites

- OrbStack is installed and running.
- The active Docker context points at OrbStack and exposes a Unix Docker socket.
- Docker host networking is available. The wrapper uses `--network host` because the repository-native verifier intentionally reaches its Compose PostgreSQL port and E2E servers on `127.0.0.1`.
- The machine has internet access for the initial image build and for repository dependency/tool downloads performed by `./verify`.

No host installation of Node, pnpm, Go, Python, PostgreSQL, or Playwright is required.

The usual verifier port/project overrides remain available through `ILETS_VERIFY_COMPOSE_PROJECT`, `ILETS_VERIFY_POSTGRES_PORT`, `ILETS_E2E_CORE_PORT`, and `ILETS_E2E_WEB_PORT`; the wrapper merely forwards them to `./verify`.
