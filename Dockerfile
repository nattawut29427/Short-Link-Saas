# Production stage
FROM alpine:3.19

WORKDIR /app

# Install ca-certificates in case the app needs to make outbound HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the pre-built binary directly from the host machine
COPY main .

# Copy config directory (requires configs/config.yaml)
COPY configs/config.yaml ./configs/config.yaml

# Expose port (default 9000, will be dynamically overridden by $PORT on GCP Cloud Run)
EXPOSE 9000

# Run the application
CMD ["./main"]
