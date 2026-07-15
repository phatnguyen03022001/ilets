# Assessment Type Schema (v1.0)

The canonical unit of the Assessment Model. Every assessment type conforms to this schema ([AM-002](decisions.md)). Referenced by `id` from Skill Leaves (`assessment_strategy`) and Curriculum Nodes — never embedded or redefined ([AM-001](decisions.md)). Independent from Practice.

## Fields
| field | type | notes |
|---|---|---|
| `id` | string, unique, stable | `AT-<nn>` — e.g. `AT-02`. |
| `name` | string | descriptive label. |
| `evidence_produced` | string | what evidence of mastery this type yields. |
| `strategy` | string | how it is administered and scored. |
| `sufficiency` | string | what constitutes enough evidence (references the reliability rule in [evidence-and-confidence.md](evidence-and-confidence.md)). |
| `confidence_required` | string | confidence threshold to declare mastery (see [evidence-and-confidence.md](evidence-and-confidence.md)). |
| `ai_reliability` | string | what AI can assess reliably via this type ([LD-002](../learning/decisions.md)). |
| `confidence_surfacing` | string | where calibrated confidence / uncertainty is surfaced. |
| `kind` | enum | `formative` \| `summative` \| `diagnostic`. |
| `aligns_to` | string | Band exit criteria / leaf `mastery_criteria` it supports. |
| `applies_to` | array | skills/leaves suited (e.g. `W-*`, receptive). |

## Type vs. instance
An assessment **type** is a reusable template (this schema). An assessment **instance** (a specific test/event run for a learner) is runtime, in `../progress/` — not defined here.

## Referencing (no duplication)
Leaves and nodes hold **only assessment-type `id`s**. Type definitions are the single source in [taxonomy.md](taxonomy.md); bindings in [binding.md](binding.md).

## Independence from Practice
Practice Types ([../practice/](../practice/)) *train*; Assessment Types *measure*. A Practice Type prepares learners for an assessment pathway (mapped in [binding.md](binding.md)), but the two layers never redefine each other.

## Versioning
`schema_version = 1.0`. Breaking changes bump the major version.
