# Glossary
Canonical terminology for the Blueprint — the single source for every defined term. Other documents link rather than redefine. Finalized 2026-07-16.

## Architecture — the six layers (responsibilities kept separate)
- **Knowledge** — what must be **known** (vocabulary, grammar, phonology).
- **Skills** — what must be **demonstrated** (the four IELTS skills, decomposed).
- **Curriculum** — **when** learning occurs (sequenced pathways).
- **Practice** — **how** learning is **trained** (practice taxonomy).
- **Assessment** — how **mastery** is **measured** (assessment model).
- **Progress** — learner **runtime state** (state + transitions).
- **Single representation** — each object defined once; referenced by `id` elsewhere.
- **Canonical / runtime separation** — object definitions (Blueprint) vs per-learner runtime instances (`progress/`).

## Canonical objects (each has a stable `id` + schema)
- **Skill Leaf** — atomic skill unit (`W/S/L/R-*`; [../skills/leaf-schema.md](../skills/leaf-schema.md)).
- **Knowledge Object** — atomic knowledge unit (`K-GRA/K-VOC/K-PHON-*`; [../knowledge/knowledge-object-schema.md](../knowledge/knowledge-object-schema.md)).
- **Curriculum Node** — a sequenced step in a band pathway (`C-B<band>-<nn>`).
- **Practice Type** — a reusable training template (`PT-*`).
- **Assessment Type** — a reusable measurement template (`AT-*`).
- **Learner State** — runtime state objects (`LeafMasteryState`, etc.).

## Learning model ([../learning/](../learning/))
- **Band** — IELTS 0–9 proficiency level; Blueprint focus **Bands 3–9** (0–2 boundary-only).
- **Mastery states** — `not_started → practicing → emerging → mastered`.
- **Mastery-gated progression** ([LD-001](../learning/decisions.md)) — advance only on demonstrated mastery.
- **Hierarchical exit criteria** ([BD-002](../bands/decisions.md)) — Task-level → Skill-level → band progression.
- **Prerequisite classification** ([LD-004](../learning/decisions.md)) — `Required` (hard gate) · `Recommended` (soft) · `Independent`.
- **Learning Progression vs Exam Preparation** ([LD-005](../learning/decisions.md)) — certified advancement vs diagnostic higher-band exposure (non-certifying).
- **Canonical learning phases** ([LD-006](../learning/decisions.md)) — Acquisition · Consolidation · Retrieval · Transfer · Fluency · **Exam Readiness**.
- **Exam Readiness vs Exam Preparation** — Exam Readiness is a *learning phase* (LD-006); Exam Preparation is the *non-certifying mode* (LD-005). Distinct.

## Assessment ([../assessment/](../assessment/))
- **Certification** — a consequence of sufficient evidence (`AT-05` portfolio), not a single event.
- **Formative / Summative / Certification** — ongoing-driving-learning / exam-mock / band-gating. Distinct.
- **Calibration defaults** ([AM-003](../assessment/decisions.md)) — minimum mastery confidence **≥0.80**; minimum evidence **≥2** independent demonstrations (configurable policy values).

## Variants ([FD-001](../product/foundational-decisions.md), [GV-002](../review/decisions.md))
- **Academic / General Training (variant)** — the two IELTS modules. Shared content default; variant-specific tagged `academic` | `general-training` | `shared`. Blueprint v1 = Academic; GT added later non-breaking.

## IELTS assessment (official terms)
- **Band Score** — 0–9 scale; half-bands by interpolation. Per-section bands; overall = average.
- **TA/TR/CC/LR/GRA** — Writing criteria (Task Achievement/Response, Coherence & Cohesion, Lexical Resource, Grammatical Range & Accuracy).
- **FC/LR/GRA/Pronunciation** — Speaking criteria.
- **Task 1 / Task 2** — Writing tasks (Task 2 essay shared; Task 1 differs by variant).
