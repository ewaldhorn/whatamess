const std = @import("std");
const expect = std.testing.expect;

// split a string on spaces
// var s = "hello world this is a test";
// var splits = std.mem.split(u8, s, " ");
// while (splits.next()) |chunk| {
//   std.debug.print("{s}\n", .{chunk});
// }

test "different ways to init strings" {
    const literal = "hello"; // *const [5:0]u8
    const array = [_]u8{'h', 'e', 'l', 'l', 'o'}; // [5]u8

    try expect(std.mem.eql(u8, literal, &array) == true);
    try expect(array.len == 5);
}

test "combining strings" {
    const left = "hello";
    const space = " ";
    const right = "world";
    const answer = left ++ space ++ right;

    try expect(std.mem.eql(u8, answer, "hello world"));
    try expect(std.mem.eql(u8, answer[0..5], "hello"));
}

test "finding stuff in strings" {
    const haystack = "this string contains a few words";

    try expect(std.mem.indexOf(u8, haystack, "this") == 0);
    try expect(std.mem.indexOf(u8, haystack, "string") == 5);
    try expect(std.mem.lastIndexOf(u8, haystack, "s") == haystack.len - 1);

    // when not found, we get a null back
    try expect(std.mem.indexOf(u8, haystack, "lol") == null);
}

test "modifying strings" {
    // note the .* dereference to create a mutable array
    var s = "dood dood".*;
    s[0] = 'm';
    s[5] = 'm';

    try expect(std.mem.eql(u8, &s, "mood mood"));

    std.mem.reverse(u8, &s);
    try expect(std.mem.eql(u8, &s, "doom doom"));
}

test "dynamic formatting of strings" {
    const allocator = std.testing.allocator;
    const count : i32 = 1001;
    const output = try std.fmt.allocPrint(allocator, "You need {d} excuses!", .{count});
    defer allocator.free(output);
    
    try expect(std.mem.eql(u8, "You need 1001 excuses!", output));
}