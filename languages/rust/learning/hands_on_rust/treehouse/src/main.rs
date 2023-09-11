#![warn(clippy::all, clippy::pedantic)]

use crate::visitor::Visitor;

mod prompts;
mod visitor;

fn main() {
    let mut allowed_visitors = vec![
        Visitor::new("bert", "Hello Bert, enjoy your treehouse."),
        Visitor::new("steve", "Hi Steve. Your milk is in the fridge."),
        Visitor::new("fred", "Wow, who invited Fred?"),
    ];

    loop {
        println!("Who goes there?");
        let name = prompts::get_user_name();
        let allowed = allowed_visitors.iter().find(|v| v.name == name);

        if let Some(v) = allowed {
            v.greet_visitor();
        } else {
            if name.is_empty() {
                break;
            }
            println!("Seems you aren't on the list, here's an application form.");
            allowed_visitors.push(Visitor::new(&name, "New Friend"));
        }
    }
    println!("The final list of visitors:");
    println!("{allowed_visitors:#?}");
}
