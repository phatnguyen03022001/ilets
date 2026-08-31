# Interfaces

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Created by `TASK-0006` revision 1 and extended only for the bounded external-dependency slice by `TASK-0009` revision 1. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Canonical target snapshot: `phatnguyen03022001/ilets@4891798cc76fbb1317e51d2d231f378f9bf9dc30`.
>
> This document migrates semantic interface boundaries plus the smallest material external-dependency inventory and its failure/fallback/exit treatment for the current behavior slice. It does not materialize OpenAPI, wire fields/status codes, generated bindings, provider activation/configuration, deployment topology, database schema, `CAP-*` build/buy state, `DEC-*` closure, or another product authority.

## INTERFACES-CONTRACTS Material contract boundaries

The bounded migrated slice has exactly two material semantic interface identities:

1. the learner/admin-facing product API owned by Core and consumed through Web; and
2. the internal Core-to-Evaluator bounded capability API.

They are stable authority/trust boundaries, not endpoint, button, DTO, transport-helper, or provider inventories. Web remains presentation authority and cannot become a second domain API. Core remains the sole learner/admin product API authority and the only application runtime allowed to own authoritative product persistence. Evaluator remains an internal bounded capability and cannot become public product, learner-state, evidence, certification, content-activation, or next-action authority.

### IFC-001 Core public product API

`IFC-001` is an `API` owned by `SYS-002` Core API with `SYS-001` Web as its material catalog peer.

The contract is semantic rather than wire-exact. It covers the public product operation classes already fixed by canonical design:

- bounded synchronous reads that return authoritative or explicitly derived product state without hidden mutation;
- bounded synchronous mutations whose durable-success acknowledgement follows the authoritative Core commit;
- asynchronously accepted operations whose accepted/pending state and logical work identity are durably committed or recoverably derivable before acknowledgement depends on downstream work;
- privileged learner/admin mutations that still pass Core authentication/access/capability and legal-state boundaries; and
- SSE/update delivery as a scoped refresh hint, never as state authority or proof of semantic completion.

Creation/submission work is idempotent where retry could duplicate learner history, evidence, accepted content revision, paid/provider work, or logically identical capability work. A preflight duplicate or stale-revision check may reject early, but the decisive idempotency/revision decision is coupled to the authoritative mutation or an equivalent serialization invariant.

Transport success is kept distinct from domain success, domain unresolved/negative outcomes, operation-contract rejection, and infrastructure/transient failure. An infrastructure or internal-capability failure cannot be translated into learner weakness, fake evidence, fake score, content-quality truth, or product-support truth.

### IFC-002 Core to Evaluator bounded capability API

`IFC-002` is an internal `API` owned by `SYS-003` Evaluator with `SYS-002` Core API as its sole material catalog peer for this slice.

Core owns authoritative work identity/state, product persistence, interpretation, and final reconciliation. Evaluator receives only bounded capability input and returns bounded result/signals plus provenance, quality, uncertainty, and evaluator/model identity where material. Browser/Web/admin clients do not call Evaluator directly, and Evaluator does not read or mutate Core-owned authoritative product persistence.

Invocation is bounded by a caller deadline and safe cancellation semantics. Retry/replacement reuses one logical work identity and preserves enough execution-attempt/fencing association for Core to reject late or superseded completion deterministically. Duplicate delivery of an already accepted completion is idempotent. A timeout or network failure does not prove remote non-execution and therefore cannot authorize blind duplicate work.

A successful internal transport response only proves that a bounded capability response arrived. Core still validates contract/work association, current execution/deletion/content eligibility, and owning policy before the result can affect Observation, EvidenceFact, Assessment, Progression, content state, or a next action.

## INTERFACES-ASYNC Asynchronous jobs and commands

No separate `IFC-*` job/command identity is justified by current canonical truth. Durable work registration, dispatch claim/fencing, retry, and reconciliation are operation semantics of `IFC-001` and `IFC-002`; the selected external dispatch service is cataloged separately as dependency `EXT-003`, but its queue/delivery state is not an independently authoritative interface or business-work identity.

For material asynchronous work:

- one logical work identity represents one accepted semantic operation;
- required pending-work/outbox/recoverable state is committed with, or deterministically recoverable from, authoritative Core state before an accepted/pending acknowledgement depends on it;
- merely observing `pending` does not grant exclusive dispatch ownership;
- where exclusivity matters, the current eligible execution attempt is decisively claimed/leased/fenced before dispatch;
- decisive dispatch admission re-checks current cancellation, deletion/tombstone, content/rights/privacy, and other material eligibility that can prohibit execution or data egress;
- bounded transient retry/backoff preserves the logical work identity and treats ambiguous outcomes as unresolved until safe reconciliation is possible;
- stale/superseded completion cannot overwrite newer authority or independently duplicate Observation, EvidenceFact, learner-state, content-revision, or semantic continuation;
- accepted completion that requires downstream semantic continuation is not semantically complete until the required continuation is committed or durably/recoverably reconstructable; and
- backlog, pending, degraded, or honestly rejected state is preferred to unbounded process/provider queues or silent quality reduction.

Scheduled/background work remains conditional on a real product/operational need. This migration does not invent a third interface or queue contract solely to make the catalog look complete; `EXT-003` records only the already-canonical selected external dispatch dependency.

## INTERFACES-TRUST Data exchange and trust

Browser/Web input is untrusted. Web may render, compose presentation, and call the Core public API, but it cannot bypass Core to authoritative persistence or Evaluator. Internal reachability alone is not authorization; service-to-service access remains constrained by the canonical runtime/security design without making deployment/provider identities part of this IFC migration.

Material data crossing either interface preserves canonical semantic identity and applicability distinctions. Stable canonical/external IDs and exact content revision identity cross boundaries unchanged where applicable. Transport representations must not collapse materially distinct states such as:

- optional input not supplied;
- canonically `NOT_APPLICABLE`;
- `UNKNOWN` / unresolved;
- required but absent or otherwise invalid; and
- present value.

Public projections expose only data the caller may legally observe and must not reveal another learner's protected resource, privileged content/security state, evaluator/system instructions, hidden assessment material, or provider secrets. Evaluator output remains an untrusted bounded signal until Core validates and interprets it under the owning policy, with provenance/uncertainty retained where consequential.

External provider egress/ingress, callbacks, processor retention/reuse, failure/fallback, and exit treatment are migrated only at the semantic dependency level below. Canonical provider lifecycle and exact selected/TBD route truth remain owned by `design/07-third-party-services.md`; this migration does not promote any route to `ACTIVE`.

## INTERFACES-EXTERNAL-DEPENDENCIES External dependency inventory

The current migrated behavior slice justifies exactly four material `EXT-*` identities. The inventory is deliberately boundary-sized: it does not create one identity per model alias, SDK/library, endpoint, pricing plan, optional candidate, observability tool, CDN/edge service, email route, payment route, or every infrastructure product named by canonical third-party design.

All concrete routes named below remain at their canonical lifecycle state. `SELECTED_FOR_IMPLEMENTATION` means selected for implementation only; it is not production activation, legal/privacy approval, data-egress approval, calibration evidence, or permission to carry real learner traffic.

### EXT-001 Clerk identity and session service

`EXT-001` is the selected external identity/session service for authenticated public product access. Clerk may hold credential/session mechanics and issue the external principal presented to Core, but Core retains the stable internal actor/learner identity, principal association, authorization/capability policy, entitlements, and all learner/product state.

Durable learner state in this slice requires authenticated durable identity. Provider roles, organizations, permissions, or metadata cannot become Core authorization truth. Replacing the identity service therefore means re-establishing external-principal association to the same Core-owned actor semantics rather than migrating product authority into the provider.

### EXT-002 Neon Launch PostgreSQL hosting

`EXT-002` is the selected managed PostgreSQL hosting route for the Core-owned authoritative product store. Hosting does not own the data semantics in `docs/DATA.md`: only Core application runtime reads/writes authoritative product state, and network/provider behavior cannot redefine transaction, evidence, progression, target, or content meaning.

The dependency is material because the migrated target/evidence/plan loop relies on durable Core state. Its managed-hosting identity does not create a database-schema identity, provider-owned product authority, or a requirement for a second live database.

### EXT-003 Google Cloud Tasks bounded dispatch

`EXT-003` is the selected bounded asynchronous dispatch service for recoverable external delivery of already accepted Core work. It is not the authority for acceptance, logical work identity, cancellation/deletion eligibility, execution ownership, or semantic completion.

Core must durably register or recoverably derive required work before learner-visible acceptance depends on dispatch. A task message/delivery is an execution attempt, not business truth; dispatch claim/fencing, idempotency, current egress eligibility, completion reconciliation, and downstream semantic continuation stay anchored in Core-owned state.

### EXT-004 External evaluator AI and speech provider boundary

`EXT-004` is the material external provider boundary behind bounded Evaluator AI/speech capabilities used by the migrated evidence/support loop. The exact provider/model route matrix remains canonical in `design/07-third-party-services.md` and is intentionally not duplicated as one catalog identity per provider or model alias.

Selected productive-evaluation, speech/transcription, realtime and translation routes remain selected-for-implementation only. Any capability whose canonical provider state remains `TBD`—including the unresolved pronunciation/acoustic route—remains unavailable as a resolved provider claim here. A successful provider response is still only bounded external output; Evaluator/Core validation, provenance, quality/uncertainty and owning Assessment/product policy determine whether it may have any downstream consequence.

## INTERFACES-DEPENDENCY-FAILURE-EXIT Dependency failure, fallback, and exit

Every external call is bounded by a caller deadline and an explicit retry classification. Transient failures may receive bounded retry/backoff when safe; permanent failures do not become retry storms; timeout, connection loss, or missing acknowledgement is an ambiguous outcome rather than proof that remote work did not execute. Retry/replacement preserves one logical work identity and sufficient execution-attempt/fencing state to prevent duplicate accepted learner work, provider cost, Observation, EvidenceFact, or semantic continuation.

Provider ingress/output is untrusted until associated with current authoritative work and validated. Provider or infrastructure failure may leave work pending, degraded, temporarily unavailable, or unresolved; it never authorizes a fake score, learner weakness, fabricated evidence, lower target, content-quality rewrite, or silent completion.

Fallback is allowed only to a pre-approved route that independently satisfies the same consequence-specific quality/calibration, privacy, security, rights, reliability, and evidence floor. A fallback label does not prove interchangeability. When no equivalent route exists, the valid outcome is delayed/unavailable/unresolved work or a semantically valid lower-consequence product path—not lower standards disguised as recovery.

External egress uses only the minimum data necessary for the declared operation. Before a route may become active, applicable retention/reuse/training, deletion/export, callback authentication/replay, quota/cost and rights constraints must be satisfied under their canonical owners. A deletion/export request being sent is not deletion/export truth; ambiguous outcomes remain pending and are retried/reconciled safely, including provider copies/derived artifacts where the canonical provider contract requires them.

Exit remains semantic and bounded rather than a speculative multi-provider framework:

- `EXT-001`: Core-owned stable actor/learner identity and authorization state remain independent of the external subject so identity-provider replacement can re-associate principals without redefining product identity.
- `EXT-002`: portability uses the smallest PostgreSQL-native export/restore/recovery path that preserves canonical state and migration compatibility; no second live database or automatic multi-cloud failover is implied.
- `EXT-003`: durable Core logical-work/recovery state makes the dispatch service replaceable without treating provider queue state as accepted business state; ambiguous in-flight attempts are reconciled before safe redrive.
- `EXT-004`: provider-neutral Evaluator capability boundaries, exact consequential provenance and consequence-specific calibration keep provider/model replacement from redefining product/evidence semantics; replacement still requires independent validation before use.

These exit rules do not select alternates, activate providers, create `CAP-*` build/buy dispositions, create `DEC-*` choices, or close `UNK-004`.

## INTERFACES-EVOLUTION Exact contract materialization and evolution

This migration establishes semantic boundaries only. It deliberately does not create `contracts/http/**`, OpenAPI, exact request/response fields, status-code mappings, SDKs, generated bindings, transport DTOs, event envelopes, or implementation adapters.

After documentation closure and an explicit contract lock/materialization step, each material runtime boundary has one exact machine-contract authority. Generated TypeScript/Go/Python bindings, when useful, are derived from that contract and validated rather than manually becoming parallel truth. Public and internal contracts may be separate, but equivalent DTO/schema truth is not independently authored by multiple runtimes.

Compatibility preserves semantic identity and the applicability/absence distinctions above across every supported version-skew direction. A change can therefore be semantically breaking even when a wire schema still validates. SSE remains part of the public HTTP contract; a separate event contract is justified only by a genuinely separate asynchronous cross-unit event boundary.
