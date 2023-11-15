const stdlib = @import("std");

pub fn main() !void {
    const stdout = stdlib.io.getStdOut().writer();
    try stdout.print("Hello {s}!\n", .{"Zig"});
}
