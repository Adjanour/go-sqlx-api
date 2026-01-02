# Go SQLX API

A standard and compliant Go backend template following Go best practices, using sqlx for database operations and the standard library for HTTP handling.

## Features

- ✅ **Standard Library First**: Uses Go's standard library (`net/http`) for HTTP server and routing
- ✅ **SQLX Integration**: Database operations with [sqlx](https://github.com/jmoiron/sqlx) for enhanced SQL capabilities
- ✅ **PostgreSQL Support**: Configured for PostgreSQL (easily adaptable to other databases)
- ✅ **RESTful API**: Clean RESTful endpoints following best practices
- ✅ **Middleware**: Logging, recovery, and CORS middleware
- ✅ **Graceful Shutdown**: Proper server shutdown handling
- ✅ **Docker Support**: Dockerfile and docker-compose for easy deployment
- ✅ **Project Structure**: Clean, idiomatic Go project layout
- ✅ **Error Handling**: Consistent error handling and responses
- ✅ **Configuration**: Environment-based configuration

## Project Structure

```
.
├── cmd/
│   └── api/              # Application entry point
│       └── main.go
├── internal/             # Private application code
│   ├── database/         # Database connection and migrations
│   ├── handlers/         # HTTP request handlers
│   ├── middleware/       # HTTP middleware
│   └── models/          # Data models and types
├── pkg/                  # Public reusable packages
│   └── response/        # HTTP response utilities
├── migrations/          # Database migrations
├── .env.example         # Environment variables template
├── .gitignore          # Git ignore rules
├── docker-compose.yml  # Docker Compose configuration
├── Dockerfile          # Docker image configuration
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── Makefile            # Build and development tasks
└── README.md           # This file
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher (or Docker)
- Make (optional, for using Makefile commands)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Adjanour/go-sqlx-api.git
cd go-sqlx-api
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Environment

Copy the example environment file and configure your settings:

```bash
cp .env.example .env
```

Edit `.env` to match your database configuration.

### 4. Run with Docker Compose (Recommended)

The easiest way to get started is using Docker Compose, which will start both PostgreSQL and the API:

```bash
docker-compose up -d
```

The API will be available at `http://localhost:8080`

### 5. Run Locally

If you have PostgreSQL running locally:

```bash
# Using Make
make run

# Or directly with Go
go run cmd/api/main.go
```

## API Endpoints

### Health Check

```bash
GET /health
```

Response:
```json
{
  "success": true,
  "data": {
    "status": "healthy",
    "database": "connected"
  }
}
```

### Users API

#### List All Users

```bash
GET /api/v1/users
```

#### Create User

```bash
POST /api/v1/users
Content-Type: application/json

{
  "username": "johndoe",
  "email": "john@example.com"
}
```

#### Get User by ID

```bash
GET /api/v1/users/{id}
```

#### Update User

```bash
PUT /api/v1/users/{id}
Content-Type: application/json

{
  "username": "newusername",
  "email": "newemail@example.com"
}
```

#### Delete User

```bash
DELETE /api/v1/users/{id}
```

## Development

### Available Make Commands

```bash
make help              # Show all available commands
make build             # Build the application
make run               # Run the application
make test              # Run tests
make test-coverage     # Run tests with coverage
make fmt               # Format code
make vet               # Run go vet
make tidy              # Tidy go modules
make clean             # Clean build artifacts
make docker-build      # Build Docker image
make docker-run        # Run Docker container
```

### Running Tests

```bash
make test

# Or with coverage
make test-coverage
```

### Code Formatting

```bash
make fmt
make vet
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_ADDRESS` | Server address and port | `:8080` |
| `DATABASE_URL` | Full database connection string | - |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database user | `postgres` |
| `DB_PASSWORD` | Database password | `postgres` |
| `DB_NAME` | Database name | `api_db` |

## Database

The application uses PostgreSQL as the default database. The connection is managed through sqlx.

### Migrations

Simple migrations are included in `internal/database/migrations.go`. For production applications, consider using a dedicated migration tool like [golang-migrate](https://github.com/golang-migrate/migrate).

## Docker

### Build Image

```bash
make docker-build
```

### Run Container

```bash
make docker-run
```

### Using Docker Compose

```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## Best Practices Implemented

1. **Standard Library First**: Minimizes external dependencies
2. **Clear Project Structure**: Follows Go project layout conventions
3. **Separation of Concerns**: Handlers, models, and database logic are separated
4. **Error Handling**: Consistent error handling throughout the application
5. **Middleware Pattern**: Composable middleware for cross-cutting concerns
6. **Graceful Shutdown**: Proper cleanup on server shutdown
7. **Configuration**: Environment-based configuration
8. **Database Connection Pooling**: Optimized database connections
9. **RESTful Design**: Clean and predictable API endpoints
10. **Type Safety**: Strong typing with Go structs

## License

MIT License - see LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request