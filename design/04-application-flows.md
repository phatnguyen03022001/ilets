STATUS: CANONICAL
OWNS: end-to-end product/system flows across web, core API, evaluator, learner state, target route, media, content supply/assignment and content-incident recovery, privileged content-operation capability semantics, runtime execution/trust/failure patterns, async result delivery, planner-stage separation, hard eligibility, and legal runtime lifecycle semantics
DEPENDS_ON: ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 01-skill-features.md, 02-practice-catalog.md, 03-media-youtube.md
DOES_NOT_OWN: API field schemas, learning/mastery truth, content semantic identity/quality truth, product coverage declaration, identity-provider implementation, concrete authorization role matrix, exact persistence topology, provider selection, framework internals, deployment technology selection, or learner-facing UX defaults

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
- exact content-revision release/operational eligibility;
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
- operational cost after semantic validity is preserved.

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

A DailyPlan is a snapshot/recommendation produced from named state, not assignment authority. It records enough provenance to reconstruct why each activity was eligible and selected at plan time, including where material the TargetProfile/target-context revision, learner/Progression state reference/version, content/release state, product support/coverage version, content-revision references, and unresolved conditions used by planning.

Before actual PracticeActivity/AssessmentActivity assignment or learner exposure, Core re-evaluates every **current** hard condition that can change. This includes where material current target context, learner/evidence state, product coverage/support scope, ContentRevision release and quarantine/operational state, rights/source eligibility, learner exposure/novelty/independence and reservation state, and delivery/capture feasibility.

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
ranking
  ↓
coherent DailyPlan + reason codes + plan-time provenance
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
```

Direct `wrong answer → fixed exercise id` mapping is not a canonical remediation policy. Detectable secondary issues may be recorded/deferred rather than surfaced when they are outside the current feedback focus.

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

A mixed Academic/GT mock is invalid for normal full-test readiness unless explicitly created as non-certifying comparison practice.

Delivery-mode practice may change interaction conditions without changing scoring/Band semantics.

# Flow J — target supported/unresolved

For every known TargetProfile condition:

```text
current admissible evidence
  ↓
SUPPORTED | unresolved evidence state
```

Target-input uncertainty and learner-evidence uncertainty remain separate. A missing target variant/Band condition is not converted into `INSUFFICIENT_EVIDENCE` about learner capability.

When all learner evidence conditions are supported, the app may state that current evidence supports the declared target profile **only if the target itself is sufficiently resolved and product-support wording remains separately truthful**.

It must never state that the learner is guaranteed an external result.

When product capability is missing for a resolved requested scope, route generation stops at the CoverageGap rather than manufacturing an invalid activity.

# Flow K — content supply and learner assignment

Content supply is reuse-first and mechanism-neutral.

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
8. every assigned activity resolves the exact immutable revision and current release/operational eligibility before learner exposure, regardless of whether an older DailyPlan referenced it;
9. the decisive assignment enforces the mutable current hard gates described in Stage 7 under a transaction, conditional write, or equivalent serialization invariant where concurrency can invalidate them; learner-specific unseen/exposure/independence/uniqueness and reservation state are included when material so concurrent requests cannot both consume the same protected opportunity incorrectly;
10. reservation/assignment for delivery is not actual learner exposure. A failed/disconnected delivery must not fabricate `seen`; actual ExposureContext follows `spec/10-CONTENT-MODEL.md`;
11. a reservation may temporarily exclude concurrent assignment until it is reconciled/released; exact reservation timeout/recovery mechanism remains implementation policy.

Content supply may be asynchronous. If the desired content is unavailable, the planner may use another genuinely eligible activity or expose an honest product/content gap; it must not silently relax semantic or evidence requirements merely to avoid generation latency/cost.

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

Privileged content operations are capability-scoped rather than defined by a canonical role taxonomy. The product semantics distinguish at least these operational capabilities:

- inspect content, provenance, validation evidence, and operational state;
- stop new assignment / quarantine content when warranted;
- activate or release content that already satisfies every applicable semantic, validation, release, and operational condition;
- retire content from new assignment;
- request content supply, regeneration, or replacement;
- request revalidation;
- resolve content reports with reconstructable operational reason/provenance.

Invariants:

1. an operational capability grant is not learning authority;
2. an operational capability grant is not Assessment/evidence authority;
3. an operational capability grant is not permission to mutate an established ContentRevision semantic payload;
4. an operational capability grant is not validation-bypass authority;
5. no capability may convert a known applicable hard failure into assignable content;
6. activation/release capability may act only on content that already satisfies the applicable semantic, validation, release, and operational policy;
7. authorization implementation may later map authenticated identities/roles to these capabilities without redefining their meanings.

Concrete role names, role hierarchy, identity-provider integration, and the role-to-capability matrix remain implementation/authorization concerns. API operations consume these capability meanings through `05-api.md` without creating a second authorization taxonomy.

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