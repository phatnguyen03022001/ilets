# Curriculum

## Purpose
Transform the completed [Skill Graph](../skills/) + [Knowledge Graph](../knowledge/) into coherent **Band 3→9 learning pathways**. Orchestrate canonical objects by `id` — never redefine them ([CR-001](decisions.md)).

## Model ([CR-001](decisions.md), [CR-002](decisions.md))
A pathway of **Curriculum Nodes** ([curriculum-node-schema.md](curriculum-node-schema.md)), organized into **Band phases (3–9)**, mastery-gated between bands ([LD-001](../learning/decisions.md)). Each node references Skill Leaves + Knowledge Objects **by `id`** and is ordered by the five factors: prerequisite dependencies, band progression, cognitive complexity, estimated load, mastery requirements.

## Sequencing policy
Within a band phase: **knowledge before the skill that needs it**; prerequisites before dependents; cognitive/load ascending; mastery gates each node. The recommended order is the Blueprint default — adaptive sequencing is runtime, in `../progress/`.

## Status
**Complete (2026-07-16).** All Band 3–9 pathways (44 nodes, schema v1.1) + [cross-band consistency review](consistency-review.md). Calibration deferred. Pending Founder review — the gate before `../practice/`.

## Structure
- `band-3.md` ✅ `band-4.md` ✅ `band-5.md` ✅ `band-6.md` ✅ `band-7.md` ✅ `band-8.md` ✅ `band-9.md` ✅ — per-band pathway nodes (44 nodes, schema v1.1).
- [consistency-review.md](consistency-review.md) — cross-band consistency review (calibration deferred).
- [curriculum-node-schema.md](curriculum-node-schema.md) — canonical Curriculum Node schema v1.1.
- [decisions.md](decisions.md) — CR-001/002/003.

## Dependencies
- **Consumes:** [../skills/](../skills/) (leaves + prereqs) + [../knowledge/](../knowledge/) (objects + `requires`) + [../bands/](../bands/) (exit criteria).
- **Feeds (canonical pathway):** [../practice/](../practice/), [../assessment/](../assessment/), [../progress/](../progress/), [../review/](../review/) — they consume, not replace.
