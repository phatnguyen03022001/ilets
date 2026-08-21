STATUS: CANONICAL
OWNS: conceptual content-instance model, reference contracts between concrete content and canonical objects, Learning Unit, Practice/Assessment Item, Stimulus, ScaffoldingProfile, ExposureContext, Error/Remediation Pattern, Feedback Artifact, Attempt, Observation, EvidenceFact, and content coverage identity semantics
DEPENDS_ON: 03-SKILLS.md, 04-KNOWLEDGE.md, 06-CURRICULUM.md, 07-PRACTICE.md, 08-ASSESSMENT.md, 09-PROGRESSION.md
DOES_NOT_OWN: canonical skill/knowledge/band rules, curriculum sequence, learning-mechanism policy, assessment eligibility/sufficiency, learner-state transitions, persistence/API/database schemas

# 10 — Content Model

## Purpose

Define how concrete learning, tutoring, and measurement objects reference the canonical learning system without becoming a second source of truth.

The specification separates reusable canonical semantics from replaceable instances:

```text
Canonical definitions
  Skill Leaf
  Knowledge Object
  Curriculum Node
  Learning Mechanism
  Practice Type
  Assessment Type

Concrete/supporting instances
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

## Reference-first rule

Every concrete object identifies the canonical target(s) and policy objects that justify its existence. Concrete content may instantiate semantics; it may not redefine them.

When variant/context affects the task or inference, the object must preserve that scope explicitly. A generic `reading` or `writing-task-1` label is insufficient when Academic and General Training require different content or capability.

# `LearningUnit`

A Learning Unit is a delivery-neutral grouping of content around a coherent curriculum purpose.

Conceptual fields:

```text
id
curriculum_node_id
objective_refs
prerequisite_refs
test_variant_scope
content_sequence
practice_item_refs
assessment_item_refs
error_pattern_refs
remediation_pattern_refs
completion_intent
```

Completion is not mastery. Multiple Learning Units may instantiate the same Curriculum Node for different delivery modes, variants, or learner contexts.

# `Stimulus`

A Stimulus is source material the learner reads, hears, views, analyzes, or responds to.

Examples include passages, recordings/transcripts, charts/tables/maps, Writing prompts, Speaking questions/cue cards, model responses, worked examples, and sentence/paragraph material.

Conceptual fields:

```text
id
kind
test_variant_scope
exam_section_or_task_context
source_or_provenance
content
language_properties
difficulty_parameters
rights_or_usage_metadata
```

Variant/context examples:

```text
ACADEMIC_READING_PASSAGE
GT_READING_SECTION_1_EVERYDAY
GT_READING_SECTION_2_WORKPLACE
GT_READING_SECTION_3_GENERAL_INTEREST
ACADEMIC_WRITING_TASK_1_VISUAL
GT_WRITING_TASK_1_LETTER
WRITING_TASK_2_SHARED_OR_VARIANT_SCOPED
LISTENING_SHARED
SPEAKING_SHARED
```

These are conceptual context classes, not mandatory wire enum spellings until contracts are materialized.

A Stimulus may be reused across Practice and Assessment only when reuse does not invalidate independence, novelty, or the inference intended from the later attempt.

# `ScaffoldingProfile`

A ScaffoldingProfile records material support available during an item/attempt.

Conceptual dimensions may include:

```text
content_support
structural_support
lexical_support
response_support
hints
worked_example_access
feedback_timing
attempt/retry support
timing support
```

The profile exists because a correct response with material support has a different inference scope from independent performance.

Scaffolding is not inherently bad. It is a learning tool whose presence must remain visible when Assessment interprets an Observation.

# `ExposureContext`

ExposureContext records material prior exposure relevant to learning, novelty, retry, or transfer inference.

Conceptual fields:

```text
item_seen_before
stimulus_seen_before
prior_feedback_exposure
similarity_or_novelty_dimensions
prior_attempt_refs
context_variation
variant_context_variation
```

Exact-item novelty is not equivalent to transfer. The material dimensions depend on the claim.

# `PracticeItem`

A Practice Item is a concrete instance of a Practice Type and may declare the Learning Mechanisms it instantiates.

Conceptual fields:

```text
id
practice_type_id
learning_mechanism_refs
test_variant_scope
exam_section_or_task_context
target_skill_leaf_ids
target_knowledge_ids
curriculum_node_id
stimulus_refs
prompt_or_instruction
response_contract
difficulty_parameters
scaffolding_profile
exposure_context
error_pattern_refs
remediation_pattern_ref
feedback_contract
answer_or_model_reference
```

Invariants:

1. `practice_type_id` resolves to `07-PRACTICE.md`.
2. mechanism refs resolve to canonical `LM-*` where declared.
3. target IDs resolve to canonical Skill/Knowledge objects.
4. variant/context is compatible with every variant-specific target.
5. `W-TA-*` visual targets cannot be instantiated as GT Task 1 and `W-GT1-*` cannot be instantiated as Academic Task 1.
6. difficulty/scaffolding may vary without changing the target.
7. generated and authored items obey the same contract.
8. item instances are replaceable.

# `AssessmentItem`

An Assessment Item is a concrete measurement instance of an Assessment Type.

Conceptual fields:

```text
id
assessment_type_id
claim_scope
test_variant_scope
exam_section_or_task_context
target_skill_leaf_ids
target_knowledge_ids
target_band
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

- the item samples the capability it claims to measure;
- Band semantics are referenced from `05-BANDS.md`;
- variant/task/context matches the claim;
- Reading scoring reference uses the correct variant conversion where a Band inference is intended;
- conditions preserve timing, assistance, partial/full-task status, delivery mode, and other material context;
- independence/novelty metadata allows `08-ASSESSMENT.md` to judge the claim correctly;
- the item itself never decides mastery or certification.

# `ErrorPattern`

An Error Pattern is reusable tutoring knowledge about a recurring learner mistake or misconception.

Conceptual fields:

```text
id
target_skill_leaf_ids
target_knowledge_ids
variant_or_context_scope where material
pattern_description
applicability_context
example_refs
detection_hints
likely_causes where evidence supports them
priority when useful
evidence_or_provenance
locale_or_l1_scope
```

An Error Pattern is a hypothesis/supporting pattern, not proof that a learner has the error. Population/L1/variant-specific patterns must declare their scope.

# `RemediationPattern`

A Remediation Pattern is a reusable tutoring strategy for a diagnosed weak target or Error Pattern.

Conceptual fields:

```text
id
target_skill_leaf_ids
target_knowledge_ids
variant_or_context_scope where material
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

It may recommend mechanisms and Practice Types but cannot redefine prerequisites, Band thresholds, or mastery rules.

# `FeedbackArtifact`

Feedback is guidance derived from an Attempt/Observation.

Conceptual fields:

```text
attempt_ref
observation_refs
target_refs
variant/context scope where material
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

Feedback must distinguish observed performance from inferred cause. It should be actionable when formative and must not perform the learner's required cognitive operation on the learner's behalf.

Feedback is runtime/supporting output, not canonical learning truth.

# `Attempt`

An Attempt is a learner-instance event against a Practice Item or Assessment Item.

Conceptual fields:

```text
learner_ref
item_ref
response
started_at
completed_at
actual_delivery_mode
scaffolding_profile
exposure_context
evaluation_refs
```

An Attempt records what the learner did. It is not automatically evidence.

# `Observation`

An Observation is a normalized measurement result derived from an Attempt before evidence admission.

Conceptual fields:

```text
id
attempt_ref
assessment_type_ref
claim_candidate_refs
target_refs
test_variant_scope
exam_section_or_task_context
criterion_outcomes
raw_result
conditions
scaffolding_profile
exposure_context
scorer_or_evaluator_provenance
uncertainty
observed_at
```

Observations preserve history. A scoring/policy change may alter downstream interpretation without rewriting the original observation.

# `EvidenceFact`

An EvidenceFact is a claim-scoped admission of an Observation under `08-ASSESSMENT.md`.

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

One Observation may produce EvidenceFacts for multiple compatible claims or no EvidenceFact at all.

EvidenceFact is historical admissible evidence. It is not MasteryEstimate, ReadinessEvaluation, or Certification.

The legacy umbrella term `EvidenceRecord` should not be used for new design because it hides the Observation-versus-EvidenceFact boundary.

# Content coverage manifest semantics

When executable content is materialized, implementation must maintain a machine-checkable **content coverage manifest** or equivalent derived index. The manifest is an implementation artifact, not a new learning authority.

Each content capability entry should be able to answer:

```text
which canonical targets are instantiated
which IELTS variant/task/section contexts are represented
which Practice/Assessment Types are supported
which response interaction is implemented
which answer key/rubric/evaluator path applies
which difficulty/transfer classes exist
which rights/provenance state applies
whether exam-readiness independent assets exist
which product/release version activates the content
```

Coverage tooling may derive condition status from this manifest, but `design/08-coverage-and-support.md` remains the owner of what conditions must be satisfied.

A feature existing in UI is not evidence that content coverage exists.

# Content composition

A typical flow is:

```text
Curriculum Node + variant overlay
     ↓
Learning Unit
     ├── explanation / Stimuli
     ├── Practice Items
     ├── Error / Remediation Patterns
     └── Assessment Items
              ↓
            Attempt
              ↓
          Observation
              ↓
          EvidenceFact
              ↓
   Mastery / Readiness evaluation
              ↓
       learner-state decision
```

This is not a required UI screen flow or persistence schema.

# Difficulty model

Concrete content may express difficulty through linguistic complexity, information load, abstraction, distractor strength, response length, target integration, timing pressure, scaffold/cue availability, novelty, transfer distance, and variant/context demand.

Difficulty metadata supports selection and calibration. It does not define a Band.

# Content quality requirements

A concrete/supporting object is acceptable only when:

- canonical references resolve;
- variant/task/context references are internally compatible;
- it does not introduce contradictory teaching/scoring rules;
- task/variant conditions match purpose;
- answer key/rubric/model is valid where applicable;
- assessment contexts do not leak answers;
- provenance/rights are known where needed;
- difficulty and scaffolding are plausible for intended use;
- exposure/reuse does not invalidate the intended evidence claim;
- error/remediation claims declare population/context scope when non-universal;
- it can be replaced without changing canonical learning truth.

# Generated content boundary

AI generation is an instance-generation mechanism, not authority. Generated exercises, explanations, feedback, error hypotheses, and remediation suggestions must pass the same reference, variant/context, rights, quality, and evidence-context contract as authored content.

# No persistence contract

Names and fields here are conceptual contracts. They do not require one-to-one SQL tables, classes, JSON payloads, services, or language-specific types. Cross-language implementation uses explicit contracts under the repository governance rules in `CONSTITUTION.md`.
