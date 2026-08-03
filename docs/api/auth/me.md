# Current User

Version: 0.1.0

Status: Stable

---

## Endpoint

GET /api/v1/auth/me

---

## Authentication

Bearer Token Required

---

## Success

200 OK

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Nitin",
    "email": "nitin@example.com",
    "role": "admin"
  }
}
```

---

## Errors

- 401 Unauthorized

---

## Business Rules

- Read user id from JWT
- Fetch latest user details

---

## Related Files

- auth_handler.go
- auth_service.go