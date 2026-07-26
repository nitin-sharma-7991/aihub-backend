# AIHub Backend Context

## Project Goal

Production-ready AI Backend in Go.

---

## Tech Stack

Go 1.25

Gin

GORM

PostgreSQL

Zap

Viper

JWT (Future)

Redis (Future)

Docker (Future)

---

## Architecture

Feature First

cmd/

internal/

modules/

shared/

server/

---

## User Module Structure

handler

service

repository

model

dto

module.go

---

## Shared Structure

config

logger

database

router

apperrors

middleware (future)

validator (future)

---

## Documentation Standard

README

docs/

api/

architecture/

diagrams/

decisions/

setup/

development/

standards/

---

## Coding Rules

Always use Context

Always use DTO

Repository never returns DTO

Handler returns JSON

Service contains business logic

Repository only DB operations

No business logic in handler

No GORM in handler

No GORM in service

---

## Current Progress

✅ User CRUD

✅ PostgreSQL

✅ Repository Pattern

✅ Dependency Injection

✅ Architecture Docs

✅ API Docs

✅ Diagram Docs

---

## Next Sprint

ADR Documentation

JWT

Middleware

Validation

Docker

Testing