STATUS: CANONICAL
OWNS: product identity, learning-product scope, product principles, variant boundary, delivery-overlay boundary, and learning-system promise boundary
DEPENDS_ON: ../OBJECTIVE.md
DOES_NOT_OWN: live IELTS format/scoring/delivery facts, learner state, skill decomposition, Band thresholds, curriculum, practice, assessment, progression, coverage status, or implementation architecture

# 00 — Product

## Product identity

This project specifies an evidence-based IELTS learning system whose primary value is learner capability and truthful readiness evidence, not feature count or content volume.

The learning specification should be complete enough that implementation can build delivery and technology around it without re-deciding the core learning model.

# Product principles

## Learning is the product

A feature is justified only when it improves the learner's ability to acquire, demonstrate, retain, transfer, or prepare to demonstrate target IELTS capability.

## Teach exactly what the target requires

Do not over-teach unrelated English merely because content is available. Do not under-teach required capability to reduce implementation scope.

## Evidence leads learning policy

External exam claims are grounded in authoritative IELTS evidence. Learning-policy claims should use credible learning evidence and label uncertainty where the evidence is weaker.

## Deliberate performance over passive consumption

Core activities have an explicit target, require a relevant learner cognitive/performance operation, and can produce feedback or evidence that changes a later decision.

## Mastery before certification

Time spent, lessons completed, or content consumed do not certify competence. Evidence/Progression owners define certification semantics.

## AI supports learning without becoming the standard

AI may tutor, generate, analyse, and evaluate within bounded roles. It does not become IELTS truth, Assessment truth, or Progression authority.

# Standard IELTS learning scope

The intended complete learning construct is:

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ official task/question families
+ Band-3→9 structured learning paths
```

Shared constructs stay shared. Variant-specific semantics exist only where IELTS externally differs.

Learning-model boundary:

- Listening is shared;
- Speaking is shared;
- Reading capabilities/question strategies are shared while variant-specific contexts/scoring remain external/contextual conditions;
- Academic Writing Task 1 uses visual-information capability;
- General Training Writing Task 1 uses letter purpose/audience/register/required-content capability;
- Writing Task 2 is substantially shared;
- shared Writing criteria/grammar/lexis/cohesion are not cloned per variant.

Academic may be released before General Training. Release order does not redefine canonical construct scope.

Current external variant facts are owned by `02-IELTS-MODEL.md`.

# Delivery-overlay boundary

The canonical learning requirement does not change merely because IELTS is delivered through a different permitted interface or setting.

A delivery mode may affect:

- exam-interface familiarity;
- response input mechanics;
- timing/navigation practice;
- accessibility/capture conditions;
- external eligibility/acceptance constraints.

It may not redefine:

- Skill identity;
- Knowledge requirements;
- Band quality standard;
- evidence truth for a learning claim except where the claim explicitly includes delivery-condition readiness.

Live IELTS delivery modes/eligibility are owned only by `02-IELTS-MODEL.md`. Product exam-readiness design consumes those facts rather than copying them here.

# Language/localization boundary

The canonical learning system is L1-agnostic.

First-language context may improve localization, explanations, examples, interference hypotheses, or remediation. It cannot create a different definition of IELTS competence.

# Human-support boundary

The core learning loop must not require mandatory human review for ordinary operation.

Human expert input may be an optional/required escalation only where the applicable Assessment/evaluator policy needs it for a particular consequence. This file does not define that evidence rule.

# Implementation boundary

Learning truth does not choose application architecture, persistence, model vendor, API design, cloud provider, or frontend framework.

Implementation consumes canonical semantics; code/protocol/storage do not become competing learning owners.

# Learning-model success

At the learning-model level, the system is successful when it can represent and govern a path from the structured Band-3 entry range toward Band 9 by:

- representing required Skills and enabling Knowledge;
- preserving genuine Academic/GT differences without duplicating shared constructs;
- enforcing genuine prerequisites;
- supporting acquisition, consolidation, retrieval, transfer, fluency, and exam readiness;
- preserving uncertainty rather than converting unknown to weakness;
- measuring capability through admissible evidence;
- progressing each skill independently;
- supporting exam preparation without treating exposure as certification.

# Model completeness vs product support

A model may be coherent while the executable product is still missing content, interactions, machine contracts, evaluator calibration, runtime, rights/privacy/security, reliability, cost controls, or empirical validation.

Current coverage/support semantics and status are owned by `../design/08-coverage-and-support.md`, not this learning-product owner.

# Promise boundary

The learning system may promise process properties it can actually control: target integrity, valid prerequisite/evidence rules, explainable next-action semantics, and truthful current state.

It must not claim that following the system guarantees a future external IELTS result.

# Non-goals

This file does not define pricing, acquisition, retention mechanics, concrete UI flow, provider selection, or technical implementation.