STATUS: CANONICAL
OWNS: third-party capability inventory, provider lifecycle/selection rules, portability boundaries, external data-sharing constraints, provider failure/degradation semantics, and external-service activation requirements
DEPENDS_ON: ../CONSTITUTION.md, 03-media-youtube.md, 04-application-flows.md, 06-implementation-stack.md
DOES_NOT_OWN: provider legal terms themselves, learning/mastery semantics, public API wire shape, deployment implementation, pricing plans, historical provider candidates, or provider internal architecture

# Third-Party Services

## Purpose

Make external dependencies explicit so provider convenience cannot become product truth, data authority, or an irreversible architecture boundary.

Canonical shape:

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

# Provider lifecycle

Every external capability/provider relationship is in exactly one of:

- `DEFERRED` — known possible external capability/provider route, but not required by the current implementation/release scope; no provider-selection work should occur until concrete demand exists;
- `TBD` — an external provider route is required for the current implementation/release direction but no provider is selected;
- `CANDIDATE` — under evaluation, no activation authority;
- `SELECTED_FOR_IMPLEMENTATION` — approved implementation choice behind the declared boundary;
- `ACTIVATION_BLOCKED` — selected but release/legal/privacy/security/calibration gates are unresolved;
- `ACTIVE` — production-enabled under the current support declaration;
- `SUSPENDED` — temporarily disabled after a material gate/provider failure;
- `RETIRED` — no longer used for new work.

The lifecycle applies to the **external-provider relationship**, not to whether the underlying product capability exists. A first-party/local implementation may exist while an external provider route remains `DEFERRED`.

The inventory `Status` column uses only the lifecycle tokens above. Technology choices, first-party availability, demand conditions, and implementation notes belong in the boundary/invariant column; prose in `Status` is invalid.

`DEFERRED` is not a missing-provider error. Promoting a capability from `DEFERRED` requires an actual product/runtime/coverage need; the existence of a row in this inventory is not itself demand.

Historical candidates in `research/` or `archive/` have **no provider status** until explicitly adopted here.

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
2. reuse a prior valid provider result when the logical input, relevant source state, model/policy version, and freshness requirements make reuse correct;
3. deduplicate retries by logical work identity so network retry cannot create duplicate provider cost;
4. cache/derived output remains non-authoritative and is invalidated when the source/policy/model/version relevant to correctness changes;
5. optional expensive work should be asynchronous, batched, rate-limited, or skipped before quality/evidence standards are lowered;
6. collect provider latency/usage/cost by capability so a new route is justified by measured value rather than convenience;
7. adding a second provider for the same capability requires an explicit reliability/quality/cost/exit benefit and must not create two semantic owners;
8. provider optimization must consume the minimum necessary learner/source data.

# Mandatory portability boundaries

The first implementation treats these as mandatory strategic boundaries:

1. **Database**;
2. **AI/model evaluation**;
3. **Identity**.

Object storage, email, analytics, observability, payments, hosting, and other external capabilities should also remain replaceable where practical, but the three above require explicit exit design from the beginning.

Portability does **not** require a generic multi-provider framework at bootstrap. For one selected provider, the minimum sufficient shape is a narrow capability interface/port, one adapter, provider-independent internal identity/state, and a credible export/exit path. Dynamic routing, provider registries, weighted failover, or multiple simultaneously active adapters require a demonstrated need.

# Initial capability inventory

| Capability | Runtime owner | Status | Required boundary/invariant |
|---|---|---|---|
| Identity / credential custody | Core API integration | `TBD` | stable internal learner identity; provider ID never becomes learning identity |
| PostgreSQL-compatible structured persistence | Core API | `TBD` | PostgreSQL semantics are selected; provider remains unselected; migrations, PITR, logical export/restore, provider exit |
| Object storage | Core API | `DEFERRED` | external object storage is needed only when retained/large cross-unit artifacts require it; private references, retention/access policy, large artifacts outside normal JSON state |
| AI / LLM productive evaluation | Evaluator | `TBD` | adapter boundary; output is Observation candidate, never certification |
| Speech-to-text / acoustic analysis | Evaluator | `TBD` | quality/provenance/uncertainty preserved; a separate provider is not required when the selected evaluator route already satisfies the capability |
| Text-to-speech / generated audio | content/media tooling | `DEFERRED` | quality/provenance fit for intended learning role; owned/licensed audio may satisfy demand without TTS |
| YouTube playback/metadata | Web + Core API | `SELECTED_FOR_IMPLEMENTATION` | eligible embed/Data API path selected; activation still requires applicable live policy/product gates; no assumed arbitrary extraction |
| Transactional email | Core API | `DEFERRED` | external email is selected only when a concrete account/product notification flow requires it; transport only; provider templates own no product truth |
| Product analytics | Web/Core API | `DEFERRED` | first-party minimal structured events may exist without an external analytics provider; no raw learner-content shadow database |
| Error monitoring | all runtime units | `DEFERRED` | native/first-party logging may be used initially; any external route redacts secrets/sensitive learner payloads by default |
| Payments / billing | Core API | `DEFERRED` | external billing remains deferred until monetization scope requires it; entitlement boundary cannot alter learning/evidence truth |
| Hosting | deployable owners | `TBD` | an external hosting route is required for deployment but provider is unselected; hosting must not redefine deployable/service ownership; co-location is allowed by `06-implementation-stack.md` |
| CDN | Web/asset delivery | `DEFERRED` | external CDN remains deferred until traffic/latency/asset requirements justify it; delivery optimization only; never content authority |
| Feature flags | Core API/Web | `DEFERRED` | first-party bounded flags/kill switches may exist without an external provider; external flag service requires demonstrated need and never becomes a hidden policy engine |

Research may contain named vendor candidates. Canonical design must not repeat those names until a provider actually enters this lifecycle.

# Identity requirements

Regardless of provider:

- credential custody stays outside learning-domain records;
- learning state references stable internal `learner_id`;
- guest→account merge requires explicit identity-safe semantics;
- same-email identities are not silently linked;
- account export/deletion capability must exist before a support declaration that depends on it;
- privileged/admin access is a separate security boundary;
- session design must preserve safe revocation without requiring unnecessary provider coupling.

Exact token/session durations are implementation/security policy, not canonical architecture constants.

# Learner data + AI processor rules

Default processor posture:

```text
training/reuse of learner content by processor = prohibited unless explicitly approved
minimum necessary context                     = required
raw learner content in analytics               = prohibited
provider/model provenance                      = required for evaluator observations
```

AI route selection follows:

```text
privacy + semantic quality + reliability eligibility
                  ↓
             eligible routes
                  ↓
      cost / latency optimisation
```

# Audio/media privacy

Learner audio is ephemeral-by-default unless the product purpose explicitly requires retention.

Prefer:

- local preview before upload where practical;
- temporary processing;
- persisted derived observations/evidence/provenance rather than permanent raw audio;
- explicit retention state when raw recording is retained;
- user-visible microphone/upload disclosure;
- deletion/backup policy that prevents deleted media from silently returning to normal use after restore.

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

Third-party retry must be:

- bounded;
- idempotent;
- classified as transient/permanent/ambiguous;
- observable;
- tied to one logical work identity;
- cost-aware after semantic correctness.

Fallback is valid only when a pre-approved route meets the same applicable quality/privacy/security floor. Otherwise the product stays delayed/unavailable or requests re-evidence.

# Database/recovery baseline

Initial structured durable product state uses PostgreSQL semantics behind a provider boundary.

Before production support, the selected implementation must demonstrate:

- point-in-time recovery;
- independent logical export/backup;
- restore verification;
- safe migration discipline;
- provider-exit test;
- preservation of accepted learner work at the application commit boundary.

Cache, broker, analytics store, search index, or vector store cannot become authority for learner/product state.

# Queue/broker rule

Asynchronous evaluation does not itself justify a broker.

Initial semantic direction:

```text
authoritative database state
+ durable work/outbox semantics
+ idempotent execution
```

Introduce dedicated dispatch infrastructure only after a measured reliability/throughput need. Dispatch infrastructure never becomes business-state authority.

# Analytics + observability

Keep separate concerns for:

- product analytics;
- service/operational telemetry;
- privileged/admin audit history;
- security events;
- provider/evaluator latency/cost;
- learner-visible accepted-work state.

Retention may differ by class. Sensitive payloads/secrets are redacted by default.

# Payments + entitlements

Tiers may differ in volume, optional depth, frequency, or expensive-capability access. They may not change canonical learning truth or silently lower evidence quality for the same claim.

Quota pressure may reduce optional work or delay noninteractive work. It cannot create a lower-quality scoring route unless that route independently passes the same eligibility floor.

# Provider failure product behavior

A provider failure maps to an explicit state such as:

- retrying;
- delayed;
- degraded-safe;
- temporarily unavailable;
- approved equivalent fallback.

It never becomes:

- fake learner failure;
- fake mastery;
- silent content loss;
- silent quality downgrade.

# Activation gate

A capability/provider may become `ACTIVE` only when the current product-support declaration confirms all applicable items:

- terms/licensing/rights compatibility;
- privacy/data-processing conditions;
- learner-content training/reuse policy;
- region/residency implications;
- security/secrets;
- availability/rate limits/quotas;
- cost assumptions + kill switch;
- deletion/export behavior;
- backup/restore implications;
- fallback/degraded behavior;
- exit/portability path;
- evaluator quality/calibration where applicable.

# Replacement invariant

Replacing a provider must not require redefining Skill, Knowledge, Band, Assessment, Progression, feature IDs, practice-mode IDs, or learner identity.

If a provider replacement changes canonical learning/product semantics, the provider boundary was incorrectly designed.