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

Each Practice Type owns: stable ID/name, primary phase, supported phases, mode, applies-to scope, core learning objective, canonical target bindings, reinforcing Knowledge Objects where relevant, and feedback pattern.

# Canonical Practice Type registry

Abbreviations: `Aq` Acquisition, `Co` Consolidation, `Re` Retrieval, `Tr` Transfer, `Fl` Fluency, `ER` Exam Readiness.

| ID | Type | Primary / supported phases | Mode / scope | Canonical targets / reinforcement | Feedback pattern |
|---|---|---|---|---|---|
| `PT-01` | Scaffolded/guided writing | Aq / Aq, Co | individual, adaptive / writing | `W-*` acquisition, esp. GRA/CC/TA/TR; reinforces `K-GRA-*`, `K-VOC-*` | immediate, scaffolded |
| `PT-02` | Sentence-combining drill | Aq / Aq, Fl | individual / writing | `W-GRA-02`, `W-GRA-03`, `W-GRA-07`; `K-GRA-003`, `004`, `020`, `021` | immediate |
| `PT-03` | Paraphrase task | Co / Co, Tr | individual / writing | `W-LR-03`, `W-CC-04`; `K-VOC-020`, `041` | prompt then increasingly self-check |
| `PT-04` | Error-correction exercise | Co / Co, Re | individual / writing | `W-GRA-*`, `W-LR-05`; grammar + spelling knowledge | immediate after learner diagnosis attempt |
| `PT-05` | Redraft after feedback | Tr / Co, Tr | individual / writing | integrated `W-TA/TR/CC/LR/GRA` | feedback on first draft, independent application in redraft |
| `PT-06` | Timed Writing task | ER / ER, Tr | timed / writing | integrated Writing Task 1/2 capability | post-task |
| `PT-07` | Pronunciation/minimal-pair drill | Aq / Aq, Re | individual / speaking | `S-P-01`, `S-P-02`; `K-PHON-010`, `011`, `012` | immediate |
| `PT-08` | Shadowing | Fl / Aq, Fl | individual / speaking | `S-P-03`, `S-P-04`, `S-FC-01`; `K-PHON-040`, `041` | immediate/light during acquisition; pattern review after fluent attempts |
| `PT-09` | Long-turn practice | Co / Co, Tr, Fl | timed / speaking | `S-FC-02`, `S-FC-04`, `S-GRA-02`, `S-GRA-03`, `S-LR-02` | calibrated, mostly post-turn |
| `PT-10` | Q&A / role-play | Co / Co, Tr | mixed / speaking | `S-FC-03`, `04`, `05`, `S-LR-01`, `04`, `S-GRA-04` | conversational; avoid breaking every utterance |
| `PT-11` | Timed mock Speaking | ER / ER | timed / speaking | all `S-*` integrated across test parts | post-task |
| `PT-12` | Skim/scan/gist-detail speed drill | Fl / Aq, Fl | timed / listening, reading | `L-COMP-01`, `02`, `R-COMP-01`, `02`; may reinforce core vocab | immediate on objective result, strategy review after set |
| `PT-13` | Comprehension question set | Co / Co, Re | individual / listening, reading | receptive `*-COMP-*`, `*-QT-*`; relevant vocabulary | immediate for objective items after response |
| `PT-14` | Note-taking from lecture/text | Tr / Co, Tr | individual / listening, reading | `L-COMP-06`, `R-COMP-05`, `R-COMP-06`; academic vocab | post-attempt comparison/feedback |
| `PT-15` | Timed section/passage practice | ER / ER, Fl | timed / listening, reading | integrated receptive leaves | post-section/passage |
| `PT-16` | Distractor/error review | Re / Re, Co | individual / listening, reading | `L-COMP-05`, `R-QT-02`, `R-QT-03` and related error patterns | explanation after learner justification attempt |
| `PT-17` | Spaced retrieval | Re / Re, Co | adaptive / knowledge + dependent skills | `K-VOC-*`, `K-GRA-*` and knowledge-dependent leaves | immediate after retrieval attempt |
| `PT-18` | Collocation/word-formation drill | Aq / Aq, Co | individual / knowledge, writing | `W-LR-02`, `W-LR-05`; `K-VOC-020`, `030` | immediate |
| `PT-19` | Gap-fill/completion | Co / Co, Re | individual / knowledge, listening, reading | `W-LR-05`, `L-QT-01`, `L-QT-05`, `R-QT-05`; relevant grammar/vocab | immediate after attempt |
| `PT-20` | Interleaved mixed set | Tr / Tr, Re | mixed / cross | any selected canonical targets | item/set feedback without destroying discrimination demand |
| `PT-21` | Adaptive practice set | Co / Co, Re, Tr | adaptive / cross | any targets selected from learner state | adaptive; timing follows selected type/phase |
| `PT-22` | Diagnostic checkpoint practice | Re / Re, Co | diagnostic / cross | sampled skills/knowledge to expose gaps | feedback after sample; measurement semantics, if used, belong to `AT-04` |
| `PT-23` | Full mock test | ER / ER | timed / cross | integrated all four skills | post-test only |

## Phase coverage invariant

- Acquisition: `PT-01`, `02`, `07`, `18`.
- Consolidation: `PT-03`, `04`, `09`, `10`, `13`, `19`, `21`.
- Retrieval: `PT-16`, `17`, `22`.
- Transfer: `PT-05`, `14`, `20`.
- Fluency: `PT-08`, `12`.
- Exam Readiness: `PT-06`, `11`, `15`, `23`.

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

1. Bind by stable IDs.
2. Practice Types never copy Skill/Knowledge definitions.
3. A node may use multiple Practice Types across phases.
4. A Practice Type may support multiple canonical objects.
5. Concrete items reference both type and target objects.
6. Adaptive selection may vary order/frequency/difficulty but may not remove required target capability.

`PT-22` and `AT-04` are intentionally different: the former is a reusable **practice/sampling format**; the latter is the **measurement interpretation** of diagnostic evidence.

# Difficulty progression

Valid difficulty dimensions include scaffold amount, lexical/syntactic complexity, distractor quality, integration, time pressure, novelty/transfer distance, response length, and cue availability. Harder practice is not automatically higher-band truth; Band ownership remains in `05-BANDS.md`.

# Deliberate-practice and AI invariants

Practice should make the target capability, success condition, error pattern, and intended next change understandable to the learner. Raw repetition volume is not a learning strategy by itself.

AI may generate examples, explain errors, provide feedback, and select practice, but must not perform the cognitive operation on the learner's behalf when that would defeat retrieval, reasoning, planning, or independent production.

Concrete prompts/items/examples/remediation messages belong to `10-CONTENT-MODEL.md`, not this taxonomy.