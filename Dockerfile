# Use the official Go image for development
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Install fresh (a tool to automatically reload your application when source files change)
RUN go install github.com/pilu/fresh@latest

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run fresh
CMD ["fresh"]
