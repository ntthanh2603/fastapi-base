# Makefile for DRUNK CLI tool

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=drunk-cli
MAIN_RUN = ./cmd/drunk/main.go


# Default target is to build the binary
all: build

start:
	@echo "Starting development server..."
	$(GOCMD) run $(MAIN_RUN)

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
swag:
	@echo "Generating Swagger documentation..."
	swag init -g $(MAIN_RUN) -o ./cmd/swag/docs
# Cross-platform builds
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-linux-amd64 -v

build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-windows-amd64.exe -v

build-mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-darwin-amd64 -v

# Build for all platforms
build-all: build-linux build-windows build-mac

# Install to GOPATH/bin
install:
	$(GOBUILD) -o $(GOPATH)/bin/$(BINARY_NAME) -v

.PHONY: all build clean build-linux build-windows build-mac build-all install 