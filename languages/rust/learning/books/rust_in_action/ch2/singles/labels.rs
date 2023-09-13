fn main() {
    'outside: for x in 0.. {
        for y in 0.. {
            if x + y > 1234 {
                break 'outside;
            } else {
                println!("Currently at {x}:{y}");
            }
        }
    }
}