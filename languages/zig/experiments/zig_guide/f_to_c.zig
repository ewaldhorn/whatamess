// Based on https://zig.guide/posts/fahrenheit-to-celsius
const std = @import("std");
const expect = std.testing.expect;

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    // get the program arguments, allocating space for them on the heap
    // remember to schedule a memory deallocation call with defer
    const args = try std.process.argsAlloc(std.heap.page_allocator);
    defer std.process.argsFree(std.heap.page_allocator, args);

    if (args.len < 2) {
        return error.ExpectedArgument;
    }

    const argument = args[1];
    const value = try std.fmt.parseFloat(f32, argument);
    const c = (value - 32) * (5.0 / 9.0);

    try stdout.print("Converting {d:.2} \u{2109} gives {d:.2} \u{2103}.\n", .{ value, c });

    // list all arguments
    // for (args, 0..) |arg, i| {
    //     try stdout.print("arg {}: {s}\n", .{ i, arg });
    // }
}
