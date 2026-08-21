STATUS: CANONICAL
OWNS: learner-state semantics, MasteryEstimate, GapEvaluation, ActionIntent, per-skill Band advancement/regression, prerequisite-gating semantics, semantic review state, certification history/current state, and Exam Preparation mode semantics
DEPENDS_ON: 01-LEARNER-MODEL.md, 05-BANDS.md, 06-CURRICULUM.md, 08-ASSESSMENT.md
DOES_NOT_OWN: Skill/Knowledge definitions, Band thresholds, curriculum ordering, Practice selection, Assessment eligibility/scoring, product planner candidate eligibility/ranking, UX scheduling, or storage schema

# 09 — Progression

## Purpose

Define when learner interpretation changes and which **semantic next-action objective** follows after Assessment has interpreted current evidence.

Progression emits learner state, `GapEvaluation`, and `ActionIntent`. It does not choose a UI feature, Practice Mode, concrete item, or ranked DailyPlan candidate.

# Per-skill progression

Listening, Reading, Writing, and Speaking progress independently.

Overall IELTS Band may be relevant to TargetProfile planning, but it is not a synchronized four-skill advancement gate.

# State layers

```text
Attempt history
  ↓
Observation / EvidenceFact         Assessment
  ↓
MasteryEstimate                    Progression
  ↓
ReadinessEvaluation                Assessment
  ↓
Certification / GapEvaluation / ActionIntent   Progression
  ↓
Product Planner                    design/04
```

No layer rewrites the historical evidence owned upstream.

# `MasteryEstimate`

A MasteryEstimate is the current uncertainty-aware interpretation of one scoped Skill Leaf or Knowledge Object.

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

Support states:

```text
unknown
learning
currently_supported
```

`unknown` is not weak. `currently_supported` is a current evidence interpretation, not permanent mastery.

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

Status:

```text
not_started
in_progress
certified
```

Certification means the internal skill-Band claim is currently `SUPPORTED` under Assessment policy. It is not an official IELTS result or guarantee.

# Band advancement

A skill-Band becomes `certified` when:

1. `05-BANDS.md` defines the threshold;
2. `08-ASSESSMENT.md` returns `SUPPORTED` for the corresponding claim;
3. no required claim condition remains blocked;
4. Progression records current state + evidence/policy provenance.

Curriculum completion count is not a certification requirement. Valid evidence may accelerate past redundant acquisition stages.

After Band N certification, Band N+1 work for that skill may become semantically eligible without waiting for other skills.

# `GapEvaluation`

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

Meanings:

- `ABILITY_GAP` — admissible current evidence is below the target capability requirement;
- `PREREQUISITE_GAP` — a Required prerequisite is materially unresolved for dependent learning;
- `EVIDENCE_GAP` — evidence is insufficient; weakness is not established;
- `CONFLICTING_EVIDENCE` — material admissible evidence supports incompatible interpretations;
- `STALE_EVIDENCE` — historical support needs refresh for a current claim;
- `SCAFFOLD_DEPENDENCE` — material support is carrying performance required independently;
- `TRANSFER_GAP` — capability does not generalize to a materially required context;
- `FLUENCY_GAP` — underlying quality is broadly present but automaticity/speed/rhythm/processing efficiency limits performance;
- `EXAM_CONDITION_GAP` — timing, integration, stamina, test-interface/input mode, or other exam-condition demand reduces performance without necessarily implying missing underlying capability.

Delivery-mode unfamiliarity belongs under `EXAM_CONDITION_GAP` when it is the material cause. Do not create a second ability taxonomy for computer/handwriting/remote delivery.

# `ActionIntent`

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

Default semantic mapping:

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

This table defines semantic intent, not a recommender score/ranking algorithm.

# Required prerequisites

Dependency classification is owned by `06-CURRICULUM.md`.

Progression semantics:

- **Required** → unresolved prerequisite produces a dependent-learning gate;
- **Recommended** → ordering signal, not a canonical hard block;
- **Independent** → no prerequisite relation.

Evidence may satisfy/resolve a prerequisite without requiring the learner to replay redundant lessons.

The product Planner consumes these gates; it does not redefine them.

# Next-action explanation

A semantic next action is reconstructable as:

```text
TargetProfile condition
+ Assessment evidence state
+ prerequisite state
→ GapEvaluation
→ ActionIntent
```

`07-PRACTICE.md` maps learning-oriented intents to Learning Mechanisms/Practice Types. Assessment-oriented intents may trigger additional measurement.

`design/04-application-flows.md` then performs product hard eligibility, candidate generation, and ranking.

Uncertainty must not be routed automatically to remediation.

# Learner behavior boundary

Skipping, preference, abandonment, or repeated friction may influence product planning but are not ability evidence.

Progression does not lower target standards because the learner avoids a task.

# Review state

Spacing/review applies only where repeated retrieval or repeated performance is meaningful.

No universal review formula applies to vocabulary, grammar, writing organization, speaking fluency, and integrated IELTS tasks.

Progression may determine that a target is due for review/reassessment. Concrete scheduling time, mode choice, and activity ranking are downstream product/Practice concerns.

# Exam Preparation mode

Exam Preparation may expose higher-demand/integrated conditions before certification for:

- diagnosis;
- timing/pacing;
- interface/input familiarity;
- strategy;
- stamina;
- target-condition readiness.

It may request timed/integrated Practice or readiness Assessment.

It never:

- unlocks a higher Band by exposure;
- satisfies a Required prerequisite by completion alone;
- turns a mock into automatic certification;
- rewrites Band thresholds;
- converts unresolved evidence into certainty.

# Staleness, conflict, regression

## Staleness

Evidence is too old for the current claim. Refresh is needed; regression is not established.

## Conflict

Material evidence supports incompatible interpretations. Resolve with discriminating evidence rather than averaging the conflict away.

## Regression

Regression requires later admissible evidence establishing that previously supported current capability is now below requirement.

When established:

1. preserve historical evidence/attainment;
2. update current MasteryEstimate;
3. move affected current certification from `certified` to `in_progress` when the claim is no longer supported;
4. classify GapEvaluation;
5. emit ActionIntent;
6. re-certify only through normal Assessment policy.

Absence of recent evidence alone never establishes regression.

# Certification history

Historical certifications remain point-in-time records with evidence/policy provenance. Regression/re-certification append history rather than rewriting prior attainment.

# Explainability invariant

Every learner-state transition and ActionIntent must be reconstructable from canonical target references, Assessment interpretation, and prerequisite state.

No transition may depend on a duplicated learning definition, ranker side effect, opaque model decision, UI completion state, or provider failure.