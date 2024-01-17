// Based on https://zig.guide/posts/fizz-buzz
// We can also write this using a switch over an integer. Here we're using @boolToInt which 
// converts bool values into a u1 value (i.e. a 1 bit unsigned integer). You may notice that we 
// haven't given div_5 an explicit type - this is because it is inferred from the value that is 
// assigned to it. We have however given div_3 a type; this is as integers may widen to larger 
// ones, meaning that they may coerce to larger integer types providing that the larger integer 
// type has at least the same range as the smaller integer type. We have done this so that the 
// operation div_3 * 2 + div_5 provides us a u2 value, or enough to fit two booleans.

const std = @import("std");

pub fn main() !void {
    var count: u8 = 1;

    while (count <= 30) : (count += 1) {
        const div_3: u2 = @intFromBool(count % 3 == 0);
        const div_5 = @intFromBool(count % 5 == 0);

        switch (div_3 * 2 + div_5) {
            0b10 => std.debug.print("Fizz\n", .{}),
            0b11 => std.debug.print("Fizz Buzz\n", .{}),
            0b01 => std.debug.print("Buzz\n", .{}),
            0b00 => std.debug.print("{}\n", .{count}),
        }
    }
}
