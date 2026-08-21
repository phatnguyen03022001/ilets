STATUS: CANONICAL
OWNS: public/internal API resource model, route groups, operation semantics, async API behavior, idempotency/error conventions, and contract-materialization rules
DEPENDS_ON: ../spec/01-LEARNER-MODEL.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 04-application-flows.md
DOES_NOT_OWN: learning truth, TargetProfile UX semantics, product coverage policy, evaluator algorithms, runtime lifecycle truth, framework selection, persistence schema, provider choice, or exact wire schema after machine contracts exist

# API

## Purpose

Define semantic API boundaries before implementation. The product exposes one learner-facing public API through Go Core API and one bounded internal evaluation API to Python.

# Principles

1. public version prefix `/v1`;
2. resource-oriented design over one-endpoint-per-button RPC;
3. stable canonical IDs cross boundaries unchanged;
4. creation/submission operations are idempotent where retry could duplicate history/cost;
5. long work returns durable pending resources rather than unbounded requests;
6. evidence states and CoverageGap are domain results, not generic failures;
7. browser never calls Python evaluator directly;
8. variant/task/context/delivery is explicit wherever it changes content, scoring, inference, target eligibility, or exam-readiness behavior;
9. API resources expose lifecycle states owned by `04-application-flows.md` rather than redefining transition rules here;
10. once materialized, machine-readable contracts own exact wire shape.

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
- stable ID/enum compatibility;
- public/internal boundary separation;
- consumer/provider conformance;
- deployed `/v1` compatibility policy;
- contract version/provenance in integration verification.

Machine schemas define transport shape, not IELTS learning truth.

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
target_overall_band
per-skill minimums
test_date
selected_skill_retake
revision/version
```

`target-support` reports product support for the exact target scope, including delivery/purpose conditions where applicable, plus blocking CoverageGap classes.

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
- a material target-condition sampling ledger that distinguishes at least:
  - `sampled`;
  - `not_sampled`;
  - `unusable`;
  - `pending_evaluation`;
- resulting evidence states/next evidence need where available.

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
- variant/context/delivery refs where material;
- evidence-role labels;
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
canonical targets
variant/task/context where material
delivery condition where exam-readiness material
stimulus/source
response conditions
scaffolding/exposure state
evidence-role label
```

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

`progress` exposes target conditions, per-skill current support/certification history, evidence freshness, and unresolved requirements.

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

- resolved test variant;
- delivery mode/interaction conditions when material;
- full-test or selected-section scope;
- actual completion/abandonment state;
- section observations/readiness outputs at their valid scope.

Full mocks cannot mix Academic and GT Reading/Task 1 accidentally.

One Skill Retake preparation scopes to one existing skill; it does not create another Skill type.

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
variant/task/context
actual delivery/input condition where material
assessment/practice context
response or secure response reference
scaffold/exposure conditions
rubric/evaluator configuration reference
requested observation families
```

## Evaluation response semantics

Python returns bounded measurement output:

```text
evaluation identity
status
Observation candidates
criterion measurements
transcript/acoustic/text-analysis refs where appropriate
uncertainty/quality flags
model/evaluator provenance
diagnostics
```

Python never returns authoritative learner certification, product support, Band advancement, or final DailyPlan state.

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
- practice/UI action implicitly mutating TargetProfile/readiness;
- CoverageGap represented as learner GapEvaluation;
- ranker returning an activity that failed hard eligibility;
- omitted variant/context/delivery when that omission changes scoring, evidence, or target meaning;
- `completed diagnostic` represented as `complete learner baseline`.