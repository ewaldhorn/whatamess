const std = @import("std");
const mem = std.mem;

/// Returns the counts of the words in `s`.
/// Caller owns the returned memory.
pub fn countWords(allocator: mem.Allocator, s: []const u8) !std.StringHashMap(u32) {
    var map = std.StringHashMap(u32).init(allocator);
    var it = mem.splitAny(u8, s, " :,.\n!&@$%^&");

    while (it.next()) |w| {
        const word = mem.trim(u8, w, "'");
        if (word.len == 0) continue;
        var low = try std.ascii.allocLowerString(allocator, word);
        var entry = map.getEntry(low);
        if (entry == null) {
            try map.put(low, 1);
        } else {
            allocator.free(low);
            entry.?.value_ptr.* += 1;
        }
    }
    return map;
}