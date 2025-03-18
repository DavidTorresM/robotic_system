#Stage 1 - compile
# Use the official Golang image as the base image
FROM golang:1.24.0-alpine AS base 

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go mod download

# Build the Go app
RUN go build -o main -ldflags="-s -w" .

#Stage 2 - the production environment
FROM alpine:3.12.0 AS final

WORKDIR /app

COPY --from=base /app/main /app/main
# Set environment variables from .env file
COPY .env .env
RUN export $(cat .env | grep -v "#" | xargs)

# Command to run the executable
CMD ["./main"]