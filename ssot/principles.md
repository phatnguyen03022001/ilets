# Principles

## Core Design Principles

### Simple over Clever
Prefer simple, maintainable code over clever, complex solutions.
- Clear code > clever code
- Readability > cleverness
- Maintainability > cleverness

### Explicit over Implicit
Make assumptions explicit, document them, don't hide them.
- Explicit requirements > hidden assumptions
- Documented decisions > implicit knowledge
- Self-documenting code > clever tricks

### Domain before Technology
Design around domain first, technology second.
- Domain problems first > tech solutions first
- Business logic > technical abstraction
- User needs > technology trends

### Evolution over Prediction
Evolve based on real usage, not predicted needs.
- Measure before optimize
- Ship early, iterate based on data
- Predictions → evidence → iteration

### One Question → One Document → One Authority
Each question needs one document, each document has one authority.
- Question → Single answer document
- Document → Single source of truth
- Authority → Single owner

## Protocols

### Gap Protocol
1. **DETECT** - Detect a gap: Implementation needs Authority but doesn't have one
2. **RECORD** - Record gap in open-decisions.md with classification
3. **STOP** - STOP implementation immediately
4. **RESUME** - Await Founder/Vision Assistant approval before resuming

### Loop Protocol
1. **IDENTIFY** - Identify a pattern or issue during implementation
2. **UPDATE** - Update Authority file to reflect findings
3. **VERIFY** - Verify update doesn't violate any Core Laws
4. **RECORD** - Record decision as ADR if it's architectural
