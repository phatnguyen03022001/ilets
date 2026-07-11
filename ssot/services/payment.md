# Payment Service

## Service Definition

The Payment Service manages all financial transactions, payment processing, billing, and revenue operations for the IELTS platform.

### Core Responsibilities
- Payment processing
- Subscription management
- Billing operations
- Revenue tracking
- Refunds and credits
- In-app purchases
- Payment analytics
- Fraud detection
- Multi-currency support

### API Endpoints
```typescript
// Core Payments
POST /api/payments/process       // Process payment
GET  /api/payments/status/:id    // Get payment status
GET  /api/payments/history      // Payment history
GET  /api/payments/transactions // Transaction list
POST /api/payments/refund       // Refund payment
GET  /api/payments/refunds      // Refund history

// Subscriptions
POST /api/payments/subscribe     // Create subscription
GET  /api/payments/subscriptions // List subscriptions
PUT  /api/payments/subscriptions/:id // Update subscription
DELETE /api/payments/subscriptions/:id // Cancel subscription
GET  /api/payments/subscriptions/:id/invoice // Get invoice

// Billing
GET  /api/payments/invoices      // List invoices
GET  /api/payments/invoices/:id // Get invoice details
POST /api/payments/invoices/:id/pay // Pay invoice
GET  /api/payments/invoices/:id/download // Download invoice
GET  /api/payments/billing-info // Get billing info
PUT  /api/payments/billing-info  // Update billing info

// In-App Purchases
POST /api/payments/iap         // Process in-app purchase
GET  /api/payments/iap/history  // Purchase history
GET  /api/payments/iap/products // Get products
POST /api/payments/iap/verify   // Verify purchase
GET  /api/payments/iap/status/:id // Purchase status

// Multi-Currency
GET  /api/payments/currencies   // Supported currencies
GET  /api/payments/rates        // Exchange rates
POST /api/payments/convert     // Convert currency
GET  /api/payments/pricing      // Multi-currency pricing

// Payment Methods
GET  /api/payments/methods     // Payment methods
POST /api/payments/methods     // Add payment method
DELETE /api/payments/methods/:id // Delete payment method
PUT  /api/payments/methods/:id // Set default method
POST /api/payments/methods/:id/verify // Verify method

// Analytics
GET  /api/payments/analytics   // Payment analytics
GET  /api/payments/revenue     // Revenue analytics
GET  /api/payments/metrics     // Performance metrics
GET  /api/payments/forecasts  // Revenue forecasts

// Fraud Detection
POST /api/payments/fraud/assess // Assess fraud risk
GET  /api/payments/fraud/history // Fraud history
POST /api/payments/fraud/block // Block transaction
GET  /api/payments/fraud/score // Risk score

// Webhooks
POST /api/payments/webhooks     // Register webhook
GET  /api/payments/webhooks     // List webhooks
DELETE /api/payments/webhooks/:id // Delete webhook
GET  /api/payments/webhooks/:id/test // Test webhook

// System
GET  /api/payments/health      // Health check
GET  /api/payments/metrics     // System metrics
POST /api/payments/settle      // Settle payments
GET  /api/payments/reconciliation // Reconciliation reports
POST /api/payments/test        // Test payment flow
```

### Data Models
```typescript
interface Payment {
  id: string
  user_id: string
  amount: number
  currency: string
  status: 'pending' | 'processing' | 'completed' | 'failed' | 'refunded' | 'cancelled'
  payment_method: string
  payment_gateway: string
  transaction_id: string
  order_id?: string
  description: string
  metadata: Record<string, any>
  created_at: Date
  completed_at?: Date
  refunded_at?: Date
  error?: string
  refund_id?: string
  subscription_id?: string
}

interface Subscription {
  id: string
  user_id: string
  plan_id: string
  status: 'active' | 'trial' | 'past_due' | 'cancelled' | 'expired'
  billing_cycle: 'monthly' | 'quarterly' | 'yearly'
  current_period: {
    start: Date
    end: Date
  }
  trial_period?: {
    start: Date
    end: Date
  }
  metadata: {
    created_at: Date
    updated_at: Date
    cancelled_at?: Date
    paused_at?: Date
    paused_reason?: string
  }
  features: Array<{
    name: string
    included: boolean
    limit?: number
  }>
  pricing: {
    amount: number
    currency: string
    period: string
  }
}

interface Invoice {
  id: string
  user_id: string
  subscription_id?: string
  number: string
  status: 'draft' | 'pending' | 'paid' | 'overdue' | 'void' | 'written_off'
  due_date: Date
  amount: number
  currency: string
  tax_amount: number
  total_amount: number
  items: Array<{
    description: string
    quantity: number
    unit_price: number
    total: number
  }>
  metadata: {
    created_at: Date
    sent_at?: Date
    paid_at?: Date
    notes?: string
  }
}

interface PaymentMethod {
  id: string
  user_id: string
  type: 'card' | 'paypal' | 'apple_pay' | 'google_pay' | 'bank_transfer'
  provider: string
  last_four?: string
  expiry_month?: number
  expiry_year?: number
  is_default: boolean
  is_active: boolean
  metadata: {
    created_at: Date
    updated_at: Date
    verified_at?: Date
  }
}

interface Transaction {
  id: string
  payment_id: string
  type: 'payment' | 'refund' | 'subscription' | 'in_app' | 'adjustment'
  amount: number
  currency: string
  status: 'pending' | 'completed' | 'failed' | 'reversed'
  gateway_response: Record<string, any>
  fees: {
    processing_fee: number
    currency_conversion_fee: number
    other_fees: number
    total_fees: number
  }
  created_at: Date
  processed_at?: Date
}

interface InAppPurchase {
  id: string
  user_id: string
  product_id: string
  purchase_id: string
  price: number
  currency: string
  status: 'pending' | 'completed' | 'failed' | 'cancelled'
  receipt: string
  metadata: {
    created_at: Date
    consumed_at?: Date
    device_info: Record<string, any>
  }
}

interface Currency {
  code: string
  name: string
  symbol: string
  exchange_rate: number
  is_active: boolean
  supported: boolean
}

interface Pricing {
  plan_id: string
  name: string
  description: string
  price: number
  currency: string
  period: string
  features: Array<{
    name: string
    included: boolean
    limit?: number
  }>
  is_popular: boolean
}

interface PaymentAnalytics {
  revenue: {
    monthly: number
    yearly: number
    growth_rate: number
    average_order_value: number
    lifetime_value: number
  }
  transactions: {
    total: number
    successful: number
    failed: number
    average_amount: number
  }
  subscriptions: {
    active: number
    churn_rate: number
    mrr: number
    ltv: number
  }
  refunds: {
    count: number
    amount: number
    rate: number
  }
  trends: {
    daily: Array<{
      date: Date
      revenue: number
      transactions: number
      new_customers: number
    }>
    monthly: Array<{
      month: string
      revenue: number
      growth_rate: number
    }>
  }
}

interface FraudAssessment {
  id: string
  transaction_id: string
  risk_score: number
  risk_level: 'low' | 'medium' | 'high' | 'critical'
  factors: Array<{
    type: string
    value: any
    weight: number
    impact: string
  }>
  recommendation: 'approve' | 'review' | 'block' | 'require_auth'
  reviewed_by?: string
  reviewed_at?: Date
  notes?: string
}

interface PaymentGateway {
  id: string
  name: string
  provider: string
  status: 'active' | 'inactive' | 'maintenance'
  currencies: string[]
  features: Array<{
    name: string
    available: boolean
  }>
  metadata: {
    connected_at: Date
    last_sync_at?: Date
    webhook_url?: string
  }
}

interface ReconciliationReport {
  id: string
  period: {
    start: Date
    end: Date
  }
  status: 'pending' | 'processing' | 'completed' | 'failed'
  totals: {
    expected: number
    actual: number
    variance: number
    variance_percentage: number
  }
  discrepancies: Array<{
    type: string
    amount: number
    reason: string
    reference: string
    status: 'pending' | 'resolved'
  }>
  metadata: {
    generated_at: Date
    reviewed_by?: string
    reviewed_at?: Date
  }
}

interface Webhook {
  id: string
  name: string
  url: string
  events: string[]
  secret?: string
  status: 'active' | 'inactive' | 'failed'
  metadata: {
    created_at: Date
    last_sent_at?: Date
    error_count: number
    error_details?: string[]
  }
}
```

### Observable Metrics
- Transaction success rate
- Revenue and growth metrics
- Subscription retention rate
- Payment method usage
- Average order value
- Fraud detection accuracy
- Refund rates
- Customer lifetime value

### Integration Points
- User Service: User billing information
- Email Service: Invoice notifications
- Analytics Service: Revenue analytics
- Notification Service: Payment alerts
- Third-party payment gateways
- Fraud detection services

### Error Handling
```typescript
const ERRORS = {
  PAYMENT_FAILED: new Error('Payment processing failed'),
  INSUFFICIENT_FUNDS: new Error('Insufficient funds'),
  INVALID_PAYMENT_METHOD: new Error('Invalid payment method'),
  SUBSCRIPTION_ERROR: new Error('Subscription operation failed'),
  INVOICE_ERROR: new Error('Invoice operation failed'),
  FRAUD_DETECTED: new Error('Potential fraud detected'),
  CURRENCY_CONVERSION_ERROR: new Error('Currency conversion failed'),
  WEBHOOK_ERROR: new Error('Webhook delivery failed'),
} as const
```

### Performance Requirements
- Payment processing < 3s
- Subscription updates < 2s
- Invoice generation < 1s
- Refund processing < 5s
- Fraud assessment < 500ms
- Webhook delivery < 1s
- 99.9% uptime for payment endpoints
- Handle 500 QPS peak load
- Support multiple payment gateways concurrently