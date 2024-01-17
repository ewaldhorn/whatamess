// Based on https://zig.guide/posts/a-guessing-game
const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    const stdin = std.io.getStdIn().reader();

    // use the 'undefined' value (random memory bytes as found by the OS)
    // convert that to a number, it should be relatively random.
    var seed: u64 = undefined;
    try std.os.getrandom(std.mem.asBytes(&seed));

    var prng = std.rand.DefaultPrng.init(seed);
    const rand = &prng.random();

    const guess_me = rand.intRangeAtMost(u8, 1, 100);
    try stdout.print("Alright, try to guess a number between 1-100: ", .{});
    var guess_count : i32 = 0;

    while (true) {
        guess_count += 1;
        const raw_line = try stdin.readUntilDelimiterAlloc(std.heap.page_allocator, '\n', 8192);
        defer std.heap.page_allocator.free(raw_line);
        const line = std.mem.trim(u8, raw_line, "\r");

        const guess = std.fmt.parseInt(u8, line, 10) catch |err| switch (err) {
            error.Overflow => {
                try stdout.writeAll("Please enter a small positive number\n");
                continue;
            },
            error.InvalidCharacter => {
                try stdout.writeAll("Please enter a valid number\n");
                continue;
            },
        };

        if (guess > guess_me) {
            try stdout.print("Too high! Try again: ", .{});
        } else if (guess < guess_me) {
            try stdout.print("Too low! Try again: ", .{});
        } else {
            try stdout.print("\n\nNicely done!\n{} is indeed the correct number!\nYou took {d} guesses!\n\n", .{guess, guess_count});
            break;
        }
    }
}
