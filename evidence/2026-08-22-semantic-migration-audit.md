TYPE: EVIDENCE
STATUS: NON-CANONICAL
OWNER: 2026-08-22 document-architecture refactor

# Semantic Migration Audit

This file records how legacy Blueprint semantics were treated during the migration from the fragmented `blueprint/` hierarchy to the active domain-owner architecture.

It has **no authority**. Canonical truth is defined by the owning active spec.

## Migration identity

Legacy baseline commit:

```text
63aed1debaa6c3fa4cfe10e754ce0d9101092432
```

First structural refactor commit:

```text
43cddc45160016beb6d2d9342acc2fd9db4b1075
```

The exact baseline root tree is preserved at:

```text
archive/legacy-2026-07-16/
```

with original tree SHA:

```text
48fb1af1c55da867a977594f50b3eae9f890ca61
```

## Invariants preserved

The frozen stable object inventories remain represented in the active specs:

- 64 Skill Leaves: 11 Listening, 12 Reading, 18 Speaking, 23 Writing;
- 46 Knowledge Objects: 29 grammar, 9 vocabulary, 8 phonology;
- 44 Curriculum Nodes;
- 23 Practice Types;
- 7 Assessment Types.

Stable IDs were preserved.

## Legacy Skill Leaf schema migration

| Legacy field/semantic | Active treatment |
|---|---|
| `id`, `name`, `skill`, `component`, atomic objective | retained in `spec/03-SKILLS.md` |
| atomicity / independence invariant | retained in `spec/03-SKILLS.md` |
| skill-to-skill prerequisites | owned by `spec/03-SKILLS.md` |
| knowledge prerequisites | resolved canonically in `spec/04-KNOWLEDGE.md` |
| `bands` / introduced-refined ranges | rehomed to `spec/05-BANDS.md` and `spec/06-CURRICULUM.md`; no longer intrinsic leaf ownership |
| `traces_to` IELTS descriptors | external fact in `spec/02-IELTS-MODEL.md`, target interpretation in `spec/05-BANDS.md` |
| per-leaf `mastery_criteria` | rehomed to Band/task/skill threshold ownership in `spec/05-BANDS.md` rather than duplicated on leaves |
| `assessment_strategy` | rehomed to `spec/08-ASSESSMENT.md` |
| `mastery_states` | rehomed to `spec/09-PROGRESSION.md` |
| `practice_item_types` | rehomed to `spec/07-PRACTICE.md` bindings by stable ID |
| `dependents` | treated as graph-derived, not separately authored canonical truth |
| `cognitive_level` | usable as planning/content metadata; no longer required as identity-level canonical truth for every leaf |
| `typical_learning_load` | legacy heuristic retained in archive; requires calibration before being promoted as an active planning value |
| `common_errors`, `remediation` | treated as concrete tutoring/content/feedback knowledge under `spec/10-CONTENT-MODEL.md` unless future evidence shows a stable domain-level rule deserves canonical promotion |
| `schema_version` | retired as document bureaucracy; Git history is the version history |

Reason: the legacy leaf object had accumulated fields owned semantically by Practice, Assessment, Bands, Progression, and content generation. The new architecture preserves those semantics at their actual owners instead of keeping a cross-domain mega-object.

## Legacy Knowledge Object schema migration

| Legacy field/semantic | Active treatment |
|---|---|
| `id`, `name`, `domain`, atomic unit | retained in `spec/04-KNOWLEDGE.md` |
| `requires` prerequisite edges | retained in `spec/04-KNOWLEDGE.md` |
| Skill→Knowledge resolution | retained in `spec/04-KNOWLEDGE.md` |
| `band_relevance` | rehomed to Bands/Curriculum; not intrinsic knowledge ownership |
| examples | treated as concrete content instances unless an example is necessary to define the concept itself |
| common misconceptions | treated as tutoring/remediation content unless future evidence justifies canonical domain-level promotion |
| `related_to` soft peer edges | legacy graph remains in archive; active promotion should occur only when a peer edge has a concrete learning/selection consequence |
| `schema_version` | retired; Git history owns version history |

## Legacy Curriculum Node schema migration

Retained canonical semantics in `spec/06-CURRICULUM.md`:

- stable ID;
- band phase;
- recommended sequence;
- Skill/Knowledge references;
- focus;
- expected learning outcome;
- sequencing rationale and prerequisite policy;
- node exit intent.

Legacy `typical_learning_duration` and `estimated_load` values were heuristic planning estimates, not empirically calibrated learner truth. The active spec preserves the ability to use such planning metadata but does not promote the legacy estimates as current calibrated facts.

Consolidation nodes are orchestration/integration nodes, **not synchronized four-skill certification gates**. Per-skill advancement is exclusively owned by `spec/09-PROGRESSION.md`.

## Legacy Practice schema migration

Retained in `spec/07-PRACTICE.md`:

- 23 stable IDs;
- objective/purpose;
- primary phase + supported phases;
- skill/knowledge scope;
- mode;
- feedback-timing semantics;
- type-vs-instance distinction;
- binding by canonical references.

Concrete item fields remain instance-level content in `spec/10-CONTENT-MODEL.md`.

## Legacy Assessment schema migration

Retained in `spec/08-ASSESSMENT.md`:

- 7 stable IDs;
- evidence produced;
- measurement strategy;
- sufficiency;
- confidence requirement and surfacing;
- kind/scope;
- band/skill alignment by reference;
- type-vs-instance distinction;
- Practice/Assessment independence.

The productive-scoring policy was strengthened: automated confidence must be empirically calibrated before production use; model self-confidence alone is not treated as calibration evidence.

## Progression correction preserved

The accepted legacy correction from synchronized four-skill advancement to **independent per-skill progression** remains canonical in `spec/09-PROGRESSION.md`.

Current certification may regress to in-progress after later valid evidence; historical attainment remains in certification history and re-certification requires fresh valid evidence.

## External IELTS recheck

During the migration, current official IELTS pages were rechecked for:

- four sections;
- Academic vs General Training sharing/differences;
- Listening four-part/40-question structure;
- Academic test timing;
- official Listening and Reading average score anchors;
- Writing criteria and Task-2 weighting;
- Speaking criteria and equal weighting.

Current external exam truth is owned by `spec/02-IELTS-MODEL.md`, not by this audit.

## Structural result

Active authority is now:

```text
CONSTITUTION.md
      ↓
OBJECTIVE.md
      ↓
spec/ canonical owners
```

`README.md` is navigation. `spec/DECISIONS.md` is rationale. `research/`, `evidence/`, and `archive/` are non-canonical.

No vendor-specific active understanding layer remains.
