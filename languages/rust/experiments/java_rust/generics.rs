fn main() {
    let container_int = Container { item: 456 };
    let container_string = Container {
        item: "Hello".to_string(),
    };
}

struct Container<T> {
    item: T,
}
