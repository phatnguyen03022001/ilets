# Curriculum Decisions

Founder decisions defining the curriculum section. Indexed from [product/foundational-decisions.md](../product/foundational-decisions.md).

---

## CR-001 — Curriculum orchestrates canonical objects (not a lesson list)
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** `curriculum/` transforms the completed [Skill Graph](../skills/) + [Knowledge Graph](../knowledge/) into coherent **learning pathways**. It **orchestrates** existing canonical objects — it does **not** redefine them.
- Every curriculum node **references Skill Leaves and Knowledge Objects by `id`**; it never duplicates their definitions.
- **Sequence by:** (1) prerequisite dependencies, (2) band progression, (3) cognitive complexity, (4) estimated learning load, (5) mastery requirements.
- Curriculum is the **canonical learning pathway**; `../practice/`, `../assessment/`, `../progress/`, and `../review/` **consume** it — they do not replace it.

**Structural rules.**
- Pathway is organized into **Band phases (3–9)**, mastery-gated between bands ([LD-001](../learning/decisions.md)); within a band, a **recommended order** (adaptive sequencing is runtime, in `../progress/`).
- **Knowledge before the skill that requires it**; prerequisites before dependents; cognitive complexity and load generally ascending within a phase.
- Each band phase exits at that band's exit criteria ([../bands/](../bands/), [BD-002](../bands/decisions.md)).

**Implication.** A Curriculum Node schema ([CR-002](#cr-002--curriculum-node-schema)) governs every node; nodes reference `id`s only.

---

## CR-003 — Curriculum Node schema v1.1 refinements
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** Schema refined to **v1.1** ([curriculum-node-schema.md](curriculum-node-schema.md)):
1. Add **`expected_outcomes`** — what the learner should be able to do after the node; a learning **intention** (not pass/fail), **complementing** `exit`.
2. Add **`typical_learning_duration`** — estimated planning effort (e.g. "2–3 hours"); a **canonical planning property** of the node, **not** learner runtime data.

**Reaffirmed:** curriculum is an **orchestration layer only** — it never duplicates Skill/Knowledge definitions; nodes reference objects by `id`.

**Cross-band review gate.** After all Bands 3–9 are implemented, a cross-band consistency review (prerequisite consistency, workload balance, cognitive progression, mastery progression, duplicate/missing paths, full Band→Curriculum→Skill→Knowledge traceability) is required **before** proceeding to `../practice/`.
