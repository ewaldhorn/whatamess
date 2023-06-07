# Basic WASM in Go

See <https://www.freecodecamp.org/news/webassembly-with-golang-is-fun-b243c0e34f02/> for the
original article I sort of followed.

## Building

Just run the `buildwasm.sh` script (it runs `GOOS=js GOARCH=wasm go build -o main.wasm`)

## JavaScript superglue

`cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`
