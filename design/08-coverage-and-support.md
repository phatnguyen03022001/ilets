STATUS: CANONICAL
OWNS: product coverage semantics, TargetCoverageSpecification, CoverageGap taxonomy, support-promotion gates, official-construct-to-feature coverage mapping, and current product-support declarations
DEPENDS_ON: ../spec/00-PRODUCT.md, ../spec/02-IELTS-MODEL.md, ../spec/03-SKILLS.md, ../spec/05-BANDS.md, ../spec/06-CURRICULUM.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, 00-learning-experience.md, 01-skill-features.md, 02-practice-catalog.md, 03-media-youtube.md, 07-third-party-services.md
DOES_NOT_OWN: IELTS external truth, learning targets, feature behavior, provider selection, learner GapEvaluation, or empirical validation results

# Coverage and Product Support

## Purpose

Define exactly what it means for the product to claim that an IELTS target is modelled, covered, supported, or validated.

This document exists to prevent statements such as “100% IELTS coverage” from being made because many documents, features, or exercises exist.

## Coverage is not a percentage

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

## Product support statuses

### `MODELLED`

The construct/requirement is represented in canonical spec with enough semantic detail to reason about it.

### `COVERED`

The complete executable learning path exists for the scoped target with no blocking CoverageGap:

```text
requirement
→ capability / knowledge / context
→ curriculum path
→ intervention / practice
→ evidence / re-evidence / transfer
→ learner experience / transition
→ viable runtime/provider path
```

### `SUPPORTED_FOR_PRODUCT`

A versioned release declaration says the product actually exposes the target and all release-critical content, rights, privacy/security, reliability, evaluator/calibration, and cost gates are satisfied.

### `VALIDATED`

Scoped empirical evidence supports the declared product outcome under named learner/product/evaluator/intervention versions and conditions.

Architecture coherence never implies `VALIDATED`.

## Coverage-condition statuses

Each material condition is independently one of:

- `UNKNOWN`;
- `DEFINED`;
- `PARTIAL`;
- `SATISFIED`;
- `BLOCKED`;
- `NOT_APPLICABLE`;
- `CALIBRATION_REQUIRED`.

## CoverageGap

A product coverage hole is not a learner weakness.

`CoverageGap` classes include:

- `MODEL_OR_SPEC`;
- `INTERVENTION_OR_ACTIVITY`;
- `CONTENT_OR_ASSET`;
- `EVIDENCE_OR_EVALUATOR`;
- `EXPERIENCE`;
- `TRANSITION`;
- `COST_OR_OPERATIONS`;
- `RIGHTS_PRIVACY_RELIABILITY`;
- `CALIBRATION_OR_VALIDATION`.

Each gap records scope, missing/failed condition, blocking consequence, dependencies, provenance/version, and the required demand class.

# Target scope

## Standard IELTS construct target

The product's complete standard-IELTS target is:

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ all official task/question families
+ band/criterion semantics needed for learner planning
+ diagnostic → learning → practice → assessment → progression → exam-readiness paths
```

Academic may be released before General Training, but the repository must not confuse release sequencing with construct completeness.

## UKVI Academic / General Training

IELTS for UKVI Academic and General Training use the same language-test construct/results as their standard counterparts while adding security/administrative requirements outside the learning construct. The learning system therefore reuses Academic/GT coverage; UKVI-specific booking/venue/compliance workflow is not a new learning curriculum.

Any visa-route minimum-score requirement belongs in a learner `TargetProfile` as an external constraint with provenance rather than being hard-coded as IELTS learning truth.

## One Skill Retake

One Skill Retake reuses one of the existing four skill constructs. It requires no new Skill ontology. Product support is represented by focused goal/reassessment/exam-preparation flows for the selected skill.

## IELTS Life Skills

IELTS for UKVI Life Skills A1/A2/B1 is **not currently in product scope**. It is a different Listening/Speaking-only pass/fail construct and does not fit the current Band-3-to-9 curriculum model.

Therefore the repository must not claim “100% of every test sold under the IELTS brand.” The intended claim, once all gates pass, is **complete standard IELTS Academic + General Training learning coverage**.

# Current coverage declaration — 2026-08-22 design state

This is a design-time declaration, not a production support claim.

| Target | Model status | Feature/practice path | Evidence/progression path | Runtime/content/calibration | Current product status |
|---|---|---|---|---|---|
| Academic Listening | strong model | defined | defined | not implemented/validated | `MODELLED`, not `COVERED` |
| Academic Reading | strong model | defined, task-family mapping requires explicit closure | defined | not implemented/validated | `MODELLED`, not `COVERED` |
| Academic Writing | strong model | defined | defined; productive evaluator unimplemented | calibration/content/runtime missing | `MODELLED`, not `COVERED` |
| Academic Speaking | strong model | defined | defined; productive evaluator unimplemented | calibration/audio/runtime missing | `MODELLED`, not `COVERED` |
| General Training Listening | shared with Academic | shared path | shared path | not implemented/validated | `MODELLED`, not `COVERED` |
| General Training Speaking | shared with Academic | shared path | shared path | not implemented/validated | `MODELLED`, not `COVERED` |
| General Training Reading | external format partly represented | product interaction reusable but GT corpus/context path incomplete | scoring/evidence path conceptually reusable | GT assets/path missing | `PARTIAL` |
| General Training Writing Task 1 | external letter task represented | dedicated learner feature/path incomplete | criterion path reusable but task-specific model incomplete | assets/evaluator calibration missing | `PARTIAL` |
| General Training Writing Task 2 | substantially shared | shared Writing path | shared criterion path | not implemented/validated | `MODELLED`, not `COVERED` |
| One Skill Retake preparation | construct shared | focused-skill journey can reuse current surfaces | normal evidence semantics reusable | release workflow not implemented | `DEFINED` |
| IELTS Life Skills | out of scope | none | none | none | `NOT_APPLICABLE` to current product target |

## Current top-level verdict

```text
Academic semantic model        STRONG / MODELLED
Academic product coverage      NOT YET COVERED
General Training               PARTIAL
Runtime support                NOT IMPLEMENTED
Validated band-outcome claim   NOT ESTABLISHED
```

Until this table has no blocking conditions for a scoped target and a TargetSupportDeclaration exists, user-facing copy must not say the product “fully supports” that target.

# Official task/question-family mapping

This table maps external IELTS families to existing product surfaces. It does not redefine the families; `02-IELTS-MODEL.md` owns external exam truth.

## Listening

| Official family | Product feature path | Practice path | Design condition |
|---|---|---|---|
| Multiple choice | `L-F01`, `L-F05` | focused distractor/choice items + `PM-L06` | `DEFINED` |
| Matching | `L-F01`, `L-F05` | focused relation/matching items + `PM-L06` | `DEFINED` |
| Plan/map/diagram labelling | `L-F06` | `PM-L05` | `SATISFIED` at design level |
| Form/note/table/flow-chart/summary completion | `L-F04` | `PM-L03` | `SATISFIED` at design level |
| Sentence completion | `L-F04` | `PM-L03` | `SATISFIED` at design level |
| Short-answer questions | `L-F04` | `PM-L03` | `SATISFIED` at design level |

A dedicated mode is not required for each task family when one interaction model truthfully supports several families.

## Academic Reading

| Official family | Product feature path | Practice path | Design condition |
|---|---|---|---|
| Multiple choice | `R-F01`, `R-F06` | `PM-R05` / `PM-R06` | `DEFINED` |
| True / False / Not Given | `R-F04` | `PM-R03` | `SATISFIED` at design level |
| Yes / No / Not Given | `R-F04` | `PM-R03` | `SATISFIED` at design level |
| Matching information | `R-F03`, `R-F07` | `PM-R02` / `PM-R06` | `DEFINED` |
| Matching headings | `R-F05` | `PM-R04` | `SATISFIED` at design level |
| Matching features | `R-F03`, `R-F06` | `PM-R02` / `PM-R05` | `DEFINED` |
| Matching sentence endings | `R-F06` | `PM-R05` | `DEFINED` |
| Sentence completion | `R-F03` | `PM-R02` / `PM-R06` | `DEFINED` |
| Summary/note/table/flow-chart completion | `R-F03`, `R-F06` | `PM-R02` / `PM-R05` | `DEFINED` |
| Diagram label completion | `R-F03` | `PM-R02` | `DEFINED` |
| Short-answer questions | `R-F03` | `PM-R02` / `PM-R06` | `DEFINED` |

Before Academic Reading becomes `COVERED`, each `DEFINED` family above requires real content/template generation, interaction support, answer/evidence logic, and transfer/readiness assets—not merely a table mapping.

## Writing

| Construct | Product feature path | Practice/evidence path | Design condition |
|---|---|---|---|
| Academic Task 1 visual information | `W-F01`, `W-F03`, `W-F07`, `W-F08`, `W-F09` | planning → draft → revision → timed independent | `DEFINED` |
| General Training Task 1 letter | current generic Writing surfaces plus missing dedicated letter-planning behavior | same Writing mechanisms, GT task-specific assets required | `PARTIAL` |
| Task 2 essay | `W-F01`, `W-F02`, `W-F06`, `W-F07`, `W-F08`, `W-F09` | component → full → timed independent | `DEFINED` |
| Writing criteria | criterion-level feedback/evaluator path | `AT-01` / `AT-05`, calibration required | `CALIBRATION_REQUIRED` for high-stakes support |

## Speaking

| Construct | Product feature path | Practice/evidence path | Design condition |
|---|---|---|---|
| Part 1 | `S-F01` | `PM-S03` | `DEFINED` |
| Part 2 | `S-F02` | `PM-S04` | `DEFINED` |
| Part 3 | `S-F03` | `PM-S05` | `DEFINED` |
| Pronunciation | `S-F04`, `S-F05`, `S-F07` | `PM-S01`, `PM-S02` | `DEFINED`; calibration needed for numeric claims |
| Whole Speaking construct | `S-F09` | `PM-S06` | `DEFINED`; evaluator/content/runtime missing |

# TargetProfile coverage

A product target is not one number.

A `TargetProfile` may contain:

```text
test_variant
purpose / receiving requirement when relevant
target_overall_band optional
minimum_listening_band optional
minimum_reading_band optional
minimum_writing_band optional
minimum_speaking_band optional
test_date optional
selected_skill_retake optional
```

The product may support a target only if every required TargetProfile condition has a covered path.

If a learner supplies only an overall target, the planner must not pretend that this uniquely determines the four skill targets. It must either:

1. collect required skill minima from the learner/receiving institution; or
2. use an explicitly labelled planning profile that keeps multiple valid score combinations possible.

Overall readiness is then evaluated from TargetProfile constraints and skill evidence—not from averaging internal mastery percentages.

# Route-to-target invariant

When a target is supported, the default product route is:

```text
TargetProfile
  ↓
provisional diagnostic
  ↓
material target gaps / unknowns
  ↓
Daily Plan
  ↓
required prerequisite + highest-value eligible action
  ↓
practice / review / re-evidence
  ↓
updated readiness conditions
  ↓
next plan
  ↺
until target conditions are supported or learner changes target
```

The learner may Swap, Skip, Shorten, or Change Skill among eligible alternatives, but:

- Required prerequisites remain blocking;
- skipped target requirements remain unsatisfied;
- learner preference never turns an uncovered/failed condition into support;
- the target remains stable until the learner explicitly changes it;
- repeated avoidance becomes a friction/preference signal, not ability evidence.

The app therefore provides a strong governed route without claiming it can force adherence or guarantee an external test result.

# No guaranteed-band promise

`SUPPORTED_FOR_PRODUCT` means the product has a complete, release-qualified path for the target. It does **not** mean every learner following the path is guaranteed to receive that IELTS score.

Allowed product language:

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
- provider/runtime/operations;
- rights/privacy/security;
- validation/research.

Only content gaps create content quantity demand.

# Support-declaration minimum

A future `TargetSupportDeclaration` must name:

- target scope and variants;
- supported bands/TargetProfiles;
- product/version boundary;
- feature/practice coverage;
- content/asset coverage;
- evaluator/calibration state;
- rights/privacy/security gates;
- third-party activation state;
- reliability/recovery state;
- known non-blocking validation backlog;
- revocation conditions.

Support is versioned and revocable when construct, provider, rights, reliability, cost, or calibration evidence changes.