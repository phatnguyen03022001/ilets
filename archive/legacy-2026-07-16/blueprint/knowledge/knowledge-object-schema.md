# Knowledge Object Schema (v1.1)

The canonical atomic unit of the Knowledge Graph. Every knowledge object conforms to this schema ([KK-002](decisions.md)). Referenced by `id` across `../curriculum/`, `../practice/`, `../assessment/`, and to resolve `../skills/` prerequisites — no parallel representations ([KK-001](decisions.md)).

> **Strict separation (Founder):** Knowledge↔Knowledge relationships (`requires`/`related_to`) live here; Skill↔Knowledge resolution lives in [resolution.md](resolution.md). The two never mix.

## Fields
| field | type | notes |
|---|---|---|
| `id` | string, unique, stable | `K-<DOMAIN>-<nnn>` — e.g. `K-GRA-004`. `GRA`=grammar, `VOC`=vocabulary, `PHON`=phonology. |
| `name` | string | human-readable label. |
| `domain` | enum | `grammar` \| `vocabulary` \| `phonology`. |
| `definition` | string | **one atomic** knowledge unit. |
| `requires` | array of `id` | explicit **prerequisite** edges (knowledge→knowledge). |
| `related_to` | array of `id` | explicit **peer** edges (associated, non-blocking). |
| `band_relevance` | range/array (3–9) | bands where this knowledge materially matters (optional). |
| `examples` | array | illustrative items — intrinsic to the knowledge (optional). |
| `common_misconceptions` | array | frequent misconceptions/errors about this knowledge — intrinsic; feeds AI tutoring, remediation, content generation (optional). |

## Atomicity (stop condition)
Decompose a knowledge object **only while it still contains multiple independently learnable knowledge units** (or cannot be taught/assessed as one). Stop when it is a single independently-learnable unit. *(Founder refinement, 2026-07-16.)*

## Relationships are explicit
`requires` and `related_to` are graph edges — dependencies modeled explicitly, never implicit ([KK-001](decisions.md)).

## Cross-graph resolution (skill → knowledge)
Which skill leaves a knowledge object satisfies is mapped in [resolution.md](resolution.md). Kept separate so the Knowledge Graph's edges stay knowledge↔knowledge.

## Versioning & changelog
- **v1.1** (2026-07-16) — added `common_misconceptions` (intrinsic; AI tutoring/remediation/content); confirmed `examples`; refined atomicity stop-condition to "independently learnable knowledge unit". Reaffirmed strict Knowledge↔Knowledge vs Skill-resolution separation.
- **v1.0** (2026-07-16) — initial schema ([KK-002](decisions.md)).
