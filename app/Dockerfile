# ── Stage 1: Build ────────────────────────────────────────────────────────────
FROM golang:1.21-alpine AS builder

# Install git for go module downloads
RUN apk add --no-cache git ca-certificates

WORKDIR /build

# Cache dependencies separately from source
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /app/server .

# ── Stage 2: Runtime ──────────────────────────────────────────────────────────
FROM scratch

# Copy CA certs for HTTPS if needed
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy compiled binary
COPY --from=builder /app/server /server

# Non-root user (numeric UID for scratch compatibility)
USER 65532:65532

# Environment variables with defaults
ENV PORT=8080 \
    APP_ENV=production

EXPOSE 8080

ENTRYPOINT ["/server"]
