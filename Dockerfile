# Step 1: Build the Go binary
FROM golang:1.21-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Ensure dependencies are fetched, and add the GOPROXY variable for reliable fetching
ENV GOPROXY=https://proxy.golang.org,direct

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy -v

# Copy the entire project to the container
COPY . .

# Build the Go app
RUN go build -o main .

# Step 2: Create a smaller image to run the Go binary
FROM alpine:latest

# Install necessary libraries (e.g., for SQLite, Postgres, etc., if needed)
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your app will run on
EXPOSE 8080

# Run the Go binary
CMD ["./main"]
