const std = @import("std");

pub fn isArmstrongNumber(num: u128) bool {
    if (num < 10) {
        return true; // 0..9 are Armstrong numbers
    } else if (num < 100) {
        return false; // 10..99 are not Armstrong numbers
    }

    // we have a potential number

    // first determine how many digits we have.
    var digits : u128 = 0;
    var source = num;
    while (source > 0) {
        digits += 1;
        source = source / 10;
    }

    // now that we know how many digits we have, we can continue
    source = num; // reset to num
    var goal : u128 = 0;

    // loop through the number, getting one digit at a time using modulus (remainder)
    // then finally divide by 10 to get to the next digit
    while (source > 0) {
        const leftOver = source % 10;
        // std.debug.print("Working with {d}\n", .{leftOver});
        goal = goal + std.math.pow(u128, leftOver, digits);
        source = source / 10;
    }

    return goal == num;
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const testing = std.testing;

test "zero is an armstrong number" {
    try testing.expect(isArmstrongNumber(0));
}

test "single-digit numbers are armstrong numbers" {
    try testing.expect(isArmstrongNumber(5));
}

test "there are no two-digit armstrong numbers" {
    try testing.expect(!isArmstrongNumber(10));
}

test "three-digit number that is an armstrong number" {
    try testing.expect(isArmstrongNumber(153));
}

test "three-digit number that is not an armstrong number" {
    try testing.expect(!isArmstrongNumber(100));
}

test "four-digit number that is an armstrong number" {
    try testing.expect(isArmstrongNumber(9_474));
}

test "four-digit number that is not an armstrong number" {
    try testing.expect(!isArmstrongNumber(9_475));
}

test "seven-digit number that is an armstrong number" {
    try testing.expect(isArmstrongNumber(9_926_315));
}

test "seven-digit number that is not an armstrong number" {
    try testing.expect(!isArmstrongNumber(9_926_314));
}

test "33-digit number that is an armstrong number" {
    try testing.expect(isArmstrongNumber(186_709_961_001_538_790_100_634_132_976_990));
}

test "38-digit number that is not an armstrong number" {
    try testing.expect(!isArmstrongNumber(99_999_999_999_999_999_999_999_999_999_999_999_999));
}

test "the largest and last armstrong number" {
    try testing.expect(isArmstrongNumber(115_132_219_018_763_992_565_095_597_973_971_522_401));
}

test "the largest 128-bit unsigned integer value is not an armstrong number" {
    try testing.expect(!isArmstrongNumber(340_282_366_920_938_463_463_374_607_431_768_211_455));
}
