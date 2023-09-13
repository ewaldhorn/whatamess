fn main() {
    for (i,j) in (0..).zip(0..) {
        if i + j > 123 {
            break;
        } else {
            println!("So {i} + {j} is {}", i+j);
        }
    }
}