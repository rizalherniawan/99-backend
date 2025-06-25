# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder

WORKDIR /app

# Download dependencies early to cache them
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build a statically linked binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

# --- Final minimal image ---
FROM alpine:latest

WORKDIR /app

# Copy only the compiled binary and migrations
COPY --from=builder /app/server .
COPY --from=builder /app/migrations ./migrations


# Default command
CMD ["./server"]
