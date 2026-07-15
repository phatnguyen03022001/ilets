# Learner State Transitions & Decisions
*Rules governing how learner state changes — referencing canonical objects **by `id` only** ([PG-001](decisions.md)). Calibration via [AM-003](../assessment/decisions.md) defaults.*

## 1. Mastery state transitions (per `LeafMasteryState`)
- `not_started → practicing`: learner completes a Practice Type event targeting the leaf (PT-* per [../practice/binding.md](../practice/binding.md)).
- `practicing → emerging`: assessment evidence (AT-*) accumulates; `confidence` rises.
- `emerging → mastered`: **`confidence ≥ 0.80` AND `evidence_count ≥ 2`** (independent demonstrations) AND no criterion below band N−1 ([AM-003](../assessment/decisions.md); [LD-001](../learning/decisions.md)). Certification via `AT-05` portfolio.
- `mastered → emerging` (regression): a later assessment shows regression below the certified level → revert; re-practice / re-assess.
- (`KnowledgeState`: `not_acquired → learning → acquired`, via `AT-03` probes — same threshold logic.)

## 2. Band progression rules ([LD-001](../learning/decisions.md))
- **Between bands: mastery-gated.** Band N+1 unlocks only when all four skill `BandCertificationState`s at N are `certified` (via `AT-05`; [BD-002](../bands/decisions.md) exit criteria).
- **Within band: adaptive** (§4) — never bypasses required objectives or mastery criteria.
- **Same outcomes, different paths** — every learner reaches the same band exit criteria via different node/leaf sequences.

## 3. Prerequisite enforcement ([LD-004](../learning/decisions.md))
- A leaf/node `unlocks` when its **`Required`** prerequisites (skill `requires` + knowledge `requires`, classified per LD-004) are satisfied (`mastery_state=mastered` / `state=acquired`).
- **`Recommended`** prereqs inform adaptive ordering but do **not** gate; **`Independent`** never gates.
- Hard gates limited to the minimum set (LD-004).

## 4. Adaptive scheduling (within band)
Select Practice/Assessment Types by learner state (references PT/AT by `id`):
- **Weak leaves** (low confidence) → acquisition/consolidation PTs (from the leaf's `practice_item_types`).
- **Due reviews** (`review_queue`) → `PT-17` spaced retrieval.
- **Near-mastery leaves** → `AT-05` certification probe (when ≥2 evidence feasible).
- **Exam-prep mode** ([LD-005](../learning/decisions.md)) → `AT-07` mock + higher-band exposure (non-certifying).
Adaptive selection is runtime; it **never bypasses mastery criteria** (LD-001 invariant).

## 5. Review scheduling ([../learning/review.md](../learning/review.md))
- **Performance-graded spacing:** schedule a knowledge/leaf item's next review when retrieval accuracy drops toward ~80–90% (just before full loss).
- Expanding intervals; `review_queue` holds due items.

## 6. Certification state
- **Certified** when `AT-05` portfolio is satisfied (≥2 demonstrations, confidence ≥0.80, no regression) → `BandCertificationState.status = certified`. This gates band progression (**Learning Progression**).
- `AT-07` mock → exam-readiness estimate, **non-certifying** ([LD-005](../learning/decisions.md)). A learner may be certified yet mock-underperform, or vice versa.

## 7. Learning recommendations
Next-best actions, referencing canonical objects by `id`:
- **Advance** → next Curriculum Node whose prerequisites are satisfied.
- **Remediate** → a regressed/weak leaf's `remediation` path + relevant PTs.
- **Review** → due spaced-retrieval items.
- **Exam-prep** → when `exam_prep_mode` is active: `AT-07` mock + higher-band exposure (LD-005).

## Traceability
Every transition is explainable via canonical objects: state changes reference `leaf_id`/`knowledge_id`; decisions reference Curriculum Nodes, Practice Types, Assessment Types; certification references Band exit criteria + `AT-05`. **No transition depends on non-canonical data.**
