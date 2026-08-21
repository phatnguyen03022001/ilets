# Review (Governance)

## Purpose
The **Blueprint Governance layer** — defines and executes repository-wide quality assurance. Not per-document review ([GV-001](decisions.md)).

## Responsibility
Verifies structural/terminological/traceability/dependency consistency, canonical-object uniqueness, duplicate + orphan detection, Academic-vs-GT separation, evidence labeling, decision-record consistency, and reference integrity. Outputs a formal **Blueprint Health Report**.

## Status
**Governance + Validation executed; Blueprint FROZEN (2026-07-16).** [health-report.md](health-report.md) (0 Critical) + [validation.md](validation.md) (independent gate): **0 Critical / 0 High** (V-HIGH-01 resolved via [PG-002](../progress/decisions.md)). Independent Red-Team review triaged in [red-team-triage.md](red-team-triage.md). Glossary finalized. See [freeze-report.md](freeze-report.md).

## Structure
- [governance.md](governance.md) — the 12-check QA catalog.
- [health-report.md](health-report.md) — executed Health Report.
- [validation.md](validation.md) — independent validation report (5 suites).
- [red-team-report.md](red-team-report.md) — raw independent Red-Team adversarial review (input record).
- [red-team-triage.md](red-team-triage.md) — architectural-owner triage of the Red-Team review.
- [challenge-test-results.md](challenge-test-results.md) — raw challenge-test results (input record).
- [freeze-report.md](freeze-report.md) — Blueprint freeze record (scope + gate evidence).
- [close-out.md](close-out.md) — repository close-out report.
- [decisions.md](decisions.md) — GV-001 governance; GV-002 Academic/GT variant convention.

## Dependencies
- **Consumes:** the entire Blueprint (all canonical layers + decisions + glossary).
- **Feeds:** the final validation + freeze gate (zero unresolved Critical issues required).
