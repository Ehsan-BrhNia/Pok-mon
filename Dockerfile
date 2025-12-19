# Build stage
FROM golang:1.25-alpine AS builder

# Install git and ca-certificates for Go modules
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o app .

# Final stage
FROM alpine:latest
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app .

# Expose port that Gin uses
EXPOSE 8081

# Run the app
CMD ["./app"]
