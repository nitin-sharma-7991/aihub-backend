# Request Flow

This document explains how an HTTP request travels through the AIHub Backend application.

The project follows a layered architecture where every layer has a single responsibility.

---

# Complete Request Lifecycle

```text
                Client
                   │
                   ▼
           HTTP Request
                   │
                   ▼
             Gin Router
                   │
                   ▼
        ┌────────────────────┐
        │    Middleware      │
        ├────────────────────┤
        │ Request ID         │
        │ Recovery           │
        │ Logger             │
        │ JWT Authentication │
        └────────────────────┘
                   │
                   ▼
               Handler
                   │
                   ▼
               Service
                   │
                   ▼
             Repository
                   │
                   ▼
             PostgreSQL
                   │
                   ▼
             Repository
                   │
                   ▼
               Service
                   │
                   ▼
               Handler
                   │
                   ▼
          Standard Response
                   │
                   ▼
             HTTP Response
```

---

# Step 1 — Client

The request originates from a client application.

Examples:

- React
- Vue
- Mobile App
- Postman
- Curl

Example:

```http
GET /api/v1/users?page=1&limit=10
Authorization: Bearer <JWT>
```

---

# Step 2 — Gin Router

The Gin router receives the request and matches it to the correct route.

Example:

```go
users := protected.Group("/users")
{
    users.GET("", userHandler.GetAll)
}
```

Responsibilities:

- Route matching
- API versioning
- Route grouping
- Middleware registration

---

# Step 3 — Middleware Pipeline

Every request passes through middleware before reaching the handler.

Current middleware execution order:

```text
Request

↓

Request ID

↓

Recovery

↓

Logger

↓

JWT Authentication

↓

Handler
```

---

## Request ID Middleware

Responsibilities:

- Generate unique request ID
- Store request ID in request context
- Make request traceable across logs

Example:

```
Request ID:
3c44e6c0-8c8d-4c74-a6e7-55c1a4f5a512
```

---

## Recovery Middleware

Responsibilities:

- Recover from panic
- Prevent server crash
- Log panic information
- Return HTTP 500

Without this middleware, one panic could terminate the server process.

---

## Logger Middleware

Responsibilities:

- Log every request
- Measure request duration
- Record HTTP status
- Record request path
- Record client IP

Example log:

```text
GET /api/v1/users

Status: 200

Duration: 14ms

RequestID: 3c44e6...
```

---

## JWT Authentication Middleware

Responsibilities:

- Read Authorization header
- Validate JWT
- Extract user information
- Store authenticated user in context

Protected routes only execute if authentication succeeds.

---

# Step 4 — Handler Layer

Handlers represent the HTTP layer.

Responsibilities:

- Bind request
- Validate request
- Parse URL parameters
- Call service
- Return HTTP response

Example:

```go
validation.BindJSON(ctx, &req)

service.Create(...)
```

Handlers never contain business logic.

---

# Step 5 — Service Layer

Services contain business rules.

Responsibilities:

- Validation
- Business decisions
- Repository orchestration
- DTO conversion
- Error mapping

Example:

```
Check duplicate email

↓

Hash password

↓

Create user

↓

Return response DTO
```

Services are independent of HTTP.

---

# Step 6 — Repository Layer

Repositories interact with PostgreSQL.

Responsibilities:

- CRUD
- Queries
- Transactions

Example:

```go
db.WithContext(ctx).
    Create(&user)
```

Repositories never contain business logic.

---

# Step 7 — PostgreSQL

The repository executes SQL operations using GORM.

Examples:

- INSERT
- SELECT
- UPDATE
- DELETE

Only repositories communicate with the database.

---

# Step 8 — Returning the Response

The response flows back through the same layers.

```text
Database

↓

Repository

↓

Service

↓

Handler

↓

Response Package

↓

Client
```

---

# Standard Response Package

All endpoints return a consistent JSON structure.

Example:

```json
{
    "success": true,
    "message": "Users fetched successfully",
    "data": [],
    "meta": {
        "page": 1,
        "limit": 10,
        "total": 45,
        "total_pages": 5
    }
}
```

Benefits:

- Consistent API contracts
- Easier frontend integration
- Predictable error handling

---

# Error Flow

Errors propagate upward through the layers.

```text
Repository

↓

Service

↓

Handler

↓

Response Package

↓

HTTP Error Response
```

Business errors are mapped to appropriate HTTP status codes.

Examples:

- 400 Bad Request
- 401 Unauthorized
- 404 Not Found
- 409 Conflict
- 500 Internal Server Error

---

# Design Principles

The request flow follows these principles:

- Single Responsibility Principle
- Separation of Concerns
- Dependency Injection
- Repository Pattern
- Feature-First Architecture

Each layer focuses on one responsibility, making the application easier to maintain, test, and scale.

---

# Benefits

This request lifecycle provides:

- Clear separation of responsibilities
- Consistent request processing
- Centralized logging
- Automatic panic recovery
- Secure authentication
- Standardized responses
- Easy debugging
- Production-ready architecture