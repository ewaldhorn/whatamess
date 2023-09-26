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
    defer allocator.free(read_buffer);

    // output everything
    // try std.io.getStdOut().writeAll(read_buffer);

    // create an iterator so that we can parse the input line for line
    var iterator = std.mem.split(u8, read_buffer, "\n");

    var all_elves = std.ArrayList(u32).init(allocator);
    defer all_elves.deinit();

    var currentCount : u32 = 0;
    // loop over the iterator
    while (iterator.next()) |amount| {
        if (amount.len == 0) {
            // std.debug.print("------\n", .{});
            try all_elves.append(currentCount);
            currentCount = 0;
        } else {
            var tmp : u32 = try std.fmt.parseInt(u32, amount, 10);
            // std.debug.print("{d}\n", .{tmp});
            currentCount += tmp;
        }
    }
    var fattest : u32 = std.mem.max(u32, all_elves.items);
    std.debug.print("The largest count was {}.\n", .{fattest});
}