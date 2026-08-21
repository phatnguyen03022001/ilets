STATUS: CANONICAL
OWNS: product coverage semantics, TargetCoverageSpecification, CoverageGap taxonomy, support-promotion gates, construct-to-product coverage mapping, and current product-support declarations
DEPENDS_ON: ../spec/00-PRODUCT.md, ../spec/02-IELTS-MODEL.md, ../spec/03-SKILLS.md, ../spec/05-BANDS.md, ../spec/06-CURRICULUM.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 01-skill-features.md, 02-practice-catalog.md, 03-media-youtube.md, 04-application-flows.md, 07-third-party-services.md
DOES_NOT_OWN: external IELTS truth, learning targets, feature behavior, provider selection, learner GapEvaluation, or empirical validation results themselves

# Coverage and Product Support

## Purpose

Define exactly when a scoped IELTS target is merely modelled, actually executable, release-supported, or empirically validated.

Document volume, feature count, Skill coverage, or an aggregate percentage cannot hide a missing required condition.

# Product-support status machine

A **scoped product target** uses only:

```text
MODELLED
  ↓
COVERED
  ↓
SUPPORTED_FOR_PRODUCT
  ↓
VALIDATED
```

## `MODELLED`

The relevant construct and product semantics are represented with enough precision to reason about the path and identify remaining gaps.

## `COVERED`

The complete executable path exists for the scoped target with no blocking CoverageGap:

```text
external requirement / official family
→ canonical capability/knowledge/context
→ curriculum route
→ intervention/practice
→ assessment/evidence/re-evidence
→ progression/next action
→ learner experience
→ content/assets
→ machine/runtime/provider path
```

## `SUPPORTED_FOR_PRODUCT`

A versioned TargetSupportDeclaration activates the target for a named product/release boundary and confirms every applicable release-critical gate.

## `VALIDATED`

Scoped empirical evidence supports the declared product outcome under named learner/product/content/evaluator/intervention versions and conditions.

Architecture coherence is never `VALIDATED` evidence.

# Condition status is separate

A **coverage condition** uses:

```text
UNKNOWN
DEFINED
PARTIAL
SATISFIED
BLOCKED
NOT_APPLICABLE
CALIBRATION_REQUIRED
```

Condition statuses never appear as target/product statuses.

An out-of-scope construct uses **OUT_OF_SCOPE** as a scope disposition, not a product-support state.

# CoverageGap

A CoverageGap is a product inability or unresolved condition, never learner weakness.

Canonical classes:

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

Each gap records:

```text
scoped target/condition
missing or failed condition
blocking consequence
dependencies
demand class
provenance/version
```

# Scope

## Complete standard IELTS learning construct

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ every in-scope official task/question family
+ material official subformats/presentation classes required for genuine content coverage
+ Band-3→9 learning interpretation
+ diagnostic → learning → practice → assessment → progression → exam-readiness
```

Release ordering may differ from construct completeness.

## Delivery modes

Delivery mode is a scope dimension, not another IELTS learning variant.

Current external delivery facts are owned by `../spec/02-IELTS-MODEL.md`. A product may be complete for a learning construct while still having a blocking readiness gap for a requested delivery mode.

## UKVI

UKVI Academic/GT reuses the corresponding language construct while adding external security/location/administrative requirements. A UKVI-purpose TargetProfile cannot imply support for an ineligible delivery route.

## One Skill Retake

One Skill Retake reuses an existing skill. It requires no fifth Skill ontology. Product support is scoped to focused preparation plus applicable eligibility/delivery conditions.

## IELTS Life Skills

IELTS Life Skills is **OUT_OF_SCOPE** for the current Band-3→9 construct.

# Required TargetCoverageSpecification conditions

For each scoped target, evaluate every applicable condition independently.

| Condition | Required meaning |
|---|---|
| `construct_model` | current external IELTS construct represented correctly |
| `official_family_coverage` | every applicable stable official family ID from `02-IELTS-MODEL.md` has an executable product/content/evidence path |
| `material_subformat_coverage` | required Presentation Classes inside a family are represented where one family label would otherwise hide a material gap |
| `skill_knowledge_model` | required canonical capabilities/knowledge exist |
| `band_threshold` | target Band/task/variant threshold is defined |
| `curriculum_route` | valid sequencing/variant route exists |
| `practice_intervention` | required learning mechanisms/modes can be executed |
| `feature_experience` | learner can perform the interaction and understand result/state |
| `content_assets` | sufficient valid stimuli/items/templates/generators cover required family/context/diversity |
| `assessment_policy` | applicable Assessment Type and executable versioned EvidenceRequirement exist |
| `evaluator_scoring` | deterministic scoring or calibrated productive evaluation exists |
| `progression_transition` | valid evidence can drive explainable state/next action |
| `variant_context` | Academic/GT task/section/context conditions are represented and sampled |
| `delivery_mode_readiness` | requested delivery-specific interaction/preparation exists when material |
| `machine_contracts` | exact interfaces exist for implemented cross-unit boundaries |
| `rights_privacy_security` | applicable source/data/consent/security requirements pass |
| `reliability_recovery` | lifecycle/idempotency/retry/failure/recovery pass |
| `accessibility_capture_quality` | access/capture failures cannot become false ability judgments |
| `cost_abuse_operations` | release operates within declared cost/rate/abuse limits |
| `observability_audit` | consequential decisions preserve version/provenance/reason reconstruction |
| `validation` | empirical outcome evidence exists when promoting to `VALIDATED` |

`NOT_APPLICABLE` requires an explicit reason for the scoped target.

# Official-family closure

Skill coverage and official-family coverage are independent dimensions.

Examples:

- `R-QT-04` may help teach several matching families, but content for `IELTS-R-QF-04` Matching Information cannot satisfy `IELTS-R-QF-06` Matching Features or `IELTS-R-QF-07` Matching Sentence Endings;
- one broad completion Skill cannot prove the existence of all required official completion-family content;
- a product may have strong Reading capability teaching while still having a blocking missing official-family interaction/content path.

For a complete standard target, every applicable stable `IELTS-*-*` family in `02-IELTS-MODEL.md` must resolve through:

```text
official family ID
→ feature interaction
→ Practice/Assessment role
→ concrete content family/subformat
→ answer/rubric/evaluator path
→ valid learner-state/readiness flow
```

# Material-subformat closure

Some official families contain materially different presentations. Coverage consumes the stable Presentation Classes from `../spec/10-CONTENT-MODEL.md` where defined.

Minimum current checks include:

- Listening `IELTS-L-QF-04`: form, note, table, flow-chart, summary completion;
- Reading `IELTS-R-QF-09`: summary, note, table, flow-chart completion;
- Academic Writing `IELTS-W-A-T1`: statistical graph/chart/table, diagram/process, and map/plan presentation classes.

A single content template cannot satisfy a multi-presentation condition merely by carrying the parent family ID.

# Content closure

Feature/UI existence is not content coverage.

When executable content exists, coverage tooling consumes the content manifest/equivalent index defined by `../spec/10-CONTENT-MODEL.md` and verifies at minimum:

```text
canonical target refs
stable official family refs
stable Content Context refs
material Presentation Class refs
variant/task/section context
practice/assessment role
interaction support
answer/rubric/evaluator route
transfer/novelty classes
rights/provenance
independent readiness assets
release activation
```

A `CONTENT_OR_ASSET` condition becomes `SATISFIED` only from executable/verified assets, not a design table.

# Evidence closure

Principles in `08-ASSESSMENT.md` are not by themselves production policy.

Before product support, each high-consequence claim resolves to a versioned executable EvidenceRequirement. Hidden model heuristics, unversioned cutoffs, or mechanical attempt-count rules are blockers.

Official-family **product coverage** does not imply that learner certification must mechanically test every family in every portfolio. Learner claim sufficiency remains an Assessment decision. This distinction prevents product completeness rules from leaking into learner mastery policy.

# Contract closure

A multi-runtime path cannot become `COVERED` while TypeScript, Go, and Python independently maintain handwritten interpretations of the same interface.

Every implemented cross-unit boundary needs one machine contract authority, validation/generated consumers where appropriate, and conformance verification.

# Delivery closure

When TargetProfile names a material delivery mode, exam-readiness coverage must expose that interaction honestly.

Delivery-mode preparation changes interaction conditions, not Skill/Band truth.

# Current design-state declaration — 2026-08-22

This is a documentation/design declaration, not production support.

| Scoped target | Model state | Remaining blocking classes | Product status |
|---|---|---|---|
| Academic Listening | strong construct/learning/design model incl. stable family IDs | official-family/subformat content, runtime/contracts, operations/validation | `MODELLED` |
| Academic Reading | strong construct/learning/design model incl. stable family IDs | official-family content, runtime/contracts, operations/validation | `MODELLED` |
| Academic Writing | strong construct/learning/design model incl. Task-1 presentation classes | productive evaluator calibration, diverse Task-1 content, runtime/contracts | `MODELLED` |
| Academic Speaking | strong construct/learning/design model incl. Part 1–3 family IDs | productive/audio calibration, content, runtime/contracts | `MODELLED` |
| GT Listening | shared strong model | official-family/subformat content, runtime/contracts, validation | `MODELLED` |
| GT Speaking | shared strong model | Part 1–3 content/evaluator/runtime verification | `MODELLED` |
| GT Reading | shared skills + explicit GT context/scoring + official families | GT family/context assets, runtime/contracts, validation | `MODELLED` |
| GT Writing Task 1 | dedicated GT capability + Band/Curriculum/product semantics | letter asset diversity, productive calibration, runtime/contracts | `MODELLED` |
| GT Writing Task 2 | shared Writing construct | content, evaluator calibration, runtime/contracts | `MODELLED` |
| One Skill Retake preparation | existing-skill reuse + target-flow semantics | release/admin/delivery workflow, runtime, validation | `MODELLED` |
| Test-centre computer exam-readiness | external delivery model represented | executable UX/content/runtime verification | `MODELLED` |
| Writing-on-Paper readiness | external delivery overlay represented | handwriting-specific product assets/verification by supported market | `MODELLED` |
| IELTS Online Academic readiness | external delivery overlay represented | remote-mode product support, purpose/acceptance handling, verification | `MODELLED` |

Scope disposition:

```text
IELTS Life Skills = OUT_OF_SCOPE
```

Top-level current truth:

```text
Academic learning/design semantics         MODELLED
General Training learning/design semantics MODELLED
Official family identities                 MODELLED
Material presentation identities           MODELLED where required
Delivery-mode semantics                    MODELLED
Academic product execution                 NOT YET COVERED
General Training product execution         NOT YET COVERED
Runtime implementation                     NOT IMPLEMENTED
Validated target-band outcome              NOT ESTABLISHED
```

No user-facing copy may claim full support until all scoped blocking conditions close and a TargetSupportDeclaration activates that exact scope.

# Construct → product mapping

This mapping proves design intent only. It does not satisfy runtime/content gates.

## Listening

| Official family ID | Product feature/practice path |
|---|---|
| `IELTS-L-QF-01` Multiple choice | `L-F01`, `L-F05`, `PM-L04`, `PM-L06` |
| `IELTS-L-QF-02` Matching | `L-F01`, `PM-L06` plus focused `PT-13` items |
| `IELTS-L-QF-03` Plan/map/diagram labelling | `L-F06`, `PM-L05` |
| `IELTS-L-QF-04` Form/note/table/flow-chart/summary completion | `L-F04`, `PM-L03`; all required `PRES-L-QF04-*` classes remain independently checkable |
| `IELTS-L-QF-05` Sentence completion | `L-F04`, `PM-L03` |
| `IELTS-L-QF-06` Short answer | `L-F04`, `PM-L03` |

## Reading

| Official family ID | Product feature/practice path |
|---|---|
| `IELTS-R-QF-01` Multiple choice | `R-F01`, `R-F06`, `PM-R05` / `PM-R06` |
| `IELTS-R-QF-02` True/False/Not Given | `R-F04`, `PM-R03` |
| `IELTS-R-QF-03` Yes/No/Not Given | `R-F04`, `PM-R03` |
| `IELTS-R-QF-04` Matching information | `R-F03`, `R-F07`, focused Reading item + `PM-R06` |
| `IELTS-R-QF-05` Matching headings | `R-F05`, `PM-R04` |
| `IELTS-R-QF-06` Matching features | `R-F03`, `R-F06`, focused Reading item |
| `IELTS-R-QF-07` Matching sentence endings | `R-F06`, focused Reading item |
| `IELTS-R-QF-08` Sentence completion | `R-F03`, focused completion item |
| `IELTS-R-QF-09` Summary/note/table/flow-chart completion | `R-F03`, `R-F06`; required `PRES-R-QF09-*` classes remain independently checkable |
| `IELTS-R-QF-10` Diagram label completion | `R-F03`, focused Reading item |
| `IELTS-R-QF-11` Short answer | `R-F03`, focused Reading item / `PM-R06` |

Variant coverage additionally requires applicable Academic or GT Content Context distribution and the correct scoring/evidence policy.

## Writing

| Official task family ID | Product path | Canonical target/evidence path |
|---|---|---|
| `IELTS-W-A-T1` | `W-F01`, `W-F03`, `W-F07`, `W-F08`, `W-F09` | `W-TA-*`; required `PRES-W-A-T1-*` content diversity; plan → draft → revision → timed independent |
| `IELTS-W-GT-T1` | variant-aware Writing features | `W-GT1-*`; recipient/purpose/register/required-point coverage → draft → timed independent |
| `IELTS-W-T2` | `W-F01`, `W-F02`, `W-F06`, `W-F07`, `W-F08`, `W-F09` | component → full → timed independent |
| Writing criteria | criterion feedback/evaluator path | `AT-01` / `AT-05`; calibration required for high-consequence support |

## Speaking

| Official family ID | Product path |
|---|---|
| `IELTS-S-P1` | `S-F01` / `PM-S03` |
| `IELTS-S-P2` | `S-F02` / `PM-S04` |
| `IELTS-S-P3` | `S-F03` / `PM-S05` |
| whole Speaking construct | `S-F09` / `PM-S06` with Part 1–3 coverage |
| pronunciation learning surface | `S-F04`, `S-F05`, `S-F07` / `PM-S01`, `PM-S02` |

# TargetProfile coverage invariant

A product target is a constraint set, not one Band number.

Support requires a covered path for every required TargetProfile condition, including variant, real per-skill minima, external purpose constraints, and delivery mode where material.

If only an overall target exists, the planner either obtains real minima or uses an explicitly labelled planning profile that preserves multiple valid score combinations.

# Route invariant

For a supported target:

```text
TargetProfile
  ↓
provisional diagnostic
  ↓
explicit unresolved conditions
  ↓
Planner hard eligibility
  ↓
ranked valid next action
  ↓
practice / review / re-evidence
  ↓
updated claim states
  ↺
```

Ranking cannot erase prerequisites, family/context/delivery compatibility, evidence truth, or CoverageGaps.

# Promise boundary

`SUPPORTED_FOR_PRODUCT` means the product has a complete release-qualified path for the declared scope. It does **not** guarantee an external IELTS result.

Allowed:

- “Your current evidence supports Band N under this profile.”
- “These are the remaining blockers to your target.”
- “This target/delivery scope is supported by the current product version.”

Forbidden without suitable empirical basis:

- “Follow this plan and you will get Band 7.”
- “100% guaranteed target Band.”
- “This exercise improved your IELTS score by 0.5 Band.”

# Gap → implementation demand

Every blocking CoverageGap maps to one demand class:

- spec/model;
- feature/interaction;
- content/assets/generator;
- evaluator/calibration;
- learner flow/transition;
- contract/integration;
- runtime/provider/operations;
- rights/privacy/security;
- validation/research.

Only actual content gaps create content-quantity demand.

# TargetSupportDeclaration minimum

A release declaration names at least:

- exact target/variant scope;
- supported Band/TargetProfile conditions;
- supported delivery modes/purpose constraints where material;
- product/release version;
- official-family coverage manifest/result;
- material-subformat coverage manifest/result;
- feature/practice coverage;
- content-manifest version;
- EvidenceRequirement/policy versions;
- evaluator/calibration state;
- machine-contract versions;
- rights/privacy/security state;
- third-party activation state;
- reliability/recovery state;
- accessibility/capture-quality gate;
- cost/abuse/operations gate;
- observability/audit gate;
- known non-blocking validation backlog;
- revocation conditions.

Support is versioned and revocable when construct, delivery, provider, rights, reliability, cost, calibration, content coverage, or validation evidence materially changes.