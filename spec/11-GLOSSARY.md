STATUS: CANONICAL
OWNS: canonical terminology and term meanings used across the active specification
DEPENDS_ON: 00-PRODUCT.md, 01-LEARNER-MODEL.md, 02-IELTS-MODEL.md, 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md, 06-CURRICULUM.md, 07-PRACTICE.md, 08-ASSESSMENT.md, 09-PROGRESSION.md, 10-CONTENT-MODEL.md
DOES_NOT_OWN: the behavioral rules or object inventories behind those terms

# 11 — Glossary

This file is the canonical vocabulary for the active specification. Behavioral detail remains in the owning domain spec.

## Authority and documentation

**Canonical** — authoritative within one explicitly owned semantic domain.

**Canonical owner** — the one active spec authorized to define a semantic.

**Supporting** — useful rationale, research, evidence, or navigation with no independent authority.

**Historical / archived** — preserved prior material with no active authority.

**SSOT** — single source of truth distributed by explicit domain ownership rather than duplicated across files.

**Semantic ownership** — responsibility for defining a rule, object, threshold, or concept.

**Reference** — use of a canonical semantic by link/ID without redefining it.

## IELTS

**IELTS Band** — proficiency score on the IELTS 0–9 scale; the structured learning program targets Bands 3–9.

**Section band** — band score for Listening, Reading, Writing, or Speaking.

**Overall band** — IELTS-reported combination of the four section bands; informational for learning progression, not a progression gate.

**Academic** — IELTS Academic variant; the fully specified initial variant.

**General Training / GT** — IELTS General Training variant; shared Listening/Speaking with variant differences in Reading/Writing.

**Task 1 / Task 2** — the two Writing tasks; exact variant semantics are owned by `02-IELTS-MODEL.md`.

**TA / TR / CC / LR / GRA / FC** — Task Achievement / Task Response / Coherence and Cohesion / Lexical Resource / Grammatical Range and Accuracy / Fluency and Coherence.

## Canonical learning domains

**Knowledge** — enabling language knowledge such as grammar, vocabulary, and phonology.

**Skill** — capability demonstrated in Listening, Reading, Writing, or Speaking.

**Band threshold** — required quality of performance for a target skill/task at a target band.

**Curriculum** — canonical orchestration of when Skill and Knowledge objects are sequenced.

**Practice** — how canonical capability is trained.

**Assessment** — how demonstrated capability is measured and how observations become admissible evidence for claims.

**Progression** — how learner state, current mastery interpretation, certification, gaps, and next actions change over time.

**Content Model** — conceptual representation of concrete learning, tutoring, measurement, and feedback instances.

## Canonical objects

**Skill Leaf** — atomic capability unit identified by stable `W-*`, `S-*`, `L-*`, or `R-*` ID.

**Knowledge Object** — atomic enabling-knowledge unit identified by stable `K-GRA-*`, `K-VOC-*`, or `K-PHON-*` ID.

**Curriculum Node** — orchestration step identified by stable `C-B<band>-<nn>` ID.

**Learning Mechanism** — learning process such as active retrieval, contrast, guided production, scaffold fading, or transfer variation; identified by stable `LM-*` ID and owned by `07-PRACTICE.md`.

**Practice Type** — reusable executable activity pattern identified by stable `PT-*`; may instantiate one or more Learning Mechanisms.

**Assessment Type** — reusable measurement strategy identified by stable `AT-*`.

## Concrete/supporting representations

**Learning Unit** — delivery-neutral grouping of concrete content around a Curriculum Node.

**Stimulus** — material presented to the learner such as a passage, recording, visual, prompt, or example.

**ScaffoldingProfile** — explicit representation of material support available during learning/performance.

**ExposureContext** — representation of prior item/stimulus/feedback exposure and material novelty/variation dimensions.

**Practice Item** — concrete instance of a Practice Type targeting canonical objects.

**Assessment Item** — concrete instance of an Assessment Type for a defined claim/target scope.

**Error Pattern** — reusable supporting hypothesis/pattern describing a recurring mistake or misconception; not proof a learner has that error.

**Remediation Pattern** — reusable tutoring strategy referencing targets, mechanisms, practice types, scaffolds, and success checks.

**Feedback Artifact** — runtime guidance tied to learner performance; not canonical learning truth.

## Evidence and epistemic state

**Attempt** — learner-instance event against a Practice or Assessment Item.

**Observation** — normalized measurement of what happened in an Attempt before evidence admission.

**EvidenceFact** — an Observation admitted for a specific claim/purpose under Assessment eligibility rules.

**EvidenceRequirement** — claim-scoped logical conditions that evidence must satisfy before the claim is supported.

**Evidence eligibility** — decision that an Observation may be used for a specific claim/purpose.

**Inference scope** — the exact claim/context/generalization that an EvidenceFact can legitimately support.

**MasteryEstimate** — current, derived, uncertainty-aware interpretation of capability from admissible evidence; not a raw evidence container.

**ReadinessEvaluation** — claim-scoped Assessment outcome: `INSUFFICIENT_EVIDENCE`, `CONFLICTING_EVIDENCE`, `STALE_EVIDENCE`, `NOT_YET_SUPPORTED`, or `SUPPORTED`.

**Insufficient evidence** — required information is missing; not evidence of weakness.

**Conflicting evidence** — material valid evidence supports incompatible interpretations and must be resolved rather than averaged away.

**Stale evidence** — historical support exists but is not current enough for a present claim; not proof of regression.

**Not yet supported** — current relevant evidence exists and falls below the scoped requirement.

**Supported** — current admissible evidence satisfies the scoped claim; never a guarantee of a future official IELTS result.

**Confidence** — calibrated uncertainty associated with a measurement/evaluation when it can affect a decision; model self-report alone is not calibration.

## Learning and progression

**Certification** — current internal recognition that a skill-band learning claim is `SUPPORTED`; not an official IELTS result.

**Certification history** — point-in-time record of prior internal certifications retained across later state change.

**Regression** — later admissible evidence shows a previously supported capability is now below the required current threshold.

**Re-certification** — fresh support/certification after established regression using the normal Assessment policy.

**GapEvaluation** — classification of what kind of learning/evidence problem exists, such as ability, prerequisite, evidence, conflict, staleness, scaffold dependence, transfer, fluency, or exam-condition gap.

**ActionIntent** — semantic next-action objective derived from learner state/GapEvaluation before choosing a Learning Mechanism or Practice Type.

**Required prerequisite** — hard dependency whose absence makes dependent learning ineffective and cannot reasonably be bypassed adaptively.

**Recommended prerequisite** — beneficial ordering dependency that does not hard-block learning.

**Independent** — no prerequisite relationship.

**Adaptive sequencing** — personalization of order, mechanism, practice, difficulty, review, scaffold, or pacing while preserving canonical targets.

**Same outcomes, different paths** — learners may follow different routes but must meet the same target standard for the same skill/band.

## Learning phases

**Acquisition** — initial construction of new knowledge or capability.

**Consolidation** — stabilization through focused varied use.

**Retrieval** — active recall/reproduction after delay.

**Transfer** — application in new or materially different contexts.

**Fluency** — increased speed/automaticity/rhythm without sacrificing target quality.

**Exam Readiness** — practice under exam-like timing, integration, stamina, and independence conditions.

**Exam Preparation** — learner mode that may expose higher-demand tasks without automatically certifying mastery.

## Assessment roles

**Formative assessment** — measurement primarily used to guide ongoing learning.

**Diagnostic assessment** — measurement used to reduce uncertainty or classify the kind of gap.

**Readiness assessment** — measurement under target/exam-like conditions used to estimate current performance.

**Mastery portfolio** — `AT-05`, cumulative claim evaluation over admissible EvidenceFacts.

**Blueprint Inference** — pedagogical claim beyond what official IELTS explicitly states.

**Evidence-Based Interpretation** — interpretation closely grounded in evidence but not a direct official statement.

**Official Evidence** — fact/standard directly established by an authoritative IELTS source.

## Representation boundaries

**Single representation** — each canonical object is defined once and referenced elsewhere by stable identity.

**Canonical/runtime separation** — distinction between reusable specification truth and learner/attempt-specific state.

**Variant overlay** — variant-specific semantics layered on shared concepts without duplicating the shared core.

**Localization overlay** — locale/L1-specific explanation/remediation layered on the L1-agnostic canonical requirement.
