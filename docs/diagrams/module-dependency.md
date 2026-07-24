# Module Dependency

```mermaid
graph LR

Main

Router

UserModule

Handler

Service

Repository

Database

Main --> Database

Main --> UserModule

Main --> Router

Router --> Handler

Handler --> Service

Service --> Repository

Repository --> Database
```