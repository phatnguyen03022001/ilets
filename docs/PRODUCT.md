# Product

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Originally created for `TASK-0001` revision 1 and extended by `TASK-0012` revision 1. This file remains non-authoritative until a later explicit canonical cutover. It cannot override `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, or `design/**`; those existing owners remain canonical.
>
> Target source snapshot: `phatnguyen03022001/ilets@23fa05c8586a9f295a3c0fe90774b78b248d61f7`.

This draft maps existing IELTS product truth into the `PRODUCT` authority domain of the adopted `agent-documents` V1 candidate. `TASK-0012` adds only the minimum actor/role identity foundation needed for the reopened product-experience milestone; `docs/catalog/project.json` owns those structural identities. This file does not close the feature catalog, freeze documentation, cut over authority, or change canonical product semantics.

## Objective

Build a complete, evidence-based IELTS learning system specification and product/runtime design for a self-directed learner with a concrete IELTS target. The product coordinates the path from target, through current evidence and explicit gaps or unknowns, to the next eligible learning action, truthful new evidence, and automatic re-planning.

The durable product value is this governed target-to-evidence-to-next-action loop. Individual AI models, SRS algorithms, media providers, evaluators, exercise formats, and realtime interaction mechanisms remain replaceable capabilities rather than the semantic center of the product.

The intended outcome is a governed route from a structured Band-3 entry range toward Band 9 across Listening, Reading, Writing, and Speaking while preserving uneven skill profiles, explicit unknown/stale/conflicting evidence, overall and per-skill target constraints, Academic and General Training differences, and a truthful distinction between learning progress, readiness, product support, and an actual external IELTS result.

The product may provide a strong governed route toward a target, but it must not imply that using the system guarantees a future external IELTS score.

**Sources:** `OBJECTIVE.md` — Purpose, Target outcome; `spec/00-PRODUCT.md` — Product identity, Promise boundary.

## Actors and roles

The product keeps three different concepts separate:

1. **participants / actor classes** describe who or what takes part in learner, support, content, or operational work;
2. **authorization** is capability-based, with named role bundles used only as default capability bundles for operator usability;
3. **commercial entitlement** controls availability of eligible product experiences and is never an RBAC role.

`LEARNER`, `COLLABORATOR`, `REVIEWER`, `ADMIN`, and `OWNER` are the ordinary default authorization bundles. The existing AI-support and optional-human-expert participation roles are not authorization bundles by themselves. A service/system identity or human may receive an explicitly granted bundle or narrower capability set only when authorized; participation never implies privilege.

**Sources:** `design/04-application-flows.md` — Actor, capability, and RBAC model; `design/00-learning-experience.md` — Entitlement-visible availability.

### Self-directed IELTS learner

The primary learner has a real IELTS target and uses the ordinary product route without mandatory teacher dependency. The learner declares the target and may later change it. Learner controls such as Swap, Skip, Shorten, or Change-skill may choose among eligible paths, but cannot erase Required prerequisites, unsupported target conditions, or uncovered product conditions. The ordinary default authorization bundle is `LEARNER`; it grants normal learner/account-scoped operations only and does not imply privileged content, support, entitlement, policy, security, or role-administration capability.

**Sources:** `OBJECTIVE.md` — Purpose, Learner route; `spec/00-PRODUCT.md` — Product identity, Human-support boundary.

### AI support role

AI may support explanation, modelling, guided practice, feedback, tutoring, generation, analysis, evaluation, Speaking interaction, and media-related learning inside bounded roles. AI does not become IELTS truth, Assessment truth, Progression authority, or the standard against which learner competence is defined.

This is a participation role, not an automatic authorization bundle. A first-party AI/service identity may receive `COLLABORATOR`, `REVIEWER`, or a narrower explicit capability grant only where current authorization policy permits it; being AI support does not imply release, entitlement, protected-data, grant-administration, security, or other privilege.

**Sources:** `OBJECTIVE.md` — Product experience; `spec/00-PRODUCT.md` — AI supports learning without becoming the standard, Human-support boundary; `design/04-application-flows.md` — Actor, capability, and RBAC model.

### Optional human expert role

Human expert input may be used for preference or coaching, or may enter as external evidence when the canonical Assessment rules admit it. Mandatory human scoring is not a prerequisite for the ordinary supported learning route merely to manufacture a consequence that the product itself cannot safely support.

Optional human learning support is likewise a participation role, not an automatic privileged authorization bundle.

**Source:** `spec/00-PRODUCT.md` — Human-support boundary.

### Privileged human operator

A privileged human operator is an authenticated person acting through one or more explicit content, operations, support, entitlement, policy, security, or grant-administration capabilities. Privilege is exercised only through the Core-owned authorization boundary; it does not become learning, Band, Assessment, evidence, content-quality, or validation-bypass authority.

**Source:** `design/04-application-flows.md` — Actor, capability, and RBAC model.

### Authorization bundles

Authorization truth is the effective capability set; these names are stable default bundles, not a second authority layer:

| Bundle | Default intent | Boundary |
| --- | --- | --- |
| `LEARNER` | normal authenticated learner/account-scoped product operations | no privileged content, support, entitlement, policy, security, or grant administration |
| `COLLABORATOR` | draft/create/edit content candidates; bounded generation/processing; classification, deduplication, correction proposals, and required non-restricted inspection | no release/activation, entitlement mutation, protected learner-data access, grant administration, or security-sensitive operation |
| `REVIEWER` | `COLLABORATOR` plus content inspection/review, validation input, permitted approve/reject outcomes, revalidation requests, and incident quarantine of new assignment | cannot bypass hard gates, grant roles, or perform unrelated learner/security administration |
| `ADMIN` | gated content release/retirement and incident resolution; bounded user support; entitlement reconciliation; operational visibility; approved operational-policy administration; purpose-scoped protected learner-data access only when separately authorized | no learning/evidence truth change, generic validation override, safety-critical financial-ceiling increase, or top-level privilege/security-policy authority |
| `OWNER` | `ADMIN` plus role/capability grant/revoke and explicitly authorized security-sensitive, recovery, destructive, or safety-critical financial-ceiling operations | still cannot rewrite immutable history or bypass hard content/evidence/product-truth gates |

One actor may hold several compatible grants, and a narrower explicit capability grant is preferred when a bundle would be excessive. `REVIEWER` approval is operational validation input, not Assessment authority; release is legal only after applicable gates pass.

**Source:** `design/04-application-flows.md` — Actor, capability, and RBAC model.

### Commercial entitlement

Commercial entitlement is a separate product-availability dimension. This migration may describe learner-visible `FREE` and `PRO` availability states, but those labels are not `ROL-*` identities and V1 defines no entitlement identity class. Entitlement may unlock otherwise eligible cost-intensive experiences, but it cannot change IELTS/Band/evidence truth, target/prerequisite semantics, learner history/data rights, or administrative privilege.

The ungated route remains a genuinely usable IELTS learning system. Eligible `PRO` entitlement may make realtime AI Speaking available only after authoritative entitlement reconciliation and only when the required provider/capability/product-support gates are actually satisfied. Exact pricing, quotas, payment-provider behavior, and commercial packaging remain outside this migration slice.

**Sources:** `design/00-learning-experience.md` — Entitlement-visible availability, Speaking interaction experience; `design/04-application-flows.md` — learner-entitlement lifecycle versus operational authorization.

## Product-level features and capabilities

This is a migration-only product-level inventory, not a closed `FTR-*` catalog and not a second authoritative feature inventory. `FTR-001` through `FTR-006` remain the six macro behavior anchors for this task. The reopened milestone intentionally defers the complete learner/admin feature and micro-skill surface decomposition to the explicit blocking design unknown in the catalog; exact feature inventories, practice-mode inventories, durations, UI defaults, and behavior remain owned by the current canonical `design/**` and `spec/**` owners until cutover.

The current product scope includes these learner-facing capability groups:

- onboarding and `TargetProfile` setup, including real overall and per-skill target constraints;
- quick and fuller diagnostic entry paths and a learner model that preserves uncertainty rather than converting unknown evidence into weakness;
- Today / Daily Plan selection from target blockers, uncertainty, due review, prerequisites, and current evidence;
- Listening, Reading, Writing, and Speaking learning and practice across supported official task/question families;
- vocabulary, grammar, and phonology acquisition, learner-saved study material, spaced retrieval, and later application in skill work;
- feedback, remediation, scaffold fading, transfer, re-evidence, fluency work, review, and exam preparation;
- media as an eligible learning source substrate when transcript, rights, and product-safety conditions permit it;
- progress and readiness interpretation that keeps Attempt, Observation, EvidenceFact, mastery/readiness, gaps, and next actions distinct;
- section and full mock flows;
- AI-supported ordinary tutoring plus bounded productive-skill and Speaking/media support without giving AI learning or evidence authority;
- focused preparation for One Skill Retake by reusing the selected skill construct rather than inventing a fifth Skill ontology.

**Sources:** `OBJECTIVE.md` — Purpose, Learner route, Product experience, Media; `spec/00-PRODUCT.md` — Product identity, Product principles, Standard IELTS learning scope, Human-support boundary.

## Scope

For the documentation migration milestone, scope is **OPEN** again because material current-milestone product/experience design has been explicitly admitted before design lock. This task establishes only the actor/RBAC/entitlement and learner/admin surface foundation; the complete learner/admin feature decomposition must be resolved before the milestone can be frozen again.

The intended complete standard-IELTS learning scope is:

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ official task/question families
+ Band-3→9 structured learning paths
+ target profile → diagnostic → learning → practice → assessment → readiness
```

The active learning specification covers the four scored skills; enabling grammar, vocabulary, and phonology; stable Skill Leaves and Knowledge Objects; Band expectations and exit criteria; curriculum sequencing and prerequisites; learning mechanisms and reusable Practice Types; Assessment, evidence, readiness, mastery, progression, regression, review, and certification semantics; and conceptual content representation needed to instantiate the system.

Academic may be released before General Training, but release order does not redefine the intended complete learning construct. Listening and Speaking are shared. Reading capabilities and question strategies are shared while variant-specific contexts/scoring remain external/contextual conditions. Academic Writing Task 1 and General Training Writing Task 1 preserve their genuine external differences; Writing Task 2 and shared Writing criteria/grammar/lexis/cohesion are not duplicated merely by variant.

IELTS for UKVI Academic/General Training reuses the corresponding learning construct while adding external administrative/security conditions. One Skill Retake reuses the selected skill construct. IELTS Life Skills is a different Listening/Speaking-only pass/fail construct and is outside the current product-learning target unless explicitly added later.

**Sources:** `OBJECTIVE.md` — Intended complete standard-IELTS scope, Scope; `spec/00-PRODUCT.md` — Standard IELTS learning scope.

## Non-goals

This migration draft preserves the current exclusions. The repository does not currently freeze:

- a final cloud/hosting provider;
- a final PostgreSQL provider;
- a final object-storage provider;
- a final identity provider;
- a final payment provider;
- a final AI/LLM/STT/TTS provider;
- final production database tables;
- Kubernetes or multi-region architecture without demonstrated need;
- pixel-perfect visual design or a final brand system;
- vendor-specific coding-agent instructions.

The product owner also does not define exact pricing, acquisition or retention mechanics, concrete UI flow, provider selection, or technical implementation. Such choices may later be selected behind canonical provider/runtime boundaries but must not redefine learning truth or product-support semantics.

**Sources:** `OBJECTIVE.md` — Non-goals; `spec/00-PRODUCT.md` — Non-goals, Implementation boundary, Commercial access boundary.

## Domain and external constraints

- External IELTS requirements are external domain truth. The learning ontology must map those requirements to capabilities without copying exam UI into the learner ontology or allowing implementation choices to redefine IELTS competence.
- Academic and General Training share common constructs and differ only where IELTS externally differs. Delivery modes may change interface familiarity, mechanics, timing/navigation practice, accessibility/capture conditions, or external eligibility/acceptance constraints, but do not redefine Skill identity, Knowledge requirements, or the Band quality standard except where a claim explicitly concerns delivery-condition readiness.
- The learning system is L1-agnostic. First-language context may improve localization, explanation, examples, interference hypotheses, or remediation, but cannot create a different IELTS competence definition.
- Evidence must remain truthful. Attempt, Observation, EvidenceFact, mastery/readiness, gaps, and next actions remain distinct; time spent, lessons completed, content consumed, AI output, or tool scores do not by themselves certify competence.
- The learner's declared target is preserved until the learner changes it. Learner agency may choose among eligible paths but cannot bypass Required prerequisites or unsupported/uncovered target conditions.
- Commercial entitlement may increase access to cost-intensive capabilities, but it must not buy a different Band standard, stronger EvidenceFact validity, hidden prerequisite bypass, fabricated certainty, administrative privilege, or ownership of learner history/data rights. An ungated ordinary route remains a genuinely usable IELTS learning system.
- Media providers are source substrates, not learning truth. Transcript/rights eligibility and safe failure behavior must be respected; arbitrary extraction or download must not be assumed.
- External providers remain replaceable behind explicit capability boundaries and must not become competing owners of learning truth.
- Product coverage and support remain distinct from learning-model coherence. The repository distinguishes `MODELLED → COVERED → SUPPORTED_FOR_PRODUCT → VALIDATED`; no global “100% IELTS” percentage may conceal a missing required condition.
- The product may control process properties such as target integrity, prerequisite/evidence rules, explainable next-action semantics, and truthful current state. It cannot guarantee a future external IELTS result.

**Sources:** `OBJECTIVE.md` — Purpose, Learner route, Coverage and support, Media, Application/runtime design; `spec/00-PRODUCT.md` — Product principles, Delivery-overlay boundary, Language/localization boundary, Human-support boundary, Commercial access boundary, Implementation boundary, Model completeness vs product support, Promise boundary.

## Migration boundary and unresolved items

No authority conflict was found between the consulted canonical PRODUCT owners for this bounded foundation slice.

`docs/**` remains **MIGRATION DRAFT / AUTHORITY NONE**. `docs/catalog/project.json` now owns the minimum migrated actor/role identities and records the milestone as `OPEN`; this file explains their product meaning only. The six existing `FTR-*` records remain macro anchors rather than the complete feature inventory. The explicit open DESIGN blocker `UNK-005` prevents false re-closure until material learner/admin surface capabilities are decomposed and traced.

This task does not establish `DOCS_READY`, design lock, canonical cutover, standards/assurance status, exact contracts, provider activation, implementation readiness, promotion, or release readiness.

**Sources:** `.agent/tasks/TASK-0012/task.yaml@4b6b29e5d3cc9b39d35267211b070d5bff44ca59`; `agent-documents@acb3f02616e700190586681306a86905792e4c07` — V1 model/reopening semantics.

## Traceability summary

| PRODUCT concern | Current canonical source at the authorized base | Draft section |
| --- | --- | --- |
| objective | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Objective |
| actors / roles | `OBJECTIVE.md`; `spec/00-PRODUCT.md`; `design/04-application-flows.md` | Actors and roles |
| features / capabilities | `OBJECTIVE.md`; `spec/00-PRODUCT.md`; `design/00-learning-experience.md`; `design/04-application-flows.md` | Product-level features and capabilities |
| scope / non-goals | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Scope; Non-goals |
| domain / external constraints | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Domain and external constraints |

The original PRODUCT migration references remain traceable to their recorded source snapshot; the `TASK-0012` actor/RBAC/entitlement and surface-foundation additions are derived from canonical `design/00-learning-experience.md` and `design/04-application-flows.md` at the accepted baseline `010a5cc9b002eaa22121047c54c8df116cfb3e27`. This traceability is migration evidence only and does not change source authority.