const std = @import("std");

const Stuff = struct {};

fn do_stuff(
  alloc: std.mem.Allocator, // Let's assume this is an arena allocator so I don't care about freeing.
  stuff: Stuff,
) !void {
  var x = std.ArrayList([]const u8).init(alloc);
  try x.appendSlice(&[_][]const u8{
    "first of something",
    "one more",
  });
  try x.append(stuff.thing);
  try x.append(try std.fmt.allocPrint(alloc, "build some string {s}.\n", .{stuff.athing}));

  var other_stuff = try std.fmt.allocPrint(alloc, "things... {s}", .{stuff.blah});

  try do_other_stuff(x.items, other_stuff);
}