fn calculate_sum(numbers: &[i32]) -> Result<i32, String> {
    let mut sum = 0;
    for num in numbers.iter().cloned() {
        sum += num;
    }
    Ok(sum)
}

fn calculate_sum_better(numbers: &[i32]) -> Result<i32, String> {
    let sum = numbers.iter().sum();
    Ok(sum)
}

fn main() {
    let numbers = vec![1, 2, 3, 4, 5];
    let sum = calculate_sum(&numbers).unwrap();
    println!("Sum is {}.", sum);

    match calculate_sum_better(&numbers) {
        Ok(sum) => println!("Better sum is also {sum}."),
        Err(err) => eprintln!("The error is: {err}"),
    }
}
