# Constitution

This document is the highest governance contract for the repository. It defines how canonical learning truth, product/runtime design, machine contracts, naming, and implementation ownership are organized.

It is not an IELTS-learning specification and must not contain domain truth that belongs in `spec/`. It is not an API/product-design owner and must not absorb semantics that belong in `design/`.

## 1. Authority

```text
USER
  ↓
CONSTITUTION.md
  ↓
OBJECTIVE.md
  ↓
spec/       canonical learning truth
  ↓
design/     canonical product/runtime translation
  ↓
contracts/  exact machine interface truth once materialized
```

The user is the final authority for constitutional changes.

`README.md`, supporting decisions, research, evidence, archive, generated code, vendor instructions, implementation artifacts, and tooling cannot override this hierarchy.

A downstream owner may operationalize an upstream semantic but may not redefine it.

## 2. Repository self-description

The repository must be understandable from the repository itself. A new reasoning session must not require a hidden prompt, vendor-specific handbook, or prior chat history to identify:

- the project's purpose;
- the authority hierarchy;
- each canonical owner;
- the dependencies between domains;
- which material is supporting or historical;
- how a learner uses the product;
- how the product maps to runtime units;
- the implementation topology;
- naming and cross-language boundary rules.

Vendor-specific operating documents are not part of the active knowledge architecture.

## 3. Single source of truth by semantic ownership

Every canonical semantic has exactly one owner.

A canonical document may reference another owner but must not restate its owned semantics as a second rule.

If a rule appears to belong to two domains, ownership must be resolved explicitly. Do not create a third reconciliation document.

Implementation code consumes canonical semantics; implementation does not become a competing owner merely because the rule is encoded there.

## 4. Canonical Markdown metadata

Every canonical Markdown owner under `spec/` or `design/` begins with exactly these four semantic-discovery fields:

```text
STATUS: CANONICAL
OWNS: ...
DEPENDS_ON: ...
DOES_NOT_OWN: ...
```

Metadata semantics are exact:

- `OWNS` names the semantic authority this document is allowed to define;
- `DOES_NOT_OWN` names nearby semantics that must resolve elsewhere rather than being redefined here;
- `DEPENDS_ON` lists the **direct canonical semantic-definition prerequisites** required to define or interpret this owner's own rules correctly.

`DEPENDS_ON` is **not** a runtime call graph, source-code import graph, transitive dependency closure, list of every file mentioned, or list of downstream consumers. A document may reference a downstream consumer, implementation artifact, evidence record, or supporting material without creating a reverse semantic dependency.

Conversely, if this owner cannot define one of its canonical rules correctly without consuming another canonical owner's semantics, that direct owner belongs in `DEPENDS_ON` even when no runtime call exists.

The canonical semantic-definition dependency graph must remain acyclic. Tooling must not infer service startup order, package imports, network direction, or data-flow direction from `DEPENDS_ON`.

Do not add document bureaucracy such as author, editor, version, created-at, or last-reviewed metadata. Git is the history mechanism.

`spec/DECISIONS.md` is supporting rationale and therefore uses supporting metadata instead.

## 5. Active authority surface

The active Markdown architecture is intentionally bounded:

```text
README.md               navigation only
CONSTITUTION.md         governance authority
OBJECTIVE.md            project-intent authority
spec/                   12 canonical learning owners + 1 supporting decision ledger
design/                 9 canonical product/runtime owners
```

The structural baseline is **25 active Markdown files**.

A new canonical file requires a genuinely new authority domain, not merely a large section or a desire to split a document.

Conversely, independently evolving semantics with genuinely distinct ownership may justify a new domain after an explicit constitutional/user decision.

The number 25 is a governance baseline, not a substitute for reasoning.

Implementation source files, machine contracts, generated code, research, evidence, and archive do not count toward the active Markdown count.

## 6. Learning plane vs design plane

`spec/` owns learning truth:

```text
what the learner is
what IELTS requires
what must be known
what must be demonstrated
Band thresholds
curriculum prerequisites
learning mechanisms and Practice Types
Assessment/evidence/mastery semantics
Progression semantics
content reference semantics
canonical vocabulary
```

`design/` owns product/runtime translation:

```text
how the learner experiences the app
which skill features exist
which practice modes the app exposes
how media is used
how end-to-end product/system flows work
how API resources/operations are organized
which runtime unit/language/framework owns implementation responsibility
which external-service capabilities/providers are allowed and how they fail/exit
how product coverage/support claims are evaluated and declared
```

A design document may use learning truth but cannot change it.

Examples:

- `design/02-practice-catalog.md` may expose 28 practice modes but cannot redefine `PT-*` semantics;
- `design/01-skill-features.md` may define a Writing workspace but cannot redefine Writing Band thresholds;
- `design/05-api.md` may carry `GapEvaluation` across HTTP but cannot redefine gap semantics;
- `design/06-implementation-stack.md` may assign Go to Progression execution but cannot change Progression policy;
- `design/07-third-party-services.md` may select a provider boundary but cannot make provider output learner truth;
- `design/08-coverage-and-support.md` may declare a target uncovered/supported but cannot redefine the IELTS construct or learner readiness standard.

## 7. Collapse documents, not semantics

Schemas, taxonomies, bindings, coverage rules, consistency rules, and review criteria remain part of the model when semantically necessary.

They belong inside the owner unless they become a genuinely independent authority domain.

Document consolidation must never mean semantic deletion.

## 8. No mini-SSOTs

The following do not become canonical merely because they are detailed, frequently referenced, or close to canonical files:

- research notes;
- evidence summaries;
- source captures;
- decision rationale;
- coverage/consistency reviews;
- archived documents;
- generated code;
- database schemas derived from canonical semantics;
- generated API bindings;
- provider catalogs;
- model prompts;
- UI component stories.

Supporting repetition does not create authority.

## 9. Supporting material

`research/` and `evidence/` are non-canonical provenance material.

Intended flow:

```text
research / evidence
        ↓
     decision
        ↓
canonical owner
```

Supporting material must not become a parallel specification hierarchy.

## 10. Archive semantics

`archive/` is historical only. It preserves superseded structures, exact prior wording, reviews, and migration provenance.

When archive and active owners differ, active owners win.

Historical material must not be silently copied back without re-establishing current ownership and validity.

## 11. Evidence discipline

Learning claims should distinguish:

- official fact / authoritative source evidence;
- evidence-based interpretation;
- Blueprint inference;
- assumption;
- recommendation.

Assumptions and inferences must not be presented as official IELTS truth.

Source authority is **claim-domain specific**. The appropriate authority for an IELTS test-format claim, a learning-science claim, an external platform/legal claim, and a framework/runtime claim may be different. A source is useful only when it actually supports the scoped claim.

Material external claims require provenance sufficient to resolve the source and inspect the support for the claim. An LLM output, summary, or LLM-produced citation is not evidence by itself; a citation is only a pointer until the referenced source is resolved, inspected, and shown to support the claim.

Conflicting or insufficient support remains unresolved until the owning domain has a justified resolution. Do not select the more convenient source, convert uncertainty into canonical truth, or silently promote an unresolved claim into an implementation requirement.

External IELTS reality belongs in `spec/02-IELTS-MODEL.md`; pedagogical interpretation belongs in the appropriate learning owner.

External platform/legal policies such as YouTube remain external authority. `design/` records the product interpretation and re-check requirement; it does not pretend to replace those policies.

## 12. Research-led change

Material learning-policy changes should be evidence-led. When credible evidence conflicts with an active rule, surface the conflict and resolve it at the correct owner.

Product defaults such as session length, practice-mode count, feature count, or framework patch versions may change from product/engineering evidence without changing learning truth.

The Constitution itself changes only by explicit user decision.

## 13. Canonical learning domains

```text
Product       why and learning-product boundaries
Learner       learner representation and epistemic state requirements
IELTS         external exam reality
Skills        what must be demonstrated
Knowledge     what must be known
Bands         proficiency thresholds and exit criteria
Curriculum    when canonical learning is sequenced
Practice      how capability is trained
Assessment    observation/evidence/claim semantics
Progression   learner state, gaps, actions, advancement/regression
Content Model concrete/supporting object reference semantics
Glossary      terminology
```

A domain may depend on another without absorbing its semantics.

## 14. Canonical product/runtime design domains

```text
Learning Experience    end-to-end learner journey, TargetProfile and route behavior
Skill Features         user-facing capabilities per skill/shared surface
Practice Catalog       concrete product practice modes
Media / YouTube        media eligibility and learning use
Application Flows      end-to-end web/API/evaluator flows
API                    public/internal resource and operation semantics
Implementation Stack   deployable/language/framework ownership
Third-Party Services   external capability/provider/portability/failure rules
Coverage and Support   product-coverage gaps, support gates and target declarations
```

`design/` may be implementation-specific where the user has intentionally frozen a first-order architecture, but must remain separable from IELTS learning truth.

## 15. Stable identity and references

Canonical learning objects use stable identifiers. Other domains reference those identifiers rather than creating parallel representations.

Structural refactors should preserve valid identifiers whenever possible.

The same concept must use the same canonical vocabulary across Go, Python, TypeScript, contracts, storage, and UI unless a boundary deliberately translates a user-facing label.

Do not create synonym drift such as `student` in one subsystem and `learner` in another when both represent the canonical Learner concept.

## 16. Canonical vs runtime separation

A Skill, Knowledge Object, Curriculum Node, Practice Type, Assessment Type, feature definition, practice-mode definition, provider capability definition, or coverage declaration is reusable canonical design.

A learner's attempts, observations, evidence, mastery estimates, feedback, scheduling state, sessions, drafts, and certification history are runtime instances referencing canonical definitions.

Do not store learner-instance truth inside canonical definitions.

## 17. Learning integrity

The system optimizes for learning correctness rather than document volume or implementation convenience:

- teach what is required for the target outcome;
- do not over-teach merely because content is available;
- do not under-teach to reduce implementation scope;
- preserve genuine prerequisites;
- preserve integration and transfer;
- do not mistake unknown for weak;
- do not mistake stale evidence for regression;
- do not mistake a practice success for automatic mastery;
- do not let product UX lower the evidence standard;
- do not describe a product CoverageGap as a learner weakness;
- do not describe a modelled path as product-supported before support gates pass.

AI and automation may execute, propose, generate, translate, or measure only within the authority of the owning canonical rules. Producing an output does not give a model, prompt, tool, generated artifact, or provider authority over learning truth, product-support truth, or learner state.

When an owning policy requires evidence, calibration, or explicit resolution, AI/automation must preserve the unresolved state rather than inventing a threshold, silently resolving a conflict, or promoting a candidate output into authority.

## 18. Conflict repair

When canonical documents conflict:

1. identify the semantic;
2. determine its owner;
3. repair the owner;
4. replace duplicate statements elsewhere with references;
5. update downstream consumers if behavior changes.

Never solve a contradiction by adding `consistency-review.md`, `final-review.md`, or another authority-like reconciliation file.

## 19. Change rule

A learning change should answer:

- What semantic changes?
- Which `spec/` owner owns it?
- Which dependencies are affected?
- What evidence/user decision justifies it?
- Does it duplicate another owner?

A product/runtime design change should answer:

- Which `design/` owner owns it?
- Does it alter learning truth or only translate it?
- Which user/system flows are affected?
- Which runtime unit owns execution?
- Does it create/change a cross-language or third-party boundary?
- Which contract/provider/coverage declaration must change?

An implementation-stack change should also answer:

- Does it add a deployable?
- Does it add a primary language?
- Does it duplicate domain logic?
- Does it add unjustified infrastructure?

## 20. Primary implementation languages

Exactly three primary application languages are approved:

```text
Go
Python
TypeScript
```

Supporting formats/tooling such as SQL, shell, YAML, JSON, OpenAPI, JSON Schema, CSS, HTML, or generated artifacts are not additional primary application languages.

Adding another primary runtime/application language requires explicit user approval.

First-order allocation is owned by `design/06-implementation-stack.md`.

## 21. Repository topology follows responsibility, not language

Implementation topology is responsibility-first:

```text
apps/       user-facing deployables and clients
services/   independently runnable backend services/workers
packages/   reusable implementation libraries when justified
contracts/  language-neutral cross-unit interface definitions
tools/      repository/development/generation/migration/release tooling
```

Forbidden top-level product organization:

```text
go/
python/
py/
typescript/
ts/
```

A unit is named for what it does, not what language implements it.

Initial approved deployable topology is defined in `design/06-implementation-stack.md`:

```text
apps/web
services/core-api
services/evaluator
```

Do not create empty taxonomy or microservice-per-feature structure merely to mirror product feature names.

Avoid cross-domain junk drawers such as:

```text
common/
shared/
utils/
helpers/
misc/
```

A reusable package must be named for the capability it owns.

## 22. Cross-language boundary rule

Go, Python, and TypeScript may not share application semantics by copy-pasting equivalent models or rules.

Cross-language interaction must cross an explicit boundary:

```text
canonical learning/design owner
      ↓
implementation owner
      ↓
machine contract
      ↓
derived language binding/client
```

`contracts/` owns exact machine-readable interface truth for genuine boundaries.

Rules:

1. contract defines transport/interface shape, not learning truth;
2. domain meaning remains owned by `spec/`/`design/`;
3. one interface has one contract authority;
4. handwritten mirror schemas across languages are forbidden when shared contract generation/validation is viable;
5. generated bindings are derived and not manually edited as authority;
6. wire-field naming is defined once;
7. stable canonical IDs survive boundaries unchanged;
8. browser never bypasses the Go public API to make Python a second product API under the initial design.

## 23. Global naming rules

Repository-level architecture names use **lowercase kebab-case**:

```text
core-api
scoring-worker
media-analyzer
```

This applies to deployable-unit directories and language-neutral architecture folders below `apps/`, `services/`, `contracts/`, and `tools/`.

Canonical terminology from `spec/11-GLOSSARY.md` controls domain vocabulary. Case may change by language; concept names must not drift.

Prefer concrete nouns/capabilities over filler such as:

```text
ThingManager
CommonService
SharedUtils
DataHelper
CoreStuff
```

Abbreviations are allowed only when established project/domain vocabulary or universal technical initialisms.

## 24. Language-specific naming and source layout

### Go

```text
package directories   compact lowercase words; no hyphens/underscores
source files           lower_snake_case.go
test files             *_test.go
exported identifiers   PascalCase
unexported identifiers camelCase
initialisms            ID, URL, HTTP, API, JSON where idiomatic
```

Runnable Go unit shape:

```text
services/<unit>/
├── go.mod
├── cmd/<binary>/
└── internal/
```

Do not create `pkg/` automatically.

### Python

```text
package/module names  snake_case
functions/variables   snake_case
classes/types         PascalCase
constants             UPPER_SNAKE_CASE
test files            test_*.py
```

Application/service shape:

```text
services/<unit>/
├── pyproject.toml
├── src/<python_package>/
└── tests/
```

Outer unit remains kebab-case; import package uses snake_case.

### TypeScript

```text
module/file names     kebab-case.ts / kebab-case.tsx
functions/variables   camelCase
classes/types/interfaces/components PascalCase
true constants        UPPER_SNAKE_CASE
test files            *.test.ts / *.test.tsx
```

Do not mix `.spec.*` and `.test.*` without an explicit tooling reason.

Framework-required structure may be used but must not create a competing domain taxonomy.

## 25. Package and module ownership

A package/module exists to own a coherent implementation responsibility, not merely to reduce file size.

Rules:

1. one package/module should have one reason to change;
2. cross-domain dependencies remain visible;
3. domain rules execute at one runtime owner;
4. another language consumes behavior/contract instead of silently reimplementing it;
5. presentation transformations may exist in clients but never become alternative scoring/progression policy;
6. cyclic deployable/package dependencies are forbidden;
7. generated code is never the conceptual owner of its source interface.

## 26. Infrastructure follows demonstrated need

Do not introduce a broker, queue technology, vector database, workflow engine, microservice split, or second public API merely because it is common architecture.

The initial design may describe durable async behavior; implementation should use the smallest mechanism that satisfies the actual reliability/throughput contract.

No implementation, repository automation, or deployment may introduce a mandatory paid external dependency, usage-billed repository capability, metered hosted execution/storage capability, or paid provider without explicit authorization from the repository's USER authority. Free/open-source local alternatives remain preferred when semantically sufficient; optional paid services may be evaluated but not silently activated, and normal dependency, license, security, privacy, quality, and provider rules still apply.

A new infrastructure subsystem should answer:

- what measured/contractual requirement requires it;
- why the current mechanism cannot satisfy that requirement;
- operational/security cost;
- rollback/exit path;
- acceptance evidence.

## 27. Verification ownership

Once implementation exists, the repository exposes one canonical root verification entrypoint:

```text
root verify
  ├── TypeScript web checks
  ├── Go Core API checks
  ├── Python evaluator checks
  ├── contract validation/generation drift
  └── cross-unit integration/E2E checks
```

Each unit owns native formatter/linter/typechecker/tests/build.

A root PASS is authoritative only when it covers every affected language and relevant contract boundary.

## 28. Final architecture test

A healthy repository lets a new session answer quickly:

- Which learning owner defines this semantic?
- Which design owner translates it into product/runtime behavior?
- Which runtime unit executes it?
- Which primary language/framework implements that unit?
- Which machine contract crosses the boundary?
- Which third-party provider boundary is involved, if any?
- Is this target MODELLED, COVERED, SUPPORTED_FOR_PRODUCT, or VALIDATED?
- Is the semantic duplicated elsewhere?
- Can material external claims be traced to inspectable provenance?
- Can direct canonical semantic dependencies be identified without treating `DEPENDS_ON` as a runtime/import graph?
- Can root verification prove the affected path?

If the answer requires guessing from folders, reading hidden agent instructions, trusting AI/tool output as authority, reconciling three language-specific copies of one rule, or trusting an undeclared provider/coverage assumption, the architecture has regressed.