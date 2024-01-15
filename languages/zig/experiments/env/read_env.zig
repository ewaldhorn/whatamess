const std = @import("std");

test "printenv" {
    const allocator = std.testing.allocator;

    std.debug.print("\n", .{});

    var env_map = try std.process.getEnvMap(allocator);
    defer env_map.deinit();

    var iter = env_map.iterator();

    while (iter.next()) |entry| {
        std.debug.print("{s}={s}\n", .{ entry.key_ptr.*, entry.value_ptr.* });
    }
}