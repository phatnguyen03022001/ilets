STATUS: CANONICAL
OWNS: canonical learning pathway, Curriculum Node identity, recommended ordering, prerequisite classification for sequence, variant route overlays, and Band-3-to-9 orchestration of Skill and Knowledge objects
DEPENDS_ON: 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md
DOES_NOT_OWN: Skill/Knowledge definitions, band thresholds, practice types, assessment sufficiency, learner-state transitions, or uncalibrated duration/load estimates

# 06 — Curriculum

## Purpose

Define **when canonical learning is sequenced**.

Curriculum orchestrates existing Skill Leaves and Knowledge Objects by stable ID. It never creates a parallel definition of them.

The active base curriculum preserves **44 stable Curriculum Node IDs** across Bands 3–9. Variant overlays modify the target set of applicable nodes without duplicating the shared Band-3-to-9 pathway.

## Canonical Curriculum Node fields

A node canonically owns stable `id`/band phase, recommended sequence position, referenced Skill and Knowledge IDs, learning focus, expected outcome, explicit recommended node dependency where needed, and node exit intent.

Legacy hour estimates and relative load values remain historical heuristics until empirically calibrated; they are not active canonical truth.

Node exit intent means the node's intended learning-completion condition. **Band certification is not owned here** and never depends on completing a synchronized four-skill consolidation node.

# Dependency semantics

Dependencies are classified as:

- **Required** — intrinsic Skill→Skill, Knowledge→Knowledge, or Skill→Knowledge prerequisite whose unresolved state hard-gates dependent learning;
- **Recommended** — useful Curriculum Node ordering but non-blocking when evidence justifies another route;
- **Independent** — no dependency relation.

Canonical interpretation:

1. every `requires` edge defined by `03-SKILLS.md` and every prerequisite/resolution edge defined by `04-KNOWLEDGE.md` is a **Required prerequisite**;
2. the `Depends` column in this file contains **only stable `C-*` Curriculum Node IDs** and records **Recommended** sequencing;
3. `Depends` never contains Skill IDs, Knowledge IDs, prose aliases, band labels, or phrases such as “prior grammar”; Required object dependencies are inherited from their canonical owners rather than duplicated here;
4. a new Required intrinsic object edge belongs in `03-SKILLS.md` or `04-KNOWLEDGE.md`; this file may recommend node order but may not silently promote that recommendation into a hard gate;
5. runtime hard-gate enforcement belongs to `09-PROGRESSION.md`;
6. a variant overlay may add/substitute variant-specific targets but may not remove shared capability genuinely required by that variant.

Therefore two implementations reading the same node registry must derive the same recommended Curriculum DAG. Human-readable Focus/Expected-outcome text never creates an implicit dependency edge.

# Deterministic selector tokens

The registry may use only the following declared compile-time selector tokens in target columns. They are **not canonical learning-object IDs** and must be fully expanded before runtime/content/assessment identity is materialized.

```text
@task1-analysis
@task1-fulfilment
@task1-global-control
@variant-reading-advanced-knowledge
@band-phase-skill-set
@band-phase-knowledge-set
```

Rules:

1. `@task1-*` selectors resolve from `TargetProfile.test_variant` using the Variant route overlay below;
2. `@variant-reading-advanced-knowledge` resolves from the selected variant using the overlay below;
3. `@band-phase-skill-set` is the deterministic union of all fully resolved Skill targets in preceding nodes of that same Band phase;
4. `@band-phase-knowledge-set` is the deterministic union of all fully resolved explicit Knowledge targets in preceding nodes of that same Band phase;
5. selector expansion is recursive until only canonical IDs remain; an unresolved selector at execution/content/assessment time is invalid;
6. no new selector token may be invented ad hoc in a node row; adding one requires updating this selector contract and its complete resolution table;
7. Required Skill/Knowledge dependencies inherited from `03-SKILLS.md`/`04-KNOWLEDGE.md` remain applicable after selector expansion and are not duplicated into the union.

# Canonical node registry

## Band 3 — foundation / structured entry

| Node | Focus | Skill targets | Knowledge targets | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B3-01` | Word classes, clause structure, simple sentences | `W-GRA-01`, `S-GRA-01` | `K-GRA-010`, `K-GRA-001`, `K-GRA-002` | — | Form accurate basic written/spoken sentence cores; foundational grammar objects acquired. |
| `C-B3-02` | Nouns, countability, determiners/articles | `W-GRA-05` | `K-GRA-061`, `K-GRA-032`, `K-GRA-040` | `C-B3-01` | Use basic countability/determiner/article choices without destroying meaning. |
| `C-B3-03` | Basic tense and agreement | `W-GRA-04` | `K-GRA-050`, `K-GRA-051`, `K-GRA-060` | `C-B3-01` | Use present/past basics and subject–verb agreement in simple production. |
| `C-B3-04` | Core vocabulary and spelling | `W-LR-01`, `W-LR-05`, `S-LR-01` | `K-VOC-010`, `K-VOC-031` | — | Produce/recognize high-frequency vocabulary with basic spelling control. |
| `C-B3-05` | Phoneme foundations and intelligibility | `S-P-01`, `S-P-05` | `K-PHON-010`, `K-PHON-011`, `K-PHON-012` | — | Produce core sound contrasts clearly enough for broad intelligibility. |
| `C-B3-06` | Receptive gist and detail | `L-COMP-01`, `L-COMP-02`, `R-COMP-01`, `R-COMP-02` | `K-VOC-010` | `C-B3-04` | Identify gist and key details in simple spoken/written input. |
| `C-B3-07` | Basic receptive question types | `L-QT-01`, `L-QT-05`, `R-QT-05` | — | `C-B3-06` | Complete basic receptive items within stated form/word limits. |
| `C-B3-08` | Foundation integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B3-01`, `C-B3-02`, `C-B3-03`, `C-B3-04`, `C-B3-05`, `C-B3-06`, `C-B3-07` | Integrate Band-3 foundations; orchestration checkpoint, not a four-skill certification gate. |

## Band 4 — basic functional control

| Node | Focus | Skill targets | Knowledge targets | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B4-01` | Writing accuracy, punctuation, paragraphing | `W-GRA-01`, `W-GRA-06`, `W-CC-01` | — | `C-B3-01`, `C-B3-03` | Produce simple, punctuated, paragraphed writing. |
| `C-B4-02` | Task foundations: Task-1 analysis and Task-2 prompt analysis | `@task1-analysis`, `W-TR-01` | — | `C-B3-08` | Identify the Task-1 construct requirements for the selected variant and deconstruct Task-2 prompts. |
| `C-B4-03` | Topic vocabulary and word choice | `W-LR-01`, `S-LR-01`, `S-LR-04` | `K-VOC-012`, `K-VOC-040` | `C-B3-04` | Discuss/write common topics with adequate and context-appropriate vocabulary. |
| `C-B4-04` | Speaking continuity and basic accuracy | `S-FC-01`, `S-GRA-04`, `S-P-05` | `K-GRA-010` | `C-B3-01`, `C-B3-05` | Sustain basic speech on familiar topics with clear overall meaning. |
| `C-B4-05` | Receptive main ideas and details | `L-COMP-01`, `L-COMP-02`, `R-COMP-01`, `R-COMP-02` | `K-VOC-012` | `C-B3-06` | Handle gist/detail in moderately more complex input. |
| `C-B4-06` | Basic receptive question types at Band-4 demand | `L-QT-01`, `L-QT-05`, `R-QT-05` | — | `C-B4-05` | Execute basic receptive item strategies at current demand. |
| `C-B4-07` | Band-4 integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B4-01`, `C-B4-02`, `C-B4-03`, `C-B4-04`, `C-B4-05`, `C-B4-06` | Integrate Band-4 work; non-certifying orchestration checkpoint. |

## Band 5 — expanding control

| Node | Focus | Skill targets | Knowledge targets | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B5-01` | Compound and complex sentence foundation | `W-GRA-02`, `W-GRA-03` | `K-GRA-003`, `K-GRA-020`, `K-GRA-004`, `K-GRA-021` | `C-B4-01` | Produce compound/complex sentences with usable control. |
| `C-B5-02` | Cohesion and logical progression | `W-CC-02`, `W-CC-03`, `W-CC-04` | `K-GRA-022`, `K-GRA-033` | `C-B4-01`, `C-B5-01` | Organize responses with progression, linking, reference and substitution. |
| `C-B5-03` | Task-2 position, development, relevance | `W-TR-02`, `W-TR-03`, `W-TR-04` | `K-VOC-012` | `C-B4-02`, `C-B5-02` | Write a relevant Task-2 response with a clear position and supported ideas. |
| `C-B5-04` | Task-1 content fulfilment | `@task1-fulfilment` | — | `C-B4-02` | Fulfil the selected variant's Task-1 content requirement with relevant support. |
| `C-B5-05` | Paraphrase, word formation, spelling | `W-LR-03`, `W-LR-05` | `K-VOC-030`, `K-VOC-031` | `C-B4-03` | Paraphrase without meaning loss and control productive word form/spelling. |
| `C-B5-06` | Speaking fluency, complexity, pronunciation | `S-FC-02`, `S-FC-03`, `S-FC-04`, `S-FC-05`, `S-GRA-02`, `S-LR-02`, `S-P-02`, `S-P-03` | `K-GRA-021`, `K-VOC-012` | `C-B4-04`, `C-B5-01` | Sustain an extended coherent turn using more complex language and controlled stress/intonation. |
| `C-B5-07` | Receptive inference, paraphrase, distractors | `L-COMP-03`, `L-COMP-04`, `L-COMP-05`, `R-COMP-03`, `R-COMP-04`, `R-COMP-05` | — | `C-B4-05` | Handle inference, paraphrase, structure and distractors more reliably. |
| `C-B5-08` | Higher-order receptive question types | `L-QT-02`, `L-QT-03`, `L-QT-04`, `R-QT-01`, `R-QT-02`, `R-QT-04`, `R-QT-06` | — | `C-B5-07` | Apply higher-order question-type strategies under realistic constraints. |
| `C-B5-09` | Band-5 integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B5-01`, `C-B5-02`, `C-B5-03`, `C-B5-04`, `C-B5-05`, `C-B5-06`, `C-B5-07`, `C-B5-08` | Integrate Band-5 capabilities; non-certifying orchestration checkpoint. |

## Band 6 — competent clarity

| Node | Focus | Skill targets | Knowledge targets | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B6-01` | Grammatical accuracy and variety | `W-GRA-03`, `W-GRA-02`, `S-GRA-03` | `K-GRA-004`, `K-GRA-003` | `C-B5-01` | Increase accuracy and structural variety in productive language. |
| `C-B6-02` | Task-1 global control and flexible reference | `@task1-global-control`, `W-CC-04` | — | `C-B5-02`, `C-B5-04` | Produce the selected variant's required global Task-1 control and reduce repetition through controlled reference. |
| `C-B6-03` | Collocation and idiomatic resource | `W-LR-02`, `S-LR-03` | `K-VOC-020`, `K-VOC-021` | `C-B4-03` | Use collocations accurately and begin appropriate less-common/idiomatic use. |
| `C-B6-04` | Extended/dense receptive content and variant transfer | `L-COMP-06`, `R-COMP-06` | `@variant-reading-advanced-knowledge` | `C-B5-07` | Handle extended/dense input and transfer shared Reading capability into the selected variant's required contexts. |
| `C-B6-05` | Writer-view classification | `R-QT-03` | — | `C-B5-08` | Distinguish writer agreement/disagreement/absence accurately. |
| `C-B6-06` | Band-6 integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B6-01`, `C-B6-02`, `C-B6-03`, `C-B6-04`, `C-B6-05` | Integrate Band-6 accuracy/clarity; non-certifying orchestration checkpoint. |

## Band 7 — flexible good control

| Node | Focus | Skill targets | Knowledge targets | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B7-01` | Structural flexibility and error-free frequency | `W-GRA-07`, `W-GRA-03` | `K-GRA-005`, `K-GRA-006`, `K-GRA-007` | `C-B6-01` | Produce varied complex structures flexibly with frequent accurate sentences. |
| `C-B7-02` | Lexical sophistication | `W-LR-04`, `S-LR-03` | `K-VOC-021` | `C-B6-03` | Use less-common/idiomatic resource with increasing precision and appropriateness. |
| `C-B7-03` | Cohesion, position, development | `W-CC-03`, `W-CC-04`, `W-TR-02`, `W-TR-03` | `K-GRA-022`, `K-GRA-033` | `C-B5-03`, `C-B6-02` | Sustain clear developed positions with flexible organization/cohesion. |
| `C-B7-04` | Receptive inference and structure | `R-COMP-04`, `R-COMP-05`, `L-COMP-03`, `L-COMP-04` | — | `C-B6-04` | Handle inference, stance, paraphrase and structure reliably. |
| `C-B7-05` | Pronunciation range and flexibility | `S-P-03`, `S-P-04` | `K-PHON-030`, `K-PHON-040` | `C-B5-06` | Sustain useful intonation, connected speech and chunking across extended turns. |
| `C-B7-06` | Band-7 integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B7-01`, `C-B7-02`, `C-B7-03`, `C-B7-04`, `C-B7-05` | Integrate Band-7 flexible control; non-certifying orchestration checkpoint. |

## Band 8 — very strong control

| Node | Focus | Skill targets | Knowledge targets | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B8-01` | Wide flexible language and near-error-free accuracy | `W-GRA-07`, `W-LR-04`, `W-GRA-01`, `W-GRA-03`, `S-GRA-03` | `K-GRA-005`, `K-GRA-007` | `C-B7-01`, `C-B7-02` | Produce wide, flexible, highly accurate language. |
| `C-B8-02` | Skillful cohesion and fully developed response | `W-CC-02`, `W-CC-03`, `W-CC-04`, `W-TR-03`, `@task1-global-control` | `K-GRA-022`, `K-GRA-033` | `C-B7-03` | Produce well-developed responses whose organization is easy to follow, including variant-appropriate Task-1 control. |
| `C-B8-03` | Effortless speaking fluency and wide resource | `S-FC-02`, `S-FC-04`, `S-GRA-03`, `S-LR-03`, `S-P-05` | `K-PHON-040`, `K-PHON-041` | `C-B7-05` | Sustain fluent, wide-resource speech that is easily understood. |
| `C-B8-04` | Complex receptive comprehension | `L-COMP-06`, `R-COMP-06`, `R-COMP-04` | — | `C-B7-04` | Handle detailed, abstract and complex input with little difficulty across required variant contexts. |
| `C-B8-05` | Band-8 integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B8-01`, `C-B8-02`, `C-B8-03`, `C-B8-04` | Integrate Band-8 high competence; non-certifying orchestration checkpoint. |

## Band 9 — ceiling

| Node | Focus | Skill targets | Knowledge targets | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B9-01` | Full flexibility and precision | `W-GRA-07`, `W-LR-04`, `S-GRA-03`, `S-LR-03`, `S-P-05` | `K-GRA-005`, `K-GRA-007`, `K-VOC-021` | `C-B8-01`, `C-B8-02`, `C-B8-03` | Approach full flexibility, precision and sustained intelligibility. |
| `C-B9-02` | Integrated exam-level mastery | `@task1-global-control`, `W-TR-03`, `W-CC-03`, `R-COMP-04`, `L-COMP-06` | — | `C-B8-05` | Integrate ceiling-level capability under independent exam-like demand for the selected variant. |
| `C-B9-03` | Ceiling integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B9-01`, `C-B9-02` | Final integration checkpoint; certification still requires evidence under Bands/Assessment/Progression. |

# Variant route overlay

The base node registry remains one shared curriculum. Selector expansion is deterministic from `TargetProfile.test_variant`.

## Selector resolution table

| Selector | Academic resolution | General Training resolution |
|---|---|---|
| `@task1-analysis` | `W-TA-01` | `W-GT1-01` |
| `@task1-fulfilment` | `W-TA-03` | `W-GT1-03` |
| `@task1-global-control` | `W-TA-02` | `W-GT1-02` |
| `@variant-reading-advanced-knowledge` | `K-VOC-011` | empty set |

The empty-set GT resolution means there is no **single additional canonical Knowledge Object** universally required merely because the learner performs dense GT Reading. GT Section-1/2/3 vocabulary/register demand is represented through existing applicable Knowledge plus concrete content/context under `10-CONTENT-MODEL.md`; implementation must not invent a hidden GT vocabulary prerequisite.

Academic Task-1 tenses/data language and concrete visual-specific lexical choices are selected through the already-owned Skill→Knowledge dependencies and concrete content requirements; they are not extra undeclared selector semantics.

## Reading context conditions

Variant-specific Reading transfer is a **context/content/readiness condition**, not a hidden Knowledge dependency:

- Academic readiness uses Academic passage/context distribution and full applicable official-family coverage;
- GT readiness samples `CTX-READING-GT-S1-EVERYDAY`, `CTX-READING-GT-S2-WORKPLACE`, and `CTX-READING-GT-S3-GENERAL-INTEREST` as required by the scoped evidence/support claim.

These context conditions are consumed by Content/Assessment/Coverage owners. They do not create prose edges in the Curriculum `Depends` graph.

### GT phase expectations

- Bands 3–4: introduce recipient/purpose identification and basic relationship/register awareness; use accessible everyday/workplace Reading contexts.
- Band 5: require all GT Task-1 prompt bullet points to be recognized/attempted and broaden Reading across Section-1/2/3 context classes.
- Band 6: require generally appropriate register and complete task-purpose fulfilment; independent timed GT Reading must include the official section/context structure.
- Bands 7–9: increase flexibility, precision, audience sensitivity, and unseen transfer without inventing extra GT-only language criteria.

## Variant-exclusion rule

A learner targeting General Training does **not** need Academic visual-specific `W-TA-*` leaves to satisfy GT Writing Task 1. A learner targeting Academic does **not** need `W-GT1-*` leaves to satisfy Academic Writing Task 1.

Shared Writing, Reading, Listening, Speaking, Grammar, Vocabulary, and Phonology targets remain reusable where the construct is shared.

# Adaptive sequencing

Runtime may reorder/interleave nodes when Required prerequisites remain satisfied, target outcomes remain complete, variant requirements remain complete, weak canonical capabilities are not bypassed, and learner evidence supports the change.

A learner may draw from different band phases by skill because progression is per skill. Integration nodes never force synchronized four-skill certification.

Recommended `Depends` edges influence default ordering and explanation; they are not converted into hard gates merely because they are machine-readable.

# Coverage and validation invariant

Every Curriculum Node target cell contains only:

- canonical Skill/Knowledge IDs;
- `—` for no explicit target; or
- one of the selector tokens declared in this document.

Every non-empty `Depends` cell contains only stable `C-*` IDs that exist in this registry.

Before a node is materialized into content/runtime work:

```text
node
  ↓
resolve variant selectors
  ↓
resolve band-phase union selectors
  ↓
validate all canonical IDs
  ↓
apply inherited Required Skill/Knowledge prerequisites
```

Static validation should fail on:

- unknown canonical IDs;
- unknown selector tokens;
- unresolved selectors at the execution boundary;
- prose/free-text in target/dependency identity positions;
- a `Depends` reference to a missing node;
- a Recommended node cycle;
- a variant expansion that includes the wrong Task-1 construct.

New nodes require a distinct orchestration purpose, not merely another lesson/exercise. Concrete lessons and exercise instances belong to `10-CONTENT-MODEL.md`.