#[macro_use]
mod browser;

use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::rc::Rc;
use std::sync::Mutex;
use wasm_bindgen::prelude::*;
use wasm_bindgen::JsCast;

#[derive(Serialize, Deserialize)]
struct Rect {
    x: u16,
    y: u16,
    w: u16,
    h: u16,
}

#[derive(Serialize, Deserialize)]
struct Cell {
    frame: Rect,
}

#[derive(Serialize, Deserialize)]
struct Sheet {
    frames: HashMap<String, Cell>,
}

// This is like the `main` function, except for JavaScript.
#[wasm_bindgen(start)]
pub fn main_js() -> Result<(), JsValue> {
    // This provides better error messages in debug mode.
    // It's disabled in release mode so it doesn't bloat up the file size.
    // #[cfg(debug_assertions)]
    console_error_panic_hook::set_once();

    let context = browser::context().expect("Could not get canvas context");

    // all of this is to wait for the image to be loaded before displaying it
    // lots of back and forth between rust and javascript
    browser::spawn_local(async move {
        // read sprite sheet data for rhb
        let json = browser::fetch_json("assets/sprite_sheets/rhb.json")
            .await
            .expect("Could not fetch rhb.json");
        let sheet: Sheet = serde_wasm_bindgen::from_value(json)
            .expect("Oh my, could not parse the json structure!"); // TODO: Rather handle the potential error

        // -----------------------------------------------------------------------------------------
        // draw using sprite sheet starts
        let (success_tx, success_rx) = futures::channel::oneshot::channel::<Result<(), JsValue>>();
        let success_tx = Rc::new(Mutex::new(Some(success_tx)));
        let error_tx = Rc::clone(&success_tx);
        let image = web_sys::HtmlImageElement::new().unwrap();
        image.set_src("assets/sprite_sheets/rhb.png");

        // wait for the image to load, let us know when it has
        let callback = Closure::once(move || {
            if let Some(success_tx) = success_tx.lock().ok().and_then(|mut opt| opt.take()) {
                _ = success_tx.send(Ok(()));
            }
        });
        let error_callback = Closure::once(move |err| {
            if let Some(error_tx) = error_tx.lock().ok().and_then(|mut opt| opt.take()) {
                _ = error_tx.send(Err(err));
            }
        });
        image.set_onload(Some(callback.as_ref().unchecked_ref()));
        image.set_onerror(Some(error_callback.as_ref().unchecked_ref()));

        _ = success_rx.await;
        let mut frame = 0;
        let interval_callback = Closure::wrap(Box::new(move || {
            frame = frame + 1;
            if frame > 8 {
                frame = 1;
            }
            let frame_name = format!("Run ({}).png", frame);
            let sprite = sheet.frames.get(&frame_name).expect("Cell not found");
            context.clear_rect(0.0, 0.0, 600.0, 600.0);
            _ = context
                .draw_image_with_html_image_element_and_sw_and_sh_and_dx_and_dy_and_dw_and_dh(
                    &image,
                    sprite.frame.x.into(),
                    sprite.frame.y.into(),
                    sprite.frame.w.into(),
                    sprite.frame.h.into(),
                    200.0,
                    320.0,
                    sprite.frame.w.into(),
                    sprite.frame.h.into(),
                );
        }) as Box<dyn FnMut()>);
        _ = browser::window()
            .unwrap()
            .set_interval_with_callback_and_timeout_and_arguments_0(
                interval_callback.as_ref().unchecked_ref(),
                50,
            );
        interval_callback.forget();

        // draw using sprite sheet ends
        // -----------------------------------------------------------------------------------------
    });

    Ok(())
}

/**
 * Responsible for loading a json file, for example to be able to use sprite sheets.
 */
async fn fetch_json(json_path: &str) -> Result<JsValue, JsValue> {
    let window = web_sys::window().unwrap();
    let resp_value = wasm_bindgen_futures::JsFuture::from(window.fetch_with_str(json_path)).await?;
    let resp: web_sys::Response = resp_value.dyn_into()?;
    wasm_bindgen_futures::JsFuture::from(resp.json()?).await
}
