STATUS: CANONICAL
OWNS: external IELTS test structure, standard variants, delivery modes, section/overall scoring facts, official task/question-family identity, official assessment criteria, and exam/administrative facts the learning system must respect
DEPENDS_ON: 00-PRODUCT.md
DOES_NOT_OWN: pedagogical skill decomposition, learning-band overlays, curriculum, practice strategy, evidence sufficiency, learner progression, content-instance identity, product-support policy, or UI/runtime behavior

# 02 — IELTS Model

## Purpose

Record the external IELTS reality that constrains the learning system.

This file answers **what IELTS externally is**. It does not decide how the product teaches, plans, admits evidence, or declares support.

External facts were rechecked against official IELTS material on **2026-08-22**. Supporting source provenance is recorded in `../evidence/2026-08-22-ielts-official-baseline.md`.

# Standard variants

The intended complete standard-IELTS construct contains two variants:

- **IELTS Academic**;
- **IELTS General Training**.

Listening and Speaking are shared between them. Reading and Writing differ where the external exam differs.

A delivery mode is not a third variant.

# Four scored sections and overall score

IELTS reports separate Band scores for:

1. Listening;
2. Reading;
3. Writing;
4. Speaking.

The overall Band is the arithmetic mean of the four section Bands rounded to the nearest whole or half Band. Current official rounding rules include:

- an average ending in `.25` rounds up to the next half Band;
- an average ending in `.75` rounds up to the next whole Band.

The overall score is an external test-result summary. It is not itself evidence that every underlying learning capability is equally strong.

# Stable external task/question-family identity

The repository assigns stable IDs to official IELTS task/question families so product coverage can be checked without turning exam UI categories into extra Skill Leaves.

These IDs identify an **external exam family**, not a pedagogical capability, Practice Type, product feature, or wire enum.

Rules:

1. a family ID remains stable while the corresponding external family remains materially the same;
2. Academic and General Training share a family ID when IELTS uses the same interaction family;
3. variant-specific task families remain distinct where the external construct differs;
4. a Skill Leaf may support several family IDs and one family may require several Skill Leaves;
5. content/coverage tooling must track family identity separately from Skill identity;
6. if IELTS materially changes its official task-family taxonomy, this owner changes first and downstream coverage is re-evaluated.

## Listening family registry

| ID | Official family |
|---|---|
| `IELTS-L-QF-01` | Multiple choice |
| `IELTS-L-QF-02` | Matching |
| `IELTS-L-QF-03` | Plan / map / diagram labelling |
| `IELTS-L-QF-04` | Form / note / table / flow-chart / summary completion |
| `IELTS-L-QF-05` | Sentence completion |
| `IELTS-L-QF-06` | Short-answer questions |

## Reading family registry

These family identities are shared by Academic and General Training; variant content/context and scoring still differ.

| ID | Official family |
|---|---|
| `IELTS-R-QF-01` | Multiple choice |
| `IELTS-R-QF-02` | Identifying information — True / False / Not Given |
| `IELTS-R-QF-03` | Identifying writer views/claims — Yes / No / Not Given |
| `IELTS-R-QF-04` | Matching information |
| `IELTS-R-QF-05` | Matching headings |
| `IELTS-R-QF-06` | Matching features |
| `IELTS-R-QF-07` | Matching sentence endings |
| `IELTS-R-QF-08` | Sentence completion |
| `IELTS-R-QF-09` | Summary / note / table / flow-chart completion |
| `IELTS-R-QF-10` | Diagram label completion |
| `IELTS-R-QF-11` | Short-answer questions |

## Writing task-family registry

| ID | Official task family | Variant |
|---|---|---|
| `IELTS-W-A-T1` | Academic Writing Task 1 — visual information | Academic |
| `IELTS-W-GT-T1` | General Training Writing Task 1 — letter | General Training |
| `IELTS-W-T2` | Writing Task 2 — essay response | shared learning core; concrete prompt remains variant-scoped |

Academic Task-1 visual presentation may include graphs/charts/tables, diagrams/processes, maps/plans, or combinations. Those stimulus-presentation classes are content-instance coverage dimensions, not additional scored Writing tasks.

## Speaking part-family registry

| ID | Official part |
|---|---|
| `IELTS-S-P1` | Speaking Part 1 — interview/familiar topics |
| `IELTS-S-P2` | Speaking Part 2 — individual long turn |
| `IELTS-S-P3` | Speaking Part 3 — extended discussion |

A whole Speaking Band claim remains holistic; these IDs exist so content/readiness coverage cannot omit a part silently.

# Delivery modes — 2026 external baseline

Delivery changes how the learner performs the test, not what language proficiency the Band scale means.

## Test-centre computer

IELTS announced on **2026-03-05** that from mid-2026 the standard paper-based test will no longer be offered, with exact transition timing varying by market. Test-centre computer delivery is therefore the default current delivery baseline.

For standard computer delivery:

- Listening, Reading, and Writing are completed through the computer interface;
- Speaking remains a human examiner interaction;
- the move to computer does not change the skills assessed, test construct, or interpretation of results.

## Test-centre computer with Writing on Paper

Selected markets may offer **Writing on Paper** as an option within IELTS on computer:

- Listening and Reading remain computer-delivered;
- Writing tasks are displayed on screen but answers are handwritten on the official answer sheet;
- Writing task types, timing, scoring, and criteria are unchanged;
- the option applies to Academic and General Training where offered;
- it is not currently offered for IELTS for UKVI;
- current guidance requires a Writing One Skill Retake to use the same Writing delivery mode as the original eligible test.

Writing on Paper is an input/delivery overlay, not a different Writing construct.

## IELTS Online Academic

IELTS Online is a remote delivery option for **IELTS Academic only** in supported countries.

Current external conditions include:

- the Academic test format, timing, questions, marking criteria, and Band interpretation remain equivalent to the Academic construct;
- Listening, Reading, and Writing are completed remotely through the online test platform;
- Speaking is conducted in real time by a trained IELTS examiner through a secure video call;
- receiving organisations decide whether they accept IELTS Online results;
- IELTS Online is not currently accepted for visa/immigration purposes;
- General Training is not currently offered as IELTS Online.

Acceptance/purpose constraints are external target conditions, not new Skill or Band definitions.

# Timing baseline

The complete test is approximately **2 hours 45 minutes** in total. Section timing is broadly:

- Listening — approximately 30 minutes;
- Reading — 60 minutes;
- Writing — 60 minutes;
- Speaking — approximately 11–14 minutes.

Delivery-specific interaction mechanics may vary and must not be confused with language-construct requirements.

# Listening

## Structure

Listening contains 40 questions in four parts, 10 questions per part. Recordings are heard once.

The context progression is:

- Part 1 — everyday/social conversation;
- Part 2 — everyday/social monologue;
- Part 3 — educational/training conversation;
- Part 4 — academic monologue.

Official question families are the six `IELTS-L-QF-*` identities defined above.

Listening is shared between Academic and General Training.

## Scoring

Each correct answer earns one mark. The raw score out of 40 is converted to the IELTS Band scale and reported in whole/half Bands.

Current official average anchors include:

| Listening Band | Average anchor /40 |
|---|---:|
| 5 | 16 |
| 6 | 23 |
| 7 | 30 |
| 8 | 35 |

IELTS states that the precise mark needed may vary slightly between test versions. These anchors therefore are not an immutable universal conversion table.

# Reading

Reading contains 40 questions and allows 60 minutes. Academic and General Training use the same Band scale but differ in text/context characteristics and typical raw-score requirements.

Official interaction families are the eleven shared `IELTS-R-QF-*` identities defined above.

Shared interaction families do not imply identical text distributions or raw-score conversions.

## Academic Reading

Academic Reading uses three passages. Texts are academically oriented and may increase in complexity.

Current official average anchors include:

| Academic Reading Band | Average anchor /40 |
|---|---:|
| 5 | 15 |
| 6 | 23 |
| 7 | 30 |
| 8 | 35 |

## General Training Reading

General Training Reading has three sections of increasing difficulty:

- **Section 1** — short texts on everyday topics needed for life in an English-speaking environment;
- **Section 2** — workplace-related texts;
- **Section 3** — one longer, more complex text on a topic of general interest.

Current official guidance describes a total text length of approximately **2150–2750 words**.

Current official average anchors include:

| GT Reading Band | Average anchor /40 |
|---|---:|
| 4 | 15 |
| 5 | 23 |
| 6 | 30 |
| 7 | 35 |

IELTS notes that General Training commonly requires more correct answers for a given Band than Academic Reading because the text characteristics differ.

# Writing

Writing contains two tasks and allows 60 minutes.

IELTS uses four Writing criteria:

- Task Achievement for Task 1 / Task Response for Task 2;
- Coherence and Cohesion;
- Lexical Resource;
- Grammatical Range and Accuracy.

The criteria are equally weighted within each task. Task 2 contributes more weight to the Writing section result than Task 1; current IELTS guidance describes Task 2 as carrying twice the weight of Task 1.

## Academic Writing

### Task 1 — `IELTS-W-A-T1`

The learner describes or explains visual information. Minimum response length is 150 words; the expected time allocation is roughly 20 minutes.

### Task 2 — `IELTS-W-T2`

The learner responds to a point of view, argument, or problem. Minimum response length is 250 words; the expected allocation is roughly 40 minutes.

## General Training Writing

### Task 1 — `IELTS-W-GT-T1`

The learner is given a situation and writes a letter of at least 150 words, normally in about 20 minutes.

The prompt supplies three bullet points/content functions to cover. Appropriate style may be:

- personal;
- semi-formal;
- formal.

Style depends on the relationship to the recipient and communicative purpose. The task may involve requesting/providing information, expressing needs or preferences, complaints, opinions, or related practical purposes.

### Task 2 — `IELTS-W-T2`

The learner writes an essay of at least 250 words in response to a point of view, argument, or problem, normally in about 40 minutes.

Task 2 substantially shares the same scored Writing criteria with Academic Task 2, while concrete prompt framing may differ.

# Speaking

Speaking is a recorded, human-examiner interaction shared across Academic and General Training.

Its three parts correspond to `IELTS-S-P1`, `IELTS-S-P2`, and `IELTS-S-P3`.

IELTS Speaking uses four equally weighted criteria:

- Fluency and Coherence;
- Lexical Resource;
- Grammatical Range and Accuracy;
- Pronunciation.

Delivery may be in person at a test centre or, for IELTS Online Academic, a secure live video call with a trained examiner. The underlying Speaking construct and criteria remain the same.

# One Skill Retake

One Skill Retake reuses one existing IELTS section; it does not create a fifth skill or a new learning construct.

Current eligibility/administrative facts include:

- the learner first completes a full eligible IELTS-on-computer test at a participating centre;
- only one skill may be retaken once per original full test;
- the retake must be completed within **60 days** of the original full test;
- availability varies by location;
- receiving organisations may apply their own acceptance requirements;
- current UKVI acceptance is available for Academic and General Training computer tests in selected locations;
- where Writing on Paper is used and a Writing One Skill Retake is available, current guidance requires the same Writing delivery mode as the original eligible test.

# UKVI boundary

IELTS for UKVI Academic/General Training uses the corresponding language-test construct while adding administrative/security/test-location conditions for UK immigration purposes.

Those conditions do not create a parallel learning curriculum or a different Band definition.

# Band scale

IELTS uses Bands 0–9 with whole and half Band reporting.

For this repository:

- the external 0–9 scale remains IELTS truth;
- Bands 3–9 are the chosen structured learning-program range, owned as a learning decision outside this file;
- detailed teaching expectations by Band belong to `05-BANDS.md`, not here.

# Official fact vs learning interpretation

For Writing and Speaking, IELTS publishes analytic criteria/descriptors that support criterion-level interpretation.

For Listening and Reading, public scoring is primarily raw-score conversion rather than a complete public analytic comprehension descriptor for every Band.

Therefore downstream learning documents must distinguish:

- **Official Evidence** — directly established external fact/descriptor/scoring guidance;
- **Evidence-Based Interpretation** — close educational interpretation;
- **Blueprint Inference** — pedagogical detail beyond explicit IELTS publication.

# Change rule

If IELTS changes a live test variant, delivery option, scoring rule, task/question family, section format, One Skill Retake condition, or other external fact relevant to learning/product behavior:

1. update this external-truth owner first;
2. refresh the dated evidence provenance;
3. review affected Skill/Band/Curriculum/Assessment/Content/design owners;
4. preserve stable family IDs when semantics remain materially the same;
5. do not preserve an obsolete external fact merely because downstream implementation already encoded it.