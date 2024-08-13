# Start from the official Golang base image
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN make build-linux

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin/automate_ai_service .
COPY --from=builder /app/.env .
COPY --from=builder /app/convo_gpt_response.json .


# Expose port 8080 to the outside world
EXPOSE 2718

# Command to run the executable
CMD ["./deep_chained_service"]