# Knowledge

## Purpose
Build a canonical **Knowledge Graph** of atomic knowledge objects that satisfies every prerequisite referenced by the [Skill Graph](../skills/). Not reference notes — a graph of atomic objects with stable IDs and **explicit relationship edges** ([KK-001](decisions.md)).

## Model ([KK-001](decisions.md), [KK-002](decisions.md), [KK-003](decisions.md))
Every knowledge item is an atomic **Knowledge Object** conforming to [knowledge-object-schema.md](knowledge-object-schema.md) v1.1: stable `id`, atomic `definition`, explicit edges (`requires`, `related_to`), and optional `examples` / `common_misconceptions`. Dependencies are graph edges, not implicit; Knowledge↔Knowledge edges are kept strictly separate from Skill↔Knowledge resolution ([resolution.md](resolution.md)). Decompose until every Skill prerequisite resolves to ≥1 object.

## Domains (per [FD-003](../product/foundational-decisions.md))
- **grammar** (`K-GRA`) · **vocabulary** (`K-VOC`) · **phonology** (`K-PHON`).

## Status
**Complete (2026-07-16).** Three domains — grammar (26), vocabulary (9), phonology (8) = **43 atomic objects**, schema v1.1 — plus [resolution.md](resolution.md), which resolves every Skill `K-*` prerequisite. Pending Founder review.

## Structure
- `grammar.md` ✅ (26) · `vocabulary.md` ✅ (9) · `phonology.md` ✅ (8) — domain knowledge graphs, schema v1.1 (43 objects).
- [resolution.md](resolution.md) — maps every Skill `K-*` prerequisite to the knowledge object(s) that satisfy it (completeness-checked).
- [knowledge-object-schema.md](knowledge-object-schema.md) — canonical Knowledge Object schema v1.1.
- [decisions.md](decisions.md) — KK-001/002/003.

## Dependencies
- **Consumes:** the `K-*` prerequisites from [../skills/](../skills/) (per [../skills/consistency-review.md](../skills/consistency-review.md)).
- **Feeds:** resolves [../skills/](../skills/) prerequisites; referenced by [../curriculum/](../curriculum/), [../practice/](../practice/), [../assessment/](../assessment/).
