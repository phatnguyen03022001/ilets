STATUS: CANONICAL
OWNS: enabling language knowledge, atomic Knowledge Object identity, knowledge-to-knowledge dependency graph, and Skill-to-Knowledge resolution
DEPENDS_ON: 03-SKILLS.md
DOES_NOT_OWN: IELTS skill capability definitions, band proficiency thresholds, curriculum sequence, practice activities, assessment sufficiency, learner knowledge state

# 04 — Knowledge

## Purpose

Define **what must be known** to support the capabilities in `03-SKILLS.md`.

Knowledge is enabling language content, not a fifth IELTS skill. The canonical knowledge domains are:

- Grammar;
- Vocabulary;
- Phonology.

The active inventory contains **46 stable atomic Knowledge Objects**:

- Grammar: 29;
- Vocabulary: 9;
- Phonology: 8.

## Atomic Knowledge Object rule

A Knowledge Object represents one independently learnable concept.

Decompose while a concept still contains multiple independently learnable units. Stop when further splitting would no longer improve learning, prerequisite reasoning, remediation, or reuse.

## Knowledge Object semantic shape

A canonical Knowledge Object owns:

- stable `id`;
- name;
- knowledge domain;
- concise definition;
- explicit `requires` edges to prerequisite Knowledge Objects;
- optional `related_to` edges for non-prerequisite relationships;
- examples where useful;
- common misconceptions.

It does not own learner acquisition state or the curriculum position where the knowledge is taught.

## Grammar graph

| ID | Knowledge Object | Requires |
|---|---|---|
| `K-GRA-010` | Word classes / parts of speech | — |
| `K-GRA-061` | Noun countability & pluralization | `K-GRA-010` |
| `K-GRA-001` | Clause structure | `K-GRA-010` |
| `K-GRA-002` | Simple sentence | `K-GRA-001` |
| `K-GRA-003` | Compound sentence / coordination | `K-GRA-002`, `K-GRA-020` |
| `K-GRA-004` | Complex sentence / subordination | `K-GRA-002`, `K-GRA-021` |
| `K-GRA-005` | Relative clauses | `K-GRA-004`, `K-GRA-031` |
| `K-GRA-006` | Conditional clauses | `K-GRA-004` |
| `K-GRA-007` | Noun clauses | `K-GRA-004` |
| `K-GRA-020` | Coordinating conjunctions | `K-GRA-010` |
| `K-GRA-021` | Subordinating conjunctions | `K-GRA-010` |
| `K-GRA-022` | Linking adverbials / conjuncts | `K-GRA-010` |
| `K-GRA-030` | Pronouns | `K-GRA-010` |
| `K-GRA-031` | Relative pronouns | `K-GRA-030` |
| `K-GRA-032` | Determiners | `K-GRA-010` |
| `K-GRA-033` | Reference & substitution for cohesion | `K-GRA-030`, `K-GRA-032` |
| `K-GRA-040` | Articles: a / an / the | `K-GRA-032` |
| `K-GRA-041` | Definite / indefinite / zero article rules | `K-GRA-040` |
| `K-GRA-050` | Present simple | `K-GRA-002` |
| `K-GRA-051` | Past simple | `K-GRA-002` |
| `K-GRA-052` | Present perfect | `K-GRA-051` |
| `K-GRA-053` | Past perfect | `K-GRA-052` |
| `K-GRA-054` | Progressive / continuous aspect | `K-GRA-050`, `K-GRA-051` |
| `K-GRA-055` | Future forms | `K-GRA-002` |
| `K-GRA-060` | Subject–verb agreement | `K-GRA-002`, `K-GRA-061` |
| `K-GRA-062` | Modal verbs | `K-GRA-002` |
| `K-GRA-063` | Passive voice | `K-GRA-002`, `K-GRA-061` |
| `K-GRA-064` | Comparatives & superlatives | `K-GRA-010` |
| `K-GRA-065` | Negation & question forms | `K-GRA-002` |

### Grammar semantic clusters

**Foundation.** Word classes, clause structure, sentence cores, noun countability, and agreement establish the minimum grammar needed for productive capability.

**Complexity.** Coordination, subordination, relative/conditional/noun clauses, and modal/passive structures provide the grammar resource needed for higher-band range and flexibility.

**Cohesion.** Conjunctions, conjuncts, pronouns, determiners, reference, and substitution support discourse relationships without becoming a Writing/Speaking capability definition themselves.

**Time and comparison.** Tense/aspect, future forms, and comparison provide content needed to describe events, data, arguments, and relationships accurately.

## Vocabulary graph

| ID | Knowledge Object | Requires |
|---|---|---|
| `K-VOC-010` | High-frequency / core vocabulary | — |
| `K-VOC-011` | Academic vocabulary | `K-VOC-010` |
| `K-VOC-012` | Topic-specific word sets | `K-VOC-010` |
| `K-VOC-020` | Collocations | `K-VOC-010` |
| `K-VOC-021` | Idioms & fixed expressions | `K-VOC-010` |
| `K-VOC-030` | Word formation / affixation | `K-VOC-010` |
| `K-VOC-031` | Spelling rules & patterns | `K-VOC-010` |
| `K-VOC-040` | Register & formality | `K-VOC-010` |
| `K-VOC-041` | Synonymy & paraphrase resources | `K-VOC-010` |

Vocabulary knowledge must be taught as usable meaning, form, collocation, register, and retrieval rather than as isolated memorized lists.

`K-VOC-012` is a canonical object representing the topic-set system. Individual topic lexicons may later be concrete content instances under `10-CONTENT-MODEL.md`; they do not require a new canonical domain.

## Phonology graph

| ID | Knowledge Object | Requires |
|---|---|---|
| `K-PHON-010` | Consonant phonemes | — |
| `K-PHON-011` | Vowel phonemes | — |
| `K-PHON-012` | Phoneme contrasts & minimal pairs | `K-PHON-010`, `K-PHON-011` |
| `K-PHON-020` | Word stress | `K-PHON-011` |
| `K-PHON-021` | Sentence stress | `K-PHON-020` |
| `K-PHON-030` | Intonation patterns | `K-PHON-021` |
| `K-PHON-040` | Connected speech | `K-PHON-010`, `K-PHON-011` |
| `K-PHON-041` | Rhythm & chunking | `K-PHON-021`, `K-PHON-040` |

The canonical phonology graph is L1-agnostic. L1-specific contrast priorities may be supplied as localization/remediation overlays without redefining these objects.

## Skill-to-Knowledge resolution

The following table is the canonical cross-domain resolution from Skill Leaf needs to Knowledge Objects. Skill definitions remain owned by `03-SKILLS.md`; this file owns which knowledge satisfies those needs.

### Writing

| Skill Leaf | Resolved Knowledge Objects |
|---|---|
| `W-CC-03` | `K-GRA-020`, `K-GRA-021`, `K-GRA-022` |
| `W-CC-04` | `K-GRA-030`, `K-GRA-032`, `K-GRA-033` |
| `W-LR-01` | `K-VOC-012` |
| `W-LR-02` | `K-VOC-020` |
| `W-LR-03` | `K-VOC-041`, `K-VOC-020` |
| `W-LR-05` | `K-VOC-030`, `K-VOC-031` |
| `W-GRA-01` | `K-GRA-001`, `K-GRA-002` |
| `W-GRA-02` | `K-GRA-003`, `K-GRA-020` |
| `W-GRA-03` | `K-GRA-004`, `K-GRA-021`, `K-GRA-005`, `K-GRA-006`, `K-GRA-007`, `K-GRA-031` |
| `W-GRA-04` | `K-GRA-050`, `K-GRA-051`, `K-GRA-052`, `K-GRA-053`, `K-GRA-054`, `K-GRA-055` |
| `W-GRA-05` | `K-GRA-032`, `K-GRA-040`, `K-GRA-041` |
| `W-TA-03` | `K-GRA-064` |
| `W-GRA-07` | `K-GRA-005`, `K-GRA-006`, `K-GRA-007`, `K-GRA-062`, `K-GRA-063` |

### Speaking

| Skill Leaf | Resolved Knowledge Objects |
|---|---|
| `S-FC-03` | `K-GRA-020`, `K-GRA-021`, `K-GRA-022` |
| `S-LR-01` | `K-VOC-012` |
| `S-LR-04` | `K-VOC-040`, `K-VOC-020` |
| `S-GRA-01` | `K-GRA-001`, `K-GRA-002` |
| `S-GRA-02` | `K-GRA-004`, `K-GRA-021`, `K-GRA-005` |
| `S-GRA-04` | `K-GRA-001`, `K-GRA-002`, `K-GRA-060`, `K-GRA-065` |
| `S-P-01` | `K-PHON-010`, `K-PHON-011`, `K-PHON-012` |
| `S-P-02` | `K-PHON-020` |
| `S-P-03` | `K-PHON-021`, `K-PHON-030` |

### Reading

| Skill Leaf | Resolved Knowledge Objects |
|---|---|
| `R-COMP-06` | `K-VOC-011` |

Listening currently has no explicit `K-*` prerequisite edge in the canonical Skill Graph; its direct prerequisites are primarily intra-skill comprehension dependencies. Future evidence may justify explicit knowledge edges, but they must be added here rather than embedded as a parallel mapping elsewhere.

## Prerequisite classification

Knowledge dependencies used as learning gates are classified by the Curriculum/Progression model as:

- **Required** — missing knowledge makes dependent learning ineffective and no reasonable adaptive workaround exists;
- **Recommended** — beneficial but not blocking;
- **Independent** — no prerequisite relation.

This file defines graph edges and resolutions. `06-CURRICULUM.md` owns when those dependencies affect sequence; `09-PROGRESSION.md` owns runtime gating behavior.

## Open calibration boundary

Object identity and dependency semantics are canonical. Empirical difficulty, time-to-learn, localization priority, and frequency estimates may be calibrated later without redefining the object graph unless the evidence shows the graph itself is wrong.
