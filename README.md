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
8. Use `research/` and `evidence/` for provenance/validation only.
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

The active Markdown baseline is **25 documents**:

```text
README.md               1
CONSTITUTION.md         1
OBJECTIVE.md            1
spec/                  13
design/                 9
--------------------------
TOTAL                  25
```

`research/`, `evidence/`, and `archive/` do not count toward the active authority surface.

## Imported founder-decision research

`research/lenbands-founder-decisions/` contains a byte-preserved import of the normalized LenBands founder set:

```text
64  platform/reliability
26  identity/privacy/access
12  AI/evaluation/cost
45  product/requirements
40  learning interventions
54  evidence/readiness
40  learner experience
24  economics/entitlements
20  coverage/support
---
325 numbered decisions
```

The separate content-rights/provenance block is imported too but is explicitly outside the 325 count.

Importing a decision does not adopt it. Active adoption happens only through the correct `spec/` or `design/` owner.

## Product baseline

The product has explicit design for:

- `TargetProfile` with overall and/or per-skill target constraints;
- diagnostic → Today plan → practice/review/re-evidence → readiness loop;
- strong route recommendations with learner Swap/Skip/Shorten agency that cannot bypass Required prerequisites;
- 40 named skill/shared feature capabilities;
- 28 user-facing practice modes;
- Listening, Reading, Writing, and Speaking flows;
- YouTube/media learning with rights/transcript eligibility;
- API and async evaluator flows;
- Go + Python + TypeScript framework allocation;
- explicit third-party capability/provider boundaries;
- condition-based coverage/support declarations instead of unsupported “100%” claims.

## Current coverage truth

Do **not** read the repository as claiming full production support yet.

Current design-time status is summarized in `design/08-coverage-and-support.md`:

```text
Academic semantic model        strong / modelled
Academic product coverage      not yet COVERED
General Training               partial
runtime implementation         not started
validated target-band outcome  not established
```

The intended complete standard-IELTS scope is Academic + General Training. IELTS Life Skills is a separate construct and is not currently in product scope.

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

## Historical snapshot

The pre-refactor Blueprint is preserved under `archive/legacy-2026-07-16/`. If archived material conflicts with active `spec/` or `design/`, active canonical owners win.