# Phonology — Knowledge Graph (`K-PHON`, schema v1.1)
*Atomic knowledge objects with explicit edges + `examples` / `common_misconceptions`. Skill-leaf resolution in [resolution.md](resolution.md).*

## Phonemes
**`K-PHON-010` Consonant phonemes** — phonology · bands 3–9
- definition: the English consonant phoneme inventory and their articulation.
- requires: — .
- related_to: K-PHON-011, K-PHON-012.
- examples: /θ/ vs /t/ (*think/tink*); /ʃ/ (*ship*).
- common_misconceptions: conflating phonemes absent in the learner's L1; final-consonant deletion.

**`K-PHON-011` Vowel phonemes** — phonology · bands 3–9
- definition: English monophthongs and diphthongs.
- requires: — .
- related_to: K-PHON-010, K-PHON-012.
- examples: /iː/ vs /ɪ/ (*sheep/ship*); /æ/ vs /e/ (*cat/bet*).
- common_misconceptions: treating tense/lax vowel pairs as identical; L1-vowel substitution.

**`K-PHON-012` Phoneme contrasts & minimal pairs** — phonology · bands 4–9
- definition: contrastive phoneme pairs (minimal pairs), prioritizing L1-relevant distinctions.
- requires: K-PHON-010, K-PHON-011.
- related_to: — .
- examples: */b/ vs /p/; /l/ vs /r/; /v/ vs /w/*.
- common_misconceptions: assuming hearing = producing; ignoring L1-specific confusions.

## Stress
**`K-PHON-020` Word stress** — phonology · bands 4–9
- definition: correct syllable stress within words; stress-shift across word families.
- requires: K-PHON-011.
- related_to: K-PHON-021.
- examples: *phoTOgrapher* (not *PHOtographer*); *REcord (n) vs reCORD (v)*.
- common_misconceptions: transferring L1 stress; stressing every syllable equally.

**`K-PHON-021` Sentence stress** — phonology · bands 5–9
- definition: stressing content/meaning-bearing words; reducing function words.
- requires: K-PHON-020.
- related_to: K-PHON-030.
- examples: *"She BOUGHT a CAR"* (content stressed).
- common_misconceptions: stressing all words; wrong emphasis changing meaning.

## Intonation & fluency
**`K-PHON-030` Intonation patterns** — phonology · bands 5–9
- definition: rising/falling contours and their discourse/attitude functions.
- requires: K-PHON-021.
- related_to: — .
- examples: rising for yes/no questions/list items; falling for statements/wh-questions.
- common_misconceptions: flat intonation; rising where falling is needed (statement read as question).

**`K-PHON-040` Connected speech** — phonology · bands 5–9
- definition: linking, assimilation, elision, and weak forms across word boundaries.
- requires: K-PHON-010, K-PHON-011.
- related_to: K-PHON-041.
- examples: *"cup of tea" → cuppa*; weak *"and" → /ən/*.
- common_misconceptions: pronouncing every word in isolation; missing weak forms in listening.

**`K-PHON-041` Rhythm & chunking** — phonology · bands 5–9
- definition: grouping speech into sense chunks; stress-timed rhythm.
- requires: K-PHON-021, K-PHON-040.
- related_to: — .
- examples: pausing at clause boundaries; natural sense-group chunking.
- common_misconceptions: word-by-word delivery; pausing mid-phrase.

---

## Coverage (resolves the `K-PHON` Skill prerequisites)
| skill prerequisite | resolves to |
|---|---|
| phoneme inventory (S-P-01) | K-PHON-010, K-PHON-011, K-PHON-012 |
| word stress (S-P-02) | K-PHON-020 |
| sentence stress & intonation (S-P-03) | K-PHON-021, K-PHON-030 |
| chunking & connected speech / rhythm (S-P-04) | K-PHON-040, K-PHON-041 |
| intelligibility (S-P-05) | supported by K-PHON-010..041 collectively |

## Open questions
- [ ] Whether to add L1-specific phoneme-contrast objects (localization-layer, per [FD-002](../product/foundational-decisions.md)) rather than embedding L1 notes in K-PHON-012.
