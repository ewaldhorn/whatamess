#[derive(Debug)]
pub(crate) enum VisitorAction {
    Accept,
    AcceptWithNote { note: String },
    Refuse,
    Probation,
}

#[derive(Debug)]
pub(crate) struct Visitor {
    pub(crate) name: String,
    action: VisitorAction,
    age: i8,
}

impl Visitor {
    pub(crate) fn new(name: &str, action: VisitorAction, age: i8) -> Self {
        Self {
            name: name.to_lowercase(),
            action,
            age,
        }
    }

    pub(crate) fn greet_visitor(&self) {
        match &self.action {
            VisitorAction::Accept => {
                println!("Welcome to the treehouse, {}.", self.name);
            }
            VisitorAction::AcceptWithNote { note } => {
                println!("Welcome to the treehouse, {}.", self.name);
                println!("{note}");
                if self.age < 21 {
                    println!("Do not serve alcohol to {}.", self.name);
                }
            }
            VisitorAction::Refuse => {
                println!("{} is not allowed in!", self.name);
            }
            VisitorAction::Probation => {
                println!("{} is now a probationary member.", self.name);
            }
        }
    }
}
