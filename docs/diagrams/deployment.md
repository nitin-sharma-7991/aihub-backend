# Deployment Architecture

```mermaid
flowchart TD

Developer

LocalMachine

GinServer

PostgreSQL

Docker

Cloud

Developer --> LocalMachine

LocalMachine --> GinServer

GinServer --> PostgreSQL

GinServer -. Future .-> Docker

Docker -. Future .-> Cloud
```

---

## Current Deployment

- Local Machine
- Gin
- PostgreSQL

---

## Planned Deployment

- Docker
- Docker Compose
- Nginx
- CI/CD
- Cloud Deployment