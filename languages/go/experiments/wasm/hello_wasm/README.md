# Basic WASM in Go

## Building

Just run the `buildwasm.sh` script (it runs `GOOS=js GOARCH=wasm go build -o main.wasm`)

## JavaScript superglue

`cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`
