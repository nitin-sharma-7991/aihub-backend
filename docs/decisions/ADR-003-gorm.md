# ADR-003: GORM ORM

Version: v0.1.0

Status: Accepted

---

## Context

The application requires an ORM to simplify database interactions and improve developer productivity.

---

## Decision

Use GORM as the ORM layer.

---

## Reasons

- Mature ecosystem
- Auto Migration
- Transactions
- Hooks
- Relationship support
- Active community

---

## Alternatives Considered

- database/sql
- sqlx
- Bun ORM

---

## Consequences

All database access should happen through repositories using GORM.