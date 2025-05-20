# natural-stupidity
The contrast between artificial intelligence and natural stupidity

## Description
A Go-based web application to manage user profiles, hierarchical roles, and employee vacation calendars for Doutor Finanças.

## Features
- User profiles with contact & personal data
- Hierarchical roles & reporting (teams, directions, board)
- Vacation & sick day booking with holiday validation
- Admin and power user consoles
- JWT-based authentication & RBAC
- Live reload with Air for efficient development
- Swagger API documentation

## Prerequisites
- Go >= 1.23
- Docker & Docker Compose
- MySQL 8.0

## Environment Variables
Copy `.env.example` to `.env` and fill in your values:
```dotenv
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=natural_stupidity

APP_PORT=8080
GIN_MODE=debug
LOG_LEVEL=debug
JWT_SECRET=your_jwt_secret
``` 

## Local Development
1. Install Air:
```bash
go install github.com/air-verse/air@latest
```
2. Start the application:
```bash
air
```
The server will run at `http://localhost:${APP_PORT}` and live-reload on code changes.

## Docker & Docker Compose
Build and start services:
```bash
docker-compose up --build
```
- App: `http://localhost:${APP_PORT}`
- Database: `localhost:${DB_PORT}`

## Testing & Quality Checks
```bash
go test -v ./...
go vet -printf=false ./...
go fmt -s -w .
```

## API Documentation
Generate Swagger docs:
```bash
swag init -g main.go
```
Access UI at `http://localhost:${APP_PORT}/docs/index.html`.

## Contributing
- Follow conventional commits format
- Run tests before submitting PRs