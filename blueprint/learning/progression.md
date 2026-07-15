# Progression Model
*Element 4. Implements [philosophy.md](philosophy.md) #3 and [LD-001](decisions.md).*

## Purpose
Define how learners advance through the Blueprint.

## Model ([LD-001](decisions.md))
- **Between bands: mastery-gated.** No advance to the next band until the current band's exit criteria are met (see [mastery.md](mastery.md)).
- **Within a band: adaptive.** The system personalizes learning sequence, practice selection, review scheduling, remediation, difficulty, and pacing.

**Invariants (binding).**
- Adaptive behavior **never bypasses required learning objectives or mastery criteria**.
- **Same outcomes, different paths** — every learner reaches the same learning outcomes.
- Every band has **explicit exit criteria**.
- Optimized for **both** educational effectiveness and real-test-date readiness.

## Evidence
- Mastery-gated progression: strongest fit for skill domains (Bloom 1968; Kulik et al. 1990). See [brief §5](learning-science-evidence-brief.md).
- Adaptive sequencing: effective when theory-driven (VanLehn 2011); variable otherwise — hence constrained by the invariants above.

## Cautions
- Over-gating demotivates.
- Test-date pressure may require letting learners *practice/expose* to higher-band content for exam simulation — reconcile with "never bypass mastery" (open question).
- Adaptive systems need transparency (the learner understands why content adapts).

## Dependencies
- Exit criteria ← `../bands/`; adaptive rules ← `../curriculum/` + `../progress/`; mastery signals ← `../assessment/` + `../progress/`.

## Open questions
- [ ] Reconcile "never bypass mastery" with test-date readiness — e.g. allow *practice/exposure* to higher-band content for exam simulation without *certifying* band progression. **Decision needed.**
