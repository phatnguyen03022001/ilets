STATUS: CANONICAL
OWNS: end-to-end product/system flows across web, core API, evaluator, learner state, target route, media, async result delivery, planner-stage separation, hard eligibility, and legal runtime lifecycle semantics
DEPENDS_ON: ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 01-skill-features.md, 02-practice-catalog.md, 03-media-youtube.md
DOES_NOT_OWN: API field schemas, learning/mastery truth, product coverage declaration, exact persistence topology, provider selection, framework internals, or learner-facing UX defaults

# Application Flows

## Purpose

Define how major product interactions traverse runtime boundaries without assigning learning authority to transport, UI, evaluator, ranker, or third-party infrastructure.

# Logical runtime units

```text
apps/web                 TypeScript learner/admin experience
services/core-api        Go public API + deterministic orchestration
services/evaluator       Python bounded AI/audio/text evaluation
```

Normal product path:

```text
Web
 ↓
Core API
 ├─ durable learner/target/session/attempt state
 ├─ deterministic scoring/policy execution
 └─ internal evaluator request
          ↓
      Evaluator
          ↓
   Observation candidates + provenance
          ↓
      Core API
          ↓
Assessment / Progression / Planner
          ↓
Web
```

The evaluator does not certify Band, mutate learner progression, declare product coverage, or choose the final next action.

# Planner decision contract

Planning is a staged decision. Later stages may not redefine truth established by earlier stages.

```text
1. resolve TargetProfile + product-support scope
        ↓
2. expand unresolved target conditions
        ↓
3. consume Assessment/Progression → GapEvaluation / ActionIntent
        ↓
4. hard eligibility filtering
        ↓
5. candidate generation
        ↓
6. ranking among eligible candidates
        ↓
7. coherent plan + stable explanation
```

## Stage 1 — target/support resolution

Resolve the current TargetProfile version, including variant and delivery mode when material, and the product-support state for that exact scope.

A product CoverageGap remains a product condition; it is never converted into learner weakness.

## Stage 2 — target-condition expansion

Expand the target into unresolved per-skill/variant/external-purpose conditions.

An overall Band alone may correspond to multiple valid four-skill combinations. The planner must not invent hidden per-skill minima. Any working per-skill planning profile follows the separate non-authoritative planning-profile semantics in `00-learning-experience.md`.

## Stage 3 — evidence interpretation

Consume canonical outputs from Assessment and Progression. Planning does not rescore an Observation, reinterpret evaluator output, or invent another gap taxonomy.

## Stage 4 — hard eligibility

Remove candidates that violate any applicable hard condition:

- Required prerequisite;
- target variant/task compatibility;
- delivery-mode compatibility for exam-readiness work when material;
- product coverage/support constraint;
- primary-purpose/evidence-candidacy compatibility for the requested action;
- rights/privacy/source eligibility;
- accessibility/capture feasibility;
- immutable lifecycle condition;
- target-purpose/acceptance condition where the product knows the requested route cannot serve it.

Examples:

- Academic visual Task-1 practice is not eligible as GT Task-1 readiness work;
- typing-only mock behavior is not sufficient delivery-mode practice when the learner explicitly targets eligible Writing on Paper;
- IELTS Online Academic exam-readiness is not eligible for a GT target or a target purpose that requires a test-centre route;
- a failed microphone capture cannot become eligible Speaking evidence by ranking it highly;
- an activity configured `NOT_EVIDENCE_CANDIDATE` cannot be selected to satisfy a `COLLECT_EVIDENCE` action merely because its expected score is favorable.

## Stage 5 — candidate generation

Generate valid learning or assessment candidates for the current ActionIntent. Multiple modes may satisfy the same intent; candidate generation may vary delivery but not the target standard.

## Stage 6 — ranking

Ranking may consider only **eligible** candidates. Useful ranking signals include:

- target urgency/test date;
- expected learning or decision value;
- due review;
- time fit;
- learner preference/friction;
- fatigue/session coherence;
- transfer/exposure diversity;
- operational cost after semantic validity is preserved.

### Ranker non-authority invariant

A ranker may reorder eligible candidates. It may never:

- make an ineligible candidate eligible;
- bypass a Required prerequisite;
- alter GapEvaluation or ReadinessEvaluation;
- reinterpret evaluator output;
- hide a CoverageGap;
- lower a target threshold;
- convert preference into ability evidence;
- ignore a material variant/delivery constraint;
- upgrade evidence candidacy/admission;
- certify a learner.

## Stage 7 — plan/explanation

A plan records enough state/target references and reason information to reconstruct why every activity was eligible and selected.

Tie-breaking should be deterministic/stable enough that unchanged learner state does not create arbitrary plan churn.

# Flow A — target setup and diagnostic

```text
1. Web collects TargetProfile + planning constraints
2. Core API validates variant/delivery/purpose combination against known external/product support
3. Core API persists TargetProfile revision
4. unsupported product conditions are surfaced as CoverageGap
5. learner selects quick/full diagnostic UX shape
6. Core API selects variant-correct diagnostic items
7. learner attempts items
8. objective L/R items may score deterministically
9. eligible W/S attempts may use Evaluator
10. Evaluator returns observations + provenance + uncertainty
11. Assessment applies diagnostic sampling/eligibility semantics
12. Progression derives only justified learner-state interpretation
13. Planner executes stages 1–7
14. Web shows sampled and unresolved conditions truthfully
```

Diagnostic is a primary activity purpose. Whether a diagnostic Observation is admissible for a higher-consequence claim is a separate Assessment decision based on its pre-declared evidence candidacy and actual attempt conditions.

A completed diagnostic run is not synonymous with a complete learner model or certification.

# Flow B — Daily Plan

```text
TargetProfile
  ↓
product support + unresolved target conditions
  ↓
GapEvaluation / ActionIntent
  ↓
hard eligibility
  ↓
valid candidate generation
  ↓
ranking
  ↓
coherent DailyPlan + reason codes
```

Representative reason codes include:

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
DELIVERY_MODE_PREPARATION
PRODUCT_COVERAGE_BLOCKED
```

`PRODUCT_COVERAGE_BLOCKED` is not a learner GapEvaluation.

Swap/Skip/Shorten/Change-skill actions operate within eligible choices. They do not mark skipped requirements satisfied.

# Flow C — ordinary practice attempt

```text
1. Web starts PracticeActivity
2. Core API returns item + target + variant/context + conditions
   + primary_activity_purpose + evidence_candidacy
3. learner performs task
4. Web submits Attempt with actual scaffold/exposure/delivery metadata
5. Core API or Evaluator produces Observation
6. NOT_EVIDENCE_CANDIDATE → feedback/remediation/review handling only
   ASSESSMENT_MAY_ADMIT  → Assessment evaluates actual claim-scoped eligibility/inference
7. Progression updates interpretation only when justified
8. Planner derives next target-relevant action
9. Web presents outcome without fabricating a micro Band change
```

Primary purpose (`TRAINING`, `DIAGNOSTIC`, or `READINESS`) does not determine evidence admission. Assessment may reject a candidate Observation after seeing actual assistance/exposure/evaluator/provenance conditions; a favorable result cannot retroactively upgrade a non-candidate activity.

Submission/retry is idempotent where network repetition could duplicate history or cost.

# Flow D — Writing/Speaking asynchronous evaluation

```text
Web submits work
  ↓
Core API persists authoritative Attempt/work state
  ↓ COMMIT
ACK accepted/pending
  ↓
Evaluator executes bounded work
  ↓
criterion observations + provenance + uncertainty
  ↓
Core API validates evaluator response
  ↓
Assessment → EvidenceFact / ReadinessEvaluation
  ↓
Progression → learner-state interpretation
  ↓
Planner → next action
  ↓
Web receives status/result via SSE or resource refresh
```

Rules:

- accepted learner work is durable before learner-visible acceptance;
- pending/unavailable are valid non-score states;
- timeout/provider failure never becomes a low learner score;
- retries preserve one logical work/evaluation identity and provenance;
- no progression occurs from failed/invalid evaluator output;
- fallback must meet the same applicable quality/privacy floor.

# Flow E — objective Listening/Reading attempt

```text
Attempt
  ↓
answer-key + instruction/word-limit validation
  ↓
raw result
  ↓
Observation with variant/context/conditions
  ↓
Assessment eligibility/inference scope when evidence-candidate
  ↓
EvidenceFact when valid
```

Reading Band inference uses the correct Academic/GT scoring/context policy.

# Flow F — feedback/remediation

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

Direct `wrong answer → fixed exercise id` mapping is not a canonical remediation policy.

# Flow G — review

```text
review request
  ↓
queue kind
  ├─ knowledge retrieval
  ├─ error remediation
  └─ re-evidence
  ↓
eligible activity
  ↓
Attempt
  ↓
update only the corresponding semantic state
```

One Review screen may present these together; backend meaning remains distinct.

# Flow H — media lesson creation

```text
1. learner supplies source URL/reference
2. Core API resolves provider metadata and source eligibility
3. rights/transcript state is established
4. authorized text may be sent to Evaluator
5. Evaluator proposes segments/targets/difficulty/prompts
6. Core API validates rights + canonical target + practice mapping
7. learner previews/saves MediaLesson
8. later attempts use the normal Practice/Attempt path
```

Evaluator access never implies authorization to download/copy arbitrary media.

# Flow I — mock/readiness

```text
TargetProfile variant + delivery target when material
  ↓
MockRun
  ↓
Listening shared
Reading variant-correct
Writing Task 1 variant-correct
Writing Task 2
Speaking shared construct
  ↓
section observations / estimates
  ↓
ReadinessEvaluation where evidence candidacy + normal eligibility permit
  ↓
GapEvaluation
  ↓
exam-preparation plan
```

`READINESS` is the primary purpose of a normal mock; it is not an automatic evidence decision. A mock Observation contributes to a claim only when pre-declared as an evidence candidate and independently admitted by normal Assessment policy.

A mixed Academic/GT mock is invalid for normal full-test readiness unless explicitly created as non-certifying comparison practice.

Delivery-mode practice may change interaction conditions without changing scoring/Band semantics.

# Flow J — target supported/unresolved

For every TargetProfile condition:

```text
current admissible evidence
  ↓
SUPPORTED | unresolved evidence state
```

When all learner evidence conditions are supported, the app may state that current evidence supports the declared target profile **only if product-support wording remains separately truthful**.

It must never state that the learner is guaranteed an external result.

When product capability is missing, route generation stops at the CoverageGap rather than manufacturing an invalid activity.

# Legal lifecycle state machines

Exact wire fields belong to machine contracts. Legal semantic transitions are owned here so services cannot invent incompatible lifecycle rules.

## LearningSession

```text
planned → in_progress → completed
   │          └────────→ abandoned
   └───────────────────→ abandoned
```

`completed` and `abandoned` are terminal for that session identity.

## Attempt

```text
draft → submitted → evaluating → evaluated
                  ├────────────→ evaluated   deterministic path
                  └────────────→ invalid
submitted ─────────────────────→ invalid
```

Rules:

- draft content may mutate only before submission under concurrency/revision rules;
- submission is idempotent;
- submitted work is immutable for that Attempt identity;
- redraft/correction creates a new version/related attempt rather than rewriting submitted history;
- infrastructure failure is not `invalid` learner performance.

## Evaluation

```text
pending → running → completed
    │         ├────→ unavailable
    │         └────→ invalid
    ├──────────────→ unavailable
    └──────────────→ invalid
```

Terminal state applies to the evaluation identity. Retry preserves logical-work linkage and must not duplicate EvidenceFacts.

## DiagnosticRun

```text
created → in_progress → completed
                    ├→ abandoned
                    └→ unavailable
```

`completed` means the sampling flow ended, not that all learner claims are known.

## MockRun

```text
created → in_progress → completed
                    ├→ abandoned
                    └→ invalid
```

Partial valid section observations preserve their actual scope; they cannot masquerade as a complete mock result.

## Media resolution/analysis

```text
requested → resolving/analyzing → ready
                         ├──────→ ineligible
                         └──────→ unavailable
```

`ineligible` is a source/rights/product condition, never learner failure.

## BandCertificationState

Per skill/Band, current state follows `../spec/09-PROGRESSION.md`:

```text
not_started → in_progress ↔ certified
```

`certified` is valid only while the corresponding current claim is `SUPPORTED`. If the current claim becomes stale, conflicting, insufficient, below requirement, or otherwise non-`SUPPORTED`, current state returns to `in_progress` while certification history remains intact.

The transition reason preserves meaning. Staleness/conflict/insufficient evidence do **not** establish regression; regression exists only when Progression's regression rule is satisfied by later admissible evidence. Once a claim has evidence/history, it does not return to `not_started` merely because current support is lost.

# Transition invariants

- invalid transitions fail closed;
- retry of an idempotent operation reuses the same logical result/identity where applicable;
- one service does not advance another owner’s lifecycle by direct database mutation;
- consequential terminal transitions record reconstructable reason/provenance;
- runtime completion state never substitutes for learning state.

# Result delivery

Immediate deterministic work may return synchronously.

Long work uses the semantic pattern:

```text
submit → accepted/pending
             ↓
       SSE status/result
```

Polling/resource refresh remains fallback. Persistent truth remains queryable independently of SSE delivery.

# Failure semantics

Keep these distinct where applicable:

```text
invalid_attempt
evaluation_pending
evaluation_unavailable
insufficient_evidence
conflicting_evidence
stale_evidence
source_unavailable
source_ineligible
target_not_supported
product_coverage_blocked
```

Infrastructure failure is never represented as score zero or generic learner failure.

# Runtime ownership invariant

```text
TypeScript
  browser interaction/rendering/capture

Go
  public API + durable product state
  deterministic scoring
  Assessment policy execution
  Progression execution
  Planner orchestration

Python
  bounded AI/audio/text evaluation and media analysis
```

Cross-language data crosses explicit contracts. The same semantic rule is not independently maintained in multiple runtimes.