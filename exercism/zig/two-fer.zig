const std = @import("std");
pub fn twoFer(buffer: []u8, name: ?[]const u8) anyerror![]u8 {
    if (name) |n| {
        const answer = try std.fmt.bufPrint(buffer, "One for {s}, one for me.", .{n});
        return answer;
    } else {
        const answer = try std.fmt.bufPrint(buffer, "One for {s}, one for me.", .{"you"});
        return answer;
    }
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const testing = std.testing;
const buffer_size = 100;

test "no name given" {
    var response: [buffer_size]u8 = undefined;
    const expected = "One for you, one for me.";
    const actual = try twoFer(&response, null);
    try testing.expectEqualStrings(expected, actual);
}

test "a name given" {
    var response: [buffer_size]u8 = undefined;
    const expected = "One for Alice, one for me.";
    const actual = try twoFer(&response, "Alice");
    try testing.expectEqualStrings(expected, actual);
}

test "another name given" {
    var response: [buffer_size]u8 = undefined;
    const expected = "One for Bob, one for me.";
    const actual = try twoFer(&response, "Bob");
    try testing.expectEqualStrings(expected, actual);
}
