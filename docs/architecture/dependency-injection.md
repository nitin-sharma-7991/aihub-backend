# Dependency Injection

Version: 0.1.0

Status: Stable

---

## Purpose

Explain how dependencies are created.

---

## Current Flow

```
main.go

↓

config.Load()

↓

logger.New()

↓

database.New()

↓

user.New()

↓

router.New()

↓

server.New()

↓

Application Starts
```

---

## User Module

```
user.New()

↓

NewRepository()

↓

NewService()

↓

NewHandler()
```

---

## Dependency Graph

```
Handler

↓

Service

↓

Repository

↓

Database
```

Dependencies are injected from top to bottom.

Objects never create their own dependencies.

---

## Benefits

✔ Loose Coupling

✔ Easy Testing

✔ Easy Mocking

✔ Replaceable Components