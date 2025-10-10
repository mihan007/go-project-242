test:
	go test -v ./...

deps:
	go mod tidy

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

lint:
	golangci-lint run ./...