extern fn b() i32;

extern "app" fn c() i32;

export fn add_two(l: i32, r: i32) i32 {
    return l + r;
}

export fn add(a: i32) i32 {
    return a + b();
}

export fn double(a: i32) i32 {
    return a + a;
}

export fn use_c() i32 {
    return c();
}
