.PHONY: help
help:
	@echo "Available commands:"
	@grep -E '^.PHONY:\s+[^#]*(##.*)?' $(MAKEFILE_LIST) | sort | cut -d ':' -f 2 | awk 'BEGIN {FS = "#"}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Align struct fields (field comments will be removed)
.PHONY: align
align:
	@go run golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest -fix ./...

# Build Docker image
.PHONY: build
build:
	@docker build -t goapp:latest .

# Remove Docker image
.PHONY: clean
clean:
	@docker rmi goapp:latest

# Generate Go code
.PHONY: gen
gen:
	@go generate ./...

# Lint Go code
.PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --enable gosec

# Run pre-commit checks
.PHONY: precommit
precommit:
	@go fmt ./...
	@go mod tidy
	@make lint
	@make test

# Start local development in container
.PHONY: start
start:
	@docker-compose up -d

# Stop the container
.PHONY: stop
stop:
	@docker-compose down

# Run unit tests
.PHONY: test
test:
	@go test -v -cover -race ./...

# Update and tidy dependencies
.PHONY: update
update:
	@go get -u
	@go mod tidy
