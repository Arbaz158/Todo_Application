# Use the official Golang image as the base image
FROM golang:1.22.3-alpine3.20

# Set the working directory to /app
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the application will run on
EXPOSE 8080

# Set the command to run the application
CMD ["./main"]

