STATUS: CANONICAL
OWNS: public/internal API resource model, route groups, operation semantics, request/response execution patterns, response/failure classes, async API behavior, idempotency/error conventions, auth/session transport contract prerequisites, content supply/validation/operations API semantics, and contract-materialization/evolution rules
DEPENDS_ON: ../spec/01-LEARNER-MODEL.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 02-practice-catalog.md, 04-application-flows.md
DOES_NOT_OWN: learning truth, external IELTS task-family definitions, content semantic identity/quality policy, TargetProfile UX semantics, product coverage policy, evaluator/generator algorithms, runtime lifecycle truth, privileged operational capability meaning or authorization role matrix, framework selection, persistence schema, provider choice, deployment technology, or exact wire schema after machine contracts exist

# API

## Purpose

Define semantic API boundaries before implementation. The product exposes one learner/admin-facing public API through Go Core API and bounded internal evaluator/generation/validation capabilities to Python where applicable.

# Principles

1. public version prefix `/v1`;
2. resource-oriented design over one-endpoint-per-button RPC;
3. stable canonical/external IDs and exact content revision identity cross boundaries unchanged;
4. creation/submission operations are idempotent where retry could duplicate history/cost;
5. long work returns durable pending resources rather than unbounded requests;
6. evidence states, target-resolution/support outcomes, content-assignment outcomes, and CoverageGap are domain results, not generic transport failures;
7. browser/admin clients never call Python evaluator/generator/validator directly;
8. variant, official family, Content Context, material Presentation Class, content revision, and delivery are explicit wherever they change content, scoring, inference, coverage, or readiness behavior;
9. API resources expose lifecycle/operational states owned by `04-application-flows.md` rather than redefining transition rules here;
10. once materialized, machine-readable contracts own exact wire shape;
11. API fields translate canonical/product semantics; they never invent thresholds, state meaning, applicability, validation authority, or evidence authority that lacks an upstream owner;
12. content operations may change release/assignment/operational state but may not mutate the semantic payload of a historical ContentRevision in place.

# Implementation-start contract gate

Parallel runtimes must not independently author equivalent DTO/schema definitions.

Before two runtime units implement the same boundary:

```text
design/05-api.md            semantic intent
        ↓
contracts/http/...          exact machine contract
        ↓
validation / generated bindings where useful
        ↓
consumer/provider implementation
```

Public and internal evaluator/content-capability APIs may use separate HTTP contract files, but each boundary has exactly one exact machine authority.

Required checks once materialized:

- schema validation;
- generated-artifact drift checks where generation is used;
- stable canonical/external ID and content-revision compatibility;
- canonical applicability/nullability preservation;
- public/internal boundary separation;
- consumer/provider conformance;
- directional compatibility for every deployed/allowed version-skew direction;
- contract version/provenance in integration verification.

Machine schemas define transport shape, not IELTS learning/exam/content-quality truth. Generated clients/server bindings are derived and are regenerated/validated from the contract rather than manually patched into authority.

## Machine-contract applicability invariant

Machine contracts preserve material distinctions among:

```text
optional input not supplied
explicitly NOT_APPLICABLE by canonical semantics
UNKNOWN / unresolved
required but absent or otherwise invalid
present value
```

These states may use different exact wire representations later; they may not be collapsed when the distinction changes meaning. A semantic dimension that is legitimately not applicable must not become required merely because a uniform schema is convenient. A materially required dimension must not become optional merely because transport permits omission.

Changing the meaning or mapping among absence, null, `NOT_APPLICABLE`, unresolved, invalid, and present values can be a breaking semantic contract change even when the JSON still validates. Generated TypeScript, Go, and Python consumers must preserve the distinction wherever material.

# Auth/session transport pre-contract gate

Architecture remains provider-neutral, but the initial public OpenAPI security scheme and auth-sensitive browser behavior must not be guessed from framework defaults.

Before the first public contract freezes security-sensitive transport, implementation explicitly chooses the initial credential/session transport well enough to determine:

- credential custody and which unit may read/verify it;
- browser cookie/header/storage behavior;
- CSRF applicability;
- CORS-with-credentials behavior;
- session/logout revocation and invalidation;
- guest→account transition/association behavior;
- learner/admin/service credential separation;
- service-to-service authentication where the actual deployment requires it;
- secret/key/token rotation appropriate to the selected mechanism.

OAuth, JWT, a hosted identity provider, opaque sessions, or another mechanism are not selected by this document. This is a required pre-contract implementation decision when exact security-scheme encoding depends on it, not permission for OpenAPI/code generation to invent the mechanism.

# Request/access ordering

Safe bounded transport framing, size checks, request-correlation setup, credential extraction, and public contract-level parsing may occur before identity is known. Protected-resource-sensitive semantic validation, existence checks, state inspection, and capability checks follow the applicable authentication/access boundary when earlier processing could disclose protected information.

Public error behavior must not accidentally reveal:

- whether another learner's protected resource exists;
- privileged content existence/details;
- admin capability configuration;
- provider secrets or internal provider/runtime state.

The exact future `403`/`404` or equivalent mapping belongs to the machine contract/security decision. The invariant is non-disclosure, not one status-code policy.

# Request/response execution patterns

These patterns translate `04-application-flows.md` into API operation behavior without pre-authoring exact OpenAPI fields.

## Pattern 1 — synchronous read

```text
Client
  ↓ bounded framing / public parsing / credential extraction
authentication where required
  ↓
authorization / access filtering where required
  ↓
resource-sensitive validation + authoritative query
  ↓
canonical/applicability projection
  ↓
response
```

Properties:

- public anonymous reads may omit authentication where explicitly permitted;
- reads have no hidden product mutation;
- the operation has a bounded deadline;
- access/error behavior does not leak protected existence;
- cache/derived reads optimize delivery only when freshness/access semantics remain correct; cache never becomes authority.

## Pattern 2 — synchronous mutation

```text
Client
  ↓ bounded framing / credential extraction
authentication where required
  ↓
capability / authorization / access filtering where required
  ↓
structural contract validation
  ↓
resource-sensitive semantic preconditions
  ↓
idempotency + concurrency checks where applicable
  ↓
authoritative transaction
  ├─ product mutation
  └─ required durable pending-work/recoverable marker where applicable
  ↓
COMMIT
  ↓
post-commit dispatch where applicable
  ↓
response
```

A response claiming durable success is emitted only after the authoritative commit. Required asynchronous continuation cannot be left to an unrecoverable post-commit registration step.

## Pattern 3 — asynchronously accepted operation

```text
POST / operation request
  ↓ bounded framing / credential extraction
authentication where required
  ↓
capability / authorization / access filtering where required
  ↓
contract + semantic preconditions
  ↓
establish logical work identity
  ↓
authoritative transaction
  ├─ persist accepted resource/work state
  └─ persist durable pending-work/recoverable marker
  ↓
COMMIT
  ↓
accepted/pending response
  ↓
bounded background/internal execution
  ↓
GET resource and/or SSE update
```

Sending an HTTP request to Python/provider is not itself durable acceptance.

## Pattern 4 — idempotent create/submission

```text
operation identity
+ actor/resource scope
+ compatible payload identity
+ current lifecycle state
        ↓
one authoritative logical outcome
```

The exact contract/bootstrap must define, for each applicable operation:

- idempotency scope and operation identity;
- actor/resource association;
- payload-compatibility rule;
- durable outcome/result/reference returned by replay;
- record retention sufficient for the expected retry horizon;
- conflict behavior for the same key/identity with incompatible payload.

No universal TTL is frozen here. Idempotency records are operational support state, not learner truth. A network retry must not duplicate an Attempt, EvidenceFact, accepted ContentRevision, paid/provider work, or logically identical generation/validation/evaluation work.

## Pattern 5 — internal Go → Python capability call

```text
Core API
  ↓ exact internal HTTP contract
caller deadline + cancellation where safe
  ↓
Evaluator / generator / validator capability
  ↓ bounded output + provenance/uncertainty
Core-side contract/work-identity validation
  ↓
owning Assessment/content/product policy interpretation
```

This is an internal capability, not a public product API. Core API is the permitted product caller; browser and admin UI do not call it directly. Deployment restricts reachability according to the selected topology. If the boundary crosses an untrusted/shared network, caller/service authentication and authorization are required. A private trusted/co-located transport may use a smaller mechanism only when its actual trust boundary makes that safe. This design does not freeze mTLS, JWT, or a service mesh.

Python output remains non-authoritative for certification, product support, learner-state transition, evidence admission, content activation, or final DailyPlan state. Python does not query or mutate the authoritative product database directly.

## Pattern 6 — privileged operation

```text
public/admin operation
  ↓ bounded framing / credential extraction
authentication
  ↓
applicable operational capability / access check
  ↓
contract + current-state preconditions
  ↓
legal mutation + required privileged audit/work marker
  ↓
COMMIT
  ↓
response / post-commit dispatch
```

Capability semantics are owned by `04-application-flows.md`; this API does not invent role hierarchy or bypass authority. Admin UI cannot manipulate authoritative DB/provider state around Core API.

## Pattern 7 — SSE update

```text
authoritative state change
  ↓
scoped event/update hint
  ↓
authorized SSE stream
  ↓
client refresh/presentation
  ↓
durable resource GET remains truth
```

Where the stream carries protected state, it authenticates/authorizes the connection and filters events to the learner/admin resources/capabilities the caller may observe. Payloads contain only the minimum information needed to identify the relevant update and must never expose another learner's resource or privileged state.

SSE events may be duplicated, delayed, reordered, or lost. Reconnect/resume cursors and event IDs are transport state, not product authority. Event/resource revision correlation may help clients discard stale hints, but a durable resource GET restores current truth. The exact event envelope belongs to one machine authority when materialized.

Because this SSE stream is delivered over the public HTTP API, its route/envelope belongs to the public HTTP contract. Do not create `contracts/events` merely for SSE. A separate event contract is justified only by a genuinely separate cross-unit asynchronous event boundary.

## Pattern 8 — inbound webhook, conditional

Inbound webhooks exist only when an actually selected external capability requires callbacks. Detailed provider authentication/replay/egress/ingress semantics are owned by `07-third-party-services.md`.

At the API boundary, a callback must be authenticated according to its provider contract, bounded/validated, associated with authoritative work, idempotent/replay-safe, and committed through normal Core policy before it can affect product state. Callback timestamps/provider success are not product authority.

# Execution trace levels

A non-authoritative trace model lets implementation/debugging follow one operation without creating another domain ontology:

```text
L0  caller / product action
L1  semantic API operation/resource
L2  exact transport contract once materialized
L3  Core API authoritative execution/transaction
L4  internal capability/provider invocation where applicable
L5  resulting product state propagation + learner/admin response
```

For a consequential operation, implementation can reconstruct where material: initiator/operation; stable resource/work/content/attempt/target identities; auth/access decision; validation/preconditions; idempotency/concurrency decision; commit point; durable async registration/dispatch state; internal/provider invocation; retry/fallback class; persisted result/failure; caller response and later SSE/resource update.

Trace/log data is operational evidence only; it never becomes learner/product semantic authority.

# Response and failure classes

These are semantic execution classes, not necessarily future public wire enums.

## A. Transport success + domain success

The transport completed and the requested domain operation/read succeeded under its contract.

## B. Transport success + domain unresolved/negative

The request executed correctly but the domain result is unresolved, pending, unavailable for the requested semantic reason, or legitimately negative, for example:

- insufficient/conflicting/stale learner evidence;
- unresolved TargetProfile condition;
- target not supported for a fully resolved scope;
- a valid CoverageGap result;
- learner-specific content unavailable because novelty/independence fails;
- evaluation/work still pending.

These states are not infrastructure errors merely because the desired product outcome is unavailable.

## C. Operation-contract rejected

The operation is not legally executable as requested, for example malformed input, unauthenticated identity, unauthorized/forbidden capability, stale revision, idempotency conflict, invalid lifecycle transition, failed precondition, or incompatible semantic combination.

## D. Infrastructure/transient failure

The requested operation cannot currently establish its intended authoritative result because an infrastructure dependency is unavailable/ambiguous, for example database outage, internal dependency/provider outage, or timeout whose remote result cannot yet be established safely.

Infrastructure failure is never low learner performance, evidence of weakness, or automatic content-quality failure. An ambiguous timeout does not prove remote failure and cannot authorize unsafe retry.

# Deadline, retry, fallback, and backpressure semantics

Every material network/provider boundary ultimately declares caller deadline, safe cancellation, transient/permanent/ambiguous retry classification, idempotency/deduplication, bounded retry/backoff when eligible, safe fallback, and capacity/backpressure behavior.

Exact timeout/retry counts/budgets are implementation/operational policy until measured evidence justifies freezing them. A retry cannot lower content/evidence/evaluator quality or turn ambiguous work into a duplicate logical operation. Rate limiting may protect abuse, cost, quotas, and capacity but cannot fabricate learner failure or silently accept undurable work.

# Public resource groups

The initial semantic API surface has 12 resource groups.

## 1. Identity, learner, target

```text
GET    /v1/me
GET    /v1/learner-profile
PATCH  /v1/learner-profile
GET    /v1/target-profile
PUT    /v1/target-profile
GET    /v1/target-support
```

`target-profile` represents supplied/material TargetProfile constraints from `00-learning-experience.md`: standard variant, delivery/purpose constraints, overall/per-skill lower-bound Band constraints, test date, selected One Skill Retake focus, and concurrency revision where applicable.

For target-relative planning/readiness/support, the standard variant plus at least one real Band constraint must be resolved. Missing target input remains unresolved rather than being silently defaulted, converted to learner evidence insufficiency, or fabricated as CoverageGap. `target-support` transports the support result owned downstream by `08-coverage-and-support.md`; it is not learner readiness.

## 2. Diagnostics

```text
POST   /v1/diagnostic-runs
GET    /v1/diagnostic-runs/{diagnostic_run_id}
```

A diagnostic exposes its lifecycle, sampling TargetProfile/reference, sample identities/observations/evaluation status, and a material sampling ledger distinguishing at least `sampled`, `not_sampled`, `unusable`, and `pending_evaluation` where applicable. `completed` means the run ended, not that the learner model/target is complete or certified.

## 3. Daily plan

```text
GET /v1/daily-plan
```

Response semantics preserve plan/TargetProfile/learner-state references, eligible activity cards, reason codes, canonical/content/variant/delivery identities where material, primary purpose, evidence candidacy, unresolved target conditions, and CoverageGap indication. Hard eligibility and ranking remain distinct as owned by `04-application-flows.md`.

## 4. Learning sessions

```text
POST   /v1/learning-sessions
GET    /v1/learning-sessions/{learning_session_id}
PATCH  /v1/learning-sessions/{learning_session_id}
```

Mutations follow the lifecycle owned by `04-application-flows.md`. Session completion never means mastery.

## 5. Practice modes/activities

```text
GET    /v1/practice-modes
POST   /v1/practice-activities
GET    /v1/practice-activities/{practice_activity_id}
```

A PracticeActivity resolves exact content revision, canonical target, applicable official family/context/presentation/variant/delivery identities, stimulus/response conditions, scaffolding/exposure state, `primary_activity_purpose`, and `evidence_candidacy`.

Purpose and evidence candidacy are orthogonal. The client/evaluator/server cannot retroactively upgrade candidacy after observing a favorable result. Actual EvidenceFact admission remains Assessment authority.

## 6. Attempts

```text
POST   /v1/attempts
GET    /v1/attempts/{attempt_id}
PATCH  /v1/attempts/{attempt_id}
POST   /v1/attempts/{attempt_id}/submissions
```

Draft mutation uses concurrency control. Submission is explicit/idempotent and preserves exact assigned revision plus material assistance, exposure/retry, delivery/input, timing, and response provenance. Accepted submission corresponds to durable authoritative state; required async continuation is already durably discoverable/recoverable before acknowledgement.

## 7. Evaluations

```text
GET /v1/evaluations/{evaluation_id}
```

Public output may expose status, learner-meaningful observations/feedback, uncertainty/quality state where useful, and retry/unavailable state. It does not expose chain-of-thought, secrets, or irrelevant provider internals.

## 8. Progress and gaps

```text
GET /v1/progress
GET /v1/gaps
```

`progress` exposes target conditions, current per-skill support/certification state, history, freshness, and unresolved requirements. `gaps` exposes learner GapEvaluation + ActionIntent. CoverageGap remains separate.

## 9. Review

```text
GET /v1/review-queue
```

Items preserve their semantic queue kind: `knowledge_retrieval`, `error_remediation`, or `re_evidence`, plus canonical target/action references where available.

## 10. Mocks

```text
POST   /v1/mock-runs
GET    /v1/mock-runs/{mock_run_id}
```

MockRun preserves `READINESS` purpose for normal mocks, evidence candidacy, exact content revisions, resolved variant, applicable family/context/presentation/delivery scope, actual completion state, and section outputs at their valid scope. Readiness purpose never admits evidence automatically.

## 11. Media

```text
POST   /v1/media-sources
GET    /v1/media-sources/{media_source_id}
POST   /v1/media-lessons
GET    /v1/media-lessons/{media_lesson_id}
```

MediaSource creation resolves provider identity/metadata, playability, rights, and transcript state through the eligible media/provider boundaries. A URL is an untrusted reference, not arbitrary fetch/download authority. Generated lesson content enters normal revision/validation policy before assignment.

## 12. Content supply and operations

Representative semantic operations:

```text
GET    /v1/admin/content-revisions
GET    /v1/admin/content-revisions/{content_revision_id}
PATCH  /v1/admin/content-revisions/{content_revision_id}
POST   /v1/admin/content-generation-requests
GET    /v1/admin/content-generation-requests/{generation_request_id}
POST   /v1/admin/content-validation-runs
GET    /v1/admin/content-validation-runs/{validation_run_id}
POST   /v1/content-reports
GET    /v1/admin/content-reports
PATCH  /v1/admin/content-reports/{content_report_id}
```

Exact fields/subresources belong to machine contracts. Privileged capability meaning, content lifecycle, quarantine/revalidation/release/retirement, and the prohibition on validation bypass are owned by `04-application-flows.md`. API operations cannot mutate historical ContentRevision semantic payload or turn generator/validator/operator output directly into active learner content.

# Server event stream

```text
GET /v1/event-stream
```

This route is part of the public HTTP API contract. It follows Pattern 7: protected streams are access-scoped, payloads are minimum-necessary, cross-learner/admin leakage is forbidden, delivery is non-authoritative, and durable resource reads recover truth after loss/reorder/reconnect. The future HTTP machine contract owns the exact SSE envelope/correlation semantics.

# Internal evaluator/content capability API

Normal product caller: Go Core API only.

```text
POST /internal/v1/evaluations
POST /internal/v1/media-analyses
POST /internal/v1/content-generations       when implemented/needed
POST /internal/v1/content-validations       when implemented/needed
GET  /internal/v1/health
```

Generation/validation endpoints are optional capability boundaries. This route group is internal, not browser/admin/public internet product surface. Deployment reachability and service authentication follow Pattern 5 and `06-implementation-stack.md`.

## Evaluation request/response semantics

Requests preserve work/attempt/content revision identities, canonical target IDs, material family/context/presentation/variant/delivery scope, actual response/assistance/exposure conditions, evaluator configuration reference, and requested observation families. Applicability is preserved rather than filled with fabricated values.

Responses contain bounded status/Observation candidates, criterion measurements, permitted derived analysis references, uncertainty/quality flags, evaluator/model provenance, and diagnostics. They never carry authoritative learner certification, product support, Band advancement, evidence candidacy/admission, content activation, or final plan state.

## Content generation/validation/media semantics

Generation receives a Core-constrained demand and returns candidate content plus provenance; it cannot activate content or invent unresolved canonical meaning. Validation receives exact revision/intended-use/policy context and returns findings/signals/provenance, not learner/evidence/product authority. Media analysis receives only authorized text/media references/data and returns bounded proposals; Core applies normal eligibility and content policy.

# Idempotency

Use a stable idempotency key/equivalent operation identity where retry can duplicate learner history or cost, including diagnostic/session creation where duplicates matter, attempt submission, mock creation, cost-bearing media work, generation, and validation.

The future exact contract/bootstrap must instantiate Pattern 4 for each applicable operation. Retry cannot create duplicate attempts, EvidenceFacts, paid work, or semantically duplicate accepted revisions for one logical operation.

# Optimistic concurrency

Mutable draft-like resources, TargetProfile, and mutable operational/admin metadata use resource-revision semantics where stale updates are possible. A stale mutation is rejected rather than silently overwriting newer state. This resource revision is not API contract version, ContentRevision, database schema version, provider/model version, or idempotency/work identity.

# Domain results vs transport/operation failures

Domain/result states such as insufficient/conflicting/stale evidence, unresolved target condition, target non-support/CoverageGap, learner-specific content unavailability, and evaluation pending are normal semantic outcomes when the request executed correctly.

Transport/operation failures cover malformed input, authentication/access failure, invalid transition/precondition, stale revision, idempotency conflict, rate limit, dependency unavailability, and internal failure. Exact public codes/status mappings belong to the future contract and must respect the non-disclosure rule above.

A `5xx`-class infrastructure failure is never a learner score, learner weakness, CoverageGap, insufficient-evidence result, or fake content-quality judgment.

# Machine-contract evolution and compatibility

Compatibility is directional. A change is safe only for the deployed caller/provider combinations that must coexist; “additive in YAML” is not sufficient evidence.

For each affected boundary, review as applicable:

```text
old client  → new server
new client  → old server
new server  → old client
old server  → new client
```

For Web↔Go this means old/new Web and old/new Go combinations. For Go↔Python it means old/new Go and old/new Python combinations whenever deployment can overlap them. Server→client response/SSE compatibility is evaluated separately from client→server request compatibility.

Compatibility analysis covers at least:

- unknown response fields and whether deployed consumers ignore/preserve/reject them safely;
- unknown/new enum, status, reason, or domain-result values;
- newly required request fields;
- newly optional fields and whether absence changes semantics;
- narrowed validation or accepted-value ranges;
- changed default behavior;
- changed null/absence/`NOT_APPLICABLE`/unresolved semantics;
- removed/renamed fields or operations;
- changed response-state meaning;
- changed lifecycle/precondition/authorization applicability;
- changed stable identifier meaning.

Adding a response field is compatible only when affected deployed consumers tolerate it. Adding an enum/status value is compatible only when the contract and every affected deployed consumer explicitly tolerate unknown/new values or otherwise handle the new value safely. Generated clients/types are not assumed tolerant; generation output and runtime behavior must be verified.

Rules:

1. transport optionality/requiredness cannot violate canonical applicability to avoid a version change;
2. deployed `/v1` behavior cannot silently break a supported consumer direction;
3. a breaking deployed boundary change requires an explicit version/rollout/migration strategy before activation;
4. internal contracts may evolve faster, but cannot assume atomic Go/Python replacement unless deployment explicitly guarantees it;
5. generated bindings are derived from one contract authority and do not eliminate runtime version skew;
6. compatibility/conformance tests cover every version-skew combination the selected rollout permits;
7. DB migrations support the application/schema compatibility window required by that rollout;
8. deprecation/new transport versions do not reinterpret historical ContentRevision, Attempt, Observation, EvidenceFact, or certification history;
9. a genuinely separate event contract, if introduced later, has its own producer/consumer compatibility rules; SSE alone remains under the public HTTP contract.

Exact version numbers, rollout technology, compatibility window, and generator library remain future contract/deployment work. This file does not mandate one SemVer algorithm.

# Pagination

Unbounded history/library/content/report collections use cursor pagination. Small stable canonical catalogs may be returned in full.

# API anti-patterns

Forbidden without an explicit architecture change:

- browser/admin calling evaluator/generator/validator directly;
- Next.js Route Handler/Server Action becoming a second product API or domain-command owner;
- endpoint per UI button;
- provider/model names in public domain routes;
- independently handwritten mirror DTOs across runtimes;
- generated binding/client treated as semantic authority or assumed compatibility evidence;
- fake score zero on evaluator failure;
- API shape becoming a second learner-state/content-quality definition;
- API-owned thresholds or hidden per-skill target constraints;
- CoverageGap represented as learner GapEvaluation or infrastructure error;
- unresolved TargetProfile input represented as learner evidence insufficiency or fabricated product non-support;
- ranker returning an activity that failed hard eligibility;
- omitting a material content revision/family/context/presentation/delivery identity;
- fabricating a family/context/presentation/variant value for a legitimately non-applicable dimension;
- client/evaluator reclassifying immutable item family or evidence candidacy;
- mutable PATCH of historical ContentRevision semantic payload;
- generator/validator output directly activating learner content;
- generic admin validation bypass for a known hard failure;
- one combined ContentStatus hiding validation/release/operational dimensions;
- `completed diagnostic` represented as `complete learner baseline`;
- retrying ambiguous/non-deduplicated mutation merely because a client timed out;
- treating additive schema or generated-code success as proof of deployed compatibility.