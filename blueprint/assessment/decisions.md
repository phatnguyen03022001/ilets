# Assessment Decisions

Founder decisions defining the assessment section. Indexed from [product/foundational-decisions.md](../product/foundational-decisions.md).

---

## AM-001 — Assessment is a canonical Assessment Model (not tests)
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** `assessment/` defines the canonical **Assessment Model** for every Skill Leaf — a reusable layer, **not** a collection of tests.
- **Independent from Practice** ([PR-001](../practice/decisions.md)): Practice *trains* learning; Assessment *measures demonstrated competence*. The two are separate.
- **Reference canonical objects by `id`** — Skill Leaves, Curriculum Nodes, Practice Types, Knowledge Objects — never duplicate their definitions.
- **Single representation:** an Assessment Type schema ([AM-002](#am-002--assessment-type-schema)) governs every type; leaf `assessment_strategy` slots are populated by reference in [binding.md](binding.md).

**The Assessment Model answers (for every leaf):**
1. what evidence demonstrates mastery;
2. what assessment strategy is appropriate;
3. what evidence is sufficient;
4. what confidence is required before mastery can be declared;
5. what AI can assess reliably;
6. where calibrated confidence / uncertainty must be surfaced;
7. how formative and summative assessments differ;
8. how assessment aligns with the Learning Blueprint and Band exit criteria.

**Alignment.** Assessment conforms to the Learning Blueprint: mastery-gating ([LD-001](../learning/decisions.md)), AI-primary feedback + calibrated confidence ([LD-002](../learning/decisions.md)), feedback timing ([LD-003](../learning/decisions.md)), Learning-Progression vs Exam-Preparation ([LD-005](../learning/decisions.md)), hierarchical exit criteria ([BD-002](../bands/decisions.md)), and Band exit criteria in [../bands/](../bands/).

---

## AM-002 — Assessment Type schema
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** Every assessment type conforms to [assessment-type-schema.md](assessment-type-schema.md) v1.0: `id`, `evidence_produced`, `strategy`, `sufficiency`, `confidence_required`, `ai_reliability`, `confidence_surfacing`, `kind` (formative/summative/diagnostic), `aligns_to`. The evidence-and-confidence rules live in [evidence-and-confidence.md](evidence-and-confidence.md); the type set in [taxonomy.md](taxonomy.md).
