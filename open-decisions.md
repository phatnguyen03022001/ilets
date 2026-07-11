# Open Decisions & Gap Analysis

## Overview
This document documents gaps, uncertainties, and decisions that need to be made before implementation.

## Gap Categories

### Type 1: Missing Services (Critical)
### Type 2: Redundant/Over-engineered Services
### Type 3: Contradictory Requirements
### Type 4: Implementation Details (Deferred)

---

## Critical Gaps (P1)

### GPD-001: Missing File Service
**Severity**: P1 - Blocks core functionality
**Context**: Media Service references "File Service" but no service exists.

**Analysis**:
- Media Service defines upload/download/delete operations
- No dedicated File Service defined in SSOT
- Other services may need file operations (user avatars, evidence uploads)

**Options**:
1. ✅ **Create dedicated File Service** (Recommended)
   - Pros: Clear separation of concerns, reusable, well-defined
   - Cons: Additional service to build
   - Trade-off: Clean architecture vs. complexity

2. ❌ Integrate into Media Service
   - Pros: Simpler, fewer services
   - Cons: Media Service becomes bloated, unclear boundaries
   - Trade-off: Simplicity vs. maintainability

3. ❌ Create library utility instead
   - Pros: No service overhead
   - Cons: Not reusable across services, harder to test

**Recommended Action**: Create File Service
- Storage: S3-compatible (Neon S3)
- CDN: Cloudflare
- Features: Upload, download, delete, optimize, CDN purge
- Integration: Called by Media Service for storage operations

**Observable Signal**: File upload/delete operations work reliably for all file types

---

### GPD-002: Missing Email Service
**Severity**: P1 - Blocks email notifications
**Context**: Notification Service references "Email Service" but no dedicated service exists.

**Analysis**:
- Notification Service lists email as a communication channel
- No Email Service defined
- Needs to integrate with email delivery service

**Options**:
1. ✅ **Create dedicated Email Service** (Recommended)
   - Pros: Clear responsibility, easy to swap providers
   - Cons: Additional service
   - Trade-off: Clean architecture vs. complexity

2. ❌ Integrate into Notification Service
   - Pros: Simpler
   - Cons: Notification Service becomes monolithic
   - Trade-off: Simplicity vs. maintainability

3. ❌ Use notification provider directly
   - Pros: No service overhead
   - Cons: Harder to test, less flexibility
   - Trade-off: Simplicity vs. testability

**Recommended Action**: Create Email Service
- Provider: Resend (cost-effective, good API)
- Features: Send, queue, retry, deliverability tracking
- Integration: Called by Notification Service for email delivery

**Observable Signal**: Email delivery success rate > 95%

---

### GPD-003: Missing Push Service
**Severity**: P1 - Blocks push notifications
**Context**: Notification Service lists push notifications but no dedicated service exists.

**Analysis**:
- Notification Service mentions push notifications
- No Push Service defined
- Needs integration with FCM (Firebase Cloud Messaging)

**Options**:
1. ✅ **Create dedicated Push Service** (Recommended)
   - Pros: Clear responsibility, FCM-specific logic isolated
   - Cons: Additional service
   - Trade-off: Clean architecture vs. complexity

2. ❌ Integrate into Notification Service
   - Pros: Simpler
   - Cons: Notification Service becomes monolithic
   - Trade-off: Simplicity vs. maintainability

**Recommended Action**: Create Push Service
- Provider: Firebase Cloud Messaging (FCM)
- Features: Send, token management, delivery tracking
- Integration: Called by Notification Service for push notifications

**Observable Signal**: Push notification delivery rate > 90%

---

### GPD-004: Missing Search Index Service
**Severity**: P1 - Blocks search functionality
**Context**: Search Service references "index content" but no index management service exists.

**Analysis**:
- Search Service has endpoints to index content
- No index management service defined
- Need to handle index updates, rebuilds, health checks

**Options**:
1. ✅ **Create Search Index Service** (Recommended)
   - Pros: Isolates search logic, easier to manage
   - Cons: Additional service
   - Trade-off: Clean architecture vs. complexity

2. ❌ Integrate into Search Service
   - Pros: Simpler
   - Cons: Search Service becomes monolithic
   - Trade-off: Simplicity vs. maintainability

3. ❌ Use library function
   - Pros: No service overhead
   - Cons: Not reusable, harder to test
   - Trade-off: Simplicity vs. testability

**Recommended Action**: Create Search Index Service
- Provider: Elasticsearch or Meilisearch
- Features: Index, update, delete, rebuild, health check
- Integration: Called by Search Service when content changes

**Observable Signal**: Search index health score > 90

---

### GPD-005: Database Per Service vs. Shared Database
**Severity**: P1 - Architectural decision
**Context**: SSOT assumes database per service, but this has implications.

**Analysis**:
- **Per-Service Databases**:
  - Pros: Isolation, clear boundaries, independent scaling
  - Cons: Join queries difficult, data consistency, complex migrations

- **Shared Database**:
  - Pros: Easy joins, consistent schema, simpler migrations
  - Cons: Service coupling, hard to scale independently

**Options**:
1. ✅ **Keep per-service databases** (SSOT) (Recommended)
   - Pros: Clean architecture, clear boundaries
   - Cons: Complex cross-service queries
   - Trade-off: Architecture vs. complexity

2. ❌ Use shared database with schema per service
   - Pros: Easier joins, consistent schema
   - Cons: Service coupling, migration pain
   - Trade-off: Simplicity vs. coupling

3. ❌ Hybrid (some services share, some separate)
   - Pros: Flexible
   - Cons: Inconsistent pattern, harder to reason about

**Recommended Action**: Keep per-service databases
- Use database federation for cross-service queries
- Accept complexity for separation of concerns
- Document query patterns (avoid N+1, use subqueries)

**Observable Signal**: Cross-service queries work efficiently (< 500ms)

---

### GPD-006: AI Model Selection & Cost Optimization
**Severity**: P1 - Critical for evaluation accuracy
**Context**: AI Service needs to select evaluation models.

**Analysis**:
- **Current**: OpenRouter (multi-model)
- **Options**:
  - GPT-4o: Best quality, most expensive
  - Claude 3.5 Sonnet: Good quality, cheaper
  - GPT-4 Turbo: Similar to 4o, cheaper
  - GPT-3.5 Turbo: Fast, cheapest, lower quality

**Options**:
1. ✅ **Use GPT-4o for all evaluations** (Recommended)
   - Pros: Highest accuracy, best for IELTS (complex rubric alignment)
   - Cons: More expensive
   - Trade-off: Quality vs. cost

2. ✅ **Tiered model approach** (Recommended)
   - Writing: GPT-4o (highest quality needed)
   - Speaking: Claude 3.5 Sonnet (good for speech)
   - Reading/Listening: GPT-3.5 Turbo (sufficient)
   - Pros: Cost optimization, right tool for each task
   - Cons: More complex routing logic
   - Trade-off: Cost vs. complexity

3. ❌ Use only GPT-3.5 Turbo
   - Pros: Cheapest
   - Cons: Lower quality, unreliable for complex rubrics
   - Trade-off: Cost vs. quality

**Recommended Action**: Tiered model approach
- Document per-task model choice
- Monitor cost per evaluation
- A/B test model combinations

**Observable Signal**: Evaluation accuracy > 85% on official samples

---

## Secondary Gaps (P2)

### GPD-007: Message Queue Implementation
**Severity**: P2 - Not blocking, but good for architecture
**Context**: Search Service research suggests message queue for background tasks.

**Analysis**:
- Current architecture: REST-only
- Some tasks are better async (background processing)
- Need to decide if queue is necessary

**Options**:
1. ✅ **Add message queue (e.g., Bull, Redis Queue)** (Recommended)
   - Pros: Decouples tasks, better error handling, can scale workers
   - Cons: Additional infrastructure
   - Trade-off: Complexity vs. reliability

2. ❌ Keep REST-only
   - Pros: Simpler, no queue infrastructure
   - Cons: Blocking operations, harder error handling
   - Trade-off: Simplicity vs. reliability

3. ❌ Only use for critical tasks
   - Pros: Hybrid approach
   - Cons: Inconsistent pattern
   - Trade-off: Flexibility vs. consistency

**Recommended Action**: Add message queue for:
- Email sending (non-critical)
- Analytics aggregation (can be batch)
- Report generation (can be async)
- Media processing (can be async)

**Observable Signal**: Background tasks complete successfully 99% of time

---

### GPD-008: CDN Provider Selection
**Severity**: P2 - Performance optimization
**Context**: Media Service needs CDN for content delivery.

**Analysis**:
- **Cloudflare CDN**:
  - Pros: Fast for Vietnam, easy to set up, generous free tier
  - Cons: Less control over edge locations

- **AWS CloudFront**:
  - Pros: Global coverage, integrates with S3
  - Cons: More complex, less optimal for Vietnam

**Options**:
1. ✅ **Cloudflare CDN** (Recommended)
   - Pros: Fastest for target audience, easy setup
   - Cons: Fewer edge locations
   - Trade-off: Performance vs. control

2. ❌ AWS CloudFront
   - Pros: More control, S3 integration
   - Cons: Slower for Vietnam users
   - Trade-off: Control vs. performance

**Recommended Action**: Cloudflare CDN
- CDN for Media Service assets
- Edge caching for static content
- Automatic origin caching

**Observable Signal**: CDN cache hit rate > 80%

---

### GPD-009: Error Tracking Service
**Severity**: P2 - Observability
**Context**: Need to track errors across all services.

**Analysis**:
- Error tracking essential for debugging
- Need centralized error logging

**Options**:
1. ✅ **Sentry** (Recommended)
   - Pros: Easy setup, good TypeScript support
   - Cons: Cost (free tier limited)
   - Trade-off: Features vs. cost

2. ❌ Self-hosted error logging
   - Pros: Free
   - Cons: More maintenance, less features
   - Trade-off: Cost vs. maintenance

**Recommended Action**: Sentry
- Track all errors across services
- Alert on critical errors
- Provide stack traces and context

**Observable Signal**: Error capture rate > 95%

---

### GPD-010: Admin Dashboard Implementation
**Severity**: P2 - Operations
**Context**: Admin Service has extensive dashboard requirements.

**Analysis**:
- Need to build comprehensive admin interface
- Can use existing dashboard components

**Options**:
1. ✅ **Use dashboard components + custom pages** (Recommended)
   - Pros: Faster development, better UX
   - Cons: Vendor lock-in (if using library)
   - Trade-off: Speed vs. lock-in

2. ❌ Build from scratch
   - Pros: Full control
   - Cons: Time-consuming, reinventing wheel
   - Trade-off: Control vs. speed

**Recommended Action**: Use dashboard library (e.g., shadcn/ui components)
- Admin dashboard using shadcn/ui
- Real-time data (WebSocket or polling)
- Pre-built components (tables, charts, forms)

**Observable Signal**: Admin dashboard response time < 2s

---

## Tertiary Gaps (P3)

### GPD-011: Analytics Database Choice
**Severity**: P3 - Optimization
**Context**: Analytics Service needs database for time-series data.

**Analysis**:
- TimescaleDB vs. Postgres with custom optimization
- Depends on data volume

**Options**:
1. ✅ **TimescaleDB extension** (Recommended if >1M events/day)
   - Pros: Optimized for time-series, built-in compression
   - Cons: Requires separate database
   - Trade-off: Performance vs. complexity

2. ❌ Custom PostgreSQL optimization
   - Pros: No extra database
   - Cons: Manual optimization, slower
   - Trade-off: Simplicity vs. performance

**Recommended Action**: Start with PostgreSQL, migrate to TimescaleDB if needed
- Start with PostgreSQL
- Add indexes, materialized views
- If >1M events/day, migrate to TimescaleDB

**Observable Signal**: Analytics query performance acceptable

---

### GPD-012: Payment Gateway Selection
**Severity**: P3 - Business-critical but deferred
**Context**: Payment Service needs to select payment gateway.

**Analysis**:
- Stripe (primary), PayPal (fallback), MoMo (local Vietnam)

**Options**:
1. ✅ **Stripe (primary)** (Recommended)
   - Pros: Widely supported, excellent API, great docs
   - Cons: May have high fees for Vietnam

2. ❌ Stripe + MoMo
   - Pros: Good coverage, local payment option
   - Cons: More integration work
   - Trade-off: Coverage vs. complexity

**Recommended Action**: Stripe for now, add MoMo if Vietnam-only users
- Stripe for global
- MoMo if needed for local Vietnam users
- Paystack if needed for Africa/Asia

**Observable Signal**: Payment success rate > 95%

---

### GPD-013: Real-time Capabilities
**Severity**: P3 - Nice to have
**Context**: Admin dashboard wants real-time updates.

**Analysis**:
- WebSocket vs. polling
- Depends on urgency

**Options**:
1. ✅ **Polling** (simpler, sufficient) (Recommended)
   - Pros: Simple, no infrastructure
   - Cons: Not truly real-time
   - Trade-off: Simplicity vs. real-time

2. ❌ WebSockets
   - Pros: True real-time
   - Cons: More complex, higher latency
   - Trade-off: Real-time vs. complexity

**Recommended Action**: Polling for now (5-second intervals)
- Use polling for admin dashboard
- If real-time needed, add WebSocket later

**Observable Signal**: Admin dashboard updates acceptable with polling

---

## Redundant/Over-engineered Services

### RPD-001: Potentially Redundant Services
**Severity**: P2 - Consideration needed

**Analysis**:
Review if any services overlap:

1. **Learning Service vs. Vocabulary/Progress Services**
   - Learning Service: Course/lesson management
   - Vocabulary Service: Word learning
   - Progress Service: Progress tracking
   - ✅ Distinct purposes, no redundancy

2. **Analytics Service vs. Progress Service**
   - Analytics Service: Business and user analytics
   - Progress Service: User progress tracking
   - ✅ Distinct purposes, no redundancy

3. **AI Service vs. All Skill Services**
   - AI Service: Evaluation and feedback
   - Skill Services: Skill-specific operations
   - ✅ AI Service is utility, Skill Services are domain-specific

**Conclusion**: All services serve distinct purposes. No redundancy.

---

## Contradictory Requirements

### CRPD-001: Service Dependencies Not Specified
**Severity**: P1 - Missing dependency graph
**Context**: Services reference each other but dependency graph not documented.

**Analysis**:
- Multiple services reference other services
- No clear dependency order defined
- Need to establish implementation sequence

**Recommended Action**: Document dependency graph
- Create diagram showing service dependencies
- Establish implementation order based on dependencies
- Services can be implemented in parallel if no dependencies

**Observable Signal**: Implementation sequence documented and followed

---

## Implementation Details (Deferred)

### ID-001: File Storage Provider
**Status**: ✅ RESOLVED - ADR-0041
**Context**: Media Service needs storage provider.

**Options**:
- Neon S3 (built-in)
- AWS S3
- MinIO (self-hosted)
- Cloudflare R2

**Decision**: Skip in MVP, implement in Phase 2 with Media Service (ADR-0041)

---

### ID-002: Email Provider
**Status**: ✅ RESOLVED - ADR-0037
**Context**: Email Service needs provider.

**Options**:
- Resend
- SendGrid
- Postmark
- Amazon SES

**Decision**: Resend (ADR-0037)

---

### ID-003: Search Engine
**Status**: ✅ RESOLVED - GPD-004
**Context**: Search Service needs search engine.

**Options**:
- Elasticsearch
- Meilisearch
- Algolia
- Typesense

**Decision**: TBD during implementation (deferred, not critical for MVP)

---

### ID-004: AI Evaluation Accuracy Baseline
**Status**: ✅ RESOLVED - ADR-0040
**Context**: Need to establish baseline accuracy before launch.

**Action**:
- Collect 100 official IELTS samples
- Get human scores
- Test AI evaluation accuracy
- Set target accuracy (e.g., 85%)
- Iteratively improve

**Decision**: 90% accuracy on official samples (ADR-0040)

---

---

## API Overlaps (LOW Priority)

### AOP-001: Dashboard Data (Analytics vs Progress)
- **Overlap**: Both provide dashboard views
- **Scope**: Analytics (platform-wide, business metrics) vs Progress (user-specific, learning progress)
- **Resolution**: Accept difference in scope, no changes needed

### AOP-002: Band Score History (Analytics vs Progress)
- **Overlap**: Both track band score data
- **Scope**: Analytics (distribution across users) vs Progress (individual user history)
- **Resolution**: Accept difference in scope, no changes needed

### AOP-003: User Activity (Identity vs Admin)
- **Overlap**: Both provide activity data
- **Scope**: Identity (current user's sessions) vs Admin (user activity logs)
- **Resolution**: Accept difference in scope, no changes needed

### AOP-004: Content Management (Learning vs Reading)
- **Overlap**: Both handle content creation
- **Scope**: Learning (general courses, lessons) vs Reading (reading passages)
- **Resolution**: Accept difference in scope, no changes needed

### AOP-005: Progress Tracking (Progress vs Learning)
- **Overlap**: Both track progress data
- **Scope**: Progress (overall user progress) vs Learning (course/lesson progress)
- **Resolution**: Accept difference in scope, no changes needed

---

## Priority Order for Resolution

### Phase 1: Critical Gaps (P1)
1. ✅ **GPD-001: File Service** - RESOLVED via ADR-0041 (deferred to Phase 2)
2. ✅ **GPD-002: Email Service** - RESOLVED via ADR-0037 (Resend)
3. ✅ **GPD-003: Push Service** - DEFERRED to Phase 2 (not critical for MVP)
4. ✅ **GPD-004: Search Index Service** - DEFERRED to Phase 2 (not critical for MVP)
5. ✅ **GPD-005: Database strategy** - RESOLVED via ADR-0043 (embedded feedback)
6. ✅ **GPD-006: AI model selection** - RESOLVED via ADR-0035/ADR-0036 (Z.ai GLM-4.5-Flash)

### Phase 2: Secondary Gaps (P2)
1. ✅ **GPD-007: Message queue** - DEFERRED (optional, not critical)
2. ✅ **GPD-008: CDN provider** - DEFERRED (performance optimization)
3. ✅ **GPD-009: Error tracking** - DEFERRED (observability)
4. ✅ **GPD-010: Admin dashboard** - DEFERRED (operations)

### Phase 3: Tertiary Gaps (P3)
1. ✅ **GPD-011: Analytics database** - DEFERRED (optimization)
2. ✅ **GPD-012: Payment gateway** - DEFERRED (business)
3. ✅ **GPD-013: Real-time capabilities** - DEFERRED (nice to have)

---

## References
- IELTS Official Band Descriptors
- Next.js Documentation
- Drizzle ORM Documentation
- Neon Database Documentation
- OpenRouter Documentation
- Stripe Documentation
