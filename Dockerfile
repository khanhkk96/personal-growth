# Stage 1: Build the Go app
FROM golang:1.23.9-alpine AS builder

# Install dependencies for swag CLI and Go
RUN apk add --no-cache git curl

WORKDIR /app

# Install swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Run swag init before building
RUN /go/bin/swag init

# Build the Go app
RUN go build -o pgw .

# Stage 2: Run the Go app in a minimal image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app .
COPY --from=builder /app/docs ./docs

# Expose the application port (optional)
EXPOSE 8000

# Run the binary
CMD ["./pgw"]

