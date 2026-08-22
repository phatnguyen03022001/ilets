STATUS: CANONICAL
OWNS: public/internal API resource model, route groups, operation semantics, async API behavior, idempotency/error conventions, and contract-materialization rules
DEPENDS_ON: ../spec/01-LEARNER-MODEL.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 02-practice-catalog.md, 04-application-flows.md
DOES_NOT_OWN: learning truth, external task-family definitions, TargetProfile UX semantics, product coverage policy, evaluator algorithms, runtime lifecycle truth, framework selection, persistence schema, provider choice, or exact wire schema after machine contracts exist

# API

## Purpose

Define semantic API boundaries before implementation. The product exposes one learner-facing public API through Go Core API and one bounded internal evaluation API to Python.

# Principles

1. public version prefix `/v1`;
2. resource-oriented design over one-endpoint-per-button RPC;
3. stable canonical/external IDs cross boundaries unchanged;
4. creation/submission operations are idempotent where retry could duplicate history/cost;
5. long work returns durable pending resources rather than unbounded requests;
6. evidence states and CoverageGap are domain results, not generic failures;
7. browser never calls Python evaluator directly;
8. variant, official family, Content Context, material Presentation Class, and delivery are explicit wherever they change content, scoring, inference, coverage, or readiness behavior;
9. API resources expose lifecycle states owned by `04-application-flows.md` rather than redefining transition rules here;
10. once materialized, machine-readable contracts own exact wire shape;
11. API fields translate canonical/product semantics; they never invent thresholds, state meaning, applicability, or evidence authority that lacks an upstream owner.

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

Public and internal evaluator APIs may use separate contract files, but each boundary has exactly one exact machine authority.

Required checks once materialized:

- schema validation;
- generated-artifact drift checks where generation is used;
- stable canonical/external ID compatibility;
- public/internal boundary separation;
- consumer/provider conformance;
- deployed `/v1` compatibility policy;
- contract version/provenance in integration verification.

Machine schemas define transport shape, not IELTS learning/exam truth.

# Public resource groups

The initial semantic API surface has 11 resource groups.

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

A target-relative TargetProfile has at least one actual Band constraint. If no Band constraint is known yet, APIs may support diagnostic/foundational flows against an unresolved target context but must not fabricate a readiness target.

`selected_skill_retake` selects focused preparation only. Eligibility-sensitive responses keep missing original-test/timing/location/delivery/purpose conditions unresolved rather than treating the selected skill as proof of One Skill Retake eligibility.

`target-support` transports the current product-support result for the exact target scope, including applicable delivery/purpose/eligibility conditions and blocking CoverageGap classes. Product-support semantics remain owned by `08-coverage-and-support.md`; referencing that downstream result here does not make API the policy owner or create a reverse semantic-definition dependency.

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
canonical target refs
external task/question-family refs where material
Content Context ref
Presentation Class refs where material
variant
delivery condition where readiness-relevant
stimulus/source
response conditions
scaffolding/exposure state
primary_activity_purpose    TRAINING | DIAGNOSTIC | READINESS
evidence_candidacy          NOT_EVIDENCE_CANDIDATE | ASSESSMENT_MAY_ADMIT
```

Primary purpose and evidence candidacy are orthogonal. `ASSESSMENT_MAY_ADMIT` is pre-attempt candidacy only; actual EvidenceFact admission remains an Assessment result after real assistance/exposure/evaluator/provenance conditions are known.

Evidence candidacy is part of the immutable activity/item configuration for an attempt. The client, evaluator, ranker, or Core API may not retroactively switch `NOT_EVIDENCE_CANDIDATE` to `ASSESSMENT_MAY_ADMIT` because the observed result is favorable.

There is no API field whose semantic meaning is “certification-contributing before Assessment”. Certification contribution is derived downstream from admitted EvidenceFacts and the applicable EvidenceRequirement.

Official family identity must not be reconstructed from a broad Skill Leaf when the leaf serves several external families.

Direct browsing may create an eligible activity but cannot satisfy unrelated target conditions.

## 6. Attempts

```text
POST   /v1/attempts
GET    /v1/attempts/{attempt_id}
PATCH  /v1/attempts/{attempt_id}
POST   /v1/attempts/{attempt_id}/submissions
```

The resource exposes the legal state from `04-application-flows.md`; this file does not define a second lifecycle diagram.

Writing drafts may mutate before submission under revision control. Submission is explicit, idempotent, and records actual:

- assistance/scaffolding;
- exposure/retry context;
- delivery mode/input mode where material;
- timestamps/response provenance required by Assessment.

Family/context/presentation/purpose/candidacy semantics come from the immutable referenced item/work configuration rather than client-supplied reclassification at submission time.

Accepted submission corresponds to durable authoritative state before success acknowledgement.

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
- resolved test variant;
- external family configuration for the run;
- Content Context distribution;
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

MediaSource creation resolves provider identity/metadata, playability, rights, and transcript state. It does not imply copying/downloading media.

MediaLesson creation requires an eligible source plus valid canonical target and Practice Mode mapping.

# Server event stream

```text
GET /v1/event-stream
```

SSE may deliver status/update notifications for evaluation, media analysis, diagnostics, mocks, and derived plan/progress refresh.

Persistent truth remains queryable through normal resources; SSE is not state authority and reconnect is safe.

# Internal evaluator API

Normal caller: Go Core API only.

```text
POST /internal/v1/evaluations
POST /internal/v1/media-analyses
GET  /internal/v1/health
```

## Evaluation request semantics

References include:

```text
evaluation/work identity
attempt identity
canonical target IDs
external task/question-family refs where material
Content Context ID
Presentation Class refs where material
variant
actual delivery/input condition where material
assessment/practice context
response or secure response reference
scaffold/exposure conditions
rubric/evaluator configuration reference
requested observation families
```

The evaluator consumes these identities; it does not infer/relabel the exam family, activity purpose, or evidence candidacy from free-form prompt/response content when authoritative item/work metadata already exists.

## Evaluation response semantics

Python returns bounded measurement output:

```text
evaluation identity
status
Observation candidates preserving target/family/context provenance
criterion measurements
transcript/acoustic/text-analysis refs where appropriate
uncertainty/quality flags
model/evaluator provenance
diagnostics
```

Python never returns authoritative learner certification, product support, Band advancement, evidence candidacy upgrades, or final DailyPlan state.

## Media-analysis semantics

Input contains permitted transcript/text/media metadata or another authorized reference. Output may propose segments, difficulty metadata, vocabulary, canonical targets, and generated prompts. Core API validates eligibility before saving product state.

# Idempotency

Require an idempotency key/equivalent stable client operation for operations where retry can duplicate learner history or paid/provider work, including:

- diagnostic run creation;
- learning-session creation when duplicate creation matters;
- attempt submission;
- mock-run creation;
- media-source/media-lesson creation with provider/evaluator cost.

Retry cannot create duplicate attempts, EvidenceFacts, or paid evaluator work.

# Optimistic concurrency

Mutable draft-like resources and TargetProfile use revision/version semantics where concurrent/stale updates are possible.

A stale mutation is rejected rather than silently overwriting newer learner state.

# Error envelope

Transport failures use one stable conceptual envelope:

```text
code
title
message
details
trace_id
```

Representative codes:

```text
invalid_request
invalid_attempt
invalid_transition
stale_revision
source_unavailable
source_not_eligible
evaluation_pending
evaluation_unavailable
insufficient_evidence
conflicting_evidence
stale_evidence
target_not_supported
product_coverage_blocked
rate_limited
```

Evidence/coverage states are often successful domain results rather than HTTP errors.

# HTTP conventions

Target semantics:

- `200` read/update success;
- `201` created;
- `202` accepted/pending asynchronous work;
- `204` successful no-body operation where useful;
- `400` malformed request;
- `401` unauthenticated;
- `403` authenticated but unauthorized/ineligible by access policy;
- `404` absent/inaccessible resource;
- `409` lifecycle/concurrency/idempotency conflict;
- `422` structurally valid input violating operation contract;
- `429` rate limit;
- `5xx` infrastructure failure, never a learner score.

Exact mapping belongs to the machine contract.

# Pagination

Unbounded history/library collections use cursor pagination. Small stable canonical catalogs may be returned in full.

# API anti-patterns

Forbidden without an explicit architecture change:

- browser calling evaluator directly;
- endpoint per UI button;
- provider/model names in public domain routes;
- independently handwritten mirror DTOs across runtimes;
- fake score zero on evaluator failure;
- API shape becoming a second learner-state definition;
- API-owned thresholds or hidden per-skill target constraints without a canonical/product owner;
- practice/UI action implicitly mutating TargetProfile/readiness;
- CoverageGap represented as learner GapEvaluation;
- ranker returning an activity that failed hard eligibility;
- omitted official family/context/presentation/delivery identity when that omission changes coverage, scoring, evidence, or target meaning;
- client/evaluator silently reclassifying immutable item family identity;
- client/evaluator/server retroactively upgrading evidence candidacy after observing performance;
- `completed diagnostic` represented as `complete learner baseline`.