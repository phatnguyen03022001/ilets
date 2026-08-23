ARG NODE_VERSION=24.19.0
ARG PYTHON_VERSION=3.14.7
ARG DOCKER_CLI_VERSION=29.7.2

FROM node:${NODE_VERSION}-bookworm-slim@sha256:3638d9a6fe4030bd716be989438248074489337ba3275657f93595428be4fc03 AS node_runtime
FROM docker:${DOCKER_CLI_VERSION}-cli@sha256:000bb62ff495f986c9f5578eb67cc2cb98b91138eda81d7762d5371eb8a497fe AS docker_cli
FROM python:${PYTHON_VERSION}-bookworm@sha256:8771427e2ac3e39208c1632f17e8b09e464333d262844a03705cc5e0023c16e2

ARG GO_VERSION=1.27.0
ARG PNPM_VERSION=11.21.0
ARG PLAYWRIGHT_VERSION=1.62.1
ARG TARGETARCH
ARG GO_AMD64_SHA256=675c26c449cbb18fc24b74650de1eabbae6e16f64326fd85a283fb3b58280685
ARG GO_ARM64_SHA256=51798d2c42d0e1c6ed7fd9f48728b4193abac9e8aad6dbac2fe96a81f5909bda

ENV PATH="/usr/local/go/bin:${PATH}" \
    PLAYWRIGHT_BROWSERS_PATH=/ms-playwright

COPY --from=node_runtime /usr/local/bin/node /usr/local/bin/node
COPY --from=node_runtime /usr/local/bin/corepack /usr/local/bin/corepack
COPY --from=node_runtime /usr/local/bin/npm /usr/local/bin/npm
COPY --from=node_runtime /usr/local/bin/npx /usr/local/bin/npx
COPY --from=node_runtime /usr/local/lib/node_modules /usr/local/lib/node_modules
COPY --from=docker_cli /usr/local/bin/docker /usr/local/bin/docker
COPY --from=docker_cli /usr/local/libexec/docker/cli-plugins/docker-compose /usr/local/libexec/docker/cli-plugins/docker-compose

RUN set -eux; \
    apt-get update; \
    apt-get install -y --no-install-recommends ca-certificates curl git; \
    rm -rf /var/lib/apt/lists/*; \
    case "$TARGETARCH" in \
      amd64) go_sha="$GO_AMD64_SHA256" ;; \
      arm64) go_sha="$GO_ARM64_SHA256" ;; \
      *) echo "unsupported target architecture: $TARGETARCH" >&2; exit 1 ;; \
    esac; \
    curl -fsSLo /tmp/go.tgz "https://go.dev/dl/go${GO_VERSION}.linux-${TARGETARCH}.tar.gz"; \
    echo "${go_sha}  /tmp/go.tgz" | sha256sum -c -; \
    tar -C /usr/local -xzf /tmp/go.tgz; \
    rm /tmp/go.tgz; \
    go version; \
    node --version; \
    python3 --version; \
    docker --version; \
    docker compose version

RUN corepack enable \
    && corepack prepare "pnpm@${PNPM_VERSION}" --activate \
    && pnpm --version

RUN npx --yes "playwright@${PLAYWRIGHT_VERSION}" install --with-deps chromium \
    && rm -rf /root/.npm

RUN git config --system --add safe.directory /workspace

WORKDIR /workspace
