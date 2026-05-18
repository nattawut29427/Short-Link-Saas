param (
    [string]$Target = "help"
)

switch ($Target) {
    "init" {
        Write-Host "Initializing..." -ForegroundColor Cyan
        go mod tidy
    }
    "dev" {
        Write-Host "Running dev server..." -ForegroundColor Cyan
        if (Get-Command air -ErrorAction SilentlyContinue) {
            air
        } else {
            Write-Warning "'air' was not found in your PATH. Falling back to standard 'go run cmd/api/main.go'..."
            go run cmd/api/main.go
        }
    }
    "swag" {
        Write-Host "Generating Swagger docs..." -ForegroundColor Cyan
        if (-not (Get-Command swag -ErrorAction SilentlyContinue)) {
            Write-Warning "'swag' was not found in your PATH. Installing..."
            go install github.com/swaggo/swag/cmd/swag@latest
        }
        swag init --parseDependency -g cmd/api/main.go
    }
    "build" {
        Write-Host "Building..." -ForegroundColor Cyan
        go build -o tmp/main.exe cmd/api/main.go
    }
    "test" {
        Write-Host "Testing..." -ForegroundColor Cyan
        go test ./... -v
    }
    "docker-up" {
        Write-Host "Starting Docker services..." -ForegroundColor Cyan
        docker compose up -d
    }
    "docker-down" {
        Write-Host "Stopping Docker services..." -ForegroundColor Cyan
        docker compose down
    }
    default {
        Write-Host "Usage: .\make.ps1 [init|dev|swag|build|test|docker-up|docker-down]" -ForegroundColor Yellow
    }
}
