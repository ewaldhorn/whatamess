env GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o static/compiled.wasm main.go
