.PHONY: all check check-format format test lint fix clean test-bazel

all: format test fix

check: check-format test lint

check-format:
	bazel test //tools/format/...:all

format:
	bazel run //tools/format

test: test-bazel

lint: lint-go

fix: fix-go

test-bazel:
	bazel test //...

# Go lint and fix isn't integrated with bazel yet. Nogo is a good option.
.PHONY: lint-go fix-go

lint-go: go.sum
	golangci-lint run ./...

fix-go: go.sum
	golangci-lint run --fix ./...

go.sum: go.mod
	go mod tidy



clean:
	bazel clean
