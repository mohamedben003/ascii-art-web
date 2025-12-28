# This builds and runs your ASCII Art web application

# Start with a Go image that has Go already installed
FROM golang:1.23.3-alpine

# Install bash (optional but helpful for beginners)
RUN apk add --no-cache bash

# Add labels (metadata) - tells others about your image
LABEL description="ASCII Art Web Generator"
LABEL version="1.0"

# Create a folder inside the container to put our app
WORKDIR /app

# Copy go.mod file (tells Go what packages we need)
COPY go.mod ./

# Download Go packages
RUN go mod download

# Copy all our code into the container
COPY . .

# Build our Go application
RUN go build -o ascii-art-web .

# Tell Docker our app uses port 8080 (just metadata)
EXPOSE 8080

# When container starts, run our application
CMD ["./ascii-art-web"]