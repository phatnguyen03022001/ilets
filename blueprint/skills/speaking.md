# Speaking — Skill Decomposition (schema v1.1)
*Decomposes the Speaking competency from [../bands/speaking.md](../bands/speaking.md) into leaf skills conforming to [leaf-schema.md](leaf-schema.md) v1.1. Consistent with the [writing.md](writing.md) reference implementation.*

## Hierarchy overview
```
Speaking
├─ S-FC  Fluency & Coherence
│   ├─ S-FC-01  Sustained speech (keep going)
│   ├─ S-FC-02  Long-turn production (Part 2)
│   ├─ S-FC-03  Spoken discourse markers & connectives
│   ├─ S-FC-04  Coherent topic development
│   └─ S-FC-05  Hesitation & self-correction management
├─ S-LR  Lexical Resource
│   ├─ S-LR-01  Topic vocabulary
│   ├─ S-LR-02  Paraphrase (oral)
│   ├─ S-LR-03  Less common / idiomatic vocabulary
│   └─ S-LR-04  Word-choice accuracy
├─ S-GRA Grammatical Range & Accuracy
│   ├─ S-GRA-01  Simple & short sentence accuracy (spoken)
│   ├─ S-GRA-02  Complex sentence forms (spoken)
│   ├─ S-GRA-03  Structural variety & flexibility
│   └─ S-GRA-04  Grammatical accuracy (error frequency)
└─ S-P   Pronunciation
    ├─ S-P-01  Individual phoneme/sound accuracy
    ├─ S-P-02  Word stress
    ├─ S-P-03  Sentence stress & intonation
    ├─ S-P-04  Chunking & connected speech (rhythm)
    └─ S-P-05  Overall intelligibility
```
*(Part coverage — Part 1 interview, Part 2 long turn, Part 3 discussion — is exercised in `../practice/`; the leaves above are the underlying abilities, independent of part.)*

---

## S-FC — Fluency & Coherence

**`S-FC-01` Sustained speech (keep going)** — speaking · FC · bands 4–9 · cognitive: apply · load: medium
- objective: produce continuous speech without excessive pausing or breakdown.
- traces_to: Speaking FC — Band 4 "unable to keep going without noticeable pauses" → Band 6 "able to keep going".
- prerequisites: — .
- mastery_criteria: sustains speech on a familiar topic for ≥1 min without breakdown.
- common_errors: long silences; restarting repeatedly; reliance on single-word answers.
- remediation: timed "keep talking" drills; 1-minute topic expands.
- consumer fields: _(populated by practice/ & assessment/; dependents derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-FC-02` Long-turn production (Part 2)** — speaking · FC · bands 5–9 · cognitive: create · load: high
- objective: produce a coherent 1–2 minute extended turn from a prompt.
- traces_to: Speaking FC — Band 5 "relies on repetition/self-correction/slow speech" → Band 7 "readily produce long turns without noticeable effort".
- prerequisites: S-FC-01.
- mastery_criteria: speaks for the full ~2 min on a Part 2 prompt with coherent development.
- common_errors: running out of content; going off-topic; stopping early.
- remediation: Part-2 planning framework; idea-expansion scaffolds.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-FC-03` Spoken discourse markers & connectives** — speaking · FC · bands 5–9 · cognitive: apply · load: medium
- objective: use a range of spoken discourse markers/connectives appropriately.
- traces_to: Speaking FC — Band 6 "uses a range of spoken discourse markers… though not always appropriately" → Band 7 "flexible use".
- prerequisites: `K-GRA` conjunctions.
- mastery_criteria: uses spoken markers/connectives flexibly and appropriately.
- common_errors: overuse of fillers ("like", "you know"); written-style linkers in speech; omission.
- remediation: spoken-marker sets; natural-connector modeling.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-FC-04` Coherent topic development** — speaking · FC · bands 5–9 · cognitive: create · load: high
- objective: develop topics coherently with relevant extension.
- traces_to: Speaking FC — Band 6 "coherence may be lost" → Band 8 "topic development is coherent, appropriate and relevant".
- prerequisites: S-FC-01.
- mastery_criteria: develops a topic with relevant, extended, logically ordered points.
- common_errors: circular/undeveloped answers; abrupt topic shifts; lack of extension.
- remediation: point→reason→example expansion; coherence editing of transcripts.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-FC-05` Hesitation & self-correction management** — speaking · FC · bands 5–9 · cognitive: apply · load: medium
- objective: keep hesitations/self-corrections natural and content-related (not language-search).
- traces_to: Speaking FC — Band 5 "hesitations… mid-sentence searches for basic lexis/grammar" → Band 8 "hesitation… mostly content related".
- prerequisites: S-FC-01.
- mastery_criteria: hesitations are mostly content-related, not word/grammar searches; self-corrections are infrequent and natural.
- common_errors: frequent word-finding pauses mid-sentence; excessive self-correction.
- remediation: fluency-over-accuracy drills; topic-familiarity building.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

## S-LR — Lexical Resource

**`S-LR-01` Topic vocabulary** — speaking · LR · bands 4–9 · cognitive: apply · load: medium
- objective: use sufficient vocabulary to discuss topics at length.
- traces_to: Speaking LR — Band 4 "resource sufficient for familiar topics only" → Band 6 "sufficient to discuss topics at length".
- prerequisites: `K-VOC` topic word lists.
- mastery_criteria: discusses a topic for ≥1 min with adequate, relevant vocabulary.
- common_errors: limited range; generic wording; inability to extend on unfamiliar topics.
- remediation: topic-cluster vocabulary; speak-about-X drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-LR-02` Paraphrase (oral)** — speaking · LR · bands 5–9 · cognitive: create · load: high
- objective: paraphrase orally when the target word is unavailable.
- traces_to: Speaking LR — Band 5 "attempts paraphrase but not always with success" → Band 6 "generally able to paraphrase successfully".
- prerequisites: S-LR-01.
- mastery_criteria: paraphrases successfully to convey intended meaning when stuck.
- common_errors: stopping when a word is missing; literal/L1 translation; meaning loss.
- remediation: circumlocution strategies; "say it another way" drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-LR-03` Less common / idiomatic vocabulary** — speaking · LR · bands 6–9 · cognitive: apply · load: high
- objective: use less common and idiomatic items appropriately.
- traces_to: Speaking LR — Band 7 "some ability to use less common and idiomatic items" → Band 8 "skilful use".
- prerequisites: S-LR-01.
- mastery_criteria: uses some less common/idiomatic items accurately and appropriately.
- common_errors: forced idioms; register-inappropriate slang; misused items.
- remediation: high-utility idiomatic sets; appropriateness checks.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-LR-04` Word-choice accuracy** — speaking · LR · bands 4–9 · cognitive: apply · load: medium
- objective: choose accurate, appropriate words; avoid frequent inappropriacies.
- traces_to: Speaking LR — Band 4 "frequent inappropriacies and errors in word choice" → Band 6 "vocabulary use may be inappropriate but meaning is clear".
- prerequisites: `K-VOC`.
- mastery_criteria: word-choice errors are occasional, not frequent; meaning clear.
- common_errors: wrong word form; L1-transfer choices; collocation errors.
- remediation: error-pattern correction; collocation practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

## S-GRA — Grammatical Range & Accuracy

**`S-GRA-01` Simple & short sentence accuracy (spoken)** — speaking · GRA · bands 4–9 · cognitive: apply · load: medium
- objective: produce accurate simple/short utterances in speech.
- traces_to: Speaking GRA — Band 4 "some short utterances are error-free" → Band 6 "errors… rarely impede communication".
- prerequisites: `K-GRA` basic sentence structure.
- mastery_criteria: simple/short spoken utterances largely error-free.
- common_errors: subject–verb agreement; tense; word order (under real-time pressure).
- remediation: real-time accuracy drills; controlled speaking practice.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-GRA-02` Complex sentence forms (spoken)** — speaking · GRA · bands 5–9 · cognitive: apply · load: high
- objective: produce a mix of simple and complex sentence forms in speech.
- traces_to: Speaking GRA — Band 5 "complex structures attempted but limited… nearly always contain errors" → Band 6 "a mix of short and complex sentence forms".
- prerequisites: S-GRA-01; `K-GRA` subordination.
- mastery_criteria: uses complex forms in speech with errors that rarely impede.
- common_errors: avoidance of complex forms; breakdown under real-time load; systematic errors.
- remediation: planned-then-spoken complex-sentence drills; gradual fluency building.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-GRA-03` Structural variety & flexibility** — speaking · GRA · bands 6–9 · cognitive: create · load: high
- objective: use a range of structures flexibly in extended speech.
- traces_to: Speaking GRA — Band 6 "limited flexibility" → Band 8 "wide range of structures, flexibly used".
- prerequisites: S-GRA-02.
- mastery_criteria: uses varied structures flexibly across a long turn; frequent error-free sentences.
- common_errors: repetitive structures; limited range despite accuracy.
- remediation: structure-rotation speaking tasks; transcript variety editing.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-GRA-04` Grammatical accuracy (error frequency)** — speaking · GRA · bands 4–9 · cognitive: apply · load: medium
- objective: keep grammatical errors infrequent enough not to impede.
- traces_to: Speaking GRA — Band 4 "errors are frequent" → Band 7 "error-free sentences are frequent".
- prerequisites: `K-GRA`.
- mastery_criteria: error-free sentences are frequent; errors rarely impede communication.
- common_errors: persistent basic errors; systematic tense/article errors.
- remediation: targeted error-pattern drilling; accuracy-focused recording review.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

## S-P — Pronunciation

**`S-P-01` Individual phoneme/sound accuracy** — speaking · P · bands 4–9 · cognitive: apply · load: high
- objective: produce individual sounds/phonemes clearly enough to be understood.
- traces_to: Speaking P — Band 4 "individual words or phonemes frequently mispronounced" → Band 6 "causes only occasional lack of clarity".
- prerequisites: `K-PHON` phoneme inventory (knowledge/).
- mastery_criteria: phoneme errors cause only occasional lack of clarity.
- common_errors: L1-transfer phoneme substitution; consonant-cluster reduction; vowel confusion.
- remediation: minimal-pair drills; phoneme-focused repetition.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-P-02` Word stress** — speaking · P · bands 5–9 · cognitive: apply · load: medium
- objective: place word-level stress correctly.
- traces_to: Speaking P — Band 6 "some effective use of… stress" → Band 8 "flexible use of stress".
- prerequisites: S-P-01.
- mastery_criteria: word stress generally correct; errors don't impede.
- common_errors: wrong syllable stress; stress-shift across word families.
- remediation: word-stress pattern drills; dictionary-stress habits.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-P-03` Sentence stress & intonation** — speaking · P · bands 5–9 · cognitive: apply · load: high
- objective: use sentence stress and intonation to convey meaning.
- traces_to: Speaking P — Band 4 "attempts to use intonation and stress, control limited" → Band 6 "some effective use… not sustained".
- prerequisites: S-P-02.
- mastery_criteria: uses sentence stress/intonation effectively, generally sustained.
- common_errors: flat intonation; wrong emphasis; question-statement intonation confusion.
- remediation: intonation modeling; meaning-through-stress drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-P-04` Chunking & connected speech (rhythm)** — speaking · P · bands 5–9 · cognitive: apply · load: high
- objective: chunk speech into appropriate sense groups with natural rhythm.
- traces_to: Speaking P — Band 4 "frequent lapses in overall rhythm" → Band 6 "chunking generally appropriate".
- prerequisites: S-P-01.
- mastery_criteria: chunking generally appropriate; rhythm sustained across long utterances.
- common_errors: word-by-word delivery; lost rhythm at speed; inappropriate pausing.
- remediation: shadowing; chunking-from-text drills.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

**`S-P-05` Overall intelligibility** — speaking · P · bands 4–9 · cognitive: apply · load: medium
- objective: be understood without undue listener effort (accent effect minimal).
- traces_to: Speaking P — Band 4 "understanding requires some effort" → Band 8 "easily understood throughout… accent has minimal effect".
- prerequisites: S-P-01, S-P-03.
- mastery_criteria: can generally be understood throughout without much effort.
- common_errors: dense mispronunciation patches; accent obscuring meaning.
- remediation: intelligibility-focused feedback; priority-phoneme targeting.
- consumer fields: _(TBD / derived)_.
- independence: ✅ atomic; teachable/practiceable/assessable/remediable.

---

## Coverage check
- **FC** → S-FC-01..05 · **LR** → S-LR-01..04 · **GRA** → S-GRA-01..04 · **P** → S-P-01..05. Each official criterion fully covered; each leaf traces to a descriptor band range and satisfies the leaf invariant.

## Dependencies
- **Consumes:** [../bands/speaking.md](../bands/speaking.md).
- **References:** `K-VOC`, `K-GRA`, `K-PHON` ([../knowledge/](../knowledge/)).
- **Feeds:** [../curriculum/](../curriculum/), [../practice/](../practice/), [../assessment/](../assessment/).
- Speaking descriptors are shared across Academic & GT.

## Open questions
- [ ] Confirm pronunciation granularity (5 leaves) vs. consolidation — currently atomic per the stop condition.
- [ ] `K-PHON` (phoneme inventory/pronunciation knowledge) to be defined in [../knowledge/](../knowledge/).
