# AI Service

## Service Definition

The AI Service provides artificial intelligence capabilities across the IELTS platform including evaluation, feedback, personalization, and content generation.

### Core Responsibilities
- AI evaluation and feedback
- Content generation
- Personalized recommendations
- Speech recognition
- Natural language processing
- Automated scoring
- Quality assurance

### API Endpoints
```typescript
// AI Evaluation
POST /api/ai/evaluate                 // Evaluate any response
GET  /api/ai/evaluate/:id           // Get evaluation results
POST /api/ai/score                   // Band score calculation
GET  /api/ai/feedback/:id           // Get feedback
POST /api/ai/improve                 // Improve content
GET  /api/ai/rubric/:skill          // Get scoring rubric

// Content Generation
POST /api/ai/generate                // Generate content
GET  /api/ai/templates              // Get templates
POST /api/ai/custom                 // Custom content
GET  /api/ai/examples               // Get examples
POST /api/ai/validate               // Validate content

// Personalization
GET  /api/ai/profile                // Get user profile
POST /api/ai/profile                // Update profile
GET  /api/ai/recommendations        // Get recommendations
POST /api/ai/adjust                 // Adjust difficulty
GET  /api/ai/preferences            // Get preferences

// Speech Processing
POST /api/ai/speech/transcribe       // Transcribe speech
GET  /api/ai/speech/analyze        // Analyze speech
POST /api/ai/speech/recognize       // Recognize speech
GET  /api/ai/speech/feedback       // Get speech feedback

// Quality Control
POST /api/ai/quality/check          // Check quality
GET  /api/ai/quality/metrics        // Get quality metrics
POST /api/ai/quality/improve        // Improve quality
GET  /api/ai/consistency            // Get consistency score

// Integration Points
POST /api/ai/integrate              // Integrate with service
GET  /api/ai/capabilities           // Get capabilities
GET  /api/ai/status                  // AI service status
POST /api/ai/logs                   // Log interactions
```

### Data Models
```typescript
interface AIEvaluation {
  id: string
  userId: string
  skill: 'writing' | 'speaking' | 'reading' | 'listening'
  input: string
  evaluation: {
    band_score: number
    score_breakdown: {
      task_achievement: number
      coherence_cohesion: number
      lexical_resource: number
      grammatical_range: number
      pronunciation?: number
      fluency?: number
    }
    feedback: {
      strengths: string[]
      improvements: string[]
      suggestions: string[]
      examples: string[]
    }
    confidence: number // 0-1
    processing_time: number
  }
  timestamp: Date
}

interface AIContent {
  id: string
  type: 'essay' | 'speech' | 'reading_passage' | 'questions' | 'vocabulary'
  prompt: string
  generated: any
  metadata: {
    word_count: number
    difficulty: number
    time_estimate: number
    band_target: number
  }
  quality_score: number
  reviewed: boolean
  created_at: Date
}

interface AIPersona {
  id: string
  userId: string
  preferences: {
    learning_style: 'visual' | 'auditory' | 'kinesthetic' | 'mixed'
    difficulty_preference: number
    pace_preference: number
    feedback_style: 'detailed' | 'concise' | 'encouraging'
    focus_areas: string[]
  }
  history: {
    skill_usage: Record<string, number>
    accuracy_rates: number[]
    improvement_trend: number
  }
  last_updated: Date
}

interface AISpeechAnalysis {
  id: string
  userId: string
  audio_url: string
  transcript: string
  analysis: {
    fluency: number
    pronunciation: number
    vocabulary: number
    grammar: number
    coherence: number
    fillers: number
    pace: number
  }
  feedback: {
    pronunciation_issues: string[]
    fluency_suggestions: string[]
    vocabulary_recommendations: string[]
    overall_assessment: string
  }
  timestamp: Date
}

interface AIQualityMetrics {
  id: string
  content_type: string
  metrics: {
    accuracy: number
    relevance: number
    coherence: number
    completeness: number
    creativity: number
  }
    score: number
    issues: QualityIssue[]
    recommendations: string[]
  }
  
  interface QualityIssue {
    severity: 'low' | 'medium' | 'high' | 'critical'
    category: string
    description: string
    suggestion: string
  }

interface AIRecommendation {
  id: string
  userId: string
  type: 'skill' | 'resource' | 'practice' | 'goal'
  content: string
  priority: number
  confidence: number
  impact: number
  created_at: Date
  expires_at?: Date
}
```

### Observable Metrics
- AI evaluation accuracy
- Response time for evaluations
- Content generation quality
- Recommendation acceptance rate
- Speech recognition accuracy
- Quality score trends
- User satisfaction ratings

### Integration Points
- All skill services: Evaluation support
- Learning Service: Content generation
- Analytics Service: AI usage patterns
- File Service: Media processing
- Database: Result storage

### Error Handling
```typescript
const ERRORS = {
  EVALUATION_FAILED: new Error('AI evaluation failed'),
  CONTENT_INVALID: new Error('Invalid content for processing'),
  SPEECH_RECOGNITION_ERROR: new Error('Speech recognition failed'),
  QUALITY_CHECK_FAILED: new Error('Quality check failed'),
  PERSONALIZATION_ERROR: new Error('Personalization failed'),
  SERVICE_UNAVAILABLE: new Error('AI service unavailable'),
} as const
```

### Performance Requirements
- Writing evaluation < 3s
- Speech recognition < 2s
- Content generation < 5s
- Quality check < 1s
- Recommendation generation < 2s
- 99.9% uptime for AI endpoints
- Average response time < 2s