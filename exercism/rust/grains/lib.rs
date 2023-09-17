pub fn square(s: u32) -> u64 {
    if s < 1 || s > 64 {
        panic!("Square must be between 1 and 64");
    }
    
    if s == 1 {
        return 1
    } else {
        return 2u64.pow(s-1)
    }
}

pub fn total() -> u64 {
    let mut tmp : u64 = 0;
    
    for i in 1..=64 {
        tmp += square(i);
    }

    tmp
}
