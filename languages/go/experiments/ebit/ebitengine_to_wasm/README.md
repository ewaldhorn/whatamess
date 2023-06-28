# Experimental WASM deployment

- Compile: env GOOS=js GOARCH=wasm go build -o ebittowasm.wasm ebittowasm`
- Then: `cp $(go env GOROOT)/misc/wasm/wasm_exec.js .`
- Should work when the index.html is opened via a server.
