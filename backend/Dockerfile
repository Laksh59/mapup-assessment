# Use official Go image
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy go files
COPY . .

# Download dependencies
RUN go mod tidy

# Build the application
RUN go build -o main .

# Expose port (change if your app uses different port)
EXPOSE 8080

# Run the app
CMD ["./main"]
