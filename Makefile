init:
	@echo "Initializing..."
	@go mod tidy

dev:
	@echo "Running..."
	@air

swag:
	@echo "Generating Swagger docs..."
	@swag init --parseDependency -g cmd/api/main.go

build:
	@echo "Building..."
	@go build -o tmp/main cmd/api/main.go

test:
	@echo "Testing..."
	@go test ./... -v

docker-up:
	@echo "Starting Docker services..."
	@docker compose up -d

docker-down:
	@echo "Stopping Docker services..."
	@docker compose down
