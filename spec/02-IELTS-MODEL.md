STATUS: CANONICAL
OWNS: external IELTS test structure, variants, section scoring model, official criteria, and exam facts the learning system must respect
DEPENDS_ON: 00-PRODUCT.md
DOES_NOT_OWN: pedagogical skill decomposition, learning-band overlays, curriculum, practice strategy, mastery evidence policy, learner progression

# 02 — IELTS Model

## Purpose

Record the external IELTS reality that constrains the learning system.

This file separates official exam facts from internal pedagogical interpretation. Other specs consume these facts; they must not redefine them.

External facts in this document were rechecked against official IELTS material during the 2026-08-22 document refactor.

## Test variants

IELTS has two relevant test variants for this Blueprint:

- **Academic** — fully specified by the initial active Blueprint;
- **General Training** — future variant overlay.

Listening and Speaking are shared between Academic and General Training. Reading and Writing differ by variant.

Shared learning concepts must remain shared. Variant-specific differences are represented only where the exam actually differs.

## Four scored sections

IELTS reports separate band scores for:

1. Listening;
2. Reading;
3. Writing;
4. Speaking.

The overall band score is computed from the four section band scores using IELTS's official averaging and rounding rules. The overall score is an exam result summary; the learning system does not treat it as a progression gate.

## Academic test timing baseline

The official Academic format is approximately 2 hours 45 minutes in total, with section timing broadly:

- Listening: approximately 30 minutes, plus answer-transfer time where applicable to the test mode;
- Reading: 60 minutes;
- Writing: 60 minutes;
- Speaking: approximately 11–14 minutes.

Delivery details can vary by test mode; learning semantics should not depend on paper-specific mechanics unless the mechanic affects demonstrated capability.

## Listening

### Structure

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

### Scoring

Each correct answer earns one mark and the raw score out of 40 is converted to the IELTS band scale.

Official IELTS publishes average anchor marks and explicitly warns that exact boundaries may vary slightly by test version.

| Listening band | Official average anchor /40 |
|---|---:|
| 5 | 16 |
| 6 | 23 |
| 7 | 30 |
| 8 | 35 |

Half bands and other whole bands are obtained through the live test's conversion, not by treating an internally interpolated table as immutable official truth.

## Reading

Reading contains 40 questions. Academic and General Training are graded on the same band scale but use different text characteristics and typical raw-score requirements.

### Academic Reading

Academic Reading uses three passages and 60 minutes.

Question families include:

- multiple choice;
- identifying information: True / False / Not Given;
- identifying writer views/claims: Yes / No / Not Given;
- matching information;
- matching headings;
- matching features;
- matching sentence endings;
- sentence/note/summary/table/flow-chart/diagram completion;
- short answer.

Official average Academic anchor marks:

| Academic Reading band | Official average anchor /40 |
|---|---:|
| 5 | 15 |
| 6 | 23 |
| 7 | 30 |
| 8 | 35 |

### General Training Reading

General Training usually requires more correct answers for the same band because its texts differ in vocabulary and style complexity.

Official average General Training anchor marks published by IELTS include:

| GT Reading band | Official average anchor /40 |
|---|---:|
| 4 | 15 |
| 5 | 23 |
| 6 | 30 |
| 7 | 35 |

The active curriculum remains Academic-first. General Training Reading becomes a variant overlay rather than a new shared-core model.

## Writing

Writing contains two tasks and both must be completed.

### Academic Writing

- Task 1: describe or explain visual information; minimum 150 words; roughly 20 minutes is the official planning expectation.
- Task 2: respond to a point of view, argument, or problem; minimum 250 words; roughly 40 minutes.

### General Training Writing

- Task 1: respond to a situation in letter form; minimum 150 words.
- Task 2: essay response to a point of view, argument, or problem; minimum 250 words.

### Writing assessment criteria

IELTS Writing uses four criteria:

- Task Achievement for Task 1 / Task Response for Task 2;
- Coherence and Cohesion;
- Lexical Resource;
- Grammatical Range and Accuracy.

Criteria are equally weighted within a task. Task 2 carries more weight than Task 1 in the Writing section score.

The detailed learning thresholds derived from these criteria are owned by `05-BANDS.md`, not here.

## Speaking

Speaking is a recorded face-to-face interview with a certified examiner and has three parts:

- Part 1 — interview/questions on familiar personal topics;
- Part 2 — individual long turn on a topic;
- Part 3 — extended discussion related to the Part-2 topic.

Speaking is shared across Academic and General Training.

### Speaking assessment criteria

IELTS Speaking uses four equally weighted criteria:

- Fluency and Coherence;
- Lexical Resource;
- Grammatical Range and Accuracy;
- Pronunciation.

Detailed learning thresholds are owned by `05-BANDS.md`.

## Band scale

IELTS uses a 0–9 band scale with whole and half bands.

For this learning Blueprint:

- Bands 0–2 are retained as diagnostic/external boundaries;
- Bands 3–9 are the structured learning range;
- official descriptor/scoring facts remain external truth;
- any statement about what should be *taught* at a band is a Blueprint learning decision and therefore belongs outside this file.

## Official fact vs Blueprint inference

For Writing and Speaking, IELTS publishes analytic criteria/descriptors that support direct criterion-level interpretation.

For Listening and Reading, the public scoring model is primarily raw-score conversion rather than a complete public analytic comprehension descriptor per band. Therefore:

- raw-score conversion and test format can be Official Evidence;
- direct implications of the format may be Evidence-Based Interpretation;
- detailed receptive learning abilities assigned to bands are Blueprint Inference and must be labeled accordingly in `05-BANDS.md`.

## External source baseline

Primary official sources for this model:

- IELTS — Academic test format: `https://ielts.org/take-a-test/test-types/ielts-academic-test`
- IELTS — Academic format in detail: `https://ielts.org/organisations/ielts-for-organisations/test-types/ielts-academic-test/academic-test-format-in-detail`
- IELTS — General Training test: `https://ielts.org/take-a-test/test-types/ielts-general-training-test`
- IELTS — scoring in detail: `https://ielts.org/take-a-test/your-results/ielts-scoring-in-detail`
- IELTS — official sample-test resources for Academic and General Training.

If IELTS changes the live exam format or scoring guidance, this file is the first canonical owner that must be updated; dependent learning specs are then reviewed for consequences.
