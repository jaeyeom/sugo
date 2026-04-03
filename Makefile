.PHONY: all check check-format check-generated format build test lint fix clean \
       test-bazel tidy lint-go fix-go coverage coverage-html coverage-report \
       clean-coverage verify-golangci-config

# --- Full local workflow ---

all: tidy format fix test build

# --- CI-friendly checks (no mutation) ---

check: check-format lint test build

# --- Tidy ---

tidy: MODULE.bazel.lock go.sum

MODULE.bazel.lock: MODULE.bazel go.sum
	bazel mod tidy

# --- Format ---

check-format:
	bazel test //tools/format/...:all

format:
	bazel run //tools/format

# --- Build ---

build:
	bazel build //...

# --- Test ---

test: test-bazel

test-bazel:
	bazel test //...

# --- Lint / Fix ---
# Go lint and fix isn't integrated with bazel yet. Nogo is a good option.

lint: lint-go

lint-go: go.sum
	golangci-lint run ./...

# fix depends on format so both don't mutate files concurrently under make -j.
fix: format
fix: fix-go

fix-go: go.sum
	golangci-lint run --fix ./...

verify-golangci-config:
	golangci-lint config verify

# --- Generated / dirty-tree check ---

check-generated:
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "ERROR: working tree is dirty after generation/formatting:"; \
		git diff --stat; \
		exit 1; \
	fi

# --- Coverage ---

COVERAGE_DIR := coverage
COVERAGE_PROFILE := $(COVERAGE_DIR)/coverage.out

coverage:
	@mkdir -p $(COVERAGE_DIR)
	go test -coverprofile=$(COVERAGE_PROFILE) ./...

coverage-html: coverage
	go tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_DIR)/coverage.html

coverage-report: coverage
	go tool cover -func=$(COVERAGE_PROFILE)

clean-coverage:
	rm -rf $(COVERAGE_DIR)

# --- Dependency management ---

go.sum: go.mod
	bazel run @rules_go//go -- mod tidy

# --- Clean ---

clean: clean-coverage
	bazel clean
