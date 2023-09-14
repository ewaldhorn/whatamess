mod boxes;

fn main() {
    println!("Square");
    boxes::draw_filled_square(5);
    println!("\nTriangle:");
    boxes::draw_triangle(5);
    println!("\nReversed Triangle:");
    boxes::draw_triangle_reversed(5);
}
