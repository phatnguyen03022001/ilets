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
7. browser never calls Python evaluator/generator/validator directly;
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

Public and internal evaluator/content-capability APIs may use separate contract files, but each boundary has exactly one exact machine authority.

Required checks once materialized:

- schema validation;
- generated-artifact drift checks where generation is used;
- stable canonical/external ID and content-revision compatibility;
- public/internal boundary separation;
- consumer/provider conformance;
- deployed `/v1` compatibility policy;
- contract version/provenance in integration verification.

Machine schemas define transport shape, not IELTS learning/exam/content-quality truth. Generated clients/server bindings are derived and are regenerated/validated from the contract rather than manually patched into authority.

## Machine-contract applicability invariant

Machine contracts must preserve canonical applicability. A semantic dimension that is legitimately `NOT_APPLICABLE` must not become required merely because a uniform transport shape is convenient. Conversely, a materially required dimension must not become optional merely because transport permits omission.

Exact representation of absence, explicit not-applicable state, conditional requirement, or validation constraints belongs to the future machine contract. Transport convenience may not fabricate or erase canonical meaning.

# Auth/session transport pre-contract gate

Architecture remains provider-neutral, but the initial public OpenAPI security scheme and auth-sensitive browser behavior must not be guessed from framework defaults.

Before the first public contract freezes security-sensitive transport, implementation must explicitly choose the initial credential/session transport well enough to determine:

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

# Request/response execution patterns

These patterns translate the runtime execution invariants in `04-application-flows.md` into API operation behavior without pre-authoring exact OpenAPI fields.

Safe bounded transport framing and credential extraction may occur before identity is known. Protected-resource-sensitive semantic processing follows the applicable authentication/access boundary so validation behavior cannot leak protected existence/data accidentally.

## Pattern 1 — synchronous read

```text
Client
  ↓ bounded transport framing / request-correlation context / credential extraction
authentication where required
  ↓
authorization / access filtering where required
  ↓
parse + structural contract validation
  ↓
resource-sensitive semantic preconditions / authoritative query
  ↓
canonical/applicability projection
  ↓
response
```

Properties:

- public anonymous reads may omit authentication where explicitly permitted;
- no hidden product mutation occurs as a side effect of a read;
- the operation has a bounded deadline;
- access policy must not accidentally leak protected resource existence;
- cache/derived reads may optimize delivery only when freshness/access semantics remain correct; cache never becomes authority.

## Pattern 2 — synchronous mutation

```text
Client
  ↓ bounded transport framing / credential extraction
authentication where required
  ↓
capability / authorization / access filtering where required
  ↓
structural contract validation
  ↓
resource-sensitive semantic preconditions
  ↓
idempotency check where applicable
  ↓
optimistic/concurrency check where applicable
  ↓
authoritative transaction
  ├─ product mutation
  └─ required durable pending-work/outbox/recoverable marker where applicable
  ↓
COMMIT
  ↓
post-commit dispatch where applicable
  ↓
response
```

A response claiming durable success is emitted only after the authoritative commit. Required asynchronous continuation cannot be left to an unrecoverable post-commit registration step. Dispatch may occur after commit because durable work identity/marker is already established or recoverably derivable.

## Pattern 3 — asynchronously accepted operation

Learner/admin initiated asynchronous work performs applicable auth/access checks and validation **before** authoritative acceptance:

```text
POST / operation request
  ↓ bounded transport framing / credential extraction
authentication where required
  ↓
capability / authorization / access filtering where required
  ↓
structural contract validation
  ↓
semantic preconditions
  ↓
establish logical work identity
  ↓
authoritative transaction
  ├─ persist accepted resource/work state
  └─ persist durable pending-work/outbox/recoverable marker
  ↓
COMMIT
  ↓
accepted/pending response
  ↓
background/internal capability dispatch/execution
  ↓
GET resource and/or SSE status
```

Expensive evaluator/generator/validator work should not keep an unbounded HTTP request open when durable asynchronous semantics are appropriate. Sending an HTTP request to Python/provider is not itself durable acceptance.

## Pattern 4 — idempotent create/submission

Conceptually:

```text
logical operation identity
+ compatible payload identity
+ current lifecycle state
        ↓
one authoritative logical outcome
```

A network retry must not duplicate an Attempt, EvidenceFact, one accepted logical ContentRevision, provider charge/work, or logically identical generation/validation/evaluation work. Retry of an operation known to be non-idempotent and not safely deduplicable is forbidden.

## Pattern 5 — internal Go → Python capability call

```text
Core API
  ↓ exact internal contract
caller deadline + cancellation where safe
  ↓
Evaluator / generator / validator capability
  ↓ bounded output + provenance/uncertainty
Core-side structural/work-identity validation
  ↓
owning Assessment/content/product policy interpretation
```

Python cannot return authoritative certification, product support, learner-state transition, evidence-candidacy upgrade, or content activation merely because the internal request succeeded. Python does not mutate Core-owned persistence directly.

## Pattern 6 — privileged operation

```text
public/admin operation
  ↓ bounded transport framing / credential extraction
authentication
  ↓
applicable operational capability / access check
  ↓
structural contract validation
  ↓
current revision/state semantic precondition
  ↓
legal mutation + required privileged audit/work marker
  ↓
COMMIT
  ↓
response / post-commit dispatch
```

The capability semantics are owned by `04-application-flows.md`; this API does not invent role hierarchy or bypass authority. Admin UI cannot manipulate the authoritative DB/provider around Core API.

## Pattern 7 — SSE update

```text
authoritative state/version change
  ↓
event/update hint
  ↓
client
  ↓
resource remains readable as current durable truth
```

SSE delivery may be duplicated, delayed, reordered, or disconnected without becoming product-state authority.

## Pattern 8 — inbound webhook, conditional

Inbound webhooks exist only when an actually selected external capability requires callbacks.

When used:

```text
provider callback
  ↓
signature/authentication verification
  ↓
replay/idempotency protection
  ↓
structural validation
  ↓
authoritative event/work association + safe freshness checks where relevant
  ↓
transaction + durable audit/work state
  ↓
COMMIT
  ↓
bounded response
```

A webhook endpoint is not created merely because webhooks are a common integration pattern. Provider callback timestamps are not trusted as causal truth merely because they are present. Webhooks cannot create learner/product truth outside normal owning policy.

# Execution trace levels

A non-authoritative trace model lets implementation/debugging follow one operation without creating another domain ontology:

```text
L0  caller / product action
L1  semantic API operation/resource
L2  exact transport contract once materialized
L3  Core API authoritative execution/transaction
L4  internal capability/provider invocation where applicable
L5  resulting canonical/product state propagation + learner/admin response
```

For a consequential operation, implementation must be able to reconstruct where material:

- initiator and semantic operation;
- stable resource/work/content/attempt/target identities crossing boundaries;
- authentication/access/capability decision;
- structural validation and semantic preconditions;
- idempotency/concurrency decision;
- authoritative transaction/commit point;
- durable asynchronous work identity/registration and dispatch state;
- internal capability/provider invocation and response authority;
- retry/fallback classification;
- persisted result/failure state;
- caller response class and later SSE/resource update.

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

Infrastructure failure is never low learner performance, evidence of weakness, or automatic content-quality failure. When a timeout is ambiguous, the server/client must not assume the remote operation failed and retry unsafely.

# Deadline, retry, fallback, and backpressure semantics

Every material network/provider boundary ultimately declares:

- caller deadline;
- cancellation propagation where safe;
- retry eligibility and classification: transient, permanent, or ambiguous;
- idempotency/deduplication behavior;
- exponential backoff/jitter where retry is appropriate;
- safe fallback behavior;
- backpressure/capacity behavior.

Exact timeout/retry counts/budgets are implementation/operational policy until measured evidence justifies freezing them. A retry cannot lower content/evidence/evaluator quality or turn ambiguous work into a duplicate logical operation.

Rate limiting may protect abuse, cost, provider quotas, and capacity. Exact limits are deployment/operations policy; rate limiting must not be used to fabricate learner failure or silently accept work that was not durably accepted.

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

`target-profile` represents, where supplied/material:

```text
test_variant
delivery_mode
purpose_or_receiving_rule
target_overall_band          lower-bound constraint
per-skill minimums           lower-bound constraints
test_date
selected_skill_retake        preparation focus, not eligibility proof
revision/version
```

Band fields preserve the TargetProfile semantics owned by `00-learning-experience.md`: `7.0` means `>= 7.0`, not exact equality. The API must not expand an overall target into hidden equal per-skill minima. Any derived working/planning profile is separately identified as non-authoritative and cannot mutate real target constraints without an explicit target update.

For **target-relative planning, readiness, or product-support evaluation**, a TargetProfile is sufficiently resolved only when it has both:

1. a resolved standard `test_variant`: Academic or General Training; and
2. at least one actual Band constraint: `target_overall_band` and/or one real per-skill minimum.

If the standard variant is unresolved, the variant-specific target remains an unresolved target condition. If no actual Band constraint is known, the readiness target remains unresolved. APIs may still support genuinely variant-independent/shared/foundational/diagnostic work where valid, but neither missing condition may be silently completed with a default.

Incomplete target input is not learner evidence insufficiency and is not automatically a product CoverageGap. `/v1/target-support` therefore cannot produce a complete target-relative support conclusion until the target-resolution minimum is met; it preserves/surfaces the unresolved target condition instead of fabricating `target_not_supported`, CoverageGap, or another product failure. Exact transport representation and HTTP behavior belong to the future machine contract.

`selected_skill_retake` selects focused preparation only. Eligibility-sensitive responses keep missing original-test/timing/location/delivery/purpose conditions unresolved rather than treating the selected skill as proof of One Skill Retake eligibility.

`target-support` transports the current product-support result for the sufficiently resolved target scope, including applicable delivery/purpose/eligibility conditions and blocking CoverageGap classes. Product-support semantics remain owned by `08-coverage-and-support.md`; referencing that downstream result here does not make API the policy owner or create a reverse semantic-definition dependency.

It is not learner readiness.

## 2. Diagnostics

```text
POST   /v1/diagnostic-runs
GET    /v1/diagnostic-runs/{diagnostic_run_id}
```

Creation selects the product diagnostic shape and optional focus constraints.

Result semantics include:

- run lifecycle state from `04-application-flows.md`;
- TargetProfile/reference used for sampling;
- Observations/evaluation status by sampled activity;
- stable family/context refs for samples where material;
- a material target-condition sampling ledger that distinguishes at least:
  - `sampled`;
  - `not_sampled`;
  - `unusable`;
  - `pending_evaluation`;
- resulting evidence states/next evidence need where available.

Diagnostic is an activity purpose, not an evidence-admission shortcut. A diagnostic result may expose admitted evidence only when the activity was an evidence candidate and normal Assessment policy admitted the resulting Observation.

A `completed` diagnostic does not mean every condition is known or certified.

## 3. Daily plan

```text
GET /v1/daily-plan
```

Optional request inputs may include duration preset and an eligible focus override.

Response semantics include:

- plan identity;
- TargetProfile revision/reference;
- learner/evidence state reference used to compute it;
- activity cards;
- reason codes;
- estimated duration;
- canonical targets;
- official family/Content Context refs where material;
- exact content revision references for resolved activities;
- delivery refs where material;
- primary activity purpose;
- evidence candidacy;
- unresolved target conditions;
- CoverageGap indicator where applicable.

DailyPlan is the output of the staged Planner contract. It must not collapse hard eligibility and ranking into one opaque recommendation score.

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

A PracticeActivity resolves:

```text
practice mode
exact content revision
canonical target refs
external task/question-family refs where material
Content Context ref where material
Presentation Class refs where material
variant where material
delivery condition where readiness-relevant
stimulus/source
response conditions
scaffolding/exposure state
primary_activity_purpose    TRAINING | DIAGNOSTIC | ASSESSMENT | READINESS
evidence_candidacy          NOT_EVIDENCE_CANDIDATE | ASSESSMENT_MAY_ADMIT
```

Primary purpose and evidence candidacy are orthogonal. `ASSESSMENT` means focused measurement/re-measurement is the product purpose; it still does not admit evidence by itself. `ASSESSMENT_MAY_ADMIT` is pre-attempt candidacy only; actual EvidenceFact admission remains an Assessment result after real assistance/exposure/evaluator/provenance conditions are known.

Evidence candidacy is part of the immutable activity/item configuration for an attempt. The client, evaluator, ranker, or Core API may not retroactively switch `NOT_EVIDENCE_CANDIDATE` to `ASSESSMENT_MAY_ADMIT` because the observed result is favorable.

There is no API field whose semantic meaning is “certification-contributing before Assessment”. Certification contribution is derived downstream from admitted EvidenceFacts and the applicable EvidenceRequirement.

Official family identity must not be reconstructed from a broad Skill Leaf when the leaf serves several external families.

Direct browsing may create an eligible activity but cannot satisfy unrelated target conditions. Activity creation fails closed or selects another eligible revision if learner-specific exposure/novelty/operational state makes the requested content ineligible.

## 6. Attempts

```text
POST   /v1/attempts
GET    /v1/attempts/{attempt_id}
PATCH  /v1/attempts/{attempt_id}
POST   /v1/attempts/{attempt_id}/submissions
```

The resource exposes the legal state from `04-application-flows.md`; this file does not define a second lifecycle diagram.

Writing drafts may mutate before submission under revision control. Submission is explicit, idempotent, and records actual:

- exact assigned content revision;
- assistance/scaffolding;
- exposure/retry context;
- delivery mode/input mode where material;
- timestamps/response provenance required by Assessment.

Family/context/presentation/purpose/candidacy semantics come from the immutable referenced content revision/work configuration rather than client-supplied reclassification at submission time.

Accepted submission corresponds to durable authoritative state before success acknowledgement. If productive evaluation or another required async continuation follows, its durable work identity/marker is committed or recoverably derivable before the accepted response can rely on it. Later content retirement/revalidation does not rewrite the Attempt's revision reference.

## 7. Evaluations

```text
GET /v1/evaluations/{evaluation_id}
```

Public output may expose:

- status;
- learner-meaningful criterion observations;
- feedback;
- uncertainty/quality state where useful;
- retry/unavailable state.

It does not expose chain-of-thought, secrets, or irrelevant provider internals.

## 8. Progress and gaps

```text
GET /v1/progress
GET /v1/gaps
```

`progress` exposes target conditions, per-skill current support/current certification state, certification history, evidence freshness, and unresolved requirements.

Current certification and certification history are distinct. A historical certification can remain visible while current status is `in_progress` because the corresponding present claim is stale, conflicting, insufficient, below requirement, or otherwise non-`SUPPORTED`. Only `09-PROGRESSION.md` determines whether a loss of current support constitutes regression.

`gaps` exposes learner GapEvaluation + explainable ActionIntent.

CoverageGap remains separate. The API does not invent one mastery percentage to replace claim states.

## 9. Review

```text
GET /v1/review-queue
```

Each item declares one semantic queue kind:

```text
knowledge_retrieval
error_remediation
re_evidence
```

and references canonical targets plus the recommended action intent/mode where available.

## 10. Mocks

```text
POST   /v1/mock-runs
GET    /v1/mock-runs/{mock_run_id}
```

A MockRun preserves:

- primary activity purpose `READINESS` for a normal mock;
- evidence candidacy for the configured run/sections;
- exact content revisions used by the run;
- resolved test variant;
- external family configuration for the run;
- Content Context distribution where material;
- material Presentation Class coverage where applicable;
- delivery mode/interaction conditions when material;
- full-test or selected-section scope;
- actual completion/abandonment state;
- section observations/readiness outputs at their valid scope.

Readiness purpose does not imply evidence admission. A mock Observation can contribute to a claim only when the configured run/section is an evidence candidate and normal Assessment policy admits the actual Observation.

Full mocks cannot mix Academic and GT Reading/Task 1 accidentally or silently omit a required Speaking part from a claimed whole-Speaking mock.

One Skill Retake preparation scopes to one existing skill; it does not create another Skill type or assert administrative eligibility.

## 11. Media

```text
POST   /v1/media-sources
GET    /v1/media-sources/{media_source_id}
POST   /v1/media-lessons
GET    /v1/media-lessons/{media_lesson_id}
```

MediaSource creation resolves provider identity/metadata, playability, rights, and transcript state. Learner/provider URL input is an untrusted reference and does not imply arbitrary fetch/download permission.

MediaLesson creation requires an eligible source plus valid canonical target and Practice Mode mapping. Generated/derived lesson content enters normal content revision/validation semantics before learner assignment.

## 12. Content supply and operations

This group exposes product/runtime content work without making the API, AI, or an operator the content-quality authority.

Representative semantic resources/operations:

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

Exact wire fields and any additional subresources belong to machine contracts. Each privileged content operation requires the applicable privileged operational capability defined by `04-application-flows.md`; this file does not define roles or a role-to-capability matrix.

These route semantics establish the boundary:

- content-revision reads expose identity, lineage/origin, canonical bindings, current applicable validation evidence, release/assignment/operational eligibility, and reconstructable provenance as authorized;
- `PATCH content-revisions` may change only authorized operational/release metadata owned downstream; it cannot mutate the immutable semantic payload of that revision;
- a material content correction creates a new ContentRevision through the applicable content-supply path rather than patching historical semantics;
- GenerationRequest represents a concrete supply demand, constraints, provenance/work identity, and resulting candidate refs; it is not evidence that generation is mandatory or that generated output is valid/active;
- validation-run resources represent execution/status/results of applicable validation work and preserve validator/policy provenance; a generator's self-check is only one possible signal;
- content reports preserve reporter/context/evidence sufficient for operational triage without mutating learner evidence or automatically proving the report correct;
- quarantine/stop-assignment, revalidation, release activation, replacement, and retirement behavior follow `04-application-flows.md`; this API must not invent a second combined ContentStatus lifecycle;
- authorization implementation determines which authenticated identities receive the applicable operational capabilities without redefining those capability meanings.

No API operation may provide a generic bypass that turns a known applicable hard validation failure into assignable learner content. Authorized operators repair the cause/change legitimate intended use, then revalidate under the applicable policy.

# Server event stream

```text
GET /v1/event-stream
```

SSE may deliver status/update notifications for evaluation, media analysis, diagnostics, mocks, content generation/validation operations, and derived plan/progress refresh.

Persistent truth remains queryable through normal resources; SSE is not state authority and reconnect is safe.

# Internal evaluator/content capability API

Normal caller: Go Core API only.

```text
POST /internal/v1/evaluations
POST /internal/v1/media-analyses
POST /internal/v1/content-generations       when this capability is implemented/needed
POST /internal/v1/content-validations       when this capability is implemented/needed
GET  /internal/v1/health
```

Generation/validation endpoints are optional capability boundaries. Their existence in semantic API design does not make AI generation or external model use a product coverage requirement.

## Evaluation request semantics

References include:

```text
evaluation/work identity
attempt identity
exact content revision identity
canonical target IDs
external task/question-family refs where material
Content Context ID where material
Presentation Class refs where material
variant where material
actual delivery/input condition where material
assessment/practice context
response or secure response reference
scaffold/exposure conditions
rubric/evaluator configuration reference
requested observation families
```

The evaluator consumes these identities; it does not infer/relabel the exam family, activity purpose, or evidence candidacy from free-form prompt/response content when authoritative item/work metadata already exists. Applicability is preserved: context-neutral content does not receive a fabricated Content Context/family/variant solely to make the request uniform, while any materially required identity remains present.

## Evaluation response semantics

Python returns bounded measurement output:

```text
evaluation identity
status
Observation candidates preserving content/target/family/context provenance
criterion measurements
transcript/acoustic/text-analysis refs where appropriate
uncertainty/quality flags
model/evaluator provenance
diagnostics
```

Python never returns authoritative learner certification, product support, Band advancement, evidence candidacy upgrades, content activation, or final DailyPlan state.

## Content generation semantics

Input is a bounded ContentDemand/GenerationRequest already constrained by Core API/product policy, such as:

```text
work identity
intended use/consequence
canonical target refs
official family/context/presentation refs where applicable
variant/delivery scope where material
difficulty/scaffold/diversity requirements
prohibited/reuse/source constraints
authorized source/context refs
```

Output is candidate semantic content plus origin/model/tool provenance and diagnostics. It is not an active ContentRevision until Core API persists the candidate/revision and applicable validation/release policy passes.

The generator may not invent unresolved canonical target meaning, relax required conditions, or classify its own output as product-supported.

## Content validation semantics

Input identifies the exact ContentRevision, intended use/consequence scope, validation policy/version, and authorized supporting material.

Output may contain validation findings/signals, similarity measurements, answer/rubric checks, language/construct checks, uncertainty, and validator provenance. It does not decide Assessment evidence admission, learner mastery, product coverage, or content release activation beyond the authority explicitly granted to the owning deterministic validation policy.

A generated candidate being checked by the same model/process that generated it does not create independent confidence merely because the process returns `pass`.

## Media-analysis semantics

Input contains permitted transcript/text/media metadata or another authorized reference. Output may propose segments, difficulty metadata, vocabulary, canonical targets, and generated prompts. Core API validates eligibility and normal content revision/validation requirements before saving assignable product state.

# Idempotency

Require an idempotency key/equivalent stable client operation for operations where retry can duplicate learner history or paid/provider work, including:

- diagnostic run creation;
- learning-session creation when duplicate creation matters;
- attempt submission;
- mock-run creation;
- media-source/media-lesson creation with provider/evaluator cost;
- content generation requests where retry could duplicate generated inventory/provider cost;
- content validation runs where retry could duplicate paid validation work.

Retry cannot create duplicate attempts, EvidenceFacts, paid evaluator/generator work, or semantically duplicated revision records for one logical accepted operation.

# Optimistic concurrency

Mutable draft-like resources, TargetProfile, and mutable operational/admin metadata use revision/version semantics where concurrent/stale updates are possible.

A stale mutation is rejected rather than silently overwriting newer state. ContentRevision semantic payload itself is immutable once established under `spec/10` semantics; optimistic concurrency is not permission to overwrite it.

# Domain results vs transport/operation failures

Domain/result semantics are represented in normal resource/operation outputs when the request itself executed correctly. Examples include:

```text
insufficient evidence
conflicting evidence
stale evidence
TargetProfile condition unresolved
target not supported for a fully resolved scope
CoverageGap
learner-specific content unavailable because novelty/independence fails
evaluation/work pending
```

These are not inherently HTTP errors.

Transport/operation failure semantics instead cover failures such as:

```text
malformed request
unauthenticated
unauthorized / forbidden capability
invalid transition / failed precondition
stale revision
idempotency conflict
rate limited
dependency unavailable
internal failure
```

Exact public failure-code catalog and HTTP mapping belong to the future machine contract. A transport error envelope may carry a stable non-sensitive code/title/message/details/correlation identity, but domain states must not be copied into that catalog merely for convenience.

# HTTP conventions

Conceptual target semantics include successful reads/creates/updates, accepted asynchronous operations, unauthenticated/forbidden access, malformed/contract-rejected input, concurrency/idempotency conflict, rate limiting, and infrastructure failure. Exact status mapping belongs to the machine contract.

A `5xx`-class infrastructure failure is never a learner score, learner weakness, CoverageGap, insufficient-evidence result, or fake content-quality judgment.

# Machine-contract evolution and compatibility

Each material machine boundary has exactly one contract authority. Contract changes are reviewed conceptually as **backward-compatible** or **breaking**; these are review concepts, not required public enums.

Generally backward-compatible changes may include:

- adding a genuinely optional/applicability-safe response field;
- adding a new independent operation that does not alter existing operation meaning.

Potentially breaking changes include:

- removing or renaming a field/operation consumed by deployed clients;
- narrowing previously valid input;
- changing enum/status/domain-result meaning;
- changing required/optional applicability in a way that changes valid semantics;
- changing stable identifier meaning;
- changing lifecycle/precondition semantics under an existing operation.

Rules:

1. transport optionality/requiredness may not violate canonical applicability merely to avoid a version change;
2. deployed public `/v1` behavior cannot silently break existing consumers;
3. a breaking deployed public contract requires an explicit version/rollout/migration strategy before activation;
4. internal contracts may evolve faster but still require consumer/provider compatibility during rollout;
5. generated bindings are regenerated/validated from the contract and never manually patched as interface authority;
6. compatibility/conformance/drift checks enter root verification once the relevant contract/deployed consumer exists;
7. deprecation or a new API version does not reinterpret historical ContentRevision, Attempt, Observation, EvidenceFact, or certification history;
8. event contracts, if later introduced, require explicit producer/consumer compatibility rules rather than relying on undocumented payload tolerance.

This file does not mandate one generic SemVer algorithm or code-generator library.

# Pagination

Unbounded history/library/content/report collections use cursor pagination. Small stable canonical catalogs may be returned in full.

# API anti-patterns

Forbidden without an explicit architecture change:

- browser calling evaluator/generator/validator directly;
- Next.js Route Handler/Server Action becoming a second product API or hidden domain-command owner;
- endpoint per UI button;
- provider/model names in public domain routes;
- independently handwritten mirror DTOs across runtimes;
- generated binding/client treated as semantic authority;
- fake score zero on evaluator failure;
- API shape becoming a second learner-state or content-quality definition;
- API-owned thresholds or hidden per-skill target constraints without a canonical/product owner;
- practice/UI action implicitly mutating TargetProfile/readiness;
- CoverageGap represented as learner GapEvaluation or infrastructure error;
- unresolved TargetProfile input represented as learner evidence insufficiency or fabricated product non-support;
- ranker returning an activity that failed hard eligibility;
- omitted content revision/family/context/presentation/delivery identity when that omission changes coverage, scoring, evidence, or target meaning;
- fabricated family/context/presentation/variant identity where the upstream semantic is legitimately not applicable;
- client/evaluator silently reclassifying immutable item family identity;
- client/evaluator/server retroactively upgrading evidence candidacy after observing performance;
- mutable PATCH of a historical ContentRevision semantic payload;
- generator/validator output directly activating learner content;
- generic admin validation bypass for a known hard failure;
- one combined API ContentStatus that hides validation vs release vs operational-safety dimensions;
- `completed diagnostic` represented as `complete learner baseline`;
- retrying an ambiguous/non-deduplicated mutation merely because the client timed out.