# Learning Decisions

Founder decisions that define the learning model. Each resolves a fork surfaced from the [evidence brief](learning-science-evidence-brief.md). Indexed from [product/foundational-decisions.md](../product/foundational-decisions.md).

---

## LD-001 — Progression model: mastery-gated across bands, adaptive within
**Date:** 2026-07-15 · **Status:** Decided (Founder) · **Resolves:** Progression fork

**Decision.** Progression **between bands is mastery-gated**: a learner does not advance to the next band until the current band's required competencies are demonstrated. **Within a band, learning is adaptive** — the system may personalize learning sequence, practice selection, review scheduling, remediation, difficulty, and pacing.

**Constraints (binding).**
- Adaptive behavior must **never bypass required learning objectives or mastery criteria**.
- **Every band defines explicit exit criteria.**
- **Every learner reaches the same learning outcomes**, even via different paths.
- Optimize for **both** educational effectiveness **and** practical readiness for real IELTS test dates.

**Evidence.** Mastery learning (Bloom 1968; Kulik et al. 1990); adaptive effectiveness when theory-driven (VanLehn 2011). See [brief §5](learning-science-evidence-brief.md).

---

## LD-002 — AI is the primary feedback provider; human review optional
**Date:** 2026-07-15 · **Status:** Decided (Founder) · **Resolves:** AI/human feedback fork

**Decision.** AI is the **primary feedback provider** throughout. The Blueprint must **not require human feedback** as a dependency for progression. AI provides immediate, actionable, evidence-based feedback for **all skills**.

**Constraints (binding).**
- For complex productive tasks (Writing, Speaking), explicitly **acknowledge current AI limitations** and provide **calibrated confidence levels** where appropriate.
- Human review is an **optional enhancement**, not a required component. The core Blueprint is **fully functional without human intervention**; optional expert review may exist in future product editions.

**Evidence.** ITS ≈ human tutoring (VanLehn 2011); large positive effects of AI in L2 (Lee & Lee 2024; Wu et al. 2024; Saarela et al. 2026); limits on nuance, motivation, hallucination — see [brief §8](learning-science-evidence-brief.md).

---

## LD-003 — Feedback timing: stage-dependent (acquisition / consolidation / transfer)
**Date:** 2026-07-15 · **Status:** Decided (Founder) · **Resolves:** Principle #7 conflict A

**Decision.** Feedback timing is determined by the learner's stage, learning objective, and task type — not by a blanket "immediate" rule.
- **Acquisition (typically Bands 3–5):** immediate, specific, actionable feedback to prevent reinforcing incorrect understanding.
- **Consolidation / transfer (typically Bands 6–9):** increasingly retrieval-first and appropriately delayed feedback to strengthen long-term retention, transfer, and exam independence.

**Principle.** The objective is the **most effective feedback for the stage**, not immediate feedback per se. The Blueprint distinguishes **acquisition / consolidation / transfer** when defining feedback strategies.

**Evidence.** Hattie & Timperley (2007); feedback-timing evidence — see [brief §4](learning-science-evidence-brief.md), `[CONTESTED]`.

---

## LD-004 — Prerequisite model: Required / Recommended / Independent (minimum hard gates)
**Date:** 2026-07-15 · **Status:** Decided (Founder) · **Resolves:** Principle #4 conflict B

**Decision.** Enforce only a **small set of high-confidence foundational prerequisites** essential for successful learning. All other dependencies are **adaptive recommendations**, not mandatory gates.

**Classification.**
- **Required (hard prerequisite)** — insufficient foundation would make subsequent learning ineffective. Each must include an **evidence-based rationale** for why it is mandatory.
- **Recommended (soft prerequisite)** — beneficial but not blocking.
- **Independent** — no prerequisite.

**Principle.** Prefer the **minimum number of hard gates** necessary to maintain learning quality, avoiding unnecessary bottlenecks.

**Evidence.** Direct L2 prerequisite evidence is `[OPEN]` (brief §6); CLT element interactivity (Sweller et al. 2019) informs what can plausibly be classified "Required".

**Implications.** The actual prerequisite *graph* (which specific items are Required/Recommended/Independent) is defined per skill in `../curriculum/` and `../knowledge/`, applying this classification.

---

## LD-005 — Learning Progression vs. Exam Preparation (two independent concepts)
**Date:** 2026-07-16 · **Status:** Decided (Founder) · **Resolves:** Progression fork #4 (test-date reconciliation)

**Decision.** The Blueprint explicitly models two independent concepts:
- **Learning Progression** — mastery-gated. A learner is **not** certified as having completed or progressed beyond a band until **all required mastery criteria** are satisfied.
- **Exam Preparation** — may expose learners to **higher-band tasks before mastery** when preparing for a fixed test date. This exposure is for **diagnosis, familiarization, and exam readiness only**.

**Constraints (binding) — Exam Preparation must never:**
- unlock higher-band mastery;
- satisfy progression requirements;
- bypass prerequisites;
- modify certification of current-band completion.

**Rationale.** Preserves educational integrity (mastery-gated certification) while remaining practical for real IELTS candidates facing fixed test dates.

**Implications.** Implemented in [progression.md](progression.md); Exam-Preparation instruments feed [assessment.md](assessment.md) (summative/diagnostic) but do not affect [mastery.md](mastery.md) certification.
