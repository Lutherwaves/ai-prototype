# Makefile
.PHONY: build run test lint docker clean deps

build:
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

lint:
	golangci-lint run

docker:
	docker build -t imagine-proto .

clean:
	rm -rf bin/

deps:
	go mod download

.PHONY: dev
dev:
	ENV=dev go run cmd/api/main.go

.PHONY: mock
mock:
	mockgen -source=internal/llm/provider/types.go -destination=internal/llm/provider/mocks/provider_mock.go

.PHONY: tidy
tidy:
	go mod tidy
