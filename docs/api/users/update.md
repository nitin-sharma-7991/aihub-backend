# Update User

Version: 0.1.0

Status: Stable

---

## Endpoint

PUT /users/{id}

---

## Request

```json
{
    "name": "Updated Name"
}
```

---

## Success

200 OK

```json
{
    "success": true,
    "message": "User updated successfully",
    "data": {
        "id": 1,
        "name": "Updated Name",
        "email": "nitin@example.com"
    }
}
```

---

## Errors

404 User Not Found

400 Invalid Request

---

## Business Rules

- User must exist
- Name updated only
- Email cannot be changed

---

## Database

UPDATE users