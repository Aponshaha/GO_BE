# E-Commerce Go Backend

A RESTful API backend built with Go for an e-commerce application.

## Project Structure

```
GO_BE/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/                 # Configuration management
│   ├── database/               # Database connection
│   ├── models/                 # Data models
│   ├── handlers/               # HTTP handlers (controllers)
│   ├── services/               # Business logic
│   ├── repositories/           # Data access layer
│   ├── middleware/             # HTTP middleware
│   └── routes/                 # Route definitions
├── pkg/
│   └── utils/                  # Shared utilities
├── migrations/                 # Database migrations
└── docker-compose.yml          # Docker services
```

## Features

- ✅ Go 1.22+ with clean architecture (handlers → services → repositories)
- ✅ High-performance RESTful API with net/http
- ✅ PostgreSQL with connection pooling
- ✅ Request/response timeouts (15s read, 15s write, 60s idle)
- ✅ Graceful shutdown with context timeout
- ✅ Middleware: logging, CORS, request tracking
- ✅ JSON encoding/decoding for fast serialization
- ✅ Environment-based configuration
- ✅ Docker containerization

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL (via Docker)

## Setup

1. **Clone and navigate to the project**

   ```bash
   cd GO_BE
   ```

2. **Copy environment file**

   ```bash
   cp .env.example .env
   ```

3. **Start the database**

   ```bash
   docker compose up -d
   ```

4. **Install dependencies**

   ```bash
   go mod download
   ```

5. **Run the application**
   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

- `GET /health` - Health check
- `GET /api/users?id={id}` - Get user by ID
- `GET /api/products` - Get all products

## Environment Variables

See `.env.example` for all available configuration options.

## Docker

Build and run with Docker:

```bash
docker compose up --build
```

## Performance

- Connection pooling for database efficiency
- Timeout configurations prevent resource leaks
- Goroutine-based concurrent request handling
- Zero-allocation JSON encoding where possible
- Structured logging with request duration tracking
