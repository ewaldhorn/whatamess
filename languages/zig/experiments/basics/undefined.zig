const std = @import("std");

pub fn main() void {
    var x: i32 = undefined;
    std.debug.print("undefined: {}\n", .{x});

    x = 10;
    std.debug.print("undefined: {}\n", .{x});
}
