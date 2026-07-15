# Curriculum Node Schema (v1.0)

The canonical unit of a learning pathway. Every curriculum node conforms to this schema ([CR-002](decisions.md)). A node **orchestrates** — it references Skill Leaves and Knowledge Objects by `id`; it never redefines them ([CR-001](decisions.md)). Consumed by `../practice/`, `../assessment/`, `../progress/`, `../review/`.

## Fields
| field | type | notes |
|---|---|---|
| `id` | string, unique, stable | `C-B<band>-<nn>` — e.g. `C-B5-03`. |
| `band` | int (3–9) | the band phase this node belongs to. |
| `sequence` | int | order within the band pathway. |
| `skill_leaves` | array of skill-leaf `id`s | the leaves this node covers (orchestration by reference — definitions live in [../skills/](../skills/)). |
| `knowledge_objects` | array of knowledge `id`s | the knowledge introduced/required here (definitions live in [../knowledge/](../knowledge/)). |
| `focus` | string | the node's learning focus — a **synthesis** of the referenced objects, **not** a restatement of them. |
| `sequencing` | string | rationale: which of the 5 factors drove this position (prereq / band / cognitive / load / mastery). |
| `exit` | string | mastery required to complete the node — references the leaves' `mastery_criteria` / `mastery_states` (not redefined). |
| `estimated_load` | enum | aggregated `low` \| `medium` \| `high` (rolled up from referenced leaves/knowledge). |

## Orchestration, not redefinition
`skill_leaves` and `knowledge_objects` hold **only `id`s**. The node's `focus` describes *how* the referenced objects combine into a teachable step; the objects' full definitions remain the single source in skills/ and knowledge/. No parallel representation ([CR-001](decisions.md)).

## Sequencing policy (the 5 factors)
Within a band pathway, order nodes so that: (1) **prerequisites** precede dependents; (2) **band** progression is respected; (3) **cognitive complexity** ascends; (4) **load** is balanced; (5) **mastery** gates (a node's exit must be achievable before the next). Each node's `sequencing` states which factors apply.

## Versioning
`schema_version = 1.0`. Breaking changes bump the major version and migrate all nodes.
