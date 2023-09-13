fn main() {
    greet_world();
}

fn greet_world() {
    println!("Hello, world!");
    let southern_germany = "Grüß Gott!";
    let japan = "ハロー・ワールド";
    let south_africa = "Yes ja!";
    let regions = [southern_germany, japan, south_africa];
    for region in regions {
        println!("{}", &region);
    }
}
