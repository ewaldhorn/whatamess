const std = @import("std");

pub fn main() !void {
    var it = std.mem.split(u8, "Hello World. This is a longish sentence.", " ");
    while (it.next()) |x| {
        std.debug.print("{s}\n", .{x});
    }
}