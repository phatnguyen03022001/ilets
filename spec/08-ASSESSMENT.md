STATUS: SUPERSEDED_NON_CANONICAL
OWNS: measurement semantics, Assessment Type taxonomy, Observation→EvidenceFact admission, claim classes and inference ceilings, diagnostic sampling semantics, claim-scoped EvidenceRequirement, confidence/calibration policy, readiness evaluation, and certification evidence construction
DEPENDS_ON: 01-LEARNER-MODEL.md, 02-IELTS-MODEL.md, 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md
DOES_NOT_OWN: Skill/Knowledge definitions, curriculum ordering, practice pedagogy, learner-state transitions, band-progression decisions, TargetProfile product composition, product coverage/support state, product UX durations, or concrete assessment-item storage

# 08 — Assessment

## Purpose

Define how demonstrated capability is measured and how observations become evidence for a scoped claim.

The required separation is:

```text
Attempt
  ↓
Observation
  ↓ eligibility / interpretation
EvidenceFact
  ↓ claim-scoped aggregation
ReadinessEvaluation
  ↓
Progression decision owned by 09-PROGRESSION
```

An Attempt is not automatically evidence. Evidence is not automatically mastery. Mastery is not automatically exam readiness.

# Core objects

## Observation

An Observation records what was measured before evidence admission. It preserves the result, conditions, scorer/evaluator output, assistance/scaffolding, exposure, variant/context, delivery conditions where material, and provenance required to interpret the attempt.

## EvidenceFact

An EvidenceFact is an Observation admitted for a specific claim/purpose under an explicit eligibility policy.

Eligibility is claim-scoped. The same Observation may be valid formative evidence and invalid for an independent Band-readiness claim.

Historical EvidenceFacts remain historical facts. Later staleness changes what they support **now**; it does not rewrite what happened.

# Claim classes and inference ceilings

Assessment distinguishes the measurement record from the consequence being claimed. The minimum claim classes are:

| Claim class | Meaning | Minimum inference boundary |
|---|---|---|
| sampled capability evidence | one admitted Observation supports only the Skill Leaf/Knowledge/criterion/task slice actually measured | never inflated into a whole-skill Band/readiness claim |
| current capability support | current admissible EvidenceFacts support a scoped Skill Leaf, Knowledge Object, or criterion-level capability | does not imply whole-skill Band |
| current per-skill Band support | admissible evidence satisfies the Band-N threshold for the complete applicable skill construct | requires a Band-scoped EvidenceRequirement; not derived from a micro/sampled activity |
| task/section readiness | current evidence supports performance under a named task/section/part and material target-like conditions | narrower than whole-skill or full-test readiness |
| target-condition readiness | current evidence supports one declared target condition, such as a per-skill minimum or delivery-condition requirement | evaluated against the exact declared condition, not a planner working target |
| full IELTS readiness | all applicable current learner-evidence conditions for the resolved TargetProfile are supported under integrated target-like conditions | does not mean guaranteed external result or product support by itself |
| historical attainment / external result | point-in-time external result or prior internal certification reference with provenance | Assessment decides only whether/how it may support a current claim; history ownership remains separate |

An Assessment attempt can always produce an Observation if measurement succeeds. It may produce zero or more EvidenceFacts when claim-scoped admission succeeds. One attempt may therefore prove a narrow sampled fact; it never gains a broader inference merely because the activity was labelled assessment/readiness or produced a high score.

A broad integrated attempt may contribute evidence across several required subscopes, but no architecture rule assumes that one attempt is sufficient for a per-skill Band or readiness claim. Sufficiency is determined only by the applicable versioned EvidenceRequirement and its evidence-backed calibration.

## EvidenceRequirement

An EvidenceRequirement defines the logical conditions that must hold for a scoped claim to be `SUPPORTED`. It is a reusable semantic policy, not a database/API schema.

A consequential requirement resolves at minimum:

```text
claim class + exact claim scope
threshold/criterion reference when applicable
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

A condition that is not material to a claim is explicitly non-applicable; omission must not silently weaken a consequential requirement.

Quantity, consistency, diversity, recency, novelty, evaluator-quality, and calibration thresholds may be numeric only when justified by evidence for that claim/evaluator/use. There is **no universal attempt count, confidence cutoff, recency window, transfer distance, or weighted score formula** for every target.

Where empirical calibration is not yet sufficient, the requirement records that calibration is required and the higher-consequence claim cannot become `SUPPORTED` merely to make the product executable.

# EvidenceRequirement materialisation gate

Before a `(variant, skill, band)` target can be `SUPPORTED_FOR_PRODUCT`, every high-consequence readiness/Band claim used by that target must resolve to a **versioned executable EvidenceRequirement**.

The executable policy must let Core API answer without a hidden model heuristic:

```text
claim identity and scope
applicable variant/task/context conditions
admissible EvidenceFacts
satisfied conditions
missing conditions
stale conditions
material conflicts
below-threshold conditions
required next evidence / refresh / discriminating condition when unresolved
applicable evaluator/calibration policy
policy version
```

Materialisation makes the canonical policy executable; it does not invent a universal threshold.

# ReadinessEvaluation

A claim evaluation returns one of:

```text
INSUFFICIENT_EVIDENCE
CONFLICTING_EVIDENCE
STALE_EVIDENCE
NOT_YET_SUPPORTED
SUPPORTED
```

- `INSUFFICIENT_EVIDENCE` — required evidence is missing.
- `CONFLICTING_EVIDENCE` — material valid evidence supports incompatible interpretations.
- `STALE_EVIDENCE` — historical support exists but is not sufficiently current for the present claim.
- `NOT_YET_SUPPORTED` — current relevant evidence does not satisfy the scoped requirement.
- `SUPPORTED` — current admissible evidence satisfies the scoped requirement.

`NOT_YET_SUPPORTED` is a claim-scoped evidence conclusion. It is not universal proof that the underlying capability is absent or false outside that scope.

These states must never be collapsed into one mastery/readiness percentage.

`SUPPORTED` is an internal evidence statement, not a guaranteed future official IELTS result.

State interpretation is condition-aware:

- missing required subscopes, missing admissible samples, or unresolved required evaluator output → `INSUFFICIENT_EVIDENCE`;
- current admissible evidence that materially disagrees → `CONFLICTING_EVIDENCE`;
- previously adequate evidence whose recency requirement no longer holds → `STALE_EVIDENCE`;
- current admissible evidence that is sufficient to conclude the scoped threshold/condition is not met → `NOT_YET_SUPPORTED`;
- every applicable condition satisfied → `SUPPORTED`.

`NOT_YET_SUPPORTED` must not be used as a convenience fallback for missing/uncalibrated evaluation. A learner ability gap is justified only when admissible evidence positively establishes below-requirement performance for that scope.

If product/evaluator capability is unavailable for a required consequence, the learner claim remains unresolved from the available evidence and the product separately exposes the applicable CoverageGap. Product inability is not negative learner evidence.

# Assessment Type registry

| ID | Type | Primary role | Scope | Certification relation |
|---|---|---|---|---|
| `AT-01` | Criterion-referenced productive performance | formative/summative productive measurement | Writing, Speaking | may contribute EvidenceFacts; one attempt never certifies by itself |
| `AT-02` | Objective receptive item set | keyed receptive measurement | Listening, Reading | may contribute EvidenceFacts within sampled scope |
| `AT-03` | Knowledge probe | enabling-knowledge measurement | Knowledge / dependent leaves | cannot replace target-skill evidence |
| `AT-04` | Diagnostic checkpoint | reduce decision-relevant uncertainty | cross-skill / Knowledge | non-certifying as a diagnostic role |
| `AT-05` | Mastery portfolio | evaluate accumulated evidence against EvidenceRequirement | any claim scope | certification evidence mechanism |
| `AT-06` | Human / human-verified productive assessment | optional expert verification | Writing, Speaking | may contribute productive EvidenceFacts |
| `AT-07` | Full mock test | integrated exam-readiness measurement | four skills | non-certifying by itself |

# Type-specific rules

## `AT-01` Productive performance

Writing/Speaking measurement preserves criterion-level visibility. A global score cannot hide a materially weak required criterion.

Automated judgment records evaluator/model/rubric version and uncertainty/provenance sufficient for calibration and audit.

## `AT-02` Receptive item set

Uses deterministic keys where valid. Full-section Band inference uses the applicable current external scoring policy rather than an invented permanent conversion.

Reading observations preserve Academic/GT variant and material section/context. Academic and GT conversion policies are not interchangeable.

## `AT-03` Knowledge probe

Measures enabling knowledge. Passing a vocabulary, grammar, or phonology probe cannot certify Listening, Reading, Writing, or Speaking.

## `AT-04` Diagnostic checkpoint

Diagnostic purpose is to reduce **decision-relevant uncertainty**, not to force a complete measurement of every capability before learning can begin.

Required invariants:

1. **No cross-skill inference** — evidence from one scored skill cannot establish another scored skill merely because the skills correlate.
2. **No variant substitution** — when a target includes a material Academic/GT-specific condition, diagnostic coverage must either sample it or leave it explicitly unresolved.
3. **No fabricated completeness** — an unsampled, unusable, failed-capture, or evaluator-pending condition is reported as unresolved evidence state, not converted to weakness or an estimated score.
4. **Quick diagnostic is provisional** — it may stop after enough information exists to produce a useful first route while preserving all unresolved target conditions.
5. **Fuller baseline is still evidence-bounded** — completion means the baseline sampling flow ended; it does not mean every target claim became known.
6. **Sampling ledger is explicit** — for each material target condition, the diagnostic result can distinguish at least `sampled`, `not_sampled`, `unusable`, and `pending_evaluation` where applicable.
7. **Stop by decision value, not arbitrary completeness** — additional sampling may stop when expected decision value is low relative to learner burden, provided the unresolved state remains visible.
8. **Certification remains normal Assessment** — a diagnostic Observation contributes to a higher-consequence claim only if the normal eligibility and EvidenceRequirement independently admit it.

Diagnostic output answers:

```text
what appears currently supported
what is observed below requirement
what is unknown/insufficient
what is conflicting
what is stale
what material condition should be sampled next
```

It does not directly emit learner remediation without Progression interpreting the evidence state.

## `AT-05` Mastery portfolio

Applies a versioned EvidenceRequirement over admissible EvidenceFacts. It never certifies by mechanically counting attempts.

Repeated demonstrations, multiple contexts, independence, recency, transfer, criterion coverage, or variant coverage are required only when the claim makes them material.

## `AT-06` Human verification

Human expert scoring/verification is an optional external/secondary evidence route. It may contribute productive EvidenceFacts when provenance, rubric fit, and normal claim-scoped eligibility are satisfied.

No ordinary `SUPPORTED_FOR_PRODUCT` learner route may require hidden mandatory human scoring merely because the deterministic/automated evaluator is insufficient. If the product cannot safely evaluate the desired consequence without human input, that automated consequence remains unsupported/unresolved until a calibrated supported evaluator path exists. Optional coaching, learner-imported expert feedback, or explicitly chosen human-service extensions do not change this invariant.

## `AT-07` Full mock

A full mock resolves one valid IELTS variant before Reading and Writing Task 1 are instantiated/scored.

It measures integrated readiness under exam-like conditions. One mock cannot bypass claim-scoped sufficiency or automatically certify a skill/Band.

# Evidence eligibility

An Observation becomes an EvidenceFact for a claim only when:

1. the task samples the claimed capability;
2. variant/task/context matches where material;
3. scoring aligns with the canonical criterion and applicable external conversion;
4. material conditions are known;
5. assistance/scaffolding is compatible with the claim;
6. the response is attributable to the learner;
7. evaluator quality is adequate for the intended consequence;
8. exposure/retry history does not invalidate the inference;
9. provenance is sufficient for audit/recomputation;
10. the Observation is not used outside its designed inference scope.

A guided attempt or same-item retry may be valuable learning evidence while remaining invalid for an independent unseen-transfer claim.

# Evidence independence and contamination

Independence is a property of the **inference supported by demonstrations**, not a count of Attempt rows.

Two performances are not materially independent merely because they have different Attempt IDs. Independence is narrowed when success can be reproduced from prior exposure rather than demonstrated capability, including material cases such as:

- the same or trivially equivalent prompt/stimulus;
- access to a prior answer/model response;
- material corrective feedback followed by reconstruction of the same task;
- memorized or substantially rehearsed response reuse;
- generated correction/continuation that performs material target work for the learner;
- another exposure relationship that makes the later success non-independent for the scoped claim.

The consequence is claim-scoped. Contaminated or assisted success may remain valid formative/recovery evidence while being ineligible, or carrying narrower inference scope, for a claim requiring unaided independent or unseen-transfer performance.

A same-task success after material corrective feedback cannot independently satisfy an EvidenceRequirement that requires unaided performance. Fresh evidence must use conditions whose novelty/assistance are sufficient for that claim; exact diversity thresholds remain versioned calibration policy rather than a universal attempt rule.

# Missing, negative, conflicting, and stale evidence

- Missing evidence is not negative evidence.
- One wrong response is not automatically an ability gap.
- Stale evidence is not proof of regression.
- Material conflicts remain explicit until discriminating evidence resolves them; do not average them away or select the more convenient interpretation.
- Historical Observations/EvidenceFacts remain immutable enough for governed reinterpretation.

# Claim-scoped sufficiency

## Capability → Band derivation

A per-skill Band claim is a threshold claim over the applicable **whole skill construct**, not an average of leaf scores. The derivation is:

```text
Band threshold from 05-BANDS
+ applicable Skill Leaves / criteria / task-part conditions
+ admissible EvidenceFacts for those subscopes
+ versioned Band-scoped EvidenceRequirement
→ current per-skill Band ReadinessEvaluation
```

Rules:

1. leaf/criterion support may satisfy required subconditions but does not itself numerically compose a Band;
2. a sampled micro-activity, SRS result, completion state, planner estimate, AI opinion, or unsupported weighted average cannot produce a Band claim;
3. receptive Band inference uses the applicable official section scoring/conversion only when the measured scope legitimately supports that inference;
4. productive Band inference preserves criterion/task/part coverage and evaluator calibration required by the Band-scoped policy;
5. unresolved required sub-capability/task/criterion/context remains an explicit missing condition rather than being imputed from stronger neighboring evidence;
6. support for Band N does not imply every canonical leaf is error-free; `05-BANDS.md` owns the threshold and permitted residual limitations.

A lower Band may be supported while a higher Band remains unresolved or not yet supported. Assessment does not infer hidden intermediate Bands unless their own threshold/policy conditions are satisfied.

## Academic Writing

A Band claim covers:

- Academic Task 1;
- Task 2;
- applicable Writing criteria;
- independence/recency/transfer conditions required by its EvidenceRequirement.

Academic visual Task-1 evidence cannot be replaced by GT letter evidence.

## General Training Writing

A Band claim covers:

- GT Task-1 capability `W-GT1-01`, `W-GT1-02`, `W-GT1-03` plus shared Writing quality;
- Task 2;
- applicable Writing criteria;
- independence/recency/transfer conditions required by its EvidenceRequirement.

Where transferable letter control is claimed, evidence must vary recipient/purpose/register materially. Academic visual Task-1 evidence cannot satisfy this condition.

## Speaking

A Speaking Band claim represents the whole construct across Parts 1, 2, and 3 rather than one rehearsed slice.

Exact demonstration count and recency are versioned calibration policy, not architectural constants.

## Listening

A Listening Band claim requires admissible timed section evidence whose scoring/context support the claimed Band. A short practice set cannot be arithmetically inflated into a full-section Band claim.

## Academic Reading

A Band claim uses Academic content/conditions and the applicable Academic score conversion for full-section inference.

## General Training Reading

A Band claim uses the GT score conversion and enough GT context sampling for the scoped readiness claim, including the material distribution across:

```text
Section 1  everyday
Section 2  workplace
Section 3  longer general-interest
```

Shared Reading capability observed on Academic material may support compatible leaf-level claims but cannot by itself establish whole GT Reading readiness.

## Knowledge

Knowledge evidence matches the object and inference required. Recognition, retrieval, and productive application may support different scopes.

# Confidence and calibration

Confidence is meaningful only when calibrated against relevant outcomes.

Rules:

- model self-reported confidence is not calibration evidence;
- uncalibrated/low-quality productive judgment cannot support a higher-consequence claim by itself;
- uncertainty may trigger re-sampling, another supported evaluator/scoring route, or optional human verification; it may also leave the claim unresolved;
- calibration is scoped by evaluator/model/rubric/task/population/variant where material;
- policy/evaluator changes preserve provenance so historical observations can be reinterpreted without rewriting them.

# Assessment roles are not interchangeable

- **Formative** — guides ongoing learning.
- **Diagnostic** — reduces uncertainty/classifies evidence need.
- **Readiness** — estimates performance under target conditions.
- **Certification evidence** — `AT-05` evaluates whether the scoped claim is `SUPPORTED`.

A full mock can teach without certifying. A diagnostic can expose a strong sample without becoming an automatic Band claim.

# Evidence collection strategy

Collect new evidence when it can materially change a decision:

- insufficient evidence → smallest useful missing sample;
- conflict → discriminating assessment;
- staleness → representative refresh;
- observed below requirement → re-evidence after useful intervention or when diagnosis needs refinement;
- supported claim → stop when added evidence has poor value relative to learner burden unless refresh is required.

# AI assessment boundary

AI may automate feedback and measurement; it is not assessment authority.

Objective items use deterministic keys where possible. Productive evaluation must be benchmarked, versioned, uncertainty-aware, and fail closed above demonstrated evaluator quality.

If no eligible evaluation route exists, the truthful measurement/result state is pending or unavailable and the learner claim remains unresolved/insufficiently evidenced. The product separately exposes the applicable evaluator/product CoverageGap; it does not emit a fake score, mandatory hidden-human dependency, or silent lower-quality fallback.

# Output consumed by Progression

A claim evaluation exposes:

```text
claim_scope
variant/context scope where material
readiness_status
supporting_evidence_fact_refs
blocking_conditions
material_conflicts
recency_state
required_next_evidence when unresolved
confidence/calibration refs where relevant
policy/evaluator version
```

`09-PROGRESSION.md` consumes this result. Progression does not rescore observations or create a second evidence policy.