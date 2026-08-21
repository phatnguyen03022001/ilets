STATUS: CANONICAL
OWNS: learner runtime-state semantics, MasteryEstimate semantics, GapEvaluation and ActionIntent semantics, per-skill band advancement/regression, prerequisite gating behavior, adaptive scheduling, review policy, certification state, and exam-preparation mode
DEPENDS_ON: 01-LEARNER-MODEL.md, 05-BANDS.md, 06-CURRICULUM.md, 08-ASSESSMENT.md
DOES_NOT_OWN: Skill/Knowledge definitions, band thresholds, curriculum object ordering, learning mechanisms/practice taxonomy, assessment eligibility/sufficiency/scoring, or concrete data-storage implementation

# 09 — Progression

## Purpose

Define when learner state changes and which semantic next-action objective follows after Assessment has interpreted the available evidence.

Progression emits learner-state decisions, `GapEvaluation`, and `ActionIntent`. It does not choose the downstream Learning Mechanism or Practice Type; that selection is owned by `07-PRACTICE.md`.

Progression is a semantic state model, not a database schema.

## Core rule: progression is per skill

Listening, Reading, Writing, and Speaking progress independently.

A learner may legitimately hold an uneven profile. The IELTS overall band may be calculated for planning/information using `02-IELTS-MODEL.md`, but it is never a hard learning-progression gate.

## State layers

The runtime model keeps these layers distinct:

```text
Attempt history
    ↓
Observation / EvidenceFact history       owned semantically by 08
    ↓
MasteryEstimate                          owned here
    ↓
ReadinessEvaluation                      produced by 08
    ↓
Certification / GapEvaluation / ActionIntent
```

Historical evidence is not overwritten merely because the current interpretation changes.

# `MasteryEstimate`

A MasteryEstimate is a current, derived, uncertainty-aware interpretation of capability for a scoped Skill Leaf or Knowledge Object.

Conceptual fields:

```text
scope_ref
current_support_state
uncertainty
supporting_evidence_fact_refs
blocking_or_conflicting_refs
computed_at
policy_version
```

Recommended support states:

```text
unknown
learning
currently_supported
```

`unknown` does not mean weak. `currently_supported` means current admissible evidence supports the scoped mastery claim; it does not mean immutable mastery forever.

# `BandCertificationState`

Per `(skill, band)`:

```text
skill
band
status
claim_evaluation_ref
certification_evidence_refs
certified_at
```

Canonical status:

```text
not_started
in_progress
certified
```

Certification is internal learning-system recognition that the skill-band claim is currently `SUPPORTED` under `08-ASSESSMENT.md`. It is not an official IELTS result or guaranteed future exam score.

Current certification and historical attainment are separate.

# Band advancement

A skill-band becomes `certified` when:

1. `05-BANDS.md` defines the target threshold;
2. `08-ASSESSMENT.md` evaluates the corresponding claim as `SUPPORTED`;
3. no required claim condition remains blocked;
4. Progression records the current certification and its evidence/provenance.

Curriculum completion count is not a certification requirement. A learner who directly demonstrates the target capability may skip unnecessary acquisition stages. Required prerequisites constrain dependent learning paths; they are not paperwork that must be completed after the capability is already demonstrated.

After Band N is certified, the system may prioritize Band N+1 work for that skill without waiting for other skills.

# GapEvaluation

A GapEvaluation classifies what the current learner state actually requires. It is not synonymous with “the learner got something wrong.”

Canonical classes:

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

Semantics:

- `ABILITY_GAP` — current admissible evidence shows below-requirement target performance.
- `PREREQUISITE_GAP` — a Required prerequisite is materially missing for the dependent learning target.
- `EVIDENCE_GAP` — capability is unresolved because evidence is insufficient, not because weakness is established.
- `CONFLICTING_EVIDENCE` — material valid evidence supports incompatible interpretations.
- `STALE_EVIDENCE` — historical support exists but needs refresh for a current claim.
- `SCAFFOLD_DEPENDENCE` — performance is carried by support that the target claim requires the learner to perform without.
- `TRANSFER_GAP` — practiced performance does not generalize to a materially different required context.
- `FLUENCY_GAP` — target quality is broadly present but speed, automaticity, rhythm, or processing efficiency limits performance.
- `EXAM_CONDITION_GAP` — timing, integration, stamina, or exam-format conditions reduce performance without necessarily implying a missing underlying skill.

# ActionIntent

GapEvaluation maps to an explicit planning intent before downstream practice/assessment selection.

Canonical intents:

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

Representative mapping:

| Gap / state | Default ActionIntent |
|---|---|
| `PREREQUISITE_GAP` | `ACQUIRE_PREREQUISITE` |
| `ABILITY_GAP` | `REMEDIATE` or `CONSOLIDATE` based on diagnosis |
| `EVIDENCE_GAP` | `COLLECT_EVIDENCE` |
| `CONFLICTING_EVIDENCE` | `RESOLVE_CONFLICT` |
| `STALE_EVIDENCE` | `REASSESS` |
| `SCAFFOLD_DEPENDENCE` | `FADE_SCAFFOLD` |
| `TRANSFER_GAP` | `EXPAND_CONTEXT` |
| `FLUENCY_GAP` | `BUILD_FLUENCY` |
| `EXAM_CONDITION_GAP` | `EXAM_PREPARE` |
| target currently supported with remaining consolidation need | `CONSOLIDATE` |
| target currently supported and next target eligible | `ADVANCE` |

This mapping is a semantic default, not a hardcoded recommender algorithm.

# Required prerequisites

Dependency classification is owned by `06-CURRICULUM.md`.

Runtime behavior:

- **Required** prerequisite → hard gate for dependent learning when its absence would make the learning ineffective;
- **Recommended** prerequisite → ordering signal, not a blocker;
- **Independent** → no gate.

Evidence can accelerate the path. If assessment already demonstrates the dependent capability or prerequisite adequately, the learner need not complete redundant instructional stages.

# Next-action policy

A next action must be explainable as:

```text
learner target
+ current evidence state
+ GapEvaluation
+ prerequisite status
→ ActionIntent
```

`07-PRACTICE.md` then maps a learning-oriented ActionIntent to suitable Learning Mechanism(s) and Practice Type(s). Evidence-oriented intents may instead trigger an Assessment action.

The system must not route every uncertainty state into remediation.

# Learner agency and plan stability

The system should make a strong recommendation while preserving practical learner control where the action set remains valid.

A learner may be allowed to swap, shorten, skip, or change skill among eligible alternatives. Such behavior changes planning context; it does not create evidence of ability or inability.

Repeated skipping or abandonment is a preference/friction signal. It may justify a different eligible activity or delivery pattern but cannot lower the canonical standard.

A current plan should remain stable enough to trust. Replanning should follow material learner/evidence state change rather than every minor scoring fluctuation.

# Review scheduling

Spacing applies only where a target is meaningfully reviewable through repeated retrieval or repeated performance.

A review system may expand or shorten intervals using performance history, but there is no universal spacing formula across vocabulary, grammar, writing organization, speaking fluency, and integrated IELTS tasks.

Suitable reviewable targets may use spaced retrieval; other targets may use targeted reassessment, transfer work, or performance practice. Exact mechanism/type selection is downstream Practice authority.

# Exam Preparation mode

Exam Preparation may expose higher-demand or integrated tasks before certification for diagnosis, pacing, familiarity, stamina, strategy, or readiness estimation.

It may request downstream timed/integrated Practice or a full readiness Assessment.

Exam Preparation must never:

- unlock a higher band by exposure alone;
- satisfy a missing Required prerequisite by completion alone;
- treat one mock as certification;
- rewrite a Band threshold;
- convert an unresolved evidence state into certainty.

# Staleness, conflict, and regression

These are different conditions.

### Staleness

Stale evidence means the current claim needs refresh. It does **not** mean the learner regressed.

### Conflict

Conflicting evidence means the system does not yet know which interpretation is reliable. It triggers a discriminating evidence intent rather than majority-vote averaging.

### Regression

Regression requires later admissible evidence showing that a previously supported capability is now below the required current threshold.

When regression is established:

1. preserve historical evidence and prior attainment;
2. update the current MasteryEstimate honestly;
3. move affected current certification from `certified` to `in_progress` when the skill-band claim is no longer supported;
4. classify the relevant GapEvaluation;
5. emit the appropriate ActionIntent;
6. re-certify only through the normal Assessment policy.

Absence of recent evidence alone never revokes historical attainment.

# Certification history

Certification history records point-in-time attainment with evidence and policy provenance. Later regression or re-certification adds history; it does not erase earlier events.

# Explainability invariant

Every state transition and ActionIntent must be reconstructable from canonical target references, evidence interpretation, and prerequisite status.

No transition may depend on duplicate learning definitions or opaque “AI decided” state change.
