#![warn(clippy::all, clippy::pedantic)]

use crate::visitor::{Visitor, VisitorAction};

mod prompts;
mod visitor;

fn main() {
    let mut allowed_visitors = vec![
        Visitor::new("bert", VisitorAction::Accept, 45),
        Visitor::new(
            "steve",
            VisitorAction::AcceptWithNote {
                note: String::from("Lactose-free milk is in the fridge."),
            },
            29,
        ),
        Visitor::new("fred", VisitorAction::Refuse, 27),
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
            allowed_visitors.push(Visitor::new(&name, VisitorAction::Probation, 18));
        }
    }
    println!("The final list of visitors:");
    println!("{allowed_visitors:#?}");
}
