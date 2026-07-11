# Principles

## 1. Documentation-First
- Write documentation before code
- Documents make decisions explicit
- Code implements decisions

## 2. SSOT as Source of Truth
- Single point of truth for all architectural decisions
- No concurrent truth sources
- All work must reference SSOT

## 3. Evolution Over Revolution
- Incremental changes based on observable signals
- Big architectural changes only when pain thresholds are hit
- Maintain backward compatibility whenever possible

## 4. Convention Over Configuration
- Establish clear naming conventions
- Default behaviors for common patterns
- Minimize custom configuration

## 5. Observability First
- Every decision must have observable signals
- Metrics before features
- Measurable outcomes

## 6. User Value First
- Every feature must deliver clear user value
- No technical debt without user benefit
- Ship-criterion must be met before scaling

## 7. Pragmatism Over Perfection
- Good enough now is better than perfect later
- Working solutions over elegant abstractions
- Evidence-based decisions

## 8. Learning as Currency
- Track metrics rigorously
- Learn from every deployment
- Iterate based on data

## 9. Security By Default
- Privacy-first design
- Minimal data collection
- Explicit consent models

## 10. Sustainability
- No service that cannot be maintained by one person
- Clear documentation for all systems
- Avoid vendor lock-in where possible