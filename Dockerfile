# Use a minimal Go runtime base image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the entire source code to the working directory
COPY . .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the main.go application
CMD ["go", "run", "main.go"]
