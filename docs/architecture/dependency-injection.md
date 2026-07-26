# Dependency Injection

AIHub uses Constructor Dependency Injection.

Example

```
Repository

↓

Service

↓

Handler

↓

Router
```

Example:

```go
repo := repository.NewUserRepository(db)

service := service.NewUserService(repo)

handler := handler.NewUserHandler(service)
```

## Benefits

- Loose coupling
- Easier testing
- Better maintainability
- Clear dependencies

## Current Dependency Graph

```
main.go

↓

database.New()

↓

user.New()

↓

Repository

↓

Service

↓

Handler

↓

Router
```

Every dependency is created only once during application startup.