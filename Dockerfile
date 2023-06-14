# Use an official Go runtime as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /build

# Copy the Go modules manifest and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application source code
COPY . .

# Build the Go application inside the container
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o authentigo

# Set the entry point command for the container
CMD ["./authentigo"]