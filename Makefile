include .env
export MASTER_DB_URL
export EC2_DB_URL

# Application settings
APP_NAME=deep_chained_service
BACKEND_SERVICE_ENTRYPOINT=./cmd/start/main.go
CORE_DB_MIGRATION=./db/migrations

# Go command and tools
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOVET=$(GOCMD) vet

# Binary paths
BINARY_PATH=./bin/$(APP_NAME)
BINARY_UNIX=$(BINARY_PATH)_unix

# Default target
all: test build

# Build targets
build:
	$(GOBUILD) -o $(BINARY_PATH) -v $(BACKEND_SERVICE_ENTRYPOINT)

build-linux:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(BINARY_UNIX) $(BACKEND_SERVICE_ENTRYPOINT)

build-htmx: tailwind-build templ-generate
	$(GOBUILD) -ldflags "-X main.Environment=production" -o $(BINARY_PATH) $(BACKEND_SERVICE_ENTRYPOINT)

# Run and development targets
run: build
	$(BINARY_PATH)

dev:
	$(GOBUILD) -o ./tmp/$(APP_NAME) $(BACKEND_SERVICE_ENTRYPOINT) && air

# Test and quality targets
test:
	$(GOTEST) -race -v -timeout 30s ./...

vet:
	$(GOVET) ./...

staticcheck:
	staticcheck ./...

# Dependency management
deps:
	$(GOMOD) tidy
	$(GOMOD) verify

# Clean up
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH) $(BINARY_UNIX)

# Database migration targets
up:
	@goose -dir=$(CORE_DB_MIGRATION) postgres $$MASTER_DB_URL up

down:
	@goose -dir=$(CORE_DB_MIGRATION) postgres $$MASTER_DB_URL down

ec2-db-up:
	@goose -dir=$(CORE_DB_MIGRATION) postgres $$EC2_DB_URL up

ec2-db-down:
	@goose -dir=$(CORE_DB_MIGRATION) postgres $$EC2_DB_URL down

create-migration:
	@read -p "Enter migration name: " name; \
	goose -dir $(CORE_DB_MIGRATION) create $$name sql

# Documentation and API
swag:
	@swag fmt
	@swag init -g $(BACKEND_SERVICE_ENTRYPOINT)

# Frontend tooling
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

templ-generate:
	templ generate

templ-watch:
	templ generate --watch

# Declare phony targets
.PHONY: all build build-linux build-htmx run dev test vet staticcheck deps clean \
        up down ec2-db-up ec2-db-down create-migration swag \
        tailwind-watch tailwind-build templ-generate templ-watch