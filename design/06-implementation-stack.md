STATUS: CANONICAL
OWNS: initial deployable-unit allocation, primary-language/framework assignment, runtime responsibility split, cross-language contract strategy, canonical-registry materialization strategy, and repository/native verification contract
DEPENDS_ON: ../CONSTITUTION.md, 04-application-flows.md, 05-api.md
DOES_NOT_OWN: learning/product truth, parser/materializer implementation, registry serialization or codegen-library choice, CI platform configuration, exact dependency patch versions, cloud/provider choice, database schema, deployment topology, evaluator model vendor, or package-manager lock state

# Implementation Stack

## Purpose

Assign approved languages/framework families to explicit runtime responsibilities so implementation does not re-decide first-order architecture or duplicate domain logic across stacks.

Patch/minor dependency selection is implementation maintenance. This architecture owns **responsibility and framework family**, not a frozen release-note snapshot.

# Initial deployable topology

```text
apps/
└── web/                  TypeScript / React / Next.js App Router

services/
├── core-api/             Go / net/http + chi
└── evaluator/            Python / FastAPI

contracts/
├── http/                 exact HTTP contracts once materialized
└── events/               only for a real asynchronous cross-unit event boundary

tools/
└── contract/content/repository tooling when justified
```

Do not create another deployable merely because a product feature has a distinct name.

These are ownership/process boundaries, not a requirement for three separately billed infrastructure stacks. Initial deployment may co-locate runnable units on the same host/platform when isolation, scaling, security, failure, and runtime constraints remain satisfied. Split infrastructure only when a demonstrated requirement justifies the added operational cost.

# Reuse-first implementation invariant

Before adding a package, service, external call, generated artifact class, or infrastructure subsystem, implementation must first determine whether the requirement can be satisfied by an existing canonical semantic, runtime owner, contract, content asset, standard-library/framework capability, or already-approved provider route.

Preferred execution order when the resulting semantics/quality are equivalent:

```text
existing canonical/runtime/content capability
        ↓
deterministic or local first-party execution
        ↓
existing eligible external-provider capability
        ↓
new external provider or infrastructure only for a demonstrated gap
```

Rules:

1. deterministic scoring/validation remains deterministic; do not call AI to replace an available exact rule;
2. browser/native/framework capabilities should be reused for presentation/capture/local interaction before introducing a duplicate backend capability when ownership remains correct;
3. derived work may be cached/reused by stable input + policy/model/version identity when doing so preserves correctness, privacy, freshness, and auditability;
4. retries/repeated requests reuse logical work identity and prior valid results where applicable instead of duplicating provider work/cost;
5. expensive noninteractive work may be delayed, batched, or omitted when optional; cost pressure must not silently lower evidence/quality standards;
6. do not pre-generate infrastructure, content quantity, or AI output merely because a taxonomy exists; concrete demand comes from actual product/content coverage requirements;
7. optimization must not move semantic authority into caches, generated files, model prompts, or provider output.

This is an implementation-economy rule, not permission to bypass required content diversity, calibration, reliability, or support gates.

# Version policy

Architecture freezes compatibility families and responsibilities, not volatile patch numbers.

At implementation/bootstrap time:

- use a currently supported Node.js LTS line;
- use a currently supported TypeScript release compatible with the selected web stack;
- use the supported React/Next.js App Router line chosen for the web application;
- use a currently supported Go release with `net/http` and `chi/v5`;
- use a currently supported Python 3 release with a compatible FastAPI/Pydantic line;
- pin exact package versions in implementation manifests/lockfiles;
- keep supported security/maintenance releases current through normal dependency maintenance.

A patch/minor upgrade does not change architecture when responsibility, runtime model, and contracts remain stable.

# TypeScript — web

## Unit

```text
apps/web/
```

## Framework family

```text
TypeScript
React
Next.js App Router
```

## Owns

- learner/admin rendering;
- route/layout composition;
- interactive Reading/Writing workspaces;
- browser microphone/recording capture;
- external embedded-player interaction;
- timers/local draft interaction/optimistic presentation;
- SSE client/reconnect behavior;
- presentation-only transformations;
- accessibility/responsive UI.

## Does not own

- mastery/progression policy;
- deterministic IELTS scoring policy;
- evidence sufficiency;
- canonical gap/action rules;
- content validation/activation authority;
- productive evaluator algorithms;
- handwritten DTO truth independent of machine contracts.

## Client-state invariant

Transient interaction may live client-side. Durable learner/product truth remains server/resource state.

A global client store must not become a second Learner/Target/Progression/content authority.

# Go — Core API + deterministic orchestration

## Unit

```text
services/core-api/
```

## Framework family

```text
Go
net/http
chi/v5
```

## Owns

- learner/admin-facing `/v1` API;
- auth/authorization integration boundary;
- durable LearnerProfile/TargetProfile product state;
- durable content identity/revision metadata and release/assignment/operational state;
- DailyPlan/LearningSession orchestration;
- PracticeActivity creation;
- Attempt intake/lifecycle;
- deterministic Listening/Reading scoring;
- deterministic content/reference/answer/structural validation where exact rules apply;
- applicable content-validation policy aggregation over deterministic and bounded validator signals;
- content demand, eligible-pool reuse, learner assignment, quarantine/retirement/revalidation orchestration;
- content generation/validation work orchestration and retry state when those capabilities are implemented;
- Assessment policy execution over canonical/materialized rules;
- Progression execution;
- Planner eligibility/ranking orchestration;
- review-queue composition;
- idempotency/concurrency enforcement;
- evaluation-work orchestration/retry state;
- media-source eligibility/product state;
- product SSE event delivery.

## Does not own

- learning/content semantic truth defined by canonical owners;
- AI rubric judgment;
- speech/audio feature extraction;
- LLM analysis/generation internals;
- browser interaction;
- duplicate Python evaluator behavior.

# Python — evaluator/media analysis

## Unit

```text
services/evaluator/
```

## Framework family

```text
Python 3
FastAPI
Pydantic-compatible typed models
uv or equivalent project/environment tooling chosen consistently
```

## Owns

- Writing criterion observation generation;
- Speaking criterion observation generation;
- eligible speech transcription;
- validated pronunciation/fluency/acoustic feature extraction;
- bounded text analysis supporting Feedback/ErrorPattern candidates;
- bounded AI-generated feedback/content candidates;
- bounded content-validation analysis/signals where deterministic checks are insufficient and the capability is explicitly invoked;
- authorized transcript/media analysis;
- evaluator/model/generator/validator provenance and uncertainty output.

## Does not own

- content release activation, assignment eligibility, or authoritative revision mutation;
- certification;
- Band advancement;
- final evidence sufficiency;
- DailyPlan selection;
- public user/admin API;
- auth/session state;
- arbitrary external-media extraction.

# One public product API

```text
Browser
  ↓
Go Core API
  ↓ internal contract where needed
Python Evaluator
```

Python is a bounded internal capability, not a second learner-facing backend or content authority.

# Contract strategy

## HTTP

Before independent runtimes implement the same boundary, materialize one exact contract under `contracts/` as required by `05-api.md`.

Typical structure:

```text
contracts/http/openapi.yaml
```

Public and internal evaluator/content-capability surfaces may be split if useful, but each boundary has one exact machine authority.

Generated clients/server bindings/validators are derived artifacts.

## Events

Do not create hypothetical event schemas.

Create event contracts only when an actual asynchronous cross-unit boundary exists and HTTP/work-resource semantics are insufficient.

## Stable identities

Canonical Skill, Knowledge, Practice, Assessment, feature, practice-mode, learner-state identifiers, and exact content revision references cross boundaries unchanged unless a deliberate presentation translation is explicitly defined.

# Canonical registry materialization

Shared canonical identities consumed by implementation are materialized through a derived pipeline rather than copied independently into each runtime:

```text
canonical Markdown owners
        ↓
repository materializer / validator
        ↓
derived machine-readable registry
        ↓
generated or validated Go / TypeScript / Python consumers
```

Invariants:

1. the canonical Markdown owner remains semantic authority;
2. a machine-readable registry is derived implementation material, not another SSOT;
3. generated language bindings, constants, and models are derived artifacts;
4. equivalent canonical enum/ID registries must not be maintained manually and independently in Go, TypeScript, and Python;
5. duplicate canonical IDs fail repository verification;
6. broken canonical references fail repository verification;
7. generated-registry or generated-binding drift fails repository verification where those artifacts are materialized;
8. derived artifacts preserve sufficient source/provenance identity to trace values back to the canonical owner and source revision;
9. stable canonical IDs remain unchanged across language and machine boundaries;
10. materialization may be incremental: an initial vertical slice need materialize only the canonical registries it actually consumes;
11. Markdown parsing, generated registries, and generated code never acquire learning/product semantic authority merely because runtimes consume them;
12. a parser/materializer/tooling defect must fail verification or leave the affected materialization unresolved; tooling may not silently reinterpret or rewrite canonical meaning.

This architecture does not freeze the materializer language, parser implementation, serialization format, generated-registry file format, generator/codegen library, or a requirement to materialize the complete ontology on day one.

Materialization of canonical registries and materialization of cross-unit machine contracts are related but distinct: registries derive stable canonical identities, while `contracts/` owns exact interface shape for genuine machine boundaries.

# Async-work baseline

Do not pre-authorize Kafka, Redis Streams, a workflow engine, or another broker merely because evaluation, generation, or validation can be asynchronous.

Initial semantic baseline:

1. Core API persists authoritative Attempt/content-work/evaluation-work lifecycle and operational state it owns;
2. expensive/provider-backed work is idempotent by stable logical identity;
3. Core API invokes Evaluator/content capability through the internal contract only when needed;
4. retry preserves logical work/Attempt/content-revision identity and cannot double-count evidence or duplicate accepted provider cost/work;
5. dedicated dispatch infrastructure is introduced only after measured reliability/throughput need.

# Persistence boundary

Final database provider/schema is not owned here.

Implementation invariants:

- one durable product fact has one runtime owner;
- Evaluator does not mutate Core-API-owned learner/progression/content operational storage directly;
- historical Attempts pin the exact content revision they used;
- large audio/media/content artifacts use explicit object references rather than opaque large JSON state where appropriate;
- storage schema is derived implementation, not domain authority.

# Native verification baseline

Each deployable owns its native checks.

## Web

At minimum:

```text
format/lint
TypeScript typecheck
unit/component tests
production build
critical browser E2E tests
```

Use one coherent formatter/linter strategy rather than overlapping competing stacks.

## Go

At minimum:

```text
gofmt check
go vet ./...
go test ./...
race tests where supported/relevant
build core-api
```

## Python

At minimum:

```text
format check
lint
one primary strict static type check
pytest
```

Tool choice may evolve without changing architecture when the verification contract remains equivalent.

## Contracts

At minimum once contracts exist:

```text
schema validation
generated-artifact drift check
consumer/provider conformance
public API → evaluator/content-capability integration tests where implemented
backward-compatibility checks where deployed contracts require them
```

# Root verification

The repository uses one root verification contract for local and automated verification. Checks may be absent while the corresponding unit/materialized artifact does not yet exist; once a unit, registry, contract, or cross-unit boundary is materialized, its applicable checks enter this same root contract.

Conceptually:

```text
verify
  ├── repository/canonical
  │     ├── canonical metadata
  │     ├── semantic dependency/reference integrity
  │     ├── canonical ID uniqueness
  │     └── generated registry/binding drift where materialized
  ├── contracts where materialized
  ├── web where materialized
  ├── core-api where materialized
  ├── evaluator where materialized
  └── cross-unit integration where materialized
```

Repository/canonical verification proves structural/reference consistency and derived-artifact agreement; it does not turn tooling or generated files into canonical semantic owners.

CI invokes this same root verification contract. CI may provide an execution environment, triggers, or reporting, but it must not become a separate hidden definition of repository correctness. Local verification and CI must agree on what constitutes PASS for the same repository state.

A cross-stack change is not PASS because only one affected ecosystem is green.

# Framework-change rule

A framework replacement requires a design change when it materially changes any of:

- deployable boundaries;
- API ownership;
- rendering/runtime model;
- persistence ownership;
- cross-language contracts;
- operational complexity.

Patch/minor maintenance within the same responsibility boundary is implementation work.

# Initial non-goals

Do not introduce at bootstrap without demonstrated need:

- microservice-per-feature topology;
- broker/queue infrastructure merely because async work exists;
- vector database as default memory/content store;
- GraphQL alongside REST without a concrete consumer need;
- separate BFF duplicating Core API semantics;
- Python and Go implementations of the same progression/scoring/content-validation rule;
- frontend-owned Band/mastery/content eligibility calculations.