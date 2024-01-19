const std = @import("std");
const expect = std.testing.expect;

fn sumFromOneTo(num: usize) usize {
    var tmp : usize = 1;

    for (1..num) |i| {
        tmp += i;
    }

    return tmp;
}

fn rangeHasNumber(begin: usize, end: usize, number: usize) bool {
    var i = begin;
    return while (i < end) : (i += 1) {
        if (i == number) {
            break true;
        }
    } else false;
}

// =========================================================================================== TESTS
test "while loop expression" {
    try expect(rangeHasNumber(0, 10, 3));
}

test "sum 1..1 will be one, as the last number is excluded" {
    const val = sumFromOneTo(1);

    try expect(val == 1);
}

test "sum 1..2 will be two, as the last number is excluded" {
    const val = sumFromOneTo(2);

    try expect(val == 2);
}

test "sum 1..3 will be four, as the last number is excluded" {
    const val = sumFromOneTo(3);

    try expect(val == 4);
}

test "sum 1..4 will be zero, as the last number is excluded" {
    const val = sumFromOneTo(4);

    try expect(val == 7);
}