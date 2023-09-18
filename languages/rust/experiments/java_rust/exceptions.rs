fn main() {
    let result = floor_divide(10.0, 4.0);
    println!("The result of dividing 10 by 4 is {:?}", result);

    let by_zero = floor_divide(10.0, 0.0);
    println!("The result of dividing 10 by 0 is {:?}", by_zero);

    match floor_divide(11.0, 5.0) {
        Ok(result) => println!("Got {}", result),
        Err(error) => println!("Error: {}", error),
    }
}

fn floor_divide(num: f32, by: f32) -> Result<i32, String> {
    // crashes the app immediately
    // if by == 0.0 {
    //     panic!("Division by zero!");
    // }
    if by == 0.0 {
        Err("Division by zero!".to_string())
    } else {
        Ok((num / by).floor() as i32)
    }
}
