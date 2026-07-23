# Get User

Version: 0.1.0

Status: Stable

---

## Endpoint

GET /users/{id}

---

## Authentication

Not Required

---

## Path Parameter

| Name | Type |
|------|------|
| id | uint |

---

## Success

200 OK

```json
{
    "success": true,
    "data": {
        "id": 1,
        "name": "Nitin",
        "email": "nitin@example.com"
    }
}
```

---

## Errors

404 User Not Found

400 Invalid ID

---

## Database

SELECT users