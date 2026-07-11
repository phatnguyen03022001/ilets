# Learning Service

## Service Definition

The Learning Service manages educational content, courses, and learning resources for the IELTS platform.

### Core Responsibilities
- Course catalog management
- Learning content organization
- Progress tracking
- Content delivery
- Learning paths
- Resource recommendations

### API Endpoints
```typescript
// Course Management
GET  /api/courses             // List all courses
GET  /api/courses/:id         // Get specific course
POST /api/courses             // Create new course
PUT  /api/courses/:id         // Update course
DELETE /api/courses/:id       // Delete course

// Learning Content
GET  /api/lessons/:courseId   // Get course lessons
GET  /api/lessons/:id         // Get specific lesson
POST /api/lessons/completion  // Mark lesson complete
GET  /api/lessons/progress   // Get user progress

// Learning Paths
GET  /api/paths               // List learning paths
GET  /api/paths/:id           // Get specific path
POST /api/paths/:enroll       // Enroll in path
GET  /api/paths/:progress     // Get path progress

// Recommendations
GET  /api/recommendations    // Get personalized recommendations
GET  /api/recommendations/content // Content-based recommendations
```

### Data Models
```typescript
interface Course {
  id: string
  title: string
  description: string
  level: 'beginner' | 'intermediate' | 'advanced'
  duration: number // minutes
  lessons: Lesson[]
  tags: string[]
  difficulty: 1-5
  publishedAt: Date
  updatedAt: Date
}

interface Lesson {
  id: string
  courseId: string
  title: string
  content: {
    type: 'text' | 'video' | 'audio' | 'interactive'
    url?: string
    duration?: number
    transcript?: string
  }
  order: number
  prerequisites: string[]
  estimatedTime: number
}

interface LearningPath {
  id: string
  name: string
  description: string
  courses: string[]
  estimatedDuration: number
  targetBandScore: number
  isEnrolled: boolean
  progress: number
}

interface Progress {
  userId: string
  courseId: string
  lessonId: string
  completedAt: Date
  score?: number
  timeSpent: number
  attempts: number
}
```

### Observable Metrics
- Course enrollment rate
- Lesson completion rate
- Average time per lesson
- Learning path completion
- Content engagement score
- Recommendation acceptance rate

### Integration Points
- Database: Course and progress data
- AI Service: Content recommendations
- Analytics Service: Learning behavior tracking
- File Service: Media content delivery

### Error Handling
```typescript
const ERRORS = {
  COURSE_NOT_FOUND: new Error('Course not found'),
  LESSON_LOCKED: new Error('Prerequisites not met'),
  ALREADY_COMPLETED: new Error('Lesson already completed'),
  INVALID_PROGRESS: new Error('Invalid progress data'),
  RECOMMENDATION_ERROR: new Error('Failed to generate recommendations'),
} as const
```

### Performance Requirements
- Course list response < 300ms
- Lesson content delivery < 100ms
- Progress update < 50ms
- Recommendation generation < 1s
- 99.95% uptime for content endpoints