# Delete User

Version: 0.1.0

Status: Stable

---

## Endpoint

DELETE /users/{id}

---

## Authentication

Not Required

---

## Success

200 OK

```json
{
    "success": true,
    "message": "User deleted successfully"
}
```

---

## Errors

404 User Not Found

400 Invalid ID

---

## Business Rules

User must exist before deletion.

---

## Database

DELETE FROM users