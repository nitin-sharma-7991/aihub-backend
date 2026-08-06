# Repository Pattern

## Overview

AIHub follows the **Repository Pattern** to isolate database operations from business logic.

The Service layer never communicates directly with the database.

Instead, all persistence operations are performed through repositories.

This separation improves maintainability, testability, and scalability.

---

# Why Repository Pattern?

Without a repository layer:

```text
Handler

↓

Service

↓

Database
```

The service would become tightly coupled to GORM and SQL.

With the Repository Pattern:

```text
Handler

↓

Service

↓

Repository

↓

PostgreSQL
```

The service focuses only on business rules, while the repository manages data persistence.

---

# Responsibilities

## Service Layer

Responsibilities:

- Business rules
- Validation
- Authorization
- DTO conversion
- Error handling

The service never executes SQL queries.

Example:

```go
user, err := s.userRepo.FindByEmail(ctx, email)
```

---

## Repository Layer

Responsibilities:

- CRUD operations
- SQL queries
- Transactions
- Pagination queries
- GORM interaction

Example:

```go
func (r *userRepository) FindByEmail(
    ctx context.Context,
    email string,
) (*model.User, error) {

    var user model.User

    err := r.db.
        WithContext(ctx).
        Where("email = ?", email).
        First(&user).
        Error

    if err != nil {
        return nil, err
    }

    return &user, nil
}
```

---

# Repository Interface

Each module exposes an interface.

Example:

```go
type UserRepository interface {

    Create(ctx context.Context, user *model.User) error

    FindByID(ctx context.Context, id uint) (*model.User, error)

    FindByEmail(ctx context.Context, email string) (*model.User, error)

    Update(ctx context.Context, user *model.User) error

    Delete(ctx context.Context, id uint) error
}
```

The service depends on the interface instead of the implementation.

---

# Current Repository Flow

```text
HTTP Request

↓

Handler

↓

Service

↓

Repository Interface

↓

Repository Implementation

↓

GORM

↓

PostgreSQL
```

---

# Current Repositories

The AIHub backend currently includes repositories for:

- User
- Authentication
- Organization
- Membership

Each repository owns all database operations related to its feature.

---

# Benefits

Using the Repository Pattern provides:

- Separation of Concerns
- Loose Coupling
- Easier Unit Testing
- Centralized Database Logic
- Better Readability
- Reusable Queries
- Easier Refactoring
- Production-Ready Architecture

---

# Design Principles

The Repository Pattern in AIHub follows these principles:

- Single Responsibility Principle
- Dependency Inversion Principle
- Feature-First Architecture
- Constructor Dependency Injection

Business logic remains inside services, while repositories are responsible only for data access.