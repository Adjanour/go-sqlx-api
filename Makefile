.PHONY: help build run test clean fmt vet lint docker-build docker-run

# Variables
BINARY_NAME=api-server
MAIN_PATH=./cmd/api
DOCKER_IMAGE=go-sqlx-api

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

run: ## Run the application
	@echo "Running..."
	@go run $(MAIN_PATH)

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

tidy: ## Tidy go modules
	@echo "Tidying go modules..."
	@go mod tidy

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	@docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE)

docker-compose-up: ## Start services with docker-compose
	@echo "Starting services with docker-compose..."
	@docker-compose up -d

docker-compose-down: ## Stop services with docker-compose
	@echo "Stopping services with docker-compose..."
	@docker-compose down

install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install golang.org/x/tools/cmd/goimports@latest

all: clean fmt vet build test ## Clean, format, vet, build and test
