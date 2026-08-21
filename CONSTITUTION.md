# Constitution

This document is the highest governance contract for the repository. It defines how canonical knowledge is owned, changed, interpreted, and how implementation code is organized. It is not an IELTS-learning specification and must not contain domain truth that belongs in `spec/`.

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

The user is the final authority for constitutional changes. `README.md`, supporting decisions, research, evidence, archived documents, agent instructions, implementation artifacts, generated code, and tooling cannot override this hierarchy.

## 2. Repository self-description

The repository must be understandable from the repository itself. A new reasoning session must not require a hidden prompt, vendor-specific handbook, or prior chat history to identify:

- the project's purpose;
- the authority hierarchy;
- each canonical domain owner;
- the dependencies between domains;
- which material is supporting or historical;
- the implementation topology once source code exists;
- the naming and cross-language boundary rules.

Vendor-specific operating documents are not part of the active knowledge architecture.

## 3. Single source of truth by semantic ownership

Every canonical semantic has exactly one owning specification.

A canonical spec may reference another canonical spec, but must not restate semantics owned by it. References are dependencies; copied rules are duplicate authority.

If a rule appears to belong to two domains, ownership must be resolved explicitly. Do not create a third reconciliation document.

Implementation code consumes canonical semantics; implementation does not become a competing semantic owner merely because the rule is encoded there.

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

Implementation directories and source files do not count toward the active Markdown authority quota.

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
- archived documents;
- generated code;
- database schemas derived from canonical semantics;
- API bindings derived from contracts.

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

The same domain concept must use the same canonical vocabulary across Go, Python, TypeScript, contracts, storage, and UI unless a boundary deliberately translates a user-facing label. Code must not create synonym drift such as `student` in one subsystem and `learner` in another when both represent the canonical Learner concept.

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

An implementation-architecture change should answer:

- Which deployable unit or package owns the responsibility?
- Does the change introduce a new cross-language boundary?
- Which contract owns that boundary?
- Does it duplicate domain logic in another language?
- Does it add a new primary implementation language?

## 18. Final knowledge-architecture test

A healthy repository allows a new session to answer, for any important rule:

- Where is the canonical owner?
- What does that owner explicitly not own?
- Which upstream specs does it depend on?
- Which downstream specs consume it?
- Is the same semantic being defined anywhere else?

If those questions cannot be answered quickly, the document architecture has regressed.

## 19. Primary implementation languages

The product is a polyglot system with exactly three approved **primary application languages**:

```text
Go
Python
TypeScript
```

This rule does not forbid supporting formats or tooling such as SQL, shell, YAML, JSON, Protocol Buffers, OpenAPI, CSS, HTML, or generated artifacts. Those are not independent primary application-language domains.

Adding another primary runtime/application language is a repository-architecture change and requires explicit user approval.

The Constitution does **not** pre-assign every product responsibility to a language. Language follows the responsibility and runtime constraints of the owning unit; semantic ownership still follows `spec/`.

## 20. Repository topology follows responsibility, not language

The implementation topology is responsibility-first:

```text
apps/       user-facing deployables and clients
services/   independently runnable backend services and workers
packages/   reusable implementation libraries
contracts/  language-neutral cross-unit interface definitions
tools/      repository, development, generation, migration, and release tooling
```

These directories are created only when real implementation units exist. Do not add empty taxonomy merely to match this diagram.

Forbidden top-level product organization:

```text
go/
python/
py/
typescript/
ts/
```

The repository must not be partitioned into three language silos. A unit is named for what it does, not what language implements it.

Examples:

```text
apps/web/
services/learning-api/
services/scoring-worker/
packages/...
```

The examples are topology examples, not pre-approved product components.

Every deployable unit has one primary language and one clear responsibility. If one responsibility requires multiple runtimes, separate them into explicit units connected by a contract rather than mixing independent application runtimes in one source package.

Avoid top-level or cross-domain junk drawers such as:

```text
common/
shared/
utils/
helpers/
misc/
```

A reusable package must be named for the capability it owns. Generic helper names are acceptable only inside a narrowly scoped owning package when the responsibility is unambiguous.

## 21. Cross-language boundary rule

Go, Python, and TypeScript may not share application semantics by copy-pasting equivalent models or business rules into each language.

Cross-language interaction must cross an explicit boundary:

```text
canonical spec
      ↓
implementation owner
      ↓
contract when another unit/language must consume it
      ↓
derived language binding / client
```

`contracts/` owns machine-readable **interface truth** for boundaries that genuinely cross units or languages. Contract technology is chosen per need; this Constitution does not force Protocol Buffers, OpenAPI, JSON Schema, or another format before evidence requires it.

Rules:

1. A contract defines transport/interface shape, not learning-domain truth.
2. Domain meaning remains owned by the relevant canonical spec.
3. One interface has one contract authority.
4. Handwritten mirror schemas in Go, Python, and TypeScript are forbidden when a shared contract can generate or validate them.
5. Generated bindings are derived artifacts and must not be manually edited as authority.
6. Wire-field naming is defined once by the owning contract; each language may expose idiomatic local identifiers only through a deliberate mapping.
7. Stable canonical IDs must survive language boundaries unchanged.

## 22. Global naming rules

Repository-level architecture names use **lowercase kebab-case**:

```text
learning-api
scoring-worker
content-generator
```

This applies to deployable-unit directories and language-neutral architecture folders below `apps/`, `services/`, `contracts/`, and `tools/`.

Canonical terminology from `spec/11-GLOSSARY.md` controls domain vocabulary. Case may change to match a language, but the concept name must not drift.

Prefer concrete nouns and capabilities over architectural filler. Do not encode implementation patterns into names unless the pattern is the real responsibility.

Avoid names such as:

```text
ThingManager
CommonService
SharedUtils
DataHelper
CoreStuff
```

unless the name genuinely describes the narrow responsibility better than a domain term.

Abbreviations are allowed only when they are established project/domain vocabulary or universal technical initialisms. Prefer `assessment` over an invented short form that another language must rediscover.

## 23. Language-specific naming and source layout

Repository topology is shared; source naming inside each unit follows the language's native conventions.

### Go

Canonical conventions:

```text
package directories  compact lowercase words; no hyphens or underscores
source files          lower_snake_case.go
test files            *_test.go
exported identifiers  PascalCase
unexported identifiers camelCase
initialisms           ID, URL, HTTP, API, JSON where idiomatic
```

A runnable Go unit should normally use the Go-native shape:

```text
services/<unit>/
├── go.mod
├── cmd/<binary>/
└── internal/
```

Do not create `pkg/` automatically. Export reusable Go packages only when another owning unit actually needs them.

### Python

Canonical conventions follow PEP-style naming:

```text
package/module names   snake_case
functions/variables    snake_case
classes/types          PascalCase
constants              UPPER_SNAKE_CASE
test files             test_*.py
```

A Python application/service should normally use a `src` layout:

```text
services/<unit>/
├── pyproject.toml
├── src/<python_package>/
└── tests/
```

The outer unit directory remains responsibility-based kebab-case; the importable Python package uses snake_case.

### TypeScript

Canonical conventions:

```text
module/directory files  kebab-case.ts / kebab-case.tsx
functions/variables     camelCase
classes/types/interfaces/components PascalCase
true constants          UPPER_SNAKE_CASE
test files              *.test.ts / *.test.tsx
```

Use `.test.*` consistently; do not mix `.spec.*` and `.test.*` conventions in the same repository without an explicit tooling reason.

A TypeScript application/package should normally place owned source under:

```text
<unit>/
├── package.json
└── src/
```

Framework-generated structure may add required directories, but it must not create a second domain taxonomy that competes with canonical vocabulary.

## 24. Package and module ownership

A package/module exists to own a coherent implementation responsibility, not merely to reduce file size.

Rules:

1. One package/module should have one reason to change.
2. Cross-domain dependencies must be visible rather than hidden behind `common` utilities.
3. Domain rules are implemented at one runtime owner and exposed through interfaces when another unit needs them.
4. A second language may consume the behavior or contract; it must not silently reimplement the rule.
5. Pure presentation transformations may exist in clients, but they must not become alternative learning/progression/scoring policy.
6. Cyclic dependencies between deployable units or reusable packages are forbidden.
7. Generated code is never the conceptual owner of the interface that produced it.

## 25. Verification ownership

Once implementation code exists, the repository must expose one canonical root verification entrypoint. The exact tool may be selected when the first implementation slice is introduced, but the semantic contract is fixed:

```text
root verify
  ├── Go checks
  ├── Python checks
  ├── TypeScript checks
  ├── contract validation/generation drift checks
  └── cross-unit integration checks where applicable
```

Each unit also owns its native formatter, linter, type checker, tests, and build checks.

A root PASS is authoritative only when it covers every affected primary language and relevant contract boundary. A green Go test cannot certify a change that also modifies Python or TypeScript behavior.

## 26. Final implementation-architecture test

Once source code exists, a healthy repository must allow a new session to answer quickly:

- Which unit owns this runtime responsibility?
- Which primary language implements it?
- Which spec owns the domain semantics it consumes?
- Which contract governs every cross-unit/cross-language boundary?
- Is any rule duplicated in another language?
- Are names derived from canonical project vocabulary?
- Can the root verification entrypoint prove the affected units together?

If the answer requires guessing from folder names or reading three language-specific copies of the same rule, the implementation architecture has regressed.
