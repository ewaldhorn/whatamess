#![warn(clippy::all, clippy::pedantic)]

use std::io::stdin;

fn main() {
    println!("Who goes there?");

    let mut name = String::new();
    stdin().read_line(&mut name).expect("Unable to read line");

    println!("Oh, hello there {name}!");
}
