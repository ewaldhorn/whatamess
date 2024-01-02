const std = @import("std");

pub fn main() void {
    var i: usize = 1;

    while (i <= 30) : (i += 1) {
        const div3 = i % 3 == 0;
        const div5 = i % 5 == 0;

        if (div3 and div5) {
            std.debug.print("FizzBuzz\n", .{});
        } else if (div3) {
            std.debug.print("Fizz\n", .{});
        } else if (div5) {
            std.debug.print("Buzz\n", .{});
        } else {
            std.debug.print("{d}\n", .{i});
        }
    }
}

// 1
// 2
// Fizz
// 4
// Buzz
// Fizz
// 7
// 8
// Fizz
// Buzz
// 11
// Fizz
// 13
// 14
// FizzBuzz
// 16
// 17
// Fizz
// 19
// Buzz
// Fizz
// 22
// 23
// Fizz
// Buzz
// 26
// Fizz
// 28
// 29
// FizzBuzz
