# Request Flow

Every HTTP request follows the same lifecycle.

```
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

GORM

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

## Handler

Responsible for

- Validation
- Parsing Request
- Returning Response

No business logic should exist here.

---

## Service

Responsible for

- Business rules
- Password hashing
- Validation
- Error mapping

---

## Repository

Responsible for

- Database queries
- CRUD operations

---

## Database

Persistent storage using PostgreSQL.