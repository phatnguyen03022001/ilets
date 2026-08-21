STATUS: CANONICAL
OWNS: canonical terminology and term meanings used across the active specification
DEPENDS_ON: 00-PRODUCT.md, 01-LEARNER-MODEL.md, 02-IELTS-MODEL.md, 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md, 06-CURRICULUM.md, 07-PRACTICE.md, 08-ASSESSMENT.md, 09-PROGRESSION.md, 10-CONTENT-MODEL.md
DOES_NOT_OWN: the behavioral rules or object inventories behind those terms

# 11 — Glossary

This file is the canonical vocabulary for the active specification. Definitions are intentionally concise; behavioral detail remains in the owning domain spec.

## Authority and documentation

**Canonical** — authoritative within one explicitly owned semantic domain.

**Canonical owner** — the one active spec authorized to define a semantic.

**Supporting** — useful rationale, research, evidence, or navigation with no independent authority.

**Historical / archived** — preserved prior material with no active authority.

**SSOT** — single source of truth. In this repository SSOT is distributed by explicit domain ownership, not by copying all truth into one file.

**Semantic ownership** — responsibility for defining a particular rule, object, threshold, or concept.

**Reference** — use of a canonical object or semantic by link/ID without redefining it.

## IELTS

**IELTS Band** — proficiency score on the IELTS 0–9 scale; the structured learning program in this Blueprint targets Bands 3–9.

**Section band** — the band score for one of Listening, Reading, Writing, or Speaking.

**Overall band** — the IELTS-reported combination of the four section bands according to the official averaging/rounding rule; informational for learning progression, not a progression gate.

**Academic** — IELTS Academic variant; the fully specified initial variant of this Blueprint.

**General Training / GT** — IELTS General Training variant; shares Listening and Speaking with Academic but differs in Reading and Writing where specified by IELTS.

**Task 1** — first Writing task; Academic uses visual-information description, General Training uses a letter/situation response.

**Task 2** — essay response to a point of view, argument, or problem.

**TA** — Task Achievement, Writing Task-1 criterion.

**TR** — Task Response, Writing Task-2 criterion.

**CC** — Coherence and Cohesion.

**LR** — Lexical Resource.

**GRA** — Grammatical Range and Accuracy.

**FC** — Fluency and Coherence, Speaking criterion.

## Canonical learning domains

**Knowledge** — what must be known: enabling language concepts such as grammar, vocabulary, and phonology.

**Skill** — what must be demonstrated in Listening, Reading, Writing, or Speaking.

**Band threshold** — the required quality of performance for a skill/task at a target IELTS band.

**Curriculum** — canonical orchestration of when Skill and Knowledge objects are sequenced.

**Practice** — how canonical capability is trained.

**Assessment** — how demonstrated capability is measured and how evidence validity/sufficiency is judged.

**Progression** — how learner runtime state advances, regresses, certifies, and selects next actions.

**Content Model** — conceptual representation of concrete learning/assessment instances that reference canonical definitions.

## Canonical objects

**Skill Leaf** — atomic capability unit, identified by a stable `W-*`, `S-*`, `L-*`, or `R-*` ID.

**Knowledge Object** — atomic enabling-knowledge unit, identified by a stable `K-GRA-*`, `K-VOC-*`, or `K-PHON-*` ID.

**Curriculum Node** — canonical orchestration step, identified by a stable `C-B<band>-<nn>` ID.

**Practice Type** — reusable training pattern, identified by a stable `PT-*` ID.

**Assessment Type** — reusable measurement strategy, identified by a stable `AT-*` ID.

**Learning Unit** — delivery-neutral grouping of concrete content around a Curriculum Node.

**Practice Item** — concrete instance of a Practice Type targeting canonical Skill/Knowledge objects.

**Assessment Item** — concrete instance of an Assessment Type targeting canonical Skill/Knowledge/Band semantics.

**Stimulus** — source material presented to the learner, such as a passage, recording, visual, prompt, or example.

**Feedback Artifact** — runtime feedback tied to an attempt and canonical targets; not canonical truth.

**Evidence Record** — normalized assessment output tied to a learner attempt and canonical target.

## Learning and mastery

**Mastery** — reliably demonstrated criterion-referenced capability, not time spent or content consumed.

**Mastery state** — current Skill Leaf runtime state: `not_started`, `practicing`, `emerging`, or `mastered`.

**Knowledge state** — current Knowledge Object runtime state: `not_acquired`, `learning`, or `acquired`.

**Certification** — current recognition that sufficient valid evidence satisfies a skill-band exit standard.

**Certification history** — point-in-time record of prior certifications retained even if current capability later regresses.

**Regression** — later valid evidence showing a previously held capability is no longer held reliably enough for the current state.

**Re-certification** — fresh certification after regression, using the normal evidence standard rather than restoring status automatically.

**Exit criterion** — observable target that must be met to leave a task/skill band state; threshold semantics are owned by `05-BANDS.md`.

## Prerequisites and adaptation

**Required prerequisite** — hard dependency whose absence makes dependent learning ineffective and cannot reasonably be worked around adaptively.

**Recommended prerequisite** — beneficial dependency used for ordering but not a hard gate.

**Independent** — no prerequisite relationship.

**Adaptive sequencing** — runtime personalization of order, difficulty, practice, review, or pacing while preserving canonical targets and Required prerequisites.

**Same outcomes, different paths** — principle that learners may follow different routes but must meet the same target standard for the same skill/band.

## Learning phases

**Acquisition** — initial construction of new knowledge or capability.

**Consolidation** — stabilization through focused varied use.

**Retrieval** — active recall/reproduction after delay.

**Transfer** — application in new or less scaffolded contexts.

**Fluency** — development of speed/automaticity without sacrificing the target quality standard.

**Exam Readiness** — practice phase using exam-like conditions for timing, integration, familiarity, and stamina.

**Exam Preparation** — learner mode that may expose higher-demand tasks for readiness/diagnosis without certifying higher mastery.

## Assessment

**Formative assessment** — measurement used primarily to guide ongoing learning.

**Diagnostic assessment** — measurement used to locate current strengths, gaps, and uncertainty.

**Readiness assessment** — exam-like measurement estimating current test performance; non-certifying by itself.

**Mastery portfolio** — `AT-05`, accumulated sufficient evidence used to support certification.

**Evidence sufficiency** — requirement that the evidence set contains enough independent valid demonstrations for the target claim.

**Confidence** — calibrated degree of certainty attached to a measurement judgment when uncertainty can affect the decision.

**Blueprint Inference** — pedagogical claim derived beyond what the official IELTS source explicitly states.

**Evidence-Based Interpretation** — interpretation closely grounded in external evidence but not a direct official statement.

**Official Evidence** — fact or standard directly established by an authoritative IELTS source.

## Representation boundaries

**Single representation** — each canonical object is defined once and referenced elsewhere by stable ID.

**Canonical/runtime separation** — distinction between reusable Blueprint definitions and per-learner/per-attempt state.

**Variant overlay** — variant-specific semantics layered on shared canonical concepts without duplicating the shared core.

**Localization overlay** — L1- or locale-specific explanation/remediation layered on the L1-agnostic canonical learning requirement.
