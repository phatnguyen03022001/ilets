STATUS: CANONICAL
OWNS: initial user-facing practice-mode catalog, mode identity/count, default duration envelopes, mode-to-learning-role mapping, variant/delivery-aware packaging, and catalog change policy
DEPENDS_ON: ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 01-skill-features.md
DOES_NOT_OWN: live IELTS facts, Learning Mechanism definitions, Practice Type semantics, evidence sufficiency, Skill/Band thresholds, concrete item content, planner eligibility/ranking, or scheduling implementation

# Practice Catalog

## Purpose

Define the concrete practice choices the learner sees. Canonical Practice Types/Mechanisms remain learning truth; this file packages them into understandable product modes.

The initial catalog contains **28 modes**:

```text
Listening   6
Reading     6
Writing     6
Speaking    6
Shared      4
----------
Total      28
```

Variant/delivery differences resolve inside an existing mode when the learner interaction remains materially the same.

# Listening — 6

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-L01` | Dictation | 5–12 min | detail discrimination, segmentation, spelling | `PT-13`, `PT-19`, Retrieval/Contrast |
| `PM-L02` | Gist Sprint | 3–6 min | main idea/discourse listening | `PT-12`, Fluency/Transfer |
| `PM-L03` | Detail & Completion | 5–10 min | explicit detail, completion, word-limit control | `PT-13`, `PT-19` |
| `PM-L04` | Paraphrase & Distractor | 5–10 min | paraphrase recognition/distractor rejection | `PT-16`, Contrast/Self-explanation |
| `PM-L05` | Map / Diagram | 5–10 min | spatial language/structured detail | `PT-13` |
| `PM-L06` | Timed Section | 10–35 min | integrated listening under exam-like conditions | `PT-15`, Exam Readiness |

# Reading — 6

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-R01` | Skim Sprint | 3–5 min | topic/purpose/paragraph gist | `PT-12`, Fluency |
| `PM-R02` | Scan & Detail Hunt | 4–8 min | explicit information location | `PT-12`, `PT-13` |
| `PM-R03` | T/F/NG + Y/N/NG | 6–12 min | evidence/stance classification | `PT-13`, `PT-16`, Contrast |
| `PM-R04` | Headings & Structure | 6–12 min | paragraph function/text organization | `PT-13`, Transfer |
| `PM-R05` | Paraphrase / Inference / Stance | 6–12 min | inferential reasoning beyond keyword matching | `PT-13`, `PT-16`, Self-explanation |
| `PM-R06` | Timed Reading | 15–60 min by scope | variant-aware passage/section/full Reading performance | `PT-15`, Exam Readiness |

## Reading variant packaging

`PM-R06` identifies selected variant and scope.

Concrete GT content uses the stable GT Reading Content Context IDs owned by `../spec/10-CONTENT-MODEL.md`.

A focused activity may target one context. A whole-Reading readiness activity must use a complete applicable variant configuration and normal Assessment policy.

# Writing — 6

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-W01` | Prompt & Plan | 5–10 min | variant-aware task/purpose/content planning | `PT-01`, Guided production |
| `PM-W02` | Sentence & Grammar Lab | 5–10 min | structural accuracy/range/correction | `PT-02`, `PT-04`, Controlled production |
| `PM-W03` | Lexical & Paraphrase Lab | 5–10 min | collocation, precision, paraphrase, spelling/register | `PT-03`, `PT-18` |
| `PM-W04` | Paragraph & Cohesion Builder | 8–15 min | organization/reference/cohesion/development | `PT-01`, `PT-05` |
| `PM-W05` | Guided Draft & Redraft | 15–30 min | scaffolded variant-aware production + revision | `PT-01`, `PT-05`, Scaffold fading |
| `PM-W06` | Timed Writing | about 20 min Task 1 / 40 min Task 2 | independent full-task performance under selected target conditions | `PT-06`, Exam Readiness |

## Writing variant packaging

Task-1 mode instantiation resolves:

```text
Academic → W-TA-01 / W-TA-02 / W-TA-03
GT       → W-GT1-01 / W-GT1-02 / W-GT1-03 + shared Writing capability
```

A GT planning activity covers recipient/purpose/required content/register rather than Academic visual-overview behavior.

## Writing delivery packaging

For exam-readiness uses, `PM-W06` records the selected supported delivery/input condition when material, such as typed computer response or an eligible handwriting-rehearsal path.

Delivery packaging changes interaction conditions, not Writing criteria or Band truth.

# Speaking — 6

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-S01` | Pronunciation Contrast | 5–10 min | phoneme/stress/intonation discrimination/production | `PT-07`, Contrast/Controlled production |
| `PM-S02` | Shadowing | 5–12 min | rhythm/connected speech/stress/fluency imitation | `PT-08`, Fluency rehearsal |
| `PM-S03` | Part-1 Quick Response | 5–10 min | spontaneous short response | `PT-10`, Retrieval/Fluency |
| `PM-S04` | Part-2 Long Turn | 5–10 min | planning + sustained long turn | `PT-09`, Transfer/Fluency |
| `PM-S05` | Part-3 Discussion | 8–15 min | abstract explanation/comparison/justification | `PT-10`, Transfer |
| `PM-S06` | Full Speaking Mock | 11–14 min | integrated Parts 1–3 readiness | `PT-11`, Exam Readiness |

# Shared — 4

| ID | Mode | Default duration | Learning role | Typical canonical backing |
|---|---|---:|---|---|
| `PM-X01` | Vocabulary / Grammar Review | 5–10 min | spaced retrieval for suitable Knowledge Objects | `PT-17`, `PT-18`, `PT-19` |
| `PM-X02` | Error Remediation | 5–15 min | classified error → fresh corrective action | `PT-04` + RemediationPattern |
| `PM-X03` | Adaptive Mixed Set | 10–20 min | interleave eligible targets from current learner state | `PT-20`, `PT-21` |
| `PM-X04` | Full IELTS Mock | about 150 min for Listening+Reading+Writing; Speaking follows the applicable separate/scheduled test configuration | variant/delivery-aware whole-test readiness | `PT-23`; primary purpose `READINESS`; evidence candidacy is configured separately |

The complete external IELTS timing is owned by `../spec/02-IELTS-MODEL.md`. Catalog duration is a planning envelope, not external-exam authority.

`PM-X04` resolves variant before Reading/Writing Task 1 are instantiated. It also records supported delivery interaction where material. A mixed variant is invalid for a normal full-test readiness claim.

# Media source rule

Media does not create another practice taxonomy. Eligible media instantiates suitable existing modes under `03-media-youtube.md`.

# Activity purpose vs evidence candidacy

A concrete activity has two orthogonal product dimensions.

## Primary activity purpose

Exactly one primary purpose describes **why the activity is being scheduled/presented**:

```text
TRAINING
DIAGNOSTIC
ASSESSMENT
READINESS
```

Meanings:

- `TRAINING` — acquisition, consolidation, retrieval, transfer, fluency, remediation, or other learning work is primary;
- `DIAGNOSTIC` — reducing decision-relevant uncertainty/classifying what should be sampled or addressed next is primary;
- `ASSESSMENT` — focused measurement/re-measurement of a scoped capability/claim is primary without necessarily simulating whole target/exam conditions;
- `READINESS` — performance under target-like integrated/timed/exam conditions is primary.

This is a product/activity-purpose classification, not an Assessment admission or certification judgment.

## Evidence candidacy

Separately, the configured activity declares whether its resulting Observation is allowed to be considered by Assessment:

```text
NOT_EVIDENCE_CANDIDATE
ASSESSMENT_MAY_ADMIT
```

`ASSESSMENT_MAY_ADMIT` is only a pre-attempt candidacy designation. After the learner performs the activity, `08-ASSESSMENT.md` still decides claim-scoped eligibility from actual task fit, assistance/scaffolding, exposure/retry history, evaluator quality, provenance, and other material conditions.

The two dimensions are orthogonal. Any purpose may be configured with either candidacy value when semantically justified; the candidacy choice must be made before observing the learner result.

Examples:

- `TRAINING + NOT_EVIDENCE_CANDIDATE` — guided acquisition/recovery work;
- `TRAINING + ASSESSMENT_MAY_ADMIT` — independent practice deliberately configured as a potential evidence source;
- `DIAGNOSTIC + ASSESSMENT_MAY_ADMIT` — a diagnostic sample that may also satisfy normal Assessment eligibility;
- `ASSESSMENT + ASSESSMENT_MAY_ADMIT` — focused re-evidence/mastery sampling;
- `READINESS + ASSESSMENT_MAY_ADMIT` — eligible timed task/mock sample;
- `READINESS + NOT_EVIDENCE_CANDIDATE` — exam-familiarization simulation whose configuration is not suitable for formal evidence.

Purpose never implies evidence admission by itself.

There is no pre-attempt `CERTIFICATION_CONTRIBUTING` role. Certification contribution exists only after Assessment admits evidence and evaluates the applicable EvidenceRequirement.

# Selection boundary

Hard eligibility and ranking are downstream product Planner concerns owned by `04-application-flows.md`.

This catalog answers **which user-facing mode packages a semantically valid selected learning action**.

A mode must not independently:

- bypass prerequisites;
- select the wrong variant/delivery configuration;
- reinterpret GapEvaluation;
- create certification;
- convert a CoverageGap into learner weakness.

# Retry purpose

Retries are labelled by purpose:

- **recovery retry** — correct misunderstanding with support;
- **faded retry** — reduce scaffold;
- **retention retry** — retrieve after delay;
- **transfer retry** — new material/context for generalization;
- **re-evidence attempt** — fresh admissible sample for a claim.

Immediate same-item retry is normally recovery, not independent transfer evidence.

# Duration boundary

Catalog durations are UX/planning defaults. They do not define dosage, mastery, or certification.

# Catalog change rule

A new mode requires at least one materially distinct learner-visible reason:

- different goal;
- different interaction model;
- task family not representable clearly by existing interaction;
- distinct learning/evidence role.

Do not create a mode solely for a topic, source, Band, variant label, delivery label, or generated template.