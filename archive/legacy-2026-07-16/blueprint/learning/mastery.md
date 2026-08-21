# Mastery Model
*Element 2 of the learning model. Implements [philosophy.md](philosophy.md) #3 (mastery-based progression) and [LD-001](decisions.md).*

## Purpose
Define what **mastery / competence** means in this Blueprint and how it is demonstrated.

## Definition
Mastery is **demonstrated competency against the target band's criteria** — not time spent, and not a norm-referenced ranking. A learner has "mastered" a band when they can reliably produce performance meeting that band's competency descriptors (defined in `../bands/`).

## Evidence
- Mastery learning: most learners reach high competence given sufficient time + appropriate instruction (Bloom 1968; Guskey 2010); moderate positive effect (Kulik et al. 1990). See [brief §2](learning-science-evidence-brief.md).
- **Criterion-referenced** assessment is the correct frame for skill mastery (IELTS band descriptors are criterion-referenced standards).

## How mastery is demonstrated
- **Per band:** explicit **exit criteria** (required by [LD-001](decisions.md)) — the observable performance evidencing the band's competencies. Defined in `../bands/`, measured via `../assessment/`.
- **Per micro-skill / knowledge item:** criterion-referenced performance on practice and formative assessment.
- A flat percentage threshold (e.g. 80–90%, per Guskey) is a *fallback heuristic for objective items only*; for productive skills, mastery is judged against descriptors, not a percentage.

## Dependencies
- Exit criteria ← `../bands/` (Band 3–9 descriptors).
- Measurement ← `../assessment/`.
- Gates progression ← [LD-001](decisions.md) and `../progress/`.

## Open questions
- [ ] Confirm whether objective-item mastery uses a fixed threshold (e.g. 80%) or a criterion-referenced standard — to settle in `../assessment/`.
- [ ] Define "reliably" (how many demonstrations, over what interval) — to settle in `../progress/`.
