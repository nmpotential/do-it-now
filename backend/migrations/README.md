# Database Migrations

All database schema changes for the Do It Now backend are managed via SQL migrations in this folder.

We use [golang-migrate/migrate](https://github.com/golang-migrate/migrate) for applying and creating migrations.

---

## How to Apply Migrations

Apply all migrations to your database:

```sh
migrate -path ./backend/migrations -database "$DB_URL" up
```

## How to Create a New Migration

Create a new migration file (replace `<migration_name>` with a descriptive name):

```sh
migrate create -ext sql -dir ./backend/migrations -seq <migration_name>
```

- Name migration files sequentially and descriptively, e.g., `001_create_tasks.sql`.
- Each migration should be idempotent and, if possible, reversible (provide a `down` migration).

---

## Migration File Naming Conventions

- Use sequential numbering: `001_`, `002_`, etc.
- Use snake_case for names: `001_create_tasks.sql`
- Add clear comments at the top and inline in each migration file to document schema changes and rationale.

---

## Tool Installation

Install the CLI:

```sh
brew install golang-migrate
# or see https://github.com/golang-migrate/migrate/tree/master/cmd/migrate for other install options
```

---

## Best Practices & Notes

- **Never edit a migration after it’s been applied to any environment.**  
  If you need to change the schema, create a new migration.
- **Document all schema changes** with SQL comments in each migration file.
- **Review and test migrations** locally before applying to production or shared environments.
- **See [`001_create_tasks.sql`](001_create_tasks.sql) for an example migration and schema documentation.**