# Red Team Triage — Architectural Owner Review
**Date:** 2026-07-16 · **Owner:** Blueprint architecture · **Input:** [red-team-report.md](red-team-report.md) (preliminary, 2/6 attacks reported successful)

## Mandate & method
Architectural ownership was retained; the Red Team report was reviewed **independently and adversarially against the source**. No finding was assumed correct. Each claim was checked against the actual Blueprint object it names (schema, prerequisites, decision text). Classifications: **ACCEPT** (genuine defect → minimum correction applied) · **REJECT** (not a defect — evidence below) · **DEFER** (real but already tracked, or downstream-owned).

## Report-reliability assessment (read before the per-finding verdicts)
The Red Team report is **low-reliability** as a whole. Four independent indicators:

1. **Self-contradiction.** The report states 4 of 6 attacks "failed due to rate limiting," yet its largest section ("Supplementary Evidence") presents detailed, citation-laden findings from exactly those "failed" attacks. Either the attacks failed (so the findings are unsupported) or they did not. The report cannot have it both ways.
2. **Unverified citations.** The Supplementary Evidence cites ~15 research works (Larsen-Freeman, Krashen, Long, Cardoso, Verspoor, Birdsong, Harley, Tarone, Bryfonski & McKay, etc.) with **no audit provenance**. This violates the Blueprint's own citation standard: [learning-science-evidence-brief.md](../learning/learning-science-evidence-brief.md) documents that a research subagent's "all citations verified" claim was **not trusted**, independently audited, and a **fabricated "Corbett et al. 2017" was removed** plus four venues corrected. Per that precedent, the Red Team's unaudited citations are treated as unverified and are **not propagated**.
3. **Factual errors verified against source** (see F5, F6, F13, F14): wrong object counts, an invented prerequisite edge, a false "circular" claim, and a false "disjoint silos" claim.
4. **Severity inflation of already-documented items.** Findings 1, 4, 7, 10, 12, 15 re-discover open items already triaged in [validation.md](validation.md) at Low/Medium (V-L-01/02/03, Band-5 density, K-VOC-012 split, cross-skill open-question) and re-file them as CRITICAL/HIGH.

Because of (1)–(4), the headline verdict ("3 CRITICAL; NOT PRODUCTION READY; 6.5/10") is **not sustained** (see Verdict).

## Classification summary

| # | Red-Team claim (their severity) | Verdict | Basis (one line) |
|---|---|---|---|
| **F2** | No certification revocation on regression (CRITICAL) | **ACCEPT (→ MEDIUM)** | Genuine gap: no explicit `BandCertificationState` transition on regression. Corrected. |
| F1 | No plateau diagnosis (CRITICAL) | **DEFER** | Already V-L-02 (Low). Time-in-band framing adopted for calibration. |
| F3 | False independence, PG-002 (CRITICAL) | **REJECT** | Conflates certification-independence (faithful to IELTS) with structural-independence; skills integrate via the shared K-* layer. |
| F4 | K-VOC-012 supernode (HIGH) | **DEFER** | Already an open question ([vocabulary.md:83](../knowledge/vocabulary.md)); knowledge-object ≠ delivery unit. |
| F5 | S-GRA-02 under-specified (HIGH) | **REJECT** | Wrong facts: 3 objects, not 2; K-GRA-005 already present. Asymmetry is defensible. |
| F6 | R-COMP-04 ↔ W-TR-04 cycle (HIGH) | **REJECT** | No edge connects them (R-COMP-04→R-COMP-03; W-TR-04→W-TR-01). No cycle. |
| F7 | Exam-prep entry safeguards (HIGH) | **DEFER** | LD-005 already constrains exam-prep; "reality-check UX" is product-scope. Builds on V-L-01. |
| F8 | Spaced-review breakdown (HIGH) | **REJECT** | Performance-graded spacing is accuracy-triggered — it *is* the adaptation for variable forgetting. |
| F9 | No partial-mastery state (HIGH) | **REJECT** | `emerging` already serves that role; "stuck <0.80" folds into V-L-02. |
| F10 | Cross-skill prerequisite silence (MEDIUM-HIGH) | **DEFER** | Zero cross-skill edges by design; shared K-* layer integrates; already an open question. |
| F11 | Overall band masking (MEDIUM) | **REJECT** | Mirrors IELTS (overall = average); per-skill bands are primary. Balance advisory is product-scope. |
| F12 | No knowledge-trap diagnostics (MEDIUM) | **DEFER** | This **is** V-L-03 (Medium, OPEN): audit critical knowledge to `Required`. |
| F13 | K-GRA false prerequisite (MEDIUM) | **REJECT** | K-GRA-010 is root (`requires: —`); K-GRA-001 requires it — a clean one-way edge, no cycle. |
| F14 | Missing K-VOC→K-GRA edges (MEDIUM) | **REJECT** | Edge exists (`K-VOC-030 related_to: K-GRA-010`); soft classification is deliberate (LD-004). |
| F15 | Node-density imbalance (LOW-MEDIUM) | **DEFER** | Already the "Band-5 density" calibration flag; density is evidence-reflective. |
| F16 | No maintenance mode (MEDIUM) | **DEFER** | Performance-graded review already spaces mastered items out; optimization/product concept. |
| F17 | Missing knowledge for Y/N/NG (MEDIUM) | **DEFER** | R-QT-03→R-COMP-04 (inference) is correct receptive chaining; K-GRA-062 as `Recommended` is a calibration item. |
| F18 | Impossible node combinations (MEDIUM) | **DEFER** | Two receptive leaves sharing K-VOC-011 = reinforcement; already the "load/cognitive tuning" flag. |
| F0 | Foundational mastery-learning challenge (CRITICAL) | **REJECT** | Strawman + unaudited citations + limitations already disclosed in the brief. |

**Tally: ACCEPT 1 · DEFER 9 · REJECT 10** (18 findings + the foundational challenge).

---

## ACCEPT — F2 (certification regression semantics)
**Genuine defect.** `BandCertificationState.status` was `not_started → in_progress → certified` with no rule for what happens to a *previously certified* band when the underlying leaves later regress. [transitions.md](../progress/transitions.md) §1 handled leaf regression (`mastered → emerging`) and §6/AM-003 stated the invariant ("no regression"), but the band-level consequence was implicit — leaving the ambiguity the Red Team names ("can a learner display 'Band 6 certified' while performing at Band 4?").

**Minimum correction applied** (not the Red Team's proposed `revoked` state):
- [transitions.md](../progress/transitions.md) §6 — added explicit rule: regression below the band floor reverts `status: certified → in_progress`; re-certification requires a fresh AT-05 portfolio; **no `revoked` state** (certification is held or not); `certification_history` retains the prior attainment as a **point-in-time record** (IELTS Test-Report-Form model).
- [learner-state-schema.md](../progress/learner-state-schema.md) — `BandCertificationState.status` and `certification_history` notes updated to match.

**Why not `revoked`:** IELTS does not revoke a Test Report Form when ability later declines; it records point-in-time attainment. The Blueprint mirrors this: current status tracks current ability; history preserves attainment. A `revoked` enum would mis-model the domain and add a state with no clean re-entry.

**Severity: MEDIUM, not CRITICAL.** No incorrect certification could be *produced* — the gap was an undescribed reversion, and `certification_history` already preserved records. Deflating CRITICAL→MEDIUM.

---

## DEFER — real observations, already tracked or downstream-owned
Each maps to an existing calibration/product backlog item (per the [freeze-report](freeze-report.md) deferred list). New calibration sharpenings from the Red Team are captured here so the calibration phase inherits them.

- **F1 plateau → V-L-02.** *Sharpening for calibration:* implement a time-in-band stagnation signal (extended time at a band with insufficient evidence growth) that triggers an AT-04 diagnostic / escalation — framed as a system-health signal, **not** a mastery criterion (mastery is demonstrated competency, not time). Plateau detection is partly a runtime/product feature; the Blueprint specifies the signal, delivery decides the UX.
- **F4 K-VOC-012 → open question ([vocabulary.md:83](../knowledge/vocabulary.md)).** Per-topic split is calibration. *Correction of the Red Team's framing:* a knowledge object is a content domain, not a delivery pacing unit — "cognitive overload from a vocabulary dump" is a curriculum-pacing concern (handled by adaptive within-band sequencing), independent of object granularity.
- **F7 exam-prep safeguards → V-L-01 + product scope.** LD-005 already makes exam-prep non-certifying and unable to unlock mastery / bypass prereqs / modify certification; diagnostic exposure **is** the reality check. "Impossible-timeline UX" is a delivery decision.
- **F10 cross-skill edges → open question ([prerequisite.md](../learning/prerequisite.md): "model synergistic dependencies as co-developing, not strictly gated").** *Calibration item:* consider `Recommended` (non-gating) cross-skill edges where transfer is well-supported (e.g., writing vocabulary ↔ reading). Not a defect — the architecture is intentionally per-skill (IELTS) with a shared knowledge layer.
- **F12 knowledge trap → V-L-03.** *Calibration action (already OPEN):* audit that critical knowledge is classified `Required` (hard-gated), not merely `Recommended`, so a learner cannot certify a high band with shaky underlying knowledge.
- **F15 node density → "Band-5 density" calibration flag.** Band 5 density is evidence-reflective (the 5→6 transition is the largest jump and the most common target band) — not a defect.
- **F16 maintenance mode → review model + product.** Performance-graded spacing already lengthens intervals for well-mastered items (accuracy stays high). A "minimum maintenance" floor is a calibration detail.
- **F17 R-QT-03 knowledge → skill/knowledge calibration.** R-QT-03 → R-COMP-04 (inference) is correct receptive chaining. *Calibration item:* consider adding K-GRA-062 (modal verbs) / stance markers as `Recommended` for writer's-view question types.
- **F18 C-B6-04 → "load/cognitive tuning" calibration flag.** Two receptive leaves sharing K-VOC-011 is efficient reinforcement within a learning phase, not a simultaneous working-memory task.

---

## REJECT — not defects (evidence)
- **F3 (PG-002 "false independence").** The finding conflates *certification* independence with *structural* independence. PG-002 makes each skill certify independently — this is **faithful to IELTS**, which awards four separate section bands; it was adopted to fix a real defect (V-HIGH-01, the synchronized gate that blocked a Band-7 reader behind Band-5 writing). The "false confidence" scenario (Reading 7 / Writing 4) is how IELTS itself reports. Skills are **not** disjoint: they integrate through the **shared knowledge layer** (K-VOC/K-GRA objects serve multiple skills, per [resolution.md](../knowledge/resolution.md)). The legitimate residual — `Recommended` cross-skill learning edges — is F10 (deferred). CRITICAL is unjustified.
- **F5 (S-GRA-02).** Materially inaccurate. [resolution.md:32](../knowledge/resolution.md) resolves S-GRA-02 → **K-GRA-004, K-GRA-021, K-GRA-005** (3 objects), not 2; and the report's own recommendation ("add K-GRA-005") targets an object **already present**. The W-GRA-03 (6) vs S-GRA-02 (3) asymmetry is defensible: spoken subordination relies on a smaller **automatized** set, while written subordination draws a wider inventory (conditionals, noun clauses) because writing permits monitoring. *(Minor doc-consistency note, not the Red Team's claim: the [grammar.md](../knowledge/grammar.md) coverage table groups S-GRA-02 under "005/06/07" while resolution.md lists 004/021/005 — reconciliation is a cleanup, defer.)*
- **F6 (R-COMP-04 ↔ W-TR-04 "cycle").** Verified at source: R-COMP-04 → `R-COMP-03`; W-TR-04 → `W-TR-01`. **No edge connects the two skills.** The Red Team posited a structural edge that does not exist, then declared it circular. There is no cycle.
- **F8 (spaced-review breakdown).** The review model is **performance-graded** — the next review is scheduled when retrieval accuracy drops toward 80–90% ([review.md](../learning/review.md)). This is accuracy-*triggered*, not calendar-fixed, so it **adapts** to each learner's forgetting rate — including a 1-hour/week learner who forgets faster (accuracy drops sooner → review scheduled sooner). The finding's premise ("intervals may exceed forgetting curves") misunderstands the design; the design exists precisely to prevent that.
- **F9 (no partial-mastery state).** The enum `not_started → practicing → emerging → mastered` already graduates progress; **`emerging` is the partial/in-progress state**. "Stuck below 0.80 confidence" folds into V-L-02 (stuck-learner escalation); "1 of 2 demonstrations" correctly remains `emerging`. A separate `partial_mastery` state would be redundant.
- **F11 (overall band masking).** Overall = average of the four section bands **is how IELTS computes the overall score**; per-skill bands are the Blueprint's operative unit (each advances independently, PG-002). "Masking" is inherent to IELTS reporting, not a Blueprint defect. An optional skill-imbalance advisory is a product/UX enhancement.
- **F13 (K-GRA-010 → K-GRA-001 "circular").** Verified: K-GRA-010 (Word classes) `requires: —` (graph root); K-GRA-001 (Clause structure) `requires: K-GRA-010`. That is a **clean one-way foundation edge**, not a cycle. Teaching word classes before clause structure is standard and defensible.
- **F14 (K-VOC → K-GRA "disjoint silos").** Factually wrong: [vocabulary.md:45](../knowledge/vocabulary.md) shows `K-VOC-030 … related_to: K-GRA-010` — the cross-reference exists. Classifying it `related_to` (soft) rather than `requires` (hard) is a **deliberate** LD-004 choice (minimum hard gates); word formation does not strictly require formal word-class grammar.
- **F0 (foundational mastery-learning challenge).** Rejects rest on three legs, each unsound: **(a) strawman** — the Blueprint is mastery-gated *across band thresholds* and **adaptive within** a band ([LD-001](../learning/decisions.md)); it explicitly models non-linear band transitions ("4→5 ≠ 7→8"), regression, and adult plateaus. It does **not** require mastering one leaf/morpheme before the next (the strict-linear model Krashen's natural-order research actually critiques). **(b) Already disclosed** — the brief marks direct L2 prerequisite evidence `[OPEN]` (§6) and carries mastery-fit cautions (§2, §5); criterion-referenced band certification is the IELTS frame, not a claim about natural acquisition order. **(c) Unaudited citations** — see report-reliability (2); they are not propagated.

---

## Verdict
- **CRITICAL findings surviving as CRITICAL: 0.** F2 → ACCEPT at MEDIUM (corrected); F1 → DEFER (V-L-02); F3 → REJECT; F0 → REJECT.
- **The freeze gate is not breached.** The Blueprint's stated bar — *0 unresolved Critical, 0 unresolved High* ([validation.md](validation.md)) — holds. The one genuine defect (F2) was an undescribed reversion rule, now made explicit; it could not produce an incorrect certification.
- **Net Blueprint change from this review:** one clarification across two progress/ files (F2). All other "findings" are either already-tracked calibration/product items (deferred) or not defects (rejected).
- **Calibration-phase inheritance:** the DEFER sharpenings above (time-in-band signal, `Required`-knowledge audit V-L-03, optional `Recommended` cross-skill / R-QT-03 edges, review-frequency floor, load tuning) are captured for the deferred calibration phase and do not affect structural correctness.
