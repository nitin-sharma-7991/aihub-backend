# Dependency Injection

## Overview

AIHub uses **Constructor Dependency Injection** to manage dependencies between application layers.

Instead of creating dependencies inside a component, they are created during application startup and injected through constructors.

This keeps the application loosely coupled, easier to test, and easier to maintain.

---

# Constructor Dependency Injection

Each layer receives its required dependency through its constructor.

```text
Repository

↓

Service

↓

Handler

↓

Router
```

Example:

```go
repo := repository.NewUserRepository(db)

service := service.NewUserService(repo)

handler := handler.NewUserHandler(service)
```

No layer is responsible for creating its own dependencies.

---

# Application Startup Flow

During application startup, all dependencies are initialized only once.

```text
main.go

↓

Load Config

↓

Initialize Logger

↓

Connect Database

↓

Run Migrations

↓

Initialize Modules

↓

Create Router

↓

Start HTTP Server
```

Each module follows the same initialization process.

```text
Repository

↓

Service

↓

Handler
```

Example:

```go
repo := repository.NewOrganizationRepository(db)

service := service.NewOrganizationService(repo)

handler := handler.NewOrganizationHandler(service)
```

---

# Dependency Flow

Dependencies always flow in one direction.

```text
Router

↓

Handler

↓

Service

↓

Repository

↓

PostgreSQL
```

No layer is allowed to depend on a higher layer.

For example:

- Router depends on Handler
- Handler depends on Service
- Service depends on Repository
- Repository depends on Database

Reverse dependencies are never allowed.

---

# Request Execution Flow

When a request reaches the application, it follows this lifecycle.

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

PostgreSQL

↓

Repository

↓

Service

↓

Handler

↓

JSON Response
```

Each layer has a single responsibility.

---

# Module Initialization

Every feature module is responsible for wiring its own dependencies.

Example:

```text
user/

↓

Repository

↓

Service

↓

Handler
```

The `module.go` file creates and injects dependencies.

Example:

```go
repo := repository.NewUserRepository(db)

service := service.NewUserService(repo)

handler := handler.NewUserHandler(service)

return &Module{
    Handler: handler,
    Service: service,
    Repo:    repo,
}
```

This keeps `main.go` clean and focused only on application startup.

---

# Why Dependency Injection?

Dependency Injection provides several advantages.

- Loose coupling
- Better maintainability
- Easier unit testing
- Clear dependency graph
- Easier refactoring
- Improved scalability

Components can be modified or replaced without affecting unrelated parts of the application.

---

# Architecture Principles

AIHub follows these architectural principles:

- Constructor Dependency Injection
- Feature-First Architecture
- Repository Pattern
- Separation of Concerns
- Single Responsibility Principle
- Modular Design

These principles make the project easier to understand, extend, and maintain as the codebase grows.