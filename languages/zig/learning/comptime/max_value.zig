const std = @import("std");

fn add_up(left: i32, comptime right: i32) i32 {
    if (right > 50) @compileError("No values over 50 allowed.");

    return left + right;
}

pub fn main() void {
    std.debug.print("Ten and Ten is {d}.\n", .{add_up(10, 10)});
    std.debug.print("Ten and Fifty is {d}.\n", .{add_up(10, 50)});

    // looks like legit code but won't compile
    std.debug.print("Ten and Sixty is {d}.\n", .{add_up(10, 60)});
}