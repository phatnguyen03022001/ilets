STATUS: CANONICAL
OWNS: end-to-end learner product journey, TargetProfile product semantics, navigation surfaces, progressive-disclosure UX, learner-facing visual/interaction language, study-session UX shapes, product timing defaults, learner agency, learner-visible AI tutoring behavior, entitlement-visible availability, learner-visible target-trajectory advisory, strategy-change explanation, and user-visible interpretation of learning/product state
DEPENDS_ON: ../spec/00-PRODUCT.md, ../spec/01-LEARNER-MODEL.md, ../spec/07-PRACTICE.md, ../spec/08-ASSESSMENT.md, ../spec/09-PROGRESSION.md, ../spec/02-IELTS-MODEL.md
DOES_NOT_OWN: skill definitions, mastery thresholds, evidence sufficiency, practice taxonomy, planner decision internals, legal runtime lifecycle transitions, coverage declarations, commercial pricing/tier policy, authorization role matrices, provider selection, API wire contracts, persistence, frameworks, deployment, pixel-level design tokens, or component implementation

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

# Learner-facing visual and interaction language

The product uses one calm, precise visual language across ordinary study. The reference direction is translated into durable qualities—restraint, clear components, high finish, and subtle depth—rather than copied from another product. Premium quality means the interface feels deliberate and trustworthy without luxury theatrics, cinematic presentation, or generic “AI product” styling.

## Primary visual character

The default learner surface is typography-led and content-first:

- hierarchy comes primarily from type, spacing, grouping, and clear primary/secondary action relationships;
- neutral surfaces, restrained borders/elevation, deliberate whitespace, and consistent component behavior create structure without turning every concept into a card;
- progressive disclosure keeps deep semantics available without making the default view dense;
- readability and long-session comfort outrank decorative novelty;
- one coherent interaction/component language is reused across Today, Skills, Practice, Review, Media, Progress, Mock, and activity workspaces.

Avoid childish edtech styling, decorative gamification without learning value, excessive gradients/glass effects, glowing/futuristic AI motifs, dense dashboard mosaics, badge/icon clutter, and visual treatment that exposes internal ontology merely because it exists.

These are product-quality semantics, not a fixed token catalogue. Exact color values, type scale, spacing scale, radii, shadows, and component implementation remain downstream design/implementation choices provided they preserve this character and the accessibility/performance boundaries below.

## Capability-first learner language

Normal learner-facing copy names the capability, action, or result rather than advertising the implementation mechanism. Suitable language includes concepts such as **Explain**, **Analyze**, **Feedback**, **Practice**, **Suggestion**, **Next step**, **Conversation**, **Writing feedback**, **Speaking practice**, **Tutor**, and **Help**.

Generic promotional labels such as **AI Tutor**, **AI Coach**, **AI Analysis**, **AI-powered**, **Powered by AI**, or **Ask AI** are not normal product chrome merely because a model participates. The learner experiences one IELTS product rather than a collection of AI-labelled features.

This naming rule is presentation-only:

- it never hides a material evaluator/product limitation or changes Assessment/evidence authority;
- explicit model/automation disclosure remains available where required for legal, privacy, provenance, safety, settings, support, or informed-use purposes;
- internal architecture may still name AI/model capabilities precisely even when ordinary learner copy does not.

## Iconography

Icons are functional aids, not decoration.

- prefer text when text communicates the action/status more clearly;
- use one restrained coherent icon family for scanning, navigation, action recognition, or status interpretation;
- do not use emoji, cartoon imagery, mixed icon packs, or an icon beside every heading/card as normal product chrome;
- internal taxonomy must not become dozens of learner-facing symbols;
- required status or action meaning is always available in text or another accessible non-icon-only form.

## Imagery and background depth

The product may use sparse high-quality imagery or atmospheric depth when it improves orientation, tone, or context without competing with learning. Suitable treatments include subtle editorial/contextual photography, abstract light/material/architectural imagery, soft texture, restrained atmospheric backgrounds, and low-contrast visual fields.

Rules:

- learning content and controls remain the foreground and text contrast/readability always wins;
- background imagery never carries required state, instruction, or evidence meaning;
- avoid stock-looking decoration, futuristic/robot/AI imagery, repeated large hero art, or constant visual spectacle;
- prefer a few memorable visual moments over persistent decoration;
- Today, activity workspaces, Writing, Speaking, and Review remain visually quiet enough for long study sessions.

## Motion and responsiveness

Motion exists to communicate state transition, hierarchy, completion, feedback, expansion/collapse, or navigation continuity. It is not spectacle.

Avoid scroll-driven showpieces, unnecessary parallax, atmospheric background video, long blocking transitions, excessive blur/compositing, or animation that delays study. Navigation and controls should feel immediate, layout should remain stable, loading should be predictable, and ordinary scrolling/input should not visibly jank or jump.

Reduced-motion preferences are respected gracefully. If visual motion conflicts with responsiveness, input latency, readability, accessibility, or stable layout, responsiveness and usability win. Concrete implementation/performance mechanisms remain owned by the Web/implementation boundary.

## Interaction micro-quality

Every ordinary surface applies the same quality bar to non-happy paths and control states:

- loading/skeleton treatment preserves useful layout shape without pretending unknown content is already known;
- empty states explain what is absent and, when useful, the next valid action rather than showing decorative emptiness;
- error, degraded, disabled, pending, and unavailable states remain distinguishable when their learner consequence differs;
- hover, focus, pressed, selected, expanded, and disabled states are perceivable and behaviorally consistent;
- keyboard/focus/screen-reader accessibility is preserved where applicable and status is not communicated by color/icon alone;
- copy remains concise and action-oriented while preserving important uncertainty/limitation semantics;
- primary actions are visually clear without making secondary/recovery actions hard to find;
- responsive/mobile layouts preserve task hierarchy, readable text, stable controls, and safe input/capture rather than merely shrinking desktop composition.

Micro-quality must remain consistent across long sessions; visual polish may not introduce avoidable fatigue, content jumping, ambiguous controls, or inaccessible interaction.

## Beginner and L1 scaffold compatibility

The same visual language supports beginner scaffolding; there is no separate beginner design system. When L1 support materially improves understanding, the product may present a simple L1-first explanation while preserving important canonical English/IELTS terminology and allowing deeper detail on demand.

A typical disclosure pattern is:

```text
Task Response  [info]
  → simple learner-language / L1 explanation
  → tiny concrete example when useful
  → deeper canonical detail on demand
```

The information control is functional and accessible rather than decorative. As learner independence grows, L1 explanation may become less prominent or user-invoked without changing the underlying canonical target. L1 support remains the localization/explanation overlay defined by `../spec/00-PRODUCT.md`; it never creates a different IELTS standard.

## Simple surface / deep system invariant

The learner interface exposes only concepts that help the learner understand state or act correctly. Internal structures such as Skill Leaves, Knowledge Objects, Practice Types, Learning Mechanisms, EvidenceRequirement, GapEvaluation, and ActionIntent do not become menu/navigation complexity or visual badges by default.

The primary learner surface remains centered on **Today, Skills, Practice, Review, Media, Progress, Mock**, and the current next action. Internal depth is reached progressively through plain learner-facing concepts and explanations rather than by exposing implementation/canonical taxonomy directly.

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

A learner entitlement may make optional cost-intensive experiences available without changing the learner model or standard. The ungated baseline remains a genuinely usable IELTS learning product: the learner can set/understand the target, see truthful current evidence/gaps/unknowns, receive an eligible ordinary next action, learn/practise/review through the supported ordinary route, and inspect applicable progress/history without buying epistemic truth.

Learner-facing rules are:

- unavailable paid capability is presented as product availability, never learner weakness;
- activation/restoration makes future eligible gated capability available only after authoritative product reconciliation;
- pending/ambiguous commercial state is shown as pending rather than fabricated active/expired access;
- expiry/downgrade may block new gated sessions while an already accepted session follows the safe bounded continuation/termination semantics in `04-application-flows.md`;
- paid access may primarily increase expensive AI/realtime use, productive-evaluation volume/depth, adaptive tutoring compute, media transformation, assessment/mock capacity, speed/priority/quota, extended history, or analytics where justified;
- free/paid labels do not alter target, prerequisites, evidence eligibility, Band standards, product-support truth, or historical state;
- losing entitlement does not erase attempts/evidence/saved learning history and does not by itself remove normal learner access to applicable history/export/deletion controls;
- basic understanding of what is missing/unresolved is not paywalled merely to create an upgrade incentive;
- learner subscription/entitlement never grants content-operation/admin authority.

Concrete pricing, quotas, marketing tier names, launch discounts, and commercial packaging remain outside this owner. Effective entitlement lifecycle and operational authorization separation are owned by `04-application-flows.md`.

# Daily study presets

Product defaults provide predictable session sizes without becoming learning dosage laws.

## Daily load semantics

A DailyPlan decides both **what** eligible work is useful now and **how much** of that work fits the learner's current planning envelope. Learner-visible load may include any justified mix of:

- due vocabulary/Knowledge retrieval and review;
- grammar or other prerequisite acquisition/practice;
- cohesion/linking-language or other focused productive work;
- topic/context exposure needed for transfer;
- Listening/Reading practice;
- Writing/Speaking production;
- reassessment/re-evidence;
- due remediation/review;
- exam-preparation/timed work.

The Planner derives load from the current TargetProfile, current evidence/GapEvaluation/ActionIntent, Required prerequisites, due review state, learner available time, test-date/target urgency, activity difficulty/load estimate, recent fatigue/session coherence, retention/performance history, and need for transfer or re-evidence. The learner may shorten/swap only within the eligibility rules already defined above; skipped required work remains visible.

Keep these meanings separate:

```text
Band requirement        = quality/capability threshold
content inventory       = available eligible opportunities
daily dosage            = versioned planning/load policy
learner-specific plan   = today's composition from current state
```

No Band implies a universal word count, exercise count, or items-per-day requirement. Numeric dosage, load estimates, review limits, and session-composition defaults may be versioned/tuned from product and learning evidence; until calibrated, they remain bounded product policy rather than invented learning truth. System-level composition semantics are owned by `04-application-flows.md`.

## Target trajectory advisory

When the target/date context and empirical basis are sufficient, the product may answer a narrower planning question:

> Given what the system currently knows, does the current study envelope appear plausibly consistent with the learner's declared target and date?

This is a product/planning advisory interpretation. It is not IELTS evidence, a Band claim, readiness evidence, a guaranteed score prediction, or authority to change TargetProfile.

The advisory may consume only justified current information such as TargetProfile, time remaining, supported and unresolved target conditions, `GapEvaluation`/`ActionIntent`, available study envelope, recent admissible progression/history, planned workload, and calibrated empirical rate/load assumptions when those assumptions actually exist.

Learner-visible interpretation preserves at least these distinctions without requiring one permanent enum vocabulary:

- there is not yet enough evidence/history/calibration to estimate trajectory honestly;
- the current planning envelope appears consistent with the target trajectory under the named projection policy;
- the trajectory is materially at risk for explained reasons;
- known required work exceeds the current time/planning envelope under the named projection policy.

Rules:

- insufficient learner evidence increases projection uncertainty; it is not assumed weakness or invented progress velocity;
- evaluator/content/product unavailability remains a separate product limitation and must not be represented as slow learner progress;
- an at-risk/insufficient-envelope message explains the material causes rather than presenting an opaque warning;
- learner-visible options may include increasing available study time, changing scheduling/focus within eligible work, collecting missing evidence, or explicitly changing the test date/target;
- the system never changes the target or date without learner action;
- do not claim that a target is impossible unless a separately justified policy can support that claim; normally report that the current envelope/projection is insufficient or unresolved;
- projection policy/version plus the material input/provenance used for a consequential learner-facing advisory remain reconstructable.

Exact forecast coefficients and rate/load assumptions are empirical/versioned product policy, not canonical Band or progression truth. System-level projection composition is owned by `04-application-flows.md`.

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

When comparable intervention history justifies a strategy change, learner-facing explanation should state that the product is changing approach because the previous approach has not produced sufficient improvement under the current learning policy, and identify the meaningful change when useful (for example scaffold, practice type, mechanism, context, or load). This explanation must not imply that the learner became weaker merely because the prior intervention produced limited improvement. When history is insufficient, do not label the learner as plateaued.

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