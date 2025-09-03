# FILE: orchestrator-cli/Dockerfile
# --- Builder Stage ---
FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build the orchestrator binary for a linux/amd64 environment
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /orchestrator .

# --- Final Image ---
FROM alpine:latest

# Install ALL runtime dependencies that our Go program calls
RUN apk add --no-cache \
    git \
    docker-cli \
    docker-compose \
    terraform \
    github-cli # This provides the 'gh' command

# Copy the compiled orchestrator binary from the builder stage
COPY --from=builder /orchestrator /usr/local/bin/orchestrator

# Copy the entire templates directory into a known location in the image
COPY templates /templates

# Set the working directory to where the user's code will be mounted
WORKDIR /app

ENTRYPOINT ["orchestrator"]
CMD ["--help"]
