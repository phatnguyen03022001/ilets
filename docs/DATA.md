# Data

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> `TASK-0005` revision 1 established the base DATA migration; `TASK-0014` revision 1 extends only the detailed feature DATA relations. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Canonical detailed-feature authority snapshot: `phatnguyen03022001/ilets@a3f093c70d1c6e6d732c6714ac594cd53fd69eda`.
>
> Model pin: `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07` (accepted unreleased V1 candidate).

This draft migrates only the material DATA semantics required by the accepted target-to-evidence-to-next-action behavior slice and its accepted detailed learner/privileged feature surface. `docs/catalog/project.json` owns the `DAT-*` inventory, kinds, ownership, and feature relations; this file explains the migrated DATA semantics. It does not define database tables/columns, API payloads, provider identities, exact retention periods, cache authority, or a new learning/evidence model.

## DATA-ENTITIES Material data identities and ownership

All `DAT-*` identities in this bounded slice are owned by `SYS-002` Core API. Canonical architecture makes Core the only application runtime with authoritative product persistence. Web may collect or display data and Evaluator may return bounded candidate results/provenance, but neither becomes the authoritative owner of learner, content, evidence, progression, entitlement, authorization, policy, work, or privileged-audit state.

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

### DAT-013 FeedbackArtifact

A `FeedbackArtifact` is the material learner-facing guidance derived from an Attempt/Observation under the canonical feedback-focus policy. It can preserve observed performance, selected feedback targets, deferred observations, error/remediation references, recommended action intent, provenance, and uncertainty without becoming EvidenceFact, learner weakness, or mastery authority.

This catalog treats it as `MATERIAL_EPHEMERAL`: the feedback object is a real semantic boundary used by review/revision/re-record behavior, while current canonical truth does not require a universal durable feedback-history contract.

Sources: `spec/10-CONTENT-MODEL.md` — FeedbackArtifact; `spec/07-PRACTICE.md` — Feedback focus and noise control.

### DAT-014 ReviewState

`ReviewState` is the current semantic due-review/reassessment condition for a reviewable Skill/Knowledge target or prior learning/evidence need. It preserves why something is due and the applicable review/re-evidence intent without inventing one universal SRS interval or reducing all review to flashcards.

The state is `MATERIAL_EPHEMERAL`: it is a current derived planning/review boundary that may be reconstructed from learner history plus applicable policy. Learner-saved words, phrases, examples, and review cards that become reusable product content remain `ContentRevision` data rather than a second Knowledge ontology.

Sources: `spec/09-PROGRESSION.md` — Review state; `spec/04-KNOWLEDGE.md` — learner-saved material; `design/04-application-flows.md` — Flow G.

### DAT-015 MediaSource

`MediaSource` is Core-owned persistent product state for an eligible media source reference and the material accepted metadata/provenance, current playability/embeddability, rights state, transcript state/reference, and source identity needed to govern media learning. It remains separate from the reusable `ContentRevision` that may reference it.

Source availability or rights/transcript eligibility may change without rewriting historical Attempts or turning source loss into learner regression. External bytes/provider state do not become product authority merely because a source is referenced.

Sources: `design/03-media-youtube.md` — MediaSource product contract and source removal/failure; `design/06-implementation-stack.md` — media-source product state.

### DAT-016 ContentCandidate

`ContentCandidate` is persistent mutable authoring/generation input state before an established immutable `ContentRevision` exists. It preserves candidate identity/provenance and review/processing context needed by the content workflow without pretending a draft is released content or canonical learning truth.

Acceptance/materialization creates or binds an established `ContentRevision`; it does not mutate the candidate identity into historical learner content. Exact authoring package, repository shard, queue, or storage schema remains outside this DATA identity.

Sources: `spec/10-CONTENT-MODEL.md` — content revision distinction and provenance; `design/04-application-flows.md` — content supply; `design/06-implementation-stack.md` — content authoring and released-materialization boundary.

### DAT-017 ValidationDecision history

A `ValidationDecision` is a persistent auditable decision about a specific ContentRevision under a named validation-policy/intended-use scope. Revalidation of unchanged semantic content creates another historical decision; a later decision does not rewrite an earlier result/reason/provenance record.

Validation history is distinct from the immutable ContentRevision and from current release/operational eligibility. A validator or generator cannot make itself authoritative by self-reporting confidence.

Sources: `spec/10-CONTENT-MODEL.md` — Validation decision semantics; `design/04-application-flows.md` — revalidation and content operations.

### DAT-018 Content release and operational eligibility state

This persistent Core-owned state represents the current release/use/assignment and operational-safety eligibility that can activate, quarantine, retire, block, or otherwise govern an established content revision for current product use. It does not rewrite the revision or collapse validation history, release eligibility, and operational safety into one historical status.

Assignment and planning must consume current eligible state; historical Attempts keep the exact revision and then-valid provenance even when later current eligibility changes. Provider availability, cache state, or implementation table status is not this identity.

Sources: `design/04-application-flows.md` — content report/quarantine/revalidation/retirement and assignment gates; `design/06-implementation-stack.md` — assignment and deletion/content-eligibility fencing.

### DAT-019 Effective entitlement state

`Effective entitlement state` is the persistent Core-owned commercial-access truth after accepted entitlement facts and any pending reconciliation are interpreted. It may distinguish current effective access from pending/expired/restored commercial state where canonical behavior requires that distinction.

Entitlement never grants an authorization role/capability, changes learner evidence or Band semantics, makes unsupported product capability valid, or deletes retained learner history when access expires.

Sources: `design/04-application-flows.md` — entitlement reconciliation; `design/08-coverage-and-support.md` — learner evidence, interaction availability, entitlement, and authorization separation.

### DAT-020 Learner/account support-managed product state

This persistent identity is the bounded non-evidence learner/account product state that authorized support operations may legitimately correct or reconcile. It exists because canonical privileged behavior permits support mutations but explicitly forbids fabricating learner activity/evidence or using support access as blanket protected-data authority.

`TargetProfile`, Attempt/Observation/EvidenceFact, entitlement, authorization grants, and policy revisions remain their own DATA identities; this support-managed state does not absorb them. Exact account fields, identity-provider metadata, or support-case schema are not defined here.

Sources: `design/04-application-flows.md` — support operations and privileged capability boundaries; `docs/BEHAVIOR.md#FTR-048` — migrated behavior only, authority NONE.

### DAT-021 Authorization grant state

`Authorization grant state` is the persistent Core-owned current role/capability/scope grant truth consumed by the authoritative privileged-operation check. It stays separate from commercial entitlement, identity authentication, learning/evidence authority, content validation authority, and ordinary policy configuration.

Grant/revoke and security-sensitive changes are consequential mutations with their own audit requirements. This identity does not define a permission-table schema or make role bundles the authorization authority by themselves.

Sources: `design/04-application-flows.md` — RBAC / capability model; `design/06-implementation-stack.md` — Admin/privileged actor → Core.

### DAT-022 Typed operating policy revision

A `Typed operating policy revision` is persistent Core-owned approved runtime policy whose value can legitimately change product admission/operation without redeploy while staying inside already-authorized canonical semantics. Consequential revisions remain reconstructable where required.

This identity excludes secrets and bootstrap/deployment environment, does not provide a raw environment-variable editor, and cannot redefine Band/Skill/Knowledge/evidence/content/progression semantics or provider calibration eligibility.

Sources: `design/06-implementation-stack.md` — Secrets, runtime policy, and deployment configuration boundary; `design/04-application-flows.md` — approved policy administration.

### DAT-023 Operational work state

`Operational work state` is persistent authoritative Core state for a logical accepted asynchronous/processing/reconciliation operation and the material current execution-attempt/fencing/recovery status needed to prevent duplicate, superseded, stranded, or resurrected outcomes. External/provider observations can be associated as non-authoritative provenance without becoming provider-owned product state.

Accepted work remains recoverable across dispatch ambiguity and worker loss; required semantic continuation is committed or reconstructable before completion is treated as final. This identity is not a queue message, provider job schema, cache, or telemetry record.

Sources: `design/04-application-flows.md` — authoritative asynchronous work; `design/06-implementation-stack.md` — AUTHORITATIVE_ASYNC_STATE and persistence discipline.

### DAT-024 Consequential privileged audit history

A consequential privileged audit record is persistent reconstructable history for a protected operation, including the material actor identity/class, effective role/capability context, target resource, prior/result state or policy revision where required, reason, and outcome/provenance. It is evidence that a privileged action occurred, not authorization to perform the action.

Privileged audit remains distinct from operational logs/metrics/traces and may be a transactional precondition for the consequential mutation it protects. Exact retention duration or log backend is not defined here.

Sources: `design/04-application-flows.md` — privileged operations and audit; `design/06-implementation-stack.md` — Observability boundary.

### DAT-025 MockRun

`MockRun` is persistent learner-run composition/history that binds one integrated mock execution to the exact eligible component ContentRevisions/Attempts, variant and material delivery/independence conditions required to reconstruct what was actually delivered. It is not a new immutable mega-content ontology and does not make one mock self-certifying.

A Speaking/full-test run can contribute only the EvidenceFacts and readiness scope independently admitted by Assessment. Later revalidation, content retirement, or learner-state change does not rewrite the run's historical component lineage.

Sources: `design/04-application-flows.md` — full IELTS mock flow; `spec/10-CONTENT-MODEL.md` — section and full-mock content assembly.

### DAT-026 LearningSession

`LearningSession` is persistent Core-owned learner interaction/session state where the product needs a durable logical session boundary distinct from individual Attempts, including optional realtime Speaking interaction. Session lifecycle can preserve accepted turns/work association and terminal completion/abandonment semantics without making a conversation transcript or provider session the learning authority.

A session may contain or coordinate Attempts, but those Attempts retain their own exact content/capture/evidence lineage. This identity defines no WebSocket/provider protocol or session table.

Sources: `design/04-application-flows.md` — legal lifecycle state and realtime Speaking flow; `design/06-implementation-stack.md` — durable session/Attempt state.

### DAT-027 CoverageGap

A `CoverageGap` is the material current product inability or unresolved product-support condition for a scoped target/condition. It is categorically separate from learner `GapEvaluation`: product inability, missing evaluator/content/runtime support, or calibration/reliability failure cannot become learner weakness.

This catalog treats CoverageGap as `MATERIAL_EPHEMERAL`: the current gap plus source/version provenance is material to planner/learner explanation, while exact durable representation and support-declaration storage remain outside this bounded migration.

Sources: `design/08-coverage-and-support.md` — CoverageGap and product-support status; `design/00-learning-experience.md` — learner-visible CoverageGap separation.

## DATA-LIFECYCLE Lifecycle and persistence

- Core API is the authoritative application owner for the target/content/attempt/evidence/progression, media/content operational, entitlement/authorization/policy, work/session/mock, support-managed account, and privileged-audit semantics in this slice. Web/Next.js presentation state and Evaluator/provider candidate output are non-authoritative until accepted through Core.
- `DAT-001` through `DAT-005`, `DAT-007` through `DAT-009`, and `DAT-015` through `DAT-026` are persistent semantic identities because canonical truth requires durable current/history, recoverability, or exact-lineage reconstruction. Their catalog identities do not imply one-table-per-DAT or any table/column layout.
- `DAT-006`, `DAT-010` through `DAT-014`, and `DAT-027` are material current boundary/snapshot values whose exact durable representation is not required by current canonical truth. They may be reconstructed or persisted by implementation without changing their semantic owner.
- A mutable `ContentCandidate` does not become an immutable ContentRevision by in-place reinterpretation. Validation decisions append historical decision records, while current content release/operational eligibility may change without rewriting either the revision or prior validation history.
- Accepted asynchronous evaluator/content/entitlement work becomes product truth only after Core validates current logical-work/execution association and current legal state. Required downstream Observation/EvidenceFact/Progression or other semantic continuation is committed or made durably recoverable before semantic completion.
- Consequential authorization/policy/content/entitlement/support/security mutations preserve the required durable audit boundary; operational telemetry cannot substitute for that audit history.
- Historical Attempt, Observation, EvidenceFact, content-revision, validation, mock/session, and attainment references stay resolvable across later current-state changes. Current projections never replace their upstream history.

Sources: `design/04-application-flows.md` — authoritative mutation/asynchronous work and privileged flows; `design/06-implementation-stack.md` — authoritative store, consistency classes, runtime policy and observability boundaries; `spec/08-ASSESSMENT.md`; `spec/09-PROGRESSION.md`; `spec/10-CONTENT-MODEL.md`.

## DATA-RETENTION Retention and deletion

Current canonical truth defines deletion/retention boundaries but does not define a universal retention duration. This migration therefore records the boundary without inventing days, legal periods, or provider-specific policy.

- One authoritative product retention/deletion decision governs the affected data scope and is propagated to authoritative state and applicable derived/external copies under their owning policies.
- Current deletion/tombstone eligibility participates in asynchronous dispatch/result reconciliation. A late evaluator/provider result cannot resurrect deleted learner data, recreate active artifact references, or create new Observation/EvidenceFact state from data no longer eligible for use.
- Derived stores and restore procedures cannot silently resurrect deleted state into active product use; restored authority is reconciled against current deletion/tombstone truth before normal use.
- Historical semantic/audit provenance is retained only when an explicit applicable product/security/data policy permits it; accidental persistence is not authority. This includes audit/validation/session/mock history only to the extent required by the applicable policy; no duration is implied.
- Entitlement loss or product downgrade is not itself a deletion instruction. Source removal or content retirement likewise changes current eligibility without erasing otherwise valid historical learner lineage.

No `N/A` is claimed for unresolved retention detail, and this section creates no provider identity or exact retention schedule.

Source: `design/06-implementation-stack.md` — Data lifecycle and deletion boundary; `design/03-media-youtube.md` — source removal/failure.

## DATA-CONSISTENCY Consistency and transaction boundaries

- A product success acknowledgement follows the authoritative commit it claims succeeded.
- Where duplicate execution or stale writes matter, decisive idempotency admission/replay and expected-revision checks are coupled to the protected authoritative mutation or an equivalent serialization invariant; preflight reads alone are insufficient.
- Required async work identity/recoverable continuation is established with or derivable from committed authoritative state before later work relies on it.
- Duplicate/superseded/late evaluator/provider results are fenced against current authoritative work/deletion/content eligibility; replay cannot create duplicate Attempt, Observation, EvidenceFact, learner-state transitions, entitlement outcomes, or released content state.
- Assignment uses the exact current eligible content revision and rechecks mutable hard gates at the decisive reservation/assignment boundary. Historical Attempt/evidence lineage continues to reference the revision actually delivered.
- Current effective entitlement, authorization grants, approved policy, content operational eligibility, and support/deletion fences are checked by their owning mutation/admission boundaries rather than cached/preflight state being treated as authority.
- Consequential privileged mutations couple required audit materialization to the protected outcome when canonical policy requires the audit as a transactional precondition.
- Network/provider calls are not one database transaction. The exact SQL/table/index/locking design remains implementation-owned.

Sources: `design/04-application-flows.md` — Authoritative durable mutation and asynchronous work; `design/06-implementation-stack.md` — Persistence/migration discipline and Admin/privileged actor boundary.

## DATA-MIGRATION Migration and backfill

Current canonical implementation truth requires explicit, ordered, versioned schema migrations once a schema exists and application/schema compatibility across the selected rollout window. The database schema and migration files remain derived implementation, not DATA semantic authority.

For this migrated semantic slice:

- migration/deploy recovery preserves already committed accepted learner and privileged work;
- a migration, backfill, recomputation, or materializer may not manufacture missing evidence, broaden an Observation/EvidenceFact inference, rewrite historical Attempts/attainment/validation/audit/mock history, or silently turn stale/conflicting/unknown state into supported/weak state;
- derived current projections such as ReadinessEvaluation, GapEvaluation, ActionIntent, ReviewState, DailyPlan, or CoverageGap may be recomputed only from authoritative histories and applicable current policy/provenance needed by their owning semantics;
- canonical registry materialization preserves source-owner identity/fingerprint/revision and does not create new semantic objects by scanning incidental tokens;
- content materialization preserves exact ContentRevision identity/lineage; candidate, validation, and release/operational identities remain distinct across migration.

This is the complete L1 migration boundary supported by current truth; it deliberately does not define DDL, table names, backfill jobs, rollout commands, or a universal migration algorithm.

Sources: `design/06-implementation-stack.md` — Persistence/migration discipline, Canonical registry materialization, and Content authoring/released-materialization boundary.

## DATA-LINEAGE Provenance, lineage, and data quality

The material learner/content lineage for this slice remains reconstructable without collapsing semantic layers:

```text
TargetProfile / target-context revision
        ↓ planning context
MediaSource → ContentCandidate → ContentRevision
                               ↘ ValidationDecision history
                                 ↓ current release/use eligibility
ContentRevision → Attempt ← LearningSession / MockRun
                    ↓ measurement
                Observation
                    ↓ claim-scoped admission + policy
                EvidenceFact
                    ↓ Assessment
             ReadinessEvaluation
                    ↓ Progression
 MasteryEstimate / BandCertificationState
 GapEvaluation → ActionIntent
       ↓             ↓
 ReviewState      DailyPlan
       ↘             ↙
      FeedbackArtifact / eligible next work
```

`CoverageGap` remains a separate product-support lineage from learner `GapEvaluation`. Historical attainment remains alongside—not inside—current certification/readiness state.

Privileged operational lineage remains separately reconstructable:

```text
accepted external/commercial facts → Effective entitlement state
current authorization grants + typed operating policy
        ↓ capability-gated operation
ContentCandidate / ContentRevision / ValidationDecision / content operational state
OperationalWork state where asynchronous/reconciled
        ↓
Consequential privileged audit history
```

Required provenance stays attached at the owning layer: exact content revision and attempt conditions; media/candidate/validation/release source identity; evaluator/scorer identity/uncertainty where material; Observation context/conditions; EvidenceFact claim scope and policy version; target-context revision; current progression evidence/policy references; entitlement/auth/policy/work provenance; mock/session component lineage; and certification/attainment/audit history provenance.

Missing, stale, conflicting, pending/unusable, below-requirement, supported, provider/runtime-failure, product-coverage, entitlement, authorization, content-eligibility, and privileged-work conditions are not interchangeable data-quality labels. Missing evidence is not negative evidence; technical/product failure is not learner evidence; a downstream projection never rewrites its upstream history.

Sources: `spec/01-LEARNER-MODEL.md`; `spec/08-ASSESSMENT.md`; `spec/09-PROGRESSION.md`; `spec/10-CONTENT-MODEL.md`; `design/03-media-youtube.md`; `design/04-application-flows.md`; `design/06-implementation-stack.md`; `design/08-coverage-and-support.md`.

## Migration boundary

`docs/DATA.md`, `docs/PRODUCT.md`, `docs/BEHAVIOR.md`, `docs/ARCHITECTURE.md`, `docs/DECISIONS.md`, and `docs/catalog/project.json` remain migration artifacts with **authority NONE**. `TASK-0014` resolved `UNK-006` after complete typed DATA mapping for `FTR-007..FTR-053`. Later bounded migrations resolved `UNK-007`, `UNK-008`, and `UNK-009` without changing DATA semantics; `UNK-001..UNK-009` are now `RESOLVED`.

Current mutable migration milestone scope state is owned exclusively by `docs/catalog/project.json`; this Markdown intentionally does not restate `OPEN`, `FROZEN`, or `SCOPE_OPEN`. `DOCS_READY` is derived by the pinned documentation model rather than owned or stored here. This file does not cut over canonical ownership, redefine IFC/EXT/CAP mappings, materialize exact contracts/providers/storage schema, or declare design lock, implementation/assurance/promotion/release readiness.