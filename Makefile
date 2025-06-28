.PHONY: all fmt vet test test-cover run-matrix-demo run-vector-demo clean

all: fmt vet test

# Format check
fmt:
	@echo "Checking formatting..."
	@fmtResult=$$(gofmt -l .); \
	if [ -n "$$fmtResult" ]; then \
		echo "These files need formatting:"; \
		echo "$$fmtResult"; \
		exit 1; \
	else \
		echo "All files are properly formatted."; \
	fi

# Vet static analysis
vet:
	@echo "Running go vet..."
	go vet ./...

# Run tests
test:
	@echo "Running unit tests..."
	go test ./...

# Run tests with coverage and report
test-cover:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out

# Run matrix demo
run-matrix-demo:
	@echo "Running matrix demo..."
	go run examples/matrix_demo/main.go

# Run vector demo
run-vector-demo:
	@echo "Running vector demo..."
	go run examples/vector_demo/main.go

# Clean coverage reports
clean:
	rm -f matrix/coverage.out vector/coverage.out coverage.out
