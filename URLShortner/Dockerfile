# Start from the official golang image, based on Alpine Linux for a small image size
FROM golang:alpine AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod ./

# Download dependencies using go mod
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Start a new stage from scratch
FROM alpine AS final

# Set environment variables for the Go binary
ENV PORT=8080

# Copy the compiled Go binary from the 'build' stage
COPY --from=build /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE $PORT

# Command to run the executable
CMD ["/app/main"]
