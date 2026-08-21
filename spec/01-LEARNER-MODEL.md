STATUS: CANONICAL
OWNS: conceptual learner representation, learner dimensions, starting and target state semantics, learner constraints and epistemic conditions that affect planning
DEPENDS_ON: 00-PRODUCT.md, 02-IELTS-MODEL.md
DOES_NOT_OWN: IELTS band definitions, canonical skill or knowledge objects, curriculum ordering, assessment sufficiency, runtime state transitions, certification rules

# 01 — Learner Model

## Purpose

Define the learner that the learning system reasons about without duplicating the domain objects used to teach, assess, or progress that learner.

The learner is not represented by one scalar band. The system reasons about an uneven, evidence-limited, changing profile. Per-learner runtime transitions are owned by `09-PROGRESSION.md`; evidence semantics are owned by `08-ASSESSMENT.md`.

## Learner profile

A learner profile contains independent dimensions.

### Current capability

Current capability is represented per IELTS skill:

```text
Listening
Reading
Writing
Speaking
```

Uneven profiles are normal. A learner may be stronger in Reading than Writing without either skill being forced to the other skill's state.

A current capability estimate is a claim supported by evidence, not a permanent fact. Unknown, stale, or conflicting evidence must not be silently converted into weakness.

### Target profile

A learner may have:

- target band per skill;
- target overall IELTS band;
- fixed test date or no fixed date;
- target test variant, initially Academic;
- temporary study priorities.

The target overall band is a planning constraint. It does not replace the per-skill requirements owned by `05-BANDS.md`.

### Skill profile

Skill capability references canonical Skill Leaves from `03-SKILLS.md`.

The learner model requires the system to represent, as evidence permits, which leaves are unobserved, developing, currently supported, weak under a defined context, or in need of fresh evidence. This file does not define the measurement rule behind those states.

### Knowledge profile

Enabling knowledge references canonical Knowledge Objects from `04-KNOWLEDGE.md`.

A Knowledge gap may explain a skill difficulty, but Knowledge is not an IELTS section score and passing a Knowledge probe cannot substitute for target-skill performance.

### Epistemic profile

The learner model must preserve uncertainty about what is known.

For any material claim the system must be able to distinguish at least:

- sufficient current evidence;
- insufficient evidence;
- conflicting evidence;
- stale evidence;
- observed below-requirement performance.

These conditions have different consequences. `08-ASSESSMENT.md` owns their evidence interpretation and `09-PROGRESSION.md` owns the resulting next action or state change.

### Learning constraints

Planning may account for:

- available study time;
- fixed exam date;
- preferred or available delivery mode;
- accessibility needs;
- prior IELTS exposure;
- first-language context for optional explanation/remediation;
- temporary focus on a weak skill;
- learner preference and repeated friction.

Constraints may alter path, pacing, or presentation. They must not alter the canonical standard for a target skill and band.

Preference, skipping, abandonment, or friction are not ability evidence by themselves.

## Program range

The structured learning program targets Bands 3–9.

Bands 0–2 remain diagnostic/external boundaries. A learner below the Band-3 structured entry range receives diagnosis and foundational remediation before entering the canonical Band-3 pathway. The band boundary is owned by `05-BANDS.md`.

## Starting state

A learner does not have to begin with every skill at the same band.

Initial diagnosis should establish only what the available evidence supports:

- current per-skill performance estimate;
- Skill Leaf strengths and gaps;
- Knowledge Object gaps;
- prerequisite gaps;
- unresolved uncertainty or conflicting evidence;
- exam-readiness constraints when a fixed date exists.

A short initial diagnostic should be allowed to produce a provisional profile rather than pretending to know the entire learner.

Diagnosis informs planning. It does not certify mastery unless the relevant evidence also satisfies `08-ASSESSMENT.md`.

## Target state

A learner reaches a target state when the required per-skill capability is currently supported by valid evidence under the owning Band and Assessment rules.

The target state is not defined by:

- number of lessons completed;
- time spent;
- practice count;
- one unusually good attempt;
- an unsupported model prediction;
- overall-band averaging alone.

## Learning progression vs exam preparation

Ordinary learning progression and exam preparation are independent dimensions.

Exam preparation may expose higher-demand or full-test tasks before mastery for diagnosis, familiarization, pacing, strategy, stamina, or readiness estimation. Exposure does not alter certified learning state unless the normal evidence rules are independently satisfied.

## L1-agnostic canonical learner

The canonical learner model is independent of first language.

First-language information may improve explanation, localization, transfer-error hypotheses, and pronunciation remediation. It cannot create a different definition of IELTS competence.

## Same outcomes, different paths

Learners may receive different practice selections, review schedules, scaffolds, remediation, difficulty progression, sequencing, and pacing.

They must still meet the same canonical target for the same skill and band.

Adaptation personalizes the path, not the standard.

## Model boundary

This file defines who the learner is conceptually. It does not define database fields, event history, scoring algorithms, or storage. Those are implementation concerns unless their semantics affect learning, assessment, or progression, in which case the relevant canonical owner defines the behavior.
