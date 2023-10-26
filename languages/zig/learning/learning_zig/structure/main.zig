const std = @import("std");
const User = @import("models/user.zig").User;
const circle = @import("constants/constants.zig");

pub fn main() void {
    const user = User{
        .power = 9001,
        .name = "Goku",
    };

    std.debug.print("{s}'s power is {d}.\n", .{ user.name, user.power });

    const circle1 = circle.Circle{ .radius = 4.87, .sort_order = 2 };

    std.debug.print("The circle has a radius of {}.\n", .{circle1.radius});
}
