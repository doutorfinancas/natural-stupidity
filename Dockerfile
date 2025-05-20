FROM golang:1.23-alpine

# Install bash and git
RUN apk add --no-cache bash git

# Set working directory
WORKDIR /app

# Install Air for live reload
RUN go install github.com/air-verse/air@latest
ENV PATH="/go/bin:${PATH}"

# Copy dependency files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Expose application port
EXPOSE ${APP_PORT}

# Run with Air
CMD ["air", "-c", ".air.toml"] 