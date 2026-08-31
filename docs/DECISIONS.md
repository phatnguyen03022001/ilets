# Decisions

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Created by `TASK-0003` revision 1 and extended through `TASK-0014` revision 1. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Current CAP/build-buy migration authority snapshot: `phatnguyen03022001/ilets@09504f4d782a5a53f9f8c57e750cc28274525667`; current accepted product-experience authority: `phatnguyen03022001/ilets@b60034a50d9a5ee5f197887ed14e7b917e919660`.

## Documentation-model adoption

**Draft outcome:** this migration uses `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07`, the accepted unreleased V1 candidate after the relation-model correction. This target pin supersedes the earlier pilot pin only for continuing the IELTS migration; it does not globally stabilize or release `agent-documents`.

The corrected feature relation state distinguishes typed `refs`, genuine demonstrated `na`, and `unresolved_ref` to an existing `OPEN` `DESIGN` `UNK-*`. Migration incompleteness is not N/A. An unresolved relation blocks documentation closure without creating a substitute neighboring-domain identity.

## Catalog migration boundary

`docs/catalog/project.json` exists only as a migration artifact with **authority NONE**. Its pathname does not grant canonical target authority. The catalog owns migration identities/typed relations; `docs/BEHAVIOR.md` owns migrated behavior explanation.

The migration includes the previously accepted `SYS-*`, `DAT-*`, `IFC-*`, `EXT-*`, and `CAP-*` inventories plus the existing material `DEC-*` records. `TASK-0013` adds only the decomposition decision required to resolve the already-approved `DECISION_REQUIRED` feature blocker and bounded relation-domain unknowns for detailed relations that current neighboring inventories do not yet fully support. `TASK-0014` extends only the DATA inventory/relations required to resolve `UNK-006`; it does not alter the IFC/EXT/CAP inventories or their remaining unknowns. These records remain authority-NONE migration state and cannot override canonical legacy owners.

## Engineering claim and evidence semantics adoption

IELTS continues to pin `phatnguyen03022001/agent-standards@3f4950f280a3a35fee81471d4b83715fa72cf9ee` for generic engineering claim/evidence semantics. This migration performs no standards assessment and claims no requirement `PASS`, maturity level, assurance `N/A`, exception, implementation readiness, promotion readiness, or release readiness. Coverage depth is documentation treatment, not engineering maturity.

## Ownership and cutover boundary

Actual IELTS truth remains with the existing canonical legacy documents. `docs/PRODUCT.md`, this file, `docs/BEHAVIOR.md`, and `docs/catalog/project.json` remain authority-NONE migration artifacts. No canonical cutover, `DOCS_READY`, documentation lock, generated PROGRAM, promotion, or release is declared.

## Migration unknowns

Not-yet-migrated material feature relations remain catalog-owned `OPEN` `DESIGN` `UNK-*` records referenced through `unresolved_ref`. A later authorized migration may replace each remaining relation only with correctly typed `refs` or genuine demonstrated `na` after the owning domain is resolved. Resolving one bounded relation domain does not imply closure of neighboring relation domains, and documentation closure remains blocked while the milestone is OPEN or downstream blockers remain.

## UNK-001 DATA relation resolution

`UNK-001` is resolved for the bounded six-macro behavior slice migrated by `TASK-0003`/`TASK-0005`. The current canonical owners support a material DATA inventory consisting of target/target-context revision, exact content revision, Attempt, Observation, EvidenceFact, current Assessment/Progression interpretations, historical attainment, and DailyPlan snapshot identities as materialized by `DAT-001` through `DAT-012` in `docs/catalog/project.json`.

Each affected macro feature uses typed `DAT-*` references rather than `unresolved_ref` or false N/A. `DAT-001` is the catalog `resolved_by_ref` attribution anchor for this non-decision question; the resolution itself is the complete typed macro feature-data mapping plus the DATA sections in `docs/DATA.md`, not a claim that `TargetProfile` alone covers every detailed feature.

This resolution does not assert database tables, API/event payloads, provider identities, exact retention durations, a durable representation for every material ephemeral value, or global DATA closure beyond that migrated slice. The new detailed feature inventory introduced later by `TASK-0013` is intentionally not projected onto these macro mappings without a dedicated DATA migration.

## UNK-002 INTERFACES relation resolution

`UNK-002` is resolved for the bounded six-macro behavior slice migrated by `TASK-0003`/`TASK-0006`. Current canonical flow/API/runtime owners support exactly two material interface identities for this slice: `IFC-001` Core public product API and `IFC-002` Core-to-Evaluator bounded capability API.

Each affected macro feature uses typed `IFC-*` references rather than `unresolved_ref` or false N/A. Flow `interface_refs` identify which of those two boundaries each migrated flow actually traverses. `IFC-001` is the catalog `resolved_by_ref` attribution anchor for this non-decision question; the resolution is the complete typed macro feature/flow interface mapping plus the semantic contract sections in `docs/INTERFACES.md`, not a claim that the public API alone covers every future detailed privileged boundary.

The async/job/command concern is resolved at L1 through the durable acceptance, logical-work, dispatch-fencing, retry, reconciliation, and recoverable-continuation semantics already owned by the two interfaces. No separate queue/job/provider interface identity is created because current canonical truth does not require an independently owned third boundary for the accepted macro slice.

This resolution does not materialize OpenAPI or exact wire schemas, generated bindings, provider authentication identities, or another semantic owner.

## UNK-003 External dependency relation resolution

`UNK-003` is resolved for the bounded six-macro behavior slice migrated by `TASK-0003`/`TASK-0009`. Current canonical third-party/API/runtime owners support exactly four material dependency identities for this slice: `EXT-001` Clerk identity/session service, `EXT-002` Neon Launch PostgreSQL hosting, `EXT-003` Google Cloud Tasks bounded dispatch, and `EXT-004` the external Evaluator AI/speech provider boundary.

The first three identities preserve the concrete routes already selected for implementation by canonical design. `EXT-004` is boundary-sized rather than model-alias-sized: the exact selected provider/model route matrix and any canonical `TBD` sub-capability remain owned by `design/07-third-party-services.md`. In particular, this migration does not turn a `TBD` pronunciation/acoustic route into a selected or available provider.

All six affected macro features use typed `EXT-*` references because each participates in the material `FLW-001` target-to-next-action loop. That feature-level relation records the material dependency footprint of its related flows; it does not claim every individual operation invokes every external service. Flow dependency references narrow the actual flow footprint: `FLW-001` and `FLW-003` use all four identities, `FLW-002` uses `EXT-001` and `EXT-002`, and `FLW-004` uses `EXT-002`, `EXT-003`, and `EXT-004`.

`EXT-001` is the catalog `resolved_by_ref` attribution anchor for this non-decision question; the resolution itself is the complete typed macro feature/flow dependency mapping plus the external-dependency and failure/fallback/exit sections in `docs/INTERFACES.md`, not a claim that identity/session service alone covers every dependency.

This resolution preserves provider lifecycle and authority boundaries: `SELECTED_FOR_IMPLEMENTATION` is not `ACTIVE`; provider output remains non-authoritative until Core/Evaluator validation and owning policy interpretation; fallback requires an independently equivalent consequence floor; timeout remains ambiguous; deletion/export and provider-held copies require reconciliation; and exit relies on stable Core identity, PostgreSQL-native portability, durable Core work identity, and provider-neutral Evaluator capability boundaries rather than speculative multi-provider infrastructure.

No provider is activated or configured here. Detailed privileged feature dependency expansion introduced by `TASK-0013` remains a separate explicit DESIGN unknown rather than being fabricated from this older macro mapping.

## DEC-BUILD-BUY Material build/buy decisions

`TASK-0010` records only three material build/buy decisions because the five capability records reduce to three genuine architectural choices: keep product semantic authority first-party; buy selected commodity infrastructure behind Core-owned authority; and keep the Evaluator project-owned while buying bounded external AI/speech execution. Provider names, prices, quotas, model benchmarks, activation state, and mutable operational research are intentionally not promoted into these decision records.

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

### DEC-004 Detailed product feature completeness surface

**Context.** Accepted `TASK-0012` reopened the milestone because the migration catalog had only six macro behavior anchors while canonical `design/01-skill-features.md` already defines 40 material learner feature capabilities and `design/04-application-flows.md` defines the material privileged surface/capability families. The pinned V1 model classifies `UNK-005` as `DECISION_REQUIRED`, so a truthful resolution requires a dedicated `DEC-*` attribution rather than pointing at an unrelated build/buy decision.

**Selected outcome.** Treat `FTR-007..FTR-046` as the exact 40 learner detailed feature-completeness units mapped one-for-one to `L-F01..L-F08`, `R-F01..R-F08`, `W-F01..W-F09`, `S-F01..S-F09`, and `X-F01..X-F06`. Treat `FTR-047..FTR-053` as the minimum seven privileged/admin/BOPS feature families: content/review/release, user support, entitlement reconciliation, operations/provider/work visibility, approved policy administration, security/access administration, and consequential audit. Retain `FTR-001..FTR-006` unchanged as historical macro route anchors rather than completeness units.

**Alternative rejected.** Continue treating six macro anchors as the complete product feature inventory, or inflate the catalog by converting navigation labels, 28 practice modes, Skill Leaves/micro-skills, Knowledge Objects, media/transcript states, `FREE`/`PRO`, providers/models, individual permissions, or the cross-feature AI tutor overlay into duplicate FTR identities.

**Rationale.** Canonical design already owns the decomposition and its abstraction boundaries. One-to-one learner capability mapping plus seven genuinely distinct privileged behavior families is the smallest complete product feature surface that preserves those owners without inventing a second ontology or capability matrix.

**Consequences.** Feature/flow/acceptance completeness can resolve independently of neighboring DATA/IFC/EXT/CAP migration completeness. Existing macro relation resolutions remain historically valid for `FTR-001..FTR-006`; newly detailed relations use exact existing refs only where their related accepted flows make those refs fully justified and otherwise remain explicit OPEN DESIGN unknowns. Milestone scope stays `OPEN`.

**Reversibility.** `COSTLY`. The mapping can evolve only when canonical product design materially changes; renumbering or collapsing these stable migration identities later would require explicit traceability migration rather than cosmetic editing.

## UNK-INVENTORY Migration unknown inventory after TASK-0014

`UNK-001..UNK-006` are resolved for their authorized migration slices. `UNK-006` resolves the complete detailed DATA relation surface only; `UNK-007`, `UNK-008`, and `UNK-009` remain the bounded downstream IFC, EXT, and CAP DESIGN blockers. Therefore milestone closure remains intentionally blocked even though product feature and DATA relation completeness are no longer ambiguous.

### UNK-004 Capability/build-buy relation resolution

`UNK-004` is resolved for all six original affected macro features because canonical runtime ownership plus the accepted `EXT-001..EXT-004` inventory supports a complete, typed, non-speculative capability mapping:

- every macro feature uses `CAP-001` for Core-owned authoritative product policy/orchestration;
- every macro feature uses `CAP-002` and `CAP-003` because authenticated durable product state remains a hybrid of Core authority with the selected identity/session and managed PostgreSQL boundaries;
- `FTR-002` through `FTR-006` use `CAP-004` and `CAP-005` where diagnostic/evaluation/support/replanning behavior materially crosses recoverable async dispatch and the bounded Evaluator provider boundary; and
- `FTR-001` does not claim those evaluator/dispatch capabilities merely because the broader `FLW-001` dependency footprint contains them.

`CAP-001` is the catalog `resolved_by_ref` attribution anchor for this non-decision question; the resolution is the complete typed macro feature-capability mapping plus `docs/ARCHITECTURE.md#ARCH-BUILD-BUY`, not a claim that one capability covers every newly decomposed feature consequence.

This resolution creates no pure `BUY` capability because the external dependencies do not own the whole capability boundary, and it creates no feature-referenced `DEFER` capability. Provider lifecycle remains canonical in `design/07-third-party-services.md`: selected routes remain only `SELECTED_FOR_IMPLEMENTATION`, pronunciation/acoustic evaluation remains `TBD`, and no provider/model is activated, calibrated, configured, or made authoritative here.

## UNK-005 Full learner/admin feature decomposition resolution

`UNK-005` is resolved by `DEC-004` because current canonical design now maps completely into the migration feature surface without abstraction inflation:

- `FTR-007..FTR-046` map one-for-one to all 40 canonical learner capability aliases in `design/01-skill-features.md`;
- `FTR-047..FTR-053` cover every material privileged/admin/BOPS surface family admitted by `TASK-0012` and `design/04-application-flows.md` without creating one feature per permission or role bundle;
- `FTR-001..FTR-006` remain unchanged historical macro anchors, so prior accepted unknown/relation lineage remains intact;
- learner detailed features reuse the existing macro learner flows where truthful, while only two new grouped privileged flows are added for capability-gated operations and consequential audit;
- grouped `ACC-008..ACC-015`, together with the existing acceptance anchors where applicable, preserve feature-identity boundaries, variant/delivery truth, evidence authority, Speaking capture/realtime limitations, media/knowledge layering, entitlement/RBAC separation, Core-only privilege, and audit requirements; and
- navigation, practice modes, ontology nodes, media/source states, entitlement states, provider/model routes, and AI tutoring remain explicitly outside FTR identity.

`DEC-004` is the required `resolved_by_ref` attribution because the pinned V1 model requires a `DECISION_REQUIRED` unknown to resolve through a `DEC-*`. This is a decomposition/migration decision only; it does not create new canonical product authority.

Resolving `UNK-005` does **not** imply neighboring relation closure. The CAP mapping for all detailed features plus the IFC and EXT mapping for newly decomposed privileged surfaces remain separate OPEN DESIGN relation unknowns below. `milestone.scope_state` therefore stays `OPEN`, and `DOCS_READY`, design lock, cutover, implementation readiness, standards PASS, provider activation, promotion, and release remain unclaimed.

## Detailed relation-domain blockers after TASK-0014

### UNK-006 Expanded detailed feature DATA relations

`UNK-006` is resolved only after every `FTR-007..FTR-053` DATA relation is replaced with typed `DAT-*` references and the complete detailed mapping is semantically audited. No detailed feature requires a genuine DATA N/A: every one reads, creates, mutates, derives, or materially depends on at least one semantic DATA identity.

The accepted `DAT-001..DAT-012` identities remain unchanged. Canonical detailed behavior requires the additional minimal identities `DAT-013..DAT-027`: FeedbackArtifact; ReviewState; MediaSource; ContentCandidate; ValidationDecision history; content release/operational eligibility; effective entitlement; support-managed learner/account product state; authorization grants; typed operating-policy revision; recoverable operational work; consequential privileged audit; MockRun; LearningSession; and CoverageGap. These are boundary-sized semantic identities, not features, screens, tables, DTOs, queue messages, provider states, caches, or one object per workflow.

The mapping intentionally reuses existing learner/evidence/progression identities where they already own the meaning: learner-saved Knowledge/review material remains normal `ContentRevision` plus Knowledge-scoped `MasteryEstimate` and `ReviewState`; Error/Remediation semantics remain content/feedback references rather than new DATA ontologies; content candidate, immutable revision, validation history, and current release/operational state stay distinct because their lifecycle/history semantics differ; entitlement, authorization, and operating policy stay distinct because commercial access, privileged access, and runtime policy are different authorities; operational work does not become provider state; audit does not become authorization or telemetry; and CoverageGap never becomes learner `GapEvaluation`.

`DAT-013` is the catalog `resolved_by_ref` attribution anchor for this non-decision question. The resolution is the complete `FTR-007..FTR-053` typed DATA mapping plus `docs/DATA.md`, not a claim that FeedbackArtifact alone covers detailed DATA. Exact persistence schemas, DTOs, provider activation/state, caches, storage buckets, and retention durations remain outside this resolution.

### UNK-007 Privileged detailed interface relations

The two accepted interface identities are sufficient for the original macro flows, and learner detailed features can reuse their exact interface footprint when they reuse those accepted flows. The newly decomposed privileged feature families, however, have not yet received a dedicated INTERFACES migration proving whether `IFC-001` alone fully represents each privileged read/mutation/audit consequence or whether another semantic boundary is material. `FTR-047..FTR-053` therefore remain behind one OPEN DESIGN IFC relation unknown rather than fabricating exact contracts or false N/A.

### UNK-008 Privileged detailed external-dependency relations

The accepted `EXT-001..EXT-004` inventory describes identity/session, PostgreSQL hosting, task dispatch, and the evaluator AI/speech boundary for the earlier behavior slice. Privileged content operations, entitlement reconciliation, provider/work visibility, policy/security administration, and audit may have material dependency relations not proven by that slice, including commercial/provider facts that must not be invented here. `FTR-047..FTR-053` therefore remain behind one OPEN DESIGN EXT relation unknown; no provider is selected or activated by this task.

### UNK-009 Expanded detailed feature capability/build-buy relations

The accepted `CAP-001..CAP-005` mappings were resolved for the six macro anchors. Projecting those capability/build-buy refs across all 47 newly detailed features would require feature-specific boundary proof and could hide missing privileged/commercial/audit capability state. `FTR-007..FTR-053` therefore use one bounded OPEN DESIGN CAP relation unknown until a dedicated migration performs that mapping. Existing CAP dispositions and `DEC-001..DEC-003` remain unchanged.

The remaining three relation-domain blockers are intentionally separate from resolved feature/DATA completeness. They keep the reopened milestone truthful without reopening `UNK-005`/`UNK-006` or manufacturing neighboring-domain coverage.