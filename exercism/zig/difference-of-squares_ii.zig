pub fn squareOfSum(number: usize) usize {
    var result: usize = 0;

    for (1..number+1) |n| {
        result += n;
    }

    return result * result;
}

pub fn sumOfSquares(number: usize) usize {
    var result: usize = 0;

    var n : usize = 1;

    while (n <= number) : (n += 1) {
        result += n * n;
    }

    return result;
}

pub fn differenceOfSquares(number: usize) usize {
    return squareOfSum(number) - sumOfSquares(number);
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const std = @import("std");
const expect = std.testing.expect;
const expectEqual = std.testing.expectEqual;

test "for 0" {
    const result = differenceOfSquares(0);

    try expect(result == 0);
}

test "square of sum up to 1" {
    const expected: usize = 1;
    const actual = squareOfSum(1);
    try expectEqual(expected, actual);
}
test "square of sum up to 5" {
    const expected: usize = 225;
    const actual = squareOfSum(5);
    try expectEqual(expected, actual);
}
test "square of sum up to 100" {
    const expected: usize = 25_502_500;
    const actual = squareOfSum(100);
    try expectEqual(expected, actual);
}
test "sum of squares up to 1" {
    const expected: usize = 1;
    const actual = sumOfSquares(1);
    try expectEqual(expected, actual);
}
test "sum of squares up to 5" {
    const expected: usize = 55;
    const actual = sumOfSquares(5);
    try expectEqual(expected, actual);
}
test "sum of squares up to 100" {
    const expected: usize = 338_350;
    const actual = sumOfSquares(100);
    try expectEqual(expected, actual);
}
test "difference of squares up to 1" {
    const expected: usize = 0;
    const actual = differenceOfSquares(1);
    try expectEqual(expected, actual);
}
test "difference of squares up to 5" {
    const expected: usize = 170;
    const actual = differenceOfSquares(5);
    try expectEqual(expected, actual);
}
test "difference of squares up to 100" {
    const expected: usize = 25_164_150;
    const actual = differenceOfSquares(100);
    try expectEqual(expected, actual);
}