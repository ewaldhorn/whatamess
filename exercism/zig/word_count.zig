const std = @import("std");
const mem = std.mem;

/// Returns the counts of the words in `s`.
/// Caller owns the returned memory.
pub fn countWords(allocator: mem.Allocator, s: []const u8) !std.StringHashMap(u32) {
    const working_string = try std.ascii.allocLowerString(allocator, s);
    defer allocator.free(working_string);

    var splitter = mem.splitAny(u8, working_string, " :!?\t\n,.");
    var result = std.StringHashMap(u32).init(allocator);
    
    // std.debug.print("\n\nWorking with: \"{s}\"\n\n", .{working_string});
    
    while (splitter.next()) |tmp_s| {
        // std.debug.print("Now looking for: {s}\n", .{tmp});
        const value = result.get(tmp_s);
        if (value) |v| {
            try result.put(tmp_s, v + 1);
        } else {
            // need to make our own copy of the tmp string, or it will get freed before we are done
            // only needs to be done ONCE per unique key, or we'll have a memory leak
            const tmp = std.fmt.allocPrint(allocator, "{s}", .{tmp_s}) catch "ERROR";
            try result.put(tmp, 1);
        }
    }

    // var iterator = result.keyIterator();
    // while (iterator.next()) |entry| {
    //     std.debug.print("Entry: {s}\n", .{entry.*});
    // }

    // std.debug.print("\n>>>>>>>>>>>>>>>>>>>\nResult: Count: {d}\n", .{result.count()});
    // const word = result.get("word");
    // if (word) |w| {
    //     std.debug.print("Found {d} occurence(s) of 'word'.\n", .{w});
    // }
    
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