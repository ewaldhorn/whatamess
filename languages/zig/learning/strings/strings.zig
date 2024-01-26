const std = @import("std");
const expect = std.testing.expect;

test "different ways to init strings" {
    const literal = "hello"; // *const [5:0]u8
    const array = [_]u8{'h', 'e', 'l', 'l', 'o'}; // [5]u8

    try expect(std.mem.eql(u8, literal, &array) == true);
}
