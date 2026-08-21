STATUS: CANONICAL
OWNS: end-to-end product/system flows across web, core API, evaluation service, learner state, media, and async result delivery
DEPENDS_ON: ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 01-skill-features.md, 03-media-youtube.md
DOES_NOT_OWN: API field schemas, learning/mastery truth, exact persistence topology, provider selection, or framework internals

# Application Flows

## Purpose

Define how the major product interactions traverse the system without assigning learning authority to transport or UI layers.

## Logical runtime units

The initial design has three application units:

```text
apps/web                 TypeScript learner/admin web experience
services/core-api        Go public API + deterministic learning orchestration
services/evaluator       Python AI/audio/text evaluation and analysis
```

The browser communicates with the Go Core API as the product authority boundary. It does not call the Python evaluator directly.

```text
Web
 │
 ▼
Core API
 ├── deterministic scoring / learner state / plans
 └── evaluator request
          │
          ▼
      Evaluator
          │
          ▼
   Observation result
          │
          ▼
      Core API
          │
          ▼
Evidence / Mastery / Gap / Plan
```

The evaluator returns observations and evaluation metadata. It does **not** certify Band, advance learner state, or select the final next action.

# Flow A — onboarding and diagnostic

```text
1. Web collects target + study constraints
2. Core API creates/updates LearnerProfile and Goal
3. Web requests a diagnostic run
4. Core API selects diagnostic activities from canonical targets
5. Learner completes attempts
6. Objective Listening/Reading items may be scored deterministically by Core API
7. Productive Writing/Speaking attempts are sent to Evaluator when automated evaluation is used
8. Evaluator returns criterion observations + provenance + uncertainty
9. Core API applies Assessment eligibility/interpretation
10. Core API updates MasteryEstimate / GapEvaluation
11. Core API creates first DailyPlan
12. Web renders current evidence state without inventing missing ability
```

Quick diagnostic and full baseline remain non-certifying unless the normal evidence contract independently supports a claim.

# Flow B — daily plan

```text
GET today
  ↓
Core API loads learner state
  ↓
Progression derives current GapEvaluation / ActionIntent
  ↓
Practice design resolves eligible mechanisms/modes
  ↓
time + learner constraints filter choices
  ↓
Core API returns DailyPlan
  ↓
Web renders explainable activity cards
```

Every recommended activity should include a machine-readable reason code such as:

```text
PREREQUISITE_GAP
ABILITY_GAP
STALE_EVIDENCE
TRANSFER_GAP
REVIEW_DUE
EXAM_CONDITION_GAP
```

plus human-readable explanation.

# Flow C — ordinary practice attempt

```text
1. Web starts PracticeActivity
2. Core API returns item/source + conditions + evidence-role label
3. Learner performs task
4. Web submits Attempt with actual assistance/scaffold/exposure metadata
5. Core API or Evaluator produces Observation
6. If training-only:
      generate feedback / remediation / review state
   If evidence-eligible:
      Assessment evaluates admissibility and scope
7. Progression updates learner-state interpretation when justified
8. Core API returns result + next action
9. Web updates the active session without pretending every activity changes Band
```

Attempt creation and submission must be idempotent so network retry cannot duplicate learner history or evaluator cost.

# Flow D — Writing or Speaking async evaluation

Productive evaluation may exceed interactive-request latency. The canonical flow is asynchronous from the browser's perspective.

```text
Web
  │ POST Attempt
  ▼
Core API
  │ persist attempt + create evaluation work
  │ return 202 / pending evaluation identity
  ▼
Evaluator
  │ criterion observations
  │ transcript/acoustic/text features where relevant
  │ model/rubric provenance + uncertainty
  ▼
Core API
  │ validate evaluation response
  │ Assessment → EvidenceFact / claim outcome
  │ Progression → Gap / state / next action
  ▼
Web receives completion via SSE or fetch refresh
```

Rules:

- pending is a valid state;
- timeout does not become a low score;
- failed evaluation is retriable without duplicating Attempt;
- evaluator provider/model changes must remain visible in provenance;
- no learner progression occurs from a failed or invalid evaluator response.

# Flow E — objective Listening / Reading attempt

For deterministic keyed items:

```text
Attempt
  ↓
answer-key validation
  ↓
item result
  ↓
Observation
  ↓
Assessment eligibility / scope
  ↓
EvidenceFact when valid
```

The evaluator service is not used merely because Python exists. Objective deterministic work stays in the owning Core API path unless a real reason requires otherwise.

# Flow F — feedback and remediation

```text
Observation(s)
  ↓
GapEvaluation
  ↓
ActionIntent
  ↓
ErrorPattern / RemediationPattern when relevant
  ↓
Learning Mechanism
  ↓
Practice Mode
  ↓
fresh Practice Item
```

Do not directly map `incorrect_answer → exercise_id`.

# Flow G — review

```text
review request
  ↓
classify queue item
  ├── knowledge retrieval
  ├── error remediation
  └── re-evidence
  ↓
select eligible activity
  ↓
Attempt
  ↓
update corresponding state only
```

The UI may show one Review tab; the backend preserves the distinct semantics.

# Flow H — YouTube/media lesson creation

```text
1. learner pastes URL
2. Web sends source request to Core API
3. Core API resolves provider/video metadata and eligibility
4. transcript state is established
5. authorized text, when available, may be sent to Evaluator for segmentation/analysis
6. Evaluator proposes segments, targets, difficulty, prompts
7. Core API validates practice-mode + target references + rights state
8. Web shows preview
9. learner saves MediaLesson
10. subsequent attempts use the normal Practice/Attempt pipeline
```

The system never requires the Python service to download arbitrary YouTube media.

# Flow I — full mock

```text
MockRun
  ↓
Listening section
Reading section
Writing tasks
Speaking session (same run or separately scheduled)
  ↓
section observations / scores
  ↓
ReadinessEvaluation
  ↓
GapEvaluation
  ↓
exam-preparation plan
```

A mock is broad readiness evidence. It does not bypass the ordinary certification policy.

# Result delivery

The web app supports two result paths:

### Immediate

Used for deterministic or low-latency results.

```text
submit → 200/201 result
```

### Pending

Used for Writing/Speaking/media analysis or other long work.

```text
submit → 202 pending
            ↓
      SSE status/result
```

Polling remains a fallback; SSE is the preferred first design because result delivery is server-to-client and does not require full bidirectional WebSocket semantics.

# Failure semantics

Failures must preserve semantic distinction:

- `invalid_attempt` — input/conditions do not support the requested operation;
- `evaluation_pending` — accepted but unfinished;
- `evaluation_unavailable` — evaluator could not produce a valid result;
- `insufficient_evidence` — assessment cannot support the claim yet;
- `source_unavailable` — media source cannot currently be used;
- `conflicting_evidence` — current admissible evidence conflicts;
- `stale_evidence` — refresh required.

Do not collapse these into generic "practice failed" or a score of zero.

# Cross-language ownership

```text
TypeScript
  owns interaction/rendering/browser capture

Go
  owns public API, learner/session/attempt orchestration,
  deterministic scoring, evidence policy execution,
  progression execution, and durable product state

Python
  owns bounded AI/audio/text evaluation and media analysis
```

The same semantic rule is not independently implemented in all three languages. Cross-language data crosses explicit contracts.