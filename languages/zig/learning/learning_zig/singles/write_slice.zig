const std = @import("std");

pub fn main() void {
    var a = [_]i32{ 1, 2, 3, 4, 5 };
    var end: usize = 4;
    const b = a[1..end];
    b[2] = 99;

    std.debug.print("a[2]: {}\n", .{a[2]});
    std.debug.print("b[2]: {}\n", .{b[2]});

    a[2] = 10;

    std.debug.print("a[2]: {}\n", .{a[2]});
    std.debug.print("b[2]: {}\n", .{b[2]});
}
