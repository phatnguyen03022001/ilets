# Notification Service

## Service Definition

The Notification Service manages all push notifications, email communications, and alerts across the IELTS platform.

### Core Responsibilities
- Push notifications
- Email delivery
- In-app alerts
- System notifications
- User preferences
- Notification templates
- Delivery tracking
- Analytics

### API Endpoints
```typescript
// Core Notifications
POST /api/notifications/send         // Send notification
GET  /api/notifications/user/:id    // Get user notifications
GET  /api/notifications/unread      // Get unread notifications
PUT  /api/notifications/read/:id    // Mark as read
DELETE /api/notifications/:id      // Delete notification
POST /api/notifications/mark-all-read // Mark all as read

// Notification Types
POST /api/notifications/push       // Send push notification
POST /api/notifications/email      // Send email
POST /api/notifications/in-app     // Send in-app notification
POST /api/notifications/sms        // Send SMS
POST /api/notifications/webhook     // Send webhook

// Templates
GET  /api/notifications/templates  // Get templates
POST /api/notifications/templates  // Create template
PUT  /api/notifications/templates/:id // Update template
DELETE /api/notifications/templates/:id // Delete template
GET  /api/notifications/templates/preview // Preview template

// User Preferences
GET  /api/notifications/preferences/:id // Get preferences
PUT  /api/notifications/preferences/:id // Update preferences
POST /api/notifications/preferences/:id/subscribe // Subscribe
POST /api/notifications/preferences/:id/unsubscribe // Unsubscribe
GET  /api/notifications/preferences/categories // Get categories

// Delivery
GET  /api/notifications/delivery   // Get delivery status
POST /api/notifications/delivery/retry // Retry delivery
GET  /api/notifications/delivery/history // Delivery history
POST /api/notifications/delivery/test // Test delivery

// Analytics
GET  /api/notifications/analytics  // Notification analytics
GET  /api/notifications/metrics    // Performance metrics
GET  /api/notifications/campaigns   // Campaign analytics
GET  /api/notifications/ab-testing  // A/B test results

// Scheduling
POST /api/notifications/schedule   // Schedule notification
GET  /api/notifications/schedule    // Get scheduled notifications
PUT  /api/notifications/schedule/:id // Update schedule
DELETE /api/notifications/schedule/:id // Cancel schedule

// Bulk Operations
POST /api/notifications/batch     // Send batch notifications
GET  /api/notifications/batch/status/:id // Batch status
POST /api/notifications/campaign   // Create campaign
GET  /api/notifications/campaigns  // List campaigns
PUT  /api/notifications/campaigns/:id // Update campaign

// System
GET  /api/notifications/health     // Health check
GET  /api/notifications/metrics     // System metrics
POST /api/notifications/cleanup     // Clean up old notifications
GET  /api/notifications/limits     // Rate limits
```

### Data Models
```typescript
interface Notification {
  id: string
  user_id: string
  type: 'push' | 'email' | 'in-app' | 'sms' | 'webhook'
  category: string
  title: string
  message: string
  data: Record<string, any>
  priority: 'low' | 'normal' | 'high' | 'urgent'
  status: 'pending' | 'sending' | 'delivered' | 'failed' | 'acknowledged'
  delivery_method: string
  scheduled_at?: Date
  sent_at?: Date
  delivered_at?: Date
  read_at?: Date
  acknowledged_at?: Date
  expires_at?: Date
  retry_count: number
  error?: string
  template_id?: string
  campaign_id?: string
  metadata: Record<string, any>
}

interface NotificationPreference {
  user_id: string
  email_enabled: boolean
  push_enabled: boolean
  sms_enabled: boolean
  categories: Array<{
    category: string
    enabled: boolean
    channels: ('email' | 'push' | 'sms')[]
    frequency: 'immediate' | 'daily' | 'weekly' | 'never'
  }>
  quiet_hours: {
    enabled: boolean
    start_time: string
    end_time: string
    days: number[]
  }
  timezone: string
}

interface NotificationTemplate {
  id: string
  name: string
  category: string
  type: 'push' | 'email' | 'in-app'
  subject?: string
  body: string
  variables: string[]
  preview: string
  is_active: boolean
  created_by: string
  created_at: Date
  updated_at: Date
  usage_count: number
  last_used?: Date
}

interface NotificationCampaign {
  id: string
  name: string
  description: string
  target_audience: {
    user_ids: string[]
    segments: string[]
    conditions: Record<string, any>
  }
  template_id?: string
  content: {
    title: string
    message: string
    data: Record<string, any>
  }
  schedule: {
    type: 'immediate' | 'scheduled' | 'recurring'
    start_at?: Date
    end_at?: Date
    frequency?: string
    timezone: string
  }
  channels: ('email' | 'push' | 'sms')[]
  status: 'draft' | 'scheduled' | 'running' | 'completed' | 'cancelled'
  metrics: {
    sent: number
    delivered: number
    opened: number
    clicked: number
    bounced: number
    unsubscribed: number
  }
  created_by: string
  created_at: Date
  started_at?: Date
  completed_at?: Date
}

interface DeliveryStatus {
  id: string
  notification_id: string
  user_id: string
  channel: string
  status: 'pending' | 'sending' | 'delivered' | 'failed' | 'retrying'
  attempt: number
  max_attempts: number
  error?: string
  response_time: number
  timestamp: Date
  metadata: Record<string, any>
}

interface NotificationAnalytics {
  overview: {
    total_sent: number
    total_delivered: number
    total_failed: number
    delivery_rate: number
    open_rate: number
    click_rate: number
    bounce_rate: number
    unsubscribe_rate: number
  }
  by_channel: Record<string, {
    sent: number
    delivered: number
    delivered_rate: number
    open_rate: number
    click_rate: number
    average_response_time: number
  }>
  by_category: Record<string, {
    sent: number
    delivered: number
    engagement_rate: number
    conversion_rate: number
  }>
  trends: {
    daily: Array<{
      date: Date
      sent: number
      delivered: number
      opened: number
      clicked: number
    }>
    hourly: Array<{
      hour: string
      sent: number
      delivered: number
      engagement_rate: number
    }>
  }
  user_behavior: {
    time_to_open: number
    time_to_click: number
    repeat_engagement: number
    preference_changes: number
  }
}

interface NotificationMetrics {
  performance: {
    delivery_success_rate: number
    average_delivery_time: number
    queue_processing_time: number
    error_rate: number
    retry_rate: number
  }
  engagement: {
    open_rate: number
    click_rate: number
    conversion_rate: number
    user_satisfaction: number
  }
  volume: {
    notifications_sent: number
    notifications_delivered: number
    active_users: number
    notifications_per_user: number
  }
  cost: {
    cost_per_notification: number
    cost_per_delivery: number
    total_monthly_cost: number
    roi: number
  }
}

interface ScheduledNotification {
  id: string
  user_id: string
  notification_id: string
  scheduled_at: Date
  status: 'pending' | 'processing' | 'delivered' | 'failed' | 'cancelled'
  processed_at?: Date
  error?: string
  metadata: Record<string, any>
}

interface NotificationSegment {
  id: string
  name: string
  description: string
  conditions: Array<{
    field: string
    operator: 'equals' | 'contains' | 'greater_than' | 'less_than' | 'in'
    value: any | any[]
  }>
  user_count: number
  created_at: Date
  updated_at: Date
}

interface ATestResult {
  id: string
  campaign_id: string
  name: string
  variants: Array<{
    id: string
    name: string
    content: string
    sent: number
    delivered: number
    opened: number
    clicked: number
    conversion_rate: number
    winner: boolean
  }>
  conclusion: string
  confidence_level: number
  started_at: Date
  completed_at: Date
}
```

### Observable Metrics
- Delivery success rate
- Open/click rates
- User engagement
- Notification preferences
- Channel performance
- Response times
- Error rates
- User satisfaction

### Integration Points
- User Service: User preferences
- Email Service: Email delivery
- Push Service: Push notifications
- SMS Service: SMS delivery
- Analytics Service: Notification analytics
- Database: Notification storage

### Error Handling
```typescript
const ERRORS = {
  DELIVERY_FAILED: new Error('Notification delivery failed'),
  INVALID_RECIPIENT: new Error('Invalid notification recipient'),
  TEMPLATE_NOT_FOUND: new Error('Notification template not found'),
  PREFERENCE_ERROR: new Error('Notification preference error'),
  RATE_LIMIT_EXCEEDED: new Error('Notification rate limit exceeded'),
  SCHEDULING_ERROR: new Error('Notification scheduling error'),
  CAMPAIGN_ERROR: new Error('Campaign operation failed'),
} as const
```

### Performance Requirements
- Notification delivery < 1s
- Queue processing < 500ms
- Template rendering < 100ms
- Bulk send < 5s for 1000 notifications
- Delivery tracking < 200ms
- 99.9% uptime for notification endpoints
- Handle 1000 QPS peak load