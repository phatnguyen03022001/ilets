STATUS: CANONICAL
OWNS: public/internal API resource model, route groups, operation semantics, async API behavior, idempotency/error conventions, and contract-materialization rules
DEPENDS_ON: ../spec/01-LEARNER-MODEL.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 04-application-flows.md
DOES_NOT_OWN: learning truth, TargetProfile product semantics, product coverage/support policy, evaluator algorithms, framework selection, persistence schema, authentication provider, or exact generated wire schemas after machine-readable contracts exist

# API

## Purpose

Define target API semantics before implementation. The product exposes one public learner-facing API boundary through the Go Core API and one bounded internal evaluation API to Python.

# API principles

1. public version prefix: `/v1`;
2. resource-oriented names over action-heavy RPC naming;
3. stable canonical IDs cross the wire unchanged;
4. creation/submission operations support idempotency;
5. long work returns a pending resource rather than holding an unbounded request;
6. missing/invalid/stale/conflicting evidence are domain states, not generic server errors;
7. product CoverageGap is distinct from learner GapEvaluation;
8. Python internal routes are never called directly by browser clients;
9. variant/task/context is explicit wherever it changes content, scoring, evidence, or route behavior;
10. once a machine-readable contract exists, it is the exact wire-shape authority and handwritten language mirrors are forbidden.

# Implementation-start contract gate

Semantic API design is not enough for parallel TypeScript/Go/Python implementation.

Before two runtime units independently implement the same boundary, the repository must materialize an exact contract under `contracts/`.

Minimum gate:

```text
design/05-api.md                 semantic intent
        ↓
contracts/http/openapi.yaml      exact HTTP wire shape
        ↓
validation / generated bindings
        ↓
Go / TypeScript / Python consumers
```

If public and internal evaluator APIs need separate contract documents, they must still have one exact authority per boundary.

Implementation may prototype inside one runtime owner before contract materialization. It may **not** let multiple languages create handwritten equivalent DTO/schema definitions and reconcile them later.

Required contract checks once materialized:

- schema validation;
- generated-client/server-binding drift detection where generation is used;
- stable ID/enum compatibility;
- public/internal boundary separation;
- representative consumer/provider conformance tests;
- backward-compatibility policy for deployed `/v1` changes;
- contract version/provenance visible in integration verification.

OpenAPI/JSON Schema carry wire shape, not IELTS learning truth. Field meaning continues to resolve to canonical `spec/`/`design/` owners.

# Public API resource groups

The initial API has **11 public resource groups**.

## 1. Identity, learner, and target

```text
GET    /v1/me
GET    /v1/learner-profile
PATCH  /v1/learner-profile
GET    /v1/target-profile
PUT    /v1/target-profile
GET    /v1/target-support
```

`target-profile` carries Academic/General Training, overall target, optional per-skill minima, target date, receiving-rule reference, and optional selected One Skill Retake focus.

`target-support` reports product support for the exact TargetProfile and blocking CoverageGap classes. It is not learner readiness.

Authentication-provider-specific endpoints are outside this API contract.

## 2. Diagnostics

```text
POST   /v1/diagnostic-runs
GET    /v1/diagnostic-runs/{diagnostic_run_id}
```

Creation selects quick/full mode and optional focus constraints. A result distinguishes completed sampling from supported claims.

## 3. Daily plan

```text
GET    /v1/daily-plan
```

Parameters may include requested duration preset and eligible focus override.

Response conceptually contains:

- plan identity;
- TargetProfile reference/version;
- generated-at learner-state reference;
- activity cards;
- reason codes;
- estimated durations;
- canonical target references;
- variant/context refs when material;
- evidence-role labels;
- unresolved target conditions;
- product CoverageGap indicator where applicable.

A DailyPlan is the output of the staged Planner decision contract in `04-application-flows.md`. The API must not collapse eligibility and ranking into one opaque recommendation field.

## 4. Learning sessions

```text
POST   /v1/learning-sessions
GET    /v1/learning-sessions/{learning_session_id}
PATCH  /v1/learning-sessions/{learning_session_id}
```

Session mutation follows the legal lifecycle in `04-application-flows.md`. Lifecycle completion never means target mastery.

## 5. Practice catalog and activities

```text
GET    /v1/practice-modes
POST   /v1/practice-activities
GET    /v1/practice-activities/{practice_activity_id}
```

`practice-modes` exposes the 28 product modes from `02-practice-catalog.md`.

A PracticeActivity resolves mode, TargetProfile variant when material, canonical targets, source/stimulus, conditions, scaffolding, and item configuration.

## 6. Attempts

```text
POST   /v1/attempts
GET    /v1/attempts/{attempt_id}
PATCH  /v1/attempts/{attempt_id}
POST   /v1/attempts/{attempt_id}/submissions
```

Lifecycle:

```text
draft
→ submitted
→ evaluating | evaluated | invalid
```

Detailed legal transitions are owned by `04-application-flows.md`.

Writing drafts may be saved before final submission. Submission is explicit so draft edits cannot accidentally become assessment evidence.

Each submission records actual attempt conditions, including scaffold/exposure/delivery metadata needed by Assessment.

A learner-visible accepted submission must correspond to durable authoritative state before success acknowledgement.

## 7. Evaluation results

```text
GET    /v1/evaluations/{evaluation_id}
```

Public output exposes learner-meaningful criterion observations, feedback, status, and uncertainty where useful.

Typical status:

```text
pending
running
completed
unavailable
invalid
```

It does not expose raw model chain-of-thought, secrets, or irrelevant provider internals.

## 8. Progress and gaps

```text
GET    /v1/progress
GET    /v1/gaps
```

`progress` exposes TargetProfile, per-skill state, supported Band claims, certification history, evidence freshness, and remaining target conditions.

`gaps` exposes learner GapEvaluation results and explainable ActionIntent recommendations.

A product CoverageGap is separate. The API must not fabricate a scalar mastery percentage to replace the claim/gap model.

## 9. Review

```text
GET    /v1/review-queue
```

Each queue item declares kind:

```text
knowledge_retrieval
error_remediation
re_evidence
```

and references the canonical target plus recommended action.

## 10. Mocks

```text
POST   /v1/mock-runs
GET    /v1/mock-runs/{mock_run_id}
```

A MockRun can be full IELTS or scoped section readiness.

For a full mock, the API must preserve one resolved test variant so Reading and Writing Task 1 cannot accidentally mix Academic and General Training.

For One Skill Retake preparation, a run may be scoped to one existing skill without inventing a fifth Skill ontology.

## 11. Media

```text
POST   /v1/media-sources
GET    /v1/media-sources/{media_source_id}
POST   /v1/media-lessons
GET    /v1/media-lessons/{media_lesson_id}
```

Creating a MediaSource resolves URL/provider metadata and transcript/rights state; it does not imply media copying.

Creating a MediaLesson requires an eligible source plus valid mode/target mapping.

# Server event stream

```text
GET /v1/event-stream
```

SSE may deliver evaluation/media/diagnostic/mock/derived-state updates. Persistent truth remains queryable through normal resources and reconnect is safe.

# Internal evaluator API

The Go Core API is the only normal caller.

```text
POST /internal/v1/evaluations
POST /internal/v1/media-analyses
GET  /internal/v1/health
```

## Evaluation request semantics

Input references:

- evaluation/work identity;
- attempt identity;
- target IDs;
- variant/task/context;
- assessment/practice context;
- response or secure response reference;
- actual conditions/scaffold state;
- rubric/evaluator configuration reference;
- requested observation families.

## Evaluation response semantics

Python returns:

```text
evaluation identity
status
Observation[]
criterion-level measurements
transcript / acoustic or textual analysis refs where appropriate
uncertainty / quality flags
model/evaluator provenance
diagnostics
```

Python does not return authoritative `certified=true`, product support, Band advancement, or learner progression state.

## Media analysis semantics

Input contains authorized/licensed transcript/text/media metadata or another permitted analysis reference.

Output may propose segments, difficulty metadata, vocabulary candidates, practice-target candidates, and generated prompt candidates. Go validates product/domain eligibility before saving a MediaLesson.

# Idempotency

The following require an idempotency key or equivalent stable client-operation identity:

- diagnostic-run creation;
- learning-session creation where retry could duplicate a session;
- attempt submission;
- mock-run creation;
- media-source/media-lesson creation where evaluator/provider cost may be incurred.

Server retries must not create duplicate learner attempts, duplicate evidence, or duplicate paid evaluator work.

# Optimistic concurrency

Mutable draft-like resources such as Writing drafts and editable TargetProfile expose revision/version semantics when concurrent updates are possible.

A stale update is rejected rather than silently overwriting a newer learner draft/target.

# Error envelope

Transport errors use one stable envelope conceptually:

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

Evidence/coverage states are often successful domain results rather than HTTP failures. Exact HTTP mapping belongs to OpenAPI when materialized.

# HTTP status conventions

Target conventions:

- `200` resource/read/update success;
- `201` resource created;
- `202` accepted for asynchronous work;
- `204` successful no-body operation where useful;
- `400` malformed request;
- `401` unauthenticated;
- `403` authenticated but unauthorized;
- `404` resource absent/inaccessible;
- `409` lifecycle/concurrency/idempotency conflict;
- `422` structurally valid request violating an operation contract;
- `429` rate limited;
- `5xx` infrastructure/service failure, never a learner score.

# Pagination

Attempt history, activity history, media libraries, and other unbounded collections use cursor pagination. Canonical small catalogs may be returned completely.

# Contract materialization

When implementation begins:

```text
design/05-api.md
      ↓ semantic intent
contracts/http/openapi.yaml
      ↓ exact wire interface
TS generated client / validators
Go server binding / validation
Python internal binding as relevant
```

Do not start independent cross-language DTO authoring and defer contract creation to cleanup.

# API anti-patterns

Forbidden without a new architectural decision:

- browser calling evaluator directly;
- one endpoint per UI button;
- provider/model names embedded into public domain routes;
- duplicated Go/Python/TypeScript request models maintained independently when contract generation/validation is viable;
- returning fake zero when evaluation fails;
- using API shape as a second learner-state definition;
- allowing practice/UI override to mutate readiness/TargetProfile implicitly;
- representing a product CoverageGap as a learner GapEvaluation;
- letting a ranker return an activity that failed hard eligibility;
- omitting variant/context from data where it changes scoring or evidence meaning.
