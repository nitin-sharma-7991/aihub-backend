# Register

Version: 0.1.0

Status: Stable

---

## Endpoint

POST /api/v1/auth/register

---

## Authentication

Not Required

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

## Success

201 Created

---

## Errors

- 400 Validation Failed
- 409 Email Already Exists

---

## Business Rules

- Validate request
- Check duplicate email
- Hash password
- Create user

---

## Related Files

- auth_handler.go
- auth_service.go
- user_repository.go