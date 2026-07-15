# Bands

## Purpose
Define the Band 3–9 descriptors — the spine that `skills/`, `assessment/`, `progress/`, and `curriculum/` implement. Per Founder directive: the band model *implements* the approved [learning model](../learning/README.md); it does not define it.

## Status
**In progress (2026-07-16).** `writing.md` complete — descriptors verified + Bands 3–9 overlay + hierarchical exit criteria ([BD-002](decisions.md)). `speaking.md` next, then the receptive skills (`listening.md`, `reading.md`) with three-category evidence labeling ([BD-003](decisions.md)).

## Scope
- **Bands 3–9** are the learning focus (curriculum, learning paths, detailed requirements, exit criteria). **Bands 0–2 are documented as boundary definitions only** — to complete the assessment model and support diagnostics; no learning content for 0–2 (see [BD-001](decisions.md)). **Band 3 is the program entry point.** Half-bands (e.g. 6.5) are awarded by interpolation, not separate descriptors.
- **Academic first** ([FD-001](../product/foundational-decisions.md)); General Training is added later without restructuring. Shared concepts modeled once; only variant-specific content (GT Writing Task 1 letter; GT Reading passages/conversion) is separated.

## Two-layer model
Each band/skill combines:
1. **Official performance descriptors** (factual, verbatim-cited) — what performance looks like at that band (the IELTS public band descriptors; Listening/Reading raw-score→band conversion).
2. **Learning overlay** (derived per CLAUDE.md Band Rules; **Bands 3–9 only**) — for each band: required knowledge, required skills, **evidence of mastery / exit criteria** (for [LD-001](../learning/decisions.md)), acceptable residual errors, and what belongs to higher bands (no premature introduction). Bands 0–2 carry official descriptors only (boundary/diagnostic), no overlay (see [BD-001](decisions.md)).

The overlay is always clearly labeled as inference, not official text.

## Planned structure
- `README.md` (this file) — band model, scoring, half-bands, the two-layer model, Academic/GT handling.
- `writing.md` ✅ — Task 1 (Academic) + Task 2 descriptors (TA/CC/LR/GRA), bands 0–9 verified + Bands 3–9 overlay + hierarchical exit criteria ([BD-002](decisions.md)).
- `speaking.md` ✅ — FC / LR / GRA / Pronunciation descriptors (bands 0–9 verified) + Bands 3–9 overlay + Speaking-Skill exit criteria ([BD-002](decisions.md)).
- `listening.md` — Academic raw-score→band conversion, Bands 3–9 + overlay.
- `reading.md` — Academic raw-score→band conversion, Bands 3–9 + overlay.
- GT variant-specific content added later under each skill per [FD-001](../product/foundational-decisions.md); convention fixed in `../review/`.

## Dependencies
- **Implements:** the approved [learning model](../learning/README.md) (mastery, progression, assessment, AI role, etc.).
- **Feeds:** exit criteria → `../learning/mastery.md` + `../progress/`; descriptors → `../assessment/` + `../skills/`.

## Structure decisions
- [BD-001](decisions.md): per-skill docs; Bands 3–9 learning focus with 0–2 as boundary-only; exit criteria in `bands/`.
- [BD-002](decisions.md): hierarchical exit criteria — Task-level → Skill-level (non-duplicative) → overall band progression (in `../progress/`).
- [BD-003](decisions.md): three-category evidence labeling (Official Evidence / Evidence-Based Interpretation / Blueprint Inference) for receptive skills.

## Open questions
- [ ] (none structural) — content open questions will be tracked per band doc once descriptors are verified.
