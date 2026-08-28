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

Historical candidates in `research/` or `archive/` have no active status until adopted here.

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

No model tool execution is selected by this architecture. If introduced later:

- prompt, learner, content, or provider text cannot grant tool authority;
- tool capabilities are explicitly allowlisted/bounded by application policy;
- network/file/secret/provider access is not inferred from model text;
- tool requests and outputs remain untrusted until application validation;
- model/tool execution cannot bypass Core authority, SSRF restrictions, provider eligibility, rights/privacy rules, or learner-data policy.

# External infrastructure triggers

External infrastructure products remain `DEFERRED` until their trigger exists.

- **Identity** — external custody only if selected auth route needs it.
- **Managed PostgreSQL** — managed hosting only if operations/recovery/cost justify it; PostgreSQL-compatible semantics remain selected independently.
- **Object storage** — external route only when retained/large artifacts require it beyond selected deployment capabilities.
- **CDN/edge cache** — eligible traffic/latency need only; never private/authoritative state authority.
- **Hosting/DNS/public edge** — concrete provider only once deployment/hostname exists; no separate gateway/LB/WAF implied.
- **WAF/DDoS service** — only if public-edge risk/deployment warrants a dedicated product.
- **Observability** — telemetry is required before support, external provider is optional.
- **Email** — only when a concrete product/account notification flow requires it.
- **AI/LLM/STT/TTS** — optional behind bounded capability ports when local/deterministic execution cannot satisfy required contract.
- **Feature-flag provider** — only if bounded first-party flags are insufficient.
- **Queue/broker/PubSub** — only after measured dispatch reliability/throughput/fan-out need beyond durable DB work/recoverable dispatch.
- **API gateway/load balancer/reverse proxy** — only when deployment/public-edge/multiple-instance routing needs it.
- **Callbacks/webhooks** — only for a selected provider requiring callback delivery.

# Mandatory portability boundaries

The first implementation treats these as strategic portability boundaries regardless of whether the first route is local/self-hosted/external:

1. Database;
2. AI/model capability used for consequential evaluation and optional bounded generation/validation when implemented;
3. Identity.

Portability does not require a generic multi-provider framework. Minimum shape is a narrow capability boundary where substitution is real, provider-independent product identity/state, and credible export/exit. Dynamic routing, weighted failover, and multiple active adapters require demonstrated need.

# Initial capability inventory

| Capability | Runtime owner | External-provider status | Required boundary/invariant |
|---|---|---|---|
| Identity / credential custody | Core API integration | `DEFERRED` | stable internal learner identity; external route not preselected; selected route supports appropriate revocation/export/security |
| PostgreSQL-compatible persistence hosting | Core API | `DEFERRED` | PostgreSQL semantics selected; hosting route remains deployment choice; provider exit/recovery if external |
| Object storage | Core API | `DEFERRED` | only when retained/large artifacts require it; private access, integrity, lifecycle, backup/orphan reconciliation |
| AI / LLM productive evaluation | Evaluator | `DEFERRED` | provider not pre-required; bounded adapter; minimum data; output never certification |
| AI / LLM bounded generation / model-assisted validation | Evaluator | `DEFERRED` | optional demand-driven candidate/signals; provenance and exact revision/policy identity preserved |
| Speech-to-text / acoustic analysis | Evaluator | `DEFERRED` | local/external routes eligible; quality/provenance/uncertainty preserved |
| Realtime speech conversation | Web + bounded AI/audio capability | `DEFERRED` | optional Speaking interaction only; provider-neutral; explicit latency/cost/capacity/privacy limits; output never examiner/evidence authority by itself |
| Text-to-speech / generated audio | content/media capability | `DEFERRED` | use only when content demand requires it and quality/provenance fit |
| YouTube playback/metadata | Web + Core API | `SELECTED_FOR_IMPLEMENTATION` | supported embed/Data API capability path; activation still requires live policy/product gates; no arbitrary extraction |
| Transactional email | Core API | `DEFERRED` | only for concrete notification/account flow; transport only |
| Product analytics | Web/Core API | `DEFERRED` | minimal classified events; no raw learner-content shadow authority |
| Operational observability | all runtime units | `DEFERRED` | telemetry concern required before support; external provider optional |
| Payments / billing | Core API | `DEFERRED` | only when product monetization actually requires it; entitlement cannot alter learning/evidence truth |
| Hosting / DNS / public edge | deployable owners | `DEFERRED` | concrete route must satisfy TLS/routing/health/security/recovery/version visibility |
| CDN / edge caching | Web/assets | `DEFERRED` | only eligible derived/static assets under privacy/rights/freshness rules |
| WAF / DDoS service | public edge | `DEFERRED` | dedicated service only when risk/deployment justifies it |
| Provider callbacks | Core API | `DEFERRED` | only selected provider callbacks; auth/replay/association/audit required |
| External queue / broker / PubSub | Core dispatch | `DEFERRED` | only beyond measured durable DB/recoverable-dispatch capability |
| Feature flags | Core API/Web | `DEFERRED` | external provider optional; flags never hidden policy authority |

Research may contain named candidates; canonical design does not repeat them until a provider enters this lifecycle.

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

Learner audio is ephemeral-by-default unless product purpose explicitly requires retention. Prefer temporary processing and persisted derived observations/provenance over permanent raw audio. Any retained audio has explicit lifecycle state and participates in deletion/backup reconciliation.

# Durable external work, retry, and fallback

Learner-visible accepted submission corresponds to durable authoritative Core state plus required recoverable continuation before acknowledgement. Provider dispatch itself may occur later and outside the transaction.

Third-party retry is bounded, deduplicated/idempotent, classified transient/permanent/ambiguous, observable, and tied to one logical work identity. A timeout does not prove remote failure or non-execution. Ambiguous outcomes remain unresolved until safe reconciliation.

Fallback is valid only when a pre-approved route independently meets the same applicable quality/privacy/security floor. Otherwise the product remains delayed/unavailable or requests appropriate replacement/re-evidence; it does not fabricate a score or lower standards.

# Database, queue, analytics, and observability provider boundaries

Managed DB/provider recovery must preserve the Core commit boundary, tested restore, migration discipline, and credible provider exit before support. Cache, broker, analytics, search/vector stores never become learner/product authority.

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