.PHONY: help up down logs shell build clean

help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

up:
	docker compose up -d
	@echo "Application running at http://localhost:8888"

down:
	docker compose down

logs:
	docker compose logs -f $(or $(service),app)

shell:
	docker compose exec app sh

build:
	docker compose build

clean:
	docker compose down -v --rmi all

restart: down up

format:
	@echo "Formatting code..."
	docker compose exec app sh -c "go fmt ./..."
	@echo "Code formatting complete!"

deps-install:
	@echo "Installing dependencies..."
	docker compose exec app sh -c "go mod tidy"
	@echo "Dependencies installed!"

