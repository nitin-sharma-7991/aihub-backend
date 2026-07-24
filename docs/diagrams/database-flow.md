# Database Flow

```mermaid
flowchart TD

Application

Repository

GORM

ConnectionPool

PostgreSQL

Application --> Repository

Repository --> GORM

GORM --> ConnectionPool

ConnectionPool --> PostgreSQL
```