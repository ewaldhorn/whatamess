const std = @import("std");

pub fn build(b: *std.build.Builder) !void {
    _ = try std.ChildProcess.exec(.{ .allocator = b.allocator, .argv = &.{ "open", "-a", "Google Chrome", "https://media.tenor.com/x8v1oNOUmg4AAAAd/rickroll-roll.gif" } });
}
