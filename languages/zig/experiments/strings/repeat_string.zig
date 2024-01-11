// As per https://en.wikipedia.org/wiki/Zig_(programming_language)
const std = @import("std");

fn repeat(allocator: *std.mem.Allocator, original: []const u8, times: usize) std.mem.Allocator.Error![]const u8 {
    var buffer = try allocator.alloc(u8, original.len * times);

    for (0..times) |i| {
        std.mem.copyForwards(u8, buffer[(original.len * i)..], original);
    }

    return buffer;
}

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    var allocator = arena.allocator();

    const original = "Hello ";
    const repeated = try repeat(&allocator, original, 3);

    // Prints "Hello Hello Hello "
    try stdout.print("{s}\n", .{repeated});
}
