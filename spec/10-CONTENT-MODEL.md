STATUS: CANONICAL
OWNS: conceptual content-instance model, reference contracts between concrete learning content and canonical objects, lesson/unit composition, Practice Item and Assessment Item representation, stimulus representation, and feedback-artifact semantics
DEPENDS_ON: 03-SKILLS.md, 04-KNOWLEDGE.md, 06-CURRICULUM.md, 07-PRACTICE.md, 08-ASSESSMENT.md
DOES_NOT_OWN: canonical skill/knowledge/band rules, curriculum sequence, practice pedagogy, assessment sufficiency, learner-state transitions, persistence/API/database schemas

# 10 — Content Model

## Purpose

Define how concrete learning content **references** the canonical learning model without becoming a second source of truth.

The specification has two levels:

```text
Canonical definitions
  Skill Leaf
  Knowledge Object
  Curriculum Node
  Practice Type
  Assessment Type

Concrete content instances
  Learning Unit
  Stimulus
  Practice Item
  Assessment Item
  Feedback Artifact
```

Concrete instances consume canonical definitions. They never redefine them.

## Reference-first rule

A concrete content object must identify the canonical objects that justify its existence.

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

## `LearningUnit`

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
completion_intent
```

Rules:

- the unit's objectives reference Skill/Knowledge semantics instead of rewriting them;
- the unit may contain explanation, worked examples, practice, retrieval, and assessment;
- a unit is not itself a new curriculum authority;
- multiple Learning Units may instantiate the same Curriculum Node for different delivery modes or learner needs;
- completing a unit does not imply mastery unless valid Assessment evidence says so.

The word "lesson" may be used as a product-level presentation name for a Learning Unit, but the canonical concept is delivery-neutral.

## `Stimulus`

A Stimulus is source material the learner reads, hears, views, or responds to.

Examples:

- Reading passage;
- Listening recording/transcript;
- chart/table/diagram/map;
- Writing Task-2 prompt;
- Speaking cue card/question set;
- worked example;
- sentence or paragraph used in a focused drill.

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

A Stimulus may be reused across Practice and Assessment Items only when reuse does not invalidate evidence independence.

## `PracticeItem`

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
feedback_contract
answer_or_model_reference where appropriate
```

### Practice Item invariants

1. `practice_type_id` must resolve to `07-PRACTICE.md`.
2. Target IDs must resolve to canonical Skill/Knowledge objects.
3. The item may vary difficulty without silently changing the canonical target.
4. Scaffolding must be represented because a heavily scaffolded item cannot later be mistaken for independent assessment evidence.
5. Generated content is valid only if it satisfies the same reference and quality contract as authored content.
6. Item instances are replaceable; canonical Practice Type semantics remain stable.

## `AssessmentItem`

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

### Assessment Item invariants

1. `assessment_type_id` resolves to `08-ASSESSMENT.md`.
2. The item samples the capability it claims to measure.
3. Target Band semantics reference `05-BANDS.md`; they are not copied into the item definition.
4. `conditions` record material factors such as timing, assistance, and full-task vs partial-task status.
5. `independence_group` or equivalent metadata must allow the Assessment layer to determine whether two demonstrations are sufficiently independent.
6. An assessment item does not decide certification; it produces evidence consumed by `AT-05` and Progression.

## `FeedbackArtifact`

A Feedback Artifact is guidance generated from a learner attempt.

Conceptual fields:

```text
attempt_ref
target_refs
observed_performance
gap_or_error
evidence_reference
feedback_message
recommended_action_refs
confidence when materially relevant
```

Feedback may be authored by AI, deterministic rules, a human, or a combination.

### Feedback invariants

- point to the canonical target rather than inventing a parallel rubric;
- distinguish observed evidence from inference;
- be actionable when the purpose is formative learning;
- avoid giving away the cognitive work when retrieval/transfer is the target;
- surface material uncertainty;
- recommend canonical next actions or Practice Types by reference where possible.

Feedback is runtime/supporting output, not canonical truth.

## `Attempt`

An Attempt is a learner-instance event against a Practice Item or Assessment Item.

The content model requires enough conceptual identity to connect:

```text
learner
item
response
conditions
time
automated/human evaluation refs
```

Attempt history, mastery consequences, and learner state belong to `09-PROGRESSION.md` and implementation runtime. This spec does not prescribe storage.

## `EvidenceRecord`

An Evidence Record is the normalized measurement output from an assessment attempt.

It references:

- assessment item/type;
- target IDs;
- criterion outcomes;
- band target where relevant;
- confidence where required;
- attempt conditions;
- time;
- scorer/evaluation provenance sufficient for validation.

The Assessment owner determines whether evidence is valid and sufficient. The Progression owner consumes the resulting decision.

## Content composition

A typical learning flow may be composed as:

```text
Curriculum Node
     ↓
Learning Unit
     ├── explanation / examples
     ├── Practice Items
     ├── retrieval / transfer items
     └── Assessment Items
              ↓
          Evidence Records
              ↓
        learner-state decision
```

This is a semantic composition, not a required UI screen flow.

## Difficulty model

Concrete content may express difficulty through dimensions such as:

- linguistic complexity;
- amount of information;
- abstraction;
- distractor strength;
- response length;
- integration across targets;
- timing pressure;
- scaffold/cue availability;
- novelty;
- transfer distance.

Difficulty metadata supports selection and calibration. It does not create a new band definition.

## Content quality requirements

A content instance is acceptable only when:

- its canonical target references are valid;
- it does not introduce contradictory teaching rules;
- its language/task conditions match its stated purpose;
- its answer key/rubric/model is internally valid where applicable;
- it does not accidentally leak answers in assessment contexts;
- it respects Academic vs General Training variant scope;
- its provenance/rights are known when external material is used;
- its difficulty is plausible for the intended use;
- it can be retired or replaced without changing canonical learning truth.

## Generated content boundary

AI-generated content is an instance-generation mechanism, not a source of authority.

Generated content must pass the same validation contract as authored content. A generated exercise or explanation cannot amend Skill, Knowledge, Band, Curriculum, Practice, Assessment, or Progression semantics.

## No persistence contract

Names and fields in this document are conceptual object contracts. They do not require matching database tables, classes, JSON payloads, or service boundaries one-to-one.

Implementation should preserve the semantics while remaining free to choose appropriate storage and API designs later.
