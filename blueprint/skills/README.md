# Skills

## Purpose
Decompose every competency from [../bands/](../bands/) into teachable, assessable, traceable skills — the foundation for `../curriculum/`, `../practice/`, `../assessment/`, and adaptive learning. Per Founder directive: decompose until **every band requirement maps to a complete learning path**.

## Decomposition model ([SK-001](decisions.md))
A **skill hierarchy**: skill → component → … → **leaf**. A node is a leaf only when it satisfies all 8 leaf criteria. Decomposition **stops** when further subdivision no longer improves teaching / practice / assessment / personalization.

## Leaf record (every leaf carries)
- **Objective** — a single learning objective.
- **Mastery** — explicit criterion (feeds [../learning/mastery.md](../learning/mastery.md) + `../assessment/`).
- **Prerequisites** — other skill IDs and/or `../knowledge/` items (if any).
- **Bands** — which band(s) it belongs to (introduced → refined).
- **Traces to** — the official IELTS requirement (cited from [../bands/](../bands/) descriptors).
- *(Leaf invariant, verified per leaf: independently teachable / practiceable / assessable.)*

## Scope
- The **four official skills**: writing, speaking, listening, reading — one doc each (mirrors [../bands/](../bands/)).
- Vocabulary & Grammar live in [../knowledge/](../knowledge/) ([FD-003](../product/foundational-decisions.md)); skill leaves concerning them reference the knowledge/ item as a prerequisite/content source.

## ID scheme
`<SKILL>-<COMPONENT>-<nn>` — e.g., `W-CC-03` = Writing, Coherence & Cohesion, leaf 03. (`W`=Writing, `S`=Speaking, `L`=Listening, `R`=Reading.)

## Status
**In progress.** `writing.md` drafted as the granularity template — pending Founder validation of the decomposition depth + leaf record format before replicating to Speaking / Listening / Reading.

## Structure
- `writing.md` / `speaking.md` / `listening.md` / `reading.md` — per-skill decompositions.
- [decisions.md](decisions.md) — SK-001 decomposition rules.

## Dependencies
- **Consumes:** [../bands/](../bands/) (competencies + exit criteria).
- **Feeds:** [../curriculum/](../curriculum/) (sequencing), [../practice/](../practice/) (item design), [../assessment/](../assessment/) (what is assessed), [../knowledge/](../knowledge/) (content prerequisites).
