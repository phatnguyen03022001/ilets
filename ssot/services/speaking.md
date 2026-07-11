# Speaking Service

## Service Definition

The Speaking Service provides IELTS Speaking practice, evaluation, and feedback for all parts of the Speaking test.

### Core Responsibilities
- Speaking practice sessions
- AI evaluation of speaking responses
- Fluency and pronunciation analysis
- Cue card practice
- Mock speaking tests
- Progress tracking
- Feedback generation

### API Endpoints
```typescript
// Speaking Practice
POST /api/speaking/practice          // Start speaking practice
POST /api/speaking/practice/:record   // Submit recording
GET  /api/speaking/practice/:id       // Get practice session
GET  /api/speaking/cuecards           // Get cue card questions
POST /api/speaking/cuecards/:practice // Submit cue card response

// Evaluation
POST /api/speaking/evaluate           // Evaluate speaking response
GET  /api/speaking/feedback/:id       // Get evaluation feedback
GET  /api/speaking/scores/:id         // Get speaking scores
GET  /api/speaking/transcript/:id     // Get audio transcript

// Mock Tests
POST /api/speaking/mock/start         // Start mock speaking test
GET  /api/speaking/mock/:id           // Get mock test details
POST /api/speaking/mock/submit/:id    // Submit mock test
GET  /api/speaking/mock/results/:id   // Get mock test results

// Progress
GET  /api/speaking/progress           // Get speaking progress
GET  /api/speaking/analysis/:id       // Get detailed analysis
GET  /api/speaking/comparisons        // Compare with band descriptors
```

### Data Models
```typescript
interface SpeakingSession {
  id: string
  userId: string
  type: 'practice' | 'mock' | 'cuecard'
  part: 1 | 2 | 3
  status: 'started' | 'recording' | 'submitted' | 'evaluated'
  recordingUrl?: string
  transcript?: string
  startTime: Date
  endTime?: Date
  responses: SpeakingResponse[]
}

interface SpeakingResponse {
  id: string
  sessionId: string
  questionId: string
  question: string
  response: string
  audioUrl?: string
  transcript?: string
  duration: number
  evaluations: SpeakingEvaluation[]
}

interface SpeakingEvaluation {
  id: string
  responseId: string
  criteria: {
    fluency: number // 1-9
    coherence: number
    lexicalResource: number
    grammatical_range: number
    pronunciation: number
  }
  band_score: number
  feedback: {
    strengths: string[]
    improvements: string[]
    suggestions: string[]
  }
  audio_analysis: {
    pace: number
    hesitation_count: number
    filler_words: string[]
    pronunciation_score: number
  }
}

interface CueCard {
  id: string
  part: 1 | 2
  topic: string
  questions: string[]
  preparation_time: number
  speaking_time: number
  band_score: number
}

interface MockSpeakingTest {
  id: string
  userId: string
  status: 'incomplete' | 'completed' | 'evaluated'
  parts: {
    part1: SpeakingResponse[]
    part2: SpeakingResponse[]
    part3: SpeakingResponse[]
  }
  overall_score: number
  overall_feedback: string
  completed_at?: Date
}
```

### Observable Metrics
- Speaking session completion rate
- Average response duration
- Fluency score improvement
- Band score progression
- Recording upload success rate
- Evaluation response time
- Cue card completion rate

### Integration Points
- AI Service: Speech recognition and evaluation
- File Service: Audio file storage
- Learning Service: Progress tracking
- Analytics Service: Speaking pattern analysis

### Error Handling
```typescript
const ERRORS = {
  INVALID_AUDIO_FORMAT: new Error('Invalid audio format'),
  EVALUATION_TIMEOUT: new Error('Evaluation timeout'),
  TRANSCRIPT_ERROR: new Error('Failed to generate transcript'),
  CRITERIA_ERROR: new Error('Invalid evaluation criteria'),
  MOCK_TEST_INCOMPLETE: new Error('Mock test not completed'),
  CUECARD_TIME_EXCEEDED: new Error('Preparation time exceeded'),
} as const
```

### Performance Requirements
- Audio upload processing < 10s
- Speech-to-text conversion < 5s
- Evaluation generation < 15s
- Mock test completion < 30min
- 99.9% uptime for evaluation endpoints