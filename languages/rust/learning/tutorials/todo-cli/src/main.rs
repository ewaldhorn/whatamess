use crate::todo::Todo;

mod todo;

fn main() {
    // let action = std::env::args().nth(1).expect("Please specify an action");
    // let item = std::env::args().nth(2).expect("Please specify an item");
    //
    // println!("{:?}, {:?}", action, item);

    let mut tasks = Todo::new();

    // should say the list is empty
    println!("Contents of Todo list:");
    tasks.list_all();

    tasks.insert(String::from("Build todo list"), None);
    tasks.insert("Bake bread".into(), Some("Need flour".into()));

    // leave a gap
    println!("\n");

    // expecting items in the list
    println!("Contents of Todo list:");
    tasks.list_all();
}
