fn main() {
    // arrays
    let my_array = [2, 4, 6, 8, 10];
    let my_big_array = [12; 50]; // 50 items, all initialized to 12
    let my_struct_array: [MyStruct; 3] = [MyStruct {}, 3]; // array of 3 MyStructs

    // tuples
    let my_tuple = (123, 345, 567);
    println!("The second item is {}", my_tuple.1);
    let my_custom_tuple = Bar(123, 456, false);

    // vector
    let mut my_vector: Vec<i32> = Vec::new();
    my_vector.push(123);
    my_vector.push(345);
    println!("The second item in the vector is {}", my_vector.get(1));
}

#[derive(Clone, Copy, Debug)]
struct MyStruct {}

// tuple struct
struct Bar(i32, i32, bool);
