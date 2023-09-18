use leptos::*;

fn main() {
    mount_to_body(|| view! { <p><h3>"Hello, world!"</h3>"This is some content."</p> })
}

pub struct Point {
    x: i32,
    y: i32,
}

impl Point {
    pub fn new(x: i32, y: i32) -> Point {
        Point { x, y }
    }
    pub fn get_x(&self) -> i32 {
        self.x
    }
    pub fn get_y(&self) -> i32 {
        self.y
    }
    pub fn set_x(&mut self, x: i32) {
        self.x = x;
    }
    pub fn set_y(&mut self, y: i32) {
        self.y = y;
    }
    pub fn add(&mut self, p: Point) {
        self.x += p.x;
        self.y += p.y;
    }
}
