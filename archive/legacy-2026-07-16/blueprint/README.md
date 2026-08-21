# IELTS Blueprint

The **Single Source of Truth (SSOT)** for every learning decision in the IELTS learning application.

This Blueprint defines *what learners must know, what they must practice, how they are assessed, and how they progress from Band 3 to Band 9*. It does **not** define implementation (architecture, APIs, databases, infrastructure). Implementation begins only after the Blueprint is reviewed and frozen. See `/CLAUDE.md` and `/OBJECTIVE.md`.

> **Status: FROZEN (2026-07-16).** The Learning Blueprint (Blueprint Scope) is the project's Single Source of Truth for the learning domain — all 11 categories complete, independently validated (**0 Critical / 0 High**), glossary finalized. See [review/freeze-report.md](review/freeze-report.md). **Product & Implementation scopes remain open** and consume this frozen Blueprint ([FD-006](product/foundational-decisions.md)).

---

## Foundational constraints (Founder decisions, 2026-07-15)

Three constraints shape the entire structure. Full rationale: [`product/foundational-decisions.md`](product/foundational-decisions.md).

1. **Academic first; General Training added later without restructuring.** v1 completes Academic. The structure must let GT be added without modifying existing artifacts. **Shared concepts are modeled once; only variant-specific content (e.g. Reading passages, Writing Task 1) is separated.**
2. **L1-agnostic canonical Blueprint; localization is a separate layer.** All Blueprint documents are written in English and remain independent of any native language. Localization (Vietnamese first) maps *onto* the canonical Blueprint without changing learning requirements.
3. **Four official skills are first-class; Vocabulary and Grammar are knowledge.** Listening, Reading, Writing, Speaking live under `skills/`; Vocabulary and Grammar live under `knowledge/` as enabling layers that feed all four skills.

---

## Top-level structure

| Category | Purpose |
|---|---|
| `product/` | Product vision, scope, principles, and foundational decisions. |
| `learning/` | Learning model, philosophy, and pedagogy (how learners acquire skill). |
| `curriculum/` | Sequencing and structure of learning across skills and bands. |
| `bands/` | Band 3→9 descriptors: what each band requires, evidence of mastery, acceptable residual errors. |
| `skills/` | The four official skills (Listening, Reading, Writing, Speaking). |
| `knowledge/` | Enabling knowledge layers (Vocabulary, Grammar, others TBD) that feed all skills. |
| `practice/` | Practice strategy, item types, and practice design per skill/band. |
| `assessment/` | Assessment model, criteria, scoring, and authenticity vs. official IELTS. |
| `progress/` | Progression rules, mastery signals, and movement between bands. |
| `review/` | Consistency review, gap analysis, and change protocol. |
| `glossary/` | Canonical terminology — the single source for every defined term. |
| `localization/` | L1-specific guidance (VN first) mapped onto the canonical Blueprint. *(New top-level — see justification below.)* |

### Why `localization/` is a new top-level category

`CLAUDE.md` requires justification for any new top-level category.

- **Why necessary:** Decision [FD-002](product/foundational-decisions.md) mandates an L1-specific guidance layer (Vietnamese first, others later).
- **Why it cannot fit an existing category:** Localization must *not* alter the canonical learning, assessment, curriculum, or knowledge content (the L1-agnostic invariant). Embedding localization inside `skills/`, `knowledge/`, etc. would either pollute the canonical source or duplicate it. A separate top-level keeps the canonical layer pure and the localization layer clearly derivative.

### Academic vs. General Training modeling

Per [FD-001](product/foundational-decisions.md), shared content lives at the category/skill level; variant-specific content lives in a clearly separated sub-location. The exact convention (e.g. `skills/reading/academic/` vs. `skills/reading/general-training/`, with shared content at `skills/reading/`) will be fixed when the first variant-specific section is authored, and recorded in `review/`.

---

## Conventions

- **Language:** English is canonical. All other languages are localization overlays.
- **Evidence policy:** Prefer official IELTS → Cambridge → British Council → IDP → peer-reviewed research → reputable references. Always label Fact / Evidence / Inference / Assumption / Recommendation. Never present assumptions as facts.
- **No duplication:** A concept is defined once; other documents link to it.
- **Decision records:** Every non-trivial decision is recorded with rationale and indexed from [`product/foundational-decisions.md`](product/foundational-decisions.md).
- **Open questions are first-class:** every section lists its open questions; unresolved items are never silently assumed.
- **Repository lifecycle:** Blueprint work lives on the `blueprint` branch; `main` holds only the latest approved, frozen Blueprint. Implementation starts only after the frozen Blueprint is merged to `main`. See [FD-005](product/foundational-decisions.md).
