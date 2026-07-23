# Architecture Documentation

Version: 0.1.0

Status: Stable

Last Updated: 23 July 2026

---

## Purpose

This directory explains the internal architecture of AIHub Backend.

The goal is to help developers understand how the application is structured, how requests travel through the system, and why specific architectural decisions were made.

This documentation focuses on maintainability, scalability, and production-ready design rather than implementation details.

---

## Architecture Documents

| Document | Description |
|----------|-------------|
| project-structure.md | Folder organization |
| request-flow.md | Complete request lifecycle |
| dependency-injection.md | Dependency creation flow |
| repository-pattern.md | Data access architecture |
| database.md | PostgreSQL integration |

---

## Design Principles

AIHub follows

- Feature-first Architecture
- Layered Design
- Dependency Injection
- Repository Pattern
- Separation of Concerns
- Clean Code

---

## Architecture Overview

Client

↓

Gin Router

↓

Middleware

↓

Handler

↓

Service

↓

Repository

↓

Database

---

Each layer has a single responsibility.

Business logic never directly accesses HTTP.

Database logic never reaches handlers.

Dependencies always flow downward.