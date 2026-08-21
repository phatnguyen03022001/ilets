# Practice

## Purpose
Define the canonical **Practice Taxonomy** — a reusable layer of practice **types** referenced by Skill Leaves and Curriculum Nodes. **Not** a collection of exercises ([PR-001](decisions.md)).

## Responsibility (architectural separation)
Practice defines **how things are trained** — distinct from Knowledge (known), Skills (demonstrated), Curriculum (when), Assessment (measured), Progress (runtime). Kept strictly separate ([PR-001](decisions.md)).

## Model ([PR-001](decisions.md), [PR-002](decisions.md), [PR-003](decisions.md))
Every practice type conforms to [practice-type-schema.md](practice-type-schema.md) v1.1 and answers: objective served, cognitive operations, skill leaves supported, knowledge reinforced, when used, mode, and a **primary learning phase** + supporting phases ([LD-006](../learning/decisions.md)). The taxonomy lives in [taxonomy.md](taxonomy.md).

## Status
**Complete (2026-07-16).** Taxonomy v1.1 (23 types, phase-classified) + [binding.md](binding.md) (leaf→type + node→type) + [coverage-review.md](coverage-review.md). Pending Founder review.

## Structure
- [taxonomy.md](taxonomy.md) — canonical Practice Types (v1.1, phase-classified).
- [binding.md](binding.md) ✅ — authoritative leaf→type + node→type bindings.
- [coverage-review.md](coverage-review.md) — coverage verification.
- [practice-type-schema.md](practice-type-schema.md) — Practice Type schema v1.1.
- [decisions.md](decisions.md) — PR-001/002/003.

## Dependencies
- **Consumes:** [../learning/practice.md](../learning/practice.md) (practice principles), [LD-002](../learning/decisions.md)/[LD-003](../learning/decisions.md) (feedback model/timing), [../skills/](../skills/) + [../knowledge/](../knowledge/) (what it supports/reinforces), [../curriculum/](../curriculum/) (when used).
- **Feeds:** `practice_item_types` on Skill Leaves; practice-type refs in Curriculum Nodes; consumed by `../assessment/` (item→evidence) and `../progress/` (adaptive scheduling).
