# Use the official Go image as the parent image
FROM golang:1.22 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY api ./api
COPY resources ./resources
COPY go.* .

# Build the Go app
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-w -s" -o velov-widget-api ./api/main.go

# Use a minimal Alpine image as the base image for the final image
FROM scratch

# Set the working directory to /app
WORKDIR /app

# Copy the built executable from the build image to the final image
COPY --from=builder /app/velov-widget-api ./
COPY --from=builder /app/resources ./resources
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose port 8080 for the application
EXPOSE 8080

# Run the application
CMD ["./velov-widget-api"]
