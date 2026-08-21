STATUS: CANONICAL
OWNS: end-to-end learner product journey, navigation surfaces, study-session shapes, product timing defaults, and user-visible interpretation of learning state
DEPENDS_ON: ../spec/01-LEARNER-MODEL.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md
DOES_NOT_OWN: skill definitions, mastery thresholds, evidence sufficiency, concrete practice taxonomy, API wire contracts, framework selection, persistence, or deployment

# Learning Experience

## Purpose

Define how a learner experiences the learning system from first launch through daily study, assessment, review, and exam preparation.

This document translates canonical learning semantics into a coherent product journey. It does not change the standard for mastery or progression.

## Primary product loop

```text
Goal
  ↓
Diagnostic
  ↓
Learner model
  ↓
Daily plan
  ↓
Learning session
  ↓
Attempt + feedback
  ↓
Evidence / mastery update
  ↓
Gap evaluation
  ↓
Next action
  ↺
```

The learner should always be able to answer four questions:

1. What am I trying to improve?
2. Why is this activity being recommended?
3. What did the last attempt show?
4. What should I do next?

## First-run flow

### 1. Goal setup

Capture only information that changes learning planning:

- IELTS variant, initially Academic;
- target test date, optional;
- target overall band, optional;
- target band per skill, optional;
- current self-estimate, explicitly non-authoritative;
- available study time;
- preferred focus or urgent skill;
- accessibility and optional first-language context.

Target setup should take approximately **3–5 minutes**.

### 2. Diagnostic choice

The initial product exposes two entry modes.

| Mode | Target duration | Purpose | Authority |
|---|---:|---|---|
| Quick start | 15 min | establish a rough first plan and identify obvious gaps/unknowns | diagnostic only, non-certifying |
| Full baseline | 60 min target window | collect broader four-skill evidence before long-term planning | diagnostic only unless normal Assessment rules independently support a claim |

These are product scheduling defaults, not mastery requirements. The diagnostic may shorten or extend when the evidence state makes that appropriate.

### 3. First plan

After diagnosis, the app must distinguish:

- observed gap;
- missing evidence;
- conflicting evidence;
- stale evidence;
- prerequisite gap;
- current strength.

Unknown must never be rendered as failure.

## Home surface

The default learner home has seven primary destinations:

1. **Today** — current plan and resume state;
2. **Skills** — Listening, Reading, Writing, Speaking profiles;
3. **Practice** — all user-facing practice modes;
4. **Review** — due knowledge, errors, and re-evidence actions;
5. **Media Lab** — authorized/eligible video and audio learning;
6. **Progress** — current estimates, evidence state, certification history, and target distance;
7. **Mock** — section or full-test readiness practice.

Streaks, points, badges, or social mechanics may be added as engagement features, but they never alter mastery, evidence, or certification.

## Daily study presets

The app exposes exactly three default study-plan sizes. A learner may stop early or continue, but plan generation uses these stable product presets.

### Quick — 10 minutes

```text
2 min  retrieval / due review
6 min  one highest-value focus activity
2 min  result + next-action summary
```

Default maximum: **3 activity blocks**.

### Standard — 25 minutes

```text
4 min  retrieval / review
8 min  acquire or remediate
9 min  independent practice / transfer
4 min  checkpoint + next plan
```

Default maximum: **4 activity blocks**.

### Deep — 45 minutes

```text
5 min   retrieval / review
10 min  acquire or remediate
12 min  focused practice
12 min  transfer / fluency / timed work
6 min   checkpoint / reflection / next plan
```

Default maximum: **5 activity blocks**.

These are scheduling defaults. They do not imply that a Skill, Knowledge Object, or Band requires a fixed number of minutes.

## Learning-session state machine

```text
planned
  ↓
in_progress
  ↓
completed | abandoned
```

An abandoned session is an adherence signal, not negative ability evidence.

A session may contain:

- due review;
- prerequisite acquisition;
- targeted remediation;
- focused practice;
- transfer;
- fluency work;
- assessment sampling;
- exam-readiness work.

The mix is driven by `GapEvaluation → ActionIntent` from Progression and the mechanism/practice rules in Practice.

## Activity card contract

Before a learner begins an activity, the card should expose:

- skill and canonical target;
- plain-language objective;
- why it is recommended;
- expected duration;
- mode: learn, practice, review, assess, or mock;
- whether the activity is scaffolded or independent;
- whether the result may contribute evidence.

After completion, the result should expose:

- what was observed;
- what changed, if anything;
- uncertainty or missing evidence;
- recommended next action;
- retry/review options.

Do not display a fabricated precise band change after every micro-activity.

## Skill-page contract

Each skill page contains:

```text
Current estimate
Target
Evidence state
Canonical capability map
Current gaps
Recommended actions
Practice shortcuts
Recent attempts
Readiness / certification state
```

Skill pages may look different because Listening/Reading are objective-receptive while Writing/Speaking are productive and criterion-based.

## Review experience

The Review surface combines three distinct queues without pretending they are one memory model:

1. Knowledge retrieval — vocabulary/grammar/phonology objects appropriate for spaced review;
2. Error remediation — recurring error/remediation patterns tied to recent attempts;
3. Re-evidence — stale, conflicting, or insufficient claims requiring a fresh independent sample.

The UI may present one queue, but the scheduler must preserve those semantics.

## Exam-preparation experience

Exam Preparation is an explicit mode that may increase:

- timed sections;
- full-task Writing;
- complete Speaking simulations;
- integrated receptive sections;
- full mocks;
- pacing/stamina work.

It does not change canonical mastery thresholds and does not convert one mock into certification.

## Progress presentation

User-facing progress should prefer explainable states over false precision.

Recommended labels include:

- **Not sampled yet**;
- **Needs more evidence**;
- **Developing**;
- **Ready to reassess**;
- **Currently supported at Band N**;
- **Evidence stale**;
- **Evidence conflicting**.

Any numeric confidence shown to users must be meaningful and calibrated; otherwise use categorical language.

## Product defaults vs learning truth

Durations, block counts, navigation order, and dashboard layout in this file are product defaults. They may be tuned by product evidence without changing IELTS or mastery semantics.

If a product experiment would change what counts as mastery, progression, prerequisite, or valid evidence, it must be resolved in the owning `spec/` document instead.