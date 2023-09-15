mod enums;
mod structs;
mod tuples;

#[derive(Debug)]
struct Data {
    answer: i32,
}

impl std::fmt::Display for Data {
    fn fmt(&self, _f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        Ok(())
    }
}
fn main() {
    let answer = 42;

    println!("Standard display is {answer}.");
    println!("Debug display is {answer:?}.");

    let data = Data { answer: 42 };
    println!("Standard display is {data} with {}.", data.answer);
    println!("Debug display is {data:?}.");

    enums::select_colour(enums::Colour::Green);

    structs::do_boxy_things();
    structs::do_more_struct_things();

    tuples::do_tuple_stuff();
}
