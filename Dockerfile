# Multi-stage build
# Stage 1: Build binary
FROM golang:latest AS builder

WORKDIR /app

COPY . .

# Compile Go binary
RUN go build -o main .

# Stage 2: Final Image
FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .

# Expose port yang digunakan oleh web service
EXPOSE 8080

# Command untuk menjalankan aplikasi ketika container dijalankan
CMD ["./main"]
