STATUS: CANONICAL
OWNS: initial deployable-unit allocation, primary-language/framework assignment, bootstrap toolchain profile, runtime responsibility split, material implementation-boundary model, system-engineering concern disposition, cross-language contract/evolution strategy, canonical-registry materialization/evolution strategy, persistence/consistency engineering baseline, configuration/data-lifecycle/security/observability/performance/deployment engineering invariants, and repository/native verification contract
DEPENDS_ON: ../CONSTITUTION.md, 04-application-flows.md, 05-api.md
DOES_NOT_OWN: learning/product truth, parser/materializer implementation, registry serialization or codegen-library choice, CI platform configuration, exact dependency patch versions, cloud/provider choice, final database schema, concrete deployment topology, external-provider lifecycle/selection/ingress/egress details, evaluator model vendor, numerical SLO/timeout/retry/scaling thresholds, or package-manager lock state

# Implementation Stack

## Purpose

Assign approved languages/framework families and implementation-engineering invariants to explicit runtime responsibilities so implementation does not re-decide first-order architecture, duplicate domain logic across stacks, collapse trust boundaries for convenience, or cargo-cult infrastructure.

Architecture freezes responsibility, bootstrap tool families, material boundary/consistency/evolution semantics, and production concern semantics. Patch versions, provider products, deployment vendors, and empirical operational thresholds remain implementation/deployment decisions.

# Initial deployable topology

```text
apps/
└── web/                  TypeScript / React / Next.js App Router

services/
├── core-api/             Go / net/http + chi
└── evaluator/            Python / FastAPI

contracts/
└── http/                 exact HTTP contracts once materialized

tools/                    contract/content/repository tooling when justified
```

A separate event-contract area is introduced only for a genuine asynchronous cross-unit event boundary. SSE over the public HTTP API remains part of the public HTTP contract as owned by `05-api.md`.

Do not create another deployable merely because a feature, operational concern, or external capability has a distinct name.

## Co-location invariant

These are logical/runtime ownership boundaries, not a requirement for three separately billed infrastructure stacks. Units may be co-located when the selected deployment preserves:

- Browser/product behavior through Go Core API;
- Next.js as Web/presentation authority rather than product-state authority;
- Core-only access to authoritative product persistence;
- explicit cross-runtime machine contracts;
- least-privilege runtime credentials/secrets;
- no undocumented shared-memory state replacing a declared boundary;
- attributable restart/failure behavior per logical unit.

Splitting or co-locating deployment changes topology/operations, not semantic authority.

# Material boundary model

A material implementation boundary is any trust/state/contract boundary where a shortcut could change authority, correctness, security, recoverability, compatibility, or coverage. It is a review concept, not a runtime enum.

For each material boundary implementation can establish, where applicable: owner; caller/callee; trust/access scope; allowed/forbidden data path; machine-contract authority; commit/consistency class; retry/idempotency/recovery; cache/freshness; privacy/security; observability/audit; version compatibility; and verification evidence.

A boundary may delegate detailed behavior to an existing owner. It does not become another semantic owner merely because implementation needs to enforce it.

## Semantic authority → implementation

```text
canonical spec/design
        ↓
materialized registry / machine contract
        ↓
derived bindings / implementation
```

Generated code, OpenAPI types, DB schemas, UI types, prompts, provider responses, caches, metrics, migrations, and storage tables are derived. None can redefine canonical meaning.

## Browser/user → Web

Browser/user input is untrusted. Web bounds presentation input and forwards authoritative reads/mutations through Go. Client state is presentation/transient unless persisted through the Go API.

Browser/Web cannot declare evidence eligibility, mastery, product support, immutable content classification, or privileged capability. Security controls such as XSS/CSP/CSRF/CORS/upload bounds follow the selected browser/session path.

## Next.js server execution → Go

Server Components, Route Handlers, Server Actions, server-side fetch, and eligible presentation caches remain inside `apps/web`.

They may render/compose presentation, perform presentation-safe calls to the exact Go public API, handle explicitly approved web/session-edge mechanics, and cache derived presentation/API state under correct access/freshness rules.

They may not:

- become a second product/domain API;
- read or write the authoritative product database directly;
- own durable learner/product/content/evidence state;
- duplicate Assessment/Progression/Planner/content policy;
- call Python as an alternate backend;
- create independent DTO/interface truth;
- bypass Go access/capability policy.

## Web → Go public API

Go owns learner/admin-facing product behavior and durable product state. One public HTTP contract governs exact IDs/applicability, auth/session transport, result/failure semantics, idempotency, concurrency, SSE, and compatibility once materialized. Browser and Next.js use that contract; neither bypasses Go to Python.

## Go → authoritative PostgreSQL store

For the initial application architecture, **only Go Core API reads from and writes to the authoritative product PostgreSQL-compatible store**.

Therefore:

- Python does not query or mutate it directly;
- Next.js does not query or mutate it directly;
- Browser cannot access it;
- read-only application-runtime access is still coupling and is not an exception;
- additional application-runtime access requires an architecture change.

This restriction does not prohibit migration tooling, backup/restore tooling, repository verification, operational database administration, or a future explicitly designed derived analytics/read model with its own non-authoritative boundary.

Database access uses parameterized SQL, explicit transaction ownership, bounded pooling, migration/application compatibility, and concurrency control appropriate to the invariant. Network/provider calls are never one atomic SQL transaction. Database schema remains derived storage implementation, not semantic authority.

## Go → Python Evaluator

Go owns authoritative work identity/state and is the permitted product caller. Python owns bounded evaluation/media/content-analysis capabilities behind one internal HTTP machine contract.

The evaluator is an **internal capability**, not public API surface:

- Browser and admin UI cannot call it directly;
- public-internet reachability is not assumed;
- deployment restricts reachability according to the selected topology;
- trust is determined by actual reachability, authorized principals, process/container isolation, deployment controls, and attack surface rather than network labels;
- `localhost`, a private IP, same VPC, same host, same cluster, or an `internal` hostname alone does not prove a trusted caller boundary;
- if an unauthorized principal/process/workload can reach the evaluator, caller/service authentication and authorization are required;
- a smaller mechanism is acceptable only when the selected isolation/reachability boundary actually makes it safe;
- this design does not freeze mTLS, JWT, or service mesh.

Python returns bounded result/signals plus provenance/model/evaluator identity and uncertainty/quality where material. Responses remain non-authoritative until Core validates/interprets them through owning policy. Python cannot directly read/write the product DB.

## Authoritative state → async work

Implementation preserves the distinct states defined by `04-application-flows.md`:

```text
authoritative mutation
required durable work identity/marker
dispatch attempt
remote/capability execution
result reconciliation
```

Required continuation is durably registered with or recoverably derivable from committed state before acknowledgement depends on it. Timeout does not prove remote non-execution. Retry/reconciliation reuses logical work/idempotency identity and preserves execution-attempt/fencing state where overlapping executions are possible.

## Cache boundary

Cache is derived. Correctness-sensitive cache identity includes relevant source identity/revision, policy/model/contract identity where material, access scope, and freshness/invalidation state.

Stale cached output cannot overwrite newer authority. Cache never becomes learner/evidence/content/product/progression/capability authority. Redis is not implied.

## Object/media storage boundary

When large artifacts are stored separately, authoritative Core state stores the reference/metadata needed to govern them; object/media storage stores bytes behind that authority. The chosen implementation addresses access scope, integrity, private-by-default behavior, lifecycle/recovery, provenance/rights, and orphan reconciliation as applicable. No provider is selected here.

### Conditional direct browser byte transfer

If a future storage route uses direct browser upload/download:

```text
Core authorizes one narrow temporary transfer grant
        ↓
Browser transfers bytes
        ↓
Core reconciles completion + integrity + authoritative metadata
        ↓
artifact becomes usable only under normal product policy
```

Direct byte transfer is not direct database authority, content activation, learner-state mutation, or permanent anonymous object access.

## Generic external runtime boundary

Any external runtime/provider crossing is untrusted ingress/egress behind an explicit provider-neutral capability/adapter/contract. It cannot transfer Core/learning/content/evidence authority to an external system.

Detailed provider lifecycle, URL/media ingress, data egress, callbacks/webhooks, provider trust, processor retention/reuse, activation/degradation, and provider-specific fallback semantics are owned downstream by `07-third-party-services.md`. This implementation owner does not restate those rules.

## Admin/privileged actor → Core

Identity, role membership if used, operational capability, learning authority, Assessment authority, and validation authority remain distinct. Admin mutations go through Core, use legal preconditions/capability checks, and create durable audit where consequential. Privilege cannot bypass hard learning/evidence/content validation rules or mutate historical semantic revisions.

## Observability boundary

Operational logs, security events, privileged audit, product analytics, traces/metrics, and learner-visible history are different data classes.

Telemetry is operational evidence, not business state. It carries enough privacy-safe request/work/build/contract/provenance identity to reconstruct consequential execution while redacting secrets and minimizing learner content. Raw audio, auth tokens, provider secrets, and sensitive full bodies are not logged by default.

Telemetry delivery failure normally does not rewrite a valid transaction; an explicitly required durable security/privileged audit record may be a precondition/transactional requirement for the consequential operation it protects.

## Time/clock boundary

Cross-language durable timestamps use timezone-aware absolute instants once serialized by machine contract. Wall clock is not causal ordering, idempotency identity, or optimistic-concurrency identity. Recency/expiry names its policy/clock where material.

## File/request-size boundary

Upload/audio/text/media ingress is bounded. Exact content type/size/duration/streaming/time limits are implementation policy based on the actual accepted formats and risks, not architecture constants.

# Internal module dependency direction

Do not freeze folder names or impose ceremony-heavy layers. Preserve conceptual direction and inversion at concrete infrastructure boundaries.

## Go

```text
transport adapter
      ↓
application / orchestration
      ↓
domain / owned deterministic policy

application/domain
      ↓ depends on
capability + persistence interfaces/ports
      ↑ implemented by
PostgreSQL / provider adapters

composition/bootstrap wires implementations
```

Application/domain code does not import concrete pgx/provider SDK policy as semantic authority. Concrete adapters implement required ports; they do not become policy owners.

## Python

```text
internal HTTP adapter
      ↓
capability application/service
      ↓
bounded evaluator implementation

provider/tool adapters implement required capability boundaries
```

Provider SDK models do not leak across the internal machine contract or become product semantics.

## Web

```text
routes / layouts / components
      ↓
presentation / application interaction
      ↓
generated or validated Go public-API client
```

UI components do not own server state transitions or duplicate server policy.

# Consistency classes

These are reasoning classes, not runtime enums.

## AUTHORITATIVE_TRANSACTIONAL

Core-owned facts committed under one authoritative transaction/invariant, including required durable work/audit markers when atomicity with a mutation is required. Durable-success response follows commit.

## AUTHORITATIVE_ASYNC_STATE

Core-owned work state progresses asynchronously while dispatch/capability execution may lag. One logical work identity remains authoritative.

## DERIVED_EVENTUALLY_CONSISTENT

Caches, projections, analytics, indexes/search, telemetry, and similar derived outputs may lag authority but preserve enough source/freshness identity to avoid overwriting newer truth.

## EXTERNAL_OBSERVED_STATE

External/provider observations require provenance/association/reconciliation before owning product policy can consume them.

Product/runtime failure behavior is owned in detail by `04-application-flows.md`. Implementation must make those outcomes possible through correct persistence, consistency, recovery, capacity, and verification; this file does not maintain a second failure list.

# Data lifecycle and deletion boundary

Retention/deletion has one authoritative product decision for the affected data scope. Implementation propagates that decision as applicable to:

- authoritative DB state;
- object/media bytes;
- caches;
- indexes/search/derived projections;
- product analytics/telemetry according to their data classification;
- external processors/providers under `07-third-party-services.md`;
- backup/restore reconciliation.

Derived stores cannot silently resurrect deleted data into active product use. Restore procedures reconcile restored data against current deletion/tombstone state before normal active use. Historical semantic records needed for integrity may be retained only under an explicit applicable product/security/data policy rather than accidental persistence.

This architecture does not invent retention durations or legal obligations.

# Configuration boundary

Configuration is derived implementation/deployment input, not semantic authority.

Configuration may select, where already authorized by canonical design:

- endpoint/location;
- runtime/deployment mode;
- capacity/concurrency;
- feature availability/kill switch;
- credential/secret reference;
- an already eligible provider/capability route;
- operational threshold owned by runtime policy.

Configuration must not redefine:

- Band semantics;
- Skill/Knowledge identity;
- evidence meaning or Assessment authority;
- content semantic identity;
- Progression rules;
- Coverage meaning.

Security-critical missing/invalid configuration fails closed. Secrets and ordinary configuration are separate classes. A consequential configuration/revision that changes evaluator/provider execution or interpretation must be reconstructable where required for audit/provenance. `.env` files or deployment variables are never policy authority.

# System-engineering concern disposition

Review classes are:

- **BOOTSTRAP_REQUIRED** — invariant resolved when the relevant runtime capability first exists;
- **PRODUCTION_GATE** — may be incremental but must pass before applicable product support;
- **CONDITIONAL** — required only after its trigger exists;
- **NOT_SELECTED_INITIAL** — named technology/pattern intentionally not selected without demonstrated need.

## Bootstrap-required concerns

When applicable:

- bounded deadlines/input/concurrency/backpressure and safe retry/idempotency;
- transaction ownership, commit-before-success, durable async recoverability, stale-write/ambiguous-outcome handling;
- parameterized SQL, bounded DB connections, migrations/schema compatibility when schema exists;
- auth/access boundaries, least privilege, secrets, production transport protection, browser/session controls, safe external-boundary enforcement;
- structured logs/correlation, health, running build/contract identity, dependency locking/reproducibility;
- contract compatibility/conformance and canonical-registry drift verification once materialized;
- security-critical configuration validation and data-lifecycle enforcement for implemented stores.

## Production-gate concerns

Before applicable `COVERED`/`SUPPORTED_FOR_PRODUCT` promotion, the release candidate resolves and verifies as applicable backup/restore, deploy/migration recovery, failure/degraded behavior, monitoring/alerts/audit, security verification, capacity/backlog/external-cost visibility, and measurable operational objectives appropriate to the release. Architecture does not invent numeric SLO/RPO/RTO values.

## Conditional / initial non-selection

Reverse proxies/gateways/load balancers, CDN, autoscaling, read replicas/sharding, distributed locks/transactions/Sagas, brokers/PubSub/DLQ, WebSockets, gRPC, circuit breakers beyond bounded retry/degradation, multi-region, chaos engineering, dedicated WAF/DDoS products, Docker, Terraform/IaC, Kubernetes/Helm/service discovery, serverless-specific handling, OAuth/JWT-specific mechanics, and external feature-flag services remain trigger-based.

Kubernetes, Helm, Kafka-class broker infrastructure, distributed transactions/Sagas, multi-region, DB sharding/read replicas, leader election, gRPC, and WebSockets are not selected initially.

# Reuse-first implementation invariant

Prefer an existing canonical/runtime/content capability, then deterministic/local first-party execution, then an already eligible external capability, and only then new infrastructure/provider when a demonstrated gap remains.

Optimization cannot move semantic authority into caches, generated files, prompts, provider output, logs, metrics, migrations, configuration, or DB schema. Cost pressure cannot lower evidence/content/evaluator quality for the same intended consequence.

# Version and dependency policy

Architecture freezes compatibility families/responsibilities, not patch numbers.

At bootstrap:

- use a currently supported Node.js LTS line and compatible TypeScript/React/Next.js App Router;
- use a currently supported Go release with `net/http` and `chi/v5`;
- use a currently supported Python 3 release with compatible FastAPI/Pydantic;
- pin exact dependency/runtime versions in manifests/lockfiles;
- maintain verified security/maintenance updates;
- avoid overlapping tools for the same concern without a concrete reason.

Dependency updates cannot silently change machine-contract or canonical semantics.

# TypeScript — Web

Unit: `apps/web/`.

Bootstrap family:

```text
TypeScript
React
Next.js App Router
ESLint
Prettier
Vitest
React Testing Library
Playwright
browser fetch / generated contract client
```

Owns learner/admin rendering, interactive workspaces, browser capture, embedded-player interaction, transient timers/drafts/optimistic presentation, SSE client behavior, presentation transformations, accessibility/responsiveness, and presentation-side Next.js server rendering/web-edge mechanics.

Does not own learning/evidence/progression policy, deterministic IELTS scoring, content eligibility, evaluator algorithms, durable learner/product truth, authoritative DB access, or independent DTO truth.

# Go — Core API + deterministic orchestration

Unit: `services/core-api/`.

Bootstrap family:

```text
Go
net/http
chi/v5
PostgreSQL-compatible persistence
pgx/v5 or equivalent direct PostgreSQL-capable driver
SQL-first persistence
standard structured logging where sufficient
```

A general-purpose ORM is not selected by default.

Owns learner/admin `/v1`, authoritative product DB access, durable learner/target/content/work/session/Attempt state, transaction/idempotency/concurrency boundaries, deterministic scoring/exact validation, Assessment/Progression/Planner execution over materialized canonical rules, content orchestration, async work orchestration, media-source product state, and product SSE delivery.

## Persistence/migration discipline

- one authoritative PostgreSQL-compatible product store initially;
- Core API is the only application runtime that reads/writes it;
- SQL is parameterized;
- transaction boundaries align to product invariants and required durable work/audit markers;
- decisive idempotency admission/replay association is atomic with the authoritative mutation/outcome it protects, or enforced by an equivalent serialization invariant; preflight lookup alone is insufficient;
- decisive optimistic-concurrency comparison is atomic with the guarded mutation; application-memory read/compare followed by an unrelated write is insufficient;
- learner-specific reservation/assignment state used to protect novelty/independence is serialized with its decisive eligibility/assignment decision as required by `04-application-flows.md`;
- async completion acceptance is guarded by authoritative logical-work/current execution-fencing state so duplicate or superseded result delivery cannot create a second learner outcome;
- provider/network calls are outside DB atomicity;
- connection management is bounded/observable;
- indexes/query optimization follow real measured access paths;
- migrations are explicit/ordered/versioned once schema exists;
- application/schema compatibility supports the selected rollout window rather than assuming lock-step deployment;
- migration/deploy recovery preserves committed accepted learner work.

No migration package is selected by this documentation pass.

# Python — Evaluator/media analysis

Unit: `services/evaluator/`.

Bootstrap family:

```text
Python 3
FastAPI
Pydantic-compatible typed contract models
uv
Ruff
Pyright
pytest
```

Owns bounded Writing/Speaking observations, eligible speech/transcription/acoustic extraction, bounded text/media analysis, bounded generated feedback/content candidates, optional model-assisted validation signals, and evaluator/model/generator/validator provenance/uncertainty.

Does not own public API, authoritative DB access, learner/progression/content operational state, content activation/assignment, certification, evidence sufficiency, Band advancement, or DailyPlan selection.

Celery, a separate broker framework, and a second web framework are not selected initially.

# Public/internal network baseline

```text
Browser / Next.js presentation
        ↓ public HTTP contract
Go Core API
        ↓ internal HTTP contract
Python Evaluator
```

SSE is selected for server→Web update hints under the public HTTP contract. WebSockets and gRPC are not selected while current semantics do not require them. HTTP generation, DNS, proxying, and hosting topology are deployment concerns rather than product truth.

# Contract strategy

`05-api.md` owns API semantics and directional compatibility. Implementation consequences are:

- each boundary has one exact machine authority;
- SSE remains in the public HTTP authority;
- generated bindings/validators are derived and drift-checked;
- every allowed old/new consumer-provider skew is tested rather than inferred from additive schema shape;
- public breaking change requires explicit rollout/migration/version handling;
- internal Go/Python rollout cannot assume simultaneous replacement unless deployment explicitly guarantees it;
- DB migration supports the application/schema skew window selected by deployment.

Exact contract files/version numbers are materialized later.

# Canonical registry materialization

```text
explicit canonical Markdown owner/source structure
        ↓
repository materializer / validator
        ↓
derived machine-readable registry
        ↓
generated or validated Go / TypeScript / Python consumers
```

Invariants:

1. each materialized registry knows its expected canonical owner and owner-specific canonical structure/source region; the Markdown owner remains semantic authority;
2. materialization extracts declared canonical objects, not arbitrary ID-looking tokens discovered by repository-wide grep/regex;
3. supporting prose/references, examples, `spec/DECISIONS.md`, research, archive, and other non-owner material do not become registry entries merely because they contain a matching-looking identifier;
4. duplicate, ambiguous, structurally unparseable, or unexpectedly sourced canonical definitions fail verification or remain unresolved rather than being guessed;
5. derived registry entries preserve source-owner identity and enough source fingerprint/repository revision provenance for drift/reconstruction;
6. a derived tooling source map/config may identify expected owner/source regions but is tooling configuration, not another semantic SSOT and cannot invent objects absent from the owner;
7. equivalent ID/enum registries are not manually recreated per language;
8. canonical IDs cross boundaries unchanged and are never recycled for unrelated meaning;
9. materialization may start only with registries consumed by the bounded implementation slice;
10. historical references remain reconstructable across later canonical evolution.

Exact parser, source-map representation, serialization, and codegen choices remain implementation decisions.

# Distinct state/version identities

Implementation must not overload one generic `version` concept for unrelated identities:

```text
canonical semantic evolution
API / machine-contract version
DB schema / migration version
ContentRevision
ValidationDecision / validation-policy version
provider / model version
runtime / software build version
resource optimistic-concurrency revision
idempotency operation identity
async work identity
```

They can be correlated where useful but have different ownership and change semantics. A DB migration does not imply canonical change; a new API version does not rewrite historical learner/content meaning; revalidation need not create a new ContentRevision; provider/model replacement cannot redefine canonical truth.

# Partial deployment and version skew

A rollout may temporarily contain, where the selected deployment strategy permits overlap:

```text
new Web + old Go
old Web + new Go
new Go + old Python
old Go + new Python
```

Therefore public/internal compatibility is verified for the actually allowed skew. A breaking boundary change cannot assume atomic simultaneous deployment unless that guarantee is explicit in the deployment contract. Generated bindings do not eliminate runtime skew.

Database migrations likewise support the application/schema compatibility window required by the rollout. This architecture does not select blue-green, canary, rolling, Kubernetes, or another rollout technology.

# Async/process correctness

Implement `04-application-flows.md` without assuming broker infrastructure:

```text
authoritative DB work state
+ durable work/recoverable registration where needed
+ idempotent bounded dispatch
+ fenced result reconciliation
+ SSE/resource status
```

Race-sensitive operations use transaction/conditional-write/equivalent serialization controls rather than timing or preflight assumptions. Async retries preserve execution identity/fencing where overlap is possible, and novelty-sensitive reservations are reconciled separately from actual ExposureContext. Cross-process correctness does not rely on in-memory mutexes alone. Memory use is bounded for requests, media, queues, caches, and external outputs.

# Security engineering baseline

- parameterized SQL;
- framework escaping and explicit sanitization for raw HTML;
- deliberate CSP/CORS and session-appropriate CSRF handling;
- external URLs/capabilities never grant arbitrary network authority;
- least-privilege runtime/admin/service access;
- secrets never committed/logged or exposed in browser-readable configuration;
- production sensitive cross-network transport protected according to its trust boundary;
- applicable data-at-rest protection satisfied by selected storage/deployment;
- OAuth/JWT-specific behavior only if selected;
- public-edge abuse/resilience addressed without requiring a dedicated WAF product.

Before public auth contract materialization, make the credential/session transport decision required by `05-api.md`.

# Health, observability, and incidents

Each runnable unit exposes deployment-appropriate process/readiness health without claiming downstream semantic correctness it did not check.

Privacy-safe telemetry may include request/work/execution correlation, runtime/unit, operation, duration, result/failure/reconciliation class, non-sensitive error code, and consequential software/contract/config/provider provenance. Stale/superseded async completions are observable where needed to reconstruct why they were rejected. Metrics cover relevant latency, failures, DB pressure, backlog, retries, provider use/cost, and capacity.

Before product support, incident handling can detect, contain/degrade safely, preserve accepted work, recover, verify, and record material cause/follow-up. Operational history is not canonical learning truth.

# Reliability, recovery, performance, deployment

Before applicable support promotion:

- backup/restore is verified for authoritative/retained consequential data;
- migration/deploy failure has rollback or forward-recovery procedure;
- network/provider ambiguity stays pending/unresolved until safely reconciled;
- capacity/backpressure prevents unbounded queues/concurrency;
- measurable latency/throughput/backlog/cost objectives exist for the intended release without architecture inventing numeric thresholds;
- builds are reproducible from pinned/locked dependencies;
- running software/contract/config identity is observable where consequential;
- deployment keeps environment config and secrets separated;
- rollout/recovery preserves the compatibility windows above.

Docker, Terraform/IaC, load balancing, autoscaling, multi-region, and Kubernetes remain conditional on actual deployment need.

# Feature flags

Bounded flags/kill switches may control authorized availability/degradation. A flag cannot redefine canonical learning/content/evidence truth, bypass validation/support standards, or mutate historical meaning. Materially persistent flags have an owner, safe default, and lifecycle.

# Architecture finding handling

Implementation/boundary findings follow the authority, conflict-repair, and change rules in `../CONSTITUTION.md`. Temporary audit labels are supporting review metadata only; they are never canonical Gap types, runtime enums, or another governance workflow.

# Verification fixture/data boundary

Tests and fixtures use synthetic, owned/licensed, or appropriately de-identified material suitable for the repository/test purpose. Verification artifacts are derived test inputs, never canonical or content authority.

Normal test fixtures must not contain production learner credentials, real auth/provider secrets, unnecessary production learner responses/audio, a copied production database dump, live provider credentials, or unauthorized copyrighted IELTS/media material. If official/external sample material is used, its provenance/rights must permit the intended repository/test use.

# Native verification baseline

Each deployable owns native checks while repository correctness remains one root contract.

## Web

```text
Prettier format check
ESLint
TypeScript typecheck
Vitest
React Testing Library where material
Next.js production build
Playwright critical E2E
security-sensitive browser/input/hidden-content projection tests where material
public-contract conformance/compatibility once materialized
```

## Go

```text
gofmt
go vet ./...
go test ./...
race tests where material
build core-api
DB/migration/query integration once persistence exists
atomic idempotency + optimistic-concurrency race tests where material
async duplicate/late/superseded completion reconciliation tests where material
novelty reservation/assignment race + exposure reconciliation tests where material
```

## Python

```text
Ruff format/lint
Pyright
pytest
internal-contract conformance/compatibility once materialized
bounded input/output/provider tests where material
```

## Contracts/registries

Once materialized:

```text
schema validation
generated-artifact drift
canonical owner/source extraction + ID/reference validation
consumer/provider conformance
public/internal directional compatibility for allowed version skew
null/applicability/unknown-enum behavior where material
learner/public hidden-content projection conformance
cross-unit integration
```

Boundary verification also exercises auth/access ordering, forbidden bypasses, internal evaluator reachability/auth according to actual principals/reachability, durable async recovery/fencing, DB ownership, SSE access isolation, data-lifecycle reconciliation, security-critical config failure, fixture data/rights constraints, and privacy-safe observability.

# Root verification

One root verification contract eventually spans:

```text
verify
  ├── repository/canonical + dependency/reference integrity
  ├── materialized registries + owner/source drift
  ├── materialized contracts + directional compatibility
  ├── web
  ├── core-api
  ├── evaluator
  ├── persistence/migrations
  └── cross-unit/boundary integration
```

Checks enter when the corresponding artifact/runtime exists. CI must invoke the same correctness contract rather than define a separate hidden PASS. A cross-stack change is not PASS because only one affected ecosystem is green.

# Framework/infrastructure change rule

A replacement requires architecture review when it materially changes deployable boundaries, API/transport ownership, rendering/runtime model, persistence ownership/consistency, cross-language contracts, provider/trust boundary, or operational complexity. Patch/minor maintenance inside the same responsibility boundary is implementation work.

# Initial non-goals

Do not introduce at bootstrap without demonstrated trigger:

```text
need async        → Kafka/broker
need cache        → Redis
need similarity   → vector database
need scaling      → Kubernetes
need reliability  → multi-region
need auth         → mandatory JWT/OAuth
need events       → event-source everything
need architecture → microservice-per-feature
```

Also do not add GraphQL beside REST, a Next.js product backend/BFF, duplicate Go/Python domain rules, direct Python/Next.js authoritative DB access, or frontend-owned Band/mastery/content eligibility without an explicit architecture change.