# Writing — Skill Decomposition (schema v1.1)
*Decomposes the Writing competency from [../bands/writing.md](../bands/writing.md) into leaf skills conforming to [leaf-schema.md](leaf-schema.md) v1.1. Reference implementation of the canonical learning object ([SK-002](decisions.md), [SK-003](decisions.md)).*

Each leaf block lists the schema fields the skills/ phase populates; consumer fields (`practice_item_types`, `assessment_strategy`, `dependents`) are owned by the leaf and populated later by the consuming section. `mastery_states` (`not_started → practicing → emerging → mastered`) is defined once in the schema.

## Hierarchy overview
```
Writing
├─ W-TA  Task Achievement (Task 1)
│   ├─ W-TA-01  Identify key features
│   ├─ W-TA-02  Write an overview
│   └─ W-TA-03  Report key features with supporting data
├─ W-TR  Task Response (Task 2)
│   ├─ W-TR-01  Analyse the prompt
│   ├─ W-TR-02  Formulate a clear position
│   ├─ W-TR-03  Develop main ideas with support
│   └─ W-TR-04  Maintain relevance
├─ W-CC  Coherence & Cohesion
│   ├─ W-CC-01  Paragraphing
│   ├─ W-CC-02  Overall logical progression
│   ├─ W-CC-03  Cohesive devices (linkers)
│   └─ W-CC-04  Reference & substitution
├─ W-LR  Lexical Resource
│   ├─ W-LR-01  Topic-specific vocabulary
│   ├─ W-LR-02  Collocation
│   ├─ W-LR-03  Paraphrase
│   ├─ W-LR-04  Less common / precise / idiomatic vocabulary
│   └─ W-LR-05  Spelling & word formation
└─ W-GRA Grammatical Range & Accuracy
    ├─ W-GRA-01  Simple sentence accuracy
    ├─ W-GRA-02  Compound sentences
    ├─ W-GRA-03  Complex sentences (subordination)
    ├─ W-GRA-04  Tense & aspect
    ├─ W-GRA-05  Articles & determiners
    ├─ W-GRA-06  Punctuation
    └─ W-GRA-07  Structural flexibility/variety
```
*Length/format compliance (≥150 words T1 / ≥250 words T2) is a shared production constraint, not a leaf; enforced in `../practice/`.*

---

## W-TA — Task Achievement (Task 1)

**`W-TA-01` Identify key features** — writing · TA · bands 4–9 · cognitive: analyze · load: medium
- objective: identify the most important features (main trends, differences, stages, comparisons) in a graph/chart/table/diagram/map.
- traces_to: Task 1 TA — Band 4 "few key features selected" → Band 8 "key features skilfully selected".
- prerequisites: — (basic graph/numeracy literacy assumed).
- mastery_criteria: reliably identifies the 3–5 key features from a novel visual, matching an expert key.
- common_errors: listing all data points; missing the main trend; focusing on detail over overview features.
- remediation: guided feature-selection from worked exemplars; compare to expert key.
- practice_item_types / assessment_strategy / dependents: _(populated by practice/ & assessment/; dependents derived)_.
- independence: ✅ atomic; independently teachable / practiceable / assessable / remediable.

**`W-TA-02` Write an overview** — writing · TA · bands 6–9 · cognitive: analyze · load: medium
- objective: write a clear overview statement summarizing main trends/differences/stages, distinct from detail.
- traces_to: Task 1 TA — Band 6 "a relevant overview is attempted" → Band 7 "presents a clear overview".
- prerequisites: W-TA-01.
- mastery_criteria: produces an accurate overview distinct from supporting detail in ≥2 Task 1 responses.
- common_errors: overview missing; overview mixed with detail; inaccurate summarization.
- remediation: overview-sentence templates; separate overview from body practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-TA-03` Report key features with supporting data** — writing · TA · bands 5–9 · cognitive: apply · load: medium
- objective: select relevant key features and support them with accurate figures/data, avoiding irrelevant/inaccurate detail.
- traces_to: Task 1 TA — Band 6 "supported using figures/data" → Band 7 "covered and clearly highlighted".
- prerequisites: W-TA-01.
- mastery_criteria: supports key features with correct data across ≥2 responses; no large irrelevant/inaccurate sections.
- common_errors: inventing/incorrect figures; over-describing minor detail; copying input verbatim.
- remediation: data-accuracy drills; select-then-report scaffolds.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

## W-TR — Task Response (Task 2)

**`W-TR-01` Analyse the prompt** — writing · TR · bands 4–9 · cognitive: analyze · load: medium
- objective: identify all parts of the essay prompt and the required task type (opinion / discussion / advantages-disadvantages / problem-solution).
- traces_to: Task 2 TR — Band 4 "misunderstanding of the prompt" → Band 6 "main parts of the prompt are addressed".
- prerequisites: — .
- mastery_criteria: correctly identifies all prompt parts + task type across ≥3 novel prompts.
- common_errors: missing a prompt part; misidentifying task type; answering a different question.
- remediation: prompt-deconstruction routine; task-type classification practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-TR-02` Formulate a clear position** — writing · TR · bands 5–9 · cognitive: create · load: high
- objective: formulate and state a clear position that directly answers the question.
- traces_to: Task 2 TR — Band 5 "expresses a position, development unclear" → Band 7 "clear and developed position".
- prerequisites: W-TR-01.
- mastery_criteria: states a clear, relevant position sustained throughout in ≥2 responses.
- common_errors: no clear position; contradictory stance; position not answering the question.
- remediation: thesis-statement practice; position-consistency checks.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-TR-03` Develop main ideas with support** — writing · TR · bands 5–9 · cognitive: create · load: high
- objective: extend each main idea with relevant reasons, examples, and evidence.
- traces_to: Task 2 TR — Band 6 "main ideas… insufficiently developed" → Band 7 "main ideas extended and supported".
- prerequisites: W-TR-02.
- mastery_criteria: each main idea extended with relevant, sufficient support in ≥2 responses.
- common_errors: undeveloped ideas; irrelevant examples; repetition instead of development.
- remediation: idea-development frameworks (PEEL/claim-reason-evidence); expansion drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-TR-04` Maintain relevance** — writing · TR · bands 5–9 · cognitive: evaluate · load: medium
- objective: keep all content relevant to the prompt; avoid irrelevant or repetitive material.
- traces_to: Task 2 TR — Band 5 "may be irrelevant detail" → Band 7 "lack of focus… in supporting ideas".
- prerequisites: W-TR-01.
- mastery_criteria: no large irrelevant/repetitive sections across ≥2 responses.
- common_errors: tangents; padding; repeated points.
- remediation: relevance-editing checklist; cut-irrelevant exercises.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

## W-CC — Coherence & Cohesion

**`W-CC-01` Paragraphing** — writing · CC · bands 4–9 · cognitive: apply · load: medium
- objective: organise writing into logical paragraphs, each with a clear topic.
- traces_to: CC — Band 4 "may be no paragraphing" → Band 7 "paragraphing generally used effectively".
- prerequisites: — .
- mastery_criteria: uses appropriate paragraphing consistently (T2: intro/body/conclusion; T1: logical grouping).
- common_errors: no/misplaced paragraphs; multi-topic paragraphs; missing topic sentences.
- remediation: paragraph-structure templates; topic-sentence practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-CC-02` Overall logical progression** — writing · CC · bands 5–9 · cognitive: apply · load: medium
- objective: arrange information/ideas in a logical sequence with clear overall progression.
- traces_to: CC — Band 5 "lack of overall progression" → Band 7 "clear progression throughout".
- prerequisites: W-CC-01.
- mastery_criteria: response shows clear logical progression throughout.
- common_errors: disordered ideas; no overall flow; abrupt transitions.
- remediation: outlining before writing; sequencing exercises.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-CC-03` Cohesive devices (linkers)** — writing · CC · bands 5–9 · cognitive: apply · load: medium
- objective: use a range of cohesive devices accurately, without over/under-use or mechanical patterns.
- traces_to: CC — Band 5 "limited/overuse of cohesive devices" → Band 7 "used flexibly".
- prerequisites: `K-GRA` conjunctions; W-CC-02.
- mastery_criteria: uses cohesive devices flexibly and accurately; no mechanical/overused linking.
- common_errors: overuse of "firstly/secondly"; mechanical linkers; wrong connector meaning.
- remediation: linker-meaning matching; varied-transition practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-CC-04` Reference & substitution** — writing · CC · bands 6–9 · cognitive: apply · load: medium
- objective: use reference (pronouns, determiners) and substitution to link ideas without repetition.
- traces_to: CC — Band 6 "reference may lack flexibility" → Band 7 "reference and substitution used flexibly".
- prerequisites: `K-GRA` reference; W-CC-02.
- mastery_criteria: uses reference/substitution to avoid repetition with minimal error.
- common_errors: ambiguous reference; noun repetition; faulty pronoun agreement.
- remediation: pronoun/reference drills; rewrite-to-avoid-repetition exercises.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

## W-LR — Lexical Resource

**`W-LR-01` Topic-specific vocabulary** — writing · LR · bands 4–9 · cognitive: apply · load: medium
- objective: use adequate, relevant vocabulary for the task topic.
- traces_to: LR — Band 4 "vocabulary is basic" → Band 6 "generally adequate and appropriate".
- prerequisites: `K-VOC` topic word lists.
- mastery_criteria: uses sufficient topic-appropriate vocabulary; meaning clear.
- common_errors: generic vocabulary; wrong word choice; limited range.
- remediation: topic word-list study; collocate-and-use practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-LR-02` Collocation** — writing · LR · bands 6–9 · cognitive: apply · load: high
- objective: use accurate collocations (word partnerships).
- traces_to: LR — Band 7 "awareness of style and collocation" → Band 8 "despite occasional inaccuracies in… collocation".
- prerequisites: W-LR-01; `K-VOC` collocations.
- mastery_criteria: uses collocations accurately; few inappropriate pairings.
- common_errors: unnatural word pairings; L1-transfer collocations; register clashes.
- remediation: collocation sets; corpus-informed practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-LR-03` Paraphrase** — writing · LR · bands 5–9 · cognitive: create · load: high
- objective: paraphrase to avoid repetition and to reword prompts/data.
- traces_to: LR/CC — avoidance of repetition (Band 6+ resource adequacy).
- prerequisites: W-LR-01.
- mastery_criteria: paraphrases successfully throughout; minimal verbatim copying of the prompt.
- common_errors: copying prompt verbatim; awkward synonym swaps; meaning change when paraphrasing.
- remediation: paraphrase technique (synonym + structure change); prompt-rewording drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-LR-04` Less common / precise / idiomatic vocabulary** — writing · LR · bands 7–9 · cognitive: apply · load: high
- objective: use less common, precise, and (appropriately) idiomatic vocabulary.
- traces_to: LR — Band 7 "some ability to use less common/idiomatic items" → Band 8 "skilful use".
- prerequisites: W-LR-01.
- mastery_criteria: uses some less common/idiomatic items accurately.
- common_errors: forced/unnatural idioms; misused "big words"; inappropriate register.
- remediation: high-utility academic/less-common lexical sets; appropriateness checks.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-LR-05` Spelling & word formation** — writing · LR · bands 4–9 · cognitive: remember · load: medium
- objective: spell accurately and form words correctly.
- traces_to: LR — spelling/word-formation across bands 4–9.
- prerequisites: `K-VOC`.
- mastery_criteria: few spelling/word-formation errors; none that impede communication.
- common_errors: high-frequency spelling errors; wrong word class (e.g., noun↔verb suffixes).
- remediation: spelling lists by error pattern; word-formation drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

## W-GRA — Grammatical Range & Accuracy

**`W-GRA-01` Simple sentence accuracy** — writing · GRA · bands 4–9 · cognitive: apply · load: medium
- objective: produce accurate simple sentences.
- traces_to: GRA — Band 4 "simple sentences predominate" → Band 9 full control.
- prerequisites: `K-GRA` basic sentence structure.
- mastery_criteria: simple sentences largely error-free.
- common_errors: subject–verb agreement; word order; missing verbs.
- remediation: sentence-core drills; error-correction sets.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-GRA-02` Compound sentences** — writing · GRA · bands 5–9 · cognitive: apply · load: medium
- objective: produce accurate compound sentences (coordination).
- traces_to: GRA — sentence-form range.
- prerequisites: W-GRA-01; `K-GRA` coordination.
- mastery_criteria: uses compound sentences correctly.
- common_errors: comma splices; run-ons; faulty coordinating conjunctions.
- remediation: coordination drills; run-on correction.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-GRA-03` Complex sentences (subordination)** — writing · GRA · bands 5–9 · cognitive: apply · load: high
- objective: produce a variety of accurate complex sentences (relative, conditional, concession, noun clauses, etc.).
- traces_to: GRA — Band 5 "complex sentences attempted but faulty" → Band 7 "a variety of complex structures".
- prerequisites: W-GRA-01; `K-GRA` subordination.
- mastery_criteria: uses ≥3 distinct complex-structure types accurately in a 250-word essay; ≤2 structural errors.
- common_errors: run-ons; faulty subordinators; overuse of "which"; sentence fragments.
- remediation: sentence-combining drills; structure-type rotation practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-GRA-04` Tense & aspect** — writing · GRA · bands 4–9 · cognitive: apply · load: medium
- objective: use appropriate tense and aspect accurately.
- traces_to: GRA — grammatical accuracy.
- prerequisites: `K-GRA` tense/aspect.
- mastery_criteria: tense/aspect appropriate and accurate throughout.
- common_errors: tense shifting; wrong aspect; past/present confusion in data description.
- remediation: tense-for-purpose drills (e.g., past for Task 1 data vs present for essays).
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-GRA-05` Articles & determiners** — writing · GRA · bands 4–9 · cognitive: apply · load: high
- objective: use articles and determiners accurately.
- traces_to: GRA — grammatical accuracy.
- prerequisites: `K-GRA` articles/determiners.
- mastery_criteria: article/determiner errors do not impede meaning.
- common_errors: a/an/the omission or misuse; zero-article errors; countable/uncountable confusion.
- remediation: article-rule drills; error-pattern correction (esp. common L1-transfer errors).
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-GRA-06` Punctuation** — writing · GRA · bands 4–9 · cognitive: apply · load: medium
- objective: punctuate accurately.
- traces_to: GRA — punctuation across bands.
- prerequisites: — .
- mastery_criteria: punctuation generally well-controlled; rarely impedes communication.
- common_errors: comma splices; missing commas; faulty capitalization/apostrophes.
- remediation: punctuation-rule drills; editing practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`W-GRA-07` Structural flexibility / variety** — writing · GRA · bands 7–9 · cognitive: create · load: high
- objective: use a wide, flexible range of structures.
- traces_to: GRA — Band 7 "a variety of complex structures… some flexibility" → Band 9 "full flexibility and control".
- prerequisites: W-GRA-02, W-GRA-03.
- mastery_criteria: uses varied structures flexibly; majority of sentences error-free.
- common_errors: repetitive sentence openings/lengths; over-reliance on one structure; flexibility without accuracy.
- remediation: sentence-variety transformation drills; stylistic editing.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

---

## Coverage check (every band requirement maps to leaves)
- **TA/TR** → W-TA-01..03, W-TR-01..04. · **CC** → W-CC-01..04. · **LR** → W-LR-01..05. · **GRA** → W-GRA-01..07.
Each official criterion is fully covered by its leaves; each leaf traces to a descriptor band range and satisfies the leaf invariant (atomic + independently teachable/practiceable/assessable/remediable). *(Self-check vs [../bands/writing.md](../bands/writing.md).)*

## Dependencies
- **Consumes:** [../bands/writing.md](../bands/writing.md).
- **References `K-VOC` / `K-GRA`** = [../knowledge/](../knowledge/) items (built later).
- **Feeds:** [../curriculum/](../curriculum/), [../practice/](../practice/) (populates `practice_item_types`), [../assessment/](../assessment/) (populates `assessment_strategy`).

## Open questions
- [ ] Confirm per-leaf `cognitive_level` (Bloom) and `typical_learning_load` assignments are reasonable starting heuristics.
- [ ] Validate that no leaf should split further (e.g., W-GRA-03 complex sentences by structure type) — currently atomic per the stop condition.
