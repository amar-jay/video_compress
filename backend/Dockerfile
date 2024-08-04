# Use the official Go image as a base image
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download the Go modules. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o server ./cmd/server.go

# Use a minimal base image to reduce the final image size
FROM debian:bullseye-slim

# Install FFmpeg
RUN apt-get update && apt-get install -y ffmpeg && rm -rf /var/lib/apt/lists/*

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary and other necessary files from the builder stage
COPY --from=builder /app/server .

# Expose port 5000 to the outside world
EXPOSE 5000

# Command to run the executable
CMD ["./main"]

