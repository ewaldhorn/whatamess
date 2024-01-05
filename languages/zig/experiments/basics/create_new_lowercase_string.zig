const std = @import("std");
const Allocator = std.mem.Allocator;

fn allocLower(allocator: Allocator, str: []const u8) ![]const u8 {
	var dest = try allocator.alloc(u8, str.len);

	for (str, 0..) |c, i| {
		dest[i] = switch (c) {
			'A'...'Z' => c + 32,
			else => c,
		};
	}

	return dest;
}