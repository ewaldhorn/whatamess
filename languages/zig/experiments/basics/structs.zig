const stdlib = @import("std");
const expect = @import("std").testing.expect;

pub fn main() void {
    const user = User {
        .power = 123,
        .name = "Roger"
    };

    stdlib.debug.print("{s}'s power is {d}.\n", .{user.name, user.power});
}

pub const User = struct {
    power: u64,
    name: []const u8,
    version: u32 = 1
};


// ========================================================================================== TESTS
test "simple struct with default version" {
    const simple = User {
        .power = 555,
        .name = "Bob"
    };

    try expect(simple.power == 555);
    try expect(stdlib.mem.eql(u8, simple.name, "Bob"));
    try expect(simple.version == 1);
}

test "simple struct with custom version" {
    const simple = User {
        .power = 555,
        .name = "Bob",
        .version = 2
    };

    try expect(simple.power == 555);
    try expect(stdlib.mem.eql(u8, simple.name, "Bob"));
    try expect(simple.version == 2);
}