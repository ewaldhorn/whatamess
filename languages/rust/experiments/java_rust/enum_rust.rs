fn main() {
    let current_job = JobStatus::Idle;

    match current_job {
        JobStatus::None => println!("Not available"),
        JobStatus::Idle => println!("Available"),
        JobStatus::Busy => println!("Busy!"),
    }
}

enum JobStatus {
    None,
    Idle,
    Busy,
}
