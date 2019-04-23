setup:
	@npm install

build-prod:
	@npm run build

serve:
	@(sleep 2 && open http://localhost:3000)&
	@go run *.go

storybook:
	@npm run storybook
