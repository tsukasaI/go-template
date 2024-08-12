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

## oapi-codegen
.PHONY: oapi-codegen
oapi-codegen:
	@oapi-codegen --config=oapi-codegen-config.yaml ./docs/openapi.yaml

## build
.PHONY: build
build:
	@go build ./src/cmd/main.go
