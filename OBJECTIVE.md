# Objective

## Purpose

Build a complete, evidence-based IELTS learning system specification and product/runtime design that a future implementation team or reasoning session can consume without making new major learning-system, learner-route, product-coverage, external-provider-boundary, or first-order architecture decisions.

The repository must define, with implementation-grade precision:

- what the learner must know and demonstrate;
- how IELTS external requirements map to capabilities without copying exam UI into the learner ontology;
- how a learner sets a target overall Band and/or per-skill minimums;
- how diagnosis, planning, learning, practice, review, re-evidence, readiness, and mocks form one closed route toward that target;
- how learner agency changes eligible delivery without silently weakening Required prerequisites or target conditions;
- which concrete skill features and practice experiences exist;
- how every official task/question family is represented in the product surface;
- how media such as eligible YouTube content is used safely;
- how attempts, observations, evidence, mastery, readiness, gaps, and next actions flow;
- how product support/coverage is declared without unsupported completeness percentages;
- how external providers are isolated behind explicit capability boundaries;
- how TypeScript, Go, and Python divide runtime responsibility without duplicating domain truth.

## Target outcome

The product should support a learner moving from a Band-3 structured entry point toward Band 9 across Listening, Reading, Writing, and Speaking while preserving:

- uneven per-skill profiles;
- explicit unknown/stale/conflicting evidence states;
- target overall/per-skill constraints;
- truthful distinction between learning progress, readiness support, product coverage, and actual external IELTS result.

The product may provide a strong governed route to a target. It must never imply that following the route guarantees a particular external test score.

## Intended complete standard-IELTS scope

The intended complete product-learning scope is:

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ official task/question families
+ Band-3→9 learning paths
+ target profile → diagnostic → learning → practice → assessment → readiness flow
```

IELTS Academic may be supported before General Training, but release ordering must not be confused with construct completeness.

IELTS for UKVI Academic/General Training reuses the same learning construct while adding external administrative/security conditions. One Skill Retake reuses the selected skill construct. IELTS Life Skills is a different Listening/Speaking-only pass/fail construct and is outside the current product-learning target unless explicitly added later.

## Scope

### Learning specification

The active learning specification covers:

- the four scored IELTS skills;
- enabling grammar, vocabulary, and phonology;
- canonical Skill Leaves and Knowledge Objects;
- Band expectations and exit criteria;
- curriculum sequencing and prerequisites;
- learning mechanisms and reusable Practice Types;
- Assessment, evidence, readiness, and mastery semantics;
- Progression, regression, review, and certification semantics;
- conceptual content representation required to instantiate the system.

### Learner route

The active design must cover:

```text
TargetProfile
→ provisional diagnostic
→ learner model
→ target blockers / uncertainty / due review
→ Daily Plan
→ coherent learning session
→ Attempt
→ Observation / Evidence
→ Mastery / Readiness
→ GapEvaluation / ActionIntent
→ next plan
↺
```

The planner must preserve the learner's declared target until the learner changes it. Swap/Skip/Shorten/Change-skill controls may choose among eligible paths but cannot make a Required prerequisite, uncovered product condition, or unsupported target disappear.

### Product experience

The active design covers:

- Today-first navigation;
- onboarding and TargetProfile setup;
- quick/full diagnostics;
- skill/practice/review/media/progress/mock surfaces;
- Quick, Standard, and Deep plan-size presets;
- 40 named feature capabilities;
- 28 user-facing practice modes;
- Listening, Reading, Writing, and Speaking interaction flows;
- feedback, remediation, scaffold fading, transfer, re-evidence, review, and exam preparation;
- section and full mock flows;
- focused preparation for One Skill Retake without inventing a fifth Skill ontology.

### Coverage and support

The repository must distinguish:

```text
MODELLED
→ COVERED
→ SUPPORTED_FOR_PRODUCT
→ VALIDATED
```

A scoped target is `COVERED` only when its executable chain is complete with no blocking CoverageGap. A product-support declaration additionally requires release-critical content, evaluator/calibration, provider, rights/privacy/security, reliability, and cost gates.

No global “100% IELTS” percentage may hide a missing required condition.

### Media

The design covers:

- YouTube as an embedded source where platform policy permits;
- transcript/rights eligibility;
- dictation, shadowing, retell/comprehension, and vocabulary-mining uses;
- media removal/source-failure behavior;
- prohibition on treating arbitrary unofficial download/transcription as an assumed capability.

### Application/runtime design

The active design covers:

- one learner-facing public API boundary;
- asynchronous productive-skill evaluation;
- TypeScript web, Go Core API, and Python Evaluator responsibility allocation;
- cross-language machine-readable contracts;
- idempotency, pending evaluation, SSE result delivery, and failure semantics;
- canonical third-party capability/provider boundaries;
- native verification expectations for all three primary languages.

## Current design-state truth

The repository is a specification/design repository. It is not yet a running learning product.

At this stage:

```text
Academic semantic model        strong / modelled
Academic product coverage      not yet COVERED
General Training               partial
runtime implementation         not started
production support declaration none
validated target-band outcome  not established
```

`design/08-coverage-and-support.md` owns the detailed current declaration and blockers.

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

These choices may later be selected behind the canonical provider/runtime boundaries. They must not redefine learning truth or product-support semantics.

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
10. minimal duplicated policy across services/languages/providers.

## Success definition

The project objective is satisfied when:

- every major learning semantic has one canonical `spec/` owner;
- every major first-order product/runtime semantic has one canonical `design/` owner;
- Bands 3–9 form coherent learning progressions across all four skills;
- a TargetProfile can express the learner's real overall/per-skill constraints;
- the planner can derive a valid route from current learner state to every supported target condition;
- Required prerequisites cannot be bypassed by UI choice;
- every supported official task/question family maps to concrete features, practice, content, and evidence paths;
- every supported target has no blocking CoverageGap and has a versioned TargetSupportDeclaration;
- Academic + General Training can reach complete standard-IELTS coverage without duplicating shared constructs;
- practice covers acquisition, consolidation, retrieval, transfer, fluency, exam readiness, and targeted remediation;
- Assessment preserves Attempt → Observation → EvidenceFact → Readiness/Progression separation;
- productive evaluator output cannot directly certify Band or advance learner state;
- the app provides an understandable concrete practice catalog rather than a vague adaptive promise;
- YouTube/media adds independent learning value while respecting external platform/rights constraints;
- each external provider is replaceable behind a declared capability boundary;
- Go, Python, and TypeScript have non-overlapping primary runtime responsibilities;
- cross-language interfaces have one contract authority;
- product copy distinguishes modelled, covered, supported, validated, and learner-ready states;
- a new reasoning session can reconstruct the complete system from repository content alone.