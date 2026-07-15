# Knowledge Decisions

Founder decisions defining the knowledge section. Indexed from [product/foundational-decisions.md](../product/foundational-decisions.md).

---

## KK-001 — Knowledge Graph rules
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** Build a canonical **Knowledge Graph** — not reference notes. Every knowledge item is an **atomic knowledge object** with a **stable identifier** and **explicit relationships** to other knowledge objects.
- Dependencies are modeled **explicitly** (graph edges), not implicitly, wherever possible.
- Decompose until **every Skill prerequisite** referenced by `../skills/` resolves to one or more atomic knowledge objects.
- The completed Knowledge Graph must fully satisfy the unresolved `K-*` dependencies flagged in [../skills/consistency-review.md](../skills/consistency-review.md) (`K-VOC`, `K-GRA`, `K-PHON`).
- **Single representation** (mirrors [../skills/](../skills/)): the knowledge object is referenced by `id`; curriculum/practice/assessment read it — no parallel representations.

**Implication.** A Knowledge Object schema ([KK-002](#kk-002--knowledge-object-schema)) governs every object; a [resolution.md](resolution.md) maps each Skill `K-*` prerequisite to the object(s) that satisfy it.

---

## KK-002 — Knowledge Object schema
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** Every knowledge object conforms to [knowledge-object-schema.md](knowledge-object-schema.md) v1.0: `id`, `name`, `domain`, `definition`, `requires` (explicit prerequisite edges), `related_to` (peer edges), `band_relevance`. Cross-graph skill-resolution lives in [resolution.md](resolution.md) (skill leaf → knowledge objects), keeping the knowledge graph's edges knowledge↔knowledge.
