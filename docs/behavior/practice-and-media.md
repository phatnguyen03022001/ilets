# PRACTICE-MEDIA Practice, learner-visible modes, and media semantics

> **CANONICAL BEHAVIOR AUTHORITY**
>
> Canonical within the BEHAVIOR authority domain under `docs/catalog/project.json`. `CONSTITUTION.md` and `OBJECTIVE.md` retain distinct authority, and `contracts/**` retains scoped exact machine-contract authority. References below to `spec/**` or `design/**`, including historical uses of “canonical”, are pre-cutover provenance only and do not create equal authority.

## Learning phases

Every Practice Type has one primary phase and may support others:

- **Acquisition** — initial construction of a new capability/knowledge pattern.
- **Consolidation** — stabilize newly acquired capability through focused varied use/correction.
- **Retrieval** — actively recall/reproduce after delay before feedback.
- **Transfer** — apply learning in a materially new, less scaffolded or more integrated context.
- **Fluency** — increase speed/automaticity/rhythm/processing efficiency without sacrificing target quality.
- **Exam Readiness** — perform under IELTS-like timing, integration, stamina and independence conditions.

Exam Readiness is a practice phase; Exam Preparation is a learner mode. Practice is selected from diagnosed need, not Band membership, and does not itself admit Assessment evidence.

## Learning Mechanism registry

| ID | Mechanism | Use when | Must not infer |
|---|---|---|---|
| `LM-01` | Worked example | model/attention guidance needed | copying proves independent performance |
| `LM-02` | Active retrieval | reviewable target should be recalled/reproduced | every target belongs in flashcards |
| `LM-03` | Spaced review | suitable target has repeated retrieval history | one spacing formula fits all objects |
| `LM-04` | Contrast / discrimination | near alternatives/cues are confused | random variety is useful by itself |
| `LM-05` | Controlled production | recognition exists but production unstable | constrained success proves free production |
| `LM-06` | Guided production | failure-relevant support is needed | support should remain indefinitely |
| `LM-07` | Scaffold fading | material support carries required performance | all scaffolds disappear at once |
| `LM-08` | Variation / transfer | generalization beyond practiced context needed | novelty alone proves transfer |
| `LM-09` | Interleaving | flexible selection/discrimination benefits from mixing stable targets | arbitrary mixing always helps |
| `LM-10` | Self-explanation | reasoning/discrimination becomes more observable through explanation | every item needs reflection |
| `LM-11` | Deliberate revision / reproduction | corrected productive output has learning value | every error requires rewrite/re-record |
| `LM-12` | Fluency rehearsal | quality is stable enough and efficiency/rhythm limits performance | speed compensates for weak quality |

AI tutoring is delivery/runtime capability, not a Learning Mechanism, and may not perform the learner operation the selected mechanism requires.

## Practice Type registry

`Aq` Acquisition, `Co` Consolidation, `Re` Retrieval, `Tr` Transfer, `Fl` Fluency, `ER` Exam Readiness.

| ID | Type | Primary / supported phases | Mode/scope | Typical mechanisms | Target role |
|---|---|---|---|---|---|
| `PT-01` | Scaffolded/guided writing | Aq / Aq,Co | individual/adaptive writing | `LM-01`,`LM-06`,`LM-07` | Writing acquisition/support |
| `PT-02` | Sentence-combining drill | Aq / Aq,Fl | individual writing | `LM-05`,`LM-12` | grammatical production/control |
| `PT-03` | Paraphrase task | Co / Co,Tr | individual writing | `LM-04`,`LM-05`,`LM-08` | lexical/cohesive flexibility |
| `PT-04` | Error-correction exercise | Co / Co,Re | individual writing | `LM-02`,`LM-04`,`LM-10` | language-error diagnosis/correction |
| `PT-05` | Redraft after feedback | Tr / Co,Tr | individual writing | `LM-11`,`LM-07`,`LM-08` | apply feedback in new production |
| `PT-06` | Timed Writing task | ER / ER,Tr | timed writing | `LM-08` | independent/timed integrated Writing |
| `PT-07` | Pronunciation/minimal-pair drill | Aq / Aq,Re | individual speaking | `LM-04`,`LM-02` | pronunciation discrimination/production |
| `PT-08` | Shadowing | Fl / Aq,Fl | individual speaking | `LM-01`,`LM-12` | pronunciation/prosody/fluency imitation where relevant |
| `PT-09` | Long-turn practice | Co / Co,Tr,Fl | timed speaking | `LM-06`,`LM-07`,`LM-08`,`LM-12` | sustained Speaking production |
| `PT-10` | Q&A / role-play | Co / Co,Tr | mixed speaking | `LM-05`,`LM-06`,`LM-08` | responsive Speaking production |
| `PT-11` | Timed mock Speaking | ER / ER | timed speaking | `LM-08` | integrated exam-like Speaking |
| `PT-12` | Skim/scan/gist-detail speed drill | Fl / Aq,Fl | timed listening/reading | `LM-04`,`LM-12` | selective receptive processing |
| `PT-13` | Comprehension question set | Co / Co,Re | individual listening/reading | `LM-02`,`LM-04` | receptive comprehension/question-family work |
| `PT-14` | Note-taking from lecture/text | Tr / Co,Tr | individual listening/reading | `LM-08`,`LM-10` | information selection/integration |
| `PT-15` | Timed section/passage practice | ER / ER,Fl | timed listening/reading | `LM-08`,`LM-12` | receptive section timing |
| `PT-16` | Distractor/error review | Re / Re,Co | individual listening/reading | `LM-04`,`LM-10` | discriminate why alternatives failed |
| `PT-17` | Spaced retrieval | Re / Re,Co | adaptive reviewable targets | `LM-02`,`LM-03` | retention of suitable Skill/Knowledge targets |
| `PT-18` | Collocation/word-formation drill | Aq / Aq,Co | individual knowledge/writing | `LM-04`,`LM-05` | lexical production/discrimination |
| `PT-19` | Gap-fill/completion | Co / Co,Re | knowledge/receptive | `LM-02`,`LM-05` | constrained retrieval/application |
| `PT-20` | Interleaved mixed set | Tr / Tr,Re | mixed cross | `LM-09`,`LM-04`,`LM-08` | flexible target selection |
| `PT-21` | Adaptive practice set | Co / Co,Re,Tr | adaptive cross | selected GapEvaluation-dependent | evidence-based activity container |
| `PT-22` | Diagnostic checkpoint practice | Re / Re,Co | diagnostic cross | uncertainty-dependent | low-friction sampling; measurement interpretation belongs to `AT-04` |
| `PT-23` | Full mock test | ER / ER | timed cross | none by default | broad readiness/re-evidence with possible learning side-effects |

### Canonical materializer anchor — receptive classification practice types

| ID | Type | Primary / supported phases | Mode / scope | Typical mechanisms | Canonical target role |
|---|---|---|---|---|---|
| `PT-13` | Comprehension question set | Co / Co, Re | individual / listening, reading | `LM-02`, `LM-04` | receptive comprehension/question-type work |
| `PT-16` | Distractor/error review | Re / Re, Co | individual / listening, reading | `LM-04`, `LM-10` | discriminate why an alternative failed |

Concrete binding resolves target families to explicit stable Skill/Knowledge IDs before execution/evidence recording.

## Gap-to-practice selection

1. Ability gap → target the demonstrated failure pattern; do not default to prerequisite reteaching.
2. Prerequisite gap → acquire/repair the actually Required prerequisite before dependent work.
3. Insufficient evidence → use low-friction diagnostic/assessment; do not label weakness.
4. Conflicting evidence → choose discriminating evidence.
5. Stale evidence → smallest useful representative refresh before remediation.
6. Scaffold dependence → identify carrying support, fade selectively, then re-evidence.
7. Transfer gap → vary context/reduce pattern repetition and use sufficiently fresh/materially different conditions when the claim requires generalization.
8. Fluency gap → rehearsal only after target quality is sufficiently stable.
9. Exam-condition gap → timed/integrated work without redefining underlying mastery.

Feedback timing follows phase: generally immediate/actionable in Acquisition; self-correction before reveal in Consolidation; authentic retrieval before feedback in Retrieval; increasingly delayed/batched in Transfer; avoid excessive interruption in Fluency; preserve independence during Exam Readiness and review primarily afterward.

Feedback is target-led. Priority is: current target; Required prerequisite/material blocker; issue invalidating interpretation/performance/inference; then other useful observations may be recorded/deferred. Detection alone does not require immediate correction. Use the smallest useful correction set; do not surface unrelated teaching or perform the learner's required operation. There is no universal maximum feedback count.

## Intervention-effectiveness semantics

Practice strategy is revisable only from comparable history scoped by canonical target, ActionIntent, Learning Mechanism, Practice Type, scaffold/support, context/transfer condition, difficulty/load, relevant attempts/outcomes, later admissible evidence where separately available, recency/comparability and learner friction/fatigue.

Too little/incomparable history keeps effectiveness unresolved; lack of improvement alone does not prove greater weakness/regression/prerequisite gap; technical/content/evaluator/provider failure is excluded from learner-efficacy inference; effective repetition may continue; justified diminishing returns may change an existing mechanism/type/scaffold/context/difficulty/load/frequency/sequence while preserving target and Required prerequisites; diagnostic uncertainty routes to discriminating evidence. Strategy change never rewrites Attempts/Observations/EvidenceFacts or the target/Band standard, and intervention history is not Assessment evidence unless separately admitted.

Practice completion is not mastery. Same-item retry is mainly recovery, not automatic transfer evidence. Learner preference/friction may influence planning but not target truth or ability evidence.

## Beginner L1-to-English Speaking scaffold

L1 support is a temporary, adaptive scaffold for a learner who cannot yet use an English-only Speaking prompt reliably. It reuses the existing Speaking targets, Practice Types, Learning Mechanisms, feedback, Assessment, and readiness semantics; it creates no new Speaking feature, curriculum, mechanism, or competence standard.

The staged path is:

```text
L1 situation / meaning cue
→ learner-produced English
→ concise correction / model / contrast / hint
→ repeat or vary with less L1/support
→ English-only prompt
→ spontaneous English response
→ fresh transfer to the appropriate Speaking target
```

Rules:

- L1 supplies meaning or task context only; the learner's target production remains English.
- A simple L1 explanation, constrained cue, partial model, contrast, or requested clarification may be used when it enables the learner to perform the English target operation. The support must not silently perform that operation for the learner.
- Scaffolding should fade as stable performance permits. Adaptive behavior may strengthen or reintroduce support when justified; support level is not a learner badge or a replacement proficiency scale.
- Material L1 assistance, supplied models, hints, corrections, prior-answer exposure, and other carrying support are recorded as attempt conditions. L1/model-supported performance cannot masquerade as independent Speaking evidence.
- Normal Assessment and readiness rules remain unchanged: a claim requiring independent/unseen/spontaneous production needs fresh evidence that satisfies its actual assistance, exposure, evaluator, provenance, and transfer conditions.
- Weak L1 translation or task-understanding quality is not a Speaking weakness unless the English capability itself was validly measured. Product/translation failure likewise cannot become learner evidence.
- The scaffold may support rapid response, role-play/Q&A, retell/summarize, or other targeted productive practice tied to a canonical Speaking target. It must not expand into a generic conversation curriculum disconnected from that target.

## Learner-visible Practice Mode registry

The locked catalog contains exactly 28 modes: 6 Listening, 6 Reading, 6 Writing, 6 Speaking and 4 shared. The count is descriptive, not a completeness target; variant/delivery differences stay inside a mode when interaction is materially unchanged.

| ID | Mode | Default duration | Role/backing |
|---|---|---:|---|
| `PM-L01` | Dictation | 5–12m | detail discrimination/segmentation/spelling; `PT-13`,`PT-19` |
| `PM-L02` | Gist Sprint | 3–6m | main idea/discourse; `PT-12` |
| `PM-L03` | Detail & Completion | 5–10m | explicit detail/completion/word limits; `PT-13`,`PT-19` |
| `PM-L04` | Paraphrase & Distractor | 5–10m | paraphrase/distractor discrimination; `PT-16` |
| `PM-L05` | Map / Diagram | 5–10m | spatial/structured detail; `PT-13` |
| `PM-L06` | Timed Section | 10–35m | integrated Listening readiness; `PT-15` |
| `PM-R01` | Skim Sprint | 3–5m | topic/purpose/paragraph gist; `PT-12` |
| `PM-R02` | Scan & Detail Hunt | 4–8m | explicit information location; `PT-12`,`PT-13` |
| `PM-R03` | T/F/NG + Y/N/NG | 6–12m | evidence/stance classification; `PT-13`,`PT-16` |
| `PM-R04` | Headings & Structure | 6–12m | paragraph function/organization; `PT-13` |
| `PM-R05` | Paraphrase / Inference / Stance | 6–12m | inferential reasoning; `PT-13`,`PT-16` |
| `PM-R06` | Timed Reading | 15–60m by scope | variant-aware passage/section/full Reading; `PT-15` |
| `PM-W01` | Prompt & Plan | 5–10m | variant-aware planning; `PT-01` |
| `PM-W02` | Sentence & Grammar Lab | 5–10m | structural accuracy/range/correction; `PT-02`,`PT-04` |
| `PM-W03` | Lexical & Paraphrase Lab | 5–10m | collocation/precision/paraphrase/spelling/register; `PT-03`,`PT-18` |
| `PM-W04` | Paragraph & Cohesion Builder | 8–15m | organization/reference/cohesion/development; `PT-01`,`PT-05` |
| `PM-W05` | Guided Draft & Redraft | 15–30m | scaffolded production/revision; `PT-01`,`PT-05` |
| `PM-W06` | Timed Writing | ≈20m Task 1 / 40m Task 2 | independent full-task performance; `PT-06` |
| `PM-S01` | Pronunciation Contrast | 5–10m | phoneme/stress/intonation; `PT-07` |
| `PM-S02` | Shadowing | 5–12m | rhythm/connected speech/stress/fluency; `PT-08` |
| `PM-S03` | Part-1 Quick Response | 5–10m | spontaneous short response; `PT-10` |
| `PM-S04` | Part-2 Long Turn | 5–10m | sustained long turn; `PT-09` |
| `PM-S05` | Part-3 Discussion | 8–15m | abstract explanation/comparison/justification; `PT-10` |
| `PM-S06` | Full Speaking Mock | 11–14m | Parts 1–3 readiness; `PT-11` |
| `PM-X01` | Knowledge Learn & Review | 5–15m | vocabulary/grammar/phonology + retrieval; `PT-17`,`PT-18`,`PT-19`, `PT-07` where phonological production is targeted |
| `PM-X02` | Error Remediation | 5–15m | classified error→fresh corrective action; `PT-04` + RemediationPattern |
| `PM-X03` | Adaptive Mixed Set | 10–20m | interleaved eligible targets; `PT-20`,`PT-21` |
| `PM-X04` | Full IELTS Mock | ≈150m L+R+W; Speaking per applicable separate/scheduled configuration | variant/delivery whole-test readiness; `PT-23` |

### Canonical materializer anchor — Reading classification mode

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-R03` | T/F/NG + Y/N/NG | 6–12 min | evidence/stance classification | `PT-13`, `PT-16`, Contrast |

`PM-R06` resolves variant/scope; full-Reading readiness uses complete applicable variant conditions. `PM-W06` records material delivery/input condition. Task-1 mode instantiation resolves Academic to `W-TA-01/02/03` and GT to `W-GT1-01/02/03` plus shared Writing. `PM-X04` resolves variant before Reading/Writing Task 1 and does not allow mixed Academic/GT as normal readiness.

### Activity purpose and evidence candidacy

Exactly one primary purpose describes why an activity is scheduled: `TRAINING`, `DIAGNOSTIC`, `ASSESSMENT`, or `READINESS`. Separately, pre-attempt evidence candidacy is exactly `NOT_EVIDENCE_CANDIDATE` or `ASSESSMENT_MAY_ADMIT`. The dimensions are orthogonal; purpose never admits evidence. `ASSESSMENT_MAY_ADMIT` only permits later claim-scoped Assessment consideration after actual assistance, exposure, evaluator quality, provenance and other conditions are known. There is no pre-attempt `CERTIFICATION_CONTRIBUTING` role.

#### Primary activity purpose

```text
TRAINING
DIAGNOSTIC
ASSESSMENT
READINESS
```

#### Evidence candidacy

```text
NOT_EVIDENCE_CANDIDATE
ASSESSMENT_MAY_ADMIT
```

Retries remain purpose-labelled: recovery retry, faded retry, retention retry, transfer retry, re-evidence attempt. Immediate same-item retry is normally recovery, not independent transfer evidence. Catalog durations are UX/planning defaults, never dosage/mastery/certification. A new PM mode requires materially distinct learner-visible goal, interaction model, unrepresentable task family, or learning/evidence role; not merely topic/source/Band/variant/delivery/template.

## Media and YouTube semantics

Media is a source, never a Skill, Knowledge Object, Learning Mechanism or Practice Type. Eligible media may instantiate existing modes for dictation, shadowing, retell/comprehension and vocabulary/collocation mining. Transcript-dependent keys/scoring require authorized text.

### Transcript-state registry

- `AUTHORIZED_CREATOR_CAPTION` — caption access through an authorized account/path with required platform permission; may support segmentation, trusted dictation keys, vocabulary extraction and lesson generation within permitted use.
- `LICENSED_TRANSCRIPT` — independently owned/licensed transcript/timed text with provenance/rights identifying that source.
- `USER_PROVIDED_TRANSCRIPT` — learner/content-author supplied text under applicable usage terms; must not be represented as official platform captions unless independently established.
- `NO_AUTHORIZED_TRANSCRIPT` — embedding/playback may still support shadowing, notes/bookmarks, retell/discussion and learner-entered vocabulary, but trustworthy caption-derived answer keys and automated transcript-dependent scoring are disabled.

The initial architecture explicitly does **not** authorize `arbitrary public URL → unofficial downloader → server media copy → unrestricted transcription/storage`. Broader extraction requires a new legal/platform/rights/storage/quota/deletion/privacy decision.

For YouTube playback, implementation must use a live supported embed/player path, preserve required controls/branding/visibility, avoid prohibited control overlays, preserve required origin/referrer behavior, not strip ads/standard player behavior, identify source truthfully and provide independent learning value. External platform policy remains external/live authority and is rechecked before implementation/release. Arbitrary public-video caption download is not assumed.

A `MediaSource` preserves provider/source ID, canonical URL, title/author/channel, language, duration, playability/embeddability, rights state, transcript state/ref and metadata snapshot. A `MediaLesson` preserves source, segment bounds, Practice Mode, canonical targets, transcript ref where required/permitted, prompts, scaffold/feedback policy, difficulty and provenance. Generated proposals become valid lessons only after target, mode, rights and transcript eligibility validate; persisted reusable lessons enter normal ContentRevision/validation/assignment semantics.

Media-based practice is non-certifying by default. A media Observation becomes formal evidence only through normal target fit, independence, provenance, scoring-quality and inference-scope admission. Preferred YouTube persistence is reference-based, not mirrored audiovisual content.

If a source disappears, becomes private or loses embedding/eligibility, preserve historical attempts/evidence at their valid historical scope, mark source unavailable/ineligible, stop new scheduling, substitute an eligible equivalent when possible, and never reinterpret source loss as learner regression.
