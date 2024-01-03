const std = @import("std");
const expect = std.testing.expect;

test "something optional is not set" {
    const home: ?[]const u8 = null;
    const result = home orelse "not set";

    try expect(std.mem.eql(u8, "not set", result));
}

test "something optional is set" {
    const home: ?[]const u8 = "is set";
    const result = home orelse "not set";

    try expect(std.mem.eql(u8, "is set", result));
}

test "undefined is also available" {
    var pseudo_uuid: [16]u8 = undefined;
    std.crypto.random.bytes(&pseudo_uuid);

    try expect(pseudo_uuid[0] != undefined);
}