# Objective

## Purpose

Build a complete, evidence-based IELTS learning system specification and product/runtime design.

The repository must define, with enough precision for later implementation:

- what a learner needs to know;
- what the learner needs to demonstrate;
- how those capabilities are trained and assessed;
- how the learner progresses from Band 3 through Band 9;
- how the learner experiences the product day to day;
- which practice experiences the application provides;
- how skill-specific features work;
- how media such as eligible YouTube content can be used safely and usefully;
- how attempts, feedback, evidence, progress, and next actions flow through the application;
- how the public/internal APIs are shaped semantically;
- how Go, Python, and TypeScript responsibilities are divided without duplicating domain truth.

## Target outcome

A future implementation team or reasoning session should be able to consume the repository without making new major learning-system or first-order product-architecture decisions.

The system should support a learner moving from a Band-3 structured entry point toward Band 9 across Listening, Reading, Writing, and Speaking while preserving uneven per-skill profiles and explicit uncertainty when evidence is missing, stale, or conflicting.

## Scope

### Learning specification

The active learning specification covers:

- IELTS Academic as the fully specified initial test variant;
- a structure that can accept General Training differences without restructuring shared learning semantics;
- the four IELTS skills;
- enabling knowledge such as grammar, vocabulary, and phonology;
- band expectations and exit criteria;
- curriculum sequencing and prerequisites;
- practice mechanisms and reusable Practice Types;
- assessment, evidence, readiness, and mastery semantics;
- learner progression, regression, review, and certification semantics;
- conceptual content representations required to instantiate the learning model.

### Product experience

The active design covers:

- onboarding and goal setup;
- quick/full diagnostics;
- Today/home/skill/practice/review/media/progress/mock surfaces;
- Quick, Standard, and Deep daily study-plan presets;
- 40 named user-facing feature capabilities;
- 28 user-facing practice modes;
- complete Listening, Reading, Writing, and Speaking product flows;
- feedback, remediation, re-evidence, review, and exam-preparation flows;
- full mocks and readiness flows.

### Media

The active design covers:

- YouTube as an embedded media source where platform policy permits;
- transcript/rights eligibility;
- dictation, shadowing, retell/comprehension, and vocabulary-mining media uses;
- media source failure/removal handling;
- prohibition on treating arbitrary unofficial media extraction as an assumed product capability.

### Application/runtime design

The active design covers:

- one learner-facing public API boundary;
- asynchronous productive-skill evaluation;
- TypeScript web, Go Core API, and Python evaluator responsibility allocation;
- cross-language machine-readable contracts;
- idempotency, pending evaluation, SSE result delivery, and failure semantics;
- native verification expectations for all three primary languages.

## Non-goals

The repository does not currently freeze:

- a specific cloud provider;
- a specific relational/object-storage vendor;
- a specific authentication provider;
- a specific payment provider;
- a specific AI/LLM/speech provider;
- final production database tables;
- final deployment topology, Kubernetes strategy, or multi-region architecture;
- pricing, commercialization, acquisition, or go-to-market strategy;
- pixel-perfect visual design or brand system;
- vendor-specific coding-agent instructions.

Those choices may later be added behind the canonical product/runtime boundaries. They must not redefine learning truth.

## Quality target

Optimize for:

1. correctness;
2. completeness;
3. internal consistency;
4. evidence and traceability;
5. explicit semantic ownership;
6. implementation clarity;
7. learner-facing explainability;
8. minimal duplicated policy across languages/services.

Do not optimize the learning model for implementation speed at the expense of learning correctness. Do not optimize architecture aesthetics at the expense of a coherent learner flow.

## Success definition

The project objective is satisfied when:

- every major learning semantic has one canonical `spec/` owner;
- every major first-order product/runtime semantic has one canonical `design/` owner;
- Bands 3–9 form a coherent learning progression for all four skills;
- skill requirements resolve to the knowledge they require;
- curriculum paths reference canonical skill and knowledge objects rather than redefining them;
- practice covers acquisition, consolidation, retrieval, transfer, fluency, exam readiness, and targeted learning mechanisms;
- assessment separates Attempt, Observation, EvidenceFact, claim interpretation, readiness, and progression;
- progression supports independent per-skill advancement and honest unknown/stale/conflicting evidence states;
- the app provides a concrete and understandable practice catalog rather than a vague "adaptive" promise;
- the learner journey is coherent from goal → diagnostic → plan → practice → evidence → next action;
- YouTube/media integration adds independent learning value while preserving external platform/rights constraints;
- the API has a single public product boundary and explicit internal evaluator boundary;
- Go, Python, and TypeScript have non-overlapping primary runtime responsibilities;
- cross-language interfaces have one contract authority instead of handwritten mirror schemas;
- Academic IELTS is fully represented and General Training can be added as a variant overlay without restructuring shared domains;
- a new reasoning session can reconstruct the complete learning/product/runtime system from `README.md` → `CONSTITUTION.md` → this Objective → relevant `spec/` → relevant `design/`.