---
name: <short-kebab-case-title>
description: Architecture decision record with observable signal
metadata:
  type: decision
  status: draft|proposed|implemented|deprecated
  date: YYYY-MM-DD
  decisions:
    - D####: <decision-id> (e.g., D0001)
  related:
    - [[DR####]]
    - [[DR####]]
---

## Context

[Why are we making this decision? What problem does it solve? What constraints are we working under?]

## Decision

[The specific decision itself — what we're choosing to do and why.]

**Key choices:**
- Choice 1
- Choice 2
- Choice 3

**Why this choice?**
[Argument in favor — explain the reasoning.]

## Consequences

### Positive (benefits)
- Benefit 1
- Benefit 2

### Negative (drawbacks)
- Drawback 1
- Drawback 2

### Neutral (unaffected)
- Impact 1

## Observable Signal (migration gate)

**Signal:** What will we watch to know if this was right?

**Trigger:**
- If signal → A: [next step]
- If signal → B: [next step]
- If signal → C: [rollback]

**Rationale:** Why this specific signal and trigger?

## Rollback Plan

How to undo this decision if it fails:

1. Step 1: [action]
2. Step 2: [action]

**Cost of rollback:** [estimate: low/medium/high]

## References

- Related decisions: [[DR####]]
- Source material: [URL or doc]
- Stakeholders: [@user1, @user2]

**Why:** This template enforces:
1. Decision density (not line count)
2. Observable signal gate (prevents over-engineering)
3. Migration path + rollback path (single ADR rule)
4. Related-decision linking (propagation control)

**How to use:**
- Create a new file per decision (not per file/component)
- Fill all 6 required sections (why now/why not later/why not another/approach/signal/rollback)
- Link related decisions with [[DR####]] syntax
- Mark status: draft → proposed → implemented → deprecated
