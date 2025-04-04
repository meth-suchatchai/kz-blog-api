# Use a minimal base image for Go applications
#FROM golang:1.20 AS builder

# Set the working directory inside the container
#WORKDIR /app

# Copy the Go command source code into the container
#COPY . .

# download dependecies
#RUN go mod tidy

# Build the Go command
#RUN GOOS=linux GOARCH=amd64 go build -o kz_api-linux-x64 ./bootstrap/*.go

# Add debugging step to confirm the binary exists
#RUN ls -l /app/kz_api-linux-x64

# Create a lightweight final image
FROM alpine

# Install tzdata package to manage timezones
RUN apk --no-cache add tzdata

# Set the timezone (for example, Asia/Bangkok)
ENV TZ=Asia/Bangkok

WORKDIR /app

# Copy the compiled Go command from the builder stage
COPY kz_api-linux-x64 /app

# Add debugging to ensure the binary is copied
RUN ls -l /app/kz_api-linux-x64

# Expose the port that the Go command listens on
EXPOSE 3100

# Run the Go command
CMD ["./kz_api-linux-x64", "start"]