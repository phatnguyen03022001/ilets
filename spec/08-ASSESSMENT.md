STATUS: CANONICAL
OWNS: how mastery is measured, Assessment Type taxonomy, type-level kind/scope/evidence semantics, evidence sufficiency, confidence policy, validity, and certification evidence construction
DEPENDS_ON: 02-IELTS-MODEL.md, 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md
DOES_NOT_OWN: Skill/Knowledge definitions, curriculum ordering, practice pedagogy, learner-state transitions, band-progression decisions, or concrete assessment-item storage

# 08 — Assessment

## Purpose

Define **how demonstrated capability is measured** and when evidence is strong enough to support mastery/certification.

Practice trains capability. Assessment measures demonstrated capability. A performance may participate in both roles, but the semantics remain separate.

## Mastery evidence principle

Mastery is criterion-referenced demonstrated competence against canonical Skill and Band requirements. It is not defined by time spent, lessons completed, practice volume, one unusually good attempt, a generic detached percentage, or an unsupported model prediction.

Objective receptive scoring may use official raw-score conversion because that conversion is part of external IELTS truth.

# Canonical Assessment Type registry

| ID | Type | Kind | Scope | Evidence / role | Certification status |
|---|---|---|---|---|---|
| `AT-01` | Criterion-referenced productive performance | formative or summative measurement | Writing, Speaking | Complete productive performance scored by relevant criteria with uncertainty where automated | contributes evidence; never certifies by one attempt |
| `AT-02` | Objective receptive item set | formative or summative measurement | Listening, Reading | keyed item results, raw score, official section-band conversion | contributes evidence; repeated full timed demonstrations required for certification |
| `AT-03` | Knowledge probe | formative measurement | Knowledge Objects / dependent leaves | retrieval, recognition, gap-fill, judgment, short production, targeted application | supports knowledge acquisition; cannot replace target-skill evidence |
| `AT-04` | Diagnostic checkpoint | diagnostic | cross-skill / knowledge | scoped mastery profile showing strengths, weaknesses, uncertainty, next sampling need | non-certifying |
| `AT-05` | Mastery portfolio | cumulative certification mechanism | all target skills/bands | aggregates valid independent demonstrations and applies sufficiency/confidence policy | **the certification evidence mechanism** |
| `AT-06` | Human / human-verified productive assessment | optional summative/verification | Writing, Speaking | expert-scored performance or expert verification of uncertain automated score | may contribute valid productive evidence; human review is optional, not mandatory |
| `AT-07` | Full mock test | summative readiness | all four skills | integrated exam-like section estimates and readiness picture | non-certifying by itself; may only contribute evidence through normal validity/sufficiency rules |

## Type semantics

### `AT-01` productive performance

Requires criterion-level visibility for Writing/Speaking. A single opaque global score is insufficient when one criterion can invalidate the target Band. Automated scoring must surface calibrated uncertainty.

### `AT-02` receptive item set

Uses deterministic keys where the item/key are valid. Raw-score-to-band conversion comes from current official truth in `02-IELTS-MODEL.md`; test-version variation must not be hidden behind a universal hard-coded table.

### `AT-03` knowledge probe

Measures enabling knowledge. Passing a vocabulary/grammar probe cannot certify an IELTS skill because knowledge supports but does not substitute for capability demonstration.

### `AT-04` diagnostic checkpoint

Answers what appears strong/weak, where evidence is insufficient, and what should be sampled next. `PT-22` may provide a diagnostic practice format; `AT-04` owns the measurement interpretation.

### `AT-05` mastery portfolio

Aggregates sufficient independent demonstrations from valid evidence records. Certification is therefore a consequence of accumulated evidence rather than a special one-off test.

### `AT-06` human verification

Optional escalation for productive skills. It can replace/verify an uncertain automated judgment, but the core system must not require a human reviewer to function.

### `AT-07` full mock

Measures readiness under integrated exam-like conditions. A single mock cannot bypass `AT-05`.

# Evidence sufficiency

## Productive skills

Band-N certification for Writing/Speaking requires at least **2 independent, time-bounded demonstrations** meeting the Band-N target, sustained rather than arising from one isolated success, with no relevant criterion falling below the Band-(N−1) floor where that floor exists.

Writing evidence must represent both Task 1 and Task 2. Speaking evidence must represent the whole construct across Parts 1, 2, and 3 rather than a rehearsed Part-2-only sample.

## Receptive skills

Band-N certification for Listening/Reading requires at least **2 independent full timed demonstrations** whose official conversion reaches Band N, with no result below Band N−1 in the qualifying set.

## Knowledge

Knowledge acquisition requires repeated evidence across at least **2 independent probes** before being treated as reliably acquired. Accuracy criteria may differ by object/probe and must remain criterion-referenced; no arbitrary universal percentage overrides a more suitable criterion.

# Confidence policy

Default minimum confidence for uncertain mastery/certification judgment:

```text
0.80
```

This is a **calibratable policy default**, not an architectural constant.

Rules:

- low-confidence productive evidence cannot certify mastery on its own;
- low-confidence evidence triggers re-probing, more evidence, or optional human verification;
- uncertainty is surfaced at the level where it can change the decision;
- productive automated confidence must be empirically calibrated before production use;
- model self-reported confidence alone is not calibration evidence.

# Formative vs readiness vs certification

- **Formative** assessment guides learning and next practice.
- **Readiness** assessment estimates likely exam-condition performance.
- **Certification** requires `AT-05` evaluation of accumulated valid evidence.

These roles are not interchangeable. A learner may perform well on one mock yet lack sustained certification evidence, or be currently certified yet underperform on one mock day.

# Assessment binding

```text
Assessment Type
      +
Skill Leaf / Knowledge Object target
      +
Band threshold when relevant
      ↓
Concrete Assessment Item / Performance
      ↓
Evidence Record
```

Binding rules:

1. bind by canonical IDs;
2. never copy Skill/Knowledge definitions into assessment types/items;
3. Band criteria are referenced from `05-BANDS.md`;
4. one integrated task may generate evidence for multiple leaves/criteria;
5. integrated evidence remains attributable enough to expose criterion-level weakness;
6. evidence records preserve target IDs and task conditions.

# Evidence validity

Evidence is valid for a decision only when:

- the task actually samples the target capability;
- scoring aligns to the canonical criterion;
- conditions are appropriate to the claim;
- the response is attributable to the learner rather than excessive assistance;
- demonstrations are sufficiently independent;
- automated judgments satisfy calibration/confidence policy;
- evidence is not used outside its designed scope.

Examples of invalid shortcuts:

- guided acquisition writing treated as independent timed Writing evidence;
- vocabulary quiz used to certify Writing Band 7;
- full mock treated as automatic mastery certification;
- low-confidence productive score silently used as a progression gate.

# AI assessment boundary

AI may be the primary automated assessor for repeated feedback/sampling, but productive judgment must be confidence-aware and empirically validated. Objective items should use deterministic keys where possible; uncertain productive judgments are re-probed or optionally human-verified. Convenience never lowers the evidence standard.

# Output consumed by Progression

A certification-ready evidence set must expose:

- target Skill/Band;
- included demonstrations;
- criterion outcomes;
- timing/order sufficient to establish independence and recency;
- confidence where relevant;
- whether sufficiency is satisfied;
- regression evidence where present.

`09-PROGRESSION.md` consumes this result. It does not rescore or redefine the assessment standard.