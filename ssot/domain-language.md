# Domain Language

## Core Terms

### Platform Components
- **Wedge**: Entry feature that drives user acquisition and validates product-market fit
- **Product**: Full platform covering all four IELTS skills
- **Ship-criterion**: Measurable targets for platform success
- **User**: Learner taking IELTS preparation

### IELTS Specific
- **Task 1**: Writing task (graph/diagram description)
- **Task 2**: Writing task (essay)
- **TA**: Task Achievement (writing criterion)
- **CC**: Coherence and Cohesion (writing criterion)
- **LR**: Lexical Resource (writing criterion)
- **GRA**: Grammatical Range and Accuracy (writing criterion)
- **Band Score**: 0-9 scale for IELTS assessment
- **Redraft**: Improved version after feedback
- **Essay**: Type 2 writing task requiring argumentative structure
- **Submission**: User's written work submitted for evaluation
- **Feedback**: AI-generated assessment with strengths, improvements, and scores
- **Learner**: User taking IELTS preparation course
- **Micro-skill**: Specific sub-skill within a major skill (e.g., verb tense in grammar)
- **Learning Method**: Strategy for learning and practicing (e.g., spaced repetition, rote memorization)
- **Progress**: Measurable improvement over time (scores, completion rate, skill acquisition)
- **Magic Link**: Passwordless login via email verification
- **Session**: Active authenticated period (JWT token validity)

### Technical Terms
- **SSOT**: Single Source of Truth
- **ADR**: Architecture Decision Record
- **Next.js**: Framework for building web applications
- **Drizzle ORM**: TypeScript ORM for PostgreSQL
- **Neon**: Serverless PostgreSQL database
- **OpenRouter**: AI service provider
- **Turbopack**: Next.js bundler for faster development

### Metrics
- **Acquisition**: Number of new users
- **Engagement**: Time spent and features used
- **Retention**: Return rate within specified period
- **Product-value**: Post-feedback behavior and score improvement

## Naming Conventions

### Files and Directories
- Use kebab-case for all files and folders
- Example: `user-management.md`, `api-routes/`

### Components
- Use PascalCase for React components
- Example: `WritingFeedback`, `BandScoreCalculator`

### Variables and Functions
- Use camelCase for JavaScript/TypeScript
- Example: `calculateBandScore`, `userProgress`

### Constants
- Use UPPER_SNAKE_CASE for constants
- Example: `MAX_REDRAFTS`, `SHIP_CRITERION`

### Database Tables
- Use snake_case for table names
- Example: `user_progress`, `writing_tasks`

## Communication Patterns

### Decision Making
- Reference existing ADRs
- Use observable signals for evolution
- Document trade-offs explicitly

### Documentation
- Keep documentation current
- Link related documents
- Use clear, concise language

### Code Comments
- Comment why, not what
- Reference ADRs for context
- Keep comments up-to-date with code
