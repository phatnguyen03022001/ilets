# IELTS Learning Blueprint

This repository is self-describing. It contains the canonical learning specification for an IELTS learning system and the supporting history needed to explain how that specification was reached.

## Read order

A new session should reconstruct the project in this order:

1. `CONSTITUTION.md` — governance, authority, ownership, and change rules.
2. `OBJECTIVE.md` — why the project exists, its scope, and its success condition.
3. `spec/00-PRODUCT.md` — product-level boundaries and principles.
4. `spec/01-LEARNER-MODEL.md` — the learner the system represents.
5. `spec/02-IELTS-MODEL.md` — the external IELTS reality the system must respect.
6. Read only the canonical domain specs relevant to the current question, following each file's `DEPENDS_ON` metadata.
7. Use `spec/DECISIONS.md` only to understand rationale. It is not a source of canonical truth.
8. Consult `research/` and `evidence/` only when provenance or validation is required.
9. Treat `archive/` as historical only.

## Authority

```text
CONSTITUTION.md
      ↓
OBJECTIVE.md
      ↓
spec/*.md canonical domain owners
```

`README.md` is navigation only. It has no authority over the Constitution, Objective, or canonical specs.

Canonical truth lives in the spec that owns the semantic. A canonical spec may reference another canonical spec, but it must not restate semantics owned by that other spec.

## Active specification

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

The structural baseline is 16 active Markdown documents: this file, `CONSTITUTION.md`, `OBJECTIVE.md`, and the 13 files under `spec/`. Markdown under `research/`, `evidence/`, and `archive/` is supporting or historical and is not part of the active-authority count.

## Project boundary

The active specification defines the learning domain: what learners need to know, what they need to demonstrate, how learning is sequenced and practiced, how mastery is assessed, and how progression works from Band 3 through Band 9.

It does not define production architecture, databases, APIs, infrastructure, deployment, pricing, go-to-market strategy, or vendor-specific agent behavior.

## Historical snapshot

The pre-refactor Blueprint is preserved under `archive/legacy-2026-07-16/`. It is retained for provenance and forensic comparison only. If archived material conflicts with the active specification, the active specification wins.
