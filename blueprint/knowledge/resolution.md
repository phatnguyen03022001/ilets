# Knowledge Resolution Map
*Maps every `K-*` prerequisite referenced by the [Skill Graph](../skills/) to the canonical Knowledge Object(s) that satisfy it ([KK-001](decisions.md)). Strictly Skill↔Knowledge (kept separate from the Knowledge↔Knowledge edges in the domain docs).*

**Inventory:** `K-GRA` 29 · `K-VOC` 9 · `K-PHON` 8 = **46 atomic knowledge objects.**

## Resolution (skill leaf → knowledge objects)

### Writing ([../skills/writing.md](../skills/writing.md))
| skill leaf | prerequisite (as stated) | resolves to |
|---|---|---|
| W-CC-03 | cohesive devices (conjunctions) | K-GRA-020, K-GRA-021, K-GRA-022 |
| W-CC-04 | reference & substitution | K-GRA-030, K-GRA-032, K-GRA-033 |
| W-LR-01 | topic word lists | K-VOC-012 |
| W-LR-02 | collocations | K-VOC-020 |
| W-LR-03 | paraphrase | K-VOC-041, K-VOC-020 |
| W-LR-05 | spelling / word formation | K-VOC-030, K-VOC-031 |
| W-GRA-01 | basic sentence structure | K-GRA-001, K-GRA-002 |
| W-GRA-02 | coordination | K-GRA-003, K-GRA-020 |
| W-GRA-03 | subordination | K-GRA-004, K-GRA-021, K-GRA-005, K-GRA-006, K-GRA-007, K-GRA-031 |
| W-GRA-04 | tense & aspect | K-GRA-050, K-GRA-051, K-GRA-052, K-GRA-053, K-GRA-054, K-GRA-055 |
| W-GRA-05 | articles & determiners | K-GRA-032, K-GRA-040, K-GRA-041 |
| W-TA-03 | report key features with data | K-GRA-064 |
| W-GRA-07 | structural flexibility | K-GRA-005, K-GRA-006, K-GRA-007, K-GRA-062, K-GRA-063 |

### Speaking ([../skills/speaking.md](../skills/speaking.md))
| skill leaf | prerequisite (as stated) | resolves to |
|---|---|---|
| S-FC-03 | discourse markers / connectives | K-GRA-020, K-GRA-021, K-GRA-022 |
| S-LR-01 | topic word lists | K-VOC-012 |
| S-LR-04 | word-choice accuracy | K-VOC-040, K-VOC-020 |
| S-GRA-01 | basic sentence structure | K-GRA-001, K-GRA-002 |
| S-GRA-02 | subordination | K-GRA-004, K-GRA-021, K-GRA-005 |
| S-GRA-04 | grammatical accuracy | K-GRA-001, K-GRA-002, K-GRA-060, K-GRA-065 |
| S-P-01 | phoneme inventory | K-PHON-010, K-PHON-011, K-PHON-012 |
| S-P-02 | word stress | K-PHON-020 |
| S-P-03 | sentence stress & intonation | K-PHON-021, K-PHON-030 |

### Reading ([../skills/reading.md](../skills/reading.md))
| skill leaf | prerequisite (as stated) | resolves to |
|---|---|---|
| R-COMP-06 | academic vocabulary | K-VOC-011 |

### Listening ([../skills/listening.md](../skills/listening.md))
No `K-*` prerequisites (receptive comprehension; prereqs are intra-skill COMP leaves).

## Completeness check
- **Every `K-*` Skill prerequisite resolves** to ≥1 knowledge object: ✅ (18 distinct skill-leaf prerequisites, all resolved).
- **Every referenced knowledge object exists** in the Knowledge Graph: ✅ — all `K-GRA-*` / `K-VOC-*` / `K-PHON-*` IDs above are defined in [grammar.md](grammar.md) / [vocabulary.md](vocabulary.md) / [phonology.md](phonology.md).
- This **fully satisfies** the unresolved `K-*` dependencies flagged in [../skills/consistency-review.md](../skills/consistency-review.md) (`K-VOC`, `K-GRA`, `K-PHON`).

## Reverse index (knowledge object → skill leaves served) — high-use objects
| knowledge object | serves skill leaves |
|---|---|
| K-GRA-002 (simple sentence) | W-GRA-01, W-GRA-02*, S-GRA-01, S-GRA-04 |
| K-GRA-004 (complex sentence) | W-GRA-03, S-GRA-02 |
| K-GRA-021 (subordinating conjunctions) | W-GRA-03, S-GRA-02, W-CC-03, S-FC-03 |
| K-GRA-040 (articles) | W-GRA-05 |
| K-VOC-012 (topic sets) | W-LR-01, S-LR-01 |
| K-VOC-020 (collocations) | W-LR-02, S-LR-04 |
| K-PHON-010/011/012 (phonemes) | S-P-01 |

(*transitively, via the requires graph.)

## Notes
- Some Skill prerequisites resolve to **multiple** knowledge objects (e.g., W-GRA-03 → 5 objects) because the skill spans several atomic knowledge units — exactly the resolution the Knowledge Graph is meant to provide.
- This map is the single Skill↔Knowledge reference; if a Skill prerequisite is later added, it must be added here (and a matching knowledge object must exist).
