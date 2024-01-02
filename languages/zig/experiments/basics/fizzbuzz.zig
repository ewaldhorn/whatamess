// Original from https://gist.github.com/andrewrk/6780fa252b693169897686e907c9da2a
const std = @import("std");

pub fn main() void {
    var i: usize = 1;

    while (i <= 30) : (i += 1) {
        if (i % 3 == 0 and i % 5 == 0) {
            std.debug.print("FizzBuzz\n", .{});
        } else if (i % 3 == 0) {
            std.debug.print("Fizz\n", .{});
        } else if (i % 5 == 0) {
            std.debug.print("Buzz\n", .{});
        } else {
            std.debug.print("{d}\n", .{i});
        }
    }
}