STATUS: CANONICAL
OWNS: end-to-end product/system flows across web, core API, evaluator, learner state, target route, media, content demand/supply/assignment and content-incident recovery, actor classification, privileged capability semantics and default RBAC capability bundles, learner-entitlement lifecycle versus operational authorization, runtime execution/trust/failure patterns, async result delivery, planner-stage separation/load composition, target-trajectory projection, intervention-effectiveness planning interpretation, hard eligibility, and legal runtime lifecycle semantics
DEPENDS_ON: ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 01-skill-features.md, 02-practice-catalog.md, 03-media-youtube.md
DOES_NOT_OWN: API field schemas, learning/mastery truth, content semantic identity/quality truth, product coverage declaration, identity-provider implementation or credential mechanics, exact persistence topology, provider selection, framework internals, deployment technology selection, or learner-facing UX defaults

# Application Flows

## Purpose

Define how major product interactions traverse runtime boundaries without assigning learning authority to transport, UI, evaluator, ranker, content generator, deployment topology, or third-party infrastructure.

# Logical runtime units

```text
apps/web                 TypeScript learner/admin experience
services/core-api        Go public API + deterministic orchestration
services/evaluator       Python bounded AI/audio/text evaluation
```

Normal product path:

```text
Web
 ↓
Core API
 ├─ durable learner/target/session/attempt state
 ├─ deterministic scoring/policy execution
 └─ internal evaluator request
          ↓
      Evaluator
          ↓
   Observation candidates + provenance
          ↓
      Core API
          ↓
Assessment / Progression / Planner
          ↓
Web
```

The evaluator does not certify Band, mutate learner progression, declare product coverage, activate content, or choose the final next action.

These flows obey the logical ownership and co-location invariant owned by `06-implementation-stack.md`. Deployment topology cannot change the legal caller/callee direction or semantic authority defined here.

# Runtime execution patterns

These patterns define legal end-to-end execution behavior. Exact HTTP fields/status mappings belong to `05-api.md` and future machine contracts; exact persistence technology belongs to implementation stack/deployment decisions.

For protected operations, safe bounded transport framing and credential extraction necessarily begin before identity is known. Protected-resource-sensitive processing then follows the access boundary. A public anonymous operation may omit authentication/authorization steps where its contract permits it.

## A. Authoritative durable mutation

```text
bounded transport framing / credential extraction
        ↓
authentication where required
        ↓
authorization / capability / access filtering where required
        ↓
structural contract validation
        ↓
resource-sensitive semantic preconditions
        ↓
optional preflight idempotency / stale-revision rejection
        ↓
authoritative transaction or equivalent serialization boundary
  ├─ decisive idempotency admission + replay/conflict decision where applicable
  ├─ decisive expected-revision check where applicable
  ├─ product/state mutation
  └─ required durable pending-work/outbox/recoverable marker where applicable
        ↓
COMMIT
        ↓
post-commit dispatch / downstream invocation where applicable
        ↓
response / success acknowledgement
```

Do not perform protected-resource-sensitive semantic processing before the applicable access boundary when doing so could leak protected existence or data.

A preflight idempotency or resource-revision lookup may reject work early, but it is not the decisive race-safety boundary. Where duplicate execution or stale writes matter, the decisive operation-identity claim/replay decision and expected-revision comparison are atomically coupled with the authoritative mutation or protected by an equivalent serialization invariant. Two concurrent requests cannot both win from the same unclaimed operation identity or the same expected mutable revision.

A learner/product success acknowledgement must not precede the authoritative durable commit it claims succeeded. When downstream asynchronous work is **required** to continue or reconcile the committed semantic operation, one durable discoverable work identity/marker must be established atomically with the authoritative state, or the required work must be deterministically and recoverably derivable from that committed state, before the acknowledgement can depend on it.

Dispatch itself does not execute inside the database transaction. A crash after commit but before dispatch must leave enough durable state for later recovery/dispatch. Optional non-semantic notifications may be best-effort when losing them does not lose required product work or semantic reconciliation.

## B. Asynchronous accepted work

Learner/admin-initiated asynchronous work composes the same access and validation boundary as other protected mutations before acceptance:

```text
bounded transport framing / credential extraction
  ↓
authentication where required
  ↓
authorization / capability / access filtering where required
  ↓
structural validation
  ↓
semantic preconditions
  ↓
optional preflight duplicate/stale rejection
  ↓
authoritative transaction or equivalent serialization boundary
  ├─ decisively admit/replay one logical operation
  ├─ establish accepted product/work state + logical work identity
  └─ persist durable pending-work/outbox/recoverable marker
  ↓
COMMIT
  ↓
acknowledge accepted/pending
  ↓
decisively confirm current dispatch eligibility + claim/fence one execution attempt where exclusive dispatch is required
  ↓
bounded dispatch/execution
  ↓
bounded eligible retry/backoff or explicitly controlled replacement/speculative attempt
  ↓
reconcile only a current legal completion
  ↓
persist accepted result + required recoverable semantic continuation
  ↓
SSE update hint and/or resource refresh
```

Provider webhooks use their own provider-authentication/replay semantics rather than learner/admin authentication, but still require validation and authoritative association before mutation.

Rules:

- an HTTP request being sent or a provider invocation being attempted is not durable work acceptance;
- reading `pending` is not an execution claim: where duplicate execution is not intentionally permitted, one dispatcher must decisively claim/lease/fence an execution attempt under authoritative state or an equivalent serialization invariant before dispatch, so concurrent dispatchers cannot both believe they exclusively own the same attempt;
- before dispatch, the decisive claim also confirms every current authoritative condition that can prohibit this execution or data egress where material, including cancellation/deletion/tombstone state and content/rights/privacy eligibility; if work is already ineligible, Core reconciles it without dispatch. A preflight read of those conditions is not the decisive race boundary;
- loss of a dispatcher after claim does not permanently consume the logical work. Recovery reconciles the claim and the known or ambiguous dispatch outcome; if remote execution may already have started, that uncertainty is treated like any other ambiguous remote outcome and redrive occurs only when idempotency/status/fencing makes it safe. A replacement/reclaim cannot let a stale prior completion satisfy current work;
- retry preserves one logical work identity and cannot duplicate accepted learner work, EvidenceFacts, content revisions, or paid/provider work;
- retry eligibility distinguishes transient, permanent, and ambiguous outcomes;
- exponential backoff and jitter are used where repeated immediate retry would worsen a transient route; exact budgets/counts remain implementation policy;
- a timeout does not prove the remote operation failed, so ambiguous outcomes remain unresolved until authoritative state can be established safely;
- when retries/replacements can overlap, Core preserves enough logical-work and execution-attempt identity/state to fence completion: only a result matching the expected authoritative work and current legal execution state may transition current work;
- intentional speculative/hedged/replacement execution, if later allowed, uses distinct execution identity/fencing and makes duplicate work/cost explicit and controlled rather than arising from a worker race;
- duplicate delivery of an already accepted result is idempotent; a late/stale result from a superseded execution cannot overwrite a newer accepted result or independently create Observation/EvidenceFact state, though it may be retained as operational provenance;
- when an accepted capability result requires downstream semantic continuation, the accepted result plus enough durable state to complete or reconstruct that continuation is committed atomically or made deterministically recoverable before the logical work is considered semantically complete; downstream replay is idempotent and cannot duplicate Observation, EvidenceFact, or learner-state transitions;
- cancellation propagates where safe, but cancellation of a caller connection does not roll back work already authoritatively committed;
- infrastructure/provider failure is not learner failure, a fake score, or a content-quality judgment.

## C. Internal capability invocation

```text
Core API authoritative work/state
        ↓
exact internal machine contract
        ↓
deadline + cancellation context
        ↓
Evaluator capability
        ↓
bounded response + provenance/quality/uncertainty
        ↓
Core API validates contract + work/execution association
        ↓
Core reconciles against current authoritative work/deletion/content-eligibility state
        ↓
accepted result + recoverable downstream continuation
        ↓
Assessment / content / product policy interprets an accepted result
```

HTTP/network success from Python establishes only that a bounded capability response arrived. It does not make evaluator/generator/validator output authoritative learner evidence, certification, content activation, product support, or a completed semantic operation. Evaluator execution must not mutate Core-API-owned product persistence directly.

Every network/provider boundary declares a caller deadline, retry eligibility/classification, idempotency behavior, safe fallback, and capacity/backpressure behavior when material. A separate circuit-breaker product is not required at bootstrap; circuit-breaker behavior may be introduced for repeatedly failing external routes when measured need justifies it.

## D. Privileged operational mutation

```text
bounded transport framing / credential extraction
        ↓
authenticated identity
        ↓
applicable privileged capability / access boundary
        ↓
structural contract validation
        ↓
current state + semantic preconditions
        ↓
legal operational mutation + required durable audit/work marker
        ↓
COMMIT
        ↓
response / post-commit dispatch
```

Capability remains distinct from identity, role label, learning authority, Assessment/evidence authority, validation-bypass authority, and historical ContentRevision mutation authority. Consequential privileged operations require durable reconstructable audit; admin UI never bypasses Core API to mutate storage/provider state directly.

### Actor, capability, and RBAC model

Authorization is capability-based underneath every role. Stable actor classes are:

| Actor class | Product meaning | Authority boundary |
|---|---|---|
| learner | authenticated person acting on their own learner/account scope | normal learner operations only unless separately granted privileged capability |
| privileged human | authenticated person with one or more explicit content/operations/security capabilities | may perform only granted operations through Core; privilege never becomes learning/evidence authority |
| service/system identity | first-party runtime/workload, including authorized AI content workers | receives only explicit capability grants/bundles; internal reachability is never privilege |
| external principal | provider/callback/source identity observed at a bounded external ingress | untrusted until authenticated/associated/reconciled; never direct product authority |

Capabilities remain the authorization truth. Named RBAC roles are stable default capability bundles for operator usability; they do not create another authority layer, and a human or first-party service identity may receive a narrower explicit grant when the bundle would be excessive. One actor may hold several compatible grants.

The minimum privileged capability families are:

- **content authoring** — draft/create/edit candidates or create a new semantic revision candidate;
- **AI content generation/processing** — request/run bounded generation, classification, deduplication, or correction proposal work;
- **content inspection/review** — inspect content/provenance/validation state and record review findings;
- **content validation decision input** — approve/reject or otherwise record an authorized validation/review outcome under the applicable policy without bypassing deterministic hard gates;
- **content release operations** — activate/release, quarantine/stop new assignment, retire, request/re-run revalidation, and resolve content incidents within legal state/preconditions;
- **user support operations** — perform bounded support mutations on learner/account product state without fabricating learner activity/evidence;
- **entitlement reconciliation** — reconcile effective commercial entitlement from accepted provider/commercial facts;
- **operational visibility** — inspect runtime/provider/content/work state appropriate to support/operations;
- **protected learner-data access** — inspect otherwise protected learner data only for an explicit authorized purpose and scope;
- **authorization administration** — grant/revoke roles/capabilities subject to the current actor's own grant authority;
- **security-sensitive operations** — perform explicitly authorized secret/config/recovery/destructive security operations.

A capability may carry resource, purpose, environment, or action scope. Possessing one capability does not imply a neighboring one: content review does not imply release; user support does not imply protected-data access; entitlement reconciliation does not imply learner/admin role; operational visibility does not imply mutation; release does not imply validation bypass.

Default bundles:

| Role bundle | Default capability intent | Explicit limits |
|---|---|---|
| `COLLABORATOR` | draft/create/edit content candidates; run/request AI generation; classify/tag; run duplicate/similarity checks; propose corrections; inspect non-restricted content/provenance needed for that work | no release/activation, entitlement mutation, protected learner-data access, grant administration, or security-sensitive operation |
| `REVIEWER` | `COLLABORATOR` plus semantic/content review, validation input, approve/reject candidate validation outcome where policy permits, request/re-run revalidation, inspect validation/provenance, and quarantine new assignment for a material content incident | cannot bypass hard gates, grant roles, or perform unrelated learner/security administration |
| `ADMIN` | operational content activation/release/retirement after gates pass; incident resolution; user-support operations; entitlement reconciliation; operational visibility; bounded protected learner-data access only when separately purpose-authorized; ordinary administrative mutations | no automatic authority to change canonical learning/evidence truth or top-level privilege/security policy |
| `OWNER` | `ADMIN` plus role/capability grant/revoke, security-sensitive configuration/operations, highest-level administrative control, and explicitly authorized destructive/recovery operations | still cannot bypass immutable history, hard content/evidence gates, or product truth |

`REVIEWER` approval is operational validation input, not learning/Assessment authority. Release is permitted only when the applicable deterministic/content/evidence/release gates already pass. No role contains a generic `override_validation` power.

AI workers are first-party service identities. They may be granted `COLLABORATOR`, `REVIEWER`, or narrower explicit capabilities and may perform most content labor—generation, editing, classification, semantic review, validation signalling, deduplication, and correction proposals—without a mandatory human-review step merely because the actor is AI. Consequential release, security, grant/revoke, protected-data, entitlement, or destructive operations still require the corresponding explicit capability and reconstructable audit; whether a human or service identity is permitted to hold such a capability is an operational security decision constrained by least privilege, not an assumption of trust.

Content activation/quarantine/retirement, protected learner-data access, entitlement reconciliation, capability-grant changes, security-sensitive configuration/operation, and other consequential privileged mutations require stronger authentication/authorization as appropriate plus durable audit sufficient to reconstruct actor identity, actor class, effective role/capability, target resource, reason, and outcome. Exact authentication factors remain implementation policy.

No privileged actor—human or service—may rewrite historical ContentRevision payloads, fabricate learner actions, bypass a known hard content/evidence failure, alter Band/Skill/Assessment truth, infer privilege from internal network position, or directly mutate another runtime's authoritative state outside normal Core operations. Learner entitlement remains completely separate from RBAC.

## E. Event/status propagation

```text
authoritative state change
        ↓
event/update hint
        ↓
SSE
        ↓
Web refresh/presentation
```

SSE is delivery, not state authority or semantic completion. Persistent current state remains queryable through resources, duplicate/reordered hints are tolerated, and reconnect cannot fabricate a transition.

## F. Scheduled/background work

Scheduled, cron-like, batch, maintenance, or background work is allowed when a real product/operational need exists; no cron platform is required merely by this design.

When used:

- duplicate-sensitive execution has one logical work identity;
- required work is durably registered/recoverably discoverable before another committed operation relies on it;
- where one execution attempt must be exclusive, authoritative claim/fencing precedes dispatch rather than relying on multiple workers observing `pending`;
- execution is bounded and observable;
- retries are safe/idempotent or are not performed;
- current/last outcome remains inspectable where operationally material;
- background work cannot silently mutate learner/evidence/content meaning outside the normal owning policies and legal transitions.

## G. Backpressure and capacity protection

When asynchronous/provider demand exceeds safe execution capacity:

- preserve already accepted authoritative work;
- bound in-process/external concurrency and pending dispatch;
- delay, rate-limit, or honestly reject new optional work according to product semantics;
- expose backlog/pending/degraded state rather than creating unbounded memory/process/provider queues;
- do not lower evaluation/content/evidence quality merely to drain backlog;
- introduce dedicated queue/broker infrastructure only when measured reliability/throughput requirements justify it.

## Distributed/network failure assumptions

- network calls are not atomic database transactions;
- cross-runtime/provider work cannot rely on distributed transactions by default;
- retry must not assume the first call failed solely because the caller timed out;
- provider/network partitions or ambiguous acknowledgement preserve pending/unknown operational truth until resolved safely;
- wall-clock timestamps may support audit/recency but do not by themselves establish global causal ordering, idempotency identity, concurrency correctness, or dispatch ownership;
- event/derived/provider state may lag authoritative Core-API state and must expose pending/stale/unavailable semantics honestly;
- distributed locks, leader election, Saga orchestration, or other distributed coordination are introduced only when an actual invariant spans multiple independent authorities and simpler single-owner transaction/work semantics cannot satisfy it.

# Failure-domain containment

Expected containment does not promise infrastructure that has not been selected:

- browser/client failure does not erase work already committed by Core API;
- SSE failure is recoverable through authoritative resource reads;
- Python/Evaluator failure leaves Go authoritative and productive evaluation pending/unavailable/invalid-at-capability-scope as appropriate, never converted to learner weakness;
- external-provider failure produces bounded degradation or unresolved work state and cannot authorize a lower-quality fallback;
- database failure before commit produces no durable-success acknowledgement;
- ambiguous database/network outcome is reconciled with authoritative idempotency/work identity before retry can create another logical operation;
- accepted capability completion cannot become permanently stranded from required Observation/Assessment/Progression continuation; recovery reconstructs or replays missing downstream work idempotently from durable authoritative state;
- cache failure falls back to authoritative state where operationally feasible and never changes truth;
- telemetry failure cannot rewrite business truth; only an explicitly required security/audit invariant may make absence of required durable audit block a consequential operation.

# Planner decision contract

Planning is a staged decision. Later stages may not redefine truth established by earlier stages.

```text
1. resolve TargetProfile + product-support scope
        ↓
2. expand unresolved target conditions
        ↓
3. consume Assessment/Progression → GapEvaluation / ActionIntent
        ↓
4. hard eligibility filtering
        ↓
5. activity candidate generation
        ↓
6. ranking among eligible candidates
        ↓
7. coherent plan + stable explanation
```

## Stage 1 — target/support resolution

Resolve the current target context and TargetProfile revision.

A target is fully resolved for target-relative planning/support only when `00-learning-experience.md`'s minimum holds: standard variant is known and at least one real Band constraint is known.

Keep these two failure classes distinct:

```text
missing/unknown learner target constraint
    = unresolved target condition
    ≠ product CoverageGap

known resolved target + product cannot serve it
    = CoverageGap / target support condition
```

If variant or all Band constraints remain unresolved, the Planner may offer only genuinely variant-independent/shared/foundational/diagnostic actions whose eligibility does not require the missing target condition. It cannot silently choose Academic/GT semantics, manufacture a Band target, or make a complete target-support/readiness claim.

Once the target scope is sufficiently resolved, resolve delivery mode/purpose/One-Skill-Retake eligibility conditions where material and the product-support state for that exact scope. Missing learner/external target information remains unresolved; an actual product inability remains a CoverageGap.

## Stage 2 — target-condition expansion

Expand the resolved/known target constraints into per-skill/variant/external-purpose conditions while preserving unresolved input explicitly.

An overall Band alone may correspond to multiple valid four-skill combinations. The planner must not invent hidden per-skill minima. Any working per-skill planning profile follows the separate non-authoritative planning-profile semantics in `00-learning-experience.md`.

## Stage 3 — evidence interpretation

Consume canonical outputs from Assessment and Progression. Planning does not rescore an Observation, reinterpret evaluator output, or invent another gap taxonomy.

## Stage 4 — hard eligibility

Remove candidates that violate any applicable hard condition:

- Required prerequisite;
- target variant/task compatibility when variant is known/material;
- delivery-mode compatibility for exam-readiness work when material;
- product coverage/support constraint;
- primary-purpose/evidence-candidacy compatibility for the requested action;
- exact content-revision release/operational eligibility for the actual intended use under the currently applicable validation policy/use scope;
- learner-specific exposure/novelty/independence eligibility where material;
- rights/privacy/source eligibility;
- accessibility/capture feasibility;
- immutable lifecycle condition;
- target-purpose/acceptance condition where the product knows the requested route cannot serve it.

Examples:

- Academic visual Task-1 practice is not eligible as GT Task-1 readiness work;
- variant-specific Task-1 work is not eligible while target variant is unresolved unless the learner explicitly selects it as non-target exploratory content;
- typing-only mock behavior is not sufficient delivery-mode practice when the learner explicitly targets eligible Writing on Paper;
- IELTS Online Academic exam-readiness is not eligible for a GT target or a target purpose that requires a test-centre route;
- a failed microphone capture cannot become eligible Speaking evidence by ranking it highly;
- an activity configured `NOT_EVIDENCE_CANDIDATE` cannot be selected to satisfy a `COLLECT_EVIDENCE` action merely because its expected score is favorable;
- content that is globally valid may still be ineligible for a learner's unseen/readiness claim after materially contaminating prior exposure.

## Stage 5 — activity candidate generation

Generate valid learning or assessment **activity candidates** for the current ActionIntent. This is planner candidate construction, not necessarily content generation.

Multiple modes may satisfy the same intent; candidate generation may vary delivery but not the target standard. Activity construction must resolve to an eligible exact ContentRevision before assignment. Existing eligible content is preferred; content supply generation/import is a separate downstream flow and is invoked only when actual supply demand remains.

When target resolution is incomplete, activity candidate generation is restricted to work that does not require inventing the unresolved condition.

## Stage 6 — ranking

Ranking may consider only **eligible** candidates. Useful ranking signals include:

- target urgency/test date;
- expected learning or decision value;
- due review;
- time fit;
- learner preference/friction;
- fatigue/session coherence;
- transfer/exposure diversity;
- operational cost after semantic validity is preserved;
- attributable intervention-effectiveness history where enough comparable learning history exists.

### Intervention effectiveness signal

The Planner may use a Practice-owned intervention-effectiveness interpretation to decide whether repeating an otherwise eligible intervention remains the best learning strategy. This signal does not create another `GapEvaluation` class and cannot reinterpret Assessment evidence.

Comparison is scoped by the existing intervention dimensions from `spec/07-PRACTICE.md`, including where material the canonical target, `ActionIntent`, Learning Mechanism, Practice Type, scaffold, context/transfer condition, difficulty/load, relevant Attempts/outcomes, subsequent admissible evidence, recency/comparability, and learner friction/fatigue.

Planner behavior is:

- insufficient or incomparable history → preserve uncertainty and rank without a fabricated plateau conclusion;
- repeated useful improvement → the same intervention may continue;
- justified diminishing returns or inadequate improvement → candidate generation/ranking may vary an existing mechanism/type/scaffold/context/difficulty/load/frequency/sequence dimension;
- evidence suggesting the diagnosis itself may be wrong → prefer an eligible discriminating evidence action rather than endlessly rotating remediation;
- technical/content/evaluator/provider failure → exclude that event from learner-efficacy inference and handle it under its real failure/product state.

Strategy variation preserves the same target standard, Required prerequisites, evidence policy, and historical Attempts/EvidenceFacts. It does not create novelty merely to appear adaptive. Where learner-facing consequence depends on the comparison, the relevant comparison basis/policy version and reason for continuing/changing strategy remain reconstructable.

### Daily load composition

After hard eligibility and candidate construction, the Planner composes **how much** work to schedule rather than treating ranking as an unlimited list. The plan may allocate load across due review/Knowledge retrieval, prerequisite acquisition, focused grammar/cohesion/language work, topic/context transfer, Listening/Reading practice, Writing/Speaking production, remediation, re-evidence, and exam-preparation work according to current need.

Composition consumes, where applicable:

```text
TargetProfile + target/test-date urgency
Assessment/Progression evidence state + GapEvaluation/ActionIntent
Required prerequisites
due review / retention state
learner available time
activity duration + difficulty/cognitive/performance load estimate
recent performance/retention history
fatigue, interruption risk, and session coherence
transfer/diversity/re-evidence need
product/capability availability and operational cost after eligibility
```

Rules:

1. Band thresholds define required quality; they do not prescribe daily counts.
2. Content inventory defines available eligible opportunities; inventory size is not dosage.
3. Daily dosage is a versioned product/planning policy over eligible work, not a new learning truth or one fixed formula for every learner.
4. The learner-specific DailyPlan is the current composition produced from that policy and current state; it may differ across learners with the same target Band.
5. Due review and Required prerequisites receive their semantic priority without forcing every due item into one session when available time/load makes that incoherent; deferred required/due work remains visible and schedulable rather than being marked satisfied.
6. Re-evidence is scheduled when Assessment/Progression calls for it; the Planner does not create repetitive assessment merely to fill time.
7. Production/timed/high-load work may be spaced or reduced when current fatigue/history would make another attempt low-value, while target/test-date urgency may legitimately increase exam-condition exposure within safe product bounds.
8. Numeric word/item/activity counts, load scores, daily maxima, and spacing/dosage coefficients are empirical/versioned policy when calibrated. The architecture does not invent claims such as a universal words-per-day requirement for Band N.
9. The plan records enough policy/provenance to explain why both the selected work and the chosen amount were reasonable under the current state.

Ranking still orders eligible candidates; load composition decides which ranked work fits the coherent session. Implementations may combine these computations internally, but must preserve the semantic distinction.

### Target trajectory advisory

After the current eligible work and coherent planning envelope are known, the product may derive a learner-facing trajectory advisory without feeding the projection back into learner truth. The question is whether the **current plan/envelope appears plausibly sufficient under the applicable projection policy** for the learner's declared target/date, not whether the product can predict an IELTS result.

The projection may consume, where justified:

```text
TargetProfile + fixed test date when supplied
time remaining
current supported / unresolved target conditions
GapEvaluation / ActionIntent
available study envelope
recent admissible progression/history
current planned workload
calibrated empirical rate/load assumptions when available
product/evaluator/content availability as a separate constraint axis
```

Rules:

1. missing history, insufficient calibration, or unresolved learner evidence that prevents a defensible projection yields an unresolved advisory rather than fabricated velocity;
2. missing learner evidence contributes uncertainty, not assumed weakness;
3. product/evaluator/content inability remains distinguishable from learner pace and may itself explain why trajectory cannot currently be established;
4. a material at-risk interpretation names its causes; a known required-work-versus-time-envelope shortfall may be reported when the policy can justify it;
5. the projection does not lower a target, invent per-skill minima, change a test date, certify readiness, or alter hard eligibility;
6. learner options may include changing available study time, scheduling/focus among eligible work, collecting missing evidence, or explicitly editing the target/date; Core never performs that target/date change implicitly;
7. do not emit `impossible` or equivalent certainty unless a separately justified policy supports it; normally preserve `unresolved`, `at risk`, or `current envelope insufficient` semantics;
8. projection-policy version, material inputs/state references, calibration provenance, computed-at context, and explanatory causes are reconstructable wherever the advisory has learner-facing consequence.

Exact trajectory coefficients, pace assumptions, and risk cutoffs are empirical/versioned policy. They are not Band, Assessment, Progression, or readiness truth.

### Ranker non-authority invariant

A ranker may reorder eligible candidates. It may never:

- make an ineligible candidate eligible;
- bypass a Required prerequisite;
- alter GapEvaluation or ReadinessEvaluation;
- reinterpret evaluator output;
- hide a CoverageGap;
- invent/resolve missing target constraints;
- lower a target threshold;
- convert preference into ability evidence;
- ignore a material variant/delivery constraint;
- override content validation/release/operational ineligibility;
- upgrade evidence candidacy/admission;
- certify a learner.

## Stage 7 — plan/explanation

A DailyPlan is a snapshot/recommendation produced from named state, not assignment authority. It records enough provenance to reconstruct why each activity was eligible, selected, and included at its planned amount, including where material the TargetProfile/target-context revision, learner/Progression state reference/version, due-review/retention state, applicable dosage/load-policy version or configuration, available-time/session constraint, content/release state, product support/coverage version, content-revision references, and unresolved conditions used by planning. When a learner-facing trajectory advisory or intervention-strategy change is emitted, its applicable projection/effectiveness policy version, material comparison/input references, and explanatory causes are likewise reconstructable rather than being an opaque model assertion.

Before actual PracticeActivity/AssessmentActivity assignment or learner exposure, Core re-evaluates every **current** hard condition that can change. This includes where material current target context, learner/evidence state, product coverage/support scope, ContentRevision release/quarantine/operational state and validation eligibility for the actual intended use under the currently applicable policy/use scope, rights/source eligibility, learner exposure/novelty/independence and reservation state, and delivery/capture feasibility.

The assignment boundary is therefore:

```text
plan candidate/reference when one exists
        ↓
current hard-eligibility re-evaluation
        ↓
decisive reservation/assignment where applicable
        ↓
learner-safe projection
        ↓
delivery / actual exposure
```

Where a mutable authoritative hard condition can change concurrently and make the assignment illegal, its decisive current-state check is atomically/conditionally coupled to reservation/assignment or protected by an equivalent serialization invariant. A candidate/preflight recheck alone is insufficient: a concurrent quarantine, release revocation, target/support change, or learner-specific reservation/exposure change must either be reflected in the assignment decision or cause that assignment to fail/reselect. This does not require one global lock or transaction across unrelated state.

A previously eligible plan item that is now ineligible is not executed merely because it remains in a saved plan; Core reselects another eligible action or exposes the truthful current blocker. Plan-time explanation remains historical provenance, not present eligibility. A full plan regeneration is unnecessary when a smaller current eligibility/reselection check is sufficient. No plan TTL is frozen.

When validation materially gates assignment, assignment-time provenance remains sufficient to reconstruct the exact ContentRevision, intended-use/consequence scope, applicable validation-policy version, and the compatible ValidationDecision/current release state that justified eligibility. Exact storage shape is not fixed and these need not all be fields on Attempt. Plan-time provenance remains separate; neither an older DailyPlan nor an unrelated older/newer ValidationDecision overrides the currently applicable policy/use scope.

Assignment records the exact current eligibility/content revision actually used; plan provenance and assignment authority remain distinct.

Tie-breaking should be deterministic/stable enough that unchanged learner state does not create arbitrary plan churn.

# Flow A — target setup and diagnostic

```text
1. Web collects known TargetProfile fields + planning constraints
2. Core API validates known variant/delivery/purpose combinations without inventing missing fields
3. Core API persists TargetProfile/target-context revision + unresolved target conditions
4. actual unsupported product conditions are surfaced as CoverageGap;
   missing learner target information remains an unresolved target condition
5. learner selects quick/full diagnostic UX shape
6. Core API selects only items compatible with the known target context;
   variant-specific sampling waits for variant resolution unless explicitly exploratory
7. learner attempts items
8. objective L/R items may score deterministically
9. eligible W/S attempts may use Evaluator
10. Evaluator returns observations + provenance + uncertainty
11. Assessment applies diagnostic sampling/eligibility semantics
12. Progression derives only justified learner-state interpretation
13. Planner executes stages 1–7
14. Web shows sampled, unresolved-target, evidence, and product-coverage conditions distinctly
```

Diagnostic is a primary activity purpose. Whether a diagnostic Observation is admissible for a higher-consequence claim is a separate Assessment decision based on its pre-declared evidence candidacy and actual attempt conditions.

A completed diagnostic run is not synonymous with a complete learner model, a resolved TargetProfile, or certification.

# Flow B — Daily Plan

```text
known TargetProfile / target context
  ↓
resolved + unresolved target conditions
  ↓
product support where evaluable
  ↓
GapEvaluation / ActionIntent where evidence permits
  ↓
hard eligibility
  ↓
valid activity candidate generation
  ↓
ranking + coherent load composition
  ↓
coherent DailyPlan + reason codes + amount/load provenance
  ↓
target-trajectory advisory when enough justified basis exists
```

Representative reason codes include:

```text
PREREQUISITE_GAP
ABILITY_GAP
INSUFFICIENT_EVIDENCE
CONFLICTING_EVIDENCE
STALE_EVIDENCE
SCAFFOLD_DEPENDENCE
TRANSFER_GAP
FLUENCY_GAP
REVIEW_DUE
EXAM_CONDITION_GAP
DELIVERY_MODE_PREPARATION
PRODUCT_COVERAGE_BLOCKED
```

`PRODUCT_COVERAGE_BLOCKED` is not a learner GapEvaluation and is not used for merely missing TargetProfile input.

Swap/Skip/Shorten/Change-skill actions operate within eligible choices. They do not mark skipped requirements satisfied. Starting an activity from this plan still crosses the current assignment eligibility boundary above.

# Flow C — ordinary practice attempt

```text
1. Web requests PracticeActivity from a current candidate/plan item
2. Core API rechecks current hard eligibility and performs decisive reservation/assignment where applicable
3. Core API returns the exact assigned item revision + target + variant/context + conditions
   + primary_activity_purpose + evidence_candidacy through the learner-safe API projection
4. learner performs task
5. Web submits Attempt with actual scaffold/exposure/delivery metadata
6. Core API or Evaluator produces Observation
7. NOT_EVIDENCE_CANDIDATE → feedback/remediation/review handling only
   ASSESSMENT_MAY_ADMIT  → Assessment evaluates actual claim-scoped eligibility/inference
8. Progression updates interpretation only when justified
9. Planner derives next target-relevant action
10. Web presents outcome without fabricating a micro Band change
```

Primary purpose (`TRAINING`, `DIAGNOSTIC`, `ASSESSMENT`, or `READINESS`) does not determine evidence admission. `ASSESSMENT` means focused measurement is the activity purpose; it does not mean the resulting Observation is already admitted. Assessment may reject a candidate Observation after seeing actual assistance/exposure/evaluator/provenance conditions; a favorable result cannot retroactively upgrade a non-candidate activity.

Submission/retry is idempotent where network repetition could duplicate history or cost. Once an Attempt begins, later content activation/retirement cannot rewrite the revision referenced by that Attempt.

# Flow D — Writing/Speaking asynchronous evaluation

```text
Web submits work
  ↓
Core API applies auth/access + structural validation + semantic preconditions
  ↓
authoritative transaction
  ├─ persist Attempt/submission state
  ├─ persist logical evaluation work identity
  └─ persist pending/recoverable dispatch state
  ↓ COMMIT
ACK accepted/pending
  ↓
current-eligible decisive execution-attempt claim/fence
  ↓
bounded dispatch / one or more explicitly identified execution attempts as policy permits
  ↓
criterion observations + provenance + uncertainty
  ↓
Core API validates contract + logical-work/execution association
  ↓
Core API fences/reconciles result against current authoritative evaluation/deletion/content-eligibility state
  ↓
persist accepted result + Observation or recoverable downstream-continuation source/marker
  ↓
Assessment → EvidenceFact / ReadinessEvaluation where eligible
  ↓
Progression → learner-state interpretation
  ↓
Planner → next action
  ↓
Web receives status/result via SSE or resource refresh
```

For the ordinary Speaking route, browser recording before submission is draft interaction state. The learner may replay or replace that draft when the activity permits. Once submitted, the Attempt is immutable; feedback-driven re-record/retry creates a new related Attempt rather than rewriting the submitted performance.

Rules:

- accepted learner work and required evaluation continuation are durably discoverable before learner-visible acceptance;
- pending/unavailable are valid non-score states;
- timeout/provider failure never becomes a low learner score and does not by itself prove an execution did no work;
- normal competing workers cannot silently double-dispatch the same exclusive execution attempt; any intentionally concurrent replacement/speculative execution is separately identified/fenced;
- a claimant/worker failure cannot permanently strand accepted evaluation work; claim recovery follows the ambiguous-dispatch rule above and never assumes that missing local acknowledgement proves remote non-execution;
- retries/replacements preserve the same logical evaluation work while retaining distinct execution-attempt identity/provenance whenever executions can overlap;
- exactly one current legal completion may satisfy the logical evaluation outcome; duplicate delivery of that accepted completion is idempotent;
- a late or stale completion from a superseded execution cannot overwrite the accepted result, reopen terminal work, or independently create Observation/EvidenceFact state;
- accepting a result does not permit a crash gap that permanently strands required semantic continuation: accepted result/Observation and/or durable recoverable continuation state must allow idempotent reconstruction of missing Assessment/EvidenceFact/Progression consequences;
- no progression occurs from failed, stale, superseded, deleted/ineligible, or invalid evaluator output;
- fallback must meet the same applicable quality/privacy floor.

## Realtime Speaking interaction

Realtime conversation is an optional learner interaction layered over eligible Speaking training/practice. Core retains authoritative session/admission/entitlement state; external realtime/audio/AI capability owns no learner truth.

```text
eligible activity + current entitlement/capability admission where gated
  ↓
logical Speaking interaction session
  ↓
AI prompt / question / follow-up
  ↓
learner spoken turn + capture provenance
  ↓
bounded transcript/acoustic/semantic interpretation where eligible
  ↓
next responsive AI turn
  ↺ until completed / abandoned / unavailable
  ↓
post-session feedback / normal Assessment path only when separately eligible
```

System rules:

1. one logical session preserves activity target, Speaking part/role-play scope, assistance/evidence configuration, entitlement admission, and reconstructable turn/capture/provider provenance where material;
2. Part 1 may use short responsive turns, Part 2 preserves preparation then learner long-turn independence, and Part 3 may use responsive follow-up discussion. Realtime follow-up must not interrupt or replace the learner operation required by the selected activity;
3. interruption/barge-in is a delivery behavior only when the configured interaction supports it; whether interruption is pedagogically/evidentially meaningful remains target/Assessment scoped;
4. learner silence/timeout is distinguished from capture/network/provider silence. Only trustworthy learner-performance conditions may become an Observation; technical ambiguity remains capture/provider failure;
5. AI latency does not consume learner-performance timing unless the activity explicitly defines a target-like interaction condition that can distinguish system delay from learner delay;
6. dropped connection moves the logical session to reconnect/degraded handling rather than creating a new learner Attempt or paid logical work automatically. Reconnect resumes/fences the same accepted session where safe; duplicated turn delivery is idempotent;
7. if reconnect cannot safely restore the interaction, preserve completed turns and mark the remaining session partial/abandoned/unavailable as appropriate. Partial completion supports only its actual scope;
8. graceful fallback to ordinary record→submit is allowed only for a learning purpose whose semantics remain valid after the mode change. It cannot silently satisfy a realtime/readiness/mock condition that required responsive interaction;
9. provider/evaluator unavailability may delay feedback or end/degrade the realtime route while leaving the ordinary Speaking route usable; it never authorizes lower evidence quality;
10. completion, abandonment, network failure, entitlement change, or paid-provider accounting does not itself admit EvidenceFacts. Normal Assessment consumes only eligible Attempts/Observations.

Conceptual lifecycle:

```text
created → active ↔ reconnecting → completed
            ├───────────────→ abandoned
            └───────────────→ unavailable
reconnecting ├──────────────→ abandoned
             └──────────────→ unavailable
```

`completed` means the configured interaction ended successfully, not that Speaking capability/readiness is supported.

# Flow E — objective Listening/Reading attempt

```text
Attempt
  ↓
answer-key + instruction/word-limit validation against exact content revision
  ↓
raw result
  ↓
Observation with variant/context/conditions
  ↓
Assessment eligibility/inference scope when evidence-candidate
  ↓
EvidenceFact when valid
```

Reading Band inference uses the correct Academic/GT scoring/context policy.

# Flow F — feedback/remediation

```text
Observation(s)
  ↓
GapEvaluation
  ↓
ActionIntent
  ↓
feedback-focus policy (`spec/07`)
  ↓
ErrorPattern / RemediationPattern when relevant
  ↓
Learning Mechanism
  ↓
Practice Mode
  ↓
fresh eligible Practice Item revision
  ↓
learner Attempt / suitable learning outcome / separately eligible re-evidence
  ↓
comparable intervention history when sufficient
  ↓
intervention-effectiveness interpretation
  ├─ useful improvement → continue when still appropriate
  ├─ diminishing / inadequate improvement → change an existing strategy dimension
  └─ diagnosis uncertain → collect discriminating evidence
  ↺
next eligible intervention selection
```

Direct `wrong answer → fixed exercise id` mapping is not a canonical remediation policy. Detectable secondary issues may be recorded/deferred rather than surfaced when they are outside the current feedback focus. The product also must not run `same GapEvaluation → essentially same remediation` indefinitely merely because the gap still exists. Strategy continuation/change follows the Practice-owned effectiveness rules above: too little comparable history stays unresolved, effective repetition may continue, product/runtime failure is not plateau, and strategy change never rewrites historical Attempts/EvidenceFacts or the learner target.

# Flow G — review

```text
review request
  ↓
queue kind
  ├─ knowledge retrieval
  ├─ error remediation
  └─ re-evidence
  ↓
eligible activity + exact content revision
  ↓
Attempt
  ↓
update only the corresponding semantic state
```

One Review screen may present these together; backend meaning remains distinct.

# Flow H — media lesson creation

```text
1. learner supplies source URL/reference
2. Core API treats it as untrusted reference and resolves provider metadata/source eligibility through the allowed media boundary
3. rights/transcript state is established
4. authorized/minimum necessary text may be sent to Evaluator
5. Evaluator proposes segments/targets/difficulty/prompts
6. Core API validates rights + canonical target + practice mapping
7. resulting content is admitted only through the applicable content validation/revision path
8. learner previews/saves MediaLesson when eligible
9. later attempts use the normal Practice/Attempt path
```

Evaluator access never implies authorization to download/copy arbitrary media. Learner/provider URLs do not imply arbitrary network access or scraping. Generated media-derived prompts do not bypass content quality or revision semantics.

# Flow I — mock/readiness

```text
resolved TargetProfile variant + delivery target when material
  ↓
MockRun
  ↓
Listening shared
Reading variant-correct
Writing Task 1 variant-correct
Writing Task 2
Speaking shared construct
  ↓
section observations / estimates
  ↓
ReadinessEvaluation where evidence candidacy + normal eligibility permit
  ↓
GapEvaluation
  ↓
exam-preparation plan
```

A normal target-relative full mock requires a resolved variant. `READINESS` is the primary purpose of a normal mock; it is not an automatic evidence decision. A mock Observation contributes to a claim only when pre-declared as an evidence candidate and independently admitted by normal Assessment policy.

A mock result is a target-like performance snapshot. Full IELTS readiness becomes `SUPPORTED` only through its versioned integrated EvidenceRequirement; one favorable mock, four independent micro-skill claims, or four per-skill Band claims do not automatically satisfy exam-condition/integration requirements that the readiness policy makes material.

Mock construction issues one integrated content demand whose component demands preserve the applicable external section/task/part composition, one coherent variant, exact revision/provenance, delivery overlay where material, and learner-specific independence/unseen conditions required by the readiness policy. Component reuse is allowed for familiarisation only when it does not contradict the intended inference.

A mixed Academic/GT mock is invalid for normal full-test readiness unless explicitly created as non-certifying comparison practice.

Delivery-mode practice may change interaction conditions without changing scoring/Band semantics.

# Flow J — target supported/unresolved

Target interpretation preserves four separate axes:

```text
TargetProfile input resolved?
learner evidence supports each applicable claim?
exam/delivery readiness condition supported where material?
product supports serving/measuring the requested scope?
```

For each known TargetProfile condition, Core consumes the exact Assessment claim evaluation; it does not infer support from completion or nearby capability.

Composition rules:

1. an explicit per-skill minimum requires the corresponding current per-skill Band claim at or above that minimum;
2. an overall-only target requires supported current Band claims for all four skills, then applies the official overall-score rule; it does not fabricate per-skill minima;
3. mixed overall + per-skill constraints require both the derived overall and every explicit minimum;
4. Academic/GT-scoped claims must use the applicable variant evidence/conversion/context;
5. material delivery readiness is evaluated separately from the underlying Band claim;
6. full IELTS readiness additionally requires the applicable integrated/exam-condition EvidenceRequirement; a collection of supported micro-capability claims is not a full-readiness shortcut.

Target-input uncertainty and learner-evidence uncertainty remain separate. A missing target variant/Band condition is not converted into `INSUFFICIENT_EVIDENCE` about learner capability.

Evidence states remain exact: missing → insufficient, material disagreement → conflicting, expired recency → stale, positively below threshold → not yet supported. Each unresolved condition yields the corresponding Progression evidence need/ActionIntent where product capability exists.

When all learner evidence conditions are supported, the app may state that **current evidence supports the declared TargetProfile** only if the target itself is sufficiently resolved. Product support remains a separate statement.

When a required evaluator/content/runtime/product capability is missing, preserve the learner claim as unresolved and surface the CoverageGap. Do not translate product inability into `ABILITY_GAP`, a low score, or mandatory hidden-human review.

It must never state that the learner is guaranteed an external result.

# Flow K — content supply and learner assignment

Content supply is reuse-first and mechanism-neutral. A `ContentDemand` describes **what eligible opportunity is needed**, not which database pool, generator, or source must satisfy it.

For a concrete learner/product need, demand resolves only the applicable dimensions:

```text
canonical Skill/Knowledge target(s)
Curriculum node/phase when the action is node-bound
Practice Type or Assessment Type
primary activity purpose + evidence candidacy
variant / Content Context / official family / Presentation Class when material
task / section / part / integrated scope when material
difficulty / scaffold / transfer-distance requirements
response interaction + answer/rubric/evaluator route
delivery interaction when material
rights/source constraints
learner exposure / novelty / independence conditions for the intended consequence
quantity/diversity/stability demand only where the owning learning/evidence/coverage policy makes it material
```

Fields that are semantically inapplicable remain inapplicable; demand construction must not invent a family, context, Band, or novelty requirement merely to fit a generic pool query.

The same ContentRevision may satisfy several different demands over time only when it independently passes the current use-specific packaging, validation, rights, release, and learner-exposure gates.

```text
canonical learning/product need
        ↓
ContentDemand
        ↓
query eligible content pool
        ↓
sufficient content for required target/context/difficulty/diversity?
    ├─ yes → reuse eligible revision
    └─ no  → obtain candidate through an applicable supply route
              ├─ authored/imported
              ├─ deterministic/local generation
              └─ bounded external/AI generation when eligible
                    ↓
              ContentRevision
                    ↓
              applicable validation
                    ↓
              pool admission / release eligibility
        ↓
current hard eligibility for the actual learner/use
        ↓
learner-specific exposure/novelty/independence eligibility
        ↓
decisive reservation/assignment under required serialization
        ↓
learner-safe projection
        ↓
delivery
        ↓
actual exposure recorded when established
        ↓
Attempt references that revision
```

Rules:

1. generation is optional; a release may be fully supplied by authored/imported/deterministic assets when they satisfy coverage and operational requirements;
2. generation/import is requested only when actual content demand remains after eligible pool reuse, or may be pre-generated/batched from demonstrated future demand so learner requests do not unnecessarily wait on generation;
3. reuse optimization must never override target/family/context/presentation/difficulty/diversity/rights/evidence requirements;
4. content validation and assignment are separate: a globally valid revision can still be ineligible for a specific learner/use because of exposure, independence, novelty, quarantine, target/product-support change, or another current hard condition;
5. corpus similarity is not a universal rejection rule; similarity facts are interpreted against the intended learning/evidence purpose;
6. generator output and validator signals cannot activate content by themselves; Core-API-owned deterministic product policy applies the owning semantic/coverage rules;
7. higher-consequence use requires the stronger applicable validation/evidence conditions defined upstream; low-consequence training does not require unrelated high-consequence checks;
8. every assigned activity resolves the exact immutable revision and current release/validation/operational eligibility for the actual intended use under the currently applicable validation policy/use scope before learner exposure, regardless of whether an older DailyPlan or unrelated ValidationDecision referenced another state;
9. the decisive assignment enforces the mutable current hard gates described in Stage 7 under a transaction, conditional write, or equivalent serialization invariant where concurrency can invalidate them; learner-specific unseen/exposure/independence/uniqueness and reservation state are included when material so concurrent requests cannot both consume the same protected opportunity incorrectly;
10. when the intended use requires proven unseen or sufficiently independent conditions, `UNKNOWN`, missing, or ambiguous material exposure is not treated as unseen. Core selects another eligible opportunity or preserves a truthful unresolved/ineligible state for that consequence; training that does not require novelty may remain eligible;
11. reservation/assignment for delivery is not actual learner exposure. A failed/disconnected delivery must not fabricate `seen`; actual ExposureContext follows `spec/10-CONTENT-MODEL.md`;
12. a reservation may temporarily exclude concurrent assignment until it is reconciled/released; exact reservation timeout/recovery mechanism remains implementation policy.
13. no separate canonical `practice_pool`, `assessment_pool`, `readiness_pool`, or `unseen_pool` is required: current eligibility is derived from purpose/candidacy, validation/use scope, release state, content identities, and learner ExposureContext; an implementation may index/cache these dimensions without creating new authority.
14. supply sufficiency is scoped to the demand set, not raw inventory count. Ten interchangeable near-duplicates do not satisfy a demand requiring materially different transfer opportunities; one high-quality revision may satisfy several training demands where reuse remains valid.
15. when an EvidenceRequirement or readiness policy requires multiple/diverse/independent opportunities, the demand carries that requirement symbolically/versionedly; product design does not guess the numeric threshold.
16. authored, licensed/imported, deterministic/template-produced, pre-generated AI-assisted, runtime-generated, learner-provided, and media-backed candidates enter the same validation/release path. Source preference may optimize cost/latency/editorial quality only after semantic eligibility is preserved.

Content supply may be asynchronous. If the desired content is unavailable, the planner may use another genuinely eligible activity or expose an honest product/content gap; it must not silently relax semantic or evidence requirements merely to avoid generation latency/cost.

## Parallel AI content-authoring/materialization

Bulk or pre-generated content may be produced by many independent first-party AI/service workers without introducing an orchestration platform. Parallelism is safe only when work is deterministically partitioned before generation.

```text
canonical schema + allowed stable canonical/external IDs
        ↓
versioned content-demand set
        ↓
deterministic non-overlapping shard contract
        ↓
independent authorized workers
        ↓
candidate content packages + provenance
        ↓
structural/reference validation
        ↓
semantic/answer/rubric/rights validation as applicable
        ↓
duplicate + purpose-aware similarity checks
        ↓
deterministic integration of accepted candidates
        ↓
normal ContentRevision validation/release lifecycle
```

A shard identifies a bounded slice from pre-existing canonical dimensions such as target IDs, official family/context/presentation, Practice/Assessment role, variant, and a deterministic shard/slot identity. Workers consume that contract; they cannot invent new Skill, Knowledge, family, context, presentation, Practice, Assessment, or Band identities to make generation convenient. Two normal workers assigned distinct shards must not be expected to coordinate with each other or edit the same candidate identity.

Candidate-package rules:

1. candidate IDs/slot identities are deterministic or collision-checked within the shard contract; duplicate candidate identity is rejected rather than last-writer-wins;
2. exact duplicate content is rejected unless the owning purpose explicitly requires identity/repetition; semantic near-duplicates are measured as facts and judged by intended training/transfer/evidence use rather than globally forbidden;
3. canonical references are validated against the current allowed registry/source before integration; unknown/invented references fail closed;
4. AI-proposed difficulty, learner level, Band suitability, scoring, or evidence claims are candidate metadata only. They gain consequence only through the owning Band/content/evaluator/calibration policy;
5. generation/review by multiple AI workers does not constitute independent evidence merely because models/threads differ; validation policy decides what signals are sufficient;
6. provenance identifies generator/worker/configuration/source/shard as needed for reconstruction, rights/privacy review, quality incidents, and replacement;
7. authored and AI-produced candidates enter the same ContentRevision immutability, ValidationDecision, release, quarantine, retirement, and historical-reference rules;
8. deterministic integration rejects shard overlap, duplicate IDs, unresolved canonical refs, invalid package structure, or another merge ambiguity rather than asking workers to negotiate;
9. a candidate package is authoring/materialization input, not runtime assignment authority and not learner evidence.

Git may store text/source candidate packages and reviewed authored source when appropriate. Large generated audio/media or other heavy binary artifacts may live behind the governed object/media boundary with integrity/provenance references rather than becoming repository blobs. `06-implementation-stack.md` owns the source/materialization boundary; runtime product storage remains a released projection, not canonical authoring authority.

No queue, registry service, agent framework, or bespoke content-orchestration platform is required by this model. Introduce one only if measured throughput/recovery/concurrency requirements cannot be met by deterministic shard production plus normal repository/runtime materialization.

# Flow L — content report, quarantine, revalidation, and retirement

Content operational safety is independent from immutable revision identity and from whether a revision previously passed validation.

Representative flow:

```text
Attempt / learner report / operational signal / validator-policy regression
        ↓
triage material risk + affected use stages
        ↓
new ValidationDecision when policy/revalidation is rerun on unchanged semantic content
        ↓
recompute current release/use eligibility
        ↓
preserve historical ContentRevision + ValidationDecisions + learner facts
        ↓
apply consequence to current reservations / in-flight work / evidence as applicable
        ↓
investigate / revalidate / replace
        ├─ same semantic revision verified → restore eligible assignment when release permits
        ├─ material content fix required → create new ContentRevision → validate → activate when eligible
        └─ no safe/valid route → retire from new assignment
                                    ↓
                         replacement/supply demand if coverage needs it
```

Current release/use recomputation resolves the applicable validation policy and intended-use/consequence scope according to `spec/10-CONTENT-MODEL.md`; it does not implement a global “newest ValidationDecision wins” rule.

Do not collapse validation state, release eligibility, operational safety, or incident consequence into one global ContentStatus/defect enum. A revision may be semantically validated yet not activated for a release, or may have been active and later be quarantined operationally pending investigation.

Current assignment disposition is therefore derived rather than one canonical status enum:

- **assignable/active for a use** — every current semantic, validation, release, rights, operational, and learner-specific gate for that use passes;
- **quarantined** — new assignment is temporarily blocked for the affected use while safety/quality/rights validity is investigated or revalidated;
- **superseded** — another semantic revision is preferred/current for a lineage; this alone does not rewrite history or determine whether the older revision may still be assigned for some use;
- **retired from new assignment** — product/release policy intentionally stops future assignment while historical references remain resolvable;
- **unavailable/ineligible for a use** — a current source, rights, validation, exposure, consequence, or product condition prevents assignment for that use.

These dispositions may differ by intended use. A revision can remain valid low-consequence training content while being ineligible for independent evidence, or remain historically valid after retirement from all new assignment.

Incident consequence is scoped by the discovered defect and the current use stage, not merely by `stop new assignment`:

- **new assignment** — a known material risk blocks new reservation/assignment whenever the affected use is no longer eligible;
- **reserved but not exposed** — reservation may be cancelled/reconciled when continuing exposure is unsafe, rights/security invalid, semantically invalid, or otherwise prohibited for that use; cancellation does not fabricate learner exposure;
- **already exposed / draft in progress** — preserve what the learner already received/worked on, but stop/replace/mark unavailable when continuing the task would be unsafe, rights/security invalid, answer-compromised, or semantically invalid; a minor non-semantic defect need not have the same consequence;
- **submitted Attempt / pending or running Evaluation** — preserve the historical submission, but block, narrow, invalidate for the affected claim/use, or re-route scoring/evaluation/evidence processing when the defect makes the intended inference invalid; infrastructure/content incident does not rewrite the Attempt as learner failure;
- **existing Observation / EvidenceFact / historical completion** — preserve historical records; when the defect changes historical inference validity, Assessment/Progression re-evaluates current support/certification using explicit incident/provenance rather than deleting or mutating history.

Incorrect answer keys/rubrics, prohibited answer leakage, canonical-binding errors, rights/privacy/security defects, independence contamination, and minor presentation defects can therefore have different consequences without requiring a universal defect taxonomy.

Operational actors exercise only applicable privileged capabilities defined below. They are operators, not learning or Assessment authorities.

No authorized operator/admin action, regardless of role level, may bypass a known applicable hard failure such as:

- invalid canonical reference or incompatible IELTS family/context/presentation binding;
- known incorrect answer key/rubric/model where required;
- prohibited answer leakage for the intended use;
- rights/privacy/security failure;
- evidence-critical exposure/independence failure for a use that requires independent evidence.

The cause must be repaired or the intended use changed legitimately, then the content/revision must pass the applicable policy again. A generic `override_validation` capability that converts a known hard failure into active learner content is forbidden.

Historical learner facts are never repaired by mutating an old revision or earlier ValidationDecision. If a later discovery changes how historical evidence should be interpreted, preserve the original Attempt/Observation/EvidenceFact and apply the owning Assessment/Progression policy with explicit provenance.

# Privileged content-operation capabilities

Flow L uses the content-authoring, AI-generation/processing, content-inspection/review, validation-input, and content-release capability families defined in the actor/RBAC section above. Content-specific execution includes candidate editing/new-revision creation, review/validation input, quarantine/stop-assignment, activation/release after gates pass, retirement, regeneration/replacement/revalidation requests, and incident/report resolution. These are capability semantics; the default RBAC bundles only package them.

Invariants:

1. an operational capability grant is not learning authority;
2. an operational capability grant is not Assessment/evidence authority;
3. an operational capability grant is not permission to mutate an established ContentRevision semantic payload;
4. an operational capability grant is not validation-bypass authority;
5. no capability may convert a known applicable hard failure into assignable content;
6. activation/release capability may act only on content that already satisfies the applicable semantic, validation, release, and operational policy;
7. authorization implementation may later map authenticated identities/roles to these capabilities without redefining their meanings.

The default RBAC bundles above are canonical authorization convenience semantics; exact persistence, identity-provider mapping, custom/narrow grants, authentication factors, and machine-contract encoding remain implementation/authorization concerns. API operations consume these capability meanings through `05-api.md` without creating a second authorization taxonomy.

# Learner entitlement vs operational authorization

Commercial learner entitlement and privileged operational authority are separate dimensions. Marketing tier names such as Free/Pro are mutable presentation; the canonical product needs only provider-neutral capability entitlement semantics.

Conceptually, Core owns the effective entitlement state and any pending provider reconciliation:

- **absent/not entitled** — the gated capability cannot admit new paid use; ordinary ungated learning remains available;
- **active** — the named gated capability may be admitted subject to all normal product/support/rate/cost/capture gates;
- **pending change** — activation/renewal/downgrade/payment information is delayed, ambiguous, or not yet reconciled; provider output alone does not grant/revoke authority;
- **expired/downgraded** — future use of capabilities no longer entitled is blocked from the effective product boundary;
- **restored/reactivated** — future eligible gated use becomes available again after authoritative reconciliation.

Rules:

1. entitlement activation/revocation is an authoritative Core-owned product decision derived from accepted commercial/provider facts, not direct authority of a payment webhook, client flag, or provider dashboard;
2. ambiguous provider/payment state does not change effective entitlement by itself; the last authoritative state remains until another Core-owned rule (including a known effective expiry) changes it or the observation is reconciled. Ambiguity must not invent free premium access or erase an unexpired accepted entitlement merely from a callback;
3. a gated session admitted while entitlement is valid carries a bounded admission snapshot/logical work identity. Later expiry/downgrade blocks new gated sessions but does not duplicate, erase, or arbitrarily rewrite already accepted learner work; the admitted session may reach a safe bounded completion or graceful termination according to current abuse/security/provider/product constraints;
4. entitlement loss/restoration never rewrites Attempts, Observations, EvidenceFacts, certification history, TargetProfile, or learner-owned saved data;
5. entitlement does not change Skill/Band thresholds, evidence eligibility/quality, prerequisites, content validation, or product CoverageGap meaning;
6. entitlement never grants content contribution/review/release/quarantine/admin/security capabilities, and privileged operational capability never implies paid learner entitlement;
7. normal access to applicable learner history plus account data export/deletion controls is governed by learner-data/product/privacy policy, not withheld merely because premium capability access expired;
8. rate/cost limits may bound an entitled optional capability, but exhaustion/degradation is product availability rather than learner weakness and cannot lower evaluator/evidence quality.

# Legal lifecycle state machines

Exact wire fields belong to machine contracts. Legal semantic transitions are owned here so services cannot invent incompatible lifecycle rules.

The state machines below apply to learner/runtime work resources. Content validation, release eligibility, and operational safety intentionally remain separate dimensions governed by Flows K–L rather than one combined lifecycle enum.

## LearningSession

```text
planned → in_progress → completed
   │          └────────→ abandoned
   └───────────────────→ abandoned
```

`completed` and `abandoned` are terminal for that session identity.

## Attempt

```text
draft → submitted ↔ evaluating → evaluated
                  ├────────────→ evaluated   deterministic path
                  └────────────→ invalid
submitted ─────────────────────→ invalid
```

Rules:

- draft response content may mutate only before submission under concurrency/revision rules;
- the referenced assigned content revision does not mutate for that Attempt identity;
- submission is idempotent;
- submitted work is immutable for that Attempt identity;
- `evaluating` is current processing state, not learner validity; if no current eligible execution is active while accepted work remains pending/retriable, processing may return to `submitted` rather than becoming `invalid` or remaining falsely stuck;
- content incident may narrow/block current evaluation/evidence consequence without rewriting submitted learner work;
- redraft/correction creates a new version/related attempt rather than rewriting submitted history;
- evaluator/infrastructure failure is not `invalid` learner performance and does not require another learner Attempt.

## Evaluation

```text
pending → running → completed
    │         ├────→ unavailable
    │         └────→ invalid
    ├──────────────→ unavailable
    └──────────────→ invalid
```

The Evaluation identity is the logical evaluation resource and may have distinct execution attempts beneath it when retry/replacement can overlap. Its terminal states apply to the Evaluation identity and never silently reopen.

Rules:

- one execution-attempt timeout is ambiguous and does not by itself prove the logical Evaluation failed or that the remote attempt did no work;
- one exclusive execution attempt is claimed/fenced before normal dispatch; multiple workers observing pending work do not acquire the same execution merely by reading it;
- loss of the worker/claim owner does not itself make the Evaluation terminal or prove dispatch did not occur; recovery follows current eligibility plus ambiguous-outcome/idempotency/fencing rules;
- Core accepts a completion only when it matches this Evaluation/logical work, the current legal execution/fencing state, and current deletion/content eligibility relevant to reconciliation;
- duplicate delivery of the already accepted completion is idempotent and cannot duplicate Observation/EvidenceFact creation;
- a late/stale completion from a superseded execution is non-authoritative for learner state and may be retained only as operational diagnostic/provenance;
- `completed` means the logical Evaluation has an accepted durable outcome under its lifecycle; when required downstream semantic artifacts are not materialized in the same commit, recoverable continuation/reconstruction remains durably guaranteed and is not optional cleanup;
- if an Evaluation becomes terminal `unavailable` or `invalid` and product policy later permits another evaluation, create another Evaluation identity linked to the same accepted Attempt/logical obligation rather than reopening the terminal identity or creating another learner Attempt;
- linked Attempt processing reflects whether eligible evaluation work remains pending/running/accepted; evaluator infrastructure failure alone does not force Attempt `invalid`.

## DiagnosticRun

```text
created → in_progress → completed
                    ├→ abandoned
                    └→ unavailable
```

`completed` means the sampling flow ended, not that all learner claims or target fields are known.

## MockRun

```text
created → in_progress → completed
                    ├→ abandoned
                    └→ invalid
```

Partial valid section observations preserve their actual scope; they cannot masquerade as a complete mock result.

## Media resolution/analysis

```text
requested → resolving/analyzing → ready
                         ├──────→ ineligible
                         └──────→ unavailable
```

`ineligible` is a source/rights/product condition, never learner failure.

## BandCertificationState

Per skill/Band, current state follows `../spec/09-PROGRESSION.md`:

```text
not_started → in_progress ↔ certified
```

`certified` is valid only while the corresponding current claim is `SUPPORTED`. If the current claim becomes stale, conflicting, insufficient, below requirement, or otherwise non-`SUPPORTED`, current state returns to `in_progress` while certification history remains intact.

The transition reason preserves meaning. Staleness/conflict/insufficient evidence do **not** establish regression; regression exists only when Progression's regression rule is satisfied by later admissible evidence. Once a claim has evidence/history, it does not return to `not_started` merely because current support is lost.

# Transition invariants

- invalid transitions fail closed;
- retry of an idempotent operation reuses the same logical result/identity where applicable;
- one service does not advance another owner’s lifecycle by direct database mutation;
- consequential terminal transitions record reconstructable reason/provenance;
- runtime completion state never substitutes for learning state;
- upstream work completion with required downstream semantics carries durable recoverability until those semantics are materialized;
- content operational state changes never rewrite historical content revision or validation-decision identity.

# Result delivery

Immediate deterministic work may return synchronously.

Long work uses the semantic pattern:

```text
submit / durable logical work
        ↓
execution attempt claimed/fenced
        ↓
dispatch attempted
        ↓
remote completion received when any
        ↓
completion accepted/rejected by Core
        ↓
required Observation/Assessment/Progression continuation durably materialized or recoverable
        ↓
learner-visible resource state
        ↓
SSE update hint
```

These points are not synonyms: dispatch success is not semantic completion; HTTP `200` from Python is not a learner-state transition; Evaluation `completed` does not imply every downstream derived artifact was already written in the same transaction, but it does require durable recoverability of required continuation; SSE delivery is only a notification hint. Polling/resource refresh remains fallback and persistent truth remains queryable independently of SSE delivery.

# Failure semantics

Keep domain/operational outcomes distinct where applicable, including:

```text
invalid_attempt
evaluation_pending
evaluation_unavailable
insufficient_evidence
conflicting_evidence
stale_evidence
source_unavailable
source_ineligible
target_not_supported
product_coverage_blocked
```

These labels do not imply one transport-error category. Pending, evidence states, target/product support outcomes, and content assignment outcomes can be valid domain results; transport/operation failure classification belongs to `05-api.md` and future machine contracts.

Content generation/validation/assignment failures should likewise remain domain/operational states rather than fake learner failure.

Infrastructure failure is never represented as score zero or generic learner failure. An unresolved target field should use target-resolution semantics rather than pretending the product rejected a fully specified target.

# Implementation-boundary reference

Runtime allocation, co-location, framework/tooling, database access ownership, and cross-runtime implementation boundaries are owned by `06-implementation-stack.md`. These flows require implementations to preserve their declared authority and exact machine contracts; no runtime may recreate or bypass the product/learning semantics owned here and upstream.