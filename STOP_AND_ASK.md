# STOP_AND_ASK.md

## Phase 4: Implementation Planning Results

**Date**: 2026-07-11
**Duration**: ~20 minutes
**Status**: IN PROGRESS

---

## Phase 3 Completion Summary

✅ **Phase 3 Complete**: Architecture validation completed
- ✅ Service boundaries validated (16 services)
- ⚠️ API overlaps identified (5 overlaps, all acceptable scope differences)
- ⚠️ Hidden dependencies identified (5 key dependencies)
- ✅ Architecture decision alignment checked (3/5 fully aligned)

✅ **Hidden Dependencies Documented**: Added to `ssot/decisions.md` (5 dependencies, 1-2 lines each)
✅ **API Overlaps Documented**: Added to `open-decisions.md` (5 overlaps, LOW priority tags)

---

## Phase 4: Implementation Planning - Writing Wedge E (Writing)

### Objective
Plan implementation sequence for Writing wedge (Wedge E) with clear "do first, do after, what's needed" breakdown.

### Analysis

#### What Makes Writing Wedge E Special?

Writing is the first skill to implement (Phase 1 wedge), foundational for:
- Band score calculation (core utility already exists)
- AI evaluation (primary evaluation pattern)
- Progress tracking (first progress integration)
- User feedback loop (critical for user retention)

---

## Implementation Plan: Writing Wedge E

### 🎯 **What to Do FIRST** (Critical Path)

#### 1. **Identity Service** (Foundation)
**Priority**: P0 - Blocks everything else
**Why**: Need authentication system before users can access any features

**Tasks**:
- [ ] Email verification setup (Resend)
- [ ] Magic-link flow implementation
- [ ] Session management (iron-session)
- [ ] User profile creation
- [ ] Role-based access (student/teacher/admin)

**Dependencies**: None
**Timeline**: 3-5 days
**Risk**: Low (magic-link is well-understood)

---

#### 2. **Band Score Calculation Utilities** (Utility)
**Priority**: P0 - Core calculation logic
**Why**: Band scores are the core metric for IELTS evaluation

**Tasks**:
- [ ] Round score calculation (0.5 precision)
- [ ] Per-criterion scoring → band conversion
- [ ] Handle edge cases (scores < 0, > 9)
- [ ] Comprehensive unit tests

**Dependencies**: Identity Service (for testing with users)
**Timeline**: 1-2 days
**Risk**: Low (already exists, just verify/refine)

---

#### 3. **AI Service - Basic Evaluation** (Core)
**Priority**: P0 - Evaluation logic
**Why**: Writing evaluation is the primary AI use case

**Tasks**:
- [ ] OpenRouter integration (GLM-4.5-Flash)
- [ ] Basic prompt template (Task 1 vs Task 2)
- [ ] Evaluation endpoint (receive writing sample)
- [ ] Band score calculation integration
- [ ] Feedback generation (strengths/improvements)

**Dependencies**: Band Score Utilities, Identity Service
**Timeline**: 4-6 days
**Risk**: Medium (AI quality critical)

**Critical Path**: OpenRouter API → Prompt Engineering → Band Score Integration

---

#### 4. **Progress Service - Writing Integration** (Tracking)
**Priority**: P0 - Progress tracking
**Why**: Track user band scores over time

**Tasks**:
- [ ] Create user_progress table (Drizzle schema)
- [ ] Write evaluation → progress update endpoint
- [ ] Band score history tracking
- [ ] Per-criterion progress (TA/CC/LR/GRA)

**Dependencies**: AI Service (for evaluation data)
**Timeline**: 2-3 days
**Risk**: Low (straightforward CRUD)

---

### 🔧 **What to Do AFTER** (Parallel or Sequential)

#### 5. **Writing Service** (Domain Logic) ⏸️ **BLOCKED**
**Priority**: P1 - Domain-specific writing operations
**Why**: Writing service owns writing-specific operations

**Tasks**:
- [ ] Writing sample CRUD (create, read, update, delete)
- [ ] Cue card management
- [ ] Writing practice session management
- [ ] Upload/download writing samples

**Dependencies**: AI Service (for evaluation), Identity Service (for auth)
**Timeline**: 3-4 days
**Risk**: Low
**Note**: Blocked until AI Service evaluation is stable

---

#### 6. **Media Service - Basic Upload** (Support) ⏸️ **BLOCKED**
**Priority**: P2 - File upload for writing samples
**Why**: Users need to upload writing samples

**Tasks**:
- [ ] S3-compatible storage setup (Neon S3)
- [ ] Upload endpoint (streaming)
- [ ] File validation (word count, format)
- [ ] CDN integration (Cloudflare)

**Dependencies**: None
**Timeline**: 2-3 days
**Risk**: Medium (storage setup)
**Note**: Can be done in parallel if Media Service has basic upload without AI integration

---

#### 7. **Notification Service - Email Notifications** (Support) ⏸️ **BLOCKED**
**Priority**: P2 - Email alerts
**Why**: Send evaluation results to users

**Tasks**:
- [ ] Email templates (evaluation results)
- [ ] Send email endpoint (Resend)
- [ ] Queue system (for async sending)
- [ ] Retry logic (failed attempts)

**Dependencies**: Identity Service (user emails), AI Service (evaluation results)
**Timeline**: 2-3 days
**Risk**: Low
**Note**: Blocked until AI Service is ready

---

### 📦 **What's Needed to Start** (Prerequisites)

#### **Immediate Prerequisites**:
1. ✅ **SSOT Service Definitions**: All Writing-related services defined (AI, Progress, Writing, Media, Notification)
2. ✅ **Band Score Utilities**: Core calculation logic ready
3. ✅ **OpenRouter Access**: API key for GLM-4.5-Flash
4. ✅ **Resend API Key**: For magic-link authentication
5. ✅ **Neon Database**: For Identity and Progress services
6. ✅ **S3 Storage**: For Media Service uploads

#### **Technical Stack**:
- Next.js 16 with App Router
- Drizzle ORM (PostgreSQL)
- OpenRouter API (GLM-4.5-Flash)
- Resend (email)
- iron-session (auth)
- TypeScript strict mode

#### **Data Models** (need database schema):
- **User**: id, email, verified, role
- **AuthToken**: id, userId, token, expiresAt
- **UserProgress**: id, userId, writingScore, writingDate, criteriaScores
- **WritingSample**: id, userId, content, wordCount, status
- **AIEvaluation**: id, writingSampleId, bandScore, criteriaScores, feedback

---

## Critical Path Summary

```
Identity Service (Foundation)
    ↓
Band Score Utilities (Validation)
    ↓
AI Service - Basic Evaluation (Core)
    ↓
Progress Service - Writing Integration (Tracking)
    ↓
[Blocked] Writing Service (Domain Logic)
[Blocked] Media Service - Basic Upload (Support)
[Blocked] Notification Service - Email (Support)
```

**Total Critical Path Time**: ~10-16 days
**Overall Timeline to MVP**: ~4-6 weeks (including parallel work)

---

## STOP_AND_ASK Points

### 🟢 **STOP_AND_ASK #1: Implementation Sequence Confirmed?**

**Question**: Is the sequence "Identity → Band Utilities → AI → Progress → [blocked services]" correct?

**Rationale**:
- Identity is foundational (no users = no features)
- Band utilities are independent and needed for AI evaluation
- AI evaluation is core feature (critical path)
- Progress tracking depends on AI evaluation data
- Writing service and Media service are blocked until AI is stable

**Your choice**: [ ] Confirm sequence  [ ] Modify sequence

---

### 🟡 **STOP_AND_ASK #2: AI Evaluation Quality Target?**

**Question**: What's the acceptable AI evaluation accuracy for Writing wedge?

**Current Target**: 85% accuracy on official IELTS samples

**Options**:
1. **85% accuracy** (baseline for MVP)
2. **90% accuracy** (higher bar)
3. **90-95% accuracy** (thorough calibration)

**Your choice**: [ ] 85% (MVP)  [ ] 90%  [ ] 90-95%

---

### 🟡 **STOP_AND_ASK #3: Parallel Work Allowed?**

**Question**: Can Media Service (uploads) be built in parallel while AI Service is in progress?

**Rationale**:
- Media Service can do basic file upload without AI integration
- Could be blocked later when AI evaluation integration needed
- Increases parallel work efficiency

**Your choice**: [ ] Build in parallel  [ ] Wait for AI Service to finish

---

## Summary

**Writing Wedge E Implementation Plan**:
- ✅ **First**: Identity Service (foundation)
- ✅ **Next**: Band Score Utilities (utility)
- ✅ **Core**: AI Service - Basic Evaluation (evaluation logic)
- ✅ **Tracking**: Progress Service - Writing Integration (progress)
- ⏸️ **Blocked**: Writing Service (domain logic), Media Service (upload), Notification Service (email)

**Total Critical Path**: 10-16 days
**Overall Timeline**: 4-6 weeks (including parallel work)

**Immediate Prerequisites**: OpenRouter key, Resend key, Neon DB, S3 storage

**Status**: ⏸️ AWAITING YOUR DECISION
