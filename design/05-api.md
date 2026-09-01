STATUS: SUPERSEDED_NON_CANONICAL
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
11. API fields translate canonical/product semantics; they never invent thresholds, state meaning, applicability, validation authority, evidence authority, or learner reveal policy that lacks an upstream owner;
12. content operations may change non-semantic operational/release state but may not mutate an established ContentRevision semantic payload.

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

# Auth/session transport selection

The initial external identity route is **Clerk**, selected for implementation under `07-third-party-services.md`. The authority split is:

```text
Clerk
  = credential custody
  + authentication
  + session issuance / revocation mechanics

Core
  = stable internal actor / learner identity
  + external-principal association
  + RBAC / capabilities
  + entitlements
  + learner/product state
  + all product authorization
```

Clerk roles, permissions, organizations, or metadata never become canonical `ADMIN` / `OWNER` / `REVIEWER` / `COLLABORATOR` authorization authority. The external Clerk subject/principal is associated with a stable Core-owned actor; Core applies the capability model in `04-application-flows.md` after authentication.

Initial public authenticated transport is:

```text
Browser / eligible Next.js presentation execution
        ↓ supported Clerk session API
short-lived Clerk-issued token
        ↓
Authorization: Bearer <token>
        ↓
Go Core API verifies authenticity + configured issuer/audience/authorized-party/expiry
and every other security condition required by the selected Clerk integration
        ↓
external principal → stable Core actor association
        ↓
normal Core product authorization
```

Auth tokens are not persisted in `localStorage`. The Core public API therefore has enough selected transport semantics for the future OpenAPI contract to encode an HTTP Bearer/JWT security scheme; exact scheme names, scopes, error/status mapping, and schema syntax remain for contract materialization.

Public/demo behavior may remain anonymous only where explicitly allowed. Durable learner state—including `TargetProfile`, Attempt history, Progression/history, and entitlement—requires authenticated durable identity. Guest→account transfer remains future evidence-driven work rather than a bootstrap identity-merging requirement.

Do not build password hashing/reset, OAuth-provider plumbing, passkey machinery, custom session rotation, a custom identity provider/session framework, or durable anonymous learner identity. Service-to-service authentication is separate from learner/admin transport and follows `06-implementation-stack.md`.

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
optional preflight idempotency / stale-revision rejection
  ↓
authoritative transaction or equivalent serialization boundary
  ├─ decisive idempotency admission/replay/conflict decision where applicable
  ├─ decisive expected-resource-revision comparison where applicable
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

Preflight lookups may reject duplicate/stale work early but are not the decisive concurrency boundary. For duplicate-sensitive or revision-guarded mutations, the operation-identity claim/replay decision and expected-revision comparison must be atomically coupled to the mutation or protected by an equivalent serialization invariant.

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
optional preflight duplicate/stale rejection
  ↓
authoritative transaction or equivalent serialization boundary
  ├─ decisively admit/replay one logical operation
  ├─ persist accepted resource/work state + logical work identity
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

Sending an HTTP request to Python/provider is not itself durable acceptance or semantic completion.

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

The decisive claim of a previously unseen operation identity, its compatible-payload association, and the authoritative mutation/outcome it protects are transactionally/atomically coupled or protected by an equivalent serialization invariant. A separate preflight lookup is only an optimization. A crash after authoritative mutation cannot leave the idempotency outcome unassociated such that retry executes the mutation again.

No universal TTL is frozen here. Idempotency records are operational support state, not learner truth. Replay after committed success returns or references the one authoritative logical outcome. The same operation identity with an incompatible payload fails closed. A network retry must not duplicate an Attempt, EvidenceFact, accepted ContentRevision, paid/provider work, or logically identical generation/validation/evaluation work.

## Pattern 5 — internal Go → Python capability call

```text
Core API
  ↓ exact internal HTTP contract
caller deadline + cancellation where safe
  ↓
Evaluator / generator / validator capability
  ↓ bounded output + provenance/uncertainty
Core-side contract + logical-work/execution association validation
  ↓
Core-side current-state/fencing/deletion-eligibility reconciliation
  ↓
owning Assessment/content/product policy interpretation
```

This is an internal capability, not a public product API. Core API is the permitted product caller; browser and admin UI do not call it directly. Deployment reachability/service authorization follows the trust-boundary rules owned by `06-implementation-stack.md`; an `internal`, private-address, or co-located label is not by itself authorization.

When retry/replacement executions may overlap, the machine contract preserves enough logical work/Evaluation association plus execution-attempt identity or equivalent fencing information for Core to reject stale/superseded completion deterministically. Duplicate delivery of an already accepted completion is idempotent.

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

SSE delivery is not proof that an operation or its required downstream semantic continuation is complete; the resource state defined by `04-application-flows.md` remains authoritative.

Because this SSE stream is delivered over the public HTTP API, its route/envelope belongs to the public HTTP contract. Do not create `contracts/events` merely for SSE. A separate event contract is justified only by a genuinely separate cross-unit asynchronous event boundary.

## Pattern 8 — inbound webhook, conditional

Inbound webhooks exist only when an actually selected external capability requires callbacks. Detailed provider authentication/replay/egress/ingress semantics are owned by `07-third-party-services.md`.

At the API boundary, a callback must be authenticated according to its provider contract, bounded/validated, associated with authoritative work, idempotent/replay-safe, reconciled against current deletion/content/work eligibility, and committed through normal Core policy before it can affect product state. Callback timestamps/provider success are not product authority.

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

For a consequential operation, implementation can reconstruct where material: initiator/operation; stable resource/work/content/attempt/target identities; auth/access decision; validation/preconditions; idempotency/concurrency decision; commit point; durable async registration; execution claim/dispatch state; internal/provider invocation; retry/fallback class; completion reconciliation; persisted result/recoverable continuation; caller response and later SSE/resource update.

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

# Learner/public content projection

An authoritative internal ContentRevision is not automatically a learner-facing DTO. Public learner/activity responses use an actor/use-specific projection of the same exact revision identity.

Before the applicable semantic policy permits reveal, learner/browser payloads exclude material whose disclosure would reveal an answer, invalidate the intended inference, or expose privileged execution state, including as applicable:

- answer keys/reference answers;
- hidden scoring references or rubrics not intended for learner display;
- model answers whose early exposure would contaminate the task;
- evaluator/system/validator instructions;
- future feedback;
- privileged validation/provenance details;
- other hidden assessment material.

Machine contracts encode the projection/reveal decision; they do **not** own semantic reveal policy. Reveal semantics are derived from the applicable upstream owners: `../spec/07-PRACTICE.md` for feedback timing/learning mechanism, `../spec/08-ASSESSMENT.md` for independence/evidence consequence, `../spec/10-CONTENT-MODEL.md` for represented answer/model/rubric material, and `04-application-flows.md` for current lifecycle/use context.

Thus a protected readiness performance cannot reveal answer/model material during the attempt; retrieval practice cannot reveal required feedback/model before the authentic retrieval attempt when Practice policy requires that ordering; and post-task review may reveal permitted model/rubric material only when the owning learning/evidence/product semantics allow it. OpenAPI later encodes these learner-safe projections but cannot invent reveal timing because a field is nullable or present in an internal schema.

Admin/evaluator projections remain capability/contract scoped. Projection does not create another ContentRevision or duplicate content truth; exact DTO fields belong to future public/internal machine contracts.

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

Response semantics preserve plan-time TargetProfile/target-context and learner/Progression references, content/release and product-support/coverage provenance where consequential, eligible activity cards, reason codes, canonical/content/variant/delivery identities, primary purpose/evidence candidacy, unresolved target conditions, and CoverageGap indication where material. Hard eligibility and ranking remain distinct as owned by `04-application-flows.md`.

A DailyPlan response is a snapshot, not an assignment grant. When a client later starts an activity from a plan reference, Core re-evaluates current hard eligibility and performs any required decisive reservation/assignment before returning the learner-safe activity projection; a stale plan item cannot force assignment.

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

PracticeActivity creation resolves current hard eligibility at assignment time rather than trusting historical plan eligibility. The resulting activity preserves the exact assigned content revision and assignment-time provenance needed by `04-application-flows.md`, then exposes the learner-safe projection plus canonical target, applicable official family/context/presentation/variant/delivery identities, stimulus/response conditions, scaffolding/exposure state, `primary_activity_purpose`, and `evidence_candidacy`.

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

Public output may expose status, learner-meaningful observations/feedback, uncertainty/quality state where useful, and retry/unavailable state. It does not expose chain-of-thought, secrets, hidden evaluator instructions, or irrelevant provider internals. Evaluation/execution retry, terminal-state, completion, and downstream-recovery semantics are owned by `04-application-flows.md`; `completed` on this resource cannot be implemented as an unrecoverable dead end with required Observation/Assessment/Progression work missing.

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
POST   /v1/admin/content-generation-requests
GET    /v1/admin/content-generation-requests/{generation_request_id}
POST   /v1/admin/content-validation-runs
GET    /v1/admin/content-validation-runs/{validation_run_id}
POST   /v1/content-reports
GET    /v1/admin/content-reports
PATCH  /v1/admin/content-reports/{content_report_id}
```

Exact fields/subresources and operational state-transition surfaces belong to machine contracts. The representative surface intentionally does not define a generic `PATCH` on an established ContentRevision: semantic payload is immutable under `spec/10-CONTENT-MODEL.md`, while ValidationDecision history and minimum validation semantics are owned there and release eligibility, operational safety/quarantine, reports, revalidation, and retirement remain runtime dimensions owned by `04-application-flows.md`.

Privileged capability meaning, content lifecycle, quarantine/revalidation/release/retirement, and the prohibition on validation bypass are owned by `04-application-flows.md`. API operations cannot turn generator/validator/operator output directly into active learner content or overwrite a historical ValidationDecision to represent revalidation.

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

Requests preserve logical work/Evaluation identity, Attempt/content revision identities, canonical target IDs, material family/context/presentation/variant/delivery scope, actual response/assistance/exposure conditions, evaluator configuration reference, and requested observation families. When execution attempts may overlap, the internal contract also preserves execution-attempt/fencing association sufficient for Core to determine whether a completion is current, duplicate, stale, or superseded without guessing. Applicability is preserved rather than filled with fabricated values.

Responses contain bounded status/Observation candidates, criterion measurements, permitted derived analysis references, uncertainty/quality flags, evaluator/model provenance, and diagnostics. They never carry authoritative learner certification, product support, Band advancement, evidence candidacy/admission, content activation, or final plan state. Core reconciles response acceptance according to current lifecycle/work/deletion/content-eligibility fencing in `04-application-flows.md` and `06-implementation-stack.md` before downstream learner semantics may be created.

## Content generation/validation/media semantics

Generation receives a Core-constrained demand and returns candidate content plus provenance; it cannot activate content or invent unresolved canonical meaning. Validation receives exact revision/intended-use/policy context and returns findings/signals/provenance, not learner/evidence/product authority. Media analysis receives only authorized text/media references/data and returns bounded proposals; Core applies normal eligibility and content policy.

# Idempotency

Use a stable idempotency key/equivalent operation identity where retry can duplicate learner history or cost, including diagnostic/session creation where duplicates matter, attempt submission, mock creation, cost-bearing media work, generation, and validation.

The future exact contract/bootstrap must instantiate Pattern 4 for each applicable operation. Retry cannot create duplicate attempts, EvidenceFacts, paid work, or semantically duplicate accepted revisions for one logical operation.

# Optimistic concurrency

Mutable draft-like resources, TargetProfile, and mutable operational/admin metadata use resource-revision semantics where stale updates are possible. A stale mutation is rejected rather than silently overwriting newer state.

The decisive expected-resource-revision comparison is enforced atomically with the authoritative mutation through a transaction, conditional write, or equivalent serialization invariant. A read/compare in application memory followed by an unrelated later write is not sufficient; preflight stale detection is only an optimization. Two simultaneous writers using the same expected revision cannot both successfully mutate from it.

This resource revision is not API contract version, ContentRevision, database schema version, provider/model version, or idempotency/work identity.

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

A cursor is derived transport state, not product authority. Where material it is interpreted only for the operation/query/filter/access scope that issued it and cannot bypass current authorization or resource visibility. Exact encoding, integrity/signing, expiration, and ordering semantics belong to the future contract/implementation.

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
- treating plan-time eligibility as assignment authority after current hard conditions changed;
- ranker returning an activity that failed hard eligibility;
- omitting a material content revision/family/context/presentation/delivery identity;
- fabricating a family/context/presentation/variant value for a legitimately non-applicable dimension;
- client/evaluator reclassifying immutable item family or evidence candidacy;
- mutating an established ContentRevision semantic payload through generic PATCH or another operational mutation;
- serializing an unrestricted internal ContentRevision representation to a learner/browser surface;
- using OpenAPI field presence/nullability to invent pedagogical/evidential answer or rubric reveal timing;
- generator/validator output directly activating learner content;
- overwriting an earlier ValidationDecision to represent a later policy run;
- generic admin validation bypass for a known hard failure;
- one combined ContentStatus hiding validation/release/operational dimensions;
- `completed diagnostic` represented as `complete learner baseline`;
- retrying ambiguous/non-deduplicated mutation merely because a client timed out;
- treating additive schema or generated-code success as proof of deployed compatibility;
- treating SSE delivery or internal HTTP success as semantic completion.