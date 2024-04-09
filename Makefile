export GOBIN := $(shell pwd)/bin

.PHONY: build
build: ## build go files from root of project
	go build -o $(GOBIN)/

.PHONY: help
help:
	@echo "Available targets:"
	@awk '/^[a-zA-Z0-9_-]+:[[:space:]]*## / {split($$0, a, /:[[:space:]]*## /); printf "  \033[34m%-15s\033[0m - %s\n", a[1], a[2]}' $(MAKEFILE_LIST)