VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
TARGET  ?= app

build: ## Build the binary
	go build -ldflags "-X main.version=$(VERSION)" -o $(TARGET) .

run: build ## Build and run
	./$(TARGET)

install: ## Install the binary
	go install -ldflags "-X main.version=$(VERSION)" .

fmt: ## Format source code
	go fmt ./...

test: ## Run tests
	go test -v ./...

vet: ## Run go vet
	go vet ./...

staticcheck: ## Run static analysis
	@command -v staticcheck >/dev/null 2>&1 || { echo "staticcheck is not installed."; echo "Install it with:"; echo "  go install honnef.co/go/tools/cmd/staticcheck@latest"; exit 1; }
	staticcheck ./...

check: vet staticcheck test ## Run all checks

clean: ## Remove artefacts
	rm -f $(TARGET)

help: ## Show this help
	@grep -E '^[a-zA-Z0-9_-]+:.*##' $(MAKEFILE_LIST) | sort | \
		awk -F ':.*## ' '{printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2}'


.PHONY: build run install fmt test vet staticcheck check clean help
.DEFAULT_GOAL := help
