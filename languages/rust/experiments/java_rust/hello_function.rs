fn main() {
    greetings();
}

fn greetings() {
    let coding = true;
    let mood = if coding { "happy" } else { "sad" };
    println!("Mood: {}", mood);
}
