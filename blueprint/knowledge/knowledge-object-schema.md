# Knowledge Object Schema (v1.0)

The canonical atomic unit of the Knowledge Graph. Every knowledge object conforms to this schema ([KK-002](decisions.md)). Referenced by `id` across `../curriculum/`, `../practice/`, `../assessment/`, and to resolve `../skills/` prerequisites — no parallel representations ([KK-001](decisions.md)).

## Fields
| field | type | notes |
|---|---|---|
| `id` | string, unique, stable | `K-<DOMAIN>-<nnn>` — e.g. `K-GRA-004`. `GRA`=grammar, `VOC`=vocabulary, `PHON`=phonology. |
| `name` | string | human-readable label. |
| `domain` | enum | `grammar` \| `vocabulary` \| `phonology`. |
| `definition` | string | **one atomic** knowledge unit. If a node holds >1 independent unit → not atomic → decompose. |
| `requires` | array of `id` | explicit **prerequisite** edges (knowledge→knowledge). The graph's dependency structure. |
| `related_to` | array of `id` | explicit **peer** edges (associated but non-blocking). |
| `band_relevance` | range/array (3–9) | bands where this knowledge materially matters (optional). |
| `examples` | array | illustrative items (optional). |

## Relationships are explicit
`requires` and `related_to` are **graph edges** — dependencies are modeled explicitly, never implicit ([KK-001](decisions.md)). An object's prerequisites are other knowledge objects (or, rarely, foundational assumptions stated as such).

## Cross-graph resolution (skill → knowledge)
Which skill leaves a knowledge object satisfies is mapped in [resolution.md](resolution.md) (skill-leaf `K-*` prerequisite → knowledge object `id`s). This keeps the Knowledge Graph's own edges (knowledge↔knowledge) separate from the cross-graph mapping (skill↔knowledge), both explicit.

## Atomicity / stop condition
Subdivide a knowledge object only while it still contains multiple independent knowledge units or cannot be taught/assessed as one. Mirror of the skills/ atomicity principle.

## Versioning
Versioned (`schema_version = 1.0`). Breaking changes bump the major version and migrate all objects.
