pub(crate) fn draw_filled_square(size: usize) {
    for _ in 0..size {
        repeat_character(String::from("*"), size);
        println!();
    }
}

pub(crate) fn draw_triangle(size: usize) {
    for i in 1..=size {
        repeat_character(String::from("*"), i);
        println!();
    }
}

pub(crate) fn draw_triangle_reversed(size: usize) {
    for i in (1..=size).rev() {
        repeat_character(String::from("*"), i);
        println!();
    }
}

fn repeat_character(character: String, times: usize) {
    print!("{}", character.repeat(times));
}
