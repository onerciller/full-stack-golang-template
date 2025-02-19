# Fullstack Golang Template

A modern, well-structured template for building Go applications using Fiber framework.

## Project Structure

```
.
├── internal/
│   ├── api/        # HTTP handlers and routing
│   ├── store/      # Repository layer (database interactions)
│   └── provider/   # Dependency injection setup
├── pkg/            # Common packages
│   └── config/     # Configuration management
├── go.mod          # Go module file
├── main.go         # Application entry point
└── README.md       # This file
```

## Features

- Clean architecture with separation of concerns
- Fiber framework for high-performance HTTP routing
- Dependency injection using samber/do
- Configuration management with Viper
- Repository pattern for database operations
- Structured logging
- API versioning
- Health check endpoint

## Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL (optional, depending on your needs)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/fullstack-golang-template.git
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on port 3000 by default.

## API Endpoints

- `GET /health` - Health check endpoint
- `GET /api/v1/users` - List users (example endpoint)

## Configuration

Configuration can be provided through environment variables or a config file. Create a `config.yaml` file in the root directory:

```yaml
server:
  port: "3000"

db:
  host: "localhost"
  port: "5432"
  user: "postgres"
  password: "your-password"
  dbname: "your-database"
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 