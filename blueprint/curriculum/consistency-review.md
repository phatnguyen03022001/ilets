# Cross-Band Consistency Review
*Performed 2026-07-16 after all Band 3–9 pathways were built (schema v1.1). Per Founder directive: verify before proceeding to `../practice/`.*

## Inventory
| band | nodes | est. duration | character |
|---|---|---|---|
| 3 | 8 | ~20–27h | entry/foundation (knowledge core + entry skills) |
| 4 | 7 | ~14–20h | first structured band (task analysis, paragraphing, sustained speech) |
| 5 | 9 | ~21–29h | dense mid-band (complex sentences, cohesion, Task 2 dev, receptive inference) |
| 6 | 6 | ~13–19h | accuracy + clarity (overview, collocation, dense receptive, Y/N/NG) |
| 7 | 6 | ~14–20h | sophistication (flexibility, less-common lexis, inference mastery) |
| 8 | 5 | ~12–18h | high competence (near-error-free, skillful, effortless) |
| 9 | 3 | ~8–12h | ceiling (full flexibility + integrated mastery) |
| **Total** | **44** | **~102–145h** | |

## 1. Prerequisite consistency ✅ (with one flag)
- **Knowledge-before-skill** is respected throughout (e.g., K-GRA-003/004 before W-GRA-02/03 in C-B5-01; K-PHON-010/011 before S-P-01 in C-B3-05).
- Intra-skill prerequisites respected (e.g., W-CC-02 before W-CC-03/04).
- ⚠ *Flag:* several Skill Leaves carry `band_relevance 4–9` but are introduced at **Band 3** (entry) in the curriculum (e.g., W-GRA-01, S-P-01, L-COMP-01/02). The curriculum is authoritative for sequencing; `band_relevance` should be reconciled (3–9) in a calibration pass — **not** a structural error.

## 2. Workload balance ✅ (with one flag)
- Distribution is plausible: Bands 3 and 5 are densest (foundations + mid-band explosion of new structures), tapering to Band 9 (ceiling refinement).
- ⚠ *Flag:* **Band 5 is dense** (~21–29h, 9 nodes) — reflects real mid-band complexity; could optionally split, but is left intact for fidelity.

## 3. Cognitive progression ✅
- Ascends as expected: Band 3 (remember/understand/apply) → Band 5 (apply/create) → Band 7 (create/evaluate) → Band 9 (create/mastery). Higher bands add `evaluate`/`create` (inference, distractor rejection, flexible generation).

## 4. Mastery progression ✅
- Every band ends in a mastery-gated consolidation/exit node (C-B*n*-last) referencing that band's exit criteria ([../bands/](../bands/)); bands chain via [LD-001](../learning/decisions.md) (mastery-gated between bands).

## 5. Duplicate / missing learning paths ✅
- **Full leaf coverage:** all 64 Skill Leaves appear in ≥1 curriculum node (verified across Writing/Speaking/Listening/Reading). No missing path.
- Leaves that span bands (e.g., W-GRA-03 at 5–9) intentionally recur as *refinement* nodes at higher bands — not duplication (different mastery level).

## 6. Traceability Band → Curriculum → Skill → Knowledge ✅
- Complete chain exists: Band pathway (e.g., `band-5.md`) → Curriculum Node (`C-B5-01`) → Skill Leaf (`W-GRA-03`) → Knowledge Object (`K-GRA-004`), with skill→knowledge resolution in [../knowledge/resolution.md](../knowledge/resolution.md). Every `id` resolves.

## Conclusion
The curriculum is **structurally consistent and fully traceable**; no missing/duplicate paths. Two non-blocking flags (`band_relevance` reconciliation; Band 5 density) are deferred to a calibration pass. **Proceed to `../practice/`** is unblocked by this review.

## Deferred: calibration pass
- Reconcile Skill Leaf `band_relevance` with curriculum introduction band (Band 3 entry leaves → 3–9).
- Optionally rebalance Band 5 density.
- Tune `typical_learning_duration` estimates empirically.
