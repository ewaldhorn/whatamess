// Division with error handling in Rust

fn perform_division(numerator: i32, denominator: i32) -> Result<f32, String> {
    if denominator == 0 {
        return Err(String::from("Division by zero not supported."));
    }

    Ok(numerator as f32 / denominator as f32)
}

fn main() {
    let num = 123;
    let dom = 0;

    match perform_division(num, dom) {
        Ok(result) => println!("Result: {result}"),
        Err(error) => println!("Error: {error}"),
    }
}
