STATUS: CANONICAL
OWNS: end-to-end product/system flows across web, core API, evaluation service, learner state, target route, media, and async result delivery
DEPENDS_ON: ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 01-skill-features.md, 03-media-youtube.md
DOES_NOT_OWN: API field schemas, learning/mastery truth, product coverage/support declaration, exact persistence topology, provider selection, or framework internals

# Application Flows

## Purpose

Define how major product interactions traverse the system without assigning learning authority to transport, UI, evaluator, or third-party layers.

## Logical runtime units

```text
apps/web                 TypeScript learner/admin web experience
services/core-api        Go public API + deterministic learning orchestration
services/evaluator       Python AI/audio/text evaluation and analysis
```

The browser communicates with Go Core API as the product authority boundary. It does not call Python Evaluator directly.

```text
Web
 │
 ▼
Core API
 ├── target / learner / plan / deterministic scoring
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

Evaluator returns observations and provenance. It does not certify Band, advance learner state, declare product coverage, or select the final next action.

# Flow A — target setup, onboarding, and diagnostic

```text
1. Web collects TargetProfile + study constraints
2. Core API validates the target against current product-support scope
3. Core API creates/updates LearnerProfile and TargetProfile
4. If target path is not product-supported, Web shows the blocking product CoverageGap truthfully
5. Web requests quick/full diagnostic
6. Core API selects diagnostic activities from canonical targets
7. Learner completes attempts
8. Objective Listening/Reading items may be scored deterministically by Core API
9. Productive Writing/Speaking attempts use Evaluator when automated evaluation is eligible
10. Evaluator returns criterion observations + provenance + uncertainty
11. Core API applies Assessment eligibility/interpretation
12. Core API updates MasteryEstimate / GapEvaluation
13. Core API creates first DailyPlan toward the unchanged TargetProfile
14. Web renders learner evidence state without inventing missing ability
```

Quick diagnostic and full baseline remain non-certifying unless normal evidence rules independently support a claim.

# Flow B — route-to-target Daily Plan

```text
TargetProfile
  ↓
Core API loads learner state
  ↓
resolve target conditions still unsatisfied / unknown / stale
  ↓
Progression derives GapEvaluation / ActionIntent
  ↓
Required prerequisites filter eligible actions
  ↓
Practice design resolves eligible mechanisms/modes
  ↓
time + accessibility + learner constraints rank choices
  ↓
Core API returns DailyPlan
  ↓
Web renders strong explainable recommendation
```

Reason codes may include:

```text
PREREQUISITE_GAP
ABILITY_GAP
INSUFFICIENT_EVIDENCE
CONFLICTING_EVIDENCE
STALE_EVIDENCE
TRANSFER_GAP
FLUENCY_GAP
REVIEW_DUE
EXAM_CONDITION_GAP
PRODUCT_COVERAGE_BLOCKED
```

`PRODUCT_COVERAGE_BLOCKED` is not a learner GapEvaluation; it tells the learner the product cannot truthfully provide the required route yet.

Learner Swap/Skip/Shorten/Change-skill actions choose among eligible actions only. They never mark a Required prerequisite or target condition satisfied.

# Flow C — ordinary practice attempt

```text
1. Web starts PracticeActivity
2. Core API returns item/source + conditions + evidence-role label
3. learner performs task
4. Web submits Attempt with actual assistance/scaffold/exposure metadata
5. Core API or Evaluator produces Observation
6. if training-only:
      generate feedback / remediation / review state
   if evidence-eligible:
      Assessment evaluates admissibility and scope
7. Progression updates learner-state interpretation when justified
8. Core API derives next target-relevant action
9. Web updates session without pretending every activity changes Band
```

Attempt creation/submission is idempotent so network retry cannot duplicate learner history or evaluator cost.

# Flow D — Writing or Speaking async evaluation

```text
Web
  │ POST submission
  ▼
Core API
  │ persist authoritative Attempt/work state
  │ COMMIT
  │ ACK 202 + evaluation identity
  ▼
Evaluator
  │ criterion observations
  │ transcript/acoustic/text features when relevant
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

- accepted learner work is durable before learner-visible success;
- pending is valid;
- timeout/provider failure does not become a low score;
- retries reuse one work/evaluation identity;
- evaluator provider/model remains visible in provenance;
- no progression occurs from failed/invalid evaluator output;
- fallback must satisfy the same approved quality/privacy floor or the work remains delayed/unavailable.

# Flow E — objective Listening / Reading attempt

```text
Attempt
  ↓
answer-key / instruction / word-limit validation
  ↓
item result
  ↓
Observation
  ↓
Assessment eligibility / scope
  ↓
EvidenceFact when valid
```

Python is not used merely because the repository has Python. Deterministic work stays in Core API when it satisfies the semantic requirement.

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

The UI may show one Review tab; backend semantics stay distinct.

# Flow H — YouTube/media lesson creation

```text
1. learner pastes URL
2. Web sends source request to Core API
3. Core API resolves provider/video metadata and eligibility
4. transcript/rights state is established
5. authorized text, when available, may be sent to Evaluator for analysis
6. Evaluator proposes segments, targets, difficulty, prompts
7. Core API validates practice-mode + target references + rights state
8. Web shows preview
9. learner saves MediaLesson
10. later attempts use normal Practice/Attempt pipeline
```

Python never receives an implied authorization to download arbitrary YouTube media.

# Flow I — full mock

```text
MockRun
  ↓
Listening
Reading
Writing
Speaking (same run or separately scheduled)
  ↓
section observations / scores
  ↓
ReadinessEvaluation
  ↓
GapEvaluation
  ↓
exam-preparation plan
```

A mock is broad readiness evidence, not a certification shortcut.

# Flow J — target reached / target unresolved

For each TargetProfile condition:

```text
current admissible evidence
  ↓
SUPPORTED | unresolved state
```

When all required target conditions are currently supported, the app may present **current evidence supports the declared target profile**.

It must not present **you are guaranteed to score this on test day**.

When a condition remains unresolved, the route continues through the appropriate action: learn/remediate, collect evidence, resolve conflict, refresh stale evidence, expand transfer, build fluency, or exam preparation.

When the product itself lacks a required path, the route stops with a product CoverageGap rather than manufacturing an activity.

# Result delivery

Immediate deterministic results may use normal synchronous success.

Long productive/media work uses:

```text
submit → 202 pending
            ↓
      SSE status/result
```

Polling remains fallback. Persistent truth is queryable through normal resources.

# Failure semantics

Failures/states preserve distinctions:

- `invalid_attempt`;
- `evaluation_pending`;
- `evaluation_unavailable`;
- `insufficient_evidence`;
- `source_unavailable`;
- `conflicting_evidence`;
- `stale_evidence`;
- `target_not_supported` / product coverage blocked.

Do not collapse these into generic “practice failed” or score zero.

# Cross-language ownership

```text
TypeScript
  interaction/rendering/browser capture

Go
  public API, target/learner/session/attempt orchestration,
  deterministic scoring, Assessment policy execution,
  Progression execution, durable product state

Python
  bounded AI/audio/text evaluation and media analysis
```

The same semantic rule is not independently implemented in all three languages. Cross-language data crosses explicit contracts.