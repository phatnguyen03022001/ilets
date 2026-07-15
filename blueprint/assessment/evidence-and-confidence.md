# Evidence & Confidence Model
*How mastery is evidenced, judged sufficient, and certified — independent from Practice ([AM-001](decisions.md)). Formalizes the cross-cutting rules referenced by every Assessment Type ([taxonomy.md](taxonomy.md)).*

## 1. What evidence demonstrates mastery
Mastery = **criterion-referenced demonstrated competence** against the leaf's `mastery_criteria` and the band descriptors in [../bands/](../bands/). Not norm-referenced; not time spent; not a single percentage (except objective receptive items, where raw-score→band conversion *is* the criterion). Evidence is observable performance meeting the band's criteria.

## 2. Sufficiency (reliability rule — [LD-001](../learning/decisions.md))
Mastery is demonstrated **reliably**, not once:
- **Productive** (Writing/Speaking): ≥2 independent, time-boxed demonstrations at band N, sustained over a short interval, with **no criterion dropping below band N−1**.
- **Receptive** (Listening/Reading): band-N raw-score range achieved across ≥2 full timed tests, sustained, no result below N−1.
- **Knowledge**: ≥ threshold accuracy across ≥2 probes, sustained.
A single demonstration is insufficient (guards against false positives).

## 3. Confidence required before declaring mastery
Mastery is declared only when the system is **High-confidence** the learner meets the criterion.
- **Proposed threshold: ≥0.80 confidence** (calibratable) that the learner meets the leaf's `mastery_criteria` at the target band.
- For AI-scored productive tasks, the **calibrated confidence** ([LD-002](../learning/decisions.md)) must reach the threshold; below it → **re-probe** (collect more evidence) or optional human verification (`AT-06`).
- Mastery is **never** declared on low-confidence evidence.

## 4. What AI assesses reliably ([LD-002](../learning/decisions.md))
- **Objective receptive items (`AT-02`):** fully reliable (deterministic scoring).
- **Knowledge recognition (`AT-03`):** reliable; free production carries calibrated confidence.
- **Productive Writing/Speaking (`AT-01`):** AI scores with calibrated confidence — reliable at global/criterion-band level; **less reliable for fine nuance** (subtle cohesion, advanced idiomatic precision).

## 5. Where calibrated confidence / uncertainty is surfaced ([LD-002](../learning/decisions.md))
- Productive tasks output **per-criterion calibrated confidence** + an explicit acknowledgement of AI limits.
- **Low-confidence evidence is flagged** and either re-probed or routed to optional human verification (`AT-06`); it cannot certify mastery on its own.
- Learner-facing: confidence is surfaced where it affects decisions (e.g., "AI is moderately confident this is Band 6 Writing; a human check is available").

## 6. Formative vs summative (and certification)
Three distinct concepts — do not conflate:
- **Formative** (dominant, ~70/30 per [assessment.md](../learning/assessment.md)): ongoing, low-stakes, drives learning; uses `AT-01/02/03/04`.
- **Summative (Exam Readiness):** band-approximating mock (`AT-07`) under exam conditions; estimates readiness; **non-certifying** ([LD-005](../learning/decisions.md)) — does not satisfy progression.
- **Mastery certification:** `AT-05` (portfolio) — reliability-rule-satisfied evidence that **certifies** band completion and gates progression (Learning Progression). Distinct from summative: a learner may pass a mock yet not be certified (insufficient sustained evidence), or be certified yet under-perform on a given mock day.

## 7. Alignment
[LD-001](../learning/decisions.md) mastery-gating (certification via `AT-05`); [LD-002](../learning/decisions.md) AI-primary + calibrated confidence; [LD-003](../learning/decisions.md) feedback timing by phase; [LD-005](../learning/decisions.md) Exam-Preparation non-certifying; [BD-002](../bands/decisions.md) hierarchical exit criteria (Task → Skill → band progression); Band exit criteria in [../bands/](../bands/).

## Proposed defaults (open to Founder confirmation)
- Confidence threshold for mastery: **≥0.80** (High).
- Receptive sufficiency: band-N range across **≥2** full tests.
- Productive sufficiency: **≥2** independent demonstrations, no criterion below N−1.
