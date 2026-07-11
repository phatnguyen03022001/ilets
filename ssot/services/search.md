# Search Service

## Service Definition

The Search Service provides comprehensive search capabilities across all platform content, user data, and learning resources.

### Core Responsibilities
- Full-text search
- Content indexing
- Filter and sort
- Search analytics
- Content discovery
- Performance optimization
- Relevance ranking

### API Endpoints
```typescript
// Core Search
POST /api/search                     // General search
GET  /api/search/suggest            // Search suggestions
GET  /api/search/autocomplete        // Autocomplete
POST /api/search/advanced           // Advanced search
GET  /api/search/filters           // Available filters
GET  /api/search/popular           // Popular searches

// Content Search
GET  /api/search/content            // Search content
GET  /api/search/writing            // Writing content
GET  /api/search/speaking           // Speaking content
GET  /api/search/reading            // Reading content
GET  /api/search/listening          // Listening content
GET  /api/search/vocabulary        // Vocabulary content
GET  /api/search/grammar            // Grammar content

// User Data Search
GET  /api/search/user/:id           // User-specific search
GET  /api/search/history            // Search history
POST /api/search/save               // Save search
GET  /api/search/bookmarks          // Search bookmarks
POST /api/search/favorite           // Favorite search

// Analytics
GET  /api/search/analytics          // Search analytics
GET  /api/search/popular-queries    // Popular queries
GET  /api/search/trending           // Trending searches
GET  /api/search/recommendations    // Search recommendations
POST /api/search/feedback           // Search feedback

// System
GET  /api/search/status             // Search service status
GET  /api/search/health             // Health check
POST /api/search/index              // Index content
GET  /api/search/metrics            // Performance metrics
```

### Data Models
```typescript
interface SearchResult {
  id: string
  type: 'content' | 'user' | 'resource' | 'question'
  title: string
  description: string
  content: string
  url: string
  relevance_score: number
  metadata: Record<string, any>
  highlights: string[]
  created_at: Date
  updated_at: Date
}

interface SearchQuery {
  id: string
  user_id: string
  query: string
  filters: SearchFilter[]
  sort: SearchSort
  limit: number
  offset: number
  results_count: number
  execution_time: number
  created_at: Date
}

interface SearchFilter {
  field: string
  operator: 'equals' | 'contains' | 'range' | 'exists' | 'in'
  value: any | any[]
  type: 'string' | 'number' | 'date' | 'boolean' | 'array'
}

interface SearchSort {
  field: string
  order: 'asc' | 'desc'
}

interface SearchSuggestion {
  text: string
  category: string
  confidence: number
  popular: boolean
  related: string[]
}

interface SearchAnalytics {
  total_searches: number
  average_response_time: number
  popular_queries: Array<{
    query: string
    count: number
    conversion_rate: number
  }>
  search_patterns: {
    peak_hours: string[]
    average_session_duration: number
    unique_queries: number
    repeat_queries: number
  }
  content_performance: {
    click_rates: Record<string, number>
    bounce_rates: Record<string, number>
    avg_time_on_page: Record<string, number>
  }
}

interface SearchHistory {
  id: string
  user_id: string
  query: string
  results_count: number
  filters: SearchFilter[]
  clicked_results: string[]
  created_at: Date
}

interface SearchBookmark {
  id: string
  user_id: string
  query: string
  filters: SearchFilter[]
  name: string
  description?: string
  created_at: Date
  last_used?: Date
  usage_count: number
}

interface SearchIndex {
  name: string
  status: 'active' | 'building' | 'failed' | 'rebuilding'
  document_count: number
  last_updated: Date
  size_mb: number
  health_score: number
}

interface SearchMetrics {
  performance: {
    average_response_time: number
    p95_response_time: number
    qps: number
    error_rate: number
  }
  relevance: {
    click_through_rate: number
    satisfaction_score: number
    improved_results: number
  }
  usage: {
    daily_searches: number
    unique_users: number
    average_queries_per_user: number
  }
}
```

### Observable Metrics
- Search response time
- Relevance accuracy
- Click-through rate
- Search success rate
- Popular query trends
- Index size and performance
- User satisfaction

### Integration Points
- Database: Content indexing
- Analytics Service: Search analytics
- Learning Service: Content discovery
- User Service: Search history
- File Service: Content storage

### Error Handling
```typescript
const ERRORS = {
  INVALID_QUERY: new Error('Invalid search query'),
  INDEX_ERROR: new Error('Search index error'),
  NO_RESULTS: new Error('No results found'),
  FILTER_ERROR: new Error('Invalid filter configuration'),
  TIMEOUT: new Error('Search timeout'),
  SERVICE_UNAVAILABLE: new Error('Search service unavailable'),
} as const
```

### Performance Requirements
- Simple search < 100ms
- Complex search < 500ms
- Autocomplete < 50ms
- Index rebuild < 5min
- Search response < 300ms
- 99.9% uptime for search endpoints
- Handle 1000 QPS peak load