# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (required for some Go modules)
RUN apk --no-cache add git

# Copy go.mod and go.sum first for better Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary for Linux
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/api/main.go

# Production stage
FROM alpine:3.19

WORKDIR /app

# Install ca-certificates for outbound HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Copy the built binary from builder
COPY --from=builder /app/main .

# Copy config files
COPY configs/config.yaml ./configs/config.yaml

# Copy public assets
COPY public/ ./public/

# Expose port (default 9000, overridden by $PORT env var)
EXPOSE 9000

# Run the application
CMD ["./main"]

