# Use the official Golang image as the base image
FROM golang:1.24.0-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main -ldflags="-s -w" .

# Set environment variables from .env file
COPY .env .env
RUN export $(cat .env | grep -v "#" | xargs)

# Command to run the executable
CMD ["./main"]