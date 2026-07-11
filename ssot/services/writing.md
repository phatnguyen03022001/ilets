# Writing Service

## Service Overview

The Writing Service manages user submissions for IELTS Writing tasks (Task 1: Letter/Report and Task 2: Essay) with domain-specific operations (CRUD) and AI-powered evaluation integration.

## Architecture Decisions

### ADR-0041: Writing Service Media Storage Deferred
**Status**: Resolved
**Date**: 2026-07-12

**Decision**: MVP scope is text-only submissions (no PDF/task prompt uploads). Media Service will be added in Phase 2 for PDF upload/download support.

**Key Points**:
- Skip PDF support in MVP (faster delivery, simpler implementation)
- Text-only submissions (textarea in UI, string in database)
- Clear scope boundary (MVP = text only, full = PDF + text)
- Users copy-paste text into textarea (no file upload)

**Why**:
- Faster MVP delivery (skip Media Service integration)
- Simpler implementation (no file handling)
- Text input is sufficient for initial validation

**Consequences**:
- ✅ Faster MVP delivery
- ✅ Simpler implementation (no file handling)
- ✅ Lower complexity (less error handling)
- ❌ Missing PDF support (incomplete user experience)
- ❌ Users must copy-paste text (less convenient)

**Observable Signal**: User feedback indicates need for PDF upload or copy-paste UX is problematic

**Rollback Plan**:
1. Add Media Service to Phase 1 parallel work
2. Implement file upload endpoints
3. Update UI for file input

---

### ADR-0042: Writing Service API Ownership
**Status**: Resolved
**Date**: 2026-07-12

**Decision**: Writing Service owns submission CRUD endpoints. AI Service only owns `/api/ai/score` (evaluation logic) and `/api/ai/feedback/:id` (feedback retrieval).

**Key Points**:
- Writing Service for CRUD (clean domain separation)
- AI Service for evaluation (clear utility vs. domain)
- No API overlap (each service has distinct responsibilities)
- Clear integration contract (AI endpoints are called by Writing Service)

**Why**:
- Clean separation of concerns (Writing = domain, AI = utility)
- Easier testing (domain logic isolated from AI integration)
- Better scalability (Writing Service can evolve independently)

**Consequences**:
- ✅ Clear service boundaries
- ✅ Easier testing and maintenance
- ✅ Independent scaling potential
- ❌ More endpoints to implement (3 extra Writing endpoints)
- ❌ Need to manage service-to-service calls

**Observable Signal**: API ownership ambiguity causes merge conflicts or unclear responsibilities

---

### ADR-0043: Writing Service Feedback Storage Strategy
**Status**: Resolved
**Date**: 2026-07-12

**Decision**: Store feedback in `submissions` table as JSONB field (embedded approach, not separate table).

**Key Points**:
- Embedded in submissions table (no separate feedbacks table)
- JSONB storage (flexible, efficient for nested data)
- MVP-only approach (simpler queries, less joins)
- Field names: `score_breakdown` (per-criterion scores), `feedback` (AI feedback text)

**Why**:
- Simpler database schema (fewer tables)
- Fewer JOINs (better performance at MVP scale)
- JSONB flexibility (easy to add fields without schema changes)
- Sufficient for MVP requirements

**Consequences**:
- ✅ Simpler database schema (1 table instead of 2)
- ✅ Faster queries (no JOINs)
- ✅ Easier to implement (less boilerplate)
- ✅ JSONB flexibility (future-proof)
- ❌ Cannot query individual feedback rows separately
- ❌ No granular feedback history (one feedback per submission)

**Observable Signal**: Feedback queries taking > 100ms or query complexity > 3 levels

---

### ADR-0044: Writing Service Task Type Support
**Status**: Resolved
**Date**: 2026-07-12

**Decision**: MVP supports both Task 1 (Letter/Report) and Task 2 (Essay) for completeness.

**Key Points**:
- Both task types in MVP (completeness over minimal scope)
- Separate prompt templates for Task 1 and Task 2
- Same 4 criteria for both (standard IELTS rubric)
- Task 1 specialization: Additional task-specific instructions (format, length, audience)
- Task 2 specialization: Standard essay structure guidance

**Why**:
- Writing is primary skill, users expect both task types
- No significant technical overhead (just different prompts)
- Better user experience (no "wait for next version")
- Aligns with IELTS test structure (both types on exam day)

**Consequences**:
- ✅ Complete writing evaluation experience (Task 1 + Task 2)
- ✅ Clear prompt structure reduces ambiguity
- ✅ Better user satisfaction
- ✅ Future-proof (no need for rewrite)
- ❌ Slightly more prompt engineering work
- ❌ Need to maintain two templates
- ❌ Evaluation quality varies by task type

**Observable Signal**: Task 1 and Task 2 accuracy within 5% of each other on official samples

---

### ADR-0045: Writing Service Retry Logic
**Status**: Resolved
**Date**: 2026-07-12

**Decision**: Retry once with exponential backoff on Z.ai API failure (1s → 2s → exhaust). 30-second timeout for scoring requests.

**Key Points**:
- Single retry (balances reliability vs. user latency)
- Exponential backoff (prevents rate limiting and server load spikes)
- 4xx errors (client errors) logged and returned immediately - no retry needed
- 5xx errors (server errors) retried - temporary failures
- Max 3 total attempts (including initial)

**Why**:
- Single retry is sufficient for transient failures (network, API overload)
- Exponential backoff prevents creating cascading failures
- 30-second timeout prevents indefinite blocking
- Clear distinction between client errors and server errors

**Consequences**:
- ✅ High reliability for transient failures
- ✅ User-friendly error messages (no silent failures)
- ✅ Prevents rate limiting issues
- ✅ Clear error boundaries
- ❌ 1-second latency for failed requests
- ❌ Needs proper logging for monitoring

**Observable Signal**: Retry rate > 20% or error response time > 5s

---

## Data Models

### Entity Definition
```typescript
interface WritingSubmission {
  id: string;
  userId: string;
  taskType: 'Task1' | 'Task2';
  taskPrompt: string;        // Task prompt text
  userContent: string;       // User's writing submission (text only in MVP)
  createdAt: Date;
  updatedAt: Date;
  score_breakdown: {
    TA: number;  // Task Achievement (1-9)
    CC: number;  // Coherence and Cohesion (1-9)
    LR: number;  // Lexical Resource (1-9)
    GRA: number; // Grammatical Range and Accuracy (1-9)
  };
  feedback: {
    overallBand: number;
    strengths: string[];
    improvements: string[];
    detailedFeedback: string;
    perCriterionFeedback: {
      TA: string;
      CC: string;
      LR: string;
      GRA: string;
    };
  };
}
```

### DTOs for API Layer
```typescript
interface CreateSubmissionDto {
  taskType: 'Task1' | 'Task2';
  taskPrompt: string;
  userContent: string;
}

interface UpdateSubmissionDto {
  userContent?: string;
}
```

## API Endpoints

### Writing Service (CRUD)
- `POST /api/writing/submissions` - Create new submission
- `GET /api/writing/submissions/:id` - Read submission by ID
- `PUT /api/writing/submissions/:id` - Update submission content
- `DELETE /api/writing/submissions/:id` - Delete submission

### AI Service (Evaluation)
- `POST /api/ai/score` - Submit writing for evaluation
- `GET /api/ai/feedback/:id` - Retrieve feedback for submission

## Integration Contract

### Writing Service → AI Service
```typescript
// Writing Service calls AI Service for scoring
POST /api/ai/score
{
  "taskType": "Task1",
  "taskPrompt": "Write a letter to...",
  "userContent": "Dear Sir/Madam,\n\nI am writing..."
}

// Response (from AI Service)
{
  "submissionId": "xxx",
  "score_breakdown": {
    "TA": 7.5,
    "CC": 8,
    "LR": 7,
    "GRA": 7.5
  },
  "feedback": {
    "overallBand": 7.5,
    "strengths": ["Good organization", "Varied vocabulary"],
    "improvements": ["Use more complex sentences"],
    "detailedFeedback": "...",
    "perCriterionFeedback": {
      "TA": "Task Achievement: Well addressed...",
      "CC": "Coherence and Cohesion: Excellent paragraphing...",
      "LR": "Lexical Resource: Good range of vocabulary...",
      "GRA": "Grammatical Range and Accuracy: Generally accurate..."
    }
  }
}
```

## Evaluation Criteria

### Task 1 (Letter/Report)
- **Task Achievement (TA)**: Addresses all parts of the task, presents a clear overview, highlights key features appropriately
- **Coherence and Cohesion (CC)**: Uses paragraphing logically, cohesive devices appropriately
- **Lexical Resource (LR)**: Uses a wide range of vocabulary with flexibility
- **Grammatical Range and Accuracy (GRA)**: Uses a wide range of structures, few errors

### Task 2 (Essay)
- **Task Response (TR)**: Addresses all parts of the task, presents a clear position throughout the essay
- **Coherence and Cohesion (CC)**: Uses paragraphs logically, uses cohesive devices effectively
- **Lexical Resource (LR)**: Uses vocabulary with flexibility and precision
- **Grammatical Range and Accuracy (GRA)**: Uses a wide range of structures, maintains good control

## Band Score Calculation

### 4 Core Criteria
All writing submissions are evaluated on the same 4 criteria:
1. **Task Achievement** (Task 1) / **Task Response** (Task 2)
2. **Coherence and Cohesion**
3. **Lexical Resource**
4. **Grammatical Range and Accuracy**

### Band Calculation
- Each criterion scored from 1 to 9
- Average of 4 criteria = overall band score
- Rounding rules (per IELTS standards)

### Utility Functions (from ADR-0034)
```typescript
// Band score calculation
export function calculateOverallBand(scores: {
  TA: number;
  CC: number;
  LR: number;
  GRA: number;
}): number {
  const sum = scores.TA + scores.CC + scores.LR + scores.GRA;
  const average = sum / 4;
  // Round to nearest 0.5 (per IELTS standards)
  const rounded = Math.round(average * 2) / 2;
  return Math.min(9, Math.max(1, rounded)); // Clamp to 1-9
}

// Parse AI response
export function parseScoreBreakdown(response: string): {
  TA: number;
  CC: number;
  LR: number;
  GRA: number;
} {
  // Parse from AI response
}
```

## Prompt Templates

### Task 1 Prompt Template
```
You are an IELTS Writing Task 1 evaluator.

Task Type: Task 1 (Letter/Report)

Instructions:
- Evaluate the submission based on the IELTS band descriptors
- Provide per-criterion scores (TA, CC, LR, GRA) from 1 to 9
- Give an overall band score rounded to nearest 0.5
- List strengths and improvements
- Provide detailed feedback per criterion

Submission:
Task Prompt: {taskPrompt}

User's Writing:
{userContent}

Output Format:
{
  "score_breakdown": {
    "TA": <number>,
    "CC": <number>,
    "LR": <number>,
    "GRA": <number>
  },
  "feedback": {
    "overallBand": <number>,
    "strengths": ["<string>", ...],
    "improvements": ["<string>", ...],
    "detailedFeedback": "<string>",
    "perCriterionFeedback": {
      "TA": "<string>",
      "CC": "<string>",
      "LR": "<string>",
      "GRA": "<string>"
    }
  }
}
```

### Task 2 Prompt Template
```
You are an IELTS Writing Task 2 evaluator.

Task Type: Task 2 (Essay)

Instructions:
- Evaluate the submission based on the IELTS band descriptors
- Provide per-criterion scores (TR, CC, LR, GRA) from 1 to 9
- Give an overall band score rounded to nearest 0.5
- List strengths and improvements
- Provide detailed feedback per criterion

Submission:
Task Prompt: {taskPrompt}

User's Writing:
{userContent}

Output Format:
{
  "score_breakdown": {
    "TR": <number>,
    "CC": <number>,
    "LR": <number>,
    "GRA": <number>
  },
  "feedback": {
    "overallBand": <number>,
    "strengths": ["<string>", ...],
    "improvements": ["<string>", ...],
    "detailedFeedback": "<string>",
    "perCriterionFeedback": {
      "TR": "<string>",
      "CC": "<string>",
      "LR": "<string>",
      "GRA": "<string>"
    }
  }
}
```

## Retry Logic

### Z.ai API Retry Strategy
- **Timeout**: 30 seconds for scoring requests
- **Retry**: Once on 5xx errors (exponential backoff: 1s → 2s)
- **No Retry**: 4xx errors (client errors) logged immediately
- **Max Attempts**: 3 total (including initial)

### Error Handling
```typescript
async function scoreWriting(
  taskType: 'Task1' | 'Task2',
  taskPrompt: string,
  userContent: string
): Promise<ScoreResult> {
  try {
    const response = await zaiClient.chat.completions.create({
      model: 'glm-4.5-flash',
      messages: [
        { role: 'system', content: getPromptTemplate(taskType) },
        { role: 'user', content: formatInput(taskPrompt, userContent) }
      ],
      temperature: 0.3,  // Lower temperature for consistent scoring
      max_tokens: 1000
    });

    return parseResponse(response.choices[0].message.content);
  } catch (error) {
    if (isServerSideError(error)) {
      // Retry once
      await sleep(1000);
      return retryScore(...);
    }
    // Client errors - log and return
    logger.error('AI scoring failed', { error });
    throw new ScoreError('Failed to score writing', 'AI_SERVICE_ERROR');
  }
}
```

## Testing Strategy

### Unit Tests
- Band score calculation utilities
- Prompt template generation
- Response parsing (JSON validation)
- Retry logic

### Integration Tests
- End-to-end writing evaluation flow
- API endpoint integration
- Database schema persistence
- Service-to-service communication

### Test Cases
```typescript
describe('Writing Service', () => {
  describe('Band Score Calculation', () => {
    it('should calculate overall band from 4 criteria', () => {
      const scores = { TA: 7, CC: 8, LR: 7, GRA: 7.5 };
      const overall = calculateOverallBand(scores);
      expect(overall).toBe(7.25);  // Average: 7.375 → 7.5
    });

    it('should clamp band score to 1-9 range', () => {
      const scores = { TA: 0, CC: 0, LR: 0, GRA: 0 };
      const overall = calculateOverallBand(scores);
      expect(overall).toBe(1);
    });
  });

  describe('Prompt Templates', () => {
    it('should generate Task 1 prompt correctly', () => {
      const taskType = 'Task1';
      const prompt = getPromptTemplate(taskType);
      expect(prompt).toContain('Task 1 (Letter/Report)');
      expect(prompt).toContain('Task Achievement');
    });

    it('should generate Task 2 prompt correctly', () => {
      const taskType = 'Task2';
      const prompt = getPromptTemplate(taskType);
      expect(prompt).toContain('Task 2 (Essay)');
      expect(prompt).toContain('Task Response');
    });
  });
});
```

## Observable Metrics

### Success Metrics
- Evaluation accuracy (band score correlation with human grading)
- Task 1 vs Task 2 accuracy parity
- User satisfaction with feedback quality

### Technical Metrics
- Retry rate (< 20% threshold)
- Average response time (< 5s)
- Error rate by type

## Dependencies

### Internal Services
- **AI Service**: Calls `/api/ai/score` for evaluation
- **Database**: Drizzle ORM with PostgreSQL

### External Services
- **Z.ai API**: GLM-4.5-Flash for evaluation (via direct API, not PaaS)
- **Anthropic Air/5.2**: Fallback when Z.ai balance is insufficient

### Related Decisions
- [[ADR-0035]] - Z.ai integration
- [[ADR-0036]] - Model management
- [[ADR-0038]] - Sandbox strategy
- [[ADR-0039]] - Implementation sequence
- [[ADR-0040]] - AI accuracy target

## References
- IELTS Official Writing Band Descriptors
- Z.ai API Documentation
- Research in `research/specific-services-research.md` (Writing section)
