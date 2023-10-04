const std = @import("std");
const mem = std.mem;

/// Returns the counts of the words in `s`.
/// Caller owns the returned memory.
pub fn countWords(allocator: mem.Allocator, s: []const u8) !std.StringHashMap(u32) {
    var working_string = try std.ascii.allocLowerString(allocator, s);
    var splitter = mem.split(u8, working_string, " :!?\t\n");
    var result = std.StringHashMap(u32).init(allocator);

    std.debug.print("\n\n<<<Working with: {s}>>>\n\n", .{working_string});

    while (splitter.next()) |tmp| {
        std.debug.print("{s}\n", .{tmp});
    }

    return result;
}
