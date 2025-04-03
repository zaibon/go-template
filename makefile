# Makefile
.PHONY: all build lint test clean

# Default target
all: build lint test

# Build the service
build:
	@echo "Building..."
	go build -v -o ./bin/service ./cmd/service/main.go

# Lint the code
lint:
	@echo "Linting..."
	go tool golangci-lint run ./...

# Run the tests
test:
	@echo "Testing..."
	go test -v ./...

# Clean up build artifacts
clean:
	@echo "Cleaning..."
	rm -f ./bin/service
	go clean