pub(crate) struct Visitor {
    pub(crate) name: String,
    greeting: String,
}

impl Visitor {
    pub(crate) fn new(name: &str, greeting: &str) -> Self {
        Self {
            name: name.to_lowercase(),
            greeting: greeting.to_string(),
        }
    }

    pub(crate) fn greet_visitor(&self) {
        println!("{}", self.greeting);
    }
}
