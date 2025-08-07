# Stage 1: Build the Go binary
FROM golang:1.24-alpine3.22 AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /docker-gs-ping


FROM alpine:3.19
WORKDIR /app


COPY --from=builder /docker-gs-ping /app/docker-gs-ping 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ 
COPY --chown=1000:1000 calc/ ./calc/
COPY --chown=1000:1000 data/ ./data/


RUN adduser -D -u 1000 appuser && \
    chown -R appuser:appuser /app
USER appuser

ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/docker-gs-ping"]