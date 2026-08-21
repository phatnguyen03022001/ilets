TYPE: EVIDENCE
STATUS: NON-CANONICAL
OWNER: spec/02-IELTS-MODEL.md provenance

# Official IELTS Evidence Baseline — 2026-08-22

This is source/provenance evidence only. `spec/02-IELTS-MODEL.md` owns current canonical external-exam semantics.

## Primary official sources rechecked

- IELTS Academic test overview and detailed format — `ielts.org`.
- IELTS General Training overview and detailed Reading/Writing format — `ielts.org`.
- IELTS scoring in detail — `ielts.org`.
- IELTS official sample-test resources — `ielts.org`.
- Updates to IELTS test delivery, published 2026-03-05 — `ielts.org/news-and-insights/updates-to-ielts-test-delivery`.
- IELTS on computer: Writing on Paper — `ielts.org/take-a-test/test-types/ielts-on-computer-writing-on-paper`.
- IELTS Online — `ielts.org/take-a-test/test-types/ielts-academic-test/ielts-online` and the corresponding organisations guidance.
- IELTS One Skill Retake guidance — `ielts.org`.

URLs are evidence pointers, not repository authority. Live external guidance must be rechecked when implementation or release depends on it.

## Confirmed construct facts

### Four scored sections

IELTS reports Listening, Reading, Writing, and Speaking section scores. Academic and General Training share Listening and Speaking; Reading and Writing differ by variant.

### Listening

Current official format confirms:

- 4 parts;
- 10 questions per part;
- 40 questions total;
- recordings heard once;
- Parts 1–2 use everyday/social situations;
- Parts 3–4 use educational/training contexts, with Part 4 an academic monologue.

Official scoring guidance publishes average Listening anchors such as Band 5 = 16/40, Band 6 = 23/40, Band 7 = 30/40, Band 8 = 35/40, while warning that exact conversion can vary slightly by test version.

### Reading

Academic Reading uses three passages in 60 minutes.

General Training Reading uses three sections of increasing difficulty:

- Section 1 — short everyday/social-survival texts;
- Section 2 — workplace texts;
- Section 3 — one longer general-interest text.

Academic and GT share major question interaction families but not corpus/context distribution or raw-score conversion.

Official average Reading anchors include:

| Band | Academic average /40 | GT average /40 |
|---|---:|---:|
| 4 | — | 15 |
| 5 | 15 | 23 |
| 6 | 23 | 30 |
| 7 | 30 | 35 |
| 8 | 35 | — |

These are published average anchors, not an immutable internal conversion table.

### Writing

Academic:

- Task 1 — visual information, minimum 150 words;
- Task 2 — point of view/argument/problem response, minimum 250 words.

General Training:

- Task 1 — situation-based letter, minimum 150 words;
- the prompt supplies three bullet points/content functions;
- appropriate style may be personal, semi-formal, or formal depending on recipient relationship and purpose;
- Task 2 — essay response, minimum 250 words.

Writing assessment uses Task Achievement/Task Response, Coherence and Cohesion, Lexical Resource, and Grammatical Range and Accuracy. Task 2 contributes twice as much as Task 1 to the Writing section score.

### Speaking

Speaking uses three parts and four equally weighted criteria: Fluency and Coherence, Lexical Resource, Grammatical Range and Accuracy, and Pronunciation.

The language construct remains human-interactive. Test-centre Speaking is conducted with a trained examiner; IELTS Online Academic conducts Speaking in real time with a trained examiner through a secure video call.

## Delivery facts relevant to exam readiness

### Test-centre computer baseline

IELTS announced on 2026-03-05 that from mid-2026 it will no longer offer the standard paper-based IELTS test, with exact market transition timing varying. The test construct and result interpretation do not change because of this delivery transition.

Therefore computer delivery is the default current test-centre exam-readiness baseline, subject to market rollout/access arrangements.

### Writing on Paper

In selected countries, IELTS on computer may offer Writing on Paper:

- Listening and Reading remain on computer;
- Writing tasks are shown on screen but responses are handwritten;
- task types, scoring, and time allocation are unchanged;
- it is available for Academic and General Training;
- it is not currently offered for IELTS for UKVI;
- if a Writing One Skill Retake follows a Writing-on-Paper original test, current guidance requires the same delivery mode.

This is a delivery/input-mode overlay, not a different Writing construct.

### IELTS Online

IELTS Online is a remote delivery option for IELTS Academic in supported countries:

- the Academic format, timing, questions, marking criteria, and Band interpretation remain the same;
- Listening, Reading, and Writing are delivered remotely online;
- Speaking is a live video call with a trained IELTS examiner;
- receiving institutions decide whether they accept IELTS Online results;
- IELTS Online is not currently accepted for visa/immigration purposes;
- General Training is not currently offered through IELTS Online.

This is a delivery/acceptance condition, not a third standard IELTS variant.

### One Skill Retake

Current official guidance establishes that One Skill Retake:

- reuses one of Listening, Reading, Writing, or Speaking;
- follows a full eligible IELTS-on-computer test at a participating centre;
- permits one skill retake once per original full test;
- must be completed within 60 days of the original test;
- remains subject to location and receiving-organisation acceptance.

## Evidence boundary

These sources establish external format, delivery, scoring, and administrative facts. They do **not** establish every pedagogical statement in `spec/05-BANDS.md`, especially detailed receptive can-do descriptions by Band. Those learning overlays remain Evidence-Based Interpretation or Blueprint Inference as labelled by the owning specification.

## Update rule

When external IELTS format, delivery, scoring, variant, One Skill Retake, or acceptance guidance changes:

1. recheck the live official source;
2. update `spec/02-IELTS-MODEL.md` first when canonical external truth changes;
3. review affected learner/design/coverage owners;
4. refresh this evidence baseline or replace it with a newer dated provenance record;
5. never preserve a stale external rule merely because it was previously documented.