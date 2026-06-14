# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency files FIRST (important)
COPY go.mod go.sum ./

RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o server main.go


# Runtime stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
