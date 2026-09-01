# IELTS Learning System

This repository is self-describing. `README.md` is orientation only; it does not create a second specification or override an owning authority.

## Read order

1. `CONSTITUTION.md` — repository governance, authority, ownership, topology, naming, and change rules.
2. `OBJECTIVE.md` — project intent, scope, quality target, and success condition.
3. Read the relevant canonical project-documentation root under `docs/`: `PRODUCT.md`, `BEHAVIOR.md`, `ARCHITECTURE.md`, `DATA.md`, `INTERFACES.md`, `QUALITY.md`, `DELIVERY.md`, or `DECISIONS.md`.
4. Follow any legal shard referenced by that root and `docs/catalog/project.json`; the shard refines only its owning docs domain.
5. Use `docs/catalog/project.json` for the canonical machine-readable project inventory, typed identities, relations, milestone state, and support references.
6. Use `contracts/` for scoped exact machine-interface authority where an exact contract has been materialized.
7. Use `spec/` and `design/` only as superseded non-canonical legacy reference retained for provenance; they no longer provide equal project-documentation authority.
8. Use `research/` and `evidence/` for supporting provenance/validation only.

## Authority

The active precedence path is unambiguous:

```text
USER
  ↓
CONSTITUTION.md                 governance authority
  ↓
OBJECTIVE.md                    project intent/scope authority
  ↓
docs/                           canonical project-documentation authority
  ├─ eight domain roots
  ├─ legal referenced shards
  └─ docs/catalog/project.json

contracts/                      scoped exact machine-contract authority
spec/ and design/               superseded non-canonical legacy reference
research/ and evidence/         supporting material only
```

`contracts/` owns exact wire/schema details only for its declared boundary. It does not redefine product, learning, evidence, lifecycle, provider, or other semantic authority owned by `CONSTITUTION.md`, `OBJECTIVE.md`, or `docs/`. Conversely, prose must not silently redefine an exact materialized machine contract inside that contract's scoped wire/schema authority.

## Canonical project-documentation domains

```text
docs/
├── PRODUCT.md
├── BEHAVIOR.md
├── ARCHITECTURE.md
├── DATA.md
├── INTERFACES.md
├── QUALITY.md
├── DELIVERY.md
├── DECISIONS.md
└── catalog/project.json
```

Legal shards under `docs/product/`, `docs/behavior/`, `docs/architecture/`, `docs/data/`, `docs/interfaces/`, and other referenced docs subtrees are canonical only through their owning domain and catalog support references. They do not create a ninth authority domain.

## Legacy specification and design

`spec/**` and `design/**` are retained unchanged except for authority demotion. Their historical semantics remain useful for provenance, but active canonical truth is now represented by the successor `docs/**` authority plane. If legacy wording differs from its successor owner, the successor authority wins.

Do not delete, rename, archive, or treat the legacy trees as current equal-canonical owners merely because their detailed historical prose remains present.

## Implementation navigation

Before implementing a bounded slice:

1. resolve `CONSTITUTION.md` and `OBJECTIVE.md`;
2. resolve the relevant `docs/**` domain owner, referenced shard, and catalog identities/relations;
3. resolve the applicable exact `contracts/**` boundary before implementing a shared machine interface;
4. preserve canonical IDs and applicability semantics across generated bindings and implementation;
5. use coverage/support state from the canonical docs/catalog plane rather than inferring readiness from code, file count, or this README.

Generated bindings, database schemas, UI types, prompts, provider outputs, caches, metrics, migrations, and runtime configuration are derived implementation artifacts unless an owning authority explicitly says otherwise.

## Supporting material

```text
research/   non-canonical research/provenance
evidence/   source/evidence records supporting canonical claims
```

Historical reviews, migration reports, superseded structures, and retired workflow artifacts remain provenance, not active project-documentation authority.

## Current project state

Do not infer coverage, implementation readiness, DESIGN LOCK, or release status from this README. Resolve those states through their owning canonical docs/catalog/task authority. If a summary conflicts with an owner, the owner wins.
