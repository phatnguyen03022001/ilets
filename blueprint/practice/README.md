# Practice

## Purpose
Define the canonical **Practice Taxonomy** — a reusable layer of practice **types** referenced by Skill Leaves and Curriculum Nodes. **Not** a collection of exercises ([PR-001](decisions.md)).

## Responsibility (architectural separation)
Practice defines **how things are trained** — distinct from Knowledge (known), Skills (demonstrated), Curriculum (when), Assessment (measured), Progress (runtime). Kept strictly separate ([PR-001](decisions.md)).

## Model ([PR-001](decisions.md), [PR-002](decisions.md))
Every practice type conforms to [practice-type-schema.md](practice-type-schema.md) v1.0 and answers: objective served, cognitive operations, skill leaves supported, knowledge reinforced, when used, purpose (acquisition/consolidation/transfer/exam-readiness), mode (individual/mixed/adaptive/timed/diagnostic). The taxonomy lives in [taxonomy.md](taxonomy.md).

## Status
**Drafted (2026-07-16).** Taxonomy v1.0 (~23 types). Pending Founder validation, then bind: populate each Skill Leaf's `practice_item_types` slot + add practice-type references to Curriculum Nodes.

## Structure
- [taxonomy.md](taxonomy.md) — canonical Practice Types (v1.0).
- [practice-type-schema.md](practice-type-schema.md) — Practice Type schema v1.0.
- [decisions.md](decisions.md) — PR-001 taxonomy rules + architectural separation; PR-002 schema.

## Dependencies
- **Consumes:** [../learning/practice.md](../learning/practice.md) (practice principles), [LD-002](../learning/decisions.md)/[LD-003](../learning/decisions.md) (feedback model/timing), [../skills/](../skills/) + [../knowledge/](../knowledge/) (what it supports/reinforces), [../curriculum/](../curriculum/) (when used).
- **Feeds:** `practice_item_types` on Skill Leaves; practice-type refs in Curriculum Nodes; consumed by `../assessment/` (item→evidence) and `../progress/` (adaptive scheduling).
