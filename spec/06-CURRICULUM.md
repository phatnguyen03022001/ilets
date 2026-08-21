STATUS: CANONICAL
OWNS: canonical learning pathway, Curriculum Node identity, prerequisite classification for sequencing, recommended ordering, and Band-3-to-9 orchestration of Skill and Knowledge objects
DEPENDS_ON: 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md
DOES_NOT_OWN: Skill or Knowledge definitions, band proficiency thresholds, practice method definitions, assessment sufficiency, learner runtime transitions

# 06 — Curriculum

## Purpose

Define **when canonical learning is sequenced**.

Curriculum orchestrates existing Skill Leaves and Knowledge Objects by stable ID. It does not rewrite their definitions.

The active curriculum preserves **44 stable Curriculum Node IDs** across Bands 3–9.

## Curriculum Node semantic shape

A Curriculum Node owns:

- stable `id`;
- band phase;
- recommended sequence position;
- referenced Skill Leaf IDs;
- referenced Knowledge Object IDs;
- learning focus;
- expected outcome;
- prerequisite relationship;
- relative learning load and planning-duration estimate where useful;
- node exit intent.

A node's exit is a learning intention for the node. Band certification itself is not owned here; it follows the Band and Assessment rules through `09-PROGRESSION.md`.

## Sequencing principles

Sequence is determined by, in order of importance:

1. **Required prerequisites**;
2. target band progression;
3. knowledge-before-dependent-skill where the knowledge is genuinely required;
4. cognitive complexity;
5. learning load and workload balance;
6. integration needs and target-band outcomes.

Within those constraints, a runtime system may adapt ordering to the learner.

## Prerequisite classification

Dependencies used by the curriculum are classified as:

- **Required** — insufficient foundation would make the dependent learning ineffective; a hard gate must have an evidence- or theory-based rationale and no reasonable adaptive workaround.
- **Recommended** — beneficial sequencing information but not a hard gate.
- **Independent** — no gating relationship.

The system should use the minimum number of hard gates needed to preserve learning integrity.

## Band phases

The curriculum is organized into Bands 3–9. A learner may be at different bands in different skills; the runtime progression model may therefore draw from different band phases by skill while still respecting object prerequisites.

The tables below define the canonical nodes and their core references. Detailed generated lessons or practice items are content instances, not additional Curriculum Nodes.

# Band 3 — foundation / structured entry

| Node | Focus | Skill references | Knowledge references |
|---|---|---|---|
| `C-B3-01` | Word classes, clause structure, simple sentences | `W-GRA-01`, `S-GRA-01` | `K-GRA-010`, `K-GRA-001`, `K-GRA-002` |
| `C-B3-02` | Nouns, countability, basic determiners/articles | `W-GRA-05` | `K-GRA-061`, `K-GRA-032`, `K-GRA-040` |
| `C-B3-03` | Basic tense and agreement | `W-GRA-04` | `K-GRA-050`, `K-GRA-051`, `K-GRA-060` |
| `C-B3-04` | Core vocabulary and spelling | `W-LR-01`, `W-LR-05`, `S-LR-01` | `K-VOC-010`, `K-VOC-031` |
| `C-B3-05` | Phoneme foundations and intelligibility | `S-P-01`, `S-P-05` | `K-PHON-010`, `K-PHON-011`, `K-PHON-012` |
| `C-B3-06` | Receptive basics: gist and detail | `L-COMP-01`, `L-COMP-02`, `R-COMP-01`, `R-COMP-02` | `K-VOC-010` |
| `C-B3-07` | Basic receptive question types | `L-QT-01`, `L-QT-05`, `R-QT-05` | — |
| `C-B3-08` | Band-3 consolidation and integration | referenced Band-3 foundation set | — |

Recommended order: `01` establishes grammar foundations; `02` and `03` build on them; vocabulary and phonology can partly run in parallel; receptive fundamentals precede receptive question-type execution; integration closes the phase.

Expected phase outcome: basic communicative meaning across all four skills with foundational grammar, vocabulary, phonology, and receptive capability established.

# Band 4 — basic functional control

| Node | Focus | Skill references | Knowledge references |
|---|---|---|---|
| `C-B4-01` | Writing accuracy, punctuation, paragraphing | `W-GRA-01`, `W-GRA-06`, `W-CC-01` | — |
| `C-B4-02` | Task foundations: key features and prompt analysis | `W-TA-01`, `W-TR-01` | — |
| `C-B4-03` | Topic vocabulary and word choice | `W-LR-01`, `S-LR-01`, `S-LR-04` | `K-VOC-012`, `K-VOC-040` |
| `C-B4-04` | Speaking continuity and basic accuracy | `S-FC-01`, `S-GRA-04`, `S-P-05` | `K-GRA-010` |
| `C-B4-05` | Receptive main ideas and details | `L-COMP-01`, `L-COMP-02`, `R-COMP-01`, `R-COMP-02` | `K-VOC-012` |
| `C-B4-06` | Basic receptive question types at Band-4 demand | `L-QT-01`, `L-QT-05`, `R-QT-05` | — |
| `C-B4-07` | Band-4 consolidation and integration | referenced Band-4 set | — |

Expected phase outcome: controlled basic writing structure, explicit task awareness, sustained basic speaking, and more reliable receptive gist/detail handling.

# Band 5 — expanding control

| Node | Focus | Skill references | Knowledge references |
|---|---|---|---|
| `C-B5-01` | Compound and complex sentence foundation | `W-GRA-02`, `W-GRA-03` | `K-GRA-003`, `K-GRA-020`, `K-GRA-004`, `K-GRA-021` |
| `C-B5-02` | Cohesion and logical progression | `W-CC-02`, `W-CC-03`, `W-CC-04` | `K-GRA-022`, `K-GRA-033` |
| `C-B5-03` | Task-2 position, development, relevance | `W-TR-02`, `W-TR-03`, `W-TR-04` | `K-VOC-012` |
| `C-B5-04` | Task-1 reporting with data | `W-TA-03` | `K-GRA-050`, `K-GRA-051` |
| `C-B5-05` | Paraphrase, word formation, spelling accuracy | `W-LR-03`, `W-LR-05` | `K-VOC-030`, `K-VOC-031` |
| `C-B5-06` | Speaking fluency, complexity, and pronunciation | `S-FC-02`, `S-FC-03`, `S-FC-04`, `S-FC-05`, `S-GRA-02`, `S-LR-02`, `S-P-02`, `S-P-03` | `K-GRA-021`, `K-VOC-012` |
| `C-B5-07` | Receptive inference, paraphrase, distractors | `L-COMP-03`, `L-COMP-04`, `L-COMP-05`, `R-COMP-03`, `R-COMP-04`, `R-COMP-05` | `K-VOC-011` |
| `C-B5-08` | Higher-order receptive question types | `L-QT-02`, `L-QT-03`, `L-QT-04`, `R-QT-01`, `R-QT-02`, `R-QT-04`, `R-QT-06` | — |
| `C-B5-09` | Band-5 consolidation and integration | referenced Band-5 set | — |

Recommended dependencies:

- `C-B5-01` before writing nodes requiring complex structure;
- `C-B5-02` before high-demand integrated writing;
- receptive comprehension `C-B5-07` before the associated strategy node `C-B5-08`.

# Band 6 — competent clarity

| Node | Focus | Skill references | Knowledge references |
|---|---|---|---|
| `C-B6-01` | Grammatical accuracy and variety | `W-GRA-03`, `W-GRA-02`, `S-GRA-03` | `K-GRA-004`, `K-GRA-003` |
| `C-B6-02` | Task-1 overview and flexible reference | `W-TA-02`, `W-CC-04` | `K-GRA-033` |
| `C-B6-03` | Collocation and idiomatic resource | `W-LR-02`, `S-LR-03` | `K-VOC-020`, `K-VOC-021` |
| `C-B6-04` | Extended and dense receptive content | `L-COMP-06`, `R-COMP-06` | `K-VOC-011` |
| `C-B6-05` | Writer-view classification: Yes/No/Not Given | `R-QT-03` | — |
| `C-B6-06` | Band-6 consolidation and integration | referenced Band-6 set | — |

# Band 7 — flexible good control

| Node | Focus | Skill references | Knowledge references |
|---|---|---|---|
| `C-B7-01` | Structural flexibility and error-free frequency | `W-GRA-07`, `W-GRA-03` | `K-GRA-005`, `K-GRA-006`, `K-GRA-007` |
| `C-B7-02` | Lexical sophistication | `W-LR-04`, `S-LR-03` | `K-VOC-021`, `K-VOC-011` |
| `C-B7-03` | Cohesion, position, and development mastery | `W-CC-03`, `W-CC-04`, `W-TR-02`, `W-TR-03` | `K-GRA-022`, `K-GRA-033` |
| `C-B7-04` | Receptive inference and structural understanding | `R-COMP-04`, `R-COMP-05`, `L-COMP-03`, `L-COMP-04` | `K-VOC-011` |
| `C-B7-05` | Pronunciation range and flexibility | `S-P-03`, `S-P-04` | `K-PHON-030`, `K-PHON-040` |
| `C-B7-06` | Band-7 consolidation and integration | referenced Band-7 set | — |

# Band 8 — very strong control

| Node | Focus | Skill references | Knowledge references |
|---|---|---|---|
| `C-B8-01` | Wide flexible language and near-error-free accuracy | `W-GRA-07`, `W-LR-04`, `W-GRA-01`, `W-GRA-03`, `S-GRA-03` | `K-GRA-005`, `K-GRA-007` |
| `C-B8-02` | Skillful cohesion and fully developed response | `W-CC-02`, `W-CC-03`, `W-CC-04`, `W-TR-03`, `W-TA-02` | `K-GRA-022`, `K-GRA-033` |
| `C-B8-03` | Effortless speaking fluency and wide resource | `S-FC-02`, `S-FC-04`, `S-GRA-03`, `S-LR-03`, `S-P-05` | `K-PHON-040`, `K-PHON-041` |
| `C-B8-04` | Effortless complex receptive comprehension | `L-COMP-06`, `R-COMP-06`, `R-COMP-04` | `K-VOC-011` |
| `C-B8-05` | Band-8 consolidation and integration | referenced Band-8 set | — |

# Band 9 — ceiling

| Node | Focus | Skill references | Knowledge references |
|---|---|---|---|
| `C-B9-01` | Full flexibility and precision | `W-GRA-07`, `W-LR-04`, `S-GRA-03`, `S-LR-03`, `S-P-05` | `K-GRA-005`, `K-GRA-007`, `K-VOC-021` |
| `C-B9-02` | Integrated exam-level mastery | `W-TA-02`, `W-TR-03`, `W-CC-03`, `R-COMP-04`, `L-COMP-06` | — |
| `C-B9-03` | Band-9 ceiling integration | full target set | — |

Band 9 introduces no new fundamental object family. It integrates and refines previously acquired capability to the ceiling threshold.

## Learning phases and curriculum placement

Curriculum sequencing works with the canonical learning phases owned as practice semantics in `07-PRACTICE.md`:

- acquisition;
- consolidation;
- retrieval;
- transfer;
- fluency;
- exam readiness.

A Curriculum Node may be practiced through multiple phases over time. The node itself does not duplicate Practice Type definitions.

## Adaptive sequencing

The tables define a canonical recommended pathway, not an immutable learner timeline.

A runtime system may reorder or interleave nodes when:

- all Required prerequisites remain satisfied;
- target-band outcomes remain complete;
- adaptation does not bypass a weak canonical capability;
- the learner's evidence supports the change.

Adaptation changes **path**, not **outcome**.

## Coverage invariant

Every Curriculum Node must reference canonical object IDs. New nodes are justified only when they provide a distinct orchestration step, not merely because a lesson or exercise is needed.

Concrete lessons and exercise instances belong to `10-CONTENT-MODEL.md`, not to this node registry.
