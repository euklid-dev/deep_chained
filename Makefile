include .env
export DB_URL
export EC2_DB_URL

DB_MIGRATION= ./db/migrations
# Go command
GOCMD=go
# Build, clean, test, and get commands
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
# Binary names, adjusted to place the output in ./bin/
BINARY_NAME=deep_chained
BINARY_PATH=./bin/$(BINARY_NAME)
BINARY_UNIX=$(BINARY_PATH)_unix
BACKEND_SERVICE_ENTRYPOINT=./cmd/backend_service/main.go

# Default to run tests and then build
all: test build

# Compile the binary to bin/ directory
build:
	$(GOBUILD) -o $(BINARY_PATH) -v $(BACKEND_SERVICE_ENTRYPOINT)

# Run tests across the project
test:
	$(GOTEST) -v ./...

# Clean up the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)
	rm -f $(BINARY_UNIX)

# Build and run the project
run: build
	$(BINARY_PATH)

# Handle dependencies using Go modules
deps:
	$(GOMOD) tidy
	$(GOMOD) verify

# Cross compilation for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(BINARY_PATH) $(BACKEND_SERVICE_ENTRYPOINT)

up:
	@goose -dir=$(DB_MIGRATION) postgres $$DB_URL up

down:
	@goose -dir=$(DB_MIGRATION) postgres $$DB_URL down

ec2-db-up:
	goose -dir=$(DB_MIGRATION) postgres $$EC2_DB_URL up

ec2-db-down:
	@goose -dir=$(DB_MIGRATION) postgres $$EC2_DB_URL down

