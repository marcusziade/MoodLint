# Name the build stage and use the official Go image to create a build environment
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /src

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the correct directory
COPY . /src

# Build the application
RUN go build -o /app ./src

# Use a minimal image for the runtime
FROM alpine:latest

# Copy the compiled application from the build stage
COPY --from=build /app /moodlint

# Execute the application
ENTRYPOINT ["/moodlint"]
