const std = @import("std");

test "Nested function with anonymous function" {
    const Writer = std.io.getStdOut().writer();
    Writer.print("\n", .{}) catch {};

    makeList(Writer, ". ");
}

fn makeList(Writer: std.fs.File.Writer, separator: []const u8) void {
    const makeItem = struct {
        var counter: usize = 0;

        fn makeItem(Wr: std.fs.File.Writer, sep: []const u8, item: []const u8) void {
            counter += 1;
            Wr.print("{d}{s}{s}\n", .{ counter, sep, item }) catch return;
        }
    }.makeItem;

    makeItem(Writer, separator, "first");
    makeItem(Writer, separator, "second");
    makeItem(Writer, separator, "third");
}