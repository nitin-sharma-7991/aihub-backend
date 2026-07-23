# Request Lifecycle

Version: 0.1.0

Status: Stable

---

## Purpose

This document explains how every HTTP request travels through AIHub.

---

## Flow

```
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

GORM

↓

PostgreSQL

↓

Repository

↓

Service

↓

Handler

↓

JSON Response
```

---

## Step 1

Client sends HTTP request.

Example

POST /users

---

## Step 2

Router finds matching endpoint.

```
POST /users
```

↓

UserHandler.Create()

---

## Step 3

Middleware executes.

Examples

Logger

Recovery

Request ID

Authentication (future)

---

## Step 4

Handler

Responsibilities

- Parse request
- Validate request
- Call service
- Return JSON

No business logic.

---

## Step 5

Service

Responsible for

Business Rules

Example

Duplicate Email Check

Password Hashing

Validation

---

## Step 6

Repository

Responsible for

Database Operations only.

No business logic.

---

## Step 7

GORM

Converts Go structs into SQL.

---

## Step 8

PostgreSQL

Stores persistent data.

---

## Response Flow

Database

↓

Repository

↓

Service

↓

Handler

↓

Client