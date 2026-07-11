# Media Service

## Service Definition

The Media Service handles all media-related functionality including file uploads, processing, storage, and delivery across the IELTS platform.

### Core Responsibilities
- File upload and management
- Audio/video processing
- Image optimization
- Media storage
- Content delivery
- Format conversion
- Media analytics
- Copyright management

### API Endpoints
```typescript
// Core Media Operations
POST /api/media/upload               // Upload media file
GET  /api/media/download/:id        // Download media file
DELETE /api/media/:id               // Delete media file
GET  /api/media/info/:id            // Get media information
PUT  /api/media/update/:id          // Update media metadata

// Media Processing
POST /api/media/process             // Process media file
GET  /api/media/queue              // Get processing queue
POST /api/media/convert            // Convert media format
GET  /api/media/status/:id         // Get processing status
POST /api/media/optimize           // Optimize media file

// Audio Processing
POST /api/media/audio/transcribe    // Transcribe audio
GET  /api/media/audio/analyze      // Analyze audio
POST /api/media/audio/compress     // Compress audio
GET  /api/media/audio/quality       // Get audio quality info
POST /api/media/audio/segment      // Segment audio

// Video Processing
POST /api/media/video/compress      // Compress video
GET  /api/media/video/analyze      // Analyze video
POST /api/media/video/trim         // Trim video
POST /api/media/video/cut          // Cut video
GET  /api/media/video/quality      // Get video quality info

// Image Processing
POST /api/media/image/resize       // Resize image
POST /api/media/image/optimize     // Optimize image
POST /api/media/image/crop         // Crop image
GET  /api/media/image/info          // Get image info
POST /api/media/image/convert       // Convert image format

// Media Storage
GET  /api/media/storage             // Get storage info
POST /api/media/storage/cleanup     // Cleanup old files
GET  /api/media/storage/usage      // Get storage usage
POST /api/media/storage/backup     // Create backup
GET  /api/media/storage/stats      // Storage statistics

// Content Delivery
GET  /api/media/cdn/url            // Get CDN URL
POST /api/media/cdn/purge          // Purge CDN cache
GET  /api/media/cdn/stats          // CDN statistics
POST /api/media/deliver           // Deliver content
GET  /api/media/stream/:id         // Stream media

// Media Management
GET  /api/media/list               // List all media
GET  /api/media/search             // Search media
POST /api/media/batch              // Batch operations
GET  /api/media/categories         // Get media categories
POST /api/media/category          // Create category
PUT  /api/media/category/:id       // Update category

// Analytics
GET  /api/media/analytics         // Media analytics
GET  /api/media/usage            // Usage statistics
GET  /api/media/performance       // Performance metrics
GET  /api/media/errors            // Error logs
GET  /api/media/reports           // Media reports

// System
GET  /api/media/health            // Health check
GET  /api/media/metrics           // Performance metrics
POST /api/media/maintenance       // Run maintenance
GET  /api/media/costs            // Cost analysis
```

### Data Models
```typescript
interface MediaFile {
  id: string
  name: string
  original_name: string
  type: 'audio' | 'video' | 'image' | 'document' | 'other'
  mime_type: string
  size: number
  duration?: number
  dimensions?: {
    width: number
    height: number
  }
  format: string
  quality: number
  url: string
  cdn_url: string
  storage_path: string
  access_level: 'public' | 'private' | 'restricted'
  metadata: Record<string, any>
  uploaded_by: string
  uploaded_at: Date
  processed_at?: Date
  status: 'uploading' | 'processing' | 'ready' | 'failed' | 'deleted'
  tags: string[]
  category_id?: string
}

interface MediaProcessing {
  id: string
  file_id: string
  type: 'transcription' | 'compression' | 'conversion' | 'optimization' | 'analysis'
  status: 'pending' | 'processing' | 'completed' | 'failed'
  progress: number
  result?: any
  error?: string
  started_at: Date
  completed_at?: Date
  queue_position: number
  priority: 'low' | 'normal' | 'high'
}

interface AudioAnalysis {
  id: string
  file_id: string
  analysis: {
    speech_rate: number
    clarity: number
    pronunciation: number
    fillers: number
    pauses: Array<{
      start: number
      end: number
      duration: number
    }>
    transcription: string
    language_detected: string
    confidence: number
  }
  created_at: Date
}

interface VideoAnalysis {
  id: string
  file_id: string
  analysis: {
    resolution: string
    frame_rate: number
    bitrate: number
    duration: number
    codec: string
    quality_score: number
    scenes: Array<{
      start: number
      end: number
      description?: string
    }>
  }
  created_at: Date
}

interface ImageAnalysis {
  id: string
  file_id: string
  analysis: {
    dimensions: {
      width: number
      height: number
    }
    format: string
    size: number
    color_profile: string
    resolution_score: number
    contains_text: boolean
    faces_detected: number
    objects_detected: string[]
  }
  created_at: Date
}

interface MediaStorage {
  total_space: number
  used_space: number
  available_space: number
  file_count: number
  by_type: Record<string, number>
  by_category: Record<string, number>
  daily_growth: number
  cost_analysis: {
    storage_cost: number
    processing_cost: number
    delivery_cost: number
    total_cost: number
  }
}

interface MediaDelivery {
  id: string
  file_id: string
  delivery_type: 'stream' | 'download' | 'cdn'
  client_ip: string
  user_agent: string
  location: string
  timestamp: Date
  bandwidth_used: number
  delivery_time: number
  success: boolean
  error?: string
}

interface MediaAnalytics {
  usage: {
    total_uploads: number
    total_downloads: number
    total_streams: number
    total_bandwidth: number
  }
  popular_files: Array<{
    file_id: string
    download_count: number
    stream_count: number
    bandwidth_used: number
  }>
  performance: {
    average_upload_time: number
    average_processing_time: number
    average_download_speed: number
    error_rate: number
  }
  trends: {
    daily_uploads: number[]
    weekly_uploads: number[]
    monthly_uploads: number[]
    popular_formats: Record<string, number>
    popular_categories: Record<string, number>
  }
}

interface MediaCategory {
  id: string
  name: string
  description: string
  allowed_types: string[]
  max_size: number
  is_active: boolean
  created_at: Date
  file_count: number
  total_size: number
}

interface MediaQueue {
  id: string
  type: 'processing' | 'upload' | 'download'
  status: 'pending' | 'active' | 'completed' | 'failed'
  items: Array<{
    id: string
    file_id?: string
    priority: number
    progress: number
    started_at?: Date
    completed_at?: Date
    error?: string
  }>
  total_items: number
  processing_items: number
  estimated_completion: Date
}

interface MediaMetrics {
  performance: {
    upload_throughput: number
    processing_speed: number
    download_throughput: number
    storage_efficiency: number
  }
  reliability: {
    success_rate: number
    error_rate: number
    uptime_percentage: number
  }
  cost: {
    cost_per_upload: number
    cost_per_download: number
    cost_per_storage_gb: number
    total_daily_cost: number
  }
}
```

### Observable Metrics
- Upload success rate
- Processing completion rate
- CDN performance
- Storage usage
- Bandwidth consumption
- Media quality scores
- Error rates
- User satisfaction

### Integration Points
- File Service: File operations
- CDN Service: Content delivery
- AI Service: Media analysis
- Analytics Service: Usage tracking
- Database: Media metadata
- Cloud Storage: File storage

### Error Handling
```typescript
const ERRORS = {
  UPLOAD_FAILED: new Error('Media upload failed'),
  PROCESSING_ERROR: new Error('Media processing failed'),
  STORAGE_ERROR: new Error('Storage error'),
  CONVERSION_ERROR: new Error('Format conversion failed'),
  CDN_ERROR: new Error('CDN delivery failed'),
  FILE_NOT_FOUND: new Error('Media file not found'),
  INVALID_FORMAT: new Error('Invalid media format'),
  QUOTA_EXCEEDED: new Error('Storage quota exceeded'),
} as const
```

### Performance Requirements
- Upload < 10s for files < 100MB
- Processing < 5min for audio/video
- CDN delivery < 1s
- Image optimization < 2s
- Queue processing < 500ms per item
- 99.9% uptime for media endpoints
- Handle 1000 concurrent uploads