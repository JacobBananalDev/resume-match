# ---- Build Stage ----
FROM golang:1.25-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the API binary
RUN go build -o resumematch ./cmd/api


# ---- Runtime Stage ----
FROM alpine:latest

WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /app/resumematch .

# Expose API port
EXPOSE 8080

# Run the API
CMD ["./resumematch"]