# IELTS Learning System

This repository is self-describing. It contains the canonical IELTS learning specification and the canonical product/runtime design that translates that learning model into a real application.

## Read order

A new session should reconstruct the project in this order:

1. `CONSTITUTION.md` — governance, authority, ownership, naming, repository topology, and cross-language rules.
2. `OBJECTIVE.md` — why the project exists, what the complete product must define, and its success condition.
3. `spec/00-PRODUCT.md` — learning-product principles and boundaries.
4. `spec/01-LEARNER-MODEL.md` — the learner and epistemic learner-state requirements.
5. `spec/02-IELTS-MODEL.md` — the external IELTS reality the system must respect.
6. Read the relevant `spec/` owners for learning truth.
7. Read the relevant `design/` owner for product/runtime translation.
8. Use `spec/DECISIONS.md` only for rationale; it is not canonical authority.
9. Consult `research/` and `evidence/` only when provenance/validation is required.
10. Treat `archive/` as historical only.

## Authority

```text
USER
  ↓
CONSTITUTION.md
  ↓
OBJECTIVE.md
  ↓
spec/      canonical learning truth
  ↓
design/    canonical product/runtime translation
  ↓
contracts/ exact machine interface truth once materialized
```

`README.md` is navigation only.

A design document may consume and operationalize learning semantics, but it may not redefine a Skill, Band threshold, prerequisite, evidence rule, mastery rule, or progression rule owned by `spec/`.

## Active learning specification

```text
spec/
├── 00-PRODUCT.md
├── 01-LEARNER-MODEL.md
├── 02-IELTS-MODEL.md
├── 03-SKILLS.md
├── 04-KNOWLEDGE.md
├── 05-BANDS.md
├── 06-CURRICULUM.md
├── 07-PRACTICE.md
├── 08-ASSESSMENT.md
├── 09-PROGRESSION.md
├── 10-CONTENT-MODEL.md
├── 11-GLOSSARY.md
└── DECISIONS.md
```

## Active product/runtime design

```text
design/
├── 00-learning-experience.md
├── 01-skill-features.md
├── 02-practice-catalog.md
├── 03-media-youtube.md
├── 04-application-flows.md
├── 05-api.md
└── 06-implementation-stack.md
```

The active Markdown baseline is now **23 documents**:

```text
README.md               1
CONSTITUTION.md         1
OBJECTIVE.md            1
spec/                  13
  canonical owners     12
  decision rationale    1
design/                 7
--------------------------
TOTAL                  23
```

Markdown under `research/`, `evidence/`, and `archive/` is supporting or historical and does not count toward the active authority surface.

## Product baseline

The product now has explicit design for:

- the end-to-end learner journey;
- Quick / Standard / Deep daily study presets;
- 40 named skill/shared feature capabilities;
- 28 user-facing practice modes;
- Listening, Reading, Writing, and Speaking interaction flows;
- YouTube/media learning inspired by useful dictation/shadowing patterns while preserving platform/rights boundaries;
- diagnostic, review, remediation, readiness, and mock flows;
- a public Go Core API and bounded Python evaluator service;
- asynchronous Writing/Speaking evaluation;
- a concrete Go + Python + TypeScript framework allocation.

These counts are product design, not IELTS learning thresholds.

## Implementation language and unit baseline

The approved primary application languages are:

```text
Go
Python
TypeScript
```

Initial deployable ownership is:

```text
apps/web                 TypeScript / Next.js
services/core-api        Go / net/http + chi
services/evaluator       Python / FastAPI
```

The repository is not organized into top-level language silos.

Future implementation follows responsibility/deployable boundaries:

```text
apps/       user-facing deployables and clients
services/   independently runnable backend services/workers
packages/   reusable implementation libraries when justified
contracts/  language-neutral cross-unit interface definitions
tools/      repository/development/generation/release tooling
```

Cross-language semantics are never maintained as three handwritten copies. Genuine boundaries use one explicit machine-readable contract, while learning meaning remains owned by `spec/` and product/runtime behavior by `design/`.

## Overall product loop

```text
Goal
  ↓
Diagnostic
  ↓
Learner model
  ↓
Daily plan
  ↓
Learning session
  ↓
Attempt
  ↓
Observation / Evidence
  ↓
Mastery / Gap
  ↓
Next action
  ↺
```

The implementation flow is:

```text
Web (TypeScript)
  ↓
Core API (Go)
  ├── deterministic learning/product orchestration
  └── Evaluator (Python) for bounded AI/audio/text analysis
          ↓
      observations
          ↓
      Core API
          ↓
 evidence / progression / next plan
```

Python does not certify Band or advance learner state. The browser does not call the evaluator directly.

## Project boundary

The repository defines both:

1. **Learning truth** — what the learner must know/demonstrate and how mastery/progression works;
2. **Product/runtime design** — how the application lets the learner study, practice, use media, submit attempts, receive feedback, and move through the system.

It still does not require a particular cloud provider, database vendor, AI model provider, auth provider, payment provider, or deployment platform. Those choices may evolve behind the design/contract boundaries.

## Historical snapshot

The pre-refactor Blueprint is preserved under `archive/legacy-2026-07-16/` for provenance and forensic comparison only. If archived material conflicts with active `spec/` or `design/`, active canonical owners win.