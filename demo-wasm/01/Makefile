all: static/main.wasm static/wasm_exec.js
	goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'
static/wasm_exec.js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" static
static/main.wasm: main.go
	GOOS=js GOARCH=wasm go build -o static/main.wasm main.go
clean:
	rm static/main.wasm