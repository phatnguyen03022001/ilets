# AI's Role in Learning
*Element 8. Implements [philosophy.md](philosophy.md) #8 (AI augments, never replaces thinking) and [LD-002](decisions.md).*

## Purpose
Define AI's role and its boundaries in the learning model.

## Decision ([LD-002](decisions.md))
AI is the **primary feedback provider** for all skills; progression **never depends on human feedback**. For Writing/Speaking, AI feedback carries **calibrated confidence** and explicitly acknowledges current limitations. Human review is **optional**, not required.

## Evidence ([brief §8](learning-science-evidence-brief.md))
- ITS effectiveness ≈ human tutoring, both ≫ classroom (VanLehn 2011) — AI feedback can be high-quality.
- AI in L2 shows **large positive effects** (Lee & Lee 2024; Wu et al. 2024; Saarela et al. 2026); effectiveness varies by application and algorithm quality.

## Boundaries (Principle #8)
- AI **augments**, never replaces, the learner's thinking. Practice is designed so the learner does the cognitive work (retrieval, production), not the AI.
- **Risks to manage:** over-reliance eroding metacognition; hallucination in AI-generated content (requires verification); lower motivation vs. human tutors; bias/fairness.
- High-volume practice + objective scoring (Reading/Listening): strong AI fit. Complex productive feedback (Writing/Speaking): AI with calibrated confidence + optional human verification.

## Dependencies
- Feedback design ← [assessment.md](assessment.md); confidence-calibration + content-verification spec ← `../assessment/`.

## Open questions
- [ ] Define the confidence-calibration scheme and the AI-generated-content verification policy — during `../assessment/`.
