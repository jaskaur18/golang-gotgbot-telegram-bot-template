# Golang Telegram Bot Template (Refactored)

This is a production-ready, clean architecture template for building Telegram bots in Go using `gotgbot`.

## Architecture

The project follows Clean Architecture principles:

- **`internal/domain`**: Core entities and repository interfaces. Pure Go, no dependencies.
- **`internal/service`**: Business logic (e.g., `UserService`, `LocalizationService`).
- **`internal/repository`**: Data access implementations (Postgres via `sqlc`, Redis).
- **`internal/handler`**: Telegram update handlers and routing.
- **`internal/app`**: Dependency Injection and application wiring.
- **`internal/config`**: Configuration management using `envconfig`.

## Prerequisites
- Go 1.22+
- PostgreSQL
- Redis

## Setup

1. Copy `.env.sample` to `.env` and fill in values:
   ```bash
   cp .env.sample .env
   ```

2. Run with Go:
   ```bash
   go run main.go
   ```

## Development

- **Database Queries**: Managed by `sqlc`. Edit `queries/` and run `sqlc generate`.
- **Migrations**: Found in `migrations/`.

## Structure

```
.
├── cmd/
│   └── bot/            # Main entry point logic
├── internal/
│   ├── app/            # Application wiring
│   ├── config/         # Config loading
│   ├── domain/         # Entities & Interfaces
│   ├── handler/        # Telegram handlers
│   ├── repository/     # DB & Redis implementations
│   ├── service/        # Business logic
│   └── pkg/            # Shared utils (Logger, Validator)
└── main.go             # Entry point
```
