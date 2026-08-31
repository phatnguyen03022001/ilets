# Decisions

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Created by `TASK-0003` revision 1 and extended through `TASK-0010` revision 1. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Current CAP/build-buy migration authority snapshot: `phatnguyen03022001/ilets@09504f4d782a5a53f9f8c57e750cc28274525667`.

## Documentation-model adoption

**Draft outcome:** this migration uses `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07`, the accepted unreleased V1 candidate after the relation-model correction. This target pin supersedes the earlier pilot pin only for continuing the IELTS migration; it does not globally stabilize or release `agent-documents`.

The corrected feature relation state distinguishes typed `refs`, genuine demonstrated `na`, and `unresolved_ref` to an existing `OPEN` `DESIGN` `UNK-*`. Migration incompleteness is not N/A. An unresolved relation blocks documentation closure without creating a substitute neighboring-domain identity.

## Catalog migration boundary

`docs/catalog/project.json` now exists only as a migration artifact with **authority NONE**. Its pathname does not grant canonical target authority. The catalog owns migration identities/typed relations; `docs/BEHAVIOR.md` owns migrated behavior explanation.

The migration now includes the previously accepted `SYS-*`, `DAT-*`, `IFC-*`, and `EXT-*` inventories plus the bounded `CAP-*` / `DEC-*` records added by `TASK-0010`. Those records remain authority-NONE migration state and cannot override the canonical legacy owners.

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


## DEC-BUILD-BUY Material build/buy decisions

`TASK-0010` records only three material decisions because the five capability records reduce to three genuine architectural choices: keep product semantic authority first-party; buy selected commodity infrastructure behind Core-owned authority; and keep the Evaluator project-owned while buying bounded external AI/speech execution. Provider names, prices, quotas, model benchmarks, activation state, and mutable operational research are intentionally not promoted into these decision records.

### DEC-001 Core product authority remains project-owned

**Context.** Canonical runtime design assigns learner/admin product behavior, product authorization, authoritative work state, deterministic Assessment/Progression/Planner policy, and final result interpretation to Core. External AI/provider output is explicitly non-authoritative.

**Selected outcome.** Build the authoritative product policy/orchestration capability in `SYS-002` (`CAP-001`) instead of delegating product semantics to an external platform or model provider.

**Alternative rejected.** Treat a provider/model/backend service as the owner of learner evidence, progression, certification, or next-action policy.

**Rationale.** The project must preserve deterministic product meaning, stable semantic identity, explainable failure behavior, and the Web/Core/Evaluator authority split. Provider convenience or model capability cannot safely substitute for these owners.

**Consequences.** External services may support execution, but Core remains the single authoritative product interpretation boundary. Provider success, confidence, or generated text cannot independently mutate learner/product truth.

**Reversibility.** `COSTLY`. A future architecture may redistribute product responsibility, but doing so requires an explicit semantic migration and architecture change rather than a provider/configuration switch.

### DEC-002 Commodity infrastructure stays behind Core-owned authority

**Context.** The current implementation selects external identity/session service, managed PostgreSQL hosting, and bounded task dispatch, while canonical design separately preserves stable Core identity/authorization, PostgreSQL product-state semantics, and durable Core logical-work/reconciliation state.

**Selected outcome.** Use `EXT-001`, `EXT-002`, and `EXT-003` as bounded external execution/custody inside the `CAP-002`/`CAP-003`/`CAP-004` hybrid capabilities. Core keeps the authority that makes each provider replaceable.

**Alternative rejected.** Promote provider roles/metadata, hosted-database product semantics, or queue/task state into product truth; or introduce duplicate identity/database/queue infrastructure merely to claim portability.

**Rationale.** These concerns are commodity enough to reuse selected external services, but their surrounding product authority is not. One common decision therefore explains all three boundaries without inventing a provider-specific decision per catalog identity.

**Consequences.** `SELECTED_FOR_IMPLEMENTATION` remains distinct from `ACTIVE`. Core must preserve stable internal identity, sole authoritative persistence access, and durable work/reconciliation semantics. No second live database, broker, dynamic failover layer, or generic infrastructure framework is implied.

**Reversibility.** `COSTLY`. Exit is explicitly feasible but may require principal re-association, PostgreSQL-native migration/recovery, or in-flight work reconciliation. Those costs do not justify transferring authority to the provider.

### DEC-003 Evaluator stays project-owned with external execution adapters

**Context.** The Evaluator is the internal bounded capability owner. Canonical third-party design selects some external productive-evaluation/speech routes while leaving pronunciation/acoustic evaluation `TBD`, and requires provider-neutral boundaries, provenance, validation, and consequence-specific calibration.

**Selected outcome.** Implement `CAP-005` as `HYBRID`: keep the bounded evaluator contract/validation surface in `SYS-003`, and use `EXT-004` for selected-or-TBD external AI/speech execution behind adapters.

**Alternative rejected.** Expose provider SDK/model schemas as the internal product contract, make a provider the product/evidence owner, or create one capability identity per provider/model alias.

**Rationale.** The project needs external model/speech execution where deterministic/local execution is insufficient, but external output is still untrusted and provider lifecycle differs by sub-capability. A project-owned Evaluator boundary preserves semantics while permitting bounded substitution.

**Consequences.** Provider/model replacement requires preserved provenance and fresh consequence-specific validation/calibration. The pronunciation/acoustic provider remains `TBD`; this decision does not select, activate, or validate that route.

**Reversibility.** `COSTLY`. The provider is replaceable behind the boundary, but consequential replacement may require adapter work, calibration, migration/reconciliation of in-flight work, and renewed eligibility evidence.

## UNK-INVENTORY Migration unknown inventory after TASK-0010

The four catalog-owned migration questions are now resolved for the bounded behavior slice:

- `UNK-001` → typed `DAT-*` feature relations from the DATA migration;
- `UNK-002` → typed `IFC-*` feature/flow relations from the INTERFACES migration;
- `UNK-003` → typed `EXT-*` feature/flow dependency relations from the external-dependency migration; and
- `UNK-004` → typed `CAP-*` feature capability relations from this CAP/build-buy migration.

Earlier slice-local statements that later unknowns remained `OPEN` describe the state at those earlier tasks and are superseded by the dedicated later resolution sections. This inventory does not assert that the milestone is closed: `milestone.scope_state` remains `OPEN`, so the pinned model is expected to continue reporting `SCOPE_OPEN` and must not be interpreted as `DOCS_READY`.

### UNK-004 Capability/build-buy relation resolution

`UNK-004` is resolved for all six affected features because canonical runtime ownership plus the accepted `EXT-001..EXT-004` inventory supports a complete, typed, non-speculative capability mapping:

- every feature uses `CAP-001` for Core-owned authoritative product policy/orchestration;
- every feature uses `CAP-002` and `CAP-003` because authenticated durable product state remains a hybrid of Core authority with the selected identity/session and managed PostgreSQL boundaries;
- `FTR-002` through `FTR-006` use `CAP-004` and `CAP-005` where diagnostic/evaluation/support/replanning behavior materially crosses recoverable async dispatch and the bounded Evaluator provider boundary; and
- `FTR-001` does not claim those evaluator/dispatch capabilities merely because the broader `FLW-001` dependency footprint contains them.

`CAP-001` is the catalog `resolved_by_ref` attribution anchor for this non-decision question; the resolution is the complete typed feature-capability mapping plus `docs/ARCHITECTURE.md#ARCH-BUILD-BUY`, not a claim that one capability covers every feature consequence.

This resolution creates no pure `BUY` capability because the external dependencies do not own the whole capability boundary, and it creates no feature-referenced `DEFER` capability. Provider lifecycle remains canonical in `design/07-third-party-services.md`: selected routes remain only `SELECTED_FOR_IMPLEMENTATION`, pronunciation/acoustic evaluation remains `TBD`, and no provider/model is activated, calibrated, configured, or made authoritative here.

All coverage/resolution gaps addressed by this migration can reach zero while milestone closure still remains intentionally blocked by `scope_state: OPEN`. No `DOCS_READY`, cutover, design lock, implementation readiness, standards PASS, promotion, or release claim follows from this resolution.
