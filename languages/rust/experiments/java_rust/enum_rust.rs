fn main() {
    let current_job = JobStatus::Idle(3);

    match current_job {
        JobStatus::None(last_checked) => println!("Not available (last checked: {})", last_checked),
        JobStatus::Idle(hours) if hours > 5 => {
            println!("Idle for long! (idle for {})", hours)
        }
        JobStatus::Idle(hours) => println!("Available (idle for {})", hours),
        JobStatus::Busy => println!("Busy!"),
    }

    let another_job = JobStatus::Busy;
    let job_status_message = match another_job {
        JobStatus::None(_) => "Is not available",
        JobStatus::Idle(_) => "Is Idle",
        JobStatus::Busy => "Is busy",
    };

    println!("The message associated with the job status: {job_status_message}");
}

enum JobStatus {
    None(String), // last checked at
    Idle(u8),     // how long it's been idle for
    Busy,
}
