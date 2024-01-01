const std = @import("std");

pub fn main() void {
    std.debug.print("Hello, {s}!\n", .{"world"});

    // perform explicit type coercion
    const inferred_constant = @as(i32, 5);

    // create a string
    const my_string : []const u8 = "this is a string in Zig";

    std.debug.print("My constant is {d}.\n", .{inferred_constant});
    std.debug.print("The string is \"{s}\".\n", .{my_string});

    // now create a String type
    const String: type = []const u8;
    const any_string: String = "This is my string!!";
    std.debug.print("The string is \"{s}\".\n", .{any_string});

    // wrapping operators to prevent over/underflow errors
    var overflow:u8 = 255;

    overflow = overflow +% 1; // overflow becomes 0, because it wrapped around
    overflow = overflow -% 1; // back to 255
    overflow = overflow *% 2; // now it is 254
}