.PHONY: build test test-verbose test-coverage clean help run-example

# Build variables
BINARY_NAME=cidrmgr
GO=go
GOFLAGS=-v

help:
	@echo "cidrmgr - Fast CIDR range management tool"
	@echo ""
	@echo "Available targets:"
	@echo "  make build           - Build the cidrmgr binary"
	@echo "  make test            - Run all tests"
	@echo "  make test-verbose    - Run tests with verbose output"
	@echo "  make test-coverage   - Run tests with coverage report"
	@echo "  make clean           - Clean build artifacts"
	@echo "  make run-example     - Build and run example"
	@echo ""

build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME)
	@echo "✓ Build complete: ./$(BINARY_NAME)"

test:
	@echo "Running tests..."
	$(GO) test -v ./...

test-verbose:
	@echo "Running tests with verbose output..."
	$(GO) test -v -race ./...

test-coverage:
	@echo "Running tests with coverage..."
	$(GO) test -v -coverprofile=coverage.out ./...
	@echo "✓ Coverage report generated: coverage.out"
	@echo "  View report with: go tool cover -html=coverage.out"

benchmark:
	@echo "Running benchmarks..."
	$(GO) test -bench=. -benchmem ./...

clean:
	@echo "Cleaning up..."
	$(GO) clean -testcache
	rm -f $(BINARY_NAME)
	rm -f coverage.out
	@echo "✓ Cleanup complete"

run-example: build
	@echo "Creating example input file..."
	@mkdir -p examples
	@cat > examples/ranges.txt << 'EOF'
# Private IPv4 ranges
192.168.0.0/24
192.168.1.0/24
192.168.0.128/25

# RFC1918 ranges
10.0.0.0/8
172.16.0.0/12

# IPv6 ranges
2001:db8::/32
fc00::/7

# Duplicate (will be removed)
192.168.0.0/24
EOF
	@echo "✓ Example file created: examples/ranges.txt"
	@echo ""
	@echo "Running merge..."
	./$(BINARY_NAME) merge -i examples/ranges.txt -o examples/merged.txt
	@echo "✓ Merged output: examples/merged.txt"
	@echo ""
	@echo "Result:"
	@cat examples/merged.txt

install: build
	@echo "Installing cidrmgr..."
	$(GO) install
	@echo "✓ Installation complete"

all: clean build test
	@echo "✓ Build and tests complete"

.DEFAULT_GOAL := help
