const std = @import("std");

pub fn main() !void {
    // Prints to stderr (it's a shortcut based on `std.io.getStdErr()`)
    std.debug.print("All your {s} are belong to us.\n", .{"codebase"});

    // stdout is for the actual output of your application, for example if you
    // are implementing gzip, then only the compressed bytes should be sent to
    // stdout, not any debugging messages.
    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);
    const stdout = bw.writer();

    try stdout.print("Run `zig build test` to run the tests.\n", .{});

    try bw.flush(); // don't forget to flush!

    // Primitives
    var x: i8 = -100; // Signed 8-bit integer
    var y: u8 = 120; // Unsigned 8-bit integer
    var z: f32 = 100.324; // 32-bit floating point

    std.debug.print("x={}\n", .{x}); // x=-100
    std.debug.print("y={}\n", .{y}); // y=120
    std.debug.print("z={d:3.2}\n", .{z}); // z=100.32
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
