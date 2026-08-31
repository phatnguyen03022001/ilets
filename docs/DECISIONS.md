# Decisions

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Maintained through `TASK-0003` revision 1. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Canonical target snapshot: `phatnguyen03022001/ilets@2f54b7c6003399ccf7abff8cda2277d68d0048e8`.

## Documentation-model adoption

**Draft outcome:** this migration uses `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07`, the accepted unreleased V1 candidate after the relation-model correction. This target pin supersedes the earlier pilot pin only for continuing the IELTS migration; it does not globally stabilize or release `agent-documents`.

The corrected feature relation state distinguishes typed `refs`, genuine demonstrated `na`, and `unresolved_ref` to an existing `OPEN` `DESIGN` `UNK-*`. Migration incompleteness is not N/A. An unresolved relation blocks documentation closure without creating a substitute neighboring-domain identity.

## Catalog migration boundary

`docs/catalog/project.json` now exists only as a migration artifact with **authority NONE**. Its pathname does not grant canonical target authority. The catalog owns migration identities/typed relations; `docs/BEHAVIOR.md` owns migrated behavior explanation.

This slice may assign directly supported `ACT-*`, `ROL-*`, `FTR-*`, `ACC-*`, `FLW-*`, and `UNK-*` identities. It does not prematurely materialize `SYS-*`, `DAT-*`, `IFC-*`, `EXT-*`, or `CAP-*`.

## Engineering claim and evidence semantics adoption

IELTS continues to pin `phatnguyen03022001/agent-standards@3f4950f280a3a35fee81471d4b83715fa72cf9ee` for generic engineering claim/evidence semantics. This migration performs no standards assessment and claims no requirement `PASS`, maturity level, assurance `N/A`, exception, implementation readiness, promotion readiness, or release readiness. Coverage depth is documentation treatment, not engineering maturity.

## Ownership and cutover boundary

Actual IELTS truth remains with the existing canonical legacy documents. `docs/PRODUCT.md`, this file, `docs/BEHAVIOR.md`, and `docs/catalog/project.json` remain authority-NONE migration artifacts. No canonical cutover, `DOCS_READY`, documentation lock, generated PROGRAM, promotion, or release is declared.

## Migration unknowns

`TASK-0003` records not-yet-migrated material feature relations as catalog-owned `OPEN` `DESIGN` `UNK-*` records referenced through `unresolved_ref`. A later authorized migration may replace each relation only with correctly typed `refs` or genuine demonstrated `na` after the owning domain is resolved. Until then, documentation closure remains blocked.

## UNK-001 DATA relation resolution

`UNK-001` is resolved for the bounded behavior slice migrated by `TASK-0003`/`TASK-0005`. The current canonical owners support a material DATA inventory consisting of target/target-context revision, exact content revision, Attempt, Observation, EvidenceFact, current Assessment/Progression interpretations, historical attainment, and DailyPlan snapshot identities as materialized by `DAT-001` through `DAT-012` in `docs/catalog/project.json`.

Each affected feature now uses typed `DAT-*` references rather than `unresolved_ref` or false N/A. `DAT-001` is the catalog `resolved_by_ref` attribution anchor for this non-decision question; the resolution itself is the complete typed feature-data mapping plus the DATA sections in `docs/DATA.md`, not a claim that `TargetProfile` alone covers every feature.

This resolution does not assert database tables, API/event payloads, provider identities, exact retention durations, a durable representation for every material ephemeral value, or global DATA closure beyond this migrated slice. `UNK-002`, `UNK-003`, and `UNK-004` remain `OPEN` `DESIGN` questions for their own authority domains.

## UNK-002 INTERFACES relation resolution

`UNK-002` is resolved for the bounded behavior slice migrated by `TASK-0003`/`TASK-0006`. Current canonical flow/API/runtime owners support exactly two material interface identities for this slice: `IFC-001` Core public product API and `IFC-002` Core-to-Evaluator bounded capability API.

Each affected feature now uses typed `IFC-*` references rather than `unresolved_ref` or false N/A. Flow `interface_refs` identify which of those two boundaries each migrated flow actually traverses. `IFC-001` is the catalog `resolved_by_ref` attribution anchor for this non-decision question; the resolution is the complete typed feature/flow interface mapping plus the semantic contract sections in `docs/INTERFACES.md`, not a claim that the public API alone covers every runtime boundary.

The async/job/command concern is resolved at L1 through the durable acceptance, logical-work, dispatch-fencing, retry, reconciliation, and recoverable-continuation semantics already owned by the two interfaces. No separate queue/job/provider interface identity is created because current canonical truth does not require an independently owned third boundary.

This resolution does not materialize OpenAPI or exact wire schemas, generated bindings, provider `EXT-*` identities, deployment/provider authentication identities, build/buy `CAP-*` state, or another semantic owner. `interfaces.external_dependencies` and `interfaces.dependency_failure_exit` remain intentionally `NONE`; `UNK-003` and `UNK-004` remain `OPEN` `DESIGN` questions for their own authority domains.

## UNK-003 External dependency relation resolution

`UNK-003` is resolved for the bounded behavior slice migrated by `TASK-0003`/`TASK-0009`. Current canonical third-party/API/runtime owners support exactly four material dependency identities for this slice: `EXT-001` Clerk identity/session service, `EXT-002` Neon Launch PostgreSQL hosting, `EXT-003` Google Cloud Tasks bounded dispatch, and `EXT-004` the external Evaluator AI/speech provider boundary.

The first three identities preserve the concrete routes already selected for implementation by canonical design. `EXT-004` is boundary-sized rather than model-alias-sized: the exact selected provider/model route matrix and any canonical `TBD` sub-capability remain owned by `design/07-third-party-services.md`. In particular, this migration does not turn a `TBD` pronunciation/acoustic route into a selected or available provider.

All six affected features now use typed `EXT-*` references rather than `unresolved_ref` or false N/A because each feature participates in the material `FLW-001` target-to-next-action loop. That feature-level relation records the material dependency footprint of its related flows; it does not claim every individual operation invokes every external service. Flow dependency references narrow the actual flow footprint: `FLW-001` and `FLW-003` use all four identities, `FLW-002` uses `EXT-001` and `EXT-002`, and `FLW-004` uses `EXT-002`, `EXT-003`, and `EXT-004`.

`EXT-001` is the catalog `resolved_by_ref` attribution anchor for this non-decision question; the resolution itself is the complete typed feature/flow dependency mapping plus the external-dependency and failure/fallback/exit sections in `docs/INTERFACES.md`, not a claim that identity/session service alone covers every dependency.

This resolution preserves provider lifecycle and authority boundaries: `SELECTED_FOR_IMPLEMENTATION` is not `ACTIVE`; provider output remains non-authoritative until Core/Evaluator validation and owning policy interpretation; fallback requires an independently equivalent consequence floor; timeout remains ambiguous; deletion/export and provider-held copies require reconciliation; and exit relies on stable Core identity, PostgreSQL-native portability, durable Core work identity, and provider-neutral Evaluator capability boundaries rather than speculative multi-provider infrastructure.

No provider is activated or configured here. No `CAP-*`, `DEC-*`, build/buy disposition, alternative-provider selection, pricing state, deployment change, or `UNK-004` closure is created. Earlier statements in this file that `UNK-003` remained open describe the prior migration slices and are superseded only by this dedicated `TASK-0009` resolution.
