.PHONY: build test install clean

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=terraform-docs-extended
VERSION=0.1.0
BUILD_FLAGS=-ldflags "-X main.Version=$(VERSION)"

# Main build target
build:
	$(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_NAME) -v

# Install to GOPATH/bin
install:
	$(GOINSTALL) $(BUILD_FLAGS) -v

# Run tests
test:
	$(GOTEST) -v ./...

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run the application with default parameters
run:
	./$(BINARY_NAME)

# Build for multiple OS targets
build-all: build-linux build-windows build-darwin

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_NAME)_linux_amd64 -v

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_NAME)_windows_amd64.exe -v

build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_NAME)_darwin_amd64 -v

# Run a demo on the sample module
demo:
	@echo "Running demo on sample module..."
	@mkdir -p demo/output
	./$(BINARY_NAME) --path=sample_module --out=demo/output/README.md
	@echo "Demo output written to demo/output/README.md"