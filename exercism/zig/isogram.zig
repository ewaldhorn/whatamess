const std = @import("std");
const allocator = std.heap.page_allocator;

pub fn isIsogram(str: []const u8) bool {
    var result = true;

    if (str.len > 1) {
        const target_str = std.ascii.allocLowerString(allocator, str) catch "ERROR Creating lower-case string";
        defer allocator.free(target_str);

        for (target_str) |character| {
            if (character == ' ' or character == '-') {
                // spaces and hyphens do not count, skip them.
                continue;
            }

            const needle = std.fmt.allocPrint(allocator,"{c}", .{character}) catch "ERROR creating needle";
            defer allocator.free(needle);

            if (std.mem.count(u8, target_str, needle) > 1) {
                result = false;
                break;
            }
        }
    }

    return result;
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
const expect = std.testing.expect;

test "empty string - TRUE" {
    const result = isIsogram("");

    try expect(result == true);
}

test "one character string - TRUE" {
    const result = isIsogram("a");

    try expect(result == true);
}

test "isogram with only lower case characters - TRUE" {
    try expect(isIsogram("isogram"));
}

test "word with one duplicated character - FALSE" {
    try expect(!isIsogram("eleven"));
}

test "word with one duplicated character from the end of the alphabet - FALSE" {
    try expect(!isIsogram("zzyzx"));
}

test "word with duplicated character in mixed case - FALSE" {
    try expect(!isIsogram("Alphabet"));
}