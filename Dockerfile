# Use the official Go image as the parent image
FROM golang:1.21.4-alpine3.18 AS build

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY api ./api
COPY resources ./resources
COPY go.* .

# Build the Go app
RUN go build -o velov-widget-api ./api/main.go

# Use a minimal Alpine image as the base image for the final image
FROM alpine:3.18

# Set the working directory to /app
WORKDIR /app

# Copy the built executable from the build image to the final image
COPY --from=build /app/velov-widget-api ./
COPY --from=build /app/resources ./resources

# Expose port 8080 for the application
EXPOSE 8080

# Run the application
CMD ["./velov-widget-api"]
