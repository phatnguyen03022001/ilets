STATUS: CANONICAL
OWNS: conceptual content-instance model, stable content-context identity, stable material presentation identity, reference contracts between concrete content and canonical objects/external task families, Learning Unit, Stimulus, Practice/Assessment Item, ScaffoldingProfile, ExposureContext, Error/Remediation Pattern, Feedback Artifact, Attempt, Observation, EvidenceFact representation, and content-coverage identity semantics
DEPENDS_ON: 02-IELTS-MODEL.md, 03-SKILLS.md, 04-KNOWLEDGE.md, 06-CURRICULUM.md, 07-PRACTICE.md, 08-ASSESSMENT.md, 09-PROGRESSION.md
DOES_NOT_OWN: external IELTS task-family definitions, Skill/Knowledge/Band truth, curriculum sequence, learning-mechanism policy, Assessment sufficiency, learner-state transitions, exact wire/storage schema, product activity-purpose/evidence-candidacy policy, or product coverage status

# 10 — Content Model

## Purpose

Define how concrete learning, tutoring, and measurement objects reference canonical learning/exam semantics without becoming another source of truth.

```text
Canonical definitions
  external IELTS task/question family
  Skill Leaf
  Knowledge Object
  Curriculum Node
  Learning Mechanism
  Practice Type
  Assessment Type
  Content Context
  Content Presentation Class

Concrete/runtime instances
  Learning Unit
  Stimulus
  ScaffoldingProfile
  ExposureContext
  Practice Item
  Assessment Item
  Error Pattern
  Remediation Pattern
  Feedback Artifact
  Attempt
  Observation
  EvidenceFact
```

# Reference-first rule

Every concrete object references the canonical targets/context/policies that justify its existence.

When variant, official task/question family, task/section context, material presentation, or delivery condition changes task meaning or inference, that scope remains explicit. Generic labels such as `reading`, `completion`, or `writing-task-1` are insufficient for coverage claims.

Conditional identity dimensions are **applicability-aware**. A field that is semantically inapplicable may be omitted/represented as not applicable by the eventual machine contract; a material field may not be omitted merely for convenience. Implementation must not fabricate a Band, Content Context, official family, Presentation Class, delivery mode, or Curriculum Node solely to satisfy a schema.

# Stable Content Context registry

Content Context identifies **where an IELTS construct is instantiated**, not a new learning skill and not a universal label for every content object.

| ID | Meaning | Variant |
|---|---|---|
| `CTX-LISTENING-SHARED` | shared IELTS Listening construct/context | shared |
| `CTX-READING-ACADEMIC` | Academic Reading passage/set context | Academic |
| `CTX-READING-GT-S1-EVERYDAY` | GT Reading Section 1 everyday context | General Training |
| `CTX-READING-GT-S2-WORKPLACE` | GT Reading Section 2 workplace context | General Training |
| `CTX-READING-GT-S3-GENERAL-INTEREST` | GT Reading Section 3 longer general-interest context | General Training |
| `CTX-WRITING-ACADEMIC-T1-VISUAL` | Academic Writing Task 1 visual-information context | Academic |
| `CTX-WRITING-GT-T1-LETTER` | GT Writing Task 1 letter context | General Training |
| `CTX-WRITING-T2` | Writing Task 2 construct; concrete prompt retains variant where material | shared/core |
| `CTX-SPEAKING-SHARED` | shared IELTS Speaking construct | shared |

Rules:

1. these IDs remain stable across content/manifests/runtime contracts;
2. exact JSON representation is a machine-contract concern, but IDs survive unchanged;
3. new context IDs require a materially distinct external/content inference context, not a topic, difficulty, Band, delivery mode, or screen;
4. delivery mode is orthogonal and does not multiply Content Context IDs;
5. a context-neutral Knowledge/foundation item may legitimately have no Content Context when IELTS task/section context does not affect its meaning or inference;
6. do not create `CTX-GENERIC`, `CTX-KNOWLEDGE`, or another filler context merely to satisfy storage/API shape;
7. once official-family/variant/task/section context changes validity, coverage, scoring, or inference, the appropriate Content Context becomes required.

# Official family reference

Official task/question-family identity is owned by `02-IELTS-MODEL.md` through stable `IELTS-*-*` IDs.

Concrete content records those IDs separately from Skill targets when material because:

- one Skill Leaf may serve multiple official families;
- one official family may require multiple capabilities;
- grouping several families under one Skill Leaf must not allow missing content for one family to disappear from coverage.

A context-neutral Knowledge/foundation item does not invent an official family merely because the product later packages it inside IELTS preparation.

# Stable Content Presentation Class registry

A Presentation Class identifies a materially different official subformat/stimulus shape inside an external family. It exists only when one family label would otherwise hide a meaningful content/interaction coverage gap.

## Listening completion family — `IELTS-L-QF-04`

| ID | Presentation |
|---|---|
| `PRES-L-QF04-FORM` | form completion |
| `PRES-L-QF04-NOTE` | note completion |
| `PRES-L-QF04-TABLE` | table completion |
| `PRES-L-QF04-FLOW-CHART` | flow-chart completion |
| `PRES-L-QF04-SUMMARY` | summary completion |

## Reading completion family — `IELTS-R-QF-09`

| ID | Presentation |
|---|---|
| `PRES-R-QF09-SUMMARY` | summary completion |
| `PRES-R-QF09-NOTE` | note completion |
| `PRES-R-QF09-TABLE` | table completion |
| `PRES-R-QF09-FLOW-CHART` | flow-chart completion |

## Academic Writing Task 1 — `IELTS-W-A-T1`

| ID | Presentation |
|---|---|
| `PRES-W-A-T1-GRAPH-CHART-TABLE` | graph/chart/table/statistical visual, including combined statistical displays |
| `PRES-W-A-T1-DIAGRAM-PROCESS` | diagram of a process, object, device, event, or comparable process representation |
| `PRES-W-A-T1-MAP-PLAN` | map/plan/spatial-change representation |

Rules:

1. Presentation Class is a content-coverage identity, not a Skill, Practice Type, or scored task;
2. a Stimulus may reference multiple classes when a task genuinely combines presentation types;
3. topic, accent, vocabulary theme, difficulty, and Band do not become presentation IDs merely for catalog convenience;
4. new presentation IDs require a material interaction/content-coverage distinction justified by the external construct or evidence policy.

# Delivery scope

Where exam-readiness behavior depends materially on external delivery, concrete items/runs may declare a delivery scope/reference resolved from `02-IELTS-MODEL.md` and downstream target context.

Delivery scope changes interaction/conditions. It does not create a second Content Context, external family, or Presentation Class.

# `LearningUnit`

Delivery-neutral grouping around a coherent curriculum purpose.

Conceptual fields:

```text
id
curriculum_node_id
objective_refs
prerequisite_refs
test_variant_scope
content_context_refs where material
external_task_family_refs where material
required_presentation_class_refs where material
delivery_mode_scope optional
content_sequence
practice_item_refs
assessment_item_refs
error_pattern_refs
remediation_pattern_refs
completion_intent
```

Completion is not mastery. Multiple units may instantiate one Curriculum Node for different variants, family/subformat coverage, delivery preparation, or learner contexts.

# `Stimulus`

Material the learner reads, hears, views, analyses, or responds to.

Conceptual fields:

```text
id
kind
test_variant_scope
content_context_id where material
external_task_family_refs where material
presentation_class_refs where material
source_or_provenance
content
language_properties
difficulty_parameters
rights_or_usage_metadata
```

Examples include passages, recordings/transcripts, charts/tables/maps, Writing prompts, Speaking questions/cue cards, model responses, worked examples, and sentence/paragraph material.

A Stimulus may be reused across Practice/Assessment only when reuse does not invalidate independence, novelty, or later inference.

# `ScaffoldingProfile`

Records material support available during learning/performance.

Conceptual dimensions:

```text
content_support
structural_support
lexical_support
response_support
hints
worked_example_access
feedback_timing
attempt_or_retry_support
timing_support
```

Support is not inherently invalid; its presence must remain visible because it changes inference scope.

# `ExposureContext`

Records material prior exposure and transfer/novelty context.

Conceptual fields:

```text
item_seen_before
stimulus_seen_before
prior_feedback_exposure
similarity_or_novelty_dimensions
prior_attempt_refs
context_variation
content_context_variation
external_family_variation where material
presentation_variation where material
```

Exact-item novelty is not equivalent to meaningful transfer.

# `PracticeItem`

Concrete instance of a Practice Type.

Conceptual fields:

```text
id
practice_type_id
learning_mechanism_refs where applicable
test_variant_scope
content_context_id where material
external_task_family_refs where material
presentation_class_refs where material
delivery_mode_scope optional
target_skill_leaf_ids
target_knowledge_ids
curriculum_node_id where the item is specifically node-bound
stimulus_refs
prompt_or_instruction
response_contract
difficulty_parameters
scaffolding_profile
exposure_context
error_pattern_refs
remediation_pattern_ref
feedback_contract
answer_or_model_reference where applicable
```

Invariants:

1. `practice_type_id` resolves to `07-PRACTICE.md`;
2. Learning Mechanism refs resolve where declared;
3. Skill/Knowledge refs resolve canonically;
4. official family refs resolve to `02-IELTS-MODEL.md` when the item instantiates an official family;
5. Content Context is present and compatible whenever variant/task/section context changes validity, coverage, or inference;
6. context-neutral Knowledge/foundation items may omit Content Context rather than inventing one;
7. required Presentation Class is represented where the family has material subformats;
8. `W-TA-*` visual targets cannot instantiate GT Task 1;
9. `W-GT1-*` cannot instantiate Academic Task 1;
10. family identity cannot be inferred only from a broad Skill Leaf when several official families share that leaf;
11. delivery scope is recorded when interaction/readiness inference depends on it;
12. Curriculum Node binding is explicit when material but is not fabricated for reusable/direct-browse items whose identity is target/type-based;
13. difficulty/scaffold may vary without changing the target;
14. authored and generated items obey the same contract;
15. item instances remain replaceable.

# `AssessmentItem`

Concrete measurement instance of an Assessment Type.

Conceptual fields:

```text
id
assessment_type_id
claim_scope
test_variant_scope
content_context_id where material
external_task_family_refs where material
presentation_class_refs where material
delivery_mode_scope optional
target_skill_leaf_ids
target_knowledge_ids
target_band when the claim is Band-scoped
stimulus_refs
prompt_or_task
response_contract
conditions
scaffolding_profile
exposure_context
scoring_reference
answer_key_or_rubric_reference
independence_group
```

Invariants:

- samples the claimed capability;
- `target_band` is required only for a Band-scoped claim; Knowledge probes, leaf-level diagnostics, and other non-Band claims must not invent a Band merely to satisfy schema shape;
- references Band semantics rather than redefining them when Band-scoped;
- official family/context/variant matches the claim when material;
- context-neutral Knowledge/foundation assessment may omit Content Context;
- material Presentation Class is preserved where relevant;
- Reading Band inference uses applicable variant scoring policy;
- timing, assistance, partial/full-task state, delivery/input mode, capture quality, and other material conditions remain visible;
- independence/novelty metadata supports correct Assessment inference;
- the item never decides mastery/certification.

# `ErrorPattern`

Reusable tutoring hypothesis about a recurring error/misconception.

Conceptual fields:

```text
id
target_skill_leaf_ids
target_knowledge_ids
content_context_refs when material
external_task_family_refs when material
presentation_class_refs when material
pattern_description
applicability_context
example_refs
detection_hints
likely_causes where evidence supports them
priority where useful
evidence_or_provenance
locale_or_l1_scope
```

It is not proof the learner has the error. Population/L1/context-specific patterns declare scope.

# `RemediationPattern`

Reusable tutoring strategy for a diagnosed target/ErrorPattern.

Conceptual fields:

```text
id
target_skill_leaf_ids
target_knowledge_ids
content_context_refs when material
external_task_family_refs when material
presentation_class_refs when material
error_pattern_refs
learning_mechanism_refs
practice_type_refs
explanation_or_stimulus_refs
scaffold_strategy
recommended_sequence
success_check
scope
evidence_or_provenance
```

It may recommend mechanisms/types but cannot redefine prerequisites, Band thresholds, or mastery rules.

# `FeedbackArtifact`

Runtime guidance derived from Attempt/Observation.

Conceptual fields:

```text
attempt_ref
observation_refs
target_refs
content_context_ref when material
external_task_family_refs when material
presentation_class_refs when material
observed_performance
gap_or_error
matched_error_pattern_refs
feedback_message
recommended_action_intent
recommended_mechanism_refs
recommended_practice_type_refs
remediation_pattern_refs
uncertainty when material
```

Feedback distinguishes observation from inferred cause and must not perform the learner’s required cognitive operation for them.

# `Attempt`

Learner-instance event against a Practice/Assessment Item.

Conceptual fields:

```text
learner_ref
item_ref
response
started_at
completed_at
actual_delivery_mode
actual_input_or_capture_conditions
scaffolding_profile
exposure_context
evaluation_refs
```

An Attempt records what happened; it is not automatically evidence.

# `Observation`

Normalized measurement result before evidence admission.

Conceptual fields:

```text
id
attempt_ref
assessment_type_ref
claim_candidate_refs
target_refs
test_variant_scope
content_context_id where material
external_task_family_refs when material
presentation_class_refs when material
actual_delivery_mode where material
criterion_outcomes
raw_result
conditions
scaffolding_profile
exposure_context
scorer_or_evaluator_provenance
uncertainty
observed_at
```

Observations preserve history. New scoring/policy may change later interpretation without rewriting the original measurement.

# `EvidenceFact`

Claim-scoped admission of an Observation under Assessment policy.

Conceptual fields:

```text
id
observation_ref
claim_scope
eligibility_status
eligibility_reason
inference_scope
policy_version
admitted_at
```

One Observation may yield multiple compatible EvidenceFacts or none.

EvidenceFact is not MasteryEstimate, ReadinessEvaluation, or Certification.

Do not reintroduce an umbrella `EvidenceRecord` that hides Observation vs EvidenceFact.

# Content coverage manifest

When executable content exists, maintain a machine-checkable content manifest/equivalent derived index.

Each entry must support answers to:

```text
canonical target refs
official task/question-family refs + applicability
Content Context ref + applicability
material Presentation Class refs + applicability
test variant
supported delivery scope where material
Practice/Assessment Type support
implemented response interaction
answer-key/rubric/evaluator route
difficulty/transfer classes
rights/provenance state
independent readiness asset state
product/release activation
```

For conditional identity dimensions, the manifest must let coverage tooling distinguish at least:

```text
applicable + present
explicitly not applicable
required but missing/unresolved
```

Exact encoding belongs to machine contracts/implementation. Absence alone must not be interpreted as `NOT_APPLICABLE`.

Coverage tooling must query **official-family and required Presentation-Class coverage independently of Skill coverage**.

Examples:

- Matching Information cannot satisfy Matching Features merely because both share a broader Reading capability;
- one form-completion template cannot prove all required `IELTS-L-QF-04` presentation coverage;
- only graph/chart Academic Task-1 assets cannot prove visual-task coverage when required process/map presentation classes are absent;
- a context-neutral grammar item does not need a fake Content Context and cannot be counted toward variant/family coverage that does require one.

The manifest is implementation truth about available content, not new learning authority.

`design/08-coverage-and-support.md` owns which conditions must pass before coverage/support promotion.

A UI feature existing is not proof of content coverage.

# Composition

```text
canonical Skill/Knowledge targets
+ Curriculum Node when specifically bound
+ external family/context/presentation when material
      ↓
Learning Unit / concrete content binding
  ├─ Stimuli
  ├─ Practice Items
  ├─ Error/Remediation Patterns
  └─ Assessment Items
          ↓
        Attempt
          ↓
      Observation
          ↓
      EvidenceFact
          ↓
 Readiness / learner-state decision
```

This is conceptual composition, not a required UI flow or database table graph.

# Difficulty

Concrete difficulty may vary by linguistic complexity, information load, abstraction, distractor strength, response length, integration, timing pressure, scaffold/cues, novelty, transfer distance, and context demand.

Difficulty metadata does not define a Band.

# Content quality requirements

A concrete object is acceptable only when:

- canonical refs resolve;
- external task-family refs resolve when the item represents an official family;
- Content Context/variant refs are compatible when context is material;
- context-neutral omission is explicit/valid rather than accidental missing metadata;
- Presentation Class is correct when material;
- delivery scope is compatible when material;
- it introduces no contradictory teaching/scoring rule;
- answer key/rubric/model is valid where applicable;
- assessment context does not leak answers;
- provenance/rights are known where needed;
- difficulty/scaffolding fit intended use;
- reuse/exposure does not invalidate inference;
- non-universal remediation/error claims declare scope;
- it can be replaced without changing canonical learning truth.

# Generated-content boundary

AI generation is an instance-generation mechanism, not authority. Generated items/explanations/feedback/error hypotheses/remediation must pass the same target, official-family, context, presentation, delivery, rights, quality, and evidence contracts as authored content.

# No persistence/wire contract

Names/fields here define conceptual semantics. They do not force one-to-one SQL tables/classes/JSON payloads.

Machine boundaries materialize exact shapes under repository contract governance while preserving stable canonical IDs and the applicability distinctions defined here.