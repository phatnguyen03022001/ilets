STATUS: CANONICAL
OWNS: third-party capability inventory, provider lifecycle/selection rules, external-infrastructure applicability triggers, portability boundaries, external ingress/egress and data-sharing constraints, provider failure/degradation semantics, external processor lifecycle/deletion obligations, and external-service activation requirements
DEPENDS_ON: ../CONSTITUTION.md, 03-media-youtube.md, 04-application-flows.md, 06-implementation-stack.md
DOES_NOT_OWN: provider legal terms themselves, learning/mastery semantics, public API wire shape, deployment implementation, pricing plans, historical provider candidates, internal runtime-boundary topology, or provider internal architecture

# Third-Party Services

## Purpose

Make external dependencies explicit so provider convenience cannot become product truth, data authority, unnecessary spend, an implicit data-egress path, or an irreversible architecture boundary.

Canonical shape when an external service is used:

```text
product capability
      ↓
provider-neutral port / contract
      ↓
provider adapter
      ↓
external service
```

A required product capability does not imply a required external provider. First-party/local/hosting-native execution remains eligible when it satisfies the same semantic, quality, security, reliability, and operational contract. No paid external service is required merely because the underlying concern exists.

# Provider lifecycle

Every external capability/provider relationship is in exactly one of:

- `DEFERRED` — external route is known but not required by current implementation/release scope;
- `TBD` — an external route is required but no provider is selected;
- `CANDIDATE` — under evaluation, no activation authority;
- `SELECTED_FOR_IMPLEMENTATION` — approved implementation choice behind the declared boundary;
- `ACTIVATION_BLOCKED` — selected but release/legal/privacy/security/calibration gates remain unresolved;
- `ACTIVE` — production-enabled under the current support declaration;
- `SUSPENDED` — temporarily disabled after a material gate/provider failure;
- `RETIRED` — no longer used for new work.

The lifecycle applies to the external relationship, not the underlying capability. `DEFERRED` is not a product failure. Promoting `DEFERRED → TBD` requires a concrete reason an external route is needed rather than a checklist item or vendor preference.

Historical candidates in Git history and non-canonical `research/` material have no active status until adopted here.

# Selection order

Provider eligibility is evaluated in this order:

```text
legal / rights / privacy
        ↓
semantic + quality fit
        ↓
security + data controls
        ↓
reliability + recoverability
        ↓
portability + exit feasibility
        ↓
cost / latency / operational efficiency
```

Cost cannot rescue an ineligible route. A cheaper fallback cannot silently lower the applicable evidence/content/privacy/security floor.

A paid external provider/service may be evaluated, but it may not move to an implementation or activation state that creates project spend without explicit authorization from the repository's USER authority. Cost approval is necessary but not sufficient: it does not override privacy, rights, security, quality, reliability, or product-support gates. This applies to AI, STT/TTS, hosted databases, identity, observability, email, storage, queues, analytics, hosting, and any other external provider route.

# External-resource economy

Preferred route when outcomes are equivalent:

```text
existing deterministic / first-party capability
        ↓
existing ACTIVE eligible provider route
        ↓
new provider only for a demonstrated gap
```

Rules:

1. do not send work externally when an exact local/deterministic path satisfies the same contract;
2. reuse a valid prior result only when logical input, source state, model/policy version, access scope, and freshness make reuse correct;
3. deduplicate retry by logical work identity so ambiguous/repeated calls do not duplicate provider cost/work;
4. cache/provider output remains non-authoritative;
5. optional expensive work may be delayed/batched/rate-limited/omitted before standards are lowered;
6. measure latency/usage/cost by capability where provider operation is material;
7. a second provider needs explicit quality/reliability/cost/exit value and never creates a second semantic owner;
8. provider optimization still sends minimum necessary data.

# External trust and data boundaries

`06-implementation-stack.md` owns the generic runtime boundary. This file owns the detailed external-provider rules.

## External data egress

Before learner/source/content data leaves controlled runtime for AI/STT/TTS/identity/storage/another processor:

- the capability/provider relationship is selected and eligible for that use;
- only minimum necessary data for the declared purpose is sent;
- applicable rights/privacy/consent/retention conditions are satisfied;
- provider/model/tool/version provenance is retained where consequential;
- provider training/reuse/retention behavior is compatible with declared policy;
- credentials/secrets remain isolated from learner/product state and browser payloads;
- retry uses one logical work identity and handles ambiguous outcomes safely;
- fallback is used only when the alternate route independently meets the same applicable floor.

Provider output is an untrusted bounded observation/signal until owning Core/Assessment/content/product policy validates/interprets it. HTTP success, provider confidence, model self-assertion, or provider status does not create learner evidence, certification, content activation, or product support.

## External URL/media ingress

Learner/provider URLs are untrusted references, not authorization for arbitrary network access.

Any external resolver/fetch path constrains as applicable:

- supported source/provider and URL schemes;
- redirects/final destination;
- private/internal/metadata-network reachability;
- request/body/media size and execution time;
- metadata/content provenance;
- rights/source/product eligibility;
- permitted extraction/download behavior.

A resolvable URL does not imply scraping/download permission. Media-specific semantics remain subject to `03-media-youtube.md`.

## Provider callback / webhook ingress

Callbacks are conditional on a selected provider that actually requires them.

When present:

```text
provider callback
  ↓
provider authentication / signature verification
  ↓
replay + idempotency protection
  ↓
bounded structural validation
  ↓
authoritative work/event association
  ↓
safe freshness checks where material
  ↓
Core-owned transaction + durable audit/state
  ↓
bounded response
```

Callback timestamps are external observations, not causal identity. Callback/provider success cannot directly advance Progression, certify Band, activate content, or mutate evidence outside normal policy.

## External identity/session boundary

Identity capability is required; external identity is optional. Provider-specific OAuth/JWT mechanics exist only if selected. The auth/session transport decision required by `05-api.md` must precede exact public security-contract encoding.

Provider identity never becomes canonical learner identity; Core-owned stable learner identity remains the product reference. Credential custody, revocation/logout, guest→account association, admin/service separation, and key/token handling follow the selected mechanism without redefining learning state.

## External object/media processor boundary

External object/media storage or processing is eligible only behind Core-owned authoritative references/metadata and the data-lifecycle rules in `06-implementation-stack.md`. A storage/provider receipt does not activate content or mutate learner state.

If direct browser byte transfer is later used, provider grants remain narrow/temporary and Core must authorize and reconcile the transfer before the artifact becomes usable. Public anonymous permanence is not implied.

# External processor retention/deletion

The authoritative product retention/deletion decision is owned at the product/runtime data boundary in `06-implementation-stack.md`. For every applicable external processor/provider, this owner must establish before activation:

- what data/material may be retained or reused;
- how deletion/export obligations for that route are executed/reconciled;
- how ambiguous/provider-failed deletion is represented and retried safely;
- how provider-side copies/backups/derived artifacts are handled according to the declared applicable policy;
- how provider change/retirement prevents deleted/ineligible data from silently returning to active product use.

No retention duration is invented here. Provider deletion completion is an external fact requiring appropriate confirmation/reconciliation, not automatic product truth merely because a request was sent.

# Conditional AI/tool capability boundary

No standalone tool-heavy realtime capability/provider route is selected for the initial product, and no model tool execution is selected merely because a realtime model can support tools. If tool execution is independently required by a later canonical product capability, this existing boundary applies:

- prompt, learner, content, or provider text cannot grant tool authority;
- tool capabilities are explicitly allowlisted/bounded by application policy;
- network/file/secret/provider access is not inferred from model text;
- tool requests and outputs remain untrusted until application validation;
- model/tool execution cannot bypass Core authority, SSRF restrictions, provider eligibility, rights/privacy rules, or learner-data policy.

# External infrastructure triggers

External infrastructure products remain `DEFERRED` until a concrete implementation/release need exists. A USER-approved `SELECTED_FOR_IMPLEMENTATION` route below records that the initial implementation has chosen a provider for that bounded concern; it does not make the route `ACTIVE`, move authority out of Core/PostgreSQL, or authorize a broader infrastructure class.

- **Identity** — external custody only if selected auth route needs it.
- **Managed PostgreSQL** — selected only as hosting for the authoritative PostgreSQL-compatible store; database semantics and Core-only application access remain independent of vendor.
- **Derived cache / rate-limit / short-lived coordination** — eligible only for acceleration/protection/coordination that can fail or be rebuilt without becoming product truth.
- **Object storage** — eligible for retained/large artifacts while Core remains authoritative for identity, lifecycle, access, retention, deletion, and product usability.
- **CDN/edge** — eligible for DNS/static/derived delivery and baseline edge protection; never private/authoritative state authority.
- **Deployable hosting** — may host several logical units on one platform while preserving their caller/callee, trust, persistence, and machine-contract boundaries.
- **Dedicated WAF/paid edge feature set** — remains conditional on demonstrated risk/deployment need; selecting an edge provider does not preselect a paid WAF plan.
- **Observability** — telemetry is required before support; selected external delivery remains non-authoritative.
- **Email** — only for a concrete notification/account flow.
- **AI/LLM/STT/TTS** — optional behind bounded capability ports when local/deterministic execution cannot satisfy required contract.
- **Feature-flag provider** — only if bounded first-party flags are insufficient.
- **Bounded asynchronous task dispatch** — may deliver already-authoritative durable work; dispatch infrastructure never becomes business-state authority, and no second broker/queue is implied.
- **API gateway/load balancer/reverse proxy** — only when deployment/public-edge/multiple-instance routing needs it beyond selected hosting/edge capabilities.
- **Callbacks/webhooks** — selected only when an adopted provider route requires callback delivery.

# Mandatory portability boundaries

The first implementation treats these as strategic portability boundaries regardless of whether the first route is local/self-hosted/external:

1. Database;
2. AI/model capability used for consequential evaluation and optional bounded generation/validation when implemented;
3. Identity.

Portability does not require a generic multi-provider framework. Minimum shape is a narrow capability boundary where substitution is real, provider-independent product identity/state, and credible export/exit. Dynamic routing, weighted failover, and multiple active adapters require demonstrated need.

# Initial capability inventory

| Capability | Runtime owner | External-provider status | Required boundary/invariant |
|---|---|---|---|
| Identity / credential custody | Core API integration | `SELECTED_FOR_IMPLEMENTATION` | Clerk owns credential/auth/session mechanics only; stable internal actor/learner identity, RBAC/capabilities, entitlement, and product authorization remain Core-owned |
| Secret storage | runtime/deployment integration | `SELECTED_FOR_IMPLEMENTATION` | Google Secret Manager stores provider/API credentials and other secrets; Core typed runtime policy remains separate and no routine Admin plaintext-secret read/export is implied |
| PostgreSQL-compatible persistence hosting | Core API | `SELECTED_FOR_IMPLEMENTATION` | authoritative product truth remains PostgreSQL/Core-owned; provider hosting cannot change DB semantics or application access ownership |
| Derived cache / rate-limit / short-lived coordination | Core API / applicable edge-safe callers | `SELECTED_FOR_IMPLEMENTATION` | non-authoritative derived acceleration/coordination only; correctness survives cache loss/staleness under owning policies |
| Bounded asynchronous task dispatch | Core dispatch | `SELECTED_FOR_IMPLEMENTATION` | delivers durable Core-owned logical work; dispatch receipt/queue state never becomes business-state or evidence authority |
| Object storage | Core API | `SELECTED_FOR_IMPLEMENTATION` | bytes may live externally; Core owns object identity/lifecycle/access/retention/deletion; narrow direct browser transfer is preferred where legal/safe |
| AI / LLM productive evaluation | Evaluator | `SELECTED_FOR_IMPLEMENTATION` | selected routes remain behind bounded adapter; minimum data; consequential use requires evaluator/model/configuration-specific calibration; output never certification |
| AI / LLM bounded generation / model-assisted validation | Evaluator | `DEFERRED` | optional demand-driven candidate/signals; provenance and exact revision/policy identity preserved |
| Speech-to-text | Evaluator | `SELECTED_FOR_IMPLEMENTATION` | transcript capability only; realtime/batch form follows interaction need; quality/provenance/uncertainty preserved; transcript never substitutes for acoustic pronunciation evidence |
| Pronunciation / acoustic evaluation | Evaluator | `TBD` | no provider selected; required acoustic consequence remains unresolved / calibration-required until an eligible calibrated route exists |
| Realtime speech conversation | Web + bounded AI/audio capability | `SELECTED_FOR_IMPLEMENTATION` | optional Speaking/tutor interaction only; provider-neutral; explicit latency/cost/capacity/privacy limits; output never examiner/evidence authority by itself |
| Pure VI↔EN translation | Evaluator | `SELECTED_FOR_IMPLEMENTATION` | translation assistance only; provider-neutral; translation output does not become Speaking/Writing evidence or learner ability truth |
| Text-to-speech / generated audio | content/media capability | `SELECTED_FOR_IMPLEMENTATION` | static/lesson audio only when content demand requires it and quality/provenance fit |
| YouTube playback/metadata | Web + Core API | `SELECTED_FOR_IMPLEMENTATION` | supported embed/Data API capability path; activation still requires live policy/product gates; no arbitrary extraction |
| Transactional email | Core API | `SELECTED_FOR_IMPLEMENTATION` | concrete transport route only; email delivery/receipt is not product-state authority |
| Product analytics | Web/Core API | `SELECTED_FOR_IMPLEMENTATION` | minimal classified events; no raw learner-content shadow authority |
| Operational observability | all runtime units | `SELECTED_FOR_IMPLEMENTATION` | logs/metrics are operational evidence only; telemetry delivery never becomes business truth |
| Payments / subscription billing | Core API | `SELECTED_FOR_IMPLEMENTATION` | provider facts remain external observations; effective entitlement is committed only by Core reconciliation |
| Deployable hosting | Web / Core API / Evaluator | `SELECTED_FOR_IMPLEMENTATION` | one hosting platform may run all three units while logical/runtime/trust/persistence boundaries remain intact |
| DNS / CDN / baseline DDoS edge | Web/assets/public edge | `SELECTED_FOR_IMPLEMENTATION` | DNS/static/derived delivery and baseline edge protection only; no paid WAF/plan feature set frozen |
| Dedicated WAF / paid edge add-ons | public edge | `DEFERRED` | only when risk/deployment proves the additional capability is needed |
| Provider callbacks / webhooks | Core API | `SELECTED_FOR_IMPLEMENTATION` | selected only for the payOS payment route initially; auth/replay/association/audit and Core reconciliation required |
| Additional external queue / broker / PubSub | Core dispatch | `DEFERRED` | Google Cloud Tasks is the sole selected bounded dispatcher; no second queue/broker is retained |
| Feature flags | Core API/Web | `DEFERRED` | bounded first-party flags are sufficient initially; external flags never hidden policy authority |

Research may contain named candidates; canonical design does not repeat them until a provider enters this lifecycle.

## Initial selected provider routes

The USER-approved initial external routes below are `SELECTED_FOR_IMPLEMENTATION`, not `ACTIVE`. Selection authorizes implementation behind the existing provider-neutral capability boundary only. Normal activation gates still apply independently to every provider/model/configuration/use.

| Capability / use | Primary selected route | Secondary selected route | Lifecycle / consequence |
|---|---|---|---|
| Text / productive grading | GPT-4o mini | DeepSeek V4 Flash — fallback / escalation | both selected for implementation; every consequential evaluator/model/configuration version requires independent calibration for its intended consequence |
| Speech-to-text | ElevenLabs Scribe v2 | Gemini 3.5 Transcribe — fallback | both selected for implementation; use realtime variants where realtime interaction requires them and batch variants for completed-audio work that does not require realtime latency |
| Static / lesson text-to-speech | Google Cloud WaveNet | Google Neural2 — fallback / higher-quality route | both selected for implementation; generated audio remains content/media output with normal quality/provenance/rights gates |
| Realtime VI↔EN tutor | Gemini 3.1 Flash Live | GPT-Realtime-2.1 Mini — fallback | both selected for implementation; this is interaction/tutoring capability, not automatically a Speaking examiner or evidence authority |
| Pure VI↔EN translation | GPT-Realtime-Translate | Gemini Live Translate — fallback | both selected for implementation; translation assistance does not establish learner language ability by itself |
| Pronunciation / acoustic evaluation | `TBD` | none | no provider selected; any consequence needing acoustic judgment remains unresolved / calibration-required until a supported evaluator path exists |

## Initial selected infrastructure and commerce routes

These USER-approved initial routes are also `SELECTED_FOR_IMPLEMENTATION`, not `ACTIVE`. They use the existing runtime/provider boundaries rather than creating a generic infrastructure abstraction.

| Capability / use | Initial selected route | Selection boundary |
|---|---|---|
| External identity / session | Clerk | credential custody, authentication, and session issuance/revocation mechanics only; external principal maps to a stable Core actor and never owns RBAC/entitlement/product authorization |
| Secret storage | Google Secret Manager | provider/API credentials and other secrets only; typed runtime operating policy remains Core-owned and deployment environment holds references/bootstrap wiring rather than policy authority |
| PostgreSQL-compatible hosting | Neon Launch | hosts the authoritative PostgreSQL-compatible product store; only Core has application-runtime DB authority |
| Derived cache / rate-limit / short-lived coordination | Upstash Redis PAYG | non-authoritative acceleration/protection/short coordination only; Redis loss/staleness cannot redefine product/evidence/entitlement truth |
| Bounded asynchronous task dispatch | Google Cloud Tasks | delivery/execution infrastructure for durable Core-owned work; task state/receipt is never business-state authority |
| Object / media storage | Cloudflare R2 | stores bytes behind Core-owned object identity/lifecycle/access/retention/deletion; prefer narrow signed direct browser byte transfer where legal/safe and Core reconciles completion/integrity |
| Deployable hosting | Google Cloud Run | selected for Web, Go Core API, and Python Evaluator; co-hosting platform choice does not collapse logical/runtime/trust/machine-contract boundaries |
| DNS / CDN / baseline DDoS edge | Cloudflare | selected baseline edge route; no paid plan, dedicated WAF feature set, or permanent add-on bundle is frozen here |
| Transactional email | Amazon SES | transport capability only; provider delivery facts do not become learner/product authority |
| Product analytics | PostHog | classified/minimized product events only; no raw learner-content shadow authority |
| Operational logs / metrics | Google Cloud Logging / Monitoring | operational telemetry only; non-authoritative and privacy-minimized under the existing observability boundary |
| Vietnam payments / subscription billing | payOS | payment/subscription observations only; effective entitlement changes only through authenticated/associated Core reconciliation and authoritative commit |

For the selected public bearer integration, the Clerk Session token configuration must include a custom `aud` claim equal to the Core `CLERK_AUDIENCE` value. Core fails closed when that audience claim is missing or mismatched; this deployment wiring does not transfer product authorization to Clerk.

Provider callbacks/webhooks are selected initially only because the payOS route requires payment-event ingress. Callback receipt, signature validity, provider charge/subscription status, or dashboard state remains external observed state until Core associates and reconciles it under the existing payment/entitlement rules.

The authority split remains:

```text
Clerk        = credential/auth/session mechanics, not product authorization authority
Secret Manager= secret custody, not runtime-policy authority
PostgreSQL   = authoritative durable product truth
Redis        = derived acceleration / bounded coordination
Cloud Tasks  = dispatch / execution delivery
R2           = bytes behind Core-owned object state
providers    = external capabilities / observations
Core         = product mutation + identity association + RBAC/entitlement/payment reconciliation authority
```

No Kubernetes, Kafka, service mesh, multi-region active-active, second queue, second cache, second initial payment provider, or generic infrastructure abstraction is selected by this pass.

## Initial production planning envelope

The selected provider set starts with operator-adjustable runtime defaults owned by `06-implementation-stack.md`:

```text
public API planning volume    ≈ up to 1.5 million requests / month
operating target              = USD 10 / month
variable-spend safety ceiling = USD 20 / month
```

The request volume remains a planning input, not a capacity/performance guarantee or autoscaling constant. The operating target guides forecasting/alerts/quota and optional-capability optimization. The safety ceiling is application-level admission control for **new discretionary variable-cost work**; it is not a provider invoice guarantee. Fixed infrastructure, metering/reporting delay, already admitted/in-flight work, estimation error, and provider billing behavior can produce a final invoice above the configured ceiling. Cloud/provider budget alerts and dashboards are observations/warnings, not concurrency-safe enforcement guarantees.

Selection and implementation should make the target plausible through usage-based/scale-to-zero services where appropriate, bounded autoscaling, application-level cost/quota admission, semantically valid reuse/caching/batching/pre-generation, and explicit capacity/cost/latency/reliability upgrade triggers. Public API traffic remains distinct from provider invocation volume—one request does not imply one model call.

Concurrency-safe reservation/admission/reconciliation and budget-state degradation are implementation invariants owned by `06-implementation-stack.md`. Provider routes may expose billing observations, quotas, alerts, or hard provider controls, but those signals do not replace Core's logical-operation cost authority. Cost pressure may delay/deny optional work or preserve unresolved state; it never lowers privacy, evidence, evaluator, or content-quality standards.

This operating policy does not override the normal activation gate. A route whose privacy, rights, security, reliability, deletion, calibration, or support requirements fail remains blocked even if it would be cheaper.

Selected-route invariants:

1. For routes with `primary`, `fallback`, `escalation`, or `higher-quality route` roles, those labels describe intended routing only; they do not establish semantic interchangeability. A secondary route is usable for a consequence only when it independently satisfies that consequence's privacy/security/rights/quality/reliability and, where applicable, calibration requirements. Retry/fallback/escalation remains under the original logical operation's quota/cost admission rather than obtaining fresh allowance.
2. A provider/model/configuration change preserves exact provider/model/configuration/prompt/rubric/evaluator provenance where consequential. Calibration for one productive/acoustic evaluator route does not transfer automatically to another provider, model, mode, prompt/rubric configuration, or materially different version; cheaper or newer alone is insufficient for consequential promotion, and promotion evidence is consequence-proportional.
3. STT is transcription. A transcript, even from a selected realtime route, cannot substitute for acoustic evidence required for pronunciation/prosody/intelligibility inference.
4. Realtime tutor output is interaction content. It gains Assessment consequence only through the normal separately eligible Attempt/Observation/evaluator path; provider conversation quality does not make the route an examiner.
5. Provider failure, timeout, degraded service, or fallback exhaustion remains product/runtime state and never becomes learner weakness, a fake score, or fabricated evidence.
6. No selected route becomes `ACTIVE` until the normal activation gate passes for its exact use. Selection alone does not approve data egress, learner-content reuse, release support, or production traffic.
7. The initial route keeps only the approved primary and secondary provider where one is named. A third provider is not retained merely for hypothetical redundancy; adding one requires demonstrated quality/reliability/cost/exit value under the normal lifecycle.
8. Mandatory shadow/canary pipelines are not required for every low-consequence provider change; verification and promotion evidence remain consequence-proportional.
9. Pricing, latency, benchmark, quota, and quality figures remain mutable research/operational evidence. They are not canonical constants or guarantees in this owner.

## Deferred / unresolved initial routes

- Paddle — `DEFERRED` future candidate for international Merchant-of-Record / tax-compliance expansion; it is not a second initial payment route and has no activation authority;
- external feature-flag provider — `DEFERRED`; bounded first-party flags remain sufficient initially;
- pronunciation / acoustic evaluator — `TBD`; applicable higher-consequence acoustic evaluation remains calibration-required until an eligible provider/evaluator path is selected and calibrated;
- model-assisted content-generation provider — `DEFERRED` unless a separately owned current implementation scope establishes a concrete requirement;
- standalone tool-heavy realtime route — not selected for the initial product.

# Learner data + AI processor rules

When an external AI/data processor is used:

```text
training/reuse of learner content by processor = prohibited unless explicitly approved
minimum necessary context                     = required
raw learner content in analytics               = prohibited
provider/model provenance                      = required where material
```

AI route selection follows privacy + semantic quality + reliability eligibility before cost/latency optimization.

Realtime conversation is not required for the ordinary Speaking route. It may be activated only when the product intentionally offers that interaction and the selected local/external capability satisfies the same privacy, semantic-boundary, failure, and cost-control rules. Product entitlement may gate access, but payment/entitlement never upgrades evaluator or authorization authority.

A realtime capability must support the product failure semantics in `04-application-flows.md`: bounded session admission, turn provenance/association, cancellation, reconnect/ambiguous-outcome handling, idempotent logical work, partial-session reporting, and graceful degradation without converting provider latency/drop into learner performance. Provider transport/streaming technology remains replaceable and is not selected here.

STT, acoustic analysis, conversational generation, and realtime transport are separable capabilities even when one provider bundles them. Failure of one does not silently authorize another signal as equivalent evidence. Transcript disagreement with retained/available audio remains uncertainty; provider transcript confidence is not acoustic truth.

Learner audio is ephemeral-by-default unless product purpose explicitly requires retention. Prefer temporary processing and persisted derived observations/provenance over permanent raw audio. Any retained audio has explicit lifecycle state and participates in deletion/backup reconciliation. External processors receive only the minimum audio/text/turn context needed for the admitted capability; processor retention/reuse/training must satisfy the declared policy before activation.

Optional high-cost realtime use may be entitlement-gated, rate-limited, capacity-bounded, suspended, or degraded while ordinary record→submit Speaking remains available. Cost/quota pressure may reduce availability, not evaluator/evidence/privacy quality for the same intended consequence.

Payment/billing provider output is external observed state. A callback, charge status, renewal flag, or provider-side subscription label cannot directly grant/revoke effective product entitlement or terminate accepted learner work; Core authenticates, associates, reconciles, and commits the effective entitlement transition under `04-application-flows.md`. Ambiguous provider state remains pending/reconcilable rather than being guessed.

# Durable external work, retry, and fallback

Learner-visible accepted submission corresponds to durable authoritative Core state plus required recoverable continuation before acknowledgement. Provider dispatch itself may occur later and outside the transaction.

Third-party retry is bounded, deduplicated/idempotent, classified transient/permanent/ambiguous, observable, and tied to one logical work identity. A timeout does not prove remote failure or non-execution. Ambiguous outcomes remain unresolved until safe reconciliation.

Fallback is valid only when a pre-approved route independently meets the same applicable quality/privacy/security floor. Otherwise the product remains delayed/unavailable or requests appropriate replacement/re-evidence; it does not fabricate a score or lower standards.

# Database, queue, analytics, and observability provider boundaries

Managed DB/provider recovery must preserve the Core commit boundary, tested restore, migration discipline, and credible provider exit before support. Where provider-account/control-plane loss or suspension is within the supported recovery envelope, that restore/exit evidence cannot depend exclusively on continued access to the failed/suspended provider control plane; the implementation may use the smallest sufficient PostgreSQL-native export/backup/restore path. This does not require a second live database, multi-cloud replication, or automatic cross-provider failover. Cache, broker, analytics, search/vector stores never become learner/product authority.

Async evaluation/generation/validation alone does not justify a broker. Dedicated dispatch infrastructure appears only after measured need; it never becomes business-state authority.

Product analytics, operational telemetry, privileged audit, security events, provider/evaluator cost/latency, and learner-visible accepted-work state remain separate classes with appropriate access/retention rules. External observability remains optional.

# Provider failure product behavior

Provider failure maps to bounded retry, delayed, degraded-safe, temporarily unavailable, or approved equivalent fallback. It never becomes fake learner weakness/mastery, silent content activation/loss, or silent quality downgrade.

# Activation gate

An external capability/provider may become `ACTIVE` only when the current release/support declaration confirms applicable:

- terms/licensing/rights compatibility;
- privacy/data processing/content reuse conditions;
- region/residency implications where material;
- security/secrets/credential boundary;
- minimum-necessary data and egress purpose;
- availability/rate limits/quotas;
- measured cost assumptions + kill/degrade path;
- deletion/export behavior;
- backup/restore implications;
- callback authenticity/replay behavior where applicable;
- fallback/degraded behavior;
- exit/portability path;
- evaluator/generation/validation quality/calibration where applicable.

# Replacement invariant

Replacing an external implementation/provider route must not require redefining Skill, Knowledge, Band, Assessment, Progression, feature/practice IDs, content semantic identity/revision, learner identity, or historical runtime meaning.

Provider/model/version changes preserve provider-neutral contract/provenance and compatibility required by consumers. If provider replacement changes canonical semantics, the provider boundary was incorrectly designed.