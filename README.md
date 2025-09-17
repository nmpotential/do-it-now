# do-it-now

A focus-enforcing task app that locks your phone until you complete your checklist. Built with Go, SwiftUI, Docker, and AWS.

---

## Local Development Setup

### Prerequisites

- [Go](https://golang.org/doc/install) (v1.21+)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Make](https://www.gnu.org/software/make/)

---

### 1. Clone the Repository

```
git clone https://github.com/yourusername/do-it-now.git
cd do-it-now
```

---

### 2. Environment Variables

- Copy `.env.example` to `.env` and update as needed.
- The default `DB_URL` is set for local Docker Compose usage.

---

### 3. Running the Stack

Build and start the backend API and Postgres database using Docker Compose via Makefile:

```
make docker-up-all
```

- This will build and start all services in the background.

Check container status:

```
docker compose ps
```

---

### 4. Database Migrations

Apply migrations to your local Postgres DB:

```
make migrate-up
```

Rollback the last migration:

```
make migrate-down-one
```

---

### 5. Health Check

Test that your backend API is running:

```
curl http://localhost:8080/healthz
```

You should see:

```
ok
```

---

### 6. Stopping the Stack

To stop and remove all containers and volumes:

```
make docker-down
```

---

### 7. Troubleshooting

- **Backend build errors:** Ensure `go.mod` and `go.sum` exist in `backend/`. Run `go mod tidy` if needed.
- **No Go files in /app:** Make sure your main Go entrypoint is in `backend/cmd/main.go` and your Dockerfile uses `RUN go build ... ./cmd`.
- **Database connection issues:** Ensure `DB_URL` uses `db` as the hostname (not `localhost`).
- **Port conflicts:** Make sure port 8080 is free or update the port mapping in `docker-compose.yml` and your Go app.

---

### 8. Makefile Commands Reference

- `make docker-up-all` — Build and start all services (API + DB)
- `make docker-down` — Stop and remove all services and volumes
- `make migrate-up` — Apply all migrations
- `make migrate-down-one` — Rollback the last migration
- `make db-shell` — Open a psql shell to the database

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
| DB_URL        | Postgres connection string         | Yes      | —       | postgres://user:password@db:5432/doitnow?sslmode=disable |
| SERVER_PORT   | HTTP server port                   | Yes      | —       | 8080                                                  |
| LOG_LEVEL     | Logging level (`info`, `debug`)    | No       | info    | info                                                  |

**Never commit real secrets.**  
Update `.env.example` if you add new config fields.

---

## Running Locally

1. Copy `.env.example` to `.env` and fill in your local values.
   - Do **not** commit `.env` (it is gitignored).
   - Update `.env.example` if you add new config fields.
   
2. Start services with Docker Compose via Makefile:

   ```
   make docker-up-all
   ```

3. The API will be available at `http://localhost:8080`.

---

## Database Migrations

All database schema changes are managed via SQL migrations in [`backend/migrations`](backend/migrations).

We use [golang-migrate/migrate](https://github.com/golang-migrate/migrate) for applying migrations.

- To apply migrations:  
  ```
  make migrate-up
  ```
- To create a new migration:
  ```
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
