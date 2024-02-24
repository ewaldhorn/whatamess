const std = @import("std");

fn get_input(guesses: i32) !i64 {
    const stdin = std.io.getStdIn().reader();
    const stdout = std.io.getStdOut().writer();

    var buf: [10]u8 = undefined;

    try stdout.print("({d}) Guess a number between 1 and 100: ", .{guesses});

    if (try stdin.readUntilDelimiterOrEof(buf[0..], '\n')) |user_input| {
        if (std.mem.eql(u8, user_input, "x")) {
            return error.Exit;
        } else {
            const x = std.fmt.parseInt(i64, user_input, 10);
            return x;
        }
    } else {
        return error.InvalidParam;
    }
}

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    var prng = std.rand.DefaultPrng.init(blk: {
        var seed: u64 = undefined;
        try std.os.getrandom(std.mem.asBytes(&seed));
        break :blk seed;
    });

    const value = prng.random().intRangeAtMost(i64, 1, 100);

    var winCon = false;
    var guesses: i32 = 0;

    while (true) {
        guesses = guesses + 1;

        const guess = get_input(guesses) catch |err| {
            switch (err) {
                error.InvalidCharacter => {
                    try stdout.print("\x1b[31mPlease enter a number.\x1b[0m\n", .{});
                    continue;
                },
                error.Exit => break,
                else => return err,
            }
        };

        if (guess == value) {
            winCon = true;
            break;
        }

        const message = if (guess < value) "\x1b[33mlow\x1b[0m" else "\x1b[31mhigh\x1b[0m";
        try stdout.print("{d} is too {s}.\n", .{ guess, message });
    }

    if (winCon) {
        try stdout.print("\x1b[32m({d}) Right! The number was {d}.\x1b[0m\n", .{ guesses, value });
    } else {
        try stdout.print("Bye!", .{});
    }
}
