# Constitution

This document is the highest governance contract for the repository. It defines how canonical knowledge is owned, changed, and interpreted. It is not an IELTS-learning specification and must not contain domain truth that belongs in `spec/`.

## 1. Authority

```text
USER
  ↓
CONSTITUTION.md
  ↓
OBJECTIVE.md
  ↓
CANONICAL SPEC OWNERS
```

The user is the final authority for constitutional changes. `README.md`, supporting decisions, research, evidence, archived documents, agent instructions, and implementation artifacts cannot override this hierarchy.

## 2. Repository self-description

The repository must be understandable from the repository itself. A new reasoning session must not require a hidden prompt, vendor-specific handbook, or prior chat history to identify:

- the project's purpose;
- the authority hierarchy;
- each canonical domain owner;
- the dependencies between domains;
- which material is supporting or historical.

Vendor-specific operating documents are not part of the active knowledge architecture.

## 3. Single source of truth by semantic ownership

Every canonical semantic has exactly one owning specification.

A canonical spec may reference another canonical spec, but must not restate semantics owned by it. References are dependencies; copied rules are duplicate authority.

If a rule appears to belong to two domains, ownership must be resolved explicitly. Do not create a third reconciliation document.

## 4. Canonical spec metadata

Every canonical domain spec begins with exactly these four semantic-discovery fields:

```text
STATUS: CANONICAL
OWNS: ...
DEPENDS_ON: ...
DOES_NOT_OWN: ...
```

Do not add document bureaucracy such as author, editor, version, created-at, or last-reviewed metadata. Git is the history mechanism.

`spec/DECISIONS.md` is supporting rationale, not a canonical domain owner, and therefore uses supporting metadata instead.

## 5. Active authority surface

The active Markdown architecture is intentionally small:

- `README.md` — navigation only;
- `CONSTITUTION.md` — governance authority;
- `OBJECTIVE.md` — project-intent authority;
- twelve numbered domain specs under `spec/` — canonical domain truth;
- `spec/DECISIONS.md` — non-authoritative rationale.

The structural baseline is 16 active Markdown files in total.

A new canonical specification file requires a genuinely new authority domain, not merely a large section or a desire to split a document.

Conversely, if two bodies of semantics can evolve independently and have genuinely distinct ownership, the Constitution may be intentionally amended to introduce a new domain. The number 16 is a governance baseline, not a substitute for reasoning.

## 6. Collapse documents, not semantics

Schemas, taxonomies, bindings, coverage rules, consistency rules, and review criteria remain part of the model when they are semantically necessary. They belong as sections inside the specification that owns them unless they become a genuinely independent authority domain.

Document consolidation must never mean semantic deletion.

## 7. No mini-SSOTs

The following do not become canonical merely because they are detailed, frequently referenced, or located near canonical specs:

- research notes;
- evidence summaries;
- source captures;
- decision rationale;
- coverage reports;
- consistency reviews;
- historical schemas;
- archived documents.

Supporting repetition does not create authority.

## 8. Supporting material

`research/` and `evidence/` are non-canonical provenance material. Their purpose is to support or challenge decisions that are ultimately reflected in an owning spec.

The intended flow is:

```text
research / evidence
        ↓
     decision
        ↓
canonical owning spec
```

Supporting material must never become a parallel specification hierarchy.

## 9. Archive semantics

`archive/` is historical only. It preserves superseded structures, exact prior wording, old reviews, and migration provenance.

Archived material has no active authority. When archive and active specs differ, active specs win. Historical material must not be silently copied back into active specs without re-establishing current ownership and validity.

## 10. Evidence discipline

Learning decisions should distinguish:

- official fact or authoritative source evidence;
- evidence-based interpretation;
- Blueprint inference;
- assumption;
- recommendation.

Assumptions and inferences must not be presented as official IELTS truth.

Where authoritative IELTS rules and internal pedagogical inference differ, the external IELTS reality is recorded in `02-IELTS-MODEL.md`; pedagogical interpretation belongs in the relevant learning owner.

## 11. Research-led change

Material learning-policy changes should be evidence-led. When credible evidence conflicts with an active principle or specification, surface the conflict and resolve it at the correct authority level rather than silently selecting whichever rule is convenient.

The Constitution itself changes only by explicit user decision.

## 12. Domain separation

The canonical domains must preserve these conceptual boundaries:

```text
Product       why and for whom
Learner       learner representation
IELTS         external exam reality
Skills        what must be demonstrated
Knowledge     what must be known
Bands         proficiency thresholds and exit criteria
Curriculum    when canonical learning is sequenced
Practice      how capability is trained
Assessment    how mastery is measured
Progression   how learner state advances or regresses
Content Model how concrete content instances reference canonical objects
Glossary      terminology
```

A domain may depend on another without absorbing its semantics.

## 13. Stable identity and references

Canonical learning objects use stable identifiers. Other domains reference those identifiers rather than creating parallel representations.

Structural refactors should preserve valid identifiers whenever possible so that traceability survives document reorganization.

## 14. Canonical vs runtime separation

Blueprint definitions and per-learner runtime state are different things.

A Skill, Knowledge Object, Curriculum Node, Practice Type, or Assessment Type is a canonical definition. A learner's attempts, mastery state, evidence, feedback, scheduling state, and certification history are runtime instances that reference canonical definitions.

Do not store learner-instance truth inside canonical learning objects.

## 15. Learning integrity

The specification must optimize for learning correctness rather than document volume or implementation convenience:

- teach what is required for the target outcome;
- do not over-teach merely because content is available;
- do not under-teach to reduce implementation scope;
- preserve prerequisite integrity where a dependency is genuinely required;
- preserve integration and transfer rather than fragmenting indefinitely.

## 16. Conflict repair

When canonical documents conflict:

1. identify the semantic in conflict;
2. determine its owning domain;
3. repair the owner;
4. replace duplicate statements elsewhere with references;
5. update dependent specs if their behavior changes.

Never solve a contradiction by adding `consistency-review.md`, `final-review.md`, or another authority-like document.

## 17. Change rule

A canonical change should answer:

- What semantic changes?
- Which spec owns it?
- Which dependent specs are affected?
- What evidence or user decision justifies it?
- Does it create duplicate ownership?
- Does it require a new domain, or only a new section?

## 18. Final test

A healthy repository allows a new session to answer, for any important rule:

- Where is the canonical owner?
- What does that owner explicitly not own?
- Which upstream specs does it depend on?
- Which downstream specs consume it?
- Is the same semantic being defined anywhere else?

If those questions cannot be answered quickly, the document architecture has regressed.
