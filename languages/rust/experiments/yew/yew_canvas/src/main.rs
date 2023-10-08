use gloo::console::log;
use rand::{thread_rng, Rng};
use wasm_bindgen::{prelude::Closure, JsCast, JsValue};
use wasm_bindgen_futures::JsFuture;
use web_sys::{
    window, Blob, CanvasRenderingContext2d, HtmlCanvasElement, ImageBitmap, Request, RequestInit,
    Response, Window,
};
use yew::prelude::*;

pub enum Msg {
    FetchOk(ImageBitmap),
    FetchFail(FetchImageError),
    Render,
}

struct Particle {
    x: f64,
    y: f64,
    velocity: f64,
    size: f64,
    speed: f64,
    max_height: usize,
}

impl Particle {
    fn new(width: usize, height: usize) -> Self {
        let mut rando = thread_rng();
        let y = 0.0;
        let x = rando.gen_range(0f64..width as f64);
        let velocity = rando.gen_range(0.5..10.0);
        let size = rando.gen_range(1.0..2.0);
        let speed = rando.gen_range(1.0..3.0);
        let max_height = height;

        Self {
            x,
            y,
            velocity,
            size,
            speed,
            max_height,
        }
    }

    fn update(&mut self, map: &[Vec<(u8, u8, u8, f64)>]) {
        let x = self.x as usize;
        let y = self.y as usize;
        let (_r, _g, _b, brightness) = map[y][x];

        self.speed = brightness;
        let delta_y = 2.0 - self.speed + self.velocity;
        self.y += delta_y;
        if self.y >= self.max_height as f64 {
            self.y = 0f64;
        }
    }

    fn render(&self, ctx: &CanvasRenderingContext2d, map: &[Vec<(u8, u8, u8, f64)>]) {
        ctx.begin_path();
        let x = self.x as usize;
        let y = self.y as usize;
        let (r, g, b, _b) = map[y][x];
        let js_rgb = format!("rgb({r}, {g}, {b})", r = r, g = g, b = b);
        ctx.set_fill_style(&JsValue::from_str(js_rgb.as_ref()));
        ctx.arc(
            self.x,
            self.y,
            self.size.into(),
            0.0,
            std::f64::consts::PI * 2.0,
        )
        .unwrap();
        ctx.fill();
    }
}

pub struct FetchImageError {
    err: JsValue,
}

impl From<JsValue> for FetchImageError {
    fn from(value: JsValue) -> Self {
        Self { err: value }
    }
}

pub async fn fetch_image(file_path: &str) -> Result<ImageBitmap, FetchImageError> {
    let mut opts = RequestInit::new();
    opts.method("GET");

    let request = Request::new_with_str_and_init(file_path, &opts)?;
    let window: Window = window().unwrap();
    let resp_js_value = JsFuture::from(window.fetch_with_request(&request)).await?;
    let resp: Response = resp_js_value.dyn_into()?;
    let blob: Blob = JsFuture::from(resp.blob()?).await?.dyn_into()?;
    let image_bmp_promise = window.create_image_bitmap_with_blob(&blob)?;
    Ok(JsFuture::from(image_bmp_promise).await?.dyn_into()?)
}

pub fn relative_brightness(r: f64, g: f64, b: f64) -> f64 {
    js_sys::Math::sqrt((r * r * 0.229) + (g * g * 0.587) + (b * b * 0.114)) / 100.0
}

struct AnimationCanvas {
    canvas: NodeRef,
    particles: Vec<Particle>,
    callback: Closure<dyn FnMut()>,
    brightness_map: Vec<Vec<(u8, u8, u8, f64)>>,
}

impl Component for AnimationCanvas {
    type Message = Msg;
    type Properties = ();

    fn create(ctx: &Context<Self>) -> Self {
        ctx.link().send_future(async {
            match fetch_image("content/ewald_devconf_2023.jpeg").await {
                Ok(image) => Msg::FetchOk(image),
                Err(err) => Msg::FetchFail(err),
            }
        });
        let component_context = ctx.link().clone();
        let callback = Closure::wrap(
            Box::new(move || component_context.send_message(Msg::Render)) as Box<dyn FnMut()>,
        );

        Self {
            canvas: NodeRef::default(),
            particles: vec![],
            callback,
            brightness_map: vec![],
        }
    }

    fn update(&mut self, ctx: &Context<Self>, msg: Self::Message) -> bool {
        match msg {
            Msg::FetchOk(image) => {
                let width: usize = image.width().try_into().unwrap();
                let height: usize = image.height().try_into().unwrap();
                let canvas: HtmlCanvasElement = self.canvas.cast().unwrap();
                canvas.set_width(width.try_into().unwrap());
                canvas.set_height(height.try_into().unwrap());
                self.particles = (0..5000).map(|_| Particle::new(width, height)).collect();

                let ctxx: CanvasRenderingContext2d =
                    canvas.get_context("2d").unwrap().unwrap().unchecked_into();

                ctxx.draw_image_with_image_bitmap(&image, 0.0, 0.0).unwrap();
                let image_data = ctxx
                    .get_image_data(0.0, 0.0, width as f64, height as f64)
                    .unwrap();
                ctxx.clear_rect(0.0, 0.0, width as f64, height as f64);
                let buffer = (*image_data.data()).clone();
                self.brightness_map = Vec::new();
                for y in 0usize..height {
                    let mut brightness_row = Vec::new();
                    brightness_row.reserve(width);
                    for x in 0usize..width {
                        let red = buffer[(y * 4usize * width) + (x * 4)];
                        let green = buffer[(y * 4usize * width) + (x * 4 + 1)];
                        let blue = buffer[(y * 4usize * width) + (x * 4 + 2)];
                        let brightness = relative_brightness(red as f64, green as f64, blue as f64);
                        brightness_row.push((red, green, blue, brightness));
                    }
                    self.brightness_map.push(brightness_row);
                }

                ctx.link().send_message(Msg::Render);
                true
            }
            Msg::FetchFail(e) => {
                log!(e.err);
                true
            }
            Msg::Render => {
                self.render();
                true
            }
        }
    }

    fn view(&self, _ctx: &Context<Self>) -> Html {
        html! {
            <div>
                <canvas id="canvas" ref={self.canvas.clone()}></canvas>
            </div>
        }
    }
}

impl AnimationCanvas {
    fn render(&mut self) {
        let canvas: HtmlCanvasElement = self.canvas.cast().unwrap();
        let ctx: CanvasRenderingContext2d =
            canvas.get_context("2d").unwrap().unwrap().unchecked_into();
        ctx.set_global_alpha(0.05);
        ctx.set_fill_style(&JsValue::from("rgb(0,0,0)"));
        ctx.fill_rect(0.0, 0.0, canvas.width().into(), canvas.height().into());
        ctx.set_global_alpha(0.5);
        let map = &self.brightness_map;

        self.particles.iter_mut().for_each(|particle| {
            particle.update(map);

            let x = particle.x as usize;
            let y = particle.y as usize;
            let (_r, _g, _b, brightness) = map[y][x];
            ctx.set_global_alpha(brightness * 0.3); // higher velocity higher alpha
            particle.render(&ctx, map);
        });

        window()
            .unwrap()
            .request_animation_frame(self.callback.as_ref().unchecked_ref())
            .unwrap();
    }
}

#[function_component]
fn App() -> Html {
    html! {
        <div>
            <AnimationCanvas/>
        </div>
    }
}

fn main() {
    yew::Renderer::<App>::new().render();
}
