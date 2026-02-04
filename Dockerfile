# Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git for fetch
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. 
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o bot main.go

# Production Stage
FROM alpine:latest  

WORKDIR /root/

# Install ca-certificates
RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary from the previous stage
COPY --from=builder /app/bot .
COPY --from=builder /app/locales ./locales
COPY --from=builder /app/.env .

# Expose port (if webhook used, otherwise not strictly needed but good practice)
EXPOSE 8080

# Command to run the executable
CMD ["./bot"]
