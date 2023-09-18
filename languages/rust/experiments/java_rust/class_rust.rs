fn main() {
    let contact = Contact::new("Jan".to_string(), 29);
    println!("Hello {}, you are {} years old!", contact.name, contact.age);
}

struct Contact {
    name: String,
    age: i32,
}

impl Contact {
    fn new(name: String, age: i32) -> Self {
        Contact { name, age }
    }

    fn getName(&self) -> String {
        format!("{}", self.name)
    }
}

trait BusinessCard {
    fn card(&self) -> String;
}

impl BusinessCard for Contact {
    fn card(&self) -> String {
        format!("Business Card for {}", self.name)
    }
}
