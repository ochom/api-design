SHELL=/bin/bash

tidy:
	@echo "Tidying up..."
	@go mod tidy

dev:
	@echo "starting dev..."
	@source env.sh && air

test:
	@echo "Running tests..."
	@go test ./...

generate:
	@echo "Generating code..."
	@go generate ./...

lint:
	@echo "Linting ..."
	@golangci-lint run --timeout 5m
