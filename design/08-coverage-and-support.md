STATUS: CANONICAL
OWNS: product coverage semantics, TargetCoverageSpecification, CoverageGap taxonomy, support-promotion gates, construct-to-product coverage mapping, runtime-concern coverage integration, and current product-support declarations
DEPENDS_ON: ../spec/00-PRODUCT.md, ../spec/02-IELTS-MODEL.md, ../spec/03-SKILLS.md, ../spec/05-BANDS.md, ../spec/06-CURRICULUM.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/10-CONTENT-MODEL.md, 00-learning-experience.md, 01-skill-features.md, 02-practice-catalog.md, 03-media-youtube.md, 04-application-flows.md, 05-api.md, 06-implementation-stack.md, 07-third-party-services.md
DOES_NOT_OWN: external IELTS truth, learning targets, feature behavior, content semantic identity/quality rules, provider selection, concrete infrastructure technology selection, learner GapEvaluation, or empirical validation results themselves

# Coverage and Product Support

## Purpose

Define exactly when a scoped IELTS target is merely modelled, actually executable, release-supported, or empirically validated.

Document volume, feature count, Skill coverage, an aggregate percentage, or the presence of fashionable infrastructure cannot hide a missing required condition.

# Product-support status machine

A **scoped product target** uses only:

```text
MODELLED
  ↓
COVERED
  ↓
SUPPORTED_FOR_PRODUCT
  ↓
VALIDATED
```

## `MODELLED`

The relevant construct and product semantics are represented with enough precision to expand the scoped target into its applicable coverage dimensions, identify unresolved applicability where it genuinely remains unknown, and name blocking gaps without inventing implementation truth.

`MODELLED` does not mean executable. Conditions may still be `UNKNOWN`, `DEFINED`, `PARTIAL`, `BLOCKED`, or `CALIBRATION_REQUIRED`.

## `COVERED`

The complete executable path exists for the scoped target with no blocking CoverageGap:

```text
external requirement / official family
→ canonical capability/knowledge/context
→ curriculum route
→ intervention/practice
→ assessment/evidence/re-evidence
→ progression/next action
→ learner experience
→ content/assets
→ machine/runtime/provider path
```

For promotion to `COVERED`, every applicable TargetCoverageSpecification condition **except `validation`** must be `SATISFIED`. An explicitly justified `NOT_APPLICABLE` condition is excluded from that target's required set. `UNKNOWN`, `DEFINED`, `PARTIAL`, `BLOCKED`, or `CALIBRATION_REQUIRED` on any applicable non-validation condition prevents `COVERED`.

`validation` is intentionally excluded because empirical product-outcome validation is the later `VALIDATED` gate. Evaluator/score calibration required for safe product execution belongs to `evaluator_scoring` and therefore can block `COVERED`.

## `SUPPORTED_FOR_PRODUCT`

A versioned TargetSupportDeclaration activates an already-`COVERED` target for a named product/release boundary and confirms the exact supported scope/version plus all applicable release-critical gates.

`SUPPORTED_FOR_PRODUCT` is not a shortcut around `COVERED`: declaration prose cannot convert an unsatisfied condition into support. `validation` may still remain unsatisfied because product support and empirically validated outcome are distinct states.

## `VALIDATED`

A target is `VALIDATED` only when it is currently `SUPPORTED_FOR_PRODUCT` **and** its scoped `validation` condition is `SATISFIED` by empirical evidence under named learner/product/content/evaluator/intervention versions and conditions.

Architecture coherence is never `VALIDATED` evidence.

# Condition status is separate

A **coverage condition** uses exactly:

```text
UNKNOWN
DEFINED
PARTIAL
SATISFIED
BLOCKED
NOT_APPLICABLE
CALIBRATION_REQUIRED
```

Condition statuses never appear as target/product statuses. They are not percentages and are not an ordinal score that may be averaged.

Exact semantics:

- `UNKNOWN` — applicability or current condition state cannot yet be established from the owning semantics/evidence. Unknown never passes a gate.
- `DEFINED` — applicability and success criteria are known, but sufficient executable/material evidence has not yet been instantiated or checked.
- `PARTIAL` — some required subconditions/artifacts have passed, while at least one applicable remainder is missing/unverified; no stronger known blocker status below applies.
- `SATISFIED` — every applicable subcondition for the scoped condition passes under the referenced current versions/evidence.
- `BLOCKED` — a known applicable hard failure or missing prerequisite prevents the condition from passing now; satisfying unrelated subparts cannot hide it.
- `NOT_APPLICABLE` — the owning semantics establish that the condition does not apply to this exact scoped target, with an explicit reason.
- `CALIBRATION_REQUIRED` — the relevant mechanism/evaluator/path materially exists, but calibration required for the intended consequence is not yet sufficient; this is a specialized blocking state, not a weaker form of `SATISFIED`.

When a condition aggregates subconditions, reduce deterministically:

1. use `NOT_APPLICABLE` only when the whole condition is explicitly non-applicable;
2. if applicability/state is unresolved, use `UNKNOWN` unless a more specific known blocker already determines failure;
3. a known hard blocker reduces to `BLOCKED`;
4. if the remaining material blocker is required calibration, use `CALIBRATION_REQUIRED`;
5. if all applicable subconditions pass, use `SATISFIED`;
6. if some pass and some remain missing/unverified without a stronger blocker, use `PARTIAL`;
7. if criteria are defined but executable/material checking has not materially begun, use `DEFINED`.

These statuses may move backward when external truth, content, contracts, provider state, calibration, rights, runtime reliability/security/operations evidence, or validation evidence changes.

An out-of-scope construct uses **OUT_OF_SCOPE** as a scope disposition, not a condition or product-support state.

# CoverageGap

A CoverageGap is a product inability or unresolved condition, never learner weakness.

Canonical classes:

- `MODEL_OR_SPEC`;
- `INTERVENTION_OR_ACTIVITY`;
- `CONTENT_OR_ASSET`;
- `EVIDENCE_OR_EVALUATOR`;
- `EXPERIENCE`;
- `TRANSITION`;
- `CONTRACT_OR_INTEGRATION`;
- `COST_OR_OPERATIONS`;
- `RIGHTS_PRIVACY_RELIABILITY`;
- `CALIBRATION_OR_VALIDATION`.

Each gap records:

```text
scoped target/condition
missing or failed condition
blocking consequence
dependencies
demand class
provenance/version
```

# Scope

## Complete standard IELTS learning construct

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ every in-scope official task/question family
+ material official subformats/presentation classes required for genuine content coverage
+ Band-3→9 learning interpretation
+ diagnostic → learning → practice → assessment → progression → exam-readiness
```

Release ordering may differ from construct completeness.

## Delivery modes

Delivery mode is a scope dimension, not another IELTS learning variant.

Current external delivery facts are owned by `../spec/02-IELTS-MODEL.md`. A product may be complete for a learning construct while still having a blocking readiness gap for a requested delivery mode.

## UKVI

UKVI Academic/GT reuses the corresponding language construct while adding external security/location/administrative requirements. A UKVI-purpose TargetProfile cannot imply support for an ineligible delivery route.

## One Skill Retake

One Skill Retake reuses an existing skill. It requires no fifth Skill ontology. Product support is scoped to focused preparation plus applicable eligibility/delivery conditions. Selecting an OSR focus is not evidence that the learner satisfies those external eligibility conditions.

## IELTS Life Skills

IELTS Life Skills is **OUT_OF_SCOPE** for the current Band-3→9 construct.

# Required TargetCoverageSpecification conditions

For each scoped target, evaluate every applicable condition independently.

| Condition | Required meaning |
|---|---|
| `construct_model` | current external IELTS construct represented correctly |
| `official_family_coverage` | every applicable stable official family ID from `02-IELTS-MODEL.md` has an executable product/content/evidence path |
| `material_subformat_coverage` | required Presentation Classes inside a family are represented where one family label would otherwise hide a material gap |
| `skill_knowledge_model` | required canonical capabilities/knowledge exist |
| `band_threshold` | target Band/task/variant threshold is defined |
| `curriculum_route` | valid sequencing/variant route exists |
| `practice_intervention` | required learning mechanisms/modes can be executed |
| `feature_experience` | learner can perform the interaction and understand result/state; runtime/browser behavior required by that experience is usable |
| `content_assets` | sufficient valid executable content supply covers required family/context/diversity; generators/templates count only where release actually relies on them |
| `assessment_policy` | applicable Assessment Type and executable versioned EvidenceRequirement exist |
| `evaluator_scoring` | deterministic scoring or calibrated productive evaluation exists |
| `progression_transition` | valid evidence can drive explainable state/next action |
| `variant_context` | Academic/GT task/section/context conditions are represented and sampled |
| `delivery_mode_readiness` | requested delivery-specific interaction/preparation exists when material |
| `machine_contracts` | exact implemented cross-unit interfaces, shared canonical identity materialization where consumed, API compatibility/conformance, and drift verification pass |
| `rights_privacy_security` | applicable source/data/consent/access/secrets/transport/storage/browser/provider/public-edge security requirements pass |
| `reliability_recovery` | applicable lifecycle, commit/idempotency/concurrency, timeout/retry/backpressure, migration/deploy, backup/restore, failure/degraded/recovery, and content/evaluator recovery behavior passes |
| `accessibility_capture_quality` | access/capture/browser/device failures cannot become false ability judgments and supported interaction remains usable |
| `cost_abuse_operations` | release operates within declared rate/abuse/provider-quota/compute/storage/traffic/capacity/scaling/external-usage constraints |
| `observability_audit` | applicable structured logs, metrics, correlation/tracing, privileged audit, consequential decision/version provenance, alerting/operational reconstruction are sufficient |
| `validation` | empirical outcome evidence exists when promoting to `VALIDATED` |

`NOT_APPLICABLE` requires an explicit reason for the scoped target. Absence of an artifact, unimplemented work, lack of data, inconvenience, or absence of a named enterprise technology is never an N/A reason.

A content generator is not a new coverage condition. Generation capability is required only when scoped product/release depends on generation to satisfy `content_assets` or another applicable operational requirement. A release with sufficient authored/imported/deterministic content can satisfy content coverage without AI/runtime generation.

# Runtime engineering closure

The system-engineering disposition model in `06-implementation-stack.md` is consumed through the existing coverage conditions above; it does not create another TargetCoverageSpecification condition or product status.

## `machine_contracts`

Where applicable this condition closes:

- exact public/internal cross-unit machine interface authority;
- generated/validated consumer/provider conformance;
- shared canonical registry/identity materialization actually consumed by runtime;
- API/application compatibility/version provenance;
- no handwritten equivalent DTO/canonical-ID truth across parallel runtimes.

## `rights_privacy_security`

Where applicable this condition closes deployed security obligations including:

- authentication, authorization/capability and least privilege;
- SQL injection, XSS, SSRF, CORS/CSP/CSRF protections appropriate to selected routes;
- secrets/key/session/token handling and revocation/rotation appropriate to the actual credential design;
- TLS/encryption in transit and applicable encryption-at-rest protection;
- privileged/admin and service-to-service access boundaries;
- provider/data-processor security and public-edge abuse/DDoS/WAF concern appropriate to deployment.

OAuth, JWT, or a dedicated WAF product is not required unless that technology is actually selected/needed; the underlying security concern still must pass.

## `reliability_recovery`

Where applicable this condition closes:

- durable commit-before-success-ACK and authoritative transaction behavior;
- idempotency, race/concurrency/stale-write correctness;
- deadlines/timeouts, retry classification/backoff, ambiguous network outcome handling, provider degradation, capacity/backpressure;
- migration/schema/deployment compatibility and rollback/forward-recovery;
- backup plus verified restore, PITR where applicable, disaster/recovery procedure;
- health/failure/degraded-state behavior, content/evaluator/work recovery;
- safe rollout/recovery independent of whether a named blue-green/canary/rolling strategy is used.

Reliable async evaluation can satisfy this gate without Kafka when durable DB work/outbox/idempotent bounded dispatch meets the same invariant.

## `cost_abuse_operations`

Where applicable this condition closes:

- rate limiting/abuse protection;
- provider quotas and duplicate-cost prevention;
- compute/storage/network/external-service usage and cost visibility;
- traffic/capacity/backlog pressure;
- scale strategy appropriate to measured demand.

Horizontal autoscaling, load balancers, CDN, read replicas, sharding, or Kubernetes are not individually required. They become applicable only when their trigger is present.

## `observability_audit`

Where applicable this condition closes:

- privacy-safe structured logging and correlation IDs;
- operational metrics including error/latency/capacity/backlog/provider health where material;
- trace correlation when needed for multi-unit diagnosis;
- privileged mutation audit and consequential decision/version/provenance reconstruction;
- actionable alerting plus supported-release incident ownership/escalation;
- measurable SLIs/SLOs appropriate to the supported release, without architecture inventing arbitrary numeric objectives.

Telemetry is derived evidence about operation. Logs/metrics/traces never replace authoritative product/learner/content/evidence state.

## Technology-independence invariant

Coverage evaluates **whether the concern is satisfied**, not whether a fashionable technology exists.

Therefore none of the following is intrinsically required for `COVERED`/support: Kubernetes, Helm, Kafka/another broker, Redis, vector database, distributed transactions, Saga orchestration, service mesh/discovery infrastructure, multi-region, database sharding/read replicas, leader election, gRPC, WebSockets, JWT, Docker, Terraform, external observability, or a paid GitHub capability.

If an applicable release cannot satisfy an invariant without one of those technologies, the demonstrated trigger may make that technology/path necessary downstream. Absence of the technology alone is not a CoverageGap.

# Derived reachability invariant

Coverage must be derivable from canonical identities and declared applicability rather than a manually maintained completeness matrix.

For each in-scope canonical target/object, coverage tooling should determine applicable downstream relationships and verify the required reachable path:

```text
Skill / enabling Knowledge where applicable
→ Curriculum route
→ trainable intervention / Practice
→ measurable Assessment path where claim requires measurement
→ Progression consequence
→ product/content/runtime conditions required by scoped target
```

The validator distinguishes:

```text
APPLICABLE + edge/path absent   = MISSING / blocking gap
explicitly NOT_APPLICABLE       = allowed only with owner-derived reason
applicability unresolved        = UNKNOWN, never silently satisfied
```

Rules:

1. absence of an edge or implementation artifact never defaults to `NOT_APPLICABLE`;
2. `NOT_APPLICABLE` is derivable from canonical semantics or carries an explicit owner-derived reason;
3. not every object requires every edge, but its required downstream purpose remains reachable;
4. an applicable Skill Leaf cannot be complete if orphaned from required Curriculum/Practice/Assessment/Progression reachability;
5. derived reports are verification artifacts, never manually edited SSOTs.

# Official-family closure

Skill coverage and official-family coverage are independent dimensions.

Examples:

- `R-QT-04` may help teach several matching families, but content for `IELTS-R-QF-04` cannot satisfy `IELTS-R-QF-06` or `IELTS-R-QF-07`;
- one broad completion Skill cannot prove existence of all required official completion-family content;
- strong Reading capability teaching can coexist with a blocking missing official-family interaction/content path.

For a complete standard target, every applicable stable `IELTS-*-*` family resolves through:

```text
official family ID
→ feature interaction
→ Practice/Assessment role
→ concrete content family/subformat
→ answer/rubric/evaluator path
→ valid learner-state/readiness flow
```

# Material-subformat closure

Coverage consumes stable Presentation Classes from `../spec/10-CONTENT-MODEL.md` where defined.

Minimum current checks include:

- Listening `IELTS-L-QF-04`: form, note, table, flow-chart, summary completion;
- Reading `IELTS-R-QF-09`: summary, note, table, flow-chart completion;
- Academic Writing `IELTS-W-A-T1`: statistical graph/chart/table, diagram/process, map/plan.

One template cannot satisfy a multi-presentation condition merely by carrying parent family ID.

# Content closure

Feature/UI existence is not content coverage. Content closure is evaluated through existing coverage conditions rather than a new product-status taxonomy.

When executable content exists, coverage tooling consumes the content manifest/equivalent index from `../spec/10-CONTENT-MODEL.md` and verifies at minimum:

```text
exact content revision + lineage/provenance where material
canonical target refs
stable official family refs + applicability where material
Content Context refs + applicability / where material
Presentation Class refs + applicability where material
variant/task/section context where material
practice/assessment purpose + evidence-candidacy compatibility
interaction support
answer/rubric/evaluator route
difficulty/transfer classes
applicable validation decision/policy refs
rights/provenance
independent readiness assets
release/assignment/operational eligibility
```

For every conditional identity dimension, tooling preserves:

```text
applicable + present
explicitly NOT_APPLICABLE with owner-derived reason
required but missing/unresolved
```

A uniform implementation shape cannot fabricate identity or hide required identity.

Content closure has four concerns through applicable existing conditions:

1. **Supply sufficiency** — enough executable eligible content exists for target/family/context/presentation/difficulty/diversity demand; authored/imported/deterministic/pre-generated/runtime-generated routes are eligible.
2. **Semantic/quality validity** — every assigned revision passes applicable content contract/universal hard/consequence-specific checks; global quality percentage cannot substitute for a failed requirement.
3. **Assignment novelty/independence where applicable** — learner-specific exposure/similarity is checked before transfer/readiness/evidence assignment when required.
4. **Operations/recovery viability** — revisions are auditable; problematic content can stop new assignment, be revalidated/replaced/retired without rewriting historical learner Attempts/evidence, and coverage-critical inventory can recover.

Generation is only an applicable sub-demand when scoped release relies on it. Absence of runtime AI generation is not a gap when sufficient eligible supply exists. A generator/template cannot satisfy `content_assets` until output demonstrates applicable quality/coverage/operations.

Similarity is not universal rejection. A near-duplicate can be valid controlled practice and invalid unseen evidence.

A `CONTENT_OR_ASSET` condition becomes `SATISFIED` only from executable/verified supply, not design tables or unproven generators. Generated canonical registries/bindings do not independently satisfy `content_assets`.

# Evidence closure

Principles in `08-ASSESSMENT.md` alone are not production policy.

Before product support, each high-consequence claim resolves to a versioned executable EvidenceRequirement. Hidden heuristics, unversioned cutoffs, mechanical attempt-count rules, or pre-attempt “certification contributing” labels are blockers.

Official-family product coverage does not imply learner certification mechanically tests every family in every portfolio. Learner claim sufficiency remains Assessment-owned.

Content used for evidence preserves exact revision and exposure/independence conditions required by applicable EvidenceRequirement. Later content correction/retirement does not rewrite historical evidence; reinterpretation follows Assessment/Progression policy with provenance.

# Contract and machine-materialization closure

A multi-runtime path cannot become `COVERED` while TypeScript, Go, and Python independently maintain handwritten interpretations of the same interface or equivalent manually copied canonical ID registries.

Every implemented cross-unit boundary needs one machine contract authority, generated/validated consumers where appropriate, and conformance verification.

For every implemented path consuming shared canonical identities, applicable derived machine-readable registry/binding materialization defined by `06-implementation-stack.md` exists and canonical-reference/ID-uniqueness/drift verification passes.

Generated registries/bindings remain derived repository/runtime verification artifacts. They are not canonical learning/product truth, content coverage evidence, product validation evidence, or independent satisfaction of `content_assets`/support declaration.

These requirements close through existing machine/runtime/contract conditions; they do not create another condition/status.

Implemented content generation/validation boundaries require contracts when they exist; unused optional generation does not require hypothetical contract.

# Delivery closure

When TargetProfile names a material delivery mode, exam-readiness coverage exposes that interaction honestly. Delivery preparation changes conditions, not Skill/Band truth.

# Current design-state declaration — 2026-08-22

This is a documentation/design declaration, not production support.

| Scoped target | Model state | Remaining blocking classes | Product status |
|---|---|---|---|
| Academic Listening | strong construct/learning/design model incl. stable family IDs | official-family/subformat content, runtime/contracts, operations/validation | `MODELLED` |
| Academic Reading | strong construct/learning/design model incl. stable family IDs | official-family content, runtime/contracts, operations/validation | `MODELLED` |
| Academic Writing | strong construct/learning/design model incl. Task-1 presentation classes | productive evaluator calibration, diverse Task-1 content, runtime/contracts | `MODELLED` |
| Academic Speaking | strong construct/learning/design model incl. Part 1–3 family IDs | productive/audio calibration, content, runtime/contracts | `MODELLED` |
| GT Listening | shared strong model | official-family/subformat content, runtime/contracts, validation | `MODELLED` |
| GT Speaking | shared strong model | Part 1–3 content/evaluator/runtime verification | `MODELLED` |
| GT Reading | shared skills + explicit GT context/scoring + official families | GT family/context assets, runtime/contracts, validation | `MODELLED` |
| GT Writing Task 1 | dedicated GT capability + Band/Curriculum/product semantics | letter asset diversity, productive calibration, runtime/contracts | `MODELLED` |
| GT Writing Task 2 | shared Writing construct | content, evaluator calibration, runtime/contracts | `MODELLED` |
| One Skill Retake preparation | existing-skill reuse + target-flow semantics | release/admin/delivery workflow, runtime, validation | `MODELLED` |
| Test-centre computer exam-readiness | external delivery model represented | executable UX/content/runtime verification | `MODELLED` |
| Writing-on-Paper readiness | external delivery overlay represented | handwriting-specific product assets/verification by supported market | `MODELLED` |
| IELTS Online Academic readiness | external delivery overlay represented | remote-mode product support, purpose/acceptance handling, verification | `MODELLED` |

Scope disposition:

```text
IELTS Life Skills = OUT_OF_SCOPE
```

Top-level current truth:

```text
Academic learning/design semantics         MODELLED
General Training learning/design semantics MODELLED
Official family identities                 MODELLED
Material presentation identities           MODELLED where required
Delivery-mode semantics                    MODELLED
Content runtime/governance semantics        MODELLED
Runtime engineering concern closure         MODELLED
Academic product execution                 NOT YET COVERED
General Training product execution         NOT YET COVERED
Runtime implementation                     NOT IMPLEMENTED
Validated target-band outcome              NOT ESTABLISHED
```

`Runtime engineering concern closure MODELLED` means the required concern semantics/dispositions are defined; it does not mean runtime/security/operations implementation or verification exists.

No user-facing copy may claim full support until all scoped blocking conditions close and a TargetSupportDeclaration activates that exact scope.

# Construct → product mapping

This mapping proves design intent only. It does not satisfy runtime/content gates.

## Listening

| Official family ID | Product feature/practice path |
|---|---|
| `IELTS-L-QF-01` Multiple choice | `L-F01`, `L-F05`, `PM-L04`, `PM-L06` |
| `IELTS-L-QF-02` Matching | `L-F01`, `PM-L06` plus focused `PT-13` items |
| `IELTS-L-QF-03` Plan/map/diagram labelling | `L-F06`, `PM-L05` |
| `IELTS-L-QF-04` Form/note/table/flow-chart/summary completion | `L-F04`, `PM-L03`; all required `PRES-L-QF04-*` classes remain independently checkable |
| `IELTS-L-QF-05` Sentence completion | `L-F04`, `PM-L03` |
| `IELTS-L-QF-06` Short answer | `L-F04`, `PM-L03` |

## Reading

| Official family ID | Product feature/practice path |
|---|---|
| `IELTS-R-QF-01` Multiple choice | `R-F01`, `R-F06`, `PM-R05` / `PM-R06` |
| `IELTS-R-QF-02` True/False/Not Given | `R-F04`, `PM-R03` |
| `IELTS-R-QF-03` Yes/No/Not Given | `R-F04`, `PM-R03` |
| `IELTS-R-QF-04` Matching information | `R-F03`, `R-F07`, focused Reading item + `PM-R06` |
| `IELTS-R-QF-05` Matching headings | `R-F05`, `PM-R04` |
| `IELTS-R-QF-06` Matching features | `R-F03`, `R-F06`, focused Reading item |
| `IELTS-R-QF-07` Matching sentence endings | `R-F06`, focused Reading item |
| `IELTS-R-QF-08` Sentence completion | `R-F03`, focused completion item |
| `IELTS-R-QF-09` Summary/note/table/flow-chart completion | `R-F03`, `R-F06`; required `PRES-R-QF09-*` classes remain independently checkable |
| `IELTS-R-QF-10` Diagram label completion | `R-F03`, focused Reading item |
| `IELTS-R-QF-11` Short answer | `R-F03`, focused Reading item / `PM-R06` |

Variant coverage additionally requires applicable Academic or GT Content Context distribution and correct scoring/evidence policy.

## Writing

| Official task family ID | Product path | Canonical target/evidence path |
|---|---|---|
| `IELTS-W-A-T1` | `W-F01`, `W-F03`, `W-F07`, `W-F08`, `W-F09` | `W-TA-*`; required `PRES-W-A-T1-*` content diversity; plan → draft → revision → timed independent |
| `IELTS-W-GT-T1` | variant-aware Writing features | `W-GT1-*`; recipient/purpose/register/required-point coverage → draft → timed independent |
| `IELTS-W-T2` | `W-F01`, `W-F02`, `W-F06`, `W-F07`, `W-F08`, `W-F09` | component → full → timed independent |
| Writing criteria | criterion feedback/evaluator path | `AT-01` / `AT-05`; calibration required for high-consequence support |

## Speaking

| Official family ID | Product path |
|---|---|
| `IELTS-S-P1` | `S-F01` / `PM-S03` |
| `IELTS-S-P2` | `S-F02` / `PM-S04` |
| `IELTS-S-P3` | `S-F03` / `PM-S05` |
| whole Speaking construct | `S-F09` / `PM-S06` with Part 1–3 coverage |
| pronunciation learning surface | `S-F04`, `S-F05`, `S-F07` / `PM-S01`, `PM-S02` |

# TargetProfile coverage invariant

A product target is a constraint set, not one Band number.

Target-relative support requires resolved standard variant plus covered path for every real required TargetProfile condition: overall/per-skill Band constraints, external purpose constraints, selected OSR eligibility conditions when claim includes them, and delivery mode where material.

If only overall Band constraint is known, planner may use explicitly non-authoritative planning-profile semantics from `00-learning-experience.md`; it cannot invent real per-skill minima.

# Route invariant

For a supported target:

```text
TargetProfile
  ↓
provisional diagnostic
  ↓
explicit unresolved conditions
  ↓
Planner hard eligibility
  ↓
ranked valid next action
  ↓
practice / review / re-evidence
  ↓
updated claim states
  ↺
```

Ranking cannot erase prerequisites, family/context/delivery compatibility, content assignment eligibility, evidence truth, or CoverageGaps.

# Promise boundary

`SUPPORTED_FOR_PRODUCT` means complete release-qualified path for declared scope. It does not guarantee external IELTS result.

Allowed:

- “Your current evidence supports Band N under this profile.”
- “These are the remaining blockers to your target.”
- “This target/delivery scope is supported by the current product version.”

Forbidden without suitable empirical basis:

- “Follow this plan and you will get Band 7.”
- “100% guaranteed target Band.”
- “This exercise improved your IELTS score by 0.5 Band.”

# Gap → implementation demand

Every blocking CoverageGap maps to one demand class:

- spec/model;
- feature/interaction;
- content/assets/supply route;
- evaluator/calibration;
- learner flow/transition;
- contract/integration;
- runtime/provider/operations;
- rights/privacy/security;
- validation/research.

Only actual content gaps create content-quantity demand. A generator is selected/implemented only when demonstrated supply requirement cannot be satisfied by existing eligible content/approved supply route under quality/operations constraints.

# TargetSupportDeclaration minimum

A release declaration names at least:

- exact target/variant scope;
- supported Band/TargetProfile conditions;
- supported delivery modes/purpose constraints where material;
- product/release version;
- official-family and material-subformat coverage results;
- feature/practice coverage;
- content-manifest version + exact active revision/inventory scope;
- applicable content-validation policy/result references;
- assignment novelty/independence gate where material;
- content incident/recovery/retirement gate;
- EvidenceRequirement/policy versions;
- evaluator/calibration state;
- machine-contract versions;
- rights/privacy/security state;
- third-party activation state;
- reliability/recovery state;
- accessibility/capture-quality gate;
- cost/abuse/operations gate;
- observability/audit gate including supported-release operational objectives/incident ownership;
- known non-blocking validation backlog;
- revocation conditions.

Support is versioned/revocable when construct, delivery, provider, rights/security, reliability/recovery, cost/capacity, calibration, content coverage/quality/operations, observability, or validation evidence materially changes.