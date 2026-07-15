# Knowledge

## Purpose
Build a canonical **Knowledge Graph** of atomic knowledge objects that satisfies every prerequisite referenced by the [Skill Graph](../skills/). Not reference notes — a graph of atomic objects with stable IDs and **explicit relationship edges** ([KK-001](decisions.md)).

## Model ([KK-001](decisions.md), [KK-002](decisions.md), [KK-003](decisions.md))
Every knowledge item is an atomic **Knowledge Object** conforming to [knowledge-object-schema.md](knowledge-object-schema.md) v1.1: stable `id`, atomic `definition`, explicit edges (`requires`, `related_to`), and optional `examples` / `common_misconceptions`. Dependencies are graph edges, not implicit; Knowledge↔Knowledge edges are kept strictly separate from Skill↔Knowledge resolution ([resolution.md](resolution.md)). Decompose until every Skill prerequisite resolves to ≥1 object.

## Domains (per [FD-003](../product/foundational-decisions.md))
- **grammar** (`K-GRA`) · **vocabulary** (`K-VOC`) · **phonology** (`K-PHON`).

## Status
**In progress.** `grammar.md` drafted as the template (most-referenced: `K-GRA`). Pending validation, then `vocabulary.md` + `phonology.md` + the resolution map.

## Structure
- `grammar.md` / `vocabulary.md` / `phonology.md` — domain knowledge graphs.
- [resolution.md](resolution.md) — maps every Skill `K-*` prerequisite to the knowledge object(s) that satisfy it.
- [knowledge-object-schema.md](knowledge-object-schema.md) — canonical Knowledge Object schema v1.0.
- [decisions.md](decisions.md) — KK-001/002.

## Dependencies
- **Consumes:** the `K-*` prerequisites from [../skills/](../skills/) (per [../skills/consistency-review.md](../skills/consistency-review.md)).
- **Feeds:** resolves [../skills/](../skills/) prerequisites; referenced by [../curriculum/](../curriculum/), [../practice/](../practice/), [../assessment/](../assessment/).
