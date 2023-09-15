fn main() {
    println!("Numbers 0..9:");
    for num in 0..10 {
        println!("{num} is {}.", to_word(num));
    }
    println!();
    print!("Number ");
    print_as_words(2);
    print!("Number ");
    print_as_words(21);
    print!("Number ");
    print_as_words(218);
    print!("Number ");
    print_as_words(2187);
    print!("Number ");
    print_as_words(21876);
    print!("Number ");
    print_as_words(218762);
    print!("Number ");
    print_as_words(2187628);
}

fn to_word(num: usize) -> String {
    String::from(match num {
        0 => "zero",
        1 => "one",
        2 => "two",
        3 => "three",
        4 => "four",
        5 => "five",
        6 => "six",
        7 => "seven",
        8 => "eight",
        9 => "nine",
        _ => "?",
    })
}

fn print_as_words(num: usize) {
    let mut current = num;
    let mut nums: Vec<usize> = vec![];

    print!("{num} as words: ");

    while current > 0 {
        let tmp = current % 10;
        nums.push(tmp);
        current /= 10;
    }

    nums.reverse();

    for num in nums {
        print!("{} ", to_word(num));
    }
    println!();
}
