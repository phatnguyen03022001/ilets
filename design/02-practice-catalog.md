STATUS: CANONICAL
OWNS: initial user-facing practice-mode catalog, mode counts, default durations, mode-to-learning-role mapping, and product-level catalog change policy
DEPENDS_ON: ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, 01-skill-features.md
DOES_NOT_OWN: Learning Mechanism definitions, Practice Type semantics, mastery/evidence rules, skill thresholds, concrete item content, or scheduling implementation

# Practice Catalog

## Purpose

Define the concrete practice choices the learner sees in the initial product.

The learning specification owns 23 canonical Practice Types and 12 Learning Mechanisms. This document owns the **user-facing catalog** that packages those semantics into understandable activities.

The initial product exposes exactly **28 practice modes**:

```text
Listening   6
Reading     6
Writing     6
Speaking    6
Shared      4
----------
Total      28
```

Adding or removing a user-facing mode is a product-design change. It does not require a new canonical Practice Type when the new surface is only another composition of existing mechanisms/types.

# Listening — 6 modes

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-L01` | Dictation | 5–12 min | detail discrimination, lexical segmentation, spelling | `PT-13`, `PT-19`, Retrieval/Contrast |
| `PM-L02` | Gist Sprint | 3–6 min | main idea and discourse-level listening | `PT-12`, Fluency/Transfer |
| `PM-L03` | Detail & Completion | 5–10 min | explicit detail, completion, word-limit control | `PT-13`, `PT-19` |
| `PM-L04` | Paraphrase & Distractor | 5–10 min | paraphrase recognition and distractor rejection | `PT-16`, Contrast/Self-explanation |
| `PM-L05` | Map / Diagram | 5–10 min | spatial language and structured detail | `PT-13` |
| `PM-L06` | Timed Section | 10–35 min | integrated exam-condition performance | `PT-15`, Exam Readiness |

# Reading — 6 modes

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-R01` | Skim Sprint | 3–5 min | topic, purpose, paragraph gist | `PT-12`, Fluency |
| `PM-R02` | Scan & Detail Hunt | 4–8 min | explicit information location | `PT-12`, `PT-13` |
| `PM-R03` | T/F/NG + Y/N/NG | 6–12 min | evidence classification and stance logic | `PT-13`, `PT-16`, Contrast |
| `PM-R04` | Headings & Structure | 6–12 min | paragraph function and text organization | `PT-13`, Transfer |
| `PM-R05` | Paraphrase / Inference / Stance | 6–12 min | inferential comprehension beyond keyword matching | `PT-13`, `PT-16`, Self-explanation |
| `PM-R06` | Timed Passage | 20 min target | integrated passage performance under time | `PT-15`, Exam Readiness |

# Writing — 6 modes

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-W01` | Prompt & Plan | 5–10 min | task analysis, position, idea organization | `PT-01`, Guided production |
| `PM-W02` | Sentence & Grammar Lab | 5–10 min | structural accuracy/range and correction | `PT-02`, `PT-04`, Controlled production |
| `PM-W03` | Lexical & Paraphrase Lab | 5–10 min | collocation, precision, paraphrase, spelling | `PT-03`, `PT-18` |
| `PM-W04` | Paragraph & Cohesion Builder | 8–15 min | paragraph logic, reference, cohesion, development | `PT-01`, `PT-05` |
| `PM-W05` | Guided Draft & Redraft | 15–30 min | scaffolded production followed by revision | `PT-01`, `PT-05`, Scaffold fading |
| `PM-W06` | Timed Writing | 20 min Task 1 / 40 min Task 2 | independent full-task performance | `PT-06`, Exam Readiness |

# Speaking — 6 modes

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-S01` | Pronunciation Contrast | 5–10 min | phoneme/stress/intonation discrimination and production | `PT-07`, Contrast/Controlled production |
| `PM-S02` | Shadowing | 5–12 min | rhythm, connected speech, stress, fluency imitation | `PT-08`, Fluency rehearsal |
| `PM-S03` | Part-1 Quick Response | 5–10 min | spontaneous short response and familiar-topic flexibility | `PT-10`, Retrieval/Fluency |
| `PM-S04` | Part-2 Long Turn | 5–10 min | planning + sustained 1–2 minute response | `PT-09`, Transfer/Fluency |
| `PM-S05` | Part-3 Discussion | 8–15 min | abstract explanation, comparison, justification | `PT-10`, Transfer |
| `PM-S06` | Full Speaking Mock | 11–14 min | integrated Parts 1–3 exam-condition performance | `PT-11`, Exam Readiness |

# Shared — 4 modes

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-X01` | Vocabulary / Grammar Review | 5–10 min | spaced retrieval for suitable Knowledge Objects | `PT-17`, `PT-18`, `PT-19` |
| `PM-X02` | Error Remediation | 5–15 min | target a classified recurring error with a new corrective action | `PT-04` plus referenced RemediationPattern |
| `PM-X03` | Adaptive Mixed Set | 10–20 min | interleave eligible targets selected from current learner state | `PT-20`, `PT-21` |
| `PM-X04` | Full IELTS Mock | approximately 165 min for integrated written sections; Speaking may be scheduled separately | whole-test readiness, pacing, stamina, integration | `PT-23`, readiness-only by default |

# Media is a source, not another practice taxonomy

YouTube/owned media does not create a parallel set of practice modes.

Eligible media can instantiate existing modes such as:

- `PM-L01` Dictation;
- `PM-L02` Gist Sprint;
- `PM-S02` Shadowing;
- `PM-S04` Retell/long-turn variants when the prompt is appropriately transformed;
- `PM-X01` vocabulary review extracted from authorized transcript/content.

Media-specific eligibility is owned by `03-media-youtube.md`.

# Evidence-role labels

Each concrete practice activity must expose one of these product labels:

```text
TRAINING_ONLY
EVIDENCE_ELIGIBLE
DIAGNOSTIC
READINESS_ONLY
```

`EVIDENCE_ELIGIBLE` means the attempt may produce an Observation/EvidenceFact if Assessment conditions are satisfied; it never means automatic mastery contribution.

# Practice selection

The product selection pipeline is:

```text
GapEvaluation
  ↓
ActionIntent
  ↓
Learning Mechanism
  ↓
eligible Practice Mode(s)
  ↓
constraints: time, skill focus, accessibility, prior exposure, fatigue
  ↓
concrete Practice Item
```

Wrong answer → mode is forbidden as a direct mapping. The system must first determine whether the problem is ability, prerequisite, evidence, conflict, staleness, scaffold dependence, transfer, fluency, or exam condition.

# Retry policy

A retry must be labeled by purpose:

- **recovery retry** — correct misunderstanding with support;
- **faded retry** — reduce scaffold;
- **retention retry** — retrieve after time;
- **transfer retry** — new item/context testing generalization;
- **re-evidence attempt** — fresh admissible sample for a claim.

A same-item immediate retry is normally recovery evidence, not independent transfer evidence.

# Duration policy

Durations in this catalog are UX defaults for planning and expectation-setting. They do not define learning dosage, mastery, or certification.

The scheduler may shorten or lengthen an activity when the concrete task demands it, but should not silently turn a focused 5-minute mode into an unbounded session.

# Catalog quality rule

The catalog should remain small enough that a learner can understand it and large enough to expose materially different learning actions.

A new mode requires at least one of:

- a distinct learner goal;
- a distinct interaction model;
- a distinct exam task family that cannot be represented clearly by an existing mode;
- a distinct learning/evidence role visible to the learner.

Do not create a new mode merely for a new topic, content source, Band, or generated template.