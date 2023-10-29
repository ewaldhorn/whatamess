extern fn b() i32;
extern fn print(val: i32) void;

extern "app" fn c() i32;

export fn add_two(l: i32, r: i32) i32 {
    return l + r;
}

export fn add(a: i32) i32 {
    return a + b();
}

export fn double(a: i32) i32 {
    consoleLog("This is a message.", "This is a message.".len);
    return a + a;
}

export fn use_c() i32 {
    return c();
}

export fn print_five() void {
    print(5);
}

extern fn consoleLog(msg: [*]const u8, len: u8) void;
