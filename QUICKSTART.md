# Quick Start Guide

This guide will help you get started with the Go SQLX API template.

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (recommended)
- PostgreSQL (if not using Docker)

## Option 1: Using Docker Compose (Recommended)

This is the easiest way to get started. Docker Compose will set up both PostgreSQL and the API:

```bash
# Start the services
docker-compose up -d

# Check the logs
docker-compose logs -f api

# The API is now running at http://localhost:8080
```

Test the health endpoint:
```bash
curl http://localhost:8080/health
```

## Option 2: Local Development

### 1. Set Up PostgreSQL

Make sure PostgreSQL is running locally or via Docker:

```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:15-alpine
```

### 2. Configure Environment

```bash
cp .env.example .env
# Edit .env if needed
```

### 3. Run the Application

```bash
# Using Make
make run

# Or directly
go run cmd/api/main.go
```

## Testing the API

### Health Check

```bash
curl http://localhost:8080/health
```

### Create a User

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"username": "john", "email": "john@example.com"}'
```

### List Users

```bash
curl http://localhost:8080/api/v1/users
```

### Get a User

```bash
curl http://localhost:8080/api/v1/users/1
```

### Update a User

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"username": "johndoe"}'
```

### Delete a User

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Development Commands

```bash
# Format code
make fmt

# Run tests
make test

# Run tests with coverage
make test-coverage

# Build the application
make build

# Clean build artifacts
make clean

# Run go vet
make vet

# Tidy dependencies
make tidy
```

## Project Structure

```
├── cmd/api/              # Application entry point (main.go)
├── internal/             # Private application code
│   ├── database/         # Database connection and migrations
│   ├── handlers/         # HTTP request handlers
│   ├── middleware/       # HTTP middleware
│   └── models/          # Data models and validation
├── pkg/                  # Public packages (response utilities)
└── migrations/          # Database migrations documentation
```

## Customizing the Template

### Adding a New Endpoint

1. Create a model in `internal/models/`
2. Add a handler function in `internal/handlers/`
3. Register the route in `cmd/api/main.go`

### Adding Middleware

1. Create a new middleware function in `internal/middleware/`
2. Add it to the middleware chain in `cmd/api/main.go`

### Database Migrations

Current migrations are in `internal/database/migrations.go`. For production:
- Consider using [golang-migrate](https://github.com/golang-migrate/migrate)
- See `migrations/README.md` for guidance

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_ADDRESS` | Server address | `:8080` |
| `DATABASE_URL` | Full database URL | - |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database user | `postgres` |
| `DB_PASSWORD` | Database password | `postgres` |
| `DB_NAME` | Database name | `api_db` |

## Troubleshooting

### Database Connection Issues

If you can't connect to the database:

1. Check PostgreSQL is running: `docker ps`
2. Verify environment variables in `.env`
3. Check logs: `docker-compose logs postgres`

### Port Already in Use

If port 8080 is already in use:

1. Change `SERVER_ADDRESS` in `.env` to `:8081` or another port
2. Update `docker-compose.yml` ports mapping

### Build Errors

If you encounter build errors:

```bash
# Clean and rebuild
make clean
go mod tidy
make build
```

## Next Steps

- Add authentication/authorization
- Implement additional business logic
- Add more comprehensive tests
- Set up CI/CD pipelines
- Configure production database
- Add monitoring and observability

## Additional Resources

- [Go Documentation](https://go.dev/doc/)
- [SQLX Documentation](https://github.com/jmoiron/sqlx)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://go.dev/doc/effective_go)
