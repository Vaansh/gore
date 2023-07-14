# syntax=docker/dockerfile:1

# Use the official Golang base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy all files into the image
COPY . .

# Run go mod
RUN go mod download

# Expose ports
EXPOSE 8000

# Run Go program, just like locally
ENTRYPOINT ["go","run","cmd/worker/main.go"]
