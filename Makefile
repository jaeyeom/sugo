.PHONY: all check format check-format test lint fix

all: format test fix

check: check-format test lint

format:
	gofmt -w .

check-format:
	test -z "$$(gofmt -l .)"

test:
	go test -v ./...

lint:
	golangci-lint run ./...

fix:
	golangci-lint run --fix ./...
