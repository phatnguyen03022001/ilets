# Behavior

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> `TASK-0003` revision 1 migration artifact. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Canonical target snapshot: `phatnguyen03022001/ilets@2f54b7c6003399ccf7abff8cda2277d68d0048e8`.
>
> Model pin: `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07` (unreleased V1 candidate).

This draft maps the smallest coherent target-to-evidence-to-next-action behavior slice. `docs/catalog/project.json` owns migration identities and typed relations; this file owns migrated behavior explanation only.

## BEHAVIOR-FUNCTIONAL Functional behavior

### FTR-001 Target profile setup

The learner may declare Academic or General Training when known, an overall target and/or real per-skill minimums, and other target conditions that materially affect preparation. Unknown target fields remain unresolved; no hidden variant or Band default is invented. The declared target remains stable until the learner explicitly changes it.

Sources: `OBJECTIVE.md` — Learner route; `design/00-learning-experience.md` — TargetProfile, First-run experience.

### FTR-002 Diagnostic entry and evidence picture

The learner may choose a quick provisional diagnostic or a fuller baseline. Sampling follows the known target context. Completion does not imply complete sampling, a complete learner model, or certification. Not-sampled, unusable, pending, conflicting, and stale conditions remain explicit.

Sources: `design/00-learning-experience.md` — Diagnostic choice; `spec/08-ASSESSMENT.md` — Diagnostic checkpoint.

### FTR-003 Eligible Daily Plan recommendation

The product recommends current work from target context, evidence state, gaps/unknowns, Required prerequisites, due review, available time, and current product/capability availability. Ranking orders only eligible candidates. A saved plan is a recommendation snapshot; current hard eligibility is rechecked before assignment.

Sources: `design/00-learning-experience.md` — Today-first home, Strong recommendation + learner agency; `design/04-application-flows.md` — Ranker non-authority invariant, Flow B.

### FTR-004 Governed learning and practice with learner agency

The learner may Swap, Skip, Shorten, or Change-skill only among eligible choices. Skipping or abandoning work does not satisfy a Required prerequisite or become ability evidence. Completion is history, not mastery. Results distinguish what was observed, what changed, what remains unresolved, and the next eligible action without fabricating a micro Band change.

Sources: `design/00-learning-experience.md` — learner agency, Learning-session experience; `design/04-application-flows.md` — Flow C; `spec/09-PROGRESSION.md` — Learner behavior boundary.

### FTR-005 Evidence interpretation and re-planning

An Attempt may produce an Observation, but only claim-scoped Assessment admission creates an EvidenceFact. Assessment preserves insufficient, conflicting, stale, below-threshold, and supported states; Progression derives only justified learner-state interpretation, GapEvaluation, and ActionIntent. Missing evidence or product/evaluator inability is not learner weakness. The planner consumes that truthful state to derive the next eligible plan without changing the target.

Sources: `spec/08-ASSESSMENT.md` — core objects, ReadinessEvaluation; `spec/09-PROGRESSION.md` — State layers, Assessment state → learner gap; `design/04-application-flows.md` — Flow J.

### FTR-006 Bounded AI and optional human support

AI may explain, hint, generate bounded practice, or summarize feedback, but the learner performs the target operation. Assistance is reduced or disabled where independent evidence requires it. AI output cannot create a Band claim, EvidenceFact, learner gap, or next-action truth by assertion. The ordinary supported route does not require a teacher; optional human expert input gains evidence consequence only through normal Assessment admission.

Sources: `docs/PRODUCT.md` — AI support role, Optional human expert role; `design/00-learning-experience.md` — AI-first ordinary route; `spec/08-ASSESSMENT.md` — Human verification.

## BEHAVIOR-STATES State and transition behavior

- Target fields may be resolved or explicitly unresolved; missing material fields are never silently filled.
- Diagnostic conditions may be sampled, not sampled, unusable, pending, conflicting, or stale.
- A Daily Plan is a recommendation snapshot; assignment rechecks current eligibility.
- Learner activity creates Attempt history; measurement creates Observation; claim-scoped admission creates EvidenceFact.
- Current claim interpretation remains one of the canonical Assessment states, and Progression changes interpretation only from justified Assessment state while retaining history.

Sources: `design/00-learning-experience.md`; `design/04-application-flows.md` — Stage 7; `spec/08-ASSESSMENT.md`; `spec/09-PROGRESSION.md`.

## BEHAVIOR-INVARIANTS Invariants and permissions

- The declared target changes only through explicit learner action.
- Learner preference, Skip/Swap/Shorten/Change-skill, friction, or abandonment cannot make ineligible work eligible, bypass Required prerequisites, or become ability evidence.
- Rankers, AI, evaluators, transport, and product availability do not own IELTS learning/evidence truth.
- Completion, content consumption, time spent, AI output, or one favorable micro-activity does not certify capability or Band.
- Missing, stale, conflicting, pending, unusable, or uncalibrated evidence remains unresolved under its actual reason.
- Product/evaluator/provider inability is a product condition, not negative learner evidence.
- No product statement guarantees a future external IELTS result.

Sources: `OBJECTIVE.md`; `design/00-learning-experience.md`; `design/04-application-flows.md`; `spec/08-ASSESSMENT.md`; `spec/09-PROGRESSION.md`.

## BEHAVIOR-FAILURES Errors, edges, and failure behavior

- Missing target fields keep the affected target consequence unresolved; only genuinely compatible foundational work may continue.
- Diagnostic completion preserves not-sampled, unusable, pending, conflicting, or stale conditions.
- If a saved plan item becomes ineligible before assignment, the product reselects another eligible action or exposes the blocker.
- Provider, evaluator, capture, content, or product inability may yield pending, degraded, unavailable, or unresolved product state; it never creates a low score or `ABILITY_GAP`.
- Stale or superseded evaluator results cannot overwrite accepted current state or independently create evidence.
- Consequential learner-facing failure states explain what happened, why it matters, and the next valid action when one exists.

Sources: `design/00-learning-experience.md` — Consequential learner-state explanation contract; `design/04-application-flows.md` — Stage 7, Flow D, Flow J; `spec/08-ASSESSMENT.md`.

## BEHAVIOR-FLOWS Critical flows

### FLW-001 Governed target-to-next-action loop

Set/confirm the known target → obtain an evidence picture → preserve blockers/unknowns → select eligible work → learn/practise/review/assess → interpret only what the attempt established → derive the next eligible recommendation. The loop never changes the target implicitly.

### FLW-002 Learner control within eligibility

The learner may choose another eligible activity, skip, shorten at a safe point, change eligible skill focus, or explicitly edit the target. Required or unsupported conditions remain visible and unsatisfied until actually resolved.

### FLW-003 Truthful unresolved and failure handling

When target input, evidence, capture, evaluator, provider, content, or product support is insufficient for the requested consequence, preserve the correct unresolved/product state and offer valid recovery when available. Product inability never becomes learner weakness or fabricated evidence.

### FLW-004 Evidence-to-replan reconciliation

Eligible measurement → Observation → claim-scoped Assessment admission/evaluation → justified Progression interpretation → currently eligible planning. No downstream layer broadens evidence inference or lowers the target to make planning easier.

Flow sources: `OBJECTIVE.md`; `design/00-learning-experience.md`; `design/04-application-flows.md` — Flows B/C/D/J; `spec/08-ASSESSMENT.md`; `spec/09-PROGRESSION.md`.

## BEHAVIOR-ACCEPTANCE Acceptance criteria

### ACC-001 Explicit target semantics

Unknown target fields remain unknown, no hidden Academic/GT or Band requirement is invented, and the stored target changes only from explicit learner action.

### ACC-002 Evidence-bounded diagnostic

Diagnostic completion preserves unsampled, unusable, pending, conflicting, and stale material conditions and makes no certification claim merely because the flow ended.

### ACC-003 Eligibility-preserving planning

Every recommended/executed activity satisfies current eligibility; ranking or learner preference cannot bypass Required prerequisites, target constraints, or current hard conditions.

### ACC-004 Activity outcome truth

Skip, abandonment, completion, time spent, or one favorable micro-activity does not itself satisfy a prerequisite, create an ability conclusion, or certify Band.

### ACC-005 Claim-scoped evidence and next action

Only Assessment-admitted evidence changes claim interpretation; Progression does not broaden inference scope; unresolved evidence/product inability stays distinct from demonstrated learner weakness; re-planning consumes the truthful state.

### ACC-006 Bounded assistance

AI/human support cannot bypass Assessment admission or become learning/evidence authority. Independent evidence preserves the learner's target operation and required assistance limits; ordinary product use does not depend on hidden mandatory human scoring.

### ACC-007 Promise boundary

No migrated behavior claims that product use, a favorable diagnostic/mock, internal support, or readiness guarantees a future official IELTS result.

## BEHAVIOR-SAFETY Safety and human control

The learner controls explicit target changes and may choose among eligible study paths, while prerequisite, evidence, target-integrity, and product-support boundaries remain non-bypassable. AI is bounded assistance; optional human input is not silently mandatory. Failure and uncertainty are surfaced without blaming the learner, inventing certainty, or hiding product limitations.

Sources: `docs/PRODUCT.md` — Actors and roles, Domain and external constraints; `design/00-learning-experience.md`; `spec/08-ASSESSMENT.md`; `spec/09-PROGRESSION.md`.

## Migration boundary

This file, `docs/PRODUCT.md`, `docs/DECISIONS.md`, and `docs/catalog/project.json` remain migration artifacts with **authority NONE**. The catalog owns `ACT-*`, `ROL-*`, `FTR-*`, `ACC-*`, `FLW-*`, and `UNK-*` migration identities and typed relations.

For material feature relations whose target identity domains are not yet migrated, the catalog uses `unresolved_ref` to an existing `OPEN` `DESIGN` `UNK-*`. Those relations intentionally block documentation closure; they are not N/A and do not create substitute `DAT-*`, `IFC-*`, `EXT-*`, or `CAP-*` identities.

This state does not claim `DOCS_READY`, documentation/design lock, implementation readiness, assurance status, maturity, promotion, or release readiness.
