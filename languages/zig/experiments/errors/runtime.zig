const std = @import("std");

test "out of bounds" {
    const a = [3]u8{ 1, 2, 3 };
    const index: u8 = 5;
    const b = a[index];
    _ = b;
}

test "out of bounds, no safety" {
    @setRuntimeSafety(false);
    const a = [3]u8{ 1, 2, 3 };
    const index: u8 = 5;
    const b = a[index];
    _ = b;
}

test "unreachable" {
    const x: i32 = 1;
    const y: u32 = if (x == 2) 5 else unreachable;
    _ = y;
}

fn asciiToUpper(x: u8) u8 {
    return switch (x) {
        'a'...'z' => x + 'A' - 'a',
        'A'...'Z' => x,
        else => unreachable,
    };
}

test "unreachable switch" {
    try std.testing.expect(asciiToUpper('a') == 'A');
    try std.testing.expect(asciiToUpper('A') == 'A');
}