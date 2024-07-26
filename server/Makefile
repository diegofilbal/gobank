.PHONY: default run build test lint clean

APP_NAME=gobank-dev

default: run

run:
	@cd api && go run main.go

dependencies:
	@go mod tidy

build:
	@go build -o $(APP_NAME) api/main.go

test:
	@cd bank && go test -v

lint:
	@golangci-lint run

clean:
	@rm -f $(APP_NAME)
