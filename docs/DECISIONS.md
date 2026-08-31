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
