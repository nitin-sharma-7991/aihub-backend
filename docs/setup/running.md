# Running the Application

Version: v0.1.0

Status: Stable

---

## Install Dependencies

```bash
go mod tidy
```

---

## Start Application

```bash
go run cmd/api/main.go
```

---

## Verify

Health Endpoint

```
GET http://localhost:8080/health
```

Expected Response

```json
{
    "status": "ok"
}
```

---

## Stop Server

Press

```
CTRL + C
```

The application performs a graceful shutdown.