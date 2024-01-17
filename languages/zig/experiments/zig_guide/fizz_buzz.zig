// Based on https://zig.guide/posts/fizz-buzz
const std = @import("std");

pub fn main() !void {
    var count: u8 = 1;

    while (count <= 30) : (count += 1) {
        if (count % 5 == 0 and count % 3 == 0) {
            std.debug.print("Fizz Buzz\n", .{});
        } else if (count % 5 == 0) {
            std.debug.print("Buzz\n", .{});
        } else if (count % 3 == 0) {
            std.debug.print("Fizz\n", .{});
        } else {
            std.debug.print("{d}\n", .{count});
        }
    }
}
