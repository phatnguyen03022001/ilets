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

## Initial test variant

The fully specified initial variant is **IELTS Academic**.

Shared concepts must be modeled once. General Training is a future variant overlay, not a parallel learning system.

The architecture must allow General Training differences to be introduced without restructuring shared domains:

- Listening is shared;
- Speaking is shared;
- Reading has variant-specific passage/scoring characteristics;
- Writing Task 1 is variant-specific;
- Writing Task 2 is substantially shared at the learning-model level.

The authoritative external variant semantics are owned by `02-IELTS-MODEL.md`.

## Language and localization boundary

The canonical learning specification is L1-agnostic and written in English.

A learner's first language may inform optional localization, explanations, examples, or interference-aware remediation, but L1-specific material must not alter the canonical learning requirement.

Localization is therefore an overlay on the learning system, not a canonical domain in the active specification.

## Delivery boundary

The learning model is delivery-agnostic. It must not require a specific channel such as:

- self-study only;
- tutor-led only;
- classroom only;
- mobile only;
- web only.

A future product may choose one or more delivery models without changing canonical learning semantics.

## Human-support boundary

The core learning system must be able to operate without mandatory human review. Human expert input may be an optional enhancement, especially for productive-skill assessment or low-confidence cases.

This is a product availability constraint only. The exact evidence and confidence rules are owned by `08-ASSESSMENT.md`.

## Implementation boundary

This specification is implementation-agnostic. It does not choose:

- application architecture;
- persistence technology;
- AI provider;
- model vendor;
- API design;
- cloud platform;
- frontend framework.

Implementation consumes canonical objects and rules; it does not become a second learning authority.

## Product success

The product succeeds, at the learning-model level, when it can support a learner from the structured Band-3 entry range toward Band 9 by:

- representing what the learner must know and demonstrate;
- teaching prerequisite knowledge before dependent capability when truly required;
- adapting within a band without bypassing required outcomes;
- training capability through appropriate learning phases;
- measuring mastery with sufficient evidence;
- progressing each skill independently;
- supporting exam preparation without falsely certifying higher mastery.

## Non-goals of this spec

This file does not define detailed business strategy, pricing, acquisition, retention mechanics, UI flows, or technical implementation. Those decisions may be made later as long as they consume rather than contradict the learning specification.
