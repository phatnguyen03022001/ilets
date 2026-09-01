# CURRICULUM-COVERAGE Curriculum and product-support semantics

> **CANONICAL PRODUCT AUTHORITY**
>
> Canonical within the PRODUCT authority domain under `docs/catalog/project.json`. `CONSTITUTION.md` and `OBJECTIVE.md` retain distinct authority, and `contracts/**` retains scoped exact machine-contract authority. References below to `spec/**` or `design/**`, including historical uses of “canonical”, are provenance only and do not create equal authority.

## Curriculum semantics

The active curriculum has **44 stable `C-*` Curriculum Node identities** across Bands 3–9. A node owns its stable id/band phase, recommended sequence position, Skill/Knowledge targets, focus, expected outcome, explicit recommended node dependencies, and node-exit intent. Legacy hour/load estimates are not canonical. Node completion is not Band certification.

Dependency classes are exact:

- **Required** — intrinsic Skill→Skill, Knowledge→Knowledge or Skill→Knowledge prerequisite. An unresolved Required edge hard-gates dependent learning. Required object edges are inherited from the Skill/Knowledge owners.
- **Recommended** — useful Curriculum Node ordering encoded by `Depends`; it is not a hard gate when evidence justifies another path.
- **Independent** — no prerequisite relation.

`Depends` contains only stable `C-*` IDs and means Recommended sequencing. It never contains Skill/Knowledge IDs or prose aliases. Human-readable focus/outcome text never creates an implicit dependency.

### Deterministic selector contract

Only these compile-time selector tokens may appear in node target columns:

```text
@task1-analysis
@task1-fulfilment
@task1-global-control
@variant-reading-advanced-knowledge
@band-phase-skill-set
@band-phase-knowledge-set
```

`@task1-*` and `@variant-reading-advanced-knowledge` resolve from the selected Academic/GT variant. `@band-phase-skill-set` is the union of all fully resolved Skill targets in preceding nodes of that Band phase; `@band-phase-knowledge-set` is the equivalent explicit Knowledge union. Expansion is recursive until only canonical IDs remain. An unresolved selector at a materialization boundary is invalid; a genuinely unresolved variant permits only shared nodes that do not require the missing variant. No ad-hoc selector may be invented.

Variant-specific selector resolution:

| Selector | Academic | General Training |
|---|---|---|
| `@task1-analysis` | `W-TA-01` | `W-GT1-01` |
| `@task1-fulfilment` | `W-TA-03` | `W-GT1-03` |
| `@task1-global-control` | `W-TA-02` | `W-GT1-02` |
| `@variant-reading-advanced-knowledge` | `K-VOC-011` | empty set |

The GT empty-set resolution means no single extra Knowledge Object is universally required solely because GT Reading is dense. Variant Reading differences remain context/content/readiness conditions. Academic uses Academic passage/corpus conditions; GT preserves Section 1 everyday, Section 2 workplace and Section 3 longer general-interest distribution. GT learners do not need Academic `W-TA-*` leaves for GT Task 1, and Academic learners do not need `W-GT1-*`; shared capabilities remain reusable.

### Stable Curriculum Node inventory

| Node | Focus | Skill targets | Knowledge targets | Recommended Depends | Exit intent |
|---|---|---|---|---|---|
| `C-B3-01` | Word classes, clause structure, simple sentences | `W-GRA-01`, `S-GRA-01` | `K-GRA-010`, `K-GRA-001`, `K-GRA-002` | — | accurate basic written/spoken sentence cores; foundation grammar acquired |
| `C-B3-02` | Nouns, countability, determiners/articles | `W-GRA-05` | `K-GRA-061`, `K-GRA-032`, `K-GRA-040` | `C-B3-01` | basic countability/determiner/article control |
| `C-B3-03` | Basic tense and agreement | `W-GRA-04` | `K-GRA-050`, `K-GRA-051`, `K-GRA-060` | `C-B3-01` | present/past basics and agreement in simple production |
| `C-B3-04` | Core vocabulary and spelling | `W-LR-01`, `W-LR-05`, `S-LR-01` | `K-VOC-010`, `K-VOC-031` | — | high-frequency vocabulary with basic spelling control |
| `C-B3-05` | Phoneme foundations and intelligibility | `S-P-01`, `S-P-05` | `K-PHON-010`, `K-PHON-011`, `K-PHON-012` | — | core sound contrasts with broad intelligibility |
| `C-B3-06` | Receptive gist and detail | `L-COMP-01`, `L-COMP-02`, `R-COMP-01`, `R-COMP-02` | `K-VOC-010` | `C-B3-04` | gist/key detail in simple spoken/written input |
| `C-B3-07` | Basic receptive question types | `L-QT-01`, `L-QT-05`, `R-QT-05` | — | `C-B3-06` | basic receptive items within form/word limits |
| `C-B3-08` | Foundation integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B3-01`, `C-B3-02`, `C-B3-03`, `C-B3-04`, `C-B3-05`, `C-B3-06`, `C-B3-07` | Band-3 integration; not a four-skill certification gate |
| `C-B4-01` | Writing accuracy, punctuation, paragraphing | `W-GRA-01`, `W-GRA-06`, `W-CC-01` | — | `C-B3-01`, `C-B3-03` | simple punctuated paragraphed writing |
| `C-B4-02` | Task-1 analysis and Task-2 prompt analysis | `@task1-analysis`, `W-TR-01` | — | `C-B3-08` | selected-variant Task-1 requirements + Task-2 prompt analysis |
| `C-B4-03` | Topic vocabulary and word choice | `W-LR-01`, `S-LR-01`, `S-LR-04` | `K-VOC-012`, `K-VOC-040` | `C-B3-04` | adequate/context-appropriate common-topic vocabulary |
| `C-B4-04` | Speaking continuity and basic accuracy | `S-FC-01`, `S-GRA-04`, `S-P-05` | `K-GRA-010` | `C-B3-01`, `C-B3-05` | basic sustained familiar-topic speech |
| `C-B4-05` | Receptive main ideas and details | `L-COMP-01`, `L-COMP-02`, `R-COMP-01`, `R-COMP-02` | `K-VOC-012` | `C-B3-06` | gist/detail in moderately more complex input |
| `C-B4-06` | Basic receptive question types at Band-4 demand | `L-QT-01`, `L-QT-05`, `R-QT-05` | — | `C-B4-05` | current-demand basic receptive strategies |
| `C-B4-07` | Band-4 integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B4-01`, `C-B4-02`, `C-B4-03`, `C-B4-04`, `C-B4-05`, `C-B4-06` | non-certifying Band-4 orchestration checkpoint |
| `C-B5-01` | Compound/complex sentence foundation | `W-GRA-02`, `W-GRA-03` | `K-GRA-003`, `K-GRA-020`, `K-GRA-004`, `K-GRA-021` | `C-B4-01` | usable compound/complex sentence control |
| `C-B5-02` | Cohesion and logical progression | `W-CC-02`, `W-CC-03`, `W-CC-04` | `K-GRA-022`, `K-GRA-033` | `C-B4-01`, `C-B5-01` | progression, linking, reference/substitution |
| `C-B5-03` | Task-2 position, development, relevance | `W-TR-02`, `W-TR-03`, `W-TR-04` | `K-VOC-012` | `C-B4-02`, `C-B5-02` | relevant Task-2 response with clear supported position |
| `C-B5-04` | Task-1 content fulfilment | `@task1-fulfilment` | — | `C-B4-02` | selected-variant Task-1 content fulfilment |
| `C-B5-05` | Paraphrase, word formation, spelling | `W-LR-03`, `W-LR-05` | `K-VOC-030`, `K-VOC-031` | `C-B4-03` | meaning-preserving paraphrase + productive word form/spelling |
| `C-B5-06` | Speaking fluency, complexity, pronunciation | `S-FC-02`, `S-FC-03`, `S-FC-04`, `S-FC-05`, `S-GRA-02`, `S-LR-02`, `S-P-02`, `S-P-03` | `K-GRA-021`, `K-VOC-012` | `C-B4-04`, `C-B5-01` | extended coherent turn with more complex language/stress/intonation |
| `C-B5-07` | Receptive inference, paraphrase, distractors | `L-COMP-03`, `L-COMP-04`, `L-COMP-05`, `R-COMP-03`, `R-COMP-04`, `R-COMP-05` | — | `C-B4-05` | more reliable inference/paraphrase/structure/distractors |
| `C-B5-08` | Higher-order receptive question types | `L-QT-02`, `L-QT-03`, `L-QT-04`, `R-QT-01`, `R-QT-02`, `R-QT-04`, `R-QT-06` | — | `C-B5-07` | higher-order question strategies under realistic constraints |
| `C-B5-09` | Band-5 integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B5-01`, `C-B5-02`, `C-B5-03`, `C-B5-04`, `C-B5-05`, `C-B5-06`, `C-B5-07`, `C-B5-08` | non-certifying Band-5 integration |
| `C-B6-01` | Grammatical accuracy and variety | `W-GRA-03`, `W-GRA-02`, `S-GRA-03` | `K-GRA-004`, `K-GRA-003` | `C-B5-01` | improved productive accuracy/variety |
| `C-B6-02` | Task-1 global control and flexible reference | `@task1-global-control`, `W-CC-04` | — | `C-B5-02`, `C-B5-04` | selected-variant global Task-1 control + reduced repetition |
| `C-B6-03` | Collocation and idiomatic resource | `W-LR-02`, `S-LR-03` | `K-VOC-020`, `K-VOC-021` | `C-B4-03` | accurate collocation and emerging appropriate less-common/idiomatic use |
| `C-B6-04` | Extended/dense receptive content and variant transfer | `L-COMP-06`, `R-COMP-06` | `@variant-reading-advanced-knowledge` | `C-B5-07` | dense/extended input + selected-variant context transfer |
| `C-B6-05` | Writer-view classification | `R-QT-03` | — | `C-B5-08` | distinguish writer agreement/disagreement/absence |
| `C-B6-06` | Band-6 integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B6-01`, `C-B6-02`, `C-B6-03`, `C-B6-04`, `C-B6-05` | non-certifying Band-6 integration |
| `C-B7-01` | Structural flexibility/error-free frequency | `W-GRA-07`, `W-GRA-03` | `K-GRA-005`, `K-GRA-006`, `K-GRA-007` | `C-B6-01` | flexible varied complex structures with frequent accuracy |
| `C-B7-02` | Lexical sophistication | `W-LR-04`, `S-LR-03` | `K-VOC-021` | `C-B6-03` | less-common/idiomatic resource with increasing precision |
| `C-B7-03` | Cohesion, position, development | `W-CC-03`, `W-CC-04`, `W-TR-02`, `W-TR-03` | `K-GRA-022`, `K-GRA-033` | `C-B5-03`, `C-B6-02` | sustained clear developed positions/cohesion |
| `C-B7-04` | Receptive inference and structure | `R-COMP-04`, `R-COMP-05`, `L-COMP-03`, `L-COMP-04` | — | `C-B6-04` | reliable inference/stance/paraphrase/structure |
| `C-B7-05` | Pronunciation range/flexibility | `S-P-03`, `S-P-04` | `K-PHON-030`, `K-PHON-040` | `C-B5-06` | useful intonation/connected speech/chunking across extended turns |
| `C-B7-06` | Band-7 integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B7-01`, `C-B7-02`, `C-B7-03`, `C-B7-04`, `C-B7-05` | non-certifying Band-7 integration |
| `C-B8-01` | Wide flexible language/near-error-free accuracy | `W-GRA-07`, `W-LR-04`, `W-GRA-01`, `W-GRA-03`, `S-GRA-03` | `K-GRA-005`, `K-GRA-007` | `C-B7-01`, `C-B7-02` | wide flexible highly accurate language |
| `C-B8-02` | Skillful cohesion/fully developed response | `W-CC-02`, `W-CC-03`, `W-CC-04`, `W-TR-03`, `@task1-global-control` | `K-GRA-022`, `K-GRA-033` | `C-B7-03` | well-developed easy-to-follow response incl. variant Task-1 control |
| `C-B8-03` | Effortless speaking fluency/wide resource | `S-FC-02`, `S-FC-04`, `S-GRA-03`, `S-LR-03`, `S-P-05` | `K-PHON-040`, `K-PHON-041` | `C-B7-05` | fluent wide-resource easily understood speech |
| `C-B8-04` | Complex receptive comprehension | `L-COMP-06`, `R-COMP-06`, `R-COMP-04` | — | `C-B7-04` | complex/abstract input with little difficulty across variant contexts |
| `C-B8-05` | Band-8 integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B8-01`, `C-B8-02`, `C-B8-03`, `C-B8-04` | non-certifying Band-8 integration |
| `C-B9-01` | Full flexibility and precision | `W-GRA-07`, `W-LR-04`, `S-GRA-03`, `S-LR-03`, `S-P-05` | `K-GRA-005`, `K-GRA-007`, `K-VOC-021` | `C-B8-01`, `C-B8-02`, `C-B8-03` | near-full flexibility/precision/sustained intelligibility |
| `C-B9-02` | Integrated exam-level mastery | `@task1-global-control`, `W-TR-03`, `W-CC-03`, `R-COMP-04`, `L-COMP-06` | — | `C-B8-05` | ceiling-level independent exam-like demand for selected variant |
| `C-B9-03` | Ceiling integration | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B9-01`, `C-B9-02` | final integration; certification remains Assessment/Progression-owned |

Adaptive sequencing may reorder/interleave nodes only while Required prerequisites, complete target outcomes and variant requirements remain satisfied. Learners may work in different band phases per skill. Integration nodes do not synchronize four-skill certification. Static materialization fails on unknown IDs/selectors, unresolved selectors where required, prose in identity positions, missing `Depends` nodes, Recommended cycles, or wrong-variant Task-1 expansion.

## Product-support semantics

A scoped product target uses exactly this status progression:

```text
MODELLED → COVERED → SUPPORTED_FOR_PRODUCT → VALIDATED
```

- `MODELLED`: construct/product semantics are represented sufficiently to expand applicable conditions and expose unresolved gaps; it is not executable support.
- `COVERED`: the complete executable path exists with no blocking CoverageGap, and every applicable TargetCoverageSpecification condition except `validation` is `SATISFIED`. `UNKNOWN`, `DEFINED`, `PARTIAL`, `BLOCKED` and `CALIBRATION_REQUIRED` on any applicable non-validation condition block COVERED.
- `SUPPORTED_FOR_PRODUCT`: a versioned TargetSupportDeclaration activates an already-COVERED target for a named product/release boundary and confirms exact scope/version and release-critical gates. Declaration prose cannot bypass COVERED.
- `VALIDATED`: current SUPPORTED_FOR_PRODUCT plus scoped empirical `validation = SATISFIED` under named learner/product/content/evaluator/intervention versions/conditions. Architecture coherence is not validation evidence.

Coverage-condition status is a separate vocabulary:

```text
UNKNOWN
DEFINED
PARTIAL
SATISFIED
BLOCKED
NOT_APPLICABLE
CALIBRATION_REQUIRED
```

`UNKNOWN` never passes; `DEFINED` means criteria known but executable/material checking not sufficiently instantiated; `PARTIAL` means some subconditions pass but remainder is missing/unverified; `SATISFIED` means every applicable subcondition passes; `BLOCKED` is a known hard failure/missing prerequisite; `NOT_APPLICABLE` requires owner-derived non-applicability; `CALIBRATION_REQUIRED` means a mechanism/path exists but required calibration for the consequence is insufficient. Reduction order preserves unresolved applicability, hard blockers and calibration rather than averaging statuses. `OUT_OF_SCOPE` is a scope disposition, not a condition/product-support status.

### CoverageGap taxonomy

A CoverageGap is product inability/unresolved product condition, never learner weakness. Canonical classes:

`MODEL_OR_SPEC`, `INTERVENTION_OR_ACTIVITY`, `CONTENT_OR_ASSET`, `EVIDENCE_OR_EVALUATOR`, `EXPERIENCE`, `TRANSITION`, `CONTRACT_OR_INTEGRATION`, `COST_OR_OPERATIONS`, `RIGHTS_PRIVACY_RELIABILITY`, `CALIBRATION_OR_VALIDATION`.

Each gap preserves scoped target/condition, missing/failed condition, blocking consequence, dependencies, demand class and provenance/version.

### Required TargetCoverageSpecification conditions

Every scoped target evaluates every applicable condition independently:

| Condition | Required meaning |
|---|---|
| `construct_model` | current external IELTS construct represented correctly |
| `official_family_coverage` | every applicable stable official-family ID has executable product/content/evidence path |
| `material_subformat_coverage` | required Presentation Classes represented where one family label would hide a material gap |
| `skill_knowledge_model` | required canonical capabilities/knowledge exist |
| `band_threshold` | target Band/task/variant threshold defined |
| `curriculum_route` | valid sequencing/variant route exists |
| `practice_intervention` | required learning mechanisms/modes executable |
| `feature_experience` | learner can perform interaction and understand state/result; required runtime/browser behavior usable |
| `content_assets` | sufficient valid executable content covers family/context/diversity; generators count only where release depends on them |
| `assessment_policy` | applicable Assessment Type and executable versioned EvidenceRequirement exist |
| `evaluator_scoring` | deterministic scoring or calibrated productive evaluation exists |
| `progression_transition` | valid evidence drives explainable state/next action |
| `variant_context` | Academic/GT task/section/context represented and sampled |
| `delivery_mode_readiness` | requested delivery-specific preparation/interaction exists where material |
| `machine_contracts` | exact implemented cross-unit interfaces/shared canonical identity materialization, skew compatibility, applicability preservation, conformance and drift verification pass |
| `rights_privacy_security` | applicable source/data/consent/access/secrets/transport/storage/browser/SSE/internal-service/provider/public-edge requirements pass |
| `reliability_recovery` | applicable lifecycle, transaction/idempotency/concurrency, durable async, retry/backpressure, migration/deploy, restore/data-lifecycle and degraded/recovery behavior passes |
| `accessibility_capture_quality` | access/capture/browser/device failure cannot become false learner-ability judgment and supported interaction is usable |
| `cost_abuse_operations` | declared rate/abuse/provider-quota/compute/storage/traffic/capacity/scaling/external-usage constraints pass |
| `observability_audit` | structured logs/metrics/correlation, privileged audit, consequential provenance, alerting and reconstruction are sufficient |
| `validation` | empirical outcome evidence exists for VALIDATED promotion |

Absence of an artifact, lack of implementation/data, inconvenience, or absence of a named technology is never a valid `NOT_APPLICABLE` reason. A content generator is required only when the scoped release actually relies on generation.

### Derived reachability and official-family closure

Coverage is derived from canonical identities/applicability, never from a manually maintained percentage matrix:

```text
Skill / enabling Knowledge where applicable
→ Curriculum route
→ trainable Practice/intervention
→ measurable Assessment path where required
→ Progression consequence
→ required product/content/runtime conditions
```

For each in-scope target/object: applicable missing edge/path is blocking; explicit owner-derived N/A is allowed; unresolved applicability remains UNKNOWN. Not every object needs every edge, but its required downstream purpose must remain reachable. Derived reports are verification artifacts, not edited SSOTs.

Skill coverage and official-family coverage are independent. A broad Skill or one completion template cannot prove all external-family or Presentation-Class coverage. A complete standard target requires every applicable stable `IELTS-*-*` family, required material Presentation Class, variant context, interaction, content supply, scoring/evaluator path and evidence consequence to close independently. Strong teaching of a shared capability cannot hide a missing official-family interaction/content path.

Learner evidence state, capture failure, evaluator calibration, provider availability, paid entitlement, privileged authorization and product CoverageGap remain distinct. Missing/stale/conflicting/below learner evidence is not by itself a CoverageGap; paid entitlement cannot make an uncovered route supported; lack of a paid optional capability cannot make a supported ordinary route unsupported; and product inability must remain an explicit CoverageGap rather than learner weakness.

### Target support declaration and promotion gates

Target-relative support applies to a target constraint set, not one Band number. It requires a resolved standard variant plus a covered path for every real required TargetProfile condition: overall/per-skill Band constraints, external-purpose constraints, selected One Skill Retake eligibility conditions when the claim includes them, and material delivery mode. If only an overall Band is known, planning may use an explicitly non-authoritative working profile but cannot invent real per-skill minima. Ranking never erases prerequisites, family/context/delivery compatibility, assignment eligibility, evidence truth, or CoverageGaps.

A `TargetSupportDeclaration` names at least:

- exact target/variant scope;
- supported Band/TargetProfile conditions;
- supported delivery modes and purpose constraints where material;
- product/release version;
- official-family and material-subformat coverage results;
- feature/practice coverage;
- content-manifest version plus exact active revision/inventory scope;
- applicable content-validation policy/result references;
- assignment novelty/independence gate where material;
- content incident/recovery/retirement gate;
- EvidenceRequirement/policy versions;
- evaluator/calibration state;
- machine-contract versions and directional-compatibility state for the selected rollout;
- rights/privacy/security state;
- third-party activation state;
- reliability/recovery/data-lifecycle state;
- accessibility/capture-quality gate;
- cost/abuse/operations gate;
- observability/audit gate, including consequential version/config/provider provenance and release-candidate operational objectives/incident ownership;
- known non-blocking validation backlog;
- revocation conditions.

Support is versioned and revocable when construct, delivery, provider, rights/security, reliability/recovery, cost/capacity, calibration, content coverage/quality/operations, observability, contract compatibility, or validation evidence materially changes. `SUPPORTED_FOR_PRODUCT` means a complete release-qualified path for the declared scope; it never guarantees an external IELTS result or authorizes claims that a plan guarantees a target Band or that one exercise caused a Band increase without suitable empirical basis.
