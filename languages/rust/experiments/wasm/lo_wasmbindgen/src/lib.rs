use wasm_bindgen::prelude::*;
#[wasm_bindgen]
pub fn get_information() -> String {
    "Hello there from WasmBindgen".to_string()
}
