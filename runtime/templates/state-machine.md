---
name: <short-kebab-case-title>
description: State machine with observable signal
metadata:
  type: decision
  state-machine:
    - entity: <entity-name>
      states: [state1, state2, state3]
      transitions: [transition1, transition2, transition3]
      initial: state1
  decisions:
    - D####: <decision-id> (e.g., D0003)
  related:
    - [[DR####]]
    - [[DR####]]
---

## Context

[What entity/state-machine are we modeling? What problem does it solve?]

**Why a state machine?**
- Reason 1
- Reason 2

**Why this state machine (not another)?**
- Reason 1
- Reason 2

## States

**State 1: <state-name>**
- Description
- Guard conditions: [when this state is active]
- Observable signals: [what we watch]

**State 2: <state-name>**
- Description
- Guard conditions: [when this state is active]
- Observable signals: [what we watch]

**State 3: <state-name>**
- Description
- Guard conditions: [when this state is active]
- Observable signals: [what we watch]

## Transitions

**Transition 1: <state-from> → <state-to>**
- Trigger: [event, signal, or condition]
- Action: [what happens]
- Guard: [when to allow this transition]
- Cost: [latency, operations, resources]

**Transition 2: <state-from> → <state-to>**
- Trigger: [event, signal, or condition]
- Action: [what happens]
- Guard: [when to allow this transition]
- Cost: [latency, operations, resources]

**Transition 3: <state-from> → <state-to>**
- Trigger: [event, signal, or condition]
- Action: [what happens]
- Guard: [when to allow this transition]
- Cost: [latency, operations, resources]

## Observable Signal (migration gate)

**Signal:** What will we watch to know if this state machine is right?

**Trigger:**
- If signal → A: [next step]
- If signal → B: [next step]
- If signal → C: [rollback]

**Rationale:** Why this specific signal and trigger?

## Rollback Plan

How to undo this state machine if it fails:

1. Step 1: [action]
2. Step 2: [action]

**Cost of rollback:** [estimate: low/medium/high]

## Implementation Notes

**Data storage:**
- Where is state persisted?
- Schema: [field definitions]

**Event sourcing:**
- Do we need event sourcing?
- Event schema: [if yes]

**State initialization:**
- How do we get to the initial state?
- Migration path: [how to populate existing data]

## References

- Related decisions: [[DR####]]
- Source material: [URL or doc]
- Stakeholders: [@user1, @user2]

**Why:** This template enforces:
1. Observable signal gate (prevents over-engineering)
2. State transition costs (latency/operations/resources)
3. Guard conditions (what prevents invalid transitions)
4. Implementation notes (data storage, event sourcing, migration)
5. Rollback plan (single ADR rule)

**How to use:**
- Create a new file per state machine (not per file/component)
- Fill all 6 required sections (why now/why not later/why not another/approach/signal/rollback)
- Define states with guard conditions and observable signals
- Define transitions with triggers, actions, guards, and costs
- Link related decisions with [[DR####]]
- Mark state-machine metadata with entity, states, transitions, initial
