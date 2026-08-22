STATUS: SUPPORTING
ROLE: FOUNDER DECISION CONSUMPTION EVIDENCE
AUTHORITY: NONE

# LenBands founder-decision closure

This register is provenance/evidence only. It is not an active architecture, product, learning, provider, or implementation owner.

Canonical owners win. Disposition text below does not redefine rules. Any change to active semantics must update the correct canonical owner first; this register merely records how the imported LenBands input was consumed.

The imported numbered register contains **325 decision rows** on this repository state. The V7 source table identifies 12 rows by boundary name rather than numeric suffix; this closure normalizes those row identities as `V7/<boundary>` solely to make one-to-one closure mechanically checkable. `content-rights-and-provenance.md` is reviewed separately after the 325 rows and is not included in that count.

Allowed dispositions: `ADOPTED`, `SUPERSEDED`, `REJECTED`, `DEFERRED`, `NOT_APPLICABLE`, `UNRESOLVED`.

## Platform and reliability
| ID | Disposition | Current canonical owner / trigger / superseding rule | Rationale |
|---|---|---|---|
| V1.1 | DEFERRED | design/08-coverage-and-support.md release market gate | Residency absence is not permanent policy; assess when a release market is declared |
| V1.2 | DEFERRED | design/08-coverage-and-support.md TargetSupportDeclaration | Vietnam launch scope is a release decision, not current semantic authority |
| V1.3 | ADOPTED | design/07-third-party-services.md selection order | Managed versus self-hosted remains an eligibility/quality/operations choice |
| V1.4 | SUPERSEDED | design/07-third-party-services.md provider lifecycle | No fixed hybrid-vendor strategy; each capability selects an eligible route |
| V1.5 | ADOPTED | design/06-implementation-stack.md deployment boundary | Compute provider is intentionally not inherited or preselected |
| V1.6 | ADOPTED | design/06-implementation-stack.md Go → PostgreSQL | PostgreSQL-compatible storage remains authoritative for product state |
| V1.7 | SUPERSEDED | design/04-application-flows.md + design/06 production gate | Accepted work durability remains; the historical recovery-minute target is not frozen |
| V1.8 | SUPERSEDED | design/07-third-party-services.md selection order | Cost minimization remains subordinate to eligibility; no near-zero constant is authority |
| V1.9 | SUPERSEDED | design/06-implementation-stack.md capacity gate | Capacity is measured for the intended release rather than frozen to learner counts |
| V1.10 | ADOPTED | design/07-third-party-services.md mandatory portability boundaries | Database, AI/model, and identity remain explicit portability boundaries |
| V1.11 | ADOPTED | design/07-third-party-services.md selection order | Privacy, quality, reliability, portability, then cost govern provider eligibility |
| V1.12 | ADOPTED | design/06-implementation-stack.md concern disposition | Bootstrap stays production-shaped without prebuilding untriggered infrastructure |
| V2.1 | ADOPTED | design/04-application-flows.md authoritative durable mutation | Durable success is acknowledged only after authoritative commit |
| V2.2 | ADOPTED | design/04-application-flows.md async work + design/07 fallback | AI failure preserves accepted work; fallback must independently meet the same floor |
| V2.3 | DEFERRED | design/06-implementation-stack.md deployment configuration | Environment topology is chosen when deployment is materialized |
| V2.4 | DEFERRED | design/06-implementation-stack.md production recovery gate | PITR is one possible recovery mechanism, not an architecture constant |
| V2.5 | SUPERSEDED | design/06-implementation-stack.md measurable operational objectives | Accepted-work durability remains stronger; numeric RPO bands require release evidence |
| V2.6 | ADOPTED | design/07-third-party-services.md learner data + AI processor rules | Learner audio is ephemeral by default unless purpose requires retention |
| V2.7 | DEFERRED | design/06-implementation-stack.md conditional infrastructure | IaC is introduced only when deployment complexity demonstrates the need |
| V2.8 | ADOPTED | design/07-third-party-services.md activation gate | Cheap hosting cannot bypass recovery, quality, privacy, or security gates |
| V2.9 | ADOPTED | design/07-third-party-services.md portability boundary | Portability does not require active multi-cloud or multi-provider deployment |
| V2.10 | ADOPTED | design/06-implementation-stack.md reuse-first invariant | Operational complexity is accepted only for a demonstrated requirement |
| V5.1 | ADOPTED | design/06-implementation-stack.md Go → authoritative PostgreSQL store | Core-owned PostgreSQL-compatible state is authoritative |
| V5.2 | ADOPTED | design/06-implementation-stack.md observability boundary | Analytics and telemetry remain derived and non-authoritative |
| V5.3 | SUPERSEDED | design/06-implementation-stack.md cache boundary | Redis is not selected; any cache is derived and cannot affect correctness |
| V5.4 | SUPERSEDED | design/04-application-flows.md async work invariant | Durable recoverable work is required without mandating a PostgreSQL outbox shape |
| V5.5 | ADOPTED | design/04-application-flows.md async retry rules | Eligible worker retries are bounded, recoverable, and state-aware |
| V5.6 | ADOPTED | design/04-application-flows.md async retry rules | Retry is classified, idempotent, bounded, and cost-aware |
| V5.7 | DEFERRED | design/06-implementation-stack.md object/media storage trigger | Object-storage topology waits for a real retained-artifact need |
| V5.8 | DEFERRED | design/06-implementation-stack.md production recovery gate | Backup mechanisms and cadence are release operational policy |
| V5.9 | DEFERRED | design/07-third-party-services.md managed DB activation gate | Independent backup failure-domain needs are resolved with the selected DB provider |
| V5.10 | ADOPTED | design/06-implementation-stack.md reliability/recovery gate | Restore must be verified; drill cadence remains operational policy |
| V5.11 | ADOPTED | design/07-third-party-services.md product analytics boundary | Analytics is minimized and raw learner content is not a shadow event store |
| V5.12 | SUPERSEDED | design/06-implementation-stack.md data lifecycle boundary | Retention is data-class/policy owned rather than a fixed raw-versus-derived duration |
| V5.13 | SUPERSEDED | design/04-application-flows.md authoritative async registration | Required work must be durable before publication without freezing one outbox technique |
| V5.14 | ADOPTED | design/06-implementation-stack.md deletion reconciliation fence | Restore must reconcile current deletion/tombstone truth before active use |
| V7/Identity | SUPERSEDED | design/07-third-party-services.md identity lifecycle | Auth0 is historical; identity custody remains provider-neutral and deferred |
| V7/API-evaluation-workers | SUPERSEDED | design/06-implementation-stack.md deployment boundary | Cloud Run/Singapore is not selected by current architecture |
| V7/Frontend | SUPERSEDED | design/06-implementation-stack.md deployment boundary | Cloudflare Workers/OpenNext is not current provider authority |
| V7/Region | SUPERSEDED | design/07-third-party-services.md activation gate | No compute/database region is selected before residency and release gates |
| V7/Postgres | SUPERSEDED | design/07-third-party-services.md managed PostgreSQL lifecycle | Neon is historical; PostgreSQL semantics remain while hosting is deferred |
| V7/Independent-backup | SUPERSEDED | design/06-implementation-stack.md recovery gate | Neon/R2-specific backup topology is replaced by provider-neutral verified recovery |
| V7/Object-storage | SUPERSEDED | design/07-third-party-services.md object-storage lifecycle | Cloudflare R2 is historical; object storage remains trigger-based |
| V7/Redis | SUPERSEDED | design/06-implementation-stack.md async/process correctness | Redis remains unselected, but PostgreSQL outbox is not the only legal mechanism |
| V7/Admin-auth | SUPERSEDED | design/04-application-flows.md privileged capabilities + design/05 auth gate | Admin access remains separated without freezing IdP or OTP details |
| V7/Staging | SUPERSEDED | design/06-implementation-stack.md deployment configuration | Scale-to-zero and DB-branch topology are deployment choices, not architecture truth |
| V7/IaC | SUPERSEDED | design/06-implementation-stack.md conditional infrastructure | OpenTofu/Terraform is not selected without a demonstrated deployment trigger |
| V7/Provider-rule | ADOPTED | design/07-third-party-services.md selection order | Provider choice uses measured eligibility and cost, not a fixed dollar cap |
| V8.1 | DEFERRED | design/06-implementation-stack.md production incident gate | Severity/escalation policy is required when operational support is materialized |
| V8.2 | SUPERSEDED | design/07-third-party-services.md observability lifecycle | PostHog/Sentry ordering is historical; external observability is unselected |
| V8.3 | DEFERRED | design/06-implementation-stack.md data lifecycle boundary | Log retention follows data class and operational policy, not a frozen adaptive schedule |
| V8.4 | ADOPTED | design/04-application-flows.md privileged mutation | Consequential privileged operations require reconstructable durable audit |
| V8.5 | SUPERSEDED | design/06-implementation-stack.md partial deployment policy | Staging/prod automation mode is not frozen before deployment strategy exists |
| V8.6 | DEFERRED | design/06-implementation-stack.md deploy recovery gate | Automatic versus manual rollback is implementation policy under safe recovery rules |
| V8.7 | ADOPTED | design/06-implementation-stack.md partial deployment/version skew | Schema changes must preserve the selected application/schema compatibility window |
| V8.8 | SUPERSEDED | design/06-implementation-stack.md production recovery gate | Restore/drill is required, but weekly/quarterly cadence is not canonical |
| V8.9 | ADOPTED | design/07-third-party-services.md fallback rule | Fallback exists only where an independently eligible route is justified |
| V8.10 | ADOPTED | design/07-third-party-services.md portability/exit gate | External routes require a credible provider exit path before activation |
| V8.11 | ADOPTED | design/06-implementation-stack.md feature flags | Material risky flags have an owner, safe default, and lifecycle |
| V8.12 | ADOPTED | design/06-implementation-stack.md feature flags + design/07 cost gate | Kill/degrade paths may protect cost and reliability without changing semantic truth |
| V8.13 | ADOPTED | design/04-application-flows.md failure-domain containment | Subsystem failure degrades honestly without inventing learner failure |
| V8.14 | ADOPTED | design/06-implementation-stack.md health/observability/incidents | Support requires operational recovery procedure and follow-up |
| V8.15 | ADOPTED | design/06-implementation-stack.md deletion boundary | Deletion/tombstone truth fences restore and late asynchronous completions |
| V8.16 | REJECTED | design/06-implementation-stack.md observability boundary | A dedicated internal status dashboard is not required; capability matters, not UI ceremony |

## Identity, privacy, and access
| ID | Disposition | Current canonical owner / trigger / superseding rule | Rationale |
|---|---|---|---|
| V3.1 | SUPERSEDED | design/05-api.md auth/session pre-contract gate | Email/password and social-login ordering are not selected before auth transport design |
| V3.2 | SUPERSEDED | design/07-third-party-services.md identity lifecycle | Auth0 credential custody is historical and not current provider authority |
| V3.3 | DEFERRED | design/05-api.md guest→account pre-contract decision | Guest trial and merge UX are decided when guest/account identity is implemented |
| V3.4 | ADOPTED | design/06-implementation-stack.md data lifecycle/deletion boundary | Deletion may retain only explicitly permitted legal/security/integrity history |
| V3.5 | DEFERRED | design/05-api.md identity/account product gate | Machine-readable learner export is decided with the concrete privacy/account lifecycle |
| V3.6 | SUPERSEDED | design/06-implementation-stack.md data lifecycle boundary | Retention is policy/data-class specific rather than fixed derived/raw durations |
| V3.7 | ADOPTED | design/07-third-party-services.md learner audio rule | Raw learner audio remains ephemeral by default |
| V3.8 | ADOPTED | design/07-third-party-services.md learner data + AI processor rules | External processor training/reuse is prohibited unless explicitly approved |
| V3.9 | ADOPTED | design/07-third-party-services.md analytics boundary | Analytics remains minimal and excludes raw learner-content shadow storage |
| V3.10 | DEFERRED | design/08-coverage-and-support.md rights/privacy/security gate | Minor support requires an explicit release scope plus legal/privacy/consent review |
| V4.1 | SUPERSEDED | design/05-api.md auth/session pre-contract gate | Mandatory email/password is not selected by current provider-neutral architecture |
| V4.2 | SUPERSEDED | design/07-third-party-services.md identity lifecycle | Auth0 is historical and remains unselected |
| V4.3 | DEFERRED | design/05-api.md auth/session pre-contract gate | Session behavior during IdP outage depends on the selected session/credential mechanism |
| V4.4 | DEFERRED | design/05-api.md auth/session pre-contract gate | Session lifetime is implementation/security policy, not a frozen 30-day constant |
| V4.5 | DEFERRED | design/05-api.md auth/session pre-contract gate | Global logout/revocation behavior is resolved when account sessions exist |
| V4.6 | DEFERRED | design/05-api.md auth/session pre-contract gate | Compromise/password-change revocation semantics depend on selected credentials |
| V4.7 | DEFERRED | design/05-api.md guest→account pre-contract decision | Explicit guest merge is required to be decided before that transition exists |
| V4.8 | DEFERRED | design/05-api.md guest→account pre-contract decision | Email collision/linking policy is deferred to identity materialization |
| V4.9 | DEFERRED | design/05-api.md auth/session pre-contract gate | Unverified-account capability limits are selected with the concrete identity flow |
| V4.10 | DEFERRED | design/05-api.md auth/session pre-contract gate | Password-reset OTP is mechanism-specific and not selected |
| V4.11 | ADOPTED | design/04-application-flows.md privileged operations | Admin capability/access is distinct from learner identity and cannot bypass Core |
| V4.12 | DEFERRED | design/05-api.md auth/session security decision | Default learner MFA policy requires the actual threat and credential model |
| V4.13 | DEFERRED | design/05-api.md auth/session pre-contract gate | Access-token lifetime is not canonical before token/session mechanism selection |
| V4.14 | DEFERRED | design/05-api.md auth/session pre-contract gate | Opaque rotating refresh-session design is an implementation choice at auth materialization |
| V4.15 | ADOPTED | design/07-third-party-services.md external identity boundary | Core-owned learner identity is stable and independent of provider identity |
| V4.16 | DEFERRED | design/05-api.md admin/session security decision | New-device OTP is mechanism-specific and deferred to admin auth implementation |

## AI, evaluation, and cost
| ID | Disposition | Current canonical owner / trigger / superseding rule | Rationale |
|---|---|---|---|
| V6.1 | SUPERSEDED | design/07-third-party-services.md provider lifecycle/selection order | Flash/model-tier names are historical; escalation remains provider-neutral and quality-first |
| V6.2 | ADOPTED | spec/08-ASSESSMENT.md uncertainty handling | Disagreement/uncertainty can trigger stronger evidence or evaluation handling by consequence |
| V6.3 | ADOPTED | design/04-application-flows.md asynchronous accepted work | AI work may be asynchronous while exposing truthful pending/unavailable states |
| V6.4 | SUPERSEDED | design/07-third-party-services.md external-resource economy | The 5–10× multiplier is not authority; expensive routes require justified marginal value |
| V6.5 | ADOPTED | spec/08-ASSESSMENT.md evaluator quality/uncertainty | High-risk or uncertain productive scoring requires stronger checking before high-consequence use |
| V6.6 | ADOPTED | spec/07-PRACTICE.md feedback policy | Feedback is staged/focused for action rather than an undifferentiated report |
| V6.7 | DEFERRED | design/07-third-party-services.md external-resource economy | Precompute/cache strategy is an optimization triggered by real reuse and correctness |
| V6.8 | ADOPTED | design/06-implementation-stack.md Python evaluator ownership | Speech evaluation is bounded, provenance-bearing, and not an opaque certification authority |
| V6.9 | DEFERRED | design/07-third-party-services.md payments/AI capability gates | Free-tier AI allowance belongs to future entitlement and sustainable-cost policy |
| V6.10 | ADOPTED | design/07-third-party-services.md selection order | Quota/cost pressure cannot lower the applicable evaluation quality floor |
| V6.11 | ADOPTED | design/04-application-flows.md content supply + spec/08 calibration | Batch/pre-generation and benchmark calibration are eligible where demand justifies them |
| V6.12 | ADOPTED | design/07-third-party-services.md selection order | Privacy, semantic quality, and reliability gate provider routing before cost |

## Product and requirements
| ID | Disposition | Current canonical owner / trigger / superseding rule | Rationale |
|---|---|---|---|
| V9.1 | ADOPTED | design/00-learning-experience.md Today-first flow | Home experience prioritizes the best current action |
| V9.2 | ADOPTED | design/01-skill-features.md feature model | IELTS skills and supporting language capabilities remain distinct |
| V9.3 | ADOPTED | design/02-practice-catalog.md learner-facing modes | Practice is presented by learning goal/interaction, not backend ontology |
| V9.4 | ADOPTED | design/01-skill-features.md Listening feature set | Listening combines task practice, targeted drills, and broader skill work |
| V9.5 | SUPERSEDED | design/07-third-party-services.md TTS capability lifecycle | No fixed TTS/human asset mix; audio route follows demand, rights, quality, and provenance |
| V9.6 | ADOPTED | spec/07-PRACTICE.md feedback policy | Listening feedback links observed error to explanation and another learning action |
| V9.7 | ADOPTED | design/02-practice-catalog.md Reading modes | Reading can progress from focused work to full timed performance |
| V9.8 | ADOPTED | design/08-coverage-and-support.md release ordering | Full families are modeled while executable coverage may close in honest waves |
| V9.9 | ADOPTED | spec/08-ASSESSMENT.md claim-scoped evidence | One Observation may support multiple scoped claims without task-label equivalence |
| V9.10 | ADOPTED | design/02-practice-catalog.md Writing modes | Writing supports component practice before independent full-task production |
| V9.11 | ADOPTED | spec/07-PRACTICE.md feedback focus | Writing criterion feedback is prioritized and action-producing |
| V9.12 | ADOPTED | spec/07-PRACTICE.md guided production/revision | Rewrite work creates fresh production instead of passive feedback consumption |
| V9.13 | ADOPTED | design/02-practice-catalog.md Speaking modes | Speaking progresses through parts/support drills to full mock when appropriate |
| V9.14 | SUPERSEDED | design/07-third-party-services.md learner audio rule | Replay/capture may exist, but download and retention are not mandatory; audio is ephemeral by default |
| V9.15 | ADOPTED | spec/08-ASSESSMENT.md calibration boundary | Pronunciation signals may be used only at calibrated, justified inference scope |
| V9.16 | SUPERSEDED | spec/04-KNOWLEDGE.md canonical Knowledge model | The historical Lexeme→Sense graph is replaced by the current canonical knowledge ontology |
| V9.17 | ADOPTED | spec/04-KNOWLEDGE.md + spec/07-PRACTICE.md | Grammar combines structured knowledge, error remediation, and contextual transfer |
| V9.18 | ADOPTED | design/00-learning-experience.md diagnostic flow | Short diagnostics create provisional state rather than complete learner truth |
| V9.19 | ADOPTED | design/02-practice-catalog.md assessment/readiness modes | Assessment granularity spans focused tasks through full mocks |
| V9.20 | ADOPTED | spec/08-ASSESSMENT.md evidence eligibility | Practice scores remain distinct from IELTS/Band evidence claims |
| V9.21 | ADOPTED | spec/09-PROGRESSION.md current learner state | Mastery/support is evidence- and context-derived, versioned, and uncertain |
| V9.22 | ADOPTED | spec/09-PROGRESSION.md staleness/review | Freshness is policy/evidence sensitive; no universal decay coefficient is frozen |
| V9.23 | ADOPTED | design/04-application-flows.md planner ranking | Plan ranking uses target urgency, evidence gaps, review, time fit, and prerequisites |
| V9.24 | ADOPTED | design/00-learning-experience.md learner agency | Recommendations preserve swap/skip/shorten/change controls within eligibility |
| V9.25 | ADOPTED | design/00-learning-experience.md progress presentation | Progress emphasizes blockers and next actions before decorative analytics |
| V9.26 | ADOPTED | design/00-learning-experience.md motivation boundary | Consistency support is non-punitive and does not make streaks learning truth |
| V9.27 | ADOPTED | design/04-application-flows.md review flow | Review preserves distinct knowledge, remediation, and re-evidence semantics |
| V9.28 | DEFERRED | design/07-third-party-services.md payments/entitlements trigger | A real free learning loop is an entitlement/release decision once monetization exists |
| V9.29 | DEFERRED | design/07-third-party-services.md payments/entitlements trigger | Premium packaging is deferred; entitlement may not redefine learning/evidence truth |
| V9.30 | ADOPTED | OBJECTIVE.md + design/04-application-flows.md | The product remains an adaptive evidence-driven IELTS learning loop |
| 10A.1 | ADOPTED | spec/00-PRODUCT.md + design/08-coverage-and-support.md | Academic and GT are modeled; release coverage can progress without semantic hard-coding |
| 10A.2 | ADOPTED | spec/01-LEARNER-MODEL.md TargetProfile | TargetProfile is a constraint set rather than one overall Band |
| 10A.3 | ADOPTED | spec/00-PRODUCT.md + design/08-coverage-and-support.md | Bands 3–9 are modeled while release support remains explicitly scoped |
| 10A.4 | ADOPTED | spec/03-SKILLS.md + spec/10-CONTENT-MODEL.md | Listening/Reading requirements derive from capabilities and contexts, not UI family labels |
| 10A.5 | ADOPTED | spec/03-SKILLS.md task-family boundary | Official task family is not automatically a competency |
| 10A.6 | ADOPTED | spec/03-SKILLS.md Writing leaves + spec/08-ASSESSMENT.md | Writing criteria are operationalized through observable scoped performance claims |
| 10A.7 | ADOPTED | spec/03-SKILLS.md Speaking leaves + spec/08-ASSESSMENT.md | Speaking Part context and criterion evidence remain separate dimensions |
| 10A.8 | ADOPTED | spec/07-PRACTICE.md timing boundary | Timing is introduced only when the learning/evidence construct requires it |
| 10A.9 | ADOPTED | spec/08-ASSESSMENT.md independence | Independent performance is required only for claims that require independence |
| 10A.10 | ADOPTED | spec/08-ASSESSMENT.md retry/transfer rules | Same-item retry is recovery; readiness/transfer requires appropriate fresh evidence |
| 10A.11 | ADOPTED | spec/08-ASSESSMENT.md receptive evidence | Practice correctness alone cannot establish receptive readiness |
| 10A.12 | ADOPTED | spec/08-ASSESSMENT.md readiness + spec/01 TargetProfile | Overall readiness follows target constraints and skill evidence, not average mastery |
| 10A.13 | ADOPTED | spec/03-SKILLS.md atomicity boundary | Requirements are kept at the smallest useful operational semantic unit |
| 10A.14 | ADOPTED | spec/05-BANDS.md source boundary | External Band descriptors inform construct truth; curriculum remains product-owned |
| 10A.15 | ADOPTED | spec/10-CONTENT-MODEL.md provenance + validation | Derived requirements/content retain provenance and validation state |

## Learning interventions
| ID | Disposition | Current canonical owner / trigger / superseding rule | Rationale |
|---|---|---|---|
| 10B.1 | ADOPTED | spec/07-PRACTICE.md gap-to-mechanism selection | Prerequisites are targeted only when evidence implicates them |
| 10B.2 | ADOPTED | spec/07-PRACTICE.md scaffold-fading mechanisms | Worked/controlled/guided/independent/transfer stages are available without mandatory ceremony |
| 10B.3 | ADOPTED | spec/07-PRACTICE.md retrieval/spacing | Suitable retained knowledge can use spaced retrieval |
| 10B.4 | ADOPTED | spec/07-PRACTICE.md spacing applicability | Spacing is limited to units for which repeated retrieval is semantically useful |
| 10B.5 | ADOPTED | spec/07-PRACTICE.md contrast mechanism | Discrimination gaps use contrast among plausible alternatives |
| 10B.6 | ADOPTED | spec/07-PRACTICE.md production progression | Production can move from recognition/control to independent unseen use |
| 10B.7 | ADOPTED | spec/07-PRACTICE.md scaffold-dependence handling | Support dependence is tracked and scaffolds are faded rather than mistaken for mastery |
| 10B.8 | ADOPTED | spec/07-PRACTICE.md transfer mechanism | Transfer uses varied/unseen contexts when the claim requires generalization |
| 10B.9 | ADOPTED | spec/08-ASSESSMENT.md evidence collection | Insufficient evidence triggers targeted diagnostic sampling rather than weakness inference |
| 10B.10 | ADOPTED | spec/08-ASSESSMENT.md conflict resolution | Conflicting evidence triggers discriminating re-evidence |
| 10B.11 | ADOPTED | spec/08-ASSESSMENT.md recency handling | Stale evidence triggers scoped refresh rather than assumed regression |
| 10B.12 | ADOPTED | spec/07-PRACTICE.md failure-specific remediation | Intervention failure is classified before changing the next action |
| 10B.13 | ADOPTED | design/04-application-flows.md planner ranking | Learner preference affects suitability/ranking, not ability truth |
| 10B.14 | ADOPTED | spec/07-PRACTICE.md diversity/transfer | Variation is added when it serves a learning or evidence purpose |
| 10B.15 | ADOPTED | spec/06-CURRICULUM.md adaptive sequencing | Strong evidence may skip or compress already-satisfied work |
| 10B.16 | ADOPTED | spec/07-PRACTICE.md decomposition | Persistent struggle may decompose the same target instead of changing standards |
| 10B.17 | ADOPTED | spec/07-PRACTICE.md retrieval mechanism | Active retrieval is preferred where retrieval is the intended mechanism |
| 10B.18 | ADOPTED | spec/07-PRACTICE.md review scheduling | FSRS-like scheduling is only appropriate for repeatable review units |
| 10B.19 | ADOPTED | spec/07-PRACTICE.md worked-example mechanism | Worked examples direct attention/comparison and fade when support is no longer needed |
| 10B.20 | ADOPTED | spec/07-PRACTICE.md contrast mechanism | Near alternatives are compared when discrimination is the learning bottleneck |
| 10B.21 | ADOPTED | spec/07-PRACTICE.md dictation boundary | Dictation targets decoding/detail/spelling rather than becoming generic Listening mastery |
| 10B.22 | ADOPTED | spec/07-PRACTICE.md shadowing boundary | Shadowing targets pronunciation/prosody/fluency rather than broad speaking certification |
| 10B.23 | ADOPTED | spec/07-PRACTICE.md controlled production | Controlled production bridges knowledge and freer output |
| 10B.24 | ADOPTED | spec/07-PRACTICE.md guided production | Guided support is explicitly faded toward independent performance |
| 10B.25 | ADOPTED | spec/07-PRACTICE.md revision practice | Rewrite/rerecord is used only when new production has learning value |
| 10B.26 | ADOPTED | spec/07-PRACTICE.md feedback focus | Feedback depth follows the diagnosed failure and next useful action |
| 10B.27 | ADOPTED | spec/08-ASSESSMENT.md inference scope | Extensive input can aid learning but is weak direct readiness evidence |
| 10B.28 | ADOPTED | spec/07-PRACTICE.md interleaving rule | Interleaving follows sufficient stability rather than replacing initial acquisition |
| 10B.29 | ADOPTED | spec/07-PRACTICE.md difficulty selection | Difficulty follows target/evidence and task conditions, not a Band label alone |
| 10B.30 | ADOPTED | spec/07-PRACTICE.md feedback timing | Feedback timing follows activity intent and evidence conditions |
| 10B.31 | ADOPTED | spec/08-ASSESSMENT.md assistance metadata | Hints/scaffolds remain visible to evidence eligibility and inference |
| 10B.32 | ADOPTED | spec/07-PRACTICE.md self-explanation mechanism | Self-explanation is used only where it improves the target learning process |
| 10B.33 | ADOPTED | design/00-learning-experience.md reflection boundary | Reflection remains lightweight and tied to meaningful learner events |
| 10B.34 | ADOPTED | spec/07-PRACTICE.md productive practice | Learner-generated output is used when feedback/evaluation is feasible |
| 10B.35 | ADOPTED | spec/07-PRACTICE.md timing boundary | Time pressure is introduced only for targets where timing is material |
| 10B.36 | ADOPTED | spec/08-ASSESSMENT.md readiness/re-evidence | Mocks are readiness/assessment work, not the default acquisition mechanism |
| 10B.37 | ADOPTED | spec/07-PRACTICE.md diversity rule | No arbitrary content-diversity quota is treated as learning truth |
| 10B.38 | ADOPTED | spec/07-PRACTICE.md mechanism selection | Multi-method sequences require a target-specific rationale, not assumed efficacy |
| 10B.39 | ADOPTED | spec/07-PRACTICE.md AI boundary | An AI tutor is an implementation surface, not a new Learning Mechanism ontology |
| 10B.40 | ADOPTED | design/00-learning-experience.md motivation boundary | Gamification may support experience but cannot establish mastery/evidence |

## Evidence and readiness
| ID | Disposition | Current canonical owner / trigger / superseding rule | Rationale |
|---|---|---|---|
| 10C.1 | ADOPTED | spec/08-ASSESSMENT.md Observation→EvidenceFact gate | An Observation is not automatically admitted evidence |
| 10C.2 | ADOPTED | spec/08-ASSESSMENT.md evidence eligibility | Eligibility considers task fit, assistance, timing, independence, quality, and provenance |
| 10C.3 | ADOPTED | spec/08-ASSESSMENT.md evidence interpretation | Positive and negative evidence are interpreted at claim scope |
| 10C.4 | ADOPTED | spec/08-ASSESSMENT.md retry history | Same-item retries preserve history rather than overwriting prior observations |
| 10C.5 | ADOPTED | spec/08-ASSESSMENT.md attribution | Evidence may support multiple claims without collapsing them |
| 10C.6 | ADOPTED | spec/08-ASSESSMENT.md task-family boundary | Task-family performance is not competency proof by label |
| 10C.7 | ADOPTED | spec/08-ASSESSMENT.md objective scoring boundary | Objective correctness does not by itself establish readiness relevance |
| 10C.8 | ADOPTED | spec/08-ASSESSMENT.md productive evaluator quality | Low-quality/uncertain W/S evaluation cannot support high-consequence claims |
| 10C.9 | ADOPTED | spec/08-ASSESSMENT.md activity/evidence boundary | Practice and assessment share evidence semantics but differ in configured conditions |
| 10C.10 | ADOPTED | spec/08-ASSESSMENT.md retry purpose | Immediate correction is recovery; delayed independent work can become new evidence |
| 10C.11 | ADOPTED | spec/08-ASSESSMENT.md same-item limitation | Same-item retest carries narrower inference because exposure changed |
| 10C.12 | ADOPTED | spec/08-ASSESSMENT.md unseen/transfer | Unseen material is required only where the claim needs transfer independence |
| 10C.13 | ADOPTED | spec/08-ASSESSMENT.md inference scope | Transfer evidence remains scoped to demonstrated contexts |
| 10C.14 | ADOPTED | spec/08-ASSESSMENT.md consistency | Consistency requirements are claim-specific; no universal attempt count is frozen |
| 10C.15 | ADOPTED | spec/08-ASSESSMENT.md historical evidence | Historical evidence remains while current support may become stale |
| 10C.16 | ADOPTED | spec/08-ASSESSMENT.md conflicts | Conflicting evidence is preserved rather than averaged away |
| 10C.17 | ADOPTED | spec/08-ASSESSMENT.md versioning | Evidence history and derived current interpretation remain version-aware |
| 10C.18 | ADOPTED | spec/08-ASSESSMENT.md policy evolution | Policy changes preserve observations/history and recompute derived decisions |
| 10C.19 | ADOPTED | spec/08-ASSESSMENT.md EvidenceRequirement | Sufficiency is logical claim policy rather than one weighted score |
| 10C.20 | ADOPTED | spec/08-ASSESSMENT.md claim-scoped dimensions | Evidence dimensions apply only when material to the claim |
| 10C.21 | ADOPTED | spec/08-ASSESSMENT.md behavior/requirement binding | Ability evidence maps to observable behavior/requirement scope |
| 10C.22 | ADOPTED | spec/08-ASSESSMENT.md unresolved mandatory conditions | Unknown required evidence conditions block support |
| 10C.23 | ADOPTED | spec/08-ASSESSMENT.md context applicability | Context requirements are enforced only where material |
| 10C.24 | ADOPTED | spec/08-ASSESSMENT.md independence applicability | Independence is mandatory only for claims that require it |
| 10C.25 | ADOPTED | spec/08-ASSESSMENT.md consistency applicability | Consistency is required only when the EvidenceRequirement says so |
| 10C.26 | ADOPTED | spec/08-ASSESSMENT.md recency applicability | Recency is claim/policy specific, not a universal time window |
| 10C.27 | ADOPTED | spec/08-ASSESSMENT.md transfer applicability | Broad claims require transfer evidence when specified |
| 10C.28 | ADOPTED | spec/08-ASSESSMENT.md inference scope | Transfer context and inference scope remain explicit |
| 10C.29 | ADOPTED | spec/08-ASSESSMENT.md evidence diversity | Diversity is required only where it changes claim validity |
| 10C.30 | ADOPTED | spec/08-ASSESSMENT.md negative evidence | Admissible negative evidence remains part of current interpretation |
| 10C.31 | ADOPTED | spec/08-ASSESSMENT.md conflict gate | Material unresolved conflict blocks a stronger supported claim |
| 10C.32 | ADOPTED | spec/08-ASSESSMENT.md status semantics | Missing evidence is INSUFFICIENT; demonstrated below-threshold is NOT_YET |
| 10C.33 | ADOPTED | spec/08-ASSESSMENT.md SUPPORTED semantics | SUPPORTED means current sufficient evidence, not a future guarantee |
| 10C.34 | ADOPTED | spec/08-ASSESSMENT.md scoped readiness | Readiness remains scoped to skill/target/profile conditions |
| 10C.35 | ADOPTED | spec/08-ASSESSMENT.md ReadinessSpecification | Readiness policy owns logical requirements rather than architecture-wide numbers |
| 10C.36 | ADOPTED | spec/08-ASSESSMENT.md evidence collection | Near-complete claims collect the smallest decision-relevant missing evidence |
| 10C.37 | ADOPTED | spec/08-ASSESSMENT.md provenance | Evidence and decisions retain reconstructable provenance |
| 10C.38 | ADOPTED | spec/08-ASSESSMENT.md recomputation | Current readiness/support can be recomputed from preserved facts and current policy |
| 10C.39 | ADOPTED | spec/08-ASSESSMENT.md re-evidence value | Retest is justified by decision value rather than ritual |
| 10C.40 | ADOPTED | spec/08-ASSESSMENT.md sampling economy | Assessment avoids collecting evidence once it no longer changes the decision |
| 10C.41 | ADOPTED | spec/08-ASSESSMENT.md stale refresh | Staleness triggers the smallest useful refresh sample |
| 10C.42 | ADOPTED | spec/08-ASSESSMENT.md conflict sampling | Conflict triggers the most discriminating useful task |
| 10C.43 | ADOPTED | spec/08-ASSESSMENT.md missing-evidence sampling | Missing evidence triggers the exact unresolved requirement where feasible |
| 10C.44 | ADOPTED | spec/08-ASSESSMENT.md remediation/re-evidence separation | Immediate remediation and later independent confirmation remain distinct |
| 10C.45 | ADOPTED | spec/08-ASSESSMENT.md unseen sampling | Unseen tasks are required only where transfer/independence calls for them |
| 10C.46 | ADOPTED | design/04-application-flows.md planner/evidence collection | Learner burden is considered after semantic/evidence eligibility |
| 10C.47 | ADOPTED | spec/08-ASSESSMENT.md exposure history | Exposure changes future admissibility without deleting historical observations |
| 10C.48 | ADOPTED | spec/08-ASSESSMENT.md retest timing | No universal cooldown is frozen |
| 10C.49 | ADOPTED | spec/08-ASSESSMENT.md information value | Diminishing information value stops or redirects repeated testing |
| 10C.50 | ADOPTED | spec/08-ASSESSMENT.md post-intervention sampling | Re-evidence targets the next unresolved claim after intervention |
| 10C.51 | ADOPTED | spec/08-ASSESSMENT.md staleness | Previously supported claims may later become stale without erasing history |
| 10C.52 | ADOPTED | spec/08-ASSESSMENT.md mock inference scope | Full mocks give broad readiness evidence but weak fine-grained attribution |
| 10C.53 | ADOPTED | design/00-learning-experience.md learner agency + spec/08 | Learners may request reassessment without bypassing evidence policy |
| 10C.54 | ADOPTED | spec/08-ASSESSMENT.md stopping rule | Sampling stops when uncertainty/coverage is resolved or further value is low |

## Learner experience
| ID | Disposition | Current canonical owner / trigger / superseding rule | Rationale |
|---|---|---|---|
| 10D.1 | ADOPTED | design/00-learning-experience.md first-run diagnostic | Unknown learners get a short provisional diagnostic and useful next action |
| 10D.2 | ADOPTED | design/00-learning-experience.md explanation policy | Recommendations include concise evidence-grounded rationale |
| 10D.3 | ADOPTED | design/00-learning-experience.md activity framing | Activities state goal, reason, and success condition before work |
| 10D.4 | ADOPTED | design/00-learning-experience.md result semantics | Activity success is not presented as mastery |
| 10D.5 | ADOPTED | design/00-learning-experience.md recovery UX | Failure leads to failure-specific recovery instead of generic repetition |
| 10D.6 | ADOPTED | design/00-learning-experience.md evidence-state UX | Insufficient evidence is explained as a targeted check, not learner weakness |
| 10D.7 | ADOPTED | design/00-learning-experience.md uncertainty UX | Conflicting evidence is shown as uncertainty plus a useful next check |
| 10D.8 | ADOPTED | design/00-learning-experience.md staleness UX | Staleness is explained as recency need, not assumed regression |
| 10D.9 | ADOPTED | design/00-learning-experience.md blocker presentation | The learner sees a small ranked set of actionable blockers |
| 10D.10 | ADOPTED | design/00-learning-experience.md presentation boundary | Displayed priority count is UX, not a canonical completeness score |
| 10D.11 | ADOPTED | design/00-learning-experience.md redirect explanation | When the plan changes target/action, the reason is surfaced |
| 10D.12 | ADOPTED | design/00-learning-experience.md scaffold UX | Support preserves the target while making scaffold dependence visible |
| 10D.13 | ADOPTED | design/00-learning-experience.md acceleration UX | Strong evidence may compress or skip already-satisfied work |
| 10D.14 | ADOPTED | design/00-learning-experience.md learner agency | Swap/skip/shorten/change controls remain available within eligibility |
| 10D.15 | ADOPTED | design/00-learning-experience.md learning/readiness distinction | Learning progress and exam-readiness evidence are presented separately |
| 10D.16 | ADOPTED | design/00-learning-experience.md promise boundary | Readiness language is cautious and never guarantees external IELTS outcome |
| 10D.17 | ADOPTED | design/00-learning-experience.md progress UX | Progress shows current state, blocker, change, and next action |
| 10D.18 | ADOPTED | design/04-application-flows.md no-eligible-action fallback | No eligible action yields a truthful blocker, not an invented exercise |
| 10D.19 | ADOPTED | spec/07-PRACTICE.md progressive feedback | Feedback can be staged to preserve attention and learning value |
| 10D.20 | ADOPTED | spec/07-PRACTICE.md feedback focus | The UI prioritizes actionable feedback rather than every detectable issue |
| 10D.21 | ADOPTED | spec/07-PRACTICE.md feedback focus | Not every error is surfaced when it would dilute the current learning target |
| 10D.22 | ADOPTED | design/00-learning-experience.md learner language | Observed performance is described without identity-level weakness labels |
| 10D.23 | ADOPTED | design/00-learning-experience.md session assembly | Sessions reflect time, intent, priority, prerequisites, and stopping conditions |
| 10D.24 | ADOPTED | design/00-learning-experience.md time-fit | Available time constrains plan packaging without redefining eligibility |
| 10D.25 | ADOPTED | design/00-learning-experience.md small-action UX | Short sessions still deliver one coherent useful action |
| 10D.26 | ADOPTED | design/00-learning-experience.md resume semantics | Resume respects activity/lifecycle/evidence state rather than blindly reopening work |
| 10D.27 | ADOPTED | design/04-application-flows.md DailyPlan provenance | Today plan remains stable unless material state changes justify replan |
| 10D.28 | ADOPTED | design/00-learning-experience.md continuity | Returning learners continue from current evidence/state rather than restart |
| 10D.29 | ADOPTED | design/00-learning-experience.md motivation | Motivation emphasizes visible competence, next action, and gentle consistency |
| 10D.30 | ADOPTED | design/00-learning-experience.md return flow | Missed days do not cause punitive reset |
| 10D.31 | ADOPTED | design/04-application-flows.md planner ranking | Skip behavior may inform friction/preferences but cannot become ability evidence |
| 10D.32 | ADOPTED | design/00-learning-experience.md struggle recovery | Repeated failure escalates recovery/decomposition and can stop unproductive repetition |
| 10D.33 | ADOPTED | design/00-learning-experience.md accessibility | Accessibility is a baseline interaction requirement |
| 10D.34 | ADOPTED | spec/08-ASSESSMENT.md assistance metadata | Accessibility/scaffold support is separated from claims that require independence |
| 10D.35 | ADOPTED | design/00-learning-experience.md speaking capture UX | Speaking permissions/retention are disclosed at the moment they matter |
| 10D.36 | ADOPTED | design/00-learning-experience.md notifications | Notifications are goal-supporting and learner-controllable |
| 10D.37 | ADOPTED | design/00-learning-experience.md notification truth | Notification claims must be grounded in actual learner/product evidence |
| 10D.38 | ADOPTED | design/04-application-flows.md failure semantics | Infrastructure failure is shown truthfully while preserving accepted work |
| 10D.39 | ADOPTED | design/04-application-flows.md async result delivery | AI delay is represented with coarse truthful pending/unavailable states |
| 10D.40 | ADOPTED | design/00-learning-experience.md completion UX | Success messaging explains result, meaning, and next action without overclaiming |

## Economics and entitlements
| ID | Disposition | Current canonical owner / trigger / superseding rule | Rationale |
|---|---|---|---|
| 10E.1 | ADOPTED | design/07-third-party-services.md selection order | Cost is considered only after semantic/quality/security/reliability eligibility |
| 10E.2 | ADOPTED | design/07-third-party-services.md external-resource economy | Usage/cost is measured by meaningful capability or logical work |
| 10E.3 | ADOPTED | design/07-third-party-services.md external-resource economy | Deterministic/local execution is preferred when it satisfies the same contract |
| 10E.4 | ADOPTED | design/07-third-party-services.md capability eligibility | AI is used where it adds required value rather than by default |
| 10E.5 | ADOPTED | design/07-third-party-services.md external-resource economy | Expensive escalation requires material risk/value justification |
| 10E.6 | ADOPTED | design/07-third-party-services.md selection order | Free and paid entitlement cannot change semantic/evidence quality floors |
| 10E.7 | ADOPTED | design/04-application-flows.md content supply | Validated reusable core content is preferred before new generation |
| 10E.8 | ADOPTED | design/04-application-flows.md planner/content supply | Personalization should primarily select/compose valid content and bounded feedback |
| 10E.9 | ADOPTED | design/06-implementation-stack.md cache boundary | Generic reusable results may be cached only when identity/freshness/access make reuse correct |
| 10E.10 | ADOPTED | design/04-application-flows.md retry rules | Retries are bounded and cannot duplicate accepted logical work |
| 10E.11 | ADOPTED | design/04-application-flows.md idempotent work identity | One accepted submission retains one logical evaluation lineage unless explicit reevaluation occurs |
| 10E.12 | ADOPTED | design/07-third-party-services.md external-resource economy | Budgets may be capability-specific without becoming semantic truth |
| 10E.13 | ADOPTED | design/07-third-party-services.md quota/degrade rule | Quota pressure sheds optional depth/volume/delay before quality |
| 10E.14 | ADOPTED | design/07-third-party-services.md activation gate | Cost anomalies require visibility plus bounded kill/degrade protection |
| 10E.15 | ADOPTED | design/07-third-party-services.md minimum-data egress | Model/provider calls send minimum necessary context |
| 10E.16 | ADOPTED | design/04-application-flows.md content supply | Generation can be progressive/batched rather than all-at-once runtime work |
| 10E.17 | ADOPTED | design/06-implementation-stack.md evaluator ownership | Speech evaluation may stage cheaper deterministic/acoustic work before costly routes |
| 10E.18 | ADOPTED | design/04-application-flows.md planner ranking | Among sufficient evidence options, lower burden/cost may rank higher |
| 10E.19 | ADOPTED | spec/08-ASSESSMENT.md human review boundary | Human review is an exception path when consequence/quality requires it |
| 10E.20 | DEFERRED | design/07-third-party-services.md payments/entitlements trigger | A meaningful free loop is a future entitlement/release decision |
| 10E.21 | DEFERRED | design/07-third-party-services.md payments/entitlements trigger | Premium packaging is future product economics and cannot redefine learning truth |
| 10E.22 | ADOPTED | design/07-third-party-services.md external-resource economy | Product value is not expressed as raw AI-token entitlement |
| 10E.23 | ADOPTED | design/08-coverage-and-support.md cost_abuse_operations | Support requires a sustainable measured execution path at the declared quality |
| 10E.24 | ADOPTED | design/08-coverage-and-support.md support gate | If economics cannot meet the required floor, redesign or withhold support rather than lower standards |

## Coverage and support
| ID | Disposition | Current canonical owner / trigger / superseding rule | Rationale |
|---|---|---|---|
| 10F.1 | ADOPTED | design/08-coverage-and-support.md TargetCoverageSpecification | Coverage is scoped to target/profile/claim conditions, not one scalar |
| 10F.2 | ADOPTED | design/08-coverage-and-support.md condition reduction | Required/alternative conditions are logical gates, not weighted percentages |
| 10F.3 | ADOPTED | design/08-coverage-and-support.md condition status enum | Coverage-condition status remains separate from target product status |
| 10F.4 | ADOPTED | design/08-coverage-and-support.md required conditions | Coverage traces the complete construct→content→runtime→operations path |
| 10F.5 | ADOPTED | design/08-coverage-and-support.md CoverageGap | Product coverage gaps remain distinct from learner gaps |
| 10F.6 | ADOPTED | design/08-coverage-and-support.md CoverageGap fields | Each gap records scope, failed condition, consequence, dependency, demand, provenance |
| 10F.7 | ADOPTED | design/08-coverage-and-support.md CoverageGap classes | Gap classes remain explicit without becoming learner taxonomy |
| 10F.8 | ADOPTED | design/08-coverage-and-support.md gap→implementation demand | Every blocking gap maps to an implementation demand class |
| 10F.9 | ADOPTED | design/08-coverage-and-support.md content closure | Content demand is derived from role/context/diversity/novelty/rights/evaluator needs |
| 10F.10 | ADOPTED | design/04-application-flows.md reuse-first content supply | Existing eligible content is reused before generating/importing new supply |
| 10F.11 | ADOPTED | design/08-coverage-and-support.md symmetry rule | Shared standards do not require equal asset volume across variants |
| 10F.12 | ADOPTED | design/08-coverage-and-support.md COVERED gate | MODELLED advances to COVERED only when every applicable non-validation condition passes |
| 10F.13 | ADOPTED | design/08-coverage-and-support.md calibration condition | Required evaluator/scoring calibration blocks COVERED only where the consequence needs it |
| 10F.14 | ADOPTED | design/08-coverage-and-support.md TargetSupportDeclaration | COVERED advances to supported only through an explicit versioned release declaration |
| 10F.15 | ADOPTED | design/08-coverage-and-support.md VALIDATED gate | SUPPORTED advances to VALIDATED only through scoped empirical evidence |
| 10F.16 | ADOPTED | design/08-coverage-and-support.md support blockers | Known release blockers remain explicit rather than hidden by aggregate completeness |
| 10F.17 | ADOPTED | design/08-coverage-and-support.md support revocation | Support is versioned and revocable when its gates change |
| 10F.18 | ADOPTED | design/08-coverage-and-support.md validation backlog | Nonblocking validation backlog is distinguished from support blockers |
| 10F.19 | ADOPTED | design/08-coverage-and-support.md coverage outputs | Coverage remains expressed through scoped specs/gaps/declarations, not completeness percentages |
| 10F.20 | ADOPTED | design/08-coverage-and-support.md technology independence | Coverage judges satisfied concerns rather than named infrastructure products |

## Rights/provenance closure — outside the 325 rows
`content-rights-and-provenance.md` was reviewed as a separate material input. These entries close its distinct durable rules without creating a second content-rights owner.

| ID | Disposition | Current canonical owner / trigger / superseding rule | Rationale |
|---|---|---|---|
| RIGHTS-1 | NOT_APPLICABLE | design/03-media-youtube.md + design/07-third-party-services.md | Spotify is not a current selected capability/provider, so the historical removal choice needs no active ban |
| RIGHTS-2 | ADOPTED | spec/10-CONTENT-MODEL.md validation + design/04 Flow K | Rights/provenance eligibility gates candidate admission and learner assignment |
| RIGHTS-3 | SUPERSEDED | spec/10-CONTENT-MODEL.md source/provenance representation | Historical FactBundle naming is unnecessary; durable source/provenance semantics remain |
| RIGHTS-4 | ADOPTED | design/04-application-flows.md Flow K + spec/10 ContentRevision | Generated or transformed content remains a candidate until immutable revision validation/release gates pass |
| RIGHTS-5 | ADOPTED | spec/10-CONTENT-MODEL.md content validation | Source integrity, rights/provenance, similarity/use scope, and quality are evaluated before consequential use |
| RIGHTS-6 | ADOPTED | spec/10-CONTENT-MODEL.md rights hard gate + design/03-media-youtube.md | Technical/public accessibility does not establish reuse eligibility; rights are checked independently |
| RIGHTS-7 | ADOPTED | design/03-media-youtube.md AI boundary + spec/10 validation | AI generation or transformation cannot override source rights/provenance eligibility |
| RIGHTS-8 | SUPERSEDED | spec/10-CONTENT-MODEL.md validation/release eligibility | Fixed GREEN/AMBER/RED labels are historical; current policy evaluates the actual intended-use rights condition |
| RIGHTS-9 | ADOPTED | spec/10-CONTENT-MODEL.md rights validation + design/07 data egress | Restricted, non-commercial, or unlicensed material cannot be used where the intended use lacks permission |
| RIGHTS-10 | ADOPTED | CONSTITUTION.md authority hierarchy + spec/10 + design/04 | External facts, source material, generated candidates, and product admission remain separate authorities |
| RIGHTS-11 | ADOPTED | spec/10-CONTENT-MODEL.md provenance + coverage manifest | Executable content retains traceable source/rights/validation provenance |

## Mechanical closure result

- expected numbered decisions: **325**
- dispositioned numbered decisions: **325**
- duplicate IDs: **0**
- missing IDs: **0**
- `UNRESOLVED`: **0**
- rights/provenance material rules reviewed separately: **11**
