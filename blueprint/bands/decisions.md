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
