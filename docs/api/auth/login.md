# Login

Version: 0.1.0

Status: Stable

---

## Endpoint

POST /api/v1/auth/login

---

## Authentication

Not Required

---

## Request

```json
{
  "email": "nitin@example.com",
  "password": "password123"
}
```

---

## Success

200 OK

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "access_token": "jwt_token",
    "token_type": "Bearer",
    "expires_in": "24h"
  }
}
```

---

## Errors

- 400 Validation Failed
- 401 Invalid Credentials

---

## Business Rules

- Verify email exists
- Verify password
- Generate JWT
- Return access token

---

## Related Files

- handler/auth_handler.go
- service/auth_service.go
- shared/security/jwt.go