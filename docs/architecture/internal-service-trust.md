# INTERNAL-SERVICE-TRUST Internal service trust boundary

## Core → Evaluator authority boundary

Go Core API owns authoritative product work identity/state and is the permitted product caller. Python Evaluator owns bounded evaluation/media/content-analysis capabilities behind the existing internal HTTP machine contract. Evaluator is an internal capability, never a public API or independent product-state authority, and it does not read/write authoritative product persistence directly.

For the selected Google Cloud Run topology, authenticated internal invocation is specifically:

```text
Go Core Cloud Run service identity
        ↓ Google-signed OIDC identity token
restricted/private Evaluator Cloud Run service
        ↓ Cloud Run IAM authorization
bounded evaluator capability
```

This selected mechanism is part of the locked trust boundary rather than a replaceable implementation detail inside the current design.

Rules:

1. Browser, learner client, Next.js presentation code and Admin UI do not call Evaluator directly; they use Core's public product boundary.
2. Network placement is **not** caller identity. `localhost`, private IP, same host/VPC, internal DNS/hostname or co-location cannot establish authorization by itself.
3. Distinct least-privilege service identities are used where material. The caller principal must be authorized for the bounded internal capability rather than trusted because it is reachable.
4. Protected Google Cloud Tasks delivery likewise uses the appropriate service identity and Google-signed OIDC/IAM route when calling a protected handler.
5. Do **not** substitute a shared `INTERNAL_SECRET`, home-made service JWT, application-defined HMAC authentication protocol, mTLS, or service mesh for this locked Core→Evaluator authentication boundary.
6. Exact Cloud Run ingress/topology configuration remains deployment-owned, but it may not weaken the IAM/principal invariant.
7. Evaluator responses are bounded non-authoritative signals with provider/model/evaluator provenance and uncertainty/quality where material. Core validates the exact machine contract and interprets the result through owning product/Assessment/content policy before any learner/evidence/content consequence.
8. Co-locating logical units to reduce cost does not collapse this caller/callee, trust, persistence or exact machine-contract boundary.

## Material boundary preservation

The implementation authority chain remains:

```text
canonical product/system semantics
        ↓
materialized canonical registry / scoped machine contract
        ↓
derived generated bindings / implementation
```

Generated OpenAPI types, UI/DB types, storage schemas, prompts, provider responses, caches, metrics and migrations never redefine higher-level product semantics. Exact machine contracts remain authoritative only for their scoped wire/schema details and do not supersede product/Assessment/content authority.

For any material implementation boundary, implementation preserves where applicable: semantic owner; caller/callee direction; trust/access scope; allowed and forbidden data path; exact contract authority; commit/consistency class; retry/idempotency/recovery; cache/freshness; privacy/security; observability/audit; version compatibility; and verification evidence. A boundary may delegate detail to an existing owner but does not become another semantic owner merely because implementation enforces it.
