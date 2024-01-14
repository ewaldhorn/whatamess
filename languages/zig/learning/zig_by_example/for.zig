const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    var array = [_]u32{ 1, 2, 3 };

    for (array) |elem| {
        print("by val: {}\n", .{elem});
    }

    for (&array) |*elem| {
        elem.* += 100;
        print("by ref: {}\n", .{elem.*});
    }

    for (array, &array) |val, *ref| {
        _ = val;
        _ = ref;
    }

    for (0.., array) |i, elem| {
        print("{}: {}\n", .{ i, elem });
    }

    for (array) |_| {}
}
