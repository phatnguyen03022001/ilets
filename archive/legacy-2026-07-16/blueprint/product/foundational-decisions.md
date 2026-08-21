# Foundational Decisions

Foundational decisions that constrain the entire Blueprint. Each records the decision, its rationale, and its structural implications. New foundational decisions are appended here.

---

## FD-001 — Test variant scope: Academic first, General Training extensible
**Date:** 2026-07-15 · **Status:** Decided (Founder) · **Type:** Product scope

**Decision.** v1 of the Blueprint fully covers **Academic**. **General Training** is part of the long-term scope and is added later.

**Constraints (binding).**
- The Blueprint structure, domain models, learning/assessment/curriculum models must allow GT to be added **without restructuring existing artifacts**.
- Shared concepts are **modeled once and reused**. Only variant-specific content (e.g. Academic vs. GT Reading passages; Writing Task 1 graph/diagram vs. letter) is separated.
- Optimize for long-term completeness, maintainability, and educational correctness — **not** for minimizing v1 scope.

**Implications.** Requires a shared-core / variant-overlay structure across `skills/`, `practice/`, `assessment/`. Variant-separation convention to be fixed when the first variant-specific section is authored (tracked in `review/`).

---

## FD-002 — Learner model: L1-agnostic canonical + optional localization
**Date:** 2026-07-15 · **Status:** Decided (Founder) · **Type:** Product scope

**Decision.** The Blueprint is **L1-agnostic**. All Blueprint documents are written in **English** as the canonical SSOT. Learning, assessment, curriculum, and knowledge models are independent of any native language.

**Localization** (Vietnamese first; Chinese, Japanese, etc. later) is an **optional, separate layer** that maps onto the canonical Blueprint **without changing learning requirements**.

**Implications.** A new top-level `localization/` category (justified in [`../README.md`](../README.md)). The knowledge layer (e.g. L1-transfer errors) is expressed generically in canonical form; L1-specific interference is a localization overlay, not a core requirement.

---

## FD-003 — Skill taxonomy: 4 official skills + Vocabulary/Grammar as knowledge
**Date:** 2026-07-15 · **Status:** Decided (Founder) · **Type:** Structure

**Decision.** The four official IELTS skills — **Listening, Reading, Writing, Speaking** — are first-class skills under `skills/`. **Vocabulary** and **Grammar** are **not** separate skills; they are enabling **knowledge** layers under `knowledge/` that feed all four skills.

**Rationale.** Matches official IELTS structure (only four skills are scored) and the `CLAUDE.md` separation of `skills/` vs `knowledge/`.

**Implications.** `skills/` has four sub-areas. `knowledge/` begins with Vocabulary and Grammar; other knowledge layers (e.g. pronunciation/phonology, discourse) are TBD during `knowledge/` discovery.

---

## FD-004 — Repository slate: salvage, then commit clean slate
**Date:** 2026-07-15 · **Status:** Decided (Founder) · **Type:** Process

**Decision.** Salvage still-valid product decisions from the prior (`ssot/`) direction into `product/`, then commit the clean-slate pivot.

**Salvaged (carried into Blueprint).** mission, product principles, band/criteria references (TA/CC/LR/GRA), and domain terms.

**Not salvaged (implementation-era; archived in git history `HEAD`).** The `ssot/` microservice architecture, OpenAPI contract, and ADRs 0034–0045 (auth, retry, sandbox, AI-model selection, DB, Z.ai round-trip). These are out of scope for the Blueprint phase. A small number are *learning-adjacent* (AI evaluation accuracy target, band-score calculation, Writing task-type support, Writing evaluation criteria) and may **inform** — but not constrain — the corresponding Blueprint sections (`assessment/`, `bands/`) when those are authored.

**Implications.** This commit records the pivot. Prior artifacts remain recoverable via `git show HEAD:<path>`.

---

## FD-005 — Branch & freeze lifecycle
**Date:** 2026-07-15 · **Status:** Decided (Founder) · **Type:** Process / governance

**Decision.** All Blueprint work proceeds on the `blueprint` branch. `main` represents only the latest **stable, approved** Blueprint. The `blueprint` branch is merged into `main` only after the Blueprint has:
- completed all planned sections;
- passed consistency review;
- passed gap analysis;
- passed validation, challenge, and stress testing;
- been approved by the Founder;
- been frozen as the project's official SSOT.

Implementation begins only after the frozen Blueprint is merged into `main`.

**Implications.** Discovery/drafting commits land on `blueprint`. The pivot (`1192da7`) is on `blueprint`; `main` still holds the prior direction until the Blueprint is frozen and merged.

---

## FD-006 — Blueprint scope boundaries (delivery / product / implementation-agnostic)
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** The Learning Blueprint defines the **canonical learning domain only**. It is:
- **delivery-agnostic** (no assumption about self-study / tutor-supported / institutional channel);
- **implementation-agnostic** (no architecture, code, tech, or UX);
- **business-model-agnostic** (no commercialization, pricing, or GTM).

**Scope split:**
- **Blueprint Scope** — the canonical learning domain (`blueprint/`). **FROZEN 2026-07-16.**
- **Product Scope** — delivery model, onboarding, entry wedge, UX, commercialization, GTM. **Intentionally open**; consumes the frozen Blueprint. `PRODUCT-002` (delivery model) and `PRODUCT-003` (entry wedge) are **deferred product-level decisions** — *not* Blueprint defects.
- **Implementation Scope** — architecture, code, infrastructure. **Intentionally open**; a later phase.

**Principle.** Future product/implementation decisions **consume** the frozen Blueprint; they must **not modify** it. Only **Blueprint Scope** is frozen today.

---

## Decision index
| ID | Title | Status |
|---|---|---|
| FD-001 | Test variant scope (Academic first, GT extensible) | Decided |
| FD-002 | Learner model (L1-agnostic + localization) | Decided |
| FD-003 | Skill taxonomy (4 skills + knowledge) | Decided |
| FD-004 | Repository slate (salvage + clean commit) | Decided |
| FD-005 | Branch & freeze lifecycle | Decided |

## Section decision logs
Non-foundational decisions live with their section and are indexed here:
- [`learning/decisions.md`](../learning/decisions.md) — LD-001 progression (mastery-gated across bands, adaptive within); LD-002 AI-primary feedback; LD-003 stage-dependent feedback timing; LD-004 prerequisite classification (Required / Recommended / Independent); LD-005 Learning Progression vs Exam Preparation; LD-006 canonical learning phases.
- [`bands/decisions.md`](../bands/decisions.md) — BD-001 bands structure & scope (per-skill docs; Bands 3–9 learning focus, 0–2 boundary-only); BD-002 hierarchical exit criteria (Task → Skill → band progression); BD-003 receptive-skills three-category evidence labeling.
- [`skills/decisions.md`](../skills/decisions.md) — SK-001 decomposition rules (8 leaf criteria + stop); SK-002 atomic + common schema; SK-003 schema v1.1 refinements (`cognitive_level`, `typical_learning_load`, `assessment_strategy`).
- [`knowledge/decisions.md`](../knowledge/decisions.md) — KK-001 Knowledge Graph rules; KK-002 Knowledge Object schema; KK-003 schema v1.1 (`common_misconceptions`, `examples`, atomicity refinement).
- [`curriculum/decisions.md`](../curriculum/decisions.md) — CR-001 orchestration (reference by `id`, sequence by 5 factors, canonical pathway); CR-002 Curriculum Node schema.
- [`practice/decisions.md`](../practice/decisions.md) — PR-001 Practice Taxonomy rules (canonical layer, referenced by `id`, 7 attributes, architectural separation); PR-002 Practice Type schema; PR-003 phase classification.
- [`assessment/decisions.md`](../assessment/decisions.md) — AM-001 Assessment Model rules (canonical layer, independent from Practice, 8 questions, reference by `id`); AM-002 Assessment Type schema; AM-003 calibration defaults (≥0.80 confidence, ≥2 demonstrations — configurable).
- [`progress/decisions.md`](../progress/decisions.md) — PG-001 Learner State Model (state + transitions + decisions; reference by `id`; runtime-vs-model separation); PG-002 per-skill band progression.
