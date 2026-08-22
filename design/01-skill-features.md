STATUS: CANONICAL
OWNS: user-facing feature capabilities for Listening, Reading, Writing, Speaking, and shared learning surfaces, including variant/delivery-aware feature behavior
DEPENDS_ON: ../spec/02-IELTS-MODEL.md, ../spec/03-SKILLS.md, ../spec/05-BANDS.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md
DOES_NOT_OWN: live IELTS format/delivery facts, skill/Band truth, practice mechanisms, evidence sufficiency, planner policy, API wire shape, media rights, or frameworks

# Skill Features

## Purpose

Define the learner-facing capability surface. Features consume canonical learning/exam semantics; they do not redefine them.

The initial product contains **40 named feature capabilities**:

```text
Listening  8
Reading    8
Writing    9
Speaking   9
Shared     6
---------
Total     40
```

Academic/GT and delivery differences normally appear as feature configuration/submodes where the interaction capability is shared. Do not inflate feature identity merely to mirror labels.

# Variant/delivery behavior rule

When the current TargetProfile carries a material `test_variant` or `delivery_mode`, exam-readiness features must resolve the applicable external conditions from `../spec/02-IELTS-MODEL.md`.

Feature configuration may change:

- content context;
- task shape;
- input mode;
- on-screen navigation/timing;
- handwriting rehearsal;
- remote-test familiarity;
- evidence/readiness condition labels.

It may not change Skill/Band truth.

# Listening — 8

| ID | Feature | Learner experience | Primary output |
|---|---|---|---|
| `L-F01` | IELTS Section Player | timed/untimed section playback with target-delivery-compatible question navigation | section attempt |
| `L-F02` | Dictation Lab | listen to short segments and reproduce details before reveal | detail/segmentation/spelling observations |
| `L-F03` | Gist & Main-Idea Drill | capture central meaning before detail review | gist observation |
| `L-F04` | Detail & Completion Drill | completion/short-answer work with form/word-limit validation | detail + form-control observations |
| `L-F05` | Paraphrase & Distractor Lab | compare prompt/audio meaning and identify distractor logic | paraphrase/distractor observations |
| `L-F06` | Map / Diagram Lab | spatial-language listening with visual labelling | spatial/detail observations |
| `L-F07` | Transcript & Error Review | after eligible independent attempt, inspect evidence/transcript and classify failure | ErrorPattern/remediation links |
| `L-F08` | Media Listening Lab | eligible media used through existing listening practice roles | practice attempt; non-certifying by default |

Listening flow:

```text
objective
→ independent listen/answer
→ result
→ evidence/transcript reveal when permitted
→ error classification
→ targeted remediation
→ fresh transfer when useful
```

Transcript visibility is not required before the first independent listening attempt.

# Reading — 8

| ID | Feature | Learner experience | Primary output |
|---|---|---|---|
| `R-F01` | Reading Workspace | variant-aware passage/section workspace with notes/highlighting/navigation | passage/section attempt |
| `R-F02` | Skim Sprint | identify topic/purpose/paragraph gist under a short time budget | gist/structure observation |
| `R-F03` | Scan & Detail Hunt | locate explicit facts/references/evidence spans | detail-location observation |
| `R-F04` | T/F/NG + Y/N/NG Lab | classify support/contradiction/absence or writer stance | classification observation |
| `R-F05` | Headings & Structure Lab | match paragraph purpose/headings/text organization | structure observation |
| `R-F06` | Paraphrase / Inference / Stance Lab | reason beyond keyword matching | inference/paraphrase observation |
| `R-F07` | Evidence Review | compare response with exact supporting/invalidating text evidence | ErrorPattern/reasoning review |
| `R-F08` | Timed Reading | variant-correct passage/section/full-Reading configuration under target conditions | readiness evidence candidate |

## Reading variant contract

Academic and GT reuse Reading capabilities and question interactions while concrete content/evidence preserves the selected external variant context.

GT readiness content resolves stable Content Context IDs from `../spec/10-CONTENT-MODEL.md`. Academic-only assets cannot silently satisfy a GT whole-Reading path.

Reading flow:

```text
variant/context
→ independent attempt
→ answer state
→ evidence review
→ error/reasoning classification
→ targeted question/context transfer work
```

The product should distinguish missed evidence, misunderstood question logic, and context-transfer failure when observations support that distinction.

# Writing — 9

| ID | Feature | Learner experience | Primary output |
|---|---|---|---|
| `W-F01` | Task Analyzer | identify variant, task/situation, constraints, required content, audience/purpose when applicable | task interpretation |
| `W-F02` | Idea & Plan Board | construct position/ideas/support/paragraph intent before drafting | plan artifact |
| `W-F03` | Task-1 Planner | Academic visual planning or GT letter recipient/purpose/register/required-point planning | variant-specific Task-1 plan |
| `W-F04` | Sentence & Grammar Workshop | combine/correct/transform sentence forms | focused productive attempt |
| `W-F05` | Lexical & Paraphrase Lab | practise collocation, precision, paraphrase, spelling, register where relevant | lexical observations |
| `W-F06` | Paragraph & Cohesion Builder | develop paragraph purpose, progression, reference, and linking | paragraph attempt |
| `W-F07` | Draft Workspace | variant-aware Task 1/2 drafting with timer, word count, plan, history | complete Writing attempt |
| `W-F08` | Rubric Feedback & Revision Diff | criterion feedback, evidence-linked annotations, revision comparison | FeedbackArtifact + revised attempt |
| `W-F09` | Timed Writing | independent full-task/combined Writing under selected variant and supported delivery-input condition | readiness/evidence candidate |

## `W-F03` variant mapping

Academic Task 1 resolves to `W-TA-01..03`.

GT Task 1 resolves to `W-GT1-01..03` plus shared Writing capability.

Academic overview/visual-feature guidance is never presented as a GT letter rule; GT letter convention is never presented as Academic visual-task truth.

## Writing delivery behavior

`W-F07` is a normal learning/drafting workspace and may be delivery-agnostic.

`W-F09` is exam-readiness oriented and therefore resolves a supported input configuration, for example:

- typed computer response;
- handwriting rehearsal for an eligible Writing-on-Paper target, with actual assistance/timing/input context recorded;
- remote computer-interface rehearsal when the target is an eligible online delivery.

A delivery-specific rehearsal does not alter Writing scoring criteria or Band semantics.

Writing feedback prioritizes task fulfilment/response and high-impact structure/language issues before local corrections.

During evidence/readiness attempts, autocomplete/rewriting/AI continuation may not perform the target cognitive work for the learner.

# Speaking — 9

| ID | Feature | Learner experience | Primary output |
|---|---|---|---|
| `S-F01` | Part-1 Quick Response | recorded familiar-topic short responses | Part-1 attempt |
| `S-F02` | Part-2 Cue Card | preparation + sustained long turn | Part-2 attempt |
| `S-F03` | Part-3 Discussion | abstract follow-up discussion | Part-3 attempt |
| `S-F04` | Pronunciation Lab | sounds, stress, chunking, intonation, connected speech | pronunciation observations |
| `S-F05` | Shadowing | imitate eligible speech for prosody/fluency targets | training attempt |
| `S-F06` | Retell & Summarize | reconstruct meaning without copying wording | transfer/fluency attempt |
| `S-F07` | Transcript & Fluency Analysis | review pauses/repairs/discourse/prosodic observations | observations + FeedbackArtifact |
| `S-F08` | Feedback & Re-record | apply one targeted change in a fresh recording | remediation/recovery attempt |
| `S-F09` | Full Speaking Mock | Parts 1–3 sequence with target-condition-aware preparation/feedback | readiness/evidence candidate |

Capture failure or poor microphone quality is evidence-quality state, not low Speaking ability.

Where external delivery uses a different interaction channel, readiness configuration may rehearse that channel; it does not redefine the human-interactive Speaking construct.

# Shared — 6

| ID | Feature | Purpose |
|---|---|---|
| `X-F01` | Daily Plan | explainable eligible recommended actions |
| `X-F02` | Gap Map | show learner gap/evidence states without mixing CoverageGap |
| `X-F03` | Review Queue | present retrieval, remediation, and re-evidence while preserving semantics |
| `X-F04` | Vocabulary / Grammar SRS | spaced retrieval for suitable Knowledge Objects |
| `X-F05` | Media Lesson Creator | create eligible media-supported practice under Media contract |
| `X-F06` | Full IELTS Mock | variant/delivery-aware integrated readiness run; never certification shortcut |

# Full-mock invariant

`X-F06` resolves the current TargetProfile before construction:

- Listening shared construct;
- Speaking shared construct;
- Reading selected variant context/scoring;
- Writing Task 1 selected variant construct;
- Task 2 applicable configuration;
- delivery interaction only where the declared supported target requires it.

A mixed variant is invalid for a normal full readiness claim.

# Feature activity-purpose/evidence boundary

A feature may host training, diagnostic, or readiness-oriented activity configurations. **Primary product purpose and evidence candidacy are separate dimensions.**

A diagnostic/readiness configuration is not automatically evidence-admissible, and a training-oriented configuration is not automatically evidence-ineligible. Concrete activity packaging declares whether Assessment may consider the resulting Observation; `08-ASSESSMENT.md` makes the actual claim-scoped admission decision after the attempt conditions are known.

A feature may never:

- upgrade a non-candidate activity into evidence after seeing a favorable result;
- treat diagnostic/readiness purpose as proof of eligibility;
- pre-label an activity as certification-contributing;
- certify a learner or lower normal Assessment requirements.

# Traceability

Every implementation feature resolves:

```text
feature ID
→ TargetProfile variant/delivery context when material
→ stable Content Context when material
→ canonical target IDs
→ Practice Type or Assessment Type
→ learner-state purpose
```

A feature without these references is product decoration, not a learning-system capability.