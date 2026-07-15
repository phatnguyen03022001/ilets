# Cross-Skill Consistency Review
*Performed 2026-07-16, after all four skills were decomposed to [leaf-schema.md](leaf-schema.md) v1.1. Per Founder directive: optimize for **cross-skill consistency**; field **calibration is deferred** to a dedicated phase that compares equivalent skills across all four — not Writing in isolation.*

## Skill Graph inventory
| skill | components | leaves | decomposition axis |
|---|---|---|---|
| Writing | TA, TR, CC, LR, GRA | 23 | official criteria (+ task split TA/TR) |
| Speaking | FC, LR, GRA, P | 18 | official criteria |
| Listening | COMP, QT | 11 | comprehension ability × question-type (Blueprint Inference) |
| Reading | COMP, QT | 12 | comprehension ability × question-type (Blueprint Inference) |
| **Total** | | **64** | |

## Dimension-by-dimension consistency

### 1. `cognitive_level` (Bloom)
- **Productive** (W/S) lean `apply`/`create`/`analyze`, with one `remember` (W-LR-05 spelling) and `evaluate` (W-TR-04 relevance). **Receptive** (L/R) lean `understand`/`analyze`/`evaluate` (input comprehension; no `create`). **This divergence is correct** — it reflects skill nature, not inconsistency.
- ✅ Internally consistent within each pair (W↔S productive; L↔R receptive).
- ⚠ *Calibration flag:* `remember` is used only once (W-LR-05). Mechanical-accuracy leaves like W-GRA-06 (punctuation) may also warrant `remember` — review in calibration.

### 2. `typical_learning_load`
- ✅ Equivalent leaves across skills carry consistent loads (e.g., W-GRA-03 complex sentences = `high` ↔ S-GRA-02 complex spoken = `high`; W-LR-02 collocation = `high` ↔ S-LR-03 idiomatic = `high`).
- ⚠ *Calibration flag:* distribution skews `high`/`medium`; `low` is rarely used. Calibration should rebalance so `load` discriminates meaningfully across the graph.

### 3. Prerequisite depth
- ✅ Productive leaves chain within-component (e.g., W-CC-02→W-CC-01) and to `K-*` knowledge items; receptive QT leaves chain to their COMP prereqs (e.g., R-QT-02→R-COMP-03). Depth is shallow (1–2 hops) and consistent.
- ⚠ *Dependency to resolve:* `K-VOC` / `K-GRA` / `K-PHON` are referenced but **not yet built** (`../knowledge/` phase). When knowledge/ is built, every `K-*` prerequisite must resolve to a real knowledge leaf.

### 4. `mastery_criteria`
- ⚠ *Calibration flag:* reliability threshold is inconsistent — productive leaves use "≥2 responses" / "≥1 min"; receptive leaves use the vaguer "reliably". Calibration should **standardize** all mastery criteria to the [LD-001](../learning/decisions.md) reliability rule (≥2 independent demonstrations, sustained, no drop below band N−1) so the Skill Graph and `../bands/` exit criteria use one definition.

### 5. `common_errors` patterns
- ✅ Recurring cross-skill patterns are coherent: **L1-transfer** (articles W-GRA-05, collocation W-LR-02, phonemes S-P-01, word-choice S-LR-04); **over-literal processing** (verbatim copying W-LR-03, literal listening/reading L-COMP-03/R-COMP-04); **distractor/mechanical traps** (L-COMP-05, R-QT-02).
- ⚠ *Calibration flag:* error-list depth varies (some leaves 3 items, some 4). Standardize depth in calibration; consider a shared **error-pattern taxonomy** (L1-transfer / accuracy / strategic / processing) referenced across skills.

## Cross-skill shared dependencies (knowledge layer)
| knowledge item | referenced by |
|---|---|
| `K-VOC` (vocabulary) | W-LR-01/02/05, S-LR-01/04, R-COMP-06 |
| `K-GRA` (grammar) | W-CC-03/04, W-GRA-01..05, S-GRA-01..04 |
| `K-PHON` (pronunciation) | S-P-01 |

These define what `../knowledge/` must provide so the Skill Graph's prerequisites resolve.

## Conclusion
The **Skill Graph is structurally consistent** (one schema, v1.1; productive by official criteria, receptive by comprehension×question-type; atomic leaves; shared knowledge layer). No structural rework needed. **Five calibration flags** above are deferred to a dedicated calibration phase that tunes field values across all four skills together.

## Deferred: calibration phase (scope)
A later review phase will, across all four skills jointly:
1. rebalance `typical_learning_load` (reduce `high` overuse);
2. standardize `mastery_criteria` to the LD-001 reliability rule;
3. review `cognitive_level` assignments (e.g., mechanical-accuracy leaves);
4. standardize `common_errors` depth + adopt a shared error-pattern taxonomy;
5. confirm prerequisite depth and resolve all `K-*` references once `../knowledge/` exists.
