# Practice Coverage Review
*Verifies the binding ([binding.md](binding.md)) is complete and consistent. Per Founder directive (2026-07-16).*

## 1. Every Skill Leaf has ≥1 suitable Practice Type ✅
All **64** Skill Leaves are bound to ≥1 Practice Type in [binding.md](binding.md) Part 1. Most have **3+ types** spanning acquisition → consolidation → transfer/exam (e.g., W-TA-03 → PT-01 scaffolded, PT-05 redraft, PT-06 timed essay).

## 2. Every Curriculum Node has an appropriate practice sequence ✅
All **44** Curriculum Nodes reference Practice Types (Part 2), phase-aligned: knowledge-foundation/acquisition nodes use PT-02/07/12/17/18/19; application/consolidation nodes use PT-03/04/09/10/13/19/21; transfer nodes use PT-05/14/20; consolidation/exit nodes use PT-22/23. The sequence across a band ascends acquisition → consolidation → exam-readiness.

## 3. Every Practice Type is referenced by ≥1 Skill Leaf ✅ (with a design note)
- **PT-01..19** — each referenced by ≥1 Skill Leaf (Part 1). ✅
- **PT-20 (interleaved), PT-21 (adaptive), PT-22 (diagnostic), PT-23 (mock)** — these are **cross-cutting / exam** types whose natural consumer is the **Curriculum Node** (multi-leaf integration/exam), not an individual leaf. They are referenced by nodes (Part 2), not leaves — **by design**. This is the intended architecture (single-leaf training ≠ multi-leaf exam simulation).

## 4. No orphan Practice Types ✅
All **23** Practice Types are referenced by ≥1 leaf (PT-01..19) or ≥1 node (PT-20..23, including PT-21 added to consolidation nodes). None orphaned.

## 5. No Skill Leaf lacks an effective training strategy ✅
Every leaf has a training path covering its relevant phases:
- Grammar/mechanics leaves → drills (PT-02/18/19) + retrieval (PT-17) + error-correction (PT-04).
- Productive-skill leaves → scaffolded (PT-01) → redraft (PT-05) → timed (PT-06/11).
- Receptive leaves → skim/scan (PT-12) → comprehension (PT-13) → note-taking (PT-14) → timed (PT-15) → distractor review (PT-16).
- Knowledge-dependent leaves → spaced retrieval (PT-17).
No leaf is left with only a single weak strategy; no leaf is untrained.

## Traceability
Full chain intact: **Curriculum Node → Practice Type → (primary phase, [LD-006](../learning/decisions.md)) ↔ Skill Leaf ↔ Knowledge Object**. Practice is referenced everywhere by `id`; no Practice definition is duplicated in Skills or Curriculum ([PR-001](decisions.md)).

## Conclusion
The Practice Taxonomy is **fully bound and coverage-complete** — no orphans, no untrained leaves, every node has a phase-appropriate sequence. Cross-cutting/exam types are node-referenced by design. Practice remains a reusable canonical layer referenced only by `id`.

## Open items (non-blocking)
- [ ] Adaptive *selection* of types per learner (which PT, when) is runtime — defined in `../progress/`.
- [ ] Calibration may tune leaf→type bindings as empirical data arrives.
