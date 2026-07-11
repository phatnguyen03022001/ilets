# STOP_AND_ASK: Writing Service Foundation Requirements

## Date
2026-07-12

## Status
✅ Awaiting founder approval before implementation

---

## 1. API Endpoints (Writing MVP)

### Core Scoring & Feedback (Z.ai GLM-4.5-Flash)
- `POST /api/ai/score` - Band score calculation for writing submission
  - Request: `{ "content": string, "task_type": "essay" | "task1" | "task2" }`
  - Response: `{ "band_score": number, "score_breakdown": {...}, "feedback": {...}, "confidence": number }`

- `GET /api/ai/feedback/:evaluationId` - Retrieve AI feedback
  - Response: `{ "id": string, "band_score": number, "score_breakdown": {...}, "feedback": {...} }`

### Submission Management (CRUD)
- `POST /api/writing/submissions` - Create new writing submission
  - Request: `{ "assessment_id": string, "content": object, "time_spent": number, "metadata": object }`
  - Response: `{ "id": string, "assessment_id": string, "user_id": string, "status": "pending" | "scored", "created_at": string }`

- `GET /api/writing/submissions/:id` - Retrieve submission by ID
  - **MISSING from AI Service** - Needs to be in Writing Service
  - Response: `{ "id": string, "assessment_id": string, "user_id": string, "content": object, "score": number, "feedback": object, "status": string, "created_at": string }`

- `PUT /api/writing/submissions/:id` - Update submission (e.g., cancel, edit before scoring)
  - **MISSING from AI Service** - Needs to be in Writing Service
  - Request: `{ "status": "pending" | "cancelled", "content": object }`

- `DELETE /api/writing/submissions/:id` - Delete submission
  - **MISSING from AI Service** - Needs to be in Writing Service

### Progress Integration
- `POST /api/writing/submissions/:id/score` - Trigger scoring (proxy to AI Service)
  - Redirects to `/api/ai/score` and stores result in database

---

## 2. Database Schema

### submissions Table (NEW)
- `id` (UUID, PK) - Unique submission identifier
- `user_id` (UUID, FK → users.id) - Link to Identity Service (JWT verified)
- `assessment_id` (UUID, FK → assessments.id) - Link to Learning Service
- `content` (JSONB) - Raw user submission (task type, paragraphs, word count)
- `status` (text, DEFAULT 'pending') - 'pending' | 'scoring' | 'scored' | 'cancelled'
- `band_score` (number) - Final band score (nullable until scored)
- `score_breakdown` (JSONB) - Per-criterion scores (task_achievement, coherence, lexical, grammar)
- `feedback` (JSONB) - AI feedback (strengths, improvements, suggestions)
- `confidence` (number, 0-1) - AI confidence in evaluation
- `processing_time_ms` (integer) - Time taken for AI scoring
- `created_at` (timestamp, DEFAULT now())
- `updated_at` (timestamp, DEFAULT now())

**Dependencies**: Users table from Identity Service

### feedbacks Table (NEW - Optional for MVP)
- `id` (UUID, PK)
- `submission_id` (UUID, FK → submissions.id)
- `criterion` (text) - 'task_achievement' | 'coherence' | 'lexical_resource' | 'grammatical_range'
- `score` (number)
- `comment` (text)
- `suggestion` (text)
- `created_at` (timestamp, DEFAULT now())

**Rationale**: Could embed feedback in submissions.score_breakdown instead of separate table (less JOINs for MVP)

---

## 3. Dependencies

### Core Stack (Already defined in Authority)
- ✅ Next.js 14+ (App Router)
- ✅ Drizzle ORM (type-safe database queries)
- ✅ Neon Postgres (cloud PostgreSQL)
- ✅ iron-session (serverless-safe session management)
- ✅ Z.ai API (GLM-4.5-Flash for scoring)

### Required New Dependencies
- ✅ Resend (magic-link auth - already in Identity Service)
- ❌ **MISSING**: Media storage (PDF upload support)
  - Gap: Writing submissions may contain task prompts (PDF, images)
  - Needs WebSearch to research alternatives (Cloudflare R2, S3-compatible, etc.)

### Optional Dependencies (Phase 2+)
- 🔄 Prompt evaluation prompt templates library
- 🔄 Band score calibration database

---

## 4. Implementation Sequence (Top-Down)

### Phase 1: Foundation (3-5 days)
**Priority Order**:
1. **Database Schema** (1 day)
   - Create `submissions` table in Neon
   - Define Drizzle schema with TypeScript types
   - Add indexes on `user_id`, `assessment_id`, `status`

2. **API Endpoints** (2 days)
   - `POST /api/writing/submissions` (create submission)
   - `GET /api/writing/submissions/:id` (retrieve - **MISSING**)
   - `PUT /api/writing/submissions/:id` (update - **MISSING**)
   - `DELETE /api/writing/submissions/:id` (delete - **MISSING**)

3. **AI Integration** (1-2 days)
   - Integrate Z.ai GLM-4.5-Flash API for scoring
   - Implement `/api/ai/score` proxy endpoint
   - Store results in `submissions` table
   - Handle errors and retries

**Dependencies**: ✅ Identity Service (auth), ✅ Band Score Utilities (already built)

### Phase 2: Progress Integration (2-3 days)
1. Update `submissions` table to include `score_breakdown` and `feedback`
2. Create `/api/progress/writing` endpoint
3. Sync AI scores to Progress Service

### Phase 3: Media Support (2-3 days)
1. Research and choose media storage solution
2. Add `attachments` table for PDF/task prompts
3. Implement file upload/download endpoints

### Phase 4: Testing & Optimization (2-3 days)
1. Unit tests for all endpoints
2. Integration tests with Z.ai API
3. Error handling and logging
4. Performance optimization

**Total Estimated Time**: 9-14 days for MVP

---

## 5. Open Decisions (To be captured in ADRs)

### Gap-001: API Endpoint Retrieval
**Description**: `GET /api/writing/submissions/:id`, `PUT /api/writing/submissions/:id`, `DELETE /api/writing/submissions/:id` endpoints are missing from AI Service but needed for Writing Service.

**Impact**: Writing Service cannot retrieve, update, or delete submissions - incomplete CRUD operations.

**Options**:
- A. Add endpoints to Writing Service (recommended - clean separation of concerns)
- B. Add to AI Service (unclear ownership, mixing concerns)

**Recommended**: Option A - Add to Writing Service

### Gap-002: Media Storage Solution
**Description**: Writing tasks may require PDF/task prompt uploads, but Media Service is in Blocked state (depends on AI Service).

**Impact**: Can't implement full submission support (task prompts missing).

**Options**:
- A. Add Media Service (parallel work, 2-3 days) - complete MVP
- B. Skip PDF support in MVP, use text-only submissions - minimal MVP
- C. Use embedded base64 (limited to small files) - quick but limited

**Recommended**: Option B - skip PDF in MVP for minimal viable feature set

### Gap-003: Database Transaction Strategy
**Description**: Submission creation and AI scoring need atomic transaction to prevent orphaned submissions.

**Impact**: Race conditions between submission creation and scoring.

**Options**:
- A. Neon serverless transactions (recommended)
- B. Manual lock mechanisms
- C. Upsert with optimistic concurrency

**Recommended**: Option A - Neon serverless transactions (built-in)

---

## 6. Next Steps

1. ✅ **Await founder approval** on:
   - API endpoint definitions
   - Database schema
   - Implementation sequence
   - Open decisions (Gap-001, Gap-002, Gap-003)

2. Once approved, create ADRs for:
   - Gap-001: Writing Service API Structure
   - Gap-003: Database Transaction Strategy

3. Begin Phase 1 implementation (Database Schema first)

---

## 7. Validation Checklist

- ✅ No code written yet (pure analysis)
- ✅ No new dependencies added without Authority
- ✅ All endpoints scoped to Writing MVP only
- ✅ Database tables limited to `submissions` and optional `feedbacks`
- ✅ Dependencies match Authority definitions
- ✅ Implementation sequence respects dependencies (Identity → Band Score → AI → Progress → Writing)

---

## 8. Questions for Founder

1. **Media Storage**: Approve Option B (skip PDF support in MVP) or should we prioritize Media Service parallel work?

2. **API Ownership**: Confirm endpoints should be in Writing Service (not AI Service)?

3. **Feedback Storage**: Use embedded `feedback` JSONB in submissions table or separate `feedbacks` table? (Recommend: embedded for MVP)

4. **Task Type Support**: Support only Task 2 essays initially, or Task 1 + Task 2? (Recommend: Task 2 only for MVP)

5. **Retry Logic**: For Z.ai API failures, retry once with exponential backoff or immediately return error? (Recommend: retry once, log failures)

---

**Approval Required**: Before writing any code, get founder confirmation on these decisions.
