#![warn(clippy::all, clippy::pedantic)]

mod prompts;

fn main() {
    println!("Who goes there?");

    let name = prompts::get_user_name();
    println!("Oh, hello there {name}!");
}
