const std = @import("std");

pub fn main() !void {
    const file = try std.fs.cwd().openFile("check_if_file_exists.zig", .{});
    defer file.close();

    const printToStdout = std.io.getStdOut().writer();

    var buffer: [64]u8 = undefined;

    // Or use heap-allocated buffer instead.
    //var buffer = try std.heap.page_allocator.alloc(u8, 64);
    //defer std.heap.page_allocator.free(buffer);

    var file_size: usize = 0;

    while (true) {
        const bytes_read = try file.read(&buffer);

        if (bytes_read == 0)
            break;

        file_size += bytes_read;

        printToStdout.print("{s}", .{buffer[0..bytes_read]}) catch {};
    }

    printToStdout.print("\nFile size is {d} bytes\n", .{file_size}) catch {};
}
