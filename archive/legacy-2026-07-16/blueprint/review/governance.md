# Blueprint Governance — QA Catalog
*The checks the Governance layer ([GV-001](decisions.md)) executes repository-wide. Each check has a definition, pass criterion, and home in the Health Report ([health-report.md](health-report.md)).*

## Checks
| # | check | verifies | pass criterion |
|---|---|---|---|
| 1 | **Structural consistency** | all required top-level categories exist; each has README + decisions where applicable | the 11 categories (+ `localization/`) per CLAUDE.md exist and are internally structured |
| 2 | **Terminology consistency** | one term per concept; no drift | terms match [../glossary/](../glossary/); no synonyms used as distinct concepts |
| 3 | **Traceability completeness** | Band → Curriculum → Skill → Knowledge (+ Practice/Assessment) chain | every canonical object is reachable; [../curriculum/consistency-review.md](../curriculum/consistency-review.md) + [../knowledge/resolution.md](../knowledge/resolution.md) hold |
| 4 | **Dependency correctness** | `requires`/prerequisite edges point to existing objects | no dangling `requires`/prerequisite references |
| 5 | **Prerequisite consistency** | prereqs respected in curriculum; LD-004 classification applied | knowledge-before-skill; Required/Recommended/Independent consistent |
| 6 | **Canonical object uniqueness** | no duplicate `id`s within a namespace | every `id` is unique within its layer |
| 7 | **Duplicate detection** | no duplicated concepts across layers | each concept defined once; cross-refs by `id`, not redefined |
| 8 | **Orphan detection** | every object is referenced (no unreachable leaves/knowledge/types) | every Skill Leaf, Knowledge Object, Practice/Assessment Type is consumed somewhere |
| 9 | **Academic vs GT separation** | variant convention ([GV-002](decisions.md)) applied | shared vs variant-specific correctly tagged/separated |
| 10 | **Evidence labeling consistency** | BD-003 three-category labels applied (Official Evidence / Evidence-Based Interpretation / Blueprint Inference) | receptive docs + inferences correctly labeled with confidence |
| 11 | **Decision-record consistency** | decision series (FD/LD/BD/SK/KK/CR/PR/AM/PG/GV) indexed + current | foundational index lists every decision; no stale entries |
| 12 | **Reference integrity** | all cross-layer `id` references resolve | every `K-*`, `W/S/L/R-*`, `PT-*`, `AT-*`, `C-B*` reference points to a real object |

## Execution
Run all 12 checks; record each finding in [health-report.md](health-report.md) with severity (`Critical`/`High`/`Medium`/`Low`), rationale, and recommended resolution. **Freeze requires zero unresolved `Critical` issues** ([GV-001](decisions.md)).
