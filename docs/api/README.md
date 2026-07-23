# API Documentation

Version: 0.1.0

Status: Stable

Last Updated: 23 July 2026

---

# Purpose

This directory contains documentation for all HTTP APIs exposed by AIHub Backend.

Each API document follows a consistent structure to make it easier for developers to understand request formats, responses, validation rules, business logic, and possible errors.

---

# API Standards

Every endpoint documentation includes:

- Purpose
- Endpoint
- HTTP Method
- Authentication
- Headers
- Request Body
- Validation Rules
- Success Response
- Error Responses
- Business Rules
- Database Operation
- Related Source Files

---

# Available APIs

## Health

| Method | Endpoint |
|---------|----------|
| GET | /health |

---

## User Module

| Method | Endpoint |
|---------|----------|
| POST | /users |
| GET | /users/{id} |
| PUT | /users/{id} |
| DELETE | /users/{id} |

---

# Response Format

Successful response

```json
{
    "success": true,
    "message": "Success",
    "data": {}
}
```

---

Error response

```json
{
    "success": false,
    "message": "Something went wrong"
}
```

---

# Versioning

Current API Version

v1

Future

/api/v1
/api/v2