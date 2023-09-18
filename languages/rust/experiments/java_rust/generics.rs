use std::cmp::{Ord, Ordering};

fn main() {
    let container_int = Container { item: 456 };
    let container_string = Container {
        item: "Hello".to_string(),
    };

    println!("Compare int: {:?}", container_int.compare_item(&123));
    println!("Generic compare: {:?}", compare_items(&223, &332));
}

struct Container<T: Ord> {
    item: T,
}

impl<T: Ord> Container<T> {
    fn compare_item(&self, other: &T) -> Ordering {
        self.item.cmp(other)
    }
}

fn compare_items<T: Ord>(item1: &T, item2: &T) -> Ordering {
    item1.cmp(item2)
}
