STATUS: CANONICAL
OWNS: learner runtime-state semantics, mastery-state transitions, per-skill band advancement and regression, prerequisite gating behavior, adaptive scheduling semantics, review scheduling, certification status, and exam-preparation mode
DEPENDS_ON: 01-LEARNER-MODEL.md, 05-BANDS.md, 06-CURRICULUM.md, 07-PRACTICE.md, 08-ASSESSMENT.md
DOES_NOT_OWN: Skill/Knowledge definitions, band thresholds, curriculum object ordering, practice taxonomy, assessment sufficiency or scoring, concrete data-storage implementation

# 09 — Progression

## Purpose

Define **when learner state changes** after valid learning and assessment events.

Progression is a semantic state model, not a database design and not a storage service.

It references canonical objects by ID and never redefines them.

## Core rule: progression is per skill

Band progression is independent for:

```text
Listening
Reading
Writing
Speaking
```

A learner may legitimately hold an uneven profile such as:

```text
Listening 6
Reading   7
Writing   5
Speaking  6
```

A stronger skill must not be blocked because another skill is weaker.

The IELTS overall band may be calculated for informational/planning purposes from the four section bands, using the external rule in `02-IELTS-MODEL.md`, but the overall band is **not a learning-progression gate**.

## Runtime state model

### `LeafMasteryState`

Per Skill Leaf:

```text
leaf_id
mastery_state
confidence
assessment_evidence_refs
last_assessed
```

Canonical mastery-state progression:

```text
not_started
    ↓
practicing
    ↓
emerging
    ↓
mastered
```

A mastered leaf may regress to `emerging` when later valid evidence shows that the capability is no longer held reliably.

The thresholds that make evidence sufficient are owned by `08-ASSESSMENT.md`.

### `KnowledgeState`

Per Knowledge Object:

```text
knowledge_id
state
confidence
evidence_refs
```

Canonical states:

```text
not_acquired
    ↓
learning
    ↓
acquired
```

Later evidence may return an inadequately retained object to `learning`.

### `BandCertificationState`

Per `(skill, band)`:

```text
skill
band
status
certification_evidence_refs
```

Canonical status:

```text
not_started
    ↓
in_progress
    ↓
certified
```

Certification is current-state truth, not an irrevocable badge.

If later valid evidence shows regression below the certified band's floor, status returns:

```text
certified → in_progress
```

Re-certification then requires a fresh valid certification evidence set under `08-ASSESSMENT.md`.

There is no separate `revoked` status. Historical attainment is preserved separately.

### `OverallLearnerState`

Conceptually includes:

- current certified band per skill;
- Skill Leaf mastery-state map;
- Knowledge Object state map;
- band certification states;
- certification history;
- exam-preparation mode/state;
- due-review queue;
- current curriculum recommendations.

This is a semantic model. It does not prescribe SQL tables, event sourcing, document storage, or any implementation technology.

## State transitions

### `not_started → practicing`

Occurs when the learner meaningfully attempts practice targeting the canonical Skill Leaf.

Merely viewing content does not necessarily count as practice.

### `practicing → emerging`

Occurs when valid assessment evidence begins to show the target capability but the certification/mastery evidence standard has not yet been satisfied.

### `emerging → mastered`

Occurs when `08-ASSESSMENT.md` reports that the Skill Leaf evidence satisfies the canonical sufficiency and confidence policy for the relevant target.

Progression consumes the assessment result; it does not reimplement the scoring rule.

### `mastered → emerging`

Occurs after later valid evidence demonstrates meaningful regression.

Regression is a learning-state change, not deletion of history.

## Band advancement

A skill advances from Band N toward Band N+1 when:

1. the current Band-N exit conditions owned by `05-BANDS.md` have valid certification evidence;
2. `08-ASSESSMENT.md` reports the evidence set as sufficient and confidence-valid;
3. the per-skill `BandCertificationState` is set to `certified` for Band N;
4. curriculum recommendations may then prioritize appropriate Band-(N+1) work for that skill.

This does not require other skills to certify Band N first.

## Certification history

Current certification and historical attainment are separate.

`certification_history` records point-in-time attainment, including:

- skill;
- band;
- evidence reference;
- certification time;
- later regression/re-certification events where applicable.

If current performance regresses, prior attainment remains in history while current certification status reverts to `in_progress`.

This avoids the false choice between erasing prior achievement and pretending current capability has not changed.

## Prerequisite gating

Curriculum dependency classification is owned by `06-CURRICULUM.md`.

Runtime behavior:

- **Required** prerequisite → hard gate until the prerequisite's required learner state is satisfied;
- **Recommended** prerequisite → influences ordering/recommendation but does not block;
- **Independent** → no gate.

Hard gates must remain minimal. Adaptation may route around Recommended dependencies but may not bypass Required ones.

## Adaptive within-band planning

Within a band, the system may personalize:

- node order where prerequisites permit;
- practice type selection;
- difficulty;
- review frequency;
- remediation;
- skill emphasis;
- pace;
- assessment timing.

Adaptation must never:

- change the canonical target outcome;
- remove a required Skill Leaf;
- redefine a Band threshold;
- treat practice completion as mastery;
- bypass Required prerequisites;
- convert exam-preparation exposure into certification.

Same outcomes, different paths.

## Next-action policy

A learner's next recommendation should be explainable using canonical references.

Representative actions:

### Advance

Select a Curriculum Node whose Required prerequisites are satisfied and whose targets move the learner toward the current per-skill band goal.

### Remediate

Target a weak or regressed Skill Leaf using an appropriate Practice Type and the enabling Knowledge Objects that explain the gap.

### Review

Schedule retrieval for due Skill/Knowledge targets where retention is at risk.

### Assess

Collect evidence when a target is near mastery or existing evidence has become stale/insufficient.

### Exam prepare

Expose the learner to timed, integrated, or higher-demand exam tasks without changing certification status unless the normal evidence standard is independently met.

## Review scheduling

Long-term retention uses performance-informed spacing rather than a fixed universal calendar.

A review system should:

- schedule retrieval before knowledge/capability is fully lost;
- expand intervals after successful retrieval;
- shorten intervals after weak retrieval;
- prioritize repeatedly regressing targets;
- keep review distinct from new acquisition.

`PT-17` is the canonical spaced-retrieval Practice Type; other practice types may also be scheduled for review when appropriate.

Exact scheduling algorithms and interval formulas are implementation/calibration concerns unless future evidence shows they must become canonical learning policy.

## Exam Preparation mode

Learning Progression and Exam Preparation are independent concepts.

Exam Preparation may:

- expose higher-band tasks before certification;
- schedule `PT-06`, `PT-11`, `PT-15`, `PT-23`;
- use `AT-07` full mocks;
- prioritize timing, task familiarity, stamina, and strategy;
- diagnose likely test-day performance.

Exam Preparation must never:

- unlock a higher band by exposure alone;
- satisfy a missing Required prerequisite;
- treat a single mock as certification;
- modify the canonical Band threshold;
- conceal uncertainty in readiness estimates.

## Regression policy

Regression can occur at Skill Leaf, Knowledge Object, or current skill-band certification level when later valid evidence shows loss of capability.

Response to regression:

1. preserve historical evidence;
2. update current state honestly;
3. identify the canonical objects involved;
4. schedule targeted review/remediation;
5. collect fresh evidence;
6. re-certify only through the normal Assessment policy.

## Explainability invariant

Every progression decision must be explainable through:

```text
learner state
+ canonical target IDs
+ prerequisite status
+ assessment evidence
+ band requirement
→ state transition / next recommendation
```

No transition may depend on an untraceable duplicate definition of the learning target.
