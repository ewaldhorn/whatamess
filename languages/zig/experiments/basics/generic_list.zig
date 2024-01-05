const std = @import("std");

fn List(comptime T: type) type {
	return struct {
		pos: usize,
		items: []T,
		allocator: Allocator,

		fn init(allocator: Allocator) !List(T) {
			return .{
				.pos = 0,
				.allocator = allocator,
				.items = try allocator.alloc(T, 4),
			};
		}
	}
};