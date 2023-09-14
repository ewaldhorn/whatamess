pub(crate) fn draw_filled_square(size: usize) {
    for _ in 0..size {
        repeat_character("*", size);
        println!();
    }
}

pub(crate) fn draw_triangle(size: usize) {
    for i in 1..=size {
        repeat_character("*", i);
        println!();
    }
}

pub(crate) fn draw_triangle_reversed(size: usize) {
    for i in (1..=size).rev() {
        repeat_character("*", i);
        println!();
    }
}

fn repeat_character(character: &str, times: usize) {
    print!("{}", character.repeat(times));
}

pub fn print_spaces(times: usize) {
    repeat_character(" ", times);
}

pub(crate) fn draw_box(size: usize) {
    repeat_character("*", size);
    println!();
    for _ in 2..size {
        print!("*");
        print_spaces(size - 2);
        println!("*");
    }
    repeat_character("*", size);
    println!();
}
