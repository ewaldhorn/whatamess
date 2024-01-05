const std = @import("std");

pub fn main() void {
    std.debug.print("Please wait: ", .{});

    for (1..10) |_| {
        std.debug.print(".", .{});
        std.time.sleep(std.time.ns_per_s * 1);
    }

    std.debug.print(" Done!\n", .{});
}