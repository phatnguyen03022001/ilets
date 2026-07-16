# Review (Governance)

## Purpose
The **Blueprint Governance layer** — defines and executes repository-wide quality assurance. Not per-document review ([GV-001](decisions.md)).

## Responsibility
Verifies structural/terminological/traceability/dependency consistency, canonical-object uniqueness, duplicate + orphan detection, Academic-vs-GT separation, evidence labeling, decision-record consistency, and reference integrity. Outputs a formal **Blueprint Health Report**.

## Status
**Governance + Validation executed (2026-07-16).** [health-report.md](health-report.md) (0 Critical) + [validation.md](validation.md) (independent gate): **0 Critical**, 1 **High** (V-HIGH-01 band-progression — Founder decision needed before freeze). Glossary finalized. Freeze pending V-HIGH-01.

## Structure
- [governance.md](governance.md) — the 12-check QA catalog.
- [health-report.md](health-report.md) — executed Health Report.
- [validation.md](validation.md) — independent validation report (5 suites).
- [decisions.md](decisions.md) — GV-001 governance; GV-002 Academic/GT variant convention.

## Dependencies
- **Consumes:** the entire Blueprint (all canonical layers + decisions + glossary).
- **Feeds:** the final validation + freeze gate (zero unresolved Critical issues required).
