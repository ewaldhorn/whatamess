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

test "optional" {
    var found_index: ?usize = null;
    const data = [_]i32{ 1, 2, 3, 4, 5, 6, 7, 8, 12 };
    for (data, 0..) |v, i| {
        if (v == 10) found_index = i;
    }
    try expect(found_index == null);
}

test "orelse" {
    const a: ?f32 = null;
    const b = a orelse 0;
    try expect(b == 0);
    try expect(@TypeOf(b) == f32);
}

test "orelse unreachable" {
    const a: ?f32 = 5;
    const b = a orelse unreachable;
    const c = a.?;
    try expect(b == c);
    try expect(@TypeOf(c) == f32);
}