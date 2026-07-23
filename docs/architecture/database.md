# Database Architecture

Version: 0.1.0

Status: Stable

---

## Database

PostgreSQL

ORM

GORM

---

## Connection Flow

```
main.go

↓

database.New()

↓

gorm.Open()

↓

Connection Pool

↓

Application
```

---

## Current Database

```
users
```

---

## Migration

Current migration strategy

AutoMigrate()

Future

golang-migrate

---

## Connection Pool

Configured

- Max Idle Connections

- Max Open Connections

- Connection Lifetime

---

## Future Improvements

- Read Replicas

- Database Transactions

- Soft Deletes

- Audit Logs

- Multi-Tenant Support