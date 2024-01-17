NAME = app
IMAGE_REPO = registry.cn-hangzhou.aliyuncs.com/toodo/brain

MODULE ?= github.com/MR5356/go-template

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

VERSION ?= $(shell git symbolic-ref --short -q HEAD)-$(shell git rev-parse --short HEAD)

.DEFAULT_GOAL := help
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

docs: ## Generate docs
	swag init
	swag fmt

build: clean deps docs ## Build the project
	go build -ldflags "-s -w" -o bin/$(NAME)

binary: clean deps docs ## Build binary
	go build -ldflags "-s -w" -o bin/$(NAME)-${GOOS}-${GOARCH}

test: deps docs ## Execute tests
	go test ./...

deps: docs ## Install dependencies using go get
	go get -d -v -t ./...

clean: ## Remove building artifacts
	rm -rf bin

image: ## Build and push docker image
	docker buildx build --platform linux/arm64,linux/amd64 -t $(IMAGE_REPO):$(VERSION) . --push
