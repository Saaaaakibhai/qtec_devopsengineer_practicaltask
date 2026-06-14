# Use official Go image (build stage)
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o server main.go


# Minimal production image
FROM alpine:latest

WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Run app
CMD ["./server"]
