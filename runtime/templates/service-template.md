# Service Template

## Service Implementation Template

### 1. Service Definition
```typescript
// Service interface definition
interface Service {
  // Core methods
  create(data: CreateDto): Promise<Entity>;
  read(id: string): Promise<Entity>;
  update(id: string, data: UpdateDto): Promise<Entity>;
  delete(id: string): Promise<void>;
  
  // Business logic
  processBusinessLogic(input: Input): Promise<Output>;
}
```

### 2. Data Models
```typescript
// Entity definition
interface Entity {
  id: string;
  createdAt: Date;
  updatedAt: Date;
  // Additional fields based on service
}

// DTOs for API layer
interface CreateDto {
  // Required fields
  requiredField: string;
  
  // Optional fields
  optionalField?: string;
}

interface UpdateDto {
  // Fields that can be updated
  optionalField?: string;
}
```

### 3. Repository Layer
```typescript
// Repository interface
interface Repository {
  save(entity: Entity): Promise<Entity>;
  findById(id: string): Promise<Entity | null>;
  update(id: string, data: Partial<Entity>): Promise<Entity>;
  delete(id: string): Promise<void>;
}
```

### 4. Service Implementation
```typescript
// Service implementation with proper error handling
class ServiceImpl implements Service {
  constructor(private repository: Repository) {}
  
  async create(data: CreateDto): Promise<Entity> {
    // Business logic
    const entity = new Entity();
    // ... populate entity
    
    return await this.repository.save(entity);
  }
  
  // ... other service methods
}
```

### 5. API Endpoints
```typescript
// Route definitions
router.post('/api/service', createValidation, serviceController.create);
router.get('/api/service/:id', serviceController.read);
router.put('/api/service/:id', updateValidation, serviceController.update);
router.delete('/api/service/:id', serviceController.delete);
```

### 6. Validation
```typescript
// Zod schemas for validation
const createSchema = z.object({
  requiredField: z.string().min(1),
  optionalField: z.string().optional(),
});

const updateSchema = z.object({
  optionalField: z.string().optional(),
});
```

### 7. Testing
```typescript
// Unit tests
describe('Service', () => {
  let service: Service;
  
  beforeEach(() => {
    // Setup test dependencies
  });
  
  it('should create entity', async () => {
    // Test implementation
  });
});
```

### 8. Observable Metrics
```typescript
// Metrics collection
interface ServiceMetrics {
  requests: number;
  errors: number;
  averageResponseTime: number;
  successRate: number;
}
```

### 9. Integration Points
- Database: PostgreSQL via Drizzle ORM
- Cache: Redis integration (when needed)
- Monitoring: Metrics collection
- Logging: Structured logging

### 10. Error Handling
```typescript
// Custom error types
class ServiceError extends Error {
  constructor(message: string, public code: string) {
    super(message);
  }
}

// Error handling in services
try {
  // Service logic
} catch (error) {
  if (error instanceof ValidationError) {
    // Handle validation errors
  } else if (error instanceof DatabaseError) {
    // Handle database errors
  }
  throw new ServiceError('Service failed', 'SERVICE_ERROR');
}
```