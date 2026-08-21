STATUS: CANONICAL
OWNS: how mastery is measured, Assessment Type taxonomy, assessment schema, evidence sufficiency, confidence policy, assessment binding, validity semantics, and certification evidence construction
DEPENDS_ON: 02-IELTS-MODEL.md, 03-SKILLS.md, 04-KNOWLEDGE.md, 05-BANDS.md
DOES_NOT_OWN: Skill/Knowledge definitions, curriculum ordering, practice pedagogy, learner-state transitions, band progression decisions, concrete assessment-item storage

# 08 — Assessment

## Purpose

Define **how demonstrated capability is measured** and when the resulting evidence is strong enough to support a mastery or band-certification decision.

Assessment is independent from Practice:

```text
Practice   trains capability
Assessment measures demonstrated capability
```

The same learner performance may sometimes serve both purposes, but the semantic roles remain distinct.

## Mastery evidence principle

Mastery is criterion-referenced demonstrated competence against the relevant Skill and Band requirements.

Mastery is not defined by:

- time spent;
- lessons completed;
- practice volume;
- one unusually good attempt;
- a generic percentage detached from the target criterion;
- an unsupported model prediction.

Objective receptive scoring may legitimately use official raw-score conversion because that is part of the IELTS scoring model.

## Assessment Type semantic shape

Every canonical Assessment Type owns:

- stable `id`;
- kind: formative, summative/readiness, diagnostic, or certification-supporting;
- target scope;
- evidence produced;
- measurement strategy;
- evidence sufficiency rule;
- confidence requirement;
- uncertainty-surfacing rule;
- alignment to canonical Skill/Band/Knowledge objects;
- AI/human scoring boundary where relevant.

Concrete questions, prompts, submissions, recordings, scores, and rubrics are assessment instances under `10-CONTENT-MODEL.md` or runtime state. They are not new Assessment Types.

## Canonical Assessment Type taxonomy

The active taxonomy contains **7 stable Assessment Types**.

### `AT-01` — Criterion-referenced productive performance

Applies to Writing and Speaking.

Evidence:

- a complete productive performance;
- criterion-level judgments aligned to the relevant Band threshold in `05-BANDS.md`;
- confidence for judgments where an automated/AI scorer is used.

Typical instances:

- Academic Writing Task 1;
- Writing Task 2;
- Speaking performance across relevant parts.

Productive scoring must preserve criterion-level visibility. A single opaque global score is insufficient when criterion failure could invalidate the target band.

### `AT-02` — Objective receptive item set

Applies to Listening and Reading.

Evidence:

- correct/incorrect response scoring against an answer key;
- raw score;
- official section-band conversion from `02-IELTS-MODEL.md`.

Objective scoring is deterministic when the item and answer key are valid. Test-version conversion variation must not be hidden behind a hard-coded universal raw-score table.

### `AT-03` — Knowledge probe

Applies to Knowledge Objects that support target capability.

Possible evidence:

- retrieval;
- recognition;
- gap-fill;
- judgment;
- short production;
- targeted application.

Knowledge acquisition supports learning but does not replace the target IELTS skill demonstration.

### `AT-04` — Diagnostic checkpoint

Produces a scoped learner profile across selected Skill Leaves and/or Knowledge Objects.

A diagnostic checkpoint answers:

- what appears strong;
- what appears weak;
- where evidence is insufficient;
- what should be sampled next.

It is not automatically a mastery certificate.

### `AT-05` — Mastery portfolio

`AT-05` is the certification evidence mechanism.

It aggregates sufficient independent demonstrations from other valid assessment events and determines whether the evidence set satisfies the canonical reliability and confidence policy.

Certification is therefore a consequence of **accumulated sufficient evidence**, not a special one-off test event.

### `AT-06` — Human / human-verified productive assessment

Optional for Writing/Speaking.

Uses expert judgment to score a productive performance or to verify an uncertain automated judgment.

Human review is an optional quality/escalation path, not a mandatory dependency for the core learning system.

### `AT-07` — Full mock test

Measures exam readiness under integrated IELTS-like conditions.

Outputs section estimates and an overall readiness picture.

`AT-07` is **non-certifying by itself**. A mock can contribute evidence only through the same validity and sufficiency rules as other evidence; a single mock result never bypasses `AT-05`.

## Evidence sufficiency

### Productive skills

To support Band-N certification for Writing or Speaking, the evidence set must contain at least **2 independent, time-bounded demonstrations** meeting the Band-N target, sustained rather than coming from one isolated success.

The portfolio must also show no relevant criterion falling below the Band-(N−1) floor where that floor exists.

For Writing, both Task 1 and Task 2 must be represented because `05-BANDS.md` owns both task-level thresholds.

For Speaking, evidence must represent the whole Speaking construct, including all three test parts rather than only rehearsed Part-2 performance.

### Receptive skills

To support Band-N certification for Listening or Reading, the evidence set must contain at least **2 independent full timed section/test demonstrations** whose official conversion reaches Band N, with no result below Band N−1 in the qualifying evidence set.

Official raw-score conversion remains external truth from `02-IELTS-MODEL.md`.

### Knowledge

Knowledge acquisition requires repeated evidence rather than one recognition event. The exact accuracy criterion may differ by Knowledge Object and probe type; it must be criterion-referenced and demonstrated across at least **2 independent probes** before the state is treated as reliably acquired.

A universal arbitrary percentage must not override a more appropriate object-specific criterion.

## Confidence policy

Default minimum confidence for a mastery/certification judgment is:

```text
0.80
```

This is a **calibratable policy default**, not an immutable architectural constant.

Rules:

- low-confidence productive evidence cannot certify mastery on its own;
- low-confidence evidence triggers more evidence collection, re-probing, or optional human verification;
- confidence should be surfaced at the level where uncertainty can change the decision;
- confidence estimates must be empirically calibrated before production use; a model's self-reported confidence is not sufficient evidence of calibration.

Changing the threshold after empirical validation is a policy calibration. It does not require a new Assessment Type or new authority domain.

## Formative, readiness, and certification are distinct

### Formative assessment

Used during learning to identify errors, guide feedback, and decide what to practice next.

Typical types: `AT-01`, `AT-02`, `AT-03`, `AT-04` in low-stakes use.

### Exam-readiness assessment

Estimates performance under exam-like conditions.

Typical type: `AT-07`.

A readiness estimate answers "How might the learner perform now under test conditions?" It does not answer "Has the learner reliably mastered the full target standard over time?"

### Certification evidence

`AT-05` answers the latter question by evaluating accumulated evidence against the sufficiency and confidence policy.

## Assessment binding

Assessment binds to canonical objects through IDs:

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

Rules:

1. Assessment Types never copy Skill or Knowledge definitions.
2. Band criteria are referenced from `05-BANDS.md`.
3. A Skill Leaf may be measured by more than one assessment strategy.
4. One integrated productive task may generate evidence for multiple leaves/criteria.
5. Integrated evidence must remain attributable enough to identify criterion-level weakness.
6. Evidence records carry the canonical IDs and conditions under which the evidence was produced.

## Validity requirements

Evidence is valid for a target decision only when:

- the task actually samples the target capability;
- scoring aligns to the canonical criterion;
- task conditions are appropriate to the claim being made;
- the response is attributable to the learner rather than excessive assistance;
- the evidence is sufficiently independent from the other demonstrations in the portfolio;
- automated judgments meet the confidence/calibration policy;
- the evidence is not being used outside the scope it was designed to support.

Examples:

- guided acquisition writing cannot be treated as equivalent to independent timed Writing evidence;
- a vocabulary quiz cannot certify Writing Band 7;
- a full mock can estimate readiness but does not automatically prove sustained mastery;
- a low-confidence productive score must not silently become a progression gate.

## AI assessment boundary

AI may be the primary automated assessor in the future product, especially for rapid feedback and repeated sampling, but productive-skill judgment must be confidence-aware and empirically validated.

Policy:

- objective scoring uses deterministic keys where possible;
- productive AI scoring reports criterion-level output and calibrated uncertainty;
- uncertain cases are re-probed or optionally human-verified;
- AI convenience never lowers the evidence standard.

## Evidence output required by Progression

A certification-ready evidence set must allow `09-PROGRESSION.md` to determine, without re-scoring the learning semantics:

- target Skill/Band;
- demonstrations included;
- criterion outcomes;
- timestamps/order sufficient to establish independence and recency;
- confidence where relevant;
- whether the sufficiency rule is satisfied;
- whether regression evidence exists.

Progression consumes this result. It does not redefine the assessment standard.
