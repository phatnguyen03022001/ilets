# Architecture

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Created by `TASK-0004` revision 1 and extended through `TASK-0010` revision 1. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Execution base: `phatnguyen03022001/ilets@64d5915a2d31ec5eec025ccaf05b11aef60e9933`.
>
> Model pin: `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07` (unreleased V1 candidate).

This draft migrates the established logical runtime responsibility layer plus the bounded capability/build-buy consequences supported by canonical `design/06-implementation-stack.md`, `design/07-third-party-services.md`, and the accepted `SYS-*` / `EXT-*` migration state. `docs/catalog/project.json` owns the typed identities and relations; this file explains the migrated architecture only.

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

## ARCH-BUILD-BUY Capability and build/buy boundaries

The migrated behavior slice needs five material capability boundaries. They are deliberately larger than vendors, libraries, model aliases, endpoints, or minor implementation concerns. The dispositions describe who implements the capability boundary; they do not promote any external route from `SELECTED_FOR_IMPLEMENTATION` to `ACTIVE`.

| Capability | Disposition | Project owner | External boundary | Material decision |
| --- | --- | --- | --- | --- |
| `CAP-001` Core-owned product policy and orchestration | `BUILD` | `SYS-002` | none | `DEC-001` |
| `CAP-002` Identity and session with Core-owned authorization | `HYBRID` | `SYS-002` | `EXT-001` | `DEC-002` |
| `CAP-003` Authoritative PostgreSQL persistence on managed hosting | `HYBRID` | `SYS-002` | `EXT-002` | `DEC-002` |
| `CAP-004` Recoverable asynchronous dispatch | `HYBRID` | `SYS-002` | `EXT-003` | `DEC-002` |
| `CAP-005` Bounded evaluator AI and speech execution | `HYBRID` | `SYS-003` | `EXT-004` | `DEC-003` |

There is no pure `BUY` record in this slice because none of the four external dependencies owns the whole product capability it supports. Identity/session, hosted persistence, dispatch, and evaluator-provider execution all remain coupled to project-owned responsibility that cannot be transferred to the provider without changing canonical architecture. There is no feature-referenced `DEFER` record; deferred/TBD provider sub-routes remain lifecycle truth inside the existing external boundary instead of becoming fake resolved feature capabilities.

### CAP-001 Core-owned product policy and orchestration

**Boundary.** The authoritative product capability is built in `SYS-002` Core API: learner/admin product behavior, product authorization, deterministic Assessment/Progression/Planner orchestration, authoritative work identity/state, and final interpretation/reconciliation remain project-owned. Web remains presentation authority and Evaluator/provider output remains bounded input rather than product truth.

**Build/buy consequence.** Frameworks, hosted infrastructure, and external model output may support this capability, but none may become the semantic owner. Moving authoritative learning/evidence/progression/next-action policy into a provider would be an architecture change, not a provider configuration change.

**Exit.** Core implementation technology may be replaced or rewritten only while canonical product semantics, authoritative identity/state, and the Web/Core/Evaluator authority split remain preserved or are explicitly re-architected. No external provider exit is required for the `BUILD` disposition itself.

### CAP-002 Identity and session with Core-owned authorization

**Boundary.** `SYS-002` jointly implements authenticated product access with `EXT-001`: the external service owns credential/session mechanics and external-principal issuance, while Core owns stable internal actor/learner identity, principal association, RBAC/capabilities, entitlement, and product authorization.

**Build/buy consequence.** This is `HYBRID`, not `BUY`, because delegating credential custody does not delegate product identity or authorization truth.

**Exit.** Provider replacement re-associates a new external principal to the same Core-owned actor/learner semantics. Provider roles, organizations, permissions, or metadata must not become migration-critical product authority.

### CAP-003 Authoritative PostgreSQL persistence on managed hosting

**Boundary.** `SYS-002` owns authoritative product persistence semantics and is the sole application runtime permitted to read/write that state; `EXT-002` supplies managed PostgreSQL hosting only. Database schema remains derived implementation detail rather than learner/evidence/progression authority.

**Build/buy consequence.** This is `HYBRID`: Core owns transaction/state semantics while the hosting platform supplies commodity execution/operations. A second live database or provider-specific product-state model is not introduced.

**Exit.** Use the smallest PostgreSQL-native export/restore/recovery path that preserves canonical state and migration compatibility. Provider exit does not move application access away from Core or require automatic multi-cloud failover.

### CAP-004 Recoverable asynchronous dispatch

**Boundary.** `SYS-002` owns acceptance, durable logical work identity/state, dispatch admission/claim/fencing, retry identity, result reconciliation, and semantic completion; `EXT-003` performs bounded external task delivery only.

**Build/buy consequence.** This is `HYBRID`: the external dispatcher cannot become business-work, evidence, cancellation, or completion authority. No second broker, Pub/Sub layer, or generic queue framework is created.

**Exit.** Durable Core work/recovery state permits replacement of the dispatcher without treating provider queue state as accepted truth. Ambiguous in-flight attempts must be reconciled before safe redrive through a replacement route.

### CAP-005 Bounded evaluator AI and speech execution

**Boundary.** `SYS-003` owns the bounded evaluator capability contract and validation/provenance surface; `EXT-004` supplies selected-or-TBD external AI/speech execution behind provider-neutral adapters. Core remains the permitted product caller and retains final product interpretation.

**Build/buy consequence.** This is `HYBRID`: external models/services can execute bounded work but cannot own learner evidence, Band certification, progression, content activation, or final next-action semantics. Provider/model aliases do not become separate `CAP-*` identities.

The canonical pronunciation/acoustic provider route remains `TBD`. `CAP-005` therefore means the evaluator boundary is resolved, not that every provider sub-capability is selected, calibrated, available, or active.

**Exit.** Provider/model replacement occurs behind the same bounded Evaluator contract with consequential provenance retained and consequence-specific validation/calibration performed before use. Dynamic routing, generic multi-provider infrastructure, or a second semantic owner is not implied.

## Migration boundary

This file and `docs/catalog/project.json` remain **AUTHORITY NONE** migration artifacts. Legacy canonical owners are unchanged.

`TASK-0010` adds only the five capability identities above and their catalog decision references. Existing `SYS-*` and `EXT-*` identities are reused without changing provider lifecycle, activation, configuration, pricing, or exact contracts. The earlier DATA/INTERFACES/external-dependency migrations remain separate accepted migration state.

`milestone.scope_state` remains `OPEN`. This state does not claim `DOCS_READY`, design lock, cutover, implementation readiness, standards PASS, promotion, or release readiness.
