# Decisions

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Maintained through `TASK-0003` revision 1. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Canonical target snapshot: `phatnguyen03022001/ilets@2f54b7c6003399ccf7abff8cda2277d68d0048e8`.

## Documentation-model adoption

**Draft outcome:** this migration uses `phatnguyen03022001/agent-documents@acb3f02616e700190586681306a86905792e4c07`, the accepted unreleased V1 candidate after the relation-model correction. This target pin supersedes the earlier pilot pin only for continuing the IELTS migration; it does not globally stabilize or release `agent-documents`.

The corrected feature relation state distinguishes typed `refs`, genuine demonstrated `na`, and `unresolved_ref` to an existing `OPEN` `DESIGN` `UNK-*`. Migration incompleteness is not N/A. An unresolved relation blocks documentation closure without creating a substitute neighboring-domain identity.

## Catalog migration boundary

`docs/catalog/project.json` now exists only as a migration artifact with **authority NONE**. Its pathname does not grant canonical target authority. The catalog owns migration identities/typed relations; `docs/BEHAVIOR.md` owns migrated behavior explanation.

This slice may assign directly supported `ACT-*`, `ROL-*`, `FTR-*`, `ACC-*`, `FLW-*`, and `UNK-*` identities. It does not prematurely materialize `SYS-*`, `DAT-*`, `IFC-*`, `EXT-*`, or `CAP-*`.

## Engineering claim and evidence semantics adoption

IELTS continues to pin `phatnguyen03022001/agent-standards@3f4950f280a3a35fee81471d4b83715fa72cf9ee` for generic engineering claim/evidence semantics. This migration performs no standards assessment and claims no requirement `PASS`, maturity level, assurance `N/A`, exception, implementation readiness, promotion readiness, or release readiness. Coverage depth is documentation treatment, not engineering maturity.

## Ownership and cutover boundary

Actual IELTS truth remains with the existing canonical legacy documents. `docs/PRODUCT.md`, this file, `docs/BEHAVIOR.md`, and `docs/catalog/project.json` remain authority-NONE migration artifacts. No canonical cutover, `DOCS_READY`, documentation lock, generated PROGRAM, promotion, or release is declared.

## Migration unknowns

`TASK-0003` records not-yet-migrated material feature relations as catalog-owned `OPEN` `DESIGN` `UNK-*` records referenced through `unresolved_ref`. A later authorized migration may replace each relation only with correctly typed `refs` or genuine demonstrated `na` after the owning domain is resolved. Until then, documentation closure remains blocked.
