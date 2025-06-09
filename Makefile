.PHONY: help up down logs shell build clean restart format deps-install test test-short test-race test-specific lint check

help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

up: ## Start the application
	docker compose up -d
	@echo "Application running at http://localhost:8888"

down: ## Stop the application
	docker compose down

logs: ## Show logs (usage: make logs [service=app])
	docker compose logs -f $(or $(service),app)

shell: ## Access application shell
	docker compose exec app sh

build: ## Build Docker images
	docker compose build

clean: ## Clean up containers, volumes, and images
	docker compose down -v --rmi all

restart: down up ## Restart the application

format: ## Format Go code
	@echo "Formatting code..."
	docker compose exec app sh -c "go fmt ./..."
	@echo "Code formatting complete!"

deps-install: ## Install/update dependencies
	@echo "Installing dependencies..."
	docker compose exec app sh -c "go mod tidy"
	@echo "Dependencies installed!"

test: ## Run all tests
	@echo "Running tests..."
	docker compose exec app go test ./... -v


test-short: ## Run short tests only
	@echo "Running short tests..."
	docker compose exec app go test ./... -short

test-specific: ## Run specific test (usage: make test-specific TEST=TestName)
	@echo "Running specific test: $(TEST)"
	docker compose exec app go test ./... -run $(TEST) -v

lint: ## Run linters (gofmt and go vet)
	@echo "Running linters..."
	@echo "Checking formatting..."
	@docker compose exec app sh -c 'if [ -n "$$(gofmt -l .)" ]; then echo "Formatting issues found:"; gofmt -l .; exit 1; fi'
	@echo "Running go vet..."
	docker compose exec app go vet ./...
	@echo "Linting complete!"

check: lint test ## Run all checks (lint + test)
