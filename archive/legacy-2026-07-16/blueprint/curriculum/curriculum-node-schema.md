# Curriculum Node Schema (v1.1)

The canonical unit of a learning pathway. Every curriculum node conforms to this schema ([CR-002](decisions.md)). A node **orchestrates** — it references Skill Leaves and Knowledge Objects by `id`; it never redefines them ([CR-001](decisions.md)). Consumed by `../practice/`, `../assessment/`, `../progress/`, `../review/`. Curriculum is an **orchestration layer only**.

## Fields
| field | type | notes |
|---|---|---|
| `id` | string, unique, stable | `C-B<band>-<nn>` — e.g. `C-B5-03`. |
| `band` | int (3–9) | the band phase this node belongs to. |
| `sequence` | int | order within the band pathway. |
| `skill_leaves` | array of skill-leaf `id`s | leaves this node covers (definitions in [../skills/](../skills/)). |
| `knowledge_objects` | array of knowledge `id`s | knowledge introduced/required here (definitions in [../knowledge/](../knowledge/)). |
| `focus` | string | the node's learning focus — a **synthesis**, not a restatement. |
| `expected_outcomes` | string | what the learner should be able to **do** after the node — a learning **intention**, not a pass/fail condition; **complements** `exit`. |
| `sequencing` | string | rationale: which factors drove this position (prereq / band / cognitive / load / mastery). |
| `exit` | string | mastery required to complete the node — references the leaves' `mastery_criteria` / `mastery_states`. |
| `typical_learning_duration` | string | estimated planning effort (e.g. "2–3 hours"). **Canonical planning property** of the node — not learner runtime data. |
| `estimated_load` | enum | aggregated `low` \| `medium` \| `high`. |

## Orchestration, not redefinition
`skill_leaves` and `knowledge_objects` hold **only `id`s**. `focus`/`expected_outcomes` describe how the referenced objects combine into a teachable step; the objects' definitions remain the single source in skills/ and knowledge/. No parallel representation.

## Sequencing policy (the 5 factors)
Order nodes so that: (1) **prerequisites** precede dependents; (2) **band** progression respected; (3) **cognitive complexity** ascends; (4) **load** balanced; (5) **mastery** gates each node. Each node's `sequencing` states which factors apply.

## `expected_outcomes` vs `exit`
`expected_outcomes` is a **learning intention** (what the learner aims to be able to do). `exit` is the **mastery gate** (the criterion referenced from the leaves). They complement, never replace, each other.

## Versioning & changelog
- **v1.1** (2026-07-16) — added `expected_outcomes` (learning intention; complements `exit`) and `typical_learning_duration` (canonical planning estimate, not runtime).
- **v1.0** (2026-07-16) — initial schema ([CR-002](decisions.md)).
