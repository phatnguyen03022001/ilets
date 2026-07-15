# Assessment Taxonomy (v1.0)
*Canonical assessment types ([assessment-type-schema.md](assessment-type-schema.md)), referenced by `id` from Skill Leaves (`assessment_strategy`) + Curriculum Nodes. Independent from Practice ([AM-001](decisions.md)).*

**`AT-01` Criterion-referenced productive performance task** · kind: formative/summative · applies_to: `W-*`, `S-*`
- evidence_produced: a band-level Writing/Speaking performance scored against the official descriptors (TA/TR/CC/LR/GRA; FC/LR/GRA/P).
- strategy: learner produces a full task (essay / Task 1 / spoken turn); AI scores per criterion against band descriptors.
- sufficiency: ≥2 independent demonstrations at band N, sustained, no criterion below N−1 ([evidence-and-confidence.md](evidence-and-confidence.md)).
- confidence_required: AI calibrated confidence ≥ threshold on each criterion.
- ai_reliability: AI scores productive tasks with calibrated confidence — reliable at global/criterion level; less reliable for fine nuance.
- confidence_surfacing: per-criterion calibrated confidence + explicit acknowledgement of AI limits; low-confidence items flagged for optional human review ([LD-002](../learning/decisions.md)).
- aligns_to: Writing/Speaking band exit criteria + the leaves' `mastery_criteria`.

**`AT-02` Objective receptive item set** · kind: formative/summative · applies_to: `L-*`, `R-*`
- evidence_produced: raw score (/40) → band conversion (Listening/Reading).
- strategy: learner answers 40 objective items; auto-scored; converted to band via the official tables.
- sufficiency: band-N raw-score range across ≥2 full timed tests, sustained, no result below N−1.
- confidence_required: effectively deterministic (objective scoring); minor test-version variance noted.
- ai_reliability: automated scoring fully reliable (objective items).
- confidence_surfacing: minimal (objective); note test-version variance.
- aligns_to: Listening/Reading band exit (score conversion) + receptive leaf mastery.

**`AT-03` Knowledge probe** · kind: formative · applies_to: knowledge-dependent leaves
- evidence_produced: recognition/production accuracy for vocabulary/grammar knowledge.
- strategy: retrieval items (recall, gap-fill, judgment, short production).
- sufficiency: ≥ threshold accuracy across ≥2 probes, sustained.
- confidence_required: high (objective recognition); calibrated for free production.
- ai_reliability: AI reliable for recognition; calibrated confidence for free production.
- confidence_surfacing: production items carry a confidence flag.
- aligns_to: knowledge-dependent leaves (the `K-*` prerequisites); supports W-LR/W-GRA/S-LR etc.

**`AT-04` Diagnostic checkpoint** · kind: diagnostic · applies_to: cross
- evidence_produced: a scoped mastery profile across multiple leaves/knowledge (identifies gaps).
- strategy: targeted item set + sampled productive tasks; AI scored + calibrated.
- sufficiency: a snapshot (not mastery certification); locates leaves needing work.
- confidence_required: per-leaf confidence reported; not a pass/fail gate.
- ai_reliability: objective + calibrated productive.
- confidence_surfacing: per-leaf confidence + flagged uncertainty.
- aligns_to: band entry/exit, Exam-Preparation entry.

**`AT-05` Mastery portfolio (accumulated demonstrations)** · kind: formative (cumulative — the certification mechanism) · applies_to: all leaves (per band exit)
- evidence_produced: **sustained** performance evidence over time — the [LD-001](../learning/decisions.md) reliability rule materialized.
- strategy: aggregate ≥2 independent demonstrations (drawn from AT-01/02/03) at band N, sustained, no drop below N−1.
- sufficiency: the LD-001 reliability rule satisfied.
- confidence_required: combined confidence across demonstrations ≥ threshold; low-confidence demonstrations are re-probed.
- ai_reliability: inherits from constituent types.
- confidence_surfacing: aggregate confidence; re-probe low-confidence evidence.
- aligns_to: **Band exit criteria ([BD-002](../bands/decisions.md))** — this is what **certifies** band completion → the progression gate (Learning Progression, **not** Exam Preparation).

**`AT-06` Human / human-verified assessment** · kind: summative (optional) · applies_to: `W-*`, `S-*` (optional)
- evidence_produced: expert-scored productive performance (highest nuance).
- strategy: a human examiner scores Writing/Speaking, or verifies low-confidence AI scores.
- sufficiency: as AT-01, human-judged.
- confidence_required: human judgment (qualitative, not a numeric confidence).
- ai_reliability: N/A (human); used where AI confidence is low.
- confidence_surfacing: N/A; optional escalation ([LD-002](../learning/decisions.md)).
- aligns_to: productive band exit (optional verification); **not** a progression dependency.

**`AT-07` Full mock test** · kind: summative (Exam Readiness — non-certifying) · applies_to: cross
- evidence_produced: a full-test band estimate under exam conditions (all four skills).
- strategy: full IELTS simulation; AI scored + calibrated on productive.
- sufficiency: a **readiness estimate**, not mastery certification.
- confidence_required: reported band estimate + confidence.
- ai_reliability: objective receptive + calibrated productive.
- confidence_surfacing: per-section confidence; Writing/Speaking calibrated.
- aligns_to: **Exam Preparation** ([LD-005](../learning/decisions.md)); does **not** satisfy progression requirements.

---

## Coverage of the 8 questions
Every type declares evidence (`evidence_produced`), strategy, sufficiency, confidence required, AI reliability, confidence surfacing, kind (formative/summative/diagnostic), and alignment — answering the Assessment Model's eight questions ([AM-001](decisions.md)). The cross-cutting rules (sufficiency, confidence threshold, AI reliability, formative-vs-summative) are formalized in [evidence-and-confidence.md](evidence-and-confidence.md).

## Open questions
- [ ] Validate the type set (7) + the evidence/confidence model before binding.
