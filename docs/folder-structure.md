# Folder Structure for the Website Project

This document outlines the folder structure for the website project, which is designed to be modular and maintainable. Each directory serves a specific purpose, promoting separation of concerns and ease of navigation.

``` plaintext
website/
  ├── cmd/
  │   └── website/
  │       └── main.go          # Application entry point
  ├── internal/
  │   ├── app/                 # Main application struct
  │   │   └── app.go           # App initialization and setup
  │   ├── config/
  │   │   └── config.go        # Configuration management
  │   ├── handlers/            # HTTP handlers (controllers)
  │   │   ├── home.go
  │   │   ├── projects.go
  │   │   ├── about.go
  │   │   └── contact.go
  │   ├── services/            # Business logic
  │   │   ├── github.go        # GitHub API integration
  │   │   └── cache.go         # Caching service
  │   ├── models/              # Data structures
  │   │   ├── project.go
  │   │   └── contact.go
  │   ├── middleware/          # HTTP middleware
  │   │   ├── logger.go
  │   │   └── cache.go
  │   └── router/              # Route definitions
  │       └── router.go
  ├── web/                     # Frontend assets
  │   ├── templates/           # HTML templates
  │   │   ├── layouts/
  │   │   │   └── base.html
  │   │   ├── pages/
  │   │   │   ├── home.html
  │   │   │   ├── projects.html
  │   │   │   └── about.html
  │   │   └── components/      # Reusable components
  │   │       ├── header.html
  │   │       └── project-card.html
  │   ├── static/              # Static assets
  │   │   ├── css/
  │   │   ├── js/
  │   │   └── img/
  │   └── assets/              # Source assets (if using build tools)
  ├── pkg/                     # Shared packages (if needed)
  │   └── utils/
  ├── scripts/                 # Build/deployment scripts
  ├── docker/
  │   └── dev/
  │       └── Dockerfile
  ├── docs/
  │   └── PRD.md
  ├── .env.example
  ├── docker-compose.yaml
  ├── Makefile
  ├── go.mod
  └── go.sum
```
