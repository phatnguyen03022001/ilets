STATUS: CANONICAL
OWNS: how capability is trained, learning phases, Learning Mechanism taxonomy, Practice Type taxonomy, type-level phase/mode/scope/binding semantics, and feedback timing for learning
DEPENDS_ON: 03-SKILLS.md, 04-KNOWLEDGE.md, 06-CURRICULUM.md, 09-PROGRESSION.md
DOES_NOT_OWN: Skill/Knowledge definitions, band thresholds, assessment sufficiency/certification, learner-state transitions, or concrete generated exercise instances

# 07 — Practice

## Purpose

Define how canonical capability is trained.

Practice is selected because of a diagnosed learning need, not because a learner merely belongs to a band. The canonical selection chain is:

```text
GapEvaluation / ActionIntent     owned by 09-PROGRESSION
            ↓
Learning Mechanism              owned here
            ↓
Practice Type                   owned here
            ↓
Concrete Practice Item          owned by 10-CONTENT-MODEL
```

A Learning Mechanism is a learning process. A Practice Type is an executable activity pattern that may instantiate one or more mechanisms. They are not interchangeable.

Practice is not Assessment. A performance may later yield admissible evidence, but `08-ASSESSMENT.md` alone decides whether an observation becomes evidence and whether the evidence supports a claim.

# Canonical learning phases

Every Practice Type has one primary phase and may support additional phases.

- **Acquisition** — initial construction of a new capability or knowledge pattern.
- **Consolidation** — stabilize newly acquired capability through focused varied use and correction.
- **Retrieval** — actively recall or reproduce after delay before feedback.
- **Transfer** — apply learning in a new, less scaffolded, or more integrated context.
- **Fluency** — increase speed, automaticity, rhythm, or processing efficiency without sacrificing target quality.
- **Exam Readiness** — perform under IELTS-like timing, integration, stamina, and independence constraints.

Exam Readiness is a practice phase. Exam Preparation is a learner mode owned by `09-PROGRESSION.md`.

# Canonical Learning Mechanisms

| ID | Mechanism | Use when | Do not infer |
|---|---|---|---|
| `LM-01` | Worked example | learner needs a model of successful performance or attention guidance | copying a model proves independent performance |
| `LM-02` | Active retrieval | a reviewable target should be recalled or reproduced from memory | every target belongs in a flashcard scheduler |
| `LM-03` | Spaced review | a suitable reviewable target has repeated retrieval history | one fixed spacing formula fits all learning objects |
| `LM-04` | Contrast / discrimination | learner confuses near alternatives or misses a discriminating cue | random variety is useful by itself |
| `LM-05` | Controlled production | recognition exists but production is unstable | constrained success proves free production |
| `LM-06` | Guided production | learner needs failure-relevant support to produce the target | support should remain indefinitely |
| `LM-07` | Scaffold fading | performance depends materially on a support that should be removed | all scaffolds should disappear at once |
| `LM-08` | Variation / transfer | learner must generalize beyond the practiced context | novelty alone proves transfer |
| `LM-09` | Interleaving | discrimination or flexible selection benefits from mixing sufficiently stable targets | arbitrary mixing always improves learning |
| `LM-10` | Self-explanation | reasoning or discrimination becomes more observable through explanation | every item needs a reflection prompt |
| `LM-11` | Deliberate revision / reproduction | corrected productive output has learning value | every error requires rewrite/re-record |
| `LM-12` | Fluency rehearsal | target quality is sufficiently stable and efficiency/rhythm is limiting performance | speed compensates for weak underlying quality |

Mechanism selection must follow the GapEvaluation and ActionIntent. A wrong answer is not automatically a remediation command; insufficient evidence, stale evidence, prerequisite failure, transfer failure, and scaffold dependence require different actions.

AI tutor behavior is delivery/runtime capability, not a Learning Mechanism. AI may instantiate mechanisms but must not replace the learner's required cognitive operation.

# Feedback timing

- Acquisition: generally immediate, specific, actionable.
- Consolidation: prompt feedback with increasing self-correction before reveal.
- Retrieval: require an authentic retrieval attempt before feedback.
- Transfer: increasingly delayed/batched when immediate intervention would destroy independence.
- Fluency: avoid excessive interruption; review patterns after the attempt.
- Exam Readiness: preserve exam-like independence during the attempt; feedback primarily post-task.

Feedback depth follows failure type. Do not correct every detectable issue when doing so would overload the learner or derail the current learning objective.

# Canonical Practice Type registry

Abbreviations: `Aq` Acquisition, `Co` Consolidation, `Re` Retrieval, `Tr` Transfer, `Fl` Fluency, `ER` Exam Readiness.

| ID | Type | Primary / supported phases | Mode / scope | Typical mechanisms | Canonical target role |
|---|---|---|---|---|---|
| `PT-01` | Scaffolded/guided writing | Aq / Aq, Co | individual, adaptive / writing | `LM-01`, `LM-06`, `LM-07` | Writing acquisition and controlled support |
| `PT-02` | Sentence-combining drill | Aq / Aq, Fl | individual / writing | `LM-05`, `LM-12` | grammatical production/control |
| `PT-03` | Paraphrase task | Co / Co, Tr | individual / writing | `LM-04`, `LM-05`, `LM-08` | lexical/cohesive flexibility |
| `PT-04` | Error-correction exercise | Co / Co, Re | individual / writing | `LM-02`, `LM-04`, `LM-10` | diagnosis and correction of language errors |
| `PT-05` | Redraft after feedback | Tr / Co, Tr | individual / writing | `LM-11`, `LM-07`, `LM-08` | apply feedback in new production |
| `PT-06` | Timed Writing task | ER / ER, Tr | timed / writing | `LM-08` | integrated Writing under independence/timing |
| `PT-07` | Pronunciation/minimal-pair drill | Aq / Aq, Re | individual / speaking | `LM-04`, `LM-02` | pronunciation discrimination/production |
| `PT-08` | Shadowing | Fl / Aq, Fl | individual / speaking | `LM-01`, `LM-12` | only pronunciation/prosody/fluency targets where imitation is relevant |
| `PT-09` | Long-turn practice | Co / Co, Tr, Fl | timed / speaking | `LM-06`, `LM-07`, `LM-08`, `LM-12` | sustained Speaking production |
| `PT-10` | Q&A / role-play | Co / Co, Tr | mixed / speaking | `LM-05`, `LM-06`, `LM-08` | responsive Speaking production |
| `PT-11` | Timed mock Speaking | ER / ER | timed / speaking | `LM-08` | integrated Speaking under exam-like conditions |
| `PT-12` | Skim/scan/gist-detail speed drill | Fl / Aq, Fl | timed / listening, reading | `LM-04`, `LM-12` | selective receptive processing |
| `PT-13` | Comprehension question set | Co / Co, Re | individual / listening, reading | `LM-02`, `LM-04` | receptive comprehension/question-type work |
| `PT-14` | Note-taking from lecture/text | Tr / Co, Tr | individual / listening, reading | `LM-08`, `LM-10` | information selection and integration |
| `PT-15` | Timed section/passage practice | ER / ER, Fl | timed / listening, reading | `LM-08`, `LM-12` | receptive work under section timing |
| `PT-16` | Distractor/error review | Re / Re, Co | individual / listening, reading | `LM-04`, `LM-10` | discriminate why an alternative failed |
| `PT-17` | Spaced retrieval | Re / Re, Co | adaptive / suitable reviewable targets | `LM-02`, `LM-03` | retention of reviewable Knowledge/Skill targets |
| `PT-18` | Collocation/word-formation drill | Aq / Aq, Co | individual / knowledge, writing | `LM-04`, `LM-05` | lexical production and discrimination |
| `PT-19` | Gap-fill/completion | Co / Co, Re | individual / knowledge, receptive skills | `LM-02`, `LM-05` | constrained retrieval/application |
| `PT-20` | Interleaved mixed set | Tr / Tr, Re | mixed / cross | `LM-09`, `LM-04`, `LM-08` | flexible selection across targets |
| `PT-21` | Adaptive practice set | Co / Co, Re, Tr | adaptive / cross | depends on selected GapEvaluation | container for evidence-based activity selection |
| `PT-22` | Diagnostic checkpoint practice | Re / Re, Co | diagnostic / cross | depends on uncertainty sampled | low-friction sampling; measurement interpretation belongs to `AT-04` |
| `PT-23` | Full mock test | ER / ER | timed / cross | none as a default learning mechanism | broad exam-readiness / re-evidence activity with learning side-effects |

Concrete bindings must resolve target families to explicit stable Skill/Knowledge IDs before execution or evidence recording.

# Selection rules

1. **Ability gap** — choose a mechanism that targets the demonstrated failure pattern; do not default to reteaching prerequisites.
2. **Prerequisite gap** — acquire or repair the required Knowledge/Skill prerequisite before dependent work when the dependency is truly Required.
3. **Insufficient evidence** — use a low-friction diagnostic or assessment action; do not label the learner weak.
4. **Conflicting evidence** — choose a discriminating task that can separate plausible explanations.
5. **Stale evidence** — collect the smallest useful representative refresh before remediation.
6. **Scaffold dependence** — identify the support carrying performance, fade it selectively, then re-evidence.
7. **Transfer gap** — vary context, reduce pattern repetition, and use unseen or materially different conditions when the claim requires generalization.
8. **Fluency gap** — use rehearsal only after target quality is sufficiently stable.
9. **Exam-condition gap** — use timed/integrated practice without redefining underlying mastery.

# Binding model

```text
Curriculum Node
      ↓
Skill Leaf / Knowledge Object
      ↓
GapEvaluation + ActionIntent
      ↓
Learning Mechanism
      ↓
Practice Type
      ↓
Concrete Practice Item
```

Practice Types reference canonical targets; they do not copy Skill/Knowledge definitions. A node may use multiple Practice Types. A Practice Type may support multiple objects. Adaptive selection may vary order, mechanism, type, difficulty, scaffold, and frequency while preserving the target standard.

# Difficulty and scaffolding

Difficulty may vary through scaffold amount, linguistic complexity, distractor quality, integration, time pressure, novelty/transfer distance, response length, and cue availability.

Harder is not automatically higher-band truth. Band ownership remains in `05-BANDS.md`.

Scaffolding must be represented explicitly enough that guided success cannot later masquerade as independent evidence. The concrete `ScaffoldingProfile` representation is owned by `10-CONTENT-MODEL.md`.

# Practice invariants

- practice completion is not mastery;
- same-item retry is primarily recovery/learning, not automatic transfer evidence;
- shadowing is used only where pronunciation/prosody/fluency makes it relevant;
- dictation-like activity is used only when decoding/segmentation/detail perception is implicated;
- full mocks are broad readiness/evidence activities, not substitutes for targeted intervention;
- learner preference may influence suitable options but cannot override target requirements;
- repeated learner friction or skipping is a planning signal, not ability evidence;
- mechanism diversity is justified by learning need, diminishing returns, or transfer—not by a novelty quota.
