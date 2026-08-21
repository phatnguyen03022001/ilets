# CLAUDE.md

# Project

**Project:** IELTS Blueprint

**Purpose**

Build the complete Blueprint for an IELTS learning application before any production implementation.

The Blueprint is the Single Source of Truth (SSOT) for all learning decisions.

**Why**

A single authoritative Blueprint prevents inconsistent learning logic, duplicated decisions, and implementation drift.

---

# Working Model

Founder decides **WHAT**.

Claude Code decides **HOW**.

When ambiguity exists:

* Stop.
* Explain the ambiguity.
* Ask clarification questions.

Never silently assume business decisions.

**Why**

Incorrect assumptions become incorrect specifications, which are significantly more expensive to correct than asking questions early.

---

# Primary Objective

Your primary responsibility is **not writing production code**.

Your primary responsibility is building a complete, internally consistent, evidence-based Blueprint.

Implementation begins only after the Blueprint has been reviewed and frozen.

**Why**

The product's competitive advantage is the quality of its learning Blueprint, not implementation speed.

---

# Blueprint Contract

The repository must contain the following top-level structure.

```
blueprint/
├── product/
├── learning/
├── curriculum/
├── bands/
├── skills/
├── knowledge/
├── practice/
├── assessment/
├── progress/
├── review/
└── glossary/
```

You may create additional subdirectories whenever justified.

Do not create new top-level Blueprint categories without explaining:

* why they are necessary;
* why they cannot fit into an existing category.

**Why**

A stable top-level structure keeps the Blueprint predictable for both humans and AI while allowing controlled growth.

---

# Research Rules

Prefer evidence in the following order whenever possible:

1. Official IELTS publications
2. Cambridge materials
3. British Council
4. IDP
5. High-quality educational references
6. Other reputable sources

When sources disagree:

* compare the evidence;
* explain the differences;
* recommend the best-supported decision;
* record remaining uncertainty.

Never invent unsupported learning requirements.

Clearly distinguish between:

* evidence;
* inference;
* assumption.

**Why**

The Blueprint should be evidence-driven rather than opinion-driven.

---

# Discovery Rules

Before writing any Blueprint section:

* identify ambiguity;
* ask clarification questions;
* wait for Founder decisions when required.

Never guess:

* product scope;
* business priorities;
* learning philosophy;
* acceptance criteria.

If confidence is low, explicitly state it.

**Why**

Discovery is part of the work. Resolving ambiguity early produces higher-quality specifications.

---

# Quality Rules

Every Blueprint should define, where applicable:

* Purpose
* Scope
* Learning Objectives
* Prerequisites
* Required Knowledge
* Required Skills
* Practice Strategy
* Assessment
* Progression Rules
* Dependencies
* Constraints
* Out of Scope
* Open Questions

Avoid duplication across Blueprint sections.

---

# Consistency

Continuously verify:

* no duplicated concepts;
* no conflicting requirements;
* no circular dependencies;
* no missing prerequisites;
* no unnecessary overlap between bands;
* consistent terminology across the repository.

Perform gap analysis before considering any Blueprint complete.

When inconsistencies are discovered:

1. identify them;
2. explain the impact;
3. recommend a resolution;
4. update all affected Blueprint sections.

**Why**

Internal consistency is more important than document completeness.

---

# Implementation Boundary

Until the Blueprint is approved and frozen, do not:

* write production code;
* design production architecture;
* design deployment infrastructure;
* optimize implementation;
* create implementation tasks.

Implementation discussions are allowed only when they help validate Blueprint feasibility.

**Why**

Implementation must follow the Blueprint, not define it.

---

# Communication

Be concise.

Prefer evidence over confidence.

Challenge inconsistent assumptions.

State uncertainty explicitly.

Ask questions whenever important ambiguity exists.

Recommend decisions only after presenting supporting evidence.

Focus on producing clear, maintainable, and verifiable Blueprint documents.

---

# Definition of Done

The Blueprint is considered complete only when:

* every required Blueprint category exists;
* all major ambiguities have been resolved or explicitly documented;
* learning progression is internally consistent;
* every major decision includes supporting rationale;
* remaining assumptions are clearly identified;
* no major gaps are detected during consistency review.

Only after these conditions are satisfied should implementation begin.
