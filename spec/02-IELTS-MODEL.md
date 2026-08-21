STATUS: CANONICAL
OWNS: external IELTS test structure, variants, section scoring model, official criteria, delivery facts, and exam facts the learning system must respect
DEPENDS_ON: 00-PRODUCT.md
DOES_NOT_OWN: pedagogical skill decomposition, learning-band overlays, curriculum, practice strategy, mastery evidence policy, learner progression

# 02 — IELTS Model

## Purpose

Record the external IELTS reality that constrains the learning system.

This file separates official exam facts from internal pedagogical interpretation. Other specs consume these facts; they must not redefine them.

External facts in this document were rechecked against official IELTS material on **2026-08-22**.

## Standard IELTS variants

The intended complete standard-IELTS construct contains two active variants:

- **IELTS Academic**;
- **IELTS General Training**.

Listening and Speaking are shared between Academic and General Training. Reading and Writing differ by variant.

Shared learning concepts must remain shared. Variant-specific differences are represented only where the exam actually differs.

Academic may be released first, but release ordering is a product-support concern and must not be expressed here as if General Training were a future or optional exam construct.

## Four scored sections

IELTS reports separate band scores for:

1. Listening;
2. Reading;
3. Writing;
4. Speaking.

The overall band score is computed from the four section band scores using IELTS's official averaging and rounding rules. The overall score is an exam result summary; the learning system does not treat it as a synchronized progression gate.

## Delivery baseline — 2026

IELTS announced that from **mid-2026** paper-based IELTS will no longer be offered, with exact rollout timing varying by market. IELTS on computer is therefore the standard delivery baseline for exam-readiness design.

Selected markets may offer a **Writing on Paper** option. This is a delivery/input-mode option, not a different language construct or scoring standard.

Canonical consequences:

- exam-readiness interaction must support computer-delivered timing/navigation as the default;
- handwriting-specific practice is an optional delivery overlay only when relevant to the learner's booked option;
- paper-era answer-transfer mechanics must not be encoded as universal capability requirements;
- a delivery-mode change does not redefine Skill, Band, Assessment, or Progression semantics.

## Test timing baseline

The overall test remains approximately 2 hours 45 minutes, with section timing broadly:

- Listening: approximately 30 minutes;
- Reading: 60 minutes;
- Writing: 60 minutes;
- Speaking: approximately 11–14 minutes.

Delivery details can vary. Learning semantics should not depend on obsolete paper-specific mechanics unless a live delivery option materially changes what the learner must do.

# Listening

## Structure

Listening contains 40 questions in four parts, 10 questions per part. Recordings are heard once.

The progression of contexts is:

- Part 1 — everyday/social conversation;
- Part 2 — everyday/social monologue;
- Part 3 — educational/training conversation;
- Part 4 — academic monologue.

Canonical question families include:

- multiple choice;
- matching;
- plan/map/diagram labelling;
- form/note/table/flow-chart completion;
- sentence completion;
- short answer.

Listening is shared between Academic and General Training.

## Scoring

Each correct answer earns one mark and the raw score out of 40 is converted to the IELTS band scale.

Official IELTS publishes average anchor marks and warns that exact boundaries may vary slightly by test version.

| Listening band | Official average anchor /40 |
|---|---:|
| 5 | 16 |
| 6 | 23 |
| 7 | 30 |
| 8 | 35 |

Half bands and other whole bands are obtained through the live test's conversion, not by treating an internally interpolated table as immutable official truth.

# Reading

Reading contains 40 questions and allows 60 minutes. Academic and General Training are graded on the same 0–9 band scale but use different text characteristics and typical raw-score requirements.

## Shared Reading question families

Both variants use the major Reading interaction families represented by IELTS material, including:

- multiple choice;
- identifying information: True / False / Not Given;
- identifying writer views/claims: Yes / No / Not Given;
- matching information;
- matching headings;
- matching features;
- matching sentence endings;
- sentence completion;
- summary/note/table/flow-chart completion;
- diagram label completion where applicable;
- short answer.

A shared question family does not imply a shared corpus/context distribution or shared raw-score conversion.

## Academic Reading

Academic Reading uses three passages and 60 minutes.

Texts are academically oriented and may increase in complexity. The product must preserve Academic passage characteristics when creating practice/readiness evidence.

Official average Academic anchor marks:

| Academic Reading band | Official average anchor /40 |
|---|---:|
| 5 | 15 |
| 6 | 23 |
| 7 | 30 |
| 8 | 35 |

## General Training Reading

General Training Reading has **three sections of increasing difficulty** and 40 questions in 60 minutes.

External context structure:

- **Section 1** — two or three short texts, or several shorter texts, on everyday topics needed for life in an English-speaking environment, such as notices, advertisements, and timetables;
- **Section 2** — two texts focused on work contexts, such as job descriptions, contracts, staff development, and training material;
- **Section 3** — one longer and more complex text on a topic of general interest, commonly descriptive or instructive and sourced from newspapers, magazines, books, or online resources.

The total GT Reading text length is approximately **2150–2750 words** according to current official format guidance.

Official average General Training anchor marks include:

| GT Reading band | Official average anchor /40 |
|---|---:|
| 4 | 15 |
| 5 | 23 |
| 6 | 30 |
| 7 | 35 |

General Training typically requires more correct answers for the same band because its text characteristics differ from Academic Reading.

Canonical product implication: GT Reading may reuse shared Reading capabilities and question engines, but a GT readiness claim must sample the GT section/context distribution and use the GT score conversion. Academic-only passage exposure is not sufficient GT variant coverage.

# Writing

Writing contains two tasks and both must be completed.

IELTS uses four Writing criteria:

- Task Achievement for Task 1 / Task Response for Task 2;
- Coherence and Cohesion;
- Lexical Resource;
- Grammatical Range and Accuracy.

Criteria are equally weighted within a task. **Task 2 contributes twice as much as Task 1** to the Writing section score.

Detailed learning thresholds derived from these criteria are owned by `05-BANDS.md`.

## Academic Writing

### Academic Task 1

The learner describes or explains visual information. The minimum response is 150 words and the expected planning allocation is roughly 20 minutes.

The external construct requires selection and communication of key information/features and an appropriate overview of the visual information.

### Academic Task 2

The learner responds to a point of view, argument, or problem. The minimum response is 250 words and the expected allocation is roughly 40 minutes.

## General Training Writing

### General Training Task 1

The learner is given a situation and writes a **letter of at least 150 words**, normally in about 20 minutes.

The prompt provides **three bullet points** describing information/functions that must be covered.

The required style may be:

- personal;
- semi-formal;
- formal.

The appropriate style depends on the relationship to the recipient and the communicative purpose.

The task may require the learner to:

- ask for and/or provide general factual information;
- express needs, wants, likes, or dislikes;
- express opinions, views, complaints, or related purposes.

Task Achievement includes how effectively the letter achieves its purpose and covers the required content. Audience, purpose, register/style, organization, and required-point coverage are therefore variant-specific capability requirements and must not be replaced by Academic visual-feature/overview behavior.

### General Training Task 2

The learner writes a semi-formal/neutral discursive essay of at least 250 words in response to a point of view, argument, or problem, normally in about 40 minutes.

The underlying Task-2 learning construct is substantially shared with Academic, while prompt domains and expected framing may differ in concrete content.

# Speaking

Speaking is a recorded face-to-face interview with a certified examiner and is shared across Academic and General Training.

It has three parts:

- Part 1 — interview/questions on familiar personal topics;
- Part 2 — individual long turn on a topic, with one minute to prepare and a target long turn of up to about two minutes;
- Part 3 — extended discussion related to the Part-2 topic.

IELTS Speaking uses four equally weighted criteria:

- Fluency and Coherence;
- Lexical Resource;
- Grammatical Range and Accuracy;
- Pronunciation.

Detailed learning thresholds are owned by `05-BANDS.md`.

# One Skill Retake

One Skill Retake reuses one existing IELTS section; it does not create a fifth skill or a new learning construct.

Current external eligibility/administrative facts include:

- the learner first completes a full IELTS test at a centre offering One Skill Retake;
- the original full test is IELTS on computer;
- only one skill may be retaken once per original test;
- the One Skill Retake must be completed within **60 days** of the original full test;
- availability and organization acceptance may vary by location/receiving institution.

These conditions affect TargetProfile/product guidance and exam-preparation flow. They do not redefine the selected skill's learning standard.

# UKVI boundary

IELTS for UKVI Academic/General Training uses the same language-test construct/results as the corresponding standard variant while adding administrative/security conditions.

Those external conditions may constrain a TargetProfile or booking guidance but do not create a parallel learning curriculum.

# Band scale

IELTS uses a 0–9 band scale with whole and half bands.

For this learning system:

- Bands 0–2 are retained as diagnostic/external boundaries;
- Bands 3–9 are the structured learning range;
- official descriptor/scoring facts remain external truth;
- statements about what should be taught at a band are learning decisions owned outside this file.

# Official fact vs Blueprint inference

For Writing and Speaking, IELTS publishes analytic criteria/descriptors that support direct criterion-level interpretation.

For Listening and Reading, the public scoring model is primarily raw-score conversion rather than a complete public analytic comprehension descriptor per band. Therefore:

- raw-score conversion and test format can be Official Evidence;
- direct implications of the format may be Evidence-Based Interpretation;
- detailed receptive learning abilities assigned to bands are Blueprint Inference and must be labelled accordingly in `05-BANDS.md`.

# External source baseline

Primary official sources for this model include:

- IELTS — Academic test format;
- IELTS — General Training test format and detailed Reading/Writing pages;
- IELTS — scoring in detail;
- IELTS — official sample-test resources;
- IELTS — One Skill Retake guidance;
- IELTS — 2026 test-delivery update announcing the transition away from paper-based IELTS.

If IELTS changes the live exam format, delivery rules, scoring guidance, or variant definitions, this file is the first canonical owner that must be updated; dependent learning/product owners are then reviewed for consequences.
