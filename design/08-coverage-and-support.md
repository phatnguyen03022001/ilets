STATUS: CANONICAL
OWNS: product coverage semantics, TargetCoverageSpecification, CoverageGap taxonomy, support-promotion gates, construct-to-product coverage mapping, and current product-support declarations
DEPENDS_ON: ../spec/00-PRODUCT.md, ../spec/02-IELTS-MODEL.md, ../spec/03-SKILLS.md, ../spec/05-BANDS.md, ../spec/06-CURRICULUM.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 01-skill-features.md, 02-practice-catalog.md, 03-media-youtube.md, 04-application-flows.md, 07-third-party-services.md
DOES_NOT_OWN: external IELTS truth, learning targets, feature behavior, provider selection, learner GapEvaluation, or empirical validation results themselves

# Coverage and Product Support

## Purpose

Define exactly when a scoped IELTS target is merely modelled, actually executable, release-supported, or empirically validated.

Document volume, feature count, or an aggregate percentage cannot hide a missing required condition.

# Product-support status machine

A **scoped product target** uses only these statuses:

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
external requirement
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

# Condition status is a different enum

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

Condition statuses must never be written into a target/product-status field.

Examples:

- `DEFINED` may describe an individual content/evidence/feature condition while the target remains `MODELLED`;
- `NOT_APPLICABLE` is a condition result with an explicit reason, not a product-support status;
- an out-of-scope construct is labelled **OUT_OF_SCOPE** as a scope disposition, not inserted into the product-status state machine.

# CoverageGap

A CoverageGap is a product inability or unresolved condition, never a learner weakness.

Canonical gap classes:

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

The intended complete learning construct is:

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ official task/question families
+ Band-3→9 learning interpretation
+ diagnostic → learning → practice → assessment → progression → exam-readiness
```

Release ordering may differ from construct completeness.

## Delivery modes

Delivery mode is a scope dimension, not another IELTS learning variant.

The external model currently distinguishes, where applicable:

- test-centre computer;
- test-centre computer + Writing on Paper;
- IELTS Online Academic.

A product may be complete for an Academic/GT learning construct while still having a blocking delivery-readiness gap for a specific requested mode. TargetSupportDeclaration must therefore name supported delivery modes when they are material to the target.

## UKVI

UKVI Academic/GT reuses the corresponding language construct while adding external security/location/administrative requirements. Product support for a UKVI-purpose TargetProfile must not imply support for an ineligible delivery mode.

## One Skill Retake

One Skill Retake reuses an existing skill. It requires no fifth Skill ontology. Product support is scoped to focused preparation plus applicable eligibility/delivery guidance.

## IELTS Life Skills

IELTS Life Skills is **OUT_OF_SCOPE** for the current Band-3→9 product-learning construct. `OUT_OF_SCOPE` is a scope disposition, not a product-support status.

# Required TargetCoverageSpecification conditions

For each scoped target, evaluate every applicable condition independently.

| Condition | Required meaning |
|---|---|
| `construct_model` | current external IELTS construct represented correctly |
| `skill_knowledge_model` | required canonical capabilities/knowledge exist |
| `band_threshold` | target Band/task/variant threshold is defined |
| `curriculum_route` | valid sequencing/variant route exists |
| `practice_intervention` | required learning mechanisms/modes can be executed |
| `feature_experience` | learner can perform the interaction and understand result/state |
| `content_assets` | sufficient valid stimuli/items/templates/generators cover required diversity |
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

# Content closure

Feature/UI existence is not content coverage.

When executable content exists, coverage tooling consumes the content manifest/equivalent index defined by `../spec/10-CONTENT-MODEL.md` and verifies at minimum:

```text
canonical target refs
variant/task/section contexts
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

Before product support, each high-consequence claim resolves to a versioned executable EvidenceRequirement. Hidden LLM heuristics, unversioned cutoffs, or mechanical attempt-count rules are blockers.

# Contract closure

A multi-runtime path cannot become `COVERED` while TypeScript, Go, and Python independently maintain handwritten interpretations of the same interface.

Every implemented cross-unit boundary needs one machine contract authority, validation/generated consumers where appropriate, and conformance verification.

# Delivery closure

When TargetProfile names a material delivery mode, exam-readiness coverage must expose that interaction honestly.

Examples:

- standard test-centre computer → computer navigation/typing/listening/reading conditions where relevant;
- Writing on Paper → handwritten Writing interaction/readiness while L/R remain computer-delivered;
- IELTS Online Academic → remote-platform readiness where product support claims it, subject to the external purpose/acceptance boundary.

Delivery-mode preparation never changes Skill/Band truth.

# Current design-state declaration — 2026-08-22

This is a documentation/design declaration, not production support.

| Scoped target | Model state | Remaining blocking classes | Product status |
|---|---|---|---|
| Academic Listening | strong construct/learning/design model | content, runtime/contracts, operational validation | `MODELLED` |
| Academic Reading | strong construct/learning/design model | content, runtime/contracts, operational validation | `MODELLED` |
| Academic Writing | strong construct/learning/design model | productive evaluator calibration, content, runtime/contracts | `MODELLED` |
| Academic Speaking | strong construct/learning/design model | productive/audio evaluator calibration, content, runtime/contracts | `MODELLED` |
| GT Listening | shared strong model | content, runtime/contracts, validation | `MODELLED` |
| GT Speaking | shared strong model | content, runtime/contracts, validation | `MODELLED` |
| GT Reading | shared skills + explicit GT section/context/scoring semantics | GT content/assets, runtime/contracts, validation | `MODELLED` |
| GT Writing Task 1 | dedicated GT capability + Band/Curriculum/product semantics | assets, productive evaluator calibration, runtime/contracts | `MODELLED` |
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
Academic learning/design semantics        MODELLED
General Training learning/design semantics MODELLED
Delivery-mode semantics                   MODELLED
Academic product execution                NOT YET COVERED
General Training product execution         NOT YET COVERED
Runtime implementation                     NOT IMPLEMENTED
Validated target-band outcome              NOT ESTABLISHED
```

No user-facing copy may claim full support until all scoped blocking conditions close and a TargetSupportDeclaration activates that exact scope.

# Construct → product mapping

This mapping proves design intent only. It does not satisfy runtime/content gates.

## Listening

| External family | Feature/practice path |
|---|---|
| Multiple choice | `L-F01`, `L-F05`, `PM-L06` |
| Matching | `L-F01`, `L-F05`, `PM-L06` |
| Plan/map/diagram labelling | `L-F06`, `PM-L05` |
| Form/note/table/flow-chart completion | `L-F04`, `PM-L03` |
| Sentence completion | `L-F04`, `PM-L03` |
| Short answer | `L-F04`, `PM-L03` |

## Reading

| External family | Feature/practice path |
|---|---|
| Multiple choice | `R-F01`, `R-F06`, `PM-R05` / `PM-R06` |
| True/False/Not Given | `R-F04`, `PM-R03` |
| Yes/No/Not Given | `R-F04`, `PM-R03` |
| Matching information | `R-F03`, `R-F07`, `PM-R02` / `PM-R06` |
| Matching headings | `R-F05`, `PM-R04` |
| Matching features/endings | `R-F03`, `R-F06`, `PM-R02` / `PM-R05` |
| Completion families | `R-F03`, `R-F06`, `PM-R02` / `PM-R05` |
| Diagram label completion | `R-F03`, `PM-R02` |
| Short answer | `R-F03`, `PM-R02` / `PM-R06` |

Variant coverage additionally requires Academic corpus/readiness assets or the GT Section-1/2/3 context distribution plus correct variant scoring/evidence policy.

## Writing

| Construct | Product path | Canonical target/evidence path |
|---|---|---|
| Academic Task 1 visual | `W-F01`, `W-F03`, `W-F07`, `W-F08`, `W-F09` | `W-TA-*`; plan → draft → revision → timed independent |
| GT Task 1 letter | variant-aware Writing features | `W-GT1-*`; recipient/purpose/register/bullet coverage → draft → timed independent |
| Task 2 essay | `W-F01`, `W-F02`, `W-F06`, `W-F07`, `W-F08`, `W-F09` | component → full → timed independent |
| Writing criteria | criterion feedback/evaluator path | `AT-01` / `AT-05`; calibration required for high-consequence support |

## Speaking

| Construct | Product path |
|---|---|
| Part 1 | `S-F01` / `PM-S03` |
| Part 2 | `S-F02` / `PM-S04` |
| Part 3 | `S-F03` / `PM-S05` |
| Pronunciation | `S-F04`, `S-F05`, `S-F07` / `PM-S01`, `PM-S02` |
| Whole Speaking | `S-F09` / `PM-S06` |

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

Ranking cannot erase prerequisites, variant/delivery compatibility, evidence truth, or CoverageGaps.

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

Support is versioned and revocable when construct, delivery, provider, rights, reliability, cost, calibration, or validation evidence materially changes.