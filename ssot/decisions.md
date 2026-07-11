# Architecture Decision Records

## ADR Index

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

## ADR-0040: AI Evaluation Accuracy Target

### Context
Need realistic accuracy target for AI evaluation in Writing wedge to balance quality and development velocity.

### Decision
- **Target**: 90% accuracy on official IELTS samples (formative feedback phase)
- **Upgrade**: Raise to 95% when real user data available
- **Baseline**: 100 official IELTS samples for initial calibration
- **Iterative improvement**: A/B test different prompts, models, calibration data

**Key choices**:
- 90% as MVP baseline (acceptable for formative feedback)
- 95% target for production grade (after real user data validation)
- 100 official samples for calibration (comprehensive coverage)

**Why this choice?**
- 90% is realistic for MVP (industry standard for AI evaluation)
- 95% achievable with real user data and iterative improvement
- 100 official samples provides solid baseline without over-engineering
- Allows progressive improvement rather than perfect from start

### Consequences

**Positive (benefits)**:
- Realistic MVP timeline (avoid perfectionism trap)
- Clear quality threshold for release
- Progressive improvement (formative → production grade)
- Data-driven optimization (real user data > official samples)

**Negative (drawbacks)**:
- Initial accuracy may feel low to stakeholders
- May need re-calibration as real user data comes in
- Trust building requires evidence collection

**Neutral (unaffected)**:
- Evaluation quality monitored continuously
- User feedback collected for improvement
- Error tracking in place for debugging

### Observable Signal (migration gate)

**Signal**: Accuracy reaches 90% on official samples OR real user data collected

**Trigger**:
- If signal → A: Release MVP with 90% accuracy
- If signal → B: Add more calibration samples, iterate prompts
- If signal → C: Collect real user data, target 95%

**Rationale**: Real user data provides more meaningful accuracy than official samples alone.

### Rollback Plan

1. Step 1: Increase calibration samples to 200
2. Step 2: Test alternative prompt strategies
3. Step 3: Re-evaluate accuracy target

**Cost of rollback**: low (additive data, no code changes)

### References

- Related decisions: [[ADR-0035]], [[ADR-0036]]
- Source material: IELTS evaluation research, industry AI evaluation standards

---

## ADR-0039: Implementation Sequence (Writing Wedge E)

### Context
Need clear implementation sequence for Writing wedge to ensure dependencies are satisfied and timeline is realistic.

### Decision
**Priority Order**:
1. **Identity Service** (3-5 days) - Foundation, must be first
2. **Band Score Utilities** (1-2 days) - Core calculation logic, independent
3. **AI Service - Basic Evaluation** (4-6 days) - Core evaluation, depends on utilities
4. **Progress Service - Writing Integration** (2-3 days) - Progress tracking, depends on AI service
5. **[Blocked] Writing Service** (3-4 days) - Domain logic, blocked until AI stable
6. **[Blocked] Media Service** (2-3 days) - Upload support, can parallelize partially
7. **[Blocked] Notification Service** (2-3 days) - Email support, blocked until AI ready

**Critical Path**: Identity → Band Score → AI → Progress → [blocked services]

**Total Critical Path**: 10-16 days
**Overall Timeline**: 4-6 weeks (including parallel work)

**Dependencies**:
- Identity Service → All services (no users = no features)
- Band Score → AI Service (evaluation needs scoring)
- AI Service → Progress Service (evaluation data)
- AI Service → Writing Service (evaluation output)
- AI Service → Notification Service (result delivery)

### Key Choices

- **Identity first**: Authentication is foundational, nothing works without users
- **Band Score utilities**: Independent, high value, blocks AI evaluation
- **AI Service core**: Primary feature, determines MVP quality
- **Progress Service secondaries**: Track evaluation results, clear dependency
- **Parallel work**: Media Service can do basic uploads while AI is being built

**Why this choice?**
- Minimizes blocking dependencies
- Clear critical path for timeline estimation
- Allows parallel work where possible (Media Service)
- Ensures core quality (AI evaluation) is validated first
- Progressive build (formative feedback → production grade)

### Consequences

**Positive (benefits)**:
- Clear implementation order (no bottlenecks)
- Realistic timeline (10-16 days core)
- Parallel work efficiency (Media Service can start early)
- Quality-first approach (AI evaluation validated first)

**Negative (drawbacks)**:
- Slower overall timeline (4-6 weeks)
- Services blocked until dependencies satisfied
- Need careful coordination for parallel work

**Neutral (unaffected)**:
- User experience: No changes
- Feature completeness: Core features only (formative feedback)
- Documentation: Already complete from Phase 4

### Observable Signal (migration gate)

**Signal**: Identity Service functional with users created and authenticated

**Trigger**:
- If signal → A: Proceed to Band Score utilities
- If signal → B: Delay, fix authentication issues
- If signal → C: Re-evaluate implementation sequence

**Rationale**: Identity is the foundation; nothing works without users.

### Rollback Plan

1. Step 1: Shuffle sequence (e.g., Media Service first)
2. Step 2: Remove parallel work constraint
3. Step 3: Simplify to critical path only

**Cost of rollback**: low (sequence changes, no code impact)

### References

- Related decisions: [[ADR-0034]] (Band Score), [[ADR-0035]] (AI), [[ADR-0036]] (Model Selection)
- Source material: Phase 4 implementation planning research

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