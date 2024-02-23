FROM golang:1.22.0

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and install Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the application
CMD ["./main"]
