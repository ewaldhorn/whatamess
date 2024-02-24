const std = @import("std");

// Check that two strings are valid anagrams.  Anagrams use the same letters in the same frequency.
// Input will always be in lower-case
// For example:  Loop, Pool
fn validateAnagram(left: []const u8, right: []const u8) bool {
    if (left.len != right.len) {
        return false;
    }

    // create buffers
    var left_buffer : [1000]u8 = undefined;
    var right_buffer : [1000]u8 = undefined;

    // create slices so we can sort them
    const left_slice : []u8 = left_buffer[0..left.len];
    const right_slice : []u8 = right_buffer[0..right.len];

    // copy contents into mutable slices
    std.mem.copyForwards(u8, left_slice, left);
    std.mem.copyForwards(u8, right_slice, right);

    // sort characters
    std.mem.sort(u8, left_slice, {}, comptime std.sort.asc(u8));
    std.mem.sort(u8, right_slice, {}, comptime std.sort.asc(u8));

    // ensure they are an exact match
    return std.mem.eql(u8, left_slice, right_slice);
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const expect = std.testing.expect;
const test_allocator = std.testing.allocator;

test "valid anagram: pool/loop" {
    try expect(validateAnagram("pool", "loop"));
}

test "valid anagram: danger/garden" {
    try expect(validateAnagram("danger", "garden"));
}

test "valid anagram: nameless/salesmen" {
    try expect(validateAnagram("nameless", "salesmen"));
}

test "invalid anagram: danger/dangers" {
    try expect(!validateAnagram("danger", "dangers"));
}

test "invalid anagram: broke/croak" {
    try expect(!validateAnagram("broke", "croak"));
}
