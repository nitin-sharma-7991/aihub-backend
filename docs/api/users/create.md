# Create User

Version: 0.1.0

Status: Stable

---

# Purpose

Creates a new user.

---

## Endpoint

POST /users

---

## Authentication

Not Required

---

## Headers

```
Content-Type: application/json
```

---

## Request

```json
{
    "name": "Nitin Sharma",
    "email": "nitin@example.com",
    "password": "password123"
}
```

---

## Validation

| Field | Rule |
|---------|------|
| name | Required, Min 3 |
| email | Required, Valid Email |
| password | Required, Min 8 |

---

## Success

201 Created

```json
{
    "success": true,
    "message": "User created successfully",
    "data": {
        "id": 1,
        "name": "Nitin Sharma",
        "email": "nitin@example.com"
    }
}
```

---

## Errors

### Duplicate Email

409 Conflict

```json
{
    "success": false,
    "message": "email already exists"
}
```

---

### Validation Error

400 Bad Request

---

## Business Rules

- Duplicate email check
- Password hashing
- Store hashed password
- Return DTO

---

## Database

INSERT users

---

## Related Files

handler/user_handler.go

service/user_service.go

repository/user_repository.go