# Default target
.DEFAULT_GOAL := help

# Colors for terminal output
YELLOW := \033[1;33m
NC := \033[0m # No Color

# Go commands
.PHONY: lint test fmt dev build

lint: ## Run Go linters
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go vet ./...
	staticcheck ./...

test: ## Run Go tests
	go build ./...
	go test -v ./...

fmt: ## Format Go code
	go fmt ./...

dev: ## Run development environment
	docker compose up --build

build: ## Build Docker image
	docker build -t 58hack .

# API commands
.PHONY: api-lint bundle docs

api-lint: bundle ## Lint API documentation
	docker run --rm -v ${PWD}:/spec redocly/cli lint --config docs/api/redoc.yaml docs/api/bundled.yaml

bundle: ## Bundle OpenAPI specification
	docker run --rm -v ${PWD}:/spec redocly/cli bundle docs/api/openapi.yaml -o docs/api/bundled.yaml

docs: ## Generate API documentation
	docker run --rm -v ${PWD}:/spec redocly/cli build-docs docs/api/openapi.yaml --o docs/api/redoc.html

# Help command
.PHONY: help
help: ## Display this help message
	@echo "Usage: make ${YELLOW}<target>${NC}"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  ${YELLOW}%-15s${NC} %s\n", $$1, $$2}' $(MAKEFILE_LIST)
