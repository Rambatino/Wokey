GO := go
GLIDE := glide
COMPOSE := docker-compose
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

.PHONY: build

setup:
	@npm install
	@$(GLIDE) install

build:
	@npm run build

serve:
	@go run *.go

storybook:
	@npm run storybook

test:
	@$(GO) test ./...

run:
	@$(GO) run *.go

run-docker:
	@$(COMPOSE) up --build
