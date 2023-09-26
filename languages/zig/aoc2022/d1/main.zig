const std = @import("std");

pub fn main() !void {
    std.debug.print("AOC2022 Day 1 in Zig.\n", .{});

    // deal with opening the input file
    const inputFilename = "input.txt";
    const input = try std.fs.cwd().openFile(inputFilename, .{});
    defer input.close();

    // now read the file
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    const read_buffer = try input.readToEndAlloc(allocator, 1024*1024);

    try std.io.getStdOut().writeAll(read_buffer);
}