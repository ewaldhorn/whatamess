const std = @import("std");

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const max_size = std.math.maxInt(usize);

    var data = try std.fs.cwd().readFileAlloc(allocator, "read_file.zig", max_size);
    defer allocator.free(data);

    const printToStdout = std.io.getStdOut().writer();

    printToStdout.print("{s}", .{data}) catch {};
}
