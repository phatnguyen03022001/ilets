STATUS: CANONICAL
OWNS: conceptual content-instance model, reference contracts between concrete learning content and canonical objects, lesson/unit composition, Practice/Assessment Item representation, Stimulus representation, reusable Error/Remediation Pattern representation, Feedback Artifact semantics, and evidence-record shape
DEPENDS_ON: 03-SKILLS.md, 04-KNOWLEDGE.md, 06-CURRICULUM.md, 07-PRACTICE.md, 08-ASSESSMENT.md
DOES_NOT_OWN: canonical skill/knowledge/band rules, curriculum sequence, practice pedagogy, assessment sufficiency, learner-state transitions, persistence/API/database schemas

# 10 — Content Model

## Purpose

Define how concrete learning/tutoring content **references** the canonical learning model without becoming a second source of truth.

The specification has two levels:

```text
Canonical definitions
  Skill Leaf
  Knowledge Object
  Curriculum Node
  Practice Type
  Assessment Type

Concrete/supporting content
  Learning Unit
  Stimulus
  Practice Item
  Assessment Item
  Error Pattern
  Remediation Pattern
  Feedback Artifact
  Attempt
  Evidence Record
```

Concrete/supporting objects consume canonical definitions. They never redefine them.

## Reference-first rule

Every content object must identify the canonical target(s) that justify its existence.

Examples:

```text
Practice Item
  → practice_type_id
  → target_skill_leaf_ids / target_knowledge_ids
  → curriculum_node_id when assigned in a pathway
```

```text
Assessment Item
  → assessment_type_id
  → target_skill_leaf_ids / target_knowledge_ids
  → target_band when the claim is band-specific
```

If a content object cannot identify what canonical target it serves, it is not valid Blueprint-aligned content.

# `LearningUnit`

A Learning Unit is a delivery-neutral grouping of content around a coherent curriculum purpose.

Conceptual fields:

```text
id
curriculum_node_id
objective_refs
prerequisite_refs
content_sequence
practice_item_refs
assessment_item_refs
error_pattern_refs where useful
remediation_pattern_refs where useful
completion_intent
```

Rules:

- objectives reference Skill/Knowledge semantics instead of rewriting them;
- a unit may contain explanations, examples, practice, retrieval, transfer, remediation, and assessment;
- a unit is not a new curriculum authority;
- multiple units may instantiate the same Curriculum Node for different delivery modes/learner needs;
- completing a unit does not imply mastery unless valid Assessment evidence says so.

The product may call a Learning Unit a "lesson"; the canonical concept remains delivery-neutral.

# `Stimulus`

A Stimulus is source material the learner reads, hears, views, analyzes, or responds to.

Examples include Reading passages, Listening recordings/transcripts, charts/tables/diagrams/maps, Writing prompts, Speaking cue cards/questions, worked examples, model responses, and sentence/paragraph material used in focused drills.

Conceptual fields:

```text
id
kind
variant
source_or_provenance
content
language_properties
difficulty_parameters
rights_or_usage_metadata where required
```

A Stimulus may be reused across Practice and Assessment only when reuse does not invalidate evidence independence.

Examples formerly attached directly to a Knowledge Object are represented as Stimuli or explanatory content referencing that Knowledge Object; the example is replaceable content, not part of object identity.

# `PracticeItem`

A Practice Item is a concrete instance of a canonical Practice Type.

Conceptual fields:

```text
id
practice_type_id
target_skill_leaf_ids
target_knowledge_ids
curriculum_node_id
stimulus_refs
prompt_or_instruction
response_contract
difficulty_parameters
scaffold_level
error_pattern_refs when targeting known errors
remediation_pattern_ref when the item is remedial
feedback_contract
answer_or_model_reference where appropriate
```

## Practice Item invariants

1. `practice_type_id` resolves to `07-PRACTICE.md`.
2. Target IDs resolve to canonical Skill/Knowledge objects.
3. Difficulty may vary without silently changing the canonical target.
4. Scaffolding is represented because heavily scaffolded work cannot later be mistaken for independent assessment evidence.
5. Generated and authored items obey the same reference/quality contract.
6. Item instances are replaceable; Practice Type semantics remain stable.

# `AssessmentItem`

An Assessment Item is a concrete measurement instance of a canonical Assessment Type.

Conceptual fields:

```text
id
assessment_type_id
target_skill_leaf_ids
target_knowledge_ids
target_band when applicable
stimulus_refs
prompt_or_task
response_contract
conditions
scoring_reference
answer_key_or_rubric_reference
independence_group
```

## Assessment Item invariants

1. `assessment_type_id` resolves to `08-ASSESSMENT.md`.
2. The item samples the capability it claims to measure.
3. Band semantics reference `05-BANDS.md`; they are not copied into the item definition.
4. `conditions` record material factors such as timing, assistance, and full-task vs partial-task status.
5. `independence_group` or equivalent metadata lets Assessment judge independence across demonstrations.
6. An item does not decide certification; it produces evidence consumed by `AT-05` and Progression.

# `ErrorPattern`

An Error Pattern is reusable tutoring/content knowledge about a recurring learner mistake or misconception. It is **not** intrinsic Skill/Knowledge identity and may change with population, L1, task context, empirical evidence, or product localization.

Conceptual fields:

```text
id
target_skill_leaf_ids
target_knowledge_ids
pattern_description
applicability_context
example_refs
detection_hints
likely_causes where evidence supports them
severity_or_priority when useful
evidence_or_provenance
locale_or_l1_scope when relevant
```

Rules:

- target references are mandatory;
- an Error Pattern may target one or many canonical objects;
- population/L1-specific patterns must declare that scope rather than changing canonical learning requirements;
- a pattern is supporting tutoring knowledge, not proof that every learner has the error;
- observations from learner attempts may instantiate/match a pattern without changing the pattern's supporting status.

This is the active home for legacy per-leaf `common_errors` and Knowledge `common_misconceptions` semantics when they are useful enough to retain or re-author.

# `RemediationPattern`

A Remediation Pattern is a reusable content/tutoring strategy for correcting a known error or weak target. It is not a prerequisite rule, Band threshold, or canonical Skill definition.

Conceptual fields:

```text
id
target_skill_leaf_ids
target_knowledge_ids
error_pattern_refs when applicable
practice_type_refs
explanation_or_stimulus_refs
scaffold_strategy
recommended_sequence
success_check
variant_or_locale_scope when relevant
evidence_or_provenance
```

Rules:

- remediation must point back to canonical targets;
- Practice Types are referenced rather than redefined;
- success checks may produce formative evidence but certification still follows `08-ASSESSMENT.md`;
- remediation may differ by learner context/L1 without altering canonical target semantics;
- remediation patterns are replaceable and can be empirically improved independently of Skill/Knowledge identity.

This is the active home for legacy leaf `remediation` semantics after the structural refactor.

# `FeedbackArtifact`

A Feedback Artifact is guidance generated from a learner attempt.

Conceptual fields:

```text
attempt_ref
target_refs
observed_performance
gap_or_error
matched_error_pattern_refs when applicable
evidence_reference
feedback_message
recommended_action_refs
remediation_pattern_refs when applicable
confidence when materially relevant
```

Feedback may be authored by AI, deterministic rules, a human, or a combination.

## Feedback invariants

- point to canonical targets rather than inventing a parallel rubric;
- distinguish observed evidence from inference;
- be actionable when formative;
- avoid giving away cognitive work when retrieval/transfer is the target;
- surface material uncertainty;
- recommend canonical next actions/Practice Types/Remediation Patterns by reference where useful.

Feedback is runtime/supporting output, not canonical learning truth.

# `Attempt`

An Attempt is a learner-instance event against a Practice Item or Assessment Item.

The semantic contract must connect:

```text
learner
item
response
conditions
time
automated/human evaluation refs
```

Attempt history and mastery consequences belong to `09-PROGRESSION.md` and runtime implementation. This spec does not prescribe storage.

# `EvidenceRecord`

An Evidence Record is the normalized measurement output from an assessment attempt.

It references:

- assessment item/type;
- target IDs;
- criterion outcomes;
- target band where relevant;
- confidence where required;
- attempt conditions;
- time;
- scorer/evaluation provenance sufficient for validation.

`08-ASSESSMENT.md` determines validity/sufficiency. `09-PROGRESSION.md` consumes the resulting decision.

# Content composition

A typical semantic flow may be composed as:

```text
Curriculum Node
     ↓
Learning Unit
     ├── explanation / examples / Stimuli
     ├── Practice Items
     ├── Error/Remediation Patterns when needed
     ├── retrieval / transfer items
     └── Assessment Items
              ↓
          Evidence Records
              ↓
        learner-state decision
```

This is not a required UI screen flow.

# Difficulty model

Concrete content may express difficulty through linguistic complexity, amount of information, abstraction, distractor strength, response length, target integration, timing pressure, scaffold/cue availability, novelty, and transfer distance.

Difficulty metadata supports selection/calibration. It does not create a Band definition.

# Content quality requirements

A content/supporting object is acceptable only when:

- canonical target references are valid;
- it does not introduce contradictory teaching rules;
- language/task conditions match stated purpose;
- answer key/rubric/model is internally valid where applicable;
- assessment contexts do not leak answers;
- Academic vs General Training scope is respected;
- external provenance/rights are known where needed;
- difficulty is plausible for intended use;
- error/remediation claims declare population/context scope when not universal;
- it can be retired/replaced without changing canonical learning truth.

# Generated content boundary

AI generation is an instance-generation mechanism, not a source of authority. Generated exercises, explanations, error hypotheses, and remediation suggestions must pass the same reference/quality contract as authored content and cannot amend Skill, Knowledge, Band, Curriculum, Practice, Assessment, or Progression semantics.

# No persistence contract

Names and fields here are conceptual contracts. They do not require matching database tables, classes, JSON payloads, or service boundaries one-to-one. Implementation may choose storage/API designs freely while preserving these semantics.