# Repository Pattern

Version: 0.1.0

Status: Stable

---

## Purpose

Separate database access from business logic.

---

## Without Repository

```
Handler

↓

Database
```

Problems

- Tight Coupling
- Difficult Testing
- Mixed Responsibilities

---

## With Repository

```
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

Database
```

---

## Current User Repository

```
Create()

FindByID()

FindByEmail()

Update()

Delete()
```

---

## Advantages

✔ Cleaner Services

✔ Testable Code

✔ Easier Database Changes

✔ Better Maintainability