const std = @import("std");

/// Writes a reversed copy of `s` to `buffer`.
pub fn reverse(buffer: []u8, s: []const u8) []u8 {
    std.mem.copyForwards(u8, buffer, s);
    std.mem.reverse(u8, buffer);
    return buffer;
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const testing = std.testing;

fn testReverse(comptime s: []const u8, expected: []const u8) !void {
    var buffer: [s.len]u8 = undefined;
    const actual = reverse(&buffer, s);
    try testing.expectEqualStrings(expected, actual);
}

test "an empty string" {
    try testReverse("", "");
}

test "a word" {
    try testReverse("robot", "tobor");
}

test "a capitalized word" {
    try testReverse("Ramen", "nemaR");
}

test "a sentence with punctuation" {
    try testReverse("I'm hungry!", "!yrgnuh m'I");
}

test "a palindrome" {
    try testReverse("racecar", "racecar");
}

test "an even-sized word" {
    try testReverse("drawer", "reward");
}
