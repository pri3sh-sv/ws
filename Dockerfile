# Use official Golang image as builder
FROM golang:1.22 as builder
# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary (main.go is in cmd/)
RUN go build -o server ./cmd/main.go

# Use a minimal runtime image for deployment
FROM debian:bookworm-slim

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Expose the port Cloud Run will use
ENV PORT=8080
EXPOSE 8080

# Command to run the application
CMD ["./server"]
