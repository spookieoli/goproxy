# how to use this file
# docker build -t goproxy .
#run as a container daemon
# docker run -d -p <ip>:<port>:<containerport> goproxy

# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /goproxy

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/goproxy", "-localAddr", "0.0.0.0:6379", "-remoteAddr", "keydb:6379"]