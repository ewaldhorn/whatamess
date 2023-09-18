fn main() {
    let mut item = Item { count: 11 };
    let item_reference = &mut item;

    do_something(item_reference);
    do_something(&mut item);
}

fun do_something(item: &mut Item) {
    println!("Doing something with {} items.", item.count);
}

struct Item {
    count: i32,
}
