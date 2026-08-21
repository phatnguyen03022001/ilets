TYPE: EVIDENCE
STATUS: NON-CANONICAL
OWNER: spec/02-IELTS-MODEL.md provenance

# Official IELTS Evidence Baseline — 2026-08-22

This is supporting evidence only. `spec/02-IELTS-MODEL.md` owns current canonical external-exam semantics.

## Sources rechecked

Primary official IELTS pages:

- Academic test overview — `https://ielts.org/take-a-test/test-types/ielts-academic-test`
- Academic format in detail — `https://ielts.org/organisations/ielts-for-organisations/test-types/ielts-academic-test/academic-test-format-in-detail`
- General Training test overview — `https://ielts.org/take-a-test/test-types/ielts-general-training-test`
- General Training sample-test resources — `https://ielts.org/take-a-test/preparation-resources/sample-test-questions/general-training-test`
- IELTS scoring in detail — `https://ielts.org/take-a-test/your-results/ielts-scoring-in-detail`

## Confirmed external facts

### Four sections

IELTS reports Listening, Reading, Writing, and Speaking section scores.

### Variant relationship

Listening and Speaking are the same for Academic and General Training.

Reading and Writing differ between Academic and General Training.

### Academic timing

Official Academic overview states a total test time of about 2 hours 45 minutes, with:

- Listening: about 30 minutes;
- Reading: 60 minutes;
- Writing: 60 minutes;
- Speaking: 11–14 minutes.

Paper/computer answer-transfer mechanics can differ and should not be promoted into invariant learning semantics unless material to the task condition.

### Listening structure

Official detailed format confirms:

- 4 parts;
- 10 questions per part;
- 40 questions total;
- recordings heard once;
- Parts 1–2 use everyday/social situations;
- Parts 3–4 use educational/training contexts, with Part 4 an academic monologue.

### Listening scoring anchors

Official scoring page publishes average marks out of 40:

| Band | Listening average anchor |
|---|---:|
| 5 | 16 |
| 6 | 23 |
| 7 | 30 |
| 8 | 35 |

IELTS explicitly states that precise marks can vary slightly between test versions.

### Reading scoring anchors

Academic Reading official average anchors:

| Band | Academic Reading average anchor |
|---|---:|
| 5 | 15 |
| 6 | 23 |
| 7 | 30 |
| 8 | 35 |

General Training Reading official average anchors:

| Band | GT Reading average anchor |
|---|---:|
| 4 | 15 |
| 5 | 23 |
| 6 | 30 |
| 7 | 35 |

IELTS states that General Training commonly requires more correct answers for a given band because the texts differ in vocabulary/style complexity.

### Writing

Official material confirms two tasks.

Academic:

- Task 1: describe visual information, minimum 150 words;
- Task 2: respond to a point of view, argument, or problem, minimum 250 words.

General Training:

- Task 1: letter/situation response, minimum 150 words;
- Task 2: essay response, minimum 250 words.

Writing assessment uses:

- Task Achievement / Task Response;
- Coherence and Cohesion;
- Lexical Resource;
- Grammatical Range and Accuracy.

Criteria are equally weighted within a task, and Task 2 carries more weight than Task 1 in the section score.

### Speaking

Official material confirms three parts and four equally weighted criteria:

- Fluency and Coherence;
- Lexical Resource;
- Grammatical Range and Accuracy;
- Pronunciation.

Speaking is shared across Academic and General Training.

## Evidence boundary

These pages establish exam format and scoring facts.

They do **not** establish every pedagogical statement in `spec/05-BANDS.md`, especially detailed receptive can-do descriptions by band. Those learning overlays are intentionally labeled Blueprint Inference rather than official IELTS descriptors.

## Update rule

If official IELTS changes format, scoring guidance, or variant behavior:

1. update `spec/02-IELTS-MODEL.md` first;
2. review dependent specs for consequences;
3. refresh or supersede this evidence note;
4. do not silently preserve stale external rules because they were historically canonical.
