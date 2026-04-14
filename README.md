# TaskFlow Backend — Saket Jasuja

---

## 🚀 1. Overview

TaskFlow is a minimal task management backend system.

It allows users to:
- Register and login using JWT authentication
- Create and manage projects
- Create, update, delete, and filter tasks within projects

The focus of this implementation is:
- Clean backend architecture
- Proper REST API design
- Production-ready setup using Docker and migrations

---

## 🛠️ Tech Stack

- Language: Go (Golang)
- Framework: Gin
- Database: PostgreSQL
- Driver: pgx (no ORM)
- Authentication: JWT + bcrypt
- Migrations: golang-migrate
- Containerization: Docker & Docker Compose

---

## 🧱 2. Architecture Decisions

### Architecture Pattern

Handler → Service → Repository → Database

- Handlers: HTTP layer
- Service: Business logic
- Repository: Database queries
- Models: Data structures

---

### Why This Approach?

- Clear separation of concerns
- Easy to maintain and scale
- Testable layers

---

### Key Decisions

- Used raw SQL instead of ORM
- JWT-based stateless authentication
- UUIDs for all primary keys
- PATCH APIs for partial updates

---

### Tradeoffs

- Simplified authorization (no roles/permissions)
- No pagination
- Minimal validation
- Basic logging

---

## ⚙️ 3. Running Locally

### Prerequisites

- Docker Desktop installed

---

### Steps

git clone https://github.com/your-username/taskflow-saket

cd taskflow-saket

cp .env.example .env

docker compose up --build

---

### API URL

http://localhost:8080

---

### Health Check

GET /health

Response:
{"status":"ok"}

---

## 🗄️ 4. Running Migrations

migrate -path backend/migrations -database "postgres://postgres:postgres@localhost:5432/taskflow?sslmode=disable" up

Note: **Run migrations manually after docker compose up**

---

## 🔐 5. Test Credentials

Email: test@example.com
Password: password123

---

## 📡 6. API Reference

### Auth

POST /auth/register

POST /auth/login

---

### Projects

GET /projects

POST /projects

GET /projects/:id

PATCH /projects/:id

DELETE /projects/:id

---

### Tasks

GET /projects/:id/tasks

POST /projects/:id/tasks

PATCH /tasks/:id

DELETE /tasks/:id

---

## 🐳 Docker

docker compose up --build

---

## 🚀 7. What I'd Do With More Time

- Add pagination
- Add role-based authorization
- Add tests
- Improve validation
- Add Swagger docs
- Structured logging
- Add caching
- Graceful shutdown
- DB indexing

---

## 🙌 Final Notes

This project demonstrates clean backend architecture, proper API design, and production-ready setup.
