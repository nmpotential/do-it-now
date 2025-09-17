# do-it-now
A focus-enforcing task app that locks your phone until you complete your checklist. Built with Go, SwiftUI, Docker, and AWS.

## Overview

Do It Now is a focus-enforcing task app that locks your phone until you complete your checklist.  
Backend is written in Go, with a SwiftUI iOS frontend.

---

## Architecture

- **Backend:** Go (REST API, clean architecture)
- **Database:** Postgres
- **Infra:** Docker, Docker Compose, GitHub Actions, AWS (EC2 + RDS)
- **Frontend:** SwiftUI (iOS)
- **Testing:** Unit, integration (real Postgres), HTTP tests

---

## Configuration & Environment

The backend is configured via environment variables.  
Copy `.env.example` to `.env` and update values as needed.

| Variable      | Description                        | Required | Default | Example Value                                         |
|---------------|------------------------------------|----------|---------|-------------------------------------------------------|
| DB_URL        | Postgres connection string         | Yes      | —       | postgres://user:password@localhost:5432/doitnow?sslmode=disable |
| SERVER_PORT   | HTTP server port                   | Yes      | —       | 8080                                                  |
| LOG_LEVEL     | Logging level (`info`, `debug`)    | No       | info    | info                                                  |


**Never commit real secrets.**  
Update `.env.example` if you add new config fields.

---

## Running Locally

1. Copy `.env.example` to `.env` and fill in your local values.
   - Do **not** commit `.env` (it is gitignored).
   - Update `.env.example` if you add new config fields.
   
2. Start services with Docker Compose:

   ```sh
   docker-compose up --build
   ```

3. The API will be available at `http://localhost:8080`.

---

## Database Migrations

All database schema changes are managed via SQL migrations in [`backend/migrations`](backend/migrations).

We use [golang-migrate/migrate](https://github.com/golang-migrate/migrate) for applying migrations.

- To apply migrations:  
  ```sh
  migrate -path ./backend/migrations -database "$DB_URL" up
   ```
- To create a new migration:
  ```sh
   migrate create -ext sql -dir ./backend/migrations -seq <migration_name>   
   ```

See [`backend/migrations/README.md`](backend/migrations/README.md) for full instructions and best practices.

---

## API Quickstart

See [`internal/http/contract.md`](backend/internal/http/contract.md) for endpoint docs.

---

## Key Decisions

- Clean architecture: business logic in `core`, thin handlers, repo interfaces.
- Context everywhere, errors as values, structured logging.
- 12-factor config via env vars.

---

## Roadmap

- [ ] MVP backend: tasks API, Postgres integration, CI/CD
- [ ] MVP iOS app: checklist, lockdown UI
- [ ] Metrics, notifications, AWS deployment

---
