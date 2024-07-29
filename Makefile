## Format Go files
.PHONY: fmt
fmt:
	@go fmt ./...

## Tidy modfile
.PHONY: tidy
tidy:
	@go mod tidy -v

## test: run all tests
.PHONY: test
test:
	@go test -v -race ./...

## lint
.PHONY: lint
lint:
	@golangci-lint run --fix

## build
.PHONY: build
build:
	@go build ./src/cli/main.go
