const stdlib = @import("std");
const expect = @import("std").testing.expect;

pub fn main() void {
    const user = User{ .power = 123, .name = "Roger" };

    stdlib.debug.print("{s}'s power is {d}.\n", .{ user.name, user.power });

    // access struct functions
    user.printDebugMessage();
    User.printDebugMessage(user); // can also call this way, if needed

    const anotherUser: User = .{ .power = 1, .name = "Sal" };
    anotherUser.printDebugMessage();
}

pub const User = struct {
    power: u64,
    name: []const u8,
    version: u32 = DEFAULT_VERSION,

    pub const DEFAULT_VERSION = 1;

    fn printDebugMessage(user: User) void {
        stdlib.debug.print("{s} is an object of version type {d}.\n", .{ user.name, user.version });
    }

    pub fn initDefault() User {
        return User{ .power = 0, .name = "Default" };
    }
};

// ========================================================================================== TESTS
test "simple struct with default version" {
    const simple = User{ .power = 555, .name = "Bob" };

    try expect(simple.power == 555);
    try expect(stdlib.mem.eql(u8, simple.name, "Bob"));
    try expect(simple.version == 1);
}

test "simple struct with custom version" {
    const simple = User{ .power = 555, .name = "Bob", .version = 2 };

    try expect(simple.power == 555);
    try expect(stdlib.mem.eql(u8, simple.name, "Bob"));
    try expect(simple.version == 2);
}

test "default init user" {
    const simple = User.initDefault();

    try expect(simple.power == 0);
    try expect(stdlib.mem.eql(u8, simple.name, "Default"));
    try expect(simple.version == 1);
}
