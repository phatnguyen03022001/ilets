STATUS: CANONICAL
OWNS: how capability is trained, six learning phases, Practice Type taxonomy, type-level phase/mode/scope/binding semantics, and feedback timing for learning
DEPENDS_ON: 03-SKILLS.md, 04-KNOWLEDGE.md, 06-CURRICULUM.md
DOES_NOT_OWN: Skill/Knowledge definitions, band thresholds, assessment sufficiency/certification, learner-state transitions, or concrete generated exercise instances

# 07 — Practice

## Purpose

Define **how canonical capability is trained**.

Practice is not Assessment. A learner performance may sometimes later yield valid evidence, but `08-ASSESSMENT.md` alone owns whether evidence is sufficient for mastery/certification.

Practice Types are reusable training patterns. Concrete items are instances under `10-CONTENT-MODEL.md`.

# Canonical learning phases

Every Practice Type has exactly one primary phase and may support additional phases.

- **Acquisition** — initial construction of a new capability/knowledge pattern; high scaffolding and prompt corrective feedback are often appropriate.
- **Consolidation** — stabilize newly acquired capability through focused varied repetition and correction.
- **Retrieval** — actively recall/reproduce after delay before feedback is revealed.
- **Transfer** — apply learning in a new, less scaffolded, or more integrated context.
- **Fluency** — increase speed/automaticity/rhythm while preserving quality.
- **Exam Readiness** — perform under IELTS-like timing, integration, stamina, and independence constraints.

Exam Readiness is a practice phase. Exam Preparation is a learner mode owned by `09-PROGRESSION.md` and remains non-certifying by exposure alone.

# Feedback timing

- Acquisition: generally immediate, specific, actionable.
- Consolidation: prompt feedback with increasing self-correction/retrieval first.
- Retrieval: require an authentic retrieval attempt before feedback.
- Transfer: increasingly delayed/batched where immediate intervention would disrupt independence.
- Fluency: avoid excessive interruption; emphasize post-attempt patterns.
- Exam Readiness: preserve exam-like independence during the attempt; feedback primarily post-task.

The goal is effective timing for the learning stage, not universal immediacy.

# Practice Type canonical fields

Each Practice Type owns stable ID/name, primary phase, supported phases, mode, applies-to scope, core learning objective, canonical target bindings, reinforcing Knowledge Objects where relevant, and feedback pattern.

# Canonical Practice Type registry

Abbreviations: `Aq` Acquisition, `Co` Consolidation, `Re` Retrieval, `Tr` Transfer, `Fl` Fluency, `ER` Exam Readiness.

| ID | Type | Primary / supported phases | Mode / scope | Canonical targets / reinforcement | Feedback pattern |
|---|---|---|---|---|---|
| `PT-01` | Scaffolded/guided writing | Aq / Aq, Co | individual, adaptive / writing | Writing acquisition, especially `W-GRA-01`, `W-GRA-02`, `W-GRA-03`, `W-CC-01`, `W-CC-02`, `W-TA-01`, `W-TR-01`; reinforces relevant `K-GRA-*`, `K-VOC-*` objects | immediate, scaffolded |
| `PT-02` | Sentence-combining drill | Aq / Aq, Fl | individual / writing | `W-GRA-02`, `W-GRA-03`, `W-GRA-07`; `K-GRA-003`, `K-GRA-004`, `K-GRA-020`, `K-GRA-021` | immediate |
| `PT-03` | Paraphrase task | Co / Co, Tr | individual / writing | `W-LR-03`, `W-CC-04`; `K-VOC-020`, `K-VOC-041` | prompt, then increasingly self-check |
| `PT-04` | Error-correction exercise | Co / Co, Re | individual / writing | Writing grammar/lexical accuracy leaves such as `W-GRA-01`–`W-GRA-07`, `W-LR-05`; relevant grammar/spelling objects | immediate after learner diagnosis attempt |
| `PT-05` | Redraft after feedback | Tr / Co, Tr | individual / writing | integrated Writing capability across Task Achievement/Response, Coherence/Cohesion, Lexical Resource, Grammar | feedback on first draft, independent application in redraft |
| `PT-06` | Timed Writing task | ER / ER, Tr | timed / writing | integrated Academic Writing Task 1/Task 2 capability | post-task |
| `PT-07` | Pronunciation/minimal-pair drill | Aq / Aq, Re | individual / speaking | `S-P-01`, `S-P-02`; `K-PHON-010`, `K-PHON-011`, `K-PHON-012` | immediate |
| `PT-08` | Shadowing | Fl / Aq, Fl | individual / speaking | `S-P-03`, `S-P-04`, `S-FC-01`; `K-PHON-040`, `K-PHON-041` | light immediate support in acquisition; pattern review after fluent attempts |
| `PT-09` | Long-turn practice | Co / Co, Tr, Fl | timed / speaking | `S-FC-02`, `S-FC-04`, `S-GRA-02`, `S-GRA-03`, `S-LR-02` | calibrated, mostly post-turn |
| `PT-10` | Q&A / role-play | Co / Co, Tr | mixed / speaking | `S-FC-03`, `S-FC-04`, `S-FC-05`, `S-LR-01`, `S-LR-04`, `S-GRA-04` | conversational; avoid breaking every utterance |
| `PT-11` | Timed mock Speaking | ER / ER | timed / speaking | integrated Speaking capability across Parts 1, 2, and 3 | post-task |
| `PT-12` | Skim/scan/gist-detail speed drill | Fl / Aq, Fl | timed / listening, reading | `L-COMP-01`, `L-COMP-02`, `R-COMP-01`, `R-COMP-02`; may reinforce `K-VOC-010` | immediate on objective result, strategy review after set |
| `PT-13` | Comprehension question set | Co / Co, Re | individual / listening, reading | receptive comprehension/question-type leaves including `L-COMP-*`, `L-QT-*`, `R-COMP-*`, `R-QT-*`; relevant vocabulary | immediate for objective items after response |
| `PT-14` | Note-taking from lecture/text | Tr / Co, Tr | individual / listening, reading | `L-COMP-06`, `R-COMP-05`, `R-COMP-06`; `K-VOC-011` | post-attempt comparison/feedback |
| `PT-15` | Timed section/passage practice | ER / ER, Fl | timed / listening, reading | integrated receptive leaves under section/passage timing | post-section/passage |
| `PT-16` | Distractor/error review | Re / Re, Co | individual / listening, reading | `L-COMP-05`, `R-QT-02`, `R-QT-03` and related reasoning errors | explanation after learner justification attempt |
| `PT-17` | Spaced retrieval | Re / Re, Co | adaptive / knowledge + dependent skills | relevant `K-VOC-*` / `K-GRA-*` objects and their dependent Skill Leaves | immediate after retrieval attempt |
| `PT-18` | Collocation/word-formation drill | Aq / Aq, Co | individual / knowledge, writing | `W-LR-02`, `W-LR-05`; `K-VOC-020`, `K-VOC-030` | immediate |
| `PT-19` | Gap-fill/completion | Co / Co, Re | individual / knowledge, listening, reading | `W-LR-05`, `L-QT-01`, `L-QT-05`, `R-QT-05`; relevant grammar/vocabulary objects | immediate after attempt |
| `PT-20` | Interleaved mixed set | Tr / Tr, Re | mixed / cross | any explicitly selected canonical target IDs | item/set feedback without destroying discrimination demand |
| `PT-21` | Adaptive practice set | Co / Co, Re, Tr | adaptive / cross | any explicitly selected targets from learner state | adaptive; timing follows selected type/phase |
| `PT-22` | Diagnostic checkpoint practice | Re / Re, Co | diagnostic / cross | explicitly sampled Skill/Knowledge targets used to expose gaps | feedback after sample; diagnostic measurement semantics, if used, belong to `AT-04` |
| `PT-23` | Full mock test | ER / ER | timed / cross | integrated all four IELTS skills under exam-like conditions | post-test only |

Wildcards such as `L-COMP-*` above denote a target family, not a stable object ID. Any concrete binding must resolve the family to explicit canonical IDs before execution or evidence recording.

## Phase coverage invariant

- Acquisition: `PT-01`, `PT-02`, `PT-07`, `PT-18`.
- Consolidation: `PT-03`, `PT-04`, `PT-09`, `PT-10`, `PT-13`, `PT-19`, `PT-21`.
- Retrieval: `PT-16`, `PT-17`, `PT-22`.
- Transfer: `PT-05`, `PT-14`, `PT-20`.
- Fluency: `PT-08`, `PT-12`.
- Exam Readiness: `PT-06`, `PT-11`, `PT-15`, `PT-23`.

Coverage does not imply equal practice volume; selection follows learner need.

# Binding model

```text
Curriculum Node
      ↓
Skill Leaf / Knowledge Object target
      ↓
Practice Type
      ↓
Concrete Practice Item
```

Rules:

1. bind by stable IDs;
2. Practice Types never copy Skill/Knowledge definitions;
3. a node may use multiple Practice Types across phases;
4. a Practice Type may support multiple canonical objects;
5. concrete items reference both type and explicit target IDs;
6. adaptive selection may vary order/frequency/difficulty but may not remove required target capability.

`PT-22` and `AT-04` are intentionally different: the former is a reusable **practice/sampling format**; the latter is the **measurement interpretation** of diagnostic evidence.

# Difficulty progression

Valid difficulty dimensions include scaffold amount, lexical/syntactic complexity, distractor quality, integration, time pressure, novelty/transfer distance, response length, and cue availability. Harder practice is not automatically higher-band truth; Band ownership remains in `05-BANDS.md`.

# Deliberate-practice and AI invariants

Practice should make the target capability, success condition, error pattern, and intended next change understandable to the learner. Raw repetition volume is not a learning strategy by itself.

AI may generate examples, explain errors, provide feedback, and select practice, but must not perform the cognitive operation on the learner's behalf when that would defeat retrieval, reasoning, planning, or independent production.

Concrete prompts/items/examples/remediation messages belong to `10-CONTENT-MODEL.md`, not this taxonomy.