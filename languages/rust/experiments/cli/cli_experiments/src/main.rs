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
}
