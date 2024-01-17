// Based on https://zig.guide/posts/a-guessing-game
const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    // use the 'undefined' value (random memory bytes as found by the OS)
    // convert that to a number, it should be relatively random.
    var seed: u64 = undefined;
    try std.os.getrandom(std.mem.asBytes(&seed));

    var prng = std.rand.DefaultPrng.init(seed);
    const rand = &prng.random();

    try stdout.print("A sort-of random number {d}\n", 
        .{rand.intRangeAtMost(u8, 1, 100)});
}
