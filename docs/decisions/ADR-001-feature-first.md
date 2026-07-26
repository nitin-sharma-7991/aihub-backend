# ADR-001: Feature-First Architecture

Version: v0.1.0

Status: Accepted

---

## Context

The project requires a scalable folder structure suitable for medium and large backend applications.

Traditional MVC architecture becomes difficult to maintain as the project grows.

---

## Decision

The project adopts a Feature-First Architecture.

Each feature owns its own:

- DTO
- Model
- Repository
- Service
- Handler
- module.go

Example

```
modules/

    user/

    auth/

    ai/

    billing/
```

---

## Benefits

- High cohesion
- Low coupling
- Easier maintenance
- Better scalability
- Independent modules
- Easier onboarding

---

## Alternatives Considered

- MVC
- Layer-based Architecture

---

## Consequences

New features can be developed independently without affecting existing modules.