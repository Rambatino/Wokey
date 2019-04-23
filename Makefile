setup:
	@npm install

build-prod:
	@npm run build

serve:
	# watch and rebuild assets
	# reload webpage on change
	@(sleep 2 && open http://localhost:3000)&
	@go run *.go

storybook:
	@npm run storybook
