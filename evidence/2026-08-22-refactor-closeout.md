TYPE: EVIDENCE
STATUS: NON-CANONICAL
OWNER: 2026-08-22 document-architecture refactor close-out

# Refactor Close-out Audit

This report records the final static audit of the repository-document refactor. It has **no specification authority**. If this report conflicts with an active canonical owner, the canonical owner wins.

## Baseline and final state

Legacy baseline commit:

```text
63aed1debaa6c3fa4cfe10e754ce0d9101092432
```

Canonical refactor close-out parent HEAD:

```text
15a353be8bbb4d0baf739cb1346a75ba5c8447b3
```

Root tree at that HEAD:

```text
bdde44b82ff4dd7582063392b3c63ed4a0f76e01
```

`git` ancestry comparison reports the refactored HEAD **ahead** of the baseline and **not behind** it. No force rewrite was used for the refactor sequence.

## Structural gate — PASS

Active root contains exactly:

```text
CONSTITUTION.md
OBJECTIVE.md
README.md
archive/
evidence/
research/
spec/
```

The active root contains no `.claude/`, `CLAUDE.md`, or `blueprint/` authority surface.

`spec/` contains exactly 13 active specification files:

```text
00-PRODUCT.md
01-LEARNER-MODEL.md
02-IELTS-MODEL.md
03-SKILLS.md
04-KNOWLEDGE.md
05-BANDS.md
06-CURRICULUM.md
07-PRACTICE.md
08-ASSESSMENT.md
09-PROGRESSION.md
10-CONTENT-MODEL.md
11-GLOSSARY.md
DECISIONS.md
```

Authority hierarchy is:

```text
CONSTITUTION.md
      ↓
OBJECTIVE.md
      ↓
owning spec/*.md
```

`README.md` is navigation. `spec/DECISIONS.md`, `research/`, `evidence/`, and `archive/` do not independently override canonical owners.

## Historical-preservation gate — PASS

The archive contains:

```text
archive/legacy-2026-07-16/
```

and that path points to the exact pre-refactor root tree:

```text
48fb1af1c55da867a977594f50b3eae9f890ca61
```

Therefore the structural refactor did not require deleting historical semantics or provenance.

## Stable-inventory gate — PASS

The active owner specs preserve the frozen stable registries and namespaces:

- 64 Skill Leaves: 11 Listening, 12 Reading, 18 Speaking, 23 Writing;
- 46 Knowledge Objects: 29 Grammar, 9 Vocabulary, 8 Phonology;
- 44 Curriculum Nodes;
- 23 Practice Types;
- 7 Assessment Types.

Stable ID namespaces remain the contract between owners and concrete content.

## Ownership gate — PASS

The final active ownership split is explicit:

- `02-IELTS-MODEL.md` — external IELTS reality;
- `03-SKILLS.md` — atomic demonstrated capabilities and Skill→Skill prerequisites;
- `04-KNOWLEDGE.md` — atomic enabling concepts, Knowledge→Knowledge prerequisites, Skill→Knowledge resolution;
- `05-BANDS.md` — Band-N quality thresholds and task/skill exits;
- `06-CURRICULUM.md` — orchestration and sequencing;
- `07-PRACTICE.md` — learning phases and reusable training types;
- `08-ASSESSMENT.md` — measurement, evidence validity/sufficiency, confidence and certification evidence;
- `09-PROGRESSION.md` — learner runtime state and transitions;
- `10-CONTENT-MODEL.md` — concrete/supporting content contracts;
- `11-GLOSSARY.md` — terminology only.

No owner is intended to duplicate another owner's behavioral authority.

## Prerequisite semantics gate — PASS

The executable interpretation is explicit:

- Skill/Knowledge `requires` edges are **Required prerequisites**;
- Curriculum `Depends` entries are **Recommended sequencing by default** unless explicitly marked Required;
- runtime hard-gate enforcement belongs to Progression;
- new hard dependencies require evidence/theory justification rather than convenience.

Integration rows using `Band-N target set` resolve deterministically to the union of preceding explicit targets in that same band phase. Integration rows are orchestration checkpoints, not synchronized four-skill certification gates.

## Progression/certification gate — PASS

- each IELTS skill certifies independently;
- overall band is informational for learning progression, not a hard gate;
- `AT-05` accumulated valid evidence is the certification mechanism;
- full mocks / Exam Preparation do not certify by exposure alone;
- later valid regression may move current certification back to in-progress;
- historical attainment remains preserved for audit/history;
- re-certification uses fresh normal evidence.

## Evidence-policy gate — PASS

The active Assessment owner keeps Practice and Assessment separate and makes these defaults explicit:

- productive certification requires repeated independent demonstrations;
- receptive certification requires repeated full timed demonstrations using official conversion;
- knowledge acquisition requires repeated probes;
- default confidence threshold for uncertain mastery judgment is 0.80 and calibratable;
- model self-reported confidence is not empirical calibration evidence;
- optional human review may verify uncertain productive judgments but is not a mandatory dependency.

## Semantic-delegation gate — PASS

Legacy fields intentionally removed from Skill/Knowledge identity now have explicit active destinations instead of relying on the archive:

- Band relevance / quality → Bands + Curriculum;
- practice binding → Practice;
- assessment strategy/evidence → Assessment;
- learner state → Progression;
- concrete examples → Stimulus / content instances;
- recurring errors and misconceptions → `ErrorPattern`;
- remediation strategies → `RemediationPattern`;
- generated/runtime feedback → `FeedbackArtifact`;
- uncalibrated load/duration heuristics remain historical/supporting until evidence justifies promotion.

`10-CONTENT-MODEL.md` now explicitly represents `ErrorPattern` and `RemediationPattern`, and `11-GLOSSARY.md` defines the corresponding terminology.

## Reference gate — PASS

Canonical object bindings use complete stable IDs. Practice may use documented family notation such as `L-COMP-*` at the reusable type level, but concrete items/evidence must resolve such families to explicit canonical IDs before execution or recording.

## External-evidence gate — PASS for documented baseline

Current official IELTS structure/scoring facts were rechecked during the refactor and recorded separately in:

```text
evidence/2026-08-22-ielts-official-baseline.md
```

The canonical external-fact owner is still `spec/02-IELTS-MODEL.md`, not the evidence record.

## Validation scope and limitation

This repository is currently a documentation/specification repository. This close-out is therefore a **static architecture and semantic-consistency audit**, not a runtime software test.

No claim is made here that application code, API schemas, persistence, model calibration, generated content quality, or production IELTS scoring has been implemented or runtime-tested.

## Verdict

```text
STRUCTURAL REFACTOR: PASS
SEMANTIC OWNERSHIP: PASS
LEGACY PRESERVATION: PASS
IMPLEMENTATION-FACING CONTRACT COMPLETENESS: PASS
RUNTIME IMPLEMENTATION/TESTING: OUT OF SCOPE
```

The repository can now be understood from the active authority surface without treating the archived Blueprint or vendor-specific instructions as required implementation documentation.
