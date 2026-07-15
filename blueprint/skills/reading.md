# Reading — Skill Decomposition (schema v1.1)
*Decomposes the Reading (Academic) competency from [../bands/reading.md](../bands/reading.md). Receptive — no public analytic descriptors (band = raw-score conversion) — so this is **Blueprint Inference** ([BD-003](../bands/decisions.md)): comprehension abilities + question-type strategies, traced to the receptive band range and official test format. Conforms to [leaf-schema.md](leaf-schema.md) v1.1.*

## Hierarchy overview
```
Reading (Academic)
├─ R-COMP  Comprehension abilities
│   ├─ R-COMP-01  Skim for gist / main idea
│   ├─ R-COMP-02  Scan for specific detail
│   ├─ R-COMP-03  Detailed comprehension
│   ├─ R-COMP-04  Inference (implications, writer's views/claims)
│   ├─ R-COMP-05  Understand text structure / paragraph purpose
│   └─ R-COMP-06  Manage dense/abstract/complex passages
└─ R-QT    Question-type strategies
    ├─ R-QT-01  Matching headings
    ├─ R-QT-02  True/False/Not Given (identifying information)
    ├─ R-QT-03  Yes/No/Not Given (writer's views)
    ├─ R-QT-04  Matching information/features/sentence endings
    ├─ R-QT-05  Completion (sentence/note/summary/diagram) & short-answer
    └─ R-QT-06  Multiple choice
```

---

## R-COMP — Comprehension abilities

**`R-COMP-01` Skim for gist / main idea** — reading · COMP · bands 4–9 · cognitive: understand · load: medium
- objective: quickly identify the main idea/overall topic of a passage.
- traces_to: Reading receptive overlay ([../bands/reading.md](../bands/reading.md)); Blueprint Inference.
- prerequisites: — .
- mastery_criteria: states the gist of an unfamiliar passage after a timed skim.
- common_errors: reading every word; missing the global topic.
- remediation: timed skimming; heading-prediction.
- consumer fields: _(populated by practice/ & assessment/; dependents derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`R-COMP-02` Scan for specific detail** — reading · COMP · bands 4–9 · cognitive: understand · load: medium
- objective: locate specific information (dates, names, figures) quickly.
- traces_to: receptive overlay; Blueprint Inference.
- prerequisites: — .
- mastery_criteria: locates targeted details accurately under time pressure.
- common_errors: slow linear reading; missing located detail.
- remediation: scanning drills; keyword→location practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`R-COMP-03` Detailed comprehension** — reading · COMP · bands 5–9 · cognitive: understand · load: high
- objective: understand detailed, explicit information across a passage.
- traces_to: receptive overlay; Blueprint Inference.
- prerequisites: R-COMP-01.
- mastery_criteria: answers detail questions accurately across a full passage.
- common_errors: partial reading; misreading nuance.
- remediation: close-reading practice; accuracy-over-speed drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`R-COMP-04` Inference (implications, writer's views/claims)** — reading · COMP · bands 5–9 · cognitive: analyze · load: high
- objective: infer meaning, implications, and the writer's views/claims not stated explicitly.
- traces_to: receptive overlay — Band 6 "some inference" → Band 7 "handle inference, writer's views reliably".
- prerequisites: R-COMP-03.
- mastery_criteria: answers inferential/view-based items reliably.
- common_errors: over-literal reading; assuming views not supported; confusing fact vs claim.
- remediation: inference-question strategy; evidence-justification drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`R-COMP-05` Understand text structure / paragraph purpose** — reading · COMP · bands 5–9 · cognitive: analyze · load: high
- objective: recognize how a text and its paragraphs are organized and why.
- traces_to: receptive overlay; Blueprint Inference.
- prerequisites: R-COMP-01.
- mastery_criteria: identifies paragraph purpose and overall organization reliably.
- common_errors: treating paragraphs as isolated; missing rhetorical function.
- remediation: structure-mapping; topic-sentence analysis.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`R-COMP-06` Manage dense/abstract/complex passages** — reading · COMP · bands 6–9 · cognitive: understand · load: high
- objective: comprehend demanding passages (abstract lexis, complex syntax, dense argument).
- traces_to: receptive overlay — Band 6 "some difficulty with dense/abstract" → Band 8 "handle detailed, complex, abstract texts with little difficulty".
- prerequisites: R-COMP-03; `K-VOC` academic vocabulary.
- mastery_criteria: answers items on the most demanding passage reliably.
- common_errors: cognitive overload; giving up on hard passages.
- remediation: vocabulary front-loading; scaffolded dense-text practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

## R-QT — Question-type strategies

**`R-QT-01` Matching headings** — reading · QT · bands 5–9 · cognitive: analyze · load: high
- objective: match headings to paragraphs based on main ideas.
- traces_to: official Reading format; receptive overlay.
- prerequisites: R-COMP-05.
- mastery_criteria: matches headings accurately across a passage.
- common_errors: keyword-matching over meaning; one-to-one bias.
- remediation: heading-strategy drills; paragraph-gist first.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`R-QT-02` True/False/Not Given (identifying information)** — reading · QT · bands 5–9 · cognitive: evaluate · load: high
- objective: classify statements as true/false/not given against the text.
- traces_to: official Reading format; receptive overlay.
- prerequisites: R-COMP-03.
- mastery_criteria: classifies statements reliably; distinguishes False from Not Given.
- common_errors: confusing False and Not Given; assuming prior knowledge.
- remediation: T/F/NG logic drills; text-only-evidence practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`R-QT-03` Yes/No/Not Given (writer's views)** — reading · QT · bands 6–9 · cognitive: evaluate · load: high
- objective: classify whether statements agree with the writer's views/claims.
- traces_to: official Reading format; receptive overlay.
- prerequisites: R-COMP-04.
- mastery_criteria: classifies view-based statements reliably.
- common_errors: treating as factual T/F/NG; missing writer stance.
- remediation: writer-view strategy; stance-spotting drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`R-QT-04` Matching information/features/sentence endings** — reading · QT · bands 5–9 · cognitive: analyze · load: high
- objective: match items to locations/features/endings across the text.
- traces_to: official Reading format; receptive overlay.
- prerequisites: R-COMP-02, R-COMP-03.
- mastery_criteria: matches items accurately across a passage.
- common_errors: one-to-one bias; losing track across paragraphs.
- remediation: matching-strategy drills; location-tracking practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`R-QT-05` Completion (sentence/note/summary/diagram) & short-answer** — reading · QT · bands 4–9 · cognitive: understand · load: medium
- objective: complete items using words from the text within limits.
- traces_to: official Reading format; receptive overlay.
- prerequisites: R-COMP-03.
- mastery_criteria: completes items accurately, within word limits, grammatically fitted.
- common_errors: exceeding word limit; wrong word form; copying wrong segment.
- remediation: word-limit drills; grammar-fit checking.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`R-QT-06` Multiple choice** — reading · QT · bands 5–9 · cognitive: apply · load: medium
- objective: select correct single/multiple answers from options.
- traces_to: official Reading format; receptive overlay.
- prerequisites: R-COMP-03, R-COMP-04.
- mastery_criteria: selects correct options reliably, including multi-answer/inference variants.
- common_errors: distractor traps; partial-correct options.
- remediation: option-elimination; evidence-justification drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

---

## Coverage check
- Comprehension → R-COMP-01..06 · Question-types → R-QT-01..06. Together they account for the Academic Reading raw-score band (every official question type is a practiceable/assessable leaf). *(Self-check vs [../bands/reading.md](../bands/reading.md).)*

## Dependencies
- **Consumes:** [../bands/reading.md](../bands/reading.md) (Academic conversion + receptive overlay).
- **References:** `K-VOC` academic vocabulary ([../knowledge/](../knowledge/)).
- **Feeds:** [../curriculum/](../curriculum/), [../practice/](../practice/), [../assessment/](../assessment/).
- **GT Reading** (different conversion + passage types) is the variant-specific part, added later per [FD-001](../product/foundational-decisions.md).

## Open questions
- [ ] Receptive decomposition is **Blueprint Inference** — confirm the comprehension×question-type split (same as Listening) is the right shape.
- [ ] Whether Yes/No/Not Given (R-QT-03) and T/F/Not Given (R-QT-02) should be separate leaves — kept separate because they test different things (facts vs writer's views) and are independently teachable/assessable.
