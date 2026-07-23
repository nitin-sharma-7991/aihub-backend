# Health API

Version: 0.1.0

Status: Stable

---

# Purpose

Checks whether the backend service is running.

---

## Endpoint

GET /health

---

## Authentication

Not Required

---

## Request Body

None

---

## Success Response

Status

200 OK

```json
{
    "status": "ok"
}
```

---

## Business Logic

No business logic.

Used by

- Docker
- Kubernetes
- Load Balancer
- Monitoring
- Developers

---

## Related Source Files

internal/shared/router/router.go

internal/shared/router/health.go