# Review (Governance)

## Purpose
The **Blueprint Governance layer** — defines and executes repository-wide quality assurance. Not per-document review ([GV-001](decisions.md)).

## Responsibility
Verifies structural/terminological/traceability/dependency consistency, canonical-object uniqueness, duplicate + orphan detection, Academic-vs-GT separation, evidence labeling, decision-record consistency, and reference integrity. Outputs a formal **Blueprint Health Report**.

## Status
**Governance defined (2026-07-16).** QA catalog ([governance.md](governance.md)) + variant convention ([GV-002](decisions.md)) in place. Next: execute the checks → [health-report.md](health-report.md), then the final validation phase.

## Structure
- [governance.md](governance.md) — the 12-check QA catalog.
- [health-report.md](health-report.md) — executed Health Report (findings + severity + resolutions) — pending.
- [decisions.md](decisions.md) — GV-001 governance; GV-002 Academic/GT variant convention.

## Dependencies
- **Consumes:** the entire Blueprint (all canonical layers + decisions + glossary).
- **Feeds:** the final validation + freeze gate (zero unresolved Critical issues required).
