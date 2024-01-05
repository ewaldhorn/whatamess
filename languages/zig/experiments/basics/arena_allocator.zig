const std = @import("std");

fn parse(allocator: Allocator, input: []const u8) !Something {
	// create an ArenaAllocator from the supplied allocator
    // note the use of init instead of alloc
	var arena = std.heap.ArenaAllocator.init(allocator);

	// this will free anything created from this arena
    // one call to free all allocations, all at once.
    // note the use of deinit instead of free
	defer arena.deinit();

	// create an std.mem.Allocator from the arena, this will be
	// the allocator we'll use internally
    // use this as normal to allocate memory, but don't bother freeing
    // we free the entire arena when we are done
	const aa = arena.allocator();

	const state = State{
		// we're using aa here!
		.buf = try aa.alloc(u8, 512),

		// we're using aa here!
		.nesting = try aa.alloc(NestType, 10),
	};

	// we're passing aa here, so we're guaranteed that
	// any other allocation will be in our arena
	return parseInternal(aa, state, input);
}