STATUS: CANONICAL
OWNS: public/internal API resource model, route groups, operation semantics, async API behavior, idempotency/error conventions, and contract-materialization rules
DEPENDS_ON: ../spec/01-LEARNER-MODEL.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 04-application-flows.md
DOES_NOT_OWN: learning truth, evaluator algorithms, framework selection, persistence schema, authentication provider, or exact generated wire schemas after machine-readable contracts exist

# API

## Purpose

Define the target API semantics before implementation. The product exposes one public learner-facing API boundary through the Go Core API and one bounded internal evaluation API to Python.

## API principles

1. public version prefix: `/v1`;
2. resource-oriented names over action-heavy RPC naming;
3. stable canonical IDs cross the wire unchanged;
4. creation/submission operations support idempotency;
5. long work returns a pending resource rather than holding an unbounded request;
6. missing/invalid/stale/conflicting evidence are domain states, not generic server errors;
7. Python internal routes are never called directly by browser clients;
8. once `contracts/http/openapi.yaml` exists, it becomes exact HTTP interface-shape authority; this document retains semantic ownership of API grouping/intent and must stop duplicating field-level schema detail.

# Public API resource groups

The initial API has **11 public resource groups**.

## 1. Identity and learner

```text
GET    /v1/me
GET    /v1/learner-profile
PATCH  /v1/learner-profile
PUT    /v1/learner-goal
```

Semantics:

- `/me` returns authenticated-user product identity/session context;
- learner profile exposes planning-relevant learner context, not hidden evaluator internals;
- learner goal owns target date/bands/study constraints.

Authentication-provider-specific endpoints are outside this API contract.

## 2. Diagnostics

```text
POST   /v1/diagnostic-runs
GET    /v1/diagnostic-runs/{diagnostic_run_id}
```

Creation request selects quick/full mode and optional focus constraints.

Run result distinguishes completed sampling from supported claims; a diagnostic run is not automatically certification.

## 3. Daily plan

```text
GET    /v1/daily-plan
```

Parameters may include requested duration preset and focus override.

Response conceptually contains:

- plan identity;
- generated-at state reference;
- activity cards;
- reason codes;
- estimated durations;
- canonical target references;
- evidence-role labels.

## 4. Learning sessions

```text
POST   /v1/learning-sessions
GET    /v1/learning-sessions/{learning_session_id}
PATCH  /v1/learning-sessions/{learning_session_id}
```

Session patch is limited to lifecycle/progress semantics such as in-progress/completed/abandoned and permitted client state. It does not mutate canonical learning definitions.

## 5. Practice catalog and activities

```text
GET    /v1/practice-modes
POST   /v1/practice-activities
GET    /v1/practice-activities/{practice_activity_id}
```

`practice-modes` exposes the 28 product practice modes from `02-practice-catalog.md`.

A PracticeActivity resolves a mode to concrete targets, source/stimulus, conditions, scaffolding, and item configuration.

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

Writing drafts may be saved before final submission. Submission is explicit so draft edits cannot accidentally become assessment evidence.

Each submission records actual attempt conditions, including scaffold/exposure metadata needed by Assessment.

## 7. Evaluation results

```text
GET    /v1/evaluations/{evaluation_id}
```

Public evaluation output exposes learner-meaningful criterion observations, feedback, status, and uncertainty where useful.

It does not expose raw model chain-of-thought, secrets, or provider internals irrelevant to learner trust.

Typical status:

```text
pending
running
completed
unavailable
invalid
```

## 8. Progress and gaps

```text
GET    /v1/progress
GET    /v1/gaps
```

`progress` exposes current per-skill state, supported band claims, certification history, and evidence freshness.

`gaps` exposes current GapEvaluation results and explainable ActionIntent recommendations.

The API must not fabricate a single scalar "mastery percentage" to replace the underlying claim/gap model.

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

and references the canonical target plus recommended practice action.

## 10. Mocks

```text
POST   /v1/mock-runs
GET    /v1/mock-runs/{mock_run_id}
```

A MockRun can represent a full IELTS mock or scoped exam-readiness section. It feeds readiness/gap semantics and never bypasses Assessment policy.

## 11. Media

```text
POST   /v1/media-sources
GET    /v1/media-sources/{media_source_id}
POST   /v1/media-lessons
GET    /v1/media-lessons/{media_lesson_id}
```

Creating a YouTube MediaSource resolves URL/provider metadata and transcript/rights state; it does not imply media copying.

Creating a MediaLesson requires an eligible source plus a valid practice-mode/target mapping.

# Server event stream

```text
GET /v1/event-stream
```

Transport: Server-Sent Events for product status updates such as:

- evaluation completed/unavailable;
- media lesson analysis completed;
- long diagnostic/mock phase updated;
- derived plan/progress refresh available.

SSE is a delivery optimization. Persistent learner truth remains queryable through normal resources.

Reconnect must be safe; clients may fall back to resource polling.

# Internal evaluator API

The Go Core API is the only normal caller.

```text
POST /internal/v1/evaluations
POST /internal/v1/media-analyses
GET  /internal/v1/health
```

## Evaluation request semantics

Input references:

- attempt identity;
- target IDs;
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

Python does not return an authoritative `certified=true` or mutate learner progression.

## Media analysis semantics

Input contains authorized/licensed transcript/text/media metadata or another permitted analysis reference.

Output may propose:

- segments;
- difficulty metadata;
- vocabulary candidates;
- practice-target candidates;
- generated prompt candidates.

Go validates product/domain eligibility before saving a MediaLesson.

# Idempotency

The following operations require an idempotency key or equivalent stable client operation identity:

- diagnostic-run creation;
- learning-session creation when network retry could duplicate a session;
- attempt submission;
- mock-run creation;
- media-source/media-lesson creation where evaluator/provider cost may be incurred.

Server retries must not create duplicate learner attempts or duplicate paid evaluator work.

# Optimistic concurrency

Mutable draft-like resources such as Writing Attempt drafts should expose a revision/version token.

A stale update is rejected rather than silently overwriting a newer learner draft.

# Error envelope

Transport errors use one stable envelope conceptually:

```text
code
title
message
details
trace_id
```

Representative domain-aware codes:

```text
invalid_request
invalid_attempt
stale_revision
source_unavailable
source_not_eligible
evaluation_pending
evaluation_unavailable
insufficient_evidence
conflicting_evidence
stale_evidence
rate_limited
```

`insufficient_evidence`, `conflicting_evidence`, and `stale_evidence` are often successful domain results rather than HTTP failures. The exact HTTP mapping belongs to the OpenAPI contract when materialized.

# HTTP status conventions

Target conventions:

- `200` resource/read/update success;
- `201` resource created;
- `202` accepted for asynchronous evaluation/analysis;
- `204` successful no-body operation where useful;
- `400` malformed request;
- `401` unauthenticated;
- `403` authenticated but unauthorized;
- `404` resource absent/inaccessible;
- `409` lifecycle/concurrency/idempotency conflict;
- `422` structurally valid request that violates operation contract;
- `429` rate limited;
- `5xx` infrastructure/service failure, never a learner score.

# Pagination

Attempt history, activity history, media libraries, and other unbounded collections use cursor pagination. Canonical Skill/Practice catalogs are small bounded resources and may be returned as complete lists.

# Contract materialization

When implementation begins:

```text
design/05-api.md
      ↓ semantic intent
contracts/http/openapi.yaml
      ↓ exact wire interface
TS generated client / validators
Go server binding / validation
Python internal client/server binding as relevant
```

The OpenAPI contract must not copy IELTS mastery rules as prose. It references stable enums/fields needed to carry semantics whose meaning remains owned by `spec/` and `design/`.

# API anti-patterns

Forbidden without a new architectural decision:

- browser calling evaluator directly;
- one endpoint per UI button;
- provider/model names embedded into public domain routes;
- duplicated Go/Python/TypeScript request models maintained independently when contract generation/validation is viable;
- returning a fake zero score when evaluation fails;
- using an API response shape as a second learner-state definition.