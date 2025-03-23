# Use the official Golang image as a base
FROM golang:1.22 as builder

WORKDIR /app

# Initialize Go module
RUN go mod init student-restapi

# Copy the source code (before downloading dependencies)
COPY . .

# Generate go.mod and go.sum by resolving dependencies
RUN go mod tidy

# Build the application
RUN go build -o main .

# Use a minimal base image for the final build
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./main"]
