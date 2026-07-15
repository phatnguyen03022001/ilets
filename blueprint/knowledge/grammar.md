# Grammar — Knowledge Graph (`K-GRA`, schema v1.0)
*Atomic knowledge objects with explicit dependency edges ([knowledge-object-schema.md](knowledge-object-schema.md)). Template for vocabulary/phonology. Skill-leaf resolution lives in [resolution.md](resolution.md).*

Each object: `id` · `definition` (atomic) · `requires` (explicit prerequisite edges) · `related_to` (peer edges) · `band_relevance`.

## Foundations
**`K-GRA-010` Word classes / parts of speech** — grammar · bands 3–9
- definition: the word classes (noun, verb, adjective, adverb, preposition, conjunction, determiner, pronoun, interjection) and their function.
- requires: — .
- related_to: (root of the grammar graph).

**`K-GRA-061` Noun countability & pluralization** — grammar · bands 3–9
- definition: countable/uncountable nouns, plural forms, irregular plurals.
- requires: K-GRA-010.
- related_to: K-GRA-032 (determiners), K-GRA-040 (articles).

## Clause & sentence structure
**`K-GRA-001` Clause structure** — grammar · bands 3–9
- definition: clause elements (subject, verb, object, complement, adjunct) and basic SVO/A ordering.
- requires: K-GRA-010.
- related_to: K-GRA-060 (agreement).

**`K-GRA-002` Simple sentence** — grammar · bands 3–9
- definition: a sentence consisting of one independent clause.
- requires: K-GRA-001.
- related_to: K-GRA-003, K-GRA-004.

**`K-GRA-003` Compound sentence (coordination)** — grammar · bands 5–9
- definition: two or more independent clauses joined by coordinating conjunctions.
- requires: K-GRA-002, K-GRA-020.
- related_to: K-GRA-004.

**`K-GRA-004` Complex sentence (subordination)** — grammar · bands 5–9
- definition: a sentence with one independent clause and at least one subordinate clause.
- requires: K-GRA-002, K-GRA-021.
- related_to: K-GRA-005, K-GRA-006, K-GRA-007.

**`K-GRA-005` Relative clauses** — grammar · bands 6–9
- definition: clauses modifying a noun via relative pronouns (defining/non-defining).
- requires: K-GRA-004, K-GRA-031.
- related_to: K-GRA-006, K-GRA-007.

**`K-GRA-006` Conditional clauses** — grammar · bands 6–9
- definition: zero/first/second/third conditionals and mixed conditionals.
- requires: K-GRA-004.
- related_to: K-GRA-055 (future), K-GRA-052 (perfect).

**`K-GRA-007` Noun clauses** — grammar · bands 6–9
- definition: subordinate clauses functioning as a noun (e.g., "that…", "what…").
- requires: K-GRA-004.
- related_to: K-GRA-005.

## Connectives
**`K-GRA-020` Coordinating conjunctions** — grammar · bands 4–9
- definition: and, but, or, so, yet, for, nor — joining equal elements.
- requires: K-GRA-010.
- related_to: K-GRA-022.

**`K-GRA-021` Subordinating conjunctions** — grammar · bands 5–9
- definition: although, because, if, while, since, whereas, etc. — introducing subordinate clauses.
- requires: K-GRA-010.
- related_to: K-GRA-004.

**`K-GRA-022` Linking adverbials / conjuncts** — grammar · bands 5–9
- definition: however, therefore, moreover, consequently — sentence-level transitions.
- requires: K-GRA-010.
- related_to: K-GRA-020.

## Reference & determiners
**`K-GRA-030` Pronouns** — grammar · bands 3–9
- definition: personal, possessive, demonstrative, reflexive pronouns.
- requires: K-GRA-010.
- related_to: K-GRA-031.

**`K-GRA-031` Relative pronouns** — grammar · bands 6–9
- definition: who, whom, which, that, whose.
- requires: K-GRA-030.
- related_to: K-GRA-005.

**`K-GRA-032` Determiners** — grammar · bands 3–9
- definition: articles, demonstratives, quantifiers, possessives — noun-phrase specifiers.
- requires: K-GRA-010.
- related_to: K-GRA-040, K-GRA-061.

**`K-GRA-033` Reference & substitution (cohesion)** — grammar · bands 6–9
- definition: using pronouns/determiners/substitution to refer back and avoid repetition.
- requires: K-GRA-030, K-GRA-032.
- related_to: — .

## Articles
**`K-GRA-040` Articles (a / an / the)** — grammar · bands 3–9
- definition: indefinite and definite articles and their core uses.
- requires: K-GRA-032.
- related_to: K-GRA-041, K-GRA-061.

**`K-GRA-041` Article rules (definite / indefinite / zero)** — grammar · bands 4–9
- definition: rules for article choice, including zero-article and generic/specific reference.
- requires: K-GRA-040.
- related_to: K-GRA-061.

## Tense & aspect
**`K-GRA-050` Present simple** — grammar · bands 3–9
- definition: present simple form and uses (habits, facts, generalizations).
- requires: K-GRA-002.
- related_to: K-GRA-054.

**`K-GRA-051` Past simple** — grammar · bands 3–9
- definition: past simple form and uses (completed past events).
- requires: K-GRA-002.
- related_to: K-GRA-052, K-GRA-054.

**`K-GRA-052` Present perfect** — grammar · bands 5–9
- definition: present perfect (past-with-present relevance).
- requires: K-GRA-051.
- related_to: K-GRA-053.

**`K-GRA-053` Past perfect** — grammar · bands 6–9
- definition: past perfect (earlier past; sequence in the past).
- requires: K-GRA-052.
- related_to: K-GRA-006.

**`K-GRA-054` Progressive / continuous aspect** — grammar · bands 4–9
- definition: be + -ing; ongoing actions (present/past/future continuous).
- requires: K-GRA-050, K-GRA-051.
- related_to: K-GRA-055.

**`K-GRA-055` Future forms** — grammar · bands 4–9
- definition: will, be going to, present continuous for future.
- requires: K-GRA-002.
- related_to: K-GRA-006.

## Agreement, voice & morphology
**`K-GRA-060` Subject–verb agreement** — grammar · bands 3–9
- definition: agreement of verb with subject (number/person), incl. tricky cases.
- requires: K-GRA-002, K-GRA-061.
- related_to: — .

**`K-GRA-062` Modal verbs** — grammar · bands 4–9
- definition: can, could, must, should, may, might, would — ability/obligation/possibility.
- requires: K-GRA-002.
- related_to: — .

**`K-GRA-063` Passive voice** — grammar · bands 6–9
- definition: be + past participle; focus on the recipient of the action.
- requires: K-GRA-002, K-GRA-061.
- related_to: K-GRA-052.

**`K-GRA-064` Comparatives & superlatives** — grammar · bands 4–9
- definition: comparative/superlative forms of adjectives and adverbs.
- requires: K-GRA-010.
- related_to: — .

**`K-GRA-065` Negation & question forms** — grammar · bands 3–9
- definition: negation and yes/no + wh-question formation.
- requires: K-GRA-002.
- related_to: — .

---

## Coverage (resolves the `K-GRA` Skill prerequisites)
| skill prerequisite (from [../skills/](../skills/)) | resolves to |
|---|---|
| basic sentence structure (W-GRA-01, S-GRA-01) | K-GRA-001, K-GRA-002 |
| coordination (W-GRA-02) | K-GRA-003, K-GRA-020 |
| conjunctions (W-CC-03) | K-GRA-020, K-GRA-021, K-GRA-022 |
| subordination (W-GRA-03, S-GRA-02) | K-GRA-004, K-GRA-021, K-GRA-005/06/07 |
| reference / substitution (W-CC-04) | K-GRA-030, K-GRA-032, K-GRA-033 |
| tense / aspect (W-GRA-04) | K-GRA-050..055 |
| articles / determiners (W-GRA-05) | K-GRA-032, K-GRA-040, K-GRA-041 |

Every `K-GRA` Skill prerequisite resolves to ≥1 atomic grammar object. Full cross-domain map in [resolution.md](resolution.md) (after vocabulary/phonology are built).

## Open questions
- [ ] Validate granularity (26 objects) + the dependency edges before replicating to vocabulary/phonology.
- [ ] Whether to add a `K-GRA-009` compound-complex sentence object (currently reachable via K-GRA-003 + K-GRA-004 composition).
