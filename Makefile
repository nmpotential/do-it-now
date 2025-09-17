DB_URL ?= postgres://user:password@localhost:5432/doitnow?sslmode=disable
MIGRATE = migrate -path backend/migrations -database $(DB_URL)
.PHONY: migrate-up migrate-down migrate-up-one migrate-down-one migrate force migrate-status db-shell docker-up docker-down	reset-db

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-up-one:
	$(MIGRATE) up 1

migrate-down-one:
	$(MIGRATE) down 1

migrate-force:
	$(MIGRATE) force

migrate-status:
	$(MIGRATE) status

db-shell:
	psql $(DB_URL)

docker-up:
	docker compose up -d db

docker-down:
	docker compose down -v

reset-db:
	docker-down docker-up migrate-force migrate-up