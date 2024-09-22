# Build stage
FROM golang:1.23

WORKDIR /app

# for refresh (hot reloac)
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download


# Copy the source code
COPY . .

# Build the application
# RUN CGO_ENABLED=0 GOOS=linux go build -o tmp/blogs ./cmd/api

# # Final stage
# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# # Copy the binary from builder
# COPY --from=builder /app/blogs .

# Command to run
# CMD ["./blogs"]

CMD ["air", "-c", ".air.toml"]
