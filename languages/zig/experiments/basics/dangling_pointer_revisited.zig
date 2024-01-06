const std = @import("std");

pub fn main() !void {
	var gpa = std.heap.GeneralPurposeAllocator(.{}){};
	const allocator = gpa.allocator();

	var lookup = std.StringHashMap(User).init(allocator);
	defer lookup.deinit();

	const goku = User{.power = 9001};

	try lookup.put("Goku", goku);

	// returns an optional, .? would panic if "Goku"
	// wasn't in our hashmap
	const entry = lookup.getPtr("Goku").?;

	std.debug.print("Goku's power is: {d}\n", .{entry.power});

	// returns true/false depending on if the item was removed
    // entry points to the copy stored in the hashmap
	_ = lookup.remove("Goku");

    // this now accesses memory that's already been freed.
	std.debug.print("Goku's power is: {d}\n", .{entry.power});

    // the original struct is still fine
    std.debug.print("Goku's power is: {d}\n", .{goku.power});
}

const User = struct {
	power: i32,
};

// ========================================================================================== TESTS
const testing = std.testing;
test "test for dangling pointer" {
	var lookup = std.StringHashMap(User).init(testing.allocator);
	defer lookup.deinit();


	const goku = User{.power = 9001};

	try lookup.put("Goku", goku);

	// returns an optional, .? would panic if "Goku"
	// wasn't in our hashmap
	const entry = lookup.getPtr("Goku").?;

	_ = lookup.remove("Goku");

	// this should fail, because entry has been removed
	// !! not easy to detect though !!
	try testing.expectEqual(@as(i32, 9001), entry.power);
}