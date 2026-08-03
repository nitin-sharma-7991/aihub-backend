# API Documentation

This directory contains documentation for all REST APIs exposed by AIHub Backend.

---

# Base URL

```
http://localhost:8080/api/v1
```

---

# Available APIs

| Module | Endpoint |
|---------|----------|
| Health | GET /health |

| Auth | POST /api/v1/auth/register |
| Auth | POST /api/v1/auth/login |
| Auth | GET /api/v1/auth/me |
| Auth | POST /api/v1/auth/logout |

| Users | POST /api/v1/users |
| Users | GET /api/v1/users/:id |
| Users | PUT /api/v1/users/:id |
| Users | DELETE /api/v1/users/:id |

| Organizations | POST /api/v1/organizations |
| Organizations | GET /api/v1/organizations |
| Organizations | GET /api/v1/organizations/:id |
| Organizations | PUT /api/v1/organizations/:id |
| Organizations | DELETE /api/v1/organizations/:id |

| Memberships | POST /api/v1/memberships |
| Memberships | GET /api/v1/memberships |
| Memberships | GET /api/v1/memberships/:id |
| Memberships | PUT /api/v1/memberships/:id |
| Memberships | DELETE /api/v1/memberships/:id |

---

# Response Format

## Success

```json
{
    "success": true,
    "message": "Operation completed successfully",
    "data": {}
}
```

---

## Error

```json
{
    "success": false,
    "message": "Error message"
}
```

---

# HTTP Status Codes

| Code | Meaning |
|------|---------|
| 200 | Success |
| 201 | Resource Created |
| 400 | Bad Request |
| 404 | Resource Not Found |
| 409 | Conflict |
| 500 | Internal Server Error |