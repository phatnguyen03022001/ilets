# Practice Decisions

Founder decisions defining the practice section. Indexed from [product/foundational-decisions.md](../product/foundational-decisions.md).

---

## PR-001 — Practice is a canonical Practice Taxonomy (not exercises)
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** `practice/` defines the canonical **Practice Taxonomy** — a reusable layer of practice **types** that every Skill Leaf and Curriculum Node can reference. It is **not** a collection of exercises/activities.
- Each practice **type** is a canonical template, referenced by `id`.
- **Curriculum nodes reference practice types by `id`** — they do not embed practice definitions.
- **Skill Leaves reference canonical practice types** (via their `practice_item_types` slot, [SK-002](../skills/decisions.md)) — they do not define their own practice independently.
- Concrete practice **items/instances** are generated elsewhere against these types; `practice/` defines the types only.

**Each practice type answers:**
1. what learning objective it serves;
2. which cognitive operations it develops;
3. which Skill Leaves it supports;
4. which Knowledge Objects it reinforces;
5. when it should be used within the curriculum;
6. its purpose — acquisition / consolidation / transfer / exam-readiness;
7. its mode — individual / mixed / adaptive / timed / diagnostic.

**Architectural separation (maintained throughout the Blueprint):**
| layer | responsibility |
|---|---|
| Knowledge | what must be **known** |
| Skills | what must be **demonstrated** |
| Curriculum | **when** things are learned |
| Practice | **how** they are trained |
| Assessment | how **mastery** is measured |
| Progress | learner **runtime state** |

These responsibilities stay strictly separate.

**Implication.** A Practice Type schema ([PR-002](#pr-002--practice-type-schema)) governs every type; the taxonomy lives in [taxonomy.md](taxonomy.md).
