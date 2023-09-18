# Wasmbindgen and Rust

Build: `cargo build --target=wasm32-unknown-unknown`

Bind: `wasm-bindgen target/wasm32-unknown-unknown/debug/lo_wasmbindgen.wasm --out-dir .`
