STATUS: CANONICAL
OWNS: canonical terminology and term meanings used across the active learning specification
DEPENDS_ON: 00-PRODUCT.md, 01-LEARNER-MODEL.md, 02-IELTS-MODEL.md, 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md, 06-CURRICULUM.md, 07-PRACTICE.md, 08-ASSESSMENT.md, 09-PROGRESSION.md, 10-CONTENT-MODEL.md
DOES_NOT_OWN: behavioral rules, object inventories, product-support statuses, API/runtime terminology, or external-provider semantics

# 11 — Glossary

This file owns learning-spec vocabulary only. Behavioral detail remains in the owning domain document.

# Authority/documentation

**Canonical** — authoritative within one explicitly owned semantic domain.

**Canonical owner** — the one active document authorized to define a semantic.

**Supporting** — rationale, research, evidence, or navigation with no independent authority.

**Historical / archived** — superseded material retained for provenance with no active authority.

**SSOT** — single source of truth implemented through one explicit owner per semantic, not one giant file.

**Semantic ownership** — responsibility for defining a rule, object, threshold, or concept.

**Reference** — use of an owned semantic by stable identity/link without redefining it.

# IELTS

**IELTS Band** — external IELTS proficiency score on the 0–9 scale; this learning program structurally teaches Bands 3–9.

**Section Band** — Band score for Listening, Reading, Writing, or Speaking.

**Overall Band** — IELTS-reported rounded average of the four section Bands; an external result/target constraint, not a synchronized learning-progression gate.

**Academic** — one of the two in-scope standard IELTS variants. It shares Listening/Speaking with General Training and differs where Reading/Writing externally differ.

**General Training / GT** — the other in-scope standard IELTS variant. It shares Listening/Speaking with Academic and has variant-specific Reading/Writing conditions.

**Delivery mode** — how an IELTS test is administered. Delivery mode may change interaction/eligibility but is not a new Skill, Band, or standard IELTS variant.

**Variant overlay** — variant-specific semantics layered on a shared learning core only where Academic and General Training genuinely differ.

**Official task/question family** — stable external identity for an IELTS question type, Writing task, or Speaking part defined by `02-IELTS-MODEL.md`. It is an exam-family identity, not a Skill Leaf.

**Task 1 / Task 2** — the two Writing tasks; exact Academic/GT external semantics are owned by `02-IELTS-MODEL.md`.

**TA / TR / CC / LR / GRA / FC** — Task Achievement / Task Response / Coherence and Cohesion / Lexical Resource / Grammatical Range and Accuracy / Fluency and Coherence.

# Learning domains

**Knowledge** — enabling language knowledge such as grammar, vocabulary, and phonology.

**Skill** — capability demonstrated in Listening, Reading, Writing, or Speaking.

**Band threshold** — required quality boundary for a target skill/task at a target Band.

**Curriculum** — canonical orchestration of when Skill/Knowledge targets are sequenced.

**Practice** — how canonical capability is trained.

**Assessment** — how demonstrated performance is measured and interpreted as evidence for claims.

**Progression** — how current learner interpretation, certification, gaps, review, and next-action intent change over time.

**Content Model** — conceptual representation of replaceable concrete learning, tutoring, measurement, and feedback instances.

# Canonical objects

**Skill Leaf** — atomic independently useful capability with a stable `L-*`, `R-*`, `W-*`, or `S-*` identity. A leaf may serve multiple official task/question families.

**Knowledge Object** — atomic enabling concept with a stable `K-GRA-*`, `K-VOC-*`, or `K-PHON-*` identity.

**Curriculum Node** — canonical orchestration step with stable `C-B<band>-<nn>` identity.

**Learning Mechanism** — learning process such as retrieval, contrast, guided production, scaffold fading, or transfer; stable `LM-*` identity.

**Practice Type** — reusable executable activity pattern with stable `PT-*` identity; may instantiate multiple Learning Mechanisms.

**Assessment Type** — reusable measurement role with stable `AT-*` identity.

**Content Context** — stable semantic identity for the IELTS variant/section/task context in which concrete content operates, for example a GT Reading section context or Academic Writing Task 1. It does not create another Skill.

**Content Presentation Class** — stable identity for a materially different subformat/stimulus presentation inside an official family when that distinction is necessary to verify real content coverage. It is not a scored task or Skill.

# Concrete/runtime representations

**Learning Unit** — delivery-neutral concrete grouping around a Curriculum Node.

**Stimulus** — passage, recording, visual, prompt, model, example, or other material presented to the learner.

**ScaffoldingProfile** — material support available during performance/learning whose presence changes inference scope.

**ExposureContext** — prior item/stimulus/feedback exposure and material novelty/variation context.

**Practice Item** — concrete instance of a Practice Type targeting canonical objects and, where material, an official family/context/presentation.

**Assessment Item** — concrete measurement instance for a defined claim/target scope.

**Content coverage manifest** — machine-checkable implementation index describing which canonical targets, official families, contexts, material presentation classes, interactions, evaluation routes, rights states, and release activations actual content provides. It is implementation evidence, not canonical learning truth.

**Error Pattern** — reusable tutoring hypothesis about a recurring mistake; not proof that a learner has the error.

**Remediation Pattern** — reusable intervention strategy referencing targets, mechanisms, practice types, scaffolds, and success checks.

**Feedback Artifact** — runtime guidance derived from performance; not canonical learning truth.

# Evidence/epistemic state

**Attempt** — learner event against a Practice or Assessment Item.

**Observation** — normalized measurement of what happened before evidence admission.

**EvidenceFact** — an Observation admitted for a specific claim under Assessment eligibility policy.

**EvidenceRequirement** — versionable claim-scoped logical conditions evidence must satisfy before the claim is supported.

**Evidence eligibility** — determination that an Observation may be used for a specific claim/purpose.

**Inference scope** — exact claim/context/generalization an EvidenceFact may legitimately support.

**MasteryEstimate** — current uncertainty-aware interpretation of a Skill Leaf/Knowledge Object from admissible evidence; not raw evidence storage.

**ReadinessEvaluation** — claim outcome: `INSUFFICIENT_EVIDENCE`, `CONFLICTING_EVIDENCE`, `STALE_EVIDENCE`, `NOT_YET_SUPPORTED`, or `SUPPORTED`.

**Insufficient evidence** — required information is missing; not weakness.

**Conflicting evidence** — material admissible evidence supports incompatible interpretations.

**Stale evidence** — prior support is too old for a present claim; not regression by itself.

**Not yet supported** — current relevant evidence exists and is below the scoped requirement.

**Supported** — current admissible evidence satisfies the scoped claim; never a guarantee of a future official IELTS result.

**Confidence** — calibrated uncertainty attached to a measurement/evaluation when decision-relevant; model self-report alone is not calibration.

# Progression

**Certification** — current internal recognition that a scoped skill-Band learning claim is `SUPPORTED`; not an official IELTS result.

**Certification history** — immutable-enough history of prior point-in-time certification.

**Regression** — later admissible evidence establishes that a previously supported current capability is now below requirement.

**Re-certification** — fresh support after established regression under normal Assessment policy.

**GapEvaluation** — classification of what kind of learning/evidence problem currently exists.

**ActionIntent** — semantic objective for the next action before choosing a concrete mechanism/mode/item.

**Required prerequisite** — hard dependency whose absence makes dependent learning ineffective and cannot reasonably be bypassed adaptively.

**Recommended prerequisite** — useful sequencing dependency that does not hard-block learning.

**Independent** — no prerequisite dependency.

**Adaptive sequencing** — personalization of path while preserving required targets and evidence standards.

**Same outcomes, different paths** — learners may follow different valid routes but the same scoped skill/variant/Band claim uses the same standard.

# Learning phases

**Acquisition** — initial construction of new capability/knowledge.

**Consolidation** — stabilization through focused varied use.

**Retrieval** — active recall/reproduction after delay.

**Transfer** — application in a new/materially different context.

**Fluency** — increased automaticity/speed/rhythm without sacrificing target quality.

**Exam Readiness** — practice under exam-like independence, timing, integration, stamina, and material delivery conditions.

**Exam Preparation** — learner mode that may expose higher-demand tasks without automatically certifying mastery.

# Assessment roles

**Formative assessment** — measurement primarily used to guide learning.

**Diagnostic assessment** — measurement used to reduce decision-relevant uncertainty.

**Readiness assessment** — measurement under target-like conditions used to estimate current target performance.

**Mastery portfolio** — `AT-05`, cumulative evaluation of admissible EvidenceFacts against the relevant EvidenceRequirement.

**Official Evidence** — fact/standard directly established by an authoritative IELTS source.

**Evidence-Based Interpretation** — educational interpretation closely grounded in evidence but not a direct official statement.

**Blueprint Inference** — pedagogical detail beyond what official IELTS explicitly establishes.

# Representation boundaries

**Single representation** — each canonical semantic/object is defined by one owner and referenced elsewhere.

**Canonical/runtime separation** — reusable specification truth remains separate from learner/attempt-specific state.

**Localization overlay** — L1/locale-specific explanation/remediation layered on an unchanged canonical IELTS requirement.