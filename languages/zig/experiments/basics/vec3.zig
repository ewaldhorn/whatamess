const std = @import("std");

const Vec3 = struct { x: f64, y: f64, z: f64 };

pub fn main() void {
    const v: Vec3 = .{ .x = 0.1, .y = 0.2, .z = 0.05 };
    std.debug.print("v: {}\n", .{v});
}
