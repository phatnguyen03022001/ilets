STATUS: CANONICAL
OWNS: how demonstrated capability is measured, Assessment Type taxonomy, Observation→EvidenceFact admission, claim-scoped evidence requirements, confidence/calibration policy, validity, readiness evaluation, and certification evidence construction
DEPENDS_ON: 02-IELTS-MODEL.md, 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md
DOES_NOT_OWN: Skill/Knowledge definitions, curriculum ordering, practice pedagogy, learner-state transitions, band-progression decisions, or concrete assessment-item storage

# 08 — Assessment

## Purpose

Define how demonstrated capability is measured and how evidence supports a scoped learning claim.

Assessment must preserve the distinction:

```text
Attempt
  ↓
Observation
  ↓  eligibility / interpretation
EvidenceFact
  ↓  claim-scoped aggregation
Mastery / Readiness evaluation
  ↓
Progression decision owned by 09-PROGRESSION
```

An Attempt is not automatically evidence. Evidence is not automatically mastery. Mastery is not automatically exam readiness.

## Core distinctions

### Observation

An Observation is what was measured from a learner attempt before evidence admission. It preserves the response/result, conditions, scorer output, timing, assistance/scaffolding context, exposure context, and provenance needed to interpret what happened.

### EvidenceFact

An EvidenceFact is an Observation admitted for a specific claim/purpose under an explicit eligibility rule.

Eligibility is claim-scoped. The same Observation may be useful formative evidence but invalid for an independent Band-readiness claim.

Historical EvidenceFacts remain historical facts. Later staleness changes what they support now; it does not rewrite what happened.

### EvidenceRequirement

An EvidenceRequirement defines what must be true for a scoped claim to be supported.

Requirements may include:

- target/criterion coverage;
- appropriate task and context coverage;
- independence from excessive support;
- consistency where the claim requires stable performance;
- recency where current capability matters;
- transfer/generalization where the claim implies it;
- minimum evaluator/scoring quality;
- resolution of material conflicts;
- required task/part coverage for integrated skills.

There is **no universal attempt count, confidence cutoff, recency window, or transfer distance** that applies to every target.

Numeric thresholds are calibration policy and must be evidence-backed for the claim/evaluator where they are used. They are not architectural constants.

### ReadinessEvaluation

A claim-scoped evaluation uses one of these semantic outcomes:

```text
INSUFFICIENT_EVIDENCE
CONFLICTING_EVIDENCE
STALE_EVIDENCE
NOT_YET_SUPPORTED
SUPPORTED
```

- `INSUFFICIENT_EVIDENCE` — required observations are missing.
- `CONFLICTING_EVIDENCE` — material valid evidence points in incompatible directions and the conflict is unresolved.
- `STALE_EVIDENCE` — historical support exists but is not sufficiently current for the present claim.
- `NOT_YET_SUPPORTED` — relevant current evidence exists and is below the requirement.
- `SUPPORTED` — current admissible evidence satisfies the scoped requirement.

These states must not be collapsed into a single percentage.

`SUPPORTED` means the internal learning claim is currently supported. It is never a guaranteed future official IELTS result.

# Canonical Assessment Type registry

| ID | Type | Kind | Scope | Evidence / role | Certification status |
|---|---|---|---|---|---|
| `AT-01` | Criterion-referenced productive performance | formative or summative | Writing, Speaking | complete productive performance scored by relevant criteria with uncertainty/provenance where automated | may contribute EvidenceFacts; never certifies by one attempt alone |
| `AT-02` | Objective receptive item set | formative or summative | Listening, Reading | keyed item results, raw score, current official section-band conversion when appropriate | may contribute EvidenceFacts within its sampled context |
| `AT-03` | Knowledge probe | formative | Knowledge Objects / dependent leaves | retrieval, recognition, judgment, short production, targeted application | supports Knowledge claims; cannot replace target-skill evidence |
| `AT-04` | Diagnostic checkpoint | diagnostic | cross-skill / knowledge | samples current strength, gap, uncertainty, and next evidence need | non-certifying |
| `AT-05` | Mastery portfolio | cumulative claim evaluator | any target skill/band/knowledge scope | applies the relevant EvidenceRequirement to accumulated EvidenceFacts | **certification evidence mechanism** |
| `AT-06` | Human / human-verified productive assessment | optional summative/verification | Writing, Speaking | expert-scored performance or expert verification of uncertain automated judgment | may contribute valid productive EvidenceFacts |
| `AT-07` | Full mock test | summative readiness | all four skills | integrated exam-like observations and section estimates | non-certifying by itself; contributes only within normal eligibility/sufficiency rules |

## Type semantics

### `AT-01` productive performance

Writing/Speaking assessment must preserve criterion-level visibility. A global score cannot hide a materially weak required criterion.

Automated productive judgment must expose evaluator/model/rubric version and uncertainty sufficient for calibration and audit.

### `AT-02` receptive item set

Uses deterministic keys where valid. Raw-score conversion references current external truth in `02-IELTS-MODEL.md`; test-version variation must not be hidden behind an invented immutable universal conversion.

### `AT-03` knowledge probe

Measures enabling knowledge only. Passing a vocabulary, grammar, or phonology probe cannot certify an IELTS skill.

### `AT-04` diagnostic checkpoint

Answers what appears supported, below requirement, unknown, conflicting, or stale and what should be sampled next. Diagnosis is not a weakness generator.

### `AT-05` mastery portfolio

Evaluates a claim-specific EvidenceRequirement over admissible EvidenceFacts. It does not count attempts mechanically.

A portfolio may require repeated demonstrations, multiple task contexts, criterion coverage, independence, or transfer when the claim makes those conditions material. The exact logical conditions belong to the claim specification, not a universal `N=2` rule.

### `AT-06` human verification

Optional escalation for productive skills. Human review may verify/replace an uncertain automated judgment but is not mandatory for the core learning loop.

### `AT-07` full mock

Measures broad readiness under integrated conditions. One mock may provide useful observations and EvidenceFacts; it cannot bypass claim-scoped sufficiency or automatically certify a skill/band.

# Evidence eligibility

An Observation becomes an EvidenceFact only when:

1. the task samples the claimed capability;
2. scoring aligns with the canonical criterion;
3. material conditions are known;
4. assistance/scaffolding is compatible with the claim;
5. the learner response is attributable to the learner;
6. evaluator quality is adequate for the intended consequence;
7. exposure/retry history does not invalidate the inference being made;
8. provenance is sufficient for audit/recomputation;
9. the observation is not used outside its designed scope.

A same-item retry may support recovery or retention while remaining weak evidence for unseen transfer. Guided success may support learning progress while remaining invalid for independent readiness.

## Negative, missing, conflicting, and stale evidence

- Missing evidence is not negative evidence.
- One wrong response is not automatically proof of an ability gap.
- Stale evidence is not proof of regression.
- Material conflicts are preserved until a discriminating assessment resolves them; do not average them away.
- Historical observations and EvidenceFacts remain immutable enough for later governed reinterpretation.

# Claim-scoped sufficiency

## Productive skills

A Writing/Speaking Band claim must include enough admissible performance to cover the construct required by that claim.

Writing Band claims must represent both Task 1 and Task 2 and preserve criterion-level outcomes. Speaking Band claims must represent the whole construct across Parts 1, 2, and 3 rather than one rehearsed slice.

Repeated independent demonstrations are generally expected for a stable high-consequence claim, but the exact count and recency requirement are calibration/policy decisions justified by evidence rather than universal constants.

## Receptive skills

Listening/Reading Band claims require admissible timed section evidence whose official conversion and sampled context support the claimed Band. The number of qualifying demonstrations, recency window, and diversity requirement depend on the claim and calibration evidence.

A short practice set cannot be arithmetically inflated into a full-section Band claim.

## Knowledge

Knowledge claims require evidence appropriate to the object and intended use. Recognition, retrieval, and productive application may carry different inference scope. No universal accuracy percentage overrides a more suitable criterion.

# Confidence and calibration

Confidence is meaningful only when it has been calibrated against relevant outcomes.

Rules:

- model self-reported confidence is not calibration evidence;
- low-quality or uncalibrated productive judgments cannot support high-consequence certification by themselves;
- uncertainty may trigger re-sampling, alternative scoring, or optional human verification;
- calibration is scoped by evaluator/model/rubric/task/population where material;
- threshold changes must preserve provenance so historical EvidenceFacts can be reinterpreted without rewriting observations.

# Practice, readiness, and certification

- **Formative assessment** guides learning.
- **Diagnostic assessment** reduces uncertainty or identifies the kind of gap.
- **Readiness assessment** estimates performance under target conditions.
- **Certification evidence** uses `AT-05` to determine whether a scoped learning claim is `SUPPORTED`.

These roles are not interchangeable.

A learner may learn from a full mock without becoming certified. A learner may hold a current certification while one later poor attempt creates conflict or a need for refresh rather than immediate automatic regression.

# Evidence collection strategy

Collect new evidence when it can materially change a decision.

- insufficient evidence → collect the smallest useful missing sample;
- conflicting evidence → use a discriminating assessment;
- stale evidence → use a representative refresh;
- observed-below-requirement → collect learning evidence only after useful intervention or when diagnosis needs refinement;
- supported claim → stop collecting when additional evidence has poor value relative to learner burden, unless the claim requires ongoing refresh.

Evidence collection is not free. Learner burden and operational cost matter after semantic sufficiency is preserved.

# AI assessment boundary

AI may be a primary automated assessor for repeated feedback and sampling, but it is not assessment truth.

Objective items use deterministic keys where possible. Productive assessment must be benchmarked, versioned, confidence-aware, and fail closed for claims above the demonstrated evaluator quality. If no eligible evaluator route exists, the correct result is delayed/insufficient/unavailable evidence—not silent substitution with a lower-quality route.

# Output consumed by Progression

A claim evaluation supplied to `09-PROGRESSION.md` exposes:

```text
claim_scope
readiness_status
supporting_evidence_fact_refs
blocking_conditions
material_conflicts
recency_state
confidence/calibration refs when relevant
policy/evaluator version
```

Progression consumes this result. It does not rescore observations or invent a second evidence policy.
