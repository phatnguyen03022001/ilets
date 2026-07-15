# Skill Leaf Schema (v1.0)

The canonical **learning object**. Every skill leaf conforms to this schema. `../curriculum/`, `../practice/`, `../assessment/`, AI tutoring, and personalization all **reference the leaf by `id`** and read the fields they need — they do **not** create parallel representations of the skill ([SK-002](decisions.md)).

## Fields

### Identification
| field | type | notes |
|---|---|---|
| `id` | string, unique, stable | e.g. `W-GRA-03`. `<SKILL>-<COMPONENT>-<nn>`. |
| `name` | string | human-readable label. |
| `skill` | enum | `writing` \| `speaking` \| `listening` \| `reading`. |
| `component` | string | official assessment criterion / sub-area (e.g. "Grammatical Range & Accuracy"). |
| `schema_version` | string | `1.0`. |

### Definition
| field | type | notes |
|---|---|---|
| `objective` | string | **one atomic** learning objective. If a node holds >1 independent objective → not a leaf → decompose ([SK-001](decisions.md)). |
| `bands` | range/array (3–9) | band(s) the leaf belongs to (introduced → refined). |
| `traces_to` | array | official IELTS requirement(s): descriptor + band citation (from `../bands/`). |

### Learning graph
| field | type | notes |
|---|---|---|
| `prerequisites` | array of `id` / `K-*` | leaf IDs and/or `../knowledge/` items that must precede. Empty if none. |
| `dependents` | array (derived later) | leaves that require this one. |

### Mastery & assessment
| field | type | notes |
|---|---|---|
| `mastery_criteria` | string | observable, criterion-referenced criterion that demonstrates mastery. |
| `assessment_evidence` | method (← `../assessment/`) | how mastery is measured (item type / observable performance). Populated by assessment/. |
| `mastery_states` | enum (fixed) | `not_started → practicing → emerging → mastered`. Defined **once** here; used by `../progress/` for personalization. Not repeated per leaf. |

### Practice & AI tutoring
| field | type | notes |
|---|---|---|
| `practice_item_types` | array (← `../practice/`) | item types targeting this leaf. Populated by practice/. |
| `common_errors` | array | frequent errors / misconceptions (feeds AI-tutoring feedback + remediation). |
| `remediation` | string / array | independent remediation path when not mastered. |

### Leaf invariant (the stop condition)
| field | type | notes |
|---|---|---|
| `independence` | attestation | the node is independently **teachable, practiceable, assessable, AND remediable**. If any fails → not a leaf → decompose ([SK-001](decisions.md), refined by [SK-002](decisions.md)). |

---

## Consumer → fields read
| consumer | reads | (also populates) |
|---|---|---|
| `../curriculum/` | id, prerequisites, bands, component | — |
| `../practice/` | id, objective, mastery_criteria | `practice_item_types` |
| `../assessment/` | id, mastery_criteria | `assessment_evidence` |
| AI tutoring | objective, common_errors, remediation, mastery_states | — |
| personalization / `../progress/` | id, mastery_states, bands, prerequisites | learner state (runtime, not on leaf) |

## Population policy
- **Now (skills/ phase):** id, name, skill, component, objective, bands, traces_to, prerequisites, mastery_criteria, common_errors, remediation, independence.
- **Later (owned by the leaf, populated by the consuming section):** `practice_item_types` (← practice/), `assessment_evidence` (← assessment/), `dependents` (derived). These stay **on the leaf object** — no parallel representation is ever created.

## Runtime vs. schema
This schema defines the leaf's **intrinsic** properties. A learner's *current state* on a leaf (mastery level, practice history) is **runtime data** held in `../progress/` (the learner model), referencing the leaf by `id` — it is not part of the leaf definition.

## Versioning
Schema is versioned (`schema_version`). Prefer backward-compatible changes; breaking changes bump the major version and migrate all leaves.
