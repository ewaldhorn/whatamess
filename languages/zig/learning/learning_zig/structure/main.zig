const std = @import("std");
const User = @import("models/user.zig").User;
const SuperDooper = @import("models/user.zig").SuperUser;
const circle = @import("constants/constants.zig");
const adder = @import("math/adders/adder.zig");
const Stuff = @import("models/inventory_item.zig");

pub fn main() void {
    const user = User{
        .power = 9001,
        .name = "Goku",
    };

    std.debug.print("\n{s}'s power is {d}.\n", .{ user.name, user.power });

    const circle1 = circle.Circle{ .radius = 4.87, .sort_order = 2 };

    std.debug.print("The circle has a radius of {d:.3}.\n", .{circle1.radius});

    std.debug.print("The sum of 5 and 6 is {d}.\n", .{adder.add(5, 6)});

    const powerTripper = SuperDooper{
        .user = User{ .name = "Power Puff", .power = 40 },
    };

    std.debug.print("The power user's name is {s}.\n", .{powerTripper.user.name});

    std.debug.print("Power user ", .{});
    powerTripper.debugMessage();

    std.debug.print("Power user (alternate syntax.) ", .{});
    SuperDooper.debugMessage(powerTripper);

    const fishPaste = Stuff.InventoryItem.init("Wedwo", 34.95);
    fishPaste.report();

    std.debug.print("\n", .{});
}
