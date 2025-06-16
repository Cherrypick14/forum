# First stage: Build the Go application
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install necessary dependencies
RUN apk add --no-cache git

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o forum main.go

# Second stage: Create a lightweight runtime image
FROM alpine:latest

# Set the maintainer label
LABEL maintainer="allangithaiga5@gmail.com"
LABEL description="forum container"

# Install only necessary dependencies
RUN apk add --no-cache sqlite

# Set working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/forum .

# Expose the necessary port
EXPOSE 8080

# Define the command to run the application
CMD ["./forum"]
