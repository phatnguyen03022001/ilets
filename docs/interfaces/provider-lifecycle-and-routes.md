# PROVIDER-LIFECYCLE-ROUTES Provider lifecycle and selected external routes

## Provider boundary and lifecycle

External capability use remains behind a provider-neutral boundary:

```text
product capability
→ provider-neutral port / contract
→ provider adapter
→ external service
```

A required capability does not automatically require an external provider. First-party/local/hosting-native execution remains eligible when it meets the same semantics, quality, security, reliability and operations contract.

Every external capability/provider relationship is in exactly one lifecycle state:

- `DEFERRED` — external route is known but not required by current implementation/release scope.
- `TBD` — external route is required but no provider is selected.
- `CANDIDATE` — under evaluation; no activation authority.
- `SELECTED_FOR_IMPLEMENTATION` — approved implementation choice behind the declared boundary.
- `ACTIVATION_BLOCKED` — selected but release/legal/privacy/security/calibration gates remain unresolved.
- `ACTIVE` — production-enabled under the current support declaration.
- `SUSPENDED` — temporarily disabled after material gate/provider failure.
- `RETIRED` — not used for new work.

Lifecycle applies to the external relationship, not the underlying product capability. `DEFERRED` is not product failure. Moving `DEFERRED → TBD` requires a concrete reason an external route is needed. Historical research/candidates have no active state merely because they exist.

Selection order is strict: legal/rights/privacy → semantic+quality fit → security/data controls → reliability/recoverability → portability/exit → cost/latency/operational efficiency. Cost cannot rescue an ineligible route. New paid spend still requires USER authority and never bypasses privacy, rights, security, quality or support gates.

Before external data egress, the selected relationship must be eligible for the use; only minimum necessary data is sent; applicable rights/privacy/consent/retention conditions pass; consequential provider/model/tool provenance is retained; provider training/reuse/retention is compatible; secrets stay isolated; retry preserves one logical work identity; fallback independently meets the same floor. Provider output is an untrusted signal until Core/Assessment/content/product policy interprets it; HTTP success, provider confidence or dashboard state never directly creates evidence, certification, content activation, entitlement or product support.

## Initial capability inventory

| Capability | Runtime owner | Locked external-provider status | Boundary / consequence |
|---|---|---|---|
| Identity / credential custody | Core integration | `SELECTED_FOR_IMPLEMENTATION` | Clerk owns credential/auth/session mechanics only; internal identity, RBAC/capabilities, entitlement and product authorization remain Core-owned |
| Secret storage | runtime/deployment integration | `SELECTED_FOR_IMPLEMENTATION` | Google Secret Manager stores secret material only; typed runtime policy remains separate |
| PostgreSQL-compatible persistence hosting | Core | `SELECTED_FOR_IMPLEMENTATION` | hosting cannot change PostgreSQL/Core authority or application access ownership |
| Derived cache / rate limit / short coordination | Core / eligible edge-safe callers | `SELECTED_FOR_IMPLEMENTATION` | non-authoritative only; correctness survives loss/staleness |
| Bounded async task dispatch | Core dispatch | `SELECTED_FOR_IMPLEMENTATION` | transports durable Core-owned work; queue receipt/state is not business/evidence authority |
| Object storage | Core | `SELECTED_FOR_IMPLEMENTATION` | bytes external; Core owns object identity/lifecycle/access/retention/deletion/use eligibility |
| AI / LLM productive evaluation | Evaluator | `SELECTED_FOR_IMPLEMENTATION` | bounded adapters; minimum data; consequence-specific evaluator/model/config calibration; output never certification |
| AI / LLM generation / model-assisted validation | Evaluator | `DEFERRED` | optional demand-driven candidates/signals with provenance |
| Speech-to-text | Evaluator | `SELECTED_FOR_IMPLEMENTATION` | transcription only; preserves quality/provenance/uncertainty; never substitutes for acoustic evidence |
| Pronunciation / acoustic evaluation | Evaluator | `TBD` | no provider selected; applicable acoustic consequence remains unresolved/calibration-required |
| Realtime speech conversation | Web + bounded AI/audio capability | `SELECTED_FOR_IMPLEMENTATION` | optional tutor/Speaking interaction only; provider-neutral and non-authoritative for evidence |
| Pure VI↔EN translation | Evaluator | `SELECTED_FOR_IMPLEMENTATION` | assistance only; translation output does not establish learner ability |
| Text-to-speech / generated audio | content/media capability | `SELECTED_FOR_IMPLEMENTATION` | static/lesson audio when demand requires; normal content gates apply |
| YouTube playback/metadata | Web + Core | `SELECTED_FOR_IMPLEMENTATION` | supported embed/Data API path; no arbitrary extraction; activation still gated |
| Transactional email | Core | `SELECTED_FOR_IMPLEMENTATION` | transport only; delivery is not product-state authority |
| Product analytics | Web/Core | `SELECTED_FOR_IMPLEMENTATION` | minimized/classified events; no raw learner-content shadow authority |
| Operational observability | all runtime units | `SELECTED_FOR_IMPLEMENTATION` | non-authoritative operational evidence |
| Payments / subscription billing | Core | `SELECTED_FOR_IMPLEMENTATION` | external commercial observations only; effective entitlement requires Core reconciliation/commit |
| Deployable hosting | Web/Core/Evaluator | `SELECTED_FOR_IMPLEMENTATION` | one platform may host all units without collapsing logical/trust/persistence boundaries |
| DNS / CDN / baseline DDoS edge | public edge | `SELECTED_FOR_IMPLEMENTATION` | delivery/baseline edge protection only; no paid WAF bundle implied |
| Dedicated WAF / paid edge add-ons | public edge | `DEFERRED` | only after demonstrated need |
| Provider callbacks / webhooks | Core | `SELECTED_FOR_IMPLEMENTATION` | selected initially for payOS ingress only; provider auth/replay/association/Core reconciliation required |
| Additional queue / broker / PubSub | Core dispatch | `DEFERRED` | Cloud Tasks remains sole selected bounded dispatcher; no second broker retained |
| Feature flags | Core/Web | `DEFERRED` | bounded first-party flags initially sufficient; flags never hidden semantic authority |

Strategic portability boundaries remain Database, consequential AI/model capability, and Identity. Portability requires a narrow replaceable boundary and credible export/exit, not generic multi-provider machinery or multiple active adapters without demonstrated need.

## Selected AI, speech, and language routes

All routes below are `SELECTED_FOR_IMPLEMENTATION`, **not `ACTIVE`**.

| Capability/use | Primary selected route | Secondary route | Consequence |
|---|---|---|---|
| Text / productive grading | GPT-4o mini | DeepSeek V4 Flash — fallback / escalation | every consequential evaluator/model/configuration independently calibrated for intended consequence |
| Speech-to-text | ElevenLabs Scribe v2 | Gemini 3.5 Transcribe — fallback | realtime variant where realtime interaction needs it; batch where completed-audio work does not need realtime latency |
| Static / lesson TTS | Google Cloud WaveNet | Google Neural2 — fallback / higher-quality route | generated audio remains content/media output with normal quality/provenance/rights gates |
| Realtime VI↔EN tutor | Gemini 3.1 Flash Live | GPT-Realtime-2.1 Mini — fallback | interaction/tutoring only; not automatically Speaking examiner/evidence authority |
| Pure VI↔EN translation | GPT-Realtime-Translate | Gemini Live Translate — fallback | assistance does not establish language ability |
| Pronunciation / acoustic evaluation | `TBD` | none | no provider selected; consequences requiring acoustic judgment remain unresolved/calibration-required |

Primary/fallback/escalation/higher-quality labels describe intended routing only, not semantic interchangeability. A secondary route may serve a consequence only after independently satisfying that consequence's privacy/security/rights/quality/reliability and calibration requirements. Retry/fallback/escalation stays under the original logical operation's quota/cost admission.

Calibration for one productive/acoustic provider/model/configuration does not transfer to another provider, model, mode, prompt/rubric configuration or materially different version. Provider/model/config/prompt/rubric/evaluator provenance remains exact where consequential. STT is transcription and cannot replace acoustic pronunciation/prosody/intelligibility evidence. Realtime tutor quality does not make the route an examiner. Provider failure/timeout/degradation/fallback exhaustion is product/runtime state, not learner weakness, fake score or fabricated evidence.

## Selected infrastructure and commerce routes

These routes are also `SELECTED_FOR_IMPLEMENTATION`, not ACTIVE:

| Capability/use | Locked selected route | Authority boundary |
|---|---|---|
| External identity / session | Clerk | credential custody/auth/session issuance/revocation only; stable Core actor and product authorization remain Core-owned |
| Secret storage | Google Secret Manager | secret custody only; typed runtime policy and authorization remain Core-owned |
| PostgreSQL hosting | Neon Launch | hosts authoritative PostgreSQL-compatible store; only Core has application-runtime DB authority |
| Derived cache / rate-limit / short coordination | Upstash Redis PAYG | non-authoritative acceleration/protection/coordination only |
| Bounded async dispatch | Google Cloud Tasks | delivery infrastructure for durable Core work; task state/receipt never business authority |
| Object/media storage | Cloudflare R2 | bytes behind Core-owned identity/lifecycle/access/deletion; narrow signed browser transfer may be used where legal/safe and reconciled by Core |
| Deployable hosting | Google Cloud Run | hosts Web/Core/Evaluator while preserving logical/runtime/trust/machine-contract boundaries |
| DNS/CDN/baseline DDoS | Cloudflare | baseline public edge; no paid WAF/add-on set frozen |
| Transactional email | Amazon SES | transport only |
| Product analytics | PostHog | classified/minimized events, no raw learner-content authority |
| Operational logs/metrics | Google Cloud Logging / Monitoring | privacy-minimized operational telemetry only |
| Vietnam payments/subscription | payOS | payment/subscription observations only; Core owns effective entitlement transition |

For the selected public bearer integration, the Clerk Session token configuration must include a custom `aud` claim exactly equal to Core `CLERK_AUDIENCE`. Core fails closed if the claim is absent or mismatched. This wiring does not transfer product authorization authority to Clerk.

Provider callbacks/webhooks are initially selected only because payOS requires payment-event ingress. Callback receipt, valid signature, provider charge/subscription status or dashboard state remains external observed state until authenticated, associated and reconciled by Core.

The authority split remains:

```text
Clerk          = credential/auth/session mechanics, not product authorization
Secret Manager = secret custody, not runtime-policy authority
PostgreSQL     = authoritative durable product truth
Redis          = derived acceleration / bounded coordination
Cloud Tasks    = dispatch / execution delivery
R2             = bytes behind Core-owned object state
providers      = external capabilities / observations
Core           = product mutation + identity association + RBAC/entitlement/payment reconciliation authority
```

No Kubernetes, Kafka, service mesh, multi-region active-active, second queue, second cache, second initial payment provider or generic infrastructure abstraction is selected by the locked design.

## Deferred / unresolved routes

- Paddle — `DEFERRED` future Merchant-of-Record/tax-compliance candidate, not a second initial payment route.
- External feature-flag provider — `DEFERRED`; bounded first-party flags remain sufficient initially.
- Pronunciation/acoustic evaluator — `TBD`; higher-consequence acoustic evaluation remains calibration-required until an eligible calibrated route exists.
- Model-assisted content-generation provider — `DEFERRED` until separately established current implementation need.
- Standalone tool-heavy realtime route — not selected for the initial product.

## Activation, privacy, failure, and exit invariants

No selected route becomes `ACTIVE` until the current release/support declaration confirms all applicable terms/licensing/rights; privacy/data-processing/content-reuse; region/residency; security/secrets; minimum-necessary egress purpose; availability/rate limits/quotas; measured cost + kill/degrade path; deletion/export; backup/restore; callback authenticity/replay; fallback/degraded behavior; exit/portability; and evaluator/generation/validation quality/calibration gates.

For external AI/data processors, training/reuse of learner content is prohibited unless explicitly approved; minimum necessary context is required; raw learner content in analytics is prohibited; provider/model provenance is required where material. Learner audio is ephemeral by default unless product purpose explicitly requires retention, in which case retained audio has an explicit lifecycle and deletion/restore/late-result semantics.

Realtime conversation is optional and does not replace ordinary record→submit Speaking. It preserves bounded session admission, turn provenance/association, cancellation, reconnect/ambiguous outcome handling, idempotent logical work, partial-session reporting and graceful degradation. STT, acoustic analysis, conversational generation and realtime transport remain separable even when bundled by one provider.

Third-party retry is bounded, idempotent/deduplicated by logical work identity and distinguishes transient, permanent and ambiguous outcomes. Timeout does not prove remote non-execution. Fallback is legal only when a pre-approved route independently meets the same applicable quality/privacy/security floor; otherwise the product remains delayed/unavailable rather than fabricating a score or lowering standards.

Replacing a provider route must not require redefining Skill, Knowledge, Band, Assessment, Progression, feature/practice IDs, content identity/revision, learner identity or historical runtime meaning. Provider replacement preserves the provider-neutral contract and required provenance; if replacing a provider changes canonical semantics, the boundary was incorrectly designed.
