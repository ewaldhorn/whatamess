use std::io::stdin;

pub fn get_user_name() -> String {
    let mut name = String::new();
    stdin().read_line(&mut name).expect("Unable to read line");
    name.trim().to_lowercase()
}
