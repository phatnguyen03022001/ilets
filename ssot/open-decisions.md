# Open Decisions

## Decisions Needing Resolution

### [ADR-0001] Database Choice
**Status**: Open  
**Question**: PostgreSQL vs MongoDB for IELTS platform
**Signal Needed**: User data complexity and query patterns

### [ADR-0002] AI Provider Selection
**Status**: Open  
**Question**: OpenRouter vs OpenAI for evaluation service
**Signal Needed**: Cost-effectiveness and evaluation accuracy

### [ADR-0003] Frontend Framework
**Status**: Open  
**Question**: Next.js 13 App Router vs React Server Components
**Signal Needed**: Development velocity and performance requirements

### [ADR-0004] Authentication Strategy
**Status**: Open  
**Question**: Custom auth vs Auth0 vs Firebase
**Signal Needed**: Security requirements and development time

### [ADR-0005] Payment Processing
**Status**: Open  
**Question**: Stripe vs PayPal integration
**Signal Needed**: Regional payment requirements and compliance

### [ADR-0006] Analytics Implementation
**Status**: Open  
**Question**: Custom analytics vs third-party service
**Signal Needed**: Custom metrics requirements and cost constraints

### [ADR-0007] CDN vs Local Storage
**Status**: Open  
**Question**: Cloudflare vs S3 vs local file storage
**Signal Needed**: Content delivery requirements and cost analysis

### [ADR-0008] Testing Strategy
**Status**: Open  
**Question**: Unit tests vs integration tests vs e2e tests
**Signal Needed**: Test coverage requirements and CI/CD pipeline needs

### [ADR-0009] Deployment Strategy
**Status**: Open  
**Question**: Vercel vs AWS vs Docker containerization
**Signal Needed**: Scaling requirements and cost optimization

### [ADR-0010] Internationalization
**Status**: Open  
**Question**: Full i18n vs English-only with localization
**Signal Needed**: Target market analysis and user demographics

## Decision Format Template

```markdown
# [ADR-XXXX] [Decision Title]

## Context
[Problem statement and background]

## Options Considered
[List all viable options]

## Rationale
1. Why now?
2. Why not later?
3. Why not another approach?

## Observable Signal
[Specific metric/condition that resolves the decision]

## Migration Path
[How to transition from current state to decision]

## Rollback Path
[How to revert if the decision fails]

## Status
[Current status: Open | In Progress | Resolved]
```

## Resolution Criteria
Each decision must meet all criteria before resolution:
1. Clear problem definition
2. At least 2 viable options considered
3. Observable signal identified
4. Migration and rollback paths defined
5. Decision documented in ADR format