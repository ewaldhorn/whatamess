const std = @import("std");

pub fn hello(allocator: std.mem.Allocator, str: []const u8) ![]const u8 {
    var list = try std.ArrayList(u8).initCapacity(allocator, str.len);
    defer list.deinit();

    for (str) |c| {
        list.appendAssumeCapacity(c);
    }

    return try list.toOwnedSlice();
}

// ________________________________________________________________________________________________ 
// ========================================================================================== TESTS 
test "can receive a string" {
    const input = "Hello, World!";

    const result = try hello(std.testing.allocator, input);
    defer std.testing.allocator.free(result);

    std.debug.print("\n\nWe got '{s}' {any}\n\n", .{result, @TypeOf(result)});
    try std.testing.expect(std.mem.eql(u8, result, input));
}