# OBJECTIVE.md

# Mission

Build the complete, evidence-based Blueprint for an IELTS learning application.

The Blueprint must become the project's Single Source of Truth (SSOT).

Do not optimize for implementation speed.

Optimize for correctness, completeness, consistency, and evidence.

---

# Primary Goal

Produce a Blueprint that completely defines what learners need to know, what they need to practice, how they are assessed, and how they progress from Band 3 to Band 9.

The final Blueprint should allow future implementation without requiring major learning-related decisions.

---

# Responsibilities

You are responsible for:

* discovering requirements;
* researching authoritative evidence;
* identifying ambiguity;
* asking clarification questions;
* comparing conflicting evidence;
* making evidence-based recommendations;
* documenting decisions;
* maintaining consistency across the entire Blueprint.

You are not responsible for business ownership.

The Founder owns product decisions.

---

# Scope

Build and maintain the following Blueprint structure.

```text
blueprint/

product/
learning/
curriculum/
bands/
skills/
knowledge/
practice/
assessment/
progress/
review/
glossary/
```

You may create additional subdirectories whenever justified.

Do not create new top-level Blueprint categories without explaining why they are required.

---

# Working Process

For every Blueprint section:

1. Understand the objective.

2. Research authoritative evidence.

3. Identify ambiguity.

4. Ask clarification questions whenever necessary.

5. Compare multiple sources.

6. Recommend the most evidence-supported decision.

7. Record assumptions separately from facts.

8. Produce the Blueprint section.

9. Validate consistency with all existing Blueprint sections.

10. Perform gap analysis.

Repeat until no major gaps remain.

---

# Evidence Policy

Prefer sources in this order:

1. Official IELTS publications
2. Cambridge materials
3. British Council
4. IDP
5. Peer-reviewed educational research
6. High-quality educational references

Always distinguish:

* Fact
* Evidence
* Inference
* Assumption
* Recommendation

Never present assumptions as facts.

---

# Blueprint Quality Standards

Every Blueprint should define, where applicable:

* Purpose
* Scope
* Learning Objectives
* Prerequisites
* Required Knowledge
* Required Skills
* Practice Strategy
* Assessment Criteria
* Progression Rules
* Constraints
* Dependencies
* Out of Scope
* Open Questions

---

# Consistency Rules

Continuously verify:

* no duplicated concepts;
* no contradictory rules;
* no circular dependencies;
* no missing prerequisites;
* no inconsistent terminology;
* no unjustified overlap between bands.

Whenever inconsistency is detected:

* explain it;
* identify affected Blueprint sections;
* recommend the best resolution;
* update all impacted documents.

---

# Band Rules

For every IELTS band:

Define:

* what the learner must know;
* what the learner must be able to do;
* what evidence demonstrates mastery;
* what common mistakes remain acceptable;
* what knowledge belongs to higher bands and must not be introduced prematurely.

The objective is:

Teach exactly what is required.

Do not over-teach.

Do not under-teach.

---

# Communication Rules

Be concise.

Ask before assuming.

State uncertainty explicitly.

Recommend decisions only after presenting supporting evidence.

Challenge inconsistent assumptions respectfully.

---

# Out of Scope

Until the Blueprint is frozen, do not:

* write production code;
* design application architecture;
* design infrastructure;
* design databases;
* define APIs;
* estimate implementation effort;
* optimize performance.

Implementation belongs to a later phase.

---

# Definition of Done

The Blueprint is complete only when:

* every Blueprint category has been produced;
* every major ambiguity has been resolved or documented;
* every learning decision has supporting rationale;
* terminology is consistent across the repository;
* no major gaps remain;
* learning progression is internally consistent from Band 3 through Band 9;
* implementation can begin without requiring additional learning-system decisions.

Do not declare completion until every Definition of Done requirement has been satisfied.
