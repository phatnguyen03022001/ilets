STATUS: CANONICAL
OWNS: measurement semantics, Assessment Type taxonomy, Observation→EvidenceFact admission, diagnostic sampling semantics, claim-scoped EvidenceRequirement, confidence/calibration policy, readiness evaluation, and certification evidence construction
DEPENDS_ON: 02-IELTS-MODEL.md, 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md
DOES_NOT_OWN: Skill/Knowledge definitions, curriculum ordering, practice pedagogy, learner-state transitions, band-progression decisions, product UX durations, or concrete assessment-item storage

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

## EvidenceRequirement

An EvidenceRequirement defines the logical conditions that must hold for a scoped claim to be `SUPPORTED`.

Depending on the claim, conditions may include:

- target/criterion coverage;
- variant/task/section/context coverage;
- task/context diversity;
- independence from material support;
- consistency when stable performance is claimed;
- recency when current capability is claimed;
- transfer/generalisation;
- evaluator/scoring quality;
- resolution of material conflicts;
- required part/task/section coverage;
- delivery-condition coverage only when the claim explicitly includes that delivery condition.

There is **no universal attempt count, confidence cutoff, recency window, or transfer distance** for every target.

Numeric thresholds are calibration policy. They must be evidence-backed and versioned for the claim/evaluator where used.

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
- `NOT_YET_SUPPORTED` — current relevant evidence exists and is below the requirement.
- `SUPPORTED` — current admissible evidence satisfies the scoped requirement.

These states must never be collapsed into one mastery/readiness percentage.

`SUPPORTED` is an internal evidence statement, not a guaranteed future official IELTS result.

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

Human expert scoring/verification is an optional escalation route. It may verify or replace an uncertain automated judgment but is not automatically required for every productive attempt.

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

# Missing, negative, conflicting, and stale evidence

- Missing evidence is not negative evidence.
- One wrong response is not automatically an ability gap.
- Stale evidence is not proof of regression.
- Material conflicts remain explicit until discriminating evidence resolves them; do not average them away.
- Historical Observations/EvidenceFacts remain immutable enough for governed reinterpretation.

# Claim-scoped sufficiency

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
- uncertainty may trigger re-sampling, alternative scoring, or human verification;
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

If no eligible evaluation route exists, the truthful result is pending/unavailable/insufficient evidence—not a fake score or silent lower-quality fallback.

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
confidence/calibration refs where relevant
policy/evaluator version
```

`09-PROGRESSION.md` consumes this result. Progression does not rescore observations or create a second evidence policy.