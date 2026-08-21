# Progress Consistency Review
*Verifies that every learner state transition is fully explainable through canonical Blueprint objects ([PG-001](decisions.md)). Per Founder directive (2026-07-16).*

## Transition → canonical-object trace
| transition ([transitions.md](transitions.md)) | depends on (canonical objects / decisions) | explainable? |
|---|---|---|
| §1 mastery (`not_started→…→mastered`) | `leaf_id` ([../skills/](../skills/)), PT-* ([../practice/binding.md](../practice/binding.md)), AT-* ([../assessment/](../assessment/)), [AM-003](../assessment/decisions.md) thresholds | ✅ |
| §2 band progression | `BandCertificationState`, `AT-05`, [BD-002](../bands/decisions.md) exit criteria, [LD-001](../learning/decisions.md) | ✅ |
| §3 prerequisite enforcement | skill/knowledge `requires` edges + [LD-004](../learning/decisions.md) classification | ✅ |
| §4 adaptive scheduling | leaf `practice_item_types`, `PT-17`, `AT-05`, `AT-07`, [LD-005](../learning/decisions.md) | ✅ |
| §5 review scheduling | [../learning/review.md](../learning/review.md) (performance-graded spacing), knowledge/leaf items | ✅ |
| §6 certification | `AT-05`, [AM-003](../assessment/decisions.md), [LD-005](../learning/decisions.md) | ✅ |
| §7 recommendations | Curriculum Nodes, leaf `remediation`, PT-*, [LD-005](../learning/decisions.md) | ✅ |

## Runtime values vs. learning objects
Runtime instance values — `confidence` (0–1), `evidence_count`, `last_assessed`, `evidence_refs`, `review_queue` — are **learner data derived from assessment events** (AT-* evidence). They are **not learning objects**: they conform to the [Learner State schema](learner-state-schema.md) but carry no learning definitions. The transition **rules** reference only canonical objects; these values are computed from events grounded in canonical objects.

## No non-canonical dependencies
No transition depends on learning content defined outside the Blueprint's canonical layers. `progress/` **redefines no** Skill / Knowledge / Curriculum / Practice / Assessment object.

## Conclusion
Every learner state transition is **fully explainable through canonical Blueprint objects**. Runtime values are instance data derived from canonical-grounded assessment evidence. `progress/` is internally consistent and ready for the **final Blueprint validation + freeze**.
