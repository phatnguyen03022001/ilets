---
name: <short-kebab-case-title>
description: Gap record — decision deferred or not yet made
metadata:
  type: decision
  status: open
  date: YYYY-MM-DD
  decisions:
    - D####: <decision-id> (e.g., D0002)
  related:
    - [[DR####]]
    - [[DR####]]
---

## Context

[What's the gap? Why hasn't a decision been made?]

**Why this is a gap:**
- Reason 1
- Reason 2

**Why we can't decide now:**
- Constraint 1
- Constraint 2

## Decision Options

**Option A:**
- Description
- Pros
- Cons
- Risk level: low/medium/high

**Option B:**
- Description
- Pros
- Cons
- Risk level: low/medium/high

**Option C:**
- Description
- Pros
- Cons
- Risk level: low/medium/high

**Option D:**
- Description
- Pros
- Cons
- Risk level: low/medium/high

## Criteria for Decision

What signal/constraint will trigger a decision?

1. Signal: [observable signal]
   - Trigger: [when signal ≥ X, make decision]

2. Constraint: [constraint that must be satisfied]
   - Trigger: [when constraint is met, make decision]

3. Dependency: [external dependency]
   - Trigger: [when dependency is resolved, make decision]

**Decision deadline:** YYYY-MM-DD or "after [milestone]"

## Open Questions

- Question 1: [unanswered question]
  - Status: pending/answered/no

- Question 2: [unanswered question]
  - Status: pending/answered/no

## Next Steps

1. [action 1]
2. [action 2]
3. [action 3]

**Owner:** @username

## References

- Related decisions: [[DR####]]
- Source material: [URL or doc]
- Stakeholders: [@user1, @user2]

**Why:** This template records deliberate postponement with:
1. Clear criteria for when to decide
2. Decision options with risk levels
3. Open questions tracked
4. Next steps with owner
5. Deadline/milestone trigger

**How to use:**
- Create a new file per gap (not per file/component)
- Mark status: open
- Keep options structured with pros/cons/risk
- Define 1-3 clear criteria (signal/constraint/dependency)
- Link related decisions with [[DR####]]
