# Vocabulary Service

## Service Definition

The Vocabulary Service manages vocabulary building, tracking, and assessment for IELTS preparation.

### Core Responsibilities
- Word collection management
- Vocabulary tracking
- Assessment generation
- Spaced repetition
- Word usage examples
- Progress monitoring
- Difficulty adjustment

### API Endpoints
```typescript
// Word Management
GET  /api/vocabulary/words           // List words
GET  /api/vocabulary/words/:id       // Get word details
POST /api/vocabulary/words           // Add new word
PUT  /api/vocabulary/words/:id       // Update word
DELETE /api/vocabulary/words/:id     // Delete word
GET  /api/vocabulary/search          // Search words

// Vocabulary Lists
GET  /api/vocabulary/lists           // Get vocabulary lists
POST /api/vocabulary/lists           // Create vocabulary list
GET  /api/vocabulary/lists/:id       // Get list details
PUT  /api/vocabulary/lists/:id       // Update list
DELETE /api/vocabulary/lists/:id     // Delete list

// Practice & Learning
POST /api/vocabulary/practice        // Start vocabulary practice
GET  /api/vocabulary/practice/:id    // Get practice session
POST /api/vocabulary/check-answer    // Check vocabulary answer
GET  /api/vocabulary/review          // Get review words
POST /api/vocabulary/mastery         // Mark word as mastered

// Spaced Repetition
GET  /api/vocabulary/schedule        // Get review schedule
POST /api/vocabulary/review          // Perform spaced repetition
GET  /api/vocabulary/due-count       // Get count of due words
GET  /api/vocabulary/progress        // SRS progress

// Assessment
POST /api/vocabulary/quiz           // Start vocabulary quiz
GET  /api/vocabulary/quiz/:id        // Get quiz results
GET  /api/vocabulary/test            // Full vocabulary test
GET  /api/vocabulary/stats           // Vocabulary statistics

// Progress & Analytics
GET  /api/vocabulary/progress        // Overall progress
GET  /api/vocabulary/analytics       // Detailed analytics
GET  /api/vocabulary/trends          // Learning trends
GET  /api/vocabulary/weaknesses     // Identify weaknesses
GET  /api/vocabulary/recommendations // Word recommendations
```

### Data Models
```typescript
interface VocabularyWord {
  id: string
  word: string
  phonetic: string
  definition: string
  part_of_speech: 'noun' | 'verb' | 'adjective' | 'adverb' | 'preposition' | 'conjunction' | 'pronoun' | 'interjection'
  difficulty: number // 1-9
  frequency: number // 1-100
  band_relevance: number // 1-9
  examples: string[]
  synonyms: string[]
  antonyms: string[]
  collocations: string[]
  etymology?: string
  added_by: string
  added_at: Date
}

interface VocabularyList {
  id: string
  name: string
  description: string
  word_count: number
  difficulty: number
  band_target: number
  topic: string
  is_public: boolean
  created_by: string
  created_at: Date
  words: string[] // word IDs
}

interface VocabularySession {
  id: string
  userId: string
  type: 'practice' | 'quiz' | 'review'
  words: string[]
  started_at: Date
  completed_at?: Date
  results: VocabularyResult[]
  accuracy: number
  average_time: number
}

interface VocabularyResult {
  word_id: string
  question_type: 'definition' | 'synonym' | 'usage' | 'spelling'
  user_answer: string
  correct_answer: string
  is_correct: boolean
  time_spent: number
  attempts: number
  mastery_level: number // 0-5
}

interface SpacedRepetition {
  id: string
  word_id: string
  user_id: string
  level: number // 0-5
  interval_days: number
  next_review: Date
  reviews: Review[]
  created_at: Date
}

interface Review {
  id: string
  spaced_id: string
  is_correct: boolean
  time_spent: number
  confidence: number
  reviewed_at: Date
  interval_before: number
}

interface VocabularyProgress {
  userId: string
  total_words_learned: number
  words_mastered: number
  current_level: number
  accuracy_rate: number
  average_review_time: number
  retention_rate: number
  streak_days: number
  words_by_band: Record<string, number>
  learning_trend: number[]
  weak_areas: string[]
  strong_areas: string[]
}

interface VocabularyAnalytics {
  userId: string
  learning_patterns: {
    peak_hours: number[]
    average_session_length: number
    words_per_session: number
    retention_rate: number
  }
  performance_metrics: {
    accuracy_by_difficulty: Record<string, number>
    time_per_word: number
    improvement_rate: number
    mastery_progress: number
  }
  recommendations: {
    next_words: string[]
    review_schedule: Date[]
    focus_areas: string[]
  }
}
```

### Observable Metrics
- Words learned per day
- Retention rate
- Mastery progression
- Quiz accuracy
- Review completion rate
- Learning session duration
- Vocabulary growth trend

### Integration Points
- Database: Word and progress data
- AI Service: Word generation and examples
- Learning Service: Progress tracking
- Analytics Service: Learning pattern analysis

### Error Handling
```typescript
const ERRORS = {
  WORD_NOT_FOUND: new Error('Word not found'),
  INVALID_ANSWER: new Error('Invalid answer format'),
  SCHEDULE_ERROR: new Error('Review schedule error'),
  MASTERY_ERROR: new Error('Failed to update mastery'),
  VOCABULARY_LIMIT: new Error('Vocabulary limit reached'),
  INVALID_INPUT: new Error('Invalid vocabulary input'),
} as const
```

### Performance Requirements
- Word lookup < 50ms
- Quiz generation < 1s
- Review scheduling < 200ms
- Progress calculation < 100ms
- Search response < 300ms
- 99.9% uptime for vocabulary endpoints