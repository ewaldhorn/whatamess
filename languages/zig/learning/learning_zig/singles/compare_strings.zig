const std = @import("std");

pub fn main() void {
    const tmp1 = "this is a string";
    var tmp2 = "This is a string";

    std.debug.print("Basic string comparisons.\n", .{});
    std.debug.print("'{s}' is equal to '{s}': {} (case matters).\n", .{ tmp1, tmp2, std.mem.eql(u8, tmp1, tmp2) });
    std.debug.print("'{s}' is equal to '{s}': {} (case doesn't matter).\n", .{ tmp1, tmp2, std.ascii.eqlIgnoreCase(tmp1, tmp2) });
}
