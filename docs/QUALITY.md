# Quality

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Created by `TASK-0007` revision 1. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Canonical target snapshot: `phatnguyen03022001/ilets@0af34c5036bff0526e52e2c22932b49f08c23e37`.
>
> Model pin: `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07` (accepted unreleased V1 candidate).

This draft migrates only the current quality constraints needed by the bounded product/runtime design. `docs/PRODUCT.md`, `docs/BEHAVIOR.md`, `docs/DATA.md`, and `docs/INTERFACES.md` retain their neighboring migration ownership; this file does not duplicate their permission matrices, learner/evidence semantics, retention values, interface inventory, or exact contract shape. It creates no provider inventory, `EXT-*`, `CAP-*`, `DEC-*`, deployment topology, workflow, implementation, assurance result, or numeric SLO.

## QUALITY-AUTHENTICATION Authentication

The selected external identity/session route owns credential custody, authentication, and session issuance/revocation mechanics only. Core owns the stable internal actor/learner association and remains the product authority after authentication. Protected public operations use the selected short-lived bearer/session transport; Core verifies the authenticity and configured security conditions required by that transport before mapping the external principal to the stable Core actor. Auth tokens are not persisted in browser `localStorage`.

Durable learner state requires authenticated durable identity. Anonymous access exists only for product paths that their owning behavior/API semantics explicitly allow; this draft does not create durable anonymous identity or a second session system. Service-to-service authentication is distinct from learner/admin transport and must establish an authorized caller identity rather than treating network placement as authentication.

Sources: `design/05-api.md` — Auth/session transport selection and request/access ordering; `docs/INTERFACES.md` — `INTERFACES-TRUST`.

## QUALITY-AUTHORIZATION-TRUST Authorization and trust boundaries

Core owns product authorization, capability checks, entitlements, and authoritative product persistence. External identity roles/metadata, Web/browser state, Evaluator output, provider state, cache state, transport success, or internal reachability cannot grant product authority. Browser/Web input remains untrusted; Web calls the Core public API and does not bypass Core to authoritative persistence or Evaluator. Core is the only application runtime allowed to read/write authoritative product persistence.

Evaluator is a bounded internal capability. Its caller must be authenticated/authorized under the selected service trust boundary; private addressing, co-location, or an `internal` label is insufficient by itself. Evaluator cannot become a public product API, read/write Core-owned product persistence directly, or turn its result into learner/evidence/content/progression truth before Core validates current work association and owning policy. Privileged operations likewise go through Core and the behavioral capability rules already owned by `docs/BEHAVIOR.md`; this section does not restate the role/permission matrix.

Sources: `design/05-api.md`; `design/06-implementation-stack.md` — material trust boundaries; `docs/BEHAVIOR.md` — `BEHAVIOR-INVARIANTS`; `docs/INTERFACES.md` — `INTERFACES-CONTRACTS`, `INTERFACES-TRUST`.

## QUALITY-SECRETS Secrets and sensitive operations

Secrets are kept distinct from Core-owned typed runtime policy and deployment/bootstrap configuration. Credentials and other values whose disclosure grants sensitive access use the selected secret-custody boundary, least-privilege runtime access, and references/bootstrap wiring rather than becoming product policy, browser-readable configuration, source-controlled values, or routine Admin-readable plaintext.

Routine operational surfaces expose only minimum safe secret metadata when needed. Credential creation, rotation, or switching requires the already-owned elevated security-sensitive capability and durable audit appropriate to the operation. Security-critical missing or invalid bootstrap configuration fails closed. Secrets, auth tokens, provider credentials, and hidden sensitive material are not logged or returned in normal learner/admin payloads.

Sources: `design/06-implementation-stack.md` — Secrets, runtime policy, and deployment configuration boundary; Security engineering baseline; `design/05-api.md` — request/access non-disclosure behavior.

## QUALITY-PRIVACY Privacy and sensitive data

External processing receives only the minimum learner/content context necessary for an already eligible purpose. Applicable rights, privacy, consent, processor retention/reuse, deletion, and data-egress conditions must be satisfied before the data leaves the controlled runtime. Credentials remain isolated from learner/product state and browser payloads, and external outputs remain bounded observations until Core accepts them under the owning policy.

Public and internal projections expose only what the authenticated/authorized caller may observe. Another learner's resources, privileged content/security state, hidden evaluator/system instructions, auth tokens, provider secrets, raw learner audio, and unnecessary full learner responses are not exposed or logged by default. Telemetry is minimized/classified separately from learner history. Retention/deletion values and authoritative lifecycle rules remain owned by `docs/DATA.md`; current deletion/tombstone eligibility still fences async dispatch and late-result reconciliation so deleted/ineligible data cannot be resurrected by a later completion.

Sources: `design/06-implementation-stack.md` — Observability boundary and data lifecycle/deletion boundary; `design/07-third-party-services.md` — external data-egress and learner-data processor rules; `docs/DATA.md` — `DATA-RETENTION`; `docs/INTERFACES.md` — `INTERFACES-TRUST`.

## QUALITY-RETRY-RECOVERY Timeouts, retries, idempotency, and recovery

Material network/capability operations are bounded by caller deadlines and safe cancellation semantics where applicable. Failure handling distinguishes transient, permanent, and ambiguous outcomes. A timeout or lost response does not prove remote non-execution; ambiguous work remains unresolved until it can be reconciled safely rather than being blindly repeated.

Retry is bounded and uses the same logical work/idempotency identity. Where duplicate execution could create learner history, evidence, accepted content, paid work, or another authoritative effect, decisive idempotency admission/replay and compatible-payload association are atomically coupled to the protected mutation/outcome or guarded by an equivalent serialization invariant. Preflight lookup is only an optimization. Required async work identity/recoverable continuation is committed with, or deterministically recoverable from, authoritative Core state before an accepted/pending acknowledgement depends on it. Accepted capability results that require downstream semantic continuation remain recoverable and replay idempotently after crashes.

Exact timeout values, retry counts, backoff budgets, and idempotency retention periods remain implementation/operational policy until measured evidence justifies them.

Sources: `design/05-api.md` — request/response patterns, idempotent create/submission, deadline/retry/fallback/backpressure; `design/06-implementation-stack.md` — authoritative state to async work; `docs/INTERFACES.md` — `INTERFACES-ASYNC`.

## QUALITY-CONCURRENCY-FAILURE Concurrency and failure boundaries

Race-sensitive authoritative operations use a transaction, conditional write, or equivalent serialization boundary rather than an unrelated read-then-write or process-local timing assumption. Expected-revision checks, duplicate-sensitive mutation, decisive assignment eligibility, exclusive execution-attempt claim/fencing, and concurrency-safe variable-cost admission are protected at the point where the decision becomes authoritative. Worker/claimant loss remains recoverable; late, duplicate, stale, superseded, deleted, quarantined, or otherwise ineligible completion is rejected or reconciled without overwriting newer authority.

Queue/dispatcher/provider state, cache, SSE hints, telemetry, timeout, and process memory are not business-state authority. Capacity pressure and dependency failure use honest pending/delayed/degraded/unavailable/recovery behavior rather than undurable acceptance, silent quality reduction, or fabricated learner outcomes. Product/evaluator/provider failure stays a product/runtime condition and never becomes learner weakness, evidence, Band state, content-quality truth, or successful semantic completion.

Sources: `design/05-api.md` — failure classes; `design/06-implementation-stack.md` — consistency classes, async/process correctness, reliability/recovery; `docs/BEHAVIOR.md` — `BEHAVIOR-FAILURES`; `docs/DATA.md` — `DATA-CONSISTENCY`.

## QUALITY-PERFORMANCE-RESOURCES Performance, load, and resources

The initial public API planning envelope is approximately **up to 1.5 million requests per month**. This is a planning input, not a capacity/performance guarantee, SLO, autoscaling constant, or assertion that one public request causes one external model/provider call. Requests, media, memory, DB connections, queues/backlogs, evaluator/provider concurrency, and external outputs remain bounded; upstream admission is constrained by sustainable downstream capacity rather than allowed to grow without limit.

Performance work follows measured evidence: observe the actual bottleneck, remove obvious query/code/contention inefficiency, tune bounded concurrency/pooling/admission, then increase vertical or horizontal capacity inside downstream bounds; topology changes require demonstrated need. Cache and derived projections may improve delivery only while access/freshness/correctness remain intact, and they never become product authority. Under exhaustion, latency and optional capability are sacrificed before correctness, durable learner/product state, evidence integrity, privacy, or security.

Exact RPS, latency targets, pool sizes, instance counts, concurrency limits, autoscaling maxima, backlog limits, and SLO/RPO/RTO numbers remain deployment/load-test/operational policy. This migration introduces no Kubernetes, Kafka-class broker, service mesh, multi-region active-active design, second queue/cache, replica/sharding scheme, or other new infrastructure.

Sources: `design/06-implementation-stack.md` — reliability/recovery/performance/deployment and capacity relations; `design/07-third-party-services.md` — Initial production planning envelope.

## QUALITY-OBSERVABILITY Observability

Operational logs, security events, privileged audit, product analytics, traces/metrics, and learner-visible history remain distinct classes. Runnable units expose appropriate health and structured correlation/provenance sufficient to reconstruct consequential execution across request, authoritative commit, async work, capability execution, reconciliation, and downstream recovery where material. Useful operational evidence includes privacy-safe request/work/execution identity, duration, result/failure/reconciliation class, backlog/retry/capacity pressure, and consequential build/contract/configuration provenance.

Telemetry is operational evidence only; it does not become learner/product state or prove semantic completion. Secrets and unnecessary sensitive learner content are redacted/minimized. Telemetry delivery failure normally does not rewrite an otherwise valid authoritative transaction, while an explicitly required durable security/privileged audit record may be a precondition or transactional requirement for the operation it protects.

Sources: `design/05-api.md` — execution trace levels; `design/06-implementation-stack.md` — Observability boundary and Health, observability, and incidents.

## QUALITY-TESTING-EVIDENCE Testing and evidence

Repository-native local verification is the correctness procedure; the current root command is `./verify`. Each materialized runtime/boundary contributes its native deterministic checks, and cross-boundary verification targets the failures the architecture actually needs to prevent: auth/access ordering and bypasses, service trust, decisive idempotency/concurrency, stale/duplicate/late completion fencing, crash/recovery paths, deletion/privacy fencing, contract/registry drift once materialized, persistence ownership, SSE isolation, and security-critical configuration failure.

Verification fixtures use synthetic, owned/licensed, or appropriately de-identified material. Normal fixtures do not contain production learner credentials, real provider/auth secrets, unnecessary production learner responses/audio, copied production databases, live provider credentials, or unauthorized copyrighted material.

Tests, logs, metrics, CI mechanics, and repository verification are evidence about implementation/repository behavior; they are not product truth, learning/evidence authority, or an `agent-standards` PASS/maturity claim. CI may invoke the same repository-native verification contract when separately authorized, but this migration creates no workflow or second correctness procedure.

Sources: `design/06-implementation-stack.md` — Verification fixture/data boundary, Native verification baseline, Root verification.

## QUALITY-COST-USAGE Cost and usage bounds

Current operator-adjustable planning defaults are:

```text
operating target                 = USD 10 / month
variable-spend safety ceiling    = USD 20 / month
```

The operating target guides forecasting, alerts, quotas, and optional-capability optimization. The variable-spend safety ceiling is an application-level admission boundary for **new discretionary variable-cost work**. Neither value is an architecture constant or invoice guarantee: fixed infrastructure, metering/reporting delay, already admitted/in-flight work, estimation error, unresolved potentially billable execution, and provider billing behavior can make final spend differ from the configured boundary.

A new cost-bearing logical operation conservatively estimates/reserves spend and is admitted atomically/serializably against available discretionary allowance before paid execution; later usage reconciles that same logical reservation. Retry, repair, fallback, escalation, replacement execution, continuation, period rollover, timeout, or ambiguous outcome cannot obtain fresh independent allowance that bypasses the original logical operation's cost policy. If current pricing/cost-policy state cannot be established safely, new discretionary paid work fails closed, remains pending/delayed, or is truthfully unavailable while healthy deterministic/near-zero-cost integrity and learner-state paths remain eligible under their own dependencies.

Cost pressure may reduce availability, batch/delay optional work, or deny new discretionary spend. It never lowers evidence/evaluator/content/privacy/security quality, substitutes an unqualified route for the same consequence, fabricates learner weakness/readiness, or discards durable learner state. Provider pricing, billing plans, and universal scaling thresholds remain outside this migration.

Sources: `design/06-implementation-stack.md` — Financial operating policy and concurrency-safe admission; `design/07-third-party-services.md` — Initial production planning envelope and external-resource economy.

## Migration boundary

`docs/QUALITY.md` and the existing `docs/*` migration set remain **AUTHORITY NONE**. This file preserves the ten V1 QUALITY coverage concerns backed by the sections above. Neighboring architecture/interface/delivery/decision concerns are now migrated by their own bounded owners, and `UNK-001..UNK-009` are `RESOLVED`; this reconciliation changes none of those requirements or relations. `milestone.scope_state` remains `OPEN` pending a post-reconciliation semantic closure/freeze audit.

This QUALITY reconciliation creates or changes no provider inventory or provider lifecycle state, no `EXT-*`/`CAP-*`/`DEC-*` identity, no exact OpenAPI/wire schema, no implementation, no workflow, no numerical SLO, no `agent-standards` PASS/level, no `DOCS_READY`, no design lock, no implementation-readiness claim, no provider activation, no promotion, and no release claim.
