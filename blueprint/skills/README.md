# Skills

## Purpose
Decompose every competency from [../bands/](../bands/) into teachable, assessable, traceable skills — the foundation for `../curriculum/`, `../practice/`, `../assessment/`, and adaptive learning. Per Founder directive: decompose until **every band requirement maps to a complete learning path**.

## Decomposition model ([SK-001](decisions.md), [SK-002](decisions.md))
A **skill hierarchy**: skill → component → … → **leaf**. Decomposition keeps the **official assessment criteria as the primary axis** and stops only when a node is a single **atomic** objective **and** is independently teachable / practiceable / assessable / **remediable** ([SK-002](decisions.md)). Optimize for atomic skills, not a target count.

## Leaf record (canonical schema v1.1)
Every leaf conforms to the **Skill Leaf Schema** ([leaf-schema.md](leaf-schema.md) v1.1) — the single learning object referenced by `../curriculum/`, `../practice/`, `../assessment/`, AI tutoring, and personalization **by `id`**, with **no parallel representations** ([SK-002](decisions.md), [SK-003](decisions.md)). Core fields: `id`, `objective`, `cognitive_level`, `typical_learning_load`, `bands`, `traces_to`, `prerequisites`, `mastery_criteria`, `common_errors`, `remediation`, `independence`. Consumer fields (`practice_item_types`, `assessment_strategy`) are owned by the leaf and populated later by the consuming section. **Strict separation** between Blueprint definition and learner runtime state (`../progress/`) is a core rule.

## Scope
- The **four official skills**: writing, speaking, listening, reading — one doc each (mirrors [../bands/](../bands/)).
- Vocabulary & Grammar live in [../knowledge/](../knowledge/) ([FD-003](../product/foundational-decisions.md)); skill leaves concerning them reference the knowledge/ item as a prerequisite/content source.

## ID scheme
`<SKILL>-<COMPONENT>-<nn>` — e.g., `W-CC-03` = Writing, Coherence & Cohesion, leaf 03. (`W`=Writing, `S`=Speaking, `L`=Listening, `R`=Reading.)

## Status
**Approved & locked (2026-07-16).** Leaf schema v1.1 ([leaf-schema.md](leaf-schema.md)) — canonical learning object. `writing.md` decomposition approved; retro-fitting to v1.1, then Speaking / Listening / Reading.

## Structure
- `writing.md` / `speaking.md` / `listening.md` / `reading.md` — per-skill decompositions (each leaf conforms to [leaf-schema.md](leaf-schema.md) v1.1).
- [leaf-schema.md](leaf-schema.md) — canonical Skill Leaf Schema v1.1.
- [decisions.md](decisions.md) — SK-001 decomposition rules; SK-002 atomic + common schema; SK-003 schema v1.1 refinements.

## Dependencies
- **Consumes:** [../bands/](../bands/) (competencies + exit criteria).
- **Feeds:** [../curriculum/](../curriculum/) (sequencing), [../practice/](../practice/) (item design), [../assessment/](../assessment/) (what is assessed), [../knowledge/](../knowledge/) (content prerequisites).
