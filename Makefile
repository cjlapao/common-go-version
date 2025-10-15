GO ?= go

TOOLS_BIN ?= $(shell $(GO) env GOBIN)
ifeq ($(TOOLS_BIN),)
TOOLS_BIN := $(shell $(GO) env GOPATH)/bin
endif

GOLANGCI_LINT ?= golangci-lint
GOSEC ?= gosec

GOLANGCI_LINT_PATH := $(shell command -v $(GOLANGCI_LINT) 2>/dev/null)
ifeq ($(GOLANGCI_LINT_PATH),)
GOLANGCI_LINT_PATH := $(TOOLS_BIN)/golangci-lint
endif

GOSEC_PATH := $(shell command -v $(GOSEC) 2>/dev/null)
ifeq ($(GOSEC_PATH),)
GOSEC_PATH := $(TOOLS_BIN)/gosec
endif

GOSEC_OUTPUT ?= gosec.sarif

.PHONY: all deps build test lint scan ci tools install-tools install-lint install-gosec clean

all: build

deps:
	$(GO) mod download

build:
	$(GO) build ./...

test:
	$(GO) test ./...

lint:
	@if [ ! -x "$(GOLANGCI_LINT_PATH)" ]; then \
		echo "golangci-lint not found. Run 'make install-lint' and ensure $(TOOLS_BIN) is on your PATH."; \
		exit 1; \
	fi
	"$(GOLANGCI_LINT_PATH)" run

scan:
	@if [ ! -x "$(GOSEC_PATH)" ]; then \
		echo "gosec not found. Run 'make install-gosec' and ensure $(TOOLS_BIN) is on your PATH."; \
		exit 1; \
	fi
	"$(GOSEC_PATH)" -fmt sarif -out $(GOSEC_OUTPUT) ./...

ci: deps build test lint scan

tools: install-tools

install-tools: install-lint install-gosec

install-lint:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

install-gosec:
	$(GO) install github.com/securego/gosec/v2/cmd/gosec@latest

clean:
	rm -f $(GOSEC_OUTPUT)
