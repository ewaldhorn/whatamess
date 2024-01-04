const std = @import("std");

const User = struct { id: i32, power: i32 };

pub fn main() void {
    std.debug.print("Zig Pointers:\n\n", .{});

    const user = User{
        .id = 1,
        .power = 100,
    };
    std.debug.print("{*}\n{*}\n{*}\n", .{ &user, &user.id, &user.power });

    const user_p = &user;
    std.debug.print("{any}\n", .{@TypeOf(user_p)});
}
