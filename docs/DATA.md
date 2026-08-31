# Data

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> `TASK-0005` revision 1 migration artifact. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Canonical target snapshot: `phatnguyen03022001/ilets@d104634d78a6463092998e4a1311adb837e708bd`.
>
> Model pin: `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07` (accepted unreleased V1 candidate).

This draft migrates only the material DATA semantics required by the accepted target-to-evidence-to-next-action behavior slice. `docs/catalog/project.json` owns the `DAT-*` inventory, kinds, ownership, and feature relations; this file explains the migrated DATA semantics. It does not define database tables/columns, API payloads, provider identities, exact retention periods, cache authority, or a new learning/evidence model.

## DATA-ENTITIES Material data identities and ownership

All `DAT-*` identities in this bounded slice are owned by `SYS-002` Core API. Canonical architecture makes Core the only application runtime with authoritative product persistence. Web may collect or display data and Evaluator may return bounded candidate results/provenance, but neither becomes the authoritative owner of learner, content, evidence, or progression state.

### DAT-001 TargetProfile and target-context revision

`TargetProfile` carries the learner-declared target conditions that materially constrain planning: known variant, real overall/per-skill Band constraints where supplied, date and other applicable target conditions, plus unresolved target inputs that remain unresolved. Core persists the target/target-context revision. An explicit learner target change creates new current target state; planning must not manufacture missing variant/Band values or rewrite the target implicitly.

This identity is target state, not learner ability, Assessment evidence, progression state, or a plan.

Sources: `spec/01-LEARNER-MODEL.md` — Target profile; `design/04-application-flows.md` — Planner Stage 1 and Flow A.

### DAT-002 ContentRevision

A `ContentRevision` is the exact immutable semantic revision of established content used for an assignment or historical learner attempt. Material semantic changes create another revision; later activation, retirement, or revalidation does not rewrite what a learner previously received. An Attempt/evidence lineage must remain able to resolve the exact revision actually presented.

This identity does not make content valid, released, independent, or evidence-eligible by itself, and it does not freeze a storage schema.

Sources: `spec/10-CONTENT-MODEL.md` — Content identity, revision, and lineage; `design/06-implementation-stack.md` — Content authoring and released-materialization boundary.

### DAT-003 Attempt history

An `Attempt` is learner-instance history against an exact Practice/Assessment content revision. It records what happened under the material response, timing, delivery/capture, scaffold, assistance, exposure, and evaluation conditions that are established for the attempt.

An Attempt is not automatically an Observation or EvidenceFact, and completion, retry, Skip/abandonment, time spent, or a favorable result does not by itself prove capability or Band.

Sources: `spec/10-CONTENT-MODEL.md` — Attempt; `spec/08-ASSESSMENT.md` — Core objects; `design/04-application-flows.md` — Flow C.

### DAT-004 Observation history

An `Observation` is the normalized measurement result before claim-scoped evidence admission. It preserves the measured result plus the material task/context, conditions, assistance/scaffolding, exposure, scorer/evaluator provenance, uncertainty, and attempt association needed to interpret what was actually observed.

Observations remain historical measurement records. A later policy or current-evidence interpretation may change what an Observation supports now without rewriting the original measurement.

Sources: `spec/08-ASSESSMENT.md` — Observation and evidence eligibility; `spec/10-CONTENT-MODEL.md` — Observation.

### DAT-005 EvidenceFact history

An `EvidenceFact` is a claim-scoped admission of an Observation under an explicit Assessment eligibility/policy context. One Observation may yield multiple compatible EvidenceFacts or none. Historical EvidenceFacts remain historical facts; later staleness changes current support, not the fact that the evidence was admitted at that time.

EvidenceFact is not `MasteryEstimate`, `ReadinessEvaluation`, `GapEvaluation`, `BandCertificationState`, or a guaranteed external IELTS outcome.

Sources: `spec/08-ASSESSMENT.md` — EvidenceFact; `spec/10-CONTENT-MODEL.md` — EvidenceFact.

### DAT-006 ReadinessEvaluation

`ReadinessEvaluation` is a material current Assessment interpretation for one scoped claim. Its canonical outcomes preserve `INSUFFICIENT_EVIDENCE`, `CONFLICTING_EVIDENCE`, `STALE_EVIDENCE`, `NOT_YET_SUPPORTED`, and `SUPPORTED` rather than collapsing them into one score or learner weakness.

This catalog treats the evaluation as `MATERIAL_EPHEMERAL`: the canonical owners require the interpretation and its policy/provenance to be reconstructable, but do not require this migration to invent a dedicated durable storage record.

Sources: `spec/08-ASSESSMENT.md` — ReadinessEvaluation; `spec/09-PROGRESSION.md` — Claim interpretation boundary.

### DAT-007 MasteryEstimate

`MasteryEstimate` is Progression's current uncertainty-aware interpretation of one scoped Skill Leaf or Knowledge Object. Its `unknown`, `learning`, and `currently_supported` projection does not erase the underlying Assessment reason: stale, conflicting, insufficient, pending/unusable, or below-requirement conditions remain distinguishable through their evidence/gap references.

It is current learner interpretation, not immutable evidence history and not a whole-skill Band certification.

Sources: `spec/09-PROGRESSION.md` — MasteryEstimate.

### DAT-008 BandCertificationState

`BandCertificationState` is the current certification state for a specific `(skill, band)` claim. A current claim may be `not_started`, `in_progress`, or `certified` under the canonical Assessment/Progression rules. Loss of current support can move the current state back to `in_progress` without deleting or rewriting historical certification.

This state is not an official IELTS result and cannot be created from completion count, planner state, AI output, or one narrow EvidenceFact.

Sources: `spec/09-PROGRESSION.md` — BandCertificationState and Band advancement.

### DAT-009 Historical attainment records

Historical attainment preserves point-in-time internal certification history and imported/external attainment with the evidence/policy/source/date/variant provenance needed by their owning rules. Historical attainment remains separate from current `BandCertificationState` and current readiness/mastery interpretation.

Staleness, conflict, later regression, or re-certification append or relate current/history state rather than rewriting prior attainment. An external result may inform a current claim only when normal Assessment scope and recency rules admit it.

Sources: `spec/08-ASSESSMENT.md` — historical attainment / external result; `spec/09-PROGRESSION.md` — Certification history.

### DAT-010 GapEvaluation

`GapEvaluation` is Progression's material current classification of the demonstrated or unresolved condition that matters for next action, including evidence, conflict, staleness, prerequisite, ability, scaffold, transfer, fluency, or exam-condition cases defined by the canonical owner.

Missing evidence, provider/evaluator failure, or product coverage inability cannot be converted into `ABILITY_GAP`. This catalog treats `GapEvaluation` as `MATERIAL_EPHEMERAL`; the semantic result is material to planning, while exact durable representation is not asserted here.

Sources: `spec/09-PROGRESSION.md` — Assessment state → learner gap and GapEvaluation.

### DAT-011 ActionIntent

`ActionIntent` is the semantic next-action objective emitted by Progression after the applicable target, Assessment interpretation, and prerequisite state are considered. It is not a concrete activity, ranking score, DailyPlan, or assignment.

This catalog treats `ActionIntent` as `MATERIAL_EPHEMERAL`: Planner consumes it as a material boundary value, while persistence shape is left to implementation.

Sources: `spec/09-PROGRESSION.md` — ActionIntent and Next-action explanation; `design/04-application-flows.md` — Planner Stage 3.

### DAT-012 DailyPlan snapshot

A `DailyPlan` is the current recommendation snapshot produced after target/support resolution, evidence/progression interpretation, hard eligibility, candidate construction, ranking, and coherent load composition. Its provenance can include target-context revision, learner/progression state references, due-review state, policy/configuration identity, eligible content revisions, and unresolved conditions.

A plan is not assignment authority. Current mutable hard eligibility is rechecked before actual assignment/exposure, so a stale saved plan cannot force execution. The canonical owners do not require this migration to choose a durable plan schema; the catalog therefore records it as `MATERIAL_EPHEMERAL`.

Sources: `design/04-application-flows.md` — Planner decision contract, Stage 7, and Flow B.

## DATA-LIFECYCLE Lifecycle and persistence

- Core API is the authoritative application owner for the persisted target/content/attempt/evidence/progression histories in this slice. Web/Next.js presentation state and Evaluator candidate output are non-authoritative until accepted through Core.
- `DAT-001` through `DAT-005` and `DAT-007` through `DAT-009` are persistent semantic identities because canonical truth requires durable current/history or exact-revision reconstruction. Their catalog identities do not imply one-table-per-DAT or any table/column layout.
- `DAT-006`, `DAT-010`, `DAT-011`, and `DAT-012` are material boundary/snapshot values whose exact persistence is not required by current canonical truth. They may be reconstructed or persisted by implementation only without changing their semantic owner.
- Accepted asynchronous evaluator/capability results become product truth only after Core validates current work/execution association and current legal state. Required downstream Observation/EvidenceFact/Progression continuation is committed or made durably recoverable before semantic completion.
- Historical Attempt, Observation, EvidenceFact, content-revision, and attainment references stay resolvable across later current-state changes. Current projections never replace their upstream history.

Sources: `design/04-application-flows.md` — authoritative mutation/asynchronous work; `design/06-implementation-stack.md` — Go → authoritative store and consistency classes; `spec/08-ASSESSMENT.md`; `spec/09-PROGRESSION.md`; `spec/10-CONTENT-MODEL.md`.

## DATA-RETENTION Retention and deletion

Current canonical truth defines deletion/retention boundaries but does not define a universal retention duration. This migration therefore records the boundary without inventing days, legal periods, or provider-specific policy.

- One authoritative product retention/deletion decision governs the affected data scope and is propagated to authoritative state and applicable derived/external copies under their owning policies.
- Current deletion/tombstone eligibility participates in asynchronous dispatch/result reconciliation. A late evaluator/provider result cannot resurrect deleted learner data, recreate active artifact references, or create new Observation/EvidenceFact state from data no longer eligible for use.
- Derived stores and restore procedures cannot silently resurrect deleted state into active product use; restored authority is reconciled against current deletion/tombstone truth before normal use.
- Historical semantic/audit provenance is retained only when an explicit applicable product/security/data policy permits it; accidental persistence is not authority.
- Entitlement loss or product downgrade is not itself a deletion instruction.

No `N/A` is claimed for unresolved retention detail, and this section creates no provider identity or exact retention schedule.

Source: `design/06-implementation-stack.md` — Data lifecycle and deletion boundary.

## DATA-CONSISTENCY Consistency and transaction boundaries

- A product success acknowledgement follows the authoritative commit it claims succeeded.
- Where duplicate execution or stale writes matter, decisive idempotency admission/replay and expected-revision checks are coupled to the protected authoritative mutation or an equivalent serialization invariant; preflight reads alone are insufficient.
- Required async work identity/recoverable continuation is established with or derivable from committed authoritative state before later work relies on it.
- Duplicate/superseded/late evaluator results are fenced against current authoritative work/deletion/content eligibility; replay cannot create duplicate Attempt, Observation, EvidenceFact, or learner-state transitions.
- Assignment uses the exact current eligible content revision and rechecks mutable hard gates at the decisive reservation/assignment boundary. Historical Attempt/evidence lineage continues to reference the revision actually delivered.
- Network/provider calls are not one database transaction. The exact SQL/table/index/locking design remains implementation-owned.

Sources: `design/04-application-flows.md` — Authoritative durable mutation and asynchronous work; `design/06-implementation-stack.md` — Persistence/migration discipline.

## DATA-MIGRATION Migration and backfill

Current canonical implementation truth requires explicit, ordered, versioned schema migrations once a schema exists and application/schema compatibility across the selected rollout window. The database schema and migration files remain derived implementation, not DATA semantic authority.

For this migrated semantic slice:

- migration/deploy recovery preserves already committed accepted learner work;
- a migration, backfill, recomputation, or materializer may not manufacture missing evidence, broaden an Observation/EvidenceFact inference, rewrite historical Attempts/attainment, or silently turn stale/conflicting/unknown state into supported/weak state;
- derived current projections may be recomputed only from the authoritative histories and applicable policy/provenance needed by their owning semantics;
- canonical registry materialization preserves source-owner identity/fingerprint/revision and does not create new semantic objects by scanning incidental tokens;
- content materialization preserves exact ContentRevision identity/lineage and historical references.

This is the complete L1 migration boundary supported by current truth; it deliberately does not define DDL, table names, backfill jobs, rollout commands, or a universal migration algorithm.

Sources: `design/06-implementation-stack.md` — Persistence/migration discipline, Canonical registry materialization, and Content authoring/released-materialization boundary.

## DATA-LINEAGE Provenance, lineage, and data quality

The material lineage for this slice remains reconstructable without collapsing semantic layers:

```text
TargetProfile / target-context revision
        ↓ planning context
ContentRevision → Attempt
                    ↓ measurement
                Observation
                    ↓ claim-scoped admission + policy
                EvidenceFact
                    ↓ Assessment
             ReadinessEvaluation
                    ↓ Progression
 MasteryEstimate / BandCertificationState
 GapEvaluation → ActionIntent
                    ↓ Planner
               DailyPlan
```

Historical attainment remains alongside—not inside—current certification/readiness state.

Required provenance stays attached at the owning layer: exact content revision and attempt conditions; evaluator/scorer identity/uncertainty where material; Observation context/conditions; EvidenceFact claim scope and policy version; target-context revision; current progression evidence/policy references; and certification/attainment history provenance.

Missing, stale, conflicting, pending/unusable, below-requirement, supported, provider/runtime-failure, and product-coverage conditions are not interchangeable data-quality labels. Missing evidence is not negative evidence; technical/product failure is not learner evidence; a downstream projection never rewrites its upstream history.

Sources: `spec/01-LEARNER-MODEL.md`; `spec/08-ASSESSMENT.md`; `spec/09-PROGRESSION.md`; `spec/10-CONTENT-MODEL.md`; `design/04-application-flows.md`.

## Migration boundary

`docs/DATA.md`, `docs/PRODUCT.md`, `docs/BEHAVIOR.md`, `docs/ARCHITECTURE.md`, `docs/DECISIONS.md`, and `docs/catalog/project.json` remain migration artifacts with **authority NONE**. This task does not cut over canonical ownership, freeze milestone scope, resolve `UNK-002`/`UNK-003`/`UNK-004`, materialize interface/dependency/capability/decision inventories, claim `DOCS_READY`, or declare implementation/assurance/promotion/release readiness.
