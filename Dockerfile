# Use the official Golang image to create a build artifact.
FROM golang:1.17 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o stock-analysis ./cmd

# Start a new stage from scratch
FROM alpine:latest  

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/stock-analysis /stock-analysis

# Expose port 8085 to the outside world
EXPOSE 8085

# Command to run the executable
CMD ["/stock-analysis"]