# IELTS-EXAM-BANDS IELTS external exam and Band semantics

> **CANONICAL PRODUCT AUTHORITY**
>
> Canonical within the PRODUCT authority domain under `docs/catalog/project.json`. `CONSTITUTION.md` and `OBJECTIVE.md` retain distinct authority, and `contracts/**` retains scoped exact machine-contract authority. References below to `spec/**` or `design/**`, including historical uses of “canonical”, are pre-cutover provenance only and do not create equal authority.

## External IELTS construct

The standard construct has two variants: IELTS Academic and IELTS General Training. Listening and Speaking are shared; Reading and Writing differ where the external exam differs. Delivery mode is orthogonal and is not another IELTS variant.

IELTS reports four section Bands: Listening, Reading, Writing and Speaking. Overall Band is the arithmetic mean of those four scores rounded to the nearest whole or half Band; `.25` rounds up to the next half Band and `.75` to the next whole Band. An overall result does not imply uniform underlying capability.

### Stable official-family identities

These IDs identify external exam families, not Skills, Practice Types, product features, or wire enums. They remain stable while the corresponding external family remains materially the same; coverage tracks them separately from Skill identity.

| ID | External family |
|---|---|
| `IELTS-L-QF-01` | Listening — Multiple choice |
| `IELTS-L-QF-02` | Listening — Matching |
| `IELTS-L-QF-03` | Listening — Plan / map / diagram labelling |
| `IELTS-L-QF-04` | Listening — Form / note / table / flow-chart / summary completion |
| `IELTS-L-QF-05` | Listening — Sentence completion |
| `IELTS-L-QF-06` | Listening — Short-answer questions |
| `IELTS-R-QF-01` | Reading — Multiple choice |
| `IELTS-R-QF-02` | Reading — True / False / Not Given |
| `IELTS-R-QF-03` | Reading — Yes / No / Not Given |
| `IELTS-R-QF-04` | Reading — Matching information |
| `IELTS-R-QF-05` | Reading — Matching headings |
| `IELTS-R-QF-06` | Reading — Matching features |
| `IELTS-R-QF-07` | Reading — Matching sentence endings |
| `IELTS-R-QF-08` | Reading — Sentence completion |
| `IELTS-R-QF-09` | Reading — Summary / note / table / flow-chart completion |
| `IELTS-R-QF-10` | Reading — Diagram label completion |
| `IELTS-R-QF-11` | Reading — Short-answer questions |
| `IELTS-W-A-T1` | Academic Writing Task 1 — visual information |
| `IELTS-W-GT-T1` | General Training Writing Task 1 — letter |
| `IELTS-W-T2` | Writing Task 2 — essay response; shared learning core, concrete prompt remains variant-scoped |
| `IELTS-S-P1` | Speaking Part 1 — interview/familiar topics |
| `IELTS-S-P2` | Speaking Part 2 — individual long turn |
| `IELTS-S-P3` | Speaking Part 3 — extended discussion |

#### Canonical materializer anchor — Reading classification families

| ID | Official family |
|---|---|
| `IELTS-R-QF-02` | Identifying information — True / False / Not Given |
| `IELTS-R-QF-03` | Identifying writer views/claims — Yes / No / Not Given |

#### Canonical materializer anchor — Reading matching headings family

| ID | Official family |
|---|---|
| `IELTS-R-QF-05` | Reading — Matching headings |

A Skill Leaf may serve several official families and one family may require several Skills. Academic Task-1 visual presentation subformats remain content-presentation coverage dimensions, not extra scored Writing tasks. Speaking remains a holistic Band construct even though Parts 1–3 have stable family identities for coverage.

## External-truth epistemic contract

Every statement that carries external IELTS or Band meaning uses the strongest class its provenance actually supports:

- **Official Evidence** — a fact, descriptor, scoring rule/guidance, format requirement, or other external truth directly established by an authoritative IELTS source.
- **Evidence-Based Interpretation** — a close educational interpretation grounded in official or other material evidence, but not itself a direct official IELTS statement.
- **Blueprint Inference** — pedagogical detail introduced by this product blueprint beyond what IELTS explicitly publishes.

Evidence-Based Interpretation and Blueprint Inference must never be presented as direct Official Evidence. Mutable Official Evidence keeps dated provenance so the product can distinguish a preserved historical baseline from a currently rechecked external fact.

### Owner-first external IELTS change propagation

When authoritative external IELTS truth changes:

1. update this external-truth owner first;
2. refresh the dated provenance/evidence for the changed fact;
3. review every affected downstream Skill, Band, Curriculum, Assessment, Content, coverage, behavior, and system-design consequence before changing derived behavior;
4. preserve a stable official-family ID only when the external family remains materially the same; and
5. never preserve an obsolete external fact merely because implementation, content, configuration, or another downstream artifact encoded it.

Downstream owners may interpret or operationalize external truth within their scope, but they do not override the external-truth owner or silently redefine IELTS competence.

### Delivery and administrative baseline

External facts were last rechecked by the legacy owner against official IELTS material on 2026-08-22. The successor representation preserves that locked baseline; live external facts must be rechecked before implementation/release.

- Test-centre computer is the default current standard-delivery baseline. Listening, Reading and Writing use computer interaction; Speaking remains a human-examiner interaction. Delivery does not change the language construct or Band interpretation.
- Selected markets may offer Writing on Paper inside IELTS on computer. Listening/Reading remain computer-delivered, Writing tasks are shown on screen and answers are handwritten. Task types, timing, scoring and criteria stay unchanged. It is not currently offered for IELTS for UKVI, and a Writing One Skill Retake uses the same Writing delivery mode as the original eligible test where that route applies.
- IELTS Online is Academic-only in supported countries. Listening/Reading/Writing are remote; Speaking is a secure live examiner video call. Receiving organisations decide acceptance; the route is not currently for visa/immigration purposes and General Training is not offered as IELTS Online.
- UKVI Academic/GT reuses the corresponding language construct while adding administrative/security/location constraints; those constraints do not create another curriculum or Band definition.
- One Skill Retake reuses one existing section. Current locked conditions include: original full eligible IELTS-on-computer test at a participating centre; one skill may be retaken once per original full test; retake within 60 days; availability varies by location; receiving organisations may impose acceptance requirements; current UKVI acceptance is available for Academic/GT computer tests in selected locations; Writing delivery must match the original eligible Writing mode where applicable.

### Timing, section structure and scoring facts

The complete test is approximately 2h45m: Listening about 30m, Reading 60m, Writing 60m, Speaking about 11–14m.

Listening has 40 questions in four parts, 10 questions each; recordings are heard once. Context progresses from everyday/social conversation, everyday/social monologue, educational/training conversation, to academic monologue. Each correct answer earns one mark and the raw /40 converts to Band. Current average anchors are Band 5≈16, 6≈23, 7≈30, 8≈35; exact marks may vary slightly by test version, so these are not a permanent conversion table.

Reading has 40 questions in 60m. Academic uses three academically oriented passages. Current average Academic anchors are Band 5≈15, 6≈23, 7≈30, 8≈35. General Training uses Section 1 everyday texts, Section 2 workplace texts, and Section 3 a longer general-interest text, with current total text guidance about 2150–2750 words; current average GT anchors are Band 4≈15, 5≈23, 6≈30, 7≈35. Academic and GT conversions are not interchangeable.

Writing has two tasks in 60m. The four criteria are Task Achievement (Task 1) / Task Response (Task 2), Coherence and Cohesion, Lexical Resource, and Grammatical Range and Accuracy. Criteria are equally weighted within each task; Task 2 carries twice Task-1 weight. Academic Task 1 is visual information, at least 150 words, roughly 20m. GT Task 1 is a letter, at least 150 words, roughly 20m, with three supplied content functions/bullet points and relationship-sensitive personal/semi-formal/formal style. Task 2 is an essay response, at least 250 words, roughly 40m.

Speaking is a recorded human-examiner interaction shared by Academic/GT. Its criteria are Fluency and Coherence, Lexical Resource, Grammatical Range and Accuracy, and Pronunciation, equally weighted. Delivery may be in person or, for IELTS Online Academic, secure live video; the underlying construct remains the same.

## Bands 3–9 learning thresholds

Bands 3–9 are the structured learning range. Bands 0–2 remain diagnostic/external boundaries without a detailed canonical curriculum. Receptive overlays below are Blueprint Inference because public Listening/Reading scoring is primarily raw-score based; productive thresholds closely interpret published analytic criteria. Band thresholds define required quality, not evidence quantity; Assessment owns sufficiency.

### Band 3 — structured entry

- Listening: extract isolated/simple meaning from predictable everyday speech; recognize familiar words/simple propositions, selected basic details, and begin separating gist from isolated words. Large comprehension gaps, unreliable extended-speech tracking, weak paraphrase/signposting, and academic monologue remain acceptable limitations.
- Reading: locate isolated explicit information in accessible text using basic word matching/scanning and begin identifying a simple overall topic. Reliable inference, integrated main idea, and dense/complex text remain outside the expected threshold.
- Writing: attempt both required tasks at a rateable length, attempt sentence forms, convey partial relevant meaning, and avoid wholly memorized production. Frequent grammar/spelling/organization/vocabulary/task-control problems remain acceptable; Academic Task 1 may only partly select/describe features, GT Task 1 may only partly control recipient/purpose/bullet points/register.
- Speaking: convey basic familiar/personal meaning, attempt simple sentences, minimally connect responses and remain partly communicative despite frequent pauses/language search. Numerous grammar, vocabulary and pronunciation problems remain acceptable.

### Band 4 — basic functional control

- Listening: identify main ideas in simpler recordings, capture common details, complete basic completion/short-answer tasks; paraphrase, distractor handling and dense academic listening remain unreliable.
- Reading: scan details, skim a basic main idea and complete simpler completion/short-answer work; inference, writer stance and dense abstract reading remain outside threshold.
- Writing: minimally address each task, show a discernible Task-2 position, some accurate simple sentences, paragraphing and basic cohesion. Academic Task 1 identifies some key visual features; GT Task 1 recognizes recipient/purpose, attempts required bullet points and some relationship-sensitive style. Frequent errors/repetition remain acceptable if meaning is accessible.
- Speaking: discuss familiar topics basically, connect simple sentences, produce some accurate short utterances and attempt basic stress/intonation while broadly intelligible; long pauses, repetition, frequent errors and pronunciation lapses remain acceptable.

### Band 5 — developing control

- Listening: follow main points/relevant details in familiar social and some educational contexts, recognize common paraphrase, begin managing distractors and a broader question-family range; dense academic detail/complex paraphrase remain unreliable.
- Reading: identify main ideas and detailed explicit information, begin inference/text-structure reasoning and use question-type strategies beyond simple completion; abstract/dense argument and nuanced stance remain unreliable. GT begins transfer across everyday/workplace contexts.
- Writing: address main task requirements though incompletely, state a Task-2 position with limited development, show underlying organization, attempt complex grammar with limited accuracy and minimally adequate task vocabulary. Academic Task 1 attempts an overview and selected features. GT Task 1 identifies purpose, covers major bullet points and uses broadly plausible register, though inconsistency remains acceptable. Noticeable errors/mechanical cohesion/underdevelopment remain compatible.
- Speaking: keep speaking despite repetition/self-correction/slower delivery, discuss familiar/unfamiliar topics with limited lexical flexibility, attempt paraphrase, and control basic forms while complex speech remains error-prone.

### Band 6 — competent clear performance

- Listening: follow main ideas/specific details across most parts, handle common paraphrase/question demands, begin following extended academic speech, and manage many distractors with some inconsistency.
- Reading: understand main ideas and detailed explicit information, manage most official question types, make some reliable inference and engage with denser/longer passages while nuance/abstraction remain harder. GT competence transfers across everyday, workplace and longer general-interest sections.
- Writing Task 2 addresses main prompt parts, presents a relevant position and sufficiently develops main ideas. Shared quality has generally coherent progression, adequate task vocabulary, simple+complex forms, with errors rarely blocking communication. Academic Task 1 covers main requirements, selects relevant visual features, attempts a relevant overview and supports description with data/detail. GT Task 1 correctly identifies recipient/purpose, addresses all bullet points with adequate relevant information, uses generally appropriate register/letter organization and avoids register choices that materially undermine purpose.
- Speaking: discuss topics at length, produce long turns though coherence may occasionally slip, generally paraphrase, use simple+complex structures, some effective stress/intonation, and remain generally understandable.

### Band 7 — good flexible control

- Listening: follow extended speech reliably, handle academic/abstract content, consistently recognize paraphrase/signposting, reject most distractors, missing only occasional nuance/demanding detail.
- Reading: handle extended complex passages reliably, infer meaning/writer stance, understand organization, and consistently execute demanding matching/stance work. GT remains reliable across the full Section-1/2/3 distribution.
- Writing Task 2 presents a clear developed position, extends/supports ideas and maintains relevance/focus. Shared quality has clear logical progression, generally effective paragraphing, flexible cohesion/reference, some less-common/precise vocabulary, varied complex structures and frequent error-free sentences. Academic Task 1 provides a clear overview, highlights key trends/differences/features and organizes detail. GT Task 1 clearly achieves purpose, develops every bullet point, sustains appropriate relationship/register with only occasional lapses, and uses opening/closing/organization suited to the communicative situation.
- Speaking: produce long turns without noticeable effort, develop topics coherently, use discourse markers flexibly, paraphrase effectively, use a useful structural range with frequent error-free sentences, and maintain effective pronunciation despite occasional lapses.

### Band 8 — very strong control

- Listening: understand detailed, complex, abstract speech with little difficulty, manage distractors/paraphrase confidently, and make only rare errors on subtle/dense detail.
- Reading: understand detailed, complex, abstract text with little difficulty; handle inference, stance, structure and demanding question types reliably with rare subtle-argument errors. GT context/register differences no longer create meaningful instability.
- Writing Task 2 addresses the prompt appropriately/sufficiently, presents a clear well-developed position and supports ideas effectively. Shared quality is easy to follow, with well-managed cohesion, wide flexible precise vocabulary, wide structural range, majority error-free sentences and only occasional non-systematic errors. Academic Task 1 satisfies the task appropriately/sufficiently and skilfully selects/highlights/illustrates key features. GT Task 1 fulfils purpose/content appropriately/sufficiently, controls register flexibly/consistently and makes audience-sensitive lexical/tone/organization/opening/closing choices with only occasional non-systematic lapses.
- Speaking: speak fluently with hesitation mostly about content, use wide flexible precise vocabulary and structural range with most sentences error-free, and sustain effective rhythm/stress/intonation/intelligibility.

### Band 9 — ceiling mastery

- Listening: comprehend virtually all spoken material including subtle, abstract, detailed and extended content; misses are exceptional rather than systematic.
- Reading: comprehend virtually all relevant text including subtle argument, abstraction and dense detail; misses are exceptional; Academic/GT context differences do not cause systematic weakness within their constructs.
- Writing Task 2 addresses the prompt in depth, maintains a clear fully developed position and fully/relevantly supports ideas. Shared quality has effortless coherence, sophisticated cohesion/paragraph control, full lexical and structural flexibility/precision, with only extremely rare lapses. Academic Task 1 fully and appropriately satisfies requirements and communicates essential visual features precisely. GT Task 1 fully achieves purpose/all prompt requirements with relationship-appropriate register/tone/organization/lexis and only exceptional lapses.
- Speaking: fluent, coherent, fully developed speech with total/near-total lexical flexibility/precision, highly accurate structural control, flexible connected-speech features, effortless intelligibility and only extremely rare ordinary high-level lapses.

## Variant bindings and exit boundaries

For an Academic Writing Band-N claim, Task 1 uses `W-TA-01`, `W-TA-02`, `W-TA-03` plus applicable shared Writing leaves. For GT, Task 1 uses `W-GT1-01`, `W-GT1-02`, `W-GT1-03` plus shared Writing capability. `W-TA-*` visual leaves are not required for GT Task 1; `W-GT1-*` letter leaves are not required for Academic Task 1. Task 2/shared quality use `W-TR-*`, `W-CC-*`, `W-LR-*`, `W-GRA-*` as applicable.

Skill-level exit semantics are:

- Listening Band N: performance converts to Band N under the applicable official Listening scoring policy, with evidence sufficiency owned by Assessment.
- Academic Reading Band N: performance converts under the Academic Reading policy and sufficiently samples that construct.
- GT Reading Band N: performance converts under the GT policy and sufficiently samples the GT Section-1/2/3 context distribution; Academic conversion/content cannot certify GT Reading.
- Academic Writing Band N: Academic Task 1 + Task 2 + shared Writing criteria all meet Band-N threshold.
- GT Writing Band N: GT Task 1, including recipient/purpose/required-point/register control, + Task 2 + shared criteria all meet Band-N threshold. Task 1 and Task 2 cannot substitute for one another; Academic/GT Task-1 evidence cannot cross-substitute.
- Speaking Band N: Band-N performance across the whole Speaking construct, demonstrated across Parts 1, 2 and 3; part-specific practice does not make the Band standard part-scoped.

A Band specification is a threshold, not an exhaustive study list. A learner at Band N may retain errors explicitly compatible with Band N; higher-band characteristics must not be promoted into mandatory Band-N exit criteria. This preserves the locked residual-error and higher-band exclusion boundary.
