# Progress Decisions

Founder decisions defining the progress section. Indexed from [product/foundational-decisions.md](../product/foundational-decisions.md).

---

## PG-001 — Progress is the canonical Learner State Model (not data storage)
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** `progress/` defines the canonical **Learner State Model** — the specification for learner state, transitions, and decisions. It does **not** store learner data (per-learner runtime instances conform to this model).
- **Defines:** learner state transitions; mastery state transitions; band progression rules; prerequisite enforcement; adaptive scheduling; review scheduling; certification state; learning recommendations.
- **References canonical objects by `id` only** — Band, Curriculum Node, Skill Leaf, Knowledge Object, Practice Type, Assessment Type. It **never redefines** any learning object.
- **Calibration defaults** (≥0.80 confidence, ≥2 demonstrations) are configurable policy values ([AM-003](../assessment/decisions.md)).
- **Runtime-vs-model separation:** the *model* lives here (schema + rules + decision logic); per-learner *runtime instances* (actual state) conform to the model but are not defined here.

**Architectural placement:** Progress = **learner runtime state** — distinct from Knowledge/Skills/Curriculum/Practice/Assessment. Certification is a *consequence of sufficient evidence*, not a single assessment event.

---

## PG-002 — Per-skill band progression (resolves V-HIGH-01)
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** Band progression is **per-skill**. Each skill's `BandCertificationState` advances **independently**: a skill certifies Band N→N+1 when its own [BD-002](../bands/decisions.md) exit criteria are met (`AT-05` portfolio), **regardless of other skills**. This aligns with IELTS (per-section bands; overall = average) and the per-skill exit criteria in [BD-002](../bands/decisions.md).
- The **overall band** = average of the four section bands — **informational, not a progression gate**.
- **Replaces** the prior synchronized "all four skills must certify Band N" rule ([transitions.md](transitions.md) §2), which blocked uneven profiles (e.g., a Band-7 reader gated by Band-5 writing).

**Implication.** Learners progress uneven skills independently (e.g., Reading 7 / Writing 5). Mastery-gating ([LD-001](../learning/decisions.md)) applies **per-skill**. Exam-Preparation ([LD-005](../learning/decisions.md)) remains non-certifying.
