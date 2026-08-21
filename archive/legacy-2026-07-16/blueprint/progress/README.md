# Progress

## Purpose
Define the canonical **Learner State Model** — learner runtime state, transitions, and decisions. **Not data storage** ([PG-001](decisions.md)): per-learner instances conform to this model.

## Responsibility (architectural separation)
Progress = **learner runtime state**. Certification is a *consequence of sufficient evidence*, not a single assessment event. Distinct from Knowledge/Skills/Curriculum/Practice/Assessment.

## Model ([PG-001](decisions.md))
- [learner-state-schema.md](learner-state-schema.md) — state objects (`LeafMasteryState`, `KnowledgeState`, `BandCertificationState`, `OverallLearnerState`).
- [transitions.md](transitions.md) — mastery/band/prerequisite/review/certification transitions + adaptive scheduling + recommendations.
- References canonical objects **by `id` only** (Band, Curriculum Node, Skill Leaf, Knowledge Object, Practice Type, Assessment Type). Never redefines a learning object.

## Status
**Complete (2026-07-16).** Model v1.0 (state schema + transitions) + [consistency-review.md](consistency-review.md). Pending Founder review → final Blueprint validation + freeze.

## Structure
- [learner-state-schema.md](learner-state-schema.md) — Learner State schema v1.0.
- [transitions.md](transitions.md) — transition rules + decisions.
- [decisions.md](decisions.md) — PG-001.
- [consistency-review.md](consistency-review.md) ✅ — transition traceability review.

## Dependencies
- **Consumes:** all canonical layers ([../bands/](../bands/), [../skills/](../skills/), [../knowledge/](../knowledge/), [../curriculum/](../curriculum/), [../practice/](../practice/), [../assessment/](../assessment/)) + learning decisions ([LD-001](../learning/decisions.md)/[004](../learning/decisions.md)/[005](../learning/decisions.md), [AM-003](../assessment/decisions.md) calibration defaults).
- **Feeds:** the runtime learner model; certification gating; learning recommendations; consumed by the final validation.
