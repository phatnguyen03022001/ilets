STATUS: CANONICAL
OWNS: product identity, learning-product scope, product principles, variant and delivery boundaries
DEPENDS_ON: ../OBJECTIVE.md
DOES_NOT_OWN: learner state, IELTS scoring rules, skill decomposition, band thresholds, curriculum, practice, assessment, progression, implementation architecture

# 00 — Product

## Product identity

This project specifies an evidence-based IELTS learning system whose primary product value is learning outcome, not feature count.

The learning specification is intended to be complete enough that a future implementation can build delivery, UX, and technology around it without re-deciding the core learning model.

## Product principles

### Learning is the product

Features are valuable only when they improve a learner's ability to demonstrate the target IELTS capability.

### Teach exactly what the target requires

The system should neither over-teach unrelated English content nor under-teach required IELTS capability. Breadth is justified by target competence, not by content availability.

### Evidence leads learning policy

Major pedagogical choices should be grounded in official IELTS material where the claim concerns the exam and in credible learning-science evidence where the claim concerns learning. When evidence is weak, the specification should expose the uncertainty rather than turn an inference into a fact.

### Deliberate practice over passive consumption

Learning activities should have an explicit objective, require relevant learner performance, and produce feedback or evidence that can change the next learning action.

### Mastery before certification

Progression is based on demonstrated competence rather than time spent or content consumed. The exact progression semantics are owned by `09-PROGRESSION.md`.

### AI supports learning without becoming the standard

AI may provide tutoring, generation, feedback, and assessment support, but canonical learning truth remains the specification and external IELTS standards. AI uncertainty must be surfaced where it materially affects assessment decisions; the assessment policy is owned by `08-ASSESSMENT.md`.

## Standard IELTS scope

The intended complete learning construct is:

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ official task/question families
+ Band-3→9 structured learning paths
```

Academic and General Training share Listening, Speaking, much of Reading capability, shared Writing criteria, and Task-2 learning capability. They diverge only where the external exam construct diverges.

Canonical rules:

- Listening is shared;
- Speaking is shared;
- Reading capability/question strategy is shared, while variant-specific corpus/context distribution and raw-score conversion are preserved;
- Academic Writing Task 1 uses visual-information capability;
- General Training Writing Task 1 uses letter-purpose/audience/register capability;
- Writing Task 2 is substantially shared at the learning-model level;
- variant-specific capability is represented as an overlay, not as a second parallel learning system.

Academic may be the first production release. Release order must never be confused with learning-construct completeness or used to erase General Training from the canonical model.

The authoritative external variant semantics are owned by `02-IELTS-MODEL.md`.

## Delivery boundary

The learning model is delivery-agnostic. It must not require a specific channel such as self-study, tutor-led, classroom, mobile, or web.

Exam-delivery preparation is different from learning truth. Current live-test delivery facts are owned by `02-IELTS-MODEL.md`; product exam-readiness surfaces translate them without changing the underlying capability standard.

For the 2026 design baseline, computer-delivered exam interaction is the default readiness target because IELTS announced the transition away from paper-based IELTS. A market-specific Writing-on-Paper option, where available, is a delivery overlay rather than a new learning construct.

## Language and localization boundary

The canonical learning specification is L1-agnostic and written in English.

A learner's first language may inform optional localization, explanations, examples, or interference-aware remediation, but L1-specific material must not alter the canonical learning requirement.

Localization is therefore an overlay on the learning system, not a canonical domain in the active specification.

## Human-support boundary

The core learning system must be able to operate without mandatory human review. Human expert input may be an optional enhancement, especially for productive-skill assessment or low-confidence cases.

This is a product availability constraint only. The exact evidence and confidence rules are owned by `08-ASSESSMENT.md`.

## Implementation boundary

This specification is implementation-agnostic. It does not choose application architecture, persistence technology, AI provider, model vendor, API design, cloud platform, or frontend framework.

Implementation consumes canonical objects and rules; it does not become a second learning authority.

## Product success

The product succeeds, at the learning-model level, when it can support a learner from the structured Band-3 entry range toward Band 9 by:

- representing what the learner must know and demonstrate;
- preserving Academic/General Training differences without duplicating shared constructs;
- teaching prerequisite knowledge before dependent capability when truly required;
- adapting within a band without bypassing required outcomes;
- training capability through appropriate learning phases;
- measuring mastery with sufficient evidence;
- progressing each skill independently;
- supporting exam preparation without falsely certifying higher mastery.

## Product coverage is not model completeness

A canonical construct can be fully modelled while the product is still uncovered because content, evaluator calibration, interaction, runtime, reliability, rights/privacy/security, cost, or validation gates are missing.

The statuses `MODELLED`, `COVERED`, `SUPPORTED_FOR_PRODUCT`, and `VALIDATED` are owned by `../design/08-coverage-and-support.md` and must never be collapsed into a marketing percentage.

## Non-goals of this spec

This file does not define detailed business strategy, pricing, acquisition, retention mechanics, UI flows, or technical implementation. Those decisions may be made later as long as they consume rather than contradict the learning specification.
