fn main() {
    let contact = Contact {
        name: "Jan",
        age: 29,
    };
    println!("Hello {}, you are {} years old!", contact.name, contact.age);
}

struct Contact {
    name: String,
    age: i32,
}
