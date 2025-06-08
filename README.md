# My Portfolio Website

[![CI](https://github.com/benidevo/website/actions/workflows/ci.yaml/badge.svg)](https://github.com/benidevo/website/actions/workflows/ci.yaml)
[![Go Version](https://img.shields.io/badge/Go-1.24.2-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A modern, modular portfolio website built with Go.

## 🎯 Overview

A portfolio website designed to showcase:

- **Projects**: Key engineering work and technical achievements
- **Experience**: Professional background and skills
- **Writing**: Technical articles and insights

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
- **Build Tools**: ESBuild

## 📁 Project Structure

``` plaintext
website/
├── cmd/
│   └── website/          # Application entry point
├── internal/             # Private application code
│   ├── app/             # Application lifecycle management
│   ├── config/          # Configuration and logging
│   ├── router/          # HTTP routing setup
│   ├── handlers/        # Request handlers
│   ├── middleware/      # HTTP middleware
│   ├── services/        # Business logic
│   └── models/          # Data structures
├── web/                 # Frontend assets
│   ├── templates/       # HTML templates
│   ├── static/          # CSS, JS, images
│   └── assets/          # Source assets
├── docker/              # Docker configurations
├── docs/                # Project documentation
└── scripts/             # Build and deployment scripts
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

| Command | Description |
|---------|-------------|
| `make up` | Start the application in development mode |
| `make down` | Stop the application |
| `make logs` | View application logs |
| `make shell` | Access container shell |
| `make build` | Build Docker image |
| `make format` | Format Go code |
| `make deps-install` | Install/update dependencies |
| `make restart` | Restart the application |
| `make clean` | Remove all containers and images |

## 🔧 Configuration

The application uses environment variables for configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `IS_DEVELOPMENT` | Development mode | `true` |
| `LOG_LEVEL` | Logging level | `info` |

## 🏗️ Architecture

The application follows a clean architecture pattern with:

- **Separation of Concerns**: Clear boundaries between HTTP handling, business logic, and data
- **Dependency Injection**: Services are injected where needed
- **Graceful Shutdown**: Proper signal handling for zero-downtime deployments
- **Structured Logging**: Consistent, queryable logs with Zerolog
- **Error Handling**: Global error recovery and user-friendly error pages

## 🚦 CI/CD

The project uses GitHub Actions for continuous integration:

- **Build**: Verifies Docker image builds successfully
- **Lint**: Checks code formatting and runs `go vet`
- **Test**: Runs unit and integration tests

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
