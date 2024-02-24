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

        const leftValue = left_hash.get(lookup_left);
        if (leftValue) |c| {
            left_hash.put(lookup_left, c + 1) catch {
                std.debug.panic("Error incrementing value in hashmap", .{});
            };
        } else {
            left_hash.put(lookup_left, 1) catch {
                std.debug.panic("Error adding value to hashmap", .{});
            };
        }

        const rightValue = right_hash.get(lookup_right);
        if (rightValue) |c| {
            right_hash.put(lookup_right, c + 1) catch {
                std.debug.panic("Error incrementing value in hashmap", .{});
            };
        } else {
            right_hash.put(lookup_right, 1) catch {
                std.debug.panic("Error adding value to hashmap", .{});
            };
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
