#![warn(clippy::all, clippy::pedantic)]

use crate::visitor::Visitor;

mod prompts;
mod visitor;

fn main() {
    println!("Who goes there?");

    let name = prompts::get_user_name();

    let allowed_visitors = [
        Visitor::new("bert", "Hello Bert, enjoy your treehouse."),
        Visitor::new("steve", "Hi Steve. Your milk is in the fridge."),
        Visitor::new("fred", "Wow, who invited Fred?"),
    ];

    let allowed = allowed_visitors.iter().find(|v| v.name == name);
    match allowed {
        Some(v) => v.greet_visitor(),
        None => println!("Seems you aren't on the list, here's an application form."),
    }
}
