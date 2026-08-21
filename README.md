# IELTS Learning System

This repository is self-describing. It contains canonical IELTS learning truth, canonical product/runtime design, and supporting research/evidence used to justify those owners.

## Read order

1. `CONSTITUTION.md` — governance, authority, ownership, naming, topology, and cross-language rules.
2. `OBJECTIVE.md` — product scope and success condition.
3. `spec/00-PRODUCT.md`, `spec/01-LEARNER-MODEL.md`, `spec/02-IELTS-MODEL.md`.
4. Read relevant `spec/` owners for learning truth.
5. Read relevant `design/` owners for product/runtime translation.
6. Read `design/08-coverage-and-support.md` before making any completeness/support claim.
7. Use `spec/DECISIONS.md` for rationale only.
8. Use `research/` and `evidence/` for provenance/validation/audit only.
9. Treat `archive/` as historical only.

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
contracts/ exact machine-interface truth once materialized
```

`README.md` is navigation only.

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
├── 06-implementation-stack.md
├── 07-third-party-services.md
└── 08-coverage-and-support.md
```

The active Markdown baseline remains **25 documents**:

```text
README.md               1
CONSTITUTION.md         1
OBJECTIVE.md            1
spec/                  13
design/                 9
--------------------------
TOTAL                  25
```

`research/`, `evidence/`, and `archive/` do not count toward the active authority surface. Audit material under `evidence/` is supporting only and cannot override a canonical owner.

## Standard IELTS learning scope

The canonical learning model targets:

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ official task/question families
+ Band-3→9 learning paths
+ target → diagnosis → learning → practice → evidence → progression → readiness
```

Shared constructs stay shared. Variant-specific differences are explicit:

- Listening shared;
- Speaking shared;
- Reading capability shared, with Academic/GT context and scoring differences preserved;
- Academic Writing Task 1 uses visual-information capability;
- General Training Writing Task 1 uses letter recipient/purpose/register/required-point capability;
- Writing Task 2 is substantially shared.

IELTS Life Skills is a separate construct and remains outside current product scope.

## Product baseline

The design has explicit contracts for:

- `TargetProfile` with overall and/or per-skill target constraints;
- diagnostic → Today plan → practice/review/re-evidence → readiness loop;
- strong route recommendations with learner Swap/Skip/Shorten agency that cannot bypass Required prerequisites;
- staged planner behavior: target support → unresolved conditions → Gap/ActionIntent → hard eligibility → candidate generation → ranking → explainable plan;
- Academic/General Training variant resolution;
- 67 canonical Skill Leaves, including dedicated GT Writing Task-1 capabilities;
- 46 enabling Knowledge Objects;
- 44 base Curriculum Nodes plus deterministic variant overlay;
- 40 named skill/shared feature capabilities;
- 28 user-facing practice modes;
- acquisition, consolidation, retrieval, review, remediation, scaffold fading, transfer, fluency, exam readiness, diagnostics, mocks, and re-evidence semantics;
- claim-scoped, versionable `EvidenceRequirement` semantics;
- explicit Attempt/Evaluation/session lifecycle invariants;
- content variant/context and future manifest requirements;
- API and async evaluator flows;
- Go + Python + TypeScript framework allocation;
- exact-contract gate before cross-language implementation;
- explicit third-party capability/provider boundaries;
- condition-based coverage/support declarations instead of unsupported “100%” claims.

## Current coverage truth

Do **not** read the repository as claiming full production support.

Current design-time status is owned by `design/08-coverage-and-support.md` and is summarized as:

```text
Academic semantic model          strong / MODELLED
General Training semantic model  modelled at learning/design level
Academic product coverage        not yet COVERED
General Training product coverage not yet COVERED
runtime implementation           not started
validated target-band outcome    not established
```

Major remaining closure work is implementation/content/calibration rather than another broad learning ontology:

- executable content/assets/templates and coverage manifest;
- exact OpenAPI/internal machine contracts;
- productive evaluator benchmarking/calibration;
- runtime persistence/idempotency/reliability/security/privacy/cost controls;
- accessibility/capture-quality verification;
- release-qualified TargetSupportDeclarations;
- empirical outcome validation before any `VALIDATED` claim.

## Implementation baseline

```text
apps/web                 TypeScript / Next.js
services/core-api        Go / net/http + chi
services/evaluator       Python / FastAPI
```

Cross-language semantics are never maintained as three handwritten copies. Machine boundaries use explicit contracts; learning meaning remains in `spec/`, product/runtime behavior in `design/`.

## Overall product loop

```text
TargetProfile
  ↓
Diagnostic
  ↓
Learner model
  ↓
Gap / uncertainty / due review
  ↓
Daily Plan
  ↓
Learning Session
  ↓
Attempt
  ↓
Observation / Evidence
  ↓
Mastery / Readiness
  ↓
Next Action
  ↺
```

The product can guarantee the integrity of this process and truthfulness of its own evidence/support states. It must not guarantee that following the plan necessarily produces a specific external IELTS score.

## Deep audit record

`evidence/architecture-learning-audit-2026-08-22.md` records the supporting audit that identified the variant, developer-ambiguity, content, evidence, contract, and lifecycle gaps addressed by this architecture pass.

It is an audit record only. When it conflicts with `spec/` or `design/`, the canonical owner wins.

## Historical snapshot

The pre-refactor Blueprint is preserved under `archive/legacy-2026-07-16/`. If archived material conflicts with active `spec/` or `design/`, active canonical owners win.
