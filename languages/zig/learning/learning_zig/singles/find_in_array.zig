const std = @import("std");

pub fn main() void {
    const collection = [_]u32{ 1, 4, 8, 2, 3, 5, 6 };

    std.debug.print("Collection contains 5: {}\n", .{contains(&collection, 5)});
    std.debug.print("Collection contains 7: {}\n", .{contains(&collection, 7)});
}

fn contains(haystack: []const u32, needle: u32) bool {
    for (haystack) |value| {
        if (needle == value) {
            return true;
        }
    }
    return false;
}
