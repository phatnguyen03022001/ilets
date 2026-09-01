# EVIDENCE-PROGRESSION-CONTENT Evidence, progression, and content semantics

## Assessment evidence semantics

The semantic pipeline is:

```text
Attempt
→ Observation
→ claim-scoped eligibility / interpretation
→ EvidenceFact
→ claim-scoped aggregation
→ ReadinessEvaluation
→ Progression decision
```

An Attempt is not automatically evidence. Evidence is not automatically mastery. Mastery is not automatically exam readiness. An Observation records what was measured plus material task/context, assistance/scaffolding, exposure, delivery, evaluator/scorer output, uncertainty and provenance. An EvidenceFact is an Observation admitted for a specific claim/purpose under explicit policy. The same Observation may be formative evidence while remaining ineligible for an independent Band/readiness claim. Historical EvidenceFacts remain historical facts; later staleness changes current support rather than rewriting history.

### Claim classes and inference ceilings

| Claim class | Meaning | Ceiling |
|---|---|---|
| sampled capability evidence | one admitted Observation supports only the actually measured Skill/Knowledge/criterion/task slice | never inflated into whole-skill Band/readiness |
| current capability support | current admissible evidence supports a scoped Skill Leaf, Knowledge Object or criterion capability | does not imply whole-skill Band |
| current per-skill Band support | admissible evidence satisfies Band-N threshold for the complete applicable skill construct | requires Band-scoped EvidenceRequirement; not a micro-activity inference |
| task/section readiness | current evidence supports a named task/section/part under material target-like conditions | narrower than whole-skill/full-test readiness |
| target-condition readiness | current evidence supports one declared target condition | exact declared condition, not Planner working target |
| full IELTS readiness | all applicable current learner-evidence conditions for resolved TargetProfile are supported under integrated target-like conditions | never an external-result guarantee or product-support statement by itself |
| historical attainment / external result | point-in-time official/external result or prior internal certification with provenance | current support only when normal Assessment scope/recency admits it |

A broad attempt may contribute several narrow EvidenceFacts, but no label or favorable score broadens inference. Sufficiency is always the applicable versioned EvidenceRequirement.

### EvidenceRequirement contract

A consequential EvidenceRequirement resolves at minimum:

```text
claim class + exact claim scope
threshold / criterion reference where applicable
required target / sub-capability / criterion coverage
eligible Assessment Type(s)
variant / task / section / context applicability
assistance + independence conditions
exposure / novelty / contamination conditions where material
evidence diversity / transfer coverage where material
quantity / stability condition where material
recency / staleness policy for current claims
evaluator / scorer quality + calibration requirement
material conflict handling / discriminating-evidence rule
delivery-condition coverage only when part of the claim
aggregation rule + policy version
```

Non-material conditions are explicitly non-applicable rather than silently omitted. There is no universal attempt count, confidence cutoff, recency window, transfer distance, diversity count or weighted-score formula. Numeric thresholds are permitted only when evidence-backed for the specific claim/evaluator/use. Insufficient empirical calibration leaves higher-consequence support unresolved; it cannot be invented to make the product executable.

Before a `(variant, skill, band)` route can be `SUPPORTED_FOR_PRODUCT`, each high-consequence Band/readiness claim it relies on has a versioned executable EvidenceRequirement. Core must be able to expose claim identity/scope, applicable conditions, admissible EvidenceFacts, satisfied/missing/stale/conflicting/below-threshold conditions, required next evidence where unresolved, evaluator/calibration policy, and policy version without hidden model heuristics.

### ReadinessEvaluation

Exactly five current claim outcomes exist:

```text
INSUFFICIENT_EVIDENCE
CONFLICTING_EVIDENCE
STALE_EVIDENCE
NOT_YET_SUPPORTED
SUPPORTED
```

Missing required evidence → `INSUFFICIENT_EVIDENCE`; materially incompatible valid evidence → `CONFLICTING_EVIDENCE`; previously adequate evidence failing current recency → `STALE_EVIDENCE`; sufficient current evidence positively establishing a scoped requirement is not met → `NOT_YET_SUPPORTED`; all applicable conditions satisfied → `SUPPORTED`. `NOT_YET_SUPPORTED` is not a convenience fallback for missing/uncalibrated evaluation. Product/evaluator inability leaves the learner claim unresolved and is represented separately as CoverageGap/product state. These states never collapse into one mastery/readiness percentage; `SUPPORTED` is not an external IELTS guarantee.

### Assessment Type registry

| ID | Type | Primary role | Scope | Certification relation |
|---|---|---|---|---|
| `AT-01` | Criterion-referenced productive performance | formative/summative productive measurement | Writing, Speaking | may contribute EvidenceFacts; one attempt never certifies alone |
| `AT-02` | Objective receptive item set | keyed receptive measurement | Listening, Reading | contributes only within sampled legitimate scope |
| `AT-03` | Knowledge probe | enabling-knowledge measurement | Knowledge / dependent leaves | cannot replace target-skill evidence |
| `AT-04` | Diagnostic checkpoint | reduce decision-relevant uncertainty | cross-skill / Knowledge | diagnostic role is non-certifying by itself |
| `AT-05` | Mastery portfolio | evaluate accumulated evidence against EvidenceRequirement | any claim scope | certification evidence mechanism |
| `AT-06` | Human / human-verified productive assessment | optional expert verification | Writing, Speaking | may contribute productive EvidenceFacts through normal eligibility |
| `AT-07` | Full mock test | integrated exam-readiness measurement | four skills | non-certifying by itself |

#### Canonical materializer anchor — objective receptive assessment

| ID | Type | Primary role | Scope | Certification relation |
|---|---|---|---|---|
| `AT-02` | Objective receptive item set | keyed receptive measurement | Listening, Reading | may contribute EvidenceFacts within sampled scope |

Productive measurement preserves criterion visibility and evaluator/model/rubric provenance/calibration. Receptive full-section Band inference uses applicable current external scoring rather than a frozen invented table and preserves Academic/GT distinction. Knowledge probes never certify target skills. Diagnostic sampling never cross-infers skills or variants, never fabricates unsampled completeness, keeps quick diagnostic provisional, preserves `sampled`, `not_sampled`, `unusable`, `pending_evaluation` where applicable, and stops based on decision value while leaving uncertainty visible. Diagnostic observations enter higher-consequence claims only through normal eligibility. Human verification remains optional; lack of safe automated evaluator cannot create hidden mandatory human dependency for ordinary product support. Full mock resolves one coherent variant and cannot bypass claim-scoped sufficiency.

### Evidence eligibility and independence

An Observation becomes evidence for a claim only when the task samples that capability; material variant/task/context matches; scoring aligns with canonical criterion/external conversion; material conditions are known; assistance/scaffolding is compatible; response is attributable to learner; evaluator quality is adequate; exposure/retry does not invalidate inference; provenance is sufficient; and use stays within designed inference scope.

Independence is an inference property, not an Attempt-row count. Different Attempt IDs are not independent when success can be reconstructed from prior exposure: same/trivially equivalent prompt, prior answer/model, material corrective feedback then same-task reconstruction, memorized/substantially rehearsed response, generated correction/continuation performing target work, or another material exposure relationship. Contaminated/assisted success may remain useful formative evidence while being ineligible or narrower for unaided/unseen claims. Same-task success after material corrective feedback cannot independently satisfy an unaided requirement; fresh evidence must meet the claim's actual novelty/assistance condition.

Missing evidence is not negative evidence; one wrong response is not automatically an ability gap; stale evidence is not regression; material conflicts remain explicit until discriminating evidence resolves them; historical Observation/EvidenceFact records remain preserved.

### Claim-scoped sufficiency and calibration

Per-skill Band support is a threshold over the complete applicable skill construct, not an average of leaf scores. It composes the Band threshold, applicable Skills/criteria/task conditions, admissible EvidenceFacts and versioned Band-scoped EvidenceRequirement. Micro activities, SRS, completion, Planner estimates, AI opinion and unsupported weighted averages cannot produce Band claims. Unresolved required subscopes remain missing rather than imputed from neighboring strength.

Academic Writing Band scope covers Academic Task 1, Task 2 and applicable criteria; GT Writing covers `W-GT1-01/02/03` + shared quality, Task 2 and criteria, with material recipient/purpose/register transfer where the claim requires it. Academic/GT Task-1 evidence cannot substitute. Speaking Band is whole Parts 1–3. Listening needs admissible timed-section evidence appropriate to the claim. Academic Reading uses Academic conditions/conversion; GT Reading uses GT conversion and sufficient Section-1 everyday / Section-2 workplace / Section-3 longer-general-interest context sampling. Knowledge evidence matches its exact recognition/retrieval/application scope.

Confidence is meaningful only when calibrated against relevant outcomes. Model self-reported confidence is not calibration evidence. Uncalibrated productive judgment cannot support higher consequences. Calibration is scoped by evaluator/model/rubric/task/population/variant where material and provenance is retained across evaluator/policy changes.

## Progression semantics

Progression consumes Assessment claim class/ReadinessEvaluation and never broadens inference. MasteryEstimate is current uncertainty-aware interpretation of one scoped Skill/Knowledge target with support states `unknown`, `learning`, `currently_supported`; `unknown` includes unresolved/stale/conflicting/insufficient/pending cases but underlying reasons remain explicit. `currently_supported` is current evidence interpretation, not permanent mastery.

Per `(skill, band)`, current `BandCertificationState` is exactly:

```text
not_started
in_progress
certified
```

`certified` requires the Band threshold, corresponding Assessment `SUPPORTED`, no remaining required claim blocker, and recorded evidence/policy provenance. A later non-SUPPORTED claim returns current state to `in_progress` while historical certification remains. After evidence/history exists, loss of current support does not return to `not_started`. Certification is internal current support, not official IELTS result.

### Assessment state → GapEvaluation → ActionIntent

Assessment-to-gap consequences:

| Assessment/current condition | Progression consequence |
|---|---|
| `INSUFFICIENT_EVIDENCE` | `EVIDENCE_GAP` → `COLLECT_EVIDENCE` |
| `CONFLICTING_EVIDENCE` | `CONFLICTING_EVIDENCE` → `RESOLVE_CONFLICT` |
| `STALE_EVIDENCE` | `STALE_EVIDENCE` → `REASSESS` |
| `NOT_YET_SUPPORTED` with admissible below-threshold evidence | `ABILITY_GAP` or the more specific demonstrated scaffold/transfer/fluency/exam-condition gap |
| `NOT_YET_SUPPORTED` because another demonstrated non-ability requirement fails | preserve that specific demonstrated condition; do not manufacture `ABILITY_GAP` |
| `SUPPORTED` | infer no deficit; consolidate/advance only where other rules permit |

Canonical GapEvaluation classes are exactly:

```text
ABILITY_GAP
PREREQUISITE_GAP
EVIDENCE_GAP
CONFLICTING_EVIDENCE
STALE_EVIDENCE
SCAFFOLD_DEPENDENCE
TRANSFER_GAP
FLUENCY_GAP
EXAM_CONDITION_GAP
```

`ABILITY_GAP` means current admissible evidence is below target capability; `PREREQUISITE_GAP` is an unresolved Required prerequisite; `EVIDENCE_GAP` is insufficiency without weakness proof; `CONFLICTING_EVIDENCE` is materially incompatible admissible evidence; `STALE_EVIDENCE` requires refresh; `SCAFFOLD_DEPENDENCE` means support carries required independent performance; `TRANSFER_GAP` means lack of generalization into material context; `FLUENCY_GAP` means quality is broadly present but automaticity/speed/rhythm/processing efficiency limits performance; `EXAM_CONDITION_GAP` is a timing/integration/stamina/interface/input/delivery-condition performance gap without necessarily missing underlying capability. Delivery-mode unfamiliarity belongs here when material; it does not create another ability taxonomy.

Canonical ActionIntent values are exactly:

```text
ACQUIRE_PREREQUISITE
REMEDIATE
COLLECT_EVIDENCE
RESOLVE_CONFLICT
REASSESS
FADE_SCAFFOLD
EXPAND_CONTEXT
BUILD_FLUENCY
EXAM_PREPARE
CONSOLIDATE
ADVANCE
```

Default mapping:

| State/gap | ActionIntent |
|---|---|
| `PREREQUISITE_GAP` | `ACQUIRE_PREREQUISITE` |
| `ABILITY_GAP` | `REMEDIATE` or `CONSOLIDATE` according to diagnosis |
| `EVIDENCE_GAP` | `COLLECT_EVIDENCE` |
| `CONFLICTING_EVIDENCE` | `RESOLVE_CONFLICT` |
| `STALE_EVIDENCE` | `REASSESS` |
| `SCAFFOLD_DEPENDENCE` | `FADE_SCAFFOLD` |
| `TRANSFER_GAP` | `EXPAND_CONTEXT` |
| `FLUENCY_GAP` | `BUILD_FLUENCY` |
| `EXAM_CONDITION_GAP` | `EXAM_PREPARE` |
| supported target with useful stability work | `CONSOLIDATE` |
| supported target with next target semantically available | `ADVANCE` |

Required prerequisites hard-gate dependent learning; Recommended is an ordering signal; Independent has no prerequisite relation. Evidence may satisfy a prerequisite without replaying redundant lessons. Uncertainty never routes automatically to remediation.

Exam Preparation may expose higher-demand integrated conditions for diagnosis, pacing, interface familiarity, strategy, stamina and target-condition readiness before certification; it never unlocks Band by exposure, satisfies Required prerequisites by completion, turns a mock into certification, rewrites thresholds, or converts uncertainty to certainty.

Staleness removes current support without proving regression. Conflict requires discriminating evidence. Regression requires later admissible evidence establishing a previously supported current capability is now below requirement; then preserve history, update current MasteryEstimate, move current certification to `in_progress`, classify GapEvaluation, emit ActionIntent, and re-certify only through normal Assessment. Historical certifications remain point-in-time records and are appended/related rather than rewritten.

## Content semantics

Concrete content references canonical learning/exam objects instead of redefining them. Material variant, official family, task/section context, Presentation Class or delivery condition remains explicit; applicability-aware omission is allowed only when genuinely non-applicable and implementations must not fabricate Band/context/family/presentation/delivery/curriculum identity to fit a schema.

Content has stable logical identity and immutable semantic revisions. Material change to stimulus, prompt/instruction, answer key, rubric/model, target binding, official family/context/presentation, response contract or another meaning/scoring/inference field creates a new ContentRevision. Historical Attempt/evidence resolves the exact revision delivered. Origin (authored/imported/deterministic/AI/learner/media) is provenance, not quality or activation authority.

ValidationDecision is historical evidence/decision about a ContentRevision under named policy/intended-use scope. Revalidation of unchanged content creates another decision, not another content revision. Later policy/decision does not rewrite earlier decision meaning. Current eligibility resolves current policy + intended-use scope + compatible decisions/findings; globally newest decision is not universal authority. Generator/validator self-confidence is input only.

### Similarity and exposure

Similarity facts are scoped at least across `Stimulus↔Stimulus`, `Prompt↔Prompt`, `PracticeItem↔PracticeItem`, `AssessmentItem↔AssessmentItem`, with distinguishable exact identity, normalized-text identity, shared source/stimulus, semantic prompt similarity, response/reasoning-pattern similarity, source-content similarity and structural/template similarity. Shared canonical target/family/topic/context is not duplication by itself. Shared Stimulus may be legitimate while prior exposure still contaminates an unseen claim. Genuine transfer requires materially relevant variation, not merely new IDs/cosmetic wording. Numerical similarity cutoffs are versioned implementation/calibration policy. Corpus redundancy and learner-specific novelty are distinct. Similarity facts may support deliberate repetition while narrowing transfer/readiness evidence.

ExposureContext preserves whether item/stimulus revision was seen, prior feedback/model exposure, similarity/novelty facts, prior attempts, context variation, Content Context variation and external-family/presentation variation where material. Reservation/delivery attempt is not actual exposure. `UNKNOWN`/ambiguous exposure is not proof of unseen for a consequence requiring independence; such use remains unresolved/ineligible while training not requiring novelty may remain allowed.

### Stable Content Context registry

| ID | Meaning | Variant |
|---|---|---|
| `CTX-LISTENING-SHARED` | shared IELTS Listening construct/context | shared |
| `CTX-READING-ACADEMIC` | Academic Reading passage/set context | Academic |
| `CTX-READING-GT-S1-EVERYDAY` | GT Reading Section 1 everyday context | General Training |
| `CTX-READING-GT-S2-WORKPLACE` | GT Reading Section 2 workplace context | General Training |
| `CTX-READING-GT-S3-GENERAL-INTEREST` | GT Reading Section 3 longer general-interest context | General Training |
| `CTX-WRITING-ACADEMIC-T1-VISUAL` | Academic Writing Task 1 visual-information context | Academic |
| `CTX-WRITING-GT-T1-LETTER` | GT Writing Task 1 letter context | General Training |
| `CTX-WRITING-T2` | Writing Task 2 construct; prompt remains variant-scoped where material | shared/core |
| `CTX-SPEAKING-SHARED` | shared IELTS Speaking construct | shared |

IDs remain stable across content/manifests/contracts. New CTX requires materially distinct external/content inference context, never merely topic/difficulty/Band/delivery/screen. Delivery is orthogonal. Context-neutral Knowledge/foundation content may legitimately omit CTX; filler `CTX-GENERIC`/`CTX-KNOWLEDGE` is forbidden. When family/variant/task/section changes validity, coverage, scoring or inference, compatible CTX becomes required.

### Stable Content Presentation Class registry

Listening `IELTS-L-QF-04`:

- `PRES-L-QF04-FORM` — form completion
- `PRES-L-QF04-NOTE` — note completion
- `PRES-L-QF04-TABLE` — table completion
- `PRES-L-QF04-FLOW-CHART` — flow-chart completion
- `PRES-L-QF04-SUMMARY` — summary completion

Reading `IELTS-R-QF-09`:

- `PRES-R-QF09-SUMMARY` — summary completion
- `PRES-R-QF09-NOTE` — note completion
- `PRES-R-QF09-TABLE` — table completion
- `PRES-R-QF09-FLOW-CHART` — flow-chart completion

Academic Writing Task 1 `IELTS-W-A-T1`:

- `PRES-W-A-T1-GRAPH-CHART-TABLE` — graph/chart/table/statistical visual, including combined statistical displays
- `PRES-W-A-T1-DIAGRAM-PROCESS` — process/object/device/event or comparable diagram representation
- `PRES-W-A-T1-MAP-PLAN` — map/plan/spatial-change representation

Presentation Class is content-coverage identity, not Skill/Practice Type/scored task. A stimulus may combine classes genuinely; topic/accent/theme/difficulty/Band never creates presentation ID. New PRES requires a material interaction/content-coverage distinction.

### Concrete content object invariants

`LearningUnit` groups coherent curriculum purpose and may reference node/objectives/prerequisites, variant, contexts/families/presentations, delivery scope, content sequence, Practice/Assessment items, Error/Remediation Patterns and completion intent; completion is not mastery.

`Stimulus` is material the learner reads/hears/views/analyses/responds to, carrying material variant/context/family/presentation, source/provenance, content, language/difficulty and rights metadata. Reuse is allowed only when it does not invalidate intended independence/novelty/inference.

`ScaffoldingProfile` preserves content, structural, lexical and response support; hints; worked-example access; feedback timing; retry support; timing support. Support is not inherently invalid, but it must remain visible because it changes inference.

`PracticeItem` references Practice Type/mechanisms, material variant/context/family/presentation/delivery, canonical Skill/Knowledge targets, node when specifically bound, stimuli, prompt, response contract, difficulty, scaffolding, exposure, error/remediation and feedback/answer-model references. Family/context/presentation may not be inferred only from broad Skill when materially distinct; Academic/GT Task-1 target identities never cross-substitute; reusable/direct-browse items need not fabricate Curriculum Node.

`AssessmentItem` additionally preserves Assessment Type, claim scope, target Band only when Band-scoped, scoring/rubric reference, conditions and independence group. It never decides mastery/certification. Knowledge/non-Band assessment must not invent a target Band.

`ErrorPattern` is a reusable tutoring hypothesis, not proof a learner has the error. `RemediationPattern` is reusable strategy referencing target/mechanism/type/scaffold/success checks and cannot redefine prerequisites/Band/mastery. `FeedbackArtifact` preserves observed performance, selected/deferred targets, gap/error/remediation and provenance without becoming evidence or mastery authority. `Attempt`, `Observation` and `EvidenceFact` retain the exact separation above; no umbrella EvidenceRecord hides it.

### Content coverage manifest and validation

When executable content exists, a machine-checkable manifest/equivalent derived index must answer, where applicable:

```text
content/revision identity + lineage
canonical target refs
official task/question-family refs + applicability
Content Context ref + applicability
material Presentation Class refs + applicability
test variant
supported delivery scope
Practice/Assessment Type support
implemented response interaction
answer-key/rubric/evaluator route
difficulty/transfer classes
origin/provenance + rights state
applicable validation decision/policy refs
independent readiness asset state
product/release activation
```

Conditional dimensions distinguish `applicable + present`, `explicitly not applicable`, and `required but missing/unresolved`; absence alone is never N/A. Official-family/Presentation-Class coverage is checked independently of Skill coverage. One related family/template/subformat cannot silently satisfy another.

Validation is origin-neutral, applicability-aware and consequence-aware. Every exposed ContentRevision passes applicable universal hard gates: canonical reference resolution/compatibility; executable structure/response contract; compatible family/context/presentation/variant/delivery where material; rights/privacy/security; answer/scoring/rubric correctness where applicable; productive prompt/response/rubric/evaluator compatibility; no prohibited leakage; sufficient provenance/integrity; no contradictory teaching/scoring rule. Open productive tasks do not require one objective answer; they require compatible construct/response/rubric/evaluator semantics.

Lower-consequence training still passes universal gates plus quality adequate for its mechanism/type. Transfer-sensitive use adds context/novelty/exposure validity. Evidence-candidate use adds construct/scoring/evaluator validity and observable conditions; Assessment still admits evidence. Readiness/Band use adds the strongest applicable calibration/provenance/independence/scoring/sufficiency rules. Generated content has neither weaker nor stronger authority; AI self-check never establishes validity by itself. No global quality percentage, validator count, similarity/confidence threshold or audit-sample count is frozen here.