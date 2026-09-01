# SKILL-LEAF-ONTOLOGY — Skill Leaves

> **CANONICAL PRODUCT AUTHORITY**
>
> Canonical within the PRODUCT authority domain under `docs/catalog/project.json`. `CONSTITUTION.md` and `OBJECTIVE.md` retain distinct authority, and `contracts/**` retains scoped exact machine-contract authority. References below to `spec/**` or `design/**`, including historical uses of “canonical”, are provenance only and do not create equal authority.

## Purpose and ownership

Define **what the learner must be able to demonstrate** across Listening, Reading, Writing, and Speaking.

This shard preserves the four-skill capability taxonomy, Skill Leaf identity, atomic objectives, skill/component membership, variant-specific capability overlays, and Skill-to-Skill prerequisite edges. Vocabulary, grammar, and phonology are enabling knowledge and are preserved in `docs/product/knowledge-objects.md` from canonical `spec/04-KNOWLEDGE.md`.

Official IELTS task/question families are external exam identities owned by canonical `spec/02-IELTS-MODEL.md`. A Skill Leaf may support several official families; family/content coverage must not be inferred from leaf existence alone.

The first-class IELTS skills are exactly Listening, Reading, Writing, and Speaking.

The active inventory contains **67 stable Skill Leaves**:

- Listening: 11;
- Reading: 12;
- Speaking: 18;
- Writing shared/Academic: 23;
- Writing General Training Task 1 overlay: 3.

General Training does not create a fifth skill and does not duplicate shared Reading/Writing capability. Variant-specific leaves exist only where the external construct genuinely differs.

## Atomic Skill Leaf rule

A Skill Leaf represents one independently useful capability. Decompose further only while a node still contains multiple independent objectives or cannot be independently taught, practiced, assessed, and remediated. Stop when further splitting no longer improves those properties. Do not split a capability merely to mirror every exam UI/question-family label.

## Canonical Skill Leaf fields

A Skill Leaf canonically owns only:

- stable `id`;
- human-readable name;
- IELTS skill and component;
- applicable variant when not shared;
- one atomic capability objective;
- Skill-to-Skill prerequisite edges when another capability must precede it.

The following are intentionally not identity-level Skill authority:

- official task/question family identity → canonical `spec/02-IELTS-MODEL.md`;
- band relevance or Band-N mastery quality → canonical `spec/05-BANDS.md`;
- enabling Knowledge Object resolution → `docs/product/knowledge-objects.md`, preserving canonical `spec/04-KNOWLEDGE.md`;
- practice binding → canonical `spec/07-PRACTICE.md`;
- assessment strategy/evidence → canonical `spec/08-ASSESSMENT.md`;
- learner mastery state → canonical `spec/09-PROGRESSION.md`;
- curriculum position → canonical `spec/06-CURRICULUM.md`;
- concrete family/subformat content coverage → canonical `spec/10-CONTENT-MODEL.md` plus `design/08-coverage-and-support.md`;
- common errors, examples, remediation scripts, scaffold choices → canonical `spec/10-CONTENT-MODEL.md`;
- cognitive-level and learning-load estimates → optional planning metadata unless future evidence promotes them to canonical policy.

## Listening

### Comprehension

| ID | Capability | Atomic objective |
|---|---|---|
| `L-COMP-01` | Main ideas / gist | Identify the central idea of spoken input. |
| `L-COMP-02` | Specific details | Capture targeted factual details accurately. |
| `L-COMP-03` | Purpose / attitude / opinion | Infer speaker purpose, attitude, or opinion from language and delivery. |
| `L-COMP-04` | Paraphrase & signposting | Recognize rewording, discourse signals, and answer-relevant signposts. |
| `L-COMP-05` | Distractor management | Reject plausible but unsupported or superseded options. |
| `L-COMP-06` | Extended academic speech | Follow dense, extended academic monologue without losing the discourse thread. |

### Question-type strategies

| ID | Capability | Atomic objective |
|---|---|---|
| `L-QT-01` | Form/note/table/flow-chart/summary completion | Fill structured completion gaps from the recording while respecting layout, grammar, and stated word limits. |
| `L-QT-02` | Multiple choice | Select correct single or multiple options while managing distractors. |
| `L-QT-03` | Matching | Match recorded information to speakers, categories, or options. |
| `L-QT-04` | Map/plan/diagram labelling | Apply spatial and directional language to a visual. |
| `L-QT-05` | Sentence completion & short answer | Produce concise answers from the recording within stated limits. |

### Listening prerequisite edges

| Dependent | Requires Skill Leaves |
|---|---|
| `L-COMP-03` | `L-COMP-01` |
| `L-COMP-04` | `L-COMP-01` |
| `L-COMP-05` | `L-COMP-02` |
| `L-COMP-06` | `L-COMP-01`, `L-COMP-04` |
| `L-QT-01` | `L-COMP-02` |
| `L-QT-02` | `L-COMP-01`, `L-COMP-05` |
| `L-QT-03` | `L-COMP-01` |
| `L-QT-04` | `L-COMP-02` |
| `L-QT-05` | `L-COMP-02` |

Listening decomposition separates comprehension capability from question-type execution strategy. Receptive analytic decomposition is Blueprint inference grounded in the official test format; official score/family truth remains in canonical `spec/02-IELTS-MODEL.md` and Band interpretation in canonical `spec/05-BANDS.md`.

## Reading

### Comprehension

| ID | Capability | Atomic objective |
|---|---|---|
| `R-COMP-01` | Skim for gist | Identify main idea and global topic efficiently. |
| `R-COMP-02` | Scan for detail | Locate targeted facts or references quickly. |
| `R-COMP-03` | Detailed comprehension | Understand explicit information accurately across a passage. |
| `R-COMP-04` | Inference / writer views | Infer implications, claims, stance, or meaning beyond literal wording. |
| `R-COMP-05` | Text structure / paragraph purpose | Recognize rhetorical organization and paragraph function. |
| `R-COMP-06` | Dense/abstract/complex passages | Maintain comprehension through demanding vocabulary, syntax, and argument. |

### Question-type strategies

| ID | Capability | Atomic objective |
|---|---|---|
| `R-QT-01` | Matching headings | Match paragraph meaning to headings by main idea rather than keyword coincidence. |
| `R-QT-02` | True / False / Not Given | Distinguish supported, contradicted, and absent factual information. |
| `R-QT-03` | Yes / No / Not Given | Distinguish agreement, disagreement, and absence for writer views/claims. |
| `R-QT-04` | Matching distributed information | Map information accurately across matching-information, matching-features, and matching-sentence-ending task structures. |
| `R-QT-05` | Completion & short answer | Extract and fit text evidence within grammatical, structural, and word-limit constraints. |
| `R-QT-06` | Multiple choice | Evaluate options against passage evidence, including inferential variants. |

### Reading prerequisite edges

| Dependent | Requires Skill Leaves |
|---|---|
| `R-COMP-03` | `R-COMP-01` |
| `R-COMP-04` | `R-COMP-03` |
| `R-COMP-05` | `R-COMP-01` |
| `R-COMP-06` | `R-COMP-03` |
| `R-QT-01` | `R-COMP-05` |
| `R-QT-02` | `R-COMP-03` |
| `R-QT-03` | `R-COMP-04` |
| `R-QT-04` | `R-COMP-02`, `R-COMP-03` |
| `R-QT-05` | `R-COMP-03` |
| `R-QT-06` | `R-COMP-03`, `R-COMP-04` |

`R-QT-02` and `R-QT-03` remain distinct because one targets factual information and the other writer stance.

`R-QT-04` and `R-QT-05` deliberately group transferable cognitive strategy rather than duplicating Skill Leaves for every official UI family. Canonical `spec/02-IELTS-MODEL.md` keeps those official families distinct and canonical `spec/10-CONTENT-MODEL.md` keeps their executable content coverage independently checkable.

Academic and General Training reuse these Reading leaves. Their difference is external corpus/context distribution and score conversion, not a second reading-cognition ontology. GT context-transfer requirements are owned by canonical `spec/06-CURRICULUM.md` and concrete variant/context tagging by canonical `spec/10-CONTENT-MODEL.md`.

## Speaking

### Fluency & Coherence

| ID | Capability | Atomic objective |
|---|---|---|
| `S-FC-01` | Sustained speech | Continue speaking without excessive breakdown. |
| `S-FC-02` | Long-turn production | Produce a coherent extended Part-2 response. |
| `S-FC-03` | Spoken discourse markers | Use connectives and spoken discourse markers appropriately. |
| `S-FC-04` | Coherent topic development | Extend and organize ideas coherently and relevantly. |
| `S-FC-05` | Hesitation/self-correction management | Keep hesitation and correction from disrupting communication. |

### Lexical Resource

| ID | Capability | Atomic objective |
|---|---|---|
| `S-LR-01` | Topic vocabulary | Sustain discussion with relevant vocabulary across common topics. |
| `S-LR-02` | Oral paraphrase | Circumlocute or re-express meaning when exact wording is unavailable. |
| `S-LR-03` | Less-common / idiomatic resource | Use less-common and idiomatic language naturally and appropriately. |
| `S-LR-04` | Word-choice accuracy | Select vocabulary with accurate meaning, collocation, and register. |

### Grammatical Range & Accuracy

| ID | Capability | Atomic objective |
|---|---|---|
| `S-GRA-01` | Simple/short sentence accuracy | Produce accurate simple spoken structures under real-time load. |
| `S-GRA-02` | Complex spoken sentence forms | Produce a useful mix of simple and complex spoken structures. |
| `S-GRA-03` | Structural variety & flexibility | Use a broad range of structures flexibly in extended speech. |
| `S-GRA-04` | Grammatical accuracy | Keep grammatical errors infrequent enough to preserve clear communication. |

### Pronunciation

| ID | Capability | Atomic objective |
|---|---|---|
| `S-P-01` | Phoneme accuracy | Produce individual English sounds clearly enough for intelligibility. |
| `S-P-02` | Word stress | Place lexical stress appropriately. |
| `S-P-03` | Sentence stress & intonation | Use prominence and pitch movement to convey meaning. |
| `S-P-04` | Chunking / connected speech / rhythm | Group speech into natural sense units with effective connected speech. |
| `S-P-05` | Overall intelligibility | Be understood without undue listener effort. |

### Speaking prerequisite edges

| Dependent | Requires Skill Leaves |
|---|---|
| `S-FC-02` | `S-FC-01` |
| `S-FC-04` | `S-FC-01` |
| `S-FC-05` | `S-FC-01` |
| `S-LR-02` | `S-LR-01` |
| `S-LR-03` | `S-LR-01` |
| `S-GRA-02` | `S-GRA-01` |
| `S-GRA-03` | `S-GRA-02` |
| `S-P-02` | `S-P-01` |
| `S-P-03` | `S-P-02` |
| `S-P-04` | `S-P-01` |
| `S-P-05` | `S-P-01`, `S-P-03` |

Other Speaking dependencies are enabling Knowledge Objects and are resolved in `docs/product/knowledge-objects.md`, preserving canonical `spec/04-KNOWLEDGE.md`.

## Writing — shared and Academic Task 1

### Task Achievement — Academic Task 1

| ID | Capability | Atomic objective |
|---|---|---|
| `W-TA-01` | Identify key visual features | Select the most important trends, differences, stages, changes, or comparisons in an Academic Task-1 visual. |
| `W-TA-02` | Write an Academic overview | Summarize the main patterns/features clearly and separately from supporting detail. |
| `W-TA-03` | Report features with supporting detail | Support selected visual features with accurate, relevant figures or non-numeric details as the visual requires. |

### Task Response — Task 2 shared core

| ID | Capability | Atomic objective |
|---|---|---|
| `W-TR-01` | Analyse the prompt | Identify every prompt requirement and the response task. |
| `W-TR-02` | Formulate a clear position | State a position that directly answers the question. |
| `W-TR-03` | Develop ideas with support | Extend main ideas with relevant reasons, examples, or evidence. |
| `W-TR-04` | Maintain relevance | Keep content focused on the prompt and avoid padding or repetition. |

### Coherence & Cohesion

| ID | Capability | Atomic objective |
|---|---|---|
| `W-CC-01` | Paragraphing | Group content into logical paragraphs with clear purpose. |
| `W-CC-02` | Logical progression | Sequence information and ideas into a clear overall progression. |
| `W-CC-03` | Cohesive devices | Use linking devices accurately and non-mechanically. |
| `W-CC-04` | Reference & substitution | Link ideas and reduce repetition through clear reference and substitution. |

### Lexical Resource

| ID | Capability | Atomic objective |
|---|---|---|
| `W-LR-01` | Topic-specific vocabulary | Use adequate, relevant vocabulary for the task topic. |
| `W-LR-02` | Collocation | Use natural and accurate word partnerships. |
| `W-LR-03` | Paraphrase | Re-express prompts, data, and ideas without meaning distortion. |
| `W-LR-04` | Precise / less-common / idiomatic vocabulary | Use sophisticated lexis when it improves precision and remains appropriate. |
| `W-LR-05` | Spelling & word formation | Control spelling and word-family form accurately. |

### Grammatical Range & Accuracy

| ID | Capability | Atomic objective |
|---|---|---|
| `W-GRA-01` | Simple sentence accuracy | Produce accurate simple sentence structures. |
| `W-GRA-02` | Compound sentences | Coordinate independent clauses accurately. |
| `W-GRA-03` | Complex sentences | Use subordination and complex structures effectively. |
| `W-GRA-04` | Tense & aspect | Select tense/aspect appropriate to communicative purpose. |
| `W-GRA-05` | Articles & determiners | Control article, determiner, and countability choices. |
| `W-GRA-06` | Punctuation | Use punctuation to support grammatical and discourse clarity. |
| `W-GRA-07` | Structural flexibility & variety | Use a broad structural repertoire without sacrificing accuracy. |

## Writing — General Training Task 1 overlay

General Training Task 1 is not represented by Academic visual-feature leaves. It has three dedicated leaves while reusing shared coherence, lexical, grammar, and general productive capability.

| ID | Capability | Atomic objective |
|---|---|---|
| `W-GT1-01` | Situation / recipient / purpose analysis | Identify the relationship to the recipient, communicative purpose, and all required prompt bullet points. |
| `W-GT1-02` | Letter register & relationship control | Select and sustain personal, semi-formal, or formal style appropriate to recipient and purpose, including suitable openings/closings and tone. |
| `W-GT1-03` | Purpose and required-point fulfilment | Achieve the letter's communicative purpose and cover every required bullet point with relevant, sufficiently developed information. |

These leaves are active only when the Writing Task-1 variant is General Training. They are assessed through the normal Writing criteria and do not create duplicate lexical/grammar/cohesion definitions.

### Writing prerequisite edges

| Dependent | Requires Skill Leaves |
|---|---|
| `W-TA-02` | `W-TA-01` |
| `W-TA-03` | `W-TA-01` |
| `W-TR-02` | `W-TR-01` |
| `W-TR-03` | `W-TR-02` |
| `W-TR-04` | `W-TR-01` |
| `W-CC-02` | `W-CC-01` |
| `W-CC-03` | `W-CC-02` |
| `W-CC-04` | `W-CC-02` |
| `W-LR-02` | `W-LR-01` |
| `W-LR-03` | `W-LR-01` |
| `W-LR-04` | `W-LR-01` |
| `W-GRA-02` | `W-GRA-01` |
| `W-GRA-03` | `W-GRA-01` |
| `W-GRA-07` | `W-GRA-02`, `W-GRA-03` |
| `W-GT1-02` | `W-GT1-01` |
| `W-GT1-03` | `W-GT1-01` |

GT register knowledge resolves through `K-VOC-040` in `docs/product/knowledge-objects.md`, preserving canonical `spec/04-KNOWLEDGE.md`; letter capability remains a Skill because the learner must select and use that knowledge appropriately under a communicative task.

Writing leaves with only Knowledge prerequisites are resolved in `docs/product/knowledge-objects.md`. Length and exam-format compliance are external task constraints, not Skill Leaves.

## Variant inclusion invariant

For `Academic` Writing Task 1, required Task-1 capability includes `W-TA-01`, `W-TA-02`, and `W-TA-03` plus applicable shared Writing leaves.

For `General Training` Writing Task 1, required Task-1 capability includes `W-GT1-01`, `W-GT1-02`, and `W-GT1-03` plus applicable shared Writing leaves; `W-TA-*` visual-specific leaves are not prerequisites for GT Task 1.

Task 2 uses the shared `W-TR-*` and shared Writing criteria/capabilities for both variants unless canonical `spec/02-IELTS-MODEL.md` later identifies a genuine construct difference requiring a separate leaf.

## Ownership invariants

- canonical `spec/02-IELTS-MODEL.md` owns official exam task/question-family identity;
- `docs/product/knowledge-objects.md`, preserving canonical `spec/04-KNOWLEDGE.md`, is the sole owner inside the successor PRODUCT representation of Skill→Knowledge resolution;
- canonical `spec/05-BANDS.md` is the sole owner of Band-N quality/exit thresholds, including variant-specific Task-1 overlays;
- canonical `spec/06-CURRICULUM.md` may sequence Skill Leaves but may not redefine them;
- canonical `spec/07-PRACTICE.md` may bind Practice Types to leaves but may not redefine them;
- canonical `spec/08-ASSESSMENT.md` may measure leaves but may not redefine them;
- canonical `spec/09-PROGRESSION.md` tracks learner state by leaf ID but may not redefine the target capability;
- canonical `spec/10-CONTENT-MODEL.md` tracks executable family/context/subformat content without creating parallel Skill definitions.

No downstream owner may create a parallel Skill Leaf definition.
