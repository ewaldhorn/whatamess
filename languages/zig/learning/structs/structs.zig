const std = @import("std");

const UserRecord = struct {
    username: []const u8,
    userid: u32,
    user_pin: []const u8,

    pub fn init(name: []const u8) UserRecord {
        return UserRecord{.username = name, .user_pin = "1234", .userid = 9999};
    }
};

const RecordWithDefaults = struct {
    username: []const u8 = "Not Supplied",
    userid: u32,
    user_pin: []const u8 = "0000",

    pub fn report(self: RecordWithDefaults) void {
        std.debug.print("{s} has an ID of {d}.\n", .{self.username, self.userid});
    }
};

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const expect = std.testing.expect;

test "can create a struct directly" {
    const user = UserRecord {.username = "Calm Pete", .userid = 1, .user_pin = "1121"};

    try expect(user.userid == 1);
    try expect(std.mem.eql(u8, user.user_pin, "1121"));
    try expect(std.mem.eql(u8, user.username, "Calm Pete"));
}

test "can create a struct with an init function" {
    const user = UserRecord.init("Bob");

    try expect(user.userid == 9999);
    try expect(std.mem.eql(u8, user.user_pin, "1234"));
    try expect(std.mem.eql(u8, user.username, "Bob"));
}

test "can create a struct with defaults" {
    const user = RecordWithDefaults {.userid = 1};

    try expect(user.userid == 1);
    try expect(std.mem.eql(u8, user.user_pin, "0000"));
    try expect(std.mem.eql(u8, user.username, "Not Supplied"));    
}

test "can create a struct and override defaults" {
    const user = RecordWithDefaults {.userid = 2, .user_pin = "1121"};

    try expect(user.userid == 2);
    try expect(std.mem.eql(u8, user.user_pin, "1121"));
    try expect(std.mem.eql(u8, user.username, "Not Supplied"));    
}