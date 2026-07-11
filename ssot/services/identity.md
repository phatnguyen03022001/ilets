# Identity Service

## Service Definition

The Identity Service manages user authentication, authorization, and identity management for the IELTS platform.

### Core Responsibilities
- User registration and authentication
- JWT token management
- Role-based access control
- Multi-factor authentication
- Profile management
- Security event logging

### API Endpoints
```typescript
// Authentication
POST /api/auth/login          // User login
POST /api/auth/register      // User registration
POST /api/auth/refresh       // Refresh token
POST /api/auth/logout        // User logout
POST /api/auth/verify        // Email verification

// Profile Management
GET  /api/profile             // Get user profile
PUT  /api/profile            // Update user profile
GET  /api/profile/preferences // Get user preferences
PUT  /api/profile/preferences // Update preferences

// Security
POST /api/auth/mfa/setup     // Setup 2FA
POST /api/auth/mfa/verify   // Verify 2FA
GET  /api/auth/sessions      // Get active sessions
DELETE /api/auth/sessions/:id // Terminate session
```

### Data Models
```typescript
interface User {
  id: string
  email: string
  username: string
  role: 'student' | 'teacher' | 'admin'
  profile: {
    firstName: string
    lastName: string
    avatar?: string
    preferences: {
      language: string
      timezone: string
      notifications: boolean
    }
  }
  createdAt: Date
  updatedAt: Date
  lastLoginAt?: Date
}

interface AuthToken {
  accessToken: string
  refreshToken: string
  expiresAt: Date
  scope: string[]
}

interface SecurityEvent {
  id: string
  userId: string
  type: 'login' | 'logout' | 'password_change' | 'mfa_setup' | 'suspicious_activity'
  ip: string
  userAgent: string
  timestamp: Date
  metadata: Record<string, any>
}
```

### Observable Metrics
- Login success/failure rate
- MFA adoption rate
- Session duration average
- Profile update frequency
- Security event count by type

### Integration Points
- Database: User data and security events
- Cache: Session management
- Monitoring: Security event tracking
- Email: Verification emails

### Error Handling
```typescript
const ERRORS = {
  INVALID_CREDENTIALS: new Error('Invalid email or password'),
  EMAIL_VERIFIED: new Error('Email already verified'),
  ACCOUNT_LOCKED: new Error('Account temporarily locked'),
  MFA_REQUIRED: new Error('Multi-factor authentication required'),
  INVALID_TOKEN: new Error('Invalid or expired token'),
  PERMISSION_DENIED: new Error('Insufficient permissions'),
} as const
```

### Performance Requirements
- Login response time < 200ms
- Profile update < 100ms
- Security event logging < 50ms (async)
- 99.9% uptime for authentication endpoints