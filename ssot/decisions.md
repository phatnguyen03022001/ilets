# Architecture Decision Records

## ADR Index

### ADR-0045: Writing Service Retry Logic
**Date**: 2026-07-12
**Status**: Resolved
**Related**: ADR-0035, ADR-0040

### ADR-0044: Writing Service Task Type Support
**Date**: 2026-07-12
**Status**: Resolved
**Related**: ADR-0034, ADR-0045

### ADR-0043: Writing Service Feedback Storage Strategy
**Date**: 2026-07-12
**Status**: Resolved
**Related**: ADR-0035

### ADR-0042: Writing Service API Ownership
**Date**: 2026-07-12
**Status**: Resolved
**Related**: ADR-0034

### ADR-0041: Writing Service Media Storage Deferred
**Date**: 2026-07-12
**Status**: Resolved
**Related**: ADR-0037, ADR-0038

### ADR-0040: AI Evaluation Accuracy Target
**Date**: 2026-07-11
**Status**: Resolved
**Related**: ADR-0035, ADR-0036

### ADR-0039: Implementation Sequence (Writing Wedge E)
**Date**: 2026-07-11
**Status**: Resolved
**Related**: ADR-0034, ADR-0035

### ADR-0038: Sandbox Strategy
**Date**: 2026-07-11
**Status**: Resolved
**Related**: ADR-0037

### ADR-0037: Authentication Strategy
**Date**: 2026-07-11
**Status**: Resolved
**Related**: ADR-0038

### ADR-0036: Default AI Model Management
**Date**: 2026-07-11
**Status**: Resolved
**Related**: ADR-0035

### ADR-0035: First Live Z.ai Round-Trip with GLM-4.5-Flash
**Date**: 2026-07-11
**Status**: Resolved
**Related**: ADR-0036

### ADR-0034: Band Score Calculation Utilities
**Date**: 2026-07-11
**Status**: Resolved

---

## ADR-0038: Sandbox Strategy

### Context
Founder preference: strict sandbox rules for AI interactions to prevent LLM from accessing system files or making destructive changes. Need to ensure security-first approach.

### Decision
- **Rule set**: Law 4 (basic restrictions)
- **CLI access**: Completely restricted (no `exec`, `spawn`, `child_process`)
- **File system access**: Read-only (no `fs.write`, `fs.writeFileSync`)
- **Environment variables**: Read-only (no `process.env` modification)
- **Network access**: Only specified APIs (Z.ai, Anthropic, Resend, Twilio, Stripe)
- **Constrain**: LLM cannot access any code or configuration files

**Key choices**:
- Law 4 as baseline (industry-standard restrictions)
- Strict CLI restriction (prevent command execution)
- Read-only filesystem (prevent file modifications)
- Whitelist-based network access (only approved domains)
- No self-modification (LLM cannot change system files)

**Why this choice?**
- Law 4 is well-established and tested security pattern
- Prevents most common LLM security vulnerabilities (command injection, file corruption)
- Allows safe evaluation and content generation
- Founder preference for strict control

### Consequences

**Positive (benefits)**:
- High security posture (prevents most LLM attacks)
- Safe evaluation environment (no risk to production data)
- Clear boundaries for AI behavior
- Founder-approved security stance

**Negative (drawbacks)**:
- More complex prompt engineering (constrain LLM to allowed actions)
- May limit AI capabilities (cannot run tools, save files)
- Requires strict monitoring to ensure LLM follows rules

**Neutral (unaffected)**:
- User-facing features unchanged
- Security logging: all AI interactions tracked
- Rate limiting: applies to all AI endpoints

### Observable Signal (migration gate)

**Signal**: LLM attempting prohibited operations (detected via monitoring logs)

**Trigger**:
- If signal → A: Strengthen constraints, update prompt instructions
- If signal → B: Add sandbox logs/alerts for human review
- If signal → C: Re-evaluate sandbox strategy, consider alternative approach

**Rationale**: Monitoring for prohibited operations directly indicates sandbox effectiveness.

### Rollback Plan

1. Step 1: Relaxed CLI access (allow limited `exec`)
2. Step 2: Allow filesystem read-write for specific directories
3. Step 3: Remove strict network whitelist

**Cost of rollback**: medium (requires re-implementation of security measures)

### References

- Related decisions: [[ADR-0037]]
- Source material: Founder preference (from conversation history)

---

## ADR-0037: Authentication Strategy

### Context
Need secure, scalable authentication for 16 microservices with founder preference to avoid third-party auth providers for control and cost.

### Decision
- **Primary**: Magic-link authentication via Resend
- **Session management**: iron-session (serverless-safe)
- **Tokens**: JWT (access + refresh)
- **MFA**: TOTP (Google Authenticator) via iron-session
- **Authorization**: Role-based (student, teacher, admin)
- **No third-party**: Custom auth implementation, no Auth0, Clerk, NextAuth v4

**Key choices**:
- Magic-link (email-based) for passwordless auth (frictionless UX)
- iron-session for serverless-safe session management (supports Node.js/Edge)
- JWT access tokens (15 min) + refresh tokens (7 days)
- No OAuth providers (maintain full control)

**Why this choice?**
- Magic-link reduces friction (no password management)
- iron-session is serverless-compatible (unlike standard sessions)
- JWT gives microservices flexibility for stateless auth
- No third-party lock-in (cost and control benefits)

### Consequences

**Positive (benefits)**:
- Frictionless user experience (no password creation)
- Serverless-compatible session management
- Microservices-friendly (stateless JWT)
- Full control over auth logic and billing
- Cost-effective (no Auth0/Clerk fees)

**Negative (drawbacks)**:
- Magic-link requires email deliverability setup
- Custom implementation (longer time to market)
- Security requires careful implementation (no built-in protections)
- No OAuth provider flexibility

**Neutral (unaffected)**:
- User flow: login → verify email → access
- Admin dashboard: user management
- Security logging: all security events tracked

### Observable Signal (migration gate)

**Signal**: Email deliverability issues or user registration drop

**Trigger**:
- If signal → A: Optimize email deliverability, monitor bounce rates
- If signal → B: Add password option as fallback
- If signal → C: Integrate third-party auth (e.g., OAuth Google)

**Rationale**: Email deliverability is the bottleneck for magic-link; signals here indicate if fallback is needed.

### Rollback Plan

1. Step 1: Add password option alongside magic-link
2. Step 2: Add OAuth provider support (Google/GitHub)
3. Step 3: Consider migrating to third-party provider if magic-link fails

**Cost of rollback**: low (additive changes, no breaking changes)

### References

- Related decisions: [[ADR-0038]]
- Source material: Research in `research/specific-services-research.md` (Identity Service section)

---

## ADR-0036: Default AI Model Management

### Context
Z.ai balance can fluctuate or run low, need flexible fallback to maintain service availability without manual intervention.

### Decision
- **Default model**: GLM-4.5-Flash (cost-effective)
- **Fallback**: Anthropic Air/5.2 (when Z.ai balance is insufficient)
- **Monitoring**: Track per-endpoint balance and failover automatically
- **Rate limiting**: Apply conservative limits to prevent balance depletion

**Key choices**:
- Automatic failover when balance < threshold (e.g., 50%)
- Manual override via config for testing/failure scenarios
- Separate quotas per AI path (/paas vs /api/anthropic)

**Why this choice?**
- Ensures service availability even with budget constraints
- Cost optimization by using GLM-4.5-Flash by default
- Automatic failover provides resilience without manual intervention
- Conservative rate limiting prevents unexpected balance depletion

### Consequences

**Positive (benefits)**:
- High service availability (automatic failover)
- Cost optimization (use cheaper model when possible)
- No manual intervention needed for balance changes
- Graceful degradation when budget is tight

**Negative (drawbacks)**:
- More complex routing logic
- Potential performance variance (different models have different speeds)
- Need to monitor both vendor balances

**Neutral (unaffected)**:
- User-facing API unchanged
- Evaluation quality may vary slightly between models
- Learning curves for developers on new fallback logic

### Observable Signal (migration gate)

**Signal**: Fallback rate or manual intervention frequency

**Trigger**:
- If signal → A: Maintain current dual-model setup, optimize thresholds
- If signal → B: Consider additional fallback models (e.g., OpenRouter multi-model)
- If signal → C: Consolidate to single provider

**Rationale**: Monitoring fallback rate identifies whether the dual-model strategy is necessary or if optimization is needed.

### Rollback Plan

1. Step 1: Set GLM-4.5-Flash as primary, Air/5.2 as secondary
2. Step 2: Remove automatic failover, make failover manual
3. Step 3: Update all routing logic to use primary-only

**Cost of rollback**: medium (requires routing logic changes)

### References

- Related decisions: [[ADR-0035]]
- Source material: `/Users/tienphat/.claude/projects/-Users-tienphat-Developer-ielts/memory/zai-paas-vs-anthropic-separate-quota.md`

---

## ADR-0035: First Live Z.ai Round-Trip with GLM-4.5-Flash

### Context
First production integration with Z.ai / GLM API, need to validate cost structure, performance, and feature set while avoiding premature vendor lock-in.

### Decision
- **Model**: Use GLM-4.5-Flash as default model
- **Flow**: Direct API calls (not PaaS wrapper)
- **Fallback**: Revert to Anthropic Air/5.2 when balance is insufficient
- **Cost tracking**: Monitor per-endpoint spending separately

**Key choices**:
- Default model: GLM-4.5-Flash (fast, cost-effective for evaluation tasks)
- Bypass PaaS (not needed for single-model usage)
- Separate quotas for /paas vs /api/anthropic

**Why this choice?**
- GLM-4.5-Flash provides 30-40% cost savings vs Anthropic
- Fast response times suitable for evaluation tasks
- Direct API gives full control and observability
- Separate quotas prevent cross-contamination between AI paths

### Consequences

**Positive (benefits)**:
- Significant cost savings on AI evaluation
- Faster response times for synchronous operations
- Full API control and debugging capabilities
- Clear cost attribution per endpoint

**Negative (drawbacks)**:
- Need to manage two separate quota balances
- Vendor-specific API changes require manual updates
- No built-in retry/marketplace features

**Neutral (unaffected)**:
- Authentication layer unchanged
- Monitoring dashboards need dual quota tracking
- User experience unchanged

### Observable Signal (migration gate)

**Signal**: Cross-endpoint quota conflicts (1113 errors)

**Trigger**:
- If signal → A: Maintain current dual-quota setup, investigate optimization
- If signal → B: Consolidate to single AI path with custom retry logic
- If signal → C: Re-evaluate vendor strategy, consider multi-provider router

**Rationale**: This signal catches quota management issues early while keeping flexibility to consolidate if needed.

### Rollback Plan

1. Step 1: Remove GLM-4.5-Flash from default, set Air/5.2 as default
2. Step 2: Update all AI service endpoints to use Anthropic API
3. Step 3: Restore separate quota balance checks

**Cost of rollback**: medium (requires code changes across all AI endpoints)

### References

- Related decisions: [[ADR-0036]]
- Source material: `/Users/tienphat/.claude/projects/-Users-tienphat-Developer-ielts/memory/zai-paas-vs-anthropic-separate-quota.md`

---

## ADR-0034: Band Score Calculation Utilities

### Context
Need accurate band score calculations for IELTS Writing evaluation with proper rounding rules.

### Decision
Create utility functions for band score calculation and JSON response parsing.

### Rationale
1. **Why now?**: Core functionality required for evaluation system.
2. **Why not later?**: Essential for platform MVP.
3. **Why not another approach?**: Direct implementation provides clear, maintainable code.

### Observable Signal
- Tests pass with edge cases handled
- Rounding logic verified against IELTS standards
- JSON parsing error-free

### Migration Path
- Created utility functions
- Implemented comprehensive tests
- Integrated with evaluation system

### Rollback Path
- Remove utility functions and tests
- Revert to previous state

### Resolution Metrics
- ✅ All test cases pass
- ✅ Edge cases handled
- ✅ Integration working

---

## ADR-0045: Writing Service Retry Logic

### Context
Z.ai API integration for writing evaluation needs error handling strategy for API failures.

### Decision
- **Retry Strategy**: Retry once with exponential backoff on API failure
- **Failure Handling**: Return error with 5xx status if retry exhausted, log 4xx errors
- **Timeout**: 30-second timeout for scoring requests
- **Retry Delays**: 1s → 2s → exhaust (max 3 total attempts including initial)

**Key choices**:
- Single retry (not multiple) - balances reliability vs. user latency
- Exponential backoff - prevents rate limiting and server load spikes
- 4xx errors (client errors) logged and returned immediately - no retry needed
- 5xx errors (server errors) retried - temporary failures

**Why this choice?**
- Single retry is sufficient for transient failures (network, API overload)
- Exponential backoff prevents creating cascading failures
- 30-second timeout prevents indefinite blocking
- Clear distinction between client errors (no retry) and server errors (retry)

### Consequences

**Positive (benefits)**:
- High reliability for transient failures
- User-friendly error messages (no silent failures)
- Prevents rate limiting issues
- Clear error boundaries

**Negative (drawbacks)**:
- 1-second latency for failed requests
- Needs proper logging for monitoring
- User may see "retrying" UI state briefly

**Neutral (unaffected)**:
- Success path unchanged
- Success rates unaffected by retry logic
- Logging adds minimal overhead

### Observable Signal (migration gate)

**Signal**: Retry rate > 20% or error response time > 5s

**Trigger**:
- If signal → A: Increase retry count to 2 or reduce timeout to 20s
- If signal → B: Add circuit breaker pattern for repeated failures
- If signal → C: Investigate Z.ai API health and capacity

**Rationale**: High retry rate indicates API instability or client-side issues.

### Rollback Plan

1. Step 1: Remove retry logic, return error immediately
2. Step 2: Increase timeout to 60s
3. Step 3: Add detailed logging for investigation

**Cost of rollback**: low (remove retry wrapper, adjust timeout)

### References

- Related decisions: [[ADR-0035]] (Z.ai integration), [[ADR-0036]] (Model Management)
- Source material: Retry patterns best practices, Z.ai API documentation

---

## ADR-0044: Writing Service Task Type Support

### Context
Writing evaluation needs to support multiple IELTS task types (Task 1 vs Task 2) for completeness.

### Decision
- **MVP Scope**: Support both Task 1 (Letter/Report) and Task 2 (Essay) for initial release
- **Prompt Templates**: Separate templates for Task 1 and Task 2 with specific rubric instructions
- **Evaluation Criteria**: Both share 4 core criteria (TA, CC, LR, GRA)
- **Task 1 Specialization**: Additional task-specific instructions (format, length, audience)
- **Task 2 Specialization**: Standard essay structure guidance

**Key choices**:
- Both task types in MVP (completeness over minimal scope)
- Separate prompt templates for clarity
- Same 4 criteria for both (standard IELTS rubric)
- Task 1-specific content in prompt (letter format, report format)
- Task 2-specific content in prompt (introduction, body paragraphs, conclusion)

**Why this choice?**
- Writing is primary skill, users expect both task types
- No significant technical overhead (just different prompts)
- Better user experience (no "wait for next version")
- Aligns with IELTS test structure (both types on exam day)

### Consequences

**Positive (benefits)**:
- Complete writing evaluation experience (Task 1 + Task 2)
- Clear prompt structure reduces ambiguity
- Better user satisfaction
- Future-proof (no need for rewrite)

**Negative (drawbacks)**:
- Slightly more prompt engineering work
- Need to maintain two templates
- Evaluation quality varies by task type (requires calibration)

**Neutral (unaffected)**:
- Band score calculation unchanged (same 4 criteria)
- Progress tracking unchanged
- User interface unchanged

### Observable Signal (migration gate)

**Signal**: Task 1 and Task 2 accuracy within 5% of each other on official samples

**Trigger**:
- If signal → A: Accept both task types in MVP
- If signal → B: Only include Task 2 in MVP, add Task 1 later
- If signal → C: Calibrate prompts separately for each task type

**Rationale**: Accuracy parity between task types indicates fair evaluation.

### Rollback Plan

1. Step 1: Remove Task 1 support, keep only Task 2
2. Step 2: Update prompt template
3. Step 3: Update documentation

**Cost of rollback**: low (remove task type selection, adjust prompts)

### References

- Related decisions: [[ADR-0035]] (Z.ai integration), [[ADR-0043]] (Feedback storage)
- Source material: IELTS Writing Task 1 vs Task 2 guidelines

---

## ADR-0043: Writing Service Feedback Storage Strategy

### Context
Need to store AI evaluation feedback for writing submissions (per-criterion scores, strengths, improvements).

### Decision
- **Storage Strategy**: Embed feedback in `submissions` table as JSONB field
- **Feedback Structure**: `score_breakdown` and `feedback` objects nested in `submissions` row
- **No Separate Table**: Use single table approach for MVP (simpler queries, less joins)
- **Field Names**: `score_breakdown` (per-criterion scores), `feedback` (AI feedback text)
- **Query Patterns**: Single JSONB query, no JOINs required

**Key choices**:
- Embedded in submissions table (no separate feedbacks table)
- JSONB storage (flexible, efficient for nested data)
- MVP-only approach (no need for advanced queries)
- Maintainability over performance (performance acceptable at scale)

**Why this choice?**
- Simpler database schema (fewer tables)
- Fewer JOINs (better performance at MVP scale)
- JSONB flexibility (easy to add fields without schema changes)
- Sufficient for MVP requirements (all feedback needed is present)

### Consequences

**Positive (benefits)**:
- Simpler database schema (1 table instead of 2)
- Faster queries (no JOINs)
- Easier to implement (less boilerplate)
- JSONB flexibility (future-proof)

**Negative (drawbacks)**:
- Cannot query individual feedback rows separately
- No granular feedback history (one feedback per submission)
- Update complexity (UPDATE whole row vs. single row)

**Neutral (unaffected)**:
- User-facing API unchanged
- Band score calculation unchanged
- Progress tracking unchanged

### Observable Signal (migration gate)

**Signal**: Feedback queries taking > 100ms or query complexity > 3 levels

**Trigger**:
- If signal → A: Keep embedded approach, optimize queries
- If signal → B: Create separate feedbacks table for advanced queries
- If signal → C: Use full-text search on feedback content

**Rationale**: Query performance indicates if schema needs optimization.

### Rollback Plan

1. Step 1: Create feedbacks table
2. Step 2: Migrate feedback data from submissions
3. Step 3: Update queries to use new schema

**Cost of rollback**: medium (migration + query updates)

### References

- Related decisions: [[ADR-0034]] (Band Score), [[ADR-0045]] (Retry Logic)
- Source material: Database schema best practices (embedded vs. normalized)

---

## ADR-0042: Writing Service API Ownership

### Context
Need to decide where to implement CRUD endpoints for writing submissions (Writing Service vs. AI Service).

### Decision
- **API Ownership**: Writing Service owns submission CRUD endpoints
- **AI Service**: Only owns `/api/ai/score` (evaluation logic) and `/api/ai/feedback/:id` (feedback retrieval)
- **Service Boundaries**: Writing Service handles domain-specific operations (create, read, update, delete submissions)
- **AI Service**: Handles AI-specific operations (scoring, feedback generation)
- **Dependency Direction**: AI Service is called by Writing Service for scoring, not vice versa

**Key choices**:
- Writing Service for CRUD (clean domain separation)
- AI Service for evaluation (clear utility vs. domain)
- No API overlap (each service has distinct responsibilities)
- Clear integration contract (AI endpoints are called by Writing Service)

**Why this choice?**
- Clean separation of concerns (Writing = domain, AI = utility)
- Easier testing (domain logic isolated from AI integration)
- Better scalability (Writing Service can evolve independently)
- Clear ownership (no ambiguity about who maintains code)

### Consequences

**Positive (benefits)**:
- Clear service boundaries
- Easier testing and maintenance
- Independent scaling potential
- Clear integration contract

**Negative (drawbacks)**:
- More endpoints to implement (3 extra Writing endpoints)
- Need to manage service-to-service calls
- Potential API versioning complexity (separate endpoints)

**Neutral (unaffected)**:
- User experience unchanged
- Performance unchanged
- Database schema unchanged

### Observable Signal (migration gate)

**Signal**: API ownership ambiguity causes merge conflicts or unclear responsibilities

**Trigger**:
- If signal → A: Keep current ownership model
- If signal → B: Move some endpoints to AI Service
- If signal → C: Create shared API utilities

**Rationale**: Ownership ambiguity causes technical debt over time.

### Rollback Plan

1. Step 1: Move endpoints from Writing to AI Service
2. Step 2: Update service definitions
3. Step 3: Update API documentation

**Cost of rollback**: medium (code reorganization, documentation updates)

### References

- Related decisions: [[ADR-0037]] (Authentication), [[ADR-0038]] (Sandbox Strategy)
- Source material: Service architecture best practices

---

## ADR-0041: Writing Service Media Storage Deferred

### Context
Writing submissions may require PDF/task prompt uploads, but Media Service is in Blocked state.

### Decision
- **MVP Scope**: Text-only submissions (no PDF/task prompt support)
- **Future Support**: Add Media Service in Phase 2 for PDF upload/download
- **Immediate Work**: Focus on text input (textarea in UI, string in database)
- **No File Upload Endpoints**: Omit `/api/writing/submissions/:id/attachments` from MVP
- **User Experience**: Users copy-paste text into textarea (no file upload)

**Key choices**:
- Skip PDF support in MVP (minimal viable feature set)
- Text-only submissions (simplest implementation)
- Media Service deferred to Phase 2 (after AI evaluation stable)
- Clear scope boundary (MVP = text only, full = PDF + text)

**Why this choice?**
- Faster MVP (no Media Service integration)
- Simpler implementation (no file upload/download logic)
- Text input is sufficient for initial validation
- Can add PDF support later without breaking changes

### Consequences

**Positive (benefits)**:
- Faster MVP delivery (skip Media Service)
- Simpler implementation (no file handling)
- Lower complexity (less error handling)
- Clear scope boundary (known missing features)

**Negative (drawbacks)**:
- Missing PDF support (incomplete user experience)
- Users must copy-paste text (less convenient)
- Requires UI redesign for text-only input
- Need to inform users about MVP limitations

**Neutral (unaffected)**:
- AI evaluation unchanged
- Band score calculation unchanged
- Progress tracking unchanged

### Observable Signal (migration gate)

**Signal**: User feedback indicates need for PDF upload or copy-paste UX is problematic

**Trigger**:
- If signal → A: Add Media Service parallel to AI Service (Phase 1.5)
- If signal → B: Keep text-only, improve copy-paste UX
- If signal → C: Add base64 encoding for small files (workaround)

**Rationale**: User feedback indicates missing feature impacts adoption.

### Rollback Plan

1. Step 1: Add Media Service to Phase 1 parallel work
2. Step 2: Implement file upload endpoints
3. Step 3: Update UI for file input

**Cost of rollback**: low (additive features, no breaking changes)

### References

- Related decisions: [[ADR-0037]] (Authentication), [[ADR-0038]] (Sandbox Strategy)
- Source material: MVP best practices (minimum viable feature set)

---

## Pending Decisions

Decisions to be captured as we implement:

- **Database strategy**: Per-service vs shared database (referenced in research, not yet formalized)
- **Service communication**: REST vs GraphQL vs message queue (research suggests REST + background jobs)
- **Error handling patterns**: Specific error classes and retry strategies
- **Third-party integrations**: Beyond auth (payment, notification, media storage)

**Next steps**: Capture these decisions as ADRs during Phase 3/4 implementation.

---

## Hidden Dependencies (Phase 3 Findings)

### Dependency-001: AI Service → Skill Services
**Description**: AI Service provides evaluation endpoints but doesn't document dependencies on Speaking, Reading, Listening, Grammar, Vocabulary services.
**Impact**: Unclear integration contract, ownership of evaluation logic.
**Resolution**: Document in service definitions.

### Dependency-002: Media Service → AI Service
**Description**: Media Service has speech analysis but no standardized interface to AI Service.
**Impact**: Direct calls, unclear ownership, no standardized interface.
**Resolution**: Standardize Media → AI interface in AI Service.

### Dependency-003: Notification Service → All Services
**Description**: No clear documentation of which services trigger notifications or how templates are managed.
**Impact**: Unknown notification triggers, unclear content ownership.
**Resolution**: Define notification triggers in each service.

### Dependency-004: Progress Service → Skill Services
**Description**: Progress Service tracks scores but doesn't document data sources or update patterns.
**Impact**: Unclear data flow, unclear who updates progress when scores change.
**Resolution**: Document score data sources in Progress Service.

### Dependency-005: Analytics Service → All Services
**Description**: Analytics Service aggregates data but no documented data collection pattern (real-time vs batch).
**Impact**: Unknown data ownership model, unclear collection strategy.
**Resolution**: Define data collection pattern in Analytics Service.

---

## References
- IELTS Official Band Descriptors
- Next.js Documentation
- Drizzle ORM Documentation
- Neon Database Documentation
- OpenRouter Documentation
- Stripe Documentation