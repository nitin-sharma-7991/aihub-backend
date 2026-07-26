# API Documentation

This directory contains documentation for all REST APIs exposed by AIHub Backend.

---

# Base URL

```
http://localhost:8080
```

---

# Available APIs

| Module | Endpoint |
|---------|----------|
| Health | GET /health |
| Users | POST /users |
| Users | GET /users/:id |
| Users | PUT /users/:id |
| Users | DELETE /users/:id |

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