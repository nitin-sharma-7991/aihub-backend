# Project Structure

Version: 0.1.0

Status: Stable

---

## Purpose

This document explains how the AIHub Backend source code is organized.

The project uses Feature-based Architecture instead of Layer-based Architecture.

---

## Folder Structure

```
AIHub Backend

├── cmd
│
├── internal
│
│    ├── modules
│    │
│    │      user
│    │
│    ├── shared
│    │
│    └── server
│
├── migrations
│
├── scripts
│
├── docs
│
└── test
```

---

## cmd

Application entrypoint.

Contains

```
main.go
```

Responsible for

- Loading configuration
- Creating dependencies
- Starting HTTP server

---

## internal

Contains all application logic.

Cannot be imported by external projects.

---

## modules

Each business feature lives inside its own module.

Example

```
user
auth
billing
notification
project
```

Every module owns its

- DTO
- Handler
- Service
- Repository
- Model

---

## shared

Contains reusable infrastructure.

Examples

- Logger
- Router
- Database
- Middleware
- Config

---

## migrations

Contains database migration files.

---

## docs

Contains project documentation.

---

## scripts

Development helper scripts.

---

## test

Integration and unit tests.

---

## Why Feature-first?

Instead of

```
handler

service

repository

model
```

we group by business feature

```
user

auth

project

billing
```

Advantages

✔ Easier maintenance

✔ Independent modules

✔ Better scalability

✔ Faster onboarding