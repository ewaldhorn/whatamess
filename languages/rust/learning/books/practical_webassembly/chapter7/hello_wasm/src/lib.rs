#[no_mangle]
pub fn add(left: i32, right: i32) -> i32 {
    left + right
}

#[no_mangle]
pub fn double(num: i32) -> i32 {
    num + num
}
