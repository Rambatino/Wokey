.PHONY: build

setup:
	@npm install

build:
	@npm run build

serve:
	@go run *.go

storybook:
	@npm run storybook
