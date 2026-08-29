STATUS: CANONICAL
OWNS: initial deployable-unit allocation, primary-language/framework assignment, stable framework/library-family ownership, dependency-selection/reuse policy, bootstrap toolchain profile, runtime responsibility split, material implementation-boundary model, feature-first source organization, cross-runtime naming/test/generated-code placement discipline, system-engineering concern disposition, cross-language contract/evolution strategy, canonical-registry materialization/evolution strategy, content-authoring/released-materialization boundary, persistence/consistency engineering baseline, configuration/data-lifecycle/security/observability/performance/deployment engineering invariants, repository/native verification contract, and repository automation/CI constraints
DEPENDS_ON: ../CONSTITUTION.md, 04-application-flows.md, 05-api.md
DOES_NOT_OWN: learning/product truth, parser/materializer implementation, canonical-registry or content-authoring-package serialization details, exact dependency patch versions, cloud/provider choice, final database schema, concrete deployment topology, external-provider lifecycle/selection/ingress/egress details, evaluator model vendor, numerical SLO/timeout/retry/scaling thresholds, package-manager lock state, or CI platform configuration

# Implementation Stack

## Purpose

Assign approved languages/framework families and implementation-engineering invariants to explicit runtime responsibilities so implementation does not re-decide first-order architecture, duplicate domain logic across stacks, collapse trust boundaries for convenience, or cargo-cult infrastructure.

Architecture freezes responsibility, stable framework/library families where ambiguity would otherwise create duplicate infrastructure, material boundary/consistency/evolution semantics, and production concern semantics. Patch versions, provider products, deployment vendors, and empirical operational thresholds remain implementation/deployment decisions.

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

The evaluator is an **internal capability**, not public API surface. For the selected Cloud Run topology, authenticated internal invocation uses **Cloud Run IAM plus Google-signed OIDC identity tokens**:

```text
Go Core Cloud Run service identity
        ↓ Google-signed OIDC identity token
restricted/private Evaluator Cloud Run service
        ↓ IAM authorization
bounded evaluator capability
```

Browser and admin UI cannot call the evaluator directly. Network placement alone is insufficient authorization: `localhost`, a private IP, same VPC/host, or an `internal` hostname does not establish caller identity. Use distinct least-privilege service identities where material. Cloud Tasks delivery to protected handlers likewise uses an appropriate service identity/OIDC route.

Do not introduce a shared `INTERNAL_SECRET`, home-made service-JWT or HMAC authentication protocol, mTLS, or service mesh for this boundary. Exact Cloud Run ingress/topology details remain deployment configuration, but they may not weaken the IAM/principal authorization invariant.

Python returns bounded result/signals plus provenance/model/evaluator identity and uncertainty/quality where material. Responses remain non-authoritative until Core validates/interprets them through owning policy. Python cannot directly read/write the product DB.

## Authoritative state → async work

Implementation preserves the distinct points defined by `04-application-flows.md`:

```text
authoritative mutation
required durable logical work identity/marker
execution-attempt claim/fence
dispatch attempt
remote/capability execution
completion received
Core result reconciliation
accepted result
required downstream semantic continuation materialized or recoverable
```

Required continuation is durably registered with or recoverably derivable from committed state before acknowledgement depends on it. Observing pending work is not dispatch ownership: where duplicate execution is not deliberately allowed, one execution attempt is decisively claimed/fenced under authoritative state or an equivalent serialization invariant before dispatch. The decisive claim also enforces current execution/data-egress eligibility where cancellation, deletion/tombstone, quarantine, rights, or privacy state can prohibit dispatch. Claim ownership itself is recoverable: worker loss cannot permanently strand logical work, and an uncertain post-claim dispatch is reconciled as an ambiguous remote outcome rather than blindly redriven. Timeout does not prove remote non-execution. Retry/reconciliation reuses logical work/idempotency identity and preserves distinct execution-attempt/fencing state where overlapping executions are possible.

An accepted capability completion that requires Observation/Assessment/Progression or another semantic continuation cannot be marked as a permanently self-contained success unless the required downstream state is already committed or is deterministically/recoverably derivable from durable accepted-result state. Replay/reconstruction is idempotent.

## Cache boundary

Cache is derived. Correctness-sensitive cache identity includes relevant source identity/revision, policy/model/contract identity where material, access scope, and freshness/invalidation state.

Stale cached output cannot overwrite newer authority. Cache never becomes learner/evidence/content/product/progression/capability authority. The selected Upstash Redis route in `07-third-party-services.md` remains disposable/non-authoritative and is accessed through the approved Go Redis family below.

## Object/media storage boundary

When large artifacts are stored separately, authoritative Core state stores the reference/metadata needed to govern them; object/media storage stores bytes behind that authority. The selected Cloudflare R2 route in `07-third-party-services.md` addresses access scope, integrity, private-by-default behavior, lifecycle/recovery, provenance/rights, and orphan reconciliation without becoming product-state authority.

### Conditional direct browser byte transfer

For the selected R2 route, prefer narrow signed direct browser upload/download where the product flow permits it:

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

Cross-language durable timestamps use timezone-aware absolute instants once serialized by machine contract. Wall clock is not causal ordering, idempotency identity, optimistic-concurrency identity, or execution-claim authority. Recency/expiry names its policy/clock where material.

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

Core-owned work state progresses asynchronously while dispatch/capability execution may lag. One logical work identity remains authoritative; execution-attempt claim/fencing and downstream recovery prevent worker races or crashes from creating divergent semantic outcomes.

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

Current deletion/tombstone/retention eligibility also participates in live async result reconciliation. A late evaluator/provider/callback completion cannot resurrect deleted learner data, recreate active object/media references, create new Observation/EvidenceFact state from data no longer eligible for use, reactivate deleted/quarantined content, or repopulate active derived stores contrary to current deletion policy. Minimum historical/audit provenance may remain only when the applicable policy permits it.

Derived stores cannot silently resurrect deleted data into active product use. Restore procedures reconcile restored data against current deletion/tombstone state before normal active use. Historical semantic records needed for integrity may be retained only under an explicit applicable product/security/data policy rather than accidental persistence.

An authoritative restore also creates a reconciliation boundary for external work that may have been dispatched before the restored snapshot. A pre-restore evaluator/provider execution or callback cannot mutate restored authoritative state merely because its resource/logical IDs still match. Implementation provides an equivalent fencing/reconciliation mechanism sufficient to reject, quarantine, or safely reconcile stale pre-restore external work against the restored current authority before mutation, while preserving retained historical/audit/provenance integrity. The exact fence representation is implementation-owned.

Deletion is therefore a reconciliation fence where consequential, not merely eventual background cleanup. This architecture does not invent retention durations or legal obligations.

Raw learner audio is ephemeral by default under `07-third-party-services.md`. When product purpose explicitly retains audio—for example learner replay/history or evidence/audit needs—the authoritative product state must identify that retained artifact/lifecycle sufficiently for access, export where applicable, deletion, provider cleanup, backup/restore reconciliation, and late-result fencing. Temporary capture/upload/processor copies are not silently promoted into permanent learner history.

Entitlement loss is not a deletion instruction. Premium capability expiry/downgrade may stop new paid processing while retained learner data/history continues under its normal data-lifecycle policy.

# Secrets, runtime policy, and deployment configuration boundary

Keep three implementation classes distinct:

```text
A. secrets
   → Google Secret Manager

B. operator-adjustable typed runtime policy
   → Core-owned authoritative configuration
   → changeable without redeploy where appropriate

C. deployment/bootstrap environment
   → static deployment/bootstrap configuration
   → not dynamic product-policy authority
```

Secrets include provider/API credentials and other values whose disclosure grants sensitive access. Deployment/bootstrap configuration may select endpoints, runtime mode, service identity references, bootstrap secret references, and static wiring. Typed runtime policy may select, where already authorized by canonical design, provider enabled/suspended state, capability/per-user quotas, approved-route role, rate-limit profile, cost admission policy, operating financial target, variable-spend safety ceiling, and bounded capacity policy.

`.env` files and deployment environment variables never become authoritative runtime policy. Admin/BOPS exposes no generic raw environment-variable editor. Routine secret operations expose only minimum safe metadata such as configured/not-configured state, reference/version/fingerprint/status; routine Admin users cannot read/export stored plaintext secrets. Credential creation, rotation, or switching requires the applicable elevated `security-sensitive operations` capability from `04-application-flows.md` and durable audit.

Runtime policy must not redefine Band semantics, Skill/Knowledge identity, evidence meaning or Assessment authority, content semantic identity, Progression rules, Coverage meaning, or provider/model calibration eligibility. Security-critical missing/invalid bootstrap configuration fails closed. Consequential runtime-policy revisions are reconstructable where required for audit/provenance.

## Financial operating policy and concurrency-safe admission

Initial runtime defaults are:

```text
operating target                 = USD 10 / month
variable-spend safety ceiling    = USD 20 / month
```

These are operator-adjustable runtime-policy defaults, not architecture constants and not invoice guarantees. The **operating target** is an optimization/planning objective for forecasts, alerts, quota decisions, and optional-capability policy. The **variable-spend safety ceiling** is an application-level admission boundary for new discretionary cost-bearing external operations.

The safety ceiling cannot guarantee a final provider invoice because fixed infrastructure, metering/reporting delay, already admitted or in-flight work, estimation error, and provider-side billing behavior can all create charges beyond the currently configured discretionary boundary. Provider budget alerts/dashboards are observations and warning controls, not the sole concurrency boundary and not enforcement guarantees.

A new cost-bearing logical operation must not independently read a remaining-budget value and then call a provider when concurrent callers could oversubscribe the same allowance. The implementation preserves an invariant equivalent to:

```text
logical paid operation
        ↓
estimate / reserve bounded spend
        ↓
atomic or serializable admission against available discretionary budget
        ↓ admitted
execute
        ↓
reconcile reservation against observed/estimated actual usage
        ↓
release / adjust reservation
```

Exact tables/counters/algorithms remain implementation choices. A reservation ledger, transactional counter, or equivalent mechanism is sufficient only if it proves the same concurrency invariant. Cost authority belongs to the logical operation: retry, repair, fallback, escalation, replacement execution, or continuation cannot obtain fresh independent allowance that bypasses the original operation's quota/admission policy.

If the next eligible paid attempt can conservatively cost more than the logical operation's currently reserved amount, the **same logical-operation reservation** is atomically/serializably extended before the next paid dispatch/execution. Successful extension permits dispatch; failed extension permits no new paid dispatch and the work remains delayed/pending/unavailable under its owning semantics. An ambiguous potentially billable execution retains enough conservative liability/reservation until reconciled or otherwise safely settled under explicit policy. Timeout, cancellation, or client disconnect does not prove zero provider cost.

A financial-period boundary does not erase unresolved prior-period liability, release ambiguous potentially billable reservations, authorize stale optional backlog, or turn a calendar reset into fresh execution eligibility. Before delayed optional paid work executes in a later period, admission revalidates where material its current usefulness/freshness, product eligibility, provider-route status, quota, financial policy, applicable price basis, and unresolved existing liability. A reservation spanning a period boundary remains one logical financial obligation and is reconciled once under its owning model; it is not duplicated merely because the calendar changed.

Where conservative cost estimation is material, the admission decision is reconstructable enough to identify as applicable the provider, capability/model/route identity, price/rate revision or equivalent policy identity, effective/verified time, budget currency, and estimation basis. Live provider-price scraping is not required and provider dashboards remain observations. If the applicable price basis is materially stale, unavailable, or uncertain, new discretionary paid work uses conservative admission or fails closed according to policy; unknown pricing never means unlimited execution. Refunds, credits, or late adjustments may reconcile later without rewriting original usage/execution history. Repeated/duplicated provider usage is reconciled idempotently, and absence of provider-reported usage does not prove zero cost.

Where material, operator cost state considers settled/observed spend, estimated unsettled spend, reserved in-flight spend, short-window burn rate, and projected month-end spend. If current cost-policy state cannot be established safely, deterministic zero/near-zero-cost core paths plus durable learner history/state, payment reconciliation, and integrity paths remain available where their own dependencies are healthy; **new discretionary paid external work fails closed, remains pending/delayed, or is truthfully unavailable according to product semantics**. Uncertainty never means unlimited paid execution.

Budget policy can represent states equivalent to `NORMAL`, `WARN`, `OVER_TARGET`, `SAFETY`, and `EXHAUSTED`; exact numeric transitions are typed runtime policy, not architecture thresholds. `NORMAL` admits optional work under ordinary policy, `WARN` increases operator visibility, `OVER_TARGET` may tighten expensive optional quotas/prefer already eligible cheaper or batch routes/delay optional work, `SAFETY` denies or suspends selected expensive discretionary work, and `EXHAUSTED` denies new discretionary external spend while keeping critical integrity-preserving paths alive.

Cost pressure never lowers evaluator/evidence/privacy/security quality, silently substitutes an uncalibrated evaluator, fabricates learner weakness/Band/readiness evidence, or discards durable learner state. If a consequential evaluator cannot run within admitted policy, honest states such as `PENDING`, `TEMPORARILY_UNAVAILABLE`, or another already-owned unresolved/degraded state are used instead of lower-quality fake completion.

## Admin / BOPS operational control

Admin/BOPS is an observation and typed-policy surface over Core; it is not direct database/provider mutation authority. Where available and material it can expose current-period API traffic, request/error/latency metrics, provider health and capability usage, AI tokens/calls/audio minutes, settled/estimated-unsettled/reserved spend, operating target, safety ceiling, quota consumption, queue/backlog state, storage growth, month-end projection, cost by capability, and useful unit-economic measures when product data exists. Telemetry/business dashboards remain non-authoritative observations.

A normal `ADMIN` may change routine approved typed operational policy only through its granted `operational policy administration` capability and within configured policy bounds. Raising a safety-critical financial ceiling requires the elevated `security-sensitive operations` capability (normally `OWNER` by default bundle). Consequential policy mutation records actor, timestamp, old revision/value, new revision/value, and reason where required in durable non-rewritable audit appropriate to the operation. Exact BOPS navigation, tables, and dashboard layout are not frozen.

# System-engineering concern disposition

Review classes are:

- **BOOTSTRAP_REQUIRED** — invariant resolved when the relevant runtime capability first exists;
- **PRODUCTION_GATE** — may be incremental but must pass before applicable product support;
- **CONDITIONAL** — required only after its trigger exists;
- **NOT_SELECTED_INITIAL** — named technology/pattern intentionally not selected without demonstrated need.

## Bootstrap-required concerns

When applicable:

- bounded deadlines/input/concurrency/backpressure and safe retry/idempotency;
- transaction ownership, commit-before-success, durable async recoverability, decisive assignment eligibility under concurrency, exclusive/recoverable execution-attempt claiming where required, accepted-result/downstream-continuation recovery, stale-write/ambiguous-outcome handling;
- parameterized SQL, bounded DB connections, migrations/schema compatibility when schema exists;
- auth/access boundaries, least privilege, secrets, production transport protection, browser/session controls, safe external-boundary enforcement;
- structured logs/correlation, health, running build/contract identity, dependency locking/reproducibility;
- contract compatibility/conformance and canonical-registry drift verification once materialized;
- security-critical configuration validation and data-lifecycle/deletion-fence enforcement for implemented stores/work.

## Production-gate concerns

Before applicable `COVERED`/`SUPPORTED_FOR_PRODUCT` promotion, the release candidate resolves and verifies as applicable backup/restore, deploy/migration recovery, failure/degraded behavior, monitoring/alerts/audit, security verification, capacity/backlog/external-cost visibility, and measurable operational objectives appropriate to the release. Architecture does not invent numeric SLO/RPO/RTO values.

## Conditional / initial non-selection

Reverse proxies/gateways/load balancers beyond the selected hosting/edge route, read replicas/sharding, distributed locks/transactions/Sagas, additional brokers/PubSub/DLQ beyond selected Cloud Tasks, WebSockets, gRPC, circuit breakers beyond bounded retry/degradation, multi-region, chaos engineering, dedicated paid WAF products, Terraform/IaC, Kubernetes/Helm/service discovery, and external feature-flag services remain trigger-based.

Local PostgreSQL launched through Compose is current development/integration-test implementation packaging. It is not production topology authority and does not select containerized production deployment.

Kubernetes, Helm, Kafka-class broker infrastructure, distributed transactions/Sagas, multi-region, DB sharding/read replicas, leader election, gRPC, and WebSockets are not selected initially.

# Reuse-first implementation invariant

Prefer, in order:

1. the standard library;
2. a framework-native facility;
3. a dependency already selected in this repository;
4. a mature focused open-source library;
5. thin project-specific glue;
6. custom generic infrastructure only as a last resort.

For a new general-purpose dependency or custom generic subsystem, implementation must establish that the concern exists, the selected stack does not already own it, the proposal is maintained and compatible, its license is acceptable, it does not require unapproved paid SaaS, it does not create a second semantic or machine truth, and its dependency cost is justified by the custom complexity it removes.

Optimization cannot move semantic authority into caches, generated files, prompts, provider output, logs, metrics, migrations, configuration, or DB schema. Cost pressure cannot lower evidence/content/evaluator quality for the same intended consequence.

Domain-specific implementation should remain custom where it encodes product semantics. This includes `TargetProfile`, `ContentRevision`, `ValidationDecision`, Attempt lifecycle, Observation, Assessment, `EvidenceRequirement`, Progression, Planner behavior, content eligibility, and IELTS scoring/inference. The anti-wheel rules below apply to commodity engineering infrastructure, not product semantics.

# Version and dependency policy

Architecture freezes compatibility families/responsibilities, not patch numbers.

At bootstrap:

- use a currently supported Node.js LTS line and compatible TypeScript/React/Next.js App Router;
- use a currently supported Go release with the approved Go family below;
- when the evaluator runtime is materialized, use a currently supported Python 3 release with the approved Python family below;
- pin exact dependency/runtime versions in manifests/lockfiles/tool scripts;
- maintain verified security/maintenance updates;
- reuse the selected owner for a concern instead of adding an overlapping library or silently changing tooling family.

Dependency updates cannot silently change machine-contract or canonical semantics. A real new concern may justify a new dependency when the existing selected stack is insufficient and the reuse-first checks above pass.

# TypeScript — Web

Unit: `apps/web/`.

Approved family ownership:

```text
runtime/framework        TypeScript strict + React 19 family + Next.js App Router on an Active LTS/currently supported security release
package/lock management    pnpm
public HTTP contract       exact OpenAPI contract
HTTP client/bindings       Hey API generated from OpenAPI
server state               TanStack Query where client-side server-state management is needed
i18n                       next-intl
styling                    Tailwind CSS v4 family
UI primitives              shadcn/ui source-distribution model; Radix primitives only as actually consumed
icons                      Lucide / lucide-react
format/lint              Prettier + ESLint
component/unit tests     Vitest + React Testing Library + user-event + jest-dom
browser E2E              Playwright
automated accessibility  axe Playwright integration
```

Exact versions remain owned by `package.json` and `pnpm-lock.yaml`.

State ownership is narrow:

```text
server state                         → TanStack Query where client-side server-state management is needed
local UI / simple form state           → React state
complex form state                     → focused form library only when complexity earns it
shareable navigation / filter state    → URL / Next.js router where appropriate
```

A new global state library requires a demonstrated state-ownership problem that these mechanisms cannot represent cleanly. Do not create a generic project state framework.

Implementation must not independently recreate query/cache lifecycle, mutation invalidation, translation plumbing, accessible primitive behavior, or HTTP DTO truth when the selected stack already owns those concerns. Form libraries are conditional on demonstrated form complexity; do not preselect one for simple forms. Do not add generic framework wrappers such as `QueryManager`, `FormEngine`, `TranslationEngine`, `UIComponentFactory`, or `HTTPClientFramework` merely to make selected libraries swappable. Thin domain/presentation adapters remain allowed.

MSW is trigger-based rather than always-installed: add it only when isolated component/integration HTTP mocking becomes useful. Redux, Zustand, Axios, a generic custom cache, a generic custom form framework, and a custom i18n framework are not selected while the ownership model above is sufficient.

Owns learner/admin rendering, interactive workspaces, browser capture, embedded-player interaction, transient timers/drafts/optimistic presentation, SSE client behavior, presentation transformations, accessibility/responsiveness, and presentation-side Next.js server rendering/web-edge mechanics.

Does not own learning/evidence/progression policy, deterministic IELTS scoring, content eligibility, evaluator algorithms, durable learner/product truth, authoritative DB access, or independent DTO truth.

## Browser audio/capture implementation boundary

Web implements the product capture semantics from `01-skill-features.md` without inventing ability conclusions. When audio is required:

- microphone permission denial, unavailable devices, capture initialization failure, device removal/change, recorder/browser interruption, clipping/inaudibility/noise signals, truncation, and upload/network failure remain distinguishable operational conditions where material;
- the UI must not mark a recording successfully submitted/evidence-usable before the required bytes/metadata are durably accepted by Core;
- local replay/re-record is presentation behavior permitted by the activity; it does not mutate a submitted Attempt;
- audio capture and STT are separate capabilities. A successfully captured recording remains the primary captured learner signal for acoustic use even when transcription fails; a transcript never silently replaces missing/unusable audio for an acoustic claim;
- material disagreement between transcript and audio is preserved as evaluator/input uncertainty rather than silently trusting whichever result is easier to process;
- browser/tab/network interruption preserves enough local/product state to retry safely where possible and never fabricates uninterrupted timing/capture;
- accessibility-compatible controls and supported alternate capture/input routes preserve keyboard/screen-reader/focus usability and surface any condition material to readiness inference; they do not redefine IELTS/Band truth;
- capture feasibility is checked before evidence-critical activity admission when practical, while an unexpected mid-attempt failure still produces an honest unusable/partial outcome rather than learner failure.

Exact browser APIs, codec/container choices, audio thresholds, retry timings, and device-quality cutoffs remain implementation/empirical decisions.

# Go — Core API + deterministic orchestration

Unit: `services/core-api/`.

Approved family ownership:

```text
runtime                  Go
HTTP                     Go standard-library net/http + chi/v5
PostgreSQL access        pgx/v5
typed SQL generation     sqlc
migrations               goose with SQL-first migrations
public OpenAPI server    oapi-codegen
Redis client             go-redis/v9
rate-limit helper        redis_rate or equivalent narrow maintained helper only when it removes justified complexity
R2 object client         AWS SDK for Go v2 S3 client against Cloudflare R2
async dispatch           official Google Cloud Tasks SDK
payments                 official payOS Go SDK where suitable for the selected payOS route
logging                  log/slog
security / IDs           Go standard-library crypto primitives where sufficient
tests                    testing + httptest + real PostgreSQL integration tests
```

Exact versions remain implementation/tooling-owned by `go.mod`, `go.sum`, and repository tool scripts.

SQL remains explicit project SQL. `sqlc` owns typed query generation/mapping and `pgx` owns PostgreSQL execution. Do not maintain a parallel handwritten `QueryRow`/`Scan` mapping implementation when the same query is already represented in `sqlc`. Do not introduce a generic `RepositoryBase`, generic query framework, ORM, migration engine, logging-wrapper framework, or config framework without demonstrated need. Business/domain SQL and transaction semantics remain project code.

Owns learner/admin `/v1`, authoritative product DB access, durable learner/target/content/work/session/Attempt state, transaction/idempotency/concurrency boundaries, deterministic scoring/exact validation, Assessment/Progression/Planner execution over materialized canonical rules, content orchestration, async work orchestration, media-source product state, and product SSE delivery.

## Persistence/migration discipline

- one authoritative PostgreSQL-compatible product store initially;
- Core API is the only application runtime that reads/writes it;
- SQL is parameterized;
- transaction boundaries align to product invariants and required durable work/audit markers;
- decisive idempotency admission/replay association is atomic with the authoritative mutation/outcome it protects, or enforced by an equivalent serialization invariant; preflight lookup alone is insufficient;
- decisive optimistic-concurrency comparison is atomic with the guarded mutation; application-memory read/compare followed by an unrelated write is insufficient;
- decisive assignment revalidates every mutable authoritative hard gate that can concurrently invalidate that assignment and couples the check to reservation/assignment through a transaction, conditional write, or equivalent serialization invariant; a stale-plan/current-state preflight alone is insufficient;
- learner-specific reservation/assignment state used to protect novelty/independence is serialized with its decisive eligibility/assignment decision as required by `04-application-flows.md`;
- where duplicate execution is not intentionally allowed, one execution attempt is authoritatively claimed/fenced before dispatch; the claim enforces current dispatch/data-egress eligibility where material and remains recoverable after claimant loss rather than permanently consuming logical work;
- async completion acceptance is guarded by authoritative logical-work/current execution/deletion/content-eligibility fencing so duplicate, superseded, or now-ineligible result delivery cannot create a second or resurrected learner outcome;
- accepted capability results that require downstream semantics commit the needed semantic artifacts atomically or persist enough authoritative/recoverable continuation state for deterministic idempotent reconstruction; upstream completion cannot permanently strand missing Observation/EvidenceFact/Progression state;
- provider/network calls are outside DB atomicity;
- connection management is bounded/observable;
- indexes/query optimization follow real measured access paths;
- migrations are explicit/ordered/versioned once schema exists;
- application/schema compatibility supports the selected rollout window rather than assuming lock-step deployment;
- migration/deploy recovery preserves committed accepted learner work.

Architecture does not make a migration product semantic truth. `goose` is the selected SQL-first PostgreSQL migration family; migration files remain explicit ordered SQL and exact versions remain implementation/tooling-owned.

# Python — Evaluator/media analysis

Unit: `services/evaluator/`.

The Python evaluator runtime is approved architecture but is not yet materialized in the repository. Do not create it merely to satisfy this profile; when implementation starts, use this family unless an implementation-time compatibility/security finding requires an explicit change.

Approved future family ownership:

```text
runtime/package          Python + uv
lockfile                   uv.lock once materialized
HTTP API                   FastAPI + Uvicorn
models                     Pydantic v2
bootstrap config           pydantic-settings where useful; never dynamic product-policy authority
outbound HTTP              httpx
quality                    Ruff + Pyright
tests                      pytest + pytest-asyncio
machine-contract models    OpenAPI-derived Pydantic v2 models
model generator            datamodel-code-generator
```

`respx` is conditional on real external HTTP adapter tests. `tenacity` is conditional on a real external retry policy. SQLAlchemy, Alembic, Celery, Redis, LangChain, and LlamaIndex are not selected by this profile.

Keep the runtime intentionally thin: internal HTTP adapter → bounded capability/evaluation service → provider adapters/evaluator functions. It owns bounded Writing/Speaking observations, eligible speech/transcription/acoustic extraction, bounded text/media analysis, bounded generated feedback/content candidates, optional model-assisted validation signals, and evaluator/model/generator/validator provenance/uncertainty.

Does not own public API, authoritative DB access, learner/progression/content operational state, content activation/assignment, certification, evidence sufficiency, Band advancement, or DailyPlan selection. Python must not access the authoritative product PostgreSQL store.

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
- learner-safe projections encode upstream reveal policy rather than deriving pedagogical/evidence reveal timing from nullable/internal fields;
- generated bindings/validators are derived and drift-checked;
- every allowed old/new consumer-provider skew is tested rather than inferred from additive schema shape;
- public breaking change requires explicit rollout/migration/version handling;
- internal Go/Python rollout cannot assume simultaneous replacement unless deployment explicitly guarantees it;
- DB migration supports the application/schema skew window selected by deployment.

Contract tooling families are `oapi-codegen`, Hey API, `datamodel-code-generator`, and `oasdiff` for OpenAPI compatibility/breaking-change verification. Runtime DTO/schema families are derived consumers and never become a second machine-contract authority.

The intended next materialization direction is:

```text
canonical design semantics
        ↓
exact machine contract
        ↓
generated bindings
        ↓
implementation

public OpenAPI
  → Go public server bindings/interfaces via oapi-codegen
  → TypeScript client/types/integration bindings via Hey API

internal evaluator OpenAPI
  → Go internal evaluator client/bindings via oapi-codegen
  → Python Pydantic v2 models via datamodel-code-generator
```

Exact generator flags remain implementation verification. This closure does not create `contracts/http/public.openapi.yaml` or `contracts/http/evaluator.openapi.yaml`; those exact contract authorities are the next bounded phase. The repository's existing bounded bootstrap `contracts/http/public-v1.json` and its generated artifacts are not regenerated or promoted into a second authority by this documentation pass.

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

Exact parser, source-map representation, serialization, and codegen choices remain implementation decisions except where an approved runtime family above names a generator family for a cross-language contract.

# Content authoring and released-materialization boundary

Content authoring may use repository-tracked source/packages plus deterministic tooling, including the parallel shard model owned by `04-application-flows.md`. The authority direction is:

```text
canonical spec/design IDs + content schema/policy
        ↓
authored / generated candidate source packages
        ↓
validation + deterministic integration
        ↓
released materialized ContentRevision/product artifacts
        ↓
Core-owned runtime database / governed object storage projection
```

Rules:

1. when parallel/bulk authoring is materialized, one repository-owned machine-checkable candidate-package/shard schema defines the exact authoring package shape and allowed identity fields from current canonical registries/policies; it is authoring/tooling contract, not a new learning SSOT;
2. repository-tracked authoring packages may be source inputs/provenance, but they do not redefine Skill/Knowledge/Band/Assessment truth;
3. the runtime database is authoritative for current operational/runtime state after release, not the canonical authoring source from which content semantics are casually edited backward;
4. released runtime rows/artifacts resolve exact ContentRevision identity/provenance and are reproducibly traceable to their accepted source/materialization path where material;
5. normal AI workers consume deterministic shard/package contracts and emit candidates; they do not require cross-worker coordination or direct runtime-DB mutation;
6. deterministic integration fails on overlapping shard identity, duplicate candidate ID, unresolved canonical reference, package/schema error, or non-deterministic merge ambiguity;
7. text/metadata/source packages may live in Git when appropriate for review/versioning. Large generated audio/media/binary assets may remain outside Git behind governed object/media storage with integrity identity and lifecycle/provenance references;
8. generated difficulty/Band/evidence labels remain non-authoritative candidate data until owning policy/calibration admits the consequence;
9. materialization tooling does not grant release authority: normal ValidationDecision/release/authorization rules still apply after structural integration;
10. no queue/registry/orchestrator/agent platform is introduced merely to parallelize independent deterministic shards. Existing repository tooling/processes are sufficient until measured operational need proves otherwise.

First-party AI workers are service identities under the RBAC/capability model in `04-application-flows.md`; being local, internal, or repository-adjacent does not grant filesystem, secret, provider, protected-data, release, or administrative privilege beyond explicit capability and execution scope.

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

They can be correlated where useful but have different ownership and change semantics. Plan-time provenance may reference several of these/current state revisions and assignment-time provenance may differ; neither creates a generic universal version. A DB migration does not imply canonical change; a new API version does not rewrite historical learner/content meaning; revalidation need not create a new ContentRevision or overwrite an earlier ValidationDecision; provider/model replacement cannot redefine canonical truth.

# Partial deployment and version skew

A rollout may temporarily contain, where the selected deployment strategy permits overlap:

```text
new Web + old Go
old Web + new Go
new Go + old Python
old Go + new Python
```

Therefore public/internal compatibility is verified for the actually allowed skew. A breaking boundary change cannot assume atomic simultaneous deployment unless that guarantee is explicit in the deployment contract. Generated bindings do not eliminate runtime skew.

Database migrations likewise support the application/schema compatibility window required by the rollout. Capacity verification for rollout/recovery also accounts for applicable overlapping downstream consumers, including old and new revisions, worker/evaluator instances, migrations/admin operations, and recovery/maintenance headroom. Steady-state configuration alone is insufficient when overlap can temporarily create greater DB/queue/provider pressure. Exact formulas and rollout technology remain deployment/verification policy; this architecture does not select blue-green, canary, rolling, Kubernetes, or another rollout technology.

# Async/process correctness

Implement `04-application-flows.md` without assuming broker infrastructure:

```text
authoritative DB logical-work state
+ durable work/recoverable registration where needed
+ decisive current-eligible execution-attempt claim/fence
+ bounded dispatch
+ fenced result reconciliation against current eligibility/deletion state
+ durable accepted-result/downstream continuation recovery
+ SSE/resource status
```

Race-sensitive operations use transaction/conditional-write/equivalent serialization controls rather than timing or preflight assumptions. Assignment-time mutable hard gates are protected with the decisive assignment rather than checked in an unrelated earlier read. Multiple workers observing pending state cannot silently double-own one exclusive execution attempt; claimant loss is recoverable, and a possibly dispatched remote execution remains ambiguous until reconciled safely rather than being blindly redriven. Intentional replacement/speculative executions use distinct identities/fences. Async retries preserve execution identity/fencing where overlap is possible, novelty-sensitive reservations are reconciled separately from actual ExposureContext, and stale DailyPlan references are rechecked against current eligibility before assignment. Cross-process correctness does not rely on in-memory mutexes alone. Memory use is bounded for requests, media, queues, caches, and external outputs.

# Security engineering baseline

- parameterized SQL;
- framework escaping and explicit sanitization for raw HTML;
- deliberate CSP/CORS and session-appropriate CSRF handling;
- external URLs/capabilities never grant arbitrary network authority;
- least-privilege runtime/admin/service access;
- secrets never committed/logged or exposed in browser-readable configuration;
- production sensitive cross-network transport protected according to its trust boundary;
- applicable data-at-rest protection satisfied by selected storage/deployment;
- public Bearer/JWT behavior follows the selected Clerk transport in `05-api.md`; custom auth/session protocols are not introduced;
- public-edge abuse/resilience addressed without requiring a dedicated WAF product.

Public contract materialization must encode the selected Clerk bearer transport from `05-api.md` rather than inventing another credential/session mechanism.

# Health, observability, and incidents

Each runnable unit exposes deployment-appropriate process/readiness health without claiming downstream semantic correctness it did not check. Selected implementation families are Go `slog` for structured application logging, OpenTelemetry for traces/metrics, and Google Cloud Logging / Monitoring as the initial operational backend.

Privacy-safe telemetry may include request/work/execution correlation, execution-claim/dispatch state, runtime/unit, operation, duration, result/failure/reconciliation class, downstream-recovery state, non-sensitive error code, and consequential software/contract/config/provider provenance. Do not log by default auth tokens, API/provider secrets, raw learner audio, unnecessary full essays/responses, provider credentials, or hidden model reasoning.

A consequential operation should be traceable across material boundaries using privacy-safe correlation/provenance, for example browser → Go request → DB transaction → durable async work → Cloud Tasks → Python evaluator → external provider → Core reconciliation → learner resource/SSE update. Metrics cover relevant latency, failures, DB pressure, backlog, retries, provider use/cost, and capacity. Telemetry remains operational evidence, not product-state authority.

Before product support, incident handling can detect, contain/degrade safely, preserve accepted work, recover, verify, and record material cause/follow-up. Operational history is not canonical learning truth.

# Reliability, recovery, performance, deployment

Before applicable support promotion:

- backup/restore is verified for authoritative/retained consequential data;
- where provider-account/control-plane loss or suspension is inside the supported recovery envelope, credible PostgreSQL restore/exit evidence does not depend exclusively on continued access to that failed/suspended provider control plane; the smallest sufficient PostgreSQL-native export/backup/restore mechanism may satisfy this, while exact storage target/schedule remain operations policy;
- migration/deploy failure has rollback or forward-recovery procedure;
- network/provider ambiguity stays pending/unresolved until safely reconciled;
- execution claiming prevents accidental duplicate dispatch when exclusive execution is required and claimant failure remains recoverable without unsafe blind redrive;
- accepted-result recovery prevents crashes from permanently stranding required downstream semantic state;
- capacity/backpressure prevents unbounded queues/concurrency;
- measurable latency/throughput/backlog/cost objectives exist for the intended release without architecture inventing numeric thresholds;
- builds are reproducible from pinned/locked dependencies;
- running software/contract/config identity is observable where consequential;
- deployment keeps bootstrap environment and Secret Manager references separated from typed runtime policy;
- rollout/recovery preserves the compatibility windows above.

Scale changes capacity/policy, not semantic ownership. Web, Go, and Python are stateless where their owned semantics permit and may horizontally autoscale on Cloud Run, but each deployable has bounded instance capacity and request/work concurrency sized to measured downstream capacity. Autoscaling never means unbounded DB, queue, AI, or provider concurrency; exact max-instance/concurrency numbers remain deployment/load-test policy.

Under capacity exhaustion, the system sacrifices latency and optional capability before correctness, evidence integrity, privacy, security, or durable learner/product state. The conceptual degradation priority is: shed/defer decorative or nonessential work; delay optional generation/transformation; throttle optional realtime; throttle expensive productive evaluation; queue eligible async work only within bounded backlog policy; deny new optional paid work when admission cannot be established; and preserve healthy critical paths such as authentication/authorization, payment reconciliation, learner history, authoritative learner/product state, accepted-work integrity, and truthful reads where their required dependencies remain healthy. This is priority under pressure, not a claim that a critical path survives a failed dependency it requires.

Upstream capacity is bounded by narrower downstream capacity rather than allowed to overwhelm it. Relational invariants are equivalent to:

```text
Web admitted request load    <= sustainable Core capacity
Core admitted concurrency    <= safe PostgreSQL + queue/dispatch + provider downstream capacity
Evaluator execution          <= provider quota/concurrency + financial admission + evaluator compute capacity
Cloud Tasks dispatch         <= sustainable evaluator/provider execution rate
```

These relations are canonical; instance counts, pool sizes, concurrency, RPS, quotas, and dispatch numbers remain deployment/load-test/runtime policy.

Scale decisions follow measured evidence: **observe → identify the measured bottleneck → remove obvious query/code/contention inefficiency → tune bounded concurrency/pooling/admission → increase vertical capacity where appropriate → increase horizontal capacity inside downstream bounds → change topology only when measured evidence shows existing capacity knobs are insufficient**. Connection-pool growth is not a default performance strategy.

For PostgreSQL specifically: observe actual query/lock/connection/compute pressure → improve query/index/schema behavior → increase compute when justified → adjust the connection budget only when justified → add read replicas only for proven stale-tolerant read pressure → partition/shard only when materially demonstrated later.

Operators treat likely bottleneck classes as monitoring hypotheses, not guaranteed first-failure points. Economic classes include realtime AI/audio, productive evaluation volume, database active duty cycle, Redis command density, and later ordinary HTTP volume. Technical classes include PostgreSQL connection/compute pressure, provider/AI concurrency, async backlog age, media/upload bursts, and runtime CPU/memory/request capacity. Observable triggers, not speculative numeric bottlenecks, drive scaling.

PostgreSQL remains authoritative durable product truth with bounded pooled application access and a bounded global connection capacity across autoscaled compute. Pool size is not a generic performance knob: observe/query/index/schema tuning precedes or accompanies justified vertical compute increase. Prefer vertical scaling before partitioning/sharding. Read replicas are not selected initially; if later justified, only explicitly stale-tolerant reads may use them where replica lag cannot violate product semantics.

Redis remains disposable/non-authoritative. Redis loss or staleness cannot lose/redefine learner progress, Attempts, payment, entitlement, evidence, or history.

Cloud Tasks preserves the already-owned durable-work/claim/fencing/idempotency/ambiguous-outcome model. Dispatch rate/concurrency and external AI/provider concurrency are bounded independently; backlog is observable; optional expensive work may be throttled or suspended before it creates unbounded provider spend. No physical four-queue topology is frozen—separate queues/classes require demonstrated independent admission/prioritization need.

R2/media keeps the existing direct-transfer boundary: Core temporary authorization → narrow signed browser byte transfer → completion/integrity reconciliation → normal policy before usability. Existing lifecycle/orphan cleanup remains required; arbitrary upload-size/expiry constants are deployment/runtime policy.

Terraform/IaC, multi-region, Kubernetes, service mesh, read replicas, partitioning/sharding, and additional broker topology remain trigger-based. Local Compose use for repository verification does not select those production technologies.

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
React Testing Library / user-event / jest-dom where material
Next.js production build
Playwright critical E2E
axe Playwright automated accessibility checks where applicable
security-sensitive browser/input/hidden-content projection + reveal-policy tests where material
stale-plan/current-eligibility user-flow tests where material
public-contract conformance/compatibility once materialized
```

## Go

```text
gofmt
go vet ./...
go test ./...
testing + httptest
race tests where material
build core-api
real PostgreSQL DB/migration/query integration once persistence exists
atomic idempotency + optimistic-concurrency race tests where material
stale-plan/current-eligibility assignment + concurrent quarantine/revocation race tests where material
exclusive execution-attempt claim / duplicate-dispatch + claimant-crash/ambiguous-dispatch recovery tests where material
pre-dispatch deletion/quarantine/rights eligibility prevents prohibited execution/egress where material
async duplicate/late/superseded/deletion-fenced completion reconciliation tests where material
crash after accepted completion → idempotent downstream continuation recovery tests where material
novelty reservation/assignment race + exposure reconciliation tests where material
content quarantine before/after submission + affected evidence consequence tests where material
ValidationDecision scope/policy selection, revalidation-history, and content-validation burden tests where material
duplicate downstream semantic replay tests where material
```

## Python

Once the evaluator exists:

```text
Ruff format/lint
Pyright
pytest + pytest-asyncio
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
learner/public hidden-content projection/reveal conformance
cross-unit integration
```

Boundary verification also exercises auth/access ordering, forbidden bypasses, internal evaluator reachability/auth according to actual principals/reachability, durable async registration/current-eligible claiming/claim-loss recovery/fencing, decisive assignment under concurrent eligibility change, content-incident in-flight consequences, ValidationDecision history/scope plus universal/consequence-specific content validation, DB ownership, SSE access/isolation/non-completion semantics, deletion/tombstone reconciliation before dispatch and against restore/late callbacks, security-critical config failure, fixture data/rights constraints, and privacy-safe observability.

# Root verification

Repository-native local verification is the correctness procedure. The current root command is:

```text
./verify
```

The root contract spans the checks whose artifacts/runtimes currently exist and grows when a corresponding runtime or boundary is materialized:

```text
./verify
  ├── repository/canonical + dependency/reference integrity
  ├── materialized registries + owner/source drift
  ├── materialized contracts + directional compatibility
  ├── web
  ├── core-api
  ├── evaluator, once materialized
  ├── persistence/migrations
  └── cross-unit/boundary integration
```

CI may invoke the same repository-native command. CI must not define a second correctness procedure, auto-invent or auto-commit generated semantic changes, be required to materialize missing generated truth, silently mutate source or `main` as part of verification, or become the only place correctness can be executed. Generated artifacts must be reproducible locally and committed coherently with the source change that generates them. A cross-stack change is not PASS because only one affected ecosystem is green.

Repository automation should use read-only repository permissions for verification unless an independently justified write operation is explicitly approved. Without explicit USER approval, automation must not introduce or require GitHub capabilities that may incur usage-based charges or metered storage/compute beyond the explicitly approved free/local path. This includes GitHub larger runners, Codespaces, paid/private Actions usage, artifact storage that can create metered billing exposure, paid GitHub security/add-on products, paid hosted build/deployment features, or another usage-billed GitHub capability. Standard free public-repository workflow execution may remain optional automation when it invokes the same local verification procedure and does not create a paid dependency. Do not encode an assumption that any GitHub feature is permanently free; paid/metered capability requires explicit USER authorization.

# Framework/infrastructure change rule

A replacement requires architecture review when it materially changes deployable boundaries, API/transport ownership, rendering/runtime model, persistence ownership/consistency, cross-language contracts, provider/trust boundary, operational complexity, or an approved concern→library/tool family. Patch/minor maintenance inside the same responsibility boundary is implementation work.

# Deliberate initial non-selections

The selected narrow stack is not permission to add overlapping infrastructure. Without a demonstrated trigger, do not introduce Prisma, GORM, Ent, GraphQL, tRPC, gRPC, Redux, Zustand, Kafka, RabbitMQ, Kubernetes, service mesh, custom authentication/session machinery, a custom component system, LangChain, LlamaIndex, CrewAI, a vector database, Elasticsearch, a generic AI routing/orchestration framework, a custom job framework, or a custom object-storage protocol.

Likewise do not add a Next.js product backend/BFF, duplicate Go/Python domain rules, direct Python/Next.js authoritative DB access, event-source-everything, microservice-per-feature, or frontend-owned Band/mastery/content eligibility. These choices are not forbidden forever; they require demonstrated need and normal architecture review rather than speculative bootstrap complexity.

# Feature-first source organization and naming discipline

This section owns application source organization, dependency direction inside deployables, shared-code promotion, language-specific naming, generated-code placement, and test placement. It refines the responsibility-first repository topology in `../CONSTITUTION.md` without imposing identical folder trees across languages.

## Master structural principle

Organize application code primarily by product feature or capability ownership so ordinary changes remain local.

Framework, transport, generated, and genuinely cross-feature infrastructure folders exist only where that framework or boundary actually owns the code. Inside a feature/capability, introduce additional layers only when dependencies, invariants, testability, or change boundaries require them.

Do not create empty architecture layers to satisfy a template. A one-package capability, a two-file feature, or a small route composition is legal when it is cohesive.

The shared mental model is:

```text
transport / framework shell
        ↓
feature or capability owner
        ↓
owned application / domain logic
        ↓
required infrastructure adapters
```

Generated machine-contract code remains a separate derived boundary. Feature-first is not feature-silo: database pools, HTTP clients, query-cache infrastructure, logging, design primitives, translations infrastructure, and provider SDK frameworks remain with their genuine cross-feature owner rather than being recreated per feature.

## TypeScript / Web

Preferred direction for new Web code:

```text
apps/web/src/
├── app/
├── features/
├── components/
│   └── ui/
├── i18n/
├── lib/
└── generated/
```

Ownership:

- `app/` owns Next.js routing, layouts, composition, and framework boundary behavior. Route files should not accumulate substantial product-feature state machines merely because they are entrypoints.
- `features/` owns product-facing feature implementation and local change behavior.
- `components/ui/` owns genuine reusable UI primitives, including selected shadcn materialization.
- `i18n/` owns cross-feature translation plumbing/catalog integration.
- `lib/` owns only precisely named, genuinely cross-cutting implementation glue. It is not a default home for code that lacks a feature owner.
- `generated/` owns generated contract/types only and is never hand-edited.

A feature may contain folders/files such as `components/`, `hooks/`, `queries.ts`, `mutations.ts`, `model.ts`, and `index.ts`, but none is mandatory. A small feature may contain only `form.tsx` and `api.ts`. Do not pre-create empty `components/`, `hooks/`, `services/`, `models/`, `types/`, or `utils/` merely for symmetry.

Feature code may depend on generated contracts, approved shared primitives, approved cross-cutting infrastructure, and framework/library owners. Feature A must not import Feature B's private implementation. Cross-feature collaboration uses one of:

1. an intentionally narrow exported feature surface;
2. a higher-level composition/orchestration owner;
3. a promoted shared semantic only after genuine shared ownership is established.

Do not solve cross-feature imports by moving arbitrary code into `shared/`.

### TypeScript naming

Use project/framework-idiomatic TypeScript/React naming:

```text
folders/files                  kebab-case where framework convention permits
React components/types/classes PascalCase
variables/functions/hooks      camelCase
hooks                          useXxx
```

Use normal semantic `camelCase` for ordinary constants. Use a special exported-constant convention only when it improves an established API/module convention; do not apply screaming constants mechanically.

Canonical concept names remain stable: `TargetProfile`, `ContentRevision`, `ValidationDecision`, `PracticeActivity`, `Attempt`, `Observation`, `EvidenceFact`, and `EvidenceRequirement` are not casually renamed to local synonyms. Generated contract field names and serialized wire identity remain contract-owned.

## Go / Core API

Preferred direction as the Core API grows:

```text
services/core-api/
├── cmd/core-api/
├── internal/
│   ├── httpapi/
│   ├── session/
│   ├── targetprofile/
│   ├── practice/
│   ├── attempt/
│   ├── content/
│   └── db/
│       └── sqlc/
└── migrations/
```

This is an ownership example, not a mandatory folder inventory. Create a capability package only when separation has real value. Do not create a package per table, package per endpoint, or visual-symmetry split.

Preferred dependency direction is:

```text
httpapi / transport
        ↓
feature / application package
        ↓
owned deterministic logic
        ↓
required persistence / provider adapter
```

`sqlc` output is infrastructure, not domain authority. HTTP handlers own transport decoding/encoding, auth/access invocation, and response mapping; they do not become the permanent home of reusable business policy merely because they call it. Feature/domain logic must not depend on `http.Request`, `http.ResponseWriter`, or generated transport objects when the logic itself is transport-independent.

Do not require every feature to contain `domain/`, `application/`, `service/`, `usecase/`, `ports/`, `adapters/`, `repository/`, or `controller/`. Introduce a boundary only when a real dependency/invariant/change boundary exists.

Current concentrated `internal/httpapi/` code may remain as bounded bootstrap while one small implementation slice is cohesive. Extract a feature/capability package before that concentration becomes an ownership trap, including when any of these triggers occurs:

- a second materially independent product capability adds substantial policy to the same transport package;
- policy must be reused by a non-HTTP caller or orchestration path;
- an invariant is materially easier to test without HTTP objects;
- handler files begin coordinating multiple capability policies rather than transport concerns;
- dependency cycles or broad imports appear because ownership is unclear.

When extraction occurs, move coherent policy/invariants, not files merely to match the example tree.

### Go naming

Use idiomatic Go:

```text
package names          short lowercase semantic words; no underscores
exported identifiers   PascalCase
unexported identifiers camelCase
```

Prefer `targetprofile.Service` when `Service` is clear in package context instead of redundant forms such as `targetprofile.TargetProfileService`. Preserve the full canonical semantic name when shortening would create ambiguity. Avoid Java-style names such as `TargetProfileManagerImpl`, `ITargetProfileRepository`, and `AbstractAttemptService`.

SQL/sqlc operation names express semantics such as `GetAttempt`, `LockAttemptForSubmission`, and `InsertObservation`, not placeholders such as `Query1`, `GetData`, or `DoUpdate`.

## Python / future Evaluator

Do not materialize the Python runtime merely to satisfy this structure. When `services/evaluator/` is implemented, prefer:

```text
services/evaluator/
└── src/evaluator/
    ├── api/
    ├── capabilities/
    │   ├── writing_evaluation/
    │   ├── speaking_evaluation/
    │   └── speech_analysis/
    ├── providers/
    ├── generated/
    └── config.py
```

Ownership:

- `api/` owns FastAPI transport only;
- `capabilities/` owns evaluator-specific feature/capability implementations;
- `providers/` owns external model/service/tool adapters only;
- `generated/` owns exact machine-contract generated models and is never hand-edited;
- `config.py` owns bounded implementation configuration, not product semantics.

Do not create these folders before the runtime or a real capability needs them.

### Python naming

Use idiomatic Python:

```text
packages/modules       snake_case
classes/types          PascalCase
functions/variables    snake_case
module constants       UPPER_SNAKE_CASE only for real module-level constants
```

Examples include `target_profile.py`, `writing_evaluation/`, `TargetProfile`, and `evaluate_submission()`. Generated aliases/wire names remain contract-owned.

## Shared semantic naming across runtimes

The rule is:

```text
SEMANTIC NAME IS SHARED.
LANGUAGE SYNTAX IS IDIOMATIC.
WIRE IDENTITY IS EXACT.
```

For example, canonical `TargetProfile` normally appears as:

```text
TypeScript  folder target-profile/   type TargetProfile   value targetProfile
Go          package targetprofile    type TargetProfile   value targetProfile
Python      module target_profile    class TargetProfile  value target_profile
HTTP/JSON   exact names from the machine contract
```

Do not normalize all languages to one casing scheme. Do not invent a local synonym for an existing canonical/domain concept unless the canonical vocabulary itself changes.

OpenAPI/JSON/event field identity is owned by the exact machine contract. Generated bindings may expose language-idiomatic identifiers while serialization remains exact. Do not maintain handwritten duplicate DTO naming truth.

## Shared-code promotion and generic-name review

Default to keeping code with its owning feature/capability. Promote code to a cross-feature location only when:

1. at least two real consumers exist;
2. their semantics are genuinely the same;
3. ownership can be stated precisely;
4. promotion does not create a generic abstraction with no stable meaning.

Prefer semantic cross-cutting names such as `api`, `i18n`, `query`, `security`, and `generated` over generic `common`, `helpers`, or `utils`.

The following names are review smells rather than automatic bans:

```text
BaseService
BaseRepository
BaseController
AbstractManager
GenericRepository
Manager
Engine
Framework
CommonUtils
SharedHelpers
Helper
Utils
Mapper
Registry
```

An abstraction with such a name must still identify its precise semantic owner and the real boundary it removes. If it cannot, keep the code feature-local or use the selected mature library.

At protected application roots, do not introduce generic dumping-ground directories such as `shared/`, `common/`, `base/`, `helpers/`, `utils/`, `services/`, or `managers/` merely to make imports convenient.

## File/package splitting

Split by reason to change, semantic ownership, dependency boundary, or testability/invariant boundary, not arbitrary line count. Do not create a 20-file feature because a style template says so, and do not retain unrelated responsibilities in one huge file merely to avoid folders.

The feature/package dependency graph remains acyclic. Do not break a cycle by moving unrelated code into `common`, `shared`, `base`, or `core`. Resolve the actual semantic owner or move orchestration to the correct higher-level owner.

## Database naming

PostgreSQL naming should remain consistent and unsurprising:

- `snake_case` tables and columns;
- plural table names only if the existing schema consistently uses that convention;
- stable explicit constraint/index names where operationally useful;
- semantic SQL/sqlc operation names.

Do not cosmetically rename the current schema merely to satisfy a style preference. Database names are storage implementation, not canonical semantic authority.

## Tests and generated code

Prefer tests near the implementation they own when the language/framework supports it:

- TypeScript feature/component tests near the feature/component;
- Go package tests in the corresponding package as `*_test.go`;
- future Python tests near the capability/module or in a compact pytest structure that preserves ownership;
- browser/DB/cross-runtime acceptance tests at the relevant integration/E2E boundary.

Do not centralize all tests into one generic tree when ownership is clearer through colocation.

Generated artifacts live in explicit generated locations. Never mix hand-authored business logic into generated files and never hand-edit generated outputs. The flow remains:

```text
OpenAPI / SQL / canonical owner
        ↓
generator / materializer
        ↓
explicit generated location
```

Root verification owns deterministic drift detection.

## Structural enforcement and scale test

Use lightweight deterministic verification only where false positives are low. Suitable checks include generated drift, protected top-level responsibility roots, forbidden language-silo roots, existing browser→Go/Python→DB boundary tests, read-only verification automation, and obvious new generic dumping-ground directories at protected application roots. Do not build a giant architecture linter or enforce subjective ownership through brittle grep.

Normal code review remains responsible for semantic ownership, abstraction quality, and whether a boundary has become worth extracting.

This structure must remain clear when adding General Training Reading, Listening practice, Writing evaluation, Speaking evaluation, diagnostic assessment, planner, content administration, account/identity, and full mocks. Those additions should create or extend obvious feature/capability owners while shared infrastructure remains centralized. They must not require a global manager class, giant capability switch, duplicated per-feature infrastructure stack, or identical language folder template.
