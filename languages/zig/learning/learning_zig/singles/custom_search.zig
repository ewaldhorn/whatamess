const std = @import("std");

pub fn main() void {
    const collection = [_]u32{ 1, 4, 8, 2, 3, 5, 6 };
    const collection2 = [_]u32{ 1, 4, 8, 2, 3, 5, 6 };
    const collection3 = [_]u32{ 1, 8, 4, 2, 3, 5, 6 };

    std.debug.print("Collections 1 and 2 are the same: {}\n", .{eql(u32, &collection, &collection2)});
    std.debug.print("Collections 2 and 3 are the same: {}\n", .{eql(u32, &collection2, &collection3)});
}

pub fn eql(comptime T: type, a: []const T, b: []const T) bool {
    // if they arent' the same length, the can't be equal
    if (a.len != b.len) return false;

    for (a, b) |a_elem, b_elem| {
        if (a_elem != b_elem) return false;
    }

    return true;
}
