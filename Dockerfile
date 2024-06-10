# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Efrizal"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .
COPY .env .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["go", "run", "main.go"]