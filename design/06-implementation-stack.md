STATUS: CANONICAL
OWNS: initial deployable-unit allocation, primary-language/framework assignment, bootstrap toolchain profile, runtime responsibility split, system-engineering concern disposition, cross-language contract strategy, canonical-registry materialization strategy, persistence/consistency engineering baseline, security/observability/performance/deployment engineering invariants, and repository/native verification contract
DEPENDS_ON: ../CONSTITUTION.md, 04-application-flows.md, 05-api.md
DOES_NOT_OWN: learning/product truth, parser/materializer implementation, registry serialization or codegen-library choice, CI platform configuration, exact dependency patch versions, cloud/provider choice, final database schema, concrete deployment topology, external-provider lifecycle/selection, evaluator model vendor, numerical SLO/timeout/retry/scaling thresholds, or package-manager lock state

# Implementation Stack

## Purpose

Assign approved languages/framework families and implementation-engineering invariants to explicit runtime responsibilities so implementation does not re-decide first-order architecture, duplicate domain logic across stacks, or cargo-cult infrastructure from generic system-design checklists.

Patch/minor dependency selection is implementation maintenance. This architecture freezes responsibility, coherent bootstrap tool families, and production concern semantics; it does not freeze volatile patch numbers, provider products, or empirical operational thresholds.

# Initial deployable topology

```text
apps/
└── web/                  TypeScript / React / Next.js App Router

services/
├── core-api/             Go / net/http + chi
└── evaluator/            Python / FastAPI

contracts/
├── http/                 exact HTTP contracts once materialized
└── events/               only for a real asynchronous cross-unit event boundary

tools/
└── contract/content/repository tooling when justified
```

Do not create another deployable merely because a product feature, operational concern, or external capability has a distinct name.

These are ownership/process boundaries, not a requirement for three separately billed infrastructure stacks. Initial deployment may co-locate runnable units when isolation, security, failure, performance, and operational requirements remain satisfied. Split infrastructure only when a demonstrated requirement justifies the added cost/complexity.

# System-engineering concern disposition

Every production/system concern is reasoned about using one of four conceptual dispositions. These are architecture-review classes, not required runtime enums or machine-contract fields.

- **BOOTSTRAP_REQUIRED** — the applicable invariant must be addressed when the relevant runtime capability first exists.
- **PRODUCTION_GATE** — the concern may be implemented incrementally, but the applicable release path cannot pass production/support gates until it is resolved and verified.
- **CONDITIONAL** — required only after a named architectural/runtime trigger exists.
- **NOT_SELECTED_INITIAL** — the concern is understood but the named technology/pattern is intentionally not selected at bootstrap because no demonstrated trigger exists.

A concern may be mandatory while a particular external product remains optional. Technology is promoted from conditional/non-selected only when the trigger is demonstrated.

## Bootstrap-required concern groups

When the relevant capability exists, implementation resolves at least:

- bounded request/network deadlines, cancellation where safe, retry classification, idempotency/deduplication, bounded retry/backoff where eligible, and backpressure;
- race/concurrency protection, transaction ownership, commit-before-success-ACK, stale-write protection, and safe handling of ambiguous network outcomes;
- versioned database migrations/schema compatibility once persistence schema exists, parameterized SQL, bounded connection management, and rollback/forward-recovery discipline;
- authentication/authorization boundary, least-privilege runtime access, secrets handling, production TLS, browser XSS/CORS/CSP/CSRF policy as applicable, and SSRF-safe external URL/media access;
- structured logging, request/work correlation, health semantics, running software/contract version identity, dependency pinning/lockfiles, reproducible build/verification;
- API/contract compatibility and canonical-registry/binding drift verification once materialized.

## Production-gate concern groups

Before the corresponding scoped product path is production-supported, the selected runtime/deployment must resolve and verify as applicable:

- backup plus restore testing, disaster-recovery/recovery procedure, migration/deploy recovery, release rollback/forward-recovery, and failure/degraded-state testing;
- monitoring, metrics, actionable alerting, operational auditability, production incident ownership/escalation, and material-incident follow-up;
- latency/throughput/capacity/tail-latency measurement, async backlog/DB pressure visibility, and external usage/cost measurement;
- security verification for deployed browser/API/storage/provider boundaries;
- measurable SLIs and release-appropriate SLOs/error-budget semantics without inventing unsupported numerical objectives.

## Conditional / initial non-selection

The following remain trigger-based rather than bootstrap requirements:

- reverse proxy, external API gateway, load balancer — when deployment/public-edge/multiple-instance routing requires them;
- CDN/edge cache — when eligible static/public asset latency/traffic justifies them;
- autoscaling/horizontal scaling — after measured load and deployment shape justify multiple instances; vertical scaling is an acceptable initial path;
- database read replicas, partitioning, sharding, replication topology — after measured read/scale/recovery need;
- distributed locks, leader election, distributed transactions, Saga coordination — only when a proven cross-process invariant cannot be satisfied by one authoritative transaction/work owner;
- dedicated message queue, Pub/Sub, DLQ, event-driven architecture — only when durable database work/outbox + bounded dispatch cannot meet measured reliability/throughput/fan-out needs;
- WebSockets/long polling — only when SSE + resource reads cannot satisfy actual client semantics;
- inbound webhooks — only for a selected provider requiring callbacks;
- gRPC — only if HTTP contract semantics are demonstrably insufficient;
- circuit-breaker mechanism beyond bounded failure/retry/degradation — after repeated-route failure data justifies it;
- multi-region, chaos engineering, dedicated failover topology — after release consequence/availability/recovery requirements justify them;
- WAF/dedicated DDoS product — when public-edge deployment/risk requires more than hosting/native protection;
- Docker — optional reproducible packaging when deployment/build benefits;
- Infrastructure as Code/Terraform — when concrete infrastructure state is complex/material enough to require reproducible managed definition;
- Kubernetes/Helm/service-discovery infrastructure — only after deployment topology/scale creates that need;
- serverless-specific cold-start/limit handling — only if serverless deployment is selected;
- OAuth/JWT-specific mechanics — only if those identity/token schemes are selected;
- external feature-flag service — only when first-party bounded flags are insufficient.

Explicit initial non-selection includes Kubernetes, Helm, dedicated broker/Kafka-class infrastructure, distributed transactions, Saga orchestration, dynamic service-discovery infrastructure, multi-region, database sharding/read replicas, leader election, gRPC, and WebSockets while current HTTP + SSE semantics remain sufficient.

# Reuse-first implementation invariant

Before adding a package, service, external call, generated artifact class, or infrastructure subsystem, implementation first determines whether the requirement can be satisfied by an existing canonical semantic, runtime owner, contract, content asset, standard-library/framework capability, or already-approved provider route.

Preferred execution order when resulting semantics/quality are equivalent:

```text
existing canonical/runtime/content capability
        ↓
deterministic or local first-party execution
        ↓
existing eligible external-provider capability
        ↓
new external provider or infrastructure only for a demonstrated gap
```

Rules:

1. deterministic scoring/validation remains deterministic; do not call AI to replace an available exact rule;
2. browser/native/framework capabilities should be reused before duplicate backend capability when ownership remains correct;
3. derived work may be cached/reused by stable input + relevant policy/model/version identity when correctness/privacy/freshness/auditability hold;
4. retries reuse logical work identity instead of duplicating provider work/cost;
5. expensive noninteractive work may be delayed, batched, or omitted when optional; cost pressure cannot lower evidence/quality standards;
6. do not pre-generate infrastructure/content/AI output merely because a taxonomy exists;
7. optimization cannot move semantic authority into caches, generated files, prompts, provider output, logs, or metrics.

# Version and dependency policy

Architecture freezes compatibility families and responsibilities, not volatile patch numbers.

At implementation/bootstrap time:

- use a currently supported Node.js LTS line and supported TypeScript/React/Next.js App Router combination;
- use a currently supported Go release with `net/http` and `chi/v5`;
- use a currently supported Python 3 release with compatible FastAPI/Pydantic;
- pin exact dependency/runtime versions in manifests/lockfiles as appropriate;
- keep security/maintenance releases current through normal verified dependency maintenance;
- prefer one coherent tool per concern and bounded dependency count; overlapping frameworks/lint/test stacks require a concrete reason.

Semantic versioning/version ranges may be used according to each ecosystem, but lockfiles/manifests and repository verification determine the exact reproducible build state. Dependency updates cannot silently change machine-contract or canonical semantics.

# TypeScript — web

## Unit

```text
apps/web/
```

## Bootstrap family

```text
TypeScript
React
Next.js App Router
ESLint                  lint
Prettier                formatting
Vitest                  unit-level TypeScript tests
React Testing Library   component behavior
Playwright              critical browser E2E
browser fetch / generated contract client
```

Equivalent replacements require a concrete maintenance/compatibility reason; do not install overlapping tools merely for preference.

## Owns

- learner/admin rendering and route/layout composition;
- interactive Reading/Writing workspaces;
- browser microphone/recording capture;
- external embedded-player interaction;
- timers/local draft interaction/optimistic presentation;
- SSE client/reconnect behavior;
- presentation-only transformations;
- accessibility/responsive UI.

## Does not own

- mastery/progression/evidence policy;
- deterministic IELTS scoring;
- canonical gap/action or content eligibility;
- productive evaluator algorithms;
- durable learner/product truth;
- handwritten DTO truth independent of machine contracts.

## Client-state and transport baseline

Transient UI interaction may live locally. Durable learner/product state remains authoritative behind the Go API. No Redux/Zustand/global client store is selected by default; a client state/query library is conditional on demonstrated interaction complexity and may not become product truth.

The browser uses the exact public contract through browser `fetch` or a generated/validated client once contracts exist. Browser never calls Python directly.

## Web security baseline

- rely on React/Next normal escaping for ordinary rendering;
- dangerous/raw HTML requires an explicit sanitization policy before use;
- deployed web surface has deliberate Content Security Policy compatible with required scripts/media/embed sources;
- CORS is explicit for the actual deployment origins and is not wildcard-by-convenience for credentialed/sensitive APIs;
- CSRF protection is resolved from the selected credential/session transport before auth is production-supported;
- secrets never enter browser bundles or client-readable configuration.

# Go — Core API + deterministic orchestration

## Unit

```text
services/core-api/
```

## Bootstrap family

```text
Go
net/http
chi/v5
PostgreSQL-compatible persistence
pgx/v5 or equivalent direct PostgreSQL-capable driver
SQL-first persistence
standard structured logging capability where sufficient
```

A general-purpose ORM is not selected by default. Persistence boundaries are named for coherent product/domain capabilities, not generic repositories/helpers.

## Owns

- learner/admin-facing `/v1` API and auth/authorization integration boundary;
- durable LearnerProfile/TargetProfile/content/work/session/Attempt product state it owns;
- explicit authoritative transaction boundaries;
- DailyPlan/LearningSession/PracticeActivity orchestration;
- deterministic Listening/Reading scoring and deterministic exact validation rules;
- content validation-policy aggregation and content demand/reuse/assignment/incident orchestration;
- Assessment/Progression/Planner execution over canonical/materialized rules;
- idempotency/concurrency enforcement and async work orchestration;
- media-source eligibility/product state and product SSE delivery.

## Does not own

- learning/content semantic truth;
- AI rubric judgment, speech feature extraction, or LLM internals;
- browser interaction;
- duplicate Python evaluator behavior.

## Persistence/query baseline

- one authoritative PostgreSQL-compatible product store is the initial structured durable-state model; final provider/schema remains implementation detail;
- SQL uses parameterized queries; raw user values are never concatenated into SQL;
- transactions align with authoritative product invariants; network/provider calls are not treated as part of one atomic DB transaction;
- optimistic concurrency is used where current resource semantics require stale-write protection; pessimistic row locking is conditional on a proven invariant requiring serialization;
- no distributed lock, distributed transaction, Saga, read replica, sharding, or leader election is selected initially;
- connection pooling/management is bounded and observable; DB saturation cannot create unbounded application concurrency;
- indexes follow real access/query paths; do not index every field by default;
- query optimization is measured/query-plan-driven; N+1 risks are reviewed/tested at repository/query boundaries where material.

## Migration/schema discipline

Once database schema exists:

- migrations are explicit, ordered/versioned, verified, and use one coherent migration mechanism;
- application/schema compatibility during rollout is intentional rather than assuming lock-step replacement;
- migration failure cannot silently corrupt or lose accepted learner work;
- rollback or forward-recovery strategy is defined for schema/application changes before production use;
- storage schema/version is derived implementation, not canonical domain authority.

No migration package is canonically selected in this documentation pass.

# Python — evaluator/media analysis

## Unit

```text
services/evaluator/
```

## Bootstrap family

```text
Python 3
FastAPI
Pydantic-compatible typed contract models
uv                      project/environment workflow
Ruff                    format/lint
Pyright                 primary strict static type checker
pytest                  tests
```

A demonstrably better supported equivalent may replace a bootstrap tool without changing runtime ownership, but do not install duplicate toolchains for the same concern.

## Owns

- Writing/Speaking criterion observation generation;
- eligible speech transcription and validated pronunciation/fluency/acoustic extraction;
- bounded text analysis for feedback/error candidates;
- bounded AI-generated feedback/content candidates;
- bounded content-validation signals where deterministic checks are insufficient and capability is invoked;
- authorized transcript/media analysis;
- evaluator/model/generator/validator provenance and uncertainty output.

## Does not own

- learner-facing public API;
- durable learner/progression/content operational state;
- content activation/assignment eligibility;
- certification, Band advancement, evidence sufficiency, or DailyPlan selection;
- provider-specific semantics as product truth.

Celery, a separate broker framework, and a second web framework are not selected initially. They require a demonstrated workload/reliability need.

# One public product API and network baseline

```text
Browser
  ↓ HTTP/TLS in production
Go Core API
  ↓ internal HTTP contract where needed
Python Evaluator
```

SSE is selected for server→browser update hints. WebSockets are not selected while bidirectional realtime semantics are absent. Long polling is only a conditional fallback for a real client/environment need. gRPC is not selected initially.

HTTP/2/HTTP/3, reverse-proxy behavior, and DNS are deployment/hosting concerns; application semantics do not depend on a particular HTTP generation. The application does not design custom TCP/UDP transport at bootstrap.

DNS/route failures and transport-level connection success/failure are operational facts, never business/learner-state truth.

# Contract strategy

## HTTP

Before independent runtimes implement the same boundary, materialize one exact contract under `contracts/` as required by `05-api.md`.

Typical structure:

```text
contracts/http/openapi.yaml
```

Each boundary has one exact machine authority. Generated clients/server bindings/validators are derived artifacts.

## Events

Do not create hypothetical event schemas. Create event contracts only when an actual asynchronous cross-unit event boundary exists and HTTP/work-resource semantics are insufficient.

## Cross-language engineering baseline

Shared implementation semantics are defined once through canonical owners/contracts/materialization rather than invented independently in each language, including where material:

- stable canonical IDs and exact content revision identity;
- request/work/correlation identity;
- timestamp/time-zone representation and the rule that timestamps are not causal/idempotency identity;
- error/response semantics;
- contract/version/provenance identity;
- null/not-applicable semantics;
- serialized enum/canonical-value source;
- trace/correlation propagation.

Exact wire representation belongs to contracts once materialized.

# Canonical registry materialization

Shared canonical identities consumed by implementation are materialized through a derived pipeline:

```text
canonical Markdown owners
        ↓
repository materializer / validator
        ↓
derived machine-readable registry
        ↓
generated or validated Go / TypeScript / Python consumers
```

Invariants:

1. canonical Markdown owner remains semantic authority;
2. machine-readable registry and generated bindings/constants/models are derived artifacts, not another SSOT;
3. equivalent canonical enum/ID registries are not manually maintained independently in Go, TypeScript, and Python;
4. duplicate canonical IDs, broken canonical references, and materialized generated-registry/binding drift fail repository verification;
5. derived artifacts preserve sufficient source/provenance identity to trace values to canonical owner/source revision;
6. stable canonical IDs remain unchanged across language/machine boundaries;
7. materialization may be incremental to registries actually consumed by the current implementation slice;
8. a parser/materializer/tool defect fails verification or leaves output unresolved; tooling cannot silently reinterpret canonical meaning.

Materializer language/parser/serialization/file/codegen format remain implementation choices.

# Async work, scheduling, and process correctness baseline

The execution patterns in `04-application-flows.md` are implemented without assuming broker infrastructure.

Initial direction:

```text
authoritative DB work state
+ durable work/outbox semantics where needed
+ idempotent bounded dispatch
+ SSE/resource status
```

Message queues, Pub/Sub, DLQs, event-driven architecture, and external workflow engines remain conditional on measured reliability/throughput/fan-out need.

Process correctness requirements include:

- race-sensitive mutations use transaction/concurrency controls rather than timing assumptions;
- lock ordering/critical sections must avoid deadlock; a deadlock/serialization retry cannot duplicate logical work;
- shared-memory/thread safety follows the selected runtime/library model; cross-process correctness never relies on in-memory mutexes alone;
- memory usage is bounded for request bodies, media, queues, caches, and provider outputs; long-lived growth/leaks are observable/testable;
- garbage collection/runtime memory behavior is implementation performance concern, not semantic correctness authority.

# Cache model

Caching is an optimization, never authority.

A correctness-sensitive cached value is valid only relative to:

```text
source identity
+ relevant source/policy/model/contract version
+ freshness/invalidation rule
```

Source/policy/version changes that affect correctness invalidate the derived cache. Cache cannot become authority for learner state, content semantic truth, product support, evidence, or progression.

CDN/edge caching is conditional and limited to eligible public/static/derived assets whose privacy, rights, freshness, and invalidation remain correct. Redis is not required merely because caching exists.

# Security engineering baseline

Security concern ownership remains implementation/product-runtime, not learning truth.

## Application/data security

- SQL injection: parameterized SQL only;
- XSS: framework escaping by default; explicit sanitization for raw HTML;
- SSRF: learner/media URLs do not grant arbitrary network access; allowed destinations/schemes/redirects/private-network reachability are constrained by intended source policy;
- CORS: explicit allowed deployment origins/access semantics;
- CSP: deliberate deployed web policy for scripts/media/embeds;
- CSRF: resolved according to actual credential/session transport;
- secrets: never committed/logged or stored as learner/product canonical state; supplied through deployment/runtime configuration;
- IAM: least privilege with runtime/service/admin privileges separated;
- TLS/encryption in transit: required for production sensitive cross-network boundaries;
- encryption at rest: applicable learner/secrets/artifact protection must be satisfied by selected storage/deployment;
- OAuth is conditional on identity route selection; JWT rotation is conditional on JWT selection, while the broader credential/session/key revocation/rotation requirement applies to whatever credential mechanism is selected;
- WAF/dedicated DDoS product is conditional, but public-edge abuse/resilience must be addressed by the selected hosting/deployment path.

No security telemetry/cache/provider may become authoritative learner evidence or product state.

# Health, observability, and incident baseline

## Health semantics

Each runnable unit exposes enough health semantics for the selected deployment to distinguish process availability from readiness to serve its required dependencies. Concrete liveness/readiness probe mechanisms are deployment-conditional; a health endpoint cannot claim downstream semantic correctness it has not verified.

## Structured operational telemetry

Telemetry supports, where privacy-safe/material:

```text
request_id
work_id
trace_id when tracing is used
runtime/service
operation
duration
outcome/result class
stable non-sensitive error code
privacy-safe content/evaluation/provider refs where needed
software/contract/version provenance where consequential
```

Do not log by default passwords, auth tokens, provider secrets, raw audio, unbounded Writing/Speaking content, or sensitive full request bodies.

Metrics can measure as applicable request volume, error/failure rate, latency distributions/tail latency, DB pool pressure, async backlog, evaluation duration, retries, provider failures, and external evaluation/generation/validation usage/cost.

Distributed tracing is conditional on multi-unit/debugging need. Trace correlation is derived telemetry, not business authority.

## SLI/SLO/error-budget semantics

A supported production release defines measurable SLIs/SLOs appropriate to its actual consequence and usage, with error-budget/reliability trade-off reasoning where useful. Exact numeric objectives are not frozen before operational evidence exists.

## Incidents

Before product support there is a viable operational process:

```text
detect
→ contain / degrade safely
→ preserve accepted learner work
→ recover
→ verify
→ record material cause + follow-up
```

Production ownership/escalation must be explicit for the supported release without inventing an enterprise on-call rota. Material incidents affecting learner data/work, evidence/content integrity, security, or sustained availability receive a post-incident record/postmortem. Operational history is not canonical learning truth.

# Reliability, backup, and disaster-recovery baseline

- backup existence alone is insufficient; restore is verified before the applicable production gate passes;
- PITR is required only where the chosen deployment/support model makes it applicable;
- recovery objectives/numerical RPO/RTO are defined from actual support needs, not guessed in architecture;
- migration/deploy failure preserves committed accepted learner work and has rollback/forward-recovery procedure;
- failover, multi-region, chaos engineering, or redundant database topology are conditional on demonstrated availability/recovery consequence;
- network partition/provider ambiguity exposes unresolved/pending state rather than false success/failure.

# Performance, capacity, and scaling model

Material implementation work falls into request/work classes:

- interactive synchronous read;
- durable synchronous mutation;
- asynchronously accepted work;
- internal evaluator/provider call;
- SSE notification;
- scheduled/background/batch work.

For each materialized class, implementation can measure/configure as applicable deadline, latency, p95/p99/tail latency where useful, throughput, concurrency, backlog, retry budget, backpressure, and cost. Architecture does not freeze arbitrary thresholds.

Vertical scaling/co-location may be the initial simplest capacity path. Horizontal scaling/autoscaling/load balancing become conditional when measured load and deployment shape require multiple instances. Cold starts/serverless limits apply only if serverless deployment is selected.

# Deployment and release baseline

Any selected deployment route must support:

- reproducible build from pinned/locked dependency state;
- runtime/software/contract version identity;
- environment/config separation and safe secrets;
- health semantics;
- migration-safe deployment;
- rollback or forward-recovery;
- observability of the running version.

CI/CD automation eventually invokes the same root `verify` correctness contract. No paid GitHub capability is required by this architecture.

Docker may be used for reproducible packaging when helpful; it is not mandatory independent of deployment need. Kubernetes is `NOT_SELECTED_INITIAL`; Helm is not applicable unless Kubernetes is selected. Terraform/IaC is conditional on concrete infrastructure requiring reproducible managed state. Blue-green, canary, and rolling deployment are conditional strategies; the invariant is safe rollout plus recovery/rollback, not a mandatory named strategy. Service-discovery infrastructure is not selected for the known bounded topology unless deployment makes dynamic discovery necessary.

Build caching is allowed as optimization but cannot hide stale generated artifacts or skip correctness checks.

# Feature flags

First-party bounded flags/kill switches may be used for safe rollout/degradation when material. A flag:

- cannot redefine canonical learning/content/evidence truth;
- cannot bypass validation/evidence/support standards;
- cannot mutate historical meaning;
- has an owner, safe default, and lifecycle/expiry when materially persistent.

External feature-flag providers remain optional/conditional.

# Native verification baseline

Each deployable owns native checks while repository correctness remains one root contract.

## Web

At minimum:

```text
Prettier format check
ESLint
TypeScript typecheck
Vitest unit tests
React Testing Library component behavior where material
Next.js production build
Playwright critical browser E2E
security-sensitive rendering/URL handling tests where material
```

## Go

At minimum:

```text
gofmt check
go vet ./...
go test ./...
race tests where material
build core-api
DB integration tests once persistence exists
migration tests once schema exists
query/index behavior tests where performance-sensitive
```

## Python

At minimum:

```text
Ruff format check
Ruff lint
Pyright strict/static type check
pytest
internal-contract/conformance tests once materialized
```

## Contracts/registries

At minimum once materialized:

```text
schema validation
generated-artifact drift check
canonical ID/reference validation
consumer/provider conformance
public API → evaluator/content-capability integration tests where implemented
backward-compatibility checks where deployed contracts require them
```

# Root verification

The repository uses one root verification contract for local and automated verification. Checks may be absent while the corresponding unit/materialized artifact does not exist; once a unit, registry, contract, persistence boundary, or cross-unit path appears, its applicable checks enter this same root contract.

```text
verify
  ├── repository/canonical
  │     ├── canonical metadata
  │     ├── semantic dependency/reference integrity
  │     ├── canonical ID uniqueness
  │     └── generated registry/binding drift where materialized
  ├── contracts where materialized
  ├── web where materialized
  ├── core-api where materialized
  ├── evaluator where materialized
  ├── persistence/migrations where materialized
  └── cross-unit integration where materialized
```

CI invokes this same root contract. CI may add execution environment/triggers/reporting but cannot create a separate hidden definition of correctness. Local verification and CI agree on PASS for the same repository state.

A cross-stack change is not PASS because only one affected ecosystem is green.

# Framework/infrastructure change rule

A framework/infrastructure replacement requires a design review/change when it materially changes deployable boundaries, API/transport ownership, rendering/runtime model, persistence ownership/consistency, cross-language contracts, provider boundary, or operational complexity.

Patch/minor maintenance inside the same responsibility boundary is implementation work.

# Initial non-goals / anti-cargo-cult guard

Do not introduce at bootstrap without demonstrated trigger:

```text
need async        → Kafka/broker
need cache        → Redis
need similarity   → vector database
need scaling      → Kubernetes
need reliability  → multi-region
need auth         → mandatory JWT
need events       → event-source everything
need architecture → microservice-per-feature
```

Also do not add GraphQL beside REST, separate BFF, duplicate Go/Python domain rules, or frontend-owned Band/mastery/content eligibility without an explicit architecture change and demonstrated need.