# AIHub Architecture

**Version:** v0.3.0

---

# Overview

AIHub follows a **Feature-First Clean Architecture** designed for scalability, maintainability, and modular development.

The architecture separates responsibilities into distinct layers so that each component has a single responsibility.

Business logic remains independent of HTTP, database, and framework-specific implementations.

---

# High Level Architecture

```text
                Client
                   │
                   ▼
             Gin HTTP Router
                   │
                   ▼
        ┌──────────────────────┐
        │      Middleware      │
        │----------------------│
        │ • Request ID         │
        │ • Recovery           │
        │ • Logger             │
        │ • Authentication     │
        └──────────────────────┘
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
```

---

# Request Lifecycle

Every incoming request follows the same pipeline.

```text
Client

↓

Gin Router

↓

Middleware

↓

Handler

↓

Service

↓

Repository

↓

Database

↓

Repository

↓

Service

↓

Handler

↓

JSON Response
```

---

# Feature-First Architecture

Instead of organizing the project by technical layers, AIHub organizes code by business features.

```text
internal/

    modules/

        auth/

        user/

        organization/

        membership/
```

Each module owns everything related to itself.

- DTOs
- Handler
- Service
- Repository
- Model
- Module Initialization

This makes every feature self-contained and easier to maintain.

---

# Layer Responsibilities

## Router

The router is responsible for:

- Registering routes
- Grouping APIs
- Applying middleware
- API versioning

Business logic should never exist inside the router.

---

## Middleware

Middleware executes before the request reaches the handler.

Current middleware includes:

- Request ID
- Recovery
- Logger
- JWT Authentication

Responsibilities:

- Logging
- Panic Recovery
- Authentication
- Request Context

---

## Handler

Handlers represent the HTTP layer.

Responsibilities:

- Parse request
- Bind JSON
- Validate request
- Call service
- Return HTTP response

Handlers should never contain business logic.

---

## Service

The Service layer contains all business logic.

Responsibilities:

- Business rules
- Validation
- DTO to Model conversion
- Repository orchestration
- Error handling

Services never interact directly with HTTP.

---

## Repository

Repositories contain only database operations.

Responsibilities:

- CRUD
- Database queries
- Transactions
- Persistence

Repositories never contain business logic.

---

## Database

AIHub currently uses:

- PostgreSQL
- GORM ORM

Repositories are the only layer allowed to communicate with the database.

---

# Shared Packages

Common reusable components are placed inside:

```text
internal/shared/
```

Current shared packages include:

- apperrors
- config
- database
- logger
- middleware
- pagination
- request
- response
- router
- security
- validation

These packages are shared across all modules.

---

# Dependency Flow

Dependencies always flow downward.

```text
Router

↓

Handler

↓

Service

↓

Repository

↓

Database
```

Rules:

- Router depends on Handlers
- Handlers depend on Services
- Services depend on Repositories
- Repositories depend on Database

Reverse dependencies are never allowed.

---

# Design Principles

AIHub follows several software engineering principles.

## Single Responsibility Principle

Each layer has one clear responsibility.

---

## Dependency Injection

Dependencies are injected during module initialization.

This makes components:

- Testable
- Replaceable
- Loosely coupled

---

## Repository Pattern

All database access is abstracted through repositories.

Benefits:

- Separation of concerns
- Easier testing
- Cleaner business logic

---

## Feature Isolation

Every feature owns its implementation.

No feature directly manipulates another feature's internal logic.

---

## Shared Infrastructure

Common utilities are extracted into the shared package to avoid duplication.

Examples include:

- Validation
- Responses
- Security
- Logging
- Pagination

---

# Current Modules

AIHub currently contains:

- Authentication
- Users
- Organizations
- Memberships

Future modules will include:

- Invitations
- Projects
- AI Providers
- API Keys
- Billing
- Audit Logs

---

# Why This Architecture?

This architecture was chosen because it provides:

- Clear separation of concerns
- High maintainability
- Easy testing
- Better scalability
- Feature independence
- Production-ready project structure

As the application grows, new modules can be added without affecting existing functionality.

---

# Future Enhancements

Upcoming architectural improvements include:

- Dynamic RBAC
- Audit Logging
- Redis Caching
- Background Jobs
- Event Bus
- WebSockets
- Docker
- CI/CD
- Monitoring
- Distributed Services