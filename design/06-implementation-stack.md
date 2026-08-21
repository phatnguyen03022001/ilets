STATUS: CANONICAL
OWNS: initial deployable-unit allocation, primary-language/framework assignment, runtime responsibility split, cross-language contract strategy, and native verification baseline
DEPENDS_ON: ../CONSTITUTION.md, 04-application-flows.md, 05-api.md
DOES_NOT_OWN: learning/product truth, exact cloud/provider choice, database schema, deployment topology, evaluator model vendor, or future framework upgrades within the same responsibility boundary

# Implementation Stack

## Purpose

Assign the approved Go, Python, and TypeScript languages to explicit runtime responsibilities so implementation does not re-decide the product architecture or duplicate domain logic across three stacks.

## Initial deployable topology

```text
apps/
└── web/                  TypeScript

services/
├── core-api/             Go
└── evaluator/            Python

contracts/
├── http/                 OpenAPI when materialized
└── events/               JSON Schema or another explicit event contract only when a real async event boundary exists

tools/
└── contract / content / repository tooling as justified
```

Do not create additional deployables merely because a feature has its own name.

# TypeScript — learner web application

## Unit

```text
apps/web/
```

## Baseline

```text
Node.js       24 LTS
TypeScript    5.x
React         19.2-compatible line
Next.js       16.x App Router; implementation baseline 16.3
```

The exact supported patch version follows current security/support releases. Do not freeze an obsolete patch merely because this document names the implementation baseline.

Current reference points at design time:

- Next.js 16.3 release: `https://nextjs.org/blog`
- App Router: `https://nextjs.org/docs/app`
- Node.js releases: `https://nodejs.org/en/blog/release`

## Why Next.js

The web product requires both highly interactive learner tools and ordinary application navigation/content surfaces. Next.js App Router provides a current React framework with server/client component boundaries, route/layout conventions, and production application tooling without creating a second backend-domain architecture.

## TypeScript owns

- learner/admin web rendering;
- route/layout composition;
- interactive Reading/Writing workspaces;
- microphone/browser recording capture;
- YouTube IFrame integration;
- timers, local draft interaction, optimistic UI;
- SSE client/reconnect behavior;
- presentation-only transformations;
- accessibility and responsive UI.

## TypeScript does not own

- learner mastery/progression policy;
- deterministic IELTS scoring policy;
- evidence sufficiency;
- canonical gap/action logic;
- productive evaluator algorithms;
- copied domain DTO truth independent of the contract.

## Client-state rule

Use local/component state for transient interaction and server/resource state for durable product truth.

Do not create a global client store that becomes a second learner-state authority.

# Go — Core API and learning orchestration

## Unit

```text
services/core-api/
```

## Baseline

```text
Go            1.26.x supported line
HTTP          net/http
Router        github.com/go-chi/chi/v5, 5.3.x line
```

Current design-time references:

- Go supported releases: `https://go.dev/doc/devel/release`
- chi: `https://go-chi.io/`
- chi releases: `https://github.com/go-chi/chi/releases`

## Why Go + chi

The Core API is the stable boundary for learner/session/attempt orchestration and deterministic policy execution. Go provides a small runtime and strong standard HTTP primitives; chi stays close to `net/http`, adds routing/middleware composition, and avoids imposing a large application framework over domain ownership.

## Go owns

- public `/v1` API;
- authentication/authorization integration boundary, without owning provider identity semantics;
- LearnerProfile/Goal durable product state;
- DailyPlan and LearningSession orchestration;
- PracticeActivity creation;
- Attempt intake/lifecycle;
- deterministic Listening/Reading answer-key scoring;
- Assessment admissibility/claim execution from canonical rules;
- Progression execution: MasteryEstimate, GapEvaluation, ActionIntent, certification state;
- review queue composition;
- idempotency/concurrency protection;
- evaluation work orchestration and retry state;
- media-source eligibility/product state;
- SSE product event delivery.

## Go does not own

- AI rubric judgment itself;
- speech/audio feature extraction;
- LLM-generated feedback/content analysis;
- browser interaction;
- duplicated Python evaluator rules.

# Python — evaluator and media analysis

## Unit

```text
services/evaluator/
```

## Baseline

```text
Python        3.14.x supported line
FastAPI       0.141.x line at design time
Pydantic      current FastAPI-compatible supported line
project tool  uv
```

Current design-time references:

- Python 3.14 releases: `https://www.python.org/doc/versions/`
- FastAPI release notes: `https://fastapi.tiangolo.com/release-notes/`
- FastAPI features/OpenAPI: `https://fastapi.tiangolo.com/features/`

## Why FastAPI

The evaluator boundary is typed HTTP with structured request/response models, streaming/async-friendly behavior, and strong alignment with OpenAPI/JSON Schema. FastAPI fits that boundary without making Python the product API owner.

## Python owns

- Writing criterion observation generation;
- Speaking criterion observation generation;
- speech transcription when the audio source is legally/product-eligible;
- pronunciation/fluency/acoustic feature extraction where validated;
- text analysis supporting ErrorPattern/FeedbackArtifact generation;
- bounded AI-generated feedback candidates;
- authorized transcript/media segmentation and analysis;
- content/prompt candidate generation when requested;
- evaluator/model provenance and uncertainty output.

## Python does not own

- certification;
- Band advancement;
- final evidence sufficiency;
- DailyPlan selection;
- public user API;
- auth/session state;
- arbitrary YouTube media extraction.

# One public API rule

```text
Browser
  ↓
Go Core API
  ↓ internal contract
Python Evaluator
```

The web app does not independently orchestrate Python services. The Python service is not exposed as a second learner-facing API.

# Contract strategy

## HTTP

When implementation starts, materialize:

```text
contracts/http/openapi.yaml
```

It owns exact HTTP wire shape for the public Go API and may include a separate internal evaluator surface or an explicitly separated internal document if the public/internal lifecycle makes that cleaner.

Generated bindings/validators are derived.

## Events

Do **not** create an event-schema directory full of hypothetical events.

Create `contracts/events/` only when an actual asynchronous cross-unit event is introduced and cannot be represented sufficiently by the HTTP/evaluation-work contract.

## Stable IDs

Canonical Skill, Knowledge, Practice, Assessment, feature, practice-mode, and learner-state identifiers cross boundaries unchanged.

# Async work baseline

Do not pre-authorize Kafka, Redis Streams, a workflow engine, or a custom distributed scheduler.

For the first implementation:

1. Core API persists the authoritative attempt/evaluation-work lifecycle;
2. evaluator work is idempotent by evaluation identity;
3. Core API invokes the Python service through the internal contract;
4. retries preserve Attempt identity and do not double-charge/double-count results;
5. introduce a dedicated broker only after measured throughput/reliability requirements justify it.

This design rule keeps durable product state in one owner while preserving a clear Python execution boundary.

# Persistence boundary

This document intentionally does not select the final database/provider.

Implementation may choose relational/object storage appropriate to the runtime, but:

- one durable product fact has one runtime owner;
- Python must not silently mutate Go-owned learner/progression tables as a shared-database shortcut;
- large audio/media artifacts use explicit object references rather than being smuggled through JSON fields;
- storage schema is derived implementation detail, not domain authority.

# Native verification

## TypeScript

Minimum unit checks:

```text
format/lint
TypeScript typecheck
unit/component tests
Next.js build
Playwright critical learner-flow tests where relevant
```

Recommended initial tools:

- ESLint or Biome selected once and used consistently;
- `tsc --noEmit`;
- Vitest for unit/component logic where appropriate;
- Playwright for end-to-end browser flows.

Do not keep two competing formatter/linter stacks without a measured reason.

## Go

```text
gofmt check
go vet ./...
go test ./...
go test -race ./... where supported
build core-api
```

Additional static analysis may be added, but standard checks remain understandable without a bespoke meta-framework.

## Python

Recommended initial baseline:

```text
ruff format --check
ruff check
pyright or equivalent strict typecheck
pytest
```

Choose one primary static type checker and use it consistently.

## Contract checks

```text
OpenAPI validation
generated-client drift check
Go/TS/Python contract conformance
integration tests for public API → evaluator boundary
```

# Root verification

The repository must eventually expose one root entrypoint conceptually equivalent to:

```text
verify
  ├── web
  ├── core-api
  ├── evaluator
  ├── contracts
  └── cross-unit integration
```

A change that crosses Go/Python/TypeScript is not PASS because one language's tests are green.

# Framework-change rule

A patch/minor framework upgrade does not require a new architecture decision when responsibilities/contracts remain stable and verification is green.

A framework replacement requires an explicit design change when it materially changes:

- deployable boundaries;
- public/internal API ownership;
- rendering/runtime model;
- persistence ownership;
- cross-language contracts;
- operational complexity.

# Initial non-goals

Do not introduce at bootstrap unless a real requirement appears:

- microservice-per-feature architecture;
- Redis/Kafka merely because async work exists;
- vector database as default memory/content infrastructure;
- GraphQL alongside REST without a concrete consumer need;
- separate BFF that duplicates Core API semantics;
- Python and Go implementations of the same progression/scoring rule;
- frontend-owned Band/mastery calculations.