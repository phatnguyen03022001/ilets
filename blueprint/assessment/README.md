# Assessment

## Purpose
Define the canonical **Assessment Model** for every Skill Leaf — *how mastery is measured*. A reusable layer, **not** tests; independent from Practice ([AM-001](decisions.md)).

## Responsibility (architectural separation)
Assessment = **how mastery is measured** — distinct from Knowledge (known), Skills (demonstrated), Curriculum (when), Practice (trained), Progress (runtime). Kept strictly separate.

## Model ([AM-001](decisions.md), [AM-002](decisions.md))
Every assessment type conforms to [assessment-type-schema.md](assessment-type-schema.md) v1.0 and answers: evidence produced, strategy, sufficiency, confidence required, AI reliability, confidence surfacing, kind (formative/summative/diagnostic), and alignment. Evidence-and-confidence rules in [evidence-and-confidence.md](evidence-and-confidence.md); types in [taxonomy.md](taxonomy.md).

## Status
**Drafted (2026-07-16).** Model v1.0 (7 assessment types + evidence/confidence model). Next: bind — populate each Skill Leaf's `assessment_strategy` slot + map band-exit→evidence, node→culmination, practice-type→assessment-pathway, then coverage review.

## Structure
- [taxonomy.md](taxonomy.md) — canonical Assessment Types (v1.0).
- [evidence-and-confidence.md](evidence-and-confidence.md) — what demonstrates mastery, sufficiency, confidence thresholds, AI reliability, calibrated-confidence surfacing, formative vs summative.
- [binding.md](binding.md) — leaf→type + exit→evidence + node→culmination + practice→assessment maps (pending).
- [coverage-review.md](coverage-review.md) — coverage verification (pending).
- [assessment-type-schema.md](assessment-type-schema.md) — Assessment Type schema v1.0.
- [decisions.md](decisions.md) — AM-001/002.

## Dependencies
- **Consumes:** [../bands/](../bands/) (exit criteria), [../skills/](../skills/) (leaf `mastery_criteria`), [../learning/](../learning/) (LD-001/002/003/005, assessment.md ~70/30), [../practice/](../practice/) (training that prepares for assessment).
- **Feeds:** `assessment_strategy` on Skill Leaves; consumed by `../progress/` (mastery state, gating) and `../review/`.
