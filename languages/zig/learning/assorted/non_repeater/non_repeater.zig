const std = @import("std");

// Find the first occurence of a non-repeat sequence of four characters.
// This will be four unique characters, where none of the characters are the same.
// possible regex: /^\%(.*\(.\).*\1\)\@!.*$
// regex in zig seems to not be quite there yet.
// so we'll do it with a simple algo, it's not a lot of comparisons after all.
pub fn main() void {
    const input = "znrznnozzroronznzona";
    std.debug.print("Examining:\n\"{s}\"\n\n", .{input});

    var pos: usize = 0;

    while (pos <= (input.len - 4)) : (pos += 1) {
        std.debug.print("{d:3} {s}", .{ pos, input[pos .. pos + 4] });

        if (no_repeats(input[pos .. pos + 4])) {
            std.debug.print(" *", .{});
        }
        std.debug.print("\n", .{});
    }

    std.debug.print("\n", .{});
}

fn no_repeats(s: []const u8) bool {
    if (s[0] != s[1] and s[0] != s[2] and s[0] != s[3] and s[1] != s[2] and s[1] != s[3] and s[2] != s[3]) {
        return true;
    }

    return false;
}
