const std = @import("std");
const expect = @import("std").testing.expect;

pub fn main() void {}

// ========================================================================================== TESTS
test "simple test" {
    const simple = true;

    try expect(simple);
}