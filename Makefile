build:
	@go build -o studybear ./app

wasm:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm ./app

run: build wasm
	./studybear