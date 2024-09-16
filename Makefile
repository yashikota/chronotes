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
	GITHUB_TOKEN=$(GITHUB_TOKEN) go build ./...
	GITHUB_TOKEN=$(GITHUB_TOKEN) go test -v ./...


fmt: ## Format Go code
	go fmt ./...

dev: ## Run development environment
	docker compose up --build

build: ## Build Docker image
	docker build -t chronotes .

# API commands
.PHONY: api-lint bundle docs tsp-install tsp

api-lint: bundle ## Lint API documentation
	docker run --rm -v ${PWD}:/spec redocly/cli lint --config docs/api/redoc.yaml docs/api/bundled.yaml

bundle: ## Bundle OpenAPI specification
	docker run --rm -v ${PWD}:/spec redocly/cli bundle docs/api/openapi.yaml -o docs/api/bundled.yaml

docs: ## Generate API documentation
	docker run --rm -v ${PWD}:/spec redocly/cli build-docs docs/api/openapi.yaml --o docs/api/redoc.html

tsp: ## Genrate Open API from Typespec
	docker run --rm -v ${PWD}:/wd --workdir="/wd" -t azsdkengsys.azurecr.io/typespec compile docs/tsp

# Docker commands
.PHONY: docker-lint

docker-lint: ## Lint Dockerfile
	docker run --rm -i hadolint/hadolint < Dockerfile

# Test comannds
# See: https://note.com/reality_eng/n/n338cc671968e
.PHONY: coverage ## Generate test coverage

coverage:
	GITHUB_TOKEN=$(GITHUB_TOKEN) go test ./... -short -v -covermode=count -coverprofile=coverage.out | tee test_output.txt
	GITHUB_TOKEN=$(GITHUB_TOKEN) go tool cover -func=coverage.out | awk '/total:/ {print "| **" $$1 "** | **" $$3 "** |"}' | tee coverage.txt
	cat test_output.txt | grep 'ok.*coverage' | awk '{sub("github.com/your/package/", "", $$2); print "| " $$2 " | " $$5 " |"}' | tee -a coverage.txt
	echo "## Test Coverage Report" > coverage_with_header.txt
	echo "| Package           | Coverage |" >> coverage_with_header.txt
	echo "|-------------------|----------|" >> coverage_with_header.txt
	cat coverage.txt >> coverage_with_header.txt
	mv coverage_with_header.txt coverage.txt
	cat coverage.txt

# Help command
.PHONY: help
help: ## Display this help message
	@echo "Usage: make ${YELLOW}<target>${NC}"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  ${YELLOW}%-15s${NC} %s\n", $$1, $$2}' $(MAKEFILE_LIST)
