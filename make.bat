@echo off
if "%~1"=="" goto help
if "%~1"=="init" goto init
if "%~1"=="dev" goto dev
if "%~1"=="swag" goto swag
if "%~1"=="build" goto build
if "%~1"=="test" goto test
if "%~1"=="docker-up" goto docker-up
if "%~1"=="docker-down" goto docker-down

:init
echo Initializing...
go mod tidy
goto :eof

:dev
echo Running dev server...
where air >nul 2>nul
if %errorlevel% neq 0 (
    echo [WARNING] 'air' was not found in your PATH.
    echo Running with standard 'go run cmd/api/main.go' instead...
    go run cmd/api/main.go
) else (
    air
)
goto :eof

:swag
echo Generating Swagger docs...
where swag >nul 2>nul
if %errorlevel% neq 0 (
    echo [WARNING] 'swag' was not found in your PATH.
    echo Installing swag...
    go install github.com/swaggo/swag/cmd/swag@latest
)
swag init --parseDependency -g cmd/api/main.go
goto :eof

:build
echo Building...
go build -o tmp/main.exe cmd/api/main.go
goto :eof

:test
echo Testing...
go test ./... -v
goto :eof

:docker-up
echo Starting Docker services...
docker compose up -d
goto :eof

:docker-down
echo Stopping Docker services...
docker compose down
goto :eof

:help
echo Usage: make [init^|dev^|swag^|build^|test^|docker-up^|docker-down]
goto :eof
