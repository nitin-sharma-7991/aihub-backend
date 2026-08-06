# Folder Structure

This document explains the directory structure of the AIHub Backend project.

The project follows a **Feature-First Architecture**, where every business feature owns its complete implementation.

---

# Project Structure

```text
aihub-backend/

├── cmd/
│   └── api/
│       └── main.go
│
├── configs/
│
├── docs/
│
├── internal/
│
│   ├── modules/
│   │
│   │   ├── auth/
│   │   ├── user/
│   │   ├── organization/
│   │   └── membership/
│   │
│   ├── server/
│   │
│   └── shared/
│       ├── apperrors/
│       ├── config/
│       ├── database/
│       ├── logger/
│       ├── middleware/
│       ├── pagination/
│       ├── request/
│       ├── response/
│       ├── router/
│       ├── security/
│       └── validation/
│
├── migrations/
│
├── scripts/
│
├── .env
├── go.mod
├── go.sum
└── README.md
```

---

# cmd/

Contains application entry points.

```text
cmd/

    api/

        main.go
```

Responsibilities:

- Load configuration
- Initialize logger
- Connect database
- Run migrations
- Initialize modules
- Configure router
- Start HTTP server
- Graceful shutdown

Business logic should never exist inside `cmd/`.

---

# internal/

Contains the application's source code.

Everything inside `internal/` is intended for use only by this project.

---

# internal/modules/

Each business feature lives inside its own module.

```text
modules/

    auth/

    user/

    organization/

    membership/
```

Every module is completely self-contained.

A typical module structure looks like:

```text
user/

    dto/
    handler/
    model/
    repository/
    service/
    module.go
```

---

## dto/

Contains request and response objects.

Examples:

- CreateUserRequest
- LoginRequest
- UserResponse

Responsibilities:

- Request validation
- API contracts
- Response serialization

---

## handler/

Represents the HTTP layer.

Responsibilities:

- Parse requests
- Bind JSON
- Validate input
- Call services
- Return HTTP responses

Handlers should never contain business logic.

---

## service/

Contains business logic.

Responsibilities:

- Business rules
- Validation
- DTO ↔ Model conversion
- Repository orchestration

Services are independent of HTTP.

---

## repository/

Contains database access logic.

Responsibilities:

- CRUD operations
- Queries
- Transactions

Repositories never contain business rules.

---

## model/

Contains GORM models.

Responsibilities:

- Database schema
- Relationships
- Table mapping

---

## module.go

Responsible for dependency injection.

Example:

```text
Repository

↓

Service

↓

Handler
```

The module wires everything together and exposes only the handler.

---

# internal/shared/

Contains reusable infrastructure shared across all modules.

---

## config/

Loads application configuration.

Responsibilities:

- Environment variables
- Configuration parsing
- Default values

---

## database/

Database initialization.

Responsibilities:

- PostgreSQL connection
- GORM setup
- Database migration

---

## logger/

Application logging.

Responsibilities:

- Zap logger initialization
- Structured logging

---

## middleware/

Reusable Gin middleware.

Current middleware:

- Request ID
- Recovery
- Logger
- JWT Authentication

---

## request/

Stores request-scoped helper utilities.

Example:

- Request ID context helpers

---

## response/

Standardized API responses.

Responsibilities:

- Success responses
- Error responses
- Pagination responses

Every API uses a consistent response format.

---

## validation/

Centralized request validation.

Responsibilities:

- JSON binding
- Validation
- Error formatting

---

## security/

Security-related helpers.

Examples:

- Password hashing
- JWT generation
- JWT validation

---

## pagination/

Reusable pagination infrastructure.

Responsibilities:

- Parse query parameters
- Normalize page and limit
- Offset calculation
- Pagination metadata

Shared by all list endpoints.

---

## apperrors/

Centralized application errors.

Examples:

- User not found
- Invalid credentials
- Email already exists
- Organization already exists

Provides consistent business errors across the application.

---

## router/

Registers all application routes.

Responsibilities:

- Route groups
- Middleware registration
- API versioning

---

# internal/server/

Responsible for HTTP server configuration.

Responsibilities:

- Server initialization
- Start server
- Graceful shutdown

---

# configs/

Contains configuration files used by different environments.

Examples:

- Development
- Staging
- Production

---

# migrations/

Contains database migration files.

Future versions may include versioned SQL migrations.

---

# docs/

Contains technical documentation.

Includes:

- Architecture
- APIs
- Setup guides
- ADRs
- Diagrams

---

# scripts/

Contains utility scripts used during development or deployment.

Examples:

- Build scripts
- Deployment scripts
- Seed scripts

---

# Why This Structure?

This folder structure provides:

- High modularity
- Clear separation of concerns
- Easy scalability
- Better maintainability
- Feature isolation
- Production-ready organization

New features can be added without affecting existing modules, making the project easier to understand and extend.