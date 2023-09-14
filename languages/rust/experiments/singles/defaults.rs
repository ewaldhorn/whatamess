#[derive(Default)]
struct Point {
    x: f32,
    y: f32,
}

impl Point {
    fn distance_from_origin(&self) -> f32 {
        (self.x.powi(2) + self.y.powi(2)).sqrt()
    }
}

fn main() {
    let origin = Point::default();
    let point_1 = Point { x: 2.4, y: 3.2 };
    let point_2 = Point { x: 0.5, y: 0.5 };

    println!(
        "Distance from origin for origin: {}",
        origin.distance_from_origin()
    );
    println!(
        "Distance from origin for point_1: {}",
        point_1.distance_from_origin()
    );
    println!(
        "Distance from origin for point_2: {}",
        point_2.distance_from_origin()
    );
}
