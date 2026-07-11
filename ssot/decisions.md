# Architecture Decision Records

## ADR Index

### ADR-0037: Authentication + Stack Decisions
**Date**: 2026-07-11  
**Status**: Resolved

### ADR-0036: Repository OS Restructuring
**Date**: 2026-07-11  
**Status**: Resolved

### ADR-0035: First Live Round-Trip Implementation
**Date**: 2026-07-11  
**Status**: Resolved

### ADR-0034: Band Score Calculation Utilities
**Date**: 2026-07-11  
**Status**: Resolved

---

## ADR-0037: Authentication + Stack Decisions

### Context
The platform needed to establish core authentication and technology stack choices for the IELTS learning platform.

### Decision
Choose Next.js 16 with App Router, PostgreSQL with Drizzle ORM, and a custom authentication system.

### Rationale
1. **Why now?**: The platform needs a stable foundation before implementing features.
2. **Why not later?**: Delaying architecture decisions creates technical debt.
3. **Why not another approach?**: Next.js 16 provides stability, Drizzle offers TypeScript-first development, and custom auth avoids vendor lock-in.

### Observable Signal
- TypeCheck passes without errors
- Build succeeds with zero warnings
- Lint passes with minimal warnings
- All tests pass with 100% coverage for core utilities

### Migration Path
- Migrated from legacy src/ structure to app/
- Fixed TypeScript path aliases
- Updated CI pipeline to workspaces/web structure
- Ensured all dependencies are compatible

### Rollback Path
- Revert to backup state before restructuring
- Restore original file locations
- Reset CI pipeline configuration

### Resolution Metrics
- ✅ Typecheck: PASSED
- ✅ Build: PASSED
- ✅ Lint: PASSED (1 info issue)
- ✅ Tests: PASSED (25/25 tests)

---

## ADR-0036: Repository OS Restructuring

### Context
The repository needed a complete restructuring to implement a Documentation-First architecture with clear separation of concerns.

### Decision
Implement Repository OS structure with ssot/, runtime/, and workspaces/ directories.

### Rationale
1. **Why now?**: Current structure was inconsistent and lacked clear documentation practices.
2. **Why not later?**: Procrastination leads to accrued technical debt.
3. **Why not another approach?**: Repository OS provides clear boundaries and documentation-first principles.

### Observable Signal
- Git status shows clean slate with new structure
- All verification steps pass (typecheck, lint, build, test)
- No old files remaining in repository

### Migration Path
- Deleted all old code files and directories
- Created new ssot/ directory structure
- Moved Next.js app to workspaces/web/
- Updated CI workflow paths

### Rollback Path
- Restore from git backup
- Revert to previous directory structure

### Resolution Metrics
- ✅ Clean repository status
- ✅ All verification steps pass
- ✅ SSOT documentation created

---

## ADR-0035: First Live Round-Trip Implementation

### Context
First implementation of the platform with working end-to-end flow from submission to evaluation.

### Decision
Implement writing evaluation system with AI feedback and band score calculation.

### Rationale
1. **Why now?**: Need to validate core platform functionality.
2. **Why not later?**: Early validation prevents wasted development effort.
3. **Why not another approach?**: Direct implementation provides fastest path to working prototype.

### Observable Signal
- All 25 tests pass
- Band score calculation works correctly
- JSON parsing utilities functional
- No TypeScript compilation errors

### Migration Path
- Implemented from scratch based on requirements
- Established development workflow
- Created test suite for validation

### Rollback Path
- Remove implementation and revert to previous state

### Resolution Metrics
- ✅ 25/25 tests pass
- ✅ Band score calculations verified
- ✅ JSON parsing validated

---

## ADR-0034: Band Score Calculation Utilities

### Context
Need accurate band score calculations for IELTS Writing evaluation with proper rounding rules.

### Decision
Create utility functions for band score calculation and JSON response parsing.

### Rationale
1. **Why now?**: Core functionality required for evaluation system.
2. **Why not later?**: Essential for platform MVP.
3. **Why not another approach?**: Direct implementation provides clear, maintainable code.

### Observable Signal
- Tests pass with edge cases handled
- Rounding logic verified against IELTS standards
- JSON parsing error-free

### Migration Path
- Created utility functions
- Implemented comprehensive tests
- Integrated with evaluation system

### Rollback Path
- Remove utility functions and tests
- Revert to previous state

### Resolution Metrics
- ✅ All test cases pass
- ✅ Edge cases handled
- ✅ Integration working

---

## Decision Format Template

```markdown
# [ADR-XXXX] [Decision Title]

## Context
[What problem led to this decision?]

## Decision
[What we chose to do]

## Rationale
1. Why now?
2. Why not later?
3. Why not another approach?

## Observable Signal
[What metric/condition triggered evolution?]

## Migration Path
[How we evolved from previous state?]

## Rollback Path
[How to revert if needed?]

## Status
[Current state of decision]
```

## New Decision Format
All future decisions must follow this template and be recorded in this file.