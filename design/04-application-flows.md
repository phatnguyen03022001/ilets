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

Logical ownership survives deployment co-location. Running Web, Core API, and Evaluator on one machine/platform does not authorize browser→Python bypass, Python mutation of Core-owned durable state, hidden shared-memory product state, or omission of the declared machine boundary between independently owned runtime units.

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
idempotency + concurrency protection where applicable
        ↓
authoritative transaction
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
establish one logical work identity
  ↓
authoritative transaction
  ├─ persist accepted product mutation/work state
  └─ persist durable pending-work/outbox/recoverable marker
  ↓
COMMIT
  ↓
acknowledge accepted/pending
  ↓
bounded asynchronous dispatch/execution
  ↓
bounded eligible retry/backoff
  ↓
persist result / unavailable / unresolved failure state
  ↓
SSE update hint and/or resource refresh
```

Provider webhooks use their own provider-authentication/replay semantics rather than learner/admin authentication, but still require validation and authoritative association before mutation.

Rules:

- an HTTP request being sent or a provider invocation being attempted is not durable work acceptance;
- retry preserves one logical work identity and cannot duplicate accepted learner work, EvidenceFacts, content revisions, or paid/provider work;
- retry eligibility distinguishes transient, permanent, and ambiguous outcomes;
- exponential backoff and jitter are used where repeated immediate retry would worsen a transient route; exact budgets/counts remain implementation policy;
- a timeout does not prove the remote operation failed, so ambiguous outcomes remain unresolved until authoritative state can be established safely;
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
Core API validates contract + work identity
        ↓
Assessment / content / product policy interprets the result
```

HTTP/network success from Python establishes only that a bounded capability response arrived. It does not make evaluator/generator/validator output authoritative learner evidence, certification, content activation, or product support. Evaluator execution must not mutate Core-API-owned product persistence directly.

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

SSE is delivery, not state authority. Persistent current state remains queryable through resources, duplicate/reordered hints are tolerated, and reconnect cannot fabricate a transition.

## F. Scheduled/background work

Scheduled, cron-like, batch, maintenance, or background work is allowed when a real product/operational need exists; no cron platform is required merely by this design.

When used:

- duplicate-sensitive execution has one logical work identity;
- required work is durably registered/recoverably discoverable before another committed operation relies on it;
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
- wall-clock timestamps may support audit/recency but do not by themselves establish global causal ordering, idempotency identity, or concurrency correctness;
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
- cache failure falls back to authoritative state where operationally feasible and never changes truth;
- telemetry failure cannot rewrite business truth; only an explicitly required security/audit invariant may make absence of required durable audit block a consequential operation;
- co-located process failure does not permit another unit to seize or mutate that unit's semantic authority through undocumented in-memory coupling.

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

A plan records enough state/target/content-revision references and reason information to reconstruct why every activity was eligible and selected, including which target conditions remained unresolved when the plan was produced.

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
coherent DailyPlan + reason codes
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

Swap/Skip/Shorten/Change-skill actions operate within eligible choices. They do not mark skipped requirements satisfied.

# Flow C — ordinary practice attempt

```text
1. Web starts PracticeActivity
2. Core API returns exact item revision + target + variant/context + conditions
   + primary_activity_purpose + evidence_candidacy
3. learner performs task
4. Web submits Attempt with actual scaffold/exposure/delivery metadata
5. Core API or Evaluator produces Observation
6. NOT_EVIDENCE_CANDIDATE → feedback/remediation/review handling only
   ASSESSMENT_MAY_ADMIT  → Assessment evaluates actual claim-scoped eligibility/inference
7. Progression updates interpretation only when justified
8. Planner derives next target-relevant action
9. Web presents outcome without fabricating a micro Band change
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
  └─ persist evaluation work identity + pending/recoverable dispatch marker
  ↓ COMMIT
ACK accepted/pending
  ↓
bounded dispatch / Evaluator execution
  ↓
criterion observations + provenance + uncertainty
  ↓
Core API validates evaluator response
  ↓
Assessment → EvidenceFact / ReadinessEvaluation
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
- timeout/provider failure never becomes a low learner score;
- retries preserve one logical work/evaluation identity and provenance;
- no progression occurs from failed/invalid evaluator output;
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
learner-specific exposure/novelty/independence gate
        ↓
exact revision assignment
        ↓
Attempt references that revision
```

Rules:

1. generation is optional; a release may be fully supplied by authored/imported/deterministic assets when they satisfy coverage and operational requirements;
2. generation/import is requested only when actual content demand remains after eligible pool reuse, or may be pre-generated/batched from demonstrated future demand so learner requests do not unnecessarily wait on generation;
3. reuse optimization must never override target/family/context/presentation/difficulty/diversity/rights/evidence requirements;
4. content validation and assignment are separate: a globally valid revision can still be ineligible for a specific learner/use because of exposure, independence, or novelty requirements;
5. corpus similarity is not a universal rejection rule; similarity facts are interpreted against the intended learning/evidence purpose;
6. generator output and validator signals cannot activate content by themselves; Core-API-owned deterministic product policy applies the owning semantic/coverage rules;
7. higher-consequence use requires the stronger applicable validation/evidence conditions defined upstream; low-consequence training does not require unrelated high-consequence checks;
8. every assigned activity resolves the exact immutable revision before learner exposure.

Content supply may be asynchronous. If the desired content is unavailable, the planner may use another genuinely eligible activity or expose an honest product/content gap; it must not silently relax semantic or evidence requirements merely to avoid generation latency/cost.

# Flow L — content report, quarantine, revalidation, and retirement

Content operational safety is independent from immutable revision identity and from whether a revision previously passed validation.

Representative flow:

```text
Attempt / learner report / operational signal / validator-policy regression
        ↓
triage material risk
        ↓
stop new assignment when warranted
        ↓
preserve historical revision + Attempt/evidence references
        ↓
investigate / revalidate
        ├─ same semantic revision verified → restore eligible assignment when release permits
        ├─ material content fix required → create new ContentRevision → validate → activate when eligible
        └─ no safe/valid route → retire from new assignment
                                    ↓
                         replacement/supply demand if coverage needs it
```

Do not collapse validation state, release eligibility, and operational safety into one global ContentStatus enum. A revision may be semantically validated yet not activated for a release, or may have been active and later be quarantined operationally pending investigation.

Operational actors exercise only applicable privileged capabilities defined below. They are operators, not learning or Assessment authorities.

No authorized operator/admin action, regardless of role level, may bypass a known applicable hard failure such as:

- invalid canonical reference or incompatible IELTS family/context/presentation binding;
- known incorrect answer key/rubric/model where required;
- prohibited answer leakage for the intended use;
- rights/privacy/security failure;
- evidence-critical exposure/independence failure for a use that requires independent evidence.

The cause must be repaired or the intended use changed legitimately, then the content/revision must pass the applicable policy again. A generic `override_validation` capability that converts a known hard failure into active learner content is forbidden.

Historical learner facts are never repaired by mutating an old revision. If a later discovery changes how historical evidence should be interpreted, preserve the original Attempt/Observation and apply the owning Assessment/Progression policy with explicit provenance.

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
3. an operational capability grant is not permission to mutate historical ContentRevision semantic payload;
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
draft → submitted → evaluating → evaluated
                  ├────────────→ evaluated   deterministic path
                  └────────────→ invalid
submitted ─────────────────────→ invalid
```

Rules:

- draft response content may mutate only before submission under concurrency/revision rules;
- the referenced assigned content revision does not mutate for that Attempt identity;
- submission is idempotent;
- submitted work is immutable for that Attempt identity;
- redraft/correction creates a new version/related attempt rather than rewriting submitted history;
- infrastructure failure is not `invalid` learner performance.

## Evaluation

```text
pending → running → completed
    │         ├────→ unavailable
    │         └────→ invalid
    ├──────────────→ unavailable
    └──────────────→ invalid
```

Terminal state applies to the evaluation identity. Retry preserves logical-work linkage and must not duplicate EvidenceFacts.

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
- content operational state changes never rewrite historical content revision identity.

# Result delivery

Immediate deterministic work may return synchronously.

Long work uses the semantic pattern:

```text
submit → durable accepted/pending state
             ↓
       async dispatch/work
             ↓
       SSE status/result
```

Polling/resource refresh remains fallback. Persistent truth remains queryable independently of SSE delivery.

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

# Runtime ownership invariant

```text
TypeScript
  browser interaction/rendering/capture
  authorized admin/operations presentation

Go
  public API + durable product state
  deterministic scoring
  Assessment policy execution
  Progression execution
  Planner orchestration
  content demand/reuse/assignment + operational eligibility orchestration

Python
  bounded AI/audio/text evaluation and media analysis
  bounded content-generation/validation capability only when invoked through declared internal contracts
```

Cross-language data crosses explicit contracts. The same semantic rule is not independently maintained in multiple runtimes. Python or another generator/validator may produce candidates/signals; Core API remains responsible for applying product policy and preserving authoritative content/assignment state.