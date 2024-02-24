const std = @import("std");

// Without using any built-in math functions:
// 
// Calculate the square root of the given number, returning a number rounded down 
// to the nearest integer.
fn getSquareRootOf(num: u32) u32 {
    var tmp: u32 = 1;

    while (tmp * tmp <= num) {
        tmp += 1;
    }

    return tmp - 1;
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const expectEqual = std.testing.expectEqual;

test "0" {
    const result = getSquareRootOf(0);

    try expectEqual(0, result);
}

test "4" {
    try expectEqual(2, getSquareRootOf(4));
}

test "8" {
    try expectEqual(2, getSquareRootOf(8));
}

test "16" {
    try expectEqual(4, getSquareRootOf(16));
}

test "100" {
    try expectEqual(10, getSquareRootOf(100));
}