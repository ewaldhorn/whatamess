const std = @import("std");

pub fn main() void {
    std.debug.print("Hello, {s}!\n", .{"world"});

    // perform explicit type coercion
    const inferred_constant = @as(i32, 5);

    // create a string
    const my_string : []const u8 = "this is a string in Zig";

    std.debug.print("My constante is {d}.\n", .{inferred_constant});
    std.debug.print("The string is \"{s}\".\n", .{my_string});
}