const std = @import("std");
const expect = std.testing.expect;

fn sumFromOneTo(num: usize) usize {
    var tmp : usize = 1;

    for (1..num) |i| {
        tmp += i;
    }

    return tmp;
}

// =========================================================================================== TESTS
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