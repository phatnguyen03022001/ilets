# Assessment Binding
*Authoritative population of each Skill Leaf's `assessment_strategy` slot, plus exit→evidence, node→culmination, and practice→assessment-pathway maps — **by `id`**, no duplication ([AM-001](decisions.md)). Assessment Type definitions remain the single source in [taxonomy.md](taxonomy.md).*

## Part 1 — Skill Leaf → Assessment Strategy
Strategy is uniform within leaf category; the table covers **all 64** leaves (every leaf's `assessment_strategy` = its row).

| leaf category (count) | assessment_strategy (`assessment_strategy` slot) |
|---|---|
| Writing productive — `W-TA/TR/CC/LR/GRA-*` (23) | `AT-01` (criterion-referenced productive) + `AT-03` (knowledge probe for LR/GRA) + `AT-05` (certification portfolio) |
| Speaking language — `S-FC/LR/GRA-*` (13) | `AT-01` + `AT-03` + `AT-05` |
| Speaking pronunciation — `S-P-*` (5) | `AT-01` (AI speech scoring, calibrated) + `AT-03` (phoneme probe) + `AT-05` |
| Listening receptive — `L-COMP/L-QT-*` (11) | `AT-02` (objective receptive) + `AT-05` |
| Reading receptive — `R-COMP/R-QT-*` (12) | `AT-02` + `AT-05` |

## Part 2 — Band Exit Criterion → Assessment Evidence
Each Band exit criterion ([BD-002](../bands/decisions.md)) is supported by:

| skill | exit evidence (Bands 3–9) |
|---|---|
| Writing (Task 1 + Task 2 + Writing-Skill) | `AT-01` per task (scored vs TA/TR/CC/LR/GRA) + `AT-05` (≥2 demonstrations per task, no criterion below N−1) |
| Speaking (Speaking-Skill) | `AT-01` (scored vs FC/LR/GRA/P) + `AT-05` |
| Listening (Listening-Skill) | `AT-02` (raw→band, ≥2 full tests, no result below N−1) + `AT-05` |
| Reading (Reading-Skill) | `AT-02` + `AT-05` |
| Bands 0–2 (boundary only) | `AT-04` diagnostic — **no certification** ([BD-001](../bands/decisions.md)) |

Certification (`AT-05`) is the Learning-Progression gate ([LD-001](../learning/decisions.md)); the `AT-07` mock is **not** certifying ([LD-005](../learning/decisions.md)).

## Part 3 — Curriculum Node → Culmination Evidence
Every node culminates in measurable evidence; exit nodes additionally gate via `AT-05`.

| node type | culmination evidence |
|---|---|
| Knowledge/acquisition nodes (e.g., C-B3-01..05, C-B5-01) | `AT-03` (knowledge probe) + formative `AT-01`/`AT-02` |
| Skill-application/consolidation nodes (e.g., C-B4-01..06, C-B5-03..08) | `AT-01`/`AT-02` (formative, per leaves) |
| Band entry diagnostics (C-B3-08 entry path) | `AT-04` (diagnostic snapshot) |
| Band consolidation/exit nodes (C-B*n*-last, n=3..9) | `AT-05` (mastery portfolio → band certification) + `AT-04` snapshot |
| Exam-Preparation usage (per [LD-005](../learning/decisions.md)) | `AT-07` (mock, non-certifying) + `AT-06` (optional human) |

## Part 4 — Practice Type → Assessment Pathway
Every Practice Type prepares learners for ≥1 assessment pathway.

| practice type | prepares for (assessment) |
|---|---|
| PT-01 scaffolded writing | `AT-01` |
| PT-02 sentence-combining | `AT-01`, `AT-03` |
| PT-03 paraphrase | `AT-01` |
| PT-04 error-correction | `AT-03`, `AT-04` |
| PT-05 redraft-after-feedback | `AT-01` |
| PT-06 timed essay | `AT-01`, `AT-07` |
| PT-07 pronunciation drill | `AT-01`, `AT-03` |
| PT-08 shadowing | `AT-01` |
| PT-09 long-turn | `AT-01` |
| PT-10 Q&A/role-play | `AT-01` |
| PT-11 timed mock speaking | `AT-01`, `AT-07` |
| PT-12 skim/scan | `AT-02` |
| PT-13 comprehension set | `AT-02` |
| PT-14 note-taking | `AT-02` |
| PT-15 timed section | `AT-02`, `AT-07` |
| PT-16 distractor review | `AT-04` |
| PT-17 spaced retrieval | `AT-03` |
| PT-18 collocation drill | `AT-03` |
| PT-19 gap-fill | `AT-03`, `AT-02` |
| PT-20 interleaved mixed | `AT-04`, `AT-07` |
| PT-21 adaptive set | `AT-04` |
| PT-22 diagnostic checkpoint | `AT-04` |
| PT-23 full mock | `AT-07` |

## Notes
- Assessment is independent from Practice ([AM-001](decisions.md)); this map shows preparation pathways only — never shared definitions.
- Assessment *events* (specific runs) and learner *evidence records* are runtime, in `../progress/`.
