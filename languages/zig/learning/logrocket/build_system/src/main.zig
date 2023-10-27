const std = @import("std");

pub fn main() !void {
    // Prints to stderr (it's a shortcut based on `std.io.getStdErr()`)
    std.debug.print("\nAll your {s} are belong to us.\n", .{"codebase"});

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
    var z: f32 = 100.324112; // 32-bit floating point
    var zz: f32 = 0.005; // 32-bit floating point

    std.debug.print("x={}\n", .{x}); // x=-100
    std.debug.print("y={}\n", .{y}); // y=120

    // number formatting with spaces and precision
    std.debug.print("z={d:12.2}\n", .{z}); // z=      100.32
    std.debug.print("z={d:12.4}\n", .{z}); // z=    100.3241
    std.debug.print("z={d:12.6}\n", .{z}); // z=  100.324112
    std.debug.print("zz={d:11.3}\n", .{zz}); //zz=     0.005

    // Characters
    const l1: u8 = 'Z';
    const l2: u8 = 'i';
    const l3: u8 = 'g';

    std.debug.print("{c}-{c}-{c}\n", .{ l1, l2, l3 }); // Z-i-g

    // Enums
    const LogType = enum { info, err, warn };

    const ltInfo = LogType.info;
    const ltErr = LogType.err;

    std.debug.print("{}\n", .{ltInfo}); // main.main.LogType.info
    std.debug.print("{}\n", .{ltErr}); // main.main.LogType.err

    const WaitingPeriods = enum(u32) { day = 1, week = 7, month = 30 };

    std.debug.print("Waiting period one is {}.\n", .{WaitingPeriods.day});
    std.debug.print("Waiting period two is {}.\n", .{WaitingPeriods.week});

    // Arrays and Slices
    const vowels = [_]u8{ 'a', 'e', 'i', 'o', 'u' };

    std.debug.print("{s}\n", .{vowels}); // aeiou
    std.debug.print("{d}\n", .{vowels.len}); // 5

    const msg = "Ziglang"; // c-style strings
    std.debug.print("{s}\n", .{msg}); // Zig
    std.debug.print("{}\n", .{@TypeOf(msg)}); // *const [7:0]u8

    const msg1 = "Zig";
    const msg2 = "lang";

    std.debug.print("{s}\n", .{msg1 ** 2}); // ZigZig
    std.debug.print("{s}\n", .{msg1 ++ msg2}); // Ziglang

    const nums = [_]u8{ 2, 5, 6, 4 };
    var x1: usize = 3;
    const slice = nums[1..x1];

    std.debug.print("{any}\n", .{slice}); // { 5, 6 }
    std.debug.print("{}\n", .{@TypeOf(slice)}); // []const u8

    // Footer
    std.debug.print("\n", .{});
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
