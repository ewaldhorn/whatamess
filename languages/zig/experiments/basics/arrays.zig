const std = @import("std");
const expect = @import("std").testing.expect;

pub fn main() void {
    // an array of 3 booleans with false as the sentinel value
    const a = [3:false]bool{ false, true, false };

    // This line is more advanced, and is not going to get explained!
    std.debug.print("{any}\n", .{std.mem.asBytes(&a).*});
}

// ========================================================================================== TESTS
test "basic array" {
    const a = [5]i32{ 1, 2, 3, 4, 5 };
    const b: [5]i32 = .{ 1, 2, 3, 4, 5 };
    const c = [_]i32{ 1, 2, 3, 4, 5 }; // use _ to let the compiler infer the length

    // verify lengths are the same
    try expect(a.len == b.len);
    try expect(b.len == c.len);

    // verify values are the same
    try expect(a[0] == b[0]);
    try expect(b[0] == c[0]);
}

test "update array values" {
    var a = [_]i32{ 1, 2, 3 };

    try expect(a[0] == 1);

    a[0] = 2;

    try expect(a[0] == 2);
}
