const std = @import("std");

fn multiDispatch(x : anytype) @TypeOf(x) {
    // note that this if statement happens at compile-time, not runtime.
    if (@TypeOf(x) == i64) {
        return x + 2;
    } else {
        return 2 * x;
    }
}

pub fn main() void {
    const one: i64 = 40;
    const two: i32 =  40;

    std.debug.print("i64-foo: {}\n", .{multiDispatch(one)});
    std.debug.print("i32-foo: {}\n", .{multiDispatch(two)});
}