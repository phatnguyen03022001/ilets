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
