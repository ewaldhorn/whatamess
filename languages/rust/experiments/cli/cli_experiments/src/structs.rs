struct ShippingBox {
    depth: i32,
    width: i32,
    height: i32,
}

impl ShippingBox {
    fn get_volume(&self) -> i32 {
        self.depth * self.height * self.width
    }
}

pub(crate) fn do_boxy_things() {
    let my_box = ShippingBox {
        depth: 3,
        width: 2,
        height: 5,
    };

    let vol = my_box.depth * my_box.height * my_box.width;
    println!(
        "A box of {}x{}x{} has a volume of {vol}.",
        my_box.depth, my_box.height, my_box.width
    );

    println!(
        "A box of {}x{}x{} has a volume of {}. (Using impl method)",
        my_box.depth,
        my_box.height,
        my_box.width,
        my_box.get_volume()
    );
}

struct GroceryItem {
    stock: i32,
    price: f64,
}

pub(crate) fn do_more_struct_things() {
    let soup = GroceryItem {
        stock: 20,
        price: 9.95,
    };

    println!("The price of soup is R{} per tin.", soup.price);
}
