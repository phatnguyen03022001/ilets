# Delivery

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Created by `TASK-0008` revision 1. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Canonical target snapshot: `phatnguyen03022001/ilets@560a2798710fcf463c2eb32022252a84438224f1`.
>
> Model pin: `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07` (accepted unreleased V1 candidate).

This draft migrates only the five current DELIVERY concerns: environments/configuration, deployment/migration/rollback, backup/restore, compatibility/versioning/platform constraints, and operational ownership. `docs/ARCHITECTURE.md`, `docs/DATA.md`, `docs/INTERFACES.md`, and `docs/QUALITY.md` retain their neighboring migration ownership. This file does not create provider inventory or lifecycle state, `EXT-*`/`CAP-*`/`DEC-*` identities, deployment manifests, workflows, migration files, exact environment files, provider configuration, machine contracts, code, or numerical SLO/RPO/RTO constants.

## DELIVERY-ENVIRONMENTS-CONFIG Environments and configuration

Environment-specific values are deployment/bootstrap inputs only. They may select runtime mode, endpoints, service-identity references, bootstrap secret references, and other static wiring already permitted by canonical design, but they do not become product, learning, evidence, content, progression, or interface authority. Security-critical missing or invalid bootstrap configuration fails closed rather than silently changing behavior.

Secrets remain under the QUALITY/security custody boundary and are not ordinary configuration. Core-owned typed runtime policy remains distinct from secrets and deployment/bootstrap input: only already-authorized adjustable policy may change through its owning Core path. Deployment environment variables or `.env`-style inputs therefore cannot become an alternate policy store, and this migration does not define environment names, secret products, exact values, files, or provider control-plane configuration.

Sources: `design/06-implementation-stack.md` — Secrets, runtime policy, and deployment configuration boundary; `docs/QUALITY.md` — `QUALITY-SECRETS`.

## DELIVERY-DEPLOYMENT-MIGRATION-ROLLBACK Deployment, migration, and rollback

Deployment preserves the three logical runtime responsibility units: Web remains presentation authority, Core remains learner/admin product authority and the only application runtime with authoritative persistence access, and Evaluator remains a bounded internal capability. Units may be split or co-located only while caller/callee direction, trust boundaries, Core-only persistence, explicit machine contracts, least-privilege runtime access, and attributable health/restart/failure behavior remain intact. Topology changes do not move semantic authority.

Once a persistence schema exists, schema migrations are explicit, ordered, and versioned implementation artifacts. Application/schema compatibility covers the selected rollout window instead of assuming lock-step replacement. A deployment or migration failure has a bounded rollback or forward-recovery path, and recovery preserves committed accepted learner work plus any accepted downstream semantic continuation that canonical DATA/QUALITY rules require to remain recoverable.

This migration does not choose DDL, migration commands, deployment products, blue/green/canary/rolling technology, rollout percentages, manifests, release scripts, or orchestration infrastructure.

Sources: `design/06-implementation-stack.md` — Co-location invariant, persistence/migration discipline, version-skew rules, and reliability/recovery/deployment; `docs/ARCHITECTURE.md` — `ARCH-TOPOLOGY`; `docs/DATA.md` — `DATA-MIGRATION`; `docs/QUALITY.md` — `QUALITY-RETRY-RECOVERY`.

## DELIVERY-BACKUP-RESTORE Backup and restore

Backup/restore covers authoritative and retained consequential product state sufficiently to recover a consistent Core-owned authority boundary. Where loss or suspension of a hosting/provider control plane is inside the supported recovery envelope, credible PostgreSQL-compatible restore/exit evidence must not depend exclusively on continued access to that failed control plane. The smallest sufficient export/backup/restore route is enough; this migration does not create a second live database, automatic multi-cloud failover, or a provider-specific backup design.

Restore is a reconciliation boundary, not a blind copy-back. Restored authority is reconciled against current deletion/tombstone policy before normal active use, so backup or derived state cannot resurrect deleted/ineligible data. External work or callbacks dispatched before the restored snapshot are rejected, quarantined, or safely reconciled against restored current authority before they can mutate product state. Historical/audit/provenance integrity is preserved only under its existing owning policy.

Exact backup schedules, storage targets, retention durations, restore commands, RPO, RTO, and provider implementation remain operations/implementation policy and are not invented here.

Sources: `design/06-implementation-stack.md` — Data lifecycle and deletion boundary plus reliability/recovery/deployment; `docs/DATA.md` — `DATA-RETENTION`.

## DELIVERY-COMPATIBILITY-VERSIONING-PLATFORMS Compatibility, versioning, and platforms

Every material runtime boundary has one exact machine-contract authority once separately materialized; generated bindings remain derived from that authority. Deployment verifies compatibility for every version-skew direction it actually supports between producer and consumer revisions. A breaking public or internal contract change cannot assume atomic simultaneous deployment unless the deployment contract explicitly guarantees that condition, and generated bindings do not remove runtime skew.

Persistence rollout likewise preserves the required application/schema compatibility window. Supported skew must retain canonical semantic identity and the INTERFACES distinctions among supplied, not applicable, unresolved/unknown, invalid/required-but-absent, and present values. A wire-valid representation is not sufficient when the semantic change is incompatible.

Platform or co-location choices may change operational placement only; they cannot change Web/Core/Evaluator ownership, Core-only authoritative application persistence, or the declared trust/contract direction. Exact provider products, patch versions, rollout topology, and machine schemas stay with their existing owners and are not materialized by this migration.

Sources: `design/06-implementation-stack.md` — Cross-language contract/evolution and version-skew rules; `docs/INTERFACES.md` — `INTERFACES-EVOLUTION`; `docs/ARCHITECTURE.md` — `ARCH-COMPONENTS` and `ARCH-TOPOLOGY`.

## DELIVERY-OPERATIONAL-OWNERSHIP Operational ownership

Operations owns runtime health/degradation detection, incident intervention, configuration/release provenance, rollback or forward recovery, restore execution, and observation of capacity/backlog/cost pressure at the delivery boundary. External/provider degradation is treated as an operational condition to diagnose and reconcile through the existing dependency/capability boundaries; this section does not select, activate, inventory, or define fallback providers.

Operational telemetry, queue/control-plane state, deployment state, provider status, health signals, and cost/capacity observations are evidence about running delivery. They are not product, learner, evidence, content, progression, interface, or persistence truth. Incidents therefore surface honest pending, degraded, unavailable, or recovery states through existing owners rather than fabricating semantic success, learner weakness, evidence, or product support.

Release/configuration provenance must be sufficient to identify the running build/contract/configuration context relevant to diagnosis and recovery without exposing secrets. Exact alert thresholds, on-call organization, numeric SLO/RPO/RTO targets, capacity constants, provider procedures, and deployment-control-plane implementation remain outside this migration.

Sources: `design/06-implementation-stack.md` — Observability boundary, production-gate concerns, and reliability/recovery/deployment; `design/07-third-party-services.md` — provider failure/degradation and portability boundaries; `docs/QUALITY.md` — `QUALITY-OBSERVABILITY`, `QUALITY-PERFORMANCE-RESOURCES`, and `QUALITY-COST-USAGE`.

## Migration boundary

`docs/DELIVERY.md` and the existing `docs/*` migration set remain **AUTHORITY NONE**. This file preserves `delivery.environments_config`, `delivery.deployment_migration_rollback`, `delivery.backup_restore`, `delivery.compatibility_versioning_platforms`, and `delivery.operational_ownership` at L1 using the sections above.

The neighboring architecture/interface/decision/unknown coverage is now migrated by its owning bounded work, and `UNK-001..UNK-009` are `RESOLVED`; this reconciliation changes none of those requirements, identities, relations, or provider lifecycle. `milestone.scope_state` remains `OPEN` pending a post-reconciliation semantic closure/freeze audit. This DELIVERY reconciliation creates or changes no provider inventory, `EXT-*`/`CAP-*`/`DEC-*` state, deployment implementation, contract materialization, workflow, migration artifact, standards PASS/level, `DOCS_READY`, design lock, implementation-readiness claim, promotion, or release claim.
