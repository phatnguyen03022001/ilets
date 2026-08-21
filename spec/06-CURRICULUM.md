STATUS: CANONICAL
OWNS: canonical learning pathway, Curriculum Node identity, recommended ordering, prerequisite classification for sequence, and Band-3-to-9 orchestration of Skill and Knowledge objects
DEPENDS_ON: 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md
DOES_NOT_OWN: Skill/Knowledge definitions, band thresholds, practice types, assessment sufficiency, learner-state transitions, or uncalibrated duration/load estimates

# 06 — Curriculum

## Purpose

Define **when canonical learning is sequenced**.

Curriculum orchestrates existing Skill Leaves and Knowledge Objects by stable ID. It never creates a parallel definition of them.

The active curriculum preserves **44 stable Curriculum Node IDs** across Bands 3–9.

## Canonical Curriculum Node fields

A node canonically owns:

- stable `id` and band phase;
- recommended sequence position;
- referenced Skill Leaf IDs;
- referenced Knowledge Object IDs;
- learning focus;
- expected outcome;
- explicit sequencing dependency where needed;
- node exit intent.

Legacy hour estimates and relative load values remain historical heuristics until empirically calibrated; they are not active canonical truth.

Node exit intent means "the node's intended learning completion condition". **Band certification is not owned here** and never depends on completing a synchronized four-skill consolidation node.

## Sequencing policy

Order is constrained by:

1. Required prerequisites;
2. target band progression;
3. knowledge-before-dependent-skill where genuinely required;
4. cognitive complexity;
5. workload balance;
6. integration needs.

Dependencies are classified as:

- **Required** — missing foundation makes dependent learning ineffective and no reasonable adaptive workaround exists;
- **Recommended** — useful ordering but non-blocking;
- **Independent** — no gate.

Use the minimum justified hard gates. Runtime enforcement is owned by `09-PROGRESSION.md`.

# Canonical node registry

`Depends` below records explicit sequencing relationships. In addition, every node inherits prerequisite semantics from the referenced Skill/Knowledge graphs.

## Band 3 — foundation / structured entry

| Node | Focus | Skill targets | Knowledge | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B3-01` | Word classes, clause structure, simple sentences | `W-GRA-01`, `S-GRA-01` | `K-GRA-010`, `001`, `002` | — | Form accurate basic written/spoken sentence cores; foundational grammar objects acquired. |
| `C-B3-02` | Nouns, countability, determiners/articles | `W-GRA-05` | `K-GRA-061`, `032`, `040` | `C-B3-01` | Use basic countability/determiner/article choices without destroying meaning. |
| `C-B3-03` | Basic tense and agreement | `W-GRA-04` | `K-GRA-050`, `051`, `060` | `C-B3-01` | Use present/past basics and subject–verb agreement in simple production. |
| `C-B3-04` | Core vocabulary and spelling | `W-LR-01`, `W-LR-05`, `S-LR-01` | `K-VOC-010`, `031` | — | Produce/recognize high-frequency vocabulary with basic spelling control. |
| `C-B3-05` | Phoneme foundations and intelligibility | `S-P-01`, `S-P-05` | `K-PHON-010`, `011`, `012` | — | Produce core sound contrasts clearly enough for broad intelligibility. |
| `C-B3-06` | Receptive gist and detail | `L-COMP-01`, `02`, `R-COMP-01`, `02` | `K-VOC-010` | `C-B3-04` recommended | Identify gist and key details in simple spoken/written input. |
| `C-B3-07` | Basic receptive question types | `L-QT-01`, `L-QT-05`, `R-QT-05` | — | `C-B3-06` | Complete basic receptive items within stated form/word limits. |
| `C-B3-08` | Foundation integration | Band-3 foundation set | — | `C-B3-01..07` recommended | Integrate Band-3 foundations; this is an orchestration checkpoint, not a four-skill certification gate. |

## Band 4 — basic functional control

| Node | Focus | Skill targets | Knowledge | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B4-01` | Writing accuracy, punctuation, paragraphing | `W-GRA-01`, `W-GRA-06`, `W-CC-01` | — | Band-3 sentence foundation | Produce simple, punctuated, paragraphed writing. |
| `C-B4-02` | Task foundations: key features and prompt analysis | `W-TA-01`, `W-TR-01` | — | Band-3 writing foundation | Identify Task-1 key features and deconstruct Task-2 prompts. |
| `C-B4-03` | Topic vocabulary and word choice | `W-LR-01`, `S-LR-01`, `S-LR-04` | `K-VOC-012`, `040` | `K-VOC-010` | Discuss/write common topics with adequate and context-appropriate vocabulary. |
| `C-B4-04` | Speaking continuity and basic accuracy | `S-FC-01`, `S-GRA-04`, `S-P-05` | `K-GRA-010` | Band-3 sentence/phoneme foundation | Sustain basic speech on familiar topics with clear overall meaning. |
| `C-B4-05` | Receptive main ideas and details | `L-COMP-01`, `02`, `R-COMP-01`, `02` | `K-VOC-012` | Band-3 receptive foundation | Handle gist/detail in moderately more complex input. |
| `C-B4-06` | Basic receptive question types at Band-4 demand | `L-QT-01`, `L-QT-05`, `R-QT-05` | — | `C-B4-05` | Execute basic receptive item strategies at the current demand level. |
| `C-B4-07` | Band-4 integration | Band-4 target set | — | `C-B4-01..06` recommended | Integrate Band-4 work; non-certifying orchestration checkpoint. |

## Band 5 — expanding control

| Node | Focus | Skill targets | Knowledge | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B5-01` | Compound and complex sentence foundation | `W-GRA-02`, `W-GRA-03` | `K-GRA-003`, `020`, `004`, `021` | `W-GRA-01`, core grammar | Produce compound/complex sentences with usable control. |
| `C-B5-02` | Cohesion and logical progression | `W-CC-02`, `03`, `04` | `K-GRA-022`, `033` | `W-CC-01`, `C-B5-01` recommended | Organize responses with progression, linking, reference and substitution. |
| `C-B5-03` | Task-2 position, development, relevance | `W-TR-02`, `03`, `04` | `K-VOC-012` | `W-TR-01`, `C-B5-02` recommended | Write a relevant Task-2 response with a clear position and supported ideas. |
| `C-B5-04` | Task-1 reporting with data | `W-TA-03` | `K-GRA-050`, `051` | `W-TA-01` | Report selected features with accurate supporting data and suitable tense. |
| `C-B5-05` | Paraphrase, word formation, spelling | `W-LR-03`, `W-LR-05` | `K-VOC-030`, `031` | `W-LR-01` | Paraphrase without meaning loss and control productive word form/spelling. |
| `C-B5-06` | Speaking fluency, complexity, pronunciation | `S-FC-02`, `03`, `04`, `05`, `S-GRA-02`, `S-LR-02`, `S-P-02`, `03` | `K-GRA-021`, `K-VOC-012` | `S-FC-01`, `S-P-01`, `C-B5-01` recommended | Sustain an extended coherent turn using more complex language and controlled stress/intonation. |
| `C-B5-07` | Receptive inference, paraphrase, distractors | `L-COMP-03`, `04`, `05`, `R-COMP-03`, `04`, `05` | `K-VOC-011` | lower receptive COMP leaves | Handle inference, paraphrase, structure and distractors more reliably. |
| `C-B5-08` | Higher-order receptive question types | `L-QT-02`, `03`, `04`, `R-QT-01`, `02`, `04`, `06` | — | `C-B5-07` | Apply higher-order question-type strategies under realistic constraints. |
| `C-B5-09` | Band-5 integration | Band-5 target set | — | `C-B5-01..08` recommended | Integrate Band-5 capabilities; non-certifying orchestration checkpoint. |

## Band 6 — competent clarity

| Node | Focus | Skill targets | Knowledge | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B6-01` | Grammatical accuracy and variety | `W-GRA-03`, `W-GRA-02`, `S-GRA-03` | `K-GRA-004`, `003` | Band-5 grammar | Increase accuracy and structural variety in productive language. |
| `C-B6-02` | Task-1 overview and flexible reference | `W-TA-02`, `W-CC-04` | `K-GRA-033` | `W-TA-03`, prior cohesion | Produce a clear overview and reduce repetition through controlled reference. |
| `C-B6-03` | Collocation and idiomatic resource | `W-LR-02`, `S-LR-03` | `K-VOC-020`, `021` | core/topic vocabulary | Use collocations accurately and begin appropriate less-common/idiomatic use. |
| `C-B6-04` | Extended and dense receptive content | `L-COMP-06`, `R-COMP-06` | `K-VOC-011` | Band-5 receptive analysis | Follow extended academic speech and dense academic reading with useful reliability. |
| `C-B6-05` | Writer-view classification | `R-QT-03` | — | `R-QT-02`, `R-COMP-04` | Distinguish writer agreement/disagreement/absence accurately. |
| `C-B6-06` | Band-6 integration | Band-6 target set | — | `C-B6-01..05` recommended | Integrate Band-6 accuracy/clarity; non-certifying orchestration checkpoint. |

## Band 7 — flexible good control

| Node | Focus | Skill targets | Knowledge | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B7-01` | Structural flexibility and error-free frequency | `W-GRA-07`, `W-GRA-03` | `K-GRA-005`, `006`, `007` | `W-GRA-03` | Produce varied complex structures flexibly with frequent accurate sentences. |
| `C-B7-02` | Lexical sophistication | `W-LR-04`, `S-LR-03` | `K-VOC-021`, `011` | Band-6 lexical foundation | Use less-common/idiomatic resource with increasing precision and appropriateness. |
| `C-B7-03` | Cohesion, position, development | `W-CC-03`, `04`, `W-TR-02`, `03` | `K-GRA-022`, `033` | Band-5/6 cohesion and response work | Sustain clear developed positions with flexible organization/cohesion. |
| `C-B7-04` | Receptive inference and structure | `R-COMP-04`, `05`, `L-COMP-03`, `04` | `K-VOC-011` | Band-5 receptive analysis | Handle inference, stance, paraphrase and structure reliably. |
| `C-B7-05` | Pronunciation range and flexibility | `S-P-03`, `S-P-04` | `K-PHON-030`, `040` | lower pronunciation leaves | Sustain useful intonation, connected speech and chunking across extended turns. |
| `C-B7-06` | Band-7 integration | Band-7 target set | — | `C-B7-01..05` recommended | Integrate Band-7 flexible control; non-certifying orchestration checkpoint. |

## Band 8 — very strong control

| Node | Focus | Skill targets | Knowledge | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B8-01` | Wide flexible language and near-error-free accuracy | `W-GRA-07`, `W-LR-04`, `W-GRA-01`, `W-GRA-03`, `S-GRA-03` | `K-GRA-005`, `007` | Band-7 productive control | Produce wide, flexible, highly accurate language. |
| `C-B8-02` | Skillful cohesion and fully developed response | `W-CC-02`, `03`, `04`, `W-TR-03`, `W-TA-02` | `K-GRA-022`, `033` | Band-7 writing integration | Produce well-developed responses whose organization is easy to follow. |
| `C-B8-03` | Effortless speaking fluency and wide resource | `S-FC-02`, `S-FC-04`, `S-GRA-03`, `S-LR-03`, `S-P-05` | `K-PHON-040`, `041` | Band-7 speaking/pronunciation | Sustain fluent, wide-resource speech that is easily understood. |
| `C-B8-04` | Complex receptive comprehension | `L-COMP-06`, `R-COMP-06`, `R-COMP-04` | `K-VOC-011` | Band-7 receptive control | Handle detailed, abstract and complex input with little difficulty. |
| `C-B8-05` | Band-8 integration | Band-8 target set | — | `C-B8-01..04` recommended | Integrate Band-8 high competence; non-certifying orchestration checkpoint. |

## Band 9 — ceiling

| Node | Focus | Skill targets | Knowledge | Depends | Expected outcome / node exit intent |
|---|---|---|---|---|---|
| `C-B9-01` | Full flexibility and precision | `W-GRA-07`, `W-LR-04`, `S-GRA-03`, `S-LR-03`, `S-P-05` | `K-GRA-005`, `007`, `K-VOC-021` | Band-8 productive control | Approach full flexibility, precision and sustained intelligibility. |
| `C-B9-02` | Integrated exam-level mastery | `W-TA-02`, `W-TR-03`, `W-CC-03`, `R-COMP-04`, `L-COMP-06` | — | Band-8 exit-level capability | Integrate ceiling-level capability under independent exam-like demand. |
| `C-B9-03` | Ceiling integration | full target set | — | `C-B9-01`, `C-B9-02` recommended | Final integration checkpoint; certification still requires evidence under Bands/Assessment/Progression. |

## Adaptive sequencing

The registry is a canonical recommended pathway, not a rigid learner timeline. Runtime may reorder or interleave nodes when Required prerequisites remain satisfied, target outcomes remain complete, weak canonical capabilities are not bypassed, and learner evidence supports the change.

A learner may draw from different band phases by skill because progression is per skill. Integration nodes never force synchronized four-skill certification.

## Coverage invariant

Every Curriculum Node references canonical objects. New nodes are justified only by a distinct orchestration step, not merely because another lesson/exercise is desired. Concrete lessons and exercise instances belong to `10-CONTENT-MODEL.md`.