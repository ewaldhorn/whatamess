const std = @import("std");
const mem = std.mem;

pub fn abbreviate(allocator: mem.Allocator, words: []const u8) mem.Allocator.Error![]u8 {
    var stack = std.ArrayList(u8).init(allocator);
    defer stack.deinit();

    try stack.append(words[0]);
    var addNext = false;

    for (1..words.len) |idx| {
        if (words[idx] == ' ' or words[idx] == '-') {
            addNext = true;
            continue;
        }

        if (words[idx] == '_') {
            continue;
        }

        if (addNext) {
            try stack.append(std.ascii.toUpper(words[idx]));
            addNext = false;
        }
    }

    return try stack.toOwnedSlice();
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const testing = std.testing;

test "basic" {
    const expected = "PNG";
    const actual = try abbreviate(testing.allocator, "Portable Network Graphics");
    defer testing.allocator.free(actual);
    try testing.expectEqualStrings(expected, actual);
}
test "lowercase words" {
    const expected = "ROR";
    const actual = try abbreviate(testing.allocator, "Ruby on Rails");
    defer testing.allocator.free(actual);
    try testing.expectEqualStrings(expected, actual);
}
test "punctuation" {
    const expected = "FIFO";
    const actual = try abbreviate(testing.allocator, "First In, First Out");
    defer testing.allocator.free(actual);
    try testing.expectEqualStrings(expected, actual);
}
test "all caps word" {
    const expected = "GIMP";
    const actual = try abbreviate(testing.allocator, "GNU Image Manipulation Program");
    defer testing.allocator.free(actual);
    try testing.expectEqualStrings(expected, actual);
}
test "punctuation without whitespace" {
    const expected = "CMOS";
    const actual = try abbreviate(testing.allocator, "Complementary metal-oxide semiconductor");
    defer testing.allocator.free(actual);
    try testing.expectEqualStrings(expected, actual);
}
test "very long abbreviation" {
    const expected = "ROTFLSHTMDCOALM";
    const actual = try abbreviate(testing.allocator, "Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me");
    defer testing.allocator.free(actual);
    try testing.expectEqualStrings(expected, actual);
}
test "consecutive delimiters" {
    const expected = "SIMUFTA";
    const actual = try abbreviate(testing.allocator, "Something - I made up from thin air");
    defer testing.allocator.free(actual);
    try testing.expectEqualStrings(expected, actual);
}
test "apostrophes" {
    const expected = "HC";
    const actual = try abbreviate(testing.allocator, "Halley's Comet");
    defer testing.allocator.free(actual);
    try testing.expectEqualStrings(expected, actual);
}
test "underscore emphasis" {
    const expected = "TRNT";
    const actual = try abbreviate(testing.allocator, "The Road _Not_ Taken");
    defer testing.allocator.free(actual);
    try testing.expectEqualStrings(expected, actual);
}
