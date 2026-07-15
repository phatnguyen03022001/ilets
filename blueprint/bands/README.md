# Bands

## Purpose
Define the Band 3–9 descriptors — the spine that `skills/`, `assessment/`, `progress/`, and `curriculum/` implement. Per Founder directive: the band model *implements* the approved [learning model](../learning/README.md); it does not define it.

## Status
**Scaffolded; content pending official-descriptor research.** The descriptors are official public facts (not synthesized). They will be transcribed **verbatim** from official sources with citations and independently verified before use.

## Scope
- **Bands 3–9** (per project objective). Bands 0–2 (non-user / intermittent user) are out of scope; half-bands (e.g. 6.5) are awarded by interpolation, not separate descriptors.
- **Academic first** ([FD-001](../product/foundational-decisions.md)); General Training is added later without restructuring. Shared concepts modeled once; only variant-specific content (GT Writing Task 1 letter; GT Reading passages/conversion) is separated.

## Two-layer model
Each band/skill combines:
1. **Official performance descriptors** (factual, verbatim-cited) — what performance looks like at that band (the IELTS public band descriptors; Listening/Reading raw-score→band conversion).
2. **Learning overlay** (derived per CLAUDE.md Band Rules) — for each band: required knowledge, required skills, **evidence of mastery / exit criteria** (for [LD-001](../learning/decisions.md)), acceptable residual errors, and what belongs to higher bands (no premature introduction).

The overlay is always clearly labeled as inference, not official text.

## Planned structure
- `README.md` (this file) — band model, scoring, half-bands, the two-layer model, Academic/GT handling.
- `writing.md` — Task 1 (Academic) + Task 2 descriptors (TA / CC / LR / GRA), Bands 3–9 + overlay.
- `speaking.md` — Fluency&Coherence / LR / GRA / Pronunciation descriptors, Bands 3–9 + overlay.
- `listening.md` — Academic raw-score→band conversion, Bands 3–9 + overlay.
- `reading.md` — Academic raw-score→band conversion, Bands 3–9 + overlay.
- GT variant-specific content added later under each skill per [FD-001](../product/foundational-decisions.md); convention fixed in `../review/`.

## Dependencies
- **Implements:** the approved [learning model](../learning/README.md) (mastery, progression, assessment, AI role, etc.).
- **Feeds:** exit criteria → `../learning/mastery.md` + `../progress/`; descriptors → `../assessment/` + `../skills/`.

## Open questions (proposed; flag if you disagree)
- [ ] Per-skill doc organization (vs. per-band) — proposed per-skill (maps to `../skills/`, keeps variant content separable).
- [ ] Bands 3–9 scope (exclude 0–2; half-bands via interpolation) — proposed.
- [ ] Exit-criteria location: here per band (referenced by `../progress/`) vs. in `../progress/` — proposed here.
