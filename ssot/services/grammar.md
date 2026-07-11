# Grammar Service

## Service Definition

The Grammar Service provides IELTS grammar practice, assessment, and error correction for grammar improvement.

### Core Responsibilities
- Grammar rule management
- Error detection and correction
- Practice exercises
- Progress tracking
- Feedback generation
- Band score alignment
- Custom grammar rules

### API Endpoints
```typescript
// Grammar Rules
GET  /api/grammar/rules              // List grammar rules
GET  /api/grammar/rules/:id         // Get specific rule
POST /api/grammar/rules             // Add new rule
PUT  /api/grammar/rules/:id         // Update rule
GET  /api/grammar/rules/search      // Search rules

// Grammar Practice
POST /api/grammar/practice          // Start grammar practice
GET  /api/grammar/practice/:id      // Get practice session
POST /api/grammar/check             // Check grammar
GET  /api/grammar/exercises         // Get exercises
GET  /api/grammar/feedback          // Get feedback

// Error Detection
POST /api/grammar/detect            // Detect grammar errors
GET  /api/grammar/corrections/:id   // Get corrections
POST /api/grammar/improve           // Improve grammar
GET  /api/grammar/examples/:id      // Get examples
GET  /api/grammar/explanations/:id  // Get explanations

// Assessment
POST /api/grammar/test              // Take grammar test
GET  /api/grammar/test/:id          // Get test results
GET  /api/grammar/evaluate          // Evaluate writing
GET  /api/grammar/score             // Get grammar score
GET  /api/grammar/band              // Get band descriptor

// Progress & Analytics
GET  /api/grammar/progress          // Grammar progress overview
GET  /api/grammar/history            // Grammar history
GET  /api/grammar/analytics         // Detailed analytics
GET  /api/grammar/weaknesses        // Identify weaknesses
GET  /api/grammar/recommendations  // Grammar recommendations
```

### Data Models
```typescript
interface GrammarRule {
  id: string
  name: string
  category: 'tense' | 'clause' | 'sentence_structure' | 'punctuation' | 'word_order' | 'agreement' | 'modals' | 'conditionals'
  description: string
  examples: string[]
  common_errors: string[]
  band_relevance: number // 1-9
  difficulty: number // 1-9
  frequency: number // 1-100
  created_at: Date
  updated_at: Date
}

interface GrammarExercise {
  id: string
  title: string
  type: 'fill_blank' | 'correct_sentence' | 'identify_error' | 'rewrite' | 'multiple_choice'
  question: string
  options?: string[]
  correct_answer: string | string[]
  explanation: string
  related_rules: string[]
  difficulty: number
  band_score: number
  time_limit?: number
}

interface GrammarSession {
  id: string
  userId: string
  type: 'practice' | 'test' | 'assessment'
  exercises: string[]
  started_at: Date
  completed_at?: Date
  results: GrammarResult[]
  score: number
  band_score: number
  time_spent: number
}

interface GrammarResult {
  id: string
  exercise_id: string
  user_answer: string
  correct_answer: string
  is_correct: boolean
  time_spent: number
  errors_detected: GrammarError[]
  feedback: string
  improvement_suggestions: string[]
}

interface GrammarError {
  id: string
  type: string
  description: string
  position?: number
  suggested_correction: string
  rule_id: string
  severity: 'low' | 'medium' | 'high'
}

interface GrammarProgress {
  userId: string
  total_exercises: number
  completed_exercises: number
  accuracy_rate: number
  average_score: number
  band_scores: {
    writing: number
    speaking: number
  }
  common_errors: Record<string, number>
  improvement_rate: number
  mastered_rules: string[]
  learning_curve: number[]
}

interface GrammarAnalytics {
  userId: string
  error_patterns: {
    by_category: Record<string, number>
    by_severity: Record<string, number>
    improvement_trend: number
    most_common_errors: string[]
  }
  performance_metrics: {
    accuracy_by_rule: Record<string, number>
    time_per_exercise: number
    improvement_rate: number
    mastery_progress: number
  }
  recommendations: {
    focus_areas: string[]
    review_rules: string[]
    upcoming_topics: string[]
  }
}

interface GrammarCorrection {
  original_text: string
  corrected_text: string
  changes: Correction[]
  explanation: string
  band_improvement?: number
  score_before?: number
  score_after?: number
}

interface Correction {
  position: number
  length: number
  original: string
  corrected: string
  rule: string
  explanation: string
}
```

### Observable Metrics
- Grammar exercise completion rate
- Error detection accuracy
- Correction acceptance rate
- Band score improvement
- Rule mastery progression
- Session completion rate
- Error reduction trend

### Integration Points
- AI Service: Grammar checking
- Learning Service: Progress tracking
- Writing Service: Integration for writing feedback
- Analytics Service: Grammar pattern analysis

### Error Handling
```typescript
const ERRORS = {
  RULE_NOT_FOUND: new Error('Grammar rule not found'),
  INVALID_INPUT: new Error('Invalid grammar input'),
  DETECTION_FAILED: new Error('Failed to detect errors'),
  CORRECTION_ERROR: new Error('Grammar correction failed'),
  ASSESSMENT_ERROR: new Error('Grammar assessment failed'),
  EXERCISE_LIMIT: new Error('Exercise limit reached'),
} as const
```

### Performance Requirements
- Grammar checking < 500ms
- Exercise generation < 1s
- Error detection < 200ms
- Correction generation < 300ms
- Score calculation < 50ms
- 99.9% uptime for grammar endpoints