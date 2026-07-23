# User Module APIs

Version: 0.1.0

Status: Stable

---

# Purpose

This directory contains all APIs related to the User module.

---

# Available Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | /users | Create User |
| GET | /users/{id} | Get User |
| PUT | /users/{id} | Update User |
| DELETE | /users/{id} | Delete User |

---

# Business Rules

- Email must be unique.
- Password is hashed before saving.
- Password is never returned.
- User ID must exist.
- Validation is performed before business logic.

---

# Layer Flow

Handler

↓

Service

↓

Repository

↓

Database