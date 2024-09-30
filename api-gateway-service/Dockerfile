# Build stage
FROM golang:1.23

WORKDIR /app

# for refresh (hot reloac)
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download


# Copy the source code
COPY . .

CMD ["air", "-c", ".air.toml"]
