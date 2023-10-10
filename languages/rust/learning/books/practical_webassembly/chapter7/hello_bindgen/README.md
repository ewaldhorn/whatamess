# Rust, Cargo, WebAssembly and Wasm-bindgen

Build: `cargo build --target=wasm32-unknown-unknown`

Bind: `wasm-bindgen target/wasm32-unknown-unknown/debug/hello_bindgen.wasm --out-dir .`
