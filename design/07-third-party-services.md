STATUS: CANONICAL
OWNS: third-party capability inventory, provider lifecycle/selection rules, external-infrastructure applicability triggers, portability boundaries, external data-sharing constraints, provider failure/degradation semantics, and external-service activation requirements
DEPENDS_ON: ../CONSTITUTION.md, 03-media-youtube.md, 04-application-flows.md, 06-implementation-stack.md
DOES_NOT_OWN: provider legal terms themselves, learning/mastery semantics, public API wire shape, deployment implementation, pricing plans, historical provider candidates, or provider internal architecture

# Third-Party Services

## Purpose

Make external dependencies explicit so provider convenience cannot become product truth, data authority, unnecessary spend, or an irreversible architecture boundary.

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

The capability is named for what the product needs, not after the vendor currently considered or selected.

A product capability being required does **not** imply that an external provider is required. First-party/local/hosting-native execution remains eligible whenever it satisfies the same semantic, quality, security, reliability, and operational contract. No paid GitHub feature or paid external service is an architecture requirement merely because the underlying concern is required.

# Provider lifecycle

Every external capability/provider relationship is in exactly one of:

- `DEFERRED` — known possible external capability/provider route, but not required by the current implementation/release scope; no provider-selection work should occur until concrete external-route demand exists;
- `TBD` — an external provider route is required for the current implementation/release direction but no provider is selected;
- `CANDIDATE` — under evaluation, no activation authority;
- `SELECTED_FOR_IMPLEMENTATION` — approved implementation choice behind the declared boundary;
- `ACTIVATION_BLOCKED` — selected but release/legal/privacy/security/calibration gates are unresolved;
- `ACTIVE` — production-enabled under the current support declaration;
- `SUSPENDED` — temporarily disabled after a material gate/provider failure;
- `RETIRED` — no longer used for new work.

The lifecycle applies to the **external-provider relationship**, not to whether the underlying product capability exists. A first-party/local implementation may exist while an external provider route remains `DEFERRED`.

The inventory `External-provider status` column uses only the lifecycle tokens above. Technology choices, first-party availability, demand conditions, and implementation notes belong in the boundary/invariant column.

`DEFERRED` is not a missing-provider error. Promoting a route from `DEFERRED` to `TBD` requires an actual reason that an external route is needed for the scoped implementation/release; the existence of a capability, checklist item, or inventory row is not that reason.

Before `DEFERRED → TBD`, the implementation decision states why the existing first-party/local/selected route cannot or should not satisfy the applicable contract, considering quality, security, reliability, operational burden, portability, and total cost.

Historical candidates in `research/` or `archive/` have no provider status until explicitly adopted here.

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

Cost cannot rescue an ineligible provider. A cheaper fallback cannot silently lower the evidence/quality floor.

# External-resource economy

External services are used only for a capability that remains necessary after existing first-party/runtime resources are considered.

Preferred route:

```text
existing first-party or deterministic capability
        ↓
existing ACTIVE/eligible provider route
        ↓
new provider only when a real capability/quality/reliability gap remains
```

Rules:

1. do not send work to AI/STT/TTS or another paid provider when an exact deterministic/local path satisfies the same contract;
2. reuse a prior valid provider result when logical input, source state, model/policy version, and freshness make reuse correct;
3. deduplicate retries by logical work identity so network retry cannot create duplicate provider cost;
4. cache/derived output remains non-authoritative and is invalidated when relevant source/policy/model/version changes;
5. optional expensive work should be asynchronous, batched, rate-limited, or skipped before quality/evidence standards are lowered;
6. collect provider latency/usage/cost by capability so a new route is justified by measured value;
7. adding a second provider for one capability requires explicit reliability/quality/cost/exit benefit and cannot create two semantic owners;
8. provider optimization consumes the minimum necessary learner/source data.

# External infrastructure trigger rules

Required product/runtime concerns do not automatically require external infrastructure products. External routes remain `DEFERRED` until their trigger exists.

- **Identity / credential custody** — identity capability is required; external identity is conditional on the selected auth route. OAuth-specific flows exist only if OAuth is selected. JWT signing/rotation is conditional on JWT selection; any selected credential/session route still requires revocation/rotation semantics appropriate to that route.
- **Managed PostgreSQL** — PostgreSQL-compatible persistence semantics are selected, but managed hosting is conditional. External management is justified by operational/recovery/cost needs, not by the database concern itself.
- **Object storage** — external storage becomes required only when retained/large artifacts cannot be handled safely/practically by the selected deployment while preserving access, retention, backup, and privacy requirements.
- **CDN / edge caching** — conditional on eligible asset traffic/latency/geographic delivery need; never used to cache private/authoritative state incorrectly.
- **DNS / hosting / public edge** — concrete DNS/hosting is required only once a deployment/hostname exists. Hosting-native routing/TLS/abuse protection may satisfy the concern; no separate external gateway/load-balancer/WAF product is implied.
- **WAF / dedicated DDoS protection** — conditional on public-edge risk/traffic/deployment; provider-native protection may satisfy the concern when adequate.
- **Observability / error monitoring** — operational telemetry is required before support, but an external observability provider is optional; first-party/native logs/metrics may satisfy the contract.
- **Email** — external transactional email is conditional on a concrete account/product notification flow.
- **Provider callbacks / inbound webhooks** — conditional on a selected external provider requiring callback delivery. Callback authenticity, replay protection, idempotency, and authoritative work association follow `05-api.md` semantics.
- **AI / LLM / STT / TTS** — external routes remain optional behind bounded ports; use only when local/deterministic capability cannot satisfy the required quality/coverage/operational contract.
- **External feature flags** — conditional; bounded first-party flags/kill switches may satisfy safe rollout/degradation needs.
- **External queue/broker/PubSub** — conditional on measured dispatch reliability, throughput, fan-out, or isolation need that DB-backed durable work/outbox semantics cannot satisfy. Async work alone is not the trigger.
- **External API gateway/load balancer/reverse proxy** — conditional on actual deployment/public-edge/multiple-instance routing need, not on API existence itself.

# Mandatory portability boundaries

The first implementation treats these as mandatory strategic boundaries regardless of whether the first concrete route is local, self-hosted, or external:

1. **Database**;
2. **AI/model capability used for consequential evaluation, and for bounded generation/validation when those routes are implemented**;
3. **Identity**.

This does not make runtime content generation mandatory. When AI/model capability is used for evaluation, generation, or validation, provider-specific semantics do not become canonical product/content truth.

Object storage, email, analytics, observability, payments, hosting, edge/security, and other external capabilities should remain replaceable where practical, but the three above require explicit exit design from the beginning.

Portability does not require a generic multi-provider framework at bootstrap. For one concrete route, minimum sufficient shape is a narrow capability interface/port where a real substitution boundary exists, provider-independent internal identity/state, and credible export/exit path. Dynamic routing, provider registries, weighted failover, or multiple simultaneously active adapters require demonstrated need.

# Initial capability inventory

| Capability | Runtime owner | External-provider status | Required boundary/invariant |
|---|---|---|---|
| Identity / credential custody | Core API integration | `DEFERRED` | identity capability required, external custody not preselected; stable internal learner identity; provider ID never becomes learning identity; selected route must meet security/revocation/export requirements; OAuth/JWT mechanics only if selected |
| PostgreSQL-compatible structured persistence | Core API | `DEFERRED` | PostgreSQL semantics selected; local/self-managed/managed deployment remains implementation choice; migrations, backup/restore/PITR where applicable, provider exit where external |
| Object storage | Core API | `DEFERRED` | external object storage only when retained/large artifacts require it; private references, retention/access/backup policy |
| AI / LLM productive evaluation | Evaluator | `DEFERRED` | external model provider not pre-required; adapter/portability boundary; output Observation candidate, never certification |
| AI / LLM bounded content generation / model-assisted validation | Evaluator | `DEFERRED` | optional route only on concrete demand; candidate/signals non-authoritative; exact revision/provenance/policy identity preserved |
| Speech-to-text / acoustic analysis | Evaluator | `DEFERRED` | local/selected evaluator/external routes eligible; quality/provenance/uncertainty preserved |
| Text-to-speech / generated audio | content/media tooling | `DEFERRED` | quality/provenance fit; owned/licensed audio may satisfy demand without TTS |
| YouTube playback/metadata | Web + Core API | `SELECTED_FOR_IMPLEMENTATION` | eligible embed/Data API capability path selected; activation still requires live policy/product gates; no assumed arbitrary extraction |
| Transactional email | Core API | `DEFERRED` | only when concrete notification/account flow requires it; transport only |
| Product analytics | Web/Core API | `DEFERRED` | first-party minimal events may exist; no raw learner-content shadow database |
| Operational observability/error monitoring | all runtime units | `DEFERRED` | telemetry concern required before support, external provider optional; redact secrets/sensitive learner payloads |
| Payments / billing | Core API | `DEFERRED` | deferred until monetization requires it; entitlement cannot alter learning/evidence truth |
| Hosting / DNS / public edge | deployable owners | `DEFERRED` | provider/topology not selected; concrete deployment must satisfy TLS, routing, health, security, recovery, version visibility; separate gateway/LB/WAF not implied |
| CDN / edge caching | Web/asset delivery | `DEFERRED` | only when eligible asset traffic/latency justifies it; delivery optimization, never authority |
| WAF / dedicated DDoS service | public edge | `DEFERRED` | dedicated product only when risk/traffic/deployment justifies; hosting-native controls may satisfy concern |
| Provider callback / inbound webhook handling | Core API | `DEFERRED` | only for selected provider requiring callbacks; signature/auth, replay/idempotency, authoritative association required |
| External queue / broker / PubSub | Core API dispatch | `DEFERRED` | only after measured reliability/throughput/fan-out need beyond durable DB work/outbox; never business-state authority |
| Feature flags | Core API/Web | `DEFERRED` | first-party bounded flags/kill switches may exist; external service only on demonstrated need and never hidden policy engine |

Research may contain named vendor candidates. Canonical design does not repeat those names until a provider actually enters this lifecycle.

# Identity requirements

Regardless of implementation/provider route:

- credential custody stays outside learning-domain records;
- learning state references stable internal `learner_id`;
- guest→account merge requires explicit identity-safe semantics;
- same-email identities are not silently linked;
- account export/deletion capability exists before a support declaration that depends on it;
- privileged/admin access is a separate security boundary;
- session design preserves safe revocation without unnecessary provider coupling;
- least-privilege service/admin access and secret/key/token rotation/revocation match the actually selected credential mechanism.

Exact token/session durations are implementation/security policy, not canonical constants.

# Learner data + AI processor rules

When an external AI/data processor is used, default processor posture is:

```text
training/reuse of learner content by processor = prohibited unless explicitly approved
minimum necessary context                     = required
raw learner content in analytics               = prohibited
provider/model provenance                      = required where material
```

AI route selection follows privacy + semantic quality + reliability eligibility before cost/latency optimization.

# Audio/media privacy

Learner audio is ephemeral-by-default unless product purpose explicitly requires retention.

Prefer local preview where practical, temporary processing, persisted derived observations/provenance rather than permanent raw audio, explicit retention state when retained, user-visible capture disclosure, and deletion/backup behavior that prevents deleted media from silently returning to normal use after restore.

# Durable submission rule

A learner-visible accepted submission corresponds to durable authoritative product state before success acknowledgement.

```text
persist Attempt / authoritative work
      ↓
commit
      ↓
ACK accepted
      ↓
async evaluation may continue
```

An evaluator/provider outage may delay evaluation. It must not lose accepted work or fabricate a score.

# Retry + fallback

Third-party retry is bounded, idempotent/deduplicated, classified transient/permanent/ambiguous, observable, tied to one logical work identity, and cost-aware after correctness. Exponential backoff/jitter applies where retrying immediately would worsen a transient route; exact retry budgets remain implementation policy.

A timeout does not prove remote failure. Ambiguous outcomes remain unresolved until safe authoritative reconciliation.

Fallback is valid only when a pre-approved route meets the same applicable quality/privacy/security floor. Otherwise product state remains delayed/unavailable or requests appropriate re-evidence/replacement.

# Database/recovery baseline

Initial structured durable product state uses PostgreSQL semantics behind a replaceable implementation boundary.

Before production support, selected implementation demonstrates as applicable point-in-time recovery, independent logical export/backup, restore verification, safe migration discipline, provider exit/recovery test, and preservation of accepted learner work at application commit boundary.

Cache, broker, analytics store, search index, or vector store cannot become authority for learner/product state.

# Queue/broker rule

Asynchronous evaluation/content generation/model-assisted validation does not itself justify a broker.

Initial semantic direction:

```text
authoritative database state
+ durable work/outbox semantics
+ idempotent bounded execution
```

Dedicated dispatch infrastructure appears only after measured reliability/throughput/fan-out need. DLQ semantics, Pub/Sub, or event-driven topology are consequences of that selected dispatch architecture, not bootstrap requirements. Dispatch infrastructure never becomes business-state authority.

# Analytics + observability

Keep separate product analytics, service/operational telemetry, privileged/admin audit, security events, provider/evaluator/generator/validator latency/cost, and learner-visible accepted-work state. Retention may differ by class. Sensitive payloads/secrets are redacted by default. External observability remains optional.

# Payments + entitlements

Tiers may differ in volume, optional depth, frequency, or expensive-capability access. They may not change canonical learning truth or silently lower evidence/content quality for the same intended consequence.

Quota pressure may reduce optional work or delay noninteractive work. It cannot create a lower-quality scoring/content-validation route unless that route independently passes the same applicable eligibility floor.

# Provider failure product behavior

An external-provider failure maps to retrying, delayed, degraded-safe, temporarily unavailable, or approved equivalent fallback. It never becomes fake learner failure/mastery, silent content loss/activation, or silent quality downgrade.

# Activation gate

An external capability/provider may become `ACTIVE` only when the current product-support declaration confirms all applicable items:

- terms/licensing/rights compatibility;
- privacy/data-processing/learner-content reuse conditions;
- region/residency implications where material;
- security/secrets/credential boundaries;
- availability/rate limits/quotas;
- cost assumptions + kill switch;
- deletion/export behavior;
- backup/restore implications;
- callback/webhook authenticity/replay behavior where applicable;
- fallback/degraded behavior;
- exit/portability path;
- evaluator/content-generation/validation quality/calibration where applicable.

# Replacement invariant

Replacing an implementation/provider route must not require redefining Skill, Knowledge, Band, Assessment, Progression, feature IDs, practice-mode IDs, content semantic identity/revision rules, or learner identity.

If provider replacement changes canonical learning/product/content semantics, the provider boundary was incorrectly designed.
