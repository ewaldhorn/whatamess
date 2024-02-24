const std = @import("std");

// Check that two strings are valid anagrams.  Anagrams use the same letters in the same frequency.
// For example:  Loop, Pool
fn validateAnagram(allocator: std.mem.Allocator, left: []const u8, right: []const u8) bool {
    if (left.len != right.len) {
        return false;
    }

    var left_hash = std.AutoHashMap(u8, u8).init(allocator);
    defer left_hash.deinit();

    var right_hash = std.AutoHashMap(u8, u8).init(allocator);
    defer right_hash.deinit();

    for (0..left.len) |idx| {
        const lookup_left = std.ascii.toLower(left[idx]);
        const lookup_right = std.ascii.toLower(right[idx]);

        const add_left = left_hash.getOrPut(lookup_left) catch { return false;};
        if (add_left.found_existing) {
            add_left.value_ptr.* += 1;
        } else {
            add_left.value_ptr.* = 1;
        }

        const add_right = right_hash.getOrPut(lookup_right) catch { return false;};
        if (add_right.found_existing) {
            add_right.value_ptr.* += 1;
        } else {
            add_right.value_ptr.* = 1;
        }
    }

    for (0..left.len) |idx| {
        const lookup = std.ascii.toLower(left[idx]);
        if (left_hash.get(lookup) != right_hash.get(lookup)) {
            return false;
        }
    }

    return true;
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const expect = std.testing.expect;
const test_allocator = std.testing.allocator;

test "valid anagram: Pool/Loop" {
    try expect(validateAnagram(test_allocator, "Pool", "Loop"));
}

test "valid anagram: Danger/Garden" {
    try expect(validateAnagram(test_allocator, "Danger", "Garden"));
}

test "valid anagram: Nameless/Salesmen" {
    try expect(validateAnagram(test_allocator, "Nameless", "Salesmen"));
}

test "invalid anagram: Danger/Dangers" {
    try expect(!validateAnagram(test_allocator, "Danger", "Dangers"));
}

test "invalid anagram: Broke/Croak" {
    try expect(!validateAnagram(test_allocator, "Broke", "Croak"));
}
