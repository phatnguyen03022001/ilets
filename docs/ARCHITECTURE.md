# Architecture

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> `TASK-0004` revision 1 migration artifact. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Execution base: `phatnguyen03022001/ilets@64d5915a2d31ec5eec025ccaf05b11aef60e9933`.
>
> Model pin: `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07` (unreleased V1 candidate).

This draft migrates only the established logical runtime responsibility layer from canonical `design/04-application-flows.md`, `design/05-api.md`, and `design/06-implementation-stack.md`. `docs/catalog/project.json` owns the `SYS-*` identities and typed `FLW-* → SYS-*` relations; this file explains the migrated architecture only.

## ARCH-COMPONENTS Components and ownership

The canonical design fixes exactly three initial logical runtime responsibility units. They are ownership/trust boundaries, not a requirement for three separately billed deployment stacks.

### SYS-001 Web

`apps/web` is the TypeScript / React / Next.js App Router learner/admin presentation unit. It bounds presentation input, renders product state, and calls the Core API for authoritative reads and mutations. Client or Next.js state remains presentation/transient unless persisted through Core.

Web does not own durable learner/product/content/evidence state, does not read or write the authoritative product store directly, does not duplicate Assessment/Progression/Planner policy, and does not call the Evaluator as an alternate backend.

### SYS-002 Core API

`services/core-api` is the Go `net/http` + `chi` public product API and deterministic orchestration unit. It owns learner/admin-facing product behavior, authoritative work identity/state, product authorization, deterministic policy execution, and the application path to durable product state.

Core is the sole application runtime permitted to read/write the authoritative product persistence boundary. It is also the permitted product caller for bounded Evaluator capabilities and remains responsible for validating/reconciling returned results before owning product policy consumes them.

### SYS-003 Evaluator

`services/evaluator` is the Python / FastAPI bounded AI/audio/text evaluation capability unit. It performs bounded evaluator/media/content-analysis work behind the internal Core-called boundary and returns signals/results with provenance and uncertainty where material.

Evaluator output is non-authoritative until Core validates and interprets it through owning policy. Evaluator does not certify Band, mutate learner progression, activate content, choose the final next action, expose a public product API, or directly read/write authoritative product persistence.

## ARCH-TOPOLOGY Runtime topology

The legal logical path is:

```text
Browser / learner-admin interaction
        ↓
SYS-001 Web
        ↓
SYS-002 Core API
   ├── authoritative product persistence
   ├── deterministic policy/orchestration
   └── bounded internal capability request
                ↓
         SYS-003 Evaluator
                ↓
      bounded result / provenance
                ↓
         SYS-002 Core API
                ↓
 Assessment / Progression / Planner
                ↓
         SYS-001 Web
```

Deployment may co-locate these logical units when caller/callee direction, Core-only persistence access, explicit cross-runtime contracts, least-privilege runtime identity, and attributable failure/restart behavior remain intact. Co-location or later deployment splitting does not move semantic authority.

## ARCH-BOUNDARIES Communication and trust boundaries

- **Browser/user → Web:** browser input is untrusted. Web owns presentation handling only and cannot promote client state into product authority.
- **Web → Core:** all learner/admin authoritative product reads and mutations cross the Core public API boundary. Web cannot bypass Core to persistence or Evaluator.
- **Core → Evaluator:** this is an internal machine-contract boundary. Core owns authoritative work state and product interpretation; Evaluator owns only the bounded capability execution. Reachability or network placement alone is not caller authorization.
- **Evaluator → Core:** returned output is bounded, provenance-bearing input to Core reconciliation, not learner/evidence/content/product authority. Stale or superseded completion cannot independently mutate current product truth.
- **Core → authoritative persistence:** only Core has application-runtime read/write authority. Web and Evaluator have no direct application access.

These are architecture-level relationships only. Exact `IFC-*`, `DAT-*`, external dependency, retry, contract-field, and persistence-lifecycle inventories remain outside this migration slice.

## ARCH-TECHNOLOGY Material technology-family choices

| Logical unit / boundary | Canonical technology-family choice | Consequence preserved by this migration |
| --- | --- | --- |
| `SYS-001` Web | TypeScript, React, Next.js App Router | presentation authority only; authoritative product behavior remains behind Core |
| `SYS-002` Core API | Go, `net/http`, `chi` | public product API, deterministic orchestration, product authorization, and sole application persistence access |
| `SYS-003` Evaluator | Python, FastAPI | bounded internal evaluator capability only; no product-state authority |
| Core persistence boundary | PostgreSQL-compatible authoritative store | application access remains Core-only; storage schema is derived rather than semantic authority |
| Cross-runtime HTTP boundaries | one exact machine contract per material boundary once materialized | parallel runtimes do not independently author equivalent wire truth; INTERFACES migration remains unresolved |

Patch versions, concrete deployment topology, provider inventory, exact persistence schema, and exact machine contracts remain with their existing canonical owners and are not materialized here.

## ARCH-FLOW-PARTICIPATION Existing flow participation

The catalog records only participation already implied by the migrated `docs/BEHAVIOR.md` flows and canonical runtime path:

| Flow | Participating migrated systems | Existing runtime basis |
| --- | --- | --- |
| `FLW-001` Governed target-to-next-action loop | `SYS-001`, `SYS-002`, `SYS-003` | learner-facing Web path through Core, with bounded evaluation where evidence/support requires it |
| `FLW-002` Learner control within eligibility | `SYS-001`, `SYS-002` | learner presentation plus Core-owned eligibility/product mutation; no Evaluator participation is required |
| `FLW-003` Truthful unresolved and failure handling | `SYS-001`, `SYS-002`, `SYS-003` | Web surfaces Core-owned product state; evaluator inability/stale results are explicit bounded failure inputs |
| `FLW-004` Evidence-to-replan reconciliation | `SYS-002`, `SYS-003` | evaluator observation candidates return to Core; Core owns Assessment/Progression/Planner reconciliation |

This table does not create new behavior, interfaces, data identities, dependencies, or provider relationships.

## Migration boundary

This file and `docs/catalog/project.json` remain **AUTHORITY NONE** migration artifacts. Legacy canonical owners are unchanged.

No `DAT-*`, `IFC-*`, `EXT-*`, `CAP-*`, or `DEC-*` identity is created by this slice. Feature relations for data, interfaces, dependencies, and capabilities remain on `UNK-001..UNK-004`; specifically, capability/build-buy remains unresolved through `UNK-004`. `architecture.build_buy` therefore remains `actual_depth: NONE`.

This state does not claim `DOCS_READY`, design lock, implementation readiness, assurance status, promotion, or release readiness.
