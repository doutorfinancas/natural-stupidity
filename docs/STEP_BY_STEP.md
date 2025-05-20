# Step-by-Step Implementation Guide

This guide provides ordered steps to build and run the Natural Stupidity project.

1. Clone Repository
   ```bash
   git clone git@github.com:doutorfinancas/natural-stupidity.git
   cd natural-stupidity
   ```

2. Setup Environment
   - Copy environment template:
     ```bash
     cp .env.example .env
     ```
   - Fill in database and secret values in `.env`.

3. Initialize Go Modules
   ```bash
   go mod tidy
   ```

4. Configure Live Reload
   - Install Air locally:
     ```bash
     go install github.com/air-verse/air@latest
     ```
   - Initialize Air config:
     ```bash
     air init
     ```
   - Review `.air.toml`.

5. Build and Run with Docker
   - Build Docker image:
     ```bash
     docker build -t natural-stupidity .
     ```
   - Run database container:
     ```bash
     docker run -d --name ns-db -e MYSQL_ROOT_PASSWORD=${DB_PASSWORD} -e MYSQL_DATABASE=${DB_NAME} -p 3306:3306 mysql:8.0
     ```
   - Run application container:
     ```bash
     docker run --rm -it --name ns-app --env-file .env -p ${APP_PORT}:${APP_PORT} -v $(pwd):/app natural-stupidity
     ```

6. Run Locally with Air
   ```bash
   air
   ```

7. Access the Application
   - API base URL: `http://localhost:${APP_PORT}`
   - Swagger UI: `http://localhost:${APP_PORT}/docs/index.html`

8. Run Tests
   ```bash
   go test -v ./...
   ```

9. Code Quality Checks
   ```bash
   go vet -printf=false ./...
   go fmt -s -w .
   ```

10. Commit and Push
    - Follow conventional commits
    - Push to feature branch 