const std = @import("std");

fn arrivalTimeDesc(minutes: u16, is_late: bool) []const u8 {
    switch (minutes) {
        0 => return "arrived",
        1, 2 => return "soon",
        3...5 => return "no more than 5 minutes",
        else => {
            if (!is_late) {
                return "sorry, it'll be a while";
            }
            // todo, something is very wrong
            return "never";
        },
    }
}

pub fn main() void {
    std.debug.print("Switches in Zig.\n", .{});
    std.debug.print("Description for 0 minutes: {s}.\n", .{arrivalTimeDesc(0, false)});
    std.debug.print("Description for 2 minutes: {s}.\n", .{arrivalTimeDesc(2, false)});
    std.debug.print("Description for 4 minutes: {s}.\n", .{arrivalTimeDesc(4, false)});
    std.debug.print("Description for 11 minutes: {s}.\n", .{arrivalTimeDesc(11, false)});
    std.debug.print("Description for 18 minutes: {s}.\n", .{arrivalTimeDesc(18, true)});
}
