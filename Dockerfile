# # Use an official Go runtime as a parent image
# FROM golang:1.20

# # Set the working directory inside the container
# WORKDIR /app

# # Copy go.mod and go.sum files first to cache dependencies
# COPY go.mod go.sum ./

# # Download and cache dependencies
# RUN go mod download

# # Copy the rest of the application code
# COPY . .

# # Build the Go application
# RUN go build -o todo-app .

# # Expose the application port (optional, depends on your app's port)
# EXPOSE 8080

# # Run the application
# CMD ["./todo-app"]
