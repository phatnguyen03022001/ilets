# Practice Binding
*Authoritative population of each Skill Leaf's `practice_item_types` slot and each Curriculum Node's practice-type references — **by `id`**, no duplication ([PR-001](decisions.md)). Practice Type definitions remain the single source in [taxonomy.md](taxonomy.md).*

## Part 1 — Skill Leaf → Practice Types
Every Skill Leaf has ≥1 suitable Practice Type (acquisition → consolidation → transfer/exam coverage).

### Writing
| leaf | practice types |
|---|---|
| W-TA-01 | PT-01, PT-13 |
| W-TA-02 | PT-01, PT-05, PT-06 |
| W-TA-03 | PT-01, PT-05, PT-06 |
| W-TR-01 | PT-01, PT-22 |
| W-TR-02 | PT-01, PT-05, PT-06 |
| W-TR-03 | PT-01, PT-05, PT-06 |
| W-TR-04 | PT-04, PT-05 |
| W-CC-01 | PT-01, PT-05 |
| W-CC-02 | PT-01, PT-05 |
| W-CC-03 | PT-02, PT-19, PT-04 |
| W-CC-04 | PT-02, PT-19, PT-04 |
| W-LR-01 | PT-17, PT-18 |
| W-LR-02 | PT-18, PT-17 |
| W-LR-03 | PT-03, PT-05 |
| W-LR-04 | PT-17, PT-18 |
| W-LR-05 | PT-18, PT-19, PT-17 |
| W-GRA-01 | PT-02, PT-19, PT-04 |
| W-GRA-02 | PT-02, PT-04 |
| W-GRA-03 | PT-02, PT-04 |
| W-GRA-04 | PT-19, PT-04, PT-17 |
| W-GRA-05 | PT-19, PT-04, PT-17 |
| W-GRA-06 | PT-04, PT-19 |
| W-GRA-07 | PT-02, PT-05 |

### Speaking
| leaf | practice types |
|---|---|
| S-FC-01 | PT-08, PT-10 |
| S-FC-02 | PT-09, PT-11 |
| S-FC-03 | PT-10, PT-08 |
| S-FC-04 | PT-09, PT-10 |
| S-FC-05 | PT-08, PT-09 |
| S-LR-01 | PT-10, PT-17 |
| S-LR-02 | PT-09, PT-10 |
| S-LR-03 | PT-17, PT-18 |
| S-LR-04 | PT-10, PT-17, PT-04 |
| S-GRA-01 | PT-08, PT-10 |
| S-GRA-02 | PT-09, PT-10 |
| S-GRA-03 | PT-09, PT-10 |
| S-GRA-04 | PT-11, PT-10 |
| S-P-01 | PT-07, PT-08 |
| S-P-02 | PT-07, PT-08 |
| S-P-03 | PT-08, PT-09 |
| S-P-04 | PT-08, PT-09 |
| S-P-05 | PT-08, PT-11 |

### Listening
| leaf | practice types |
|---|---|
| L-COMP-01 | PT-12, PT-13 |
| L-COMP-02 | PT-12, PT-13 |
| L-COMP-03 | PT-13, PT-14 |
| L-COMP-04 | PT-13, PT-16 |
| L-COMP-05 | PT-16, PT-13 |
| L-COMP-06 | PT-14, PT-15 |
| L-QT-01 | PT-19, PT-13 |
| L-QT-02 | PT-13, PT-15 |
| L-QT-03 | PT-13, PT-15 |
| L-QT-04 | PT-13, PT-15 |
| L-QT-05 | PT-19, PT-13 |

### Reading
| leaf | practice types |
|---|---|
| R-COMP-01 | PT-12, PT-13 |
| R-COMP-02 | PT-12, PT-13 |
| R-COMP-03 | PT-13, PT-14 |
| R-COMP-04 | PT-13, PT-16 |
| R-COMP-05 | PT-14, PT-13 |
| R-COMP-06 | PT-14, PT-15 |
| R-QT-01 | PT-13, PT-15 |
| R-QT-02 | PT-13, PT-16, PT-15 |
| R-QT-03 | PT-13, PT-16, PT-15 |
| R-QT-04 | PT-13, PT-15 |
| R-QT-05 | PT-19, PT-13 |
| R-QT-06 | PT-13, PT-15 |

## Part 2 — Curriculum Node → Practice Types
Every Curriculum Node references practice types aligned to its phase (acquisition → consolidation → transfer → exam-readiness as bands ascend).

### Band 3
| node | practice types |
|---|---|
| C-B3-01 | PT-02, PT-19, PT-17 |
| C-B3-02 | PT-19, PT-17, PT-04 |
| C-B3-03 | PT-19, PT-04, PT-17 |
| C-B3-04 | PT-17, PT-18, PT-19 |
| C-B3-05 | PT-07, PT-08 |
| C-B3-06 | PT-12, PT-13 |
| C-B3-07 | PT-19, PT-13 |
| C-B3-08 | PT-22 |

### Band 4
| node | practice types |
|---|---|
| C-B4-01 | PT-01, PT-04, PT-19 |
| C-B4-02 | PT-01, PT-22 |
| C-B4-03 | PT-17, PT-18, PT-10, PT-21 |
| C-B4-04 | PT-08, PT-10 |
| C-B4-05 | PT-12, PT-13 |
| C-B4-06 | PT-19, PT-13 |
| C-B4-07 | PT-22 |

### Band 5
| node | practice types |
|---|---|
| C-B5-01 | PT-02, PT-19, PT-04 |
| C-B5-02 | PT-02, PT-19, PT-04 |
| C-B5-03 | PT-01, PT-05, PT-06 |
| C-B5-04 | PT-01, PT-05, PT-06 |
| C-B5-05 | PT-03, PT-18, PT-19, PT-21 |
| C-B5-06 | PT-07, PT-08, PT-09, PT-10 |
| C-B5-07 | PT-13, PT-16, PT-14 |
| C-B5-08 | PT-13, PT-15 |
| C-B5-09 | PT-22, PT-23 |

### Band 6
| node | practice types |
|---|---|
| C-B6-01 | PT-02, PT-04, PT-05 |
| C-B6-02 | PT-05, PT-03 |
| C-B6-03 | PT-18, PT-17, PT-21 |
| C-B6-04 | PT-14, PT-15 |
| C-B6-05 | PT-13, PT-16 |
| C-B6-06 | PT-22, PT-23 |

### Band 7
| node | practice types |
|---|---|
| C-B7-01 | PT-02, PT-05, PT-20 |
| C-B7-02 | PT-17, PT-18, PT-05, PT-21 |
| C-B7-03 | PT-05, PT-20 |
| C-B7-04 | PT-13, PT-14, PT-16 |
| C-B7-05 | PT-08, PT-09 |
| C-B7-06 | PT-22, PT-23 |

### Band 8
| node | practice types |
|---|---|
| C-B8-01 | PT-05, PT-20, PT-06 |
| C-B8-02 | PT-05, PT-06 |
| C-B8-03 | PT-09, PT-11, PT-08 |
| C-B8-04 | PT-15, PT-14 |
| C-B8-05 | PT-23, PT-22 |

### Band 9
| node | practice types |
|---|---|
| C-B9-01 | PT-05, PT-20, PT-06 |
| C-B9-02 | PT-23, PT-20 |
| C-B9-03 | PT-23, PT-22 |

## Notes
- Cross-cutting/exam types (PT-20 interleaved, PT-21 adaptive, PT-22 diagnostic, PT-23 mock) are referenced primarily by **curriculum nodes** (multi-leaf integration/exam), which is their natural consumer — see [coverage-review.md](coverage-review.md).
- Adaptive scheduling (which specific type at which moment for which learner) is runtime, in `../progress/`.
