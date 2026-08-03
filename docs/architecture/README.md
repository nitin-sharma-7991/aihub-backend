# Architecture Documentation

This directory documents the overall architecture of the AIHub Backend project.

The project follows a Feature-First Architecture combined with layered architecture principles.

## Documents

| File | Description |
|------|-------------|
| feature-first.md | Why Feature-First Architecture was chosen |
| folder-structure.md | Project folder organization |
| request-flow.md | HTTP request lifecycle |
| dependency-injection.md | Dependency injection strategy |

## Architecture Overview

```
Client
    │
    ▼
Gin Router
    │
    ▼
Middleware
(JWT / Logger / Recovery)
    │
    ▼
Handler
    │
    ▼
Service
    │
    ▼
Repository Interface
    │
    ▼
Repository Implementation
    │
    ▼
PostgreSQL
```

The responsibility of each layer is clearly separated to improve maintainability and scalability.