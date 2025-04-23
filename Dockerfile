FROM golang:1.18-alpine AS builder

# Install git and gcc, needed for Go dependencies
RUN apk add --no-cache git gcc musl-dev

# Install terraform-docs
RUN wget -O- https://github.com/terraform-docs/terraform-docs/releases/download/v0.16.0/terraform-docs-v0.16.0-linux-amd64.tar.gz | tar xz -C /tmp && \
    mv /tmp/terraform-docs /usr/local/bin/

# Set working directory
WORKDIR /app

# Copy Go module files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /terraform-docs-extended

# Create a minimal runtime image
FROM alpine:3.16

# Install terraform-docs in the runtime image
COPY --from=builder /usr/local/bin/terraform-docs /usr/local/bin/terraform-docs

# Copy the compiled binary
COPY --from=builder /terraform-docs-extended /usr/local/bin/terraform-docs-extended

# Set entrypoint
ENTRYPOINT ["terraform-docs-extended"]

# Default command (can be overridden)
CMD ["--help"]