# Stage 1 : Build
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod first for better layer caching
COPY app/go.mod ./

# Download dependencies
RUN go mod download

# Copy application source
COPY app/ .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Stage 2 : Runtime
FROM alpine

WORKDIR /app

# Copy only the compiled binary
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
