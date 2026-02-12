# E-commerce Microservices Backend

A microservices-based e-commerce backend built with Go and PostgreSQL.

## Architecture

This project consists of 4 microservices:

1. **Product Service** - Manages product catalog, categories, and product information
2. **Inventory Service** - Handles stock management, warehouses, and inventory tracking
3. **Order Service** - Processes orders, order items, and order lifecycle
4. **Payment Service** - Manages payments, payment methods, and transactions

Each service has its own PostgreSQL database.

## Prerequisites

- Docker and Docker Compose
- Go 1.24.5 or higher

## Setup

### 1. Start Databases with Docker

```bash
docker-compose up -d
```

This will start 4 PostgreSQL databases:
- Product DB: `localhost:5432`
- Inventory DB: `localhost:5433`
- Order DB: `localhost:5434`
- Payment DB: `localhost:5435`

### 2. Environment Variables

Copy the example environment file:

```bash
cp .env.example .env
```

Edit `.env` if you need to change database credentials.

### 3. Install Dependencies

```bash
go mod download
```

### 4. Verify Database Connections

Each service has a database connection module in `services/{service}/db/db.go`. You can test the connection by importing and calling:

```go
import "main.go/services/product/db"

config := db.LoadConfig()
err := db.Connect(config)
if err != nil {
    log.Fatal("Failed to connect:", err)
}
defer db.Close()
```

## Database Schemas

Each service has its schema defined in:
- `services/product/db/schema.sql`
- `services/inventory/db/schema.sql`
- `services/order/db/schema.sql`
- `services/payment/db/schema.sql`

Schemas are automatically initialized when Docker containers start.

## Project Structure

```
EcomBackend/
├── docker-compose.yml          # Docker configuration for all databases
├── .env.example                 # Environment variables template
├── services/
│   ├── product/
│   │   ├── db/
│   │   │   ├── schema.sql      # Product database schema
│   │   │   └── db.go           # Database connection
│   │   └── models/
│   │       └── models.go       # Product models
│   ├── inventory/
│   │   ├── db/
│   │   │   ├── schema.sql
│   │   │   └── db.go
│   │   └── models/
│   │       └── models.go
│   ├── order/
│   │   ├── db/
│   │   │   ├── schema.sql
│   │   │   └── db.go
│   │   └── models/
│   │       └── models.go
│   └── payment/
│       ├── db/
│       │   ├── schema.sql
│       │   └── db.go
│       └── models/
│           └── models.go
```

## Database Connection

Each service uses `pgx/v5` for PostgreSQL connection pooling. Connection configuration is loaded from environment variables:

- `{SERVICE}_DB_HOST` - Database host (default: localhost)
- `{SERVICE}_DB_PORT` - Database port
- `{SERVICE}_DB_USER` - Database user
- `{SERVICE}_DB_PASSWORD` - Database password
- `{SERVICE}_DB_NAME` - Database name
- `{SERVICE}_DB_SSLMODE` - SSL mode (default: disable)

## Next Steps

- [ ] Create repository/data access layers
- [ ] Implement HTTP handlers and routes
- [ ] Add API endpoints for each service
- [ ] Set up service-to-service communication
- [ ] Add authentication and authorization
- [ ] Implement logging and monitoring

## Commands

```bash
# Start databases
docker-compose up -d

# Stop databases
docker-compose down

# View logs
docker-compose logs -f

# Restart a specific database
docker-compose restart product-db
```
