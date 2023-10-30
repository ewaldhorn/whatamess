const std = @import("std");

pub const User = struct {
    power: u64,
    name: []const u8,
};

// Struct with default values and containing another struct
pub const SuperUser = struct {
    user: User,
    isAdmin: bool = true,

    pub fn debugMessage(sup: SuperUser) void {
        std.debug.print("IsAdmin: {}.\n", .{sup.isAdmin});
    }
};
