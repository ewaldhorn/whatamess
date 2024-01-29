const std = @import("std");

const UserRecord = struct {
    username: []const u8,
    userid: u32,
    user_pin: []const u8
};

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const expect = std.testing.expect;

test "can create a struct" {
    const user = UserRecord {.username = "Calm Pete", .userid = 1, .user_pin = "1121"};

    try expect(user.userid == 1);
    try expect(std.mem.eql(u8, user.user_pin, "1121"));
    try expect(std.mem.eql(u8, user.username, "Calm Pete"));
}