# Skill Leaf Schema (v1.1)

The canonical **learning object**. Every skill leaf conforms to this schema. `../curriculum/`, `../practice/`, `../assessment/`, AI tutoring, and personalization all **reference the leaf by `id`** and read the fields they need — they do **not** create parallel representations of the skill ([SK-002](decisions.md)).

> **Core architectural rules (Founder, 2026-07-16):** (1) **strict separation** between Blueprint definition (intrinsic, here) and learner runtime state (`../progress/`); (2) **single representation** — one leaf object referenced by `id`, never duplicated.

## Fields

### Identification
| field | type | notes |
|---|---|---|
| `id` | string, unique, stable | e.g. `W-GRA-03`. `<SKILL>-<COMPONENT>-<nn>`. |
| `name` | string | human-readable label. |
| `skill` | enum | `writing` \| `speaking` \| `listening` \| `reading`. |
| `component` | string | official assessment criterion / sub-area (e.g. "Grammatical Range & Accuracy"). |
| `schema_version` | string | `1.1`. |

### Definition
| field | type | notes |
|---|---|---|
| `objective` | string | **one atomic** learning objective. If a node holds >1 independent objective → not a leaf → decompose ([SK-001](decisions.md)). |
| `cognitive_level` | enum (Bloom) | `remember` \| `understand` \| `apply` \| `analyze` \| `evaluate` \| `create` — the cognitive operation the skill requires. |
| `typical_learning_load` | enum | `low` \| `medium` \| `high` — relative expected effort / learning time (heuristic; calibrated empirically later). |
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
| `assessment_strategy` | method (← `../assessment/`) | **how** mastery is assessed (item type / observable task). Blueprint-defined (intrinsic). *Learner evidence* (actual results/submissions) is **runtime** data in `../progress/`, not here. |
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
| `../curriculum/` | id, prerequisites, bands, component, cognitive_level, typical_learning_load | — |
| `../practice/` | id, objective, mastery_criteria, cognitive_level | `practice_item_types` |
| `../assessment/` | id, mastery_criteria | `assessment_strategy` |
| AI tutoring | objective, cognitive_level, common_errors, remediation, mastery_states | — |
| personalization / `../progress/` | id, mastery_states, bands, prerequisites, typical_learning_load | learner state (runtime, not on leaf) |

## Population policy
- **Now (skills/ phase):** id, name, skill, component, objective, cognitive_level, typical_learning_load, bands, traces_to, prerequisites, mastery_criteria, common_errors, remediation, independence.
- **Later (owned by the leaf, populated by the consuming section):** `practice_item_types` (← practice/), `assessment_strategy` (← assessment/), `dependents` (derived). These stay **on the leaf object** — no parallel representation is ever created.

## Runtime vs. schema (strict separation)
This schema defines the leaf's **intrinsic** properties (the Blueprint definition). A learner's *current state* on a leaf (mastery level, evidence, practice history) is **runtime data** held in `../progress/`, referencing the leaf by `id`. The two never mix.

## Versioning & changelog
Schema is versioned (`schema_version`). Prefer backward-compatible changes; breaking changes bump the major version and migrate all leaves.
- **v1.1** (2026-07-16) — added `cognitive_level` (Bloom) and `typical_learning_load`; renamed `assessment_evidence` → `assessment_strategy` (Blueprint defines *how* mastery is assessed; learner evidence is runtime). Reaffirmed strict runtime separation + single-representation as core rules.
- **v1.0** (2026-07-16) — initial schema ([SK-002](decisions.md)).
