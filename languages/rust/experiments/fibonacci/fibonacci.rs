use std::hint::black_box;

fn fibonacci(n: u64) -> u64 {
    match n {
        1 | 0 => n,
        _ => fibonacci(n - 1) + fibonacci(n - 2),
    }
}

fn main() {
    let mut total: f64 = 0.0;
    for _ in 1..=20 {
        let start = std::time::Instant::now();
        black_box(fibonacci(black_box(40)));
        let elapsed = start.elapsed().as_secs_f64();
        total += elapsed;
    }
    let avg_time = total / 20.0;
    println!("Average time taken: {} s", avg_time);
}
