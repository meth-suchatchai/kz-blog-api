# Use a minimal base image for Go applications
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go command source code into the container
COPY . .

# Build the Go command
RUN go build -o api ./bootstrap/*.go

# Create a lightweight final image
FROM alpine

# Copy the compiled Go command from the builder stage
COPY --from=builder /app/api /usr/local/bin/api

# Expose the port that the Go command listens on
EXPOSE 8080

# Run the Go command
CMD ["api"]