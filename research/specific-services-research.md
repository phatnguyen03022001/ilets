# Specific Service Research & Best Practices

## Admin Service

### Research Sources
- [Admin Dashboard Best Practices](https://www.atlassian.com/blog/admin-dashboard-best-practices)
- [Admin UX Guidelines](https://www.nngroup.com/articles/admin-interfaces/)

### Key Findings

#### Monitoring Needs
- **System Health**: CPU, memory, disk, network
- **Performance Metrics**: Response times, throughput, error rates
- **Business Metrics**: User count, revenue, engagement
- **Security Metrics**: Failed logins, suspicious activity, threats

#### Recommended Libraries
- **Health Checks**:
  - `next-api-health` (custom implementation)
  - or `@goodix/health-check`
- **Metrics**:
  - `prom-client` for custom metrics
  - `nextjs-pwa` for PWA support

#### Common Pitfalls
- ❌ Over-engineered dashboards (too much data at once)
- ❌ Missing context (no comparison to baseline)
- ❌ Slow updates (minutes behind real-time)
- ❌ No drill-down capabilities

#### Best Practices
- ✅ Use real-time data (WebSocket or polling)
- ✅ Provide drill-down to specific details
- ✅ Set up alerts for critical thresholds
- ✅ Mobile-responsive dashboard
- ✅ Export functionality for reports

## Analytics Service

### Research Sources
- [Analytics Best Practices](https://www.amplitude.com/blog/analytics-best-practices)
- [Data Visualization](https://www.tableau.com/blog/data-viz-guide)

### Key Findings

#### Core Analytics Needed
1. **User Analytics**:
   - DAU/MAU metrics
   - Retention curves
   - Cohort analysis
   - User behavior flow

2. **Performance Analytics**:
   - Skill-wise progress
   - Band score improvement
   - Time-on-task analysis
   - Feature adoption

3. **Business Analytics**:
   - Revenue tracking
   - Conversion funnel
   - Churn analysis
   - LTV/CAC

#### Data Storage
- **Primary**: PostgreSQL (Drizzle ORM)
- **Analytics Database**: TimescaleDB (time-series optimization)
- **Alternative**: ClickHouse (if >10M events/day)

#### Recommendations
- ✅ Aggregate data nightly (not real-time)
- ✅ Use materialized views for fast queries
- ✅ Provide pre-built dashboards
- ✅ Custom analytics via SQL API

## Identity Service

### Research Sources
- [Auth Best Practices](https://www.okta.com/identity-101/articles/what-is-auth0/)
- [JWT Security](https://tools.ietf.org/html/rfc7519)

### Key Findings

#### Authentication Flow
1. **Login**: Email + password → JWT (access + refresh)
2. **Session Management**:
   - Access token: 15 minutes
   - Refresh token: 7 days
   - Store in httpOnly cookie (secure)
3. **Multi-Factor**:
   - TOTP (Google Authenticator)
   - SMS backup codes

#### Security Measures
- ✅ Password hashing: Argon2id (not bcrypt)
- ✅ Rate limiting: 5 attempts per 15 minutes
- ✅ IP tracking for suspicious activity
- ✅ Email verification required
- ✅ Account lockout after 5 failed attempts

#### Recommended Libraries
- **Auth**: `next-auth` (Auth.js v5) - community standard
- **Password Hash**: `argon2` package
- **Rate Limiting**: `@upstash/ratelimit` (Redis-based)

#### Common Pitfalls
- ❌ Storing JWTs in localStorage (XSS risk)
- ❌ No refresh token rotation
- ❌ Missing CSRF protection
- ❌ No rate limiting on login

## AI Service

### Research Sources
- [LLM Evaluation](https://promptfoo.dev/)
- [RAG Best Practices](https://www.pinecone.io/learn/rag/)

### Key Findings

#### Evaluation Pipeline
1. **Input Processing**:
   - Token counting
   - Format validation
   - Context extraction

2. **AI Evaluation**:
   - Model: GPT-4o or Claude 3.5 Sonnet
   - Prompt engineering with rubric
   - Structured output (JSON)

3. **Scoring**:
   - Per-criterion scoring
   - Band score calculation
   - Confidence scoring

4. **Feedback Generation**:
   - Identify strengths
   - Highlight improvements
   - Provide examples

#### RAG Implementation
- **Vector Database**: pgvector (PostgreSQL extension)
- **Embedding Model**: OpenAI text-embedding-3-large
- **Storage**: Band descriptors, rubrics, examples

#### Quality Assurance
- ✅ Human-in-the-loop for initial training
- ✅ Regular calibration with official samples
- ✅ A/B testing different prompts
- ✅ Track metric distribution (not just average)

## Grammar Service

### Research Sources
- [Grammar Checkers](https://www.grammarly.com/blog/grammar-checking/)
- [ESL Learning](https://www.esl-library.com/)

### Key Findings

#### Grammar Rules Database
- **Structure**:
  - Category (tense, clause, etc.)
  - Examples (correct + incorrect)
  - Common errors
  - Band relevance (1-9)
  - Difficulty (1-9)

#### Error Detection
- **Approach**: Rule-based + AI hybrid
  - Rule-based for common patterns (spacing, punctuation)
  - AI for complex errors (coherence, phrasing)

#### Practice Generation
- **Question Types**:
  - Fill-in-the-blank
  - Error identification
  - Rewrite sentences
  - Multiple choice

#### Recommended Tools
- **Grammar Rules**: Custom JSON database
- **Error Detection**: `gramm-y` or `typosaurus`
- **Practice Generation**: Template-based + AI enrichment

## Learning Service

### Research Sources
- [Learning Path Design](https://www.atlassian.com/blog/learning-path-design)
- [Course Structure](https://www.interactivecloud.com/blog/e-learning-course-design)

### Key Findings

#### Course Structure
1. **Learning Path**:
   - Pre-requisite tracking
   - Estimated duration
   - Target band score
   - Progress tracking

2. **Lesson Types**:
   - Video (with captions)
   - Text content
   - Interactive exercises
   - Audio files

3. **Progression**:
   - Automatic unlocking based on completion
   - Time-based hints
   - Mastery checks

#### Content Delivery
- ✅ Adaptive difficulty
- ✅ Spaced repetition
- ✅ Progress bars
- ✅ Time estimates

#### Recommendation Strategy
- **Content-based**: Based on weak areas
- **Goal-based**: Align with target band score
- **Learning-style**: Based on user preferences

## Listening Service

### Research Sources
- [Listening Practice](https://www.cambridgeenglish.org/exams-and-qualifications/english-for-educators/)
- [Audio Processing](https://ffmpeg.org/)

### Key Findings

#### Audio Management
- **Formats**: MP3, WAV, M4A (auto-convert)
- **Processing**:
  - Transcription (for evaluation)
  - Segmentation (for replay)
  - Speed adjustment
  - Audio enhancement

#### Practice Features
- **Multiple accents**: British, Australian, American, Canadian
- **Playback controls**: Play/pause, rewind 30s, speed (0.8x-1.5x)
- **Transcript availability**: For reference
- **Answer submission**: Multiple choice, completion, matching

#### Evaluation
- **Speech-to-Text**: OpenAI Whisper v3-large
- **Answer Matching**: Exact + fuzzy matching
- **Band Score**: Calculate based on accuracy

#### Common Pitfalls
- ❌ Low audio quality reduces accuracy
- ❌ No transcript for reference
- ❌ Too long sessions (>15 minutes)
- ❌ Multiple accents confusing

## Media Service

### Research Sources
- [Media Storage](https://aws.amazon.com/s3/)
- [CDN Best Practices](https://developer.mozilla.org/en-US/docs/Web/Performance/Using_CDN)

### Key Findings

#### File Storage
- **Primary**: S3-compatible storage (Neon supports S3)
- **CDN**: Cloudflare CDN (fastest for Vietnam region)
- **Formats**:
  - Audio: MP3 (192kbps), WAV (for editing)
  - Video: MP4 (H.264)
  - Images: WebP (90% smaller than PNG/JPG)

#### Processing Pipeline
1. **Upload**: Stream processing (don't buffer full file)
2. **Optimization**: Auto-convert to optimized formats
3. **Thumbnail**: Generate preview images
4. **CDN Purge**: Clear cache after updates
5. **Analytics**: Track downloads/streaming

#### Storage Requirements
- **Estimates**:
  - Audio (1 min): ~2MB
  - Video (1 min): ~20MB
  - Image: ~100KB
- **Quotas**: 100MB/user for uploads

#### Cost Optimization
- ✅ Use CDN for delivery (reduces origin hits)
- ✅ Auto-delete old files (30 days retention)
- ✅ Compression before storage
- ✅ Multi-part upload for large files

## Notification Service

### Research Sources
- [Push Notifications](https://firebase.google.com/docs/messaging)
- [Email Deliverability](https://www.sendgrid.com/resources/education/deliverability-guide)

### Key Findings

#### Notification Channels
1. **Push**: Firebase Cloud Messaging (FCM)
2. **Email**: Resend (reliable, cost-effective)
3. **SMS**: Twilio (for critical alerts)
4. **In-app**: Real-time web notifications

#### Templates
- **Categories**:
  - Achievement (badges, milestones)
  - Reminder (practice, deadline)
  - Feedback (evaluation results)
  - Marketing (promotions)

#### Delivery Optimization
- ✅ Batch sending (10,000/minute max)
- ✅ Background queue (not blocking requests)
- ✅ Retry failed attempts (3 attempts)
- ✅ Deduplication (prevent spam)

#### User Preferences
- ✅ Channel selection (email/push/SMS)
- ✅ Frequency (immediate/daily/weekly)
- ✅ Quiet hours (no notifications after 9pm)
- ✅ Categories opt-out

## Payment Service

### Research Sources
- [Payment Best Practices](https://stripe.com/docs/payments)
- [Fraud Detection](https://www.klarna.com/us/blog/payment-fraud-detection/)

### Key Findings

#### Payment Gateways
- **Primary**: Stripe (widely supported, great API)
- **Fallback**: PayPal (for international users)
- **Local**: MoMo (Vietnam, if needed)

#### Transaction Flow
1. **Order**: Create order in system
2. **Payment**: Redirect to gateway
3. **Callback**: Verify transaction
4. **Webhook**: Handle webhook notifications
5. **Refund**: Support refunds within 30 days

#### Subscription Management
- **Plans**:
  - Free tier (limited practice)
  - Pro (unlimited practice, advanced analytics)
  - Premium (all features + personalized coaching)

#### Security
- ✅ PCI DSS compliance (Stripe handles this)
- ✅ 3D Secure for credit cards
- ✅ Fraud detection (Stripe Radar)
- ✅ Sarbanes-Oxley compliant

#### Tax & Compliance
- ✅ VAT handling (if EU/UK users)
- ✅ Revenue recognition
- ✅ Invoice generation
- ✅ Financial reporting

## Progress Service

### Research Sources
- [Progress Tracking](https://www.atlassian.com/blog/productivity/progress-tracking)
- [Goal Setting](https://www.mindtools.com/pages/article/newHLE_07.htm)

### Key Findings

#### Progress Tracking
1. **Overall Progress**:
   - Current band score
   - Target band score
   - Improvement rate
   - Sessions completed

2. **Skill-wise Progress**:
   - Writing: Band scores, TA/CC/LR/GRA
   - Speaking: Fluency, pronunciation, coherence
   - Reading: Accuracy, time management
   - Listening: Accuracy, accent recognition

3. **Goals & Milestones**:
   - Short-term (weekly targets)
   - Long-term (band score goals)
   - Achievement badges

#### Band Score History
- **Storage**: Time-series database (TimescaleDB)
- **Visualization**: Line chart of scores over time
- **Breakdown**: Per-criterion trends

#### Recommendations
- ✅ Focus on weak areas
- ✅ Provide study plan
- ✅ Suggest practice resources
- ✅ Adjust difficulty based on progress

## Reading Service

### Research Sources
- [Reading Comprehension](https://www.readingrockets.org/articles/reading-comprehension)
- [Assessment Design](https://www.ets.org/toefl/reading-comprehension/)

### Key Findings

#### Passage Management
- **Authenticity**: Use official Cambridge IELTS materials
- **Difficulty**: Band-aligned (1-9)
- **Topics**: Academic, general training, mixed
- **Word Count**: 500-700 per passage

#### Question Types
- **True/False/Not Given**: Common, but tricky
- **Matching**: Headings, sentences, features
- **Multiple Choice**: Single or multiple correct
- **Completion**: Gap filling

#### Evaluation
- **Answer Validation**:
  - Exact match for completion
  - Partial match for multiple choice
  - Semantic match for T/F/NG
- **Time Tracking**: Monitor time management
- **Band Score**: Calculate based on accuracy

#### Common Pitfalls
- ❌ Poor passage quality
- ❌ Inconsistent question types
- ❌ No answer key verification
- ❌ Time limit not enforced

## Search Service

### Research Sources
- [Search Best Practices](https://www.elastic.co/blog/full-text-search-best-practices)
- [Search Relevance](https://www.vividsolutions.com/blog/search-engine-improvements)

### Key Findings

#### Search Engine
- **Primary**: Elasticsearch (full-text search)
- **Alternatives**:
  - Meilisearch (lightweight)
  - Algolia (managed, expensive)
  - Typesense (open-source)

#### Features
1. **Full-Text Search**:
   - Keyword matching
   - Fuzzy search (typos)
   - Phrase search

2. **Filters**:
   - Difficulty level
   - Skill type
   - Band score
   - Topic

3. **Sorting**:
   - Relevance score
   - Popularity
   - Date

#### Search Optimization
- ✅ Pre-index content (don't index on demand)
- ✅ Boost keywords (band descriptors, key terms)
- ✅ Caching (Redis)
- ✅ Analytics (popular queries)

#### Autocomplete
- **Sources**:
  - Common queries
  - Popular topics
  - Band descriptors
- **Delay**: 300ms for smooth UX

## Speaking Service

### Research Sources
- [Speaking Evaluation](https://www.ielts.org/for-organisations/ielts-scoring-in-detail)
- [Speech Analysis](https://www.naturalreaders.com/blog/voice-analysis/)

### Key Findings

#### Recording Features
- **Audio Quality**: 128kbps minimum
- **Duration**: Per question limit (2-3 minutes)
- **Recording**: Real-time recording + upload
- **Feedback**: Immediate transcript generation

#### Evaluation Criteria
1. **Fluency & Coherence**:
   - Speaking rate (150-160 wpm ideal)
   - Hesitation detection
   - Fillers usage (avoid excessive "um", "ah")
   - Paragraph structure

2. **Lexical Resource**:
   - Vocabulary variety
   - Collocations
   - Idiom usage (appropriate context)

3. **Grammatical Range**:
   - Sentence variety
   - Tense consistency
   - Grammar accuracy

4. **Pronunciation**:
   - Phoneme accuracy
   - Stress and intonation
   - Clarity

#### Mock Tests
- **Structure**: 4 parts (simulated)
- **Timing**: 11-14 minutes total
- **Evaluation**: Band score calculation
- **Feedback**: Detailed breakdown + examples

#### Common Pitfalls
- ❌ Microphone quality issues
- ❌ Background noise
- ❌ Time limit not enforced
- ❌ Cue card topics not clear

## Vocabulary Service

### Research Sources
- [Vocabulary Learning](https://www.vocabulario.com/vocabulary-learning-techniques)
- [Spaced Repetition](https://www.supermemo.com/en/blog/spaced-repetition-systems)

### Key Findings

#### Spaced Repetition System (SRS)
- **Algorithm**: SM-2 (SuperMemo 2)
- **Review Schedule**:
  - 1 day after first learning
  - 4 days later
  - 7 days later
  - 14 days later
  - 30 days later
- **Intervals**: Double every correct response

#### Word Management
- **Structure**:
  - Word (headword)
  - Phonetic (IPA)
  - Definition
  - Examples
  - Synonyms/antonyms
  - Collocations
  - Band relevance (1-9)
  - Frequency (1-100)

#### Practice Modes
1. **Flashcards**: Flip card, type answer
2. **Quiz**: Multiple choice, fill-in-the-blank
3. **Context**: See word in sentences
4. **Review**: Scheduled SRS reviews

#### Word Difficulty
- **Levels**: Beginner (1-4), Intermediate (5-7), Advanced (8-9)
- **Band Target**: Align with user's IELTS goal
- **Adaptive**: Adjust based on performance

#### Recommended Libraries
- **SRS Algorithm**: `spaced-repetition-2` or custom implementation
- **Word Database**: JSON file with 2000-3000 high-frequency words
- **Practice Generation**: Template-based

## Integration Patterns Summary

### Service-to-Service Communication
- **Synchronous**: REST API calls (direct, simple)
- **Asynchronous**: Message queue (background jobs)
- **Shared Database**: Federation for cross-service queries

### Data Flow
1. **User Actions**: Frontend → Service
2. **Business Logic**: Service → AI Service (evaluation)
3. **Background Tasks**: Queue → Worker (processing)
4. **Analytics**: Service → Analytics Service (aggregation)

### Error Handling
- **Retry**: Exponential backoff (3 attempts)
- **Fallback**: Return cached or default values
- **Alert**: Log error, notify admins
- **User Feedback**: Explain what went wrong

## Performance Considerations

### Database
- **Optimization**: Add indexes on frequently queried fields
- **Caching**: Redis for hot data (user progress, settings)
- **Query**: Limit result sets (pagination)

### API
- **Rate Limiting**: 100 requests/minute per user
- **Timeout**: 30 seconds (fallback after timeout)
- **Caching**: CDN caching for static responses

### Frontend
- **Optimistic UI**: Show changes immediately, rollback on error
- **Skeleton Loading**: Placeholder while loading
- **Code Splitting**: Lazy load heavy components
- **Debouncing**: Delay search input (300ms)

## Security Considerations

### Authentication
- JWT tokens (access + refresh)
- Session management (httpOnly cookies)
- CSRF protection
- Rate limiting

### Authorization
- Role-based access (student, teacher, admin)
- Resource ownership checks
- API key rotation

### Data Protection
- Encryption at rest (PostgreSQL SSL)
- Encryption in transit (TLS 1.3)
- PII handling (GDPR compliance)
- Data retention policies

## References
- [Next.js Security](https://nextjs.org/docs/security)
- [Drizzle ORM Security](https://orm.drizzle.team/docs/security)
- [Stripe Security](https://stripe.com/docs/security)
- [Firebase Cloud Messaging](https://firebase.google.com/docs/messaging)
- [Resend Email API](https://resend.com/docs/api-reference)
