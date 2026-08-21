STATUS: CANONICAL
OWNS: the four-skill capability taxonomy, Skill Leaf identity, atomic decomposition rules, intrinsic skill objectives, and skill-to-skill prerequisite semantics
DEPENDS_ON: 02-IELTS-MODEL.md
DOES_NOT_OWN: knowledge object definitions, band-specific mastery thresholds, practice types, assessment strategies, learner mastery state, curriculum ordering

# 03 — Skills

## Purpose

Define **what the learner must be able to demonstrate** across Listening, Reading, Writing, and Speaking.

Skills are capabilities. Vocabulary, grammar, and phonology are enabling knowledge and are owned by `04-KNOWLEDGE.md`.

## Canonical skill set

The first-class IELTS skills are exactly:

```text
Listening
Reading
Writing
Speaking
```

The active inventory contains **64 stable Skill Leaves**:

- Listening: 11;
- Reading: 12;
- Speaking: 18;
- Writing: 23.

The structural refactor preserves these IDs so curriculum, knowledge resolution, practice, assessment, and historical evidence remain traceable.

## Atomic Skill Leaf rule

A Skill Leaf represents one independently useful learning capability.

Decompose a capability further only when the current node still contains multiple independent objectives or cannot be independently:

1. taught;
2. practiced;
3. assessed;
4. remediated.

Stop decomposing when further splitting no longer improves those four properties.

Do not target a preferred number of leaves.

## Skill Leaf semantic shape

A canonical Skill Leaf owns intrinsic capability semantics such as:

- stable `id`;
- human-readable name;
- IELTS skill and component;
- one learning objective;
- relevant cognitive operation;
- relative learning load where useful for planning;
- trace to the external IELTS requirement that motivates the capability;
- skill-to-skill prerequisites;
- a statement of the enabling knowledge it requires;
- common capability errors;
- remediation direction;
- atomicity attestation.

A Skill Leaf does **not** own:

- the Band-N proficiency threshold for that skill;
- learner mastery state;
- accumulated learner evidence;
- canonical practice type definitions;
- canonical assessment type definitions;
- curriculum position.

Those semantics belong to their respective owners.

## Listening

Listening decomposition intentionally separates general comprehension capabilities from question-type execution strategies. Because receptive analytic abilities are not fully published by IELTS as per-band descriptors, the pedagogical decomposition is Blueprint inference grounded in the official test format.

### Comprehension

| ID | Capability | Core objective |
|---|---|---|
| `L-COMP-01` | Main ideas / gist | Identify the central idea of spoken input. |
| `L-COMP-02` | Specific details | Capture targeted factual details accurately. |
| `L-COMP-03` | Purpose / attitude / opinion | Infer speaker purpose, attitude, or opinion from language and delivery. |
| `L-COMP-04` | Paraphrase & signposting | Recognize rewording, discourse signals, and answer-relevant signposts. |
| `L-COMP-05` | Distractor management | Reject plausible but unsupported or superseded options. |
| `L-COMP-06` | Extended academic speech | Follow dense, extended academic monologue without losing the discourse thread. |

Key skill prerequisites:

- `L-COMP-03` builds on gist comprehension;
- `L-COMP-04` builds on gist comprehension;
- `L-COMP-05` builds on detail comprehension;
- `L-COMP-06` builds on gist and paraphrase/signposting.

### Question-type strategies

| ID | Capability | Core objective |
|---|---|---|
| `L-QT-01` | Form/note/table/flow-chart completion | Fill gaps from the recording while respecting form and word limits. |
| `L-QT-02` | Multiple choice | Select correct single or multiple options while managing distractors. |
| `L-QT-03` | Matching | Match recorded information to speakers, categories, or options. |
| `L-QT-04` | Map/plan/diagram labelling | Apply spatial and directional language to a visual. |
| `L-QT-05` | Sentence completion & short answer | Produce concise answers from the recording within stated limits. |

Question-type leaves depend on the relevant comprehension leaves rather than replacing them.

## Reading

Reading likewise separates comprehension capability from question-type strategy.

### Comprehension

| ID | Capability | Core objective |
|---|---|---|
| `R-COMP-01` | Skim for gist | Identify main idea and global topic efficiently. |
| `R-COMP-02` | Scan for detail | Locate targeted facts or references quickly. |
| `R-COMP-03` | Detailed comprehension | Understand explicit information accurately across a passage. |
| `R-COMP-04` | Inference / writer views | Infer implications, claims, stance, or meaning beyond literal wording. |
| `R-COMP-05` | Text structure / paragraph purpose | Recognize rhetorical organization and paragraph function. |
| `R-COMP-06` | Dense/abstract/complex passages | Maintain comprehension through demanding academic vocabulary, syntax, and argument. |

Key skill prerequisites:

- detailed comprehension develops from gist-level reading;
- inference depends on detailed comprehension;
- structural analysis depends on gist comprehension;
- dense-text capability depends on detailed comprehension plus enabling academic vocabulary owned by `04-KNOWLEDGE.md`.

### Question-type strategies

| ID | Capability | Core objective |
|---|---|---|
| `R-QT-01` | Matching headings | Match paragraph meaning to headings based on main idea rather than keyword coincidence. |
| `R-QT-02` | True / False / Not Given | Distinguish supported, contradicted, and absent factual information. |
| `R-QT-03` | Yes / No / Not Given | Distinguish agreement, disagreement, and absence for writer views/claims. |
| `R-QT-04` | Matching information/features/endings | Map distributed information across the passage accurately. |
| `R-QT-05` | Completion & short answer | Extract and fit text evidence within grammatical and word-limit constraints. |
| `R-QT-06` | Multiple choice | Evaluate options against passage evidence, including inferential variants. |

`R-QT-02` and `R-QT-03` remain distinct because one targets factual information and the other writer stance.

## Speaking

Speaking capability is organized around the official criterion families while remaining task-part independent except where a capability is inherently tied to the long turn.

### Fluency & Coherence

| ID | Capability | Core objective |
|---|---|---|
| `S-FC-01` | Sustained speech | Continue speaking without excessive breakdown. |
| `S-FC-02` | Long-turn production | Produce a coherent extended Part-2 response. |
| `S-FC-03` | Spoken discourse markers | Use connectives and spoken discourse markers appropriately. |
| `S-FC-04` | Coherent topic development | Extend and organize ideas coherently and relevantly. |
| `S-FC-05` | Hesitation/self-correction management | Keep hesitation and correction from disrupting communication. |

### Lexical Resource

| ID | Capability | Core objective |
|---|---|---|
| `S-LR-01` | Topic vocabulary | Sustain discussion with relevant vocabulary across common topics. |
| `S-LR-02` | Oral paraphrase | Circumlocute or re-express meaning when exact wording is unavailable. |
| `S-LR-03` | Less-common / idiomatic resource | Use less-common and idiomatic language naturally and appropriately. |
| `S-LR-04` | Word-choice accuracy | Select vocabulary with accurate meaning, collocation, and register. |

### Grammatical Range & Accuracy

| ID | Capability | Core objective |
|---|---|---|
| `S-GRA-01` | Simple/short sentence accuracy | Produce accurate simple spoken structures under real-time load. |
| `S-GRA-02` | Complex spoken sentence forms | Produce a useful mix of simple and complex spoken structures. |
| `S-GRA-03` | Structural variety & flexibility | Use a broad range of structures flexibly in extended speech. |
| `S-GRA-04` | Grammatical accuracy | Keep grammatical errors infrequent enough to preserve clear communication. |

### Pronunciation

| ID | Capability | Core objective |
|---|---|---|
| `S-P-01` | Phoneme accuracy | Produce individual English sounds clearly enough for intelligibility. |
| `S-P-02` | Word stress | Place lexical stress appropriately. |
| `S-P-03` | Sentence stress & intonation | Use prominence and pitch movement to convey meaning. |
| `S-P-04` | Chunking / connected speech / rhythm | Group speech into natural sense units with effective connected speech. |
| `S-P-05` | Overall intelligibility | Be understood without undue listener effort. |

Speaking-part coverage is trained and assessed by the Practice and Assessment owners. The skill taxonomy itself describes the underlying capability.

## Writing

Writing capability is decomposed by Task-1 achievement, Task-2 response, and the cross-task analytic criteria.

### Task Achievement — Academic Task 1

| ID | Capability | Core objective |
|---|---|---|
| `W-TA-01` | Identify key features | Select the most important trends, differences, stages, or comparisons in a visual. |
| `W-TA-02` | Write an overview | Summarize the main patterns clearly and separately from supporting detail. |
| `W-TA-03` | Report features with data | Support selected features with accurate, relevant figures or details. |

### Task Response — Task 2

| ID | Capability | Core objective |
|---|---|---|
| `W-TR-01` | Analyse the prompt | Identify every prompt requirement and the response task. |
| `W-TR-02` | Formulate a clear position | State a position that directly answers the question. |
| `W-TR-03` | Develop ideas with support | Extend main ideas with relevant reasons, examples, or evidence. |
| `W-TR-04` | Maintain relevance | Keep content focused on the prompt and avoid padding or repetition. |

### Coherence & Cohesion

| ID | Capability | Core objective |
|---|---|---|
| `W-CC-01` | Paragraphing | Group content into logical paragraphs with clear purpose. |
| `W-CC-02` | Logical progression | Sequence information and ideas into a clear overall progression. |
| `W-CC-03` | Cohesive devices | Use linking devices accurately and non-mechanically. |
| `W-CC-04` | Reference & substitution | Link ideas and reduce repetition through clear reference and substitution. |

### Lexical Resource

| ID | Capability | Core objective |
|---|---|---|
| `W-LR-01` | Topic-specific vocabulary | Use adequate, relevant vocabulary for the task topic. |
| `W-LR-02` | Collocation | Use natural and accurate word partnerships. |
| `W-LR-03` | Paraphrase | Re-express prompts, data, and ideas without meaning distortion. |
| `W-LR-04` | Precise / less-common / idiomatic vocabulary | Use sophisticated lexis when it improves precision and remains appropriate. |
| `W-LR-05` | Spelling & word formation | Control spelling and word-family form accurately. |

### Grammatical Range & Accuracy

| ID | Capability | Core objective |
|---|---|---|
| `W-GRA-01` | Simple sentence accuracy | Produce accurate simple sentence structures. |
| `W-GRA-02` | Compound sentences | Coordinate independent clauses accurately. |
| `W-GRA-03` | Complex sentences | Use subordination and complex structures effectively. |
| `W-GRA-04` | Tense & aspect | Select tense/aspect appropriate to communicative purpose. |
| `W-GRA-05` | Articles & determiners | Control article, determiner, and countability choices. |
| `W-GRA-06` | Punctuation | Use punctuation to support grammatical and discourse clarity. |
| `W-GRA-07` | Structural flexibility & variety | Use a broad structural repertoire without sacrificing accuracy. |

Length and exam-format compliance are task constraints from the IELTS model, not separate Skill Leaves.

## Skill-to-knowledge relationship

Skill Leaves may require enabling knowledge. The exact canonical Knowledge Object mapping is owned by `04-KNOWLEDGE.md`.

Examples of the relationship, without redefining the knowledge objects:

- writing grammar capabilities require grammar knowledge;
- Writing/Speaking lexical capabilities require vocabulary knowledge;
- Speaking pronunciation capabilities require phonology knowledge;
- dense Academic Reading requires academic vocabulary.

## Band relationship

This file intentionally does not attach a canonical Band-N mastery definition to each leaf.

`05-BANDS.md` owns the proficiency thresholds and exit criteria. A Skill Leaf can be relevant across multiple bands while the expected quality of its performance rises with the band.

## Downstream consumers

- `04-KNOWLEDGE.md` resolves enabling knowledge dependencies;
- `05-BANDS.md` defines required quality at each band;
- `06-CURRICULUM.md` sequences leaves;
- `07-PRACTICE.md` targets leaves through Practice Types;
- `08-ASSESSMENT.md` measures leaves;
- `09-PROGRESSION.md` tracks learner state by Skill Leaf ID.

None of those consumers may create a parallel definition of a Skill Leaf.
