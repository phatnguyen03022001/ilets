# Product

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Created only for `TASK-0001` revision 1. This file is non-authoritative until a later explicit canonical cutover. It cannot override `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, or `design/**`; those existing owners remain canonical.
>
> Target source snapshot: `phatnguyen03022001/ilets@23fa05c8586a9f295a3c0fe90774b78b248d61f7`.

This draft maps existing IELTS product truth into the `PRODUCT` authority domain of the adopted `agent-documents` V1 candidate. It does not create catalog identities, close documentation, or change product semantics.

## Objective

Build a complete, evidence-based IELTS learning system specification and product/runtime design for a self-directed learner with a concrete IELTS target. The product coordinates the path from target, through current evidence and explicit gaps or unknowns, to the next eligible learning action, truthful new evidence, and automatic re-planning.

The durable product value is this governed target-to-evidence-to-next-action loop. Individual AI models, SRS algorithms, media providers, evaluators, exercise formats, and realtime interaction mechanisms remain replaceable capabilities rather than the semantic center of the product.

The intended outcome is a governed route from a structured Band-3 entry range toward Band 9 across Listening, Reading, Writing, and Speaking while preserving uneven skill profiles, explicit unknown/stale/conflicting evidence, overall and per-skill target constraints, Academic and General Training differences, and a truthful distinction between learning progress, readiness, product support, and an actual external IELTS result.

The product may provide a strong governed route toward a target, but it must not imply that using the system guarantees a future external IELTS score.

**Sources:** `OBJECTIVE.md` — Purpose, Target outcome; `spec/00-PRODUCT.md` — Product identity, Promise boundary.

## Actors and roles

### Self-directed IELTS learner

The primary learner has a real IELTS target and uses the ordinary product route without mandatory teacher dependency. The learner declares the target and may later change it. Learner controls such as Swap, Skip, Shorten, or Change-skill may choose among eligible paths, but cannot erase Required prerequisites, unsupported target conditions, or uncovered product conditions.

**Sources:** `OBJECTIVE.md` — Purpose, Learner route; `spec/00-PRODUCT.md` — Product identity, Human-support boundary.

### AI support role

AI may support explanation, modelling, guided practice, feedback, tutoring, generation, analysis, evaluation, Speaking interaction, and media-related learning inside bounded roles. AI does not become IELTS truth, Assessment truth, Progression authority, or the standard against which learner competence is defined.

**Sources:** `OBJECTIVE.md` — Product experience; `spec/00-PRODUCT.md` — AI supports learning without becoming the standard, Human-support boundary.

### Optional human expert role

Human expert input may be used for preference or coaching, or may enter as external evidence when the canonical Assessment rules admit it. Mandatory human scoring is not a prerequisite for the ordinary supported learning route merely to manufacture a consequence that the product itself cannot safely support.

**Source:** `spec/00-PRODUCT.md` — Human-support boundary.

## Product-level features and capabilities

This is a migration-only product-level inventory, not a closed `FTR-*` catalog and not a second authoritative feature inventory. Exact feature inventories, practice-mode inventories, durations, UI defaults, and behavior remain owned by the current canonical `design/**` and `spec/**` owners until cutover.

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

No authority conflict was found between the consulted canonical PRODUCT owners for the five PRODUCT-domain concerns in this slice.

This draft intentionally does **not** establish a closed feature catalog, actor/role IDs, milestone coverage state, `DOCS_READY`, or canonical document references. `agent-documents` assigns those structural identities to `docs/catalog/project.json`, and `TASK-0001` explicitly forbids creating that catalog in this slice. Exact downstream feature inventories and behavior also remain with their current canonical owners until a later explicit migration/cutover task.

This is a bounded migration state, not a claim that documentation closure or implementation readiness has been established.

**Sources:** `.agent/tasks/TASK-0001/task.yaml@23fa05c8586a9f295a3c0fe90774b78b248d61f7`; `agent-documents@e6728594b0371e2e941c1457fc8efdc14a90deee` — `DOCUMENT_MODEL.md` and `templates/docs/PRODUCT.md`.

## Traceability summary

| PRODUCT concern | Current canonical source at the authorized base | Draft section |
| --- | --- | --- |
| objective | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Objective |
| actors / roles | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Actors and roles |
| features / capabilities | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Product-level features and capabilities |
| scope / non-goals | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Scope; Non-goals |
| domain / external constraints | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Domain and external constraints |

All target-source references in this table resolve to `phatnguyen03022001/ilets@23fa05c8586a9f295a3c0fe90774b78b248d61f7`. This traceability is migration evidence only and does not change source authority.