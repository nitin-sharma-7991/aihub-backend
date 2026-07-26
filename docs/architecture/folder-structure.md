# Folder Structure

```
cmd/
    api/

internal/

    modules/

        user/

            dto/
            handler/
            model/
            repository/
            service/
            module.go

    shared/

        config/
        database/
        logger/
        router/
        apperrors/

    server/

docs/
```

## cmd/

Application entry point.

## internal/

Contains all application source code.

## modules/

Business features.

Every feature owns its own:

- DTO
- Handler
- Service
- Repository
- Model

## shared/

Reusable components shared across all modules.

## server/

HTTP server initialization.