const std = @import("std");

pub fn buzz(val: u32) []const u8 {
    if (val % 15 == 0) {
        return "FizzBuzz";
    } else if (val % 5 == 0) {
        return "Buzz";
    } else if (val % 3 == 0) {
        return "Fizz";
    }

    var buffer: [5]u8 = std.mem.zeroes([5]u8);
    const t = std.fmt.bufPrint(buffer[0..], "{d}", .{val}) catch "Error printing to buffer";
    std.debug.print("{s}", .{t});
    return t;
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const expect = std.testing.expect;
const eql = std.mem.eql;

test "get a number back - 1" {
    const result = buzz(1);

    try expect(eql(u8, "1", result));
}

test "get FizzBuzz back" {
    const result = buzz(15);

    try expect(eql(u8, "FizzBuzz", result));
}

test "get Fizz back" {
    const result = buzz(3);

    try expect(eql(u8, "Fizz", result));
}

test "get Buzz back" {
    const result = buzz(5);

    try expect(eql(u8, "Buzz", result));
}
