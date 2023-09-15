pub(crate) enum Colour {
    Red,
    Green,
    Blue,
}

pub(crate) fn select_colour(col: Colour) {
    match col {
        Colour::Red => println!("You selected RED!"),
        Colour::Green => println!("You selected GREEN!"),
        Colour::Blue => println!("You selected BLUE!"),
    }
}
