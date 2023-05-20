# Use an official Go runtime as a parent image
FROM golang:1.20-alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Install build-base package
RUN apk add --no-cache build-base

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Build the app
RUN make build


RUN echo "Running the app..."

# Start the app
CMD make run

