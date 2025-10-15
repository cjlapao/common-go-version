GO ?= go
GOLANGCI_LINT ?= golangci-lint
GOSEC ?= gosec
GOSEC_OUTPUT ?= gosec.sarif

.PHONY: all deps build test lint scan ci tools

all: build

# Download Go module dependencies
deps:
	$(GO) mod download

# Compile all packages
build:
	$(GO) build ./...

# Run the full test suite
test:
	$(GO) test ./...

# Run golangci-lint with default settings
lint:
	@command -v $(GOLANGCI_LINT) >/dev/null 2>&1 || { \
		echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	}
	$(GOLANGCI_LINT) run

# Run gosec static analysis matching CI configuration
scan:
	@command -v $(GOSEC) >/dev/null 2>&1 || { \
		echo "gosec not found. Install with: go install github.com/securego/gosec/v2/cmd/gosec@latest"; \
		exit 1; \
	}
	$(GOSEC) -fmt sarif -out $(GOSEC_OUTPUT) ./...

# Convenience target to mirror the GitHub PR validation pipeline
ci: deps build test lint scan

# Install commonly used tooling locally
TOOLS_BIN ?= $(shell $(GO) env GOBIN)
ifeq ($(TOOLS_BIN),)
TOOLS_BIN := $(shell $(GO) env GOPATH)/bin
endif

.PHONY: install-tools install-lint install-gosec clean

install-tools: install-lint install-gosec

install-lint:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

install-gosec:
	$(GO) install github.com/securego/gosec/v2/cmd/gosec@latest

# Remove build artifacts
clean:
	rm -f $(GOSEC_OUTPUT)
