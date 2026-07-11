# Constitution

## Laws of the Repository

### 1. Substrate Law
> Substrate must not leak into the surface. No framework/SSOT/naming-law/abstraction before ≥2 real domains conflict over it.

### 2. Artifact Law
> Every artifact must reduce future cost. A file/folder/doc that produces no decision → DELETE.

### 3. Progress Law
> Research ≠ progress. Decision = progress. Every research effort must end in 1–3 decisions.

### 4. Dependency Law
> No new layer/tool/dependency unless it solves a REAL pain hit during building.

### 5. Decision Format Law
> Every architecture decision must answer the 6 questions:
> 1. Why now?
> 2. Why not later?
> 3. Why not another approach?
> 4. What observable signal will trigger evolution?
> 5. Migration path?
> 6. Rollback path?
> → If question 4 can't be answered = defer.

### 6. Density Law
> Decision Density is the measure — not line count, file count, or folder count.

## Decision Format Template

```markdown
# [ADR-XXXX] Decision Title

## Context
[What problem led to this decision?]

## Decision
[What we chose to do]

## Rationale
1. Why now?
2. Why not later?
3. Why not another approach?

## Observable Signal
[What metric/condition will trigger evolution?]

## Migration Path
[How to evolve from current state?]

## Rollback Path
[How to revert if needed?]
```

## Enforcement
- All architectural changes must use this format
- No decisions without observable signals
- Regular pruning of unused artifacts
- Documentation must serve clear purposes