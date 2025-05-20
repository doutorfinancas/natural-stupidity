# Technical Architecture and Decisions

## Languages and Frameworks
- Go (>= 1.23)
- Gin Gonic for HTTP server
- GORM for ORM
- Uber Dig for dependency injection
- godotenv for .env support
- Air for live reloading (https://github.com/air-verse/air)
- Bootstrap + htmx for frontend interactions

## Project Structure
```
.
├── api
│   ├── request
│   └── response
├── repository
│   └── model
├── service
├── docs
│   ├── PLAN.md
│   ├── STEP_BY_STEP.md
│   └── TECHNICAL.md
├── main.go
├── Dockerfile
├── .env.example
└── .air.toml
```

## Environment Variables
Load via `godotenv` in `main.go`. Example variables defined in `.env.example`.

## Database
- MySQL 8.0
- Use GORM AutoMigrate on startup
- Entities implement `model.Model` interface
- Relationships:
  - User ↔ Role (Many-to-One)
  - User ↔ Team (Many-to-One)
  - Team ↔ Direction (Many-to-One)
  - User ↔ Vacations (One-to-Many)

## Dockerization
- **Dockerfile**:
  - Base image: `golang:1.23-alpine`
  - Installs Air for hot reload
  - Copies source and dependencies
  - Exposes `${APP_PORT}`
- **docker-compose** (optional):
  - `app` service (builds Dockerfile, mounts code, uses Air)
  - `db` service (MySQL)

## Live Reload
- Air watches Go files and rebuilds/ restarts automatically
- Configuration in `.air.toml`

## Testing Strategy
- Unit tests: Table-driven, mocks for repositories
- Integration tests: Dockerized MySQL instance
- Use `go test`, `go vet`, `gofmt` for quality checks

## Logging and Observability
- Use structured logs (Logrus or Zap)
- Include request trace IDs
- Metrics endpoint `/metrics` for Prometheus

## API Documentation
- Use Swagger annotations in handlers
- Generate docs with `swag init -g main.go`
- Serve Swagger UI at `/docs/index.html` 