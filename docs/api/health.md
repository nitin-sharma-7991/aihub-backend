# Health API

Checks whether the application is running.

---

## Endpoint

```
GET /health
```

---

## Response

```json
{
    "status": "ok"
}
```

---

## Status Code

```
200 OK
```

---

## Purpose

This endpoint is commonly used by:

- Load Balancers
- Docker
- Kubernetes
- Monitoring Systems
- CI/CD Pipelines