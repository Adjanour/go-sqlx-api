# Database Migrations

This directory is for database migrations. Currently, migrations are embedded in the code at `internal/database/migrations.go`.

For production applications, consider using a dedicated migration tool like:

- [golang-migrate](https://github.com/golang-migrate/migrate)
- [goose](https://github.com/pressly/goose)
- [sql-migrate](https://github.com/rubenv/sql-migrate)

## Migration Best Practices

1. **Keep migrations simple**: Each migration should do one thing
2. **Make migrations reversible**: Always include up and down migrations
3. **Test migrations**: Test on a copy of production data
4. **Version control**: Track all migrations in version control
5. **Never modify existing migrations**: Create new migrations to change schema
6. **Use transactions**: Wrap migrations in transactions when possible

## Example Migration Files

If using golang-migrate, you would structure migrations like:

```
migrations/
├── 000001_create_users_table.up.sql
├── 000001_create_users_table.down.sql
├── 000002_add_user_indexes.up.sql
└── 000002_add_user_indexes.down.sql
```

### Example Up Migration (000001_create_users_table.up.sql)

```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

### Example Down Migration (000001_create_users_table.down.sql)

```sql
DROP TABLE IF EXISTS users;
```
