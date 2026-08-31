# Decisions

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Created by `TASK-0003` revision 1 and extended through `TASK-0017` revision 1. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Current CAP/build-buy migration authority snapshot: `phatnguyen03022001/ilets@09504f4d782a5a53f9f8c57e750cc28274525667`; current accepted product-experience authority: `phatnguyen03022001/ilets@b60034a50d9a5ee5f197887ed14e7b917e919660`.

## Documentation-model adoption

**Draft outcome:** this migration uses `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07`, the accepted unreleased V1 candidate after the relation-model correction. This target pin supersedes the earlier pilot pin only for continuing the IELTS migration; it does not globally stabilize or release `agent-documents`.

The corrected feature relation state distinguishes typed `refs`, genuine demonstrated `na`, and `unresolved_ref` to an existing `OPEN` `DESIGN` `UNK-*`. Migration incompleteness is not N/A. An unresolved relation blocks documentation closure without creating a substitute neighboring-domain identity.

## Catalog migration boundary

`docs/catalog/project.json` exists only as a migration artifact with **authority NONE**. Its pathname does not grant canonical target authority. The catalog owns migration identities/typed relations; `docs/BEHAVIOR.md` owns migrated behavior explanation.

The migration includes the previously accepted `SYS-*`, `DAT-*`, `IFC-*`, `EXT-*`, and `CAP-*` inventories plus the existing material `DEC-*` records. `TASK-0013` adds only the decomposition decision required to resolve the already-approved `DECISION_REQUIRED` feature blocker and bounded relation-domain unknowns for detailed relations that current neighboring inventories do not yet fully support. `TASK-0014` extends only the DATA inventory/relations required to resolve `UNK-006`. `TASK-0015` resolves only the privileged detailed IFC relations and `UNK-007`; it preserves both accepted IFC identities and does not alter EXT/CAP mappings. `TASK-0016` resolves only the privileged detailed EXT relations and `UNK-008`, adding the minimum external dependency identities justified by current privileged behavior. `TASK-0017` resolves only the 47 detailed CAP relations and `UNK-009`, preserving `CAP-001..CAP-005` and adding only the three distinct build/buy boundaries supported by accepted `EXT-005..EXT-007`. These records remain authority-NONE migration state and cannot override canonical legacy owners.

## Engineering claim and evidence semantics adoption

IELTS continues to pin `phatnguyen03022001/agent-standards@3f4950f280a3a35fee81471d4b83715fa72cf9ee` for generic engineering claim/evidence semantics. This migration performs no standards assessment and claims no requirement `PASS`, maturity level, assurance `N/A`, exception, implementation readiness, promotion readiness, or release readiness. Coverage depth is documentation treatment, not engineering maturity.

## Ownership and cutover boundary

Actual IELTS truth remains with the existing canonical legacy documents. `docs/PRODUCT.md`, this file, `docs/BEHAVIOR.md`, and `docs/catalog/project.json` remain authority-NONE migration artifacts. No canonical cutover, `DOCS_READY`, documentation lock, generated PROGRAM, promotion, or release is declared.

## Migration unknowns

All material feature relations covered by `UNK-001..UNK-009` are now migrated: those unknowns are `RESOLVED` and the catalog has no remaining `unresolved_ref` for those bounded relation domains. The resolved mappings use correctly typed refs or demonstrated N/A as accepted by their owning migrations. `milestone.scope_state` remains `OPEN` only because the post-reconciliation semantic closure/freeze audit has not yet passed; this section does not freeze scope or claim `DOCS_READY`.

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

No provider is activated or configured here. The detailed privileged dependency relation that `TASK-0013` left open is resolved separately by `TASK-0016` under `UNK-008`; that later mapping does not rewrite this older macro relation footprint.

## DEC-BUILD-BUY Material build/buy decisions

`TASK-0010` records three material build/buy decisions for the original five capability boundaries: keep product semantic authority first-party; buy selected commodity identity/persistence/dispatch behind Core-owned authority; and keep the Evaluator project-owned while buying bounded external AI/speech execution. `TASK-0017` adds three further decisions only because accepted `EXT-005..EXT-007` introduce materially different artifact-custody, commercial-reconciliation, and secret-custody responsibility/exit semantics. Provider names, prices, quotas, model benchmarks, activation state, and mutable operational research are intentionally not promoted into these decision records.

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

**Selected outcome.** Treat `FTR-007..FTR-046` as the exact 40 learner detailed feature-completeness units mapped one-for-one to `L-F01..L-F08`, `R-F01..R-F08`, `W-F01..W-F09`, `S-F01..S-F09`, and `X-F01..X-F06`. Treat `FTR-047..FTR-053` as the minimum seven privileged/admin/BOPS feature families: content/review/release, user support, entitlement reconciliation, operations/provider/work visibility, approved policy administration, security/access administration, and consequential audit. Retain `FTR-001..FTR-006` unchanged as historical/macro route anchors rather than completeness units.

**Alternative rejected.** Continue treating six macro anchors as the complete product feature inventory, or inflate the catalog by converting navigation labels, 28 practice modes, Skill Leaves/micro-skills, Knowledge Objects, media/transcript states, `FREE`/`PRO`, providers/models, individual permissions, or the cross-feature AI tutor overlay into duplicate FTR identities.

**Rationale.** Canonical design already owns the decomposition and its abstraction boundaries. One-to-one learner capability mapping plus seven genuinely distinct privileged behavior families is the smallest complete product feature surface that preserves those owners without inventing a second ontology or capability matrix.

**Consequences.** Feature/flow/acceptance completeness can resolve independently of neighboring DATA/IFC/EXT/CAP migration completeness. Existing macro relation resolutions remain historically valid for `FTR-001..FTR-006`; newly detailed relations were resolved independently by `TASK-0014..TASK-0017` using exact refs or demonstrated N/A where justified. Milestone scope stays `OPEN` pending the post-reconciliation semantic closure/freeze audit.

**Reversibility.** `COSTLY`. The mapping can evolve only when canonical product design materially changes; renumbering or collapsing these stable migration identities later would require explicit traceability migration rather than cosmetic editing.

### DEC-005 Governed object/media custody stays behind Core artifact authority

**Context.** Canonical implementation and third-party design permit large or retained content/media bytes to live in the accepted `EXT-005` object/media storage boundary while Core remains authoritative for artifact identity, metadata/reference, eligibility, access, lifecycle, retention/deletion, integrity/provenance reconciliation, and product usability.

**Selected outcome.** Implement `CAP-006` as `HYBRID`: keep governed artifact semantics and lifecycle in `SYS-002`, and use `EXT-005` only for external byte custody and bounded transfer behind Core authorization/reconciliation.

**Alternative rejected.** Treat storage bucket/object keys, signed-upload success, CDN state, or provider metadata as content/product authority; or create a separate capability per bucket, object class, media type, transfer route, or storage vendor.

**Rationale.** Byte custody is commodity infrastructure, but object/media lifecycle and usability are product/security/data semantics. Their exit path also requires artifact-integrity, deletion/orphan, and in-flight-transfer reconciliation that is materially different from identity, database hosting, or task dispatch.

**Consequences.** The selected external route remains only `SELECTED_FOR_IMPLEMENTATION`; no activation, public permanence, data-egress approval, storage plan, bucket, credential, or spend is authorized here. Direct browser transfer remains narrow, temporary, and Core-authorized where later implemented.

**Reversibility.** `COSTLY`. Provider replacement can preserve Core-owned artifact identity while moving eligible bytes, but must reconcile integrity, deletion/tombstone state, orphaned/temporary objects, and ambiguous transfers before new custody becomes usable.

### DEC-006 Effective commercial entitlement remains Core-owned

**Context.** Canonical privileged behavior defines entitlement reconciliation as accepting bounded commercial/provider facts and committing effective product access in Core. The accepted `EXT-006` boundary supplies payment/subscription observations only; learner commercial entitlement is separate from RBAC and learning/evidence truth.

**Selected outcome.** Implement `CAP-007` as `HYBRID`: keep effective entitlement, stable learner association, and product-access consequence in `SYS-002`, while `EXT-006` supplies externally observed payment/subscription facts and provider events that Core reconciles.

**Alternative rejected.** Treat provider subscription/checkout state as authoritative product entitlement, map provider roles/plans directly into RBAC or learning truth, or create one capability per webhook, billing product, payment method, plan, or commerce endpoint.

**Rationale.** Commercial processing can be bought, but the meaning of paid access inside this product cannot. Reconciliation, replay/idempotency, downgrade/expiry semantics, and exit from a commerce provider are materially different from general product policy or storage custody.

**Consequences.** Provider observations remain untrusted external facts until authenticated/associated/reconciled. No checkout, webhook, billing credential, price, provider activation, spend, or commercial policy is created by this migration.

**Reversibility.** `COSTLY`. Provider change preserves Core-owned effective entitlement/history while outstanding callbacks, refunds/cancellations, ambiguous events, and provider-account association are reconciled before cutover.

### DEC-007 Secret custody remains external while security authority stays in Core

**Context.** Canonical runtime design separates secrets, Core-owned typed runtime policy, and deployment/bootstrap configuration. `EXT-007` is the accepted external secret-custody boundary, while canonical privileged behavior reserves credential creation/rotation/switching and other security-sensitive operations for stronger authorization plus reconstructable audit.

**Selected outcome.** Implement `CAP-008` as `HYBRID`: keep security authorization and typed secret-operation policy in `SYS-002`, and use `EXT-007` only to custody/version credential material behind least-privilege runtime access.

**Alternative rejected.** Make a secret manager, environment-variable store, provider console, or plaintext export path the owner of authorization/runtime policy; or create a capability per credential, secret version, provider, environment variable, or administrative button.

**Rationale.** Secret material has disclosure/rotation/revocation and least-privilege exit semantics that differ materially from ordinary operating policy, content objects, commercial observations, or identity/session mechanics. Core must retain the permission/audit meaning even if custody changes.

**Consequences.** Routine administration exposes only minimum safe metadata; plaintext read/export is not introduced. No secret, credential, provider activation, environment edit, deployment configuration, or spend is materialized by this migration.

**Reversibility.** `COSTLY`. Exit requires controlled rotation or rebinding of references and runtime identities, reconciliation/revocation of old versions, and preservation of Core-owned authorization/audit semantics rather than routine plaintext migration through the admin surface.

## UNK-INVENTORY Migration unknown inventory after TASK-0017

`UNK-001..UNK-009` are resolved for their authorized migration slices. All 47 detailed feature relations now use typed CAP references, and no blocking design unknown or `unresolved_ref` remains. Milestone closure remains intentionally blocked only by `milestone.scope_state: OPEN`; no `DOCS_READY` claim is made.

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

Resolving `UNK-005` does **not** imply neighboring relation closure. DATA, IFC, EXT, and CAP detailed relation domains were resolved independently by `TASK-0014..TASK-0017`; `milestone.scope_state` still stays `OPEN`, and `DOCS_READY`, design lock, cutover, implementation readiness, standards PASS, provider activation, promotion, and release remain unclaimed.

## Detailed relation-domain resolutions after TASK-0017

### UNK-006 Expanded detailed feature DATA relations

`UNK-006` is resolved only after every `FTR-007..FTR-053` DATA relation is replaced with typed `DAT-*` references and the complete detailed mapping is semantically audited. No detailed feature requires a genuine DATA N/A: every one reads, creates, mutates, derives, or materially depends on at least one semantic DATA identity.

The accepted `DAT-001..DAT-012` identities remain unchanged. Canonical detailed behavior requires the additional minimal identities `DAT-013..DAT-027`: FeedbackArtifact; ReviewState; MediaSource; ContentCandidate; ValidationDecision history; content release/operational eligibility; effective entitlement; support-managed learner/account product state; authorization grants; typed operating-policy revision; recoverable operational work; consequential privileged audit; MockRun; LearningSession; and CoverageGap. These are boundary-sized semantic identities, not features, screens, tables, DTOs, queue messages, provider states, caches, or one object per workflow.

The mapping intentionally reuses existing learner/evidence/progression identities where they already own the meaning: learner-saved Knowledge/review material remains normal `ContentRevision` plus Knowledge-scoped `MasteryEstimate` and `ReviewState`; Error/Remediation semantics remain content/feedback references rather than new DATA ontologies; content candidate, immutable revision, validation history, and current release/operational state stay distinct because their lifecycle/history semantics differ; entitlement, authorization, and operating policy stay distinct because commercial access, privileged access, and runtime policy are different authorities; operational work does not become provider state; audit does not become authorization or telemetry; and CoverageGap never becomes learner `GapEvaluation`.

`DAT-013` is the catalog `resolved_by_ref` attribution anchor for this non-decision question. The resolution is the complete `FTR-007..FTR-053` typed DATA mapping plus `docs/DATA.md`, not a claim that FeedbackArtifact alone covers detailed DATA. Exact persistence schemas, DTOs, provider activation/state, caches, storage buckets, and retention durations remain outside this resolution.

### UNK-007 Privileged detailed interface relation resolution

Canonical API/runtime ownership proves that no additional semantic IFC is material for the seven privileged surface families. The existing `IFC-001` Core public product API already owns the learner/admin-facing Web-to-Core boundary, including capability-gated privileged reads and mutations; route grouping, role labels, buttons, DTOs, provider calls, and audit records do not create another runtime/trust/ownership boundary.

`FTR-047..FTR-052` therefore use `IFC-001`. Their grouped `FLW-005` crosses `SYS-001` Web to `SYS-002` Core and uses `IFC-001`; authentication, effective capability/scope checks, legal preconditions, Core-owned reads/mutations, durable work/audit markers, and the actual authorized outcome all remain inside that same public product API contract.

`FTR-053` is genuinely interface-N/A at this feature layer. Its only grouped flow, `FLW-006`, is a `SYS-002`-internal reconciliation of a consequential privileged attempt/outcome into durable reconstructable audit. Canonical design demonstrates no separate caller/callee, runtime, trust, or ownership boundary for that reconciliation, so `FLW-006.interface_refs` remains empty instead of fabricating a cross-runtime IFC. The initiating privileged operation that may require audit is already represented by the applicable `FTR-047..FTR-052` relation and `FLW-005` crossing through `IFC-001`.

`IFC-001` is the catalog `resolved_by_ref` attribution anchor for this non-decision question. The resolution is the complete seven-feature mapping plus the affected grouped-flow interface refs, not a claim that every privileged consequence itself crosses Web/Core. `IFC-002` remains exclusively the Core-to-Evaluator bounded capability boundary and is not added to privileged relations without an actual evaluator crossing.

This resolution creates no admin API identity, exact route, OpenAPI/wire schema, DTO, webhook/provider contract, queue/command contract, implementation code, or new runtime. `UNK-008` and `UNK-009` were resolved independently by later bounded relation migrations; milestone scope remains `OPEN`.

### UNK-008 Privileged detailed external-dependency relation resolution

Canonical privileged behavior plus current third-party/runtime ownership support exactly three additional material dependency identities for the seven privileged surface families: `EXT-005` external object/media storage for content operations that handle governed large/retained artifacts, `EXT-006` external payment/subscription observations for entitlement reconciliation, and `EXT-007` external secret custody for explicitly authorized secret-sensitive security operations. The accepted `EXT-001..EXT-004` identities remain unchanged and are not reused merely because a privileged actor authenticates, work can be asynchronous, Core state is hosted externally, or AI exists elsewhere.

The complete detailed mapping is intentionally sparse. `FTR-047` uses `EXT-005`; `FTR-049` uses `EXT-006`; and `FTR-052` uses `EXT-007`. `FTR-048`, `FTR-050`, `FTR-051`, and `FTR-053` use demonstrated dependency N/A because their current semantic operation remains Core-owned and no material external crossing is required: support mutates Core product/account state; provider/work visibility consumes Core-owned operational state while external telemetry remains optional operational evidence; typed policy administration explicitly excludes raw secret editing; and consequential audit reconciliation is Core-internal.

`FLW-005.dependency_refs` remains empty because it is the grouped capability-gated privileged-operation flow shared by features with different conditional external crossings. Promoting the union of object storage, payment, and secret custody onto that grouped flow would incorrectly force every related privileged feature to depend on every provider boundary under the pinned feature-flow subset rule. `FLW-006.dependency_refs` also remains empty because consequential privileged audit reconciliation is wholly Core-internal; an audited operation's earlier provider use does not turn audit reconciliation itself into an external crossing.

`EXT-005` is the catalog `resolved_by_ref` attribution anchor for this non-decision question. Resolution is the complete seven-feature dependency mapping plus the three structural EXT sections and failure/fallback/exit treatment in `docs/INTERFACES.md`, not a claim that object/media storage alone covers privileged dependencies.

The broader canonical provider inventory is deliberately not promoted: no transactional-email EXT is justified by current support semantics; selected external operational telemetry does not become a product dependency merely because operations can display provider/work state; deferred AI generation/model-assisted validation does not justify `EXT-004` reuse; and neither authentication, persistence hosting nor possible asynchronous dispatch is attached without a material feature-specific dependency crossing. Provider lifecycle remains owned by `design/07-third-party-services.md`; catalog identity does not make any route `ACTIVE`, approve data egress, authorize spend, configure credentials, or materialize webhook/OpenAPI/DTO contracts.

### UNK-009 Expanded detailed feature capability/build-buy relation resolution

`UNK-009` is resolved only after every `FTR-007..FTR-053` capability relation is replaced with exact `CAP-*` references and the resulting boundary inventory is audited against canonical runtime, detailed behavior, accepted DATA/IFC/EXT relations, and pinned V1 build/buy semantics.

All 40 detailed learner features use `CAP-001`, `CAP-002`, and `CAP-003`: Core owns their product behavior/policy, protected product access uses the existing identity/session-plus-Core-authorization capability, and their durable authoritative learner/content/evidence/product state remains behind the Core-owned managed-PostgreSQL capability. This is not a mechanical copy of macro mappings: `CAP-004`/`CAP-005` are added only to the eight detailed learner features whose current semantics materially require recoverable bounded evaluator/media work rather than merely sharing broad learner flows: `FTR-030`, `FTR-031`, `FTR-035`, `FTR-038`, `FTR-039`, `FTR-040`, `FTR-045`, and `FTR-046`. The remaining 32 detailed learner features remain usable through Core-owned deterministic/content/capture behavior without making optional AI tutoring, realtime interaction, or a broader flow footprint into a required capability relation.

The seven privileged mappings remain equally sparse and feature-specific:

- `FTR-047` = `CAP-001`, `CAP-002`, `CAP-003`, `CAP-006` for Core-authorized content/review/release behavior plus governed external object/media byte custody;
- `FTR-048` = `CAP-001`, `CAP-002`, `CAP-003` for Core-owned support state only;
- `FTR-049` = `CAP-001`, `CAP-002`, `CAP-003`, `CAP-007` for Core-owned effective entitlement plus external commercial observations;
- `FTR-050` = `CAP-001`, `CAP-002`, `CAP-003` for authorized Core-owned operational visibility; optional telemetry is not a product build/buy dependency;
- `FTR-051` = `CAP-001`, `CAP-002`, `CAP-003` for typed Core-owned operating-policy administration; raw secret editing remains excluded;
- `FTR-052` = `CAP-001`, `CAP-002`, `CAP-003`, `CAP-008` for Core-owned security/access authority plus external secret custody; and
- `FTR-053` = `CAP-001`, `CAP-003` because consequential audit reconciliation is a Core-internal durable product/security responsibility after the initiating authorization boundary and has no separate live identity/session or external provider capability crossing of its own.

Exactly three new capability identities are required. `CAP-006`, `CAP-007`, and `CAP-008` are all `HYBRID` because `SYS-002` retains semantic authority while accepted `EXT-005`, `EXT-006`, and `EXT-007` supply distinct external execution/custody. They are not merged because object/media bytes, commercial observations/effective-entitlement reconciliation, and sensitive credential custody have different responsibility and exit semantics. They are not split further by feature, role, provider, bucket, webhook, credential, model, endpoint, table, queue, permission, or UI surface.

`CAP-006` is the catalog `resolved_by_ref` attribution anchor for this non-decision question. `DEC-005..DEC-007` record only the new material build/buy choices; `DEC-001..DEC-004` and `CAP-001..CAP-005` remain unchanged. Selected/TBD provider lifecycle remains owned by canonical `design/07-third-party-services.md`: this relation resolution does not activate/configure any provider, authorize spend/data egress, choose the unresolved pronunciation route, or create exact contracts, schema, implementation code, deployment topology, or new runtime components.

With all 47 detailed CAP relations typed and `UNK-009` resolved, there is no remaining `UNRESOLVED_RELATION`, blocking unknown, or CAP/build-buy resolution defect in the migration catalog. `milestone.scope_state` deliberately remains `OPEN`, so the sole intentional documentation-closure blocker is `SCOPE_OPEN`; this task does not freeze documentation or claim `DOCS_READY`.
