# Bands Decisions

Founder decisions defining the bands section. Indexed from [product/foundational-decisions.md](../product/foundational-decisions.md).

---

## BD-001 — Bands section structure & scope
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.**
1. **One document per IELTS skill** (`writing.md`, `speaking.md`, `listening.md`, `reading.md`) — aligns with the official assessment model and lets Academic / General Training variants evolve independently.
2. **Band scope:** the learning Blueprint focuses on **Bands 3–9** (curriculum, learning paths, detailed learning requirements, exit criteria). **Bands 0–2 are documented as boundary definitions only** — to complete the assessment model and support learner diagnostics. **Do not** create curriculum, learning paths, or detailed learning requirements for Bands 0–2.
3. **Exit criteria belong in `bands/`** (per band); `progress/` references them, it does not redefine them.

**Implications.**
- Band docs include official descriptors for **all bands 0–9** (assessment completeness), but the **learning overlay** (required knowledge/skills, exit criteria, residual errors, higher-band exclusions) applies **only to Bands 3–9**.
- **Band 3 is the entry point** of the structured learning program. Bands 0–2 are diagnostic boundary markers (a learner below Band 3 is outside the structured range).
- Per-skill + variant-separable structure supports [FD-001](../product/foundational-decisions.md) (GT added later without restructuring).

---

## BD-002 — Hierarchical exit criteria (Task → Skill → Band progression)
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** Exit criteria are organized in a three-level hierarchy:
1. **Task-level mastery** — per-task exit criteria (e.g., Writing: Task 1 exit criteria *and* Task 2 exit criteria).
2. **Skill-level mastery** — a **Skill Exit Criteria** that certifies overall skill mastery, based on successful performance across the skill's tasks. It must **not duplicate** the task-level criteria; it defines the **aggregation condition** (all of the skill's task exit criteria met reliably, sustained, with no criterion dropping below band N−1).
3. **Overall band progression** — certification across all four skills; defined in `../progress/`, **not** in the skill docs.

**Implications.**
- Writing defines **Task 1 / Task 2 / Writing Skill** exit criteria per band.
- Speaking applies the hierarchy "if appropriate": Speaking has a single holistic descriptor set (not task-split), so its exit criteria are primarily at the **Speaking-Skill level** (part-level breakdown optional, deferred).
- Listening/Reading: exit criteria are at the skill level (score-based + comprehension inference per [BD-003](#bd-003--receptive-skills-three-category-evidence-labeling)).

---

## BD-003 — Receptive skills: three-category evidence labeling
**Date:** 2026-07-16 · **Status:** Decided (Founder)

**Decision.** For **Listening and Reading** (which rely on score conversion rather than public analytic performance descriptors), every claim is labeled as exactly one of:
- **Official Evidence** — verbatim from official IELTS documentation (conversion tables, official statements), with citation.
- **Evidence-Based Interpretation** — interpretation grounded in official docs / Cambridge materials, with the supporting evidence cited.
- **Blueprint Inference** — inference beyond official documentation; explicitly labeled with supporting evidence **and a confidence level** (High / Medium / Low).

**Rationale.** Maintain traceability; never blur the categories. **Traceability is more important than making stronger claims.**

**Implication.** The receptive learning overlay (what comprehension abilities a given band implies) is largely **Blueprint Inference** and must carry evidence + a confidence level. Productive-skill overlays (Writing, Speaking) are likewise labeled Blueprint Inference where they go beyond the official descriptors.
