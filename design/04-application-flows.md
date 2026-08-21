STATUS: CANONICAL
OWNS: end-to-end product/system flows across web, core API, evaluation service, learner state, target route, media, async result delivery, planning-stage separation, and legal runtime lifecycle semantics
DEPENDS_ON: ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 01-skill-features.md, 03-media-youtube.md
DOES_NOT_OWN: API field schemas, learning/mastery truth, product coverage/support declaration, exact persistence topology, provider selection, or framework internals

# Application Flows

## Purpose

Define how major product interactions traverse the system without assigning learning authority to transport, UI, evaluator, ranker, or third-party layers.

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

# Planner decision contract

Planning is a staged decision. Implementations must preserve the stage boundaries because later ranking is not allowed to redefine earlier truth.

```text
1. TargetProfile + target-support check
        ↓
2. expand unresolved target conditions
        ↓
3. interpret learner evidence → GapEvaluation / ActionIntent
        ↓
4. resolve Required prerequisites + hard eligibility
        ↓
5. generate eligible learning / assessment candidates
        ↓
6. rank eligible candidates
        ↓
7. select coherent plan + stable explanation
```

## Stage 1 — target-support check

Resolve current product support for the exact TargetProfile. A blocking product CoverageGap remains a product condition; it is never converted into learner weakness.

## Stage 2 — target-condition expansion

Expand the TargetProfile into the per-skill/variant conditions still required. Overall-band planning may keep multiple valid score combinations; it must not invent hidden per-skill minima.

## Stage 3 — evidence interpretation

Consume Assessment/Progression outputs. Planning does not rescore evidence or create a second gap taxonomy.

## Stage 4 — hard eligibility

Filter out candidates that violate:

- Required prerequisites;
- target variant/task compatibility;
- product coverage/support constraints;
- evidence-role constraints;
- rights/privacy/source eligibility;
- learner accessibility constraints that make the activity unusable;
- lifecycle constraints such as already-submitted immutable work.

## Stage 5 — candidate generation

Generate one or more valid actions for the current ActionIntent. Candidate generation may use alternate Practice Modes or Assessment actions but cannot weaken the target.

## Stage 6 — ranking

Ranking may consider, among eligible candidates:

- target urgency/test date;
- expected decision/learning value;
- due review;
- time fit;
- learner preference/friction;
- fatigue/session coherence;
- exposure diversity/transfer value;
- operational cost after semantic validity is preserved.

### Ranker non-authority invariant

A ranker may reorder **eligible** candidates only. It may never:

- make an ineligible candidate eligible;
- bypass a Required prerequisite;
- alter a GapEvaluation or ReadinessEvaluation;
- reinterpret evaluator output;
- hide a CoverageGap;
- lower a target threshold;
- turn preference into ability evidence;
- certify a learner.

## Stage 7 — plan and explanation

The selected plan preserves the TargetProfile version/state reference and emits reason codes that reconstruct why each activity is present.

Deterministic tie-breaking should be stable enough that unchanged learner state does not create arbitrary plan churn.

# Flow A — target setup, onboarding, and diagnostic

```text
1. Web collects TargetProfile + study constraints
2. Core API validates the target against current product-support scope
3. Core API creates/updates LearnerProfile and TargetProfile
4. If target path is not product-supported, Web shows the blocking product CoverageGap truthfully
5. Web requests quick/full diagnostic
6. Core API resolves test variant and selects diagnostic activities from canonical targets
7. Learner completes attempts
8. Objective Listening/Reading items may be scored deterministically by Core API
9. Productive Writing/Speaking attempts use Evaluator when automated evaluation is eligible
10. Evaluator returns criterion observations + provenance + uncertainty
11. Core API applies Assessment eligibility/interpretation
12. Core API updates MasteryEstimate / GapEvaluation
13. Core API executes the Planner decision contract
14. Web renders learner evidence state without inventing missing ability
```

Quick diagnostic and full baseline remain non-certifying unless normal EvidenceRequirement rules independently support a claim.

# Flow B — route-to-target Daily Plan

```text
TargetProfile
  ↓
target support + unresolved target conditions
  ↓
GapEvaluation / ActionIntent
  ↓
Required prerequisite + variant eligibility
  ↓
eligible candidate generation
  ↓
rank eligible candidates
  ↓
coherent DailyPlan + reason codes
```

Representative reason codes:

```text
PREREQUISITE_GAP
ABILITY_GAP
INSUFFICIENT_EVIDENCE
CONFLICTING_EVIDENCE
STALE_EVIDENCE
SCAFFOLD_DEPENDENCE
TRANSFER_GAP
FLUENCY_GAP
REVIEW_DUE
EXAM_CONDITION_GAP
PRODUCT_COVERAGE_BLOCKED
```

`PRODUCT_COVERAGE_BLOCKED` is not a learner GapEvaluation.

Learner Swap/Skip/Shorten/Change-skill actions select among eligible actions only. Skipped target requirements remain unresolved.

# Flow C — ordinary practice attempt

```text
1. Web starts PracticeActivity
2. Core API returns item/source + variant/context + conditions + evidence-role label
3. learner performs task
4. Web submits Attempt with actual assistance/scaffold/exposure/delivery metadata
5. Core API or Evaluator produces Observation
6. if training-only:
      generate feedback / remediation / review state
   if evidence-eligible:
      Assessment evaluates admissibility and scope
7. Progression updates learner-state interpretation when justified
8. Core API derives next target-relevant action through the Planner contract
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
Observation with variant/context
  ↓
Assessment eligibility / scope
  ↓
EvidenceFact when valid
```

Reading Band inference uses the correct Academic/GT context and scoring policy. Python is not used merely because the repository has Python.

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
TargetProfile variant
  ↓
MockRun
  ↓
Listening shared
Reading variant-correct
Writing Task 1 variant-correct
Writing Task 2
Speaking shared (same run or separately scheduled)
  ↓
section observations / scores
  ↓
ReadinessEvaluation
  ↓
GapEvaluation
  ↓
exam-preparation plan
```

A mock is broad readiness evidence, not a certification shortcut. A mixed Academic/GT mock is invalid for normal readiness unless explicitly created as non-certifying comparison practice.

# Flow J — target reached / target unresolved

For each TargetProfile condition:

```text
current admissible evidence
  ↓
SUPPORTED | unresolved state
```

When all required target conditions are currently supported, the app may present **current evidence supports the declared target profile**.

It must not present **you are guaranteed to score this on test day**.

When a condition remains unresolved, the route continues through the appropriate action. When the product itself lacks a required path, the route stops with a product CoverageGap rather than manufacturing an activity.

# Legal lifecycle state machines

Exact wire fields belong to API contracts, but legal semantic transitions are owned here so different services do not invent incompatible lifecycle behavior.

## LearningSession

```text
planned → in_progress → completed
                    └→ abandoned
planned ─────────────→ abandoned
```

`completed` and `abandoned` are terminal for that session identity. Resume uses an allowed in-progress session or a new session according to product policy; it does not rewrite history.

## Attempt

```text
draft → submitted → evaluating → evaluated
                  ├────────────→ evaluated       deterministic path
                  └────────────→ invalid
submitted ─────────────────────→ invalid
```

Rules:

- `draft` may mutate only before submission under revision/concurrency rules;
- submission is idempotent;
- submitted learner work is immutable as that Attempt identity;
- correction/redraft creates a new version/attempt relation rather than rewriting submitted history;
- infrastructure failure is not `invalid` learner performance.

## Evaluation

```text
pending → running → completed
              ├──→ unavailable
              └──→ invalid
pending ──────────→ unavailable
pending ──────────→ invalid
```

`completed`, `unavailable`, and `invalid` are terminal for that evaluation identity. A retry after unavailable uses the same logical work identity with explicit attempt/retry provenance; it must not duplicate learner evidence.

## DiagnosticRun

```text
created → in_progress → completed
                    └→ abandoned
                    └→ unavailable
```

Completion means sampling flow finished, not certification.

## MockRun

```text
created → in_progress → completed
                    └→ abandoned
                    └→ invalid
```

A partial/abandoned mock may preserve valid section observations at their actual inference scope but cannot masquerade as a complete full-mock result.

## MediaSource / MediaLesson analysis

```text
requested → resolving/analyzing → ready
                         ├──────→ ineligible
                         └──────→ unavailable
```

`ineligible` is a rights/source/product condition, not learner failure.

## BandCertificationState

Current per-skill Band state follows `09-PROGRESSION.md`:

```text
not_started → in_progress → certified
                     ↑          │
                     └──────────┘  only when later admissible evidence establishes regression
```

Staleness alone does not execute the regression transition.

# Transition enforcement invariants

- Invalid lifecycle transitions fail closed.
- Retry of the same idempotent operation returns/reuses the same logical result where applicable.
- One service cannot advance another service-owned lifecycle by direct database mutation.
- Every terminal transition records enough provenance/reason to reconstruct why it occurred.
- Lifecycle state never substitutes for learning state: `completed` activity does not mean `mastered` target.

# Result delivery

Immediate deterministic results may use synchronous success.

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
- `source_ineligible`;
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
  Progression execution, planner eligibility/ranking orchestration,
  durable product state

Python
  bounded AI/audio/text evaluation and media analysis
```

The same semantic rule is not independently implemented in all three languages. Cross-language data crosses explicit contracts.
