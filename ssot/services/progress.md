# Progress Service

## Service Definition

The Progress Service tracks, analyzes, and reports user learning progress across all IELTS skills and platform activities.

### Core Responsibilities
- Multi-skill progress tracking
- Band score calculation
- Performance analytics
- Goal setting and tracking
- Milestone management
- Progress reporting
- Achievement system

### API Endpoints
```typescript
// Progress Overview
GET  /api/progress/overview           // Get overall progress summary
GET  /api/progress/skills            // Get skill-wise progress
GET  /api/progress/timeline          // Get progress timeline
GET  /api/progress/band-scores       // Get band score history

// Skills Progress
GET  /api/progress/writing           // Writing progress
GET  /api/progress/speaking         // Speaking progress
GET  /api/progress/reading           // Reading progress
GET  /api/progress/listening         // Listening progress
GET  /api/progress/vocabulary       // Vocabulary progress
GET  /api/progress/grammar           // Grammar progress

// Goals & Milestones
GET  /api/progress/goals             // User goals
POST /api/progress/goals             // Create goal
PUT  /api/progress/goals/:id         // Update goal
GET  /api/progress/milestones        // Milestones
POST /api/progress/milestones        // Create milestone
GET  /api/progress/goals/progress    // Goal progress

// Analytics & Insights
GET  /api/progress/analytics         // Progress analytics
GET  /api/progress/insights          // Learning insights
GET  /api/progress/predictions       // Progress predictions
GET  /api/progress/recommendations   // Recommendations
GET  /api/progress/weaknesses       // Identify weaknesses

// Reports
GET  /api/progress/report            // Generate progress report
GET  /api/progress/report/:id       // Get specific report
POST /api/progress/share             // Share progress report
GET  /api/progress/export            // Export progress data
GET  /api/progress/certificates      // Get achievement certificates

// Streaks & Achievements
GET  /api/progress/streaks          // Get streaks
POST /api/progress/streaks/reset     // Reset streak
GET  /api/progress/achievements      // Get achievements
GET  /api/progress/leaderboard       // Get leaderboard
```

### Data Models
```typescript
interface UserProgress {
  userId: string
  overall_band_score: number
  skill_progress: {
    writing: SkillProgress
    speaking: SkillProgress
    reading: SkillProgress
    listening: SkillProgress
    vocabulary: SkillProgress
    grammar: SkillProgress
  }
  goals: Goal[]
  milestones: Milestone[]
  achievements: Achievement[]
  streaks: Streak[]
  total_time_spent: number
  total_sessions: number
  started_at: Date
  last_updated: Date
}

interface SkillProgress {
  current_band_score: number
  target_band_score: number
  improvement_rate: number
  sessions_completed: number
  average_accuracy: number
  time_spent: number
  last_activity: Date
  learning_curve: number[]
  weak_areas: string[]
  strong_areas: string[]
  band_history: BandScore[]
}

interface BandScore {
  id: string
  date: Date
  band_score: number
  test_type: 'practice' | 'mock' | 'assessment'
  score_breakdown: {
    writing?: number
    speaking?: number
    reading?: number
    listening?: number
  }
  comments?: string
}

interface Goal {
  id: string
  title: string
  description: string
  target_band_score: number
  target_date: Date
  status: 'active' | 'completed' | 'paused' | 'failed'
  progress: number
  milestones: string[]
  created_at: Date
  updated_at: Date
}

interface Milestone {
  id: string
  title: string
  description: string
  band_score: number
  requirements: {
    skill: string
    band_score: number
    time_spent: number
    sessions: number
  }
  is_achieved: boolean
  achieved_at?: Date
  created_at: Date
}

interface Achievement {
  id: string
  name: string
  description: string
  icon: string
  requirements: AchievementRequirement[]
  earned_at?: Date
  rarity: 'common' | 'rare' | 'epic' | 'legendary'
  category: 'skill' | 'streak' | 'progress' | 'consistency'
}

interface AchievementRequirement {
  type: string
  value: number
  description: string
}

interface Streak {
  id: string
  type: 'daily' | 'weekly' | 'monthly'
  current: number
  max: number
  last_active: Date
  is_active: boolean
  history: Date[]
}

interface ProgressAnalytics {
  userId: string
  overall_trend: {
    direction: 'improving' | 'stable' | 'declining'
    rate: number
    timeframe: number
  }
  skill_comparison: {
    strongest: string
    weakest: string
    average_gap: number
  }
  learning_patterns: {
    peak_times: string[]
    session_frequency: number
    average_session_length: number
  }
  predictions: {
    estimated_band_score: number
    time_to_target: number
    confidence_level: number
  }
  recommendations: string[]
}

interface ProgressReport {
  id: string
  userId: string
  period: {
    start: Date
    end: Date
  }
  summary: {
    overall_band_score: number
    improvement: number
    skills_improved: string[]
    skills_needing_attention: string[]
  }
  detailed_analysis: {
    writing: AnalysisSection
    speaking: AnalysisSection
    reading: AnalysisSection
    listening: AnalysisSection
    vocabulary: AnalysisSection
    grammar: AnalysisSection
  }
  recommendations: Recommendation[]
  generated_at: Date
}

interface AnalysisSection {
  current_band_score: number
  target_band_score: number
  progress: number
  strengths: string[]
  weaknesses: string[]
  recommendations: string[]
}

interface Recommendation {
  priority: 'high' | 'medium' | 'low'
  action: string
  expected_improvement: number
  timeframe: string
}
```

### Observable Metrics
- Overall band score improvement
- Individual skill progression
- Goal completion rate
- Achievement unlock rate
- Streak retention
- Progress report generation
- Recommendation acceptance

### Integration Points
- Learning Service: Activity tracking
- All skill services: Progress data
- Analytics Service: Pattern analysis
- Notification Service: Achievement alerts
- File Service: Report generation

### Error Handling
```typescript
const ERRORS = {
  PROGRESS_NOT_FOUND: new Error('Progress data not found'),
  INVALID_DATA: new Error('Invalid progress data'),
  CALCULATION_ERROR: new Error('Progress calculation failed'),
  GOAL_ERROR: new Error('Goal operation failed'),
  REPORT_GENERATION_ERROR: new Error('Failed to generate report'),
  ANALYSIS_ERROR: new Error('Progress analysis failed'),
} as const
```

### Performance Requirements
- Progress overview < 100ms
- Skill progress < 50ms per skill
- Report generation < 2s
- Analytics calculation < 3s
- Goal updates < 200ms
- 99.9% uptime for progress endpoints