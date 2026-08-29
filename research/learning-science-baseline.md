TYPE: RESEARCH
STATUS: NON-CANONICAL
OWNER: learning-policy evidence

# Learning-Science Baseline

This is supporting research, not specification authority.

Canonical policy lives in the owning `spec/*.md` file. This research exists to explain or challenge that policy.

## Provenance

The pre-refactor repository contained an independently audited learning-science brief preserved in Git history at immutable commit `43cddc45160016beb6d2d9342acc2fd9db4b1075`, path:

```text
archive/legacy-2026-07-16/blueprint/learning/learning-science-evidence-brief.md
```

That audit is important because the original research draft contained citation errors, including one unverifiable/fabricated citation and multiple incorrect venues. The legacy audit corrected those issues and explicitly warned against trusting precise effect sizes without primary-source verification.

This baseline therefore records **directions of evidence**, not publication-ready quantitative claims.

## Strong / broadly supported directions

### Deliberate practice quality

Focused practice with a clear goal, learner performance, and useful feedback is preferable to undirected repetition or raw hour accumulation.

Relevant legacy sources include Ericsson et al. (1993) and the later corrective perspective from Macnamara et al. (2016) that practice volume alone explains only part of performance variation.

Canonical consequences:

- `spec/00-PRODUCT.md` — learning value over feature/activity volume;
- `spec/07-PRACTICE.md` — deliberate-practice rule.

### Mastery learning and criterion-referenced progress

Mastery-learning literature supports advancement based on demonstrated criterion attainment rather than norm ranking or time spent.

Relevant legacy sources include Bloom (1968), Kulik et al. (1990), and Guskey (2010).

Canonical consequences:

- `spec/08-ASSESSMENT.md` — criterion-referenced evidence;
- `spec/09-PROGRESSION.md` — evidence-gated advancement.

### Retrieval practice

Active retrieval generally supports long-term retention better than passive restudy.

Relevant legacy sources include Roediger & Karpicke (2006) and Adesope et al. (2017).

Canonical consequence:

- Retrieval is one of the six learning phases in `spec/07-PRACTICE.md`.

### Spacing

Distributed review is generally preferable to massed repetition for long-term retention. Optimal spacing varies by retention horizon, material, and learner.

Relevant legacy sources include Cepeda et al. (2008) and forgetting-curve replication work such as Murre & Dros (2015).

Canonical consequence:

- `spec/09-PROGRESSION.md` uses performance-informed review spacing rather than a single fixed universal interval schedule.

### Formative assessment

Frequent assessment used to inform next learning actions is generally valuable when feedback is specific and connected to a target.

Relevant legacy sources include Black & Wiliam (1998), Hattie & Timperley (2007), and later review literature.

Canonical consequences:

- `spec/08-ASSESSMENT.md` distinguishes formative, readiness, and certification roles;
- `spec/07-PRACTICE.md` treats feedback as part of learning rather than mere scoring.

### Cognitive load and scaffolding

Novices often benefit from more explicit structure and examples; scaffolding should fade as expertise grows. Unnecessary/extraneous load should be reduced without removing productive learning effort such as retrieval and transfer.

Relevant legacy sources include Sweller, van Merriënboer & Paas (2019) and worked-example research.

Canonical consequences:

- `spec/07-PRACTICE.md` supports scaffold fading and difficulty manipulation;
- `spec/10-CONTENT-MODEL.md` requires scaffold level to be visible on concrete practice instances.

## Evidence with important uncertainty

### Exact L2 prerequisite graph

Direct evidence for an optimal IELTS/L2 prerequisite graph is limited.

The legacy audit explicitly removed a purported prerequisite-mapping citation that could not be verified. Therefore prerequisite edges should be conservative and justified, not treated as a scientifically complete dependency graph.

Canonical consequence:

- `spec/06-CURRICULUM.md` uses `Required`, `Recommended`, and `Independent`, with minimum hard gates.

### Feedback timing

Evidence does not support a universal rule that all useful feedback must be immediate. Immediate correction can support initial acquisition, while delayed/retrieval-first feedback may better preserve independent performance during retrieval, transfer, fluency, and exam simulation.

Canonical consequence:

- stage-dependent timing in `spec/07-PRACTICE.md`.

### Content grain size

There is no universal evidence-backed leaf size for language learning. Fine decomposition can support novices and remediation; excessive fragmentation can impair integration and transfer.

Canonical consequence:

- atomicity is functional, not numerical, in `spec/03-SKILLS.md` and `spec/04-KNOWLEDGE.md`.

### AI in language learning

Evidence for AI-assisted language learning is promising but heterogeneous by task, model, study quality, and outcome. High-stakes productive scoring requires stronger calibration than low-stakes generation or objective-item feedback.

Canonical consequences:

- AI may support high-volume practice;
- `spec/08-ASSESSMENT.md` requires calibrated uncertainty for productive mastery decisions and does not treat model self-confidence as calibration evidence.

## Research discipline

When using this file:

1. distinguish consensus, contested findings, and open questions;
2. verify primary sources before publishing exact effect sizes or strong causal claims;
3. do not turn indirect evidence from another domain into a hard IELTS prerequisite without explicit reasoning;
4. if credible new evidence conflicts with a canonical learning rule, raise the conflict and update the owning spec only after the decision is made.

## Key legacy references

The audited legacy brief discusses, among others:

- Bloom — mastery learning;
- Kulik, Kulik & Bangert-Drowns — mastery-learning meta-analysis;
- Ericsson, Krampe & Tesch-Römer — deliberate practice;
- Macnamara, Moreau & Hambrick — limits of deliberate-practice variance;
- Roediger & Karpicke — retrieval/testing effect;
- Adesope, Trevisan & Sundararajan — practice-testing meta-analysis;
- Cepeda and colleagues — distributed practice/spacing;
- Dunlosky and colleagues — learning-technique review;
- Sweller, van Merriënboer & Paas — cognitive load theory;
- Black & Wiliam — formative assessment;
- Hattie & Timperley — feedback;
- VanLehn — tutoring systems;
- DeKeyser — adult L2 skill acquisition;
- Swain — output in language learning.

For full bibliographic detail and the citation-audit history, consult the historical brief at the immutable Git reference above rather than treating this summary as a bibliography.
