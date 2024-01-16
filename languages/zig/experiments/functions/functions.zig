const std = @import("std");

pub fn main() void {}

fn add(a: i32, b: i32) i32 {
    return a + b;
}

// ========================================================================================== TESTS
test "add works (7,8)" {
    try std.testing.expectEqual(15, add(8, 7));
}

test "add works (0,-1)" {
    try std.testing.expectEqual(-1, add(0, -1));
}

test "add works (0,0)" {
    try std.testing.expectEqual(0, add(0, 0));
}

test "add works (-1,0)" {
    try std.testing.expectEqual(-1, add(-1, 0));
}

test "add works (-1,-1)" {
    try std.testing.expectEqual(-2, add(-1, -1));
}
