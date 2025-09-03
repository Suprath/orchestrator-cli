# FILE: orchestrator-cli/Dockerfile
# --- Builder Stage ---
FROM golang:1.24-alpine AS builder
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
    docker-compose

# Install Terraform
RUN apk add --no-cache curl unzip && \
    curl -LO https://releases.hashicorp.com/terraform/1.5.7/terraform_1.5.7_linux_amd64.zip && \
    unzip terraform_1.5.7_linux_amd64.zip -d /usr/local/bin && \
    rm terraform_1.5.7_linux_amd64.zip

# Install GitHub CLI
RUN apk add --no-cache curl tar && \
    curl -LO https://github.com/cli/cli/releases/download/v2.39.1/gh_2.39.1_linux_amd64.tar.gz && \
    tar -xzf gh_2.39.1_linux_amd64.tar.gz -C /tmp && \
    mv /tmp/gh_2.39.1_linux_amd64/bin/gh /usr/local/bin && \
    rm -rf /tmp/gh_2.39.1_linux_amd64 && \
    rm gh_2.39.1_linux_amd64.tar.gz

# Copy the compiled orchestrator binary from the builder stage
COPY --from=builder /orchestrator /usr/local/bin/orchestrator

# Copy the entire templates directory into a known location in the image
COPY internal/templates /templates

# Set the working directory to where the user's code will be mounted
WORKDIR /app

ENTRYPOINT ["orchestrator"]
CMD ["--help"]