# Create Organization

Version: 0.1.0

Status: Stable

---

## Endpoint

POST /api/v1/organizations

---

## Authentication

Bearer Token Required

---

## Request

```json
{
  "name": "OpenAI",
  "slug": "openai",
  "description": "AI Research Company"
}
```

---

## Success

201 Created

---

## Errors

- 400 Validation Failed
- 409 Organization Already Exists

---

## Business Rules

- Slug must be unique
- Store organization
- Return DTO

---

## Database

INSERT organizations