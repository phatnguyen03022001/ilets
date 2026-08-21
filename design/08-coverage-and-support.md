STATUS: CANONICAL
OWNS: product coverage semantics, TargetCoverageSpecification, CoverageGap taxonomy, support-promotion gates, official-construct-to-feature coverage mapping, and current product-support declarations
DEPENDS_ON: ../spec/00-PRODUCT.md, ../spec/02-IELTS-MODEL.md, ../spec/03-SKILLS.md, ../spec/05-BANDS.md, ../spec/06-CURRICULUM.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 01-skill-features.md, 02-practice-catalog.md, 03-media-youtube.md, 07-third-party-services.md
DOES_NOT_OWN: IELTS external truth, learning targets, feature behavior, provider selection, learner GapEvaluation, or empirical validation results

# Coverage and Product Support

## Purpose

Define exactly what it means for the product to claim that an IELTS target is modelled, covered, supported, or validated.

This document prevents statements such as “100% IELTS coverage” from being made because many documents, features, or exercises exist.

# Coverage is not a percentage

Coverage is evaluated as logical conditions over a scoped target.

```text
TargetCoverageSpecification
        ↓
material conditions
        ↓
CoverageGap evaluation
        ↓
MODELLED / COVERED
        ↓
TargetSupportDeclaration
        ↓
SUPPORTED_FOR_PRODUCT
        ↓
scoped empirical evidence
        ↓
VALIDATED
```

No weighted aggregate may hide a missing required condition.

# Product support statuses

## `MODELLED`

The construct/requirement is represented in canonical spec with enough semantic detail to reason about it end-to-end.

## `COVERED`

The complete executable learning path exists for the scoped target with no blocking CoverageGap:

```text
requirement
→ capability / knowledge / context
→ curriculum path
→ intervention / practice
→ evidence / re-evidence / transfer
→ learner experience / transition
→ content/assets
→ viable runtime/provider path
```

## `SUPPORTED_FOR_PRODUCT`

A versioned release declaration says the product actually exposes the target and all release-critical content, rights, privacy/security, reliability, evaluator/calibration, operational, and cost gates are satisfied.

## `VALIDATED`

Scoped empirical evidence supports the declared product outcome under named learner/product/evaluator/intervention versions and conditions.

Architecture coherence never implies `VALIDATED`.

# Coverage-condition statuses

Each material condition is independently one of:

- `UNKNOWN`;
- `DEFINED`;
- `PARTIAL`;
- `SATISFIED`;
- `BLOCKED`;
- `NOT_APPLICABLE`;
- `CALIBRATION_REQUIRED`.

# CoverageGap

A product coverage hole is not a learner weakness.

`CoverageGap` classes include:

- `MODEL_OR_SPEC`;
- `INTERVENTION_OR_ACTIVITY`;
- `CONTENT_OR_ASSET`;
- `EVIDENCE_OR_EVALUATOR`;
- `EXPERIENCE`;
- `TRANSITION`;
- `CONTRACT_OR_INTEGRATION`;
- `COST_OR_OPERATIONS`;
- `RIGHTS_PRIVACY_RELIABILITY`;
- `CALIBRATION_OR_VALIDATION`.

Each gap records scope, missing/failed condition, blocking consequence, dependencies, provenance/version, and required demand class.

# Target scope

## Complete standard IELTS construct target

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ all official task/question families
+ Band-3→9 learning interpretation
+ diagnostic → learning → practice → assessment → progression → exam-readiness paths
```

Academic may be released before General Training, but release sequencing is not construct completeness.

## UKVI Academic / General Training

UKVI reuses the corresponding language-test construct while adding administrative/security requirements outside the learning curriculum.

## One Skill Retake

One Skill Retake reuses one existing skill construct. Product support is focused goal/reassessment/exam-preparation flow plus current external eligibility/availability guidance, not a fifth Skill ontology.

## IELTS Life Skills

IELTS Life Skills remains outside current product-learning scope because it is a different Listening/Speaking-only pass/fail construct.

Therefore a future claim must be phrased as **complete standard IELTS Academic + General Training learning coverage**, not “every test sold under the IELTS brand.”

# Required TargetCoverageSpecification conditions

For a scoped `(variant, target-profile/band conditions)` target, coverage evaluation must test every applicable condition below.

| Condition | Required meaning |
|---|---|
| `construct_model` | external IELTS requirement is represented correctly |
| `skill_knowledge_model` | required capabilities/knowledge have stable canonical identity |
| `band_threshold` | target threshold/variant overlay is defined |
| `curriculum_route` | valid Band-3→target sequencing/variant route exists |
| `practice_intervention` | required acquisition/consolidation/retrieval/transfer/fluency/exam-readiness actions are executable where applicable |
| `feature_experience` | learner can perform the required interaction and understand state/result |
| `content_assets` | sufficient valid stimuli/items/templates/generators exist for required target/context diversity |
| `assessment_policy` | correct Assessment Type and versioned executable EvidenceRequirement exist |
| `evaluator_scoring` | deterministic scoring or calibrated productive evaluator route exists |
| `progression_transition` | evidence can drive explainable state/next-action behavior |
| `variant_context` | variant-specific task/section/context conditions are represented and sampled |
| `machine_contracts` | exact cross-unit interfaces exist for implemented boundaries |
| `rights_privacy_security` | applicable source/data/consent/security conditions pass |
| `reliability_recovery` | lifecycle, idempotency, retry, failure/recovery behavior pass |
| `accessibility_capture_quality` | learner can access interaction and low-quality capture is handled without false ability inference |
| `cost_abuse_operations` | release can operate within declared cost/rate/abuse constraints |
| `observability_audit` | consequential decisions preserve version/provenance/reason reconstruction |
| `validation` | empirical evidence exists when status is promoted to `VALIDATED` |

A condition may be `NOT_APPLICABLE` only with an explicit reason for the scoped target.

# Content coverage closure

A UI feature/table mapping is not content coverage.

When content is implemented, coverage tooling should consume the content coverage manifest or equivalent index defined by `spec/10-CONTENT-MODEL.md` and verify at least:

```text
canonical target refs
variant/task/section contexts
practice/assessment role
interaction implementation
answer/rubric/evaluator path
transfer/novelty classes
rights/provenance
independent readiness assets
release activation
```

Only then can a `CONTENT_OR_ASSET` condition become `SATISFIED`.

# Evidence policy closure

A high-consequence `(variant, skill, band)` claim cannot be product-supported merely because `08-ASSESSMENT.md` describes evidence principles.

Before support, the exact claim must resolve to a versioned executable `EvidenceRequirement`. Hidden LLM judgment, an unversioned threshold, or “two good attempts seems enough” is not an acceptable production policy.

# Contract closure

A multi-language runtime target cannot become `COVERED` while TypeScript, Go, and Python maintain independent handwritten interpretations of the same interface.

Implemented cross-unit boundaries require one exact machine contract authority under `contracts/`, generated/validated consumers, and conformance tests.

# Current coverage declaration — 2026-08-22 audit state

This is a design-time declaration, not a production support claim.

The canonical learning model now closes the major Academic/General Training semantic gap, including GT Reading context distribution and GT Writing Task-1 letter capabilities. Runtime/content/evaluator/contract/validation work is still missing.

| Target | Model status | Feature/practice path | Evidence/progression path | Runtime/content/calibration | Current product status |
|---|---|---|---|---|---|
| Academic Listening | strong model | defined | defined | not implemented/validated | `MODELLED`, not `COVERED` |
| Academic Reading | strong model | defined | defined | content/runtime not implemented/validated | `MODELLED`, not `COVERED` |
| Academic Writing | strong model | defined | variant-correct; productive EvidenceRequirement/evaluator must be materialized/calibrated | content/runtime/calibration missing | `MODELLED`, not `COVERED` |
| Academic Speaking | strong model | defined | defined; productive EvidenceRequirement/evaluator must be materialized/calibrated | audio/content/runtime/calibration missing | `MODELLED`, not `COVERED` |
| General Training Listening | shared strong model | shared path | shared path | not implemented/validated | `MODELLED`, not `COVERED` |
| General Training Speaking | shared strong model | shared path | shared path | not implemented/validated | `MODELLED`, not `COVERED` |
| General Training Reading | shared Reading skills + explicit GT section/context/score overlay modelled | variant-aware Reading modes/features defined | GT-specific claim conditions defined | GT assets/runtime not implemented/validated | `MODELLED`, not `COVERED` |
| General Training Writing Task 1 | dedicated `W-GT1-*` capability and Band/Curriculum overlay modelled | variant-aware planner/draft/timed path defined | GT-specific evidence non-substitution defined | assets/evaluator/runtime/calibration missing | `MODELLED`, not `COVERED` |
| General Training Writing Task 2 | shared model | shared Writing path | shared criterion path | not implemented/validated | `MODELLED`, not `COVERED` |
| One Skill Retake preparation | construct shared | focused-skill journey reusable | normal evidence semantics reusable | release/admin guidance workflow not implemented | `DEFINED` |
| IELTS Life Skills | out of scope | none | none | none | `NOT_APPLICABLE` to current target |

## Current top-level verdict

```text
Academic semantic model        STRONG / MODELLED
General Training semantic model STRONGER / MODELLED at learning-design level
Academic product coverage      NOT YET COVERED
General Training product coverage NOT YET COVERED
Runtime support                NOT IMPLEMENTED
Validated band-outcome claim   NOT ESTABLISHED
```

No user-facing copy may say a target is fully supported until its blocking conditions are closed and a versioned TargetSupportDeclaration exists.

# Official task/question-family mapping

This maps external families to product surfaces. It does not redefine external exam truth.

## Listening

| Official family | Product feature/practice path | Design condition |
|---|---|---|
| Multiple choice | `L-F01`, `L-F05`, `PM-L06` | `DEFINED` |
| Matching | `L-F01`, `L-F05`, `PM-L06` | `DEFINED` |
| Plan/map/diagram labelling | `L-F06`, `PM-L05` | `DEFINED` |
| Form/note/table/flow-chart completion | `L-F04`, `PM-L03` | `DEFINED` |
| Sentence completion | `L-F04`, `PM-L03` | `DEFINED` |
| Short answer | `L-F04`, `PM-L03` | `DEFINED` |

## Reading — shared interactions, variant content

| Official family | Product feature/practice path | Design condition |
|---|---|---|
| Multiple choice | `R-F01`, `R-F06`, `PM-R05`/`PM-R06` | `DEFINED` |
| True / False / Not Given | `R-F04`, `PM-R03` | `DEFINED` |
| Yes / No / Not Given | `R-F04`, `PM-R03` | `DEFINED` |
| Matching information | `R-F03`, `R-F07`, `PM-R02`/`PM-R06` | `DEFINED` |
| Matching headings | `R-F05`, `PM-R04` | `DEFINED` |
| Matching features/endings | `R-F03`, `R-F06`, `PM-R02`/`PM-R05` | `DEFINED` |
| Completion families | `R-F03`, `R-F06`, `PM-R02`/`PM-R05` | `DEFINED` |
| Diagram label completion | `R-F03`, `PM-R02` | `DEFINED` |
| Short answer | `R-F03`, `PM-R02`/`PM-R06` | `DEFINED` |

Variant closure additionally requires:

- Academic content corpus/transfer/readiness assets;
- GT Section-1 everyday, Section-2 workplace, Section-3 general-interest assets;
- correct variant raw-score/evidence policy.

## Writing

| Construct | Product path | Practice/evidence path | Design condition |
|---|---|---|---|
| Academic Task 1 visual | `W-F01`, `W-F03`, `W-F07`, `W-F08`, `W-F09` | `W-TA-*`; plan → draft → revision → timed independent | `DEFINED` |
| General Training Task 1 letter | same variant-aware Writing features | `W-GT1-*`; audience/purpose/register/bullet points → draft → revision → timed independent | `DEFINED` |
| Task 2 essay | `W-F01`, `W-F02`, `W-F06`, `W-F07`, `W-F08`, `W-F09` | component → full → timed independent | `DEFINED` |
| Writing criteria | criterion-level feedback/evaluator path | `AT-01` / `AT-05` | `CALIBRATION_REQUIRED` for high-consequence support |

## Speaking

| Construct | Product path | Practice/evidence path | Design condition |
|---|---|---|---|
| Part 1 | `S-F01` | `PM-S03` | `DEFINED` |
| Part 2 | `S-F02` | `PM-S04` | `DEFINED` |
| Part 3 | `S-F03` | `PM-S05` | `DEFINED` |
| Pronunciation | `S-F04`, `S-F05`, `S-F07` | `PM-S01`, `PM-S02` | `DEFINED`; calibration needed for quantitative claims |
| Whole Speaking | `S-F09` | `PM-S06` | `DEFINED`; evaluator/content/runtime missing |

# TargetProfile coverage invariant

A product target is not one number.

A TargetProfile may contain variant, overall target, per-skill minima, test date, receiving constraint, and selected One Skill Retake focus.

The product supports a target only if every required TargetProfile condition has a covered path.

If only an overall target is provided, the planner must not pretend this uniquely determines four skill targets. It must either collect real per-skill minima or use an explicitly labelled planning profile that preserves multiple valid score combinations.

# Route-to-target invariant

When a target is supported:

```text
TargetProfile
  ↓
provisional diagnostic
  ↓
material target gaps / unknowns
  ↓
Daily Plan eligibility
  ↓
ranked valid next action
  ↓
practice / review / re-evidence
  ↓
updated readiness conditions
  ↺
```

Ranking cannot make an invalid action valid. Required prerequisites, variant compatibility, evidence truth, and CoverageGaps remain upstream constraints.

# No guaranteed-band promise

`SUPPORTED_FOR_PRODUCT` means the product has a complete release-qualified path. It does **not** guarantee that a learner receives that external IELTS score.

Allowed:

- “Your current evidence supports Band N under this profile.”
- “These are the remaining blockers to your target.”
- “This target is supported by the current product path.”

Forbidden without suitable empirical basis:

- “Follow this plan and you will get Band 7.”
- “100% guaranteed target band.”
- “You improved 0.5 band from this exercise.”

# Demand outputs

Every blocking CoverageGap produces one demand class before implementation work is created:

- spec/model change;
- feature/interaction capability;
- content/asset/template/generator;
- evaluator/calibration;
- learner-flow/transition;
- contract/integration;
- provider/runtime/operations;
- rights/privacy/security;
- validation/research.

Only content gaps create content quantity demand.

# TargetSupportDeclaration minimum

A future declaration must name:

- target scope and variants;
- supported bands/TargetProfiles;
- product/version boundary;
- feature/practice coverage;
- content/asset coverage manifest version;
- EvidenceRequirement/policy versions;
- evaluator/calibration state;
- machine-contract versions;
- rights/privacy/security gates;
- third-party activation state;
- reliability/recovery state;
- accessibility/capture-quality gate;
- cost/abuse/operations gate;
- observability/audit gate;
- known non-blocking validation backlog;
- revocation conditions.

Support is versioned and revocable when construct, provider, rights, reliability, cost, calibration, or validation evidence changes.
