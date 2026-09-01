# KNOWLEDGE-OBJECT-ONTOLOGY — Knowledge Objects

> **CANONICAL PRODUCT AUTHORITY**
>
> Canonical within the PRODUCT authority domain under `docs/catalog/project.json`. `CONSTITUTION.md` and `OBJECTIVE.md` retain distinct authority, and `contracts/**` retains scoped exact machine-contract authority. References below to `spec/**` or `design/**`, including historical uses of “canonical”, are provenance only and do not create equal authority.

## Purpose and ownership

Define **what must be known** to support the capabilities preserved in `docs/product/skill-leaves.md` from historical `spec/03-SKILLS.md` provenance.

Knowledge is enabling language content, not a fifth IELTS skill. The canonical domains are Grammar, Vocabulary, and Phonology.

This shard preserves enabling language knowledge, atomic Knowledge Object identity, concise definitions, Knowledge-to-Knowledge prerequisite edges, and Skill-to-Knowledge resolution.

The active inventory contains **46 stable atomic Knowledge Objects**:

- Grammar: 29;
- Vocabulary: 9;
- Phonology: 8.

## Atomic Knowledge Object rule

A Knowledge Object represents one independently learnable concept. Decompose while a concept still contains multiple independently learnable units. Stop when further splitting no longer improves learning, prerequisite reasoning, remediation, or reuse.

## Canonical Knowledge Object fields

A Knowledge Object canonically owns:

- stable `id`;
- name;
- domain;
- concise definition;
- explicit `requires` edges to prerequisite Knowledge Objects.

The following are intentionally not canonical object fields unless future evidence promotes them:

- examples and worked instances → historical `spec/10-CONTENT-MODEL.md` provenance and current DATA/content authority;
- misconception/remediation catalogs → historical `spec/10-CONTENT-MODEL.md` provenance and current DATA/content authority;
- soft `related_to` associations with no gating/selection consequence → historical/supporting graph only;
- band relevance → current IELTS/Bands and Curriculum PRODUCT shards;
- learner acquisition state → current DATA/BEHAVIOR authority.

## Grammar graph

| ID | Object | Canonical definition | Requires |
|---|---|---|---|
| `K-GRA-010` | Word classes / parts of speech | Noun, verb, adjective, adverb, preposition, conjunction, determiner, pronoun and related functional categories. | — |
| `K-GRA-061` | Noun countability & pluralization | Countable/uncountable noun behavior, regular/irregular plurals, and number marking. | `K-GRA-010` |
| `K-GRA-001` | Clause structure | Core clause elements such as subject, verb, object, complement, adjunct and their basic ordering. | `K-GRA-010` |
| `K-GRA-002` | Simple sentence | A complete sentence built from one independent clause. | `K-GRA-001` |
| `K-GRA-003` | Compound sentence / coordination | Two or more independent clauses joined by coordination. | `K-GRA-002`, `K-GRA-020` |
| `K-GRA-004` | Complex sentence / subordination | An independent clause combined with one or more subordinate clauses. | `K-GRA-002`, `K-GRA-021` |
| `K-GRA-005` | Relative clauses | Clauses that modify a noun using relative forms, including defining and non-defining patterns. | `K-GRA-004`, `K-GRA-031` |
| `K-GRA-006` | Conditional clauses | Conditional relationships expressed through if/conditional clause patterns, including common mixed forms. | `K-GRA-004` |
| `K-GRA-007` | Noun clauses | Subordinate clauses functioning as noun-like constituents. | `K-GRA-004` |
| `K-GRA-020` | Coordinating conjunctions | Coordinators such as and/but/or/so that join equal grammatical units. | `K-GRA-010` |
| `K-GRA-021` | Subordinating conjunctions | Subordinators such as although/because/if/while that introduce dependent clauses. | `K-GRA-010` |
| `K-GRA-022` | Linking adverbials / conjuncts | Sentence/discourse transitions such as however, therefore, moreover and consequently. | `K-GRA-010` |
| `K-GRA-030` | Pronouns | Personal, possessive, demonstrative, reflexive and related pronoun systems. | `K-GRA-010` |
| `K-GRA-031` | Relative pronouns | Relative forms such as who, whom, which, that and whose. | `K-GRA-030` |
| `K-GRA-032` | Determiners | Articles, demonstratives, quantifiers, possessives and other noun-phrase specifiers. | `K-GRA-010` |
| `K-GRA-033` | Reference & substitution for cohesion | Reference and substitution devices used to link discourse and avoid unnecessary repetition. | `K-GRA-030`, `K-GRA-032` |
| `K-GRA-040` | Articles: a / an / the | Core indefinite and definite article forms and basic use. | `K-GRA-032` |
| `K-GRA-041` | Definite / indefinite / zero article rules | Conditions governing article choice, including zero article and generic/specific reference. | `K-GRA-040` |
| `K-GRA-050` | Present simple | Present-simple form and its core uses for facts, habits, states and generalizations. | `K-GRA-002` |
| `K-GRA-051` | Past simple | Past-simple form and its core uses for completed past events/states. | `K-GRA-002` |
| `K-GRA-052` | Present perfect | Present-perfect form connecting prior events/states to present relevance. | `K-GRA-051` |
| `K-GRA-053` | Past perfect | Past-perfect form for an earlier past event/state relative to another past reference point. | `K-GRA-052` |
| `K-GRA-054` | Progressive / continuous aspect | `be + -ing` aspect for ongoing or temporally bounded situations across time references. | `K-GRA-050`, `K-GRA-051` |
| `K-GRA-055` | Future forms | Common future-reference forms including will, going to and present-continuous patterns. | `K-GRA-002` |
| `K-GRA-060` | Subject–verb agreement | Agreement of finite verb form with subject person/number, including common complex cases. | `K-GRA-002`, `K-GRA-061` |
| `K-GRA-062` | Modal verbs | Modal auxiliaries such as can, could, must, should, may, might and would and their core meanings. | `K-GRA-002` |
| `K-GRA-063` | Passive voice | Passive construction using an appropriate form of `be` plus past participle to background the agent/focus the affected entity. | `K-GRA-002`, `K-GRA-061` |
| `K-GRA-064` | Comparatives & superlatives | Comparative and superlative morphology/syntax for adjectives and adverbs. | `K-GRA-010` |
| `K-GRA-065` | Negation & question forms | Core negation plus yes/no and wh-question formation, including auxiliary behavior and word order. | `K-GRA-002` |

## Vocabulary graph

| ID | Object | Canonical definition | Requires |
|---|---|---|---|
| `K-VOC-010` | High-frequency / core vocabulary | High-frequency general English word families required for basic comprehension and production. | — |
| `K-VOC-011` | Academic vocabulary | Formal/academic lexical resource common in academic reading and writing. | `K-VOC-010` |
| `K-VOC-012` | Topic-specific word sets | Reusable lexical sets organized around common IELTS-relevant topics. | `K-VOC-010` |
| `K-VOC-020` | Collocations | Conventional recurrent word partnerships and their usage constraints. | `K-VOC-010` |
| `K-VOC-021` | Idioms & fixed expressions | Fixed or semi-fixed multiword expressions whose meaning/use is conventionalized. | `K-VOC-010` |
| `K-VOC-030` | Word formation / affixation | Productive relationships among word-family forms through prefixes, suffixes and derivation. | `K-VOC-010` |
| `K-VOC-031` | Spelling rules & patterns | English orthographic conventions and recurring spelling patterns needed for accurate written production. | `K-VOC-010` |
| `K-VOC-040` | Register & formality | Lexical choice according to formal, neutral, informal and context-appropriate register. | `K-VOC-010` |
| `K-VOC-041` | Synonymy & paraphrase resources | Lexical alternatives, connotation and usage constraints used to re-express meaning accurately. | `K-VOC-010` |

`K-VOC-012` represents the topic-set system. Individual topic lexicons are concrete content under the current DATA/content authority, not new canonical domains.

Learner-saved words, phrases, collocations, examples, and personal review cards are also concrete content/review instances. They may reference one or more canonical Knowledge Objects, but saving an item never creates a new `K-*` identity.

## Phonology graph

| ID | Object | Canonical definition | Requires |
|---|---|---|---|
| `K-PHON-010` | Consonant phonemes | English consonant sound inventory and articulatory distinctions relevant to intelligible production/perception. | — |
| `K-PHON-011` | Vowel phonemes | English monophthong/diphthong inventory and contrastive vowel distinctions. | — |
| `K-PHON-012` | Phoneme contrasts & minimal pairs | Contrastive sound pairs used to perceive and produce meaning-distinguishing phonemes. | `K-PHON-010`, `K-PHON-011` |
| `K-PHON-020` | Word stress | Placement of primary lexical stress and common stress shifts across word families. | `K-PHON-011` |
| `K-PHON-021` | Sentence stress | Prominence patterns across utterances, including content/function-word stress behavior. | `K-PHON-020` |
| `K-PHON-030` | Intonation patterns | Pitch contours and their discourse, pragmatic, or attitudinal functions. | `K-PHON-021` |
| `K-PHON-040` | Connected speech | Linking, reduction, assimilation, elision and weak-form behavior across word boundaries. | `K-PHON-010`, `K-PHON-011` |
| `K-PHON-041` | Rhythm & chunking | Grouping speech into sense units and maintaining intelligible rhythmic organization. | `K-PHON-021`, `K-PHON-040` |

The phonology graph is L1-agnostic. L1-specific contrast priorities belong to localized/remediation content, not this graph.

## Skill→Knowledge resolution

This is the sole PRODUCT authority for the mapping from Skill Leaf knowledge needs to Knowledge Objects. Historical `spec/04-KNOWLEDGE.md` remains provenance only.

An edge in this mapping means the Knowledge Object is a **universal intrinsic prerequisite for the Skill Leaf across the leaf's applicable canonical scope**. It is therefore eligible to become a Required prerequisite downstream. Knowledge that is useful only for a particular variant, topic, presentation class, prompt, stimulus, or content instance does **not** belong in this universal mapping; Curriculum/Content may select it explicitly where material.

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
| `W-GRA-07` | `K-GRA-005`, `K-GRA-006`, `K-GRA-007`, `K-GRA-062`, `K-GRA-063` |
| `W-GT1-02` | `K-VOC-040` |

`W-GT1-01` and `W-GT1-03` are task-capability leaves whose direct knowledge needs are covered by the shared lexical/grammar/cohesion targets selected for the concrete letter; they do not require a new letter-specific Knowledge Object.

`W-TA-03` has no universal single Knowledge prerequisite beyond shared Writing knowledge. Comparatives (`K-GRA-064`), tense/aspect, data language, process language, and spatial language are selected when the concrete Academic Task-1 presentation/stimulus makes them material; a process/map task must not be hard-gated by comparatives merely because another Task-1 presentation uses them heavily.

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

No Reading Skill Leaf currently has an additional **universal variant-independent** Knowledge prerequisite in this mapping.

In particular, shared `R-COMP-06` does not universally require `K-VOC-011` Academic vocabulary: Academic and General Training reuse the leaf, while their corpus/context vocabulary differs. The Academic route may explicitly schedule `K-VOC-011`; GT uses the applicable general/topic/workplace lexical content and context distribution without inheriting a hidden Academic-vocabulary hard gate.

General Training Reading context diversity is a curriculum/content/evidence condition, not a separate Knowledge Object.

Listening currently has no explicit Knowledge prerequisite edge in the frozen graph; its direct prerequisites are intra-skill capability edges in `docs/product/skill-leaves.md`.

## Dependency semantics

This shard preserves universal Knowledge prerequisite edges. The Curriculum PRODUCT authority may add explicit context/variant learning targets and Recommended node ordering, but it may not reinterpret a context-specific useful Knowledge Object as an intrinsic Required edge for every instance of a shared Skill Leaf.

Required / Recommended / Independent gate semantics are consumed by the Curriculum PRODUCT authority and enforced by BEHAVIOR/DATA progression semantics.

Object identity and prerequisite semantics are canonical. Difficulty, time-to-learn, frequency, examples, misconceptions, localization priority, context-specific lexical sets, and remediation catalogs are empirical/content concerns unless future evidence explicitly promotes them.
