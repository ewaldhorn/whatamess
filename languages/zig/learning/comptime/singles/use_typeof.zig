const std = @import("std");

fn doSomething(x: anytype) @TypeOf(x) {
    // TypeOf happens at comptime here:
    if (@TypeOf(x) == i32) {
        return x + 1;
    } else if (@TypeOf(x) == i64) {
        return x + 2;
    } else {
        return x + 3;
    }
}

pub fn main() void {
    const one : i32 = 0;
    const two : i64 = 0;
    const three : f32 = 0.0;

    std.debug.print("One   is {d}.\n", .{doSomething(one)});
    std.debug.print("Two   is {d}.\n", .{doSomething(two)});
    std.debug.print("Three is {d}.\n", .{doSomething(three)});
}