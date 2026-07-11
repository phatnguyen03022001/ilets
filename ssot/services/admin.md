# Admin Service

## Service Definition

The Admin Service provides administrative control, monitoring, and management capabilities for the IELTS platform operators.

### Core Responsibilities
- Platform administration
- User management
- System monitoring
- Content moderation
- Performance optimization
- Security management
- Operations analytics
- Admin tools

### API Endpoints
```typescript
// User Management
GET  /api/admin/users                 // List users
GET  /api/admin/users/:id            // Get user details
POST /api/admin/users                // Create user
PUT  /api/admin/users/:id            // Update user
DELETE /api/admin/users/:id          // Delete user
POST /api/admin/users/:id/suspend   // Suspend user
POST /api/admin/users/:id/reactivate // Reactivate user
GET  /api/admin/users/:id/activity   // User activity

// Content Management
GET  /api/admin/content              // List content
GET  /api/admin/content/:id         // Get content
POST /api/admin/content              // Create content
PUT  /api/admin/content/:id         // Update content
DELETE /api/admin/content/:id       // Delete content
GET  /api/admin/content/flagged     // Flagged content
POST /api/admin/content/:id/approve // Approve content
POST /api/admin/content/:id/reject  // Reject content

// System Monitoring
GET  /api/admin/system/health        // System health
GET  /api/admin/system/status        // System status
GET  /api/admin/system/metrics       // System metrics
GET  /api/admin/system/logs         // System logs
POST /api/admin/system/diagnose     // Diagnose issues
GET  /api/admin/system/alerts       // System alerts
POST /api/admin/system/resolve       // Resolve alert

// Performance
GET  /api/admin/performance         // Performance metrics
GET  /api/admin/performance/cpu    // CPU metrics
GET  /api/admin/performance/memory  // Memory metrics
GET  /api/admin/performance/disk   // Disk metrics
GET  /api/admin/performance/network // Network metrics
GET  /api/admin/performance/api     // API performance
GET  /api/admin/performance/database // Database performance

// Analytics
GET  /api/admin/analytics           // Platform analytics
GET  /api/admin/analytics/users     // User analytics
GET  /api/admin/analytics/content  // Content analytics
GET  /api/admin/analytics/revenue  // Revenue analytics
GET  /api/admin/analytics/issues    // Issues analytics
GET  /api/admin/analytics/trends   // Trend analytics

// Security
GET  /api/admin/security            // Security status
GET  /api/admin/security/audits     // Security audits
GET  /api/admin/security/users     // User security
GET  /api/admin/security/activity   // Security activity
GET  /api/admin/security/threats    // Threat detection
POST /api/admin/security/ban        // Ban user
GET  /api/admin/security/logs       // Security logs

// Operations
GET  /api/admin/operations          // Operations dashboard
POST /api/admin/operations/maintenance // Maintenance
GET  /api/admin/operations/schedule // Maintenance schedule
POST /api/admin/operations/backup   // Create backup
GET  /api/admin/operations/backups  // Backup history
POST /api/admin/operations/cleanup // Cleanup old data
GET  /api/admin/operations/reports  // Operation reports

// Configuration
GET  /api/admin/config              // Platform config
PUT  /api/admin/config             // Update config
GET  /api/admin/config/validate    // Validate config
POST /api/admin/config/deploy      // Deploy config
GET  /api/admin/config/history     // Config history
GET  /api/admin/config/audit       // Config audit

// Reporting
GET  /api/admin/reports            // Available reports
GET  /api/admin/reports/daily     // Daily report
GET  /api/admin/reports/weekly    // Weekly report
GET  /api/admin/reports/monthly   // Monthly report
POST /api/admin/reports/custom     // Custom report
GET  /api/admin/reports/export     // Export reports
GET  /api/admin/reports/schedule   // Scheduled reports

// Tools
GET  /api/admin/tools              // Available tools
POST /api/admin/tools/diagnose    // Run diagnostics
POST /api/admin/tools/fix         // Fix common issues
GET  /api/admin/tools/backup      // Backup tools
GET  /api/admin/tools/restore     // Restore tools
POST /api/admin/tools/migrate     // Migration tools
GET  /api/admin/tools/status      // Tool status
```

### Data Models
```typescript
interface AdminUser {
  id: string
  username: string
  email: string
  role: 'admin' | 'super_admin' | 'moderator'
  permissions: string[]
  status: 'active' | 'suspended' | 'banned' | 'pending'
  created_at: Date
  last_login: Date
  metadata: {
    login_count: number
    failed_logins: number
    ip_addresses: string[]
    devices: string[]
  }
}

interface AdminUserActivity {
  id: string
  user_id: string
  action: string
  details: Record<string, any>
  timestamp: Date
  ip_address: string
  user_agent: string
}

interface SystemHealth {
  status: 'healthy' | 'degraded' | 'critical'
  components: Array<{
    name: string
    status: 'healthy' | 'degraded' | 'critical'
    metrics: {
      cpu_usage: number
      memory_usage: number
      disk_usage: number
      response_time: number
      error_rate: number
    }
    last_check: Date
  }>
  overall_score: number
}

interface SystemAlert {
  id: string
  type: 'warning' | 'error' | 'info'
  severity: 'low' | 'medium' | 'high' | 'critical'
  message: string
  component: string
  metrics: Record<string, any>
  timestamp: Date
  resolved: boolean
  resolved_by?: string
  resolved_at?: Date
}

interface SystemMetrics {
  timestamp: Date
  metrics: {
    cpu: {
      usage: number
      cores: number
      load_average: number[]
    }
    memory: {
      total: number
      used: number
      free: number
      swap_usage: number
    }
    disk: {
      total: number
      used: number
      free: number
      io_ops: number
    }
    network: {
      incoming: number
      outgoing: number
      packets_in: number
      packets_out: number
    }
    application: {
      active_connections: number
      response_time: number
      error_rate: number
      throughput: number
    }
  }
}

interface ContentModeration {
  id: string
  content_id: string
  content_type: string
  flagged_by: string
  flagged_at: Date
  reason: string
  severity: 'low' | 'medium' | 'high'
  status: 'pending' | 'reviewed' | 'approved' | 'rejected'
  reviewed_by?: string
  reviewed_at?: Date
  action_taken?: string
  metadata: Record<string, any>
}

interface PerformanceMetrics {
  period: {
    start: Date
    end: Date
  }
  overall_score: number
  by_category: {
    api: {
      response_time_avg: number
      error_rate: number
      throughput: number
    }
    database: {
      query_time_avg: number
      connection_usage: number
      index_efficiency: number
    }
    cache: {
      hit_rate: number
      eviction_rate: number
      size: number
    }
    frontend: {
      load_time: number
      render_time: number
      user_interactions: number
    }
  }
  trends: {
    daily: Array<{
      date: Date
      score: number
      issues: string[]
    }>
    weekly: Array<{
      week: string
      score: number
      improvement: number
    }>
  }
}

interface AnalyticsData {
  overview: {
    total_users: number
    active_users: number
    total_content: number
    reported_issues: number
    system_health: number
  }
  users: {
    registration_trend: Array<{
      date: Date
      count: number
    }>
    retention: Record<string, number>
    engagement: Record<string, number>
    demographics: Record<string, number>
  }
  content: {
    creation_trend: Array<{
      date: Date
      count: number
    }>
    moderation_stats: {
      total_flagged: number
      approved: number
      rejected: number
      pending: number
    }
    performance: Record<string, number>
  }
  revenue: {
    monthly: Array<{
      month: string
      revenue: number
      growth: number
    }>
    user_value: Record<string, number>
    conversion_rates: Record<string, number>
  }
  issues: {
    total: number
    resolved: number
    pending: number
    by_type: Record<string, number>
    resolution_time_avg: number
  }
}

interface SecurityAudit {
  id: string
  type: 'user' | 'system' | 'data' | 'config'
  level: 'low' | 'medium' | 'high' | 'critical'
  performed_by: string
  performed_at: Date
  findings: Array<{
    id: string
    severity: string
    description: string
    recommendation: string
    status: 'open' | 'resolved'
  }>
  summary: {
    total_findings: number
    critical_findings: number
    resolved_findings: number
  }
}

interface SecurityThreat {
  id: string
  type: 'brute_force' | 'malicious_activity' | 'data_breach' | 'ddos'
  severity: 'low' | 'medium' | 'high' | 'critical'
  detected_at: Date
  source_ip: string
  target: string
  description: string
  status: 'active' | 'investigating' | 'resolved'
  actions_taken: string[]
}

interface AdminConfig {
  key: string
  value: any
  type: 'string' | 'number' | 'boolean' | 'object' | 'array'
  category: string
  description: string
  modified_by: string
  modified_at: Date
  validation_rules?: {
    type?: string
    min?: number
    max?: number
    required?: boolean
    options?: any[]
  }
}

interface AdminReport {
  id: string
  name: string
  type: 'standard' | 'custom'
  description: string
  data: any
  generated_by: string
  generated_at: Date
  format: 'pdf' | 'csv' | 'json' | 'html'
  scheduled?: boolean
  schedule?: {
    frequency: 'daily' | 'weekly' | 'monthly'
    time: string
    day?: string
  }
  last_sent?: Date
}

interface AdminTool {
  name: string
  description: string
  category: string
  status: 'available' | 'running' | 'failed'
  parameters: Record<string, any>
  last_run?: Date
  result?: any
  logs?: string[]
}
```

### Observable Metrics
- System health and uptime
- User activity and engagement
- Content moderation efficiency
- Performance bottlenecks
- Security incidents
- Revenue and growth metrics
- Admin task completion rates
- Platform stability metrics

### Integration Points
- User Service: User management
- Content Service: Content moderation
- Analytics Service: Reporting
- Security Service: Security monitoring
- Payment Service: Revenue tracking
- Database: System metrics

### Error Handling
```typescript
const ERRORS = {
  UNAUTHORIZED: new Error('Unauthorized admin access'),
  INSUFFICIENT_PERMISSIONS: new Error('Insufficient permissions'),
  SYSTEM_ERROR: new Error('System error'),
  CONFIG_ERROR: new Error('Configuration error'),
  VALIDATION_ERROR: new Error('Validation error'),
  DATA_CORRUPTION: new Error('Data corruption detected'),
  SERVICE_UNAVAILABLE: new Error('Admin service unavailable'),
} as const
```

### Performance Requirements
- Admin dashboard < 2s
- User management < 1s
- System monitoring < 500ms
- Content moderation < 3s
- Analytics reports < 5s
- Security alerts < 100ms
- 99.9% uptime for admin endpoints
- Support concurrent admin operations