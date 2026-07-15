# Grammar — Knowledge Graph (`K-GRA`, schema v1.1)
*Atomic knowledge objects with explicit dependency edges + `examples` / `common_misconceptions` ([KK-003](decisions.md)). Skill-leaf resolution lives in [resolution.md](resolution.md).*

## Foundations
**`K-GRA-010` Word classes / parts of speech** — grammar · bands 3–9
- definition: the word classes (noun, verb, adjective, adverb, preposition, conjunction, determiner, pronoun, interjection) and their function.
- requires: — . related_to: (root of the grammar graph).
- examples: *noun: decision; verb: decide; adj: decisive*.
- common_misconceptions: treating word class as fixed (ignoring derivation); confusing adjectives/adverbs.

**`K-GRA-061` Noun countability & pluralization** — grammar · bands 3–9
- definition: countable/uncountable nouns, plural forms, irregular plurals.
- requires: K-GRA-010. related_to: K-GRA-032, K-GRA-040.
- examples: *information (uncountable); children; analyses*.
- common_misconceptions: pluralizing uncountables (*informations*); irregular-plural errors (*childs*).

## Clause & sentence structure
**`K-GRA-001` Clause structure** — grammar · bands 3–9
- definition: clause elements (subject, verb, object, complement, adjunct) and basic SVO/A ordering.
- requires: K-GRA-010. related_to: K-GRA-060.
- examples: *SVO: "The graph shows data."*
- common_misconceptions: missing the verb; object/complement confusion.

**`K-GRA-002` Simple sentence** — grammar · bands 3–9
- definition: a sentence consisting of one independent clause.
- requires: K-GRA-001. related_to: K-GRA-003, K-GRA-004.
- examples: *"Prices rose sharply."*
- common_misconceptions: sentence fragments; two main verbs without coordination.

**`K-GRA-003` Compound sentence (coordination)** — grammar · bands 5–9
- definition: two or more independent clauses joined by coordinating conjunctions.
- requires: K-GRA-002, K-GRA-020. related_to: K-GRA-004.
- examples: *"Prices rose, but profits fell."*
- common_misconceptions: comma splices (*Prices rose, profits fell*); run-ons.

**`K-GRA-004` Complex sentence (subordination)** — grammar · bands 5–9
- definition: a sentence with one independent clause and at least one subordinate clause.
- requires: K-GRA-002, K-GRA-021. related_to: K-GRA-005, K-GRA-006, K-GRA-007.
- examples: *"Although prices rose, profits fell."*
- common_misconceptions: faulty subordinator; subordinate clause as a fragment.

**`K-GRA-005` Relative clauses** — grammar · bands 6–9
- definition: clauses modifying a noun via relative pronouns (defining/non-defining).
- requires: K-GRA-004, K-GRA-031. related_to: K-GRA-006, K-GRA-007.
- examples: *"The region that grew fastest was X."*
- common_misconceptions: missing relative pronoun where needed; comma errors in non-defining clauses.

**`K-GRA-006` Conditional clauses** — grammar · bands 6–9
- definition: zero/first/second/third conditionals and mixed conditionals.
- requires: K-GRA-004. related_to: K-GRA-055, K-GRA-052.
- examples: *"If demand increased, prices would rise."*
- common_misconceptions: tense mismatch across clause types; overusing "would" in the if-clause.

**`K-GRA-007` Noun clauses** — grammar · bands 6–9
- definition: subordinate clauses functioning as a noun (e.g., "that…", "what…").
- requires: K-GRA-004. related_to: K-GRA-005.
- examples: *"What surprised me was the speed."*
- common_misconceptions: word-order errors inside the clause; dropping "that" where required.

## Connectives
**`K-GRA-020` Coordinating conjunctions** — grammar · bands 4–9
- definition: and, but, or, so, yet, for, nor — joining equal elements.
- requires: K-GRA-010. related_to: K-GRA-022.
- examples: *"and, but, so"*.
- common_misconceptions: using a subordinator where a coordinator fits; comma-splice with "so".

**`K-GRA-021` Subordinating conjunctions** — grammar · bands 5–9
- definition: although, because, if, while, since, whereas, etc. — introducing subordinate clauses.
- requires: K-GRA-010. related_to: K-GRA-004.
- examples: *"although, because, whereas"*.
- common_misconceptions: treating them as sentence-linkers needing a semicolon; wrong meaning (since/while).

**`K-GRA-022` Linking adverbials / conjuncts** — grammar · bands 5–9
- definition: however, therefore, moreover, consequently — sentence-level transitions.
- requires: K-GRA-010. related_to: K-GRA-020.
- examples: *"however; therefore; in addition"*.
- common_misconceptions: punctuation errors (*However, he left.* vs run-on); overuse as paragraph filler.

## Reference & determiners
**`K-GRA-030` Pronouns** — grammar · bands 3–9
- definition: personal, possessive, demonstrative, reflexive pronouns.
- requires: K-GRA-010. related_to: K-GRA-031.
- examples: *"she, hers, this, herself"*.
- common_misconceptions: pronoun-antecedent disagreement; ambiguous reference.

**`K-GRA-031` Relative pronouns** — grammar · bands 6–9
- definition: who, whom, which, that, whose.
- requires: K-GRA-030. related_to: K-GRA-005.
- examples: *"who (people), which (things), that"*.
- common_misconceptions: who/which mix-up; whom overuse in modern style.

**`K-GRA-032` Determiners** — grammar · bands 3–9
- definition: articles, demonstratives, quantifiers, possessives — noun-phrase specifiers.
- requires: K-GRA-010. related_to: K-GRA-040, K-GRA-061.
- examples: *"the, this, some, my, each"*.
- common_misconceptions: determiner-noun disagreement (*each of the students are*); quantifier/countability clashes.

**`K-GRA-033` Reference & substitution (cohesion)** — grammar · bands 6–9
- definition: using pronouns/determiners/substitution to refer back and avoid repetition.
- requires: K-GRA-030, K-GRA-032. related_to: — .
- examples: *"the former / the latter; do so; one"*.
- common_misconceptions: ambiguous referent; faulty substitution.

## Articles
**`K-GRA-040` Articles (a / an / the)** — grammar · bands 3–9
- definition: indefinite and definite articles and their core uses.
- requires: K-GRA-032. related_to: K-GRA-041, K-GRA-061.
- examples: *"a test, an exam, the results"*.
- common_misconceptions: article omission; "the" overuse for all specifics.

**`K-GRA-041` Article rules (definite / indefinite / zero)** — grammar · bands 4–9
- definition: rules for article choice, including zero-article and generic/specific reference.
- requires: K-GRA-040. related_to: K-GRA-061.
- examples: *zero article: "Society is changing."*
- common_misconceptions: *"the society" (generic); "a water" (uncountable)*.

## Tense & aspect
**`K-GRA-050` Present simple** — grammar · bands 3–9
- definition: present simple form and uses (habits, facts, generalizations).
- requires: K-GRA-002. related_to: K-GRA-054.
- examples: *"The graph shows…"*
- common_misconceptions: 3rd-person -s omission; using present for finished past.

**`K-GRA-051` Past simple** — grammar · bands 3–9
- definition: past simple form and uses (completed past events).
- requires: K-GRA-002. related_to: K-GRA-052, K-GRA-054.
- examples: *"Sales increased in 2010."*
- common_misconceptions: irregular-verb forms; tense markers missing.

**`K-GRA-052` Present perfect** — grammar · bands 5–9
- definition: present perfect (past-with-present relevance).
- requires: K-GRA-051. related_to: K-GRA-053.
- examples: *"Demand has risen since 2010."*
- common_misconceptions: confusing with past simple; *since/for* errors.

**`K-GRA-053` Past perfect** — grammar · bands 6–9
- definition: past perfect (earlier past; sequence in the past).
- requires: K-GRA-052. related_to: K-GRA-006.
- examples: *"By 2010, prices had doubled."*
- common_misconceptions: using it without a later past reference; overuse.

**`K-GRA-054` Progressive / continuous aspect** — grammar · bands 4–9
- definition: be + -ing; ongoing actions (present/past/future continuous).
- requires: K-GRA-050, K-GRA-051. related_to: K-GRA-055.
- examples: *"is rising; were falling."*
- common_misconceptions: stative verbs in -ing (*is knowing*); missing auxiliary.

**`K-GRA-055` Future forms** — grammar · bands 4–9
- definition: will, be going to, present continuous for future.
- requires: K-GRA-002. related_to: K-GRA-006.
- examples: *"will rise; is going to fall."*
- common_misconceptions: will vs going-to confusion; overusing "will" for plans.

## Agreement, voice & morphology
**`K-GRA-060` Subject–verb agreement** — grammar · bands 3–9
- definition: agreement of verb with subject (number/person), incl. tricky cases.
- requires: K-GRA-002, K-GRA-061. related_to: — .
- examples: *"The data show…; each of the charts shows…"*
- common_misconceptions: agreement with intervening nouns; *data is/are*.

**`K-GRA-062` Modal verbs** — grammar · bands 4–9
- definition: can, could, must, should, may, might, would — ability/obligation/possibility.
- requires: K-GRA-002. related_to: — .
- examples: *"must, should, might"*.
- common_misconceptions: *can to go; modals + -s*.

**`K-GRA-063` Passive voice** — grammar · bands 6–9
- definition: be + past participle; focus on the recipient of the action.
- requires: K-GRA-002, K-GRA-061. related_to: K-GRA-052.
- examples: *"The data were collected…"*
- common_misconceptions: be/participle errors; overuse in argumentative writing.

**`K-GRA-064` Comparatives & superlatives** — grammar · bands 4–9
- definition: comparative/superlative forms of adjectives and adverbs.
- requires: K-GRA-010. related_to: — .
- examples: *"higher, the highest, more rapidly"*.
- common_misconceptions: double comparatives (*more higher*); irregular forms.

**`K-GRA-065` Negation & question forms** — grammar · bands 3–9
- definition: negation and yes/no + wh-question formation.
- requires: K-GRA-002. related_to: — .
- examples: *"does not…; why did…?"*
- common_misconceptions: auxiliary omission; word-order in wh-questions.

---

## Coverage (resolves the `K-GRA` Skill prerequisites)
| skill prerequisite (from [../skills/](../skills/)) | resolves to |
|---|---|
| basic sentence structure (W-GRA-01, S-GRA-01) | K-GRA-001, K-GRA-002 |
| coordination (W-GRA-02) | K-GRA-003, K-GRA-020 |
| conjunctions (W-CC-03, S-FC-03) | K-GRA-020, K-GRA-021, K-GRA-022 |
| subordination (W-GRA-03, S-GRA-02) | K-GRA-004, K-GRA-021, K-GRA-005/06/07 |
| reference / substitution (W-CC-04) | K-GRA-030, K-GRA-032, K-GRA-033 |
| tense / aspect (W-GRA-04) | K-GRA-050..055 |
| articles / determiners (W-GRA-05) | K-GRA-032, K-GRA-040, K-GRA-041 |
| grammatical accuracy (S-GRA-04) | K-GRA-001, K-GRA-002, K-GRA-060 |

## Open questions
- [ ] Whether to add a compound-complex sentence object (currently reachable via composition of K-GRA-003 + K-GRA-004).
