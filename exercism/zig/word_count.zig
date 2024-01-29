const std = @import("std");
const mem = std.mem;

/// Returns the counts of the words in `s`.
/// Caller owns the returned memory.
pub fn countWords(allocator: mem.Allocator, s: []const u8) !std.StringHashMap(u32) {
    const working_string = try std.ascii.allocLowerString(allocator, s);
    errdefer allocator.free(working_string);

    var splitter = mem.splitAny(u8, working_string, " :!?\t\n,.");
    var result = std.StringHashMap(u32).init(allocator);
    
    // std.debug.print("\n\nWorking with: \"{s}\"\n\n", .{working_string});
    
    while (splitter.next()) |tmp| {
        // std.debug.print("Now looking for: {s}\n", .{tmp});
        const value = result.get(tmp);
        if (value) |v| {
            try result.put(tmp, v + 1);
        } else {
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