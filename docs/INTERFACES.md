# Interfaces

> **MIGRATION DRAFT — AUTHORITY NONE**
>
> Created by `TASK-0006` revision 1. `CONSTITUTION.md`, `OBJECTIVE.md`, `spec/**`, and `design/**` remain canonical until an explicit later cutover.
>
> Canonical target snapshot: `phatnguyen03022001/ilets@46473042a5b8c6fc3093a7a938e18b6e281a9aa8`.
>
> This document migrates semantic interface boundaries only. It does not materialize OpenAPI, wire fields/status codes, generated bindings, provider identities, deployment topology, database schema, or another product authority.

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

No separate `IFC-*` job/command identity is justified by current canonical truth. Durable work registration, dispatch claim/fencing, retry, and reconciliation are operation semantics of `IFC-001` and `IFC-002`; the canonical slice does not require an independently owned queue, broker, worker API, scheduler, or provider boundary.

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

Scheduled/background work remains conditional on a real product/operational need. This migration therefore does not invent a queue/provider identity or a third interface solely to make the catalog look complete.

## INTERFACES-TRUST Data exchange and trust

Browser/Web input is untrusted. Web may render, compose presentation, and call the Core public API, but it cannot bypass Core to authoritative persistence or Evaluator. Internal reachability alone is not authorization; service-to-service access remains constrained by the canonical runtime/security design without making deployment/provider identities part of this IFC migration.

Material data crossing either interface preserves canonical semantic identity and applicability distinctions. Stable canonical/external IDs and exact content revision identity cross boundaries unchanged where applicable. Transport representations must not collapse materially distinct states such as:

- optional input not supplied;
- canonically `NOT_APPLICABLE`;
- `UNKNOWN` / unresolved;
- required but absent or otherwise invalid; and
- present value.

Public projections expose only data the caller may legally observe and must not reveal another learner's protected resource, privileged content/security state, evaluator/system instructions, hidden assessment material, or provider secrets. Evaluator output remains an untrusted bounded signal until Core validates and interprets it under the owning policy, with provenance/uncertainty retained where consequential.

External provider egress/ingress, callbacks, provider lifecycle, processor retention/reuse, provider failure/fallback, and exit treatment remain outside this bounded IFC inventory. `UNK-003` continues to own that unresolved external-dependency migration rather than creating `EXT-*` records here.

## INTERFACES-EVOLUTION Exact contract materialization and evolution

This migration establishes semantic boundaries only. It deliberately does not create `contracts/http/**`, OpenAPI, exact request/response fields, status-code mappings, SDKs, generated bindings, transport DTOs, event envelopes, or implementation adapters.

After documentation closure and an explicit contract lock/materialization step, each material runtime boundary has one exact machine-contract authority. Generated TypeScript/Go/Python bindings, when useful, are derived from that contract and validated rather than manually becoming parallel truth. Public and internal contracts may be separate, but equivalent DTO/schema truth is not independently authored by multiple runtimes.

Compatibility preserves semantic identity and the applicability/absence distinctions above across every supported version-skew direction. A change can therefore be semantically breaking even when a wire schema still validates. SSE remains part of the public HTTP contract; a separate event contract is justified only by a genuinely separate asynchronous cross-unit event boundary.
