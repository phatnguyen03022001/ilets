STATUS: CANONICAL
OWNS: end-to-end learner product journey, TargetProfile product semantics, navigation surfaces, progressive-disclosure UX, study-session UX shapes, product timing defaults, learner agency, learner-visible AI tutoring behavior, entitlement-visible availability, and user-visible interpretation of learning/product state
DEPENDS_ON: ../spec/01-LEARNER-MODEL.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/02-IELTS-MODEL.md
DOES_NOT_OWN: skill definitions, mastery thresholds, evidence sufficiency, practice taxonomy, planner decision internals, legal runtime lifecycle transitions, coverage declarations, commercial pricing/tier policy, authorization role matrices, provider selection, API wire contracts, persistence, frameworks, or deployment

# Learning Experience

## Purpose

Define how a learner experiences the system from target setup through daily learning, assessment, review, and exam preparation.

This document owns learner-visible product semantics. It does not redefine learning truth, evidence policy, planner internals, runtime state machines, or product coverage.

# Learner-visible product loop

```text
set target
  ↓
get an initial evidence picture
  ↓
see the highest-value eligible next work
  ↓
learn / practise / review / assess
  ↓
see what the attempt showed
  ↓
see what remains unknown or blocking
  ↓
continue toward the unchanged target
```

The learner should always be able to answer:

1. What target am I working toward?
2. What is currently blocking or uncertain?
3. Why is this activity recommended now?
4. What did the last attempt actually show?
5. What should I do next?

# Learner mental model and progressive disclosure

The product presents one simple route over a deeper model:

```text
target
→ four-skill current state
→ blockers / unknowns / due review
→ required capability or micro-skill
→ enabling knowledge + evidence detail when needed
→ recommended work
→ next action
```

`Micro-skill` is learner-facing language for a canonical Skill Leaf; it is not a second ontology. Grammar, vocabulary, and phonology appear as enabling knowledge when they materially explain or block a skill outcome, not as a fifth IELTS score.

Progressive-disclosure rules:

1. default surfaces show target, current state, the few material blockers/unknowns, and next action;
2. expanding a skill reveals its capability map, required Knowledge Objects, evidence state, and recent attempts;
3. evidence/provenance/policy detail remains available for explanation without becoming the default learner view;
4. internal IDs and ontology structure are hidden unless useful for support/debugging;
5. no completeness percentage is inferred from the fraction of visible nodes completed.

# `TargetProfile`

A learner goal is a constraint profile, not one overall Band number.

Conceptual fields:

```text
test_variant                 Academic | General Training
delivery_mode                optional; external mode relevant to booked/intended test
purpose_or_receiving_rule    optional external constraint reference
target_overall_band          optional minimum overall-band constraint
minimum_listening_band       optional minimum section-band constraint
minimum_reading_band         optional minimum section-band constraint
minimum_writing_band         optional minimum section-band constraint
minimum_speaking_band        optional minimum section-band constraint
test_date                    optional
selected_skill_retake        optional preparation focus on one existing skill
```

## Target-resolution minimum

A persisted TargetProfile used for **target-relative planning, readiness, or product-support evaluation** contains:

1. a resolved standard `test_variant`: Academic or General Training; and
2. at least one actual Band constraint: `target_overall_band` and/or one real per-skill minimum.

If the learner does not yet know the variant, the product keeps the variant-specific target unresolved. Shared/foundational diagnostic or learning work may still be offered where it is genuinely variant-independent, but the system must not choose Academic/GT Task-1/Reading conditions silently or claim complete target readiness/support.

If the learner knows the variant but no Band constraint yet, diagnostic/foundational work may still be offered against that partial context, but the product must not fabricate a readiness target.

## Band-constraint semantics

Band constraints are **lower bounds**, not exact-equality goals:

```text
target_overall_band = 7.0       means required/desired overall result >= 7.0
minimum_writing_band = 6.5      means Writing result >= 6.5
```

An overall target alone does **not** imply four equal per-skill minima. Multiple section-Band combinations can satisfy the same official overall Band after the applicable rounding rule.

If planning requires working per-skill targets while only an overall target is known, the product may propose a separately labelled **planning profile**. A planning profile:

- is a product planning choice, not an external requirement;
- does not mutate TargetProfile per-skill minima unless the learner explicitly adopts them;
- must preserve that multiple valid four-skill combinations may satisfy the overall target;
- cannot be used as evidence that a receiving organisation requires those per-skill values.

An overall-target readiness statement may be derived only from current supported section-Band claims plus the applicable official overall-score rule and any real per-skill minima. Self-estimates, planner working targets, or unsupported point predictions cannot be substituted for section evidence.

## Target-readiness composition

Learner-facing target state is composed from claim-scoped Assessment results; it is not one synthetic mastery score.

- **per-skill minimum** — the corresponding current per-skill Band claim must be `SUPPORTED` at or above the declared minimum;
- **overall-only target** — current supported Band claims for all four scored skills must exist so the official overall-score rule can be applied; the derived overall must meet the declared minimum; no hidden equal per-skill minima are invented;
- **mixed overall + per-skill constraints** — both the derived overall condition and every explicit per-skill minimum must be satisfied;
- **variant-specific condition** — evidence/claim scope must match Academic/GT where the construct or scoring differs;
- **delivery-condition readiness** — required only when the declared target includes a material delivery interaction; it is an additional exam-condition claim, not a different Band standard.

Any required claim that is insufficient, conflicting, or stale keeps the corresponding target condition unresolved. Evidence positively below threshold is shown as not yet supported. A product CoverageGap is shown separately as **product path unavailable/not yet supported**, never as learner weakness.

The UI distinguishes:

```text
current capability evidence
current per-skill Band support
declared TargetProfile condition support
exam-condition / full-test readiness
product support for serving that route
historical official/internal attainment
```

A learner may have current Band support while still having an exam-condition gap, or may have historical attainment while current evidence needs refresh.

## General rules

1. overall target, per-skill minima, or both may be supplied;
2. real receiving-organisation/visa/employer minima are represented explicitly rather than inferred;
3. an overall target alone does not uniquely determine four skill targets;
4. the product must not fabricate hidden per-skill requirements;
5. a balanced planning profile may be proposed only under the separate planning-profile semantics above;
6. the target remains stable until the learner explicitly changes it;
7. current readiness is evaluated against TargetProfile conditions, not an internal mastery average;
8. delivery mode is included only when it changes exam-preparation interaction or an external acceptance/eligibility condition;
9. delivery mode never changes the canonical Band standard.

## One Skill Retake focus

`selected_skill_retake` selects preparation focus on one existing IELTS skill. It does **not** by itself assert that the learner is eligible for One Skill Retake.

When the product makes an eligibility-sensitive support/readiness statement, applicable original-test, timing-window, participating-location, delivery-mode, purpose/acceptance, and other current external conditions must be resolved from explicit known data/current external truth. Missing eligibility information remains unresolved; it is never inferred from the selected skill alone.

Current external variant/delivery/One-Skill-Retake facts are owned by `../spec/02-IELTS-MODEL.md`. Current product support for a TargetProfile is owned by `08-coverage-and-support.md`.

# First-run experience

## Goal setup

Capture only information that changes planning or target interpretation:

- Academic or General Training when known;
- overall target and/or real per-skill minima when known;
- receiving/purpose requirement when relevant;
- intended/booked delivery mode when known and material;
- test date if fixed;
- selected One Skill Retake focus if applicable, without implying eligibility;
- self-estimate, clearly non-authoritative;
- available study time;
- accessibility requirements;
- optional L1 context and temporary focus preference.

Unknown target fields remain unknown; onboarding convenience never licenses a hidden default variant or Band requirement.

Target setup should normally remain a short flow; the initial UX target is roughly **3–5 minutes** when the learner already knows their requirements.

## Diagnostic choice

The product exposes two entry shapes:

| Mode | UX target | Learner promise |
|---|---:|---|
| Quick start | about 15 min | enough sampling for a useful provisional route; unresolved conditions remain explicit |
| Full baseline | about 60 min target window | broader sampling across the target profile; still bounded by actual evidence obtained |

These are UX defaults, not Assessment sufficiency rules.

A completed diagnostic may still contain `not sampled`, unusable capture, pending productive evaluation, conflicting, or stale conditions. Exact diagnostic evidence semantics are owned by `../spec/08-ASSESSMENT.md`.

# Today-first home

The default landing surface is **Today**: a strong current recommendation rather than a content library that forces the learner to invent a route.

Primary destinations:

1. **Today** — current plan and resume state;
2. **Skills** — Listening, Reading, Writing, Speaking profiles;
3. **Practice** — concrete user-facing practice modes;
4. **Review** — due retrieval, remediation, and re-evidence;
5. **Media** — eligible embed/source-backed learning;
6. **Progress** — target conditions, evidence state, certification history, blockers;
7. **Mock** — section/full-test readiness activity.

Direct Practice browsing is allowed but does not silently satisfy unrelated target conditions or replace the governed route.

# Strong recommendation + learner agency

The learner sees a strong recommendation produced by the Planner contract in `04-application-flows.md`.

This experience must expose enough reason information to explain the recommendation without exposing implementation internals.

The learner may:

- Swap to another **eligible** activity;
- Skip an activity;
- Shorten at a safe stopping point;
- Change skill focus among eligible options;
- edit the TargetProfile.

UX invariants:

- Required prerequisites remain unresolved when skipped;
- a skipped target requirement remains visible;
- preference cannot certify an unrelated target;
- repeated skip/abandonment is a friction/adherence signal, not ability evidence;
- an uncovered product condition is shown as a product limitation, never learner weakness;
- the UI cannot make an ineligible activity eligible.

Eligibility, ranking, reason-code construction, and legal plan execution are owned by `04-application-flows.md`.

# AI-first ordinary route

AI assistance is available as a cross-surface tutor, not a separate learning authority. In eligible learning/practice/review contexts it may:

- explain a concept or error in learner-appropriate language;
- show a worked example or contrast;
- ask guiding questions or give progressively stronger hints;
- generate or adapt bounded practice that still validates against canonical targets/content rules;
- summarize feedback and recurring error patterns;
- help turn suitable vocabulary/grammar/phonology material into later review.

The learner must perform the target cognitive/performance operation. During independent evidence/readiness work, assistance is reduced or disabled according to the activity/evidence configuration. AI output may explain Assessment/Progression results but cannot create a Band claim, EvidenceFact, gap, or next-action truth by assertion.

Ordinary learning must remain usable without a teacher. Human help is an optional coaching preference or a consequence-specific Assessment escalation, not the default route dependency.

# Speaking interaction experience

The default Speaking experience remains usable through record → submit → feedback without realtime AI. Realtime conversation is an optional higher-cost interaction for eligible learners, not the primary route the rest of the product depends on.

Learner-visible realtime behavior:

- show whose turn it is and preserve Part 1/2/3 activity semantics rather than presenting one generic conversation;
- show reconnecting/delayed/degraded states when AI/network latency interrupts responsiveness;
- after a dropped session, resume the same logical interaction when safe or explain that only a partial session was completed;
- offer record/submit fallback when the selected learning purpose remains valid without realtime; otherwise preserve the unmet readiness/interaction condition rather than pretending fallback was equivalent;
- distinguish learner silence from “we could not hear/capture you”;
- post-session feedback reflects only turns/signals actually available and never upgrades the session into evidence by presentation.

Capture/product messages remain distinct enough for the learner to act correctly, for example:

```text
Microphone permission/device unavailable
Recording interrupted or audio not usable — retry capture
Audio captured; transcription unavailable
Audio captured; pronunciation/acoustic analysis unavailable
Realtime conversation temporarily unavailable — ordinary Speaking still available
This capability requires an active entitlement
Product cannot currently verify this evidence condition
```

None of these messages is displayed as a low Speaking score.

# Entitlement-visible availability

A learner entitlement may make an optional cost-intensive experience available, such as realtime AI conversation, without changing the learner model or standard. Learner-facing rules are:

- unavailable paid capability is presented as product availability, never learner weakness;
- activation/restoration makes future eligible gated capability available only after authoritative product reconciliation;
- pending/ambiguous commercial state is shown as pending rather than fabricated active/expired access;
- expiry/downgrade may block new gated sessions while an already accepted session follows the safe bounded continuation/termination semantics in `04-application-flows.md`;
- free/paid labels do not alter target, prerequisites, evidence eligibility, Band standards, product-support truth, or historical state;
- losing entitlement does not erase attempts/evidence/saved learning history and does not by itself remove normal learner access to applicable history/export/deletion controls;
- learner subscription/entitlement never grants content-operation/admin authority.

Concrete pricing/tier definitions remain outside this owner. Effective entitlement lifecycle and operational authorization separation are owned by `04-application-flows.md`.

# Daily study presets

Product defaults provide predictable session sizes without becoming learning dosage laws.

## Quick — 10 minutes

Typical envelope:

```text
brief review / retrieval
one highest-value action
result + next-action summary
```

Default maximum: **3 activity blocks**.

## Standard — 25 minutes

Typical envelope:

```text
review
acquire/remediate where needed
independent practice/transfer
checkpoint + next plan
```

Default maximum: **4 activity blocks**.

## Deep — 45 minutes

Typical envelope:

```text
review
acquire/remediate
focused practice
transfer/fluency/timed work
checkpoint
```

Default maximum: **5 activity blocks**.

The Planner may compose these differently according to ActionIntent, evidence need, target urgency, and safe stopping points. Exact per-activity duration in the practice catalog remains a product default, not mastery policy.

# Learning-session experience

A learner-visible session may be shown as planned, in progress, completed, or abandoned according to the canonical runtime lifecycle in `04-application-flows.md`.

This document does not define legal state transitions.

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

Completion/abandonment is session history, not a mastery judgment.

# Activity card contract

Before an activity, show enough information for informed action:

- target skill/capability in plain language;
- why now;
- expected duration;
- primary activity purpose in learner language;
- relevant target variant/context;
- scaffold/independence state where material;
- whether the configured activity is an evidence candidate under normal Assessment admission.

After an activity, show:

- what was observed;
- what changed, if anything;
- what remains uncertain/missing;
- the next recommended action;
- retry/review alternatives when useful.

Do not display a fabricated precise Band change after each micro-activity.

# Skill-page contract

Each skill page presents, with progressive disclosure:

```text
Target condition
Current per-skill Band support or unresolved state
Capability / micro-skill support map
Exact missing/stale/conflicting evidence conditions
Task/section/exam-condition readiness where relevant
Current blockers / unknowns
Recommended evidence or learning action
Recent Attempts / Observations
EvidenceFact detail on demand
Historical certification / external result separately
```

A single sampled result is labelled by its actual scope (for example, “sampled T/F/NG classification evidence”), not as “Reading Band evidence” unless the Band-scoped requirement legitimately admits that inference.

For each consequential claim, the learner can expand **Why? / What is missing?** to see plain-language required sub-capabilities/criteria/parts, currently supporting evidence, missing/stale/conflicting conditions, and the next evidence need. Internal policy IDs may remain behind deeper detail.

Listening/Reading and Writing/Speaking may present different measurement detail because their external assessment mechanics differ.

# Review experience

One Review surface may combine presentation of three semantic queues:

1. **Knowledge retrieval** — reviewable vocabulary/grammar/phonology;
2. **Error remediation** — diagnosed recurring error/remediation work;
3. **Re-evidence** — stale, conflicting, or insufficient claims.

Presentation may be unified. Scheduling/evidence meaning must remain distinct.

# Vocabulary and knowledge experience

Vocabulary/grammar/phonology form one enabling-knowledge experience that connects acquisition to later skill use. A learner may reach it from Today, a Skill blocker, Review, Media, feedback, or an explicit save action.

Learner flow:

```text
encounter or recommended knowledge target
→ understand in context / contrast / example
→ optional save to personal study set
→ active retrieval after delay
→ controlled application
→ later use inside Listening/Reading/Writing/Speaking
```

A saved vocabulary item may preserve the expression, learner-facing meaning, useful example/context, source/provenance, and applicable canonical Knowledge references. The saved item is content, not a new Knowledge Object.

Review rules:

- suitable items may enter spaced retrieval; not every encountered word/error does;
- the system should not auto-add every unknown token or flood Review for collection metrics;
- retrieval history can schedule review but does not certify target-skill capability;
- successful card recall does not replace contextual receptive/productive evidence;
- learner-created notes/examples remain visibly learner-created when provenance matters.

# Exam-preparation experience

Exam Preparation may increase:

- computer-interface familiarity for current test-centre delivery;
- typing/navigation/timing practice where relevant;
- optional handwriting Writing practice when the learner targets an eligible Writing-on-Paper delivery;
- remote-platform familiarity when targeting eligible IELTS Online Academic;
- timed receptive sections;
- complete Writing tasks;
- complete Speaking simulations;
- full mocks and pacing/stamina work;
- selected One Skill Retake preparation.

Delivery-mode preparation is not a new learning standard. It cannot rewrite Band thresholds or turn one mock into certification.

# Progress presentation

Progress is action-first. The learner should see:

```text
where I am
what blocks the target
what changed
what evidence is missing
what product condition is unavailable, if any
what to do next
```

Suitable labels include:

- **Not sampled yet**;
- **Sampled evidence only — no Band claim**;
- **Needs more evidence**;
- **Developing**;
- **Ready to reassess**;
- **Current evidence supports Band N**;
- **Evidence stale — refresh needed**;
- **Evidence conflicting**;
- **Product cannot currently verify this claim**;
- **Product path not yet supported**.

Historical certification may remain visible separately even when current evidence is stale/conflicting and current certification is no longer active.

Numeric confidence appears only when it is meaningful and calibrated.

# Promise boundary

The product may promise process integrity it controls: preserving the target, applying prerequisites/evidence/support rules, providing valid next actions, and reporting evidence/support state truthfully.

It cannot promise a future external IELTS result.

Forbidden without suitable empirical evidence:

```text
Follow this plan and you will get Band N.
```

Allowed pattern:

```text
Your target is Band N.
These conditions remain unresolved.
This is the highest-value eligible next action.
Your current evidence supports / does not yet support the scoped claim.
```

# Product-default boundary

Durations, navigation order, dashboard layout, and activity-block counts are mutable product defaults. They may change through product evidence without redefining IELTS, Skill, Band, Assessment, Progression, or Coverage truth.