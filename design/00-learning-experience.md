STATUS: CANONICAL
OWNS: end-to-end learner product journey, TargetProfile semantics, navigation surfaces, study-session shapes, product timing defaults, route-to-target behavior, and user-visible interpretation of learning state
DEPENDS_ON: ../spec/01-LEARNER-MODEL.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md
DOES_NOT_OWN: skill definitions, mastery thresholds, evidence sufficiency, concrete practice taxonomy, coverage support declarations, API wire contracts, framework selection, persistence, or deployment

# Learning Experience

## Purpose

Define how a learner experiences the learning system from target setup through daily study, assessment, review, and exam preparation.

This document translates canonical learning semantics into a coherent product journey. It does not change the standard for mastery or progression and does not guarantee an external IELTS result.

# Primary product loop

```text
TargetProfile
  ↓
Diagnostic
  ↓
Learner model
  ↓
Target gap / unknown / due review
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

The learner should always be able to answer:

1. What target am I working toward?
2. What is currently blocking or uncertain?
3. Why is this activity recommended now?
4. What did the last attempt actually show?
5. What should I do next?

# TargetProfile

A learner goal is a constraint profile, not one overall-band number.

Conceptual fields:

```text
test_variant                 Academic | General Training
purpose_or_receiving_rule    optional external constraint reference
target_overall_band          optional
minimum_listening_band       optional
minimum_reading_band         optional
minimum_writing_band         optional
minimum_speaking_band        optional
test_date                    optional
selected_skill_retake        optional
```

Rules:

1. the learner may set an overall target, per-skill minimums, or both;
2. if a university/visa/employer requirement has per-skill minima, those constraints are represented explicitly;
3. an overall target alone does not uniquely determine four skill targets;
4. the product must not fabricate hidden per-skill requirements from one overall number;
5. the planner may propose a balanced planning profile, but it must be labelled as a product planning choice rather than external IELTS truth;
6. the target remains stable until the learner explicitly changes it;
7. current readiness is evaluated against TargetProfile conditions, not internal mastery averages.

Whether a TargetProfile is currently covered/supported is evaluated downstream by `08-coverage-and-support.md`; this document owns the learner-facing target and route semantics, not support declaration.

# First-run flow

## 1. Goal setup

Capture only information that changes planning:

- IELTS variant;
- target overall band and/or required skill minima;
- receiving requirement when relevant;
- target test date, optional;
- selected One Skill Retake focus, optional;
- current self-estimate, explicitly non-authoritative;
- available study time;
- preferred focus or urgent skill;
- accessibility and optional first-language context.

Target setup should normally take **3–5 minutes**.

## 2. Diagnostic choice

The initial product exposes two entry modes.

| Mode | Target duration | Purpose | Authority |
|---|---:|---|---|
| Quick start | 15 min | establish a provisional plan and locate obvious gaps/unknowns | diagnostic only, non-certifying |
| Full baseline | 60 min target window | collect broader four-skill evidence before long-term planning | diagnostic only unless normal Assessment rules independently support a claim |

A short diagnostic never pretends to know the whole learner. It may shorten or extend when the evidence state justifies that.

## 3. First plan

The plan distinguishes:

- observed ability gap;
- prerequisite gap;
- missing evidence;
- conflicting evidence;
- stale evidence;
- scaffold dependence;
- transfer gap;
- fluency gap;
- exam-condition gap;
- product CoverageGap.

A product CoverageGap is never rendered as learner weakness.

# Today-first home

The default landing surface is **Today**. It provides a strong current recommendation rather than asking the learner to browse a library before learning can begin.

Primary destinations:

1. **Today** — current plan and resume state;
2. **Skills** — Listening, Reading, Writing, Speaking profiles;
3. **Practice** — user-facing practice modes;
4. **Review** — due knowledge, errors, and re-evidence;
5. **Media Lab** — eligible video/audio learning;
6. **Progress** — target conditions, current evidence, certification history, blockers;
7. **Mock** — section or full-test readiness practice.

The learner can browse Practice directly, but direct browsing does not silently replace the recommended route or satisfy missing target conditions.

# Strong recommendation + learner agency

The system preserves both route integrity and learner control.

The Today plan is a strong ranked recommendation derived from:

```text
TargetProfile
+ current gaps / unknowns
+ due review
+ Required prerequisites
+ available time
+ evidence value
+ learner constraints
→ eligible next actions
```

The learner may:

- Swap to another eligible activity;
- Skip an activity;
- Shorten the session at a safe stopping point;
- Change skill focus;
- edit the TargetProfile.

But:

- Required prerequisites remain blocking;
- skipped requirements remain unresolved;
- a preferred activity cannot certify an unrelated target;
- repeated skip/abandonment is a friction/preference signal, not ability evidence;
- the planner should seek eligible alternatives before repeating a low-value failed intervention;
- the system must not weaken the target merely because the learner struggles.

# Daily study presets

The app exposes three default plan sizes while allowing coherent adaptation.

## Quick — 10 minutes

Typical envelope:

```text
2 min  due review / retrieval
6 min  one highest-value action
2 min  result + next-action summary
```

Default maximum: **3 activity blocks**.

## Standard — 25 minutes

Typical envelope:

```text
4 min  review / retrieval
8 min  acquire or remediate
9 min  independent practice / transfer
4 min  checkpoint + next plan
```

Default maximum: **4 activity blocks**.

## Deep — 45 minutes

Typical envelope:

```text
5 min   review
10 min  acquire or remediate
12 min  focused practice
12 min  transfer / fluency / timed work
6 min   checkpoint / next plan
```

Default maximum: **5 activity blocks**.

These are scheduling defaults, not a fixed pedagogical recipe. Session composition follows the current ActionIntent, prerequisites, evidence value, and safe stopping points.

# Learning-session state

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
- scaffold fading;
- transfer;
- fluency work;
- assessment sampling;
- exam-readiness work.

# Activity card contract

Before an activity, show:

- target skill/capability;
- plain-language objective;
- why now;
- expected duration;
- mode: learn, practice, review, assess, or mock;
- scaffold/independence state;
- whether the result may be evidence-eligible.

After an activity, show:

- what was observed;
- what changed, if anything;
- uncertainty/missing evidence;
- the next recommended action;
- retry/review alternatives when useful.

Do not display a fabricated precise band change after every micro-activity.

# Skill-page contract

Each skill page contains:

```text
Target condition
Current estimate / readiness state
Evidence state
Canonical capability map
Current blockers / unknowns
Recommended actions
Practice shortcuts
Recent attempts
Certification history
```

Skill pages may differ because Listening/Reading are objective-receptive while Writing/Speaking are productive and criterion-based.

# Review experience

The Review surface combines three queues while preserving their semantics:

1. Knowledge retrieval — vocabulary/grammar/phonology suitable for review;
2. Error remediation — recurring error/remediation patterns;
3. Re-evidence — stale, conflicting, or insufficient claims.

The UI may merge presentation, but scheduling/evidence semantics remain distinct.

# Exam-preparation experience

Exam Preparation may increase:

- timed sections;
- full-task Writing;
- complete Speaking simulations;
- integrated receptive sections;
- full mocks;
- pacing/stamina work;
- focused preparation for a selected One Skill Retake.

It does not change canonical mastery thresholds and does not convert one mock into certification.

# Progress presentation

Progress is action-first. The learner should see:

```text
where I am
what blocks the target
what changed
what evidence is missing
what to do next
```

Recommended labels include:

- **Not sampled yet**;
- **Needs more evidence**;
- **Developing**;
- **Ready to reassess**;
- **Currently supported at Band N**;
- **Evidence stale**;
- **Evidence conflicting**;
- **Product path not yet supported** when a CoverageGap blocks the target.

Numeric confidence is shown only when meaningful and calibrated.

# No guaranteed-band promise

The app guarantees only process semantics it controls: preserving the target, following prerequisite/evidence rules, providing eligible next actions, and telling the truth about support/readiness state.

It cannot guarantee an external test result.

Forbidden learner promise without appropriate empirical evidence:

```text
Follow this plan and you will get Band N.
```

Allowed:

```text
Your target is Band N.
These are the unresolved conditions.
This is the highest-value next action.
Your current evidence supports / does not yet support the target claim.
```

# Product defaults vs learning truth

Durations, block counts, navigation order, and dashboard layout are product defaults. They may change through product evidence without redefining IELTS, mastery, prerequisites, or coverage truth.