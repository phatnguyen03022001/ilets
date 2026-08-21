STATUS: CANONICAL
OWNS: third-party capability inventory, provider-selection/activation rules, portability boundaries, external data-sharing constraints, provider failure/degradation semantics, and initial external-service requirements
DEPENDS_ON: ../CONSTITUTION.md, 03-media-youtube.md, 04-application-flows.md, 06-implementation-stack.md
DOES_NOT_OWN: provider legal terms themselves, learning/mastery semantics, public API wire shape, deployment implementation, pricing plans, or a provider's internal architecture

# Third-Party Services

## Purpose

Make every external dependency explicit before implementation so provider convenience cannot silently become product truth, data authority, or an irreversible architecture boundary.

A third-party service is selected for a capability. The capability is not named after the provider.

```text
product capability
      ↓
provider-neutral port / contract
      ↓
provider adapter
      ↓
external service
```

## Provider lifecycle

Every external capability uses one of these states:

- `TBD` — capability is required but no provider is selected;
- `CANDIDATE` — provider is being evaluated and has no activation authority;
- `SELECTED_FOR_IMPLEMENTATION` — approved implementation choice behind the declared boundary;
- `ACTIVATION_BLOCKED` — selection exists but release/privacy/security/legal/calibration gates are unresolved;
- `ACTIVE` — production-enabled under the current support declaration;
- `SUSPENDED` — temporarily disabled because a gate or provider condition failed;
- `RETIRED` — no longer used for new product work.

Historical provider choices imported from LenBands research are not automatically selected here.

## Selection order

A provider is eligible only after hard requirements pass:

```text
legal / rights / privacy
        ↓
semantic and quality fit
        ↓
reliability / recoverability
        ↓
security / data controls
        ↓
portability / exit feasibility
        ↓
cost and operational efficiency
```

Cost cannot rescue an ineligible provider. A cheaper fallback cannot silently reduce the scoring/evidence quality floor.

## Mandatory portability boundaries

The initial product treats these boundaries as strategically portable:

1. **Database**;
2. **AI / model evaluation**;
3. **Identity**.

Object storage, email, analytics, observability, payments, and hosting should also use replaceable interfaces where practical, but the three boundaries above require explicit exit design from the first implementation.

# Initial service inventory

| Capability | Runtime owner | Provider status | Required boundary / rule |
|---|---|---|---|
| Identity / credential custody | Go Core API integration boundary | `TBD` | internal stable `learner_id`; provider IDs never become learning identity |
| PostgreSQL | Go Core API | technology `SELECTED_FOR_IMPLEMENTATION`; provider `TBD` | standard PostgreSQL; canonical structured product state; migrations + PITR + export/exit |
| Object storage | Go Core API | `TBD` | private object storage; explicit retention/access namespaces; large audio/media by reference |
| AI / LLM evaluation | Python Evaluator | `TBD` | provider adapter; model result produces observations, never direct certification |
| Speech-to-text / speech features | Python Evaluator | `TBD` | staged speech pipeline; provenance and uncertainty preserved |
| Text-to-speech / generated audio | content/media tooling | `TBD` | may supplement practice; benchmark/high-value assessment content must meet stronger provenance/quality rules |
| YouTube playback/metadata | Web + Core API | YouTube `SELECTED_FOR_IMPLEMENTATION` for eligible embed/source use | IFrame/Data API compliance; no assumed arbitrary caption/audio extraction |
| Transactional email | Core API | `TBD` | verification/reset/notification only; no domain truth in provider templates |
| Product analytics | Web/Core API | `TBD` | structured minimal events; no raw learner-content firehose; analytics never owns product facts |
| Error monitoring | all runtime units | `TBD` | operational diagnostics; redact learner content/secrets by default |
| Payments / billing | Core API | `TBD` | entitlement adapter; billing state cannot alter learning truth or evidence quality |
| Hosting / CDN | deployable owners | `TBD` | deployable portability; provider runtime must not redefine service boundaries |
| Feature flags | Core API/Web | `TBD` or first-party | limited risky-feature flags and emergency kill switches; not a general policy engine |

## Historical candidates from LenBands

The imported 325-decision research includes historical directions such as Auth0, Neon PostgreSQL, Cloud Run, Cloudflare R2, PostHog, and a Cloudflare/OpenNext frontend direction. These are **research candidates/provenance only** for this repository.

They may be selected only after the current IELTS provider gate above is evaluated. No historical provider name is production activation authority.

# Identity requirements

Regardless of provider:

- credentials are managed through a dedicated identity boundary rather than stored in learning-domain tables;
- learning state references stable internal `learner_id`;
- guest use may be supported, but guest→account merge requires explicit confirmation;
- account deletion/export paths must exist before public support claims that require them;
- same-email identities are not silently linked;
- privileged/admin access is a separate security boundary;
- valid application sessions should not require a live IdP round-trip on every request when safe revocation semantics permit otherwise.

Exact token/session durations are implementation/calibration policy and are not frozen here.

# Learner-data and AI processor rules

Default external-processing policy:

```text
processor training/reuse of learner content = prohibited unless explicitly approved
minimum necessary context                 = required
raw learner content in analytics           = prohibited
provider/model provenance                  = required for evaluation observations
```

AI routing may escalate to stronger/more expensive models only when uncertainty, risk, or incremental value justifies it.

Provider routing is:

```text
privacy + quality + reliability eligibility
           ↓
eligible route set
           ↓
cost / latency optimization
```

Never invert this ordering.

# Audio and media privacy

Learner audio is **ephemeral-by-default**.

The product should prefer:

- local preview/replay before upload where practical;
- temporary server processing;
- persisted derived observation/evidence/provenance rather than permanent raw audio;
- explicit retention state when an activity genuinely requires retained recording;
- user-visible disclosure before microphone/upload use.

Deletion and backup/recovery policy must prevent deleted media from silently returning to normal product use after restore.

# Durable submission rule

A learner-visible successful submission must correspond to durable authoritative product state.

For accepted Writing/Speaking attempts:

```text
persist Attempt / authoritative work state
        ↓
commit
        ↓
ACK success to learner
        ↓
async evaluation may continue
```

An evaluator/provider outage must not lose accepted work. It may move evaluation into a truthful delayed/retry state.

# Retry and fallback

Third-party retries must be:

- bounded;
- idempotent;
- classified as transient / permanent / ambiguous;
- cost-aware;
- tied to one evaluation/work identity;
- observable.

Fallback is allowed only to a pre-approved route satisfying the same minimum semantic/quality/privacy requirements for the claim. Otherwise the system delays, requests re-evidence, or enters a degraded state instead of fabricating a lower-quality result.

# Database and recovery baseline

PostgreSQL is the selected canonical structured-store technology for initial implementation.

Required operational properties before production support:

- point-in-time recovery;
- independent logical export/backup outside the primary database provider boundary;
- restore verification;
- expand/contract-safe migration discipline;
- provider exit test;
- accepted learner submissions preserved at the application commit boundary.

A cache, broker, analytics store, or vector store may never become authority for canonical learner/product state.

# Queue / broker rule

Do not provision Redis, Kafka, or another broker merely because evaluation is asynchronous.

Initial direction remains:

```text
authoritative PostgreSQL state
+ durable evaluation-work/outbox semantics
+ idempotent worker execution
```

A dedicated broker is introduced only after measured throughput/reliability requirements justify it. The broker remains dispatch infrastructure, never business-state authority.

# Analytics and observability

Product analytics uses the minimum structured event set required to answer product questions. It must not become a shadow learner-content database.

Operational observability should separate:

- application/service telemetry;
- privileged admin audit history;
- security events;
- provider/evaluator cost and latency;
- user-facing accepted-work state.

Retention can differ by class. Secrets and sensitive learner payloads are redacted by default.

# Payments and entitlements

Free and paid product tiers may differ in volume, depth, personalization, or expensive-evaluation access. They may not differ in semantic truth or minimum evidence/quality standards for a claim.

Quota pressure may:

- reduce optional explanation depth;
- delay noninteractive work;
- limit volume/frequency;
- use a cheaper route only if it remains fully eligible.

Quota pressure may not silently lower IELTS scoring/evidence integrity.

# Provider failure product behavior

A provider failure is mapped to one of:

- retrying;
- delayed;
- degraded-safe;
- temporarily unavailable;
- fallback on an approved equivalent route.

It never becomes:

- fake learner failure;
- fake mastery;
- silent content loss;
- silent quality downgrade.

# Activation checklist

A third-party capability may become `ACTIVE` only when the current support declaration confirms, where applicable:

- terms/licensing/rights compatibility;
- privacy and data-processing terms;
- learner-content training/reuse policy;
- region/residency consequences;
- security and secret handling;
- availability/rate limits/quotas;
- cost assumptions and kill switches;
- deletion/export behavior;
- backup/restore implications;
- fallback/degraded behavior;
- provider-exit path;
- quality/calibration gate for evaluator providers.

# Replacement invariant

Replacing an external provider must not require redefining Skill, Knowledge, Band, Assessment, Progression, feature IDs, or learner identity. If provider replacement changes those semantics, the integration boundary was designed incorrectly.