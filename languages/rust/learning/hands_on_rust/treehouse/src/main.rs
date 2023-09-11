#![warn(clippy::all, clippy::pedantic)]

mod prompts;

fn main() {
    println!("Who goes there?");

    let name = prompts::get_user_name();
    let mut allowed = false;
    println!("Oh, hello there {name}!");

    let allowed_visitors = ["dan", "sally", "rico", "tammy"];

    for visitor in &allowed_visitors {
        if &name == visitor {
            allowed = true;
        }
    }

    if allowed {
        println!("Please come in!");
    } else {
        println!("Seems you aren't on the list, here's an application form.");
    }

    println!(
        "\nThe first name in allowed_visitors is {}.",
        allowed_visitors[0]
    );
}
