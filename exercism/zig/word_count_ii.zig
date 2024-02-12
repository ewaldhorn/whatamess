const std = @import("std");
const mem = std.mem;

/// Returns the counts of the words in `s`.
/// Caller owns the returned memory.
pub fn countWords(allocator: mem.Allocator, s: []const u8) !std.StringHashMap(u32) {

    var splitter = mem.splitAny(u8, s, " :!?\t\n,.&@$%^");
    var result = std.StringHashMap(u32).init(allocator);

    while (splitter.next()) |tmp_s| {
        if (tmp_s.len == 0) {
            // nothing to do this iteration, move along
            continue;
        }

        const word_start: usize = if (tmp_s[0] == '\'') 1 else 0;
        const word_end: usize = if (tmp_s[tmp_s.len - 1] == '\'') tmp_s.len - 1 else tmp_s.len;

        if (word_start >= word_end) {
            // nothing to do, empty
            continue;
        }
        const lowercased_string = try std.ascii.allocLowerString(allocator, tmp_s[word_start..word_end]);
        const value = result.get(lowercased_string);
        if (value) |v| {
            try result.put(lowercased_string, v + 1);
            allocator.free(lowercased_string);
        } else {
            // need to make our own copy of the tmp string, or it will get freed before we are done
            // only needs to be done ONCE per unique key, or we'll have a memory leak
            try result.put(lowercased_string, 1);
        }
    }

    return result;
}

// ________________________________________________________________________________________________
// ========================================================================================== TESTS
const test_allocator = std.testing.allocator;

fn freeKeysAndDeinit(self: *std.StringHashMap(u32)) void {
    var iter = self.keyIterator();
    while (iter.next()) |key_ptr| {
        self.allocator.free(key_ptr.*);
    }
    self.deinit();
}

test "count one word" {
    const s = "word";
    var map = try countWords(test_allocator, s);
    defer freeKeysAndDeinit(&map);

    try std.testing.expectEqual(@as(u32, 1), map.count());
    try std.testing.expectEqual(@as(?u32, 1), map.get("word"));
}

test "count one of each word" {
    const s = "one of each";
    var map = try countWords(test_allocator, s);
    defer freeKeysAndDeinit(&map);

    try std.testing.expectEqual(@as(u32, 3), map.count());
    try std.testing.expectEqual(@as(?u32, 1), map.get("one"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("of"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("each"));
}

test "multiple occurrences of a word" {
    const s = "one fish two fish red fish blue fish";
    var map = try countWords(test_allocator, s);
    defer freeKeysAndDeinit(&map);

    try std.testing.expectEqual(@as(u32, 5), map.count());
    try std.testing.expectEqual(@as(?u32, 1), map.get("one"));
    try std.testing.expectEqual(@as(?u32, 4), map.get("fish"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("two"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("red"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("blue"));
}

test "handles cramped lists" {
    const s = "one,two,three";
    var map = try countWords(test_allocator, s);
    defer freeKeysAndDeinit(&map);

    try std.testing.expectEqual(@as(u32, 3), map.count());
    try std.testing.expectEqual(@as(?u32, 1), map.get("one"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("two"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("three"));
}

test "handles expanded lists" {
    const s = "one,\ntwo,\nthree";
    var map = try countWords(test_allocator, s);
    defer freeKeysAndDeinit(&map);

    try std.testing.expectEqual(@as(u32, 3), map.count());
    try std.testing.expectEqual(@as(?u32, 1), map.get("one"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("two"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("three"));
}

test "ignore punctuation" {
    const s = "car: carpet as java: javascript!!&@$%^&";
    var map = try countWords(test_allocator, s);
    defer freeKeysAndDeinit(&map);

    try std.testing.expectEqual(@as(u32, 5), map.count());
    try std.testing.expectEqual(@as(?u32, 1), map.get("car"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("carpet"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("as"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("java"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("javascript"));
}

test "with apostrophes" {
    const s = "'First: don't laugh. Then: don't cry. You're getting it.'";
    var map = try countWords(test_allocator, s);
    defer freeKeysAndDeinit(&map);

    try std.testing.expectEqual(@as(u32, 8), map.count());
    try std.testing.expectEqual(@as(?u32, 1), map.get("first"));
    try std.testing.expectEqual(@as(?u32, 2), map.get("don't"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("laugh"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("then"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("cry"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("you're"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("getting"));
    try std.testing.expectEqual(@as(?u32, 1), map.get("it"));
}
