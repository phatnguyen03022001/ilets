# Skills Decisions

Founder decisions defining the skills section. Indexed from [product/foundational-decisions.md](../product/foundational-decisions.md).

---

## SK-001 — Skill decomposition rules (leaf-node criteria)
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** Decompose every competency from `../bands/` into a skill hierarchy until every band requirement maps to a complete learning path. Do **not** start from an arbitrary number of micro-skills.

Every **leaf node** in the hierarchy must satisfy **all** of:
1. has a single learning objective;
2. can be taught independently;
3. can be practiced independently;
4. can be assessed independently;
5. has explicit mastery criteria;
6. has clear prerequisites (if any);
7. belongs to one or more bands;
8. traces back to one or more official IELTS requirements.

**Stop condition.** Stop subdividing only when further subdivision no longer improves teaching, practice, assessment, or personalization.

**Implications.**
- The skills hierarchy is the foundation for `../curriculum/`, `../practice/`, `../assessment/`, and adaptive learning.
- Each leaf carries: objective, mastery criteria, prerequisites, band(s), and a trace to an official IELTS requirement. "Independently teachable / practiceable / assessable" is the **leaf invariant** (the stop condition) — verified per leaf, not re-stated everywhere.
- Skills that concern grammar/vocabulary reference the corresponding `../knowledge/` item (built later) as a **content source / prerequisite**; the skill itself is about *applying* the language in the skill context (per [FD-003](../product/foundational-decisions.md)).
- Leaf granularity is validated on Writing before replication to Speaking / Listening / Reading.

---

## SK-002 — Atomic decomposition + common leaf schema
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.**
1. **Optimize for atomic learning skills, not a target number of micro-skills.** Keep the official assessment criteria as the **primary decomposition axis**. Decompose further **only** when a node still contains multiple independent learning objectives, or cannot be taught, practiced, assessed, **or remediated** independently. *(Refines [SK-001](#sk-001--skill-decomposition-rules-leaf-node-criteria): adds "remediable" to the leaf invariant.)*
2. **Common leaf schema.** Every skill leaf conforms to one canonical schema — [leaf-schema.md](leaf-schema.md) v1.0 — the single "learning object." Curriculum, practice, assessment, AI tutoring, and personalization all reference the leaf **by `id`** and read the fields they need; they do **not** create parallel representations of the skill.

**Implications.**
- `writing.md` is retro-fitted to schema v1.0 (reference implementation).
- Consumer-specific fields (`practice_item_types`, `assessment_evidence`) are **owned by the leaf** and populated incrementally by `../practice/` and `../assessment/` — still the single representation.
- `mastery_states` (the progress/personalization scale) is a fixed enum defined once in the schema, not repeated per leaf.
