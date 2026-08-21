# Objective

## Purpose

Build a complete, evidence-based learning specification for an IELTS learning system.

The specification must define, with enough precision for later implementation, what a learner needs to know, what the learner needs to be able to demonstrate, how those capabilities are trained and assessed, and how the learner progresses from Band 3 through Band 9.

## Target outcome

A future implementation team or reasoning session should be able to consume the repository without making new major learning-system decisions.

The repository should support a learner moving from a Band-3 program entry point toward Band 9 across Listening, Reading, Writing, and Speaking while preserving uneven per-skill profiles.

## Scope

The active specification covers:

- IELTS Academic as the fully specified initial test variant;
- a structure that can accept General Training variant differences without restructuring shared learning semantics;
- the four IELTS skills;
- enabling knowledge such as grammar, vocabulary, and phonology;
- band expectations and exit criteria;
- curriculum sequencing and prerequisites;
- practice and feedback strategy;
- assessment and mastery evidence;
- learner progression, regression, review, and certification semantics;
- conceptual content representations required to instantiate the learning model.

## Non-goals

This repository does not currently define:

- production application architecture;
- database technology or schemas;
- APIs;
- deployment or infrastructure;
- authentication or authorization;
- pricing, commercialization, or go-to-market strategy;
- detailed UX or visual design;
- vendor-specific agent instructions.

Those systems may later consume the learning specification. They do not redefine it.

## Quality target

Optimize for:

1. correctness;
2. completeness;
3. internal consistency;
4. evidence and traceability;
5. explicit semantic ownership;
6. implementation-independent clarity.

Do not optimize the learning model for implementation speed at the expense of learning correctness.

## Success definition

The project objective is satisfied when:

- every major learning semantic has one canonical owner;
- Bands 3–9 form a coherent learning progression for all four skills;
- skill requirements resolve to the knowledge they require;
- curriculum paths reference canonical skill and knowledge objects rather than redefining them;
- practice covers acquisition, consolidation, retrieval, transfer, fluency, and exam readiness;
- assessment defines sufficient, confidence-aware evidence for mastery;
- progression supports independent per-skill advancement and regression/re-certification;
- Academic IELTS is fully represented and General Training can be added as a variant overlay without restructuring shared domains;
- a new reasoning session can reconstruct the learning system from `README.md` → `CONSTITUTION.md` → this Objective → relevant canonical specs.
