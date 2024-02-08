pub fn main() void {
    const std2 = @import("std");
    var buffer: [20]u8 = undefined;
    const len = std2.fmt.formatIntBuf(&buffer, @as(i64, 3), 10, .lower, .{ .fill = '0', .width = 4 });
    std2.debug.print("{s}\n", .{buffer[0..len]});

    const len2 = std2.fmt.formatIntBuf(&buffer, @as(u64, 3), 10, .lower, .{ .fill = '0', .width = 4 });
    std2.debug.print("{s}\n", .{buffer[0..len2]});
}
