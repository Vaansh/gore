# syntax=docker/dockerfile:1

# Use the official Golang base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy all files into the image
COPY . .

# Run go mod
RUN go mod download

# Expose ports
EXPOSE 8080

# Run Go program, just like locally
ENTRYPOINT ["go","run","cmd/api/main.go"]
