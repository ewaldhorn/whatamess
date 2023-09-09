use std::rc::Rc;
use std::sync::Mutex;

use rand::prelude::*;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
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
    #[cfg(debug_assertions)]
    console_error_panic_hook::set_once();

    let window = web_sys::window().unwrap();
    let document = window.document().unwrap();
    let canvas = document
        .get_element_by_id("canvas")
        .unwrap()
        .dyn_into::<web_sys::HtmlCanvasElement>()
        .unwrap();
    let context = canvas
        .get_context("2d")
        .unwrap()
        .unwrap()
        .dyn_into::<web_sys::CanvasRenderingContext2d>()
        .unwrap();

    let mut rng = thread_rng();
    let first_color = (
        rng.gen_range(0..255),
        rng.gen_range(0..255),
        rng.gen_range(0..255),
    );

    sierpinski(
        &context,
        [(300.0, 0.0), (0.0, 600.0), (600.0, 600.0)],
        first_color,
        8,
    );

    // all of this is to wait for the image to be loaded before displaying it
    // lots of back and forth between rust and javascript
    wasm_bindgen_futures::spawn_local(async move {
        let (success_tx, success_rx) = futures::channel::oneshot::channel::<Result<(), JsValue>>();
        let success_tx = Rc::new(Mutex::new(Some(success_tx)));
        let error_tx = Rc::clone(&success_tx);
        let image = web_sys::HtmlImageElement::new().unwrap();
        image.set_src("assets/images/rhb/Idle (1).png");

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
        _ = context.draw_image_with_html_image_element(&image, 0.0, 0.0);

        // read sprite sheet data for rhb
        let json = fetch_json("assets/sprite_sheets/rhb.json").await.unwrap();
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
        let sprite = sheet.frames.get("Run (1).png").expect("Cell not found");
        _ = context.draw_image_with_html_image_element_and_sw_and_sh_and_dx_and_dy_and_dw_and_dh(
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
        // draw using sprite sheet ends
        // -----------------------------------------------------------------------------------------
    });

    Ok(())
}

fn draw_triangle(
    context: &web_sys::CanvasRenderingContext2d,
    points: [(f64, f64); 3],
    color: (u8, u8, u8),
) {
    let color_str = format!("rgb({}, {}, {})", color.0, color.1, color.2);
    context.set_fill_style(&wasm_bindgen::JsValue::from_str(&color_str));

    let [top, left, right] = points;
    context.move_to(top.0, top.1);
    context.begin_path();
    context.line_to(left.0, left.1);
    context.line_to(right.0, right.1);
    context.line_to(top.0, top.1);
    context.close_path();
    context.stroke();
    context.fill();
}

fn sierpinski(
    context: &web_sys::CanvasRenderingContext2d,
    points: [(f64, f64); 3],
    color: (u8, u8, u8),
    depth: u8,
) {
    draw_triangle(&context, points, color);
    let depth = depth - 1;
    let [top, left, right] = points;
    if depth > 0 {
        let mut rng = thread_rng();
        let next_color = (
            rng.gen_range(0..255),
            rng.gen_range(0..255),
            rng.gen_range(0..255),
        );

        let left_middle = midpoint(top, left);
        let right_middle = midpoint(top, right);
        let bottom_middle = midpoint(left, right);
        sierpinski(
            &context,
            [top, left_middle, right_middle],
            next_color,
            depth,
        );
        sierpinski(
            &context,
            [left_middle, left, bottom_middle],
            next_color,
            depth,
        );
        sierpinski(
            &context,
            [right_middle, bottom_middle, right],
            next_color,
            depth,
        );
    }
}

fn midpoint(point_1: (f64, f64), point_2: (f64, f64)) -> (f64, f64) {
    ((point_1.0 + point_2.0) / 2.0, (point_1.1 + point_2.1) / 2.0)
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
