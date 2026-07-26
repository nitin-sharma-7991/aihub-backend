# ADR-002: PostgreSQL

Version: v0.1.0

Status: Accepted

---

## Context

The application requires a production-grade relational database capable of handling structured data efficiently.

---

## Decision

PostgreSQL has been selected as the primary database.

---

## Reasons

- Open Source
- ACID Compliant
- Excellent Performance
- Strong Indexing
- JSON Support
- Widely Used in Production

---

## Alternatives Considered

- MySQL
- SQLite

---

## Consequences

All persistent application data will be stored in PostgreSQL.