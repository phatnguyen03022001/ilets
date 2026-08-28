# Objective

## Purpose

Build a complete, evidence-based IELTS learning-system specification and product/runtime design that an implementation team or future reasoning session can consume without re-deciding major learning semantics, learner-route semantics, product-coverage semantics, external-provider boundaries, or first-order runtime responsibilities.

The primary learner is a self-directed IELTS candidate who has a real target but otherwise faces fragmented preparation: vocabulary/SRS, pronunciation, listening/shadowing, generic AI, media, and exam-practice tools may each perform useful local work while leaving the learner to reconcile diagnosis, priorities, dosage, evidence, and next actions manually. The product owns that coordination problem end to end:

```text
target
→ know current evidence
→ know exact gaps / unknowns
→ know what to learn next
→ perform the right learning / practice modality
→ collect truthful evidence
→ re-plan automatically
```

Individual tools or modalities remain replaceable capabilities inside this loop. The durable product differentiation is the governed target-to-evidence-to-next-action system, not possession of a particular AI model, media source, SRS algorithm, or exercise format.

The repository must define, with implementation-grade precision:

- what the learner must know and demonstrate;
- how external IELTS requirements map to capabilities without copying exam UI into the learner ontology;
- how target overall Band and/or per-skill minimums constrain planning;
- how diagnosis, planning, learning, practice, review, re-evidence, readiness, and mocks form one closed route toward the target;
- how learner agency changes eligible delivery without weakening Required prerequisites or target conditions;
- which learner-facing capabilities and practice experiences exist;
- how the ordinary learner route can teach, practise, explain, review, and adapt without mandatory teacher dependency while keeping AI inside bounded non-authoritative roles;
- how essentially every repeatable digital workflow materially useful to standard IELTS preparation is represented coherently without treating feature count as completeness;
- how every supported official task/question family reaches content, interaction, evidence, and progression paths;
- how attempts, observations, evidence, mastery/readiness, gaps, and next actions remain distinct;
- how product coverage/support is declared without unsupported completeness percentages;
- how external providers remain replaceable behind explicit capability boundaries;
- how runtime units divide responsibility without duplicating domain truth.

## Target outcome

The product should support a learner moving from a Band-3 structured entry point toward Band 9 across Listening, Reading, Writing, and Speaking while preserving:

- uneven per-skill profiles;
- explicit unknown, stale, and conflicting evidence states;
- overall and per-skill target constraints;
- variant-correct Academic and General Training learning routes;
- truthful distinction between learning progress, readiness, product support, and an actual external IELTS result.

The product may provide a strong governed route to a target. It must never imply that following the route guarantees a particular external test score.

## Intended complete standard-IELTS scope

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ official task/question families
+ Band-3→9 learning paths
+ target profile → diagnostic → learning → practice → assessment → readiness
```

Academic may be released before General Training, but release ordering must not be confused with construct completeness.

IELTS for UKVI Academic/General Training reuses the corresponding learning construct while adding external administrative/security conditions. One Skill Retake reuses the selected skill construct. IELTS Life Skills is a different Listening/Speaking-only pass/fail construct and is outside the current product-learning target unless explicitly added later.

## Scope

### Learning specification

The active learning specification covers:

- the four scored IELTS skills;
- enabling grammar, vocabulary, and phonology;
- stable Skill Leaves and Knowledge Objects;
- Band expectations and exit criteria;
- curriculum sequencing and prerequisites;
- learning mechanisms and reusable Practice Types;
- Assessment, evidence, readiness, and mastery semantics;
- Progression, regression, review, and certification semantics;
- conceptual content representation required to instantiate the system.

### Learner route

The product/runtime design must close this route:

```text
TargetProfile
→ diagnostic
→ learner model
→ target blockers / uncertainty / due review
→ eligible Daily Plan
→ coherent learning session
→ Attempt
→ Observation / EvidenceFact
→ Mastery / Readiness evaluation
→ GapEvaluation / ActionIntent
→ next plan
↺
```

The planner must preserve the learner's declared target until the learner changes it. Swap/Skip/Shorten/Change-skill controls may choose among eligible paths but cannot make a Required prerequisite, uncovered product condition, or unsupported target disappear.

### Product experience

The active design must provide concrete, understandable learner surfaces for:

- onboarding and `TargetProfile` setup;
- quick and fuller diagnostic entry paths;
- Today/plan, skills, practice, review, media, progress, and mock flows;
- Listening, Reading, Writing, and Speaking interactions;
- feedback, remediation, scaffold fading, transfer, re-evidence, fluency, review, and exam preparation;
- vocabulary/grammar/phonology acquisition, learner-saved study material, spaced retrieval, and later application in skill work;
- AI-supported ordinary tutoring and truthful Speaking/media interaction without making AI a learning/evidence authority;
- section and full mock flows;
- focused preparation for One Skill Retake without inventing a fifth Skill ontology.

Exact feature inventories, practice-mode inventories, durations, and UI defaults are owned by the relevant `design/` documents rather than this Objective.

### Coverage and support

The repository must distinguish:

```text
MODELLED
→ COVERED
→ SUPPORTED_FOR_PRODUCT
→ VALIDATED
```

A scoped target is `COVERED` only when its executable chain is complete with no blocking CoverageGap. Product support additionally requires release-critical content, evaluator/calibration, contract, provider, rights/privacy/security, reliability, accessibility/capture-quality, operational, and cost gates.

No global “100% IELTS” percentage may hide a missing required condition.

The current mutable coverage declaration is owned only by `design/08-coverage-and-support.md`; this Objective does not duplicate that status.

### Media

The design covers media as an eligible source substrate without making media providers part of learning truth. Transcript/rights eligibility, safe media failure behavior, and the prohibition on assumed arbitrary extraction/download must be explicit.

### Application/runtime design

The design must provide:

- one learner-facing public product API boundary;
- bounded asynchronous productive-skill evaluation;
- non-overlapping runtime ownership across the approved application languages;
- explicit machine-readable contracts at genuine cross-unit boundaries before independent implementations diverge;
- idempotency, lifecycle, pending/unavailable states, and recovery semantics;
- provider-neutral external capability boundaries;
- repository-wide verification across every affected runtime and contract boundary.

Exact framework versions, route inventories, provider selections, and deployment details are owned downstream by their `design/` or future machine-contract owners.

## Non-goals

The repository does not currently freeze:

- a final cloud/hosting provider;
- a final PostgreSQL provider;
- a final object-storage provider;
- a final identity provider;
- a final payment provider;
- a final AI/LLM/STT/TTS provider;
- final production database tables;
- Kubernetes/multi-region architecture without demonstrated need;
- pixel-perfect visual design or brand system;
- vendor-specific coding-agent instructions.

These choices may later be selected behind canonical provider/runtime boundaries. They must not redefine learning truth or product-support semantics.

## Quality target

Optimize for:

1. correctness;
2. completeness;
3. internal consistency;
4. evidence and traceability;
5. explicit semantic ownership;
6. implementation clarity;
7. learner-facing explainability;
8. target-route integrity;
9. measurable coverage/support gates;
10. minimal duplicated policy across documents, services, languages, and providers.

## Success definition

The project objective is satisfied when:

- every major learning semantic has one canonical `spec/` owner;
- every major first-order product/runtime semantic has one canonical `design/` owner;
- Bands 3–9 form coherent learning progressions across all four skills;
- Academic and General Training reach complete standard-IELTS modelling without duplicating shared constructs;
- a `TargetProfile` can express real overall/per-skill constraints and material exam-delivery constraints when relevant;
- the planner can derive a valid route from current learner state to every supported target condition;
- Required prerequisites cannot be bypassed by UI choice or ranking;
- every supported official task/question family maps to concrete feature, practice, content, assessment, and transition paths;
- every supported target has no blocking CoverageGap and has a versioned TargetSupportDeclaration;
- practice covers acquisition, consolidation, retrieval, transfer, fluency, exam readiness, and targeted remediation;
- Assessment preserves Attempt → Observation → EvidenceFact → Readiness/Progression separation;
- productive evaluator output cannot directly certify Band or advance learner state;
- media adds learning value while respecting external platform/rights constraints;
- external providers remain replaceable behind declared capability boundaries;
- cross-language interfaces have one machine-contract authority;
- product copy distinguishes learner evidence state from product support state;
- a new reasoning session can reconstruct the system from repository content without hidden prompts or stale review documents.