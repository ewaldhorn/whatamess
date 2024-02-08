const std = @import("std");

fn thing(a: std.mem.Allocator) !void {
    std.fmt.allocPrint(a, "", .{});
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const allocator = arena.allocator();
    try thing(allocator);
}