STATUS: CANONICAL
OWNS: how capability is trained, canonical learning phases, Practice Type taxonomy, practice-type schema, feedback timing for learning, and binding rules from practice to canonical learning objects
DEPENDS_ON: 03-SKILLS.md, 04-KNOWLEDGE.md, 06-CURRICULUM.md
DOES_NOT_OWN: Skill or Knowledge definitions, band thresholds, assessment evidence sufficiency, certification, learner-state transitions, concrete generated exercise instances

# 07 — Practice

## Purpose

Define **how canonical capability is trained**.

Practice is not assessment. Practice may generate useful performance evidence, but the rules that decide whether evidence is sufficient for mastery belong to `08-ASSESSMENT.md`.

Practice is also not a library of individual exercises. This spec owns reusable **Practice Types**. Concrete exercise instances conform to those types and are represented through `10-CONTENT-MODEL.md`.

## Canonical learning phases

Every Practice Type supports one or more phases and has exactly one primary phase.

### Acquisition

Initial construction of a new capability or knowledge pattern. Use high scaffolding, explicit modeling, and immediate corrective feedback where appropriate.

### Consolidation

Stabilize newly acquired capability through focused, varied repetition and error correction.

### Retrieval

Actively recall or reproduce knowledge/capability after delay. Retrieval should require the learner to generate the response rather than merely reconsume the material.

### Transfer

Apply learning in a new, less scaffolded, or more integrated context. Transfer deliberately reduces cue dependence and increases variation.

### Fluency

Build speed, automaticity, rhythm, and reduced cognitive effort without losing the underlying quality standard.

### Exam Readiness

Practice under IELTS-like conditions for timing, integration, stamina, task familiarity, and readiness estimation.

Exam Readiness is a learning phase. **Exam Preparation** is a learner mode owned by `09-PROGRESSION.md`; it may select Exam-Readiness practice without certifying higher mastery.

## Feedback timing

Feedback timing is stage-dependent rather than universally immediate.

- **Acquisition:** immediate, specific, actionable feedback is usually preferred to prevent stabilizing a wrong model.
- **Consolidation:** feedback may remain prompt but should increasingly require self-correction or retrieval before revealing the answer.
- **Retrieval:** allow an authentic retrieval attempt before feedback.
- **Transfer:** increasingly delay or batch feedback where immediate intervention would disrupt independent performance.
- **Fluency:** avoid excessive interruption; prioritize post-attempt patterns unless a persistent error blocks the fluency objective.
- **Exam Readiness:** preserve exam-like independence during the attempt; feedback is primarily post-task.

The objective is the most effective feedback for the learning stage, not the shortest possible response time.

## Practice Type semantic shape

A Practice Type owns:

- stable `id`;
- name;
- primary phase;
- supported phases;
- mode: individual, mixed, adaptive, timed, diagnostic, or equivalent;
- skill/knowledge scope;
- learning objective served;
- relevant cognitive operations;
- canonical Skill Leaf references;
- canonical Knowledge Object references;
- recommended use in the curriculum;
- learning-feedback pattern.

Practice Types reference canonical learning objects by ID and never redefine them.

## Canonical Practice Type taxonomy

The active taxonomy contains **23 stable Practice Types**.

### Writing

| ID | Type | Primary phase | Core use |
|---|---|---|---|
| `PT-01` | Scaffolded / guided writing | Acquisition | Build a new writing capability with support that can fade. |
| `PT-02` | Sentence-combining drill | Acquisition | Develop compound/complex construction toward automatic use. |
| `PT-03` | Paraphrase task | Consolidation | Re-express prompts, data, and ideas without meaning loss. |
| `PT-04` | Error-correction exercise | Consolidation | Diagnose and correct targeted grammar/lexical errors. |
| `PT-05` | Redraft after feedback | Transfer | Improve a full response by applying feedback to a new draft. |
| `PT-06` | Timed Writing task | Exam Readiness | Produce Academic Task 1 or Task 2 under exam-like timing. |

Typical bindings:

- `PT-01` → acquisition across `W-*`, especially grammar, cohesion, and task response;
- `PT-02` → `W-GRA-02`, `W-GRA-03`, `W-GRA-07` plus their grammar objects;
- `PT-03` → `W-LR-03`, `W-CC-04` and paraphrase/collocation knowledge;
- `PT-04` → grammar/lexical accuracy leaves;
- `PT-05` → integrated Writing criteria;
- `PT-06` → integrated Writing under test conditions.

### Speaking

| ID | Type | Primary phase | Core use |
|---|---|---|---|
| `PT-07` | Pronunciation / minimal-pair drill | Acquisition | Build phoneme and stress discrimination/production. |
| `PT-08` | Shadowing | Fluency | Develop rhythm, stress, connected speech, and continuity. |
| `PT-09` | Long-turn practice | Consolidation | Produce coherent 1–2 minute Part-2-style turns. |
| `PT-10` | Q&A / role-play | Consolidation | Practice responsive Part-1/Part-3-style spoken interaction. |
| `PT-11` | Timed mock Speaking | Exam Readiness | Complete a full Speaking sequence without mid-attempt coaching. |

Typical bindings include `S-P-*`, `S-FC-*`, productive grammar/lexical leaves, and the corresponding phonology/grammar/vocabulary objects.

### Listening and Reading

| ID | Type | Primary phase | Core use |
|---|---|---|---|
| `PT-12` | Skim / scan / gist-detail speed drill | Fluency | Build efficient location and global-understanding behavior. |
| `PT-13` | Comprehension question set | Consolidation | Apply comprehension and question-type capability to texts/recordings. |
| `PT-14` | Note-taking from lecture/text | Transfer | Capture structure, key points, and detail from extended input. |
| `PT-15` | Timed section / passage practice | Exam Readiness | Complete a full receptive unit under timing. |
| `PT-16` | Distractor / error review | Retrieval | Explain why an answer was wrong and retrieve the correct reasoning. |

`PT-12` may target Listening/Reading gist-detail leaves differently in concrete instances; the Practice Type remains one reusable training pattern.

### Knowledge

| ID | Type | Primary phase | Core use |
|---|---|---|---|
| `PT-17` | Spaced retrieval | Retrieval | Maintain vocabulary/grammar knowledge through scheduled active recall. |
| `PT-18` | Collocation / word-formation drill | Acquisition | Acquire lexical partnerships and productive morphology. |
| `PT-19` | Gap-fill / completion | Consolidation | Reinforce form and accuracy in context. |

Practice on Knowledge Objects exists to support actual skill performance. Knowledge practice is not a substitute for demonstrating the target IELTS skill.

### Cross-cutting and exam

| ID | Type | Primary phase | Core use |
|---|---|---|---|
| `PT-20` | Interleaved mixed set | Transfer | Discriminate and apply across mixed skills/topics instead of blocked repetition. |
| `PT-21` | Adaptive practice set | Consolidation | Select appropriate practice from learner state without changing target outcomes. |
| `PT-22` | Diagnostic checkpoint practice | Retrieval | Sample current strengths/gaps to guide next learning action. |
| `PT-23` | Full mock test | Exam Readiness | Simulate the full IELTS experience across all four skills. |

`PT-23` is non-certifying by itself. Whether its output contributes assessment evidence is governed by `08-ASSESSMENT.md`.

## Phase coverage invariant

Every canonical phase must have at least one suitable primary Practice Type.

Current coverage:

- Acquisition — `PT-01`, `PT-02`, `PT-07`, `PT-18`;
- Consolidation — `PT-03`, `PT-04`, `PT-09`, `PT-10`, `PT-13`, `PT-19`, `PT-21`;
- Retrieval — `PT-16`, `PT-17`, `PT-22`;
- Transfer — `PT-05`, `PT-14`, `PT-20`;
- Fluency — `PT-08`, `PT-12`;
- Exam Readiness — `PT-06`, `PT-11`, `PT-15`, `PT-23`.

This does not require equal volume per phase. Coverage must follow learning need.

## Practice binding rules

A concrete practice selection resolves:

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

1. Practice Types are referenced by ID.
2. Skill/Knowledge definitions are never embedded into a Practice Type.
3. A Curriculum Node may use multiple Practice Types across phases.
4. A Practice Type may support multiple canonical objects.
5. Concrete practice items carry references back to both their type and target objects.
6. Adaptive selection may change frequency, difficulty, or order but may not remove required target capability.

## Difficulty progression

Difficulty should rise by manipulating task demand rather than silently changing the target semantic. Valid dimensions include:

- amount of scaffolding;
- lexical/syntactic complexity;
- distractor quality;
- integration across leaves;
- time pressure;
- novelty and transfer distance;
- response length;
- cue availability.

A more difficult exercise is not automatically a higher-band requirement. Band ownership remains in `05-BANDS.md`.

## Deliberate practice rule

Practice should be focused enough that the learner can identify:

- what capability is being trained;
- what successful performance looks like;
- what error pattern occurred;
- what next attempt should change.

Raw repetition volume without targeted adaptation or feedback is not a canonical learning strategy.

## AI role in practice

AI may be the primary provider of explanations, generated examples, targeted feedback, and adaptive practice selection.

AI should not remove the learner's need to perform the cognitive operation being learned. Assistance must be reduced when it would substitute for retrieval, reasoning, planning, or independent production.

Assessment confidence and certification consequences remain outside this file.
