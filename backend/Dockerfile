# Backend Dockerfile - Go with multi-stage build
# Produces a minimal production image

# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git (needed for some Go modules)
RUN apk add --no-cache git

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary with optimizations
# CGO_ENABLED=0 for static binary, GOOS=linux for Linux target
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server ./monolith.go

# Stage 2: Minimal runtime image
FROM alpine:3.19 AS production

WORKDIR /app

# Install CA certificates for HTTPS calls (Stripe, WorkOS, etc.)
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user for security
RUN adduser -D -g '' appuser

# Copy binary from builder
COPY --from=builder /app/server .

# Change ownership
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose the backend port
EXPOSE 8173

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8173/health || exit 1

# Run the server
CMD ["./server"]
