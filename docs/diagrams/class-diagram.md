# User Module

```mermaid
classDiagram

class UserHandler

class UserService

class UserRepository

class User

UserHandler --> UserService

UserService --> UserRepository

UserRepository --> User
```