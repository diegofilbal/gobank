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

package:
	@zip -r $(APP_NAME).zip api bank go.mod go.sum utils.go

clean:
	@rm -f $(APP_NAME)
