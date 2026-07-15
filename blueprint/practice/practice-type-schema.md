# Practice Type Schema (v1.0)

The canonical unit of the Practice Taxonomy. Every practice type conforms to this schema ([PR-002](decisions.md)). Referenced by `id` from Skill Leaves (`practice_item_types`) and Curriculum Nodes — never embedded or redefined ([PR-001](decisions.md)).

## Fields
| field | type | notes |
|---|---|---|
| `id` | string, unique, stable | `PT-<nn>` — e.g. `PT-03`. |
| `name` | string | descriptive label. |
| `objective_served` | string | the learning objective this type serves. |
| `cognitive_operations` | array (Bloom) | `remember` \| `understand` \| `apply` \| `analyze` \| `evaluate` \| `create`. |
| `supports` | array | the Skill Leaf **categories/components** this type suits (e.g. `W-GRA-*`, receptive `*-COMP-*`). Characterization, not an exhaustive binding — leaves select types via `practice_item_types`. |
| `reinforces` | array | the Knowledge **domains** it reinforces (e.g. `K-GRA`, `K-VOC`). |
| `when_used` | string | where it fits in the curriculum (acquisition phase / consolidation / pre-exam / exit). |
| `purpose` | enum | `acquisition` \| `consolidation` \| `transfer` \| `exam_readiness` (aligns with [LD-003](../learning/decisions.md) stages + exam readiness per [LD-005](../learning/decisions.md)). |
| `mode` | enum | `individual` \| `mixed` \| `adaptive` \| `timed` \| `diagnostic`. |
| `applies_to` | array | skills/domains: writing, speaking, listening, reading, knowledge, cross. |
| `feedback_model` | string | how feedback is delivered (AI-primary per [LD-002](../learning/decisions.md); timing per [LD-003](../learning/decisions.md)). |

## Type vs. item
A practice **type** is a reusable template (this schema). A practice **item/instance** is a concrete exercise generated against a type for a specific leaf — items are runtime/content, not defined here.

## Referencing (no duplication)
Leaves and curriculum nodes hold **only practice-type `id`s**. The type's definition is the single source in [taxonomy.md](taxonomy.md).

## Versioning
`schema_version = 1.0`. Breaking changes bump the major version.
