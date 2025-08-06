# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app

# Copy and download dependencies first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /docker-gs-ping

# Stage 2: Minimal runtime image
FROM alpine:3.19
WORKDIR /app

# Copy pre-built binary and assets
COPY --from=builder /docker-gs-ping /app/docker-gs-ping 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ 
COPY --chown=1000:1000 templates/ ./templates/
COPY --chown=1000:1000 static/ ./static/
COPY --chown=1000:1000 calc/ ./calc/
COPY --chown=1000:1000 data/ ./data/

# Security best practices
RUN adduser -D -u 1000 appuser && \
    chown -R appuser:appuser /app
USER appuser

ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/docker-gs-ping"]