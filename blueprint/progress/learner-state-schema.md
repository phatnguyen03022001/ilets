# Learner State Schema (v1.0)

The canonical shape of learner runtime state. Per-learner instances conform to this model ([PG-001](decisions.md)); the model itself references canonical objects **by `id` only** and defines no learning content.

## State objects

### `LeafMasteryState` — per Skill Leaf
| field | type | notes |
|---|---|---|
| `leaf_id` | skill-leaf `id` | the leaf this state concerns. |
| `mastery_state` | enum | `not_started` → `practicing` → `emerging` → `mastered` (the fixed enum from [../skills/leaf-schema.md](../skills/leaf-schema.md)). |
| `confidence` | float 0–1 | current confidence the learner meets the leaf's `mastery_criteria`. |
| `evidence_count` | int | independent demonstrations recorded (Assessment Type `AT-05` aggregate). |
| `last_assessed` | timestamp | last assessment event. |
| `evidence_refs` | array | references to assessment evidence (runtime records). |

### `KnowledgeState` — per Knowledge Object
| field | type | notes |
|---|---|---|
| `knowledge_id` | knowledge `id` | the object this state concerns. |
| `state` | enum | `not_acquired` → `learning` → `acquired`. |
| `confidence` | float 0–1 | retrieval/production confidence. |

### `BandCertificationState` — per (band, skill)
| field | type | notes |
|---|---|---|
| `band` | int 3–9 | the band. |
| `skill` | enum | writing \| speaking \| listening \| reading. |
| `status` | enum | `not_started` → `in_progress` → `certified`. |
| `exit_evidence_refs` | array | the `AT-05` portfolio evidence certifying this exit ([BD-002](../bands/decisions.md)). |

### `OverallLearnerState`
| field | type | notes |
|---|---|---|
| `current_band` | per skill + overall | learner's certified band per skill. |
| `leaf_states` | map `leaf_id`→`LeafMasteryState` | mastery across the Skill Graph. |
| `knowledge_states` | map `knowledge_id`→`KnowledgeState` | acquisition across the Knowledge Graph. |
| `certification_history` | array | band certifications over time. |
| `exam_prep_mode` | bool | whether Exam-Preparation is active ([LD-005](../learning/decisions.md)) — diagnostic exposure, non-certifying. |
| `review_queue` | array | due spaced-retrieval items ([../learning/review.md](../learning/review.md)). |

## Runtime, not content
All `*_id` fields reference canonical objects; this schema carries **no learning definitions**. A learner's actual state is a runtime instance conforming to this model.

## Versioning
`schema_version = 1.0`. Breaking changes bump the major version.
