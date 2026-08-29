# IELTS Learning System

This repository is self-describing. It contains canonical IELTS learning truth, canonical product/runtime design, and non-canonical research/evidence/history used to support or explain those owners.

## Read order

1. `CONSTITUTION.md` — governance, authority, ownership, naming, topology, and cross-language rules.
2. `OBJECTIVE.md` — project intent, scope, quality target, and success condition.
3. `spec/00-PRODUCT.md`, `spec/01-LEARNER-MODEL.md`, `spec/02-IELTS-MODEL.md` — product-learning boundary, learner representation, and external IELTS reality.
4. Read the relevant `spec/` owner for learning truth.
5. Read the relevant `design/` owner for product/runtime translation.
6. Read `design/08-coverage-and-support.md` before making any completeness, support, or readiness-of-product claim.
7. Use `spec/DECISIONS.md` for rationale only.
8. Use `research/` and `evidence/` for provenance/validation only.

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

`README.md` is navigation only. It does not own product status, canonical inventory counts, learning policy, runtime policy, or implementation readiness.

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

Use the owner named in each file's `OWNS` metadata. `DECISIONS.md` is supporting rationale, not a canonical learning owner.

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

Use these owners by responsibility:

- learner journey and `TargetProfile` → `design/00-learning-experience.md`;
- skill-facing product capabilities → `design/01-skill-features.md`;
- user-facing practice modes → `design/02-practice-catalog.md`;
- media/YouTube product semantics → `design/03-media-youtube.md`;
- planner/system flows and runtime lifecycle → `design/04-application-flows.md`;
- API resource/operation semantics → `design/05-api.md`;
- deployables, languages, frameworks, canonical materialization, verification → `design/06-implementation-stack.md`;
- external provider boundaries → `design/07-third-party-services.md`;
- current coverage/support declarations and gates → `design/08-coverage-and-support.md`.

## Implementation navigation

Before implementing a bounded slice:

1. read `CONSTITUTION.md` and `OBJECTIVE.md`;
2. read the relevant `spec/` semantic owners;
3. read the relevant `design/` product/runtime owners;
4. read `design/05-api.md` before materializing a shared HTTP boundary;
5. read `design/06-implementation-stack.md` for runtime ownership, canonical-registry materialization, and repository verification rules;
6. read `design/08-coverage-and-support.md` before making an implementation-readiness, coverage, or product-support claim.

Navigation warnings:

- `MODELLED` does not mean implementation-ready or `COVERED`;
- do not independently hand-author equivalent DTO/schema truth across parallel runtimes before the exact machine contract exists;
- generated registries/bindings are derived artifacts, not canonical authority.

The detailed rules remain owned by the referenced canonical documents; this section only points to them.

## Supporting and historical material

```text
research/   current non-canonical research/provenance; never canonical

evidence/   source/evidence records supporting current canonical claims; never canonical
```

Historical reviews, migration reports, superseded structures, and retired workflow artifacts live in Git history rather than the active tree unless a current provenance, legal, or recovery need requires them to remain tracked.

## Current project state

Do not infer current coverage or implementation readiness from this README.

- project intent → `OBJECTIVE.md`;
- current product coverage/support → `design/08-coverage-and-support.md`;
- external IELTS facts → `spec/02-IELTS-MODEL.md`;
- exact machine contracts → `contracts/` once materialized.

If a summary conflicts with its owner, the owner wins.