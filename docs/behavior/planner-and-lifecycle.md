# PLANNER-LIFECYCLE Planner, assignment, flow, and lifecycle semantics

## Seven-stage Planner decision contract

Planning is staged; later stages never redefine truth produced upstream:

```text
1. resolve TargetProfile + product-support scope
2. expand unresolved target conditions
3. consume Assessment/Progression → GapEvaluation / ActionIntent
4. hard-eligibility filtering
5. activity-candidate generation
6. ranking among eligible candidates
7. coherent plan + stable explanation
```

### Stage 1 — target/support resolution

Target-relative planning/support is fully resolved only when the standard variant is known and at least one real Band constraint is known. Missing learner target input is an unresolved target condition, not a CoverageGap. A known resolved target the product cannot serve is a CoverageGap. While variant/all Band constraints remain unresolved, only genuinely shared/foundational/diagnostic actions that do not require the missing condition may be offered; Core cannot invent Academic/GT, a Band target, or complete support/readiness. Delivery mode, purpose and One-Skill-Retake conditions remain explicit where material.

### Stage 2 — target-condition expansion

Expand known target constraints into per-skill/variant/external-purpose conditions while preserving unresolved input. An overall Band may have multiple valid four-skill combinations; Planner cannot manufacture hidden per-skill minima or promote a non-authoritative working profile into target truth.

### Stage 3 — evidence interpretation

Consume Assessment and Progression outputs. Planner never rescores Observations, reinterprets evaluator output, broadens inference, or creates another GapEvaluation taxonomy.

### Stage 4 — hard eligibility

Remove candidates violating any applicable hard condition: Required prerequisite; target variant/task compatibility; material delivery-mode compatibility; product coverage/support; primary-purpose/evidence-candidacy compatibility; exact current ContentRevision validation/release/operational eligibility for the intended use; learner exposure/novelty/independence; rights/privacy/source; accessibility/capture feasibility; immutable lifecycle; or known purpose/acceptance restrictions. A ranker cannot rescue an ineligible candidate.

Examples remain invariant: Academic Task-1 work cannot serve GT Task-1 readiness; variant-specific work cannot be target-relative while variant is unresolved; typing-only work cannot satisfy required handwriting readiness; IELTS Online Academic cannot serve GT/visa-purpose constraints it does not support; failed capture cannot become Speaking evidence; `NOT_EVIDENCE_CANDIDATE` cannot satisfy `COLLECT_EVIDENCE`; prior exposure may invalidate an otherwise globally valid readiness item.

### Stage 5 — candidate generation

Construct valid learning/assessment activity candidates for the current ActionIntent. Candidate generation is not necessarily content generation. Delivery may vary but target standard may not. Actual activity must resolve to an eligible exact ContentRevision before assignment; reuse eligible existing content before supply generation/import. Unresolved target conditions restrict generation to actions that do not invent them.

### Stage 6 — ranking, load, and trajectory

Ranking considers eligible candidates only. Permitted signals include target/test-date urgency, expected learning/decision value, due review, time fit, preference/friction, fatigue/session coherence, transfer/exposure diversity, operational cost after semantic validity, and attributable intervention-effectiveness history where sufficiently comparable.

Intervention-effectiveness signal preserves Practice semantics: insufficient/incomparable history remains unresolved; effective repetition may continue; justified diminishing returns may vary an existing mechanism/type/scaffold/context/difficulty/load/frequency/sequence; diagnosis uncertainty favors discriminating evidence; technical/content/evaluator/provider failure is not learner plateau. Target, prerequisites, evidence policy and history remain unchanged.

Load composition decides how much eligible work fits a coherent plan; inventory/ranking is not dosage. It may combine due review, prerequisites, focused language work, transfer, receptive/productive practice, remediation, re-evidence and exam preparation using target urgency, evidence/gap/intent, prerequisites, due state, available time, duration/load estimates, recent history, fatigue/coherence, diversity/re-evidence need and available capabilities/cost. Band thresholds do not prescribe daily counts; numeric counts/load maxima/coefficients are empirical versioned policy, not architecture constants. Required/due work deferred for coherent load remains visible rather than marked satisfied.

Trajectory advisory is non-authoritative. With enough calibrated basis it may report whether the current study envelope appears plausibly sufficient for the declared target/date. Missing history/calibration/evidence yields unresolved advisory, not fabricated velocity. Product inability remains separate from learner pace. Projection never lowers target, invents per-skill minima, changes test date, certifies readiness or alters eligibility. It normally preserves `unresolved`, `at risk`, or `current envelope insufficient`; causal inputs/policy/calibration remain reconstructable.

Ranker non-authority: it may reorder eligible candidates but never make an ineligible candidate eligible, bypass prerequisites, alter Gap/ReadinessEvaluation, reinterpret evaluator output, hide CoverageGap, invent missing target constraints, lower thresholds, convert preference to ability evidence, ignore variant/delivery, override content ineligibility, upgrade evidence candidacy/admission, or certify.

### Stage 7 — plan/explanation and assignment boundary

DailyPlan is a recommendation snapshot, not assignment authority. It preserves enough target-context, learner/progression, due-review, policy/load, time/session, content/release, support/coverage, exact revision and unresolved-condition provenance to explain selection/amount; trajectory/effectiveness consequences additionally preserve their policy/comparison basis.

Before actual PracticeActivity/AssessmentActivity assignment or exposure, Core re-evaluates every mutable current hard condition. The boundary is:

```text
plan/candidate when present
→ current hard-eligibility re-evaluation
→ decisive reservation/assignment where applicable
→ learner-safe projection
→ delivery / actual exposure
```

Where a mutable authoritative condition can concurrently invalidate assignment, the decisive current-state check is atomically/conditionally coupled to reservation/assignment or equivalent serialization. A stale plan or preflight check cannot override a concurrent quarantine, release revocation, target/support change, or learner-specific reservation/exposure change. Now-ineligible plan items are reselected or reported truthfully. Assignment provenance preserves exact revision, intended-use scope, applicable validation-policy version and compatible validation/release state where material. Plan provenance and assignment authority are distinct; tie-breaking is stable enough to avoid arbitrary churn.

## Material end-to-end flows A–L

### Flow A — target setup and diagnostic

Web collects known target/planning fields; Core validates known variant/delivery/purpose combinations without inventing missing inputs, persists target revision + unresolved conditions, distinguishes CoverageGap from missing input, selects only compatible diagnostic items, scores deterministic L/R where valid and uses Evaluator for eligible W/S, then Assessment interprets sampling/evidence, Progression derives justified state, Planner runs stages 1–7, and Web exposes sampled/unresolved/evidence/product-coverage states separately. Completed diagnostic is not a complete learner model, resolved target or certification.

### Flow B — Daily Plan

Known target → resolved/unresolved target conditions → product support where evaluable → GapEvaluation/ActionIntent → hard eligibility → valid candidates → ranking + load composition → coherent DailyPlan/reasons/provenance → trajectory advisory only when justified. Swap/Skip/Shorten/Change-skill stay inside eligible choices and do not satisfy skipped requirements; start still crosses assignment-time eligibility.

### Flow C — ordinary practice attempt

Core rechecks and decisively assigns an exact eligible revision, projecting target/variant/context/conditions + primary purpose + evidence candidacy. Learner Attempt records actual scaffold/exposure/delivery; Core/Evaluator produces Observation. `NOT_EVIDENCE_CANDIDATE` stays feedback/remediation/review-only; `ASSESSMENT_MAY_ADMIT` only permits Assessment to evaluate actual claim-scoped eligibility. Progression changes only when justified and Planner derives next action. Favorable result cannot retroactively upgrade candidacy; network retry is idempotent; content lifecycle changes never rewrite an Attempt's referenced revision.

### Flow D — Writing/Speaking asynchronous evaluation and optional realtime interaction

Submission first crosses auth/access, structural validation and semantic preconditions. One authoritative transaction persists Attempt/submission, logical evaluation identity and durable pending/recoverable dispatch state before accepted/pending acknowledgement. A current-eligible execution attempt is then claimed/fenced, bounded dispatch occurs, evaluator output returns criterion observations/provenance/uncertainty, Core validates contract/work/execution association and current deletion/content eligibility, accepts exactly one legal durable result plus Observation or recoverable continuation, then Assessment/Progression/Planner continue and SSE/resource refresh delivers status.

Timeout/provider failure is not a low score and does not prove no remote work occurred. Competing workers cannot silently double-dispatch exclusive work; replacements have distinct fenced execution identity; claim loss cannot strand accepted work; duplicate accepted completion is idempotent; stale/superseded completion cannot overwrite current result or independently create Observation/EvidenceFact; accepted result must leave required downstream semantics materialized or durably reconstructable. No progression uses failed/stale/superseded/deleted/ineligible output; fallback meets the same floor.

Ordinary Speaking recording is draft until submission; submitted Attempt is immutable and re-record creates a related Attempt. Optional realtime Speaking preserves one logical session, target/part, assistance/evidence configuration, entitlement admission and turn/capture/provider provenance. Part 2 preserves preparation + independent long turn; technical silence/latency/network failure is not learner performance. Reconnect fences/resumes the same accepted session where safe; partial sessions support only actual scope; record→submit fallback is allowed only if semantics remain valid and cannot satisfy a realtime/readiness condition that required interaction. Realtime completion/provider accounting does not admit EvidenceFacts.

### Flow E — objective Listening/Reading attempt

Exact revision answer key + instruction/word-limit validation → raw result → Observation with variant/context/conditions → Assessment eligibility/inference where candidate → EvidenceFact when valid. Reading Band inference uses the correct Academic/GT scoring/context policy.

### Flow F — feedback/remediation

Observation → GapEvaluation → ActionIntent → feedback-focus policy → Error/RemediationPattern where relevant → Learning Mechanism → Practice Mode → fresh eligible PracticeItem revision → learner Attempt/outcome/separate re-evidence → comparable intervention history when sufficient → continue, vary strategy or collect discriminating evidence. `wrong answer → fixed exercise id` is not canonical; secondary issues may be deferred; same Gap must not force endless identical remediation; product/runtime failure is not plateau.

### Flow G — review

Review request resolves queue kind independently as knowledge retrieval, error remediation, or re-evidence; uses an eligible exact revision; resulting Attempt updates only the corresponding semantic state. One UI can combine queues without collapsing backend meaning.

### Flow H — media lesson creation

Learner source URL/reference is untrusted. Core resolves permitted metadata/source eligibility, establishes rights/transcript state, sends only authorized/minimum necessary text to Evaluator where applicable, receives proposed segments/targets/difficulty/prompts, validates rights + canonical target + practice mapping, admits resulting content only through normal validation/revision, then permits preview/save and normal Practice/Attempt use. URL/evaluator access never grants arbitrary download/scrape rights.

### Flow I — mock/readiness

Resolved variant + material delivery target → MockRun → Listening shared + variant-correct Reading + variant-correct Writing Task 1 + Writing Task 2 + Speaking shared → scoped observations → ReadinessEvaluation only where candidacy/eligibility permit → GapEvaluation → exam-preparation plan. `READINESS` purpose is not evidence admission. One favorable mock or separate micro/per-skill claims do not bypass an integrated EvidenceRequirement. Mock composition preserves coherent variant, exact components/provenance, delivery and learner-specific independence. Mixed Academic/GT is invalid for normal full-test readiness.

### Flow J — target supported/unresolved

Keep four axes distinct: target input resolution; learner evidence support; material exam/delivery readiness; product support. Explicit per-skill minima require corresponding supported current Band claims; overall-only target requires supported current Bands for all four skills then official overall calculation without invented minima; mixed constraints require both; variant claims use correct variant evidence/conversion/context; delivery is separate; full IELTS readiness additionally requires integrated/exam-condition EvidenceRequirement. Missing target input is not learner evidence insufficiency. Product inability preserves learner claim unresolved and exposes CoverageGap, never ABILITY_GAP/fake score/hidden-human dependency. Support is never an external-result guarantee.

### Flow K — content supply and learner assignment

ContentDemand describes the eligible opportunity needed across only applicable canonical targets, node/phase, Practice/Assessment Type, purpose/candidacy, variant/context/family/presentation, task/section/part/integrated scope, difficulty/scaffold/transfer, interaction/scoring/evaluator, delivery, rights, learner exposure/novelty/independence and policy-required quantity/diversity. Reuse eligible supply first; only actual residual demand justifies authoring/import/deterministic/AI candidate supply. Candidate becomes immutable ContentRevision, passes applicable validation/release, then current learner/use hard gates and decisive reservation/assignment before delivery; actual exposure is recorded only when established; Attempt references the exact revision.

Generation is optional; reuse never overrides semantics/rights/evidence. Validation and learner-specific assignment are separate. Generator/validator output cannot self-activate content. Higher consequence uses stronger gates. `UNKNOWN` exposure is not unseen. Reservation is not exposure. No canonical practice/assessment/readiness/unseen pools are required. Supply sufficiency is demand/diversity scoped, not raw count. All source mechanisms use the same content lifecycle. If no eligible supply exists, choose another genuinely eligible activity or expose a truthful gap, never relax requirements for latency/cost.

Parallel AI authoring is deterministic-shard based, using pre-existing canonical IDs and non-overlapping shard/slot identity. Workers cannot invent new semantic IDs; duplicate identity, unknown refs, overlap or ambiguous integration fail closed. AI claims about difficulty/Band/scoring/evidence are candidate metadata until owning policy validates them. Multiple models/threads do not create independent validation authority merely by multiplicity. Candidate packages are authoring input, not assignment authority or learner evidence; no orchestration platform is required without measured need.

### Flow L — content incident, quarantine, revalidation, retirement

Report/signal/policy regression → triage risk/use scope → new ValidationDecision for revalidation of unchanged content where applicable → recompute current release/use eligibility → preserve historical ContentRevision/ValidationDecisions/learner facts → apply current consequence to reservations/in-flight/evidence → investigate/revalidate/replace. Unchanged verified revision may become assignable again; material fix creates a new revision; no safe route retires new assignment and may trigger replacement demand.

Do not collapse validation, release, operational safety and incident consequence into one global status. Current use disposition may be assignable, quarantined, superseded, retired-from-new-assignment or unavailable/ineligible, and may differ by use. Consequence is stage/scoped: block new assignment, reconcile unexposed reservation, preserve but stop/replace affected exposed draft where needed, preserve submitted Attempt while blocking/narrowing/re-routing invalid inference, and preserve historical Observation/EvidenceFact while Assessment/Progression re-evaluates current support if historical inference changes. No privileged role may override known hard semantic/scoring/leakage/rights/privacy/security/independence failure or mutate historical content/learner facts.

## Legal runtime lifecycle state machines

Exact wire encoding remains machine-contract authority; these are the legal semantic transitions.

### LearningSession

```text
planned → in_progress → completed
   │          └────────→ abandoned
   └───────────────────→ abandoned
```

`completed` and `abandoned` are terminal for that session identity.

### Attempt

```text
draft → submitted ↔ evaluating → evaluated
                  ├────────────→ evaluated   deterministic path
                  └────────────→ invalid
submitted ─────────────────────→ invalid
```

Draft response can mutate only before submission. Assigned content revision is fixed for the Attempt. Submission is idempotent and submitted work immutable. `evaluating` is processing state, not learner validity; accepted retriable work may return to `submitted` when no eligible execution is active. Content incident may narrow evidence without rewriting submission. Redraft creates related Attempt. Evaluator/infrastructure failure alone is not `invalid` learner performance.

### Evaluation

```text
pending → running → completed
    │         ├────→ unavailable
    │         └────→ invalid
    ├──────────────→ unavailable
    └──────────────→ invalid
```

Evaluation is one logical resource with distinct execution attempts beneath retry/replacement. Terminal states do not reopen. Timeout is ambiguous, exclusive attempt is claimed/fenced, owner loss does not imply no dispatch, Core accepts only current legal associated completion, duplicate accepted delivery is idempotent, stale superseded completion is non-authoritative, and `completed` requires accepted durable result plus any required semantic continuation already committed or durably reconstructable. Later policy-permitted re-evaluation creates a new Evaluation linked to the same accepted Attempt rather than reopening terminal identity or creating a learner Attempt.

### DiagnosticRun

```text
created → in_progress → completed
                    ├→ abandoned
                    └→ unavailable
```

Completion means sampling ended, not that target/model is fully known.

### MockRun

```text
created → in_progress → completed
                    ├→ abandoned
                    └→ invalid
```

Partial valid section observations retain only their actual scope.

### Media resolution/analysis

```text
requested → resolving/analyzing → ready
                         ├──────→ ineligible
                         └──────→ unavailable
```

`ineligible` is source/rights/product state, never learner failure.

### Realtime Speaking interaction

```text
created → active ↔ reconnecting → completed
            ├───────────────→ abandoned
            └───────────────→ unavailable
reconnecting ├──────────────→ abandoned
             └──────────────→ unavailable
```

Completion means interaction ended successfully, not that Speaking is supported.

### BandCertificationState

```text
not_started → in_progress ↔ certified
```

`certified` is valid only while the current claim is SUPPORTED. Any non-SUPPORTED current claim returns current state to `in_progress` while history remains. Staleness/conflict/insufficiency are not regression; regression requires later admissible below-requirement evidence. Once meaningful evidence/history exists, current state does not return to `not_started` merely because support is lost.

## Transition, result and failure invariants

Invalid transitions fail closed. Idempotent retry reuses logical identity/result. One service never advances another owner's lifecycle by direct DB mutation. Consequential terminal transitions preserve reason/provenance. Runtime completion never substitutes for learning state. Upstream completion with required downstream semantics stays recoverable until materialized. Content operational changes never rewrite historical revision or ValidationDecision identity.

For long work: durable logical work → execution claim/fence → dispatch → remote completion if any → Core acceptance/rejection → required Observation/Assessment/Progression continuation materialized or recoverable → learner-visible resource state → SSE update hint. Dispatch/HTTP success/SSE are not semantic completion or learner-state authority; persistent resource state remains independently queryable.

Keep failures distinct: `invalid_attempt`, `evaluation_pending`, `evaluation_unavailable`, `insufficient_evidence`, `conflicting_evidence`, `stale_evidence`, `source_unavailable`, `source_ineligible`, `target_not_supported`, `product_coverage_blocked`. Infrastructure/content/provider failure never becomes score zero or generic learner failure, and unresolved target input uses target-resolution semantics rather than a fake rejection.