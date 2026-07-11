# Listening Service

## Service Definition

The Listening Service provides IELTS Listening practice, assessment, and skill development through audio-based exercises.

### Core Responsibilities
- Audio content management
- Listening comprehension practice
- Automated evaluation
- Progress tracking
- Time management
- Score calculation
- Accent recognition

### API Endpoints
```typescript
// Audio Content
GET  /api/listening/tracks           // List audio tracks
GET  /api/listening/tracks/:id       // Get specific track
POST /api/listening/tracks           // Upload audio track
PUT  /api/listening/tracks/:id       // Update track
GET  /api/listening/transcripts/:id  // Get transcript

// Listening Practice
POST /api/listening/start            // Start listening session
POST /api/listening/submit/:id       // Submit answers
GET  /api/listening/results/:id      // Get results
GET  /api/listening/replay/:id       // Replay audio segment
POST /api/listening/pause/:id        // Pause/resume session

// Questions & Evaluation
GET  /api/listening/questions/:id    // Get questions
POST /api/listening/evaluate         // Evaluate responses
GET  /api/listening/feedback/:id     // Get feedback
GET  /api/listening/transcripts/:id  // Get transcript

// Progress & Analysis
GET  /api/listening/progress        // Listening progress overview
GET  /api/listening/history         // Listening history
GET  /api/listening/analytics       // Detailed analytics
GET  /api/listening/recommendations // Get recommendations
GET  /api/listening/skills          // Skill breakdown

// Accent & Speed
GET  /api/listening/accents         // Available accents
POST /api/listening/adjust-speed     // Adjust playback speed
GET  /api/listening/speed-history    // Speed adjustment history
```

### Data Models
```typescript
interface ListeningTrack {
  id: string
  title: string
  description: string
  audio_url: string
  duration: number
  transcript: string
  difficulty: number // 1-9
  band_target: number
  type: 'conversation' | 'monologue' | 'lecture' | 'news'
  accents: string[]
  speaker_count: number
  topics: string[]
  estimated_time: number
  uploaded_at: Date
}

interface ListeningQuestion {
  id: string
  track_id: string
  question: string
  type: 'multiple_choice' | 'completion' | 'matching' | 'map_labeling' | 'note_completion'
  options?: string[]
  correct_answer: string | string[]
  explanation: string
  transcript_segment?: string
  start_time?: number
  end_time?: number
  skills: ['listening_for_main_ideas' | 'listening_for_detail' | 'listening_for_implicit_meaning']
}

interface ListeningSession {
  id: string
  userId: string
  trackId: string
  startTime: Date
  endTime?: Date
  timeSpent: number
  pauseTime?: number
  current_position: number
  playback_speed: number
  answers: ListeningAnswer[]
  score: number
  band_score: number
  status: 'incomplete' | 'completed' | 'paused'
}

interface ListeningAnswer {
  id: string
  questionId: string
  user_answer: string | string[]
  is_correct: boolean
  time_spent: number
  transcript_segment?: string
  audio_position: number
}

interface ListeningProgress {
  userId: string
  total_tracks: number
  completed_tracks: number
  average_score: number
  band_scores: {
    conversations: number
    monologues: number
    lectures: number
    news: number
  }
  accuracy: number
  time_efficiency: number
  common_errors: string[]
  skills_improved: string[]
  accent_mastery: Record<string, number>
}

interface ListeningAnalytics {
  userId: string
  listening_sessions: number
  average_session_time: number
  score_progression: number[]
  speed_adaptation: number[]
  error_patterns: {
    question_type_error_rate: Record<string, number>
    common_mistakes: string[]
    improvement_trend: number
  }
  accent_progression: Record<string, number[]>
}
```

### Observable Metrics
- Listening session completion rate
- Average time per track
- Question accuracy rate
- Band score improvement
- Playback speed adaptation
- Accent recognition accuracy
- Listening comprehension efficiency

### Integration Points
- File Service: Audio file storage
- AI Service: Speech recognition
- Learning Service: Progress tracking
- Analytics Service: Listening pattern analysis

### Error Handling
```typescript
const ERRORS = {
  TRACK_NOT_FOUND: new Error('Audio track not found'),
  INVALID_AUDIO_FORMAT: new Error('Invalid audio format'),
  EVALUATION_FAILED: new Error('Failed to evaluate answers'),
  TRANSCRIPT_UNAVAILABLE: new Error('Transcript not available'),
  SESSION_EXPIRED: new Error('Listening session expired'),
  INVALID_SPEED: new Error('Invalid playback speed'),
} as const
```

### Performance Requirements
- Audio streaming < 2s
- Answer evaluation < 100ms
- Transcript retrieval < 300ms
- Score calculation < 50ms
- Session management < 200ms
- 99.9% uptime for listening endpoints