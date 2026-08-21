STATUS: SUPPORTING
PURPOSE: Evidence-led architecture/learning completeness audit. This file does not override canonical owners.
DATE: 2026-08-22

# Deep architecture and IELTS learning coverage audit

## Audit objective

Test whether a new implementation team can reconstruct the intended product without inventing major IELTS-learning, learner-route, state-machine, variant, evidence, or first-order runtime decisions.

The audit uses the repository authority hierarchy. A finding is closed only by changing the canonical owner that owns the semantic; this file is only a review record.

## Executive verdict

The repository already has a strong architecture of authority and a coherent closed learning loop. The main remaining design-time risks are not missing broad concepts; they are incomplete variant closure and places where prose permits two reasonable implementations.

Current audit verdict:

- Academic learning model: strong.
- General Training: externally identified but not fully modelled through Skill/Band/Curriculum/Product feature closure.
- Practice/assessment/progression: strong conceptual separation.
- Developer execution precision: medium-high; several decision contracts remain semantic rather than deterministic.
- Runtime/product support: intentionally not implemented and therefore not COVERED/SUPPORTED_FOR_PRODUCT.
- Empirical target-band outcome validation: not established.

## P0 — canonical learning gaps

### G-001 — General Training Writing Task 1 is not a closed learning construct

Owner chain:

`spec/02-IELTS-MODEL.md` → `spec/03-SKILLS.md` → `spec/05-BANDS.md` → `spec/06-CURRICULUM.md` → `design/01-skill-features.md` → `design/08-coverage-and-support.md`

Problem:

The objective requires complete standard IELTS Academic + General Training. Current Skills/Bands/Curriculum are Academic-first and do not give GT Task 1 stable capability identity, Band threshold interpretation, or explicit route orchestration.

Required closure:

- represent GT Task 1 audience/purpose/required bullet-point coverage;
- represent personal/semi-formal/formal register control;
- preserve shared Writing criteria instead of cloning the whole Writing model;
- bind GT-specific capabilities into the curriculum only when `TargetProfile.test_variant = GENERAL_TRAINING`;
- expose a concrete GT letter-planning/product path;
- require both GT Task 1 and Task 2 evidence for a GT Writing Band claim.

Acceptance test:

Given a General Training target, a planner can identify, teach, practice, assess, re-evidence, and explain every GT Task 1 requirement without using an Academic visual-overview capability as a substitute.

### G-002 — General Training Reading context progression is not route-explicit

Owner chain:

`spec/02-IELTS-MODEL.md` → `spec/06-CURRICULUM.md` / `spec/10-CONTENT-MODEL.md` → `design/01-skill-features.md` → `design/08-coverage-and-support.md`

Problem:

GT Reading uses three sections with distinct context classes: everyday survival material, workplace material, then one longer general-interest text. Shared Reading capabilities/question families are reusable, but the content/transfer route must deliberately cover all three context classes.

Required closure:

- preserve shared Reading Skill Leaves;
- add variant-context coverage requirements rather than duplicating Reading skills;
- require unseen transfer across all GT section context classes before a GT Reading coverage claim;
- use the GT scoring conversion for GT claims.

Acceptance test:

A learner cannot become product-ready for GT Reading using only Academic-style passage assets, even if the question-type engine is shared.

### G-003 — 2026 delivery reality needs an explicit exam-readiness boundary

Owner chain:

`spec/02-IELTS-MODEL.md` → `design/00-learning-experience.md` / `design/01-skill-features.md`

External fact rechecked 2026-08-22:

IELTS announced that from mid-2026 paper-based IELTS will no longer be offered, with exact rollout timing varying by market; selected markets may provide a Writing-on-Paper option. This changes delivery preparation, not the underlying language construct.

Required closure:

- computer-delivered interaction is the default exam-readiness baseline;
- optional handwriting practice is a delivery overlay only when relevant to the learner's booked market/test option;
- do not encode obsolete paper answer-transfer mechanics as universal learning truth.

## P1 — developer ambiguity gaps

### G-101 — eligibility and ranking are not the same decision

The route docs define eligible next actions and state that time/accessibility/preferences rank choices, but they do not provide a required separation between:

1. target-condition expansion;
2. hard eligibility filtering;
3. ActionIntent routing;
4. candidate generation;
5. candidate ranking;
6. final explanation/reason code.

Required architecture invariant:

A ranker may reorder eligible candidates. It may never make an ineligible candidate eligible, remove a Required prerequisite, convert product coverage into learner weakness, or alter evidence truth.

### G-102 — state transitions need explicit legal-transition tests

The semantic states exist, but implementation should encode legal transition matrices for at least:

- LearningSession;
- Attempt;
- Evaluation;
- DiagnosticRun;
- MockRun;
- MediaSource/MediaLesson analysis;
- current BandCertificationState.

Invalid transitions must fail closed and remain idempotent under retry.

### G-103 — evidence requirement is deliberately claim-scoped but needs materialized policy before high-consequence support

The canonical Assessment model correctly rejects a universal attempt count/confidence threshold. Before a target is SUPPORTED_FOR_PRODUCT, every high-consequence claim must nevertheless resolve to a versioned `EvidenceRequirement` with executable conditions.

Acceptance test:

For any supported `(variant, skill, band)` claim, Core API can answer exactly why evidence is sufficient/insufficient/conflicting/stale without a hidden model heuristic.

### G-104 — content coverage must be manifest-driven

A task family is not covered because a UI exists. Product coverage requires real content/templates/generators, interaction behavior, scoring/evaluator paths, transfer assets, rights/provenance, and readiness assets.

Recommended implementation artifact when content work begins:

`content-manifest` keyed by stable canonical IDs and variant/context tags, validated against Coverage conditions.

### G-105 — exact cross-language contracts are not materialized

`design/05-api.md` defines semantic API intent, but exact OpenAPI is intentionally absent. This is acceptable in design state, but implementation must not begin parallel handwritten DTO design in TypeScript/Go/Python.

Gate:

Materialize `contracts/http/openapi.yaml` before cross-language implementation diverges. Generated bindings/validators are derived, not authority.

## P1 — learning-system completeness checks

The following learning modes are already represented conceptually and must remain covered by the implementation/content layer:

- acquisition;
- consolidation;
- retrieval;
- spaced review where appropriate;
- deliberate correction/revision;
- scaffold fading;
- transfer to unseen/materially different contexts;
- interleaving where discrimination benefits;
- fluency/automaticity work after quality is stable;
- exam-condition timing, stamina, and integration;
- diagnostic sampling;
- re-evidence for stale/conflicting/insufficient claims;
- productive criterion feedback;
- receptive deterministic scoring;
- error-pattern remediation;
- vocabulary/grammar/phonology enabling knowledge;
- media-supported practice where rights/eligibility permit;
- full and section mocks;
- One Skill Retake focused preparation using the existing four-skill ontology.

No implementation may collapse these into one generic `practice` state when the distinction changes evidence, scheduling, scaffolding, or next-action semantics.

## P2 — architecture closure gates before production support

A target cannot move to `SUPPORTED_FOR_PRODUCT` until all applicable gates are executable and release-qualified:

1. canonical construct model;
2. Skill/Knowledge/Band/Curriculum route;
3. learner experience and feature path;
4. practice/intervention path;
5. assessment and versioned EvidenceRequirement;
6. progression/next-action path;
7. content/assets/templates/generators;
8. exact API/internal contracts;
9. evaluator calibration for productive claims;
10. rights/privacy/security;
11. provider failure/fallback semantics;
12. reliability/idempotency/concurrency;
13. cost/abuse controls;
14. accessibility and capture-quality handling;
15. variant-specific mock/readiness path;
16. observable audit/provenance;
17. empirical validation when claiming `VALIDATED`.

A missing required gate is a CoverageGap, not a percentage deduction.

## Developer non-negotiables

- Unknown is not weak.
- Stale is not regression.
- One wrong answer is not automatically an ability gap.
- Practice completion is not mastery.
- Guided success is not independent evidence.
- Same-item retry is not transfer evidence.
- Evaluator output is an Observation candidate, not certification.
- Python never advances learner state.
- Browser never calls Python as a second product API.
- Product CoverageGap is never rendered as learner weakness.
- Overall Band is a target/result constraint, not a synchronized progression gate.
- Academic and General Training share constructs where IELTS shares them and diverge only where the exam diverges.
- A content/template count never substitutes for condition-based coverage.

## External facts used in this audit

Rechecked against official IELTS material on 2026-08-22:

- General Training Reading has three sections of increasing difficulty with everyday, workplace, and longer general-interest contexts.
- General Training Writing Task 1 is a >=150-word letter responding to a situation; the required style may be personal, semi-formal, or formal and the prompt supplies three content bullet points.
- Writing Task 2 contributes twice as much as Task 1.
- One Skill Retake reuses one existing skill; eligibility includes a full IELTS-on-computer test, one retake per original test, and completion within 60 days, subject to local availability/acceptance.
- IELTS announced computer delivery as the standard rollout from mid-2026, with market timing variation and an optional Writing-on-Paper offer in selected markets.

Canonical external-exam facts belong in `spec/02-IELTS-MODEL.md`; this file only records why the audit demanded changes.
