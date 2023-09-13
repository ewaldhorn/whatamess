use std::io;

use rand::Rng;

fn main() {
    let secret_number = rand::thread_rng().gen_range(1..=100);
    let mut guesses = 0;

    loop {
        println!("Pick a number (1..100):");
        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read input.");

        let guess: u32 = guess.trim().parse().expect("That's not a number!");
        guesses += 1;
        match guess.cmp(&secret_number) {
            std::cmp::Ordering::Less => println!("Too small!"),
            std::cmp::Ordering::Equal => break,
            std::cmp::Ordering::Greater => println!("Too big!"),
        }
    }

    println!("Got it in {guesses} guesses!");
}
