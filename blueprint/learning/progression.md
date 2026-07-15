# Progression Model
*Element 4. Implements [philosophy.md](philosophy.md) #3 and decisions [LD-001](decisions.md), [LD-005](decisions.md).*

## Purpose
Define how learners advance through the Blueprint.

## Two independent concepts ([LD-005](decisions.md))
The Blueprint separates **certified learning advancement** from **exam-readiness practice**.

### Learning Progression (mastery-gated)
- **Between bands: mastery-gated.** No certification of band completion/advancement until the band's exit criteria are met (see [mastery.md](mastery.md)).
- **Within a band: adaptive.** The system personalizes sequence, practice selection, review scheduling, remediation, difficulty, and pacing.
- **Invariants:** adaptive behavior never bypasses required objectives or mastery criteria; **same outcomes, different paths**; every band has explicit exit criteria.

### Exam Preparation (not progression)
- May expose learners to **higher-band tasks before mastery** when a fixed test date approaches — for **diagnosis, familiarization, and exam readiness only**.
- **Must never:** unlock higher-band mastery; satisfy progression requirements; bypass prerequisites; or modify certification of current-band completion.
- Feeds diagnostic/summative assessment ([assessment.md](assessment.md)); does **not** affect [mastery.md](mastery.md) certification.

## Evidence
- Mastery-gated progression: strongest fit for skill domains (Bloom 1968; Kulik et al. 1990). See [brief §5](learning-science-evidence-brief.md).
- Adaptive sequencing: effective when theory-driven (VanLehn 2011); variable otherwise — hence constrained by the invariants above.

## Dependencies
- Exit criteria ← `../bands/`; adaptive rules ← `../curriculum/` + `../progress/`; mastery signals ← `../assessment/` + `../progress/`; exam-prep instruments ← `../practice/` + `../assessment/`.

## Open questions
- [ ] Define the trigger/rules for entering "Exam Preparation" mode (e.g., proximity to a learner-declared test date) — during `../progress/`.
