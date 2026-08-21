STATUS: SUPPORTING
ROLE: DECISION RATIONALE
AUTHORITY: NONE

# Decisions

This file explains why important active choices exist. It is a navigation and rationale aid, not a source of canonical truth.

If a statement here conflicts with a canonical owner, **the canonical owner wins** and this ledger must be corrected.

Historical wording and superseded process decisions remain preserved under `archive/legacy-2026-07-16/`.

## D-001 — Self-describing, vendor-neutral repository

**Decision:** The active repository must be understandable without `CLAUDE.md`, `AGENTS.md`, hidden prompts, or prior chat history.

**Canonical owner:** `../CONSTITUTION.md`

**Rationale:** Understanding the project is a property of the repository knowledge architecture, not of whichever coding/reasoning vendor happens to open it.

**Evidence / provenance:** 2026-08 document-architecture refactor; legacy vendor instructions preserved in archive.

## D-002 — Small explicit authority surface

**Decision:** Use the frozen active structure of README + Constitution + Objective + 12 numbered domain owners + this supporting decision ledger.

**Canonical owner:** `../CONSTITUTION.md`

**Rationale:** The legacy Blueprint fragmented semantics across category READMEs, per-band files, per-skill files, schemas, bindings, decisions, consistency reviews, and coverage reviews. Explicit ownership reduces contradictory mini-SSOTs while preserving semantics inside their owner.

**Evidence / provenance:** Legacy tree under `archive/legacy-2026-07-16/`.

## D-003 — Academic first; General Training as variant overlay

**Decision:** Fully specify Academic first and model General Training differences as overlays on shared concepts.

**Canonical owner:** `00-PRODUCT.md`; external variant truth in `02-IELTS-MODEL.md`

**Rationale:** Listening and Speaking are shared across the variants, while Reading and Writing contain variant differences. Duplicating the whole learning model would create unnecessary divergence.

**Legacy decision:** `FD-001`

## D-004 — Canonical learning system is L1-agnostic

**Decision:** Canonical learning requirements are independent of learner first language; localization/remediation may use L1-specific overlays.

**Canonical owner:** `00-PRODUCT.md`

**Rationale:** IELTS competence is not different by learner L1, while explanations and predictable transfer errors can be localized without changing the target standard.

**Legacy decision:** `FD-002`

## D-005 — Four IELTS skills; vocabulary/grammar/phonology are enabling knowledge

**Decision:** Listening, Reading, Writing, and Speaking are the first-class skills. Language systems such as vocabulary, grammar, and phonology are Knowledge Objects supporting those skills.

**Canonical owners:** `03-SKILLS.md`, `04-KNOWLEDGE.md`

**Rationale:** Aligns the skill model with IELTS section scoring while retaining explicit teachable prerequisites.

**Legacy decision:** `FD-003`

## D-006 — Learning domain is implementation-agnostic

**Decision:** The Blueprint defines canonical learning truth independent of app architecture, delivery technology, infrastructure, and business model.

**Canonical owner:** `00-PRODUCT.md`

**Rationale:** Technology should consume the learning system rather than shape or duplicate it.

**Legacy decision:** `FD-006`

## D-007 — Mastery-based progression with adaptive within-band paths

**Decision:** Certification/advancement requires mastery evidence, while within-band sequencing may adapt to learner state.

**Canonical owner:** `09-PROGRESSION.md`; mastery threshold in `08-ASSESSMENT.md`

**Rationale:** Preserve a stable target standard while allowing personalized path, pacing, remediation, and practice.

**Legacy decision:** `LD-001`

## D-008 — AI-primary operation without mandatory human dependency

**Decision:** The core system can provide AI-supported tutoring/feedback/assessment without requiring a human reviewer; human expert review is optional, particularly for uncertain productive judgments.

**Canonical owners:** `00-PRODUCT.md`, `08-ASSESSMENT.md`

**Rationale:** Preserve scalability and immediate learning support while making productive-skill uncertainty explicit instead of pretending automated judgment is infallible.

**Legacy decision:** `LD-002`

## D-009 — Feedback timing is stage-dependent

**Decision:** Acquisition generally uses immediate corrective feedback; retrieval/transfer/exam-readiness increasingly preserve independent performance before feedback.

**Canonical owner:** `07-PRACTICE.md`

**Rationale:** A universal immediate-feedback rule conflicts with retrieval and transfer goals.

**Legacy decision:** `LD-003`

## D-010 — Required / Recommended / Independent prerequisites

**Decision:** Use the minimum justified hard gates. A dependency is Required only when missing foundation makes dependent learning ineffective and no reasonable adaptive workaround exists.

**Canonical owner:** `06-CURRICULUM.md`; runtime behavior in `09-PROGRESSION.md`

**Rationale:** Preserve prerequisite integrity without turning every beneficial ordering into a blocking gate.

**Legacy decision:** `LD-004`

## D-011 — Learning Progression and Exam Preparation are independent

**Decision:** Higher-band/exam-like exposure may occur before mastery for diagnosis and readiness, but exposure cannot certify, unlock, or bypass learning requirements.

**Canonical owner:** `09-PROGRESSION.md`

**Rationale:** Learners with fixed test dates need realistic test exposure without corrupting mastery semantics.

**Legacy decision:** `LD-005`

## D-012 — Six canonical learning phases

**Decision:** Acquisition, Consolidation, Retrieval, Transfer, Fluency, Exam Readiness.

**Canonical owner:** `07-PRACTICE.md`

**Rationale:** Practice types need a stable learning-purpose vocabulary broad enough to distinguish initial learning, retention, transfer, automaticity, and test simulation.

**Legacy decision:** `LD-006`

## D-013 — Atomic Skill Leaves with stable IDs

**Decision:** Decompose until a leaf is independently teachable, practiceable, assessable, and remediable; do not target an arbitrary leaf count.

**Canonical owner:** `03-SKILLS.md`

**Rationale:** Atomicity should serve learning and remediation, not document aesthetics. The frozen inventory has 64 stable leaves.

**Legacy decisions:** `SK-001`, `SK-002`, `SK-003`

## D-014 — Atomic Knowledge Graph with explicit dependency edges

**Decision:** Use stable atomic Knowledge Objects and explicit Knowledge→Knowledge plus Skill→Knowledge resolution.

**Canonical owner:** `04-KNOWLEDGE.md`

**Rationale:** Skill prerequisites must resolve to actual learnable knowledge rather than vague labels such as "grammar" or "vocabulary". The frozen inventory has 46 objects.

**Legacy decisions:** `KK-001`, `KK-002`, `KK-003`

## D-015 — Bands 3–9 are the structured learning range

**Decision:** Bands 0–2 are diagnostic/external boundaries only; detailed curriculum begins at Band 3 and continues through Band 9.

**Canonical owner:** `05-BANDS.md`

**Rationale:** Preserve full IELTS scale awareness without creating detailed curriculum below the chosen structured entry floor.

**Legacy decision:** `BD-001`

## D-016 — Hierarchical exit criteria

**Decision:** Where relevant, distinguish task-level thresholds from skill-level thresholds, then let Progression consume valid certification evidence.

**Canonical owner:** `05-BANDS.md`

**Rationale:** Prevent task-specific requirements from being collapsed into vague overall mastery while avoiding duplicated progression logic.

**Legacy decision:** `BD-002`

## D-017 — Receptive learning overlays must expose inference status

**Decision:** Listening/Reading raw-score facts remain official evidence; detailed per-band receptive abilities are labeled Blueprint Inference with confidence rather than presented as official analytic descriptors.

**Canonical owners:** `02-IELTS-MODEL.md`, `05-BANDS.md`

**Rationale:** Public IELTS receptive scoring does not provide the same analytic band-descriptor structure available for Writing/Speaking.

**Legacy decision:** `BD-003`

## D-018 — Curriculum orchestrates canonical objects

**Decision:** Curriculum Nodes reference Skill Leaves and Knowledge Objects by ID and own sequencing, not parallel learning definitions.

**Canonical owner:** `06-CURRICULUM.md`

**Rationale:** The pathway must organize the graph without becoming another copy of it. The frozen inventory has 44 node IDs.

**Legacy decisions:** `CR-001`, `CR-003`

## D-019 — Practice is a reusable taxonomy, not an exercise library

**Decision:** Define reusable Practice Types; concrete exercises are instances.

**Canonical owner:** `07-PRACTICE.md`

**Rationale:** The learning strategy should remain stable while concrete content can be generated, authored, replaced, and localized. The frozen taxonomy has 23 Practice Types.

**Legacy decisions:** `PR-001`, `PR-003`

## D-020 — Assessment is a measurement model, not a test collection

**Decision:** Define reusable Assessment Types, evidence rules, confidence, and validity separately from concrete assessment items.

**Canonical owner:** `08-ASSESSMENT.md`

**Rationale:** Practice and assessment serve different purposes; concrete test items should not own mastery semantics. The frozen taxonomy has 7 Assessment Types.

**Legacy decisions:** `AM-001`, `AM-002`

## D-021 — Certification uses repeated evidence and confidence-aware productive scoring

**Decision:** Default certification policy requires repeated independent demonstrations; current default minimum confidence for uncertain mastery judgment is 0.80. These are calibratable policy values, not architectural constants.

**Canonical owner:** `08-ASSESSMENT.md`

**Rationale:** One-off performance creates false positives; productive AI judgment requires explicit uncertainty and empirical calibration.

**Legacy decision:** `AM-003`

## D-022 — Learner state model is semantic, not storage

**Decision:** Define mastery, knowledge, certification, review, and recommendation state without choosing a database or persistence architecture.

**Canonical owner:** `09-PROGRESSION.md`

**Rationale:** Runtime semantics are learning-domain truth; storage is implementation.

**Legacy decision:** `PG-001`

## D-023 — Per-skill band progression

**Decision:** Each IELTS skill certifies independently. Overall band is informational for progression rather than a hard gate.

**Canonical owner:** `09-PROGRESSION.md`

**Rationale:** Uneven IELTS profiles are normal. Synchronizing all four skills to the same band blocks legitimate progress.

**Legacy decision:** `PG-002`; this superseded an earlier synchronized four-skill gating rule.

## D-024 — Current certification can regress; history remains

**Decision:** If later valid evidence shows regression, current certification returns to in-progress and requires fresh re-certification; prior attainment remains in history.

**Canonical owner:** `09-PROGRESSION.md`

**Rationale:** Current learning state should reflect current capability without erasing historical attainment.

**Legacy provenance:** accepted Red-Team F2 resolution in the frozen legacy Blueprint.

## D-025 — Concrete content is reference-based

**Decision:** Learning Units, Practice Items, Assessment Items, Stimuli, and Feedback Artifacts reference canonical object IDs instead of embedding parallel domain definitions.

**Canonical owner:** `10-CONTENT-MODEL.md`

**Rationale:** Concrete content must be replaceable and generatable without changing canonical learning truth.

## Retired process decisions

Legacy repository-process and vendor-specific decisions such as Claude-specific operating rules, branch/worktree workflow guidance, and staged Blueprint document conventions are intentionally **not migrated as active learning/governance truth**.

They remain available in `archive/legacy-2026-07-16/` for historical provenance.
