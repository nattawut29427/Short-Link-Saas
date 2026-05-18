# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application as a statically linked binary (optimized for speed and low RAM)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main cmd/api/main.go

# Production stage
FROM alpine:3.19

WORKDIR /app

# Install ca-certificates in case the app needs to make outbound HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/main .

# Copy config directory (requires configs/config.yaml)
COPY --from=builder /app/configs/config.yaml ./configs/config.yaml

# Expose port (default 9000, will be dynamically overridden by $PORT on GCP Cloud Run)
EXPOSE 9000

# Run the application
CMD ["./main"]
