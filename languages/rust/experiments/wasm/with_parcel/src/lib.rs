use wasm_bindgen::prelude::*;
#[wasm_bindgen]
pub fn get_version_information() -> String {
    "Parcel Served Version".to_string()
}
