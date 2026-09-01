# Constitution

This document is the highest governance contract for the repository. It defines how canonical project semantics, machine contracts, naming, and implementation ownership are organized.

It is not a project-domain specification and must not absorb semantics owned by the canonical `docs/**` authority plane. Exact materialized wire/schema truth belongs to `contracts/**` only within its declared machine-interface boundary.

## 1. Authority

```text
USER
  ↓
CONSTITUTION.md
  ↓
OBJECTIVE.md
  ↓
docs/       canonical project-documentation authority
  ├─ eight authority-domain roots
  ├─ legal referenced shards
  └─ docs/catalog/project.json
  ↓ scoped materialization at genuine machine boundaries
contracts/  exact machine-interface truth within the declared boundary
```

The user is the final authority for constitutional changes.

`README.md`, supporting decisions, research, evidence, historical material, generated code, vendor instructions, implementation artifacts, and tooling cannot override this hierarchy.

`spec/**` and `design/**` are superseded non-canonical legacy material. They may remain as provenance until separately retired, but they do not own current project semantics and cannot override the canonical `docs/**` authority plane.

A downstream owner may operationalize an upstream semantic but may not redefine it. A scoped machine contract may define exact wire/schema representation without becoming the owner of the project semantic it carries; conversely, prose must not silently redefine an exact materialized machine contract inside that contract's scoped boundary.

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

## 4. Canonical documentation ownership

The canonical project-documentation plane is `docs/**`. Its top-level authority roots are exactly the eight domains named in section 13. A legal shard under `docs/**` is canonical only through its owning domain and the owning root/catalog references; a shard does not create another authority domain merely because it contains detailed semantics.

`docs/catalog/project.json` is the canonical machine-readable documentation inventory for the adopted documentation model, including typed identities, relations, milestone state, and support references. It does not become a ninth semantic authority domain.

Canonical ownership remains semantic rather than metadata-driven. Documents may use ownership/status metadata where useful, but no legacy `STATUS`/`OWNS`/`DEPENDS_ON`/`DOES_NOT_OWN` header convention under `spec/**` or `design/**` can confer current authority after the docs cutover.

Direct semantic dependencies mean the prerequisites required to define or interpret an owner's rules correctly. They are not a runtime call graph, source-code import graph, transitive dependency closure, list of every file mentioned, or list of downstream consumers. A document may reference a downstream consumer, implementation artifact, evidence record, or supporting material without creating a reverse semantic dependency.

The canonical semantic-definition dependency graph must remain acyclic. Tooling must not infer service startup order, package imports, network direction, or data-flow direction from semantic-document dependencies.

Do not add document bureaucracy such as author, editor, version, created-at, or last-reviewed metadata merely to signal authority. Git is the history mechanism.

## 5. Active authority surface

The active authority architecture is intentionally bounded:

```text
README.md               navigation only
CONSTITUTION.md         governance authority
OBJECTIVE.md            project-intent authority
docs/                   canonical project-documentation authority
  eight domain roots    PRODUCT, BEHAVIOR, ARCHITECTURE, DATA,
                        INTERFACES, QUALITY, DELIVERY, DECISIONS
  legal shards          same-domain refinements referenced by owner/catalog
  catalog/project.json  canonical documentation inventory/relations/state
contracts/              scoped exact machine-interface authority
```

`spec/**` and `design/**` are superseded non-canonical legacy reference. `research/**` and `evidence/**` are non-canonical provenance material. Superseded history belongs in Git history by default unless a current provenance, legal, or recovery need justifies a tracked supporting artifact.

A new canonical shard requires a real ownership or navigation need inside an existing domain; it does not create a new authority domain. A ninth authority domain requires an explicit constitutional/user decision establishing genuinely distinct ownership.

Implementation source files, generated code, research, evidence, legacy material, and historical material do not become authority merely because they are detailed or frequently referenced.

## 6. Canonical semantic plane

The eight `docs/**` domains jointly own the current project semantics:

```text
PRODUCT       product boundaries, learner/product ontology, IELTS/Band truth,
              curriculum/coverage and product-level capability semantics
BEHAVIOR      functional behavior, practice, planning, flows and lifecycle semantics
ARCHITECTURE  runtime responsibility, topology, technology-family and build/buy boundaries
DATA          data, evidence, progression and content semantics
INTERFACES    interface, async and external-dependency/provider-boundary semantics
QUALITY       security, privacy, reliability, performance, observability, testing and cost quality
DELIVERY      environment, deployment, migration, rollback, backup, compatibility and operations
DECISIONS     material decisions and unresolved-question ownership
```

Detailed semantic groups such as Skills, Knowledge Objects, Bands, Curriculum Nodes, Practice Types, Assessment Types, Progression, content contexts, learner experience, provider lifecycle, or internal trust remain owned by the smallest applicable existing root or legal shard inside these eight domains. They do not create parallel top-level authority domains.

Legacy `spec/**` and `design/**` material may be consulted as provenance but cannot define, amend, or override a current semantic. If legacy wording differs from the canonical successor owner, the canonical `docs/**` owner wins.

## 7. Collapse documents, not semantics

Schemas, taxonomies, bindings, coverage rules, consistency rules, and review criteria remain part of the model when semantically necessary.

They belong inside the owner unless they become a genuinely independent authority domain.

Document consolidation must never mean semantic deletion.

## 8. No mini-SSOTs

The following do not become canonical merely because they are detailed, frequently referenced, or close to canonical files:

- research notes;
- evidence summaries;
- source captures;
- decision rationale outside its canonical decision owner;
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

## 10. Historical material

Superseded structures, exact prior wording, reviews, migrations, and retired workflow/process artifacts live in Git history unless a current provenance, legal, or recovery need justifies a tracked supporting artifact.

Historical material cannot override active owners and must not be silently copied back without re-establishing current ownership and validity.

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

External IELTS reality and Band semantics belong to the canonical PRODUCT owner `docs/product/ielts-exam-and-bands.md`; pedagogical interpretation belongs in the smallest applicable canonical `docs/**` owner.

External platform/legal policies such as YouTube remain external authority. The applicable canonical `docs/**` owner records the product interpretation and re-check requirement; repository documentation does not pretend to replace those external policies.

## 12. Research-led change

Material learning-policy changes should be evidence-led. When credible evidence conflicts with an active rule, surface the conflict and resolve it at the correct owner.

Product defaults such as session length, practice-mode count, feature count, or framework patch versions may change from product/engineering evidence without changing learning truth.

The Constitution itself changes only by explicit user decision.

## 13. Canonical documentation domains

The top-level canonical project-documentation domains are exactly:

```text
PRODUCT
BEHAVIOR
ARCHITECTURE
DATA
INTERFACES
QUALITY
DELIVERY
DECISIONS
```

A semantic belongs to the smallest correct owner inside one of these domains. A domain may depend on another without absorbing its semantics. Detailed learning/product/runtime taxonomies remain semantic groups inside the applicable owner; they are not additional top-level authority domains.

## 14. Legal shards and scoped machine truth

Legal `docs/**` shards refine only their owning domain and are canonical through the owning root and `docs/catalog/project.json` support references. They may evolve independently where ownership warrants it without creating a ninth domain.

`docs/catalog/project.json` owns the canonical machine-readable documentation inventory, typed identities, relations, milestone state, and support references for the adopted model.

`contracts/**` owns exact machine-readable interface truth only for genuine materialized boundaries. A contract defines exact transport/schema shape, not product, learning, evidence, progression, provider, coverage, or other project-domain meaning unless the canonical semantic owner explicitly defines that meaning there within the scoped contract boundary.

`spec/**` and `design/**` remain superseded non-canonical legacy reference and cannot own current semantics.

## 15. Stable identity and references

Canonical project objects use stable identifiers. Other domains reference those identifiers rather than creating parallel representations.

Structural refactors should preserve valid identifiers whenever possible.

The same concept must use the same canonical vocabulary across Go, Python, TypeScript, contracts, storage, and UI unless a boundary deliberately translates a user-facing label.

Do not create synonym drift such as `student` in one subsystem and `learner` in another when both represent the canonical Learner concept.

## 16. Canonical vs runtime separation

A Skill, Knowledge Object, Curriculum Node, Practice Type, Assessment Type, feature definition, practice-mode definition, provider capability definition, or coverage declaration is reusable canonical definition owned in the applicable `docs/**` domain.

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

A canonical semantic change should answer:

- What semantic changes?
- Which of the eight `docs/**` authority domains owns it, and which smallest existing owner inside that domain defines it?
- Which semantic dependencies are affected?
- What evidence/user decision justifies it?
- Does it duplicate another owner?

A product/runtime implementation change should also answer:

- Which canonical BEHAVIOR, ARCHITECTURE, DATA, INTERFACES, QUALITY, DELIVERY, PRODUCT, or DECISIONS owner governs the affected consequence?
- Does it alter canonical project meaning or only translate/implement it?
- Which user/system flows are affected?
- Which runtime unit owns execution?
- Does it create/change a cross-language or third-party boundary?
- Which scoped contract/provider/coverage declaration must change?

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

First-order runtime responsibility and technology-family allocation are owned by the canonical ARCHITECTURE authority, currently `docs/ARCHITECTURE.md` plus its legal referenced shards/catalog relations where applicable.

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

Within a deployable, application code is organized primarily by product feature/capability ownership. Framework/transport shells and genuine cross-feature infrastructure remain at their owning boundaries; additional layers exist only for real dependency, invariant, or change boundaries. Pure horizontal layer-first organization and empty architecture layers are not repository defaults. Detailed source organization must remain consistent with the canonical ARCHITECTURE authority and existing coherent repository conventions; it cannot create competing semantic ownership.

Initial approved deployable topology is defined by `docs/ARCHITECTURE.md`:

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
canonical docs semantic owner
      ↓
implementation owner
      ↓
machine contract
      ↓
derived language binding/client
```

`contracts/` owns exact machine-readable interface truth for genuine boundaries.

Rules:

1. contract defines transport/interface shape, not project-domain truth;
2. domain meaning remains owned by the applicable canonical `docs/**` owner;
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

Canonical terminology is owned by the canonical semantic owner that defines the concept. There is no independent glossary authority. Legacy `spec/11-GLOSSARY.md` is superseded non-canonical reference and cannot override a current `docs/**` owner. Case may change by language; concept names must not drift.

Prefer concrete nouns/capabilities over filler such as:

```text
ThingManager
CommonService
SharedUtils
DataHelper
CoreStuff
```

Abbreviations are allowed only when established project/domain vocabulary or universal technical initialisms.

## 24. Cross-language semantic naming

Repository naming follows three invariants:

```text
SEMANTIC NAME IS SHARED.
LANGUAGE SYNTAX IS IDIOMATIC.
WIRE IDENTITY IS EXACT.
```

Canonical/domain vocabulary remains stable across Go, Python, TypeScript, storage references, and UI code while each implementation language uses its idiomatic package/module/file/identifier casing. Exact serialized HTTP/JSON/event names remain owned by the machine contract.

Detailed language-specific source layout, naming, test placement, generated-code placement, and shared-code promotion rules must remain consistent with the canonical ARCHITECTURE authority and repository conventions; they cannot create a second semantic owner.

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

- Which canonical `docs/**` owner defines this semantic?
- Which canonical behavior/architecture/interface owner translates or carries it into product/runtime behavior where applicable?
- Which runtime unit executes it?
- Which primary language/framework implements that unit?
- Which machine contract crosses the boundary?
- Which third-party provider boundary is involved, if any?
- Is this target MODELLED, COVERED, SUPPORTED_FOR_PRODUCT, or VALIDATED?
- Is the semantic duplicated elsewhere?
- Can material external claims be traced to inspectable provenance?
- Can direct canonical semantic dependencies be identified without treating them as a runtime/import graph?
- Can root verification prove the affected path?

If the answer requires guessing from folders, reading hidden agent instructions, trusting AI/tool output as authority, consulting superseded legacy owners for current truth, reconciling three language-specific copies of one rule, or trusting an undeclared provider/coverage assumption, the architecture has regressed.
