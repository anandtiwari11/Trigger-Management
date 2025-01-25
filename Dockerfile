# Step 1: Build stage
FROM golang:1.20-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o event-trigger .

# Step 2: Final stage - Running the application
FROM alpine:latest

# Install necessary certificates for PostgreSQL connection
RUN apk --no-cache add ca-certificates

# Set the working directory in the final container
WORKDIR /root/

# Copy the compiled Go binary from the builder image
COPY --from=builder /app/event-trigger .

# Copy the .env file
COPY .env .

# Expose the port the app will run on
EXPOSE 8080

# Run the application
CMD ["./event-trigger"]
