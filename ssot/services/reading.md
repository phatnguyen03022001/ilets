# Reading Service

## Service Definition

The Reading Service manages IELTS Reading practice, assessment, and progress tracking for all reading comprehension tasks.

### Core Responsibilities
- Reading passage management
- Question generation and validation
- Automated evaluation
- Progress tracking
- Time management
- Score calculation
- Vocabulary analysis

### API Endpoints
```typescript
// Reading Passages
GET  /api/reading/passages           // List available passages
GET  /api/reading/passages/:id       // Get specific passage
POST /api/reading/passages           // Create new passage
PUT  /api/reading/passages/:id       // Update passage
GET  /api/reading/passages/search    // Search passages

// Reading Practice
POST /api/reading/start              // Start reading session
POST /api/reading/submit/:id         // Submit answers
GET  /api/reading/results/:id       // Get results
GET  /api/reading/timing/:id         // Get timing stats
POST /api/reading/review/:id        // Request manual review

// Questions & Answers
GET  /api/reading/questions/:id     // Get passage questions
POST /api/reading/check             // Check answers
GET  /api/reading/explanations/:id   // Get explanations
GET  /api/reading/vocabulary/:id    // Get vocabulary words

// Assessment
GET  /api/reading/assess            // Start assessment
GET  /api/reading/assess/:id        // Get assessment details
POST /api/reading/assess/submit/:id  // Submit assessment
GET  /api/reading/band-scores/:id   // Get band score breakdown

// Progress
GET  /api/reading/progress           // Reading progress overview
GET  /api/reading/history            // Reading history
GET  /api/reading/weaknesses         // Identify weaknesses
GET  /api/reading/recommendations   // Get recommendations
```

### Data Models
```typescript
interface ReadingPassage {
  id: string
  title: string
  content: string
  word_count: number
  difficulty: number // 1-9
  band_target: number
  type: 'academic' | 'general' | 'mixed'
  topics: string[]
  author: string
  source?: string
  published_at: Date
  estimated_time: number
}

interface ReadingQuestion {
  id: string
  passage_id: string
  question: string
  type: 'multiple_choice' | 'true_false' | 'matching' | 'completion' | 'short_answer'
  options?: string[]
  correct_answer: string | string[]
  explanation: string
  skills: ['heading' | 'matching' | 'tf_ng' | 'multiple' | 'completion']
  difficulty: number
}

interface ReadingSession {
  id: string
  userId: string
  passageId: string
  startTime: Date
  endTime?: Date
  timeSpent: number
  answers: Answer[]
  score: number
  band_score: number
  status: 'incomplete' | 'completed' | 'reviewed'
}

interface Answer {
  id: string
  questionId: string
  user_answer: string | string[]
  is_correct: boolean
  time_spent: number
  confidence?: number
}

interface ReadingProgress {
  userId: string
  total_passages: number
  completed_passages: number
  average_score: number
  band_scores: {
    academic: number
    general: number
  }
  accuracy: number
  time_management: number
  common_errors: string[]
  skills_improved: string[]
}

interface ReadingAssessment {
  id: string
  userId: string
  type: 'diagnostic' | 'progress' | 'practice'
  passages: string[]
  start_time: Date
  end_time?: Date
  results: {
    overall_score: number
    band_score: number
    passage_scores: number[]
    skill_breakdown: {
      heading: number
      matching: number
      tf_ng: number
      multiple: number
      completion: number
    }
  }
}
```

### Observable Metrics
- Reading passage completion rate
- Average time per passage
- Question accuracy rate
- Band score improvement
- Time management efficiency
- Vocabulary retention rate
- Error pattern analysis

### Integration Points
- Database: Passage and question data
- AI Service: Reading analysis
- Learning Service: Progress tracking
- Analytics Service: Reading pattern analysis

### Error Handling
```typescript
const ERRORS = {
  PASSAGE_NOT_FOUND: new Error('Passage not found'),
  INVALID_ANSERS: new Error('Invalid answer format'),
  TIME_EXCEEDED: new Error('Reading time exceeded'),
  ASSESSMENT_INCOMPLETE: new Error('Assessment not completed'),
  SCORE_CALCULATION_ERROR: new Error('Failed to calculate score'),
  VOCABULARY_ERROR: new Error('Vocabulary analysis failed'),
} as const
```

### Performance Requirements
- Passage loading < 200ms
- Answer checking < 50ms
- Score calculation < 100ms
- Assessment generation < 2s
- Search response < 300ms
- 99.95% uptime for reading endpoints