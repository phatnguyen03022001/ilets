STATUS: CANONICAL
OWNS: user-facing feature capabilities for Listening, Reading, Writing, Speaking, and shared learning surfaces
DEPENDS_ON: ../spec/02-IELTS-MODEL.md, ../spec/03-SKILLS.md, ../spec/05-BANDS.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, 00-learning-experience.md
DOES_NOT_OWN: skill/band truth, practice mechanism definitions, mastery/evidence rules, API wire shape, media rights policy, or implementation frameworks

# Skill Features

## Purpose

Define the initial product feature surface for each IELTS skill. Features are learner-facing capabilities; they consume canonical Skill, Practice, Assessment, and Progression semantics.

The initial product retains **40 named feature capabilities**:

```text
Listening  8
Reading    8
Writing    9
Speaking   9
Shared     6
---------
Total     40
```

Variant completeness is represented as feature submodes/conditions where the interaction model is genuinely shared. Do not inflate feature count merely to mirror Academic/General Training labels.

# Delivery interaction baseline

Exam-readiness features are **computer-first** because IELTS announced computer delivery as the standard rollout from mid-2026, with market timing variation.

A selected-market Writing-on-Paper option may be supported as an input-mode overlay. It does not create separate Skill, Band, Practice, or Assessment truth.

Features must therefore distinguish:

- learning interaction that is delivery-agnostic;
- computer exam-readiness behavior such as on-screen navigation/input;
- optional handwriting practice when the learner's real booked delivery option makes it relevant.

# Listening — 8 features

| ID | Feature | Learner experience | Primary output |
|---|---|---|---|
| `L-F01` | IELTS Section Player | play official-style/owned section material with computer-style question navigation and timed mode | section attempt |
| `L-F02` | Dictation Lab | listen to short segments and type exactly what was heard; reveal progressively after attempt | lexical/detail/spelling observations |
| `L-F03` | Gist & Main-Idea Drill | listen once, choose or produce the central meaning before detail review | gist observation |
| `L-F04` | Detail & Completion Drill | form/note/table/sentence/short-answer completion with word-limit validation | detail + form-control observations |
| `L-F05` | Paraphrase & Distractor Lab | compare prompt/audio phrasing, identify distractor logic, justify final choice | paraphrase/distractor observations |
| `L-F06` | Map / Diagram Lab | spatial-language listening with plan/map/diagram labeling | spatial/detail observations |
| `L-F07` | Transcript & Error Review | after the independent attempt, inspect transcript/evidence and classify why an answer failed | ErrorPattern + remediation links |
| `L-F08` | Media Listening Lab | use eligible YouTube/owned media for dictation, gist, retell, detail, or vocabulary practice | practice attempt; non-certifying by default |

Listening flow:

```text
pre-task objective
→ independent listen/answer
→ result
→ transcript/evidence reveal when allowed
→ error classification
→ targeted replay/remediation
→ optional fresh transfer item
```

The product must not require transcript visibility before the first independent listening attempt.

# Reading — 8 features

| ID | Feature | Learner experience | Primary output |
|---|---|---|---|
| `R-F01` | Reading Workspace | variant-aware IELTS Reading workspace with question navigation, highlighting/notes, timed or untimed mode | passage/section attempt |
| `R-F02` | Skim Sprint | locate topic, purpose, and paragraph-level main ideas under a short time budget | gist/structure observation |
| `R-F03` | Scan & Detail Hunt | locate explicit details, names, dates, numbers, definitions, and evidence spans | detail-location observation |
| `R-F04` | T/F/NG + Y/N/NG Lab | distinguish contradiction, agreement, and absence of evidence | stance/evidence classification |
| `R-F05` | Headings & Structure Lab | match headings, paragraph purposes, and text structure | structure observation |
| `R-F06` | Paraphrase / Inference / Stance Lab | map lexical paraphrase and infer writer meaning/position without keyword matching | inference/paraphrase observation |
| `R-F07` | Evidence Review | reveal and compare the exact text evidence supporting/invalidating the learner response | ErrorPattern + reasoning review |
| `R-F08` | Timed Reading | complete an Academic passage or GT section/full Reading configuration under target timing | readiness evidence when valid |

## Reading variant contract

Academic and GT share Reading capabilities/question interaction families, but the workspace/content selector must preserve external variant context.

For **Academic**, readiness content must represent Academic passage characteristics.

For **General Training**, the product must be able to construct the official three-section context progression:

1. everyday/social-survival short texts;
2. workplace texts;
3. one longer, more complex general-interest text.

A GT readiness route cannot be declared complete when only Academic-style passages exist.

Reading flow:

```text
variant/context selection
→ passage/section strategy
→ independent question attempt
→ answer state
→ evidence-span review
→ reasoning/error classification
→ targeted question-type/context transfer practice
```

The app must distinguish "wrong because evidence was missed" from "wrong because the question rule was misunderstood" and from "performance failed to transfer into the required variant context".

# Writing — 9 features

| ID | Feature | Learner experience | Primary output |
|---|---|---|---|
| `W-F01` | Task Analyzer | identify variant, task type, command/situation, required content, constraints, audience/purpose where applicable, and response objective | task interpretation |
| `W-F02` | Idea & Plan Board | construct position, ideas, support, paragraph intent, and sequence before drafting | plan artifact |
| `W-F03` | Task-1 Planner | Academic: key visual features/comparisons/grouping/overview; GT: recipient/purpose/three bullet points/register/opening/closing/organization | variant-specific Task-1 plan |
| `W-F04` | Sentence & Grammar Workshop | sentence combining, correction, transformation, range/accuracy practice | focused productive attempt |
| `W-F05` | Lexical & Paraphrase Lab | paraphrase prompt/ideas, collocation, precision, word formation, spelling, and register where task-relevant | lexical observations |
| `W-F06` | Paragraph & Cohesion Builder | construct paragraph intent, development, reference, linking, and progression appropriate to the selected task | paragraph attempt |
| `W-F07` | Draft Workspace | distraction-light variant-aware Task 1/2 editor with timer, word count, plan panel, and version history | complete writing attempt |
| `W-F08` | Rubric Feedback & Revision Diff | criterion-level feedback, evidence-linked annotations, error patterns, revision comparison | FeedbackArtifact + revised attempt |
| `W-F09` | Timed Writing | 20-minute Task 1, 40-minute Task 2, or combined exam-condition session using the selected variant | readiness/evidence candidate |

## `W-F03` exact variant behavior

Academic Task 1 planning resolves to:

```text
W-TA-01 identify key visual features
W-TA-02 plan overview
W-TA-03 select supporting data/details
```

General Training Task 1 planning resolves to:

```text
W-GT1-01 recipient + relationship + purpose + all prompt bullet points
W-GT1-02 personal/semi-formal/formal register + suitable tone/opening/closing
W-GT1-03 purpose fulfilment + relevant development of every required point
```

The UI must never show Academic "overview/key trend" guidance as a required GT letter rule, and must never teach GT letter conventions as Academic visual-task criteria.

Writing feedback hierarchy:

```text
task fulfillment / response
→ variant-specific Task-1 requirement when applicable
→ organization / cohesion
→ lexical resource / register where applicable
→ grammar
→ local errors
```

Do not flood the learner with every local correction before identifying the highest-impact gap.

During assessment or exam-readiness attempts, no live autocomplete, rewriting, or AI-generated continuation may perform the target cognitive work for the learner.

# Speaking — 9 features

| ID | Feature | Learner experience | Primary output |
|---|---|---|---|
| `S-F01` | Part-1 Quick Response | short familiar-topic questions with response recording | Part-1 attempt |
| `S-F02` | Part-2 Cue Card | one-minute preparation followed by a 1–2 minute long turn | Part-2 attempt |
| `S-F03` | Part-3 Discussion | increasingly abstract follow-up questions with multi-turn response | Part-3 attempt |
| `S-F04` | Pronunciation Lab | sound contrasts, stress, chunking, intonation, connected speech, intelligibility | pronunciation observations |
| `S-F05` | Shadowing | imitate eligible source speech for rhythm, stress, connected speech, and fluency | practice attempt; not direct IELTS certification evidence |
| `S-F06` | Retell & Summarize | listen/watch/read, then retell meaning without copying wording | transfer + fluency attempt |
| `S-F07` | Transcript & Fluency Analysis | transcript, pauses, repetitions, repair, discourse markers, and selected acoustic/prosodic feedback | observations + FeedbackArtifact |
| `S-F08` | Feedback & Re-record | one targeted change followed by a fresh re-recording | remediation/recovery evidence |
| `S-F09` | Full Speaking Mock | Parts 1–3 in IELTS-like order/timing with post-session feedback | readiness/evidence candidate |

Speaking analysis must distinguish language ability from microphone/audio failure. Low-quality capture produces an evidence-quality problem, not a low Speaking judgment.

# Shared — 6 features

| ID | Feature | Purpose |
|---|---|---|
| `X-F01` | Daily Plan | explainable set of recommended actions for selected study duration |
| `X-F02` | Gap Map | show Ability, Prerequisite, Evidence, Conflict, Staleness, Scaffold, Transfer, Fluency, and Exam-Condition gaps |
| `X-F03` | Review Queue | combine knowledge review, error remediation, and re-evidence while preserving their separate semantics |
| `X-F04` | Vocabulary / Grammar SRS | spaced retrieval for suitable Knowledge Objects; not a universal IELTS-skill scheduler |
| `X-F05` | Media Lesson Creator | turn eligible/authorized media into practice using the Media design contract |
| `X-F06` | Full IELTS Mock | variant-aware integrated readiness session; never a certification shortcut |

# Full-mock variant invariant

`X-F06` must resolve the declared TargetProfile variant before mock construction:

- Listening shared;
- Speaking shared;
- Reading uses Academic or GT format/content/scoring as selected;
- Writing Task 1 uses Academic visual or GT letter construct as selected;
- Writing Task 2 uses the applicable variant prompt/configuration without duplicating shared learning truth.

# Feature-state rules

Every feature must declare whether an attempt is:

- training-only;
- evidence-eligible under conditions;
- diagnostic;
- readiness-only;
- potentially certification-contributing.

A UI feature cannot upgrade its own evidence role. `08-ASSESSMENT.md` remains the authority.

# Feature-to-domain traceability

Every implementation feature must resolve to:

```text
feature ID
→ test variant/context when material
→ canonical Skill/Knowledge target(s)
→ Practice Type or Assessment Type
→ learner-state purpose
```

A feature that cannot identify those references is product decoration, not a learning-system capability.
