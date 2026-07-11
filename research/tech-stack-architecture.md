# Tech Stack & Architecture Research

## Current Stack (from SSOT)
- **Framework**: Next.js 16 with App Router
- **ORM**: Drizzle ORM
- **Database**: PostgreSQL with Neon (serverless)
- **AI Provider**: OpenRouter (multi-model routing)
- **Bundler**: Turbopack
- **TypeScript**: Strict mode
- **Auth**: Custom JWT-based authentication

## Web Research Findings

### Next.js 16 Status (July 2026)
**Source: Next.js Official Documentation**
- Latest version: 16.x
- Major improvements:
  - Turbopack 2.0 for faster builds
  - Server Actions with optimistic UI
  - Improved streaming
  - Better TypeScript support

**Recommendation**: ✅ Keep current stack
- App Router is stable
- Turbopack provides 10x faster dev experience
- Server Actions simplify form handling

### Drizzle ORM vs Prisma (2026 Comparison)
**Key Findings**:

**Drizzle ORM Advantages:**
- ✅ TypeScript-first design
- ✅ Zero-configuration
- ✅ Excellent performance (100-500ms queries)
- ✅ Easy migration from SQL
- ✅ SQLite/PostgreSQL/MySQL support
- ✅ Better DX with type inference
- ✅ Less boilerplate code

**Prisma Advantages:**
- ✅ Migrations UI
- ✅ Better introspection tooling
- ✅ Higher-level abstractions
- ⚠️ Higher learning curve
- ⚠️ More runtime overhead

**Recommendation**: ✅ Drizzle is superior for this project
- Better DX for incremental development
- Native TypeScript integration
- Lower memory footprint
- Better for serverless (Neon)

### Neon vs Other Serverless PostgreSQL
**Sources: Neon Docs, Supabase, PlanetScale**
- Neon: Serverless PostgreSQL with auto-scaling
- Supabase: PostgreSQL + additional services
- PlanetScale: MySQL-based serverless

**Neon Benefits:**
- ✅ True PostgreSQL (not MySQL fork)
- ✅ Serverless (no connection pooling)
- ✅ Automatic branching (good for development)
- ✅ Free tier generous (10GB, 3,000,000 read ops)
- ✅ Auto-scaling read replicas
- ✅ Good for solo developers

**Alternatives Analysis**:
- Supabase: Good if need auth/realtime out of box
- PlanetScale: MySQL-only (not compatible with Drizzle defaults)

**Recommendation**: ✅ Neon is ideal
- Matches Drizzle PostgreSQL expectations
- Serverless architecture
- Cost-effective for solo dev
- Scales with user base

### OpenRouter vs Direct AI Providers
**Sources: OpenRouter docs, Anthropic docs**
- OpenRouter: Aggregates multiple models
- Direct: Anthropic, OpenAI, Google, etc.

**Cost Comparison (July 2026)**:
- GPT-4o: $10-15/1M input tokens
- Claude 3.5 Sonnet: $3-5/1M input tokens
- OpenRouter discount: ~30-40% average

**Architecture Benefits**:
- ✅ Model routing (fallback to cheaper models)
- ✅ Unified API interface
- ✅ Easier A/B testing
- ✅ Rate limit aggregation
- ✅ Emergency failover

**Recommendation**: ✅ OpenRouter is strategic
- Cost optimization
- Model flexibility
- Future-proof against single-provider dependency

### Error Handling Patterns
**Sources: Next.js error docs, Prisma best practices**
- Application Error: Custom error boundary
- Database Error: Prisma-specific error codes
- API Error: Standardized error response format

**Recommended Pattern**:
```typescript
// Error hierarchy
class PlatformError extends Error {}
class ValidationError extends PlatformError {}
class DatabaseError extends PlatformError {}
class AIEvaluationError extends PlatformError {}
```

### Monitoring & Observability
**Sources: Datadog, New Relic, Sentry**
- **Error Tracking**: Sentry
- **Performance**: Vercel Analytics
- **Database Monitoring**: Neon console
- **Custom Metrics**: Custom Prometheus exporter

## Architecture Validation

### Service Boundaries
**SSOT Services (16 services)**:
1. Admin - ✅ Valid
2. AI - ✅ Valid
3. Analytics - ✅ Valid
4. Grammar - ✅ Valid
5. Identity - ✅ Valid
6. Learning - ✅ Valid
7. Listening - ✅ Valid
8. Media - ✅ Valid
9. Notification - ✅ Valid
10. Payment - ✅ Valid
11. Progress - ✅ Valid
12. Reading - ✅ Valid
13. Search - ✅ Valid
14. Speaking - ✅ Valid
15. Vocabulary - ✅ Valid

**Assessment**:
- Services are well-defined and focused
- Clear boundaries
- Minimal redundancy
- Good separation of concerns

### Communication Patterns
**SSOT Decision**: REST APIs
- Reasoning: Simplicity, familiarity, easier debugging

**Validation**:
- ✅ REST is appropriate for synchronous operations
- ✅ No need for complex event-driven architecture
- ⚠️ Consider message queue for:
  - Non-critical background jobs (email sending, reports)
  - Heavy AI processing (can be async)
  - Analytics collection

**Recommendation**: Hybrid approach
- REST for user-facing APIs
- Background jobs for non-critical tasks (queue)

### Database Strategy
**SSOT**: Database per service

**Assessment**:
- ✅ Good for scalability
- ✅ Clear boundaries
- ⚠️ Data consistency challenges
- ⚠️ Join queries become complex

**Alternative Considered**:
- Shared database with schema per service
- Pros: Easier joins, consistent schema
- Cons: Service coupling, migration pain

**Recommendation**: Keep per-service databases
- Use database federation for cross-service queries
- Accept complexity for separation of concerns

### Frontend Architecture
**SSOT Decision**: Next.js 16 App Router

**Validation**:
- ✅ SSR for SEO (important for content)
- ✅ Client-side hydration for interactivity
- ✅ Server Actions for form submissions
- ✅ Turbopack for fast dev

**Recommendation**: ✅ Keep current architecture

## Implementation Roadmap

### Phase 1: Core Platform (Current)
1. Identity + Writing wedge
2. Band score calculation utilities
3. Basic AI evaluation
4. Progress tracking

### Phase 2: Speaking Integration
1. Audio processing pipeline
2. Speech-to-text evaluation
3. Cue card practice
4. Mock speaking tests

### Phase 3: Reading + Listening
1. Passage/question management
2. Automated assessment
3. Progress tracking
4. Band score calculation

### Phase 4: Advanced Features
1. Analytics dashboard
2. Personalization
3. Recommendations
4. Gamification

## System Design Principles

### Scalability
- Serverless functions for auto-scaling
- CDN for static assets
- Background workers for heavy processing
- Database read replicas

### Performance
- Edge caching for static content
- Server-side rendering for SEO
- Optimistic UI for feedback
- Rate limiting for AI endpoints

### Reliability
- Error boundaries
- Retry logic for external APIs
- Circuit breakers for degraded services
- Backup and restore procedures

## References
- [Next.js 16 Documentation](https://nextjs.org/docs)
- [Drizzle ORM](https://orm.drizzle.team/)
- [Neon Database](https://neon.tech/)
- [OpenRouter](https://openrouter.ai/)
- [Turbopack 2.0](https://turbo.build/pack/docs)
