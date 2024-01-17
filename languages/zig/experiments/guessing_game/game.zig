// Based on https://zig.guide/posts/a-guessing-game
const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    var prng = std.rand.DefaultPrng.init(1919191919);
    const rand = prng.random();

    try stdout.print("A sort-of random number {d}\n", 
        .{rand.intRangeAtMost(u8, 1, 100)});
}
