const std = @import("std");

const User = @import("user.zig").User;

pub fn main() void {
    const user = User {
        .power = 123,
        .name = "Bob"
    };

    std.debug.print("{s}'s power is {d}.\n", .{user.name, user.power});
}