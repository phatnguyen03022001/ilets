# Prerequisite Model
*Element 3. Implements [philosophy.md](philosophy.md) #4 and [LD-004](decisions.md).*

## Purpose
Define how prerequisites and learning progressions are structured and classified.

## Model ([LD-004](decisions.md))
Dependencies are classified into three tiers; only **Required** items are hard gates:
- **Required (hard)** — insufficient foundation would make subsequent learning ineffective. Each carries an **evidence-based rationale**.
- **Recommended (soft)** — beneficial, not blocking.
- **Independent** — no prerequisite.

**Principle:** minimum number of hard gates necessary; prefer adaptive recommendations over mandatory gates.

## Evidence
- Direct evidence on optimal L2 prerequisite structures is `[OPEN]` (brief §6).
- Indirect support: cognitive-load **element interactivity** (Sweller et al. 2019) — content with high element interactivity plausibly needs foundational prerequisites to avoid overload.
- Hierarchical prerequisites are validated mainly in mathematics; L2 inferences (e.g. core vocabulary/grammar before complex production) are reasonable but **not empirically established** and must be justified per item.

## How a dependency is classified
A dependency is **Required** only when all hold:
1. Insufficient foundation demonstrably makes the dependent learning *ineffective* (not merely harder).
2. An evidence- or theory-based rationale can be written for why it is mandatory.
3. No reasonable adaptive path can compensate.

Otherwise it is **Recommended** or **Independent**.

## Dependencies
- The actual prerequisite *graph* per skill is defined in `../curriculum/` and `../knowledge/`, applying this classification.

## Open questions
- [ ] Define the canonical set of **Required** foundational prerequisites (candidate: core vocabulary, core grammar) — during `../knowledge/` + `../curriculum/`.
- [ ] Handle circular / synergistic dependencies (e.g. listening ↔ speaking) — model as co-developing, not strictly gated.
