# Assessment

## Purpose
Define the canonical **Assessment Model** for every Skill Leaf — *how mastery is measured*. A reusable layer, **not** tests; independent from Practice ([AM-001](decisions.md)).

## Responsibility (architectural separation)
Assessment = **how mastery is measured** — distinct from Knowledge (known), Skills (demonstrated), Curriculum (when), Practice (trained), Progress (runtime). Kept strictly separate.

## Model ([AM-001](decisions.md), [AM-002](decisions.md))
Every assessment type conforms to [assessment-type-schema.md](assessment-type-schema.md) v1.0 and answers: evidence produced, strategy, sufficiency, confidence required, AI reliability, confidence surfacing, kind (formative/summative/diagnostic), and alignment. Evidence-and-confidence rules in [evidence-and-confidence.md](evidence-and-confidence.md); types in [taxonomy.md](taxonomy.md).

## Status
**Complete (2026-07-16).** Model v1.0 (7 types + evidence/confidence) + [binding.md](binding.md) (4 maps) + [coverage-review.md](coverage-review.md). Pending Founder review.

## Structure
- [taxonomy.md](taxonomy.md) — canonical Assessment Types (v1.0).
- [evidence-and-confidence.md](evidence-and-confidence.md) — mastery evidence, sufficiency, confidence thresholds, AI reliability, calibrated-confidence surfacing, formative vs summative.
- [binding.md](binding.md) ✅ — leaf→strategy + exit→evidence + node→culmination + practice→assessment maps.
- [coverage-review.md](coverage-review.md) ✅ — coverage verification.
- [assessment-type-schema.md](assessment-type-schema.md) — Assessment Type schema v1.0.
- [decisions.md](decisions.md) — AM-001/002.

## Dependencies
- **Consumes:** [../bands/](../bands/) (exit criteria), [../skills/](../skills/) (leaf `mastery_criteria`), [../learning/](../learning/) (LD-001/002/003/005, assessment.md ~70/30), [../practice/](../practice/) (training that prepares for assessment).
- **Feeds:** `assessment_strategy` on Skill Leaves; consumed by `../progress/` (mastery state, gating) and `../review/`.
