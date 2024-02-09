const std = @import("std");

fn Vector2D(comptime T: type) type {
    return struct { x: T, y: T };
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
test "i32" {
    const v = Vector2D(i32){.x=1, .y=2};

    try std.testing.expect(v.x == 1);
    try std.testing.expect(v.y == 2);
}

test "f32" {
    const v = Vector2D(f32){.x=1.5, .y=2.7};

    try std.testing.expect(v.x == 1.5);
    try std.testing.expect(v.y == 2.7);
}