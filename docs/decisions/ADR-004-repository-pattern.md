# ADR-004: Repository Pattern

Version: v0.1.0

Status: Accepted

---

## Context

Business logic should remain independent of database implementation.

---

## Decision

Introduce a Repository layer between the Service layer and the Database.

```
Handler

↓

Service

↓

Repository

↓

Database
```

---

## Benefits

- Separation of concerns
- Easier testing
- Cleaner business logic
- Better maintainability

---

## Consequences

Only repositories communicate with the database.

Business logic remains inside services.