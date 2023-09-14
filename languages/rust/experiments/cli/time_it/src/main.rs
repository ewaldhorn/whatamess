use std::time::Instant;

fn burn_some_cpu_cycles_sorry_planet() {
    let mut _total = 0;

    for _ in 0..10000 {
        for _ in 0..5000 {
            _total += 1;
        }
    }
}

fn main() {
    let start = Instant::now();
    println!("Starting...");
    burn_some_cpu_cycles_sorry_planet();
    let elapsed_time = start.elapsed().as_micros();
    println!("The elapsed time is around {}ms", elapsed_time);
}
