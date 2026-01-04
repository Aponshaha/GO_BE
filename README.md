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

- ✅ Clean architecture with separation of concerns
- ✅ PostgreSQL database integration
- ✅ RESTful API endpoints
- ✅ Middleware support (logging, CORS)
- ✅ Environment-based configuration
- ✅ Docker support
- ✅ Graceful shutdown

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

- `GET /health` - Health check endpoint
- `GET /api/health` - Alternative health check
- `GET /api/users?id={id}` - Get user by ID
- `GET /api/products` - Get all products

## Environment Variables

See `.env.example` for all available configuration options.

## Docker

Build and run with Docker:
```bash
docker compose up --build
```

## Development

Run with hot reload (requires air or similar tool):
```bash
air
```

