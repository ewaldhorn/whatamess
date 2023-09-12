use image::{ImageBuffer, Rgb};

const WIDTH: u32 = 800;
const HEIGHT: u32 = 600;

fn main() {
    let scale_x = 3.0 / WIDTH as f64;
    let scale_y = 3.0 / HEIGHT as f64;

    let mut img = ImageBuffer::new(WIDTH, HEIGHT);

    for (x, y, pixel) in img.enumerate_pixels_mut() {
        *pixel = Rgb([100_u8, 100_u8, 100_u8]);
    }

    _ = img.save("julia.png");
}
