# Practice Type Schema (v1.1)

The canonical unit of the Practice Taxonomy. Every practice type conforms to this schema ([PR-002](decisions.md)). Referenced by `id` from Skill Leaves (`practice_item_types`) and Curriculum Nodes — never embedded or redefined ([PR-001](decisions.md)).

## Fields
| field | type | notes |
|---|---|---|
| `id` | string, unique, stable | `PT-<nn>` — e.g. `PT-03`. |
| `name` | string | descriptive label. |
| `objective_served` | string | the learning objective this type serves. |
| `cognitive_operations` | array (Bloom) | `remember` \| `understand` \| `apply` \| `analyze` \| `evaluate` \| `create`. |
| `primary_phase` | enum | the **one** canonical phase this type primarily serves (see below). |
| `phases` | array | all canonical phases it supports (`primary_phase` ∈ `phases`). |
| `supports` | array | Skill Leaf categories/components it suits (e.g. `W-GRA-*`). Characterization — leaves select types via `practice_item_types`. |
| `reinforces` | array | Knowledge domains it reinforces (e.g. `K-GRA`). |
| `when_used` | string | where it fits in the curriculum (acquisition / consolidation / pre-exam / exit). |
| `mode` | enum | `individual` \| `mixed` \| `adaptive` \| `timed` \| `diagnostic`. |
| `applies_to` | array | skills/domains: writing, speaking, listening, reading, knowledge, cross. |
| `feedback_model` | string | how feedback is delivered (AI-primary per [LD-002](../learning/decisions.md); timing per [LD-003](../learning/decisions.md)). |

## Canonical learning phases ([LD-006](../learning/decisions.md))
`acquisition` · `consolidation` · `retrieval` · `transfer` · `fluency` · `exam_readiness`. Every type declares a `primary_phase` (one) and the full `phases` set it supports.

## Type vs. item
A practice **type** is a reusable template (this schema). A practice **item/instance** is a concrete exercise generated against a type for a specific leaf — items are runtime/content, not defined here.

## Referencing (no duplication)
Leaves and curriculum nodes hold **only practice-type `id`s**. The type's definition is the single source in [taxonomy.md](taxonomy.md). Authoritative leaf↔type and node↔type bindings live in [binding.md](binding.md) (the population of those slots by reference).

## Versioning & changelog
- **v1.1** (2026-07-16) — replaced `purpose` with `primary_phase` + `phases` classified against the canonical learning phases ([LD-006](../learning/decisions.md), [PR-003](decisions.md)).
- **v1.0** (2026-07-16) — initial schema ([PR-002](decisions.md)).
