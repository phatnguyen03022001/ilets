STATUS: CANONICAL
OWNS: conceptual learner representation, learner dimensions, starting and target state semantics, learner constraints that affect planning
DEPENDS_ON: 00-PRODUCT.md, 02-IELTS-MODEL.md
DOES_NOT_OWN: IELTS band definitions, canonical skill or knowledge objects, curriculum ordering, assessment sufficiency, runtime state transitions, certification rules

# 01 — Learner Model

## Purpose

Define the learner that the learning system reasons about without duplicating the domain objects used to teach, assess, or progress that learner.

This is a conceptual model. Per-learner runtime state and transitions are owned by `09-PROGRESSION.md`.

## Learner profile

A learner is represented by a profile with the following independent dimensions.

### Current capability

The learner may have a different current level in each IELTS skill:

```text
Listening
Reading
Writing
Speaking
```

Uneven profiles are normal. The model must not collapse a learner into a single overall band for progression purposes.

A learner can therefore be, for example, stronger in Reading than Writing without the stronger skill being artificially gated by the weaker one.

### Target outcome

The learner may have:

- a target band per skill;
- a target overall IELTS band;
- a test date or no fixed test date;
- a target test variant, initially Academic.

A target overall band is a planning objective. It does not replace the per-skill competence requirements owned by `05-BANDS.md`.

### Skill mastery profile

The learner's actual capability is represented at runtime by references to canonical Skill Leaves from `03-SKILLS.md`.

The learner model does not redefine those leaves. It only requires that the system can describe which capabilities are strong, weak, unstarted, developing, or mastered.

### Knowledge profile

The learner's enabling knowledge is represented at runtime by references to canonical Knowledge Objects from `04-KNOWLEDGE.md`.

Knowledge gaps may explain skill difficulties, but knowledge acquisition is not itself an IELTS section score.

### Evidence profile

The learner accumulates assessment evidence over time. Evidence belongs to learner runtime state and references canonical Assessment Types.

A learner profile must allow the system to distinguish:

- a single good performance;
- repeated reliable performance;
- low-confidence performance;
- current regression after prior attainment.

The sufficiency rules are owned by `08-ASSESSMENT.md`; the state consequences are owned by `09-PROGRESSION.md`.

### Learning constraints

Planning may account for constraints such as:

- available study time;
- fixed exam date;
- preferred or available delivery mode;
- accessibility needs;
- prior exposure to IELTS;
- first-language context for optional explanations or remediation;
- temporary focus on a specific weak skill.

Constraints may change sequencing or emphasis but must not change the canonical mastery requirement for a target band.

## Program range

The structured learning program targets Bands 3–9.

Bands 0–2 remain useful as diagnostic boundaries in the IELTS model, but the Blueprint does not define a detailed Bands 0–2 curriculum.

A learner below the Band-3 structured entry range should receive diagnosis and foundational remediation before entering the canonical Band-3 pathway. The specific band boundary is owned by `05-BANDS.md`.

## Starting state

A learner does not have to begin with every skill at the same band.

Initial diagnosis should establish, as evidence permits:

- current per-skill band estimate;
- relevant Skill Leaf strengths and gaps;
- relevant Knowledge Object gaps;
- prerequisite gaps;
- exam-readiness constraints if a fixed test date exists.

Diagnosis informs planning. It does not itself certify mastery unless the assessment evidence also satisfies the certification requirements.

## Target state

A learner reaches a target state when the required per-skill capability has been demonstrated according to the owning band and assessment rules.

The target state is therefore not defined by:

- number of lessons completed;
- time spent;
- practice count alone;
- a model's unsupported prediction;
- the learner's overall-band average alone.

## Learning progression vs exam preparation

The learner can be in ordinary learning progression, exam-preparation mode, or both simultaneously for different purposes.

Exam-preparation mode may expose the learner to higher-band or full-test tasks before mastery for:

- diagnosis;
- familiarization;
- pacing;
- test strategy;
- readiness estimation.

This exposure must not alter the learner's certified learning state. The transition rules are owned by `09-PROGRESSION.md`.

## L1-agnostic canonical learner

The canonical learner model is independent of first language.

First-language information may improve explanation, transfer-error detection, pronunciation remediation, or localization. It remains contextual input and cannot create a different canonical definition of IELTS competence.

## Same outcomes, different paths

Learners may receive different:

- practice selections;
- review schedules;
- remediation;
- difficulty progression;
- within-band sequencing;
- pacing.

They must still reach the same canonical target outcomes for the same target skill and band.

Adaptation personalizes the path, not the standard.

## Model boundary

This file defines who the learner is conceptually. It intentionally does not define database fields, event history, transition algorithms, or storage. Those are runtime/implementation concerns unless their semantics affect learning progression, in which case `09-PROGRESSION.md` owns them.
