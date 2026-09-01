STATUS: SUPERSEDED_NON_CANONICAL
OWNS: user-facing feature capabilities for Listening, Reading, Writing, Speaking, and shared learning surfaces, including variant/delivery-aware feature behavior and Speaking capture/interaction capability boundaries
DEPENDS_ON: ../spec/02-IELTS-MODEL.md, ../spec/03-SKILLS.md, ../spec/05-BANDS.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md
DOES_NOT_OWN: live IELTS format/delivery facts, skill/Band truth, practice mechanisms, evidence sufficiency, planner policy, API wire shape, media rights, or frameworks

# Skill Features

## Purpose

Define the learner-facing capability surface. Features consume canonical learning/exam semantics; they do not redefine them.

The current design inventory contains **40 named feature capabilities**:

```text
Listening  8
Reading    8
Writing    9
Speaking   9
Shared     6
---------
Total     40
```

Academic/GT and delivery differences normally appear as feature configuration/submodes where the interaction capability is shared. The number 40 is not a completeness target; add/merge/remove feature identities only when learner interaction semantics require it.

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

## Speaking capture and interaction boundary

The ordinary digital Speaking route uses browser microphone capture when available and remains complete without realtime AI:

```text
record
→ replay/re-record before submission when allowed
→ submit one learner performance
→ process / evaluate within the configured consequence
→ feedback
→ later review / fresh re-record where useful
```

Part-specific behavior remains recognizable:

- Part 1 supports short spontaneous responses and follow-up practice;
- Part 2 preserves preparation followed by a sustained learner long turn;
- Part 3 supports responsive abstract discussion and follow-up questions;
- general role-play/conversation may train responsiveness/fluency when mapped to canonical targets, but does not become an IELTS part merely because it is conversational.

### Beginner L1-to-English speaking scaffold

For learners who cannot yet respond reliably from an English-only prompt, existing Speaking features may use the learner's configured L1 as a temporary acquisition scaffold. This is a configuration of existing guided/controlled production and scaffold-fading mechanisms from `../spec/07-PRACTICE.md`, not another Speaking feature, IELTS construct, or evidence standard.

A typical progression is:

```text
L1 situation / meaning cue
→ learner produces the target response in English
→ concise correction, model, contrast, or hint when needed
→ repeat or vary with less L1/support
→ English-only prompt
→ spontaneous English response
→ fresh transfer into the applicable Part 1/2/3 or other canonical Speaking target
```

Rules:

1. L1 carries meaning or task context; it must not perform the learner's required English production for them.
2. Support may include a simple L1 explanation, constrained cue, partial model, contrast, or learner-requested clarification, but the actual target response remains learner-produced English.
3. As performance stabilizes, the product reduces L1 and other scaffold strength rather than making bilingual prompting permanent by default.
4. A learner may remain on or return to stronger support when current learning need justifies it; scaffold fading is adaptive, not a one-way level badge.
5. Success under material L1/model support records that assistance state and cannot masquerade as independent English Speaking evidence.
6. Independent Assessment/readiness claims use the normal assistance/independence rules from `../spec/08-ASSESSMENT.md`; bilingual scaffolded training does not lower those standards.
7. Translation quality, wording choice, or failure to understand an L1 prompt is not itself a Speaking weakness unless the scoped English capability was actually measured.
8. The same pattern may serve simple rapid-response, role-play, retell, or targeted productive practice when bound to a canonical target; it does not create a generic conversation curriculum outside the IELTS learning model.

The learner-facing language/scaffolding presentation remains governed by `00-learning-experience.md`, including preserved English/IELTS terminology and progressive disclosure.

Inference follows the signal actually available:

- transcript/text may support lexical, grammatical, discourse/content observations within evaluator quality limits;
- trustworthy turn/timing metadata may support bounded fluency observations where the intended claim allows it;
- phoneme quality, word/sentence stress, intonation, connected speech, and overall intelligibility require suitable acoustic evidence plus an eligible evaluator/calibration path;
- transcript-only, failed STT, missing audio, or acoustically unusable audio must not claim pronunciation quality;
- capture/provider failure is not learner silence, hesitation, pronunciation weakness, or low Speaking ability.

An optional realtime AI conversation overlay may add responsive turn-taking, context-aware follow-ups, role-play, interruption/barge-in where the interaction supports it, and lower-latency fluency practice. It is a delivery capability, not an examiner, Skill, Assessment standard, or learning authority.

Realtime rules:

1. AI turns/prompts and learner spoken turns preserve reconstructable order/provenance when the session matters to feedback/evidence;
2. learner silence is interpreted as learner performance only when trustworthy capture establishes that silence occurred under an activity whose normal Assessment conditions make it material; capture uncertainty is never converted into silence evidence;
3. AI latency/network interruption may pause, reconnect, degrade, or end the interaction without becoming learner failure;
4. fallback from realtime to record/submit is allowed only when it still performs the selected learning purpose. A readiness/mock condition requiring responsive target-like interaction remains unresolved rather than silently changing activity meaning;
5. partial realtime sessions preserve only the scope actually completed; completion/abandonment does not imply evidence admission;
6. realtime interaction is evidence-eligible only when independently configured as a candidate and normal Assessment/capture/independence/evaluator conditions pass. Cost or interaction difficulty never upgrades evidence status.

Where external delivery uses a different interaction channel, readiness configuration may rehearse that channel; it does not redefine the human-interactive Speaking construct.

## Capture-quality product semantics

Browser/device/audio handling must preserve these distinct outcomes conceptually rather than collapsing them into a Speaking score:

- **usable performance** — required audio/interaction signal is sufficiently captured for the configured use;
- **capture unavailable** — permission denied, no eligible device, device/init failure, or an unsupported browser/device path prevents capture;
- **capture uncertain/unusable** — clipping, inaudibility, severe noise, interruption, truncation, device change, or another condition prevents trustworthy interpretation for the intended consequence;
- **transcription unavailable/conflicting** — audio may remain usable while STT fails or transcript materially disagrees with the audio;
- **acoustic evaluation unavailable** — usable audio exists but the required acoustic evaluator/calibration path cannot currently support the requested inference;
- **product capability unsupported** — the requested supported interaction/evidence route does not exist for the scoped product release.

The product may request permission/device correction, retry capture, re-record, another supported device/interaction, delayed evaluation, or another eligible assessment path according to the failure class. It must not lower the evidence standard or convert technical/accessibility failure into a learner gap.

Accessibility or assistive interaction may change presentation, controls, timing mechanics, capture method, or delivery workflow when supported. Any material difference relevant to exam/readiness inference remains visible in the attempt conditions; accommodation never silently changes Skill/Band truth or makes an otherwise ineligible evidence claim eligible.

# Shared — 6

| ID | Feature | Purpose |
|---|---|---|
| `X-F01` | Daily Plan | explainable eligible recommended actions |
| `X-F02` | Gap Map | show learner gap/evidence states without mixing CoverageGap |
| `X-F03` | Review Queue | present retrieval, remediation, and re-evidence while preserving semantics |
| `X-F04` | Knowledge Lab & SRS | vocabulary/grammar/phonology acquisition, learner-saved study material, and spaced retrieval for suitable Knowledge Objects |
| `X-F05` | Media Lesson Creator | create eligible media-supported practice under Media contract |
| `X-F06` | Full IELTS Mock | variant/delivery-aware integrated readiness run; never certification shortcut |

# Cross-feature AI tutor capability

AI tutoring is a delivery overlay across eligible features rather than a 41st feature. It may explain, model, hint, question, generate bounded practice, and summarize feedback while preserving the target operation and the activity assistance/evidence state defined upstream. A feature remains usable through non-realtime/asynchronous interaction unless its own learning purpose materially requires realtime exchange.

# Traditional workflow correspondence

The current feature surface digitally represents the major repeatable standard-preparation workflows without treating this table as a completeness score:

| Traditional workflow class | Product path |
|---|---|
| teacher explanation / worked example | AI tutor overlay + guided feature activity |
| homework / focused drills | `Practice` modes across `L/R/W/S` |
| vocabulary notebook / grammar review | `X-F04` + `X-F03` |
| teacher correction / redraft / re-record | `W-F08`, `S-F07`, `S-F08` |
| speaking partner / responsive practice | `S-F01..03`, optional realtime AI overlay |
| reasoning / rapid-response practice | `R-F06`, `R-F07`, `S-F01`, `S-F03` |
| timed section/task work | `L-F01`, `R-F08`, `W-F09`, `S-F09` |
| progress conference / study plan | `X-F01`, `X-F02`, Progress surface |
| authentic-media homework | `L-F08`, `S-F05`, `S-F06`, `X-F05` |
| section/full mock | skill timed features + `X-F06` |

Remaining material coverage is judged by target/construct/workflow reachability and `08-coverage-and-support.md`, not by increasing the number 40.

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

A feature may host `TRAINING`, `DIAGNOSTIC`, focused `ASSESSMENT`, or `READINESS` activity configurations. **Primary product purpose and evidence candidacy are separate dimensions**, with exact purpose/candidacy semantics owned by `02-practice-catalog.md`.

A diagnostic/assessment/readiness configuration is not automatically evidence-admissible, and a training-oriented configuration is not automatically evidence-ineligible. Concrete activity packaging declares whether Assessment may consider the resulting Observation; `08-ASSESSMENT.md` makes the actual claim-scoped admission decision after the attempt conditions are known.

A feature may never:

- upgrade a non-candidate activity into evidence after seeing a favorable result;
- treat diagnostic/assessment/readiness purpose as proof of eligibility;
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
