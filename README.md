# My Portfolio Website

[![CI](https://github.com/benidevo/website/actions/workflows/ci.yaml/badge.svg)](https://github.com/benidevo/website/actions/workflows/ci.yaml)
[![Go Version](https://img.shields.io/badge/Go-1.24.2-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A modern, modular portfolio website built with Go.

## 🛠️ Technology Stack

### Backend

- **Language**: Go 1.24.2
- **Web Framework**: Gin
- **Logging**: Zerolog
- **Architecture**: Monolithic with clean separation of concerns

### Frontend

- **UI Library**: HTMX
- **JavaScript**: Alpine.js
- **Styling**: Tailwind CSS

## 📁 Project Structure

``` plaintext
website/
├── cmd/
│   └── website/          # Application entry point
├── internal/             # Private application code
│   ├── app/             # Application lifecycle management
│   ├── config/          # Configuration, logging, and dependency setup
│   ├── client/          # External API clients (GitHub)
│   ├── handlers/        # HTTP request handlers
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data structures and domain models
│   ├── repository/      # Data access layer (GitHub API + in-memory)
│   ├── router/          # HTTP routing and template setup
│   └── services/        # Business logic and service orchestration
├── web/                 # Frontend assets
│   ├── templates/       # HTML templates (layouts, pages, partials)
│   ├── static/          # CSS, JS, images
│   └── assets/          # Source assets
├── docker/              # Docker configurations
└── data/                # Static data files
```

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose
- Make (optional but recommended)
- Go 1.24.2 (for local development without Docker)

### Development Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/benidevo/website.git
   cd website
   ```

2. **Start the development environment**

   ```bash
   make up
   ```

   The application will be available at <http://localhost:8888>

3. **View logs**

   ```bash
   make logs
   ```

4. **Access the container shell**

   ```bash
   make shell
   ```

### Available Make Commands

**Application Management:**

| Command | Description |
|---------|-------------|
| `make up` | Start the application in development mode |
| `make down` | Stop the application |
| `make logs` | View application logs |
| `make shell` | Access container shell |
| `make build` | Build Docker image |
| `make restart` | Restart the application |
| `make clean` | Remove all containers and images |

**Development Tools:**

| Command | Description |
|---------|-------------|
| `make format` | Format Go code |
| `make deps-install` | Install/update dependencies |
| `make lint` | Run linters (gofmt and go vet) |
| `make check` | Run all checks (lint + test) |

**Testing:**

| Command | Description |
|---------|-------------|
| `make test` | Run all tests |
| `make test-short` | Run short tests only |
| `make test-specific TEST=TestName` | Run specific test by name |

## 🔧 Configuration

The application uses environment variables for configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `IS_DEVELOPMENT` | Development mode | `true` |
| `LOG_LEVEL` | Logging level | `info` |
| `GITHUB_OWNER` | GitHub repository owner | |
| `GITHUB_REPOSITORY` | GitHub repository name | |
| `GITHUB_TOKEN` | GitHub API token (optional) | |
| `GITHUB_BASE_URL` | GitHub API base URL | `https://api.github.com` |

## 🏗️ Architecture

The application follows a clean architecture pattern with dependency injection

### Key Architectural Principles

- **Dependency Injection**: Clear dependency contracts with setup functions
- **Repository Pattern**: Abstracted data access with GitHub API and in-memory implementations
- **Service Layer**: Business logic separated from HTTP concerns
- **Error Handling**: Global panic recovery with proper HTTP status codes (404, 500)
- **Template Rendering**: Multi-template system with shared layouts and partials
- **Graceful Shutdown**: Proper signal handling for zero-downtime deployments
- **Structured Logging**: Consistent, queryable logs with Zerolog

## 🧪 Testing

The project includes comprehensive unit tests for all major components:

- **Services**: Business logic and service orchestration
- **Handlers**: HTTP request handling and validation
- **Repositories**: Data access layer implementations
- **Client**: External API interactions
- **Router**: Middleware and routing setup
- **App**: Application lifecycle management

Run tests locally with:

```bash
make test              # Run all tests
```

## 🚦 CI/CD

The project uses GitHub Actions for continuous integration:

- **Build**: Verifies Docker image builds successfully
- **Lint**: Checks code formatting and runs `go vet`
- **Test**: Runs unit tests
- **Deploy**: Deploys to production on successful builds

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
