build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run: build
	bin/hexlet-path-size

.PHONY: build run

test:
	go mod tidy
	go test -v ./tests

install:
	go install

lint:
	golangci-lint run ./...