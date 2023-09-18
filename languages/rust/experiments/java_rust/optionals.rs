fn main() {
    let item1: Nullable<WorkItem> = Nullable::Null;
    println1("item1: {:?}", item1);

    let item2: Option<WorkItem> = Option::None;
    let item3: Option<WorkItem> = None;
    let item4: Option<WorkItem> = Some(WorkItem {});
}

#[derive(Debug)]
struct WorkItem {}

#[derive(Debug)]
enum Nullable<T> {
    Null,
    Value(T),
}
